// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetVpcPublicPatRuleResult {
    private String createdAt;
    /**
     * @return The ID of the public gateway.
     * 
     */
    private String gatewayId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String organizationId;
    private String patRuleId;
    /**
     * @return The Private IP to forward data to (IP address).
     * 
     */
    private String privateIp;
    /**
     * @return The Private port to translate to.
     * 
     */
    private Integer privatePort;
    /**
     * @return The Protocol the rule should apply to. Possible values are both, tcp and udp.
     * 
     */
    private String protocol;
    /**
     * @return The Public port to listen on.
     * 
     */
    private Integer publicPort;
    private String updatedAt;
    private @Nullable String zone;

    private GetVpcPublicPatRuleResult() {}
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return The ID of the public gateway.
     * 
     */
    public String gatewayId() {
        return this.gatewayId;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String organizationId() {
        return this.organizationId;
    }
    public String patRuleId() {
        return this.patRuleId;
    }
    /**
     * @return The Private IP to forward data to (IP address).
     * 
     */
    public String privateIp() {
        return this.privateIp;
    }
    /**
     * @return The Private port to translate to.
     * 
     */
    public Integer privatePort() {
        return this.privatePort;
    }
    /**
     * @return The Protocol the rule should apply to. Possible values are both, tcp and udp.
     * 
     */
    public String protocol() {
        return this.protocol;
    }
    /**
     * @return The Public port to listen on.
     * 
     */
    public Integer publicPort() {
        return this.publicPort;
    }
    public String updatedAt() {
        return this.updatedAt;
    }
    public Optional<String> zone() {
        return Optional.ofNullable(this.zone);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetVpcPublicPatRuleResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createdAt;
        private String gatewayId;
        private String id;
        private String organizationId;
        private String patRuleId;
        private String privateIp;
        private Integer privatePort;
        private String protocol;
        private Integer publicPort;
        private String updatedAt;
        private @Nullable String zone;
        public Builder() {}
        public Builder(GetVpcPublicPatRuleResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createdAt = defaults.createdAt;
    	      this.gatewayId = defaults.gatewayId;
    	      this.id = defaults.id;
    	      this.organizationId = defaults.organizationId;
    	      this.patRuleId = defaults.patRuleId;
    	      this.privateIp = defaults.privateIp;
    	      this.privatePort = defaults.privatePort;
    	      this.protocol = defaults.protocol;
    	      this.publicPort = defaults.publicPort;
    	      this.updatedAt = defaults.updatedAt;
    	      this.zone = defaults.zone;
        }

        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            this.createdAt = Objects.requireNonNull(createdAt);
            return this;
        }
        @CustomType.Setter
        public Builder gatewayId(String gatewayId) {
            this.gatewayId = Objects.requireNonNull(gatewayId);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder organizationId(String organizationId) {
            this.organizationId = Objects.requireNonNull(organizationId);
            return this;
        }
        @CustomType.Setter
        public Builder patRuleId(String patRuleId) {
            this.patRuleId = Objects.requireNonNull(patRuleId);
            return this;
        }
        @CustomType.Setter
        public Builder privateIp(String privateIp) {
            this.privateIp = Objects.requireNonNull(privateIp);
            return this;
        }
        @CustomType.Setter
        public Builder privatePort(Integer privatePort) {
            this.privatePort = Objects.requireNonNull(privatePort);
            return this;
        }
        @CustomType.Setter
        public Builder protocol(String protocol) {
            this.protocol = Objects.requireNonNull(protocol);
            return this;
        }
        @CustomType.Setter
        public Builder publicPort(Integer publicPort) {
            this.publicPort = Objects.requireNonNull(publicPort);
            return this;
        }
        @CustomType.Setter
        public Builder updatedAt(String updatedAt) {
            this.updatedAt = Objects.requireNonNull(updatedAt);
            return this;
        }
        @CustomType.Setter
        public Builder zone(@Nullable String zone) {
            this.zone = zone;
            return this;
        }
        public GetVpcPublicPatRuleResult build() {
            final var o = new GetVpcPublicPatRuleResult();
            o.createdAt = createdAt;
            o.gatewayId = gatewayId;
            o.id = id;
            o.organizationId = organizationId;
            o.patRuleId = patRuleId;
            o.privateIp = privateIp;
            o.privatePort = privatePort;
            o.protocol = protocol;
            o.publicPort = publicPort;
            o.updatedAt = updatedAt;
            o.zone = zone;
            return o;
        }
    }
}