// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Kubernetes
{
    public static class GetCluster
    {
        /// <summary>
        /// Gets information about a Kubernetes Cluster.
        /// </summary>
        public static Task<GetClusterResult> InvokeAsync(GetClusterArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("scaleway:kubernetes/getCluster:getCluster", args ?? new GetClusterArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a Kubernetes Cluster.
        /// </summary>
        public static Output<GetClusterResult> Invoke(GetClusterInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetClusterResult>("scaleway:kubernetes/getCluster:getCluster", args ?? new GetClusterInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a Kubernetes Cluster.
        /// </summary>
        public static Output<GetClusterResult> Invoke(GetClusterInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetClusterResult>("scaleway:kubernetes/getCluster:getCluster", args ?? new GetClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClusterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The cluster ID. Only one of `name` and `cluster_id` should be specified.
        /// </summary>
        [Input("clusterId")]
        public string? ClusterId { get; set; }

        /// <summary>
        /// The cluster name. Only one of `name` and `cluster_id` should be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of the project the cluster is associated with.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region in which the cluster exists.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetClusterArgs()
        {
        }
        public static new GetClusterArgs Empty => new GetClusterArgs();
    }

    public sealed class GetClusterInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The cluster ID. Only one of `name` and `cluster_id` should be specified.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// The cluster name. Only one of `name` and `cluster_id` should be specified.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project the cluster is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region in which the cluster exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetClusterInvokeArgs()
        {
        }
        public static new GetClusterInvokeArgs Empty => new GetClusterInvokeArgs();
    }


    [OutputType]
    public sealed class GetClusterResult
    {
        /// <summary>
        /// The list of [admission plugins](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) enabled on the cluster.
        /// </summary>
        public readonly ImmutableArray<string> AdmissionPlugins;
        public readonly ImmutableArray<string> ApiserverCertSans;
        /// <summary>
        /// The URL of the Kubernetes API server.
        /// </summary>
        public readonly string ApiserverUrl;
        /// <summary>
        /// The auto upgrade configuration.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetClusterAutoUpgradeResult> AutoUpgrades;
        /// <summary>
        /// The configuration options for the [Kubernetes cluster autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler).
        /// </summary>
        public readonly ImmutableArray<Outputs.GetClusterAutoscalerConfigResult> AutoscalerConfigs;
        public readonly string? ClusterId;
        /// <summary>
        /// The Container Network Interface (CNI) for the Kubernetes cluster.
        /// </summary>
        public readonly string Cni;
        /// <summary>
        /// The creation date of the cluster.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// A description for the Kubernetes cluster.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The list of [feature gates](https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/) enabled on the cluster.
        /// </summary>
        public readonly ImmutableArray<string> FeatureGates;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetClusterKubeconfigResult> Kubeconfigs;
        public readonly string? Name;
        public readonly ImmutableArray<Outputs.GetClusterOpenIdConnectConfigResult> OpenIdConnectConfigs;
        /// <summary>
        /// The ID of the organization the cluster is associated with.
        /// </summary>
        public readonly string OrganizationId;
        /// <summary>
        /// The ID of the private network of the cluster.
        /// </summary>
        public readonly string PrivateNetworkId;
        public readonly string? ProjectId;
        /// <summary>
        /// The region in which the cluster is.
        /// </summary>
        public readonly string? Region;
        /// <summary>
        /// The status of the Kubernetes cluster.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The tags associated with the Kubernetes cluster.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The type of the Kubernetes cluster.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The last update date of the cluster.
        /// </summary>
        public readonly string UpdatedAt;
        /// <summary>
        /// True if a newer Kubernetes version is available.
        /// </summary>
        public readonly bool UpgradeAvailable;
        /// <summary>
        /// The version of the Kubernetes cluster.
        /// </summary>
        public readonly string Version;
        /// <summary>
        /// The DNS wildcard that points to all ready nodes.
        /// </summary>
        public readonly string WildcardDns;

        [OutputConstructor]
        private GetClusterResult(
            ImmutableArray<string> admissionPlugins,

            ImmutableArray<string> apiserverCertSans,

            string apiserverUrl,

            ImmutableArray<Outputs.GetClusterAutoUpgradeResult> autoUpgrades,

            ImmutableArray<Outputs.GetClusterAutoscalerConfigResult> autoscalerConfigs,

            string? clusterId,

            string cni,

            string createdAt,

            string description,

            ImmutableArray<string> featureGates,

            string id,

            ImmutableArray<Outputs.GetClusterKubeconfigResult> kubeconfigs,

            string? name,

            ImmutableArray<Outputs.GetClusterOpenIdConnectConfigResult> openIdConnectConfigs,

            string organizationId,

            string privateNetworkId,

            string? projectId,

            string? region,

            string status,

            ImmutableArray<string> tags,

            string type,

            string updatedAt,

            bool upgradeAvailable,

            string version,

            string wildcardDns)
        {
            AdmissionPlugins = admissionPlugins;
            ApiserverCertSans = apiserverCertSans;
            ApiserverUrl = apiserverUrl;
            AutoUpgrades = autoUpgrades;
            AutoscalerConfigs = autoscalerConfigs;
            ClusterId = clusterId;
            Cni = cni;
            CreatedAt = createdAt;
            Description = description;
            FeatureGates = featureGates;
            Id = id;
            Kubeconfigs = kubeconfigs;
            Name = name;
            OpenIdConnectConfigs = openIdConnectConfigs;
            OrganizationId = organizationId;
            PrivateNetworkId = privateNetworkId;
            ProjectId = projectId;
            Region = region;
            Status = status;
            Tags = tags;
            Type = type;
            UpdatedAt = updatedAt;
            UpgradeAvailable = upgradeAvailable;
            Version = version;
            WildcardDns = wildcardDns;
        }
    }
}
