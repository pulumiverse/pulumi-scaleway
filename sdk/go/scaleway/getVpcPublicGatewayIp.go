// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about a Public Gateway public flexible IP address.
//
// For further information, please see the [API documentation](https://www.scaleway.com/en/developers/api/public-gateway/#path-ips-list-ips).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/network"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := network.NewPublicGatewayIp(ctx, "main", nil)
//			if err != nil {
//				return err
//			}
//			_ = network.LookupPublicGatewayIpOutput(ctx, network.GetPublicGatewayIpOutputArgs{
//				IpId: main.ID(),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: scaleway.index/getvpcpublicgatewayip.getVpcPublicGatewayIp has been deprecated in favor of scaleway.network/getpublicgatewayip.getPublicGatewayIp
func LookupVpcPublicGatewayIp(ctx *pulumi.Context, args *LookupVpcPublicGatewayIpArgs, opts ...pulumi.InvokeOption) (*LookupVpcPublicGatewayIpResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpcPublicGatewayIpResult
	err := ctx.Invoke("scaleway:index/getVpcPublicGatewayIp:getVpcPublicGatewayIp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcPublicGatewayIp.
type LookupVpcPublicGatewayIpArgs struct {
	IpId *string `pulumi:"ipId"`
}

// A collection of values returned by getVpcPublicGatewayIp.
type LookupVpcPublicGatewayIpResult struct {
	Address   string `pulumi:"address"`
	CreatedAt string `pulumi:"createdAt"`
	// The provider-assigned unique ID for this managed resource.
	Id             string   `pulumi:"id"`
	IpId           *string  `pulumi:"ipId"`
	OrganizationId string   `pulumi:"organizationId"`
	ProjectId      string   `pulumi:"projectId"`
	Reverse        string   `pulumi:"reverse"`
	Tags           []string `pulumi:"tags"`
	UpdatedAt      string   `pulumi:"updatedAt"`
	Zone           string   `pulumi:"zone"`
}

func LookupVpcPublicGatewayIpOutput(ctx *pulumi.Context, args LookupVpcPublicGatewayIpOutputArgs, opts ...pulumi.InvokeOption) LookupVpcPublicGatewayIpResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupVpcPublicGatewayIpResultOutput, error) {
			args := v.(LookupVpcPublicGatewayIpArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway:index/getVpcPublicGatewayIp:getVpcPublicGatewayIp", args, LookupVpcPublicGatewayIpResultOutput{}, options).(LookupVpcPublicGatewayIpResultOutput), nil
		}).(LookupVpcPublicGatewayIpResultOutput)
}

// A collection of arguments for invoking getVpcPublicGatewayIp.
type LookupVpcPublicGatewayIpOutputArgs struct {
	IpId pulumi.StringPtrInput `pulumi:"ipId"`
}

func (LookupVpcPublicGatewayIpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcPublicGatewayIpArgs)(nil)).Elem()
}

// A collection of values returned by getVpcPublicGatewayIp.
type LookupVpcPublicGatewayIpResultOutput struct{ *pulumi.OutputState }

func (LookupVpcPublicGatewayIpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcPublicGatewayIpResult)(nil)).Elem()
}

func (o LookupVpcPublicGatewayIpResultOutput) ToLookupVpcPublicGatewayIpResultOutput() LookupVpcPublicGatewayIpResultOutput {
	return o
}

func (o LookupVpcPublicGatewayIpResultOutput) ToLookupVpcPublicGatewayIpResultOutputWithContext(ctx context.Context) LookupVpcPublicGatewayIpResultOutput {
	return o
}

func (o LookupVpcPublicGatewayIpResultOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayIpResult) string { return v.Address }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayIpResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayIpResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVpcPublicGatewayIpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayIpResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayIpResultOutput) IpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayIpResult) *string { return v.IpId }).(pulumi.StringPtrOutput)
}

func (o LookupVpcPublicGatewayIpResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayIpResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayIpResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayIpResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayIpResultOutput) Reverse() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayIpResult) string { return v.Reverse }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayIpResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayIpResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupVpcPublicGatewayIpResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayIpResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayIpResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayIpResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcPublicGatewayIpResultOutput{})
}
