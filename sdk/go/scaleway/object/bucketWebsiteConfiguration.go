// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package object

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
// $ pulumi import scaleway:object/bucketWebsiteConfiguration:BucketWebsiteConfiguration some_bucket fr-par/some-bucket
// ```
//
// ~> **Important:** The `project_id` attribute has a particular behavior with s3 products because the s3 API is scoped by project.
//
// If you are using a project different from the default one, you have to specify the project ID at the end of the import command.
//
// bash
//
// ```sh
// $ pulumi import scaleway:object/bucketWebsiteConfiguration:BucketWebsiteConfiguration some_bucket fr-par/some-bucket@xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxx
// ```
type BucketWebsiteConfiguration struct {
	pulumi.CustomResourceState

	// The name of the bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The name of the error file for the website detailed below.
	ErrorDocument BucketWebsiteConfigurationErrorDocumentPtrOutput `pulumi:"errorDocument"`
	// The name of the index file for the website detailed below.
	IndexDocument BucketWebsiteConfigurationIndexDocumentOutput `pulumi:"indexDocument"`
	// The projectId you want to attach the resource to
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The region you want to attach the resource to
	Region pulumi.StringOutput `pulumi:"region"`
	// The domain of the website endpoint. This is used to create DNS alias [records](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/).
	WebsiteDomain pulumi.StringOutput `pulumi:"websiteDomain"`
	// The website endpoint.
	WebsiteEndpoint pulumi.StringOutput `pulumi:"websiteEndpoint"`
}

// NewBucketWebsiteConfiguration registers a new resource with the given unique name, arguments, and options.
func NewBucketWebsiteConfiguration(ctx *pulumi.Context,
	name string, args *BucketWebsiteConfigurationArgs, opts ...pulumi.ResourceOption) (*BucketWebsiteConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.IndexDocument == nil {
		return nil, errors.New("invalid value for required argument 'IndexDocument'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("scaleway:index/objectBucketWebsiteConfiguration:ObjectBucketWebsiteConfiguration"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BucketWebsiteConfiguration
	err := ctx.RegisterResource("scaleway:object/bucketWebsiteConfiguration:BucketWebsiteConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketWebsiteConfiguration gets an existing BucketWebsiteConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketWebsiteConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketWebsiteConfigurationState, opts ...pulumi.ResourceOption) (*BucketWebsiteConfiguration, error) {
	var resource BucketWebsiteConfiguration
	err := ctx.ReadResource("scaleway:object/bucketWebsiteConfiguration:BucketWebsiteConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketWebsiteConfiguration resources.
type bucketWebsiteConfigurationState struct {
	// The name of the bucket.
	Bucket *string `pulumi:"bucket"`
	// The name of the error file for the website detailed below.
	ErrorDocument *BucketWebsiteConfigurationErrorDocument `pulumi:"errorDocument"`
	// The name of the index file for the website detailed below.
	IndexDocument *BucketWebsiteConfigurationIndexDocument `pulumi:"indexDocument"`
	// The projectId you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// The region you want to attach the resource to
	Region *string `pulumi:"region"`
	// The domain of the website endpoint. This is used to create DNS alias [records](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/).
	WebsiteDomain *string `pulumi:"websiteDomain"`
	// The website endpoint.
	WebsiteEndpoint *string `pulumi:"websiteEndpoint"`
}

type BucketWebsiteConfigurationState struct {
	// The name of the bucket.
	Bucket pulumi.StringPtrInput
	// The name of the error file for the website detailed below.
	ErrorDocument BucketWebsiteConfigurationErrorDocumentPtrInput
	// The name of the index file for the website detailed below.
	IndexDocument BucketWebsiteConfigurationIndexDocumentPtrInput
	// The projectId you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// The region you want to attach the resource to
	Region pulumi.StringPtrInput
	// The domain of the website endpoint. This is used to create DNS alias [records](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/).
	WebsiteDomain pulumi.StringPtrInput
	// The website endpoint.
	WebsiteEndpoint pulumi.StringPtrInput
}

func (BucketWebsiteConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketWebsiteConfigurationState)(nil)).Elem()
}

type bucketWebsiteConfigurationArgs struct {
	// The name of the bucket.
	Bucket string `pulumi:"bucket"`
	// The name of the error file for the website detailed below.
	ErrorDocument *BucketWebsiteConfigurationErrorDocument `pulumi:"errorDocument"`
	// The name of the index file for the website detailed below.
	IndexDocument BucketWebsiteConfigurationIndexDocument `pulumi:"indexDocument"`
	// The projectId you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// The region you want to attach the resource to
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a BucketWebsiteConfiguration resource.
type BucketWebsiteConfigurationArgs struct {
	// The name of the bucket.
	Bucket pulumi.StringInput
	// The name of the error file for the website detailed below.
	ErrorDocument BucketWebsiteConfigurationErrorDocumentPtrInput
	// The name of the index file for the website detailed below.
	IndexDocument BucketWebsiteConfigurationIndexDocumentInput
	// The projectId you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// The region you want to attach the resource to
	Region pulumi.StringPtrInput
}

func (BucketWebsiteConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketWebsiteConfigurationArgs)(nil)).Elem()
}

type BucketWebsiteConfigurationInput interface {
	pulumi.Input

	ToBucketWebsiteConfigurationOutput() BucketWebsiteConfigurationOutput
	ToBucketWebsiteConfigurationOutputWithContext(ctx context.Context) BucketWebsiteConfigurationOutput
}

func (*BucketWebsiteConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketWebsiteConfiguration)(nil)).Elem()
}

func (i *BucketWebsiteConfiguration) ToBucketWebsiteConfigurationOutput() BucketWebsiteConfigurationOutput {
	return i.ToBucketWebsiteConfigurationOutputWithContext(context.Background())
}

func (i *BucketWebsiteConfiguration) ToBucketWebsiteConfigurationOutputWithContext(ctx context.Context) BucketWebsiteConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketWebsiteConfigurationOutput)
}

// BucketWebsiteConfigurationArrayInput is an input type that accepts BucketWebsiteConfigurationArray and BucketWebsiteConfigurationArrayOutput values.
// You can construct a concrete instance of `BucketWebsiteConfigurationArrayInput` via:
//
//	BucketWebsiteConfigurationArray{ BucketWebsiteConfigurationArgs{...} }
type BucketWebsiteConfigurationArrayInput interface {
	pulumi.Input

	ToBucketWebsiteConfigurationArrayOutput() BucketWebsiteConfigurationArrayOutput
	ToBucketWebsiteConfigurationArrayOutputWithContext(context.Context) BucketWebsiteConfigurationArrayOutput
}

type BucketWebsiteConfigurationArray []BucketWebsiteConfigurationInput

func (BucketWebsiteConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketWebsiteConfiguration)(nil)).Elem()
}

func (i BucketWebsiteConfigurationArray) ToBucketWebsiteConfigurationArrayOutput() BucketWebsiteConfigurationArrayOutput {
	return i.ToBucketWebsiteConfigurationArrayOutputWithContext(context.Background())
}

func (i BucketWebsiteConfigurationArray) ToBucketWebsiteConfigurationArrayOutputWithContext(ctx context.Context) BucketWebsiteConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketWebsiteConfigurationArrayOutput)
}

// BucketWebsiteConfigurationMapInput is an input type that accepts BucketWebsiteConfigurationMap and BucketWebsiteConfigurationMapOutput values.
// You can construct a concrete instance of `BucketWebsiteConfigurationMapInput` via:
//
//	BucketWebsiteConfigurationMap{ "key": BucketWebsiteConfigurationArgs{...} }
type BucketWebsiteConfigurationMapInput interface {
	pulumi.Input

	ToBucketWebsiteConfigurationMapOutput() BucketWebsiteConfigurationMapOutput
	ToBucketWebsiteConfigurationMapOutputWithContext(context.Context) BucketWebsiteConfigurationMapOutput
}

type BucketWebsiteConfigurationMap map[string]BucketWebsiteConfigurationInput

func (BucketWebsiteConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketWebsiteConfiguration)(nil)).Elem()
}

func (i BucketWebsiteConfigurationMap) ToBucketWebsiteConfigurationMapOutput() BucketWebsiteConfigurationMapOutput {
	return i.ToBucketWebsiteConfigurationMapOutputWithContext(context.Background())
}

func (i BucketWebsiteConfigurationMap) ToBucketWebsiteConfigurationMapOutputWithContext(ctx context.Context) BucketWebsiteConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketWebsiteConfigurationMapOutput)
}

type BucketWebsiteConfigurationOutput struct{ *pulumi.OutputState }

func (BucketWebsiteConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketWebsiteConfiguration)(nil)).Elem()
}

func (o BucketWebsiteConfigurationOutput) ToBucketWebsiteConfigurationOutput() BucketWebsiteConfigurationOutput {
	return o
}

func (o BucketWebsiteConfigurationOutput) ToBucketWebsiteConfigurationOutputWithContext(ctx context.Context) BucketWebsiteConfigurationOutput {
	return o
}

// The name of the bucket.
func (o BucketWebsiteConfigurationOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketWebsiteConfiguration) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// The name of the error file for the website detailed below.
func (o BucketWebsiteConfigurationOutput) ErrorDocument() BucketWebsiteConfigurationErrorDocumentPtrOutput {
	return o.ApplyT(func(v *BucketWebsiteConfiguration) BucketWebsiteConfigurationErrorDocumentPtrOutput {
		return v.ErrorDocument
	}).(BucketWebsiteConfigurationErrorDocumentPtrOutput)
}

// The name of the index file for the website detailed below.
func (o BucketWebsiteConfigurationOutput) IndexDocument() BucketWebsiteConfigurationIndexDocumentOutput {
	return o.ApplyT(func(v *BucketWebsiteConfiguration) BucketWebsiteConfigurationIndexDocumentOutput {
		return v.IndexDocument
	}).(BucketWebsiteConfigurationIndexDocumentOutput)
}

// The projectId you want to attach the resource to
func (o BucketWebsiteConfigurationOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketWebsiteConfiguration) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The region you want to attach the resource to
func (o BucketWebsiteConfigurationOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketWebsiteConfiguration) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The domain of the website endpoint. This is used to create DNS alias [records](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/).
func (o BucketWebsiteConfigurationOutput) WebsiteDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketWebsiteConfiguration) pulumi.StringOutput { return v.WebsiteDomain }).(pulumi.StringOutput)
}

// The website endpoint.
func (o BucketWebsiteConfigurationOutput) WebsiteEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketWebsiteConfiguration) pulumi.StringOutput { return v.WebsiteEndpoint }).(pulumi.StringOutput)
}

type BucketWebsiteConfigurationArrayOutput struct{ *pulumi.OutputState }

func (BucketWebsiteConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketWebsiteConfiguration)(nil)).Elem()
}

func (o BucketWebsiteConfigurationArrayOutput) ToBucketWebsiteConfigurationArrayOutput() BucketWebsiteConfigurationArrayOutput {
	return o
}

func (o BucketWebsiteConfigurationArrayOutput) ToBucketWebsiteConfigurationArrayOutputWithContext(ctx context.Context) BucketWebsiteConfigurationArrayOutput {
	return o
}

func (o BucketWebsiteConfigurationArrayOutput) Index(i pulumi.IntInput) BucketWebsiteConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BucketWebsiteConfiguration {
		return vs[0].([]*BucketWebsiteConfiguration)[vs[1].(int)]
	}).(BucketWebsiteConfigurationOutput)
}

type BucketWebsiteConfigurationMapOutput struct{ *pulumi.OutputState }

func (BucketWebsiteConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketWebsiteConfiguration)(nil)).Elem()
}

func (o BucketWebsiteConfigurationMapOutput) ToBucketWebsiteConfigurationMapOutput() BucketWebsiteConfigurationMapOutput {
	return o
}

func (o BucketWebsiteConfigurationMapOutput) ToBucketWebsiteConfigurationMapOutputWithContext(ctx context.Context) BucketWebsiteConfigurationMapOutput {
	return o
}

func (o BucketWebsiteConfigurationMapOutput) MapIndex(k pulumi.StringInput) BucketWebsiteConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BucketWebsiteConfiguration {
		return vs[0].(map[string]*BucketWebsiteConfiguration)[vs[1].(string)]
	}).(BucketWebsiteConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketWebsiteConfigurationInput)(nil)).Elem(), &BucketWebsiteConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketWebsiteConfigurationArrayInput)(nil)).Elem(), BucketWebsiteConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketWebsiteConfigurationMapInput)(nil)).Elem(), BucketWebsiteConfigurationMap{})
	pulumi.RegisterOutputType(BucketWebsiteConfigurationOutput{})
	pulumi.RegisterOutputType(BucketWebsiteConfigurationArrayOutput{})
	pulumi.RegisterOutputType(BucketWebsiteConfigurationMapOutput{})
}
