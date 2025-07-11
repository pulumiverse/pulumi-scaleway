// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// The `observability.Token` resource allows you to create and manage your Cockpit [tokens](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#tokens).
//
// Refer to Cockpit's [product documentation](https://www.scaleway.com/en/docs/observability/cockpit/concepts/) and [API documentation](https://www.scaleway.com/en/developers/api/cockpit/regional-api) for more information.
//
// ## Example Usage
//
// ### Use a Cockpit token
//
// The following commands allow you to:
//
// - create a Scaleway Project named `my-project`
// - create a Cockpit token named `my-awesome-token` inside the Project
// - assign `read` permissions to the token for metrics and logs
// - disable `write` permissions for metrics and logs
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/account"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/observability"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			project, err := account.NewProject(ctx, "project", &account.ProjectArgs{
//				Name: pulumi.String("my-project"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = observability.NewToken(ctx, "main", &observability.TokenArgs{
//				ProjectId: project.ID(),
//				Name:      pulumi.String("my-awesome-token"),
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
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/account"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/observability"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			project, err := account.NewProject(ctx, "project", &account.ProjectArgs{
//				Name: pulumi.String("my-project"),
//			})
//			if err != nil {
//				return err
//			}
//			// Create a token that can read metrics and logs but not write
//			_, err = observability.NewToken(ctx, "main", &observability.TokenArgs{
//				ProjectId: project.ID(),
//				Name:      pulumi.String("my-awesome-token"),
//				Scopes: &observability.TokenScopesArgs{
//					QueryMetrics: pulumi.Bool(true),
//					WriteMetrics: pulumi.Bool(false),
//					QueryLogs:    pulumi.Bool(true),
//					WriteLogs:    pulumi.Bool(false),
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
// This section explains how to import a Cockpit token using the `{region}/{id}` format.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/cockpitToken:CockpitToken main fr-par/11111111-1111-1111-1111-111111111111
// ```
//
// Deprecated: scaleway.index/cockpittoken.CockpitToken has been deprecated in favor of scaleway.observability/token.Token
type CockpitToken struct {
	pulumi.CustomResourceState

	// The date and time of the creation of the Cockpit Token (Format ISO 8601)
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The name of the token.
	Name pulumi.StringOutput `pulumi:"name"`
	// ) The ID of the Project the Cockpit is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// ) The region where the Cockpit token is located.
	Region pulumi.StringOutput `pulumi:"region"`
	// Scopes allowed, each with default values:
	Scopes CockpitTokenScopesOutput `pulumi:"scopes"`
	// The secret key of the token.
	SecretKey pulumi.StringOutput `pulumi:"secretKey"`
	// The date and time of the last update of the Cockpit Token (Format ISO 8601)
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewCockpitToken registers a new resource with the given unique name, arguments, and options.
func NewCockpitToken(ctx *pulumi.Context,
	name string, args *CockpitTokenArgs, opts ...pulumi.ResourceOption) (*CockpitToken, error) {
	if args == nil {
		args = &CockpitTokenArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"secretKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CockpitToken
	err := ctx.RegisterResource("scaleway:index/cockpitToken:CockpitToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCockpitToken gets an existing CockpitToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCockpitToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CockpitTokenState, opts ...pulumi.ResourceOption) (*CockpitToken, error) {
	var resource CockpitToken
	err := ctx.ReadResource("scaleway:index/cockpitToken:CockpitToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CockpitToken resources.
type cockpitTokenState struct {
	// The date and time of the creation of the Cockpit Token (Format ISO 8601)
	CreatedAt *string `pulumi:"createdAt"`
	// The name of the token.
	Name *string `pulumi:"name"`
	// ) The ID of the Project the Cockpit is associated with.
	ProjectId *string `pulumi:"projectId"`
	// ) The region where the Cockpit token is located.
	Region *string `pulumi:"region"`
	// Scopes allowed, each with default values:
	Scopes *CockpitTokenScopes `pulumi:"scopes"`
	// The secret key of the token.
	SecretKey *string `pulumi:"secretKey"`
	// The date and time of the last update of the Cockpit Token (Format ISO 8601)
	UpdatedAt *string `pulumi:"updatedAt"`
}

type CockpitTokenState struct {
	// The date and time of the creation of the Cockpit Token (Format ISO 8601)
	CreatedAt pulumi.StringPtrInput
	// The name of the token.
	Name pulumi.StringPtrInput
	// ) The ID of the Project the Cockpit is associated with.
	ProjectId pulumi.StringPtrInput
	// ) The region where the Cockpit token is located.
	Region pulumi.StringPtrInput
	// Scopes allowed, each with default values:
	Scopes CockpitTokenScopesPtrInput
	// The secret key of the token.
	SecretKey pulumi.StringPtrInput
	// The date and time of the last update of the Cockpit Token (Format ISO 8601)
	UpdatedAt pulumi.StringPtrInput
}

func (CockpitTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*cockpitTokenState)(nil)).Elem()
}

type cockpitTokenArgs struct {
	// The name of the token.
	Name *string `pulumi:"name"`
	// ) The ID of the Project the Cockpit is associated with.
	ProjectId *string `pulumi:"projectId"`
	// ) The region where the Cockpit token is located.
	Region *string `pulumi:"region"`
	// Scopes allowed, each with default values:
	Scopes *CockpitTokenScopes `pulumi:"scopes"`
}

// The set of arguments for constructing a CockpitToken resource.
type CockpitTokenArgs struct {
	// The name of the token.
	Name pulumi.StringPtrInput
	// ) The ID of the Project the Cockpit is associated with.
	ProjectId pulumi.StringPtrInput
	// ) The region where the Cockpit token is located.
	Region pulumi.StringPtrInput
	// Scopes allowed, each with default values:
	Scopes CockpitTokenScopesPtrInput
}

func (CockpitTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cockpitTokenArgs)(nil)).Elem()
}

type CockpitTokenInput interface {
	pulumi.Input

	ToCockpitTokenOutput() CockpitTokenOutput
	ToCockpitTokenOutputWithContext(ctx context.Context) CockpitTokenOutput
}

func (*CockpitToken) ElementType() reflect.Type {
	return reflect.TypeOf((**CockpitToken)(nil)).Elem()
}

func (i *CockpitToken) ToCockpitTokenOutput() CockpitTokenOutput {
	return i.ToCockpitTokenOutputWithContext(context.Background())
}

func (i *CockpitToken) ToCockpitTokenOutputWithContext(ctx context.Context) CockpitTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CockpitTokenOutput)
}

// CockpitTokenArrayInput is an input type that accepts CockpitTokenArray and CockpitTokenArrayOutput values.
// You can construct a concrete instance of `CockpitTokenArrayInput` via:
//
//	CockpitTokenArray{ CockpitTokenArgs{...} }
type CockpitTokenArrayInput interface {
	pulumi.Input

	ToCockpitTokenArrayOutput() CockpitTokenArrayOutput
	ToCockpitTokenArrayOutputWithContext(context.Context) CockpitTokenArrayOutput
}

type CockpitTokenArray []CockpitTokenInput

func (CockpitTokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CockpitToken)(nil)).Elem()
}

func (i CockpitTokenArray) ToCockpitTokenArrayOutput() CockpitTokenArrayOutput {
	return i.ToCockpitTokenArrayOutputWithContext(context.Background())
}

func (i CockpitTokenArray) ToCockpitTokenArrayOutputWithContext(ctx context.Context) CockpitTokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CockpitTokenArrayOutput)
}

// CockpitTokenMapInput is an input type that accepts CockpitTokenMap and CockpitTokenMapOutput values.
// You can construct a concrete instance of `CockpitTokenMapInput` via:
//
//	CockpitTokenMap{ "key": CockpitTokenArgs{...} }
type CockpitTokenMapInput interface {
	pulumi.Input

	ToCockpitTokenMapOutput() CockpitTokenMapOutput
	ToCockpitTokenMapOutputWithContext(context.Context) CockpitTokenMapOutput
}

type CockpitTokenMap map[string]CockpitTokenInput

func (CockpitTokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CockpitToken)(nil)).Elem()
}

func (i CockpitTokenMap) ToCockpitTokenMapOutput() CockpitTokenMapOutput {
	return i.ToCockpitTokenMapOutputWithContext(context.Background())
}

func (i CockpitTokenMap) ToCockpitTokenMapOutputWithContext(ctx context.Context) CockpitTokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CockpitTokenMapOutput)
}

type CockpitTokenOutput struct{ *pulumi.OutputState }

func (CockpitTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CockpitToken)(nil)).Elem()
}

func (o CockpitTokenOutput) ToCockpitTokenOutput() CockpitTokenOutput {
	return o
}

func (o CockpitTokenOutput) ToCockpitTokenOutputWithContext(ctx context.Context) CockpitTokenOutput {
	return o
}

// The date and time of the creation of the Cockpit Token (Format ISO 8601)
func (o CockpitTokenOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *CockpitToken) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The name of the token.
func (o CockpitTokenOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CockpitToken) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ) The ID of the Project the Cockpit is associated with.
func (o CockpitTokenOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *CockpitToken) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// ) The region where the Cockpit token is located.
func (o CockpitTokenOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *CockpitToken) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Scopes allowed, each with default values:
func (o CockpitTokenOutput) Scopes() CockpitTokenScopesOutput {
	return o.ApplyT(func(v *CockpitToken) CockpitTokenScopesOutput { return v.Scopes }).(CockpitTokenScopesOutput)
}

// The secret key of the token.
func (o CockpitTokenOutput) SecretKey() pulumi.StringOutput {
	return o.ApplyT(func(v *CockpitToken) pulumi.StringOutput { return v.SecretKey }).(pulumi.StringOutput)
}

// The date and time of the last update of the Cockpit Token (Format ISO 8601)
func (o CockpitTokenOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *CockpitToken) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type CockpitTokenArrayOutput struct{ *pulumi.OutputState }

func (CockpitTokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CockpitToken)(nil)).Elem()
}

func (o CockpitTokenArrayOutput) ToCockpitTokenArrayOutput() CockpitTokenArrayOutput {
	return o
}

func (o CockpitTokenArrayOutput) ToCockpitTokenArrayOutputWithContext(ctx context.Context) CockpitTokenArrayOutput {
	return o
}

func (o CockpitTokenArrayOutput) Index(i pulumi.IntInput) CockpitTokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CockpitToken {
		return vs[0].([]*CockpitToken)[vs[1].(int)]
	}).(CockpitTokenOutput)
}

type CockpitTokenMapOutput struct{ *pulumi.OutputState }

func (CockpitTokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CockpitToken)(nil)).Elem()
}

func (o CockpitTokenMapOutput) ToCockpitTokenMapOutput() CockpitTokenMapOutput {
	return o
}

func (o CockpitTokenMapOutput) ToCockpitTokenMapOutputWithContext(ctx context.Context) CockpitTokenMapOutput {
	return o
}

func (o CockpitTokenMapOutput) MapIndex(k pulumi.StringInput) CockpitTokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CockpitToken {
		return vs[0].(map[string]*CockpitToken)[vs[1].(string)]
	}).(CockpitTokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CockpitTokenInput)(nil)).Elem(), &CockpitToken{})
	pulumi.RegisterInputType(reflect.TypeOf((*CockpitTokenArrayInput)(nil)).Elem(), CockpitTokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CockpitTokenMapInput)(nil)).Elem(), CockpitTokenMap{})
	pulumi.RegisterOutputType(CockpitTokenOutput{})
	pulumi.RegisterOutputType(CockpitTokenArrayOutput{})
	pulumi.RegisterOutputType(CockpitTokenMapOutput{})
}
