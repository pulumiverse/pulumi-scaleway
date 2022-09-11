// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetFlexibleIpResult {
    private String createdAt;
    private String description;
    private @Nullable String flexibleIpId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String ipAddress;
    private String macAddress;
    /**
     * @return (Defaults to provider `organization_id`) The ID of the organization the IP is in.
     * 
     */
    private String organizationId;
    /**
     * @return (Defaults to provider `project_id`) The ID of the project the IP is in.
     * 
     */
    private String projectId;
    /**
     * @return The reverse domain associated with this IP.
     * 
     */
    private String reverse;
    /**
     * @return The associated server ID if any
     * 
     */
    private String serverId;
    private List<String> tags;
    private String updatedAt;
    private String zone;

    private GetFlexibleIpResult() {}
    public String createdAt() {
        return this.createdAt;
    }
    public String description() {
        return this.description;
    }
    public Optional<String> flexibleIpId() {
        return Optional.ofNullable(this.flexibleIpId);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<String> ipAddress() {
        return Optional.ofNullable(this.ipAddress);
    }
    public String macAddress() {
        return this.macAddress;
    }
    /**
     * @return (Defaults to provider `organization_id`) The ID of the organization the IP is in.
     * 
     */
    public String organizationId() {
        return this.organizationId;
    }
    /**
     * @return (Defaults to provider `project_id`) The ID of the project the IP is in.
     * 
     */
    public String projectId() {
        return this.projectId;
    }
    /**
     * @return The reverse domain associated with this IP.
     * 
     */
    public String reverse() {
        return this.reverse;
    }
    /**
     * @return The associated server ID if any
     * 
     */
    public String serverId() {
        return this.serverId;
    }
    public List<String> tags() {
        return this.tags;
    }
    public String updatedAt() {
        return this.updatedAt;
    }
    public String zone() {
        return this.zone;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetFlexibleIpResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createdAt;
        private String description;
        private @Nullable String flexibleIpId;
        private String id;
        private @Nullable String ipAddress;
        private String macAddress;
        private String organizationId;
        private String projectId;
        private String reverse;
        private String serverId;
        private List<String> tags;
        private String updatedAt;
        private String zone;
        public Builder() {}
        public Builder(GetFlexibleIpResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createdAt = defaults.createdAt;
    	      this.description = defaults.description;
    	      this.flexibleIpId = defaults.flexibleIpId;
    	      this.id = defaults.id;
    	      this.ipAddress = defaults.ipAddress;
    	      this.macAddress = defaults.macAddress;
    	      this.organizationId = defaults.organizationId;
    	      this.projectId = defaults.projectId;
    	      this.reverse = defaults.reverse;
    	      this.serverId = defaults.serverId;
    	      this.tags = defaults.tags;
    	      this.updatedAt = defaults.updatedAt;
    	      this.zone = defaults.zone;
        }

        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            this.createdAt = Objects.requireNonNull(createdAt);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder flexibleIpId(@Nullable String flexibleIpId) {
            this.flexibleIpId = flexibleIpId;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder ipAddress(@Nullable String ipAddress) {
            this.ipAddress = ipAddress;
            return this;
        }
        @CustomType.Setter
        public Builder macAddress(String macAddress) {
            this.macAddress = Objects.requireNonNull(macAddress);
            return this;
        }
        @CustomType.Setter
        public Builder organizationId(String organizationId) {
            this.organizationId = Objects.requireNonNull(organizationId);
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        @CustomType.Setter
        public Builder reverse(String reverse) {
            this.reverse = Objects.requireNonNull(reverse);
            return this;
        }
        @CustomType.Setter
        public Builder serverId(String serverId) {
            this.serverId = Objects.requireNonNull(serverId);
            return this;
        }
        @CustomType.Setter
        public Builder tags(List<String> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }
        @CustomType.Setter
        public Builder updatedAt(String updatedAt) {
            this.updatedAt = Objects.requireNonNull(updatedAt);
            return this;
        }
        @CustomType.Setter
        public Builder zone(String zone) {
            this.zone = Objects.requireNonNull(zone);
            return this;
        }
        public GetFlexibleIpResult build() {
            final var o = new GetFlexibleIpResult();
            o.createdAt = createdAt;
            o.description = description;
            o.flexibleIpId = flexibleIpId;
            o.id = id;
            o.ipAddress = ipAddress;
            o.macAddress = macAddress;
            o.organizationId = organizationId;
            o.projectId = projectId;
            o.reverse = reverse;
            o.serverId = serverId;
            o.tags = tags;
            o.updatedAt = updatedAt;
            o.zone = zone;
            return o;
        }
    }
}