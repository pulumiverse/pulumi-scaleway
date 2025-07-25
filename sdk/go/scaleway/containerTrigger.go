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

// The `containers.Trigger` resource allows you to create and manage triggers for Scaleway [Serverless Containers](https://www.scaleway.com/en/docs/serverless/containers/).
//
// Refer to the Containers triggers [documentation](https://www.scaleway.com/en/docs/serverless/containers/how-to/add-trigger-to-a-container/) and [API documentation](https://www.scaleway.com/en/developers/api/serverless-containers/#path-triggers-list-all-triggers) for more information.
//
// ## Example Usage
//
// ### SQS
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/containers"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := containers.NewTrigger(ctx, "main", &containers.TriggerArgs{
//				ContainerId: pulumi.Any(mainScalewayContainer.Id),
//				Name:        pulumi.String("my-trigger"),
//				Sqs: &containers.TriggerSqsArgs{
//					ProjectId: pulumi.Any(mainScalewayMnqSqs.ProjectId),
//					Queue:     pulumi.String("MyQueue"),
//					Region:    pulumi.Any(mainScalewayMnqSqs.Region),
//				},
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
// ### NATS
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/containers"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := containers.NewTrigger(ctx, "main", &containers.TriggerArgs{
//				ContainerId: pulumi.Any(mainScalewayContainer.Id),
//				Name:        pulumi.String("my-trigger"),
//				Nats: &containers.TriggerNatsArgs{
//					AccountId: pulumi.Any(mainScalewayMnqNatsAccount.Id),
//					Subject:   pulumi.String("MySubject"),
//					Region:    pulumi.Any(mainScalewayMnqNatsAccount.Region),
//				},
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
// Container Triggers can be imported using `{region}/{id}`, as shown below:
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/containerTrigger:ContainerTrigger main fr-par/11111111-1111-1111-1111-111111111111
// ```
//
// Deprecated: scaleway.index/containertrigger.ContainerTrigger has been deprecated in favor of scaleway.containers/trigger.Trigger
type ContainerTrigger struct {
	pulumi.CustomResourceState

	// The unique identifier of the container to create a trigger for.
	ContainerId pulumi.StringOutput `pulumi:"containerId"`
	// The description of the trigger.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The unique name of the trigger. If not provided, a random name is generated.
	Name pulumi.StringOutput `pulumi:"name"`
	// The configuration for the Scaleway NATS account used by the trigger
	Nats ContainerTriggerNatsPtrOutput `pulumi:"nats"`
	// `region`). The region in which the namespace is created.
	Region pulumi.StringOutput `pulumi:"region"`
	// The configuration of the Scaleway SQS queue used by the trigger
	Sqs ContainerTriggerSqsPtrOutput `pulumi:"sqs"`
}

// NewContainerTrigger registers a new resource with the given unique name, arguments, and options.
func NewContainerTrigger(ctx *pulumi.Context,
	name string, args *ContainerTriggerArgs, opts ...pulumi.ResourceOption) (*ContainerTrigger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContainerId == nil {
		return nil, errors.New("invalid value for required argument 'ContainerId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ContainerTrigger
	err := ctx.RegisterResource("scaleway:index/containerTrigger:ContainerTrigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerTrigger gets an existing ContainerTrigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerTriggerState, opts ...pulumi.ResourceOption) (*ContainerTrigger, error) {
	var resource ContainerTrigger
	err := ctx.ReadResource("scaleway:index/containerTrigger:ContainerTrigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerTrigger resources.
type containerTriggerState struct {
	// The unique identifier of the container to create a trigger for.
	ContainerId *string `pulumi:"containerId"`
	// The description of the trigger.
	Description *string `pulumi:"description"`
	// The unique name of the trigger. If not provided, a random name is generated.
	Name *string `pulumi:"name"`
	// The configuration for the Scaleway NATS account used by the trigger
	Nats *ContainerTriggerNats `pulumi:"nats"`
	// `region`). The region in which the namespace is created.
	Region *string `pulumi:"region"`
	// The configuration of the Scaleway SQS queue used by the trigger
	Sqs *ContainerTriggerSqs `pulumi:"sqs"`
}

type ContainerTriggerState struct {
	// The unique identifier of the container to create a trigger for.
	ContainerId pulumi.StringPtrInput
	// The description of the trigger.
	Description pulumi.StringPtrInput
	// The unique name of the trigger. If not provided, a random name is generated.
	Name pulumi.StringPtrInput
	// The configuration for the Scaleway NATS account used by the trigger
	Nats ContainerTriggerNatsPtrInput
	// `region`). The region in which the namespace is created.
	Region pulumi.StringPtrInput
	// The configuration of the Scaleway SQS queue used by the trigger
	Sqs ContainerTriggerSqsPtrInput
}

func (ContainerTriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerTriggerState)(nil)).Elem()
}

type containerTriggerArgs struct {
	// The unique identifier of the container to create a trigger for.
	ContainerId string `pulumi:"containerId"`
	// The description of the trigger.
	Description *string `pulumi:"description"`
	// The unique name of the trigger. If not provided, a random name is generated.
	Name *string `pulumi:"name"`
	// The configuration for the Scaleway NATS account used by the trigger
	Nats *ContainerTriggerNats `pulumi:"nats"`
	// `region`). The region in which the namespace is created.
	Region *string `pulumi:"region"`
	// The configuration of the Scaleway SQS queue used by the trigger
	Sqs *ContainerTriggerSqs `pulumi:"sqs"`
}

// The set of arguments for constructing a ContainerTrigger resource.
type ContainerTriggerArgs struct {
	// The unique identifier of the container to create a trigger for.
	ContainerId pulumi.StringInput
	// The description of the trigger.
	Description pulumi.StringPtrInput
	// The unique name of the trigger. If not provided, a random name is generated.
	Name pulumi.StringPtrInput
	// The configuration for the Scaleway NATS account used by the trigger
	Nats ContainerTriggerNatsPtrInput
	// `region`). The region in which the namespace is created.
	Region pulumi.StringPtrInput
	// The configuration of the Scaleway SQS queue used by the trigger
	Sqs ContainerTriggerSqsPtrInput
}

func (ContainerTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerTriggerArgs)(nil)).Elem()
}

type ContainerTriggerInput interface {
	pulumi.Input

	ToContainerTriggerOutput() ContainerTriggerOutput
	ToContainerTriggerOutputWithContext(ctx context.Context) ContainerTriggerOutput
}

func (*ContainerTrigger) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerTrigger)(nil)).Elem()
}

func (i *ContainerTrigger) ToContainerTriggerOutput() ContainerTriggerOutput {
	return i.ToContainerTriggerOutputWithContext(context.Background())
}

func (i *ContainerTrigger) ToContainerTriggerOutputWithContext(ctx context.Context) ContainerTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerTriggerOutput)
}

// ContainerTriggerArrayInput is an input type that accepts ContainerTriggerArray and ContainerTriggerArrayOutput values.
// You can construct a concrete instance of `ContainerTriggerArrayInput` via:
//
//	ContainerTriggerArray{ ContainerTriggerArgs{...} }
type ContainerTriggerArrayInput interface {
	pulumi.Input

	ToContainerTriggerArrayOutput() ContainerTriggerArrayOutput
	ToContainerTriggerArrayOutputWithContext(context.Context) ContainerTriggerArrayOutput
}

type ContainerTriggerArray []ContainerTriggerInput

func (ContainerTriggerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerTrigger)(nil)).Elem()
}

func (i ContainerTriggerArray) ToContainerTriggerArrayOutput() ContainerTriggerArrayOutput {
	return i.ToContainerTriggerArrayOutputWithContext(context.Background())
}

func (i ContainerTriggerArray) ToContainerTriggerArrayOutputWithContext(ctx context.Context) ContainerTriggerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerTriggerArrayOutput)
}

// ContainerTriggerMapInput is an input type that accepts ContainerTriggerMap and ContainerTriggerMapOutput values.
// You can construct a concrete instance of `ContainerTriggerMapInput` via:
//
//	ContainerTriggerMap{ "key": ContainerTriggerArgs{...} }
type ContainerTriggerMapInput interface {
	pulumi.Input

	ToContainerTriggerMapOutput() ContainerTriggerMapOutput
	ToContainerTriggerMapOutputWithContext(context.Context) ContainerTriggerMapOutput
}

type ContainerTriggerMap map[string]ContainerTriggerInput

func (ContainerTriggerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerTrigger)(nil)).Elem()
}

func (i ContainerTriggerMap) ToContainerTriggerMapOutput() ContainerTriggerMapOutput {
	return i.ToContainerTriggerMapOutputWithContext(context.Background())
}

func (i ContainerTriggerMap) ToContainerTriggerMapOutputWithContext(ctx context.Context) ContainerTriggerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerTriggerMapOutput)
}

type ContainerTriggerOutput struct{ *pulumi.OutputState }

func (ContainerTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerTrigger)(nil)).Elem()
}

func (o ContainerTriggerOutput) ToContainerTriggerOutput() ContainerTriggerOutput {
	return o
}

func (o ContainerTriggerOutput) ToContainerTriggerOutputWithContext(ctx context.Context) ContainerTriggerOutput {
	return o
}

// The unique identifier of the container to create a trigger for.
func (o ContainerTriggerOutput) ContainerId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerTrigger) pulumi.StringOutput { return v.ContainerId }).(pulumi.StringOutput)
}

// The description of the trigger.
func (o ContainerTriggerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerTrigger) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The unique name of the trigger. If not provided, a random name is generated.
func (o ContainerTriggerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerTrigger) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The configuration for the Scaleway NATS account used by the trigger
func (o ContainerTriggerOutput) Nats() ContainerTriggerNatsPtrOutput {
	return o.ApplyT(func(v *ContainerTrigger) ContainerTriggerNatsPtrOutput { return v.Nats }).(ContainerTriggerNatsPtrOutput)
}

// `region`). The region in which the namespace is created.
func (o ContainerTriggerOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerTrigger) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The configuration of the Scaleway SQS queue used by the trigger
func (o ContainerTriggerOutput) Sqs() ContainerTriggerSqsPtrOutput {
	return o.ApplyT(func(v *ContainerTrigger) ContainerTriggerSqsPtrOutput { return v.Sqs }).(ContainerTriggerSqsPtrOutput)
}

type ContainerTriggerArrayOutput struct{ *pulumi.OutputState }

func (ContainerTriggerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerTrigger)(nil)).Elem()
}

func (o ContainerTriggerArrayOutput) ToContainerTriggerArrayOutput() ContainerTriggerArrayOutput {
	return o
}

func (o ContainerTriggerArrayOutput) ToContainerTriggerArrayOutputWithContext(ctx context.Context) ContainerTriggerArrayOutput {
	return o
}

func (o ContainerTriggerArrayOutput) Index(i pulumi.IntInput) ContainerTriggerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ContainerTrigger {
		return vs[0].([]*ContainerTrigger)[vs[1].(int)]
	}).(ContainerTriggerOutput)
}

type ContainerTriggerMapOutput struct{ *pulumi.OutputState }

func (ContainerTriggerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerTrigger)(nil)).Elem()
}

func (o ContainerTriggerMapOutput) ToContainerTriggerMapOutput() ContainerTriggerMapOutput {
	return o
}

func (o ContainerTriggerMapOutput) ToContainerTriggerMapOutputWithContext(ctx context.Context) ContainerTriggerMapOutput {
	return o
}

func (o ContainerTriggerMapOutput) MapIndex(k pulumi.StringInput) ContainerTriggerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ContainerTrigger {
		return vs[0].(map[string]*ContainerTrigger)[vs[1].(string)]
	}).(ContainerTriggerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerTriggerInput)(nil)).Elem(), &ContainerTrigger{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerTriggerArrayInput)(nil)).Elem(), ContainerTriggerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerTriggerMapInput)(nil)).Elem(), ContainerTriggerMap{})
	pulumi.RegisterOutputType(ContainerTriggerOutput{})
	pulumi.RegisterOutputType(ContainerTriggerArrayOutput{})
	pulumi.RegisterOutputType(ContainerTriggerMapOutput{})
}
