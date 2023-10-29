// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetMnqSqsResult {
    /**
     * @return The endpoint of the SQS service for this project.
     * 
     */
    private String endpoint;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String projectId;
    private @Nullable String region;

    private GetMnqSqsResult() {}
    /**
     * @return The endpoint of the SQS service for this project.
     * 
     */
    public String endpoint() {
        return this.endpoint;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<String> projectId() {
        return Optional.ofNullable(this.projectId);
    }
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetMnqSqsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String endpoint;
        private String id;
        private @Nullable String projectId;
        private @Nullable String region;
        public Builder() {}
        public Builder(GetMnqSqsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.endpoint = defaults.endpoint;
    	      this.id = defaults.id;
    	      this.projectId = defaults.projectId;
    	      this.region = defaults.region;
        }

        @CustomType.Setter
        public Builder endpoint(String endpoint) {
            this.endpoint = Objects.requireNonNull(endpoint);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder projectId(@Nullable String projectId) {
            this.projectId = projectId;
            return this;
        }
        @CustomType.Setter
        public Builder region(@Nullable String region) {
            this.region = region;
            return this;
        }
        public GetMnqSqsResult build() {
            final var o = new GetMnqSqsResult();
            o.endpoint = endpoint;
            o.id = id;
            o.projectId = projectId;
            o.region = region;
            return o;
        }
    }
}