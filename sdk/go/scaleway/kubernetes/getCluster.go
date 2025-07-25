// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about a Kubernetes Cluster.
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupClusterResult
	err := ctx.Invoke("scaleway:kubernetes/getCluster:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCluster.
type LookupClusterArgs struct {
	// The cluster ID. Only one of `name` and `clusterId` should be specified.
	ClusterId *string `pulumi:"clusterId"`
	// The cluster name. Only one of `name` and `clusterId` should be specified.
	Name *string `pulumi:"name"`
	// The ID of the project the cluster is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region in which the cluster exists.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getCluster.
type LookupClusterResult struct {
	// The list of [admission plugins](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) enabled on the cluster.
	AdmissionPlugins  []string `pulumi:"admissionPlugins"`
	ApiserverCertSans []string `pulumi:"apiserverCertSans"`
	// The URL of the Kubernetes API server.
	ApiserverUrl string `pulumi:"apiserverUrl"`
	// The auto upgrade configuration.
	AutoUpgrades []GetClusterAutoUpgrade `pulumi:"autoUpgrades"`
	// The configuration options for the [Kubernetes cluster autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler).
	AutoscalerConfigs []GetClusterAutoscalerConfig `pulumi:"autoscalerConfigs"`
	ClusterId         *string                      `pulumi:"clusterId"`
	// The Container Network Interface (CNI) for the Kubernetes cluster.
	Cni string `pulumi:"cni"`
	// The creation date of the cluster.
	CreatedAt string `pulumi:"createdAt"`
	// A description for the Kubernetes cluster.
	Description string `pulumi:"description"`
	// The list of [feature gates](https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/) enabled on the cluster.
	FeatureGates []string `pulumi:"featureGates"`
	// The provider-assigned unique ID for this managed resource.
	Id                   string                          `pulumi:"id"`
	Kubeconfigs          []GetClusterKubeconfig          `pulumi:"kubeconfigs"`
	Name                 *string                         `pulumi:"name"`
	OpenIdConnectConfigs []GetClusterOpenIdConnectConfig `pulumi:"openIdConnectConfigs"`
	// The ID of the organization the cluster is associated with.
	OrganizationId string `pulumi:"organizationId"`
	// The ID of the private network of the cluster.
	PrivateNetworkId string  `pulumi:"privateNetworkId"`
	ProjectId        *string `pulumi:"projectId"`
	// The region in which the cluster is.
	Region *string `pulumi:"region"`
	// The status of the Kubernetes cluster.
	Status string `pulumi:"status"`
	// The tags associated with the Kubernetes cluster.
	Tags []string `pulumi:"tags"`
	// The type of the Kubernetes cluster.
	Type string `pulumi:"type"`
	// The last update date of the cluster.
	UpdatedAt string `pulumi:"updatedAt"`
	// True if a newer Kubernetes version is available.
	UpgradeAvailable bool `pulumi:"upgradeAvailable"`
	// The version of the Kubernetes cluster.
	Version string `pulumi:"version"`
	// The DNS wildcard that points to all ready nodes.
	WildcardDns string `pulumi:"wildcardDns"`
}

func LookupClusterOutput(ctx *pulumi.Context, args LookupClusterOutputArgs, opts ...pulumi.InvokeOption) LookupClusterResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupClusterResultOutput, error) {
			args := v.(LookupClusterArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway:kubernetes/getCluster:getCluster", args, LookupClusterResultOutput{}, options).(LookupClusterResultOutput), nil
		}).(LookupClusterResultOutput)
}

// A collection of arguments for invoking getCluster.
type LookupClusterOutputArgs struct {
	// The cluster ID. Only one of `name` and `clusterId` should be specified.
	ClusterId pulumi.StringPtrInput `pulumi:"clusterId"`
	// The cluster name. Only one of `name` and `clusterId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the project the cluster is associated with.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// `region`) The region in which the cluster exists.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterArgs)(nil)).Elem()
}

// A collection of values returned by getCluster.
type LookupClusterResultOutput struct{ *pulumi.OutputState }

func (LookupClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterResult)(nil)).Elem()
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutput() LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutputWithContext(ctx context.Context) LookupClusterResultOutput {
	return o
}

// The list of [admission plugins](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) enabled on the cluster.
func (o LookupClusterResultOutput) AdmissionPlugins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []string { return v.AdmissionPlugins }).(pulumi.StringArrayOutput)
}

func (o LookupClusterResultOutput) ApiserverCertSans() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []string { return v.ApiserverCertSans }).(pulumi.StringArrayOutput)
}

// The URL of the Kubernetes API server.
func (o LookupClusterResultOutput) ApiserverUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ApiserverUrl }).(pulumi.StringOutput)
}

// The auto upgrade configuration.
func (o LookupClusterResultOutput) AutoUpgrades() GetClusterAutoUpgradeArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []GetClusterAutoUpgrade { return v.AutoUpgrades }).(GetClusterAutoUpgradeArrayOutput)
}

// The configuration options for the [Kubernetes cluster autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler).
func (o LookupClusterResultOutput) AutoscalerConfigs() GetClusterAutoscalerConfigArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []GetClusterAutoscalerConfig { return v.AutoscalerConfigs }).(GetClusterAutoscalerConfigArrayOutput)
}

func (o LookupClusterResultOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.ClusterId }).(pulumi.StringPtrOutput)
}

// The Container Network Interface (CNI) for the Kubernetes cluster.
func (o LookupClusterResultOutput) Cni() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Cni }).(pulumi.StringOutput)
}

// The creation date of the cluster.
func (o LookupClusterResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// A description for the Kubernetes cluster.
func (o LookupClusterResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Description }).(pulumi.StringOutput)
}

// The list of [feature gates](https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/) enabled on the cluster.
func (o LookupClusterResultOutput) FeatureGates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []string { return v.FeatureGates }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) Kubeconfigs() GetClusterKubeconfigArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []GetClusterKubeconfig { return v.Kubeconfigs }).(GetClusterKubeconfigArrayOutput)
}

func (o LookupClusterResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) OpenIdConnectConfigs() GetClusterOpenIdConnectConfigArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []GetClusterOpenIdConnectConfig { return v.OpenIdConnectConfigs }).(GetClusterOpenIdConnectConfigArrayOutput)
}

// The ID of the organization the cluster is associated with.
func (o LookupClusterResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

// The ID of the private network of the cluster.
func (o LookupClusterResultOutput) PrivateNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.PrivateNetworkId }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// The region in which the cluster is.
func (o LookupClusterResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

// The status of the Kubernetes cluster.
func (o LookupClusterResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Status }).(pulumi.StringOutput)
}

// The tags associated with the Kubernetes cluster.
func (o LookupClusterResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The type of the Kubernetes cluster.
func (o LookupClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

// The last update date of the cluster.
func (o LookupClusterResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

// True if a newer Kubernetes version is available.
func (o LookupClusterResultOutput) UpgradeAvailable() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupClusterResult) bool { return v.UpgradeAvailable }).(pulumi.BoolOutput)
}

// The version of the Kubernetes cluster.
func (o LookupClusterResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Version }).(pulumi.StringOutput)
}

// The DNS wildcard that points to all ready nodes.
func (o LookupClusterResultOutput) WildcardDns() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.WildcardDns }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterResultOutput{})
}
