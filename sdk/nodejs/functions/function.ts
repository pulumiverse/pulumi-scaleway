// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Functions can be imported using, `{region}/{id}`, as shown below:
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:functions/function:Function main fr-par/11111111-1111-1111-1111-111111111111
 * ```
 */
export class Function extends pulumi.CustomResource {
    /**
     * Get an existing Function resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FunctionState, opts?: pulumi.CustomResourceOptions): Function {
        return new Function(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:functions/function:Function';

    /**
     * Returns true if the given object is an instance of Function.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Function {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Function.__pulumiType;
    }

    /**
     * The CPU limit in mVCPU for your function.
     */
    public /*out*/ readonly cpuLimit!: pulumi.Output<number>;
    public readonly deploy!: pulumi.Output<boolean | undefined>;
    /**
     * The description of the function.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The native domain name of the function.
     */
    public /*out*/ readonly domainName!: pulumi.Output<string>;
    /**
     * The [environment variables](https://www.scaleway.com/en/docs/compute/functions/concepts/#environment-variables) of the function.
     */
    public readonly environmentVariables!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Handler of the function, depends on the runtime. Refer to the [dedicated documentation](https://www.scaleway.com/en/developers/api/serverless-functions/#path-functions-create-a-new-function) for the list of supported runtimes.
     */
    public readonly handler!: pulumi.Output<string>;
    /**
     * Allows both HTTP and HTTPS (`enabled`) or redirect HTTP to HTTPS (`redirected`). Defaults to `enabled`.
     */
    public readonly httpOption!: pulumi.Output<string | undefined>;
    /**
     * The maximum number of instances this function can scale to. Default to 20. Your function will scale automatically based on the incoming workload, but will never exceed the configured `maxScale` value.
     */
    public readonly maxScale!: pulumi.Output<number | undefined>;
    /**
     * The memory resources in MB to allocate to each function. Defaults to 256 MB.
     */
    public readonly memoryLimit!: pulumi.Output<number | undefined>;
    /**
     * The minimum number of function instances running continuously. Defaults to 0. Functions are billed when executed, and using a `minScale` greater than 0 will cause your function to run constantly.
     */
    public readonly minScale!: pulumi.Output<number | undefined>;
    /**
     * The unique name of the function name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The Functions namespace ID of the function.
     *
     * > **Important** Updating the `name` argument will recreate the function.
     */
    public readonly namespaceId!: pulumi.Output<string>;
    /**
     * The organization ID the function is associated with.
     */
    public /*out*/ readonly organizationId!: pulumi.Output<string>;
    /**
     * The privacy type defines the way to authenticate to your function. Please check our dedicated [section](https://www.scaleway.com/en/developers/api/serverless-functions/#protocol-9dd4c8).
     */
    public readonly privacy!: pulumi.Output<string>;
    /**
     * The ID of the Private Network the function is connected to.
     *
     * > **Important** This feature is currently in beta and requires a namespace with VPC integration activated by setting the `activateVpcIntegration` attribute to `true`.
     */
    public readonly privateNetworkId!: pulumi.Output<string | undefined>;
    /**
     * `projectId`) The ID of the project the functions namespace is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * `region`). The region in which the namespace should be created.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Runtime of the function. Runtimes can be fetched using [specific route](https://www.scaleway.com/en/developers/api/serverless-functions/#path-functions-get-a-function)
     */
    public readonly runtime!: pulumi.Output<string>;
    /**
     * Execution environment of the function.
     */
    public readonly sandbox!: pulumi.Output<string>;
    /**
     * The [secret environment variables](https://www.scaleway.com/en/docs/compute/functions/concepts/#secrets) of the function.
     */
    public readonly secretEnvironmentVariables!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The list of tags associated with the function.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The maximum amount of time your function can spend processing a request before being stopped. Defaults to 300s.
     */
    public readonly timeout!: pulumi.Output<number>;
    /**
     * Path to the zip file containing your function sources to upload.
     */
    public readonly zipFile!: pulumi.Output<string | undefined>;
    /**
     * The hash of your source zip file, changing it will re-apply function. Can be any string
     */
    public readonly zipHash!: pulumi.Output<string | undefined>;

    /**
     * Create a Function resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FunctionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FunctionArgs | FunctionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FunctionState | undefined;
            resourceInputs["cpuLimit"] = state ? state.cpuLimit : undefined;
            resourceInputs["deploy"] = state ? state.deploy : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["environmentVariables"] = state ? state.environmentVariables : undefined;
            resourceInputs["handler"] = state ? state.handler : undefined;
            resourceInputs["httpOption"] = state ? state.httpOption : undefined;
            resourceInputs["maxScale"] = state ? state.maxScale : undefined;
            resourceInputs["memoryLimit"] = state ? state.memoryLimit : undefined;
            resourceInputs["minScale"] = state ? state.minScale : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespaceId"] = state ? state.namespaceId : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["privacy"] = state ? state.privacy : undefined;
            resourceInputs["privateNetworkId"] = state ? state.privateNetworkId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["runtime"] = state ? state.runtime : undefined;
            resourceInputs["sandbox"] = state ? state.sandbox : undefined;
            resourceInputs["secretEnvironmentVariables"] = state ? state.secretEnvironmentVariables : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
            resourceInputs["zipFile"] = state ? state.zipFile : undefined;
            resourceInputs["zipHash"] = state ? state.zipHash : undefined;
        } else {
            const args = argsOrState as FunctionArgs | undefined;
            if ((!args || args.handler === undefined) && !opts.urn) {
                throw new Error("Missing required property 'handler'");
            }
            if ((!args || args.namespaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'namespaceId'");
            }
            if ((!args || args.privacy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privacy'");
            }
            if ((!args || args.runtime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'runtime'");
            }
            resourceInputs["deploy"] = args ? args.deploy : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["environmentVariables"] = args ? args.environmentVariables : undefined;
            resourceInputs["handler"] = args ? args.handler : undefined;
            resourceInputs["httpOption"] = args ? args.httpOption : undefined;
            resourceInputs["maxScale"] = args ? args.maxScale : undefined;
            resourceInputs["memoryLimit"] = args ? args.memoryLimit : undefined;
            resourceInputs["minScale"] = args ? args.minScale : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespaceId"] = args ? args.namespaceId : undefined;
            resourceInputs["privacy"] = args ? args.privacy : undefined;
            resourceInputs["privateNetworkId"] = args ? args.privateNetworkId : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["runtime"] = args ? args.runtime : undefined;
            resourceInputs["sandbox"] = args ? args.sandbox : undefined;
            resourceInputs["secretEnvironmentVariables"] = args?.secretEnvironmentVariables ? pulumi.secret(args.secretEnvironmentVariables) : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["zipFile"] = args ? args.zipFile : undefined;
            resourceInputs["zipHash"] = args ? args.zipHash : undefined;
            resourceInputs["cpuLimit"] = undefined /*out*/;
            resourceInputs["domainName"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "scaleway:index/function:Function" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        const secretOpts = { additionalSecretOutputs: ["secretEnvironmentVariables"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Function.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Function resources.
 */
export interface FunctionState {
    /**
     * The CPU limit in mVCPU for your function.
     */
    cpuLimit?: pulumi.Input<number>;
    deploy?: pulumi.Input<boolean>;
    /**
     * The description of the function.
     */
    description?: pulumi.Input<string>;
    /**
     * The native domain name of the function.
     */
    domainName?: pulumi.Input<string>;
    /**
     * The [environment variables](https://www.scaleway.com/en/docs/compute/functions/concepts/#environment-variables) of the function.
     */
    environmentVariables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Handler of the function, depends on the runtime. Refer to the [dedicated documentation](https://www.scaleway.com/en/developers/api/serverless-functions/#path-functions-create-a-new-function) for the list of supported runtimes.
     */
    handler?: pulumi.Input<string>;
    /**
     * Allows both HTTP and HTTPS (`enabled`) or redirect HTTP to HTTPS (`redirected`). Defaults to `enabled`.
     */
    httpOption?: pulumi.Input<string>;
    /**
     * The maximum number of instances this function can scale to. Default to 20. Your function will scale automatically based on the incoming workload, but will never exceed the configured `maxScale` value.
     */
    maxScale?: pulumi.Input<number>;
    /**
     * The memory resources in MB to allocate to each function. Defaults to 256 MB.
     */
    memoryLimit?: pulumi.Input<number>;
    /**
     * The minimum number of function instances running continuously. Defaults to 0. Functions are billed when executed, and using a `minScale` greater than 0 will cause your function to run constantly.
     */
    minScale?: pulumi.Input<number>;
    /**
     * The unique name of the function name.
     */
    name?: pulumi.Input<string>;
    /**
     * The Functions namespace ID of the function.
     *
     * > **Important** Updating the `name` argument will recreate the function.
     */
    namespaceId?: pulumi.Input<string>;
    /**
     * The organization ID the function is associated with.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * The privacy type defines the way to authenticate to your function. Please check our dedicated [section](https://www.scaleway.com/en/developers/api/serverless-functions/#protocol-9dd4c8).
     */
    privacy?: pulumi.Input<string>;
    /**
     * The ID of the Private Network the function is connected to.
     *
     * > **Important** This feature is currently in beta and requires a namespace with VPC integration activated by setting the `activateVpcIntegration` attribute to `true`.
     */
    privateNetworkId?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the functions namespace is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`). The region in which the namespace should be created.
     */
    region?: pulumi.Input<string>;
    /**
     * Runtime of the function. Runtimes can be fetched using [specific route](https://www.scaleway.com/en/developers/api/serverless-functions/#path-functions-get-a-function)
     */
    runtime?: pulumi.Input<string>;
    /**
     * Execution environment of the function.
     */
    sandbox?: pulumi.Input<string>;
    /**
     * The [secret environment variables](https://www.scaleway.com/en/docs/compute/functions/concepts/#secrets) of the function.
     */
    secretEnvironmentVariables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The list of tags associated with the function.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The maximum amount of time your function can spend processing a request before being stopped. Defaults to 300s.
     */
    timeout?: pulumi.Input<number>;
    /**
     * Path to the zip file containing your function sources to upload.
     */
    zipFile?: pulumi.Input<string>;
    /**
     * The hash of your source zip file, changing it will re-apply function. Can be any string
     */
    zipHash?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Function resource.
 */
export interface FunctionArgs {
    deploy?: pulumi.Input<boolean>;
    /**
     * The description of the function.
     */
    description?: pulumi.Input<string>;
    /**
     * The [environment variables](https://www.scaleway.com/en/docs/compute/functions/concepts/#environment-variables) of the function.
     */
    environmentVariables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Handler of the function, depends on the runtime. Refer to the [dedicated documentation](https://www.scaleway.com/en/developers/api/serverless-functions/#path-functions-create-a-new-function) for the list of supported runtimes.
     */
    handler: pulumi.Input<string>;
    /**
     * Allows both HTTP and HTTPS (`enabled`) or redirect HTTP to HTTPS (`redirected`). Defaults to `enabled`.
     */
    httpOption?: pulumi.Input<string>;
    /**
     * The maximum number of instances this function can scale to. Default to 20. Your function will scale automatically based on the incoming workload, but will never exceed the configured `maxScale` value.
     */
    maxScale?: pulumi.Input<number>;
    /**
     * The memory resources in MB to allocate to each function. Defaults to 256 MB.
     */
    memoryLimit?: pulumi.Input<number>;
    /**
     * The minimum number of function instances running continuously. Defaults to 0. Functions are billed when executed, and using a `minScale` greater than 0 will cause your function to run constantly.
     */
    minScale?: pulumi.Input<number>;
    /**
     * The unique name of the function name.
     */
    name?: pulumi.Input<string>;
    /**
     * The Functions namespace ID of the function.
     *
     * > **Important** Updating the `name` argument will recreate the function.
     */
    namespaceId: pulumi.Input<string>;
    /**
     * The privacy type defines the way to authenticate to your function. Please check our dedicated [section](https://www.scaleway.com/en/developers/api/serverless-functions/#protocol-9dd4c8).
     */
    privacy: pulumi.Input<string>;
    /**
     * The ID of the Private Network the function is connected to.
     *
     * > **Important** This feature is currently in beta and requires a namespace with VPC integration activated by setting the `activateVpcIntegration` attribute to `true`.
     */
    privateNetworkId?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the functions namespace is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`). The region in which the namespace should be created.
     */
    region?: pulumi.Input<string>;
    /**
     * Runtime of the function. Runtimes can be fetched using [specific route](https://www.scaleway.com/en/developers/api/serverless-functions/#path-functions-get-a-function)
     */
    runtime: pulumi.Input<string>;
    /**
     * Execution environment of the function.
     */
    sandbox?: pulumi.Input<string>;
    /**
     * The [secret environment variables](https://www.scaleway.com/en/docs/compute/functions/concepts/#secrets) of the function.
     */
    secretEnvironmentVariables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The list of tags associated with the function.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The maximum amount of time your function can spend processing a request before being stopped. Defaults to 300s.
     */
    timeout?: pulumi.Input<number>;
    /**
     * Path to the zip file containing your function sources to upload.
     */
    zipFile?: pulumi.Input<string>;
    /**
     * The hash of your source zip file, changing it will re-apply function. Can be any string
     */
    zipHash?: pulumi.Input<string>;
}
