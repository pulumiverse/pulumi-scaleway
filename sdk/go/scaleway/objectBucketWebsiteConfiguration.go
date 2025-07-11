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

// The `object.BucketWebsiteConfiguration` resource allows you to deploy and manage a bucket website with [Scaleway Object storage](https://www.scaleway.com/en/docs/object-storage/).
//
// Refer to the [dedicated documentation](https://www.scaleway.com/en/docs/object-storage/how-to/use-bucket-website/) for more information on bucket websites.
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
//			test, err := object.NewBucket(ctx, "test", &object.BucketArgs{
//				Name: pulumi.String("my-bucket"),
//				Acl:  pulumi.String("public-read"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = object.NewItem(ctx, "some_file", &object.ItemArgs{
//				Bucket:      test.Name,
//				Key:         pulumi.String("index.html"),
//				File:        pulumi.String("index.html"),
//				Visibility:  pulumi.String("public-read"),
//				ContentType: pulumi.String("text/html"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = object.NewBucketWebsiteConfiguration(ctx, "test", &object.BucketWebsiteConfigurationArgs{
//				Bucket: test.Name,
//				IndexDocument: &object.BucketWebsiteConfigurationIndexDocumentArgs{
//					Suffix: pulumi.String("index.html"),
//				},
//				ErrorDocument: &object.BucketWebsiteConfigurationErrorDocumentArgs{
//					Key: pulumi.String("error.html"),
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
// ### With A Bucket Policy
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/object"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := object.NewBucket(ctx, "main", &object.BucketArgs{
//				Name: pulumi.String("MyBucket"),
//				Acl:  pulumi.String("public-read"),
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"Version": "2012-10-17",
//				"Id":      "MyPolicy",
//				"Statement": []map[string]interface{}{
//					map[string]interface{}{
//						"Sid":       "GrantToEveryone",
//						"Effect":    "Allow",
//						"Principal": "*",
//						"Action": []string{
//							"s3:GetObject",
//						},
//						"Resource": []string{
//							"<bucket-name>/*",
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			_, err = object.NewBucketPolicy(ctx, "main", &object.BucketPolicyArgs{
//				Bucket: main.ID(),
//				Policy: pulumi.String(json0),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = object.NewBucketWebsiteConfiguration(ctx, "main", &object.BucketWebsiteConfigurationArgs{
//				Bucket: main.ID(),
//				IndexDocument: &object.BucketWebsiteConfigurationIndexDocumentArgs{
//					Suffix: pulumi.String("index.html"),
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
// Bucket website configurations can be imported using the `{region}/{bucketName}` identifier, as shown below:
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/objectBucketWebsiteConfiguration:ObjectBucketWebsiteConfiguration some_bucket fr-par/some-bucket
// ```
//
// ~> **Important:** The `project_id` attribute has a particular behavior with s3 products because the s3 API is scoped by project.
//
// If you are using a project different from the default one, you have to specify the project ID at the end of the import command.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/objectBucketWebsiteConfiguration:ObjectBucketWebsiteConfiguration some_bucket fr-par/some-bucket@xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxx
// ```
//
// Deprecated: scaleway.index/objectbucketwebsiteconfiguration.ObjectBucketWebsiteConfiguration has been deprecated in favor of scaleway.object/bucketwebsiteconfiguration.BucketWebsiteConfiguration
type ObjectBucketWebsiteConfiguration struct {
	pulumi.CustomResourceState

	// The name of the bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The name of the error file for the website detailed below.
	ErrorDocument ObjectBucketWebsiteConfigurationErrorDocumentPtrOutput `pulumi:"errorDocument"`
	// The name of the index file for the website detailed below.
	IndexDocument ObjectBucketWebsiteConfigurationIndexDocumentOutput `pulumi:"indexDocument"`
	// The projectId you want to attach the resource to
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The region you want to attach the resource to
	Region pulumi.StringOutput `pulumi:"region"`
	// The domain of the website endpoint. This is used to create DNS alias [records](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/).
	WebsiteDomain pulumi.StringOutput `pulumi:"websiteDomain"`
	// The website endpoint.
	WebsiteEndpoint pulumi.StringOutput `pulumi:"websiteEndpoint"`
}

// NewObjectBucketWebsiteConfiguration registers a new resource with the given unique name, arguments, and options.
func NewObjectBucketWebsiteConfiguration(ctx *pulumi.Context,
	name string, args *ObjectBucketWebsiteConfigurationArgs, opts ...pulumi.ResourceOption) (*ObjectBucketWebsiteConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.IndexDocument == nil {
		return nil, errors.New("invalid value for required argument 'IndexDocument'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ObjectBucketWebsiteConfiguration
	err := ctx.RegisterResource("scaleway:index/objectBucketWebsiteConfiguration:ObjectBucketWebsiteConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetObjectBucketWebsiteConfiguration gets an existing ObjectBucketWebsiteConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetObjectBucketWebsiteConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ObjectBucketWebsiteConfigurationState, opts ...pulumi.ResourceOption) (*ObjectBucketWebsiteConfiguration, error) {
	var resource ObjectBucketWebsiteConfiguration
	err := ctx.ReadResource("scaleway:index/objectBucketWebsiteConfiguration:ObjectBucketWebsiteConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ObjectBucketWebsiteConfiguration resources.
type objectBucketWebsiteConfigurationState struct {
	// The name of the bucket.
	Bucket *string `pulumi:"bucket"`
	// The name of the error file for the website detailed below.
	ErrorDocument *ObjectBucketWebsiteConfigurationErrorDocument `pulumi:"errorDocument"`
	// The name of the index file for the website detailed below.
	IndexDocument *ObjectBucketWebsiteConfigurationIndexDocument `pulumi:"indexDocument"`
	// The projectId you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// The region you want to attach the resource to
	Region *string `pulumi:"region"`
	// The domain of the website endpoint. This is used to create DNS alias [records](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/).
	WebsiteDomain *string `pulumi:"websiteDomain"`
	// The website endpoint.
	WebsiteEndpoint *string `pulumi:"websiteEndpoint"`
}

type ObjectBucketWebsiteConfigurationState struct {
	// The name of the bucket.
	Bucket pulumi.StringPtrInput
	// The name of the error file for the website detailed below.
	ErrorDocument ObjectBucketWebsiteConfigurationErrorDocumentPtrInput
	// The name of the index file for the website detailed below.
	IndexDocument ObjectBucketWebsiteConfigurationIndexDocumentPtrInput
	// The projectId you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// The region you want to attach the resource to
	Region pulumi.StringPtrInput
	// The domain of the website endpoint. This is used to create DNS alias [records](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/).
	WebsiteDomain pulumi.StringPtrInput
	// The website endpoint.
	WebsiteEndpoint pulumi.StringPtrInput
}

func (ObjectBucketWebsiteConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*objectBucketWebsiteConfigurationState)(nil)).Elem()
}

type objectBucketWebsiteConfigurationArgs struct {
	// The name of the bucket.
	Bucket string `pulumi:"bucket"`
	// The name of the error file for the website detailed below.
	ErrorDocument *ObjectBucketWebsiteConfigurationErrorDocument `pulumi:"errorDocument"`
	// The name of the index file for the website detailed below.
	IndexDocument ObjectBucketWebsiteConfigurationIndexDocument `pulumi:"indexDocument"`
	// The projectId you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// The region you want to attach the resource to
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a ObjectBucketWebsiteConfiguration resource.
type ObjectBucketWebsiteConfigurationArgs struct {
	// The name of the bucket.
	Bucket pulumi.StringInput
	// The name of the error file for the website detailed below.
	ErrorDocument ObjectBucketWebsiteConfigurationErrorDocumentPtrInput
	// The name of the index file for the website detailed below.
	IndexDocument ObjectBucketWebsiteConfigurationIndexDocumentInput
	// The projectId you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// The region you want to attach the resource to
	Region pulumi.StringPtrInput
}

func (ObjectBucketWebsiteConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*objectBucketWebsiteConfigurationArgs)(nil)).Elem()
}

type ObjectBucketWebsiteConfigurationInput interface {
	pulumi.Input

	ToObjectBucketWebsiteConfigurationOutput() ObjectBucketWebsiteConfigurationOutput
	ToObjectBucketWebsiteConfigurationOutputWithContext(ctx context.Context) ObjectBucketWebsiteConfigurationOutput
}

func (*ObjectBucketWebsiteConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectBucketWebsiteConfiguration)(nil)).Elem()
}

func (i *ObjectBucketWebsiteConfiguration) ToObjectBucketWebsiteConfigurationOutput() ObjectBucketWebsiteConfigurationOutput {
	return i.ToObjectBucketWebsiteConfigurationOutputWithContext(context.Background())
}

func (i *ObjectBucketWebsiteConfiguration) ToObjectBucketWebsiteConfigurationOutputWithContext(ctx context.Context) ObjectBucketWebsiteConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectBucketWebsiteConfigurationOutput)
}

// ObjectBucketWebsiteConfigurationArrayInput is an input type that accepts ObjectBucketWebsiteConfigurationArray and ObjectBucketWebsiteConfigurationArrayOutput values.
// You can construct a concrete instance of `ObjectBucketWebsiteConfigurationArrayInput` via:
//
//	ObjectBucketWebsiteConfigurationArray{ ObjectBucketWebsiteConfigurationArgs{...} }
type ObjectBucketWebsiteConfigurationArrayInput interface {
	pulumi.Input

	ToObjectBucketWebsiteConfigurationArrayOutput() ObjectBucketWebsiteConfigurationArrayOutput
	ToObjectBucketWebsiteConfigurationArrayOutputWithContext(context.Context) ObjectBucketWebsiteConfigurationArrayOutput
}

type ObjectBucketWebsiteConfigurationArray []ObjectBucketWebsiteConfigurationInput

func (ObjectBucketWebsiteConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectBucketWebsiteConfiguration)(nil)).Elem()
}

func (i ObjectBucketWebsiteConfigurationArray) ToObjectBucketWebsiteConfigurationArrayOutput() ObjectBucketWebsiteConfigurationArrayOutput {
	return i.ToObjectBucketWebsiteConfigurationArrayOutputWithContext(context.Background())
}

func (i ObjectBucketWebsiteConfigurationArray) ToObjectBucketWebsiteConfigurationArrayOutputWithContext(ctx context.Context) ObjectBucketWebsiteConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectBucketWebsiteConfigurationArrayOutput)
}

// ObjectBucketWebsiteConfigurationMapInput is an input type that accepts ObjectBucketWebsiteConfigurationMap and ObjectBucketWebsiteConfigurationMapOutput values.
// You can construct a concrete instance of `ObjectBucketWebsiteConfigurationMapInput` via:
//
//	ObjectBucketWebsiteConfigurationMap{ "key": ObjectBucketWebsiteConfigurationArgs{...} }
type ObjectBucketWebsiteConfigurationMapInput interface {
	pulumi.Input

	ToObjectBucketWebsiteConfigurationMapOutput() ObjectBucketWebsiteConfigurationMapOutput
	ToObjectBucketWebsiteConfigurationMapOutputWithContext(context.Context) ObjectBucketWebsiteConfigurationMapOutput
}

type ObjectBucketWebsiteConfigurationMap map[string]ObjectBucketWebsiteConfigurationInput

func (ObjectBucketWebsiteConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectBucketWebsiteConfiguration)(nil)).Elem()
}

func (i ObjectBucketWebsiteConfigurationMap) ToObjectBucketWebsiteConfigurationMapOutput() ObjectBucketWebsiteConfigurationMapOutput {
	return i.ToObjectBucketWebsiteConfigurationMapOutputWithContext(context.Background())
}

func (i ObjectBucketWebsiteConfigurationMap) ToObjectBucketWebsiteConfigurationMapOutputWithContext(ctx context.Context) ObjectBucketWebsiteConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectBucketWebsiteConfigurationMapOutput)
}

type ObjectBucketWebsiteConfigurationOutput struct{ *pulumi.OutputState }

func (ObjectBucketWebsiteConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectBucketWebsiteConfiguration)(nil)).Elem()
}

func (o ObjectBucketWebsiteConfigurationOutput) ToObjectBucketWebsiteConfigurationOutput() ObjectBucketWebsiteConfigurationOutput {
	return o
}

func (o ObjectBucketWebsiteConfigurationOutput) ToObjectBucketWebsiteConfigurationOutputWithContext(ctx context.Context) ObjectBucketWebsiteConfigurationOutput {
	return o
}

// The name of the bucket.
func (o ObjectBucketWebsiteConfigurationOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectBucketWebsiteConfiguration) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// The name of the error file for the website detailed below.
func (o ObjectBucketWebsiteConfigurationOutput) ErrorDocument() ObjectBucketWebsiteConfigurationErrorDocumentPtrOutput {
	return o.ApplyT(func(v *ObjectBucketWebsiteConfiguration) ObjectBucketWebsiteConfigurationErrorDocumentPtrOutput {
		return v.ErrorDocument
	}).(ObjectBucketWebsiteConfigurationErrorDocumentPtrOutput)
}

// The name of the index file for the website detailed below.
func (o ObjectBucketWebsiteConfigurationOutput) IndexDocument() ObjectBucketWebsiteConfigurationIndexDocumentOutput {
	return o.ApplyT(func(v *ObjectBucketWebsiteConfiguration) ObjectBucketWebsiteConfigurationIndexDocumentOutput {
		return v.IndexDocument
	}).(ObjectBucketWebsiteConfigurationIndexDocumentOutput)
}

// The projectId you want to attach the resource to
func (o ObjectBucketWebsiteConfigurationOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectBucketWebsiteConfiguration) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The region you want to attach the resource to
func (o ObjectBucketWebsiteConfigurationOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectBucketWebsiteConfiguration) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The domain of the website endpoint. This is used to create DNS alias [records](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/).
func (o ObjectBucketWebsiteConfigurationOutput) WebsiteDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectBucketWebsiteConfiguration) pulumi.StringOutput { return v.WebsiteDomain }).(pulumi.StringOutput)
}

// The website endpoint.
func (o ObjectBucketWebsiteConfigurationOutput) WebsiteEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectBucketWebsiteConfiguration) pulumi.StringOutput { return v.WebsiteEndpoint }).(pulumi.StringOutput)
}

type ObjectBucketWebsiteConfigurationArrayOutput struct{ *pulumi.OutputState }

func (ObjectBucketWebsiteConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectBucketWebsiteConfiguration)(nil)).Elem()
}

func (o ObjectBucketWebsiteConfigurationArrayOutput) ToObjectBucketWebsiteConfigurationArrayOutput() ObjectBucketWebsiteConfigurationArrayOutput {
	return o
}

func (o ObjectBucketWebsiteConfigurationArrayOutput) ToObjectBucketWebsiteConfigurationArrayOutputWithContext(ctx context.Context) ObjectBucketWebsiteConfigurationArrayOutput {
	return o
}

func (o ObjectBucketWebsiteConfigurationArrayOutput) Index(i pulumi.IntInput) ObjectBucketWebsiteConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ObjectBucketWebsiteConfiguration {
		return vs[0].([]*ObjectBucketWebsiteConfiguration)[vs[1].(int)]
	}).(ObjectBucketWebsiteConfigurationOutput)
}

type ObjectBucketWebsiteConfigurationMapOutput struct{ *pulumi.OutputState }

func (ObjectBucketWebsiteConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectBucketWebsiteConfiguration)(nil)).Elem()
}

func (o ObjectBucketWebsiteConfigurationMapOutput) ToObjectBucketWebsiteConfigurationMapOutput() ObjectBucketWebsiteConfigurationMapOutput {
	return o
}

func (o ObjectBucketWebsiteConfigurationMapOutput) ToObjectBucketWebsiteConfigurationMapOutputWithContext(ctx context.Context) ObjectBucketWebsiteConfigurationMapOutput {
	return o
}

func (o ObjectBucketWebsiteConfigurationMapOutput) MapIndex(k pulumi.StringInput) ObjectBucketWebsiteConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ObjectBucketWebsiteConfiguration {
		return vs[0].(map[string]*ObjectBucketWebsiteConfiguration)[vs[1].(string)]
	}).(ObjectBucketWebsiteConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectBucketWebsiteConfigurationInput)(nil)).Elem(), &ObjectBucketWebsiteConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectBucketWebsiteConfigurationArrayInput)(nil)).Elem(), ObjectBucketWebsiteConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectBucketWebsiteConfigurationMapInput)(nil)).Elem(), ObjectBucketWebsiteConfigurationMap{})
	pulumi.RegisterOutputType(ObjectBucketWebsiteConfigurationOutput{})
	pulumi.RegisterOutputType(ObjectBucketWebsiteConfigurationArrayOutput{})
	pulumi.RegisterOutputType(ObjectBucketWebsiteConfigurationMapOutput{})
}
