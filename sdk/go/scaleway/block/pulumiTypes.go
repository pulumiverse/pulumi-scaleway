// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package block

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

var _ = internal.GetEnvOrDefault

type SnapshotImport struct {
	// Bucket containing qcow
	Bucket string `pulumi:"bucket"`
	// Key of the qcow file in the specified bucket
	Key string `pulumi:"key"`
}

// SnapshotImportInput is an input type that accepts SnapshotImportArgs and SnapshotImportOutput values.
// You can construct a concrete instance of `SnapshotImportInput` via:
//
//	SnapshotImportArgs{...}
type SnapshotImportInput interface {
	pulumi.Input

	ToSnapshotImportOutput() SnapshotImportOutput
	ToSnapshotImportOutputWithContext(context.Context) SnapshotImportOutput
}

type SnapshotImportArgs struct {
	// Bucket containing qcow
	Bucket pulumi.StringInput `pulumi:"bucket"`
	// Key of the qcow file in the specified bucket
	Key pulumi.StringInput `pulumi:"key"`
}

func (SnapshotImportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SnapshotImport)(nil)).Elem()
}

func (i SnapshotImportArgs) ToSnapshotImportOutput() SnapshotImportOutput {
	return i.ToSnapshotImportOutputWithContext(context.Background())
}

func (i SnapshotImportArgs) ToSnapshotImportOutputWithContext(ctx context.Context) SnapshotImportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotImportOutput)
}

func (i SnapshotImportArgs) ToSnapshotImportPtrOutput() SnapshotImportPtrOutput {
	return i.ToSnapshotImportPtrOutputWithContext(context.Background())
}

func (i SnapshotImportArgs) ToSnapshotImportPtrOutputWithContext(ctx context.Context) SnapshotImportPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotImportOutput).ToSnapshotImportPtrOutputWithContext(ctx)
}

// SnapshotImportPtrInput is an input type that accepts SnapshotImportArgs, SnapshotImportPtr and SnapshotImportPtrOutput values.
// You can construct a concrete instance of `SnapshotImportPtrInput` via:
//
//	        SnapshotImportArgs{...}
//
//	or:
//
//	        nil
type SnapshotImportPtrInput interface {
	pulumi.Input

	ToSnapshotImportPtrOutput() SnapshotImportPtrOutput
	ToSnapshotImportPtrOutputWithContext(context.Context) SnapshotImportPtrOutput
}

type snapshotImportPtrType SnapshotImportArgs

func SnapshotImportPtr(v *SnapshotImportArgs) SnapshotImportPtrInput {
	return (*snapshotImportPtrType)(v)
}

func (*snapshotImportPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotImport)(nil)).Elem()
}

func (i *snapshotImportPtrType) ToSnapshotImportPtrOutput() SnapshotImportPtrOutput {
	return i.ToSnapshotImportPtrOutputWithContext(context.Background())
}

func (i *snapshotImportPtrType) ToSnapshotImportPtrOutputWithContext(ctx context.Context) SnapshotImportPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotImportPtrOutput)
}

type SnapshotImportOutput struct{ *pulumi.OutputState }

func (SnapshotImportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SnapshotImport)(nil)).Elem()
}

func (o SnapshotImportOutput) ToSnapshotImportOutput() SnapshotImportOutput {
	return o
}

func (o SnapshotImportOutput) ToSnapshotImportOutputWithContext(ctx context.Context) SnapshotImportOutput {
	return o
}

func (o SnapshotImportOutput) ToSnapshotImportPtrOutput() SnapshotImportPtrOutput {
	return o.ToSnapshotImportPtrOutputWithContext(context.Background())
}

func (o SnapshotImportOutput) ToSnapshotImportPtrOutputWithContext(ctx context.Context) SnapshotImportPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SnapshotImport) *SnapshotImport {
		return &v
	}).(SnapshotImportPtrOutput)
}

// Bucket containing qcow
func (o SnapshotImportOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v SnapshotImport) string { return v.Bucket }).(pulumi.StringOutput)
}

// Key of the qcow file in the specified bucket
func (o SnapshotImportOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v SnapshotImport) string { return v.Key }).(pulumi.StringOutput)
}

type SnapshotImportPtrOutput struct{ *pulumi.OutputState }

func (SnapshotImportPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotImport)(nil)).Elem()
}

func (o SnapshotImportPtrOutput) ToSnapshotImportPtrOutput() SnapshotImportPtrOutput {
	return o
}

func (o SnapshotImportPtrOutput) ToSnapshotImportPtrOutputWithContext(ctx context.Context) SnapshotImportPtrOutput {
	return o
}

func (o SnapshotImportPtrOutput) Elem() SnapshotImportOutput {
	return o.ApplyT(func(v *SnapshotImport) SnapshotImport {
		if v != nil {
			return *v
		}
		var ret SnapshotImport
		return ret
	}).(SnapshotImportOutput)
}

// Bucket containing qcow
func (o SnapshotImportPtrOutput) Bucket() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnapshotImport) *string {
		if v == nil {
			return nil
		}
		return &v.Bucket
	}).(pulumi.StringPtrOutput)
}

// Key of the qcow file in the specified bucket
func (o SnapshotImportPtrOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnapshotImport) *string {
		if v == nil {
			return nil
		}
		return &v.Key
	}).(pulumi.StringPtrOutput)
}

type GetSnapshotImport struct {
	// Bucket containing qcow
	Bucket string `pulumi:"bucket"`
	// Key of the qcow file in the specified bucket
	Key string `pulumi:"key"`
}

// GetSnapshotImportInput is an input type that accepts GetSnapshotImportArgs and GetSnapshotImportOutput values.
// You can construct a concrete instance of `GetSnapshotImportInput` via:
//
//	GetSnapshotImportArgs{...}
type GetSnapshotImportInput interface {
	pulumi.Input

	ToGetSnapshotImportOutput() GetSnapshotImportOutput
	ToGetSnapshotImportOutputWithContext(context.Context) GetSnapshotImportOutput
}

type GetSnapshotImportArgs struct {
	// Bucket containing qcow
	Bucket pulumi.StringInput `pulumi:"bucket"`
	// Key of the qcow file in the specified bucket
	Key pulumi.StringInput `pulumi:"key"`
}

func (GetSnapshotImportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSnapshotImport)(nil)).Elem()
}

func (i GetSnapshotImportArgs) ToGetSnapshotImportOutput() GetSnapshotImportOutput {
	return i.ToGetSnapshotImportOutputWithContext(context.Background())
}

func (i GetSnapshotImportArgs) ToGetSnapshotImportOutputWithContext(ctx context.Context) GetSnapshotImportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSnapshotImportOutput)
}

// GetSnapshotImportArrayInput is an input type that accepts GetSnapshotImportArray and GetSnapshotImportArrayOutput values.
// You can construct a concrete instance of `GetSnapshotImportArrayInput` via:
//
//	GetSnapshotImportArray{ GetSnapshotImportArgs{...} }
type GetSnapshotImportArrayInput interface {
	pulumi.Input

	ToGetSnapshotImportArrayOutput() GetSnapshotImportArrayOutput
	ToGetSnapshotImportArrayOutputWithContext(context.Context) GetSnapshotImportArrayOutput
}

type GetSnapshotImportArray []GetSnapshotImportInput

func (GetSnapshotImportArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSnapshotImport)(nil)).Elem()
}

func (i GetSnapshotImportArray) ToGetSnapshotImportArrayOutput() GetSnapshotImportArrayOutput {
	return i.ToGetSnapshotImportArrayOutputWithContext(context.Background())
}

func (i GetSnapshotImportArray) ToGetSnapshotImportArrayOutputWithContext(ctx context.Context) GetSnapshotImportArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSnapshotImportArrayOutput)
}

type GetSnapshotImportOutput struct{ *pulumi.OutputState }

func (GetSnapshotImportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSnapshotImport)(nil)).Elem()
}

func (o GetSnapshotImportOutput) ToGetSnapshotImportOutput() GetSnapshotImportOutput {
	return o
}

func (o GetSnapshotImportOutput) ToGetSnapshotImportOutputWithContext(ctx context.Context) GetSnapshotImportOutput {
	return o
}

// Bucket containing qcow
func (o GetSnapshotImportOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotImport) string { return v.Bucket }).(pulumi.StringOutput)
}

// Key of the qcow file in the specified bucket
func (o GetSnapshotImportOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotImport) string { return v.Key }).(pulumi.StringOutput)
}

type GetSnapshotImportArrayOutput struct{ *pulumi.OutputState }

func (GetSnapshotImportArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSnapshotImport)(nil)).Elem()
}

func (o GetSnapshotImportArrayOutput) ToGetSnapshotImportArrayOutput() GetSnapshotImportArrayOutput {
	return o
}

func (o GetSnapshotImportArrayOutput) ToGetSnapshotImportArrayOutputWithContext(ctx context.Context) GetSnapshotImportArrayOutput {
	return o
}

func (o GetSnapshotImportArrayOutput) Index(i pulumi.IntInput) GetSnapshotImportOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetSnapshotImport {
		return vs[0].([]GetSnapshotImport)[vs[1].(int)]
	}).(GetSnapshotImportOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotImportInput)(nil)).Elem(), SnapshotImportArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotImportPtrInput)(nil)).Elem(), SnapshotImportArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetSnapshotImportInput)(nil)).Elem(), GetSnapshotImportArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetSnapshotImportArrayInput)(nil)).Elem(), GetSnapshotImportArray{})
	pulumi.RegisterOutputType(SnapshotImportOutput{})
	pulumi.RegisterOutputType(SnapshotImportPtrOutput{})
	pulumi.RegisterOutputType(GetSnapshotImportOutput{})
	pulumi.RegisterOutputType(GetSnapshotImportArrayOutput{})
}
