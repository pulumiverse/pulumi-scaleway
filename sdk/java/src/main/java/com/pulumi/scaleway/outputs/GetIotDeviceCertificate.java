// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetIotDeviceCertificate {
    private String crt;
    private String key;

    private GetIotDeviceCertificate() {}
    public String crt() {
        return this.crt;
    }
    public String key() {
        return this.key;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIotDeviceCertificate defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String crt;
        private String key;
        public Builder() {}
        public Builder(GetIotDeviceCertificate defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.crt = defaults.crt;
    	      this.key = defaults.key;
        }

        @CustomType.Setter
        public Builder crt(String crt) {
            this.crt = Objects.requireNonNull(crt);
            return this;
        }
        @CustomType.Setter
        public Builder key(String key) {
            this.key = Objects.requireNonNull(key);
            return this;
        }
        public GetIotDeviceCertificate build() {
            final var o = new GetIotDeviceCertificate();
            o.crt = crt;
            o.key = key;
            return o;
        }
    }
}