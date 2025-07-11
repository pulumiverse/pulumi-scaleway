// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

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
func LookupPublicGatewayIp(ctx *pulumi.Context, args *LookupPublicGatewayIpArgs, opts ...pulumi.InvokeOption) (*LookupPublicGatewayIpResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPublicGatewayIpResult
	err := ctx.Invoke("scaleway:network/getPublicGatewayIp:getPublicGatewayIp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPublicGatewayIp.
type LookupPublicGatewayIpArgs struct {
	IpId *string `pulumi:"ipId"`
}

// A collection of values returned by getPublicGatewayIp.
type LookupPublicGatewayIpResult struct {
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

func LookupPublicGatewayIpOutput(ctx *pulumi.Context, args LookupPublicGatewayIpOutputArgs, opts ...pulumi.InvokeOption) LookupPublicGatewayIpResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupPublicGatewayIpResultOutput, error) {
			args := v.(LookupPublicGatewayIpArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway:network/getPublicGatewayIp:getPublicGatewayIp", args, LookupPublicGatewayIpResultOutput{}, options).(LookupPublicGatewayIpResultOutput), nil
		}).(LookupPublicGatewayIpResultOutput)
}

// A collection of arguments for invoking getPublicGatewayIp.
type LookupPublicGatewayIpOutputArgs struct {
	IpId pulumi.StringPtrInput `pulumi:"ipId"`
}

func (LookupPublicGatewayIpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPublicGatewayIpArgs)(nil)).Elem()
}

// A collection of values returned by getPublicGatewayIp.
type LookupPublicGatewayIpResultOutput struct{ *pulumi.OutputState }

func (LookupPublicGatewayIpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPublicGatewayIpResult)(nil)).Elem()
}

func (o LookupPublicGatewayIpResultOutput) ToLookupPublicGatewayIpResultOutput() LookupPublicGatewayIpResultOutput {
	return o
}

func (o LookupPublicGatewayIpResultOutput) ToLookupPublicGatewayIpResultOutputWithContext(ctx context.Context) LookupPublicGatewayIpResultOutput {
	return o
}

func (o LookupPublicGatewayIpResultOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayIpResult) string { return v.Address }).(pulumi.StringOutput)
}

func (o LookupPublicGatewayIpResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayIpResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPublicGatewayIpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayIpResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPublicGatewayIpResultOutput) IpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicGatewayIpResult) *string { return v.IpId }).(pulumi.StringPtrOutput)
}

func (o LookupPublicGatewayIpResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayIpResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupPublicGatewayIpResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayIpResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupPublicGatewayIpResultOutput) Reverse() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayIpResult) string { return v.Reverse }).(pulumi.StringOutput)
}

func (o LookupPublicGatewayIpResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPublicGatewayIpResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupPublicGatewayIpResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayIpResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupPublicGatewayIpResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayIpResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPublicGatewayIpResultOutput{})
}
