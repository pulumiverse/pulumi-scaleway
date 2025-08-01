// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Scaleway IAM Applications. For more information refer to the [IAM API documentation](https://www.scaleway.com/en/developers/api/iam/#applications-83ce5e).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/iam"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iam.NewApplication(ctx, "main", &iam.ApplicationArgs{
//				Name:        pulumi.String("My application"),
//				Description: pulumi.String("a description"),
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
// Applications can be imported using the `{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:iam/application:Application main 11111111-1111-1111-1111-111111111111
// ```
type Application struct {
	pulumi.CustomResourceState

	// The date and time of the creation of the application.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The description of the iam application.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether the application is editable.
	Editable pulumi.BoolOutput `pulumi:"editable"`
	// The name of the iam application.
	Name pulumi.StringOutput `pulumi:"name"`
	// `organizationId`) The ID of the organization the application is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// The tags associated with the application.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The date and time of the last update of the application.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		args = &ApplicationArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("scaleway:index/iamApplication:IamApplication"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Application
	err := ctx.RegisterResource("scaleway:iam/application:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("scaleway:iam/application:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type applicationState struct {
	// The date and time of the creation of the application.
	CreatedAt *string `pulumi:"createdAt"`
	// The description of the iam application.
	Description *string `pulumi:"description"`
	// Whether the application is editable.
	Editable *bool `pulumi:"editable"`
	// The name of the iam application.
	Name *string `pulumi:"name"`
	// `organizationId`) The ID of the organization the application is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// The tags associated with the application.
	Tags []string `pulumi:"tags"`
	// The date and time of the last update of the application.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type ApplicationState struct {
	// The date and time of the creation of the application.
	CreatedAt pulumi.StringPtrInput
	// The description of the iam application.
	Description pulumi.StringPtrInput
	// Whether the application is editable.
	Editable pulumi.BoolPtrInput
	// The name of the iam application.
	Name pulumi.StringPtrInput
	// `organizationId`) The ID of the organization the application is associated with.
	OrganizationId pulumi.StringPtrInput
	// The tags associated with the application.
	Tags pulumi.StringArrayInput
	// The date and time of the last update of the application.
	UpdatedAt pulumi.StringPtrInput
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	// The description of the iam application.
	Description *string `pulumi:"description"`
	// The name of the iam application.
	Name *string `pulumi:"name"`
	// `organizationId`) The ID of the organization the application is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// The tags associated with the application.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// The description of the iam application.
	Description pulumi.StringPtrInput
	// The name of the iam application.
	Name pulumi.StringPtrInput
	// `organizationId`) The ID of the organization the application is associated with.
	OrganizationId pulumi.StringPtrInput
	// The tags associated with the application.
	Tags pulumi.StringArrayInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}

type ApplicationInput interface {
	pulumi.Input

	ToApplicationOutput() ApplicationOutput
	ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput
}

func (*Application) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (i *Application) ToApplicationOutput() ApplicationOutput {
	return i.ToApplicationOutputWithContext(context.Background())
}

func (i *Application) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOutput)
}

// ApplicationArrayInput is an input type that accepts ApplicationArray and ApplicationArrayOutput values.
// You can construct a concrete instance of `ApplicationArrayInput` via:
//
//	ApplicationArray{ ApplicationArgs{...} }
type ApplicationArrayInput interface {
	pulumi.Input

	ToApplicationArrayOutput() ApplicationArrayOutput
	ToApplicationArrayOutputWithContext(context.Context) ApplicationArrayOutput
}

type ApplicationArray []ApplicationInput

func (ApplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Application)(nil)).Elem()
}

func (i ApplicationArray) ToApplicationArrayOutput() ApplicationArrayOutput {
	return i.ToApplicationArrayOutputWithContext(context.Background())
}

func (i ApplicationArray) ToApplicationArrayOutputWithContext(ctx context.Context) ApplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationArrayOutput)
}

// ApplicationMapInput is an input type that accepts ApplicationMap and ApplicationMapOutput values.
// You can construct a concrete instance of `ApplicationMapInput` via:
//
//	ApplicationMap{ "key": ApplicationArgs{...} }
type ApplicationMapInput interface {
	pulumi.Input

	ToApplicationMapOutput() ApplicationMapOutput
	ToApplicationMapOutputWithContext(context.Context) ApplicationMapOutput
}

type ApplicationMap map[string]ApplicationInput

func (ApplicationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Application)(nil)).Elem()
}

func (i ApplicationMap) ToApplicationMapOutput() ApplicationMapOutput {
	return i.ToApplicationMapOutputWithContext(context.Background())
}

func (i ApplicationMap) ToApplicationMapOutputWithContext(ctx context.Context) ApplicationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationMapOutput)
}

type ApplicationOutput struct{ *pulumi.OutputState }

func (ApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (o ApplicationOutput) ToApplicationOutput() ApplicationOutput {
	return o
}

func (o ApplicationOutput) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return o
}

// The date and time of the creation of the application.
func (o ApplicationOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The description of the iam application.
func (o ApplicationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether the application is editable.
func (o ApplicationOutput) Editable() pulumi.BoolOutput {
	return o.ApplyT(func(v *Application) pulumi.BoolOutput { return v.Editable }).(pulumi.BoolOutput)
}

// The name of the iam application.
func (o ApplicationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// `organizationId`) The ID of the organization the application is associated with.
func (o ApplicationOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// The tags associated with the application.
func (o ApplicationOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Application) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The date and time of the last update of the application.
func (o ApplicationOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type ApplicationArrayOutput struct{ *pulumi.OutputState }

func (ApplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Application)(nil)).Elem()
}

func (o ApplicationArrayOutput) ToApplicationArrayOutput() ApplicationArrayOutput {
	return o
}

func (o ApplicationArrayOutput) ToApplicationArrayOutputWithContext(ctx context.Context) ApplicationArrayOutput {
	return o
}

func (o ApplicationArrayOutput) Index(i pulumi.IntInput) ApplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Application {
		return vs[0].([]*Application)[vs[1].(int)]
	}).(ApplicationOutput)
}

type ApplicationMapOutput struct{ *pulumi.OutputState }

func (ApplicationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Application)(nil)).Elem()
}

func (o ApplicationMapOutput) ToApplicationMapOutput() ApplicationMapOutput {
	return o
}

func (o ApplicationMapOutput) ToApplicationMapOutputWithContext(ctx context.Context) ApplicationMapOutput {
	return o
}

func (o ApplicationMapOutput) MapIndex(k pulumi.StringInput) ApplicationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Application {
		return vs[0].(map[string]*Application)[vs[1].(string)]
	}).(ApplicationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationInput)(nil)).Elem(), &Application{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationArrayInput)(nil)).Elem(), ApplicationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationMapInput)(nil)).Elem(), ApplicationMap{})
	pulumi.RegisterOutputType(ApplicationOutput{})
	pulumi.RegisterOutputType(ApplicationArrayOutput{})
	pulumi.RegisterOutputType(ApplicationMapOutput{})
}
