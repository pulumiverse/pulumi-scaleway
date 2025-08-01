// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Ipam
{
    public static class GetIps
    {
        /// <summary>
        /// Gets information about multiple IP addresses managed by Scaleway's IP Address Management (IPAM) service.
        /// 
        /// For more information about IPAM, see the main [documentation](https://www.scaleway.com/en/docs/vpc/concepts/#ipam).
        /// 
        /// ## Examples
        /// 
        /// ### By tag
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var byTag = Scaleway.Ipam.GetIps.Invoke(new()
        ///     {
        ///         Tags = new[]
        ///         {
        ///             "tag",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ### By type and resource
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
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
        ///     var redis01 = new Scaleway.Redis.Cluster("redis01", new()
        ///     {
        ///         Name = "my_redis_cluster",
        ///         Version = "7.0.5",
        ///         NodeType = "RED1-XS",
        ///         UserName = "my_initial_user",
        ///         Password = "thiZ_is_v&amp;ry_s3cret",
        ///         ClusterSize = 3,
        ///         PrivateNetworks = new[]
        ///         {
        ///             new Scaleway.Redis.Inputs.ClusterPrivateNetworkArgs
        ///             {
        ///                 Id = pn01.Id,
        ///             },
        ///         },
        ///     });
        /// 
        ///     var byTypeAndResource = Scaleway.Ipam.GetIps.Invoke(new()
        ///     {
        ///         Type = "ipv4",
        ///         Resource = new Scaleway.Ipam.Inputs.GetIpsResourceInputArgs
        ///         {
        ///             Id = redis01.Id,
        ///             Type = "redis_cluster",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetIpsResult> InvokeAsync(GetIpsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIpsResult>("scaleway:ipam/getIps:getIps", args ?? new GetIpsArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about multiple IP addresses managed by Scaleway's IP Address Management (IPAM) service.
        /// 
        /// For more information about IPAM, see the main [documentation](https://www.scaleway.com/en/docs/vpc/concepts/#ipam).
        /// 
        /// ## Examples
        /// 
        /// ### By tag
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var byTag = Scaleway.Ipam.GetIps.Invoke(new()
        ///     {
        ///         Tags = new[]
        ///         {
        ///             "tag",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ### By type and resource
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
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
        ///     var redis01 = new Scaleway.Redis.Cluster("redis01", new()
        ///     {
        ///         Name = "my_redis_cluster",
        ///         Version = "7.0.5",
        ///         NodeType = "RED1-XS",
        ///         UserName = "my_initial_user",
        ///         Password = "thiZ_is_v&amp;ry_s3cret",
        ///         ClusterSize = 3,
        ///         PrivateNetworks = new[]
        ///         {
        ///             new Scaleway.Redis.Inputs.ClusterPrivateNetworkArgs
        ///             {
        ///                 Id = pn01.Id,
        ///             },
        ///         },
        ///     });
        /// 
        ///     var byTypeAndResource = Scaleway.Ipam.GetIps.Invoke(new()
        ///     {
        ///         Type = "ipv4",
        ///         Resource = new Scaleway.Ipam.Inputs.GetIpsResourceInputArgs
        ///         {
        ///             Id = redis01.Id,
        ///             Type = "redis_cluster",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetIpsResult> Invoke(GetIpsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIpsResult>("scaleway:ipam/getIps:getIps", args ?? new GetIpsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about multiple IP addresses managed by Scaleway's IP Address Management (IPAM) service.
        /// 
        /// For more information about IPAM, see the main [documentation](https://www.scaleway.com/en/docs/vpc/concepts/#ipam).
        /// 
        /// ## Examples
        /// 
        /// ### By tag
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var byTag = Scaleway.Ipam.GetIps.Invoke(new()
        ///     {
        ///         Tags = new[]
        ///         {
        ///             "tag",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ### By type and resource
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
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
        ///     var redis01 = new Scaleway.Redis.Cluster("redis01", new()
        ///     {
        ///         Name = "my_redis_cluster",
        ///         Version = "7.0.5",
        ///         NodeType = "RED1-XS",
        ///         UserName = "my_initial_user",
        ///         Password = "thiZ_is_v&amp;ry_s3cret",
        ///         ClusterSize = 3,
        ///         PrivateNetworks = new[]
        ///         {
        ///             new Scaleway.Redis.Inputs.ClusterPrivateNetworkArgs
        ///             {
        ///                 Id = pn01.Id,
        ///             },
        ///         },
        ///     });
        /// 
        ///     var byTypeAndResource = Scaleway.Ipam.GetIps.Invoke(new()
        ///     {
        ///         Type = "ipv4",
        ///         Resource = new Scaleway.Ipam.Inputs.GetIpsResourceInputArgs
        ///         {
        ///             Id = redis01.Id,
        ///             Type = "redis_cluster",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetIpsResult> Invoke(GetIpsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetIpsResult>("scaleway:ipam/getIps:getIps", args ?? new GetIpsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIpsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Defines whether to filter only for IPs which are attached to a resource.
        /// </summary>
        [Input("attached")]
        public bool? Attached { get; set; }

        /// <summary>
        /// The linked MAC address to filter for.
        /// </summary>
        [Input("macAddress")]
        public string? MacAddress { get; set; }

        /// <summary>
        /// The ID of the Private Network to filter for.
        /// </summary>
        [Input("privateNetworkId")]
        public string? PrivateNetworkId { get; set; }

        /// <summary>
        /// The ID of the Project to filter for.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// The region to filter for.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// Filter for a resource attached to the IP, using resource ID, type or name.
        /// </summary>
        [Input("resource")]
        public Inputs.GetIpsResourceArgs? Resource { get; set; }

        [Input("tags")]
        private List<string>? _tags;

        /// <summary>
        /// The IP tags to filter for.
        /// </summary>
        public List<string> Tags
        {
            get => _tags ?? (_tags = new List<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of IP to filter for (`ipv4` or `ipv6`).
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        /// <summary>
        /// Only IPs that are zonal, and in this zone, will be returned.
        /// </summary>
        [Input("zonal")]
        public string? Zonal { get; set; }

        public GetIpsArgs()
        {
        }
        public static new GetIpsArgs Empty => new GetIpsArgs();
    }

    public sealed class GetIpsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Defines whether to filter only for IPs which are attached to a resource.
        /// </summary>
        [Input("attached")]
        public Input<bool>? Attached { get; set; }

        /// <summary>
        /// The linked MAC address to filter for.
        /// </summary>
        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        /// <summary>
        /// The ID of the Private Network to filter for.
        /// </summary>
        [Input("privateNetworkId")]
        public Input<string>? PrivateNetworkId { get; set; }

        /// <summary>
        /// The ID of the Project to filter for.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region to filter for.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Filter for a resource attached to the IP, using resource ID, type or name.
        /// </summary>
        [Input("resource")]
        public Input<Inputs.GetIpsResourceInputArgs>? Resource { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The IP tags to filter for.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of IP to filter for (`ipv4` or `ipv6`).
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Only IPs that are zonal, and in this zone, will be returned.
        /// </summary>
        [Input("zonal")]
        public Input<string>? Zonal { get; set; }

        public GetIpsInvokeArgs()
        {
        }
        public static new GetIpsInvokeArgs Empty => new GetIpsInvokeArgs();
    }


    [OutputType]
    public sealed class GetIpsResult
    {
        public readonly bool? Attached;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of found IPs.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIpsIpResult> Ips;
        /// <summary>
        /// The associated MAC address.
        /// </summary>
        public readonly string? MacAddress;
        public readonly string OrganizationId;
        public readonly string? PrivateNetworkId;
        /// <summary>
        /// The ID of the Project the resource is associated with.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// The region of the IP.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The list of public IPs attached to the resource.
        /// </summary>
        public readonly Outputs.GetIpsResourceResult? Resource;
        /// <summary>
        /// The tags associated with the IP.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The type of resource.
        /// </summary>
        public readonly string? Type;
        public readonly string Zonal;

        [OutputConstructor]
        private GetIpsResult(
            bool? attached,

            string id,

            ImmutableArray<Outputs.GetIpsIpResult> ips,

            string? macAddress,

            string organizationId,

            string? privateNetworkId,

            string projectId,

            string region,

            Outputs.GetIpsResourceResult? resource,

            ImmutableArray<string> tags,

            string? type,

            string zonal)
        {
            Attached = attached;
            Id = id;
            Ips = ips;
            MacAddress = macAddress;
            OrganizationId = organizationId;
            PrivateNetworkId = privateNetworkId;
            ProjectId = projectId;
            Region = region;
            Resource = resource;
            Tags = tags;
            Type = type;
            Zonal = zonal;
        }
    }
}
