// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancers

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

type Certificate struct {
	pulumi.CustomResourceState

	// Main domain of the certificate
	CommonName pulumi.StringOutput `pulumi:"commonName"`
	// The custom type certificate type configuration
	CustomCertificate CertificateCustomCertificatePtrOutput `pulumi:"customCertificate"`
	// The identifier (SHA-1) of the certificate
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The load-balancer ID
	LbId pulumi.StringOutput `pulumi:"lbId"`
	// The Let's Encrypt type certificate configuration
	Letsencrypt CertificateLetsencryptPtrOutput `pulumi:"letsencrypt"`
	// The name of the load-balancer certificate
	Name pulumi.StringOutput `pulumi:"name"`
	// The not valid after validity bound timestamp
	NotValidAfter pulumi.StringOutput `pulumi:"notValidAfter"`
	// The not valid before validity bound timestamp
	NotValidBefore pulumi.StringOutput `pulumi:"notValidBefore"`
	// Certificate status
	Status pulumi.StringOutput `pulumi:"status"`
	// The alternative domain names of the certificate
	SubjectAlternativeNames pulumi.StringArrayOutput `pulumi:"subjectAlternativeNames"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LbId == nil {
		return nil, errors.New("invalid value for required argument 'LbId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("scaleway:index/loadbalancerCertificate:LoadbalancerCertificate"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Certificate
	err := ctx.RegisterResource("scaleway:loadbalancers/certificate:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificate gets an existing Certificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("scaleway:loadbalancers/certificate:Certificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Certificate resources.
type certificateState struct {
	// Main domain of the certificate
	CommonName *string `pulumi:"commonName"`
	// The custom type certificate type configuration
	CustomCertificate *CertificateCustomCertificate `pulumi:"customCertificate"`
	// The identifier (SHA-1) of the certificate
	Fingerprint *string `pulumi:"fingerprint"`
	// The load-balancer ID
	LbId *string `pulumi:"lbId"`
	// The Let's Encrypt type certificate configuration
	Letsencrypt *CertificateLetsencrypt `pulumi:"letsencrypt"`
	// The name of the load-balancer certificate
	Name *string `pulumi:"name"`
	// The not valid after validity bound timestamp
	NotValidAfter *string `pulumi:"notValidAfter"`
	// The not valid before validity bound timestamp
	NotValidBefore *string `pulumi:"notValidBefore"`
	// Certificate status
	Status *string `pulumi:"status"`
	// The alternative domain names of the certificate
	SubjectAlternativeNames []string `pulumi:"subjectAlternativeNames"`
}

type CertificateState struct {
	// Main domain of the certificate
	CommonName pulumi.StringPtrInput
	// The custom type certificate type configuration
	CustomCertificate CertificateCustomCertificatePtrInput
	// The identifier (SHA-1) of the certificate
	Fingerprint pulumi.StringPtrInput
	// The load-balancer ID
	LbId pulumi.StringPtrInput
	// The Let's Encrypt type certificate configuration
	Letsencrypt CertificateLetsencryptPtrInput
	// The name of the load-balancer certificate
	Name pulumi.StringPtrInput
	// The not valid after validity bound timestamp
	NotValidAfter pulumi.StringPtrInput
	// The not valid before validity bound timestamp
	NotValidBefore pulumi.StringPtrInput
	// Certificate status
	Status pulumi.StringPtrInput
	// The alternative domain names of the certificate
	SubjectAlternativeNames pulumi.StringArrayInput
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	// The custom type certificate type configuration
	CustomCertificate *CertificateCustomCertificate `pulumi:"customCertificate"`
	// The load-balancer ID
	LbId string `pulumi:"lbId"`
	// The Let's Encrypt type certificate configuration
	Letsencrypt *CertificateLetsencrypt `pulumi:"letsencrypt"`
	// The name of the load-balancer certificate
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// The custom type certificate type configuration
	CustomCertificate CertificateCustomCertificatePtrInput
	// The load-balancer ID
	LbId pulumi.StringInput
	// The Let's Encrypt type certificate configuration
	Letsencrypt CertificateLetsencryptPtrInput
	// The name of the load-balancer certificate
	Name pulumi.StringPtrInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}

type CertificateInput interface {
	pulumi.Input

	ToCertificateOutput() CertificateOutput
	ToCertificateOutputWithContext(ctx context.Context) CertificateOutput
}

func (*Certificate) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (i *Certificate) ToCertificateOutput() CertificateOutput {
	return i.ToCertificateOutputWithContext(context.Background())
}

func (i *Certificate) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOutput)
}

// CertificateArrayInput is an input type that accepts CertificateArray and CertificateArrayOutput values.
// You can construct a concrete instance of `CertificateArrayInput` via:
//
//	CertificateArray{ CertificateArgs{...} }
type CertificateArrayInput interface {
	pulumi.Input

	ToCertificateArrayOutput() CertificateArrayOutput
	ToCertificateArrayOutputWithContext(context.Context) CertificateArrayOutput
}

type CertificateArray []CertificateInput

func (CertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Certificate)(nil)).Elem()
}

func (i CertificateArray) ToCertificateArrayOutput() CertificateArrayOutput {
	return i.ToCertificateArrayOutputWithContext(context.Background())
}

func (i CertificateArray) ToCertificateArrayOutputWithContext(ctx context.Context) CertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateArrayOutput)
}

// CertificateMapInput is an input type that accepts CertificateMap and CertificateMapOutput values.
// You can construct a concrete instance of `CertificateMapInput` via:
//
//	CertificateMap{ "key": CertificateArgs{...} }
type CertificateMapInput interface {
	pulumi.Input

	ToCertificateMapOutput() CertificateMapOutput
	ToCertificateMapOutputWithContext(context.Context) CertificateMapOutput
}

type CertificateMap map[string]CertificateInput

func (CertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Certificate)(nil)).Elem()
}

func (i CertificateMap) ToCertificateMapOutput() CertificateMapOutput {
	return i.ToCertificateMapOutputWithContext(context.Background())
}

func (i CertificateMap) ToCertificateMapOutputWithContext(ctx context.Context) CertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateMapOutput)
}

type CertificateOutput struct{ *pulumi.OutputState }

func (CertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (o CertificateOutput) ToCertificateOutput() CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return o
}

// Main domain of the certificate
func (o CertificateOutput) CommonName() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.CommonName }).(pulumi.StringOutput)
}

// The custom type certificate type configuration
func (o CertificateOutput) CustomCertificate() CertificateCustomCertificatePtrOutput {
	return o.ApplyT(func(v *Certificate) CertificateCustomCertificatePtrOutput { return v.CustomCertificate }).(CertificateCustomCertificatePtrOutput)
}

// The identifier (SHA-1) of the certificate
func (o CertificateOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// The load-balancer ID
func (o CertificateOutput) LbId() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.LbId }).(pulumi.StringOutput)
}

// The Let's Encrypt type certificate configuration
func (o CertificateOutput) Letsencrypt() CertificateLetsencryptPtrOutput {
	return o.ApplyT(func(v *Certificate) CertificateLetsencryptPtrOutput { return v.Letsencrypt }).(CertificateLetsencryptPtrOutput)
}

// The name of the load-balancer certificate
func (o CertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The not valid after validity bound timestamp
func (o CertificateOutput) NotValidAfter() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.NotValidAfter }).(pulumi.StringOutput)
}

// The not valid before validity bound timestamp
func (o CertificateOutput) NotValidBefore() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.NotValidBefore }).(pulumi.StringOutput)
}

// Certificate status
func (o CertificateOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The alternative domain names of the certificate
func (o CertificateOutput) SubjectAlternativeNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringArrayOutput { return v.SubjectAlternativeNames }).(pulumi.StringArrayOutput)
}

type CertificateArrayOutput struct{ *pulumi.OutputState }

func (CertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Certificate)(nil)).Elem()
}

func (o CertificateArrayOutput) ToCertificateArrayOutput() CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) ToCertificateArrayOutputWithContext(ctx context.Context) CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) Index(i pulumi.IntInput) CertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Certificate {
		return vs[0].([]*Certificate)[vs[1].(int)]
	}).(CertificateOutput)
}

type CertificateMapOutput struct{ *pulumi.OutputState }

func (CertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Certificate)(nil)).Elem()
}

func (o CertificateMapOutput) ToCertificateMapOutput() CertificateMapOutput {
	return o
}

func (o CertificateMapOutput) ToCertificateMapOutputWithContext(ctx context.Context) CertificateMapOutput {
	return o
}

func (o CertificateMapOutput) MapIndex(k pulumi.StringInput) CertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Certificate {
		return vs[0].(map[string]*Certificate)[vs[1].(string)]
	}).(CertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateInput)(nil)).Elem(), &Certificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateArrayInput)(nil)).Elem(), CertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateMapInput)(nil)).Elem(), CertificateMap{})
	pulumi.RegisterOutputType(CertificateOutput{})
	pulumi.RegisterOutputType(CertificateArrayOutput{})
	pulumi.RegisterOutputType(CertificateMapOutput{})
}
