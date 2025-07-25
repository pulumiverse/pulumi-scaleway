// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package object

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// The `object.Bucket` resource allows you to create and manage buckets for [Scaleway Object storage](https://www.scaleway.com/en/docs/object-storage/).
//
// Refer to the [dedicated documentation](https://www.scaleway.com/en/docs/object-storage/how-to/create-a-bucket/) for more information on Object Storage buckets.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/object"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := object.NewBucket(ctx, "some_bucket", &object.BucketArgs{
//				Name: pulumi.String("some-unique-name"),
//				Tags: pulumi.StringMap{
//					"key": pulumi.String("value"),
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
// ### Creating the bucket in a specific project
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/object"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := object.NewBucket(ctx, "some_bucket", &object.BucketArgs{
//				Name:      pulumi.String("some-unique-name"),
//				ProjectId: pulumi.String("11111111-1111-1111-1111-111111111111"),
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
// ### Using object lifecycle
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/object"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := object.NewBucket(ctx, "main", &object.BucketArgs{
//				Name:   pulumi.String("mybuckectid"),
//				Region: pulumi.String("fr-par"),
//				LifecycleRules: object.BucketLifecycleRuleArray{
//					&object.BucketLifecycleRuleArgs{
//						Id:      pulumi.String("id1"),
//						Prefix:  pulumi.String("path1/"),
//						Enabled: pulumi.Bool(true),
//						Expiration: &object.BucketLifecycleRuleExpirationArgs{
//							Days: pulumi.Int(365),
//						},
//						Transitions: object.BucketLifecycleRuleTransitionArray{
//							&object.BucketLifecycleRuleTransitionArgs{
//								Days:         pulumi.Int(120),
//								StorageClass: pulumi.String("GLACIER"),
//							},
//						},
//					},
//					&object.BucketLifecycleRuleArgs{
//						Id:      pulumi.String("id2"),
//						Prefix:  pulumi.String("path2/"),
//						Enabled: pulumi.Bool(true),
//						Expiration: &object.BucketLifecycleRuleExpirationArgs{
//							Days: pulumi.Int(50),
//						},
//					},
//					&object.BucketLifecycleRuleArgs{
//						Id:      pulumi.String("id3"),
//						Prefix:  pulumi.String("path3/"),
//						Enabled: pulumi.Bool(false),
//						Tags: pulumi.StringMap{
//							"tagKey":    pulumi.String("tagValue"),
//							"terraform": pulumi.String("hashicorp"),
//						},
//						Expiration: &object.BucketLifecycleRuleExpirationArgs{
//							Days: pulumi.Int(1),
//						},
//					},
//					&object.BucketLifecycleRuleArgs{
//						Id:      pulumi.String("id4"),
//						Enabled: pulumi.Bool(true),
//						Tags: pulumi.StringMap{
//							"tag1": pulumi.String("value1"),
//						},
//						Transitions: object.BucketLifecycleRuleTransitionArray{
//							&object.BucketLifecycleRuleTransitionArgs{
//								Days:         pulumi.Int(1),
//								StorageClass: pulumi.String("GLACIER"),
//							},
//						},
//					},
//					&object.BucketLifecycleRuleArgs{
//						Enabled:                            pulumi.Bool(true),
//						AbortIncompleteMultipartUploadDays: pulumi.Int(30),
//					},
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
// Buckets can be imported using the `{region}/{bucketName}` identifier, as shown below:
//
// bash
//
// ```sh
// $ pulumi import scaleway:object/bucket:Bucket some_bucket fr-par/some-bucket
// ```
//
// ~> **Important:** The `project_id` attribute has a particular behavior with s3 products because the s3 API is scoped by project.
//
// If you are using a project different from the default one, you have to specify the project ID at the end of the import command.
//
// bash
//
// ```sh
// $ pulumi import scaleway:object/bucket:Bucket some_bucket fr-par/some-bucket@11111111-1111-1111-1111-111111111111
// ```
type Bucket struct {
	pulumi.CustomResourceState

	// (Deprecated) The canned ACL you want to apply to the bucket.
	//
	// > **Note:** The `acl` attribute is deprecated. See object.BucketAcl resource documentation. Refer to the [official canned ACL documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) for more information on the different roles.
	//
	// Deprecated: ACL attribute is deprecated. Please use the resource object.BucketAcl instead.
	Acl pulumi.StringPtrOutput `pulumi:"acl"`
	// API URL of the bucket
	ApiEndpoint pulumi.StringOutput       `pulumi:"apiEndpoint"`
	CorsRules   BucketCorsRuleArrayOutput `pulumi:"corsRules"`
	// The endpoint URL of the bucket.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// Boolean that, when set to true, allows the deletion of all objects (including locked objects) when the bucket is destroyed. This operation is irreversible, and the objects cannot be recovered. The default is false.
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
	LifecycleRules BucketLifecycleRuleArrayOutput `pulumi:"lifecycleRules"`
	// The name of the bucket.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable object lock
	ObjectLockEnabled pulumi.BoolPtrOutput `pulumi:"objectLockEnabled"`
	// `projectId`) The ID of the project the bucket is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The [region](https://www.scaleway.com/en/developers/api/#region-definition) in which the bucket will be created.
	Region pulumi.StringOutput `pulumi:"region"`
	// A list of tags (key/value) for the bucket.
	//
	// * > **Important:** The Scaleway console does not support `key/value` tags yet, so only the tags' values will be displayed.
	// If you make any change to your bucket's tags using the console, it will overwrite them with the format `value/value`.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Allow multiple versions of an object in the same bucket
	Versioning BucketVersioningOutput `pulumi:"versioning"`
}

// NewBucket registers a new resource with the given unique name, arguments, and options.
func NewBucket(ctx *pulumi.Context,
	name string, args *BucketArgs, opts ...pulumi.ResourceOption) (*Bucket, error) {
	if args == nil {
		args = &BucketArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("scaleway:index/objectBucket:ObjectBucket"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Bucket
	err := ctx.RegisterResource("scaleway:object/bucket:Bucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucket gets an existing Bucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketState, opts ...pulumi.ResourceOption) (*Bucket, error) {
	var resource Bucket
	err := ctx.ReadResource("scaleway:object/bucket:Bucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bucket resources.
type bucketState struct {
	// (Deprecated) The canned ACL you want to apply to the bucket.
	//
	// > **Note:** The `acl` attribute is deprecated. See object.BucketAcl resource documentation. Refer to the [official canned ACL documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) for more information on the different roles.
	//
	// Deprecated: ACL attribute is deprecated. Please use the resource object.BucketAcl instead.
	Acl *string `pulumi:"acl"`
	// API URL of the bucket
	ApiEndpoint *string          `pulumi:"apiEndpoint"`
	CorsRules   []BucketCorsRule `pulumi:"corsRules"`
	// The endpoint URL of the bucket.
	Endpoint *string `pulumi:"endpoint"`
	// Boolean that, when set to true, allows the deletion of all objects (including locked objects) when the bucket is destroyed. This operation is irreversible, and the objects cannot be recovered. The default is false.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
	LifecycleRules []BucketLifecycleRule `pulumi:"lifecycleRules"`
	// The name of the bucket.
	Name *string `pulumi:"name"`
	// Enable object lock
	ObjectLockEnabled *bool `pulumi:"objectLockEnabled"`
	// `projectId`) The ID of the project the bucket is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The [region](https://www.scaleway.com/en/developers/api/#region-definition) in which the bucket will be created.
	Region *string `pulumi:"region"`
	// A list of tags (key/value) for the bucket.
	//
	// * > **Important:** The Scaleway console does not support `key/value` tags yet, so only the tags' values will be displayed.
	// If you make any change to your bucket's tags using the console, it will overwrite them with the format `value/value`.
	Tags map[string]string `pulumi:"tags"`
	// Allow multiple versions of an object in the same bucket
	Versioning *BucketVersioning `pulumi:"versioning"`
}

type BucketState struct {
	// (Deprecated) The canned ACL you want to apply to the bucket.
	//
	// > **Note:** The `acl` attribute is deprecated. See object.BucketAcl resource documentation. Refer to the [official canned ACL documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) for more information on the different roles.
	//
	// Deprecated: ACL attribute is deprecated. Please use the resource object.BucketAcl instead.
	Acl pulumi.StringPtrInput
	// API URL of the bucket
	ApiEndpoint pulumi.StringPtrInput
	CorsRules   BucketCorsRuleArrayInput
	// The endpoint URL of the bucket.
	Endpoint pulumi.StringPtrInput
	// Boolean that, when set to true, allows the deletion of all objects (including locked objects) when the bucket is destroyed. This operation is irreversible, and the objects cannot be recovered. The default is false.
	ForceDestroy pulumi.BoolPtrInput
	// Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
	LifecycleRules BucketLifecycleRuleArrayInput
	// The name of the bucket.
	Name pulumi.StringPtrInput
	// Enable object lock
	ObjectLockEnabled pulumi.BoolPtrInput
	// `projectId`) The ID of the project the bucket is associated with.
	ProjectId pulumi.StringPtrInput
	// The [region](https://www.scaleway.com/en/developers/api/#region-definition) in which the bucket will be created.
	Region pulumi.StringPtrInput
	// A list of tags (key/value) for the bucket.
	//
	// * > **Important:** The Scaleway console does not support `key/value` tags yet, so only the tags' values will be displayed.
	// If you make any change to your bucket's tags using the console, it will overwrite them with the format `value/value`.
	Tags pulumi.StringMapInput
	// Allow multiple versions of an object in the same bucket
	Versioning BucketVersioningPtrInput
}

func (BucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketState)(nil)).Elem()
}

type bucketArgs struct {
	// (Deprecated) The canned ACL you want to apply to the bucket.
	//
	// > **Note:** The `acl` attribute is deprecated. See object.BucketAcl resource documentation. Refer to the [official canned ACL documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) for more information on the different roles.
	//
	// Deprecated: ACL attribute is deprecated. Please use the resource object.BucketAcl instead.
	Acl       *string          `pulumi:"acl"`
	CorsRules []BucketCorsRule `pulumi:"corsRules"`
	// Boolean that, when set to true, allows the deletion of all objects (including locked objects) when the bucket is destroyed. This operation is irreversible, and the objects cannot be recovered. The default is false.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
	LifecycleRules []BucketLifecycleRule `pulumi:"lifecycleRules"`
	// The name of the bucket.
	Name *string `pulumi:"name"`
	// Enable object lock
	ObjectLockEnabled *bool `pulumi:"objectLockEnabled"`
	// `projectId`) The ID of the project the bucket is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The [region](https://www.scaleway.com/en/developers/api/#region-definition) in which the bucket will be created.
	Region *string `pulumi:"region"`
	// A list of tags (key/value) for the bucket.
	//
	// * > **Important:** The Scaleway console does not support `key/value` tags yet, so only the tags' values will be displayed.
	// If you make any change to your bucket's tags using the console, it will overwrite them with the format `value/value`.
	Tags map[string]string `pulumi:"tags"`
	// Allow multiple versions of an object in the same bucket
	Versioning *BucketVersioning `pulumi:"versioning"`
}

// The set of arguments for constructing a Bucket resource.
type BucketArgs struct {
	// (Deprecated) The canned ACL you want to apply to the bucket.
	//
	// > **Note:** The `acl` attribute is deprecated. See object.BucketAcl resource documentation. Refer to the [official canned ACL documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) for more information on the different roles.
	//
	// Deprecated: ACL attribute is deprecated. Please use the resource object.BucketAcl instead.
	Acl       pulumi.StringPtrInput
	CorsRules BucketCorsRuleArrayInput
	// Boolean that, when set to true, allows the deletion of all objects (including locked objects) when the bucket is destroyed. This operation is irreversible, and the objects cannot be recovered. The default is false.
	ForceDestroy pulumi.BoolPtrInput
	// Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
	LifecycleRules BucketLifecycleRuleArrayInput
	// The name of the bucket.
	Name pulumi.StringPtrInput
	// Enable object lock
	ObjectLockEnabled pulumi.BoolPtrInput
	// `projectId`) The ID of the project the bucket is associated with.
	ProjectId pulumi.StringPtrInput
	// The [region](https://www.scaleway.com/en/developers/api/#region-definition) in which the bucket will be created.
	Region pulumi.StringPtrInput
	// A list of tags (key/value) for the bucket.
	//
	// * > **Important:** The Scaleway console does not support `key/value` tags yet, so only the tags' values will be displayed.
	// If you make any change to your bucket's tags using the console, it will overwrite them with the format `value/value`.
	Tags pulumi.StringMapInput
	// Allow multiple versions of an object in the same bucket
	Versioning BucketVersioningPtrInput
}

func (BucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketArgs)(nil)).Elem()
}

type BucketInput interface {
	pulumi.Input

	ToBucketOutput() BucketOutput
	ToBucketOutputWithContext(ctx context.Context) BucketOutput
}

func (*Bucket) ElementType() reflect.Type {
	return reflect.TypeOf((**Bucket)(nil)).Elem()
}

func (i *Bucket) ToBucketOutput() BucketOutput {
	return i.ToBucketOutputWithContext(context.Background())
}

func (i *Bucket) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketOutput)
}

// BucketArrayInput is an input type that accepts BucketArray and BucketArrayOutput values.
// You can construct a concrete instance of `BucketArrayInput` via:
//
//	BucketArray{ BucketArgs{...} }
type BucketArrayInput interface {
	pulumi.Input

	ToBucketArrayOutput() BucketArrayOutput
	ToBucketArrayOutputWithContext(context.Context) BucketArrayOutput
}

type BucketArray []BucketInput

func (BucketArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bucket)(nil)).Elem()
}

func (i BucketArray) ToBucketArrayOutput() BucketArrayOutput {
	return i.ToBucketArrayOutputWithContext(context.Background())
}

func (i BucketArray) ToBucketArrayOutputWithContext(ctx context.Context) BucketArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketArrayOutput)
}

// BucketMapInput is an input type that accepts BucketMap and BucketMapOutput values.
// You can construct a concrete instance of `BucketMapInput` via:
//
//	BucketMap{ "key": BucketArgs{...} }
type BucketMapInput interface {
	pulumi.Input

	ToBucketMapOutput() BucketMapOutput
	ToBucketMapOutputWithContext(context.Context) BucketMapOutput
}

type BucketMap map[string]BucketInput

func (BucketMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bucket)(nil)).Elem()
}

func (i BucketMap) ToBucketMapOutput() BucketMapOutput {
	return i.ToBucketMapOutputWithContext(context.Background())
}

func (i BucketMap) ToBucketMapOutputWithContext(ctx context.Context) BucketMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketMapOutput)
}

type BucketOutput struct{ *pulumi.OutputState }

func (BucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Bucket)(nil)).Elem()
}

func (o BucketOutput) ToBucketOutput() BucketOutput {
	return o
}

func (o BucketOutput) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return o
}

// (Deprecated) The canned ACL you want to apply to the bucket.
//
// > **Note:** The `acl` attribute is deprecated. See object.BucketAcl resource documentation. Refer to the [official canned ACL documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) for more information on the different roles.
//
// Deprecated: ACL attribute is deprecated. Please use the resource object.BucketAcl instead.
func (o BucketOutput) Acl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringPtrOutput { return v.Acl }).(pulumi.StringPtrOutput)
}

// API URL of the bucket
func (o BucketOutput) ApiEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.ApiEndpoint }).(pulumi.StringOutput)
}

func (o BucketOutput) CorsRules() BucketCorsRuleArrayOutput {
	return o.ApplyT(func(v *Bucket) BucketCorsRuleArrayOutput { return v.CorsRules }).(BucketCorsRuleArrayOutput)
}

// The endpoint URL of the bucket.
func (o BucketOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// Boolean that, when set to true, allows the deletion of all objects (including locked objects) when the bucket is destroyed. This operation is irreversible, and the objects cannot be recovered. The default is false.
func (o BucketOutput) ForceDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.BoolPtrOutput { return v.ForceDestroy }).(pulumi.BoolPtrOutput)
}

// Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
func (o BucketOutput) LifecycleRules() BucketLifecycleRuleArrayOutput {
	return o.ApplyT(func(v *Bucket) BucketLifecycleRuleArrayOutput { return v.LifecycleRules }).(BucketLifecycleRuleArrayOutput)
}

// The name of the bucket.
func (o BucketOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable object lock
func (o BucketOutput) ObjectLockEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.BoolPtrOutput { return v.ObjectLockEnabled }).(pulumi.BoolPtrOutput)
}

// `projectId`) The ID of the project the bucket is associated with.
func (o BucketOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The [region](https://www.scaleway.com/en/developers/api/#region-definition) in which the bucket will be created.
func (o BucketOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// A list of tags (key/value) for the bucket.
//
// * > **Important:** The Scaleway console does not support `key/value` tags yet, so only the tags' values will be displayed.
// If you make any change to your bucket's tags using the console, it will overwrite them with the format `value/value`.
func (o BucketOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Allow multiple versions of an object in the same bucket
func (o BucketOutput) Versioning() BucketVersioningOutput {
	return o.ApplyT(func(v *Bucket) BucketVersioningOutput { return v.Versioning }).(BucketVersioningOutput)
}

type BucketArrayOutput struct{ *pulumi.OutputState }

func (BucketArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bucket)(nil)).Elem()
}

func (o BucketArrayOutput) ToBucketArrayOutput() BucketArrayOutput {
	return o
}

func (o BucketArrayOutput) ToBucketArrayOutputWithContext(ctx context.Context) BucketArrayOutput {
	return o
}

func (o BucketArrayOutput) Index(i pulumi.IntInput) BucketOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Bucket {
		return vs[0].([]*Bucket)[vs[1].(int)]
	}).(BucketOutput)
}

type BucketMapOutput struct{ *pulumi.OutputState }

func (BucketMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bucket)(nil)).Elem()
}

func (o BucketMapOutput) ToBucketMapOutput() BucketMapOutput {
	return o
}

func (o BucketMapOutput) ToBucketMapOutputWithContext(ctx context.Context) BucketMapOutput {
	return o
}

func (o BucketMapOutput) MapIndex(k pulumi.StringInput) BucketOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Bucket {
		return vs[0].(map[string]*Bucket)[vs[1].(string)]
	}).(BucketOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketInput)(nil)).Elem(), &Bucket{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketArrayInput)(nil)).Elem(), BucketArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketMapInput)(nil)).Elem(), BucketMap{})
	pulumi.RegisterOutputType(BucketOutput{})
	pulumi.RegisterOutputType(BucketArrayOutput{})
	pulumi.RegisterOutputType(BucketMapOutput{})
}
