// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

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
//
// Deprecated: scaleway.index/getlbroutes.getLbRoutes has been deprecated in favor of scaleway.loadbalancers/getroutes.getRoutes
func GetLbRoutes(ctx *pulumi.Context, args *GetLbRoutesArgs, opts ...pulumi.InvokeOption) (*GetLbRoutesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLbRoutesResult
	err := ctx.Invoke("scaleway:index/getLbRoutes:getLbRoutes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLbRoutes.
type GetLbRoutesArgs struct {
	// The frontend ID (the origin of the redirection), to filter for. Routes with a matching frontend ID are listed.
	FrontendId *string `pulumi:"frontendId"`
	ProjectId  *string `pulumi:"projectId"`
	// `zone`) The zone in which the routes exist.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getLbRoutes.
type GetLbRoutesResult struct {
	FrontendId *string `pulumi:"frontendId"`
	// The provider-assigned unique ID for this managed resource.
	Id             string `pulumi:"id"`
	OrganizationId string `pulumi:"organizationId"`
	ProjectId      string `pulumi:"projectId"`
	// List of retrieved routes
	Routes []GetLbRoutesRoute `pulumi:"routes"`
	Zone   string             `pulumi:"zone"`
}

func GetLbRoutesOutput(ctx *pulumi.Context, args GetLbRoutesOutputArgs, opts ...pulumi.InvokeOption) GetLbRoutesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetLbRoutesResultOutput, error) {
			args := v.(GetLbRoutesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway:index/getLbRoutes:getLbRoutes", args, GetLbRoutesResultOutput{}, options).(GetLbRoutesResultOutput), nil
		}).(GetLbRoutesResultOutput)
}

// A collection of arguments for invoking getLbRoutes.
type GetLbRoutesOutputArgs struct {
	// The frontend ID (the origin of the redirection), to filter for. Routes with a matching frontend ID are listed.
	FrontendId pulumi.StringPtrInput `pulumi:"frontendId"`
	ProjectId  pulumi.StringPtrInput `pulumi:"projectId"`
	// `zone`) The zone in which the routes exist.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (GetLbRoutesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLbRoutesArgs)(nil)).Elem()
}

// A collection of values returned by getLbRoutes.
type GetLbRoutesResultOutput struct{ *pulumi.OutputState }

func (GetLbRoutesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLbRoutesResult)(nil)).Elem()
}

func (o GetLbRoutesResultOutput) ToGetLbRoutesResultOutput() GetLbRoutesResultOutput {
	return o
}

func (o GetLbRoutesResultOutput) ToGetLbRoutesResultOutputWithContext(ctx context.Context) GetLbRoutesResultOutput {
	return o
}

func (o GetLbRoutesResultOutput) FrontendId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLbRoutesResult) *string { return v.FrontendId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLbRoutesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbRoutesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLbRoutesResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbRoutesResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o GetLbRoutesResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbRoutesResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// List of retrieved routes
func (o GetLbRoutesResultOutput) Routes() GetLbRoutesRouteArrayOutput {
	return o.ApplyT(func(v GetLbRoutesResult) []GetLbRoutesRoute { return v.Routes }).(GetLbRoutesRouteArrayOutput)
}

func (o GetLbRoutesResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbRoutesResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLbRoutesResultOutput{})
}
