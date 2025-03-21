// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about a Flexible IP.
//
// Deprecated: scaleway.index/getflexibleip.getFlexibleIp has been deprecated in favor of scaleway.elasticmetal/getip.getIp
func LookupFlexibleIp(ctx *pulumi.Context, args *LookupFlexibleIpArgs, opts ...pulumi.InvokeOption) (*LookupFlexibleIpResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFlexibleIpResult
	err := ctx.Invoke("scaleway:index/getFlexibleIp:getFlexibleIp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFlexibleIp.
type LookupFlexibleIpArgs struct {
	FlexibleIpId *string `pulumi:"flexibleIpId"`
	// The IP address.
	// Only one of `ipAddress` and `ipId` should be specified.
	IpAddress *string `pulumi:"ipAddress"`
	// (Defaults to provider `projectId`) The ID of the project the IP is in.
	ProjectId *string `pulumi:"projectId"`
}

// A collection of values returned by getFlexibleIp.
type LookupFlexibleIpResult struct {
	CreatedAt    string  `pulumi:"createdAt"`
	Description  string  `pulumi:"description"`
	FlexibleIpId *string `pulumi:"flexibleIpId"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	IpAddress *string `pulumi:"ipAddress"`
	IsIpv6    bool    `pulumi:"isIpv6"`
	// (Defaults to provider `organizationId`) The ID of the organization the IP is in.
	OrganizationId string `pulumi:"organizationId"`
	// (Defaults to provider `projectId`) The ID of the project the IP is in.
	ProjectId string `pulumi:"projectId"`
	// The reverse domain associated with this IP.
	Reverse string `pulumi:"reverse"`
	// The associated server ID if any
	ServerId  string   `pulumi:"serverId"`
	Status    string   `pulumi:"status"`
	Tags      []string `pulumi:"tags"`
	UpdatedAt string   `pulumi:"updatedAt"`
	Zone      string   `pulumi:"zone"`
}

func LookupFlexibleIpOutput(ctx *pulumi.Context, args LookupFlexibleIpOutputArgs, opts ...pulumi.InvokeOption) LookupFlexibleIpResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupFlexibleIpResultOutput, error) {
			args := v.(LookupFlexibleIpArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway:index/getFlexibleIp:getFlexibleIp", args, LookupFlexibleIpResultOutput{}, options).(LookupFlexibleIpResultOutput), nil
		}).(LookupFlexibleIpResultOutput)
}

// A collection of arguments for invoking getFlexibleIp.
type LookupFlexibleIpOutputArgs struct {
	FlexibleIpId pulumi.StringPtrInput `pulumi:"flexibleIpId"`
	// The IP address.
	// Only one of `ipAddress` and `ipId` should be specified.
	IpAddress pulumi.StringPtrInput `pulumi:"ipAddress"`
	// (Defaults to provider `projectId`) The ID of the project the IP is in.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
}

func (LookupFlexibleIpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFlexibleIpArgs)(nil)).Elem()
}

// A collection of values returned by getFlexibleIp.
type LookupFlexibleIpResultOutput struct{ *pulumi.OutputState }

func (LookupFlexibleIpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFlexibleIpResult)(nil)).Elem()
}

func (o LookupFlexibleIpResultOutput) ToLookupFlexibleIpResultOutput() LookupFlexibleIpResultOutput {
	return o
}

func (o LookupFlexibleIpResultOutput) ToLookupFlexibleIpResultOutputWithContext(ctx context.Context) LookupFlexibleIpResultOutput {
	return o
}

func (o LookupFlexibleIpResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlexibleIpResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupFlexibleIpResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlexibleIpResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupFlexibleIpResultOutput) FlexibleIpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFlexibleIpResult) *string { return v.FlexibleIpId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFlexibleIpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlexibleIpResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFlexibleIpResultOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFlexibleIpResult) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o LookupFlexibleIpResultOutput) IsIpv6() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupFlexibleIpResult) bool { return v.IsIpv6 }).(pulumi.BoolOutput)
}

// (Defaults to provider `organizationId`) The ID of the organization the IP is in.
func (o LookupFlexibleIpResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlexibleIpResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

// (Defaults to provider `projectId`) The ID of the project the IP is in.
func (o LookupFlexibleIpResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlexibleIpResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// The reverse domain associated with this IP.
func (o LookupFlexibleIpResultOutput) Reverse() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlexibleIpResult) string { return v.Reverse }).(pulumi.StringOutput)
}

// The associated server ID if any
func (o LookupFlexibleIpResultOutput) ServerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlexibleIpResult) string { return v.ServerId }).(pulumi.StringOutput)
}

func (o LookupFlexibleIpResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlexibleIpResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupFlexibleIpResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFlexibleIpResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupFlexibleIpResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlexibleIpResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupFlexibleIpResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlexibleIpResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFlexibleIpResultOutput{})
}
