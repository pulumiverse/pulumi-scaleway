// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetMnqSqsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetMnqSqsArgs Empty = new GetMnqSqsArgs();

    /**
     * `project_id`) The ID of the project for which sqs is enabled.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return `project_id`) The ID of the project for which sqs is enabled.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * `region`). The region in which sqs is enabled.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return `region`). The region in which sqs is enabled.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    private GetMnqSqsArgs() {}

    private GetMnqSqsArgs(GetMnqSqsArgs $) {
        this.projectId = $.projectId;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetMnqSqsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetMnqSqsArgs $;

        public Builder() {
            $ = new GetMnqSqsArgs();
        }

        public Builder(GetMnqSqsArgs defaults) {
            $ = new GetMnqSqsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param projectId `project_id`) The ID of the project for which sqs is enabled.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId `project_id`) The ID of the project for which sqs is enabled.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param region `region`). The region in which sqs is enabled.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region `region`). The region in which sqs is enabled.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        public GetMnqSqsArgs build() {
            return $;
        }
    }

}