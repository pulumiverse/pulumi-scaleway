// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates and manages Scaleway Compute Baremetal servers. For more information, see the [API documentation](https://www.scaleway.com/en/developers/api/elastic-metal/).
 *
 * ## Example Usage
 *
 * ### Without install config
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const myOffer = scaleway.elasticmetal.getOffer({
 *     zone: "fr-par-2",
 *     name: "EM-B112X-SSD",
 * });
 * const myServer = new scaleway.elasticmetal.Server("my_server", {
 *     zone: "fr-par-2",
 *     offer: myOffer.then(myOffer => myOffer.offerId),
 *     installConfigAfterward: true,
 * });
 * ```
 *
 * ### With custom partitioning
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const config = new pulumi.Config();
 * const configCustomPartitioning = config.get("configCustomPartitioning") || "{\"disks\":[{\"device\":\"/dev/nvme0n1\",\"partitions\":[{\"label\":\"uefi\",\"number\":1,\"size\":536870912,\"useAllAvailableSpace\":false},{\"label\":\"boot\",\"number\":2,\"size\":536870912,\"useAllAvailableSpace\":false},{\"label\":\"root\",\"number\":3,\"size\":1018839433216,\"useAllAvailableSpace\":false}]},{\"device\":\"/dev/nvme1n1\",\"partitions\":[{\"label\":\"boot\",\"number\":1,\"size\":536870912,\"useAllAvailableSpace\":false},{\"label\":\"data\",\"number\":2,\"size\":1018839433216,\"useAllAvailableSpace\":false}]}],\"filesystems\":[{\"device\":\"/dev/nvme0n1p1\",\"format\":\"fat32\",\"mountpoint\":\"/boot/efi\"},{\"device\":\"/dev/nvme0n1p2\",\"format\":\"ext4\",\"mountpoint\":\"/boot\"},{\"device\":\"/dev/nvme0n1p3\",\"format\":\"ext4\",\"mountpoint\":\"/\"},{\"device\":\"/dev/nvme1n1p2\",\"format\":\"ext4\",\"mountpoint\":\"/data\"}],\"raids\":[]}";
 * const myOs = scaleway.elasticmetal.getOs({
 *     zone: "fr-par-1",
 *     name: "Ubuntu",
 *     version: "22.04 LTS (Jammy Jellyfish)",
 * });
 * const mySshKey = new scaleway.iam.SshKey("my_ssh_key", {
 *     name: "my_ssh_key",
 *     publicKey: "ssh XXXXXXXXXXX",
 * });
 * const myOffer = scaleway.elasticmetal.getOffer({
 *     zone: "fr-par-1",
 *     name: "EM-B220E-NVME",
 *     subscriptionPeriod: "hourly",
 * });
 * const myServer = new scaleway.elasticmetal.Server("my_server", {
 *     name: "my_super_server",
 *     zone: "fr-par-1",
 *     description: "test a description",
 *     offer: myOffer.then(myOffer => myOffer.offerId),
 *     os: myOs.then(myOs => myOs.osId),
 *     partitioning: configCustomPartitioning,
 *     tags: [
 *         "terraform-test",
 *         "scaleway_baremetal_server",
 *         "minimal",
 *     ],
 *     sshKeyIds: [mySshKey.id],
 * });
 * ```
 *
 * ### Migrate from hourly to monthly plan
 *
 * To migrate from an hourly to a monthly subscription for a Scaleway Baremetal server, it is important to understand that the migration can only be done by using the data source.
 * You cannot directly modify the subscriptionPeriod of an existing scaleway.elasticmetal.getOffer resource. Instead, you must define the monthly offer using the data source and then update the server configuration accordingly.
 *
 * ### Hourly Plan Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const myOffer = scaleway.elasticmetal.getOffer({
 *     zone: "fr-par-1",
 *     name: "EM-B220E-NVME",
 *     subscriptionPeriod: "hourly",
 * });
 * const myServer = new scaleway.elasticmetal.Server("my_server", {
 *     name: "UpdateSubscriptionPeriod",
 *     offer: myOffer.then(myOffer => myOffer.offerId),
 *     zone: "%s",
 *     installConfigAfterward: true,
 * });
 * ```
 *
 * ### Monthly Plan Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const myOffer = scaleway.elasticmetal.getOffer({
 *     zone: "fr-par-1",
 *     name: "EM-B220E-NVME",
 *     subscriptionPeriod: "monthly",
 * });
 * const myServer = new scaleway.elasticmetal.Server("my_server", {
 *     name: "UpdateSubscriptionPeriod",
 *     offer: myOffer.then(myOffer => myOffer.offerId),
 *     zone: "fr-par-1",
 *     installConfigAfterward: true,
 * });
 * ```
 *
 * **Important**  Once you migrate to a monthly subscription, you cannot downgrade back to an hourly plan. Ensure that the monthly plan meets your needs before making the switch.
 *
 * ## Import
 *
 * Baremetal servers can be imported using the `{zone}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:elasticmetal/server:Server web fr-par-2/11111111-1111-1111-1111-111111111111
 * ```
 */
export class Server extends pulumi.CustomResource {
    /**
     * Get an existing Server resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServerState, opts?: pulumi.CustomResourceOptions): Server {
        return new Server(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:elasticmetal/server:Server';

    /**
     * Returns true if the given object is an instance of Server.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Server {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Server.__pulumiType;
    }

    /**
     * A description for the server.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The domain of the server.
     */
    public /*out*/ readonly domain!: pulumi.Output<string>;
    /**
     * The hostname of the server.
     */
    public readonly hostname!: pulumi.Output<string | undefined>;
    /**
     * If True, this boolean allows to create a server without the install config if you want to provide it later.
     */
    public readonly installConfigAfterward!: pulumi.Output<boolean | undefined>;
    /**
     * (List of) The IPs of the server.
     */
    public /*out*/ readonly ips!: pulumi.Output<outputs.elasticmetal.ServerIp[]>;
    /**
     * (List of) The IPv4 addresses of the server.
     */
    public /*out*/ readonly ipv4s!: pulumi.Output<outputs.elasticmetal.ServerIpv4[]>;
    /**
     * (List of) The IPv6 addresses of the server.
     */
    public /*out*/ readonly ipv6s!: pulumi.Output<outputs.elasticmetal.ServerIpv6[]>;
    /**
     * The name of the server.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The offer UUID of the baremetal server.
     * Use [this endpoint](https://www.scaleway.com/en/developers/api/elastic-metal/#path-servers-get-a-specific-elastic-metal-server) to find the right offer.
     *
     * > **Important:** Updates to `offer` will recreate the server.
     */
    public readonly offer!: pulumi.Output<string>;
    /**
     * The ID of the offer.
     */
    public /*out*/ readonly offerId!: pulumi.Output<string>;
    /**
     * The name of the offer.
     */
    public /*out*/ readonly offerName!: pulumi.Output<string>;
    /**
     * The options to enable on the server.
     * > The `options` block supports:
     */
    public readonly options!: pulumi.Output<outputs.elasticmetal.ServerOption[] | undefined>;
    /**
     * The organization ID the server is associated with.
     */
    public /*out*/ readonly organizationId!: pulumi.Output<string>;
    /**
     * The UUID of the os to install on the server.
     * Use [this endpoint](https://www.scaleway.com/en/developers/api/elastic-metal/#path-os-list-available-oses) to find the right OS ID.
     * > **Important:** Updates to `os` will reinstall the server.
     */
    public readonly os!: pulumi.Output<string | undefined>;
    /**
     * The name of the os.
     */
    public /*out*/ readonly osName!: pulumi.Output<string>;
    /**
     * The partitioning schema in JSON format
     */
    public readonly partitioning!: pulumi.Output<string | undefined>;
    /**
     * Password used for the installation. May be required depending on used os.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * The list of private IPv4 and IPv6 addresses associated with the resource.
     */
    public readonly privateIps!: pulumi.Output<outputs.elasticmetal.ServerPrivateIp[]>;
    /**
     * The private networks to attach to the server. For more information, see [the documentation](https://www.scaleway.com/en/docs/compute/elastic-metal/how-to/use-private-networks/)
     */
    public readonly privateNetworks!: pulumi.Output<outputs.elasticmetal.ServerPrivateNetwork[] | undefined>;
    /**
     * `projectId`) The ID of the project the server is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * If True, this boolean allows to reinstall the server on install config changes.
     * > **Important:** Updates to `sshKeyIds`, `user`, `password`, `serviceUser` or `servicePassword` will not take effect on the server, it requires to reinstall it. To do so please set 'reinstall_on_config_changes' argument to true.
     */
    public readonly reinstallOnConfigChanges!: pulumi.Output<boolean | undefined>;
    /**
     * Password used for the service to install. May be required depending on used os.
     */
    public readonly servicePassword!: pulumi.Output<string | undefined>;
    /**
     * User used for the service to install.
     */
    public readonly serviceUser!: pulumi.Output<string>;
    /**
     * List of SSH keys allowed to connect to the server.
     */
    public readonly sshKeyIds!: pulumi.Output<string[] | undefined>;
    /**
     * The tags associated with the server.
     */
    public readonly tags!: pulumi.Output<string[]>;
    /**
     * User used for the installation.
     */
    public readonly user!: pulumi.Output<string>;
    /**
     * `zone`) The zone in which the server should be created.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a Server resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServerArgs | ServerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServerState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["hostname"] = state ? state.hostname : undefined;
            resourceInputs["installConfigAfterward"] = state ? state.installConfigAfterward : undefined;
            resourceInputs["ips"] = state ? state.ips : undefined;
            resourceInputs["ipv4s"] = state ? state.ipv4s : undefined;
            resourceInputs["ipv6s"] = state ? state.ipv6s : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["offer"] = state ? state.offer : undefined;
            resourceInputs["offerId"] = state ? state.offerId : undefined;
            resourceInputs["offerName"] = state ? state.offerName : undefined;
            resourceInputs["options"] = state ? state.options : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["os"] = state ? state.os : undefined;
            resourceInputs["osName"] = state ? state.osName : undefined;
            resourceInputs["partitioning"] = state ? state.partitioning : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["privateIps"] = state ? state.privateIps : undefined;
            resourceInputs["privateNetworks"] = state ? state.privateNetworks : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["reinstallOnConfigChanges"] = state ? state.reinstallOnConfigChanges : undefined;
            resourceInputs["servicePassword"] = state ? state.servicePassword : undefined;
            resourceInputs["serviceUser"] = state ? state.serviceUser : undefined;
            resourceInputs["sshKeyIds"] = state ? state.sshKeyIds : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["user"] = state ? state.user : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as ServerArgs | undefined;
            if ((!args || args.offer === undefined) && !opts.urn) {
                throw new Error("Missing required property 'offer'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["hostname"] = args ? args.hostname : undefined;
            resourceInputs["installConfigAfterward"] = args ? args.installConfigAfterward : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["offer"] = args ? args.offer : undefined;
            resourceInputs["options"] = args ? args.options : undefined;
            resourceInputs["os"] = args ? args.os : undefined;
            resourceInputs["partitioning"] = args ? args.partitioning : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["privateIps"] = args ? args.privateIps : undefined;
            resourceInputs["privateNetworks"] = args ? args.privateNetworks : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["reinstallOnConfigChanges"] = args ? args.reinstallOnConfigChanges : undefined;
            resourceInputs["servicePassword"] = args?.servicePassword ? pulumi.secret(args.servicePassword) : undefined;
            resourceInputs["serviceUser"] = args ? args.serviceUser : undefined;
            resourceInputs["sshKeyIds"] = args ? args.sshKeyIds : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["user"] = args ? args.user : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["domain"] = undefined /*out*/;
            resourceInputs["ips"] = undefined /*out*/;
            resourceInputs["ipv4s"] = undefined /*out*/;
            resourceInputs["ipv6s"] = undefined /*out*/;
            resourceInputs["offerId"] = undefined /*out*/;
            resourceInputs["offerName"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["osName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "scaleway:index/baremetalServer:BaremetalServer" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        const secretOpts = { additionalSecretOutputs: ["password", "servicePassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Server.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Server resources.
 */
export interface ServerState {
    /**
     * A description for the server.
     */
    description?: pulumi.Input<string>;
    /**
     * The domain of the server.
     */
    domain?: pulumi.Input<string>;
    /**
     * The hostname of the server.
     */
    hostname?: pulumi.Input<string>;
    /**
     * If True, this boolean allows to create a server without the install config if you want to provide it later.
     */
    installConfigAfterward?: pulumi.Input<boolean>;
    /**
     * (List of) The IPs of the server.
     */
    ips?: pulumi.Input<pulumi.Input<inputs.elasticmetal.ServerIp>[]>;
    /**
     * (List of) The IPv4 addresses of the server.
     */
    ipv4s?: pulumi.Input<pulumi.Input<inputs.elasticmetal.ServerIpv4>[]>;
    /**
     * (List of) The IPv6 addresses of the server.
     */
    ipv6s?: pulumi.Input<pulumi.Input<inputs.elasticmetal.ServerIpv6>[]>;
    /**
     * The name of the server.
     */
    name?: pulumi.Input<string>;
    /**
     * The offer UUID of the baremetal server.
     * Use [this endpoint](https://www.scaleway.com/en/developers/api/elastic-metal/#path-servers-get-a-specific-elastic-metal-server) to find the right offer.
     *
     * > **Important:** Updates to `offer` will recreate the server.
     */
    offer?: pulumi.Input<string>;
    /**
     * The ID of the offer.
     */
    offerId?: pulumi.Input<string>;
    /**
     * The name of the offer.
     */
    offerName?: pulumi.Input<string>;
    /**
     * The options to enable on the server.
     * > The `options` block supports:
     */
    options?: pulumi.Input<pulumi.Input<inputs.elasticmetal.ServerOption>[]>;
    /**
     * The organization ID the server is associated with.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * The UUID of the os to install on the server.
     * Use [this endpoint](https://www.scaleway.com/en/developers/api/elastic-metal/#path-os-list-available-oses) to find the right OS ID.
     * > **Important:** Updates to `os` will reinstall the server.
     */
    os?: pulumi.Input<string>;
    /**
     * The name of the os.
     */
    osName?: pulumi.Input<string>;
    /**
     * The partitioning schema in JSON format
     */
    partitioning?: pulumi.Input<string>;
    /**
     * Password used for the installation. May be required depending on used os.
     */
    password?: pulumi.Input<string>;
    /**
     * The list of private IPv4 and IPv6 addresses associated with the resource.
     */
    privateIps?: pulumi.Input<pulumi.Input<inputs.elasticmetal.ServerPrivateIp>[]>;
    /**
     * The private networks to attach to the server. For more information, see [the documentation](https://www.scaleway.com/en/docs/compute/elastic-metal/how-to/use-private-networks/)
     */
    privateNetworks?: pulumi.Input<pulumi.Input<inputs.elasticmetal.ServerPrivateNetwork>[]>;
    /**
     * `projectId`) The ID of the project the server is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * If True, this boolean allows to reinstall the server on install config changes.
     * > **Important:** Updates to `sshKeyIds`, `user`, `password`, `serviceUser` or `servicePassword` will not take effect on the server, it requires to reinstall it. To do so please set 'reinstall_on_config_changes' argument to true.
     */
    reinstallOnConfigChanges?: pulumi.Input<boolean>;
    /**
     * Password used for the service to install. May be required depending on used os.
     */
    servicePassword?: pulumi.Input<string>;
    /**
     * User used for the service to install.
     */
    serviceUser?: pulumi.Input<string>;
    /**
     * List of SSH keys allowed to connect to the server.
     */
    sshKeyIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The tags associated with the server.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * User used for the installation.
     */
    user?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the server should be created.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Server resource.
 */
export interface ServerArgs {
    /**
     * A description for the server.
     */
    description?: pulumi.Input<string>;
    /**
     * The hostname of the server.
     */
    hostname?: pulumi.Input<string>;
    /**
     * If True, this boolean allows to create a server without the install config if you want to provide it later.
     */
    installConfigAfterward?: pulumi.Input<boolean>;
    /**
     * The name of the server.
     */
    name?: pulumi.Input<string>;
    /**
     * The offer UUID of the baremetal server.
     * Use [this endpoint](https://www.scaleway.com/en/developers/api/elastic-metal/#path-servers-get-a-specific-elastic-metal-server) to find the right offer.
     *
     * > **Important:** Updates to `offer` will recreate the server.
     */
    offer: pulumi.Input<string>;
    /**
     * The options to enable on the server.
     * > The `options` block supports:
     */
    options?: pulumi.Input<pulumi.Input<inputs.elasticmetal.ServerOption>[]>;
    /**
     * The UUID of the os to install on the server.
     * Use [this endpoint](https://www.scaleway.com/en/developers/api/elastic-metal/#path-os-list-available-oses) to find the right OS ID.
     * > **Important:** Updates to `os` will reinstall the server.
     */
    os?: pulumi.Input<string>;
    /**
     * The partitioning schema in JSON format
     */
    partitioning?: pulumi.Input<string>;
    /**
     * Password used for the installation. May be required depending on used os.
     */
    password?: pulumi.Input<string>;
    /**
     * The list of private IPv4 and IPv6 addresses associated with the resource.
     */
    privateIps?: pulumi.Input<pulumi.Input<inputs.elasticmetal.ServerPrivateIp>[]>;
    /**
     * The private networks to attach to the server. For more information, see [the documentation](https://www.scaleway.com/en/docs/compute/elastic-metal/how-to/use-private-networks/)
     */
    privateNetworks?: pulumi.Input<pulumi.Input<inputs.elasticmetal.ServerPrivateNetwork>[]>;
    /**
     * `projectId`) The ID of the project the server is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * If True, this boolean allows to reinstall the server on install config changes.
     * > **Important:** Updates to `sshKeyIds`, `user`, `password`, `serviceUser` or `servicePassword` will not take effect on the server, it requires to reinstall it. To do so please set 'reinstall_on_config_changes' argument to true.
     */
    reinstallOnConfigChanges?: pulumi.Input<boolean>;
    /**
     * Password used for the service to install. May be required depending on used os.
     */
    servicePassword?: pulumi.Input<string>;
    /**
     * User used for the service to install.
     */
    serviceUser?: pulumi.Input<string>;
    /**
     * List of SSH keys allowed to connect to the server.
     */
    sshKeyIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The tags associated with the server.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * User used for the installation.
     */
    user?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the server should be created.
     */
    zone?: pulumi.Input<string>;
}
