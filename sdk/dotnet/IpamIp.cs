// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway
{
    /// <summary>
    /// Books and manages IPAM IPs.
    /// 
    /// For more information about IPAM, see the main [documentation](https://www.scaleway.com/en/docs/vpc/concepts/#ipam).
    /// 
    /// ## Example Usage
    /// 
    /// ### Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var vpc01 = new Scaleway.Network.Vpc("vpc01", new()
    ///     {
    ///         Name = "my vpc",
    ///     });
    /// 
    ///     var pn01 = new Scaleway.Network.PrivateNetwork("pn01", new()
    ///     {
    ///         VpcId = vpc01.Id,
    ///         Ipv4Subnet = new Scaleway.Network.Inputs.PrivateNetworkIpv4SubnetArgs
    ///         {
    ///             Subnet = "172.16.32.0/22",
    ///         },
    ///     });
    /// 
    ///     var ip01 = new Scaleway.Ipam.Ip("ip01", new()
    ///     {
    ///         Sources = new[]
    ///         {
    ///             new Scaleway.Ipam.Inputs.IpSourceArgs
    ///             {
    ///                 PrivateNetworkId = pn01.Id,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Request a specific IPv4 address
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var vpc01 = new Scaleway.Network.Vpc("vpc01", new()
    ///     {
    ///         Name = "my vpc",
    ///     });
    /// 
    ///     var pn01 = new Scaleway.Network.PrivateNetwork("pn01", new()
    ///     {
    ///         VpcId = vpc01.Id,
    ///         Ipv4Subnet = new Scaleway.Network.Inputs.PrivateNetworkIpv4SubnetArgs
    ///         {
    ///             Subnet = "172.16.32.0/22",
    ///         },
    ///     });
    /// 
    ///     var ip01 = new Scaleway.Ipam.Ip("ip01", new()
    ///     {
    ///         Address = "172.16.32.7",
    ///         Sources = new[]
    ///         {
    ///             new Scaleway.Ipam.Inputs.IpSourceArgs
    ///             {
    ///                 PrivateNetworkId = pn01.Id,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Request an IPv6 address
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var vpc01 = new Scaleway.Network.Vpc("vpc01", new()
    ///     {
    ///         Name = "my vpc",
    ///     });
    /// 
    ///     var pn01 = new Scaleway.Network.PrivateNetwork("pn01", new()
    ///     {
    ///         VpcId = vpc01.Id,
    ///         Ipv6Subnets = new[]
    ///         {
    ///             new Scaleway.Network.Inputs.PrivateNetworkIpv6SubnetArgs
    ///             {
    ///                 Subnet = "fd46:78ab:30b8:177c::/64",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var ip01 = new Scaleway.Ipam.Ip("ip01", new()
    ///     {
    ///         IsIpv6 = true,
    ///         Sources = new[]
    ///         {
    ///             new Scaleway.Ipam.Inputs.IpSourceArgs
    ///             {
    ///                 PrivateNetworkId = pn01.Id,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Book an IP for a custom resource
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var vpc01 = new Scaleway.Network.Vpc("vpc01", new()
    ///     {
    ///         Name = "my vpc",
    ///     });
    /// 
    ///     var pn01 = new Scaleway.Network.PrivateNetwork("pn01", new()
    ///     {
    ///         VpcId = vpc01.Id,
    ///         Ipv4Subnet = new Scaleway.Network.Inputs.PrivateNetworkIpv4SubnetArgs
    ///         {
    ///             Subnet = "172.16.32.0/22",
    ///         },
    ///     });
    /// 
    ///     var ip01 = new Scaleway.Ipam.Ip("ip01", new()
    ///     {
    ///         Address = "172.16.32.7",
    ///         Sources = new[]
    ///         {
    ///             new Scaleway.Ipam.Inputs.IpSourceArgs
    ///             {
    ///                 PrivateNetworkId = pn01.Id,
    ///             },
    ///         },
    ///         CustomResources = new[]
    ///         {
    ///             new Scaleway.Ipam.Inputs.IpCustomResourceArgs
    ///             {
    ///                 MacAddress = "bc:24:11:74:d0:6a",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// IPAM IPs can be imported using `{region}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/ipamIp:IpamIp ip_demo fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [Obsolete(@"scaleway.index/ipamip.IpamIp has been deprecated in favor of scaleway.ipam/ip.Ip")]
    [ScalewayResourceType("scaleway:index/ipamIp:IpamIp")]
    public partial class IpamIp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Request a specific IP in the requested source pool
        /// </summary>
        [Output("address")]
        public Output<string> Address { get; private set; } = null!;

        /// <summary>
        /// Date and time of IP's creation (RFC 3339 format).
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The custom resource in which to book the IP
        /// </summary>
        [Output("customResources")]
        public Output<ImmutableArray<Outputs.IpamIpCustomResource>> CustomResources { get; private set; } = null!;

        /// <summary>
        /// Defines whether to request an IPv6 address instead of IPv4.
        /// </summary>
        [Output("isIpv6")]
        public Output<bool?> IsIpv6 { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the Project the IP is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// `region`) The region of the IP.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The IP resource.
        /// </summary>
        [Output("resources")]
        public Output<ImmutableArray<Outputs.IpamIpResource>> Resources { get; private set; } = null!;

        /// <summary>
        /// The reverse DNS for this IP.
        /// </summary>
        [Output("reverses")]
        public Output<ImmutableArray<Outputs.IpamIpReverse>> Reverses { get; private set; } = null!;

        /// <summary>
        /// The source in which to book the IP.
        /// </summary>
        [Output("sources")]
        public Output<ImmutableArray<Outputs.IpamIpSource>> Sources { get; private set; } = null!;

        /// <summary>
        /// The tags associated with the IP.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Date and time of IP's last update (RFC 3339 format).
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// The zone of the IP.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a IpamIp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IpamIp(string name, IpamIpArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/ipamIp:IpamIp", name, args ?? new IpamIpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IpamIp(string name, Input<string> id, IpamIpState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/ipamIp:IpamIp", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IpamIp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IpamIp Get(string name, Input<string> id, IpamIpState? state = null, CustomResourceOptions? options = null)
        {
            return new IpamIp(name, id, state, options);
        }
    }

    public sealed class IpamIpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Request a specific IP in the requested source pool
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        [Input("customResources")]
        private InputList<Inputs.IpamIpCustomResourceArgs>? _customResources;

        /// <summary>
        /// The custom resource in which to book the IP
        /// </summary>
        public InputList<Inputs.IpamIpCustomResourceArgs> CustomResources
        {
            get => _customResources ?? (_customResources = new InputList<Inputs.IpamIpCustomResourceArgs>());
            set => _customResources = value;
        }

        /// <summary>
        /// Defines whether to request an IPv6 address instead of IPv4.
        /// </summary>
        [Input("isIpv6")]
        public Input<bool>? IsIpv6 { get; set; }

        /// <summary>
        /// `project_id`) The ID of the Project the IP is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region of the IP.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("sources", required: true)]
        private InputList<Inputs.IpamIpSourceArgs>? _sources;

        /// <summary>
        /// The source in which to book the IP.
        /// </summary>
        public InputList<Inputs.IpamIpSourceArgs> Sources
        {
            get => _sources ?? (_sources = new InputList<Inputs.IpamIpSourceArgs>());
            set => _sources = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the IP.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public IpamIpArgs()
        {
        }
        public static new IpamIpArgs Empty => new IpamIpArgs();
    }

    public sealed class IpamIpState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Request a specific IP in the requested source pool
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// Date and time of IP's creation (RFC 3339 format).
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("customResources")]
        private InputList<Inputs.IpamIpCustomResourceGetArgs>? _customResources;

        /// <summary>
        /// The custom resource in which to book the IP
        /// </summary>
        public InputList<Inputs.IpamIpCustomResourceGetArgs> CustomResources
        {
            get => _customResources ?? (_customResources = new InputList<Inputs.IpamIpCustomResourceGetArgs>());
            set => _customResources = value;
        }

        /// <summary>
        /// Defines whether to request an IPv6 address instead of IPv4.
        /// </summary>
        [Input("isIpv6")]
        public Input<bool>? IsIpv6 { get; set; }

        /// <summary>
        /// `project_id`) The ID of the Project the IP is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region of the IP.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("resources")]
        private InputList<Inputs.IpamIpResourceGetArgs>? _resources;

        /// <summary>
        /// The IP resource.
        /// </summary>
        public InputList<Inputs.IpamIpResourceGetArgs> Resources
        {
            get => _resources ?? (_resources = new InputList<Inputs.IpamIpResourceGetArgs>());
            set => _resources = value;
        }

        [Input("reverses")]
        private InputList<Inputs.IpamIpReverseGetArgs>? _reverses;

        /// <summary>
        /// The reverse DNS for this IP.
        /// </summary>
        public InputList<Inputs.IpamIpReverseGetArgs> Reverses
        {
            get => _reverses ?? (_reverses = new InputList<Inputs.IpamIpReverseGetArgs>());
            set => _reverses = value;
        }

        [Input("sources")]
        private InputList<Inputs.IpamIpSourceGetArgs>? _sources;

        /// <summary>
        /// The source in which to book the IP.
        /// </summary>
        public InputList<Inputs.IpamIpSourceGetArgs> Sources
        {
            get => _sources ?? (_sources = new InputList<Inputs.IpamIpSourceGetArgs>());
            set => _sources = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the IP.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Date and time of IP's last update (RFC 3339 format).
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// The zone of the IP.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public IpamIpState()
        {
        }
        public static new IpamIpState Empty => new IpamIpState();
    }
}
