// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// The `containers.Namespace` data source is used to retrieve information about a Serverless Containers namespace.
//
// Refer to the Serverless Containers [product documentation](https://www.scaleway.com/en/docs/serverless/containers/) and [API documentation](https://www.scaleway.com/en/developers/api/serverless-containers/) for more information.
//
// ## Retrieve a Serverless Containers namespace
//
// The following commands allow you to:
//
// - retrieve a namespace by its name
// - retrieve a namespace by its ID
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
//			// Get info by namespace name
//			_, err := containers.LookupNamespace(ctx, &containers.LookupNamespaceArgs{
//				Name: pulumi.StringRef("my-namespace-name"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			// Get info by namespace ID
//			_, err = containers.LookupNamespace(ctx, &containers.LookupNamespaceArgs{
//				NamespaceId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: scaleway.index/getcontainernamespace.getContainerNamespace has been deprecated in favor of scaleway.containers/getnamespace.getNamespace
func LookupContainerNamespace(ctx *pulumi.Context, args *LookupContainerNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupContainerNamespaceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupContainerNamespaceResult
	err := ctx.Invoke("scaleway:index/getContainerNamespace:getContainerNamespace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getContainerNamespace.
type LookupContainerNamespaceArgs struct {
	// The name of the namespace. Only one of `name` and `namespaceId` should be specified.
	Name *string `pulumi:"name"`
	// The unique identifier of the namespace. Only one of `name` and `namespaceId` should be specified.
	NamespaceId *string `pulumi:"namespaceId"`
	// `projectId`) The unique identifier of the project with which the namespace is associated.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region in which the namespace exists.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getContainerNamespace.
type LookupContainerNamespaceResult struct {
	ActivateVpcIntegration bool `pulumi:"activateVpcIntegration"`
	// The description of the namespace.
	Description     string `pulumi:"description"`
	DestroyRegistry bool   `pulumi:"destroyRegistry"`
	// The environment variables of the namespace.
	EnvironmentVariables map[string]string `pulumi:"environmentVariables"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        *string `pulumi:"name"`
	NamespaceId *string `pulumi:"namespaceId"`
	// The unique identifier of the organization with which the namespace is associated.
	OrganizationId string  `pulumi:"organizationId"`
	ProjectId      *string `pulumi:"projectId"`
	Region         *string `pulumi:"region"`
	// The registry endpoint of the namespace.
	RegistryEndpoint string `pulumi:"registryEndpoint"`
	// The unique identifier of the registry namespace of the Serverless Containers namespace.
	RegistryNamespaceId        string            `pulumi:"registryNamespaceId"`
	SecretEnvironmentVariables map[string]string `pulumi:"secretEnvironmentVariables"`
	Tags                       []string          `pulumi:"tags"`
}

func LookupContainerNamespaceOutput(ctx *pulumi.Context, args LookupContainerNamespaceOutputArgs, opts ...pulumi.InvokeOption) LookupContainerNamespaceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupContainerNamespaceResultOutput, error) {
			args := v.(LookupContainerNamespaceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway:index/getContainerNamespace:getContainerNamespace", args, LookupContainerNamespaceResultOutput{}, options).(LookupContainerNamespaceResultOutput), nil
		}).(LookupContainerNamespaceResultOutput)
}

// A collection of arguments for invoking getContainerNamespace.
type LookupContainerNamespaceOutputArgs struct {
	// The name of the namespace. Only one of `name` and `namespaceId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The unique identifier of the namespace. Only one of `name` and `namespaceId` should be specified.
	NamespaceId pulumi.StringPtrInput `pulumi:"namespaceId"`
	// `projectId`) The unique identifier of the project with which the namespace is associated.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// `region`) The region in which the namespace exists.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupContainerNamespaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerNamespaceArgs)(nil)).Elem()
}

// A collection of values returned by getContainerNamespace.
type LookupContainerNamespaceResultOutput struct{ *pulumi.OutputState }

func (LookupContainerNamespaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerNamespaceResult)(nil)).Elem()
}

func (o LookupContainerNamespaceResultOutput) ToLookupContainerNamespaceResultOutput() LookupContainerNamespaceResultOutput {
	return o
}

func (o LookupContainerNamespaceResultOutput) ToLookupContainerNamespaceResultOutputWithContext(ctx context.Context) LookupContainerNamespaceResultOutput {
	return o
}

func (o LookupContainerNamespaceResultOutput) ActivateVpcIntegration() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) bool { return v.ActivateVpcIntegration }).(pulumi.BoolOutput)
}

// The description of the namespace.
func (o LookupContainerNamespaceResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupContainerNamespaceResultOutput) DestroyRegistry() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) bool { return v.DestroyRegistry }).(pulumi.BoolOutput)
}

// The environment variables of the namespace.
func (o LookupContainerNamespaceResultOutput) EnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) map[string]string { return v.EnvironmentVariables }).(pulumi.StringMapOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupContainerNamespaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupContainerNamespaceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupContainerNamespaceResultOutput) NamespaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) *string { return v.NamespaceId }).(pulumi.StringPtrOutput)
}

// The unique identifier of the organization with which the namespace is associated.
func (o LookupContainerNamespaceResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupContainerNamespaceResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o LookupContainerNamespaceResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

// The registry endpoint of the namespace.
func (o LookupContainerNamespaceResultOutput) RegistryEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) string { return v.RegistryEndpoint }).(pulumi.StringOutput)
}

// The unique identifier of the registry namespace of the Serverless Containers namespace.
func (o LookupContainerNamespaceResultOutput) RegistryNamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) string { return v.RegistryNamespaceId }).(pulumi.StringOutput)
}

func (o LookupContainerNamespaceResultOutput) SecretEnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) map[string]string { return v.SecretEnvironmentVariables }).(pulumi.StringMapOutput)
}

func (o LookupContainerNamespaceResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContainerNamespaceResultOutput{})
}
