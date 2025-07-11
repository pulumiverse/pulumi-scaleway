// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tem

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages blocklisted email addresses for a Scaleway Transactional Email Domain.
// For more information refer to the [API documentation](https://www.scaleway.com/en/developers/api/transactional-email/#post-transactional-email-v1alpha1-regions-region-blocklists).
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
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/tem"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := tem.NewBlockedList(ctx, "test", &tem.BlockedListArgs{
//				DomainId: pulumi.String("fr-par/12345678-1234-1234-1234-123456789abc"),
//				Email:    pulumi.String("spam@example.com"),
//				Type:     pulumi.String("mailbox_full"),
//				Reason:   pulumi.String("Spam detected"),
//				Region:   pulumi.String("fr-par"),
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
// Blocklists can be imported using the `{region}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:tem/blockedList:BlockedList test fr-par/11111111-1111-1111-1111-111111111111
// ```
type BlockedList struct {
	pulumi.CustomResourceState

	// The ID of the domain affected by the blocklist. Must be in the format `{region}/{domain_id}`.
	DomainId pulumi.StringOutput `pulumi:"domainId"`
	// The email address to block.
	Email pulumi.StringOutput `pulumi:"email"`
	// The ID of the project this blocklist belongs to. Defaults to the provider's project ID.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Reason for blocking the email address.
	Reason pulumi.StringPtrOutput `pulumi:"reason"`
	// The region in which the blocklist is created. Defaults to the provider's region.
	Region pulumi.StringOutput `pulumi:"region"`
	// Type of the blocklist. Possible values are:
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBlockedList registers a new resource with the given unique name, arguments, and options.
func NewBlockedList(ctx *pulumi.Context,
	name string, args *BlockedListArgs, opts ...pulumi.ResourceOption) (*BlockedList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainId == nil {
		return nil, errors.New("invalid value for required argument 'DomainId'")
	}
	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BlockedList
	err := ctx.RegisterResource("scaleway:tem/blockedList:BlockedList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBlockedList gets an existing BlockedList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBlockedList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlockedListState, opts ...pulumi.ResourceOption) (*BlockedList, error) {
	var resource BlockedList
	err := ctx.ReadResource("scaleway:tem/blockedList:BlockedList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BlockedList resources.
type blockedListState struct {
	// The ID of the domain affected by the blocklist. Must be in the format `{region}/{domain_id}`.
	DomainId *string `pulumi:"domainId"`
	// The email address to block.
	Email *string `pulumi:"email"`
	// The ID of the project this blocklist belongs to. Defaults to the provider's project ID.
	ProjectId *string `pulumi:"projectId"`
	// Reason for blocking the email address.
	Reason *string `pulumi:"reason"`
	// The region in which the blocklist is created. Defaults to the provider's region.
	Region *string `pulumi:"region"`
	// Type of the blocklist. Possible values are:
	Type *string `pulumi:"type"`
}

type BlockedListState struct {
	// The ID of the domain affected by the blocklist. Must be in the format `{region}/{domain_id}`.
	DomainId pulumi.StringPtrInput
	// The email address to block.
	Email pulumi.StringPtrInput
	// The ID of the project this blocklist belongs to. Defaults to the provider's project ID.
	ProjectId pulumi.StringPtrInput
	// Reason for blocking the email address.
	Reason pulumi.StringPtrInput
	// The region in which the blocklist is created. Defaults to the provider's region.
	Region pulumi.StringPtrInput
	// Type of the blocklist. Possible values are:
	Type pulumi.StringPtrInput
}

func (BlockedListState) ElementType() reflect.Type {
	return reflect.TypeOf((*blockedListState)(nil)).Elem()
}

type blockedListArgs struct {
	// The ID of the domain affected by the blocklist. Must be in the format `{region}/{domain_id}`.
	DomainId string `pulumi:"domainId"`
	// The email address to block.
	Email string `pulumi:"email"`
	// The ID of the project this blocklist belongs to. Defaults to the provider's project ID.
	ProjectId *string `pulumi:"projectId"`
	// Reason for blocking the email address.
	Reason *string `pulumi:"reason"`
	// The region in which the blocklist is created. Defaults to the provider's region.
	Region *string `pulumi:"region"`
	// Type of the blocklist. Possible values are:
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a BlockedList resource.
type BlockedListArgs struct {
	// The ID of the domain affected by the blocklist. Must be in the format `{region}/{domain_id}`.
	DomainId pulumi.StringInput
	// The email address to block.
	Email pulumi.StringInput
	// The ID of the project this blocklist belongs to. Defaults to the provider's project ID.
	ProjectId pulumi.StringPtrInput
	// Reason for blocking the email address.
	Reason pulumi.StringPtrInput
	// The region in which the blocklist is created. Defaults to the provider's region.
	Region pulumi.StringPtrInput
	// Type of the blocklist. Possible values are:
	Type pulumi.StringInput
}

func (BlockedListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blockedListArgs)(nil)).Elem()
}

type BlockedListInput interface {
	pulumi.Input

	ToBlockedListOutput() BlockedListOutput
	ToBlockedListOutputWithContext(ctx context.Context) BlockedListOutput
}

func (*BlockedList) ElementType() reflect.Type {
	return reflect.TypeOf((**BlockedList)(nil)).Elem()
}

func (i *BlockedList) ToBlockedListOutput() BlockedListOutput {
	return i.ToBlockedListOutputWithContext(context.Background())
}

func (i *BlockedList) ToBlockedListOutputWithContext(ctx context.Context) BlockedListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlockedListOutput)
}

// BlockedListArrayInput is an input type that accepts BlockedListArray and BlockedListArrayOutput values.
// You can construct a concrete instance of `BlockedListArrayInput` via:
//
//	BlockedListArray{ BlockedListArgs{...} }
type BlockedListArrayInput interface {
	pulumi.Input

	ToBlockedListArrayOutput() BlockedListArrayOutput
	ToBlockedListArrayOutputWithContext(context.Context) BlockedListArrayOutput
}

type BlockedListArray []BlockedListInput

func (BlockedListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BlockedList)(nil)).Elem()
}

func (i BlockedListArray) ToBlockedListArrayOutput() BlockedListArrayOutput {
	return i.ToBlockedListArrayOutputWithContext(context.Background())
}

func (i BlockedListArray) ToBlockedListArrayOutputWithContext(ctx context.Context) BlockedListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlockedListArrayOutput)
}

// BlockedListMapInput is an input type that accepts BlockedListMap and BlockedListMapOutput values.
// You can construct a concrete instance of `BlockedListMapInput` via:
//
//	BlockedListMap{ "key": BlockedListArgs{...} }
type BlockedListMapInput interface {
	pulumi.Input

	ToBlockedListMapOutput() BlockedListMapOutput
	ToBlockedListMapOutputWithContext(context.Context) BlockedListMapOutput
}

type BlockedListMap map[string]BlockedListInput

func (BlockedListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BlockedList)(nil)).Elem()
}

func (i BlockedListMap) ToBlockedListMapOutput() BlockedListMapOutput {
	return i.ToBlockedListMapOutputWithContext(context.Background())
}

func (i BlockedListMap) ToBlockedListMapOutputWithContext(ctx context.Context) BlockedListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlockedListMapOutput)
}

type BlockedListOutput struct{ *pulumi.OutputState }

func (BlockedListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlockedList)(nil)).Elem()
}

func (o BlockedListOutput) ToBlockedListOutput() BlockedListOutput {
	return o
}

func (o BlockedListOutput) ToBlockedListOutputWithContext(ctx context.Context) BlockedListOutput {
	return o
}

// The ID of the domain affected by the blocklist. Must be in the format `{region}/{domain_id}`.
func (o BlockedListOutput) DomainId() pulumi.StringOutput {
	return o.ApplyT(func(v *BlockedList) pulumi.StringOutput { return v.DomainId }).(pulumi.StringOutput)
}

// The email address to block.
func (o BlockedListOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *BlockedList) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// The ID of the project this blocklist belongs to. Defaults to the provider's project ID.
func (o BlockedListOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *BlockedList) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Reason for blocking the email address.
func (o BlockedListOutput) Reason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlockedList) pulumi.StringPtrOutput { return v.Reason }).(pulumi.StringPtrOutput)
}

// The region in which the blocklist is created. Defaults to the provider's region.
func (o BlockedListOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *BlockedList) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Type of the blocklist. Possible values are:
func (o BlockedListOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BlockedList) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type BlockedListArrayOutput struct{ *pulumi.OutputState }

func (BlockedListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BlockedList)(nil)).Elem()
}

func (o BlockedListArrayOutput) ToBlockedListArrayOutput() BlockedListArrayOutput {
	return o
}

func (o BlockedListArrayOutput) ToBlockedListArrayOutputWithContext(ctx context.Context) BlockedListArrayOutput {
	return o
}

func (o BlockedListArrayOutput) Index(i pulumi.IntInput) BlockedListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BlockedList {
		return vs[0].([]*BlockedList)[vs[1].(int)]
	}).(BlockedListOutput)
}

type BlockedListMapOutput struct{ *pulumi.OutputState }

func (BlockedListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BlockedList)(nil)).Elem()
}

func (o BlockedListMapOutput) ToBlockedListMapOutput() BlockedListMapOutput {
	return o
}

func (o BlockedListMapOutput) ToBlockedListMapOutputWithContext(ctx context.Context) BlockedListMapOutput {
	return o
}

func (o BlockedListMapOutput) MapIndex(k pulumi.StringInput) BlockedListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BlockedList {
		return vs[0].(map[string]*BlockedList)[vs[1].(string)]
	}).(BlockedListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BlockedListInput)(nil)).Elem(), &BlockedList{})
	pulumi.RegisterInputType(reflect.TypeOf((*BlockedListArrayInput)(nil)).Elem(), BlockedListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BlockedListMapInput)(nil)).Elem(), BlockedListMap{})
	pulumi.RegisterOutputType(BlockedListOutput{})
	pulumi.RegisterOutputType(BlockedListArrayOutput{})
	pulumi.RegisterOutputType(BlockedListMapOutput{})
}
