// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetBaremetalOfferCpu {
    private Integer coreCount;
    /**
     * @return Frequency of the memory in MHz.
     * 
     */
    private Integer frequency;
    /**
     * @return The offer name. Only one of `name` and `offer_id` should be specified.
     * 
     */
    private String name;
    private Integer threadCount;

    private GetBaremetalOfferCpu() {}
    public Integer coreCount() {
        return this.coreCount;
    }
    /**
     * @return Frequency of the memory in MHz.
     * 
     */
    public Integer frequency() {
        return this.frequency;
    }
    /**
     * @return The offer name. Only one of `name` and `offer_id` should be specified.
     * 
     */
    public String name() {
        return this.name;
    }
    public Integer threadCount() {
        return this.threadCount;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBaremetalOfferCpu defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer coreCount;
        private Integer frequency;
        private String name;
        private Integer threadCount;
        public Builder() {}
        public Builder(GetBaremetalOfferCpu defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.coreCount = defaults.coreCount;
    	      this.frequency = defaults.frequency;
    	      this.name = defaults.name;
    	      this.threadCount = defaults.threadCount;
        }

        @CustomType.Setter
        public Builder coreCount(Integer coreCount) {
            this.coreCount = Objects.requireNonNull(coreCount);
            return this;
        }
        @CustomType.Setter
        public Builder frequency(Integer frequency) {
            this.frequency = Objects.requireNonNull(frequency);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder threadCount(Integer threadCount) {
            this.threadCount = Objects.requireNonNull(threadCount);
            return this;
        }
        public GetBaremetalOfferCpu build() {
            final var o = new GetBaremetalOfferCpu();
            o.coreCount = coreCount;
            o.frequency = frequency;
            o.name = name;
            o.threadCount = threadCount;
            return o;
        }
    }
}