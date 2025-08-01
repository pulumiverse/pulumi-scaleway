// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package containers

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
// $ pulumi import scaleway:containers/trigger:Trigger main fr-par/11111111-1111-1111-1111-111111111111
// ```
type Trigger struct {
	pulumi.CustomResourceState

	// The unique identifier of the container to create a trigger for.
	ContainerId pulumi.StringOutput `pulumi:"containerId"`
	// The description of the trigger.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The unique name of the trigger. If not provided, a random name is generated.
	Name pulumi.StringOutput `pulumi:"name"`
	// The configuration for the Scaleway NATS account used by the trigger
	Nats TriggerNatsPtrOutput `pulumi:"nats"`
	// `region`). The region in which the namespace is created.
	Region pulumi.StringOutput `pulumi:"region"`
	// The configuration of the Scaleway SQS queue used by the trigger
	Sqs TriggerSqsPtrOutput `pulumi:"sqs"`
}

// NewTrigger registers a new resource with the given unique name, arguments, and options.
func NewTrigger(ctx *pulumi.Context,
	name string, args *TriggerArgs, opts ...pulumi.ResourceOption) (*Trigger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContainerId == nil {
		return nil, errors.New("invalid value for required argument 'ContainerId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("scaleway:index/containerTrigger:ContainerTrigger"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Trigger
	err := ctx.RegisterResource("scaleway:containers/trigger:Trigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrigger gets an existing Trigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TriggerState, opts ...pulumi.ResourceOption) (*Trigger, error) {
	var resource Trigger
	err := ctx.ReadResource("scaleway:containers/trigger:Trigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Trigger resources.
type triggerState struct {
	// The unique identifier of the container to create a trigger for.
	ContainerId *string `pulumi:"containerId"`
	// The description of the trigger.
	Description *string `pulumi:"description"`
	// The unique name of the trigger. If not provided, a random name is generated.
	Name *string `pulumi:"name"`
	// The configuration for the Scaleway NATS account used by the trigger
	Nats *TriggerNats `pulumi:"nats"`
	// `region`). The region in which the namespace is created.
	Region *string `pulumi:"region"`
	// The configuration of the Scaleway SQS queue used by the trigger
	Sqs *TriggerSqs `pulumi:"sqs"`
}

type TriggerState struct {
	// The unique identifier of the container to create a trigger for.
	ContainerId pulumi.StringPtrInput
	// The description of the trigger.
	Description pulumi.StringPtrInput
	// The unique name of the trigger. If not provided, a random name is generated.
	Name pulumi.StringPtrInput
	// The configuration for the Scaleway NATS account used by the trigger
	Nats TriggerNatsPtrInput
	// `region`). The region in which the namespace is created.
	Region pulumi.StringPtrInput
	// The configuration of the Scaleway SQS queue used by the trigger
	Sqs TriggerSqsPtrInput
}

func (TriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerState)(nil)).Elem()
}

type triggerArgs struct {
	// The unique identifier of the container to create a trigger for.
	ContainerId string `pulumi:"containerId"`
	// The description of the trigger.
	Description *string `pulumi:"description"`
	// The unique name of the trigger. If not provided, a random name is generated.
	Name *string `pulumi:"name"`
	// The configuration for the Scaleway NATS account used by the trigger
	Nats *TriggerNats `pulumi:"nats"`
	// `region`). The region in which the namespace is created.
	Region *string `pulumi:"region"`
	// The configuration of the Scaleway SQS queue used by the trigger
	Sqs *TriggerSqs `pulumi:"sqs"`
}

// The set of arguments for constructing a Trigger resource.
type TriggerArgs struct {
	// The unique identifier of the container to create a trigger for.
	ContainerId pulumi.StringInput
	// The description of the trigger.
	Description pulumi.StringPtrInput
	// The unique name of the trigger. If not provided, a random name is generated.
	Name pulumi.StringPtrInput
	// The configuration for the Scaleway NATS account used by the trigger
	Nats TriggerNatsPtrInput
	// `region`). The region in which the namespace is created.
	Region pulumi.StringPtrInput
	// The configuration of the Scaleway SQS queue used by the trigger
	Sqs TriggerSqsPtrInput
}

func (TriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerArgs)(nil)).Elem()
}

type TriggerInput interface {
	pulumi.Input

	ToTriggerOutput() TriggerOutput
	ToTriggerOutputWithContext(ctx context.Context) TriggerOutput
}

func (*Trigger) ElementType() reflect.Type {
	return reflect.TypeOf((**Trigger)(nil)).Elem()
}

func (i *Trigger) ToTriggerOutput() TriggerOutput {
	return i.ToTriggerOutputWithContext(context.Background())
}

func (i *Trigger) ToTriggerOutputWithContext(ctx context.Context) TriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerOutput)
}

// TriggerArrayInput is an input type that accepts TriggerArray and TriggerArrayOutput values.
// You can construct a concrete instance of `TriggerArrayInput` via:
//
//	TriggerArray{ TriggerArgs{...} }
type TriggerArrayInput interface {
	pulumi.Input

	ToTriggerArrayOutput() TriggerArrayOutput
	ToTriggerArrayOutputWithContext(context.Context) TriggerArrayOutput
}

type TriggerArray []TriggerInput

func (TriggerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Trigger)(nil)).Elem()
}

func (i TriggerArray) ToTriggerArrayOutput() TriggerArrayOutput {
	return i.ToTriggerArrayOutputWithContext(context.Background())
}

func (i TriggerArray) ToTriggerArrayOutputWithContext(ctx context.Context) TriggerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerArrayOutput)
}

// TriggerMapInput is an input type that accepts TriggerMap and TriggerMapOutput values.
// You can construct a concrete instance of `TriggerMapInput` via:
//
//	TriggerMap{ "key": TriggerArgs{...} }
type TriggerMapInput interface {
	pulumi.Input

	ToTriggerMapOutput() TriggerMapOutput
	ToTriggerMapOutputWithContext(context.Context) TriggerMapOutput
}

type TriggerMap map[string]TriggerInput

func (TriggerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Trigger)(nil)).Elem()
}

func (i TriggerMap) ToTriggerMapOutput() TriggerMapOutput {
	return i.ToTriggerMapOutputWithContext(context.Background())
}

func (i TriggerMap) ToTriggerMapOutputWithContext(ctx context.Context) TriggerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerMapOutput)
}

type TriggerOutput struct{ *pulumi.OutputState }

func (TriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Trigger)(nil)).Elem()
}

func (o TriggerOutput) ToTriggerOutput() TriggerOutput {
	return o
}

func (o TriggerOutput) ToTriggerOutputWithContext(ctx context.Context) TriggerOutput {
	return o
}

// The unique identifier of the container to create a trigger for.
func (o TriggerOutput) ContainerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.ContainerId }).(pulumi.StringOutput)
}

// The description of the trigger.
func (o TriggerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The unique name of the trigger. If not provided, a random name is generated.
func (o TriggerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The configuration for the Scaleway NATS account used by the trigger
func (o TriggerOutput) Nats() TriggerNatsPtrOutput {
	return o.ApplyT(func(v *Trigger) TriggerNatsPtrOutput { return v.Nats }).(TriggerNatsPtrOutput)
}

// `region`). The region in which the namespace is created.
func (o TriggerOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The configuration of the Scaleway SQS queue used by the trigger
func (o TriggerOutput) Sqs() TriggerSqsPtrOutput {
	return o.ApplyT(func(v *Trigger) TriggerSqsPtrOutput { return v.Sqs }).(TriggerSqsPtrOutput)
}

type TriggerArrayOutput struct{ *pulumi.OutputState }

func (TriggerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Trigger)(nil)).Elem()
}

func (o TriggerArrayOutput) ToTriggerArrayOutput() TriggerArrayOutput {
	return o
}

func (o TriggerArrayOutput) ToTriggerArrayOutputWithContext(ctx context.Context) TriggerArrayOutput {
	return o
}

func (o TriggerArrayOutput) Index(i pulumi.IntInput) TriggerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Trigger {
		return vs[0].([]*Trigger)[vs[1].(int)]
	}).(TriggerOutput)
}

type TriggerMapOutput struct{ *pulumi.OutputState }

func (TriggerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Trigger)(nil)).Elem()
}

func (o TriggerMapOutput) ToTriggerMapOutput() TriggerMapOutput {
	return o
}

func (o TriggerMapOutput) ToTriggerMapOutputWithContext(ctx context.Context) TriggerMapOutput {
	return o
}

func (o TriggerMapOutput) MapIndex(k pulumi.StringInput) TriggerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Trigger {
		return vs[0].(map[string]*Trigger)[vs[1].(string)]
	}).(TriggerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerInput)(nil)).Elem(), &Trigger{})
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerArrayInput)(nil)).Elem(), TriggerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerMapInput)(nil)).Elem(), TriggerMap{})
	pulumi.RegisterOutputType(TriggerOutput{})
	pulumi.RegisterOutputType(TriggerArrayOutput{})
	pulumi.RegisterOutputType(TriggerMapOutput{})
}
