// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Scaleway MongoDB® instance.
// For more information refer to the [product documentation](https://www.scaleway.com/en/docs/managed-mongodb-databases/).
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
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/mongodb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := mongodb.NewInstance(ctx, "main", &mongodb.InstanceArgs{
//				Name:           pulumi.String("test-mongodb-basic1"),
//				Version:        pulumi.String("7.0.12"),
//				NodeType:       pulumi.String("MGDB-PLAY2-NANO"),
//				NodeNumber:     pulumi.Int(1),
//				UserName:       pulumi.String("my_initial_user"),
//				Password:       pulumi.String("thiZ_is_v&ry_s3cret"),
//				VolumeSizeInGb: pulumi.Int(5),
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
// ### Private Network
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/mongodb"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/network"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := network.NewPrivateNetwork(ctx, "pn01", &network.PrivateNetworkArgs{
//				Name:   pulumi.String("my_private_network"),
//				Region: pulumi.String("fr-par"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = mongodb.NewInstance(ctx, "main", &mongodb.InstanceArgs{
//				Name:           pulumi.String("test-mongodb-basic1"),
//				Version:        pulumi.String("7.0.12"),
//				NodeType:       pulumi.String("MGDB-PLAY2-NANO"),
//				NodeNumber:     pulumi.Int(1),
//				UserName:       pulumi.String("my_initial_user"),
//				Password:       pulumi.String("thiZ_is_v&ry_s3cret"),
//				VolumeSizeInGb: pulumi.Int(5),
//				PrivateNetwork: &mongodb.InstancePrivateNetworkArgs{
//					PnId: pulumi.Any(pn02.Id),
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
// ### Private Network and Public Network
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/mongodb"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/network"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := network.NewPrivateNetwork(ctx, "pn01", &network.PrivateNetworkArgs{
//				Name:   pulumi.String("my_private_network"),
//				Region: pulumi.String("fr-par"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = mongodb.NewInstance(ctx, "main", &mongodb.InstanceArgs{
//				Name:           pulumi.String("test-mongodb-basic1"),
//				Version:        pulumi.String("7.0.12"),
//				NodeType:       pulumi.String("MGDB-PLAY2-NANO"),
//				NodeNumber:     pulumi.Int(1),
//				UserName:       pulumi.String("my_initial_user"),
//				Password:       pulumi.String("thiZ_is_v&ry_s3cret"),
//				VolumeSizeInGb: pulumi.Int(5),
//				PrivateNetwork: &mongodb.InstancePrivateNetworkArgs{
//					PnId: pulumi.Any(pn02.Id),
//				},
//				PublicNetwork: &mongodb.InstancePublicNetworkArgs{},
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
// ### Restore From Snapshot
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/mongodb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := mongodb.NewInstance(ctx, "restored_instance", &mongodb.InstanceArgs{
//				SnapshotId: pulumi.Any(pn.IdscalewayMongodbSnapshot.MainSnapshot.Id),
//				Name:       pulumi.String("restored-mongodb-from-snapshot"),
//				NodeType:   pulumi.String("MGDB-PLAY2-NANO"),
//				NodeNumber: pulumi.Int(1),
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
// MongoDB® instance can be imported using the `id`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:mongodb/instance:Instance main fr-par-1/11111111-1111-1111-1111-111111111111
// ```
type Instance struct {
	pulumi.CustomResourceState

	// The date and time of the creation of the MongoDB® instance.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Name of the MongoDB® instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// Number of nodes in the instance
	NodeNumber pulumi.IntOutput `pulumi:"nodeNumber"`
	// The type of MongoDB® intance to create.
	NodeType pulumi.StringOutput `pulumi:"nodeType"`
	// Password of the user.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The private IPv4 address associated with the instance.
	PrivateIps InstancePrivateIpArrayOutput `pulumi:"privateIps"`
	// Private Network endpoints of the Database Instance.
	PrivateNetwork InstancePrivateNetworkPtrOutput `pulumi:"privateNetwork"`
	// The projectId you want to attach the resource to
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Public network endpoint configuration (no arguments).
	// > **Important** If neither privateNetwork nor publicNetwork is specified, a public network endpoint is created by default.
	PublicNetwork InstancePublicNetworkOutput `pulumi:"publicNetwork"`
	// The region you want to attach the resource to
	Region pulumi.StringOutput `pulumi:"region"`
	// Map of settings to define for the instance.
	Settings pulumi.StringMapOutput `pulumi:"settings"`
	// Snapshot ID to restore the MongoDB® instance from.
	SnapshotId pulumi.StringPtrOutput `pulumi:"snapshotId"`
	// List of tags attached to the MongoDB® instance.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The PEM-encoded TLS certificate for the MongoDB® instance, if available.
	TlsCertificate pulumi.StringOutput `pulumi:"tlsCertificate"`
	// The date and time of the last update of the MongoDB® instance.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// Name of the user created when the intance is created.
	UserName pulumi.StringPtrOutput `pulumi:"userName"`
	// MongoDB® version of the instance.
	Version pulumi.StringOutput `pulumi:"version"`
	// Volume size in GB.
	VolumeSizeInGb pulumi.IntOutput `pulumi:"volumeSizeInGb"`
	// Volume type of the instance.
	VolumeType pulumi.StringPtrOutput `pulumi:"volumeType"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NodeNumber == nil {
		return nil, errors.New("invalid value for required argument 'NodeNumber'")
	}
	if args.NodeType == nil {
		return nil, errors.New("invalid value for required argument 'NodeType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("scaleway:index/mongoDbInstance:MongoDbInstance"),
		},
	})
	opts = append(opts, aliases)
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Instance
	err := ctx.RegisterResource("scaleway:mongodb/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("scaleway:mongodb/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// The date and time of the creation of the MongoDB® instance.
	CreatedAt *string `pulumi:"createdAt"`
	// Name of the MongoDB® instance.
	Name *string `pulumi:"name"`
	// Number of nodes in the instance
	NodeNumber *int `pulumi:"nodeNumber"`
	// The type of MongoDB® intance to create.
	NodeType *string `pulumi:"nodeType"`
	// Password of the user.
	Password *string `pulumi:"password"`
	// The private IPv4 address associated with the instance.
	PrivateIps []InstancePrivateIp `pulumi:"privateIps"`
	// Private Network endpoints of the Database Instance.
	PrivateNetwork *InstancePrivateNetwork `pulumi:"privateNetwork"`
	// The projectId you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// Public network endpoint configuration (no arguments).
	// > **Important** If neither privateNetwork nor publicNetwork is specified, a public network endpoint is created by default.
	PublicNetwork *InstancePublicNetwork `pulumi:"publicNetwork"`
	// The region you want to attach the resource to
	Region *string `pulumi:"region"`
	// Map of settings to define for the instance.
	Settings map[string]string `pulumi:"settings"`
	// Snapshot ID to restore the MongoDB® instance from.
	SnapshotId *string `pulumi:"snapshotId"`
	// List of tags attached to the MongoDB® instance.
	Tags []string `pulumi:"tags"`
	// The PEM-encoded TLS certificate for the MongoDB® instance, if available.
	TlsCertificate *string `pulumi:"tlsCertificate"`
	// The date and time of the last update of the MongoDB® instance.
	UpdatedAt *string `pulumi:"updatedAt"`
	// Name of the user created when the intance is created.
	UserName *string `pulumi:"userName"`
	// MongoDB® version of the instance.
	Version *string `pulumi:"version"`
	// Volume size in GB.
	VolumeSizeInGb *int `pulumi:"volumeSizeInGb"`
	// Volume type of the instance.
	VolumeType *string `pulumi:"volumeType"`
}

type InstanceState struct {
	// The date and time of the creation of the MongoDB® instance.
	CreatedAt pulumi.StringPtrInput
	// Name of the MongoDB® instance.
	Name pulumi.StringPtrInput
	// Number of nodes in the instance
	NodeNumber pulumi.IntPtrInput
	// The type of MongoDB® intance to create.
	NodeType pulumi.StringPtrInput
	// Password of the user.
	Password pulumi.StringPtrInput
	// The private IPv4 address associated with the instance.
	PrivateIps InstancePrivateIpArrayInput
	// Private Network endpoints of the Database Instance.
	PrivateNetwork InstancePrivateNetworkPtrInput
	// The projectId you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// Public network endpoint configuration (no arguments).
	// > **Important** If neither privateNetwork nor publicNetwork is specified, a public network endpoint is created by default.
	PublicNetwork InstancePublicNetworkPtrInput
	// The region you want to attach the resource to
	Region pulumi.StringPtrInput
	// Map of settings to define for the instance.
	Settings pulumi.StringMapInput
	// Snapshot ID to restore the MongoDB® instance from.
	SnapshotId pulumi.StringPtrInput
	// List of tags attached to the MongoDB® instance.
	Tags pulumi.StringArrayInput
	// The PEM-encoded TLS certificate for the MongoDB® instance, if available.
	TlsCertificate pulumi.StringPtrInput
	// The date and time of the last update of the MongoDB® instance.
	UpdatedAt pulumi.StringPtrInput
	// Name of the user created when the intance is created.
	UserName pulumi.StringPtrInput
	// MongoDB® version of the instance.
	Version pulumi.StringPtrInput
	// Volume size in GB.
	VolumeSizeInGb pulumi.IntPtrInput
	// Volume type of the instance.
	VolumeType pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Name of the MongoDB® instance.
	Name *string `pulumi:"name"`
	// Number of nodes in the instance
	NodeNumber int `pulumi:"nodeNumber"`
	// The type of MongoDB® intance to create.
	NodeType string `pulumi:"nodeType"`
	// Password of the user.
	Password *string `pulumi:"password"`
	// The private IPv4 address associated with the instance.
	PrivateIps []InstancePrivateIp `pulumi:"privateIps"`
	// Private Network endpoints of the Database Instance.
	PrivateNetwork *InstancePrivateNetwork `pulumi:"privateNetwork"`
	// The projectId you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// Public network endpoint configuration (no arguments).
	// > **Important** If neither privateNetwork nor publicNetwork is specified, a public network endpoint is created by default.
	PublicNetwork *InstancePublicNetwork `pulumi:"publicNetwork"`
	// The region you want to attach the resource to
	Region *string `pulumi:"region"`
	// Map of settings to define for the instance.
	Settings map[string]string `pulumi:"settings"`
	// Snapshot ID to restore the MongoDB® instance from.
	SnapshotId *string `pulumi:"snapshotId"`
	// List of tags attached to the MongoDB® instance.
	Tags []string `pulumi:"tags"`
	// Name of the user created when the intance is created.
	UserName *string `pulumi:"userName"`
	// MongoDB® version of the instance.
	Version *string `pulumi:"version"`
	// Volume size in GB.
	VolumeSizeInGb *int `pulumi:"volumeSizeInGb"`
	// Volume type of the instance.
	VolumeType *string `pulumi:"volumeType"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Name of the MongoDB® instance.
	Name pulumi.StringPtrInput
	// Number of nodes in the instance
	NodeNumber pulumi.IntInput
	// The type of MongoDB® intance to create.
	NodeType pulumi.StringInput
	// Password of the user.
	Password pulumi.StringPtrInput
	// The private IPv4 address associated with the instance.
	PrivateIps InstancePrivateIpArrayInput
	// Private Network endpoints of the Database Instance.
	PrivateNetwork InstancePrivateNetworkPtrInput
	// The projectId you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// Public network endpoint configuration (no arguments).
	// > **Important** If neither privateNetwork nor publicNetwork is specified, a public network endpoint is created by default.
	PublicNetwork InstancePublicNetworkPtrInput
	// The region you want to attach the resource to
	Region pulumi.StringPtrInput
	// Map of settings to define for the instance.
	Settings pulumi.StringMapInput
	// Snapshot ID to restore the MongoDB® instance from.
	SnapshotId pulumi.StringPtrInput
	// List of tags attached to the MongoDB® instance.
	Tags pulumi.StringArrayInput
	// Name of the user created when the intance is created.
	UserName pulumi.StringPtrInput
	// MongoDB® version of the instance.
	Version pulumi.StringPtrInput
	// Volume size in GB.
	VolumeSizeInGb pulumi.IntPtrInput
	// Volume type of the instance.
	VolumeType pulumi.StringPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//	InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//	InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

// The date and time of the creation of the MongoDB® instance.
func (o InstanceOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Name of the MongoDB® instance.
func (o InstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Number of nodes in the instance
func (o InstanceOutput) NodeNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.NodeNumber }).(pulumi.IntOutput)
}

// The type of MongoDB® intance to create.
func (o InstanceOutput) NodeType() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.NodeType }).(pulumi.StringOutput)
}

// Password of the user.
func (o InstanceOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The private IPv4 address associated with the instance.
func (o InstanceOutput) PrivateIps() InstancePrivateIpArrayOutput {
	return o.ApplyT(func(v *Instance) InstancePrivateIpArrayOutput { return v.PrivateIps }).(InstancePrivateIpArrayOutput)
}

// Private Network endpoints of the Database Instance.
func (o InstanceOutput) PrivateNetwork() InstancePrivateNetworkPtrOutput {
	return o.ApplyT(func(v *Instance) InstancePrivateNetworkPtrOutput { return v.PrivateNetwork }).(InstancePrivateNetworkPtrOutput)
}

// The projectId you want to attach the resource to
func (o InstanceOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Public network endpoint configuration (no arguments).
// > **Important** If neither privateNetwork nor publicNetwork is specified, a public network endpoint is created by default.
func (o InstanceOutput) PublicNetwork() InstancePublicNetworkOutput {
	return o.ApplyT(func(v *Instance) InstancePublicNetworkOutput { return v.PublicNetwork }).(InstancePublicNetworkOutput)
}

// The region you want to attach the resource to
func (o InstanceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Map of settings to define for the instance.
func (o InstanceOutput) Settings() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringMapOutput { return v.Settings }).(pulumi.StringMapOutput)
}

// Snapshot ID to restore the MongoDB® instance from.
func (o InstanceOutput) SnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.SnapshotId }).(pulumi.StringPtrOutput)
}

// List of tags attached to the MongoDB® instance.
func (o InstanceOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The PEM-encoded TLS certificate for the MongoDB® instance, if available.
func (o InstanceOutput) TlsCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.TlsCertificate }).(pulumi.StringOutput)
}

// The date and time of the last update of the MongoDB® instance.
func (o InstanceOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// Name of the user created when the intance is created.
func (o InstanceOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.UserName }).(pulumi.StringPtrOutput)
}

// MongoDB® version of the instance.
func (o InstanceOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

// Volume size in GB.
func (o InstanceOutput) VolumeSizeInGb() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.VolumeSizeInGb }).(pulumi.IntOutput)
}

// Volume type of the instance.
func (o InstanceOutput) VolumeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.VolumeType }).(pulumi.StringPtrOutput)
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].([]*Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].(map[string]*Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceArrayInput)(nil)).Elem(), InstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMapInput)(nil)).Elem(), InstanceMap{})
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}
