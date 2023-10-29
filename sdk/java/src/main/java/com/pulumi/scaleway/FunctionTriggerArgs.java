// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.scaleway.inputs.FunctionTriggerNatsArgs;
import com.pulumi.scaleway.inputs.FunctionTriggerSqsArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FunctionTriggerArgs extends com.pulumi.resources.ResourceArgs {

    public static final FunctionTriggerArgs Empty = new FunctionTriggerArgs();

    /**
     * The description of the trigger.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the trigger.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The ID of the function to create a trigger for
     * 
     */
    @Import(name="functionId", required=true)
    private Output<String> functionId;

    /**
     * @return The ID of the function to create a trigger for
     * 
     */
    public Output<String> functionId() {
        return this.functionId;
    }

    /**
     * The unique name of the trigger. Default to a generated name.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The unique name of the trigger. Default to a generated name.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The configuration for the Scaleway&#39;s Nats used by the trigger
     * 
     */
    @Import(name="nats")
    private @Nullable Output<FunctionTriggerNatsArgs> nats;

    /**
     * @return The configuration for the Scaleway&#39;s Nats used by the trigger
     * 
     */
    public Optional<Output<FunctionTriggerNatsArgs>> nats() {
        return Optional.ofNullable(this.nats);
    }

    /**
     * `region`). The region in which the namespace should be created.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return `region`). The region in which the namespace should be created.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The configuration of the Scaleway&#39;s SQS used by the trigger
     * 
     */
    @Import(name="sqs")
    private @Nullable Output<FunctionTriggerSqsArgs> sqs;

    /**
     * @return The configuration of the Scaleway&#39;s SQS used by the trigger
     * 
     */
    public Optional<Output<FunctionTriggerSqsArgs>> sqs() {
        return Optional.ofNullable(this.sqs);
    }

    private FunctionTriggerArgs() {}

    private FunctionTriggerArgs(FunctionTriggerArgs $) {
        this.description = $.description;
        this.functionId = $.functionId;
        this.name = $.name;
        this.nats = $.nats;
        this.region = $.region;
        this.sqs = $.sqs;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FunctionTriggerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FunctionTriggerArgs $;

        public Builder() {
            $ = new FunctionTriggerArgs();
        }

        public Builder(FunctionTriggerArgs defaults) {
            $ = new FunctionTriggerArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The description of the trigger.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the trigger.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param functionId The ID of the function to create a trigger for
         * 
         * @return builder
         * 
         */
        public Builder functionId(Output<String> functionId) {
            $.functionId = functionId;
            return this;
        }

        /**
         * @param functionId The ID of the function to create a trigger for
         * 
         * @return builder
         * 
         */
        public Builder functionId(String functionId) {
            return functionId(Output.of(functionId));
        }

        /**
         * @param name The unique name of the trigger. Default to a generated name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The unique name of the trigger. Default to a generated name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param nats The configuration for the Scaleway&#39;s Nats used by the trigger
         * 
         * @return builder
         * 
         */
        public Builder nats(@Nullable Output<FunctionTriggerNatsArgs> nats) {
            $.nats = nats;
            return this;
        }

        /**
         * @param nats The configuration for the Scaleway&#39;s Nats used by the trigger
         * 
         * @return builder
         * 
         */
        public Builder nats(FunctionTriggerNatsArgs nats) {
            return nats(Output.of(nats));
        }

        /**
         * @param region `region`). The region in which the namespace should be created.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region `region`). The region in which the namespace should be created.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param sqs The configuration of the Scaleway&#39;s SQS used by the trigger
         * 
         * @return builder
         * 
         */
        public Builder sqs(@Nullable Output<FunctionTriggerSqsArgs> sqs) {
            $.sqs = sqs;
            return this;
        }

        /**
         * @param sqs The configuration of the Scaleway&#39;s SQS used by the trigger
         * 
         * @return builder
         * 
         */
        public Builder sqs(FunctionTriggerSqsArgs sqs) {
            return sqs(Output.of(sqs));
        }

        public FunctionTriggerArgs build() {
            $.functionId = Objects.requireNonNull($.functionId, "expected parameter 'functionId' to be non-null");
            return $;
        }
    }

}