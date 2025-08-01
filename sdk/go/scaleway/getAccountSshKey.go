// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// The `account.SshKey` data source is used to retrieve information about a the SSH key of a Scaleway account.
//
// Refer to the Organizations and Projects [documentation](https://www.scaleway.com/en/docs/organizations-and-projects/how-to/create-ssh-key/) and [API documentation](https://www.scaleway.com/en/developers/api/iam/#path-ssh-keys) for more information.
//
// Deprecated: scaleway.index/getaccountsshkey.getAccountSshKey has been deprecated in favor of scaleway.account/getsshkey.getSshKey
func LookupAccountSshKey(ctx *pulumi.Context, args *LookupAccountSshKeyArgs, opts ...pulumi.InvokeOption) (*LookupAccountSshKeyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAccountSshKeyResult
	err := ctx.Invoke("scaleway:index/getAccountSshKey:getAccountSshKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccountSshKey.
type LookupAccountSshKeyArgs struct {
	// The name of the SSH key.
	Name *string `pulumi:"name"`
	// `projectId`) The unique identifier of the project with which the SSH key is associated.
	ProjectId *string `pulumi:"projectId"`
	// The unique identifier of the SSH key.
	//
	// > **Note** You must specify at least one: `name` and/or `sshKeyId`.
	SshKeyId *string `pulumi:"sshKeyId"`
}

// A collection of values returned by getAccountSshKey.
type LookupAccountSshKeyResult struct {
	CreatedAt   string `pulumi:"createdAt"`
	Disabled    bool   `pulumi:"disabled"`
	Fingerprint string `pulumi:"fingerprint"`
	// The provider-assigned unique ID for this managed resource.
	Id   string  `pulumi:"id"`
	Name *string `pulumi:"name"`
	// The unique identifier of the Organization with which the SSH key is associated.
	OrganizationId string  `pulumi:"organizationId"`
	ProjectId      *string `pulumi:"projectId"`
	// The string of the SSH public key.
	PublicKey string  `pulumi:"publicKey"`
	SshKeyId  *string `pulumi:"sshKeyId"`
	UpdatedAt string  `pulumi:"updatedAt"`
}

func LookupAccountSshKeyOutput(ctx *pulumi.Context, args LookupAccountSshKeyOutputArgs, opts ...pulumi.InvokeOption) LookupAccountSshKeyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAccountSshKeyResultOutput, error) {
			args := v.(LookupAccountSshKeyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway:index/getAccountSshKey:getAccountSshKey", args, LookupAccountSshKeyResultOutput{}, options).(LookupAccountSshKeyResultOutput), nil
		}).(LookupAccountSshKeyResultOutput)
}

// A collection of arguments for invoking getAccountSshKey.
type LookupAccountSshKeyOutputArgs struct {
	// The name of the SSH key.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// `projectId`) The unique identifier of the project with which the SSH key is associated.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// The unique identifier of the SSH key.
	//
	// > **Note** You must specify at least one: `name` and/or `sshKeyId`.
	SshKeyId pulumi.StringPtrInput `pulumi:"sshKeyId"`
}

func (LookupAccountSshKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountSshKeyArgs)(nil)).Elem()
}

// A collection of values returned by getAccountSshKey.
type LookupAccountSshKeyResultOutput struct{ *pulumi.OutputState }

func (LookupAccountSshKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountSshKeyResult)(nil)).Elem()
}

func (o LookupAccountSshKeyResultOutput) ToLookupAccountSshKeyResultOutput() LookupAccountSshKeyResultOutput {
	return o
}

func (o LookupAccountSshKeyResultOutput) ToLookupAccountSshKeyResultOutputWithContext(ctx context.Context) LookupAccountSshKeyResultOutput {
	return o
}

func (o LookupAccountSshKeyResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountSshKeyResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupAccountSshKeyResultOutput) Disabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAccountSshKeyResult) bool { return v.Disabled }).(pulumi.BoolOutput)
}

func (o LookupAccountSshKeyResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountSshKeyResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAccountSshKeyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountSshKeyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAccountSshKeyResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountSshKeyResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The unique identifier of the Organization with which the SSH key is associated.
func (o LookupAccountSshKeyResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountSshKeyResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupAccountSshKeyResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountSshKeyResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// The string of the SSH public key.
func (o LookupAccountSshKeyResultOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountSshKeyResult) string { return v.PublicKey }).(pulumi.StringOutput)
}

func (o LookupAccountSshKeyResultOutput) SshKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountSshKeyResult) *string { return v.SshKeyId }).(pulumi.StringPtrOutput)
}

func (o LookupAccountSshKeyResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountSshKeyResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccountSshKeyResultOutput{})
}
