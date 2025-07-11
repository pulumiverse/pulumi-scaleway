// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package databases

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Read Replicas.
// For more information refer to the [API documentation](https://www.scaleway.com/en/developers/api/managed-database-postgre-mysql/).
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
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/databases"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			instance, err := databases.NewInstance(ctx, "instance", &databases.InstanceArgs{
//				Name:          pulumi.String("test-rdb-rr-update"),
//				NodeType:      pulumi.String("db-dev-s"),
//				Engine:        pulumi.String("PostgreSQL-14"),
//				IsHaCluster:   pulumi.Bool(false),
//				DisableBackup: pulumi.Bool(true),
//				UserName:      pulumi.String("my_initial_user"),
//				Password:      pulumi.String("thiZ_is_v&ry_s3cret"),
//				Tags: pulumi.StringArray{
//					pulumi.String("terraform-test"),
//					pulumi.String("scaleway_rdb_read_replica"),
//					pulumi.String("minimal"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = databases.NewReadReplica(ctx, "replica", &databases.ReadReplicaArgs{
//				InstanceId:   instance.ID(),
//				DirectAccess: &databases.ReadReplicaDirectAccessArgs{},
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
// ### Private network with static endpoint
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/databases"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/network"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			instance, err := databases.NewInstance(ctx, "instance", &databases.InstanceArgs{
//				Name:          pulumi.String("rdb_instance"),
//				NodeType:      pulumi.String("db-dev-s"),
//				Engine:        pulumi.String("PostgreSQL-14"),
//				IsHaCluster:   pulumi.Bool(false),
//				DisableBackup: pulumi.Bool(true),
//				UserName:      pulumi.String("my_initial_user"),
//				Password:      pulumi.String("thiZ_is_v&ry_s3cret"),
//			})
//			if err != nil {
//				return err
//			}
//			pn, err := network.NewPrivateNetwork(ctx, "pn", nil)
//			if err != nil {
//				return err
//			}
//			_, err = databases.NewReadReplica(ctx, "replica", &databases.ReadReplicaArgs{
//				InstanceId: instance.ID(),
//				PrivateNetwork: &databases.ReadReplicaPrivateNetworkArgs{
//					PrivateNetworkId: pn.ID(),
//					ServiceIp:        pulumi.String("192.168.1.254/24"),
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
// ### Private network with IPAM
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/databases"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/network"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			instance, err := databases.NewInstance(ctx, "instance", &databases.InstanceArgs{
//				Name:          pulumi.String("rdb_instance"),
//				NodeType:      pulumi.String("db-dev-s"),
//				Engine:        pulumi.String("PostgreSQL-14"),
//				IsHaCluster:   pulumi.Bool(false),
//				DisableBackup: pulumi.Bool(true),
//				UserName:      pulumi.String("my_initial_user"),
//				Password:      pulumi.String("thiZ_is_v&ry_s3cret"),
//			})
//			if err != nil {
//				return err
//			}
//			pn, err := network.NewPrivateNetwork(ctx, "pn", nil)
//			if err != nil {
//				return err
//			}
//			_, err = databases.NewReadReplica(ctx, "replica", &databases.ReadReplicaArgs{
//				InstanceId: instance.ID(),
//				PrivateNetwork: &databases.ReadReplicaPrivateNetworkArgs{
//					PrivateNetworkId: pn.ID(),
//					EnableIpam:       pulumi.Bool(true),
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
// Read Replicas can be imported using the `{region}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:databases/readReplica:ReadReplica rr fr-par/11111111-1111-1111-1111-111111111111
// ```
type ReadReplica struct {
	pulumi.CustomResourceState

	// Creates a direct access endpoint to rdb replica.
	DirectAccess ReadReplicaDirectAccessPtrOutput `pulumi:"directAccess"`
	// UUID of the rdb instance.
	//
	// > **Important:** The replica musts contains at least one `directAccess` or `privateNetwork`. It can contain both.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Create an endpoint in a Private Netork.
	PrivateNetwork ReadReplicaPrivateNetworkPtrOutput `pulumi:"privateNetwork"`
	// `region`) The region
	// in which the Read Replica should be created.
	Region pulumi.StringOutput `pulumi:"region"`
	// Defines whether to create the replica in the same availability zone as the main instance nodes or not.
	SameZone pulumi.BoolPtrOutput `pulumi:"sameZone"`
}

// NewReadReplica registers a new resource with the given unique name, arguments, and options.
func NewReadReplica(ctx *pulumi.Context,
	name string, args *ReadReplicaArgs, opts ...pulumi.ResourceOption) (*ReadReplica, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("scaleway:index/databaseReadReplica:DatabaseReadReplica"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ReadReplica
	err := ctx.RegisterResource("scaleway:databases/readReplica:ReadReplica", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReadReplica gets an existing ReadReplica resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReadReplica(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReadReplicaState, opts ...pulumi.ResourceOption) (*ReadReplica, error) {
	var resource ReadReplica
	err := ctx.ReadResource("scaleway:databases/readReplica:ReadReplica", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReadReplica resources.
type readReplicaState struct {
	// Creates a direct access endpoint to rdb replica.
	DirectAccess *ReadReplicaDirectAccess `pulumi:"directAccess"`
	// UUID of the rdb instance.
	//
	// > **Important:** The replica musts contains at least one `directAccess` or `privateNetwork`. It can contain both.
	InstanceId *string `pulumi:"instanceId"`
	// Create an endpoint in a Private Netork.
	PrivateNetwork *ReadReplicaPrivateNetwork `pulumi:"privateNetwork"`
	// `region`) The region
	// in which the Read Replica should be created.
	Region *string `pulumi:"region"`
	// Defines whether to create the replica in the same availability zone as the main instance nodes or not.
	SameZone *bool `pulumi:"sameZone"`
}

type ReadReplicaState struct {
	// Creates a direct access endpoint to rdb replica.
	DirectAccess ReadReplicaDirectAccessPtrInput
	// UUID of the rdb instance.
	//
	// > **Important:** The replica musts contains at least one `directAccess` or `privateNetwork`. It can contain both.
	InstanceId pulumi.StringPtrInput
	// Create an endpoint in a Private Netork.
	PrivateNetwork ReadReplicaPrivateNetworkPtrInput
	// `region`) The region
	// in which the Read Replica should be created.
	Region pulumi.StringPtrInput
	// Defines whether to create the replica in the same availability zone as the main instance nodes or not.
	SameZone pulumi.BoolPtrInput
}

func (ReadReplicaState) ElementType() reflect.Type {
	return reflect.TypeOf((*readReplicaState)(nil)).Elem()
}

type readReplicaArgs struct {
	// Creates a direct access endpoint to rdb replica.
	DirectAccess *ReadReplicaDirectAccess `pulumi:"directAccess"`
	// UUID of the rdb instance.
	//
	// > **Important:** The replica musts contains at least one `directAccess` or `privateNetwork`. It can contain both.
	InstanceId string `pulumi:"instanceId"`
	// Create an endpoint in a Private Netork.
	PrivateNetwork *ReadReplicaPrivateNetwork `pulumi:"privateNetwork"`
	// `region`) The region
	// in which the Read Replica should be created.
	Region *string `pulumi:"region"`
	// Defines whether to create the replica in the same availability zone as the main instance nodes or not.
	SameZone *bool `pulumi:"sameZone"`
}

// The set of arguments for constructing a ReadReplica resource.
type ReadReplicaArgs struct {
	// Creates a direct access endpoint to rdb replica.
	DirectAccess ReadReplicaDirectAccessPtrInput
	// UUID of the rdb instance.
	//
	// > **Important:** The replica musts contains at least one `directAccess` or `privateNetwork`. It can contain both.
	InstanceId pulumi.StringInput
	// Create an endpoint in a Private Netork.
	PrivateNetwork ReadReplicaPrivateNetworkPtrInput
	// `region`) The region
	// in which the Read Replica should be created.
	Region pulumi.StringPtrInput
	// Defines whether to create the replica in the same availability zone as the main instance nodes or not.
	SameZone pulumi.BoolPtrInput
}

func (ReadReplicaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*readReplicaArgs)(nil)).Elem()
}

type ReadReplicaInput interface {
	pulumi.Input

	ToReadReplicaOutput() ReadReplicaOutput
	ToReadReplicaOutputWithContext(ctx context.Context) ReadReplicaOutput
}

func (*ReadReplica) ElementType() reflect.Type {
	return reflect.TypeOf((**ReadReplica)(nil)).Elem()
}

func (i *ReadReplica) ToReadReplicaOutput() ReadReplicaOutput {
	return i.ToReadReplicaOutputWithContext(context.Background())
}

func (i *ReadReplica) ToReadReplicaOutputWithContext(ctx context.Context) ReadReplicaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadReplicaOutput)
}

// ReadReplicaArrayInput is an input type that accepts ReadReplicaArray and ReadReplicaArrayOutput values.
// You can construct a concrete instance of `ReadReplicaArrayInput` via:
//
//	ReadReplicaArray{ ReadReplicaArgs{...} }
type ReadReplicaArrayInput interface {
	pulumi.Input

	ToReadReplicaArrayOutput() ReadReplicaArrayOutput
	ToReadReplicaArrayOutputWithContext(context.Context) ReadReplicaArrayOutput
}

type ReadReplicaArray []ReadReplicaInput

func (ReadReplicaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReadReplica)(nil)).Elem()
}

func (i ReadReplicaArray) ToReadReplicaArrayOutput() ReadReplicaArrayOutput {
	return i.ToReadReplicaArrayOutputWithContext(context.Background())
}

func (i ReadReplicaArray) ToReadReplicaArrayOutputWithContext(ctx context.Context) ReadReplicaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadReplicaArrayOutput)
}

// ReadReplicaMapInput is an input type that accepts ReadReplicaMap and ReadReplicaMapOutput values.
// You can construct a concrete instance of `ReadReplicaMapInput` via:
//
//	ReadReplicaMap{ "key": ReadReplicaArgs{...} }
type ReadReplicaMapInput interface {
	pulumi.Input

	ToReadReplicaMapOutput() ReadReplicaMapOutput
	ToReadReplicaMapOutputWithContext(context.Context) ReadReplicaMapOutput
}

type ReadReplicaMap map[string]ReadReplicaInput

func (ReadReplicaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReadReplica)(nil)).Elem()
}

func (i ReadReplicaMap) ToReadReplicaMapOutput() ReadReplicaMapOutput {
	return i.ToReadReplicaMapOutputWithContext(context.Background())
}

func (i ReadReplicaMap) ToReadReplicaMapOutputWithContext(ctx context.Context) ReadReplicaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadReplicaMapOutput)
}

type ReadReplicaOutput struct{ *pulumi.OutputState }

func (ReadReplicaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReadReplica)(nil)).Elem()
}

func (o ReadReplicaOutput) ToReadReplicaOutput() ReadReplicaOutput {
	return o
}

func (o ReadReplicaOutput) ToReadReplicaOutputWithContext(ctx context.Context) ReadReplicaOutput {
	return o
}

// Creates a direct access endpoint to rdb replica.
func (o ReadReplicaOutput) DirectAccess() ReadReplicaDirectAccessPtrOutput {
	return o.ApplyT(func(v *ReadReplica) ReadReplicaDirectAccessPtrOutput { return v.DirectAccess }).(ReadReplicaDirectAccessPtrOutput)
}

// UUID of the rdb instance.
//
// > **Important:** The replica musts contains at least one `directAccess` or `privateNetwork`. It can contain both.
func (o ReadReplicaOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadReplica) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Create an endpoint in a Private Netork.
func (o ReadReplicaOutput) PrivateNetwork() ReadReplicaPrivateNetworkPtrOutput {
	return o.ApplyT(func(v *ReadReplica) ReadReplicaPrivateNetworkPtrOutput { return v.PrivateNetwork }).(ReadReplicaPrivateNetworkPtrOutput)
}

// `region`) The region
// in which the Read Replica should be created.
func (o ReadReplicaOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadReplica) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Defines whether to create the replica in the same availability zone as the main instance nodes or not.
func (o ReadReplicaOutput) SameZone() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ReadReplica) pulumi.BoolPtrOutput { return v.SameZone }).(pulumi.BoolPtrOutput)
}

type ReadReplicaArrayOutput struct{ *pulumi.OutputState }

func (ReadReplicaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReadReplica)(nil)).Elem()
}

func (o ReadReplicaArrayOutput) ToReadReplicaArrayOutput() ReadReplicaArrayOutput {
	return o
}

func (o ReadReplicaArrayOutput) ToReadReplicaArrayOutputWithContext(ctx context.Context) ReadReplicaArrayOutput {
	return o
}

func (o ReadReplicaArrayOutput) Index(i pulumi.IntInput) ReadReplicaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReadReplica {
		return vs[0].([]*ReadReplica)[vs[1].(int)]
	}).(ReadReplicaOutput)
}

type ReadReplicaMapOutput struct{ *pulumi.OutputState }

func (ReadReplicaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReadReplica)(nil)).Elem()
}

func (o ReadReplicaMapOutput) ToReadReplicaMapOutput() ReadReplicaMapOutput {
	return o
}

func (o ReadReplicaMapOutput) ToReadReplicaMapOutputWithContext(ctx context.Context) ReadReplicaMapOutput {
	return o
}

func (o ReadReplicaMapOutput) MapIndex(k pulumi.StringInput) ReadReplicaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReadReplica {
		return vs[0].(map[string]*ReadReplica)[vs[1].(string)]
	}).(ReadReplicaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReadReplicaInput)(nil)).Elem(), &ReadReplica{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReadReplicaArrayInput)(nil)).Elem(), ReadReplicaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReadReplicaMapInput)(nil)).Elem(), ReadReplicaMap{})
	pulumi.RegisterOutputType(ReadReplicaOutput{})
	pulumi.RegisterOutputType(ReadReplicaArrayOutput{})
	pulumi.RegisterOutputType(ReadReplicaMapOutput{})
}
