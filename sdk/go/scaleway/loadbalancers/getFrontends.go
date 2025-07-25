// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancers

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about multiple Load Balancer frontends.
//
// For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/reference-content/configuring-frontends/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-frontends).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/loadbalancers"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// Find frontends that share the same LB ID
//			_, err := loadbalancers.GetFrontends(ctx, &loadbalancers.GetFrontendsArgs{
//				LbId: lb01.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			// Find frontends by LB ID and name
//			_, err = loadbalancers.GetFrontends(ctx, &loadbalancers.GetFrontendsArgs{
//				LbId: lb01.Id,
//				Name: pulumi.StringRef("tf-frontend-datasource"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetFrontends(ctx *pulumi.Context, args *GetFrontendsArgs, opts ...pulumi.InvokeOption) (*GetFrontendsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetFrontendsResult
	err := ctx.Invoke("scaleway:loadbalancers/getFrontends:getFrontends", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFrontends.
type GetFrontendsArgs struct {
	// The Load Balancer ID this frontend is attached to. Frontends with a matching ID are listed.
	LbId string `pulumi:"lbId"`
	// The frontend name to filter for. Frontends with a matching name are listed.
	Name      *string `pulumi:"name"`
	ProjectId *string `pulumi:"projectId"`
	// `zone`) The zone in which the frontends exist.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getFrontends.
type GetFrontendsResult struct {
	// List of retrieved frontends
	Frontends []GetFrontendsFrontend `pulumi:"frontends"`
	// The provider-assigned unique ID for this managed resource.
	Id             string  `pulumi:"id"`
	LbId           string  `pulumi:"lbId"`
	Name           *string `pulumi:"name"`
	OrganizationId string  `pulumi:"organizationId"`
	ProjectId      string  `pulumi:"projectId"`
	Zone           string  `pulumi:"zone"`
}

func GetFrontendsOutput(ctx *pulumi.Context, args GetFrontendsOutputArgs, opts ...pulumi.InvokeOption) GetFrontendsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetFrontendsResultOutput, error) {
			args := v.(GetFrontendsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway:loadbalancers/getFrontends:getFrontends", args, GetFrontendsResultOutput{}, options).(GetFrontendsResultOutput), nil
		}).(GetFrontendsResultOutput)
}

// A collection of arguments for invoking getFrontends.
type GetFrontendsOutputArgs struct {
	// The Load Balancer ID this frontend is attached to. Frontends with a matching ID are listed.
	LbId pulumi.StringInput `pulumi:"lbId"`
	// The frontend name to filter for. Frontends with a matching name are listed.
	Name      pulumi.StringPtrInput `pulumi:"name"`
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// `zone`) The zone in which the frontends exist.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (GetFrontendsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFrontendsArgs)(nil)).Elem()
}

// A collection of values returned by getFrontends.
type GetFrontendsResultOutput struct{ *pulumi.OutputState }

func (GetFrontendsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFrontendsResult)(nil)).Elem()
}

func (o GetFrontendsResultOutput) ToGetFrontendsResultOutput() GetFrontendsResultOutput {
	return o
}

func (o GetFrontendsResultOutput) ToGetFrontendsResultOutputWithContext(ctx context.Context) GetFrontendsResultOutput {
	return o
}

// List of retrieved frontends
func (o GetFrontendsResultOutput) Frontends() GetFrontendsFrontendArrayOutput {
	return o.ApplyT(func(v GetFrontendsResult) []GetFrontendsFrontend { return v.Frontends }).(GetFrontendsFrontendArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFrontendsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFrontendsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetFrontendsResultOutput) LbId() pulumi.StringOutput {
	return o.ApplyT(func(v GetFrontendsResult) string { return v.LbId }).(pulumi.StringOutput)
}

func (o GetFrontendsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFrontendsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetFrontendsResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetFrontendsResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o GetFrontendsResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetFrontendsResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o GetFrontendsResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetFrontendsResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFrontendsResultOutput{})
}
