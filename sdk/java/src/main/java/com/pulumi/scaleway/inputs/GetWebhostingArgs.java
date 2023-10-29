// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetWebhostingArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetWebhostingArgs Empty = new GetWebhostingArgs();

    /**
     * The hosting domain name. Only one of `domain` and `webhosting_id` should be specified.
     * 
     */
    @Import(name="domain")
    private @Nullable Output<String> domain;

    /**
     * @return The hosting domain name. Only one of `domain` and `webhosting_id` should be specified.
     * 
     */
    public Optional<Output<String>> domain() {
        return Optional.ofNullable(this.domain);
    }

    /**
     * The ID of the organization the hosting is associated with.
     * 
     */
    @Import(name="organizationId")
    private @Nullable Output<String> organizationId;

    /**
     * @return The ID of the organization the hosting is associated with.
     * 
     */
    public Optional<Output<String>> organizationId() {
        return Optional.ofNullable(this.organizationId);
    }

    /**
     * `project_id`) The ID of the project the hosting is associated with.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return `project_id`) The ID of the project the hosting is associated with.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * The hosting id. Only one of `domain` and `webhosting_id` should be specified.
     * 
     */
    @Import(name="webhostingId")
    private @Nullable Output<String> webhostingId;

    /**
     * @return The hosting id. Only one of `domain` and `webhosting_id` should be specified.
     * 
     */
    public Optional<Output<String>> webhostingId() {
        return Optional.ofNullable(this.webhostingId);
    }

    private GetWebhostingArgs() {}

    private GetWebhostingArgs(GetWebhostingArgs $) {
        this.domain = $.domain;
        this.organizationId = $.organizationId;
        this.projectId = $.projectId;
        this.webhostingId = $.webhostingId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetWebhostingArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetWebhostingArgs $;

        public Builder() {
            $ = new GetWebhostingArgs();
        }

        public Builder(GetWebhostingArgs defaults) {
            $ = new GetWebhostingArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param domain The hosting domain name. Only one of `domain` and `webhosting_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder domain(@Nullable Output<String> domain) {
            $.domain = domain;
            return this;
        }

        /**
         * @param domain The hosting domain name. Only one of `domain` and `webhosting_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder domain(String domain) {
            return domain(Output.of(domain));
        }

        /**
         * @param organizationId The ID of the organization the hosting is associated with.
         * 
         * @return builder
         * 
         */
        public Builder organizationId(@Nullable Output<String> organizationId) {
            $.organizationId = organizationId;
            return this;
        }

        /**
         * @param organizationId The ID of the organization the hosting is associated with.
         * 
         * @return builder
         * 
         */
        public Builder organizationId(String organizationId) {
            return organizationId(Output.of(organizationId));
        }

        /**
         * @param projectId `project_id`) The ID of the project the hosting is associated with.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId `project_id`) The ID of the project the hosting is associated with.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param webhostingId The hosting id. Only one of `domain` and `webhosting_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder webhostingId(@Nullable Output<String> webhostingId) {
            $.webhostingId = webhostingId;
            return this;
        }

        /**
         * @param webhostingId The hosting id. Only one of `domain` and `webhosting_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder webhostingId(String webhostingId) {
            return webhostingId(Output.of(webhostingId));
        }

        public GetWebhostingArgs build() {
            return $;
        }
    }

}