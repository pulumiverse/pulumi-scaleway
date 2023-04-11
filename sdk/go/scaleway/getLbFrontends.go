// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about multiple Load Balancer Frontends.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-scaleway/sdk/go/scaleway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := scaleway.GetLbFrontends(ctx, &scaleway.GetLbFrontendsArgs{
//				LbId: scaleway_lb.Lb01.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.GetLbFrontends(ctx, &scaleway.GetLbFrontendsArgs{
//				LbId: scaleway_lb.Lb01.Id,
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
func GetLbFrontends(ctx *pulumi.Context, args *GetLbFrontendsArgs, opts ...pulumi.InvokeOption) (*GetLbFrontendsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetLbFrontendsResult
	err := ctx.Invoke("scaleway:index/getLbFrontends:getLbFrontends", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLbFrontends.
type GetLbFrontendsArgs struct {
	// The load-balancer ID this frontend is attached to. frontends with a LB ID like it are listed.
	LbId string `pulumi:"lbId"`
	// The frontend name used as filter. Frontends with a name like it are listed.
	Name      *string `pulumi:"name"`
	ProjectId *string `pulumi:"projectId"`
	// `zone`) The zone in which frontends exist.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getLbFrontends.
type GetLbFrontendsResult struct {
	// List of found frontends
	Frontends []GetLbFrontendsFrontend `pulumi:"frontends"`
	// The provider-assigned unique ID for this managed resource.
	Id             string  `pulumi:"id"`
	LbId           string  `pulumi:"lbId"`
	Name           *string `pulumi:"name"`
	OrganizationId string  `pulumi:"organizationId"`
	ProjectId      string  `pulumi:"projectId"`
	Zone           string  `pulumi:"zone"`
}

func GetLbFrontendsOutput(ctx *pulumi.Context, args GetLbFrontendsOutputArgs, opts ...pulumi.InvokeOption) GetLbFrontendsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLbFrontendsResult, error) {
			args := v.(GetLbFrontendsArgs)
			r, err := GetLbFrontends(ctx, &args, opts...)
			var s GetLbFrontendsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLbFrontendsResultOutput)
}

// A collection of arguments for invoking getLbFrontends.
type GetLbFrontendsOutputArgs struct {
	// The load-balancer ID this frontend is attached to. frontends with a LB ID like it are listed.
	LbId pulumi.StringInput `pulumi:"lbId"`
	// The frontend name used as filter. Frontends with a name like it are listed.
	Name      pulumi.StringPtrInput `pulumi:"name"`
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// `zone`) The zone in which frontends exist.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (GetLbFrontendsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLbFrontendsArgs)(nil)).Elem()
}

// A collection of values returned by getLbFrontends.
type GetLbFrontendsResultOutput struct{ *pulumi.OutputState }

func (GetLbFrontendsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLbFrontendsResult)(nil)).Elem()
}

func (o GetLbFrontendsResultOutput) ToGetLbFrontendsResultOutput() GetLbFrontendsResultOutput {
	return o
}

func (o GetLbFrontendsResultOutput) ToGetLbFrontendsResultOutputWithContext(ctx context.Context) GetLbFrontendsResultOutput {
	return o
}

// List of found frontends
func (o GetLbFrontendsResultOutput) Frontends() GetLbFrontendsFrontendArrayOutput {
	return o.ApplyT(func(v GetLbFrontendsResult) []GetLbFrontendsFrontend { return v.Frontends }).(GetLbFrontendsFrontendArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLbFrontendsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbFrontendsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLbFrontendsResultOutput) LbId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbFrontendsResult) string { return v.LbId }).(pulumi.StringOutput)
}

func (o GetLbFrontendsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLbFrontendsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetLbFrontendsResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbFrontendsResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o GetLbFrontendsResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbFrontendsResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o GetLbFrontendsResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbFrontendsResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLbFrontendsResultOutput{})
}