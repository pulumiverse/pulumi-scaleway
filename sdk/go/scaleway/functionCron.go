// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// The `functions.Cron` resource allows you to create and manage CRON triggers for Scaleway [Serverless Functions](https://www.scaleway.com/en/docs/serverless/functions/).
//
// Refer to the Functions CRON triggers [documentation](https://www.scaleway.com/en/docs/serverless/functions/how-to/add-trigger-to-a-function/) and [API documentation](https://www.scaleway.com/en/developers/api/serverless-functions/#path-triggers-list-all-triggers) for more information.
//
// ## Example Usage
//
// The following command allows you to add a CRON trigger to a Serverless Function.
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/functions"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := functions.NewNamespace(ctx, "main", &functions.NamespaceArgs{
//				Name: pulumi.String("test-cron"),
//			})
//			if err != nil {
//				return err
//			}
//			mainFunction, err := functions.NewFunction(ctx, "main", &functions.FunctionArgs{
//				Name:        pulumi.String("test-cron"),
//				NamespaceId: main.ID(),
//				Runtime:     pulumi.String("node14"),
//				Privacy:     pulumi.String("private"),
//				Handler:     pulumi.String("handler.handle"),
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"test": "scw",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			_, err = functions.NewCron(ctx, "main", &functions.CronArgs{
//				Name:       pulumi.String("test-cron"),
//				FunctionId: mainFunction.ID(),
//				Schedule:   pulumi.String("0 0 * * *"),
//				Args:       pulumi.String(json0),
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON1, err := json.Marshal(map[string]interface{}{
//				"my_var": "terraform",
//			})
//			if err != nil {
//				return err
//			}
//			json1 := string(tmpJSON1)
//			_, err = functions.NewCron(ctx, "func", &functions.CronArgs{
//				FunctionId: mainFunction.ID(),
//				Schedule:   pulumi.String("0 1 * * *"),
//				Args:       pulumi.String(json1),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Function Cron can be imported using `{region}/{id}`, as shown below:
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/functionCron:FunctionCron main fr-par/11111111-1111-1111-1111-111111111111
// ```
//
// Deprecated: scaleway.index/functioncron.FunctionCron has been deprecated in favor of scaleway.functions/cron.Cron
type FunctionCron struct {
	pulumi.CustomResourceState

	// The key-value mapping to define arguments that will be passed to your function’s event object
	Args pulumi.StringOutput `pulumi:"args"`
	// The unique identifier of the function to link to your CRON trigger.
	FunctionId pulumi.StringOutput `pulumi:"functionId"`
	// The name of the function CRON trigger. If not provided, a random name is generated.
	Name pulumi.StringOutput `pulumi:"name"`
	// `region`) The region
	// in which the function was created.
	Region pulumi.StringOutput `pulumi:"region"`
	// CRON format string (refer to the [CRON schedule reference](https://www.scaleway.com/en/docs/serverless/functions/reference-content/cron-schedules/) for more information).
	Schedule pulumi.StringOutput `pulumi:"schedule"`
	// The CRON status.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewFunctionCron registers a new resource with the given unique name, arguments, and options.
func NewFunctionCron(ctx *pulumi.Context,
	name string, args *FunctionCronArgs, opts ...pulumi.ResourceOption) (*FunctionCron, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Args == nil {
		return nil, errors.New("invalid value for required argument 'Args'")
	}
	if args.FunctionId == nil {
		return nil, errors.New("invalid value for required argument 'FunctionId'")
	}
	if args.Schedule == nil {
		return nil, errors.New("invalid value for required argument 'Schedule'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FunctionCron
	err := ctx.RegisterResource("scaleway:index/functionCron:FunctionCron", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunctionCron gets an existing FunctionCron resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunctionCron(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionCronState, opts ...pulumi.ResourceOption) (*FunctionCron, error) {
	var resource FunctionCron
	err := ctx.ReadResource("scaleway:index/functionCron:FunctionCron", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FunctionCron resources.
type functionCronState struct {
	// The key-value mapping to define arguments that will be passed to your function’s event object
	Args *string `pulumi:"args"`
	// The unique identifier of the function to link to your CRON trigger.
	FunctionId *string `pulumi:"functionId"`
	// The name of the function CRON trigger. If not provided, a random name is generated.
	Name *string `pulumi:"name"`
	// `region`) The region
	// in which the function was created.
	Region *string `pulumi:"region"`
	// CRON format string (refer to the [CRON schedule reference](https://www.scaleway.com/en/docs/serverless/functions/reference-content/cron-schedules/) for more information).
	Schedule *string `pulumi:"schedule"`
	// The CRON status.
	Status *string `pulumi:"status"`
}

type FunctionCronState struct {
	// The key-value mapping to define arguments that will be passed to your function’s event object
	Args pulumi.StringPtrInput
	// The unique identifier of the function to link to your CRON trigger.
	FunctionId pulumi.StringPtrInput
	// The name of the function CRON trigger. If not provided, a random name is generated.
	Name pulumi.StringPtrInput
	// `region`) The region
	// in which the function was created.
	Region pulumi.StringPtrInput
	// CRON format string (refer to the [CRON schedule reference](https://www.scaleway.com/en/docs/serverless/functions/reference-content/cron-schedules/) for more information).
	Schedule pulumi.StringPtrInput
	// The CRON status.
	Status pulumi.StringPtrInput
}

func (FunctionCronState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionCronState)(nil)).Elem()
}

type functionCronArgs struct {
	// The key-value mapping to define arguments that will be passed to your function’s event object
	Args string `pulumi:"args"`
	// The unique identifier of the function to link to your CRON trigger.
	FunctionId string `pulumi:"functionId"`
	// The name of the function CRON trigger. If not provided, a random name is generated.
	Name *string `pulumi:"name"`
	// `region`) The region
	// in which the function was created.
	Region *string `pulumi:"region"`
	// CRON format string (refer to the [CRON schedule reference](https://www.scaleway.com/en/docs/serverless/functions/reference-content/cron-schedules/) for more information).
	Schedule string `pulumi:"schedule"`
}

// The set of arguments for constructing a FunctionCron resource.
type FunctionCronArgs struct {
	// The key-value mapping to define arguments that will be passed to your function’s event object
	Args pulumi.StringInput
	// The unique identifier of the function to link to your CRON trigger.
	FunctionId pulumi.StringInput
	// The name of the function CRON trigger. If not provided, a random name is generated.
	Name pulumi.StringPtrInput
	// `region`) The region
	// in which the function was created.
	Region pulumi.StringPtrInput
	// CRON format string (refer to the [CRON schedule reference](https://www.scaleway.com/en/docs/serverless/functions/reference-content/cron-schedules/) for more information).
	Schedule pulumi.StringInput
}

func (FunctionCronArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionCronArgs)(nil)).Elem()
}

type FunctionCronInput interface {
	pulumi.Input

	ToFunctionCronOutput() FunctionCronOutput
	ToFunctionCronOutputWithContext(ctx context.Context) FunctionCronOutput
}

func (*FunctionCron) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionCron)(nil)).Elem()
}

func (i *FunctionCron) ToFunctionCronOutput() FunctionCronOutput {
	return i.ToFunctionCronOutputWithContext(context.Background())
}

func (i *FunctionCron) ToFunctionCronOutputWithContext(ctx context.Context) FunctionCronOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionCronOutput)
}

// FunctionCronArrayInput is an input type that accepts FunctionCronArray and FunctionCronArrayOutput values.
// You can construct a concrete instance of `FunctionCronArrayInput` via:
//
//	FunctionCronArray{ FunctionCronArgs{...} }
type FunctionCronArrayInput interface {
	pulumi.Input

	ToFunctionCronArrayOutput() FunctionCronArrayOutput
	ToFunctionCronArrayOutputWithContext(context.Context) FunctionCronArrayOutput
}

type FunctionCronArray []FunctionCronInput

func (FunctionCronArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FunctionCron)(nil)).Elem()
}

func (i FunctionCronArray) ToFunctionCronArrayOutput() FunctionCronArrayOutput {
	return i.ToFunctionCronArrayOutputWithContext(context.Background())
}

func (i FunctionCronArray) ToFunctionCronArrayOutputWithContext(ctx context.Context) FunctionCronArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionCronArrayOutput)
}

// FunctionCronMapInput is an input type that accepts FunctionCronMap and FunctionCronMapOutput values.
// You can construct a concrete instance of `FunctionCronMapInput` via:
//
//	FunctionCronMap{ "key": FunctionCronArgs{...} }
type FunctionCronMapInput interface {
	pulumi.Input

	ToFunctionCronMapOutput() FunctionCronMapOutput
	ToFunctionCronMapOutputWithContext(context.Context) FunctionCronMapOutput
}

type FunctionCronMap map[string]FunctionCronInput

func (FunctionCronMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FunctionCron)(nil)).Elem()
}

func (i FunctionCronMap) ToFunctionCronMapOutput() FunctionCronMapOutput {
	return i.ToFunctionCronMapOutputWithContext(context.Background())
}

func (i FunctionCronMap) ToFunctionCronMapOutputWithContext(ctx context.Context) FunctionCronMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionCronMapOutput)
}

type FunctionCronOutput struct{ *pulumi.OutputState }

func (FunctionCronOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionCron)(nil)).Elem()
}

func (o FunctionCronOutput) ToFunctionCronOutput() FunctionCronOutput {
	return o
}

func (o FunctionCronOutput) ToFunctionCronOutputWithContext(ctx context.Context) FunctionCronOutput {
	return o
}

// The key-value mapping to define arguments that will be passed to your function’s event object
func (o FunctionCronOutput) Args() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionCron) pulumi.StringOutput { return v.Args }).(pulumi.StringOutput)
}

// The unique identifier of the function to link to your CRON trigger.
func (o FunctionCronOutput) FunctionId() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionCron) pulumi.StringOutput { return v.FunctionId }).(pulumi.StringOutput)
}

// The name of the function CRON trigger. If not provided, a random name is generated.
func (o FunctionCronOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionCron) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// `region`) The region
// in which the function was created.
func (o FunctionCronOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionCron) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// CRON format string (refer to the [CRON schedule reference](https://www.scaleway.com/en/docs/serverless/functions/reference-content/cron-schedules/) for more information).
func (o FunctionCronOutput) Schedule() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionCron) pulumi.StringOutput { return v.Schedule }).(pulumi.StringOutput)
}

// The CRON status.
func (o FunctionCronOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionCron) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type FunctionCronArrayOutput struct{ *pulumi.OutputState }

func (FunctionCronArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FunctionCron)(nil)).Elem()
}

func (o FunctionCronArrayOutput) ToFunctionCronArrayOutput() FunctionCronArrayOutput {
	return o
}

func (o FunctionCronArrayOutput) ToFunctionCronArrayOutputWithContext(ctx context.Context) FunctionCronArrayOutput {
	return o
}

func (o FunctionCronArrayOutput) Index(i pulumi.IntInput) FunctionCronOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FunctionCron {
		return vs[0].([]*FunctionCron)[vs[1].(int)]
	}).(FunctionCronOutput)
}

type FunctionCronMapOutput struct{ *pulumi.OutputState }

func (FunctionCronMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FunctionCron)(nil)).Elem()
}

func (o FunctionCronMapOutput) ToFunctionCronMapOutput() FunctionCronMapOutput {
	return o
}

func (o FunctionCronMapOutput) ToFunctionCronMapOutputWithContext(ctx context.Context) FunctionCronMapOutput {
	return o
}

func (o FunctionCronMapOutput) MapIndex(k pulumi.StringInput) FunctionCronOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FunctionCron {
		return vs[0].(map[string]*FunctionCron)[vs[1].(string)]
	}).(FunctionCronOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionCronInput)(nil)).Elem(), &FunctionCron{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionCronArrayInput)(nil)).Elem(), FunctionCronArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionCronMapInput)(nil)).Elem(), FunctionCronMap{})
	pulumi.RegisterOutputType(FunctionCronOutput{})
	pulumi.RegisterOutputType(FunctionCronArrayOutput{})
	pulumi.RegisterOutputType(FunctionCronMapOutput{})
}
