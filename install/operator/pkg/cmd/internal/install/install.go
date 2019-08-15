/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package install

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/syndesisio/syndesis/install/operator/pkg"
	"github.com/syndesisio/syndesis/install/operator/pkg/cmd/internal"
	"github.com/syndesisio/syndesis/install/operator/pkg/generator"
	"github.com/syndesisio/syndesis/install/operator/pkg/util"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/yaml"
)

type Install struct {
	// cli parsed config
	*internal.Options
	wait      bool
	eject     string
	image     string
	devImages bool

	// processing state
	ejectedResources []unstructured.Unstructured
}

func New(parent *internal.Options) *cobra.Command {

	o := Install{Options: parent}

	cmd := cobra.Command{
		Use:               "install",
		Short:             "installs the all resources",
		PersistentPreRunE: o.before,
		PersistentPostRun: o.after,
		Run: func(cmd *cobra.Command, args []string) {

			err := o.installClusterResources()
			util.ExitOnError(err)
			err = o.installOperatorResources()
			util.ExitOnError(err)
			err = o.installApplication()
			util.ExitOnError(err)
		},
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "cluster",
		Short: "install the cluster level resources (requires cluster admin privileges)",
		Run: func(cmd *cobra.Command, args []string) {
			err := o.installClusterResources()
			util.ExitOnError(err)
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "operator",
		Short: "install the operator resources (requires namespace admin privileges)",
		Run: func(cmd *cobra.Command, args []string) {
			err := o.installOperatorResources()
			util.ExitOnError(err)
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "app",
		Short: "install the syndesis application (requires namespace admin privileges)",
		Run: func(cmd *cobra.Command, args []string) {
			err := o.installApplication()
			util.ExitOnError(err)
		},
	})

	cmd.PersistentFlags().StringVarP(&o.eject, "eject", "e", "", "eject configuration that would be applied to the cluster in the specified format instead of installing the configuration. One of: json|yaml")
	cmd.PersistentFlags().StringVarP(&o.image, "image", "", pkg.DefaultOperatorImage, "sets operator image that gets installed")
	cmd.PersistentFlags().BoolVarP(&o.wait, "wait", "w", false, "waits for the application to be running")
	cmd.PersistentFlags().BoolVarP(&o.devImages, "dev", "", false, "enable development mode by loading images from image stream tags.")

	return &cmd
}

func (o *Install) before(_ *cobra.Command, args []string) (err error) {
	switch o.eject {
	case "":
	case "yaml":
	case "json":
	default:
		return fmt.Errorf("invalid output format: %s", o.eject)
	}

	if len(args) > 0 {
		return fmt.Errorf("unexpected argument: %s", args[0])
	}

	if o.eject != "" {
		o.ejectedResources = []unstructured.Unstructured{}
	}

	return
}

func (o *Install) after(cmd *cobra.Command, args []string) {
	if o.ejectedResources == nil {
		return
	}
	value := util.UnstructuredsToRuntimeObject(o.ejectedResources)
	if value == nil {
		return
	}

	switch o.eject {
	case "yaml":
		data, err := yaml.Marshal(value)
		util.ExitOnError(err)
		fmt.Print(string(data))
	case "json":
		data, err := json.Marshal(value)
		util.ExitOnError(err)
		fmt.Print(string(data))
	default:
		panic(o.eject)
	}
}

func (o *Install) Println(a ...interface{}) (int, error) {
	if o.ejectedResources != nil {
		return 0, nil
	}
	return fmt.Println(a...)
}

type RenderScope struct {
    Image     string
    Namespace string
    DevImages bool
    Role      string
}

func (o *Install) install(action string, resources []unstructured.Unstructured) error {
	updateCounter := 0
	createCounter := 0
	client, err := o.NewClient()
	if err != nil {
		return err
	}
	for _, res := range resources {
		if o.ejectedResources != nil {
			o.ejectedResources = append(o.ejectedResources, res)
		} else {
			_, result, err := util.CreateOrUpdate(o.Context, client, &res)
			if err != nil {
				return errors.Wrap(err, util.Dump(res))
			}

			switch result {
			case controllerutil.OperationResultUpdated:
				createCounter += 1
			case controllerutil.OperationResultCreated:
				createCounter += 1
			}
		}
	}
	if createCounter == 0 && updateCounter == 0 {
		o.Println(action + " previously installed")
	} else if updateCounter != 0 {
		o.Println(action + " updated successfully")
	} else {
		o.Println(action + " installed successfully")
	}
	return nil
}

func (o *Install) render(fromFile string) ([]unstructured.Unstructured, error) {
    resources, err := generator.Render(fromFile, RenderScope{
        Namespace: o.Namespace,
        Image:     o.image,
        DevImages: o.devImages,
        Role:      "syndesis-operator",
    })
    return resources, err
}
