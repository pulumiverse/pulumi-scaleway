// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticmetal

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about a baremetal operating system.
// For more information, see the [API documentation](https://www.scaleway.com/en/developers/api/elastic-metal/#path-os-list-available-oses).
//
// You can also use the [scaleway-cli](https://github.com/scaleway/scaleway-cli) with `scw baremetal os list` to list all available operating systems.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/elasticmetal"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// Get info by os name and version
//			_, err := elasticmetal.GetOs(ctx, &elasticmetal.GetOsArgs{
//				Name:    pulumi.StringRef("Ubuntu"),
//				Version: pulumi.StringRef("20.04 LTS (Focal Fossa)"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			// Get info by os id
//			_, err = elasticmetal.GetOs(ctx, &elasticmetal.GetOsArgs{
//				OsId: pulumi.StringRef("03b7f4ba-a6a1-4305-984e-b54fafbf1681"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetOs(ctx *pulumi.Context, args *GetOsArgs, opts ...pulumi.InvokeOption) (*GetOsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetOsResult
	err := ctx.Invoke("scaleway:elasticmetal/getOs:getOs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOs.
type GetOsArgs struct {
	// The os name. Only one of `name` and `osId` should be specified.
	Name *string `pulumi:"name"`
	// The operating system id. Only one of `name` and `osId` should be specified.
	OsId *string `pulumi:"osId"`
	// The os version.
	Version *string `pulumi:"version"`
	// `zone`) The zone in which the os exists.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getOs.
type GetOsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id      string  `pulumi:"id"`
	Name    *string `pulumi:"name"`
	OsId    *string `pulumi:"osId"`
	Version *string `pulumi:"version"`
	Zone    string  `pulumi:"zone"`
}

func GetOsOutput(ctx *pulumi.Context, args GetOsOutputArgs, opts ...pulumi.InvokeOption) GetOsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetOsResultOutput, error) {
			args := v.(GetOsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway:elasticmetal/getOs:getOs", args, GetOsResultOutput{}, options).(GetOsResultOutput), nil
		}).(GetOsResultOutput)
}

// A collection of arguments for invoking getOs.
type GetOsOutputArgs struct {
	// The os name. Only one of `name` and `osId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The operating system id. Only one of `name` and `osId` should be specified.
	OsId pulumi.StringPtrInput `pulumi:"osId"`
	// The os version.
	Version pulumi.StringPtrInput `pulumi:"version"`
	// `zone`) The zone in which the os exists.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (GetOsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOsArgs)(nil)).Elem()
}

// A collection of values returned by getOs.
type GetOsResultOutput struct{ *pulumi.OutputState }

func (GetOsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOsResult)(nil)).Elem()
}

func (o GetOsResultOutput) ToGetOsResultOutput() GetOsResultOutput {
	return o
}

func (o GetOsResultOutput) ToGetOsResultOutputWithContext(ctx context.Context) GetOsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetOsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetOsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetOsResultOutput) OsId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOsResult) *string { return v.OsId }).(pulumi.StringPtrOutput)
}

func (o GetOsResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOsResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func (o GetOsResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetOsResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOsResultOutput{})
}
