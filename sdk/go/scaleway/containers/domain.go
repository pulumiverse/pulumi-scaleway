// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package containers

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// The `containers.Domain` resource allows you to create and manage domain name bindings for Scaleway [Serverless Containers](https://www.scaleway.com/en/docs/serverless/containers/).
//
// Refer to the Containers domain [documentation](https://www.scaleway.com/en/docs/serverless-containers/how-to/add-a-custom-domain-to-a-container/) and the [API documentation](https://www.scaleway.com/en/developers/api/serverless-containers/#path-domains-list-all-domain-name-bindings) for more information.
//
// ## Example Usage
//
// The commands below shows how to bind a custom domain name to a container.
//
// ### Simple
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
//			app, err := containers.NewContainer(ctx, "app", nil)
//			if err != nil {
//				return err
//			}
//			_, err = containers.NewDomain(ctx, "app", &containers.DomainArgs{
//				ContainerId: app.ID(),
//				Hostname:    pulumi.String("container.domain.tld"),
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
// ### Complete example with domain
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/containers"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/domain"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := containers.NewNamespace(ctx, "main", &containers.NamespaceArgs{
//				Name:        pulumi.String("my-ns-test"),
//				Description: pulumi.String("test container"),
//			})
//			if err != nil {
//				return err
//			}
//			app, err := containers.NewContainer(ctx, "app", &containers.ContainerArgs{
//				Name:        pulumi.String("app"),
//				NamespaceId: main.ID(),
//				RegistryImage: main.RegistryEndpoint.ApplyT(func(registryEndpoint string) (string, error) {
//					return fmt.Sprintf("%v/nginx:alpine", registryEndpoint), nil
//				}).(pulumi.StringOutput),
//				Port:           pulumi.Int(80),
//				CpuLimit:       pulumi.Int(140),
//				MemoryLimit:    pulumi.Int(256),
//				MinScale:       pulumi.Int(1),
//				MaxScale:       pulumi.Int(1),
//				Timeout:        pulumi.Int(600),
//				MaxConcurrency: pulumi.Int(80),
//				Privacy:        pulumi.String("public"),
//				Protocol:       pulumi.String("http1"),
//				Deploy:         pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			appRecord, err := domain.NewRecord(ctx, "app", &domain.RecordArgs{
//				DnsZone: pulumi.String("domain.tld"),
//				Name:    pulumi.String("subdomain"),
//				Type:    pulumi.String("CNAME"),
//				Data: app.DomainName.ApplyT(func(domainName string) (string, error) {
//					return fmt.Sprintf("%v.", domainName), nil
//				}).(pulumi.StringOutput),
//				Ttl: pulumi.Int(3600),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = containers.NewDomain(ctx, "app", &containers.DomainArgs{
//				ContainerId: app.ID(),
//				Hostname: pulumi.All(appRecord.Name, appRecord.DnsZone).ApplyT(func(_args []interface{}) (string, error) {
//					name := _args[0].(string)
//					dnsZone := _args[1].(string)
//					return fmt.Sprintf("%v.%v", name, dnsZone), nil
//				}).(pulumi.StringOutput),
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
// Container domain binding can be imported using `{region}/{id}`, as shown below:
//
// bash
//
// ```sh
// $ pulumi import scaleway:containers/domain:Domain main fr-par/11111111-1111-1111-1111-111111111111
// ```
type Domain struct {
	pulumi.CustomResourceState

	// The unique identifier of the container.
	ContainerId pulumi.StringOutput `pulumi:"containerId"`
	// The hostname with a CNAME record.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// `region`) The region in which the container exists.
	Region pulumi.StringOutput `pulumi:"region"`
	// The URL used to query the container.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContainerId == nil {
		return nil, errors.New("invalid value for required argument 'ContainerId'")
	}
	if args.Hostname == nil {
		return nil, errors.New("invalid value for required argument 'Hostname'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("scaleway:index/containerDomain:ContainerDomain"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Domain
	err := ctx.RegisterResource("scaleway:containers/domain:Domain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainState, opts ...pulumi.ResourceOption) (*Domain, error) {
	var resource Domain
	err := ctx.ReadResource("scaleway:containers/domain:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Domain resources.
type domainState struct {
	// The unique identifier of the container.
	ContainerId *string `pulumi:"containerId"`
	// The hostname with a CNAME record.
	Hostname *string `pulumi:"hostname"`
	// `region`) The region in which the container exists.
	Region *string `pulumi:"region"`
	// The URL used to query the container.
	Url *string `pulumi:"url"`
}

type DomainState struct {
	// The unique identifier of the container.
	ContainerId pulumi.StringPtrInput
	// The hostname with a CNAME record.
	Hostname pulumi.StringPtrInput
	// `region`) The region in which the container exists.
	Region pulumi.StringPtrInput
	// The URL used to query the container.
	Url pulumi.StringPtrInput
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	// The unique identifier of the container.
	ContainerId string `pulumi:"containerId"`
	// The hostname with a CNAME record.
	Hostname string `pulumi:"hostname"`
	// `region`) The region in which the container exists.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// The unique identifier of the container.
	ContainerId pulumi.StringInput
	// The hostname with a CNAME record.
	Hostname pulumi.StringInput
	// `region`) The region in which the container exists.
	Region pulumi.StringPtrInput
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainArgs)(nil)).Elem()
}

type DomainInput interface {
	pulumi.Input

	ToDomainOutput() DomainOutput
	ToDomainOutputWithContext(ctx context.Context) DomainOutput
}

func (*Domain) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (i *Domain) ToDomainOutput() DomainOutput {
	return i.ToDomainOutputWithContext(context.Background())
}

func (i *Domain) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainOutput)
}

// DomainArrayInput is an input type that accepts DomainArray and DomainArrayOutput values.
// You can construct a concrete instance of `DomainArrayInput` via:
//
//	DomainArray{ DomainArgs{...} }
type DomainArrayInput interface {
	pulumi.Input

	ToDomainArrayOutput() DomainArrayOutput
	ToDomainArrayOutputWithContext(context.Context) DomainArrayOutput
}

type DomainArray []DomainInput

func (DomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Domain)(nil)).Elem()
}

func (i DomainArray) ToDomainArrayOutput() DomainArrayOutput {
	return i.ToDomainArrayOutputWithContext(context.Background())
}

func (i DomainArray) ToDomainArrayOutputWithContext(ctx context.Context) DomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainArrayOutput)
}

// DomainMapInput is an input type that accepts DomainMap and DomainMapOutput values.
// You can construct a concrete instance of `DomainMapInput` via:
//
//	DomainMap{ "key": DomainArgs{...} }
type DomainMapInput interface {
	pulumi.Input

	ToDomainMapOutput() DomainMapOutput
	ToDomainMapOutputWithContext(context.Context) DomainMapOutput
}

type DomainMap map[string]DomainInput

func (DomainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Domain)(nil)).Elem()
}

func (i DomainMap) ToDomainMapOutput() DomainMapOutput {
	return i.ToDomainMapOutputWithContext(context.Background())
}

func (i DomainMap) ToDomainMapOutputWithContext(ctx context.Context) DomainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainMapOutput)
}

type DomainOutput struct{ *pulumi.OutputState }

func (DomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (o DomainOutput) ToDomainOutput() DomainOutput {
	return o
}

func (o DomainOutput) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return o
}

// The unique identifier of the container.
func (o DomainOutput) ContainerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.ContainerId }).(pulumi.StringOutput)
}

// The hostname with a CNAME record.
func (o DomainOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// `region`) The region in which the container exists.
func (o DomainOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The URL used to query the container.
func (o DomainOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type DomainArrayOutput struct{ *pulumi.OutputState }

func (DomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Domain)(nil)).Elem()
}

func (o DomainArrayOutput) ToDomainArrayOutput() DomainArrayOutput {
	return o
}

func (o DomainArrayOutput) ToDomainArrayOutputWithContext(ctx context.Context) DomainArrayOutput {
	return o
}

func (o DomainArrayOutput) Index(i pulumi.IntInput) DomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Domain {
		return vs[0].([]*Domain)[vs[1].(int)]
	}).(DomainOutput)
}

type DomainMapOutput struct{ *pulumi.OutputState }

func (DomainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Domain)(nil)).Elem()
}

func (o DomainMapOutput) ToDomainMapOutput() DomainMapOutput {
	return o
}

func (o DomainMapOutput) ToDomainMapOutputWithContext(ctx context.Context) DomainMapOutput {
	return o
}

func (o DomainMapOutput) MapIndex(k pulumi.StringInput) DomainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Domain {
		return vs[0].(map[string]*Domain)[vs[1].(string)]
	}).(DomainOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainInput)(nil)).Elem(), &Domain{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainArrayInput)(nil)).Elem(), DomainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainMapInput)(nil)).Elem(), DomainMap{})
	pulumi.RegisterOutputType(DomainOutput{})
	pulumi.RegisterOutputType(DomainArrayOutput{})
	pulumi.RegisterOutputType(DomainMapOutput{})
}
