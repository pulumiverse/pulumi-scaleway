// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about a MongoDB® Instance.
//
// For further information refer to the Managed Databases for MongoDB® [API documentation](https://developers.scaleway.com/en/products/mongodb/api/)
//
// Deprecated: scaleway.index/getmongodbinstance.getMongoDbInstance has been deprecated in favor of scaleway.mongodb/getinstance.getInstance
func LookupMongoDbInstance(ctx *pulumi.Context, args *LookupMongoDbInstanceArgs, opts ...pulumi.InvokeOption) (*LookupMongoDbInstanceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupMongoDbInstanceResult
	err := ctx.Invoke("scaleway:index/getMongoDbInstance:getMongoDbInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMongoDbInstance.
type LookupMongoDbInstanceArgs struct {
	// The MongoDB® instance ID.
	//
	// > **Note** You must specify at least one: `name` or `instanceId`.
	InstanceId *string `pulumi:"instanceId"`
	// The name of the MongoDB® instance.
	Name *string `pulumi:"name"`
	// The ID of the project the MongoDB® instance is in. Can be used to filter instances when using `name`.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region in which the MongoDB® Instance exists.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getMongoDbInstance.
type LookupMongoDbInstanceResult struct {
	// The date and time the MongoDB® instance was created.
	CreatedAt string `pulumi:"createdAt"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	InstanceId *string `pulumi:"instanceId"`
	// The name of the MongoDB® instance.
	Name *string `pulumi:"name"`
	// The number of nodes in the MongoDB® cluster.
	NodeNumber int `pulumi:"nodeNumber"`
	// The type of MongoDB® node.
	NodeType        string                             `pulumi:"nodeType"`
	Password        string                             `pulumi:"password"`
	PrivateIps      []GetMongoDbInstancePrivateIp      `pulumi:"privateIps"`
	PrivateNetworks []GetMongoDbInstancePrivateNetwork `pulumi:"privateNetworks"`
	// The ID of the project the instance belongs to.
	ProjectId *string `pulumi:"projectId"`
	// The details of the public network configuration, if applicable.
	PublicNetworks []GetMongoDbInstancePublicNetwork `pulumi:"publicNetworks"`
	Region         *string                           `pulumi:"region"`
	Settings       map[string]string                 `pulumi:"settings"`
	SnapshotId     string                            `pulumi:"snapshotId"`
	// A list of tags attached to the MongoDB® instance.
	Tags           []string `pulumi:"tags"`
	TlsCertificate string   `pulumi:"tlsCertificate"`
	UpdatedAt      string   `pulumi:"updatedAt"`
	UserName       string   `pulumi:"userName"`
	// The version of MongoDB® running on the instance.
	Version string `pulumi:"version"`
	// The size of the attached volume, in GB.
	VolumeSizeInGb int `pulumi:"volumeSizeInGb"`
	// The type of volume attached to the MongoDB® instance.
	VolumeType string `pulumi:"volumeType"`
}

func LookupMongoDbInstanceOutput(ctx *pulumi.Context, args LookupMongoDbInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupMongoDbInstanceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupMongoDbInstanceResultOutput, error) {
			args := v.(LookupMongoDbInstanceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway:index/getMongoDbInstance:getMongoDbInstance", args, LookupMongoDbInstanceResultOutput{}, options).(LookupMongoDbInstanceResultOutput), nil
		}).(LookupMongoDbInstanceResultOutput)
}

// A collection of arguments for invoking getMongoDbInstance.
type LookupMongoDbInstanceOutputArgs struct {
	// The MongoDB® instance ID.
	//
	// > **Note** You must specify at least one: `name` or `instanceId`.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// The name of the MongoDB® instance.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the project the MongoDB® instance is in. Can be used to filter instances when using `name`.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// `region`) The region in which the MongoDB® Instance exists.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupMongoDbInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMongoDbInstanceArgs)(nil)).Elem()
}

// A collection of values returned by getMongoDbInstance.
type LookupMongoDbInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupMongoDbInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMongoDbInstanceResult)(nil)).Elem()
}

func (o LookupMongoDbInstanceResultOutput) ToLookupMongoDbInstanceResultOutput() LookupMongoDbInstanceResultOutput {
	return o
}

func (o LookupMongoDbInstanceResultOutput) ToLookupMongoDbInstanceResultOutputWithContext(ctx context.Context) LookupMongoDbInstanceResultOutput {
	return o
}

// The date and time the MongoDB® instance was created.
func (o LookupMongoDbInstanceResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDbInstanceResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupMongoDbInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDbInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMongoDbInstanceResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMongoDbInstanceResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

// The name of the MongoDB® instance.
func (o LookupMongoDbInstanceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMongoDbInstanceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The number of nodes in the MongoDB® cluster.
func (o LookupMongoDbInstanceResultOutput) NodeNumber() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMongoDbInstanceResult) int { return v.NodeNumber }).(pulumi.IntOutput)
}

// The type of MongoDB® node.
func (o LookupMongoDbInstanceResultOutput) NodeType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDbInstanceResult) string { return v.NodeType }).(pulumi.StringOutput)
}

func (o LookupMongoDbInstanceResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDbInstanceResult) string { return v.Password }).(pulumi.StringOutput)
}

func (o LookupMongoDbInstanceResultOutput) PrivateIps() GetMongoDbInstancePrivateIpArrayOutput {
	return o.ApplyT(func(v LookupMongoDbInstanceResult) []GetMongoDbInstancePrivateIp { return v.PrivateIps }).(GetMongoDbInstancePrivateIpArrayOutput)
}

func (o LookupMongoDbInstanceResultOutput) PrivateNetworks() GetMongoDbInstancePrivateNetworkArrayOutput {
	return o.ApplyT(func(v LookupMongoDbInstanceResult) []GetMongoDbInstancePrivateNetwork { return v.PrivateNetworks }).(GetMongoDbInstancePrivateNetworkArrayOutput)
}

// The ID of the project the instance belongs to.
func (o LookupMongoDbInstanceResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMongoDbInstanceResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// The details of the public network configuration, if applicable.
func (o LookupMongoDbInstanceResultOutput) PublicNetworks() GetMongoDbInstancePublicNetworkArrayOutput {
	return o.ApplyT(func(v LookupMongoDbInstanceResult) []GetMongoDbInstancePublicNetwork { return v.PublicNetworks }).(GetMongoDbInstancePublicNetworkArrayOutput)
}

func (o LookupMongoDbInstanceResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMongoDbInstanceResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupMongoDbInstanceResultOutput) Settings() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMongoDbInstanceResult) map[string]string { return v.Settings }).(pulumi.StringMapOutput)
}

func (o LookupMongoDbInstanceResultOutput) SnapshotId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDbInstanceResult) string { return v.SnapshotId }).(pulumi.StringOutput)
}

// A list of tags attached to the MongoDB® instance.
func (o LookupMongoDbInstanceResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMongoDbInstanceResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupMongoDbInstanceResultOutput) TlsCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDbInstanceResult) string { return v.TlsCertificate }).(pulumi.StringOutput)
}

func (o LookupMongoDbInstanceResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDbInstanceResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupMongoDbInstanceResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDbInstanceResult) string { return v.UserName }).(pulumi.StringOutput)
}

// The version of MongoDB® running on the instance.
func (o LookupMongoDbInstanceResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDbInstanceResult) string { return v.Version }).(pulumi.StringOutput)
}

// The size of the attached volume, in GB.
func (o LookupMongoDbInstanceResultOutput) VolumeSizeInGb() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMongoDbInstanceResult) int { return v.VolumeSizeInGb }).(pulumi.IntOutput)
}

// The type of volume attached to the MongoDB® instance.
func (o LookupMongoDbInstanceResultOutput) VolumeType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDbInstanceResult) string { return v.VolumeType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMongoDbInstanceResultOutput{})
}
