// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.scaleway.DocumentdbPrivateNetworkEndpointArgs;
import com.pulumi.scaleway.Utilities;
import com.pulumi.scaleway.inputs.DocumentdbPrivateNetworkEndpointState;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * Database Instance Endpoint can be imported using the `{region}/{endpoint_id}`, e.g. bash
 * 
 * ```sh
 *  $ pulumi import scaleway:index/documentdbPrivateNetworkEndpoint:DocumentdbPrivateNetworkEndpoint end fr-par/11111111-1111-1111-1111-111111111111
 * ```
 * 
 */
@ResourceType(type="scaleway:index/documentdbPrivateNetworkEndpoint:DocumentdbPrivateNetworkEndpoint")
public class DocumentdbPrivateNetworkEndpoint extends com.pulumi.resources.CustomResource {
    /**
     * Hostname of the endpoint.
     * 
     */
    @Export(name="hostname", refs={String.class}, tree="[0]")
    private Output<String> hostname;

    /**
     * @return Hostname of the endpoint.
     * 
     */
    public Output<String> hostname() {
        return this.hostname;
    }
    /**
     * UUID of the documentdb instance.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return UUID of the documentdb instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * IPv4 address on the network.
     * 
     */
    @Export(name="ip", refs={String.class}, tree="[0]")
    private Output<String> ip;

    /**
     * @return IPv4 address on the network.
     * 
     */
    public Output<String> ip() {
        return this.ip;
    }
    /**
     * The IP network address within the private subnet. This must be an IPv4 address with a
     * CIDR notation. The IP network address within the private subnet is determined by the IP Address Management (IPAM)
     * service if not set.
     * 
     */
    @Export(name="ipNet", refs={String.class}, tree="[0]")
    private Output<String> ipNet;

    /**
     * @return The IP network address within the private subnet. This must be an IPv4 address with a
     * CIDR notation. The IP network address within the private subnet is determined by the IP Address Management (IPAM)
     * service if not set.
     * 
     */
    public Output<String> ipNet() {
        return this.ipNet;
    }
    /**
     * Name of the endpoint.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the endpoint.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Port in the Private Network.
     * 
     */
    @Export(name="port", refs={Integer.class}, tree="[0]")
    private Output<Integer> port;

    /**
     * @return Port in the Private Network.
     * 
     */
    public Output<Integer> port() {
        return this.port;
    }
    /**
     * The ID of the private network.
     * 
     */
    @Export(name="privateNetworkId", refs={String.class}, tree="[0]")
    private Output<String> privateNetworkId;

    /**
     * @return The ID of the private network.
     * 
     */
    public Output<String> privateNetworkId() {
        return this.privateNetworkId;
    }
    /**
     * The region you want to attach the resource to
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region you want to attach the resource to
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The zone you want to attach the resource to
     * 
     */
    @Export(name="zone", refs={String.class}, tree="[0]")
    private Output<String> zone;

    /**
     * @return The zone you want to attach the resource to
     * 
     */
    public Output<String> zone() {
        return this.zone;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DocumentdbPrivateNetworkEndpoint(String name) {
        this(name, DocumentdbPrivateNetworkEndpointArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DocumentdbPrivateNetworkEndpoint(String name, DocumentdbPrivateNetworkEndpointArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DocumentdbPrivateNetworkEndpoint(String name, DocumentdbPrivateNetworkEndpointArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/documentdbPrivateNetworkEndpoint:DocumentdbPrivateNetworkEndpoint", name, args == null ? DocumentdbPrivateNetworkEndpointArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DocumentdbPrivateNetworkEndpoint(String name, Output<String> id, @Nullable DocumentdbPrivateNetworkEndpointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/documentdbPrivateNetworkEndpoint:DocumentdbPrivateNetworkEndpoint", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static DocumentdbPrivateNetworkEndpoint get(String name, Output<String> id, @Nullable DocumentdbPrivateNetworkEndpointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DocumentdbPrivateNetworkEndpoint(name, id, state, options);
    }
}