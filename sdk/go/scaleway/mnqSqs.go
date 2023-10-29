// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/lbrlabs/pulumi-scaleway/sdk/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Activate Scaleway Messaging and queuing SQS for a project.
// For further information please check
// our [documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/sqs-overview/)
//
// ## Examples
//
// ### Basic
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
//			_, err := scaleway.NewMnqSqs(ctx, "main", nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewMnqSqs(ctx, "forProject", &scaleway.MnqSqsArgs{
//				ProjectId: pulumi.Any(scaleway_account_project.Main.Id),
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
// SQS status can be imported using the `{region}/{project_id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/mnqSqs:MnqSqs main fr-par/11111111111111111111111111111111
//
// ```
type MnqSqs struct {
	pulumi.CustomResourceState

	// The endpoint of the SQS service for this project.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// `projectId`) The ID of the project the sqs will be enabled for.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// `region`). The region
	// in which sqs will be enabled.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewMnqSqs registers a new resource with the given unique name, arguments, and options.
func NewMnqSqs(ctx *pulumi.Context,
	name string, args *MnqSqsArgs, opts ...pulumi.ResourceOption) (*MnqSqs, error) {
	if args == nil {
		args = &MnqSqsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MnqSqs
	err := ctx.RegisterResource("scaleway:index/mnqSqs:MnqSqs", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMnqSqs gets an existing MnqSqs resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMnqSqs(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MnqSqsState, opts ...pulumi.ResourceOption) (*MnqSqs, error) {
	var resource MnqSqs
	err := ctx.ReadResource("scaleway:index/mnqSqs:MnqSqs", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MnqSqs resources.
type mnqSqsState struct {
	// The endpoint of the SQS service for this project.
	Endpoint *string `pulumi:"endpoint"`
	// `projectId`) The ID of the project the sqs will be enabled for.
	ProjectId *string `pulumi:"projectId"`
	// `region`). The region
	// in which sqs will be enabled.
	Region *string `pulumi:"region"`
}

type MnqSqsState struct {
	// The endpoint of the SQS service for this project.
	Endpoint pulumi.StringPtrInput
	// `projectId`) The ID of the project the sqs will be enabled for.
	ProjectId pulumi.StringPtrInput
	// `region`). The region
	// in which sqs will be enabled.
	Region pulumi.StringPtrInput
}

func (MnqSqsState) ElementType() reflect.Type {
	return reflect.TypeOf((*mnqSqsState)(nil)).Elem()
}

type mnqSqsArgs struct {
	// `projectId`) The ID of the project the sqs will be enabled for.
	ProjectId *string `pulumi:"projectId"`
	// `region`). The region
	// in which sqs will be enabled.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a MnqSqs resource.
type MnqSqsArgs struct {
	// `projectId`) The ID of the project the sqs will be enabled for.
	ProjectId pulumi.StringPtrInput
	// `region`). The region
	// in which sqs will be enabled.
	Region pulumi.StringPtrInput
}

func (MnqSqsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mnqSqsArgs)(nil)).Elem()
}

type MnqSqsInput interface {
	pulumi.Input

	ToMnqSqsOutput() MnqSqsOutput
	ToMnqSqsOutputWithContext(ctx context.Context) MnqSqsOutput
}

func (*MnqSqs) ElementType() reflect.Type {
	return reflect.TypeOf((**MnqSqs)(nil)).Elem()
}

func (i *MnqSqs) ToMnqSqsOutput() MnqSqsOutput {
	return i.ToMnqSqsOutputWithContext(context.Background())
}

func (i *MnqSqs) ToMnqSqsOutputWithContext(ctx context.Context) MnqSqsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MnqSqsOutput)
}

func (i *MnqSqs) ToOutput(ctx context.Context) pulumix.Output[*MnqSqs] {
	return pulumix.Output[*MnqSqs]{
		OutputState: i.ToMnqSqsOutputWithContext(ctx).OutputState,
	}
}

// MnqSqsArrayInput is an input type that accepts MnqSqsArray and MnqSqsArrayOutput values.
// You can construct a concrete instance of `MnqSqsArrayInput` via:
//
//	MnqSqsArray{ MnqSqsArgs{...} }
type MnqSqsArrayInput interface {
	pulumi.Input

	ToMnqSqsArrayOutput() MnqSqsArrayOutput
	ToMnqSqsArrayOutputWithContext(context.Context) MnqSqsArrayOutput
}

type MnqSqsArray []MnqSqsInput

func (MnqSqsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MnqSqs)(nil)).Elem()
}

func (i MnqSqsArray) ToMnqSqsArrayOutput() MnqSqsArrayOutput {
	return i.ToMnqSqsArrayOutputWithContext(context.Background())
}

func (i MnqSqsArray) ToMnqSqsArrayOutputWithContext(ctx context.Context) MnqSqsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MnqSqsArrayOutput)
}

func (i MnqSqsArray) ToOutput(ctx context.Context) pulumix.Output[[]*MnqSqs] {
	return pulumix.Output[[]*MnqSqs]{
		OutputState: i.ToMnqSqsArrayOutputWithContext(ctx).OutputState,
	}
}

// MnqSqsMapInput is an input type that accepts MnqSqsMap and MnqSqsMapOutput values.
// You can construct a concrete instance of `MnqSqsMapInput` via:
//
//	MnqSqsMap{ "key": MnqSqsArgs{...} }
type MnqSqsMapInput interface {
	pulumi.Input

	ToMnqSqsMapOutput() MnqSqsMapOutput
	ToMnqSqsMapOutputWithContext(context.Context) MnqSqsMapOutput
}

type MnqSqsMap map[string]MnqSqsInput

func (MnqSqsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MnqSqs)(nil)).Elem()
}

func (i MnqSqsMap) ToMnqSqsMapOutput() MnqSqsMapOutput {
	return i.ToMnqSqsMapOutputWithContext(context.Background())
}

func (i MnqSqsMap) ToMnqSqsMapOutputWithContext(ctx context.Context) MnqSqsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MnqSqsMapOutput)
}

func (i MnqSqsMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*MnqSqs] {
	return pulumix.Output[map[string]*MnqSqs]{
		OutputState: i.ToMnqSqsMapOutputWithContext(ctx).OutputState,
	}
}

type MnqSqsOutput struct{ *pulumi.OutputState }

func (MnqSqsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MnqSqs)(nil)).Elem()
}

func (o MnqSqsOutput) ToMnqSqsOutput() MnqSqsOutput {
	return o
}

func (o MnqSqsOutput) ToMnqSqsOutputWithContext(ctx context.Context) MnqSqsOutput {
	return o
}

func (o MnqSqsOutput) ToOutput(ctx context.Context) pulumix.Output[*MnqSqs] {
	return pulumix.Output[*MnqSqs]{
		OutputState: o.OutputState,
	}
}

// The endpoint of the SQS service for this project.
func (o MnqSqsOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqSqs) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the sqs will be enabled for.
func (o MnqSqsOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqSqs) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// `region`). The region
// in which sqs will be enabled.
func (o MnqSqsOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqSqs) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type MnqSqsArrayOutput struct{ *pulumi.OutputState }

func (MnqSqsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MnqSqs)(nil)).Elem()
}

func (o MnqSqsArrayOutput) ToMnqSqsArrayOutput() MnqSqsArrayOutput {
	return o
}

func (o MnqSqsArrayOutput) ToMnqSqsArrayOutputWithContext(ctx context.Context) MnqSqsArrayOutput {
	return o
}

func (o MnqSqsArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*MnqSqs] {
	return pulumix.Output[[]*MnqSqs]{
		OutputState: o.OutputState,
	}
}

func (o MnqSqsArrayOutput) Index(i pulumi.IntInput) MnqSqsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MnqSqs {
		return vs[0].([]*MnqSqs)[vs[1].(int)]
	}).(MnqSqsOutput)
}

type MnqSqsMapOutput struct{ *pulumi.OutputState }

func (MnqSqsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MnqSqs)(nil)).Elem()
}

func (o MnqSqsMapOutput) ToMnqSqsMapOutput() MnqSqsMapOutput {
	return o
}

func (o MnqSqsMapOutput) ToMnqSqsMapOutputWithContext(ctx context.Context) MnqSqsMapOutput {
	return o
}

func (o MnqSqsMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*MnqSqs] {
	return pulumix.Output[map[string]*MnqSqs]{
		OutputState: o.OutputState,
	}
}

func (o MnqSqsMapOutput) MapIndex(k pulumi.StringInput) MnqSqsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MnqSqs {
		return vs[0].(map[string]*MnqSqs)[vs[1].(string)]
	}).(MnqSqsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MnqSqsInput)(nil)).Elem(), &MnqSqs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MnqSqsArrayInput)(nil)).Elem(), MnqSqsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MnqSqsMapInput)(nil)).Elem(), MnqSqsMap{})
	pulumi.RegisterOutputType(MnqSqsOutput{})
	pulumi.RegisterOutputType(MnqSqsArrayOutput{})
	pulumi.RegisterOutputType(MnqSqsMapOutput{})
}