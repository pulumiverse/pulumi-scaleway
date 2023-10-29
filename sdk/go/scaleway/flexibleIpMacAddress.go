// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/lbrlabs/pulumi-scaleway/sdk/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Creates and manages Scaleway Flexible IP Mac Addresses.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/flexible-ip/api).
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
//			mainFlexibleIp, err := scaleway.NewFlexibleIp(ctx, "mainFlexibleIp", nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewFlexibleIpMacAddress(ctx, "mainFlexibleIpMacAddress", &scaleway.FlexibleIpMacAddressArgs{
//				FlexibleIpId: mainFlexibleIp.ID(),
//				Type:         pulumi.String("kvm"),
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
// ### Duplicate on many other flexible IPs
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
//			myOffer, err := scaleway.GetBaremetalOffer(ctx, &scaleway.GetBaremetalOfferArgs{
//				Name: pulumi.StringRef("EM-B112X-SSD"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			base, err := scaleway.NewBaremetalServer(ctx, "base", &scaleway.BaremetalServerArgs{
//				Offer:                  *pulumi.String(myOffer.OfferId),
//				InstallConfigAfterward: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			ip01, err := scaleway.NewFlexibleIp(ctx, "ip01", &scaleway.FlexibleIpArgs{
//				ServerId: base.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			ip02, err := scaleway.NewFlexibleIp(ctx, "ip02", &scaleway.FlexibleIpArgs{
//				ServerId: base.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			ip03, err := scaleway.NewFlexibleIp(ctx, "ip03", &scaleway.FlexibleIpArgs{
//				ServerId: base.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewFlexibleIpMacAddress(ctx, "main", &scaleway.FlexibleIpMacAddressArgs{
//				FlexibleIpId: ip01.ID(),
//				Type:         pulumi.String("kvm"),
//				FlexibleIpIdsToDuplicates: pulumi.StringArray{
//					ip02.ID(),
//					ip03.ID(),
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
// Flexible IP Mac Addresses can be imported using the `{zone}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/flexibleIpMacAddress:FlexibleIpMacAddress main fr-par-1/11111111-1111-1111-1111-111111111111
//
// ```
type FlexibleIpMacAddress struct {
	pulumi.CustomResourceState

	// The Virtual MAC address.
	Address pulumi.StringOutput `pulumi:"address"`
	// The date at which the Virtual Mac Address was created (RFC 3339 format).
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The ID of the flexible IP for which to generate a virtual MAC.
	FlexibleIpId pulumi.StringOutput `pulumi:"flexibleIpId"`
	// The IDs of the flexible IPs on which to duplicate the virtual MAC.
	// > **Important:** The flexible IPs need to be attached to the same server for the operation to work.
	FlexibleIpIdsToDuplicates pulumi.StringArrayOutput `pulumi:"flexibleIpIdsToDuplicates"`
	// The Virtual MAC status.
	Status pulumi.StringOutput `pulumi:"status"`
	// The type of the virtual MAC.
	Type pulumi.StringOutput `pulumi:"type"`
	// The date at which the Virtual Mac Address was last updated (RFC 3339 format).
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// The zone of the Virtual Mac Address.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewFlexibleIpMacAddress registers a new resource with the given unique name, arguments, and options.
func NewFlexibleIpMacAddress(ctx *pulumi.Context,
	name string, args *FlexibleIpMacAddressArgs, opts ...pulumi.ResourceOption) (*FlexibleIpMacAddress, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FlexibleIpId == nil {
		return nil, errors.New("invalid value for required argument 'FlexibleIpId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FlexibleIpMacAddress
	err := ctx.RegisterResource("scaleway:index/flexibleIpMacAddress:FlexibleIpMacAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlexibleIpMacAddress gets an existing FlexibleIpMacAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlexibleIpMacAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlexibleIpMacAddressState, opts ...pulumi.ResourceOption) (*FlexibleIpMacAddress, error) {
	var resource FlexibleIpMacAddress
	err := ctx.ReadResource("scaleway:index/flexibleIpMacAddress:FlexibleIpMacAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FlexibleIpMacAddress resources.
type flexibleIpMacAddressState struct {
	// The Virtual MAC address.
	Address *string `pulumi:"address"`
	// The date at which the Virtual Mac Address was created (RFC 3339 format).
	CreatedAt *string `pulumi:"createdAt"`
	// The ID of the flexible IP for which to generate a virtual MAC.
	FlexibleIpId *string `pulumi:"flexibleIpId"`
	// The IDs of the flexible IPs on which to duplicate the virtual MAC.
	// > **Important:** The flexible IPs need to be attached to the same server for the operation to work.
	FlexibleIpIdsToDuplicates []string `pulumi:"flexibleIpIdsToDuplicates"`
	// The Virtual MAC status.
	Status *string `pulumi:"status"`
	// The type of the virtual MAC.
	Type *string `pulumi:"type"`
	// The date at which the Virtual Mac Address was last updated (RFC 3339 format).
	UpdatedAt *string `pulumi:"updatedAt"`
	// The zone of the Virtual Mac Address.
	Zone *string `pulumi:"zone"`
}

type FlexibleIpMacAddressState struct {
	// The Virtual MAC address.
	Address pulumi.StringPtrInput
	// The date at which the Virtual Mac Address was created (RFC 3339 format).
	CreatedAt pulumi.StringPtrInput
	// The ID of the flexible IP for which to generate a virtual MAC.
	FlexibleIpId pulumi.StringPtrInput
	// The IDs of the flexible IPs on which to duplicate the virtual MAC.
	// > **Important:** The flexible IPs need to be attached to the same server for the operation to work.
	FlexibleIpIdsToDuplicates pulumi.StringArrayInput
	// The Virtual MAC status.
	Status pulumi.StringPtrInput
	// The type of the virtual MAC.
	Type pulumi.StringPtrInput
	// The date at which the Virtual Mac Address was last updated (RFC 3339 format).
	UpdatedAt pulumi.StringPtrInput
	// The zone of the Virtual Mac Address.
	Zone pulumi.StringPtrInput
}

func (FlexibleIpMacAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*flexibleIpMacAddressState)(nil)).Elem()
}

type flexibleIpMacAddressArgs struct {
	// The ID of the flexible IP for which to generate a virtual MAC.
	FlexibleIpId string `pulumi:"flexibleIpId"`
	// The IDs of the flexible IPs on which to duplicate the virtual MAC.
	// > **Important:** The flexible IPs need to be attached to the same server for the operation to work.
	FlexibleIpIdsToDuplicates []string `pulumi:"flexibleIpIdsToDuplicates"`
	// The type of the virtual MAC.
	Type string `pulumi:"type"`
	// The zone of the Virtual Mac Address.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a FlexibleIpMacAddress resource.
type FlexibleIpMacAddressArgs struct {
	// The ID of the flexible IP for which to generate a virtual MAC.
	FlexibleIpId pulumi.StringInput
	// The IDs of the flexible IPs on which to duplicate the virtual MAC.
	// > **Important:** The flexible IPs need to be attached to the same server for the operation to work.
	FlexibleIpIdsToDuplicates pulumi.StringArrayInput
	// The type of the virtual MAC.
	Type pulumi.StringInput
	// The zone of the Virtual Mac Address.
	Zone pulumi.StringPtrInput
}

func (FlexibleIpMacAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flexibleIpMacAddressArgs)(nil)).Elem()
}

type FlexibleIpMacAddressInput interface {
	pulumi.Input

	ToFlexibleIpMacAddressOutput() FlexibleIpMacAddressOutput
	ToFlexibleIpMacAddressOutputWithContext(ctx context.Context) FlexibleIpMacAddressOutput
}

func (*FlexibleIpMacAddress) ElementType() reflect.Type {
	return reflect.TypeOf((**FlexibleIpMacAddress)(nil)).Elem()
}

func (i *FlexibleIpMacAddress) ToFlexibleIpMacAddressOutput() FlexibleIpMacAddressOutput {
	return i.ToFlexibleIpMacAddressOutputWithContext(context.Background())
}

func (i *FlexibleIpMacAddress) ToFlexibleIpMacAddressOutputWithContext(ctx context.Context) FlexibleIpMacAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexibleIpMacAddressOutput)
}

func (i *FlexibleIpMacAddress) ToOutput(ctx context.Context) pulumix.Output[*FlexibleIpMacAddress] {
	return pulumix.Output[*FlexibleIpMacAddress]{
		OutputState: i.ToFlexibleIpMacAddressOutputWithContext(ctx).OutputState,
	}
}

// FlexibleIpMacAddressArrayInput is an input type that accepts FlexibleIpMacAddressArray and FlexibleIpMacAddressArrayOutput values.
// You can construct a concrete instance of `FlexibleIpMacAddressArrayInput` via:
//
//	FlexibleIpMacAddressArray{ FlexibleIpMacAddressArgs{...} }
type FlexibleIpMacAddressArrayInput interface {
	pulumi.Input

	ToFlexibleIpMacAddressArrayOutput() FlexibleIpMacAddressArrayOutput
	ToFlexibleIpMacAddressArrayOutputWithContext(context.Context) FlexibleIpMacAddressArrayOutput
}

type FlexibleIpMacAddressArray []FlexibleIpMacAddressInput

func (FlexibleIpMacAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FlexibleIpMacAddress)(nil)).Elem()
}

func (i FlexibleIpMacAddressArray) ToFlexibleIpMacAddressArrayOutput() FlexibleIpMacAddressArrayOutput {
	return i.ToFlexibleIpMacAddressArrayOutputWithContext(context.Background())
}

func (i FlexibleIpMacAddressArray) ToFlexibleIpMacAddressArrayOutputWithContext(ctx context.Context) FlexibleIpMacAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexibleIpMacAddressArrayOutput)
}

func (i FlexibleIpMacAddressArray) ToOutput(ctx context.Context) pulumix.Output[[]*FlexibleIpMacAddress] {
	return pulumix.Output[[]*FlexibleIpMacAddress]{
		OutputState: i.ToFlexibleIpMacAddressArrayOutputWithContext(ctx).OutputState,
	}
}

// FlexibleIpMacAddressMapInput is an input type that accepts FlexibleIpMacAddressMap and FlexibleIpMacAddressMapOutput values.
// You can construct a concrete instance of `FlexibleIpMacAddressMapInput` via:
//
//	FlexibleIpMacAddressMap{ "key": FlexibleIpMacAddressArgs{...} }
type FlexibleIpMacAddressMapInput interface {
	pulumi.Input

	ToFlexibleIpMacAddressMapOutput() FlexibleIpMacAddressMapOutput
	ToFlexibleIpMacAddressMapOutputWithContext(context.Context) FlexibleIpMacAddressMapOutput
}

type FlexibleIpMacAddressMap map[string]FlexibleIpMacAddressInput

func (FlexibleIpMacAddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FlexibleIpMacAddress)(nil)).Elem()
}

func (i FlexibleIpMacAddressMap) ToFlexibleIpMacAddressMapOutput() FlexibleIpMacAddressMapOutput {
	return i.ToFlexibleIpMacAddressMapOutputWithContext(context.Background())
}

func (i FlexibleIpMacAddressMap) ToFlexibleIpMacAddressMapOutputWithContext(ctx context.Context) FlexibleIpMacAddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexibleIpMacAddressMapOutput)
}

func (i FlexibleIpMacAddressMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*FlexibleIpMacAddress] {
	return pulumix.Output[map[string]*FlexibleIpMacAddress]{
		OutputState: i.ToFlexibleIpMacAddressMapOutputWithContext(ctx).OutputState,
	}
}

type FlexibleIpMacAddressOutput struct{ *pulumi.OutputState }

func (FlexibleIpMacAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlexibleIpMacAddress)(nil)).Elem()
}

func (o FlexibleIpMacAddressOutput) ToFlexibleIpMacAddressOutput() FlexibleIpMacAddressOutput {
	return o
}

func (o FlexibleIpMacAddressOutput) ToFlexibleIpMacAddressOutputWithContext(ctx context.Context) FlexibleIpMacAddressOutput {
	return o
}

func (o FlexibleIpMacAddressOutput) ToOutput(ctx context.Context) pulumix.Output[*FlexibleIpMacAddress] {
	return pulumix.Output[*FlexibleIpMacAddress]{
		OutputState: o.OutputState,
	}
}

// The Virtual MAC address.
func (o FlexibleIpMacAddressOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *FlexibleIpMacAddress) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// The date at which the Virtual Mac Address was created (RFC 3339 format).
func (o FlexibleIpMacAddressOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *FlexibleIpMacAddress) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The ID of the flexible IP for which to generate a virtual MAC.
func (o FlexibleIpMacAddressOutput) FlexibleIpId() pulumi.StringOutput {
	return o.ApplyT(func(v *FlexibleIpMacAddress) pulumi.StringOutput { return v.FlexibleIpId }).(pulumi.StringOutput)
}

// The IDs of the flexible IPs on which to duplicate the virtual MAC.
// > **Important:** The flexible IPs need to be attached to the same server for the operation to work.
func (o FlexibleIpMacAddressOutput) FlexibleIpIdsToDuplicates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FlexibleIpMacAddress) pulumi.StringArrayOutput { return v.FlexibleIpIdsToDuplicates }).(pulumi.StringArrayOutput)
}

// The Virtual MAC status.
func (o FlexibleIpMacAddressOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *FlexibleIpMacAddress) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The type of the virtual MAC.
func (o FlexibleIpMacAddressOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FlexibleIpMacAddress) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The date at which the Virtual Mac Address was last updated (RFC 3339 format).
func (o FlexibleIpMacAddressOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *FlexibleIpMacAddress) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// The zone of the Virtual Mac Address.
func (o FlexibleIpMacAddressOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *FlexibleIpMacAddress) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type FlexibleIpMacAddressArrayOutput struct{ *pulumi.OutputState }

func (FlexibleIpMacAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FlexibleIpMacAddress)(nil)).Elem()
}

func (o FlexibleIpMacAddressArrayOutput) ToFlexibleIpMacAddressArrayOutput() FlexibleIpMacAddressArrayOutput {
	return o
}

func (o FlexibleIpMacAddressArrayOutput) ToFlexibleIpMacAddressArrayOutputWithContext(ctx context.Context) FlexibleIpMacAddressArrayOutput {
	return o
}

func (o FlexibleIpMacAddressArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*FlexibleIpMacAddress] {
	return pulumix.Output[[]*FlexibleIpMacAddress]{
		OutputState: o.OutputState,
	}
}

func (o FlexibleIpMacAddressArrayOutput) Index(i pulumi.IntInput) FlexibleIpMacAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FlexibleIpMacAddress {
		return vs[0].([]*FlexibleIpMacAddress)[vs[1].(int)]
	}).(FlexibleIpMacAddressOutput)
}

type FlexibleIpMacAddressMapOutput struct{ *pulumi.OutputState }

func (FlexibleIpMacAddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FlexibleIpMacAddress)(nil)).Elem()
}

func (o FlexibleIpMacAddressMapOutput) ToFlexibleIpMacAddressMapOutput() FlexibleIpMacAddressMapOutput {
	return o
}

func (o FlexibleIpMacAddressMapOutput) ToFlexibleIpMacAddressMapOutputWithContext(ctx context.Context) FlexibleIpMacAddressMapOutput {
	return o
}

func (o FlexibleIpMacAddressMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*FlexibleIpMacAddress] {
	return pulumix.Output[map[string]*FlexibleIpMacAddress]{
		OutputState: o.OutputState,
	}
}

func (o FlexibleIpMacAddressMapOutput) MapIndex(k pulumi.StringInput) FlexibleIpMacAddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FlexibleIpMacAddress {
		return vs[0].(map[string]*FlexibleIpMacAddress)[vs[1].(string)]
	}).(FlexibleIpMacAddressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FlexibleIpMacAddressInput)(nil)).Elem(), &FlexibleIpMacAddress{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlexibleIpMacAddressArrayInput)(nil)).Elem(), FlexibleIpMacAddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlexibleIpMacAddressMapInput)(nil)).Elem(), FlexibleIpMacAddressMap{})
	pulumi.RegisterOutputType(FlexibleIpMacAddressOutput{})
	pulumi.RegisterOutputType(FlexibleIpMacAddressArrayOutput{})
	pulumi.RegisterOutputType(FlexibleIpMacAddressMapOutput{})
}