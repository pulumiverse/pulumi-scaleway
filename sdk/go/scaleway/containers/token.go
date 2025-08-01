// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package containers

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// The `containers.Token` resource allows you to create and manage authentication tokens for Scaleway [Serverless Containers](https://www.scaleway.com/en/docs/serverless/containers/).
//
// Refer to the Containers tokens [documentation](https://www.scaleway.com/en/docs/serverless/containers/how-to/create-auth-token-from-console/) and [API documentation](https://www.scaleway.com/en/developers/api/serverless-containers/#path-tokens-list-all-tokens) for more information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/containers"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := containers.NewNamespace(ctx, "main", &containers.NamespaceArgs{
//				Name: pulumi.String("test-container-token-ns"),
//			})
//			if err != nil {
//				return err
//			}
//			mainContainer, err := containers.NewContainer(ctx, "main", &containers.ContainerArgs{
//				NamespaceId: main.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			// Namespace Token
//			_, err = containers.NewToken(ctx, "namespace", &containers.TokenArgs{
//				NamespaceId: main.ID(),
//				ExpiresAt:   pulumi.String("2022-10-18T11:35:15+02:00"),
//			})
//			if err != nil {
//				return err
//			}
//			// Container Token
//			_, err = containers.NewToken(ctx, "container", &containers.TokenArgs{
//				ContainerId: mainContainer.ID(),
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
// Tokens can be imported using `{region}/{id}`, as shown below:
//
// bash
//
// ```sh
// $ pulumi import scaleway:containers/token:Token main fr-par/11111111-1111-1111-1111-111111111111
// ```
type Token struct {
	pulumi.CustomResourceState

	// The unique identifier of the container.
	//
	// > Only one of `namespaceId` or `containerId` must be set.
	ContainerId pulumi.StringPtrOutput `pulumi:"containerId"`
	// The description of the token.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The expiration date of the token.
	ExpiresAt pulumi.StringPtrOutput `pulumi:"expiresAt"`
	// The unique identifier of the Containers namespace.
	NamespaceId pulumi.StringPtrOutput `pulumi:"namespaceId"`
	// `region`). The region in which the namespace is created.
	//
	// > **Important** Updating any of the arguments above will recreate the token.
	Region pulumi.StringOutput `pulumi:"region"`
	// The token.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewToken registers a new resource with the given unique name, arguments, and options.
func NewToken(ctx *pulumi.Context,
	name string, args *TokenArgs, opts ...pulumi.ResourceOption) (*Token, error) {
	if args == nil {
		args = &TokenArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("scaleway:index/containerToken:ContainerToken"),
		},
	})
	opts = append(opts, aliases)
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"value",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Token
	err := ctx.RegisterResource("scaleway:containers/token:Token", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetToken gets an existing Token resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TokenState, opts ...pulumi.ResourceOption) (*Token, error) {
	var resource Token
	err := ctx.ReadResource("scaleway:containers/token:Token", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Token resources.
type tokenState struct {
	// The unique identifier of the container.
	//
	// > Only one of `namespaceId` or `containerId` must be set.
	ContainerId *string `pulumi:"containerId"`
	// The description of the token.
	Description *string `pulumi:"description"`
	// The expiration date of the token.
	ExpiresAt *string `pulumi:"expiresAt"`
	// The unique identifier of the Containers namespace.
	NamespaceId *string `pulumi:"namespaceId"`
	// `region`). The region in which the namespace is created.
	//
	// > **Important** Updating any of the arguments above will recreate the token.
	Region *string `pulumi:"region"`
	// The token.
	Value *string `pulumi:"value"`
}

type TokenState struct {
	// The unique identifier of the container.
	//
	// > Only one of `namespaceId` or `containerId` must be set.
	ContainerId pulumi.StringPtrInput
	// The description of the token.
	Description pulumi.StringPtrInput
	// The expiration date of the token.
	ExpiresAt pulumi.StringPtrInput
	// The unique identifier of the Containers namespace.
	NamespaceId pulumi.StringPtrInput
	// `region`). The region in which the namespace is created.
	//
	// > **Important** Updating any of the arguments above will recreate the token.
	Region pulumi.StringPtrInput
	// The token.
	Value pulumi.StringPtrInput
}

func (TokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*tokenState)(nil)).Elem()
}

type tokenArgs struct {
	// The unique identifier of the container.
	//
	// > Only one of `namespaceId` or `containerId` must be set.
	ContainerId *string `pulumi:"containerId"`
	// The description of the token.
	Description *string `pulumi:"description"`
	// The expiration date of the token.
	ExpiresAt *string `pulumi:"expiresAt"`
	// The unique identifier of the Containers namespace.
	NamespaceId *string `pulumi:"namespaceId"`
	// `region`). The region in which the namespace is created.
	//
	// > **Important** Updating any of the arguments above will recreate the token.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a Token resource.
type TokenArgs struct {
	// The unique identifier of the container.
	//
	// > Only one of `namespaceId` or `containerId` must be set.
	ContainerId pulumi.StringPtrInput
	// The description of the token.
	Description pulumi.StringPtrInput
	// The expiration date of the token.
	ExpiresAt pulumi.StringPtrInput
	// The unique identifier of the Containers namespace.
	NamespaceId pulumi.StringPtrInput
	// `region`). The region in which the namespace is created.
	//
	// > **Important** Updating any of the arguments above will recreate the token.
	Region pulumi.StringPtrInput
}

func (TokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tokenArgs)(nil)).Elem()
}

type TokenInput interface {
	pulumi.Input

	ToTokenOutput() TokenOutput
	ToTokenOutputWithContext(ctx context.Context) TokenOutput
}

func (*Token) ElementType() reflect.Type {
	return reflect.TypeOf((**Token)(nil)).Elem()
}

func (i *Token) ToTokenOutput() TokenOutput {
	return i.ToTokenOutputWithContext(context.Background())
}

func (i *Token) ToTokenOutputWithContext(ctx context.Context) TokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenOutput)
}

// TokenArrayInput is an input type that accepts TokenArray and TokenArrayOutput values.
// You can construct a concrete instance of `TokenArrayInput` via:
//
//	TokenArray{ TokenArgs{...} }
type TokenArrayInput interface {
	pulumi.Input

	ToTokenArrayOutput() TokenArrayOutput
	ToTokenArrayOutputWithContext(context.Context) TokenArrayOutput
}

type TokenArray []TokenInput

func (TokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Token)(nil)).Elem()
}

func (i TokenArray) ToTokenArrayOutput() TokenArrayOutput {
	return i.ToTokenArrayOutputWithContext(context.Background())
}

func (i TokenArray) ToTokenArrayOutputWithContext(ctx context.Context) TokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenArrayOutput)
}

// TokenMapInput is an input type that accepts TokenMap and TokenMapOutput values.
// You can construct a concrete instance of `TokenMapInput` via:
//
//	TokenMap{ "key": TokenArgs{...} }
type TokenMapInput interface {
	pulumi.Input

	ToTokenMapOutput() TokenMapOutput
	ToTokenMapOutputWithContext(context.Context) TokenMapOutput
}

type TokenMap map[string]TokenInput

func (TokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Token)(nil)).Elem()
}

func (i TokenMap) ToTokenMapOutput() TokenMapOutput {
	return i.ToTokenMapOutputWithContext(context.Background())
}

func (i TokenMap) ToTokenMapOutputWithContext(ctx context.Context) TokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenMapOutput)
}

type TokenOutput struct{ *pulumi.OutputState }

func (TokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Token)(nil)).Elem()
}

func (o TokenOutput) ToTokenOutput() TokenOutput {
	return o
}

func (o TokenOutput) ToTokenOutputWithContext(ctx context.Context) TokenOutput {
	return o
}

// The unique identifier of the container.
//
// > Only one of `namespaceId` or `containerId` must be set.
func (o TokenOutput) ContainerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Token) pulumi.StringPtrOutput { return v.ContainerId }).(pulumi.StringPtrOutput)
}

// The description of the token.
func (o TokenOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Token) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The expiration date of the token.
func (o TokenOutput) ExpiresAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Token) pulumi.StringPtrOutput { return v.ExpiresAt }).(pulumi.StringPtrOutput)
}

// The unique identifier of the Containers namespace.
func (o TokenOutput) NamespaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Token) pulumi.StringPtrOutput { return v.NamespaceId }).(pulumi.StringPtrOutput)
}

// `region`). The region in which the namespace is created.
//
// > **Important** Updating any of the arguments above will recreate the token.
func (o TokenOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Token) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The token.
func (o TokenOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *Token) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

type TokenArrayOutput struct{ *pulumi.OutputState }

func (TokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Token)(nil)).Elem()
}

func (o TokenArrayOutput) ToTokenArrayOutput() TokenArrayOutput {
	return o
}

func (o TokenArrayOutput) ToTokenArrayOutputWithContext(ctx context.Context) TokenArrayOutput {
	return o
}

func (o TokenArrayOutput) Index(i pulumi.IntInput) TokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Token {
		return vs[0].([]*Token)[vs[1].(int)]
	}).(TokenOutput)
}

type TokenMapOutput struct{ *pulumi.OutputState }

func (TokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Token)(nil)).Elem()
}

func (o TokenMapOutput) ToTokenMapOutput() TokenMapOutput {
	return o
}

func (o TokenMapOutput) ToTokenMapOutputWithContext(ctx context.Context) TokenMapOutput {
	return o
}

func (o TokenMapOutput) MapIndex(k pulumi.StringInput) TokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Token {
		return vs[0].(map[string]*Token)[vs[1].(string)]
	}).(TokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TokenInput)(nil)).Elem(), &Token{})
	pulumi.RegisterInputType(reflect.TypeOf((*TokenArrayInput)(nil)).Elem(), TokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TokenMapInput)(nil)).Elem(), TokenMap{})
	pulumi.RegisterOutputType(TokenOutput{})
	pulumi.RegisterOutputType(TokenArrayOutput{})
	pulumi.RegisterOutputType(TokenMapOutput{})
}
