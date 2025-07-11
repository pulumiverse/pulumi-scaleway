// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancers

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about multiple Load Balancer routes.
//
// For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/how-to/create-manage-routes/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-route).
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
//			// Find routes that share the same frontend ID
//			_, err := loadbalancers.GetRoutes(ctx, &loadbalancers.GetRoutesArgs{
//				FrontendId: pulumi.StringRef(frt01.Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			// Find routes by frontend ID and zone
//			_, err = loadbalancers.GetRoutes(ctx, &loadbalancers.GetRoutesArgs{
//				FrontendId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//				Zone:       pulumi.StringRef("fr-par-2"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetRoutes(ctx *pulumi.Context, args *GetRoutesArgs, opts ...pulumi.InvokeOption) (*GetRoutesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRoutesResult
	err := ctx.Invoke("scaleway:loadbalancers/getRoutes:getRoutes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRoutes.
type GetRoutesArgs struct {
	// The frontend ID (the origin of the redirection), to filter for. Routes with a matching frontend ID are listed.
	FrontendId *string `pulumi:"frontendId"`
	ProjectId  *string `pulumi:"projectId"`
	// `zone`) The zone in which the routes exist.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getRoutes.
type GetRoutesResult struct {
	FrontendId *string `pulumi:"frontendId"`
	// The provider-assigned unique ID for this managed resource.
	Id             string `pulumi:"id"`
	OrganizationId string `pulumi:"organizationId"`
	ProjectId      string `pulumi:"projectId"`
	// List of retrieved routes
	Routes []GetRoutesRoute `pulumi:"routes"`
	Zone   string           `pulumi:"zone"`
}

func GetRoutesOutput(ctx *pulumi.Context, args GetRoutesOutputArgs, opts ...pulumi.InvokeOption) GetRoutesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetRoutesResultOutput, error) {
			args := v.(GetRoutesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway:loadbalancers/getRoutes:getRoutes", args, GetRoutesResultOutput{}, options).(GetRoutesResultOutput), nil
		}).(GetRoutesResultOutput)
}

// A collection of arguments for invoking getRoutes.
type GetRoutesOutputArgs struct {
	// The frontend ID (the origin of the redirection), to filter for. Routes with a matching frontend ID are listed.
	FrontendId pulumi.StringPtrInput `pulumi:"frontendId"`
	ProjectId  pulumi.StringPtrInput `pulumi:"projectId"`
	// `zone`) The zone in which the routes exist.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (GetRoutesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRoutesArgs)(nil)).Elem()
}

// A collection of values returned by getRoutes.
type GetRoutesResultOutput struct{ *pulumi.OutputState }

func (GetRoutesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRoutesResult)(nil)).Elem()
}

func (o GetRoutesResultOutput) ToGetRoutesResultOutput() GetRoutesResultOutput {
	return o
}

func (o GetRoutesResultOutput) ToGetRoutesResultOutputWithContext(ctx context.Context) GetRoutesResultOutput {
	return o
}

func (o GetRoutesResultOutput) FrontendId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRoutesResult) *string { return v.FrontendId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRoutesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRoutesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRoutesResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetRoutesResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o GetRoutesResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetRoutesResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// List of retrieved routes
func (o GetRoutesResultOutput) Routes() GetRoutesRouteArrayOutput {
	return o.ApplyT(func(v GetRoutesResult) []GetRoutesRoute { return v.Routes }).(GetRoutesRouteArrayOutput)
}

func (o GetRoutesResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetRoutesResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRoutesResultOutput{})
}
