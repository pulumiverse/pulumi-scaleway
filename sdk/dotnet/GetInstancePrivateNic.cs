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
    [Obsolete(@"scaleway.index/getinstanceprivatenic.getInstancePrivateNic has been deprecated in favor of scaleway.instance/getprivatenic.getPrivateNic")]
    public static class GetInstancePrivateNic
    {
        /// <summary>
        /// Gets information about an instance private NIC.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var byNicId = Scaleway.Instance.GetPrivateNic.Invoke(new()
        ///     {
        ///         ServerId = "11111111-1111-1111-1111-111111111111",
        ///         PrivateNicId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        ///     var byPnId = Scaleway.Instance.GetPrivateNic.Invoke(new()
        ///     {
        ///         ServerId = "11111111-1111-1111-1111-111111111111",
        ///         PrivateNetworkId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        ///     var byTags = Scaleway.Instance.GetPrivateNic.Invoke(new()
        ///     {
        ///         ServerId = "11111111-1111-1111-1111-111111111111",
        ///         Tags = new[]
        ///         {
        ///             "mytag",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetInstancePrivateNicResult> InvokeAsync(GetInstancePrivateNicArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstancePrivateNicResult>("scaleway:index/getInstancePrivateNic:getInstancePrivateNic", args ?? new GetInstancePrivateNicArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about an instance private NIC.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var byNicId = Scaleway.Instance.GetPrivateNic.Invoke(new()
        ///     {
        ///         ServerId = "11111111-1111-1111-1111-111111111111",
        ///         PrivateNicId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        ///     var byPnId = Scaleway.Instance.GetPrivateNic.Invoke(new()
        ///     {
        ///         ServerId = "11111111-1111-1111-1111-111111111111",
        ///         PrivateNetworkId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        ///     var byTags = Scaleway.Instance.GetPrivateNic.Invoke(new()
        ///     {
        ///         ServerId = "11111111-1111-1111-1111-111111111111",
        ///         Tags = new[]
        ///         {
        ///             "mytag",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetInstancePrivateNicResult> Invoke(GetInstancePrivateNicInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstancePrivateNicResult>("scaleway:index/getInstancePrivateNic:getInstancePrivateNic", args ?? new GetInstancePrivateNicInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about an instance private NIC.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var byNicId = Scaleway.Instance.GetPrivateNic.Invoke(new()
        ///     {
        ///         ServerId = "11111111-1111-1111-1111-111111111111",
        ///         PrivateNicId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        ///     var byPnId = Scaleway.Instance.GetPrivateNic.Invoke(new()
        ///     {
        ///         ServerId = "11111111-1111-1111-1111-111111111111",
        ///         PrivateNetworkId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        ///     var byTags = Scaleway.Instance.GetPrivateNic.Invoke(new()
        ///     {
        ///         ServerId = "11111111-1111-1111-1111-111111111111",
        ///         Tags = new[]
        ///         {
        ///             "mytag",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetInstancePrivateNicResult> Invoke(GetInstancePrivateNicInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstancePrivateNicResult>("scaleway:index/getInstancePrivateNic:getInstancePrivateNic", args ?? new GetInstancePrivateNicInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstancePrivateNicArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the private network
        /// Only one of `private_nic_id` and `private_network_id` should be specified.
        /// </summary>
        [Input("privateNetworkId")]
        public string? PrivateNetworkId { get; set; }

        /// <summary>
        /// The ID of the instance server private nic
        /// Only one of `private_nic_id` and `private_network_id` should be specified.
        /// </summary>
        [Input("privateNicId")]
        public string? PrivateNicId { get; set; }

        /// <summary>
        /// The server's id
        /// </summary>
        [Input("serverId", required: true)]
        public string ServerId { get; set; } = null!;

        [Input("tags")]
        private List<string>? _tags;

        /// <summary>
        /// The tags associated with the private NIC.
        /// As datasource only returns one private NIC, the search with given tags must return only one result
        /// </summary>
        public List<string> Tags
        {
            get => _tags ?? (_tags = new List<string>());
            set => _tags = value;
        }

        /// <summary>
        /// `zone`) The zone in which the private nic exists.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetInstancePrivateNicArgs()
        {
        }
        public static new GetInstancePrivateNicArgs Empty => new GetInstancePrivateNicArgs();
    }

    public sealed class GetInstancePrivateNicInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the private network
        /// Only one of `private_nic_id` and `private_network_id` should be specified.
        /// </summary>
        [Input("privateNetworkId")]
        public Input<string>? PrivateNetworkId { get; set; }

        /// <summary>
        /// The ID of the instance server private nic
        /// Only one of `private_nic_id` and `private_network_id` should be specified.
        /// </summary>
        [Input("privateNicId")]
        public Input<string>? PrivateNicId { get; set; }

        /// <summary>
        /// The server's id
        /// </summary>
        [Input("serverId", required: true)]
        public Input<string> ServerId { get; set; } = null!;

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the private NIC.
        /// As datasource only returns one private NIC, the search with given tags must return only one result
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// `zone`) The zone in which the private nic exists.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetInstancePrivateNicInvokeArgs()
        {
        }
        public static new GetInstancePrivateNicInvokeArgs Empty => new GetInstancePrivateNicInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstancePrivateNicResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> IpIds;
        public readonly ImmutableArray<string> IpamIpIds;
        public readonly string MacAddress;
        public readonly ImmutableArray<Outputs.GetInstancePrivateNicPrivateIpResult> PrivateIps;
        public readonly string? PrivateNetworkId;
        public readonly string? PrivateNicId;
        public readonly string ServerId;
        public readonly ImmutableArray<string> Tags;
        public readonly string? Zone;

        [OutputConstructor]
        private GetInstancePrivateNicResult(
            string id,

            ImmutableArray<string> ipIds,

            ImmutableArray<string> ipamIpIds,

            string macAddress,

            ImmutableArray<Outputs.GetInstancePrivateNicPrivateIpResult> privateIps,

            string? privateNetworkId,

            string? privateNicId,

            string serverId,

            ImmutableArray<string> tags,

            string? zone)
        {
            Id = id;
            IpIds = ipIds;
            IpamIpIds = ipamIpIds;
            MacAddress = macAddress;
            PrivateIps = privateIps;
            PrivateNetworkId = privateNetworkId;
            PrivateNicId = privateNicId;
            ServerId = serverId;
            Tags = tags;
            Zone = zone;
        }
    }
}
