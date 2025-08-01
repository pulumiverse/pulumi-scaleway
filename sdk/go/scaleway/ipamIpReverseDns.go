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
// $ pulumi import scaleway:index/ipamIpReverseDns:IpamIpReverseDns main fr-par/11111111-1111-1111-1111-111111111111
// ```
//
// Deprecated: scaleway.index/ipamipreversedns.IpamIpReverseDns has been deprecated in favor of scaleway.ipam/ipreversedns.IpReverseDns
type IpamIpReverseDns struct {
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

// NewIpamIpReverseDns registers a new resource with the given unique name, arguments, and options.
func NewIpamIpReverseDns(ctx *pulumi.Context,
	name string, args *IpamIpReverseDnsArgs, opts ...pulumi.ResourceOption) (*IpamIpReverseDns, error) {
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
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IpamIpReverseDns
	err := ctx.RegisterResource("scaleway:index/ipamIpReverseDns:IpamIpReverseDns", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpamIpReverseDns gets an existing IpamIpReverseDns resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpamIpReverseDns(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpamIpReverseDnsState, opts ...pulumi.ResourceOption) (*IpamIpReverseDns, error) {
	var resource IpamIpReverseDns
	err := ctx.ReadResource("scaleway:index/ipamIpReverseDns:IpamIpReverseDns", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpamIpReverseDns resources.
type ipamIpReverseDnsState struct {
	// The IP corresponding to the hostname.
	Address *string `pulumi:"address"`
	// The reverse domain name.
	Hostname *string `pulumi:"hostname"`
	// The IPAM IP ID.
	IpamIpId *string `pulumi:"ipamIpId"`
	// `region`) The region of the IP reverse DNS.
	Region *string `pulumi:"region"`
}

type IpamIpReverseDnsState struct {
	// The IP corresponding to the hostname.
	Address pulumi.StringPtrInput
	// The reverse domain name.
	Hostname pulumi.StringPtrInput
	// The IPAM IP ID.
	IpamIpId pulumi.StringPtrInput
	// `region`) The region of the IP reverse DNS.
	Region pulumi.StringPtrInput
}

func (IpamIpReverseDnsState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipamIpReverseDnsState)(nil)).Elem()
}

type ipamIpReverseDnsArgs struct {
	// The IP corresponding to the hostname.
	Address string `pulumi:"address"`
	// The reverse domain name.
	Hostname string `pulumi:"hostname"`
	// The IPAM IP ID.
	IpamIpId string `pulumi:"ipamIpId"`
	// `region`) The region of the IP reverse DNS.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a IpamIpReverseDns resource.
type IpamIpReverseDnsArgs struct {
	// The IP corresponding to the hostname.
	Address pulumi.StringInput
	// The reverse domain name.
	Hostname pulumi.StringInput
	// The IPAM IP ID.
	IpamIpId pulumi.StringInput
	// `region`) The region of the IP reverse DNS.
	Region pulumi.StringPtrInput
}

func (IpamIpReverseDnsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipamIpReverseDnsArgs)(nil)).Elem()
}

type IpamIpReverseDnsInput interface {
	pulumi.Input

	ToIpamIpReverseDnsOutput() IpamIpReverseDnsOutput
	ToIpamIpReverseDnsOutputWithContext(ctx context.Context) IpamIpReverseDnsOutput
}

func (*IpamIpReverseDns) ElementType() reflect.Type {
	return reflect.TypeOf((**IpamIpReverseDns)(nil)).Elem()
}

func (i *IpamIpReverseDns) ToIpamIpReverseDnsOutput() IpamIpReverseDnsOutput {
	return i.ToIpamIpReverseDnsOutputWithContext(context.Background())
}

func (i *IpamIpReverseDns) ToIpamIpReverseDnsOutputWithContext(ctx context.Context) IpamIpReverseDnsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpamIpReverseDnsOutput)
}

// IpamIpReverseDnsArrayInput is an input type that accepts IpamIpReverseDnsArray and IpamIpReverseDnsArrayOutput values.
// You can construct a concrete instance of `IpamIpReverseDnsArrayInput` via:
//
//	IpamIpReverseDnsArray{ IpamIpReverseDnsArgs{...} }
type IpamIpReverseDnsArrayInput interface {
	pulumi.Input

	ToIpamIpReverseDnsArrayOutput() IpamIpReverseDnsArrayOutput
	ToIpamIpReverseDnsArrayOutputWithContext(context.Context) IpamIpReverseDnsArrayOutput
}

type IpamIpReverseDnsArray []IpamIpReverseDnsInput

func (IpamIpReverseDnsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpamIpReverseDns)(nil)).Elem()
}

func (i IpamIpReverseDnsArray) ToIpamIpReverseDnsArrayOutput() IpamIpReverseDnsArrayOutput {
	return i.ToIpamIpReverseDnsArrayOutputWithContext(context.Background())
}

func (i IpamIpReverseDnsArray) ToIpamIpReverseDnsArrayOutputWithContext(ctx context.Context) IpamIpReverseDnsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpamIpReverseDnsArrayOutput)
}

// IpamIpReverseDnsMapInput is an input type that accepts IpamIpReverseDnsMap and IpamIpReverseDnsMapOutput values.
// You can construct a concrete instance of `IpamIpReverseDnsMapInput` via:
//
//	IpamIpReverseDnsMap{ "key": IpamIpReverseDnsArgs{...} }
type IpamIpReverseDnsMapInput interface {
	pulumi.Input

	ToIpamIpReverseDnsMapOutput() IpamIpReverseDnsMapOutput
	ToIpamIpReverseDnsMapOutputWithContext(context.Context) IpamIpReverseDnsMapOutput
}

type IpamIpReverseDnsMap map[string]IpamIpReverseDnsInput

func (IpamIpReverseDnsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpamIpReverseDns)(nil)).Elem()
}

func (i IpamIpReverseDnsMap) ToIpamIpReverseDnsMapOutput() IpamIpReverseDnsMapOutput {
	return i.ToIpamIpReverseDnsMapOutputWithContext(context.Background())
}

func (i IpamIpReverseDnsMap) ToIpamIpReverseDnsMapOutputWithContext(ctx context.Context) IpamIpReverseDnsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpamIpReverseDnsMapOutput)
}

type IpamIpReverseDnsOutput struct{ *pulumi.OutputState }

func (IpamIpReverseDnsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpamIpReverseDns)(nil)).Elem()
}

func (o IpamIpReverseDnsOutput) ToIpamIpReverseDnsOutput() IpamIpReverseDnsOutput {
	return o
}

func (o IpamIpReverseDnsOutput) ToIpamIpReverseDnsOutputWithContext(ctx context.Context) IpamIpReverseDnsOutput {
	return o
}

// The IP corresponding to the hostname.
func (o IpamIpReverseDnsOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *IpamIpReverseDns) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// The reverse domain name.
func (o IpamIpReverseDnsOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *IpamIpReverseDns) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// The IPAM IP ID.
func (o IpamIpReverseDnsOutput) IpamIpId() pulumi.StringOutput {
	return o.ApplyT(func(v *IpamIpReverseDns) pulumi.StringOutput { return v.IpamIpId }).(pulumi.StringOutput)
}

// `region`) The region of the IP reverse DNS.
func (o IpamIpReverseDnsOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *IpamIpReverseDns) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type IpamIpReverseDnsArrayOutput struct{ *pulumi.OutputState }

func (IpamIpReverseDnsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpamIpReverseDns)(nil)).Elem()
}

func (o IpamIpReverseDnsArrayOutput) ToIpamIpReverseDnsArrayOutput() IpamIpReverseDnsArrayOutput {
	return o
}

func (o IpamIpReverseDnsArrayOutput) ToIpamIpReverseDnsArrayOutputWithContext(ctx context.Context) IpamIpReverseDnsArrayOutput {
	return o
}

func (o IpamIpReverseDnsArrayOutput) Index(i pulumi.IntInput) IpamIpReverseDnsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpamIpReverseDns {
		return vs[0].([]*IpamIpReverseDns)[vs[1].(int)]
	}).(IpamIpReverseDnsOutput)
}

type IpamIpReverseDnsMapOutput struct{ *pulumi.OutputState }

func (IpamIpReverseDnsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpamIpReverseDns)(nil)).Elem()
}

func (o IpamIpReverseDnsMapOutput) ToIpamIpReverseDnsMapOutput() IpamIpReverseDnsMapOutput {
	return o
}

func (o IpamIpReverseDnsMapOutput) ToIpamIpReverseDnsMapOutputWithContext(ctx context.Context) IpamIpReverseDnsMapOutput {
	return o
}

func (o IpamIpReverseDnsMapOutput) MapIndex(k pulumi.StringInput) IpamIpReverseDnsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpamIpReverseDns {
		return vs[0].(map[string]*IpamIpReverseDns)[vs[1].(string)]
	}).(IpamIpReverseDnsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpamIpReverseDnsInput)(nil)).Elem(), &IpamIpReverseDns{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpamIpReverseDnsArrayInput)(nil)).Elem(), IpamIpReverseDnsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpamIpReverseDnsMapInput)(nil)).Elem(), IpamIpReverseDnsMap{})
	pulumi.RegisterOutputType(IpamIpReverseDnsOutput{})
	pulumi.RegisterOutputType(IpamIpReverseDnsArrayOutput{})
	pulumi.RegisterOutputType(IpamIpReverseDnsMapOutput{})
}
