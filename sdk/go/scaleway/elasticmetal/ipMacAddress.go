// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticmetal

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Scaleway Flexible IP Mac Addresses.
// For more information, see the [API documentation](https://www.scaleway.com/en/developers/api/elastic-metal-flexible-ip/).
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
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/elasticmetal"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := elasticmetal.NewIp(ctx, "main", nil)
//			if err != nil {
//				return err
//			}
//			_, err = elasticmetal.NewIpMacAddress(ctx, "main", &elasticmetal.IpMacAddressArgs{
//				FlexibleIpId: main.ID(),
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
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/elasticmetal"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			myOffer, err := elasticmetal.GetOffer(ctx, &elasticmetal.GetOfferArgs{
//				Name: pulumi.StringRef("EM-B112X-SSD"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			base, err := elasticmetal.NewServer(ctx, "base", &elasticmetal.ServerArgs{
//				Name:                   pulumi.String("TestAccScalewayBaremetalServer_WithoutInstallConfig"),
//				Offer:                  pulumi.String(myOffer.OfferId),
//				InstallConfigAfterward: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			ip01, err := elasticmetal.NewIp(ctx, "ip01", &elasticmetal.IpArgs{
//				ServerId: base.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			ip02, err := elasticmetal.NewIp(ctx, "ip02", &elasticmetal.IpArgs{
//				ServerId: base.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			ip03, err := elasticmetal.NewIp(ctx, "ip03", &elasticmetal.IpArgs{
//				ServerId: base.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = elasticmetal.NewIpMacAddress(ctx, "main", &elasticmetal.IpMacAddressArgs{
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
// Flexible IP Mac Addresses can be imported using the `{zone}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:elasticmetal/ipMacAddress:IpMacAddress main fr-par-1/11111111-1111-1111-1111-111111111111
// ```
type IpMacAddress struct {
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

// NewIpMacAddress registers a new resource with the given unique name, arguments, and options.
func NewIpMacAddress(ctx *pulumi.Context,
	name string, args *IpMacAddressArgs, opts ...pulumi.ResourceOption) (*IpMacAddress, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FlexibleIpId == nil {
		return nil, errors.New("invalid value for required argument 'FlexibleIpId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("scaleway:index/flexibleIpMacAddress:FlexibleIpMacAddress"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IpMacAddress
	err := ctx.RegisterResource("scaleway:elasticmetal/ipMacAddress:IpMacAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpMacAddress gets an existing IpMacAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpMacAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpMacAddressState, opts ...pulumi.ResourceOption) (*IpMacAddress, error) {
	var resource IpMacAddress
	err := ctx.ReadResource("scaleway:elasticmetal/ipMacAddress:IpMacAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpMacAddress resources.
type ipMacAddressState struct {
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

type IpMacAddressState struct {
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

func (IpMacAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipMacAddressState)(nil)).Elem()
}

type ipMacAddressArgs struct {
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

// The set of arguments for constructing a IpMacAddress resource.
type IpMacAddressArgs struct {
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

func (IpMacAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipMacAddressArgs)(nil)).Elem()
}

type IpMacAddressInput interface {
	pulumi.Input

	ToIpMacAddressOutput() IpMacAddressOutput
	ToIpMacAddressOutputWithContext(ctx context.Context) IpMacAddressOutput
}

func (*IpMacAddress) ElementType() reflect.Type {
	return reflect.TypeOf((**IpMacAddress)(nil)).Elem()
}

func (i *IpMacAddress) ToIpMacAddressOutput() IpMacAddressOutput {
	return i.ToIpMacAddressOutputWithContext(context.Background())
}

func (i *IpMacAddress) ToIpMacAddressOutputWithContext(ctx context.Context) IpMacAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpMacAddressOutput)
}

// IpMacAddressArrayInput is an input type that accepts IpMacAddressArray and IpMacAddressArrayOutput values.
// You can construct a concrete instance of `IpMacAddressArrayInput` via:
//
//	IpMacAddressArray{ IpMacAddressArgs{...} }
type IpMacAddressArrayInput interface {
	pulumi.Input

	ToIpMacAddressArrayOutput() IpMacAddressArrayOutput
	ToIpMacAddressArrayOutputWithContext(context.Context) IpMacAddressArrayOutput
}

type IpMacAddressArray []IpMacAddressInput

func (IpMacAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpMacAddress)(nil)).Elem()
}

func (i IpMacAddressArray) ToIpMacAddressArrayOutput() IpMacAddressArrayOutput {
	return i.ToIpMacAddressArrayOutputWithContext(context.Background())
}

func (i IpMacAddressArray) ToIpMacAddressArrayOutputWithContext(ctx context.Context) IpMacAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpMacAddressArrayOutput)
}

// IpMacAddressMapInput is an input type that accepts IpMacAddressMap and IpMacAddressMapOutput values.
// You can construct a concrete instance of `IpMacAddressMapInput` via:
//
//	IpMacAddressMap{ "key": IpMacAddressArgs{...} }
type IpMacAddressMapInput interface {
	pulumi.Input

	ToIpMacAddressMapOutput() IpMacAddressMapOutput
	ToIpMacAddressMapOutputWithContext(context.Context) IpMacAddressMapOutput
}

type IpMacAddressMap map[string]IpMacAddressInput

func (IpMacAddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpMacAddress)(nil)).Elem()
}

func (i IpMacAddressMap) ToIpMacAddressMapOutput() IpMacAddressMapOutput {
	return i.ToIpMacAddressMapOutputWithContext(context.Background())
}

func (i IpMacAddressMap) ToIpMacAddressMapOutputWithContext(ctx context.Context) IpMacAddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpMacAddressMapOutput)
}

type IpMacAddressOutput struct{ *pulumi.OutputState }

func (IpMacAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpMacAddress)(nil)).Elem()
}

func (o IpMacAddressOutput) ToIpMacAddressOutput() IpMacAddressOutput {
	return o
}

func (o IpMacAddressOutput) ToIpMacAddressOutputWithContext(ctx context.Context) IpMacAddressOutput {
	return o
}

// The Virtual MAC address.
func (o IpMacAddressOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *IpMacAddress) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// The date at which the Virtual Mac Address was created (RFC 3339 format).
func (o IpMacAddressOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IpMacAddress) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The ID of the flexible IP for which to generate a virtual MAC.
func (o IpMacAddressOutput) FlexibleIpId() pulumi.StringOutput {
	return o.ApplyT(func(v *IpMacAddress) pulumi.StringOutput { return v.FlexibleIpId }).(pulumi.StringOutput)
}

// The IDs of the flexible IPs on which to duplicate the virtual MAC.
// > **Important:** The flexible IPs need to be attached to the same server for the operation to work.
func (o IpMacAddressOutput) FlexibleIpIdsToDuplicates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IpMacAddress) pulumi.StringArrayOutput { return v.FlexibleIpIdsToDuplicates }).(pulumi.StringArrayOutput)
}

// The Virtual MAC status.
func (o IpMacAddressOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *IpMacAddress) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The type of the virtual MAC.
func (o IpMacAddressOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IpMacAddress) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The date at which the Virtual Mac Address was last updated (RFC 3339 format).
func (o IpMacAddressOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IpMacAddress) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// The zone of the Virtual Mac Address.
func (o IpMacAddressOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *IpMacAddress) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type IpMacAddressArrayOutput struct{ *pulumi.OutputState }

func (IpMacAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpMacAddress)(nil)).Elem()
}

func (o IpMacAddressArrayOutput) ToIpMacAddressArrayOutput() IpMacAddressArrayOutput {
	return o
}

func (o IpMacAddressArrayOutput) ToIpMacAddressArrayOutputWithContext(ctx context.Context) IpMacAddressArrayOutput {
	return o
}

func (o IpMacAddressArrayOutput) Index(i pulumi.IntInput) IpMacAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpMacAddress {
		return vs[0].([]*IpMacAddress)[vs[1].(int)]
	}).(IpMacAddressOutput)
}

type IpMacAddressMapOutput struct{ *pulumi.OutputState }

func (IpMacAddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpMacAddress)(nil)).Elem()
}

func (o IpMacAddressMapOutput) ToIpMacAddressMapOutput() IpMacAddressMapOutput {
	return o
}

func (o IpMacAddressMapOutput) ToIpMacAddressMapOutputWithContext(ctx context.Context) IpMacAddressMapOutput {
	return o
}

func (o IpMacAddressMapOutput) MapIndex(k pulumi.StringInput) IpMacAddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpMacAddress {
		return vs[0].(map[string]*IpMacAddress)[vs[1].(string)]
	}).(IpMacAddressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpMacAddressInput)(nil)).Elem(), &IpMacAddress{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpMacAddressArrayInput)(nil)).Elem(), IpMacAddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpMacAddressMapInput)(nil)).Elem(), IpMacAddressMap{})
	pulumi.RegisterOutputType(IpMacAddressOutput{})
	pulumi.RegisterOutputType(IpMacAddressArrayOutput{})
	pulumi.RegisterOutputType(IpMacAddressMapOutput{})
}
