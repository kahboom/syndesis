package io.syndesis.connector.jms.springboot;

import javax.annotation.Generated;

/**
 * Respond to requests using JMS Messages
 * 
 * Generated by camel-package-maven-plugin - do not edit this file!
 */
@Generated("org.apache.camel.maven.connector.SpringBootAutoConfigurationMojo")
public class ActiveMQRespondConnectorConfigurationCommon {

    /**
     * The maximum number of connections available to endpoints started under
     * this component
     */
    private Integer connectionCount = 1;
    /**
     * The kind of destination to use
     */
    private String destinationType = "queue";
    /**
     * DestinationName is a JMS queue or topic name. By default the
     * destinationName is interpreted as a queue name.
     */
    private String destinationName;
    /**
     * Sets the durable subscription Id required for durable topics.
     */
    private String durableSubscriptionId;
    /**
     * Sets the JMS Message selector syntax.
     */
    private String messageSelector;
    /**
     * Sets the reply to destination name used for InOut producer endpoints.
     */
    private String namedReplyTo;
    /**
     * Flag used to enable/disable message persistence.
     */
    private boolean persistent = true;
    /**
     * Broker URL
     */
    private String brokerUrl;
    /**
     * Authorization credential user name
     */
    private String username;
    /**
     * Authorization credential password
     */
    private String password;
    /**
     * Client ID for durable subscriptions
     */
    private String clientID;
    /**
     * Skip Certificate check for development environment
     */
    private Boolean skipCertificateCheck;
    /**
     * AMQ Broker X.509 PEM Certificate
     */
    private String brokerCertificate;
    /**
     * AMQ Client X.509 PEM Certificate
     */
    private String clientCertificate;

    public Integer getConnectionCount() {
        return connectionCount;
    }

    public void setConnectionCount(Integer connectionCount) {
        this.connectionCount = connectionCount;
    }

    public String getDestinationType() {
        return destinationType;
    }

    public void setDestinationType(String destinationType) {
        this.destinationType = destinationType;
    }

    public String getDestinationName() {
        return destinationName;
    }

    public void setDestinationName(String destinationName) {
        this.destinationName = destinationName;
    }

    public String getDurableSubscriptionId() {
        return durableSubscriptionId;
    }

    public void setDurableSubscriptionId(String durableSubscriptionId) {
        this.durableSubscriptionId = durableSubscriptionId;
    }

    public String getMessageSelector() {
        return messageSelector;
    }

    public void setMessageSelector(String messageSelector) {
        this.messageSelector = messageSelector;
    }

    public String getNamedReplyTo() {
        return namedReplyTo;
    }

    public void setNamedReplyTo(String namedReplyTo) {
        this.namedReplyTo = namedReplyTo;
    }

    public boolean isPersistent() {
        return persistent;
    }

    public void setPersistent(boolean persistent) {
        this.persistent = persistent;
    }

    public String getBrokerUrl() {
        return brokerUrl;
    }

    public void setBrokerUrl(String brokerUrl) {
        this.brokerUrl = brokerUrl;
    }

    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    public String getClientID() {
        return clientID;
    }

    public void setClientID(String clientID) {
        this.clientID = clientID;
    }

    public Boolean getSkipCertificateCheck() {
        return skipCertificateCheck;
    }

    public void setSkipCertificateCheck(Boolean skipCertificateCheck) {
        this.skipCertificateCheck = skipCertificateCheck;
    }

    public String getBrokerCertificate() {
        return brokerCertificate;
    }

    public void setBrokerCertificate(String brokerCertificate) {
        this.brokerCertificate = brokerCertificate;
    }

    public String getClientCertificate() {
        return clientCertificate;
    }

    public void setClientCertificate(String clientCertificate) {
        this.clientCertificate = clientCertificate;
    }
}