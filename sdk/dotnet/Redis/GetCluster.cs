// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Redis
{
    public static class GetCluster
    {
        /// <summary>
        /// Gets information about a Redis™ cluster.
        /// 
        /// For further information refer to the Managed Database for Redis™ [API documentation](https://developers.scaleway.com/en/products/redis/api/v1alpha1/#clusters-a85816).
        /// </summary>
        public static Task<GetClusterResult> InvokeAsync(GetClusterArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("scaleway:redis/getCluster:getCluster", args ?? new GetClusterArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a Redis™ cluster.
        /// 
        /// For further information refer to the Managed Database for Redis™ [API documentation](https://developers.scaleway.com/en/products/redis/api/v1alpha1/#clusters-a85816).
        /// </summary>
        public static Output<GetClusterResult> Invoke(GetClusterInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetClusterResult>("scaleway:redis/getCluster:getCluster", args ?? new GetClusterInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a Redis™ cluster.
        /// 
        /// For further information refer to the Managed Database for Redis™ [API documentation](https://developers.scaleway.com/en/products/redis/api/v1alpha1/#clusters-a85816).
        /// </summary>
        public static Output<GetClusterResult> Invoke(GetClusterInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetClusterResult>("scaleway:redis/getCluster:getCluster", args ?? new GetClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClusterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Redis cluster ID.
        /// 
        /// &gt; **Note** You must specify at least one: `name` and/or `cluster_id`.
        /// </summary>
        [Input("clusterId")]
        public string? ClusterId { get; set; }

        /// <summary>
        /// The name of the Redis cluster.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of the project the Redis cluster is associated with.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// `region`) The zone in which the server exists.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetClusterArgs()
        {
        }
        public static new GetClusterArgs Empty => new GetClusterArgs();
    }

    public sealed class GetClusterInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Redis cluster ID.
        /// 
        /// &gt; **Note** You must specify at least one: `name` and/or `cluster_id`.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// The name of the Redis cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project the Redis cluster is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`) The zone in which the server exists.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetClusterInvokeArgs()
        {
        }
        public static new GetClusterInvokeArgs Empty => new GetClusterInvokeArgs();
    }


    [OutputType]
    public sealed class GetClusterResult
    {
        /// <summary>
        /// List of acl rules.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetClusterAclResult> Acls;
        /// <summary>
        /// The PEM of the certificate used by redis, only when `tls_enabled` is true.
        /// </summary>
        public readonly string Certificate;
        public readonly string? ClusterId;
        /// <summary>
        /// The number of nodes in the Redis Cluster.
        /// </summary>
        public readonly int ClusterSize;
        /// <summary>
        /// The date and time of creation of the Redis Cluster.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        /// <summary>
        /// The type of Redis Cluster (e.g. `RED1-M`).
        /// </summary>
        public readonly string NodeType;
        /// <summary>
        /// Password of the first user of the Redis Cluster.
        /// </summary>
        public readonly string Password;
        public readonly ImmutableArray<Outputs.GetClusterPrivateIpResult> PrivateIps;
        /// <summary>
        /// List of private networks endpoints of the Redis Cluster.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetClusterPrivateNetworkResult> PrivateNetworks;
        public readonly string? ProjectId;
        /// <summary>
        /// Public network details.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetClusterPublicNetworkResult> PublicNetworks;
        /// <summary>
        /// Map of settings for redis cluster.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Settings;
        /// <summary>
        /// The tags associated with the Redis Cluster.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// Whether TLS is enabled or not.
        /// </summary>
        public readonly bool TlsEnabled;
        /// <summary>
        /// The date and time of the last update of the Redis Cluster.
        /// </summary>
        public readonly string UpdatedAt;
        /// <summary>
        /// The first user of the Redis Cluster.
        /// </summary>
        public readonly string UserName;
        /// <summary>
        /// Redis's Cluster version (e.g. `6.2.7`).
        /// </summary>
        public readonly string Version;
        public readonly string? Zone;

        [OutputConstructor]
        private GetClusterResult(
            ImmutableArray<Outputs.GetClusterAclResult> acls,

            string certificate,

            string? clusterId,

            int clusterSize,

            string createdAt,

            string id,

            string? name,

            string nodeType,

            string password,

            ImmutableArray<Outputs.GetClusterPrivateIpResult> privateIps,

            ImmutableArray<Outputs.GetClusterPrivateNetworkResult> privateNetworks,

            string? projectId,

            ImmutableArray<Outputs.GetClusterPublicNetworkResult> publicNetworks,

            ImmutableDictionary<string, string> settings,

            ImmutableArray<string> tags,

            bool tlsEnabled,

            string updatedAt,

            string userName,

            string version,

            string? zone)
        {
            Acls = acls;
            Certificate = certificate;
            ClusterId = clusterId;
            ClusterSize = clusterSize;
            CreatedAt = createdAt;
            Id = id;
            Name = name;
            NodeType = nodeType;
            Password = password;
            PrivateIps = privateIps;
            PrivateNetworks = privateNetworks;
            ProjectId = projectId;
            PublicNetworks = publicNetworks;
            Settings = settings;
            Tags = tags;
            TlsEnabled = tlsEnabled;
            UpdatedAt = updatedAt;
            UserName = userName;
            Version = version;
            Zone = zone;
        }
    }
}
