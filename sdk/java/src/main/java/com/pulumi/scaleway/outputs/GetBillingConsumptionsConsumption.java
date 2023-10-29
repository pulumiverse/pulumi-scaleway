// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetBillingConsumptionsConsumption {
    private String category;
    private String description;
    private String operationPath;
    private String projectId;
    private String value;

    private GetBillingConsumptionsConsumption() {}
    public String category() {
        return this.category;
    }
    public String description() {
        return this.description;
    }
    public String operationPath() {
        return this.operationPath;
    }
    public String projectId() {
        return this.projectId;
    }
    public String value() {
        return this.value;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBillingConsumptionsConsumption defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String category;
        private String description;
        private String operationPath;
        private String projectId;
        private String value;
        public Builder() {}
        public Builder(GetBillingConsumptionsConsumption defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.category = defaults.category;
    	      this.description = defaults.description;
    	      this.operationPath = defaults.operationPath;
    	      this.projectId = defaults.projectId;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder category(String category) {
            this.category = Objects.requireNonNull(category);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder operationPath(String operationPath) {
            this.operationPath = Objects.requireNonNull(operationPath);
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        @CustomType.Setter
        public Builder value(String value) {
            this.value = Objects.requireNonNull(value);
            return this;
        }
        public GetBillingConsumptionsConsumption build() {
            final var o = new GetBillingConsumptionsConsumption();
            o.category = category;
            o.description = description;
            o.operationPath = operationPath;
            o.projectId = projectId;
            o.value = value;
            return o;
        }
    }
}