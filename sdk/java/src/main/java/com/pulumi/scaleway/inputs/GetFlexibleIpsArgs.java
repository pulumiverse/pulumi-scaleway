// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetFlexibleIpsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetFlexibleIpsArgs Empty = new GetFlexibleIpsArgs();

    /**
     * (Defaults to provider `project_id`) The ID of the project the IP is in.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return (Defaults to provider `project_id`) The ID of the project the IP is in.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * List of server IDs used as filter. IPs with these exact server IDs are listed.
     * 
     */
    @Import(name="serverIds")
    private @Nullable Output<List<String>> serverIds;

    /**
     * @return List of server IDs used as filter. IPs with these exact server IDs are listed.
     * 
     */
    public Optional<Output<List<String>>> serverIds() {
        return Optional.ofNullable(this.serverIds);
    }

    /**
     * List of tags used as filter. IPs with these exact tags are listed.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return List of tags used as filter. IPs with these exact tags are listed.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * `zone`) The zone in which IPs exist.
     * 
     */
    @Import(name="zone")
    private @Nullable Output<String> zone;

    /**
     * @return `zone`) The zone in which IPs exist.
     * 
     */
    public Optional<Output<String>> zone() {
        return Optional.ofNullable(this.zone);
    }

    private GetFlexibleIpsArgs() {}

    private GetFlexibleIpsArgs(GetFlexibleIpsArgs $) {
        this.projectId = $.projectId;
        this.serverIds = $.serverIds;
        this.tags = $.tags;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetFlexibleIpsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetFlexibleIpsArgs $;

        public Builder() {
            $ = new GetFlexibleIpsArgs();
        }

        public Builder(GetFlexibleIpsArgs defaults) {
            $ = new GetFlexibleIpsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param projectId (Defaults to provider `project_id`) The ID of the project the IP is in.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId (Defaults to provider `project_id`) The ID of the project the IP is in.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param serverIds List of server IDs used as filter. IPs with these exact server IDs are listed.
         * 
         * @return builder
         * 
         */
        public Builder serverIds(@Nullable Output<List<String>> serverIds) {
            $.serverIds = serverIds;
            return this;
        }

        /**
         * @param serverIds List of server IDs used as filter. IPs with these exact server IDs are listed.
         * 
         * @return builder
         * 
         */
        public Builder serverIds(List<String> serverIds) {
            return serverIds(Output.of(serverIds));
        }

        /**
         * @param serverIds List of server IDs used as filter. IPs with these exact server IDs are listed.
         * 
         * @return builder
         * 
         */
        public Builder serverIds(String... serverIds) {
            return serverIds(List.of(serverIds));
        }

        /**
         * @param tags List of tags used as filter. IPs with these exact tags are listed.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags List of tags used as filter. IPs with these exact tags are listed.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags List of tags used as filter. IPs with these exact tags are listed.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param zone `zone`) The zone in which IPs exist.
         * 
         * @return builder
         * 
         */
        public Builder zone(@Nullable Output<String> zone) {
            $.zone = zone;
            return this;
        }

        /**
         * @param zone `zone`) The zone in which IPs exist.
         * 
         * @return builder
         * 
         */
        public Builder zone(String zone) {
            return zone(Output.of(zone));
        }

        public GetFlexibleIpsArgs build() {
            return $;
        }
    }

}