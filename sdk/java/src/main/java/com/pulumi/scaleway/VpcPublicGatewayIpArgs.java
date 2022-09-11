// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VpcPublicGatewayIpArgs extends com.pulumi.resources.ResourceArgs {

    public static final VpcPublicGatewayIpArgs Empty = new VpcPublicGatewayIpArgs();

    /**
     * `project_id`) The ID of the project the public gateway ip is associated with.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return `project_id`) The ID of the project the public gateway ip is associated with.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * The reverse domain name for the IP address
     * 
     */
    @Import(name="reverse")
    private @Nullable Output<String> reverse;

    /**
     * @return The reverse domain name for the IP address
     * 
     */
    public Optional<Output<String>> reverse() {
        return Optional.ofNullable(this.reverse);
    }

    /**
     * The tags associated with the public gateway IP.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return The tags associated with the public gateway IP.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * `zone`) The zone in which the public gateway ip should be created.
     * 
     */
    @Import(name="zone")
    private @Nullable Output<String> zone;

    /**
     * @return `zone`) The zone in which the public gateway ip should be created.
     * 
     */
    public Optional<Output<String>> zone() {
        return Optional.ofNullable(this.zone);
    }

    private VpcPublicGatewayIpArgs() {}

    private VpcPublicGatewayIpArgs(VpcPublicGatewayIpArgs $) {
        this.projectId = $.projectId;
        this.reverse = $.reverse;
        this.tags = $.tags;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VpcPublicGatewayIpArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VpcPublicGatewayIpArgs $;

        public Builder() {
            $ = new VpcPublicGatewayIpArgs();
        }

        public Builder(VpcPublicGatewayIpArgs defaults) {
            $ = new VpcPublicGatewayIpArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param projectId `project_id`) The ID of the project the public gateway ip is associated with.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId `project_id`) The ID of the project the public gateway ip is associated with.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param reverse The reverse domain name for the IP address
         * 
         * @return builder
         * 
         */
        public Builder reverse(@Nullable Output<String> reverse) {
            $.reverse = reverse;
            return this;
        }

        /**
         * @param reverse The reverse domain name for the IP address
         * 
         * @return builder
         * 
         */
        public Builder reverse(String reverse) {
            return reverse(Output.of(reverse));
        }

        /**
         * @param tags The tags associated with the public gateway IP.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The tags associated with the public gateway IP.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags The tags associated with the public gateway IP.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param zone `zone`) The zone in which the public gateway ip should be created.
         * 
         * @return builder
         * 
         */
        public Builder zone(@Nullable Output<String> zone) {
            $.zone = zone;
            return this;
        }

        /**
         * @param zone `zone`) The zone in which the public gateway ip should be created.
         * 
         * @return builder
         * 
         */
        public Builder zone(String zone) {
            return zone(Output.of(zone));
        }

        public VpcPublicGatewayIpArgs build() {
            return $;
        }
    }

}