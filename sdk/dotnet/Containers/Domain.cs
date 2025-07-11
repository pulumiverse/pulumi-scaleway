// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Containers
{
    /// <summary>
    /// The `scaleway.containers.Domain` resource allows you to create and manage domain name bindings for Scaleway [Serverless Containers](https://www.scaleway.com/en/docs/serverless/containers/).
    /// 
    /// Refer to the Containers domain [documentation](https://www.scaleway.com/en/docs/serverless-containers/how-to/add-a-custom-domain-to-a-container/) and the [API documentation](https://www.scaleway.com/en/developers/api/serverless-containers/#path-domains-list-all-domain-name-bindings) for more information.
    /// 
    /// ## Example Usage
    /// 
    /// The commands below shows how to bind a custom domain name to a container.
    /// 
    /// ### Simple
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var app = new Scaleway.Containers.Container("app");
    /// 
    ///     var appDomain = new Scaleway.Containers.Domain("app", new()
    ///     {
    ///         ContainerId = app.Id,
    ///         Hostname = "container.domain.tld",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Complete example with domain
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Scaleway.Containers.Namespace("main", new()
    ///     {
    ///         Name = "my-ns-test",
    ///         Description = "test container",
    ///     });
    /// 
    ///     var app = new Scaleway.Containers.Container("app", new()
    ///     {
    ///         Name = "app",
    ///         NamespaceId = main.Id,
    ///         RegistryImage = main.RegistryEndpoint.Apply(registryEndpoint =&gt; $"{registryEndpoint}/nginx:alpine"),
    ///         Port = 80,
    ///         CpuLimit = 140,
    ///         MemoryLimit = 256,
    ///         MinScale = 1,
    ///         MaxScale = 1,
    ///         Timeout = 600,
    ///         MaxConcurrency = 80,
    ///         Privacy = "public",
    ///         Protocol = "http1",
    ///         Deploy = true,
    ///     });
    /// 
    ///     var appRecord = new Scaleway.Domain.Record("app", new()
    ///     {
    ///         DnsZone = "domain.tld",
    ///         Name = "subdomain",
    ///         Type = "CNAME",
    ///         Data = app.DomainName.Apply(domainName =&gt; $"{domainName}."),
    ///         Ttl = 3600,
    ///     });
    /// 
    ///     var appDomain = new Scaleway.Containers.Domain("app", new()
    ///     {
    ///         ContainerId = app.Id,
    ///         Hostname = Output.Tuple(appRecord.Name, appRecord.DnsZone).Apply(values =&gt;
    ///         {
    ///             var name = values.Item1;
    ///             var dnsZone = values.Item2;
    ///             return $"{name}.{dnsZone}";
    ///         }),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Container domain binding can be imported using `{region}/{id}`, as shown below:
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:containers/domain:Domain main fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:containers/domain:Domain")]
    public partial class Domain : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The unique identifier of the container.
        /// </summary>
        [Output("containerId")]
        public Output<string> ContainerId { get; private set; } = null!;

        /// <summary>
        /// The hostname with a CNAME record.
        /// </summary>
        [Output("hostname")]
        public Output<string> Hostname { get; private set; } = null!;

        /// <summary>
        /// `region`) The region in which the container exists.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The URL used to query the container.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a Domain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Domain(string name, DomainArgs args, CustomResourceOptions? options = null)
            : base("scaleway:containers/domain:Domain", name, args ?? new DomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Domain(string name, Input<string> id, DomainState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:containers/domain:Domain", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "scaleway:index/containerDomain:ContainerDomain" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Domain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Domain Get(string name, Input<string> id, DomainState? state = null, CustomResourceOptions? options = null)
        {
            return new Domain(name, id, state, options);
        }
    }

    public sealed class DomainArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique identifier of the container.
        /// </summary>
        [Input("containerId", required: true)]
        public Input<string> ContainerId { get; set; } = null!;

        /// <summary>
        /// The hostname with a CNAME record.
        /// </summary>
        [Input("hostname", required: true)]
        public Input<string> Hostname { get; set; } = null!;

        /// <summary>
        /// `region`) The region in which the container exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public DomainArgs()
        {
        }
        public static new DomainArgs Empty => new DomainArgs();
    }

    public sealed class DomainState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique identifier of the container.
        /// </summary>
        [Input("containerId")]
        public Input<string>? ContainerId { get; set; }

        /// <summary>
        /// The hostname with a CNAME record.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// `region`) The region in which the container exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The URL used to query the container.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public DomainState()
        {
        }
        public static new DomainState Empty => new DomainState();
    }
}
