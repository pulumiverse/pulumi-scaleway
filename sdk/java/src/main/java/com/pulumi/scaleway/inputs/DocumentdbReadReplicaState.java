// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.scaleway.inputs.DocumentdbReadReplicaDirectAccessArgs;
import com.pulumi.scaleway.inputs.DocumentdbReadReplicaPrivateNetworkArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DocumentdbReadReplicaState extends com.pulumi.resources.ResourceArgs {

    public static final DocumentdbReadReplicaState Empty = new DocumentdbReadReplicaState();

    /**
     * Creates a direct access endpoint to documentdb replica.
     * 
     */
    @Import(name="directAccess")
    private @Nullable Output<DocumentdbReadReplicaDirectAccessArgs> directAccess;

    /**
     * @return Creates a direct access endpoint to documentdb replica.
     * 
     */
    public Optional<Output<DocumentdbReadReplicaDirectAccessArgs>> directAccess() {
        return Optional.ofNullable(this.directAccess);
    }

    /**
     * UUID of the documentdb instance.
     * 
     * &gt; **Important:** The replica musts contains at least one of `direct_access` or `private_network`. It can contain both.
     * 
     */
    @Import(name="instanceId")
    private @Nullable Output<String> instanceId;

    /**
     * @return UUID of the documentdb instance.
     * 
     * &gt; **Important:** The replica musts contains at least one of `direct_access` or `private_network`. It can contain both.
     * 
     */
    public Optional<Output<String>> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    /**
     * Create an endpoint in a private network.
     * 
     */
    @Import(name="privateNetwork")
    private @Nullable Output<DocumentdbReadReplicaPrivateNetworkArgs> privateNetwork;

    /**
     * @return Create an endpoint in a private network.
     * 
     */
    public Optional<Output<DocumentdbReadReplicaPrivateNetworkArgs>> privateNetwork() {
        return Optional.ofNullable(this.privateNetwork);
    }

    /**
     * `region`) The region
     * in which the Database read replica should be created.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return `region`) The region
     * in which the Database read replica should be created.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    private DocumentdbReadReplicaState() {}

    private DocumentdbReadReplicaState(DocumentdbReadReplicaState $) {
        this.directAccess = $.directAccess;
        this.instanceId = $.instanceId;
        this.privateNetwork = $.privateNetwork;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DocumentdbReadReplicaState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DocumentdbReadReplicaState $;

        public Builder() {
            $ = new DocumentdbReadReplicaState();
        }

        public Builder(DocumentdbReadReplicaState defaults) {
            $ = new DocumentdbReadReplicaState(Objects.requireNonNull(defaults));
        }

        /**
         * @param directAccess Creates a direct access endpoint to documentdb replica.
         * 
         * @return builder
         * 
         */
        public Builder directAccess(@Nullable Output<DocumentdbReadReplicaDirectAccessArgs> directAccess) {
            $.directAccess = directAccess;
            return this;
        }

        /**
         * @param directAccess Creates a direct access endpoint to documentdb replica.
         * 
         * @return builder
         * 
         */
        public Builder directAccess(DocumentdbReadReplicaDirectAccessArgs directAccess) {
            return directAccess(Output.of(directAccess));
        }

        /**
         * @param instanceId UUID of the documentdb instance.
         * 
         * &gt; **Important:** The replica musts contains at least one of `direct_access` or `private_network`. It can contain both.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(@Nullable Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId UUID of the documentdb instance.
         * 
         * &gt; **Important:** The replica musts contains at least one of `direct_access` or `private_network`. It can contain both.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param privateNetwork Create an endpoint in a private network.
         * 
         * @return builder
         * 
         */
        public Builder privateNetwork(@Nullable Output<DocumentdbReadReplicaPrivateNetworkArgs> privateNetwork) {
            $.privateNetwork = privateNetwork;
            return this;
        }

        /**
         * @param privateNetwork Create an endpoint in a private network.
         * 
         * @return builder
         * 
         */
        public Builder privateNetwork(DocumentdbReadReplicaPrivateNetworkArgs privateNetwork) {
            return privateNetwork(Output.of(privateNetwork));
        }

        /**
         * @param region `region`) The region
         * in which the Database read replica should be created.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region `region`) The region
         * in which the Database read replica should be created.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        public DocumentdbReadReplicaState build() {
            return $;
        }
    }

}