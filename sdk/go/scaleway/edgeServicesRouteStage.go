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

// Creates and manages Scaleway Edge Services Route Stages.
//
// ## Example Usage
//
// ### Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := scaleway.NewEdgeServicesRouteStage(ctx, "main", &scaleway.EdgeServicesRouteStageArgs{
//				PipelineId: pulumi.Any(mainScalewayEdgeServicesPipeline.Id),
//				WafStageId: pulumi.Any(waf.Id),
//				Rules: scaleway.EdgeServicesRouteStageRuleArray{
//					&scaleway.EdgeServicesRouteStageRuleArgs{
//						BackendStageId: pulumi.Any(backend.Id),
//						RuleHttpMatch: &scaleway.EdgeServicesRouteStageRuleRuleHttpMatchArgs{
//							MethodFilters: pulumi.StringArray{
//								pulumi.String("get"),
//								pulumi.String("post"),
//							},
//							PathFilter: &scaleway.EdgeServicesRouteStageRuleRuleHttpMatchPathFilterArgs{
//								PathFilterType: pulumi.String("regex"),
//								Value:          pulumi.String(".*"),
//							},
//						},
//					},
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
// Route stages can be imported using the `{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/edgeServicesRouteStage:EdgeServicesRouteStage basic 11111111-1111-1111-1111-111111111111
// ```
type EdgeServicesRouteStage struct {
	pulumi.CustomResourceState

	// The date and time of the creation of the route stage.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The ID of the pipeline.
	PipelineId pulumi.StringOutput `pulumi:"pipelineId"`
	// `projectId`) The ID of the project the route stage is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The list of rules to be checked against every HTTP request. The first matching rule will forward the request to its specified backend stage. If no rules are matched, the request is forwarded to the WAF stage defined by `wafStageId`.
	Rules EdgeServicesRouteStageRuleArrayOutput `pulumi:"rules"`
	// The date and time of the last update of the route stage.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// The ID of the WAF stage HTTP requests should be forwarded to when no rules are matched.
	WafStageId pulumi.StringPtrOutput `pulumi:"wafStageId"`
}

// NewEdgeServicesRouteStage registers a new resource with the given unique name, arguments, and options.
func NewEdgeServicesRouteStage(ctx *pulumi.Context,
	name string, args *EdgeServicesRouteStageArgs, opts ...pulumi.ResourceOption) (*EdgeServicesRouteStage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PipelineId == nil {
		return nil, errors.New("invalid value for required argument 'PipelineId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EdgeServicesRouteStage
	err := ctx.RegisterResource("scaleway:index/edgeServicesRouteStage:EdgeServicesRouteStage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEdgeServicesRouteStage gets an existing EdgeServicesRouteStage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEdgeServicesRouteStage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EdgeServicesRouteStageState, opts ...pulumi.ResourceOption) (*EdgeServicesRouteStage, error) {
	var resource EdgeServicesRouteStage
	err := ctx.ReadResource("scaleway:index/edgeServicesRouteStage:EdgeServicesRouteStage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EdgeServicesRouteStage resources.
type edgeServicesRouteStageState struct {
	// The date and time of the creation of the route stage.
	CreatedAt *string `pulumi:"createdAt"`
	// The ID of the pipeline.
	PipelineId *string `pulumi:"pipelineId"`
	// `projectId`) The ID of the project the route stage is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The list of rules to be checked against every HTTP request. The first matching rule will forward the request to its specified backend stage. If no rules are matched, the request is forwarded to the WAF stage defined by `wafStageId`.
	Rules []EdgeServicesRouteStageRule `pulumi:"rules"`
	// The date and time of the last update of the route stage.
	UpdatedAt *string `pulumi:"updatedAt"`
	// The ID of the WAF stage HTTP requests should be forwarded to when no rules are matched.
	WafStageId *string `pulumi:"wafStageId"`
}

type EdgeServicesRouteStageState struct {
	// The date and time of the creation of the route stage.
	CreatedAt pulumi.StringPtrInput
	// The ID of the pipeline.
	PipelineId pulumi.StringPtrInput
	// `projectId`) The ID of the project the route stage is associated with.
	ProjectId pulumi.StringPtrInput
	// The list of rules to be checked against every HTTP request. The first matching rule will forward the request to its specified backend stage. If no rules are matched, the request is forwarded to the WAF stage defined by `wafStageId`.
	Rules EdgeServicesRouteStageRuleArrayInput
	// The date and time of the last update of the route stage.
	UpdatedAt pulumi.StringPtrInput
	// The ID of the WAF stage HTTP requests should be forwarded to when no rules are matched.
	WafStageId pulumi.StringPtrInput
}

func (EdgeServicesRouteStageState) ElementType() reflect.Type {
	return reflect.TypeOf((*edgeServicesRouteStageState)(nil)).Elem()
}

type edgeServicesRouteStageArgs struct {
	// The ID of the pipeline.
	PipelineId string `pulumi:"pipelineId"`
	// `projectId`) The ID of the project the route stage is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The list of rules to be checked against every HTTP request. The first matching rule will forward the request to its specified backend stage. If no rules are matched, the request is forwarded to the WAF stage defined by `wafStageId`.
	Rules []EdgeServicesRouteStageRule `pulumi:"rules"`
	// The ID of the WAF stage HTTP requests should be forwarded to when no rules are matched.
	WafStageId *string `pulumi:"wafStageId"`
}

// The set of arguments for constructing a EdgeServicesRouteStage resource.
type EdgeServicesRouteStageArgs struct {
	// The ID of the pipeline.
	PipelineId pulumi.StringInput
	// `projectId`) The ID of the project the route stage is associated with.
	ProjectId pulumi.StringPtrInput
	// The list of rules to be checked against every HTTP request. The first matching rule will forward the request to its specified backend stage. If no rules are matched, the request is forwarded to the WAF stage defined by `wafStageId`.
	Rules EdgeServicesRouteStageRuleArrayInput
	// The ID of the WAF stage HTTP requests should be forwarded to when no rules are matched.
	WafStageId pulumi.StringPtrInput
}

func (EdgeServicesRouteStageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*edgeServicesRouteStageArgs)(nil)).Elem()
}

type EdgeServicesRouteStageInput interface {
	pulumi.Input

	ToEdgeServicesRouteStageOutput() EdgeServicesRouteStageOutput
	ToEdgeServicesRouteStageOutputWithContext(ctx context.Context) EdgeServicesRouteStageOutput
}

func (*EdgeServicesRouteStage) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeServicesRouteStage)(nil)).Elem()
}

func (i *EdgeServicesRouteStage) ToEdgeServicesRouteStageOutput() EdgeServicesRouteStageOutput {
	return i.ToEdgeServicesRouteStageOutputWithContext(context.Background())
}

func (i *EdgeServicesRouteStage) ToEdgeServicesRouteStageOutputWithContext(ctx context.Context) EdgeServicesRouteStageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeServicesRouteStageOutput)
}

// EdgeServicesRouteStageArrayInput is an input type that accepts EdgeServicesRouteStageArray and EdgeServicesRouteStageArrayOutput values.
// You can construct a concrete instance of `EdgeServicesRouteStageArrayInput` via:
//
//	EdgeServicesRouteStageArray{ EdgeServicesRouteStageArgs{...} }
type EdgeServicesRouteStageArrayInput interface {
	pulumi.Input

	ToEdgeServicesRouteStageArrayOutput() EdgeServicesRouteStageArrayOutput
	ToEdgeServicesRouteStageArrayOutputWithContext(context.Context) EdgeServicesRouteStageArrayOutput
}

type EdgeServicesRouteStageArray []EdgeServicesRouteStageInput

func (EdgeServicesRouteStageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EdgeServicesRouteStage)(nil)).Elem()
}

func (i EdgeServicesRouteStageArray) ToEdgeServicesRouteStageArrayOutput() EdgeServicesRouteStageArrayOutput {
	return i.ToEdgeServicesRouteStageArrayOutputWithContext(context.Background())
}

func (i EdgeServicesRouteStageArray) ToEdgeServicesRouteStageArrayOutputWithContext(ctx context.Context) EdgeServicesRouteStageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeServicesRouteStageArrayOutput)
}

// EdgeServicesRouteStageMapInput is an input type that accepts EdgeServicesRouteStageMap and EdgeServicesRouteStageMapOutput values.
// You can construct a concrete instance of `EdgeServicesRouteStageMapInput` via:
//
//	EdgeServicesRouteStageMap{ "key": EdgeServicesRouteStageArgs{...} }
type EdgeServicesRouteStageMapInput interface {
	pulumi.Input

	ToEdgeServicesRouteStageMapOutput() EdgeServicesRouteStageMapOutput
	ToEdgeServicesRouteStageMapOutputWithContext(context.Context) EdgeServicesRouteStageMapOutput
}

type EdgeServicesRouteStageMap map[string]EdgeServicesRouteStageInput

func (EdgeServicesRouteStageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EdgeServicesRouteStage)(nil)).Elem()
}

func (i EdgeServicesRouteStageMap) ToEdgeServicesRouteStageMapOutput() EdgeServicesRouteStageMapOutput {
	return i.ToEdgeServicesRouteStageMapOutputWithContext(context.Background())
}

func (i EdgeServicesRouteStageMap) ToEdgeServicesRouteStageMapOutputWithContext(ctx context.Context) EdgeServicesRouteStageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeServicesRouteStageMapOutput)
}

type EdgeServicesRouteStageOutput struct{ *pulumi.OutputState }

func (EdgeServicesRouteStageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeServicesRouteStage)(nil)).Elem()
}

func (o EdgeServicesRouteStageOutput) ToEdgeServicesRouteStageOutput() EdgeServicesRouteStageOutput {
	return o
}

func (o EdgeServicesRouteStageOutput) ToEdgeServicesRouteStageOutputWithContext(ctx context.Context) EdgeServicesRouteStageOutput {
	return o
}

// The date and time of the creation of the route stage.
func (o EdgeServicesRouteStageOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeServicesRouteStage) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The ID of the pipeline.
func (o EdgeServicesRouteStageOutput) PipelineId() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeServicesRouteStage) pulumi.StringOutput { return v.PipelineId }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the route stage is associated with.
func (o EdgeServicesRouteStageOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeServicesRouteStage) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The list of rules to be checked against every HTTP request. The first matching rule will forward the request to its specified backend stage. If no rules are matched, the request is forwarded to the WAF stage defined by `wafStageId`.
func (o EdgeServicesRouteStageOutput) Rules() EdgeServicesRouteStageRuleArrayOutput {
	return o.ApplyT(func(v *EdgeServicesRouteStage) EdgeServicesRouteStageRuleArrayOutput { return v.Rules }).(EdgeServicesRouteStageRuleArrayOutput)
}

// The date and time of the last update of the route stage.
func (o EdgeServicesRouteStageOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeServicesRouteStage) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// The ID of the WAF stage HTTP requests should be forwarded to when no rules are matched.
func (o EdgeServicesRouteStageOutput) WafStageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeServicesRouteStage) pulumi.StringPtrOutput { return v.WafStageId }).(pulumi.StringPtrOutput)
}

type EdgeServicesRouteStageArrayOutput struct{ *pulumi.OutputState }

func (EdgeServicesRouteStageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EdgeServicesRouteStage)(nil)).Elem()
}

func (o EdgeServicesRouteStageArrayOutput) ToEdgeServicesRouteStageArrayOutput() EdgeServicesRouteStageArrayOutput {
	return o
}

func (o EdgeServicesRouteStageArrayOutput) ToEdgeServicesRouteStageArrayOutputWithContext(ctx context.Context) EdgeServicesRouteStageArrayOutput {
	return o
}

func (o EdgeServicesRouteStageArrayOutput) Index(i pulumi.IntInput) EdgeServicesRouteStageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EdgeServicesRouteStage {
		return vs[0].([]*EdgeServicesRouteStage)[vs[1].(int)]
	}).(EdgeServicesRouteStageOutput)
}

type EdgeServicesRouteStageMapOutput struct{ *pulumi.OutputState }

func (EdgeServicesRouteStageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EdgeServicesRouteStage)(nil)).Elem()
}

func (o EdgeServicesRouteStageMapOutput) ToEdgeServicesRouteStageMapOutput() EdgeServicesRouteStageMapOutput {
	return o
}

func (o EdgeServicesRouteStageMapOutput) ToEdgeServicesRouteStageMapOutputWithContext(ctx context.Context) EdgeServicesRouteStageMapOutput {
	return o
}

func (o EdgeServicesRouteStageMapOutput) MapIndex(k pulumi.StringInput) EdgeServicesRouteStageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EdgeServicesRouteStage {
		return vs[0].(map[string]*EdgeServicesRouteStage)[vs[1].(string)]
	}).(EdgeServicesRouteStageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeServicesRouteStageInput)(nil)).Elem(), &EdgeServicesRouteStage{})
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeServicesRouteStageArrayInput)(nil)).Elem(), EdgeServicesRouteStageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeServicesRouteStageMapInput)(nil)).Elem(), EdgeServicesRouteStageMap{})
	pulumi.RegisterOutputType(EdgeServicesRouteStageOutput{})
	pulumi.RegisterOutputType(EdgeServicesRouteStageArrayOutput{})
	pulumi.RegisterOutputType(EdgeServicesRouteStageMapOutput{})
}
