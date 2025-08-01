// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mnq

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Scaleway Messaging and Queuing SQS queues.
// For further information, see
// our [main documentation](https://www.scaleway.com/en/docs/messaging/how-to/create-manage-queues/).
//
// ## Example Usage
//
// ### Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/mnq"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := mnq.NewSqs(ctx, "main", nil)
//			if err != nil {
//				return err
//			}
//			mainSqsCredentials, err := mnq.NewSqsCredentials(ctx, "main", &mnq.SqsCredentialsArgs{
//				ProjectId: main.ProjectId,
//				Name:      pulumi.String("sqs-credentials"),
//				Permissions: &mnq.SqsCredentialsPermissionsArgs{
//					CanManage:  pulumi.Bool(true),
//					CanReceive: pulumi.Bool(false),
//					CanPublish: pulumi.Bool(false),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = mnq.NewSqsQueue(ctx, "main", &mnq.SqsQueueArgs{
//				ProjectId:   main.ProjectId,
//				Name:        pulumi.String("my-queue"),
//				SqsEndpoint: main.Endpoint,
//				AccessKey:   mainSqsCredentials.AccessKey,
//				SecretKey:   mainSqsCredentials.SecretKey,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type SqsQueue struct {
	pulumi.CustomResourceState

	// The access key of the SQS queue.
	AccessKey pulumi.StringOutput `pulumi:"accessKey"`
	// Specifies whether to enable content-based deduplication. Defaults to `false`.
	ContentBasedDeduplication pulumi.BoolOutput `pulumi:"contentBasedDeduplication"`
	// Whether the queue is a FIFO queue. If true, the queue name must end with .fifo. Defaults to `false`.
	FifoQueue pulumi.BoolOutput `pulumi:"fifoQueue"`
	// The number of seconds the queue retains a message. Must be between 60 and 1_209_600. Defaults to 345_600.
	MessageMaxAge pulumi.IntPtrOutput `pulumi:"messageMaxAge"`
	// The maximum size of a message. Should be in bytes. Must be between 1024 and 262_144. Defaults to 262_144.
	MessageMaxSize pulumi.IntPtrOutput `pulumi:"messageMaxSize"`
	// The unique name of the SQS queue. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// `projectId`) The ID of the Project in which SQS is enabled.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The number of seconds to wait for a message to arrive in the queue before returning. Must be between 0 and 20. Defaults to 0.
	ReceiveWaitTimeSeconds pulumi.IntPtrOutput `pulumi:"receiveWaitTimeSeconds"`
	// `region`). The region in which SQS is enabled.
	Region pulumi.StringOutput `pulumi:"region"`
	// The secret key of the SQS queue.
	SecretKey pulumi.StringOutput `pulumi:"secretKey"`
	// The endpoint of the SQS queue. Can contain a {region} placeholder. Defaults to `https://sqs.mnq.{region}.scaleway.com`.
	SqsEndpoint pulumi.StringPtrOutput `pulumi:"sqsEndpoint"`
	// The URL of the queue.
	Url pulumi.StringOutput `pulumi:"url"`
	// The number of seconds a message is hidden from other consumers. Must be between 0 and 43_200. Defaults to 30.
	VisibilityTimeoutSeconds pulumi.IntPtrOutput `pulumi:"visibilityTimeoutSeconds"`
}

// NewSqsQueue registers a new resource with the given unique name, arguments, and options.
func NewSqsQueue(ctx *pulumi.Context,
	name string, args *SqsQueueArgs, opts ...pulumi.ResourceOption) (*SqsQueue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessKey == nil {
		return nil, errors.New("invalid value for required argument 'AccessKey'")
	}
	if args.SecretKey == nil {
		return nil, errors.New("invalid value for required argument 'SecretKey'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("scaleway:index/mnqSqsQueue:MnqSqsQueue"),
		},
	})
	opts = append(opts, aliases)
	if args.AccessKey != nil {
		args.AccessKey = pulumi.ToSecret(args.AccessKey).(pulumi.StringInput)
	}
	if args.SecretKey != nil {
		args.SecretKey = pulumi.ToSecret(args.SecretKey).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"accessKey",
		"secretKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SqsQueue
	err := ctx.RegisterResource("scaleway:mnq/sqsQueue:SqsQueue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqsQueue gets an existing SqsQueue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqsQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqsQueueState, opts ...pulumi.ResourceOption) (*SqsQueue, error) {
	var resource SqsQueue
	err := ctx.ReadResource("scaleway:mnq/sqsQueue:SqsQueue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqsQueue resources.
type sqsQueueState struct {
	// The access key of the SQS queue.
	AccessKey *string `pulumi:"accessKey"`
	// Specifies whether to enable content-based deduplication. Defaults to `false`.
	ContentBasedDeduplication *bool `pulumi:"contentBasedDeduplication"`
	// Whether the queue is a FIFO queue. If true, the queue name must end with .fifo. Defaults to `false`.
	FifoQueue *bool `pulumi:"fifoQueue"`
	// The number of seconds the queue retains a message. Must be between 60 and 1_209_600. Defaults to 345_600.
	MessageMaxAge *int `pulumi:"messageMaxAge"`
	// The maximum size of a message. Should be in bytes. Must be between 1024 and 262_144. Defaults to 262_144.
	MessageMaxSize *int `pulumi:"messageMaxSize"`
	// The unique name of the SQS queue. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// `projectId`) The ID of the Project in which SQS is enabled.
	ProjectId *string `pulumi:"projectId"`
	// The number of seconds to wait for a message to arrive in the queue before returning. Must be between 0 and 20. Defaults to 0.
	ReceiveWaitTimeSeconds *int `pulumi:"receiveWaitTimeSeconds"`
	// `region`). The region in which SQS is enabled.
	Region *string `pulumi:"region"`
	// The secret key of the SQS queue.
	SecretKey *string `pulumi:"secretKey"`
	// The endpoint of the SQS queue. Can contain a {region} placeholder. Defaults to `https://sqs.mnq.{region}.scaleway.com`.
	SqsEndpoint *string `pulumi:"sqsEndpoint"`
	// The URL of the queue.
	Url *string `pulumi:"url"`
	// The number of seconds a message is hidden from other consumers. Must be between 0 and 43_200. Defaults to 30.
	VisibilityTimeoutSeconds *int `pulumi:"visibilityTimeoutSeconds"`
}

type SqsQueueState struct {
	// The access key of the SQS queue.
	AccessKey pulumi.StringPtrInput
	// Specifies whether to enable content-based deduplication. Defaults to `false`.
	ContentBasedDeduplication pulumi.BoolPtrInput
	// Whether the queue is a FIFO queue. If true, the queue name must end with .fifo. Defaults to `false`.
	FifoQueue pulumi.BoolPtrInput
	// The number of seconds the queue retains a message. Must be between 60 and 1_209_600. Defaults to 345_600.
	MessageMaxAge pulumi.IntPtrInput
	// The maximum size of a message. Should be in bytes. Must be between 1024 and 262_144. Defaults to 262_144.
	MessageMaxSize pulumi.IntPtrInput
	// The unique name of the SQS queue. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// `projectId`) The ID of the Project in which SQS is enabled.
	ProjectId pulumi.StringPtrInput
	// The number of seconds to wait for a message to arrive in the queue before returning. Must be between 0 and 20. Defaults to 0.
	ReceiveWaitTimeSeconds pulumi.IntPtrInput
	// `region`). The region in which SQS is enabled.
	Region pulumi.StringPtrInput
	// The secret key of the SQS queue.
	SecretKey pulumi.StringPtrInput
	// The endpoint of the SQS queue. Can contain a {region} placeholder. Defaults to `https://sqs.mnq.{region}.scaleway.com`.
	SqsEndpoint pulumi.StringPtrInput
	// The URL of the queue.
	Url pulumi.StringPtrInput
	// The number of seconds a message is hidden from other consumers. Must be between 0 and 43_200. Defaults to 30.
	VisibilityTimeoutSeconds pulumi.IntPtrInput
}

func (SqsQueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqsQueueState)(nil)).Elem()
}

type sqsQueueArgs struct {
	// The access key of the SQS queue.
	AccessKey string `pulumi:"accessKey"`
	// Specifies whether to enable content-based deduplication. Defaults to `false`.
	ContentBasedDeduplication *bool `pulumi:"contentBasedDeduplication"`
	// Whether the queue is a FIFO queue. If true, the queue name must end with .fifo. Defaults to `false`.
	FifoQueue *bool `pulumi:"fifoQueue"`
	// The number of seconds the queue retains a message. Must be between 60 and 1_209_600. Defaults to 345_600.
	MessageMaxAge *int `pulumi:"messageMaxAge"`
	// The maximum size of a message. Should be in bytes. Must be between 1024 and 262_144. Defaults to 262_144.
	MessageMaxSize *int `pulumi:"messageMaxSize"`
	// The unique name of the SQS queue. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// `projectId`) The ID of the Project in which SQS is enabled.
	ProjectId *string `pulumi:"projectId"`
	// The number of seconds to wait for a message to arrive in the queue before returning. Must be between 0 and 20. Defaults to 0.
	ReceiveWaitTimeSeconds *int `pulumi:"receiveWaitTimeSeconds"`
	// `region`). The region in which SQS is enabled.
	Region *string `pulumi:"region"`
	// The secret key of the SQS queue.
	SecretKey string `pulumi:"secretKey"`
	// The endpoint of the SQS queue. Can contain a {region} placeholder. Defaults to `https://sqs.mnq.{region}.scaleway.com`.
	SqsEndpoint *string `pulumi:"sqsEndpoint"`
	// The number of seconds a message is hidden from other consumers. Must be between 0 and 43_200. Defaults to 30.
	VisibilityTimeoutSeconds *int `pulumi:"visibilityTimeoutSeconds"`
}

// The set of arguments for constructing a SqsQueue resource.
type SqsQueueArgs struct {
	// The access key of the SQS queue.
	AccessKey pulumi.StringInput
	// Specifies whether to enable content-based deduplication. Defaults to `false`.
	ContentBasedDeduplication pulumi.BoolPtrInput
	// Whether the queue is a FIFO queue. If true, the queue name must end with .fifo. Defaults to `false`.
	FifoQueue pulumi.BoolPtrInput
	// The number of seconds the queue retains a message. Must be between 60 and 1_209_600. Defaults to 345_600.
	MessageMaxAge pulumi.IntPtrInput
	// The maximum size of a message. Should be in bytes. Must be between 1024 and 262_144. Defaults to 262_144.
	MessageMaxSize pulumi.IntPtrInput
	// The unique name of the SQS queue. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// `projectId`) The ID of the Project in which SQS is enabled.
	ProjectId pulumi.StringPtrInput
	// The number of seconds to wait for a message to arrive in the queue before returning. Must be between 0 and 20. Defaults to 0.
	ReceiveWaitTimeSeconds pulumi.IntPtrInput
	// `region`). The region in which SQS is enabled.
	Region pulumi.StringPtrInput
	// The secret key of the SQS queue.
	SecretKey pulumi.StringInput
	// The endpoint of the SQS queue. Can contain a {region} placeholder. Defaults to `https://sqs.mnq.{region}.scaleway.com`.
	SqsEndpoint pulumi.StringPtrInput
	// The number of seconds a message is hidden from other consumers. Must be between 0 and 43_200. Defaults to 30.
	VisibilityTimeoutSeconds pulumi.IntPtrInput
}

func (SqsQueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqsQueueArgs)(nil)).Elem()
}

type SqsQueueInput interface {
	pulumi.Input

	ToSqsQueueOutput() SqsQueueOutput
	ToSqsQueueOutputWithContext(ctx context.Context) SqsQueueOutput
}

func (*SqsQueue) ElementType() reflect.Type {
	return reflect.TypeOf((**SqsQueue)(nil)).Elem()
}

func (i *SqsQueue) ToSqsQueueOutput() SqsQueueOutput {
	return i.ToSqsQueueOutputWithContext(context.Background())
}

func (i *SqsQueue) ToSqsQueueOutputWithContext(ctx context.Context) SqsQueueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqsQueueOutput)
}

// SqsQueueArrayInput is an input type that accepts SqsQueueArray and SqsQueueArrayOutput values.
// You can construct a concrete instance of `SqsQueueArrayInput` via:
//
//	SqsQueueArray{ SqsQueueArgs{...} }
type SqsQueueArrayInput interface {
	pulumi.Input

	ToSqsQueueArrayOutput() SqsQueueArrayOutput
	ToSqsQueueArrayOutputWithContext(context.Context) SqsQueueArrayOutput
}

type SqsQueueArray []SqsQueueInput

func (SqsQueueArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SqsQueue)(nil)).Elem()
}

func (i SqsQueueArray) ToSqsQueueArrayOutput() SqsQueueArrayOutput {
	return i.ToSqsQueueArrayOutputWithContext(context.Background())
}

func (i SqsQueueArray) ToSqsQueueArrayOutputWithContext(ctx context.Context) SqsQueueArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqsQueueArrayOutput)
}

// SqsQueueMapInput is an input type that accepts SqsQueueMap and SqsQueueMapOutput values.
// You can construct a concrete instance of `SqsQueueMapInput` via:
//
//	SqsQueueMap{ "key": SqsQueueArgs{...} }
type SqsQueueMapInput interface {
	pulumi.Input

	ToSqsQueueMapOutput() SqsQueueMapOutput
	ToSqsQueueMapOutputWithContext(context.Context) SqsQueueMapOutput
}

type SqsQueueMap map[string]SqsQueueInput

func (SqsQueueMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SqsQueue)(nil)).Elem()
}

func (i SqsQueueMap) ToSqsQueueMapOutput() SqsQueueMapOutput {
	return i.ToSqsQueueMapOutputWithContext(context.Background())
}

func (i SqsQueueMap) ToSqsQueueMapOutputWithContext(ctx context.Context) SqsQueueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqsQueueMapOutput)
}

type SqsQueueOutput struct{ *pulumi.OutputState }

func (SqsQueueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqsQueue)(nil)).Elem()
}

func (o SqsQueueOutput) ToSqsQueueOutput() SqsQueueOutput {
	return o
}

func (o SqsQueueOutput) ToSqsQueueOutputWithContext(ctx context.Context) SqsQueueOutput {
	return o
}

// The access key of the SQS queue.
func (o SqsQueueOutput) AccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v *SqsQueue) pulumi.StringOutput { return v.AccessKey }).(pulumi.StringOutput)
}

// Specifies whether to enable content-based deduplication. Defaults to `false`.
func (o SqsQueueOutput) ContentBasedDeduplication() pulumi.BoolOutput {
	return o.ApplyT(func(v *SqsQueue) pulumi.BoolOutput { return v.ContentBasedDeduplication }).(pulumi.BoolOutput)
}

// Whether the queue is a FIFO queue. If true, the queue name must end with .fifo. Defaults to `false`.
func (o SqsQueueOutput) FifoQueue() pulumi.BoolOutput {
	return o.ApplyT(func(v *SqsQueue) pulumi.BoolOutput { return v.FifoQueue }).(pulumi.BoolOutput)
}

// The number of seconds the queue retains a message. Must be between 60 and 1_209_600. Defaults to 345_600.
func (o SqsQueueOutput) MessageMaxAge() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqsQueue) pulumi.IntPtrOutput { return v.MessageMaxAge }).(pulumi.IntPtrOutput)
}

// The maximum size of a message. Should be in bytes. Must be between 1024 and 262_144. Defaults to 262_144.
func (o SqsQueueOutput) MessageMaxSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqsQueue) pulumi.IntPtrOutput { return v.MessageMaxSize }).(pulumi.IntPtrOutput)
}

// The unique name of the SQS queue. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
func (o SqsQueueOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqsQueue) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
func (o SqsQueueOutput) NamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *SqsQueue) pulumi.StringOutput { return v.NamePrefix }).(pulumi.StringOutput)
}

// `projectId`) The ID of the Project in which SQS is enabled.
func (o SqsQueueOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *SqsQueue) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The number of seconds to wait for a message to arrive in the queue before returning. Must be between 0 and 20. Defaults to 0.
func (o SqsQueueOutput) ReceiveWaitTimeSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqsQueue) pulumi.IntPtrOutput { return v.ReceiveWaitTimeSeconds }).(pulumi.IntPtrOutput)
}

// `region`). The region in which SQS is enabled.
func (o SqsQueueOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *SqsQueue) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The secret key of the SQS queue.
func (o SqsQueueOutput) SecretKey() pulumi.StringOutput {
	return o.ApplyT(func(v *SqsQueue) pulumi.StringOutput { return v.SecretKey }).(pulumi.StringOutput)
}

// The endpoint of the SQS queue. Can contain a {region} placeholder. Defaults to `https://sqs.mnq.{region}.scaleway.com`.
func (o SqsQueueOutput) SqsEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqsQueue) pulumi.StringPtrOutput { return v.SqsEndpoint }).(pulumi.StringPtrOutput)
}

// The URL of the queue.
func (o SqsQueueOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *SqsQueue) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// The number of seconds a message is hidden from other consumers. Must be between 0 and 43_200. Defaults to 30.
func (o SqsQueueOutput) VisibilityTimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqsQueue) pulumi.IntPtrOutput { return v.VisibilityTimeoutSeconds }).(pulumi.IntPtrOutput)
}

type SqsQueueArrayOutput struct{ *pulumi.OutputState }

func (SqsQueueArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SqsQueue)(nil)).Elem()
}

func (o SqsQueueArrayOutput) ToSqsQueueArrayOutput() SqsQueueArrayOutput {
	return o
}

func (o SqsQueueArrayOutput) ToSqsQueueArrayOutputWithContext(ctx context.Context) SqsQueueArrayOutput {
	return o
}

func (o SqsQueueArrayOutput) Index(i pulumi.IntInput) SqsQueueOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SqsQueue {
		return vs[0].([]*SqsQueue)[vs[1].(int)]
	}).(SqsQueueOutput)
}

type SqsQueueMapOutput struct{ *pulumi.OutputState }

func (SqsQueueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SqsQueue)(nil)).Elem()
}

func (o SqsQueueMapOutput) ToSqsQueueMapOutput() SqsQueueMapOutput {
	return o
}

func (o SqsQueueMapOutput) ToSqsQueueMapOutputWithContext(ctx context.Context) SqsQueueMapOutput {
	return o
}

func (o SqsQueueMapOutput) MapIndex(k pulumi.StringInput) SqsQueueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SqsQueue {
		return vs[0].(map[string]*SqsQueue)[vs[1].(string)]
	}).(SqsQueueOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SqsQueueInput)(nil)).Elem(), &SqsQueue{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqsQueueArrayInput)(nil)).Elem(), SqsQueueArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqsQueueMapInput)(nil)).Elem(), SqsQueueMap{})
	pulumi.RegisterOutputType(SqsQueueOutput{})
	pulumi.RegisterOutputType(SqsQueueArrayOutput{})
	pulumi.RegisterOutputType(SqsQueueMapOutput{})
}
