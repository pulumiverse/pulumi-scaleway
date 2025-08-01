// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticmetal

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about a baremetal option.
// For more information, see the [API documentation](https://developers.scaleway.com/en/products/baremetal/api).
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
//			// Get info by option name
//			_, err := elasticmetal.GetOption(ctx, &elasticmetal.GetOptionArgs{
//				Name: pulumi.StringRef("Remote Access"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			// Get info by option id
//			_, err = elasticmetal.GetOption(ctx, &elasticmetal.GetOptionArgs{
//				OptionId: pulumi.StringRef("931df052-d713-4674-8b58-96a63244c8e2"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetOption(ctx *pulumi.Context, args *GetOptionArgs, opts ...pulumi.InvokeOption) (*GetOptionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetOptionResult
	err := ctx.Invoke("scaleway:elasticmetal/getOption:getOption", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOption.
type GetOptionArgs struct {
	// The option name. Only one of `name` and `optionId` should be specified.
	Name *string `pulumi:"name"`
	// The option id. Only one of `name` and `optionId` should be specified.
	OptionId *string `pulumi:"optionId"`
	// `zone`) The zone in which the option exists.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getOption.
type GetOptionResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Is false if the option could not be added or removed.
	Manageable bool `pulumi:"manageable"`
	// The name of the option.
	Name     *string `pulumi:"name"`
	OptionId *string `pulumi:"optionId"`
	Zone     string  `pulumi:"zone"`
}

func GetOptionOutput(ctx *pulumi.Context, args GetOptionOutputArgs, opts ...pulumi.InvokeOption) GetOptionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetOptionResultOutput, error) {
			args := v.(GetOptionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway:elasticmetal/getOption:getOption", args, GetOptionResultOutput{}, options).(GetOptionResultOutput), nil
		}).(GetOptionResultOutput)
}

// A collection of arguments for invoking getOption.
type GetOptionOutputArgs struct {
	// The option name. Only one of `name` and `optionId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The option id. Only one of `name` and `optionId` should be specified.
	OptionId pulumi.StringPtrInput `pulumi:"optionId"`
	// `zone`) The zone in which the option exists.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (GetOptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOptionArgs)(nil)).Elem()
}

// A collection of values returned by getOption.
type GetOptionResultOutput struct{ *pulumi.OutputState }

func (GetOptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOptionResult)(nil)).Elem()
}

func (o GetOptionResultOutput) ToGetOptionResultOutput() GetOptionResultOutput {
	return o
}

func (o GetOptionResultOutput) ToGetOptionResultOutputWithContext(ctx context.Context) GetOptionResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetOptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOptionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Is false if the option could not be added or removed.
func (o GetOptionResultOutput) Manageable() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOptionResult) bool { return v.Manageable }).(pulumi.BoolOutput)
}

// The name of the option.
func (o GetOptionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOptionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetOptionResultOutput) OptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOptionResult) *string { return v.OptionId }).(pulumi.StringPtrOutput)
}

func (o GetOptionResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetOptionResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOptionResultOutput{})
}
