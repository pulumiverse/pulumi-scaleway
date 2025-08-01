// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Scaleway VPC Public Gateway public (flexible) IP addresses.
// For more information, see the [API documentation](https://www.scaleway.com/en/developers/api/public-gateway/#path-ips-list-ips).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/domain"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/network"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := network.NewPublicGatewayIp(ctx, "main", &network.PublicGatewayIpArgs{
//				Reverse: pulumi.String("tf.example.com"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = domain.NewRecord(ctx, "tf_A", &domain.RecordArgs{
//				DnsZone:  pulumi.String("example.com"),
//				Name:     pulumi.String("tf"),
//				Type:     pulumi.String("A"),
//				Data:     main.Address,
//				Ttl:      pulumi.Int(3600),
//				Priority: pulumi.Int(1),
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
// Public Gateway IPs can be imported using `{zone}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:network/publicGatewayIp:PublicGatewayIp main fr-par-1/11111111-1111-1111-1111-111111111111
// ```
type PublicGatewayIp struct {
	pulumi.CustomResourceState

	// The IP address itself.
	Address pulumi.StringOutput `pulumi:"address"`
	// The date and time of the creation of the Public Gateway IP.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The Organization ID the Public Gateway IP is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// `projectId`) The ID of the Project the Public Gateway IP is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The reverse domain name for the IP address
	Reverse pulumi.StringOutput `pulumi:"reverse"`
	// The tags associated with the Public Gateway IP.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The date and time of the last update of the Public Gateway IP.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// `zone`) The zone in which the Public Gateway IP should be created.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewPublicGatewayIp registers a new resource with the given unique name, arguments, and options.
func NewPublicGatewayIp(ctx *pulumi.Context,
	name string, args *PublicGatewayIpArgs, opts ...pulumi.ResourceOption) (*PublicGatewayIp, error) {
	if args == nil {
		args = &PublicGatewayIpArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("scaleway:index/vpcPublicGatewayIp:VpcPublicGatewayIp"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PublicGatewayIp
	err := ctx.RegisterResource("scaleway:network/publicGatewayIp:PublicGatewayIp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublicGatewayIp gets an existing PublicGatewayIp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicGatewayIp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicGatewayIpState, opts ...pulumi.ResourceOption) (*PublicGatewayIp, error) {
	var resource PublicGatewayIp
	err := ctx.ReadResource("scaleway:network/publicGatewayIp:PublicGatewayIp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublicGatewayIp resources.
type publicGatewayIpState struct {
	// The IP address itself.
	Address *string `pulumi:"address"`
	// The date and time of the creation of the Public Gateway IP.
	CreatedAt *string `pulumi:"createdAt"`
	// The Organization ID the Public Gateway IP is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// `projectId`) The ID of the Project the Public Gateway IP is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The reverse domain name for the IP address
	Reverse *string `pulumi:"reverse"`
	// The tags associated with the Public Gateway IP.
	Tags []string `pulumi:"tags"`
	// The date and time of the last update of the Public Gateway IP.
	UpdatedAt *string `pulumi:"updatedAt"`
	// `zone`) The zone in which the Public Gateway IP should be created.
	Zone *string `pulumi:"zone"`
}

type PublicGatewayIpState struct {
	// The IP address itself.
	Address pulumi.StringPtrInput
	// The date and time of the creation of the Public Gateway IP.
	CreatedAt pulumi.StringPtrInput
	// The Organization ID the Public Gateway IP is associated with.
	OrganizationId pulumi.StringPtrInput
	// `projectId`) The ID of the Project the Public Gateway IP is associated with.
	ProjectId pulumi.StringPtrInput
	// The reverse domain name for the IP address
	Reverse pulumi.StringPtrInput
	// The tags associated with the Public Gateway IP.
	Tags pulumi.StringArrayInput
	// The date and time of the last update of the Public Gateway IP.
	UpdatedAt pulumi.StringPtrInput
	// `zone`) The zone in which the Public Gateway IP should be created.
	Zone pulumi.StringPtrInput
}

func (PublicGatewayIpState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicGatewayIpState)(nil)).Elem()
}

type publicGatewayIpArgs struct {
	// `projectId`) The ID of the Project the Public Gateway IP is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The reverse domain name for the IP address
	Reverse *string `pulumi:"reverse"`
	// The tags associated with the Public Gateway IP.
	Tags []string `pulumi:"tags"`
	// `zone`) The zone in which the Public Gateway IP should be created.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a PublicGatewayIp resource.
type PublicGatewayIpArgs struct {
	// `projectId`) The ID of the Project the Public Gateway IP is associated with.
	ProjectId pulumi.StringPtrInput
	// The reverse domain name for the IP address
	Reverse pulumi.StringPtrInput
	// The tags associated with the Public Gateway IP.
	Tags pulumi.StringArrayInput
	// `zone`) The zone in which the Public Gateway IP should be created.
	Zone pulumi.StringPtrInput
}

func (PublicGatewayIpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicGatewayIpArgs)(nil)).Elem()
}

type PublicGatewayIpInput interface {
	pulumi.Input

	ToPublicGatewayIpOutput() PublicGatewayIpOutput
	ToPublicGatewayIpOutputWithContext(ctx context.Context) PublicGatewayIpOutput
}

func (*PublicGatewayIp) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicGatewayIp)(nil)).Elem()
}

func (i *PublicGatewayIp) ToPublicGatewayIpOutput() PublicGatewayIpOutput {
	return i.ToPublicGatewayIpOutputWithContext(context.Background())
}

func (i *PublicGatewayIp) ToPublicGatewayIpOutputWithContext(ctx context.Context) PublicGatewayIpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicGatewayIpOutput)
}

// PublicGatewayIpArrayInput is an input type that accepts PublicGatewayIpArray and PublicGatewayIpArrayOutput values.
// You can construct a concrete instance of `PublicGatewayIpArrayInput` via:
//
//	PublicGatewayIpArray{ PublicGatewayIpArgs{...} }
type PublicGatewayIpArrayInput interface {
	pulumi.Input

	ToPublicGatewayIpArrayOutput() PublicGatewayIpArrayOutput
	ToPublicGatewayIpArrayOutputWithContext(context.Context) PublicGatewayIpArrayOutput
}

type PublicGatewayIpArray []PublicGatewayIpInput

func (PublicGatewayIpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PublicGatewayIp)(nil)).Elem()
}

func (i PublicGatewayIpArray) ToPublicGatewayIpArrayOutput() PublicGatewayIpArrayOutput {
	return i.ToPublicGatewayIpArrayOutputWithContext(context.Background())
}

func (i PublicGatewayIpArray) ToPublicGatewayIpArrayOutputWithContext(ctx context.Context) PublicGatewayIpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicGatewayIpArrayOutput)
}

// PublicGatewayIpMapInput is an input type that accepts PublicGatewayIpMap and PublicGatewayIpMapOutput values.
// You can construct a concrete instance of `PublicGatewayIpMapInput` via:
//
//	PublicGatewayIpMap{ "key": PublicGatewayIpArgs{...} }
type PublicGatewayIpMapInput interface {
	pulumi.Input

	ToPublicGatewayIpMapOutput() PublicGatewayIpMapOutput
	ToPublicGatewayIpMapOutputWithContext(context.Context) PublicGatewayIpMapOutput
}

type PublicGatewayIpMap map[string]PublicGatewayIpInput

func (PublicGatewayIpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PublicGatewayIp)(nil)).Elem()
}

func (i PublicGatewayIpMap) ToPublicGatewayIpMapOutput() PublicGatewayIpMapOutput {
	return i.ToPublicGatewayIpMapOutputWithContext(context.Background())
}

func (i PublicGatewayIpMap) ToPublicGatewayIpMapOutputWithContext(ctx context.Context) PublicGatewayIpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicGatewayIpMapOutput)
}

type PublicGatewayIpOutput struct{ *pulumi.OutputState }

func (PublicGatewayIpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicGatewayIp)(nil)).Elem()
}

func (o PublicGatewayIpOutput) ToPublicGatewayIpOutput() PublicGatewayIpOutput {
	return o
}

func (o PublicGatewayIpOutput) ToPublicGatewayIpOutputWithContext(ctx context.Context) PublicGatewayIpOutput {
	return o
}

// The IP address itself.
func (o PublicGatewayIpOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicGatewayIp) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// The date and time of the creation of the Public Gateway IP.
func (o PublicGatewayIpOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicGatewayIp) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The Organization ID the Public Gateway IP is associated with.
func (o PublicGatewayIpOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicGatewayIp) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// `projectId`) The ID of the Project the Public Gateway IP is associated with.
func (o PublicGatewayIpOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicGatewayIp) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The reverse domain name for the IP address
func (o PublicGatewayIpOutput) Reverse() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicGatewayIp) pulumi.StringOutput { return v.Reverse }).(pulumi.StringOutput)
}

// The tags associated with the Public Gateway IP.
func (o PublicGatewayIpOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PublicGatewayIp) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The date and time of the last update of the Public Gateway IP.
func (o PublicGatewayIpOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicGatewayIp) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// `zone`) The zone in which the Public Gateway IP should be created.
func (o PublicGatewayIpOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicGatewayIp) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type PublicGatewayIpArrayOutput struct{ *pulumi.OutputState }

func (PublicGatewayIpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PublicGatewayIp)(nil)).Elem()
}

func (o PublicGatewayIpArrayOutput) ToPublicGatewayIpArrayOutput() PublicGatewayIpArrayOutput {
	return o
}

func (o PublicGatewayIpArrayOutput) ToPublicGatewayIpArrayOutputWithContext(ctx context.Context) PublicGatewayIpArrayOutput {
	return o
}

func (o PublicGatewayIpArrayOutput) Index(i pulumi.IntInput) PublicGatewayIpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PublicGatewayIp {
		return vs[0].([]*PublicGatewayIp)[vs[1].(int)]
	}).(PublicGatewayIpOutput)
}

type PublicGatewayIpMapOutput struct{ *pulumi.OutputState }

func (PublicGatewayIpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PublicGatewayIp)(nil)).Elem()
}

func (o PublicGatewayIpMapOutput) ToPublicGatewayIpMapOutput() PublicGatewayIpMapOutput {
	return o
}

func (o PublicGatewayIpMapOutput) ToPublicGatewayIpMapOutputWithContext(ctx context.Context) PublicGatewayIpMapOutput {
	return o
}

func (o PublicGatewayIpMapOutput) MapIndex(k pulumi.StringInput) PublicGatewayIpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PublicGatewayIp {
		return vs[0].(map[string]*PublicGatewayIp)[vs[1].(string)]
	}).(PublicGatewayIpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PublicGatewayIpInput)(nil)).Elem(), &PublicGatewayIp{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicGatewayIpArrayInput)(nil)).Elem(), PublicGatewayIpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicGatewayIpMapInput)(nil)).Elem(), PublicGatewayIpMap{})
	pulumi.RegisterOutputType(PublicGatewayIpOutput{})
	pulumi.RegisterOutputType(PublicGatewayIpArrayOutput{})
	pulumi.RegisterOutputType(PublicGatewayIpMapOutput{})
}
