// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about a Kubernetes Cluster's Pool.
//
// Deprecated: scaleway.index/getkubernetesnodepool.getKubernetesNodePool has been deprecated in favor of scaleway.kubernetes/getpool.getPool
func LookupKubernetesNodePool(ctx *pulumi.Context, args *LookupKubernetesNodePoolArgs, opts ...pulumi.InvokeOption) (*LookupKubernetesNodePoolResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupKubernetesNodePoolResult
	err := ctx.Invoke("scaleway:index/getKubernetesNodePool:getKubernetesNodePool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKubernetesNodePool.
type LookupKubernetesNodePoolArgs struct {
	// The cluster ID. Required when `name` is set.
	ClusterId *string `pulumi:"clusterId"`
	// The pool name. Only one of `name` and `poolId` should be specified. `clusterId` should be specified with `name`.
	Name *string `pulumi:"name"`
	// The pool's ID. Only one of `name` and `poolId` should be specified.
	PoolId *string `pulumi:"poolId"`
	// `region`) The region in which the pool exists.
	Region *string `pulumi:"region"`
	// The size of the pool.
	Size *int `pulumi:"size"`
}

// A collection of values returned by getKubernetesNodePool.
type LookupKubernetesNodePoolResult struct {
	// True if the autohealing feature is enabled for this pool.
	Autohealing bool `pulumi:"autohealing"`
	// True if the autoscaling feature is enabled for this pool.
	Autoscaling bool    `pulumi:"autoscaling"`
	ClusterId   *string `pulumi:"clusterId"`
	// The container runtime of the pool.
	ContainerRuntime string `pulumi:"containerRuntime"`
	// The creation date of the pool.
	CreatedAt   string `pulumi:"createdAt"`
	CurrentSize int    `pulumi:"currentSize"`
	// The provider-assigned unique ID for this managed resource.
	Id          string            `pulumi:"id"`
	KubeletArgs map[string]string `pulumi:"kubeletArgs"`
	// The maximum size of the pool, used by the autoscaling feature.
	MaxSize int `pulumi:"maxSize"`
	// The minimum size of the pool, used by the autoscaling feature.
	MinSize int `pulumi:"minSize"`
	// The name of the node.
	Name *string `pulumi:"name"`
	// The commercial type of the pool instances.
	NodeType string `pulumi:"nodeType"`
	// (List of) The nodes in the default pool.
	Nodes []GetKubernetesNodePoolNode `pulumi:"nodes"`
	// [placement group](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) the nodes of the pool are attached to.
	PlacementGroupId   string  `pulumi:"placementGroupId"`
	PoolId             *string `pulumi:"poolId"`
	PublicIpDisabled   bool    `pulumi:"publicIpDisabled"`
	Region             *string `pulumi:"region"`
	RootVolumeSizeInGb int     `pulumi:"rootVolumeSizeInGb"`
	RootVolumeType     string  `pulumi:"rootVolumeType"`
	SecurityGroupId    string  `pulumi:"securityGroupId"`
	// The size of the pool.
	Size *int `pulumi:"size"`
	// The status of the node.
	Status string `pulumi:"status"`
	// The tags associated with the pool.
	Tags []string `pulumi:"tags"`
	// The last update date of the pool.
	UpdatedAt       string                               `pulumi:"updatedAt"`
	UpgradePolicies []GetKubernetesNodePoolUpgradePolicy `pulumi:"upgradePolicies"`
	// The version of the pool.
	Version          string `pulumi:"version"`
	WaitForPoolReady bool   `pulumi:"waitForPoolReady"`
	Zone             string `pulumi:"zone"`
}

func LookupKubernetesNodePoolOutput(ctx *pulumi.Context, args LookupKubernetesNodePoolOutputArgs, opts ...pulumi.InvokeOption) LookupKubernetesNodePoolResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupKubernetesNodePoolResultOutput, error) {
			args := v.(LookupKubernetesNodePoolArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway:index/getKubernetesNodePool:getKubernetesNodePool", args, LookupKubernetesNodePoolResultOutput{}, options).(LookupKubernetesNodePoolResultOutput), nil
		}).(LookupKubernetesNodePoolResultOutput)
}

// A collection of arguments for invoking getKubernetesNodePool.
type LookupKubernetesNodePoolOutputArgs struct {
	// The cluster ID. Required when `name` is set.
	ClusterId pulumi.StringPtrInput `pulumi:"clusterId"`
	// The pool name. Only one of `name` and `poolId` should be specified. `clusterId` should be specified with `name`.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The pool's ID. Only one of `name` and `poolId` should be specified.
	PoolId pulumi.StringPtrInput `pulumi:"poolId"`
	// `region`) The region in which the pool exists.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The size of the pool.
	Size pulumi.IntPtrInput `pulumi:"size"`
}

func (LookupKubernetesNodePoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKubernetesNodePoolArgs)(nil)).Elem()
}

// A collection of values returned by getKubernetesNodePool.
type LookupKubernetesNodePoolResultOutput struct{ *pulumi.OutputState }

func (LookupKubernetesNodePoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKubernetesNodePoolResult)(nil)).Elem()
}

func (o LookupKubernetesNodePoolResultOutput) ToLookupKubernetesNodePoolResultOutput() LookupKubernetesNodePoolResultOutput {
	return o
}

func (o LookupKubernetesNodePoolResultOutput) ToLookupKubernetesNodePoolResultOutputWithContext(ctx context.Context) LookupKubernetesNodePoolResultOutput {
	return o
}

// True if the autohealing feature is enabled for this pool.
func (o LookupKubernetesNodePoolResultOutput) Autohealing() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupKubernetesNodePoolResult) bool { return v.Autohealing }).(pulumi.BoolOutput)
}

// True if the autoscaling feature is enabled for this pool.
func (o LookupKubernetesNodePoolResultOutput) Autoscaling() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupKubernetesNodePoolResult) bool { return v.Autoscaling }).(pulumi.BoolOutput)
}

func (o LookupKubernetesNodePoolResultOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKubernetesNodePoolResult) *string { return v.ClusterId }).(pulumi.StringPtrOutput)
}

// The container runtime of the pool.
func (o LookupKubernetesNodePoolResultOutput) ContainerRuntime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesNodePoolResult) string { return v.ContainerRuntime }).(pulumi.StringOutput)
}

// The creation date of the pool.
func (o LookupKubernetesNodePoolResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesNodePoolResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupKubernetesNodePoolResultOutput) CurrentSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupKubernetesNodePoolResult) int { return v.CurrentSize }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupKubernetesNodePoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesNodePoolResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupKubernetesNodePoolResultOutput) KubeletArgs() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupKubernetesNodePoolResult) map[string]string { return v.KubeletArgs }).(pulumi.StringMapOutput)
}

// The maximum size of the pool, used by the autoscaling feature.
func (o LookupKubernetesNodePoolResultOutput) MaxSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupKubernetesNodePoolResult) int { return v.MaxSize }).(pulumi.IntOutput)
}

// The minimum size of the pool, used by the autoscaling feature.
func (o LookupKubernetesNodePoolResultOutput) MinSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupKubernetesNodePoolResult) int { return v.MinSize }).(pulumi.IntOutput)
}

// The name of the node.
func (o LookupKubernetesNodePoolResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKubernetesNodePoolResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The commercial type of the pool instances.
func (o LookupKubernetesNodePoolResultOutput) NodeType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesNodePoolResult) string { return v.NodeType }).(pulumi.StringOutput)
}

// (List of) The nodes in the default pool.
func (o LookupKubernetesNodePoolResultOutput) Nodes() GetKubernetesNodePoolNodeArrayOutput {
	return o.ApplyT(func(v LookupKubernetesNodePoolResult) []GetKubernetesNodePoolNode { return v.Nodes }).(GetKubernetesNodePoolNodeArrayOutput)
}

// [placement group](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) the nodes of the pool are attached to.
func (o LookupKubernetesNodePoolResultOutput) PlacementGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesNodePoolResult) string { return v.PlacementGroupId }).(pulumi.StringOutput)
}

func (o LookupKubernetesNodePoolResultOutput) PoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKubernetesNodePoolResult) *string { return v.PoolId }).(pulumi.StringPtrOutput)
}

func (o LookupKubernetesNodePoolResultOutput) PublicIpDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupKubernetesNodePoolResult) bool { return v.PublicIpDisabled }).(pulumi.BoolOutput)
}

func (o LookupKubernetesNodePoolResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKubernetesNodePoolResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupKubernetesNodePoolResultOutput) RootVolumeSizeInGb() pulumi.IntOutput {
	return o.ApplyT(func(v LookupKubernetesNodePoolResult) int { return v.RootVolumeSizeInGb }).(pulumi.IntOutput)
}

func (o LookupKubernetesNodePoolResultOutput) RootVolumeType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesNodePoolResult) string { return v.RootVolumeType }).(pulumi.StringOutput)
}

func (o LookupKubernetesNodePoolResultOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesNodePoolResult) string { return v.SecurityGroupId }).(pulumi.StringOutput)
}

// The size of the pool.
func (o LookupKubernetesNodePoolResultOutput) Size() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupKubernetesNodePoolResult) *int { return v.Size }).(pulumi.IntPtrOutput)
}

// The status of the node.
func (o LookupKubernetesNodePoolResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesNodePoolResult) string { return v.Status }).(pulumi.StringOutput)
}

// The tags associated with the pool.
func (o LookupKubernetesNodePoolResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupKubernetesNodePoolResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The last update date of the pool.
func (o LookupKubernetesNodePoolResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesNodePoolResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupKubernetesNodePoolResultOutput) UpgradePolicies() GetKubernetesNodePoolUpgradePolicyArrayOutput {
	return o.ApplyT(func(v LookupKubernetesNodePoolResult) []GetKubernetesNodePoolUpgradePolicy { return v.UpgradePolicies }).(GetKubernetesNodePoolUpgradePolicyArrayOutput)
}

// The version of the pool.
func (o LookupKubernetesNodePoolResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesNodePoolResult) string { return v.Version }).(pulumi.StringOutput)
}

func (o LookupKubernetesNodePoolResultOutput) WaitForPoolReady() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupKubernetesNodePoolResult) bool { return v.WaitForPoolReady }).(pulumi.BoolOutput)
}

func (o LookupKubernetesNodePoolResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesNodePoolResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKubernetesNodePoolResultOutput{})
}
