// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ipam

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Manage the reverse DNS of IP addresses managed by Scaleway's IP Address Management (IPAM) service.
//
// For more information about IPAM, see the main [documentation](https://www.scaleway.com/en/docs/vpc/concepts/#ipam).
//
// ## Import
//
// IPAM IP reverse DNS can be imported using `{region}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:ipam/ipReverseDns:IpReverseDns main fr-par/11111111-1111-1111-1111-111111111111
// ```
type IpReverseDns struct {
	pulumi.CustomResourceState

	// The IP corresponding to the hostname.
	Address pulumi.StringOutput `pulumi:"address"`
	// The reverse domain name.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// The IPAM IP ID.
	IpamIpId pulumi.StringOutput `pulumi:"ipamIpId"`
	// `region`) The region of the IP reverse DNS.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewIpReverseDns registers a new resource with the given unique name, arguments, and options.
func NewIpReverseDns(ctx *pulumi.Context,
	name string, args *IpReverseDnsArgs, opts ...pulumi.ResourceOption) (*IpReverseDns, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Address == nil {
		return nil, errors.New("invalid value for required argument 'Address'")
	}
	if args.Hostname == nil {
		return nil, errors.New("invalid value for required argument 'Hostname'")
	}
	if args.IpamIpId == nil {
		return nil, errors.New("invalid value for required argument 'IpamIpId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("scaleway:index/ipamIpReverseDns:IpamIpReverseDns"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IpReverseDns
	err := ctx.RegisterResource("scaleway:ipam/ipReverseDns:IpReverseDns", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpReverseDns gets an existing IpReverseDns resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpReverseDns(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpReverseDnsState, opts ...pulumi.ResourceOption) (*IpReverseDns, error) {
	var resource IpReverseDns
	err := ctx.ReadResource("scaleway:ipam/ipReverseDns:IpReverseDns", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpReverseDns resources.
type ipReverseDnsState struct {
	// The IP corresponding to the hostname.
	Address *string `pulumi:"address"`
	// The reverse domain name.
	Hostname *string `pulumi:"hostname"`
	// The IPAM IP ID.
	IpamIpId *string `pulumi:"ipamIpId"`
	// `region`) The region of the IP reverse DNS.
	Region *string `pulumi:"region"`
}

type IpReverseDnsState struct {
	// The IP corresponding to the hostname.
	Address pulumi.StringPtrInput
	// The reverse domain name.
	Hostname pulumi.StringPtrInput
	// The IPAM IP ID.
	IpamIpId pulumi.StringPtrInput
	// `region`) The region of the IP reverse DNS.
	Region pulumi.StringPtrInput
}

func (IpReverseDnsState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipReverseDnsState)(nil)).Elem()
}

type ipReverseDnsArgs struct {
	// The IP corresponding to the hostname.
	Address string `pulumi:"address"`
	// The reverse domain name.
	Hostname string `pulumi:"hostname"`
	// The IPAM IP ID.
	IpamIpId string `pulumi:"ipamIpId"`
	// `region`) The region of the IP reverse DNS.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a IpReverseDns resource.
type IpReverseDnsArgs struct {
	// The IP corresponding to the hostname.
	Address pulumi.StringInput
	// The reverse domain name.
	Hostname pulumi.StringInput
	// The IPAM IP ID.
	IpamIpId pulumi.StringInput
	// `region`) The region of the IP reverse DNS.
	Region pulumi.StringPtrInput
}

func (IpReverseDnsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipReverseDnsArgs)(nil)).Elem()
}

type IpReverseDnsInput interface {
	pulumi.Input

	ToIpReverseDnsOutput() IpReverseDnsOutput
	ToIpReverseDnsOutputWithContext(ctx context.Context) IpReverseDnsOutput
}

func (*IpReverseDns) ElementType() reflect.Type {
	return reflect.TypeOf((**IpReverseDns)(nil)).Elem()
}

func (i *IpReverseDns) ToIpReverseDnsOutput() IpReverseDnsOutput {
	return i.ToIpReverseDnsOutputWithContext(context.Background())
}

func (i *IpReverseDns) ToIpReverseDnsOutputWithContext(ctx context.Context) IpReverseDnsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpReverseDnsOutput)
}

// IpReverseDnsArrayInput is an input type that accepts IpReverseDnsArray and IpReverseDnsArrayOutput values.
// You can construct a concrete instance of `IpReverseDnsArrayInput` via:
//
//	IpReverseDnsArray{ IpReverseDnsArgs{...} }
type IpReverseDnsArrayInput interface {
	pulumi.Input

	ToIpReverseDnsArrayOutput() IpReverseDnsArrayOutput
	ToIpReverseDnsArrayOutputWithContext(context.Context) IpReverseDnsArrayOutput
}

type IpReverseDnsArray []IpReverseDnsInput

func (IpReverseDnsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpReverseDns)(nil)).Elem()
}

func (i IpReverseDnsArray) ToIpReverseDnsArrayOutput() IpReverseDnsArrayOutput {
	return i.ToIpReverseDnsArrayOutputWithContext(context.Background())
}

func (i IpReverseDnsArray) ToIpReverseDnsArrayOutputWithContext(ctx context.Context) IpReverseDnsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpReverseDnsArrayOutput)
}

// IpReverseDnsMapInput is an input type that accepts IpReverseDnsMap and IpReverseDnsMapOutput values.
// You can construct a concrete instance of `IpReverseDnsMapInput` via:
//
//	IpReverseDnsMap{ "key": IpReverseDnsArgs{...} }
type IpReverseDnsMapInput interface {
	pulumi.Input

	ToIpReverseDnsMapOutput() IpReverseDnsMapOutput
	ToIpReverseDnsMapOutputWithContext(context.Context) IpReverseDnsMapOutput
}

type IpReverseDnsMap map[string]IpReverseDnsInput

func (IpReverseDnsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpReverseDns)(nil)).Elem()
}

func (i IpReverseDnsMap) ToIpReverseDnsMapOutput() IpReverseDnsMapOutput {
	return i.ToIpReverseDnsMapOutputWithContext(context.Background())
}

func (i IpReverseDnsMap) ToIpReverseDnsMapOutputWithContext(ctx context.Context) IpReverseDnsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpReverseDnsMapOutput)
}

type IpReverseDnsOutput struct{ *pulumi.OutputState }

func (IpReverseDnsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpReverseDns)(nil)).Elem()
}

func (o IpReverseDnsOutput) ToIpReverseDnsOutput() IpReverseDnsOutput {
	return o
}

func (o IpReverseDnsOutput) ToIpReverseDnsOutputWithContext(ctx context.Context) IpReverseDnsOutput {
	return o
}

// The IP corresponding to the hostname.
func (o IpReverseDnsOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *IpReverseDns) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// The reverse domain name.
func (o IpReverseDnsOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *IpReverseDns) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// The IPAM IP ID.
func (o IpReverseDnsOutput) IpamIpId() pulumi.StringOutput {
	return o.ApplyT(func(v *IpReverseDns) pulumi.StringOutput { return v.IpamIpId }).(pulumi.StringOutput)
}

// `region`) The region of the IP reverse DNS.
func (o IpReverseDnsOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *IpReverseDns) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type IpReverseDnsArrayOutput struct{ *pulumi.OutputState }

func (IpReverseDnsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpReverseDns)(nil)).Elem()
}

func (o IpReverseDnsArrayOutput) ToIpReverseDnsArrayOutput() IpReverseDnsArrayOutput {
	return o
}

func (o IpReverseDnsArrayOutput) ToIpReverseDnsArrayOutputWithContext(ctx context.Context) IpReverseDnsArrayOutput {
	return o
}

func (o IpReverseDnsArrayOutput) Index(i pulumi.IntInput) IpReverseDnsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpReverseDns {
		return vs[0].([]*IpReverseDns)[vs[1].(int)]
	}).(IpReverseDnsOutput)
}

type IpReverseDnsMapOutput struct{ *pulumi.OutputState }

func (IpReverseDnsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpReverseDns)(nil)).Elem()
}

func (o IpReverseDnsMapOutput) ToIpReverseDnsMapOutput() IpReverseDnsMapOutput {
	return o
}

func (o IpReverseDnsMapOutput) ToIpReverseDnsMapOutputWithContext(ctx context.Context) IpReverseDnsMapOutput {
	return o
}

func (o IpReverseDnsMapOutput) MapIndex(k pulumi.StringInput) IpReverseDnsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpReverseDns {
		return vs[0].(map[string]*IpReverseDns)[vs[1].(string)]
	}).(IpReverseDnsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpReverseDnsInput)(nil)).Elem(), &IpReverseDns{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpReverseDnsArrayInput)(nil)).Elem(), IpReverseDnsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpReverseDnsMapInput)(nil)).Elem(), IpReverseDnsMap{})
	pulumi.RegisterOutputType(IpReverseDnsOutput{})
	pulumi.RegisterOutputType(IpReverseDnsArrayOutput{})
	pulumi.RegisterOutputType(IpReverseDnsMapOutput{})
}
