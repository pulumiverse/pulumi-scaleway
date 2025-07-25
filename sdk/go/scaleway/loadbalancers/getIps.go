// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancers

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about multiple Load Balancer IP addresses.
//
// For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/how-to/create-manage-flex-ips/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-ip-addresses-list-ip-addresses).
func GetIps(ctx *pulumi.Context, args *GetIpsArgs, opts ...pulumi.InvokeOption) (*GetIpsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetIpsResult
	err := ctx.Invoke("scaleway:loadbalancers/getIps:getIps", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIps.
type GetIpsArgs struct {
	// The IP CIDR range to filter for. IPs within a matching CIDR block are listed.
	IpCidrRange *string `pulumi:"ipCidrRange"`
	// The IP type used as a filter.
	IpType *string `pulumi:"ipType"`
	// The ID of the Project the Load Balancer is associated with.
	ProjectId *string `pulumi:"projectId"`
	// List of tags used as filter. IPs with these exact tags are listed.
	Tags []string `pulumi:"tags"`
	// `zone`) The zone in which the IPs exist.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getIps.
type GetIpsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	IpCidrRange *string `pulumi:"ipCidrRange"`
	IpType      *string `pulumi:"ipType"`
	// List of retrieved IPs
	Ips []GetIpsIp `pulumi:"ips"`
	// The ID of the Organization the Load Balancer is associated with.
	OrganizationId string `pulumi:"organizationId"`
	// The ID of the Project the Load Balancer is associated with.
	ProjectId string   `pulumi:"projectId"`
	Tags      []string `pulumi:"tags"`
	// The zone of the Load Balancer.
	Zone string `pulumi:"zone"`
}

func GetIpsOutput(ctx *pulumi.Context, args GetIpsOutputArgs, opts ...pulumi.InvokeOption) GetIpsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetIpsResultOutput, error) {
			args := v.(GetIpsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway:loadbalancers/getIps:getIps", args, GetIpsResultOutput{}, options).(GetIpsResultOutput), nil
		}).(GetIpsResultOutput)
}

// A collection of arguments for invoking getIps.
type GetIpsOutputArgs struct {
	// The IP CIDR range to filter for. IPs within a matching CIDR block are listed.
	IpCidrRange pulumi.StringPtrInput `pulumi:"ipCidrRange"`
	// The IP type used as a filter.
	IpType pulumi.StringPtrInput `pulumi:"ipType"`
	// The ID of the Project the Load Balancer is associated with.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// List of tags used as filter. IPs with these exact tags are listed.
	Tags pulumi.StringArrayInput `pulumi:"tags"`
	// `zone`) The zone in which the IPs exist.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (GetIpsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpsArgs)(nil)).Elem()
}

// A collection of values returned by getIps.
type GetIpsResultOutput struct{ *pulumi.OutputState }

func (GetIpsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpsResult)(nil)).Elem()
}

func (o GetIpsResultOutput) ToGetIpsResultOutput() GetIpsResultOutput {
	return o
}

func (o GetIpsResultOutput) ToGetIpsResultOutputWithContext(ctx context.Context) GetIpsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetIpsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetIpsResultOutput) IpCidrRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIpsResult) *string { return v.IpCidrRange }).(pulumi.StringPtrOutput)
}

func (o GetIpsResultOutput) IpType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIpsResult) *string { return v.IpType }).(pulumi.StringPtrOutput)
}

// List of retrieved IPs
func (o GetIpsResultOutput) Ips() GetIpsIpArrayOutput {
	return o.ApplyT(func(v GetIpsResult) []GetIpsIp { return v.Ips }).(GetIpsIpArrayOutput)
}

// The ID of the Organization the Load Balancer is associated with.
func (o GetIpsResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpsResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

// The ID of the Project the Load Balancer is associated with.
func (o GetIpsResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpsResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o GetIpsResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpsResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The zone of the Load Balancer.
func (o GetIpsResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpsResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIpsResultOutput{})
}
