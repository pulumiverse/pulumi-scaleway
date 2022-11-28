// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetInstanceSnapshotArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetInstanceSnapshotArgs Empty = new GetInstanceSnapshotArgs();

    /**
     * The snapshot name.
     * Only one of `name` and `snapshot_id` should be specified.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The snapshot name.
     * Only one of `name` and `snapshot_id` should be specified.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The snapshot id.
     * Only one of `name` and `snapshot_id` should be specified.
     * 
     */
    @Import(name="snapshotId")
    private @Nullable Output<String> snapshotId;

    /**
     * @return The snapshot id.
     * Only one of `name` and `snapshot_id` should be specified.
     * 
     */
    public Optional<Output<String>> snapshotId() {
        return Optional.ofNullable(this.snapshotId);
    }

    /**
     * `zone`) The zone in which the snapshot exists.
     * 
     */
    @Import(name="zone")
    private @Nullable Output<String> zone;

    /**
     * @return `zone`) The zone in which the snapshot exists.
     * 
     */
    public Optional<Output<String>> zone() {
        return Optional.ofNullable(this.zone);
    }

    private GetInstanceSnapshotArgs() {}

    private GetInstanceSnapshotArgs(GetInstanceSnapshotArgs $) {
        this.name = $.name;
        this.snapshotId = $.snapshotId;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetInstanceSnapshotArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetInstanceSnapshotArgs $;

        public Builder() {
            $ = new GetInstanceSnapshotArgs();
        }

        public Builder(GetInstanceSnapshotArgs defaults) {
            $ = new GetInstanceSnapshotArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The snapshot name.
         * Only one of `name` and `snapshot_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The snapshot name.
         * Only one of `name` and `snapshot_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param snapshotId The snapshot id.
         * Only one of `name` and `snapshot_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder snapshotId(@Nullable Output<String> snapshotId) {
            $.snapshotId = snapshotId;
            return this;
        }

        /**
         * @param snapshotId The snapshot id.
         * Only one of `name` and `snapshot_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder snapshotId(String snapshotId) {
            return snapshotId(Output.of(snapshotId));
        }

        /**
         * @param zone `zone`) The zone in which the snapshot exists.
         * 
         * @return builder
         * 
         */
        public Builder zone(@Nullable Output<String> zone) {
            $.zone = zone;
            return this;
        }

        /**
         * @param zone `zone`) The zone in which the snapshot exists.
         * 
         * @return builder
         * 
         */
        public Builder zone(String zone) {
            return zone(Output.of(zone));
        }

        public GetInstanceSnapshotArgs build() {
            return $;
        }
    }

}