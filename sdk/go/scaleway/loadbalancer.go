// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway Load-Balancers.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/lb/zoned_api).
//
// ## Examples
//
// ### Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-scaleway/sdk/go/scaleway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		ip, err := scaleway.NewLoadbalancerIp(ctx, "ip", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = scaleway.NewLoadbalancer(ctx, "base", &scaleway.LoadbalancerArgs{
// 			IpId:      ip.ID(),
// 			Zone:      pulumi.String("fr-par-1"),
// 			Type:      pulumi.String("LB-S"),
// 			ReleaseIp: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## IP ID
//
// Since v1.15.0, `ipId` is a required field. This means that now a separate `LoadbalancerIp` is required.
// When importing, the IP needs to be imported as well as the LB.
// When upgrading to v1.15.0, you will need to create a new `LoadbalancerIp` resource and import it.
//
// For instance, if you had the following:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-scaleway/sdk/go/scaleway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := scaleway.NewLoadbalancer(ctx, "base", &scaleway.LoadbalancerArgs{
// 			Type: pulumi.String("LB-S"),
// 			Zone: pulumi.String("fr-par-1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// You will need to update it to:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-scaleway/sdk/go/scaleway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		ip, err := scaleway.NewLoadbalancerIp(ctx, "ip", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = scaleway.NewLoadbalancer(ctx, "base", &scaleway.LoadbalancerArgs{
// 			IpId:      ip.ID(),
// 			Zone:      pulumi.String("fr-par-1"),
// 			Type:      pulumi.String("LB-S"),
// 			ReleaseIp: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Private Network
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-scaleway/sdk/go/scaleway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		ip01, err := scaleway.NewLoadbalancerIp(ctx, "ip01", nil)
// 		if err != nil {
// 			return err
// 		}
// 		pnLB01, err := scaleway.NewVpcPrivateNetwork(ctx, "pnLB01", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = scaleway.NewLoadbalancer(ctx, "lb01", &scaleway.LoadbalancerArgs{
// 			IpId:      ip01.ID(),
// 			Type:      pulumi.String("LB-S"),
// 			ReleaseIp: pulumi.Bool(false),
// 			PrivateNetworks: LoadbalancerPrivateNetworkArray{
// 				&LoadbalancerPrivateNetworkArgs{
// 					PrivateNetworkId: pnLB01.ID(),
// 					StaticConfigs: pulumi.StringArray{
// 						pulumi.String("172.16.0.100"),
// 						pulumi.String("172.16.0.101"),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// > **Important:** Updates to `privateNetwork` will recreate the attachment.
//
// - `privateNetworkId` - (Required) The ID of the Private Network to associate.
//
// - `staticConfig` - (Optional) Define two local ip address of your choice for each load balancer instance. See below.
//
// - `dhcpConfig` - (Optional) Set to true if you want to let DHCP assign IP addresses. See below.
//
// > **Important:**  Only one of staticConfig and dhcpConfig may be set.
//
// - `zone` - (Defaults to provider `zone`) The zone in which the private network was created.
//
// ## Import
//
// Load-Balancer can be imported using the `{zone}/{id}`, e.g. bash
//
// ```sh
//  $ pulumi import scaleway:index/loadbalancer:Loadbalancer lb01 fr-par-1/11111111-1111-1111-1111-111111111111
// ```
//
//  Be aware that you will also need to import the `scaleway_lb_ip` resource.
type Loadbalancer struct {
	pulumi.CustomResourceState

	// The load-balance public IP Address
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The ID of the associated IP. See below.
	IpId pulumi.StringOutput `pulumi:"ipId"`
	// The name of the load-balancer.
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization ID the load-balancer is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// List of private network to connect with your load balancer
	PrivateNetworks LoadbalancerPrivateNetworkArrayOutput `pulumi:"privateNetworks"`
	// `projectId`) The ID of the project the load-balancer is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The region of the resource
	Region pulumi.StringOutput `pulumi:"region"`
	// The releaseIp allow release the ip address associated with the load-balancers.
	ReleaseIp pulumi.BoolPtrOutput `pulumi:"releaseIp"`
	// The tags associated with the load-balancers.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The type of the load-balancer.
	Type pulumi.StringOutput `pulumi:"type"`
	// `zone`) The zone in which the IP should be reserved.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewLoadbalancer registers a new resource with the given unique name, arguments, and options.
func NewLoadbalancer(ctx *pulumi.Context,
	name string, args *LoadbalancerArgs, opts ...pulumi.ResourceOption) (*Loadbalancer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpId == nil {
		return nil, errors.New("invalid value for required argument 'IpId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource Loadbalancer
	err := ctx.RegisterResource("scaleway:index/loadbalancer:Loadbalancer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadbalancer gets an existing Loadbalancer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadbalancer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadbalancerState, opts ...pulumi.ResourceOption) (*Loadbalancer, error) {
	var resource Loadbalancer
	err := ctx.ReadResource("scaleway:index/loadbalancer:Loadbalancer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Loadbalancer resources.
type loadbalancerState struct {
	// The load-balance public IP Address
	IpAddress *string `pulumi:"ipAddress"`
	// The ID of the associated IP. See below.
	IpId *string `pulumi:"ipId"`
	// The name of the load-balancer.
	Name *string `pulumi:"name"`
	// The organization ID the load-balancer is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// List of private network to connect with your load balancer
	PrivateNetworks []LoadbalancerPrivateNetwork `pulumi:"privateNetworks"`
	// `projectId`) The ID of the project the load-balancer is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The region of the resource
	Region *string `pulumi:"region"`
	// The releaseIp allow release the ip address associated with the load-balancers.
	ReleaseIp *bool `pulumi:"releaseIp"`
	// The tags associated with the load-balancers.
	Tags []string `pulumi:"tags"`
	// The type of the load-balancer.
	Type *string `pulumi:"type"`
	// `zone`) The zone in which the IP should be reserved.
	Zone *string `pulumi:"zone"`
}

type LoadbalancerState struct {
	// The load-balance public IP Address
	IpAddress pulumi.StringPtrInput
	// The ID of the associated IP. See below.
	IpId pulumi.StringPtrInput
	// The name of the load-balancer.
	Name pulumi.StringPtrInput
	// The organization ID the load-balancer is associated with.
	OrganizationId pulumi.StringPtrInput
	// List of private network to connect with your load balancer
	PrivateNetworks LoadbalancerPrivateNetworkArrayInput
	// `projectId`) The ID of the project the load-balancer is associated with.
	ProjectId pulumi.StringPtrInput
	// The region of the resource
	Region pulumi.StringPtrInput
	// The releaseIp allow release the ip address associated with the load-balancers.
	ReleaseIp pulumi.BoolPtrInput
	// The tags associated with the load-balancers.
	Tags pulumi.StringArrayInput
	// The type of the load-balancer.
	Type pulumi.StringPtrInput
	// `zone`) The zone in which the IP should be reserved.
	Zone pulumi.StringPtrInput
}

func (LoadbalancerState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadbalancerState)(nil)).Elem()
}

type loadbalancerArgs struct {
	// The ID of the associated IP. See below.
	IpId string `pulumi:"ipId"`
	// The name of the load-balancer.
	Name *string `pulumi:"name"`
	// List of private network to connect with your load balancer
	PrivateNetworks []LoadbalancerPrivateNetwork `pulumi:"privateNetworks"`
	// `projectId`) The ID of the project the load-balancer is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The releaseIp allow release the ip address associated with the load-balancers.
	ReleaseIp *bool `pulumi:"releaseIp"`
	// The tags associated with the load-balancers.
	Tags []string `pulumi:"tags"`
	// The type of the load-balancer.
	Type string `pulumi:"type"`
	// `zone`) The zone in which the IP should be reserved.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a Loadbalancer resource.
type LoadbalancerArgs struct {
	// The ID of the associated IP. See below.
	IpId pulumi.StringInput
	// The name of the load-balancer.
	Name pulumi.StringPtrInput
	// List of private network to connect with your load balancer
	PrivateNetworks LoadbalancerPrivateNetworkArrayInput
	// `projectId`) The ID of the project the load-balancer is associated with.
	ProjectId pulumi.StringPtrInput
	// The releaseIp allow release the ip address associated with the load-balancers.
	ReleaseIp pulumi.BoolPtrInput
	// The tags associated with the load-balancers.
	Tags pulumi.StringArrayInput
	// The type of the load-balancer.
	Type pulumi.StringInput
	// `zone`) The zone in which the IP should be reserved.
	Zone pulumi.StringPtrInput
}

func (LoadbalancerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadbalancerArgs)(nil)).Elem()
}

type LoadbalancerInput interface {
	pulumi.Input

	ToLoadbalancerOutput() LoadbalancerOutput
	ToLoadbalancerOutputWithContext(ctx context.Context) LoadbalancerOutput
}

func (*Loadbalancer) ElementType() reflect.Type {
	return reflect.TypeOf((**Loadbalancer)(nil)).Elem()
}

func (i *Loadbalancer) ToLoadbalancerOutput() LoadbalancerOutput {
	return i.ToLoadbalancerOutputWithContext(context.Background())
}

func (i *Loadbalancer) ToLoadbalancerOutputWithContext(ctx context.Context) LoadbalancerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadbalancerOutput)
}

type LoadbalancerOutput struct{ *pulumi.OutputState }

func (LoadbalancerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Loadbalancer)(nil)).Elem()
}

func (o LoadbalancerOutput) ToLoadbalancerOutput() LoadbalancerOutput {
	return o
}

func (o LoadbalancerOutput) ToLoadbalancerOutputWithContext(ctx context.Context) LoadbalancerOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoadbalancerInput)(nil)).Elem(), &Loadbalancer{})
	pulumi.RegisterOutputType(LoadbalancerOutput{})
}