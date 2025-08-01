// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about an IOT Hub.
func LookupHub(ctx *pulumi.Context, args *LookupHubArgs, opts ...pulumi.InvokeOption) (*LookupHubResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupHubResult
	err := ctx.Invoke("scaleway:iot/getHub:getHub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getHub.
type LookupHubArgs struct {
	// The Hub ID.
	// Only one of the `name` and `hubId` should be specified.
	HubId *string `pulumi:"hubId"`
	// The name of the Hub.
	// Only one of the `name` and `hubId` should be specified.
	Name *string `pulumi:"name"`
	// The ID of the project the hub is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region in which the hub exists.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getHub.
type LookupHubResult struct {
	ConnectedDeviceCount   int     `pulumi:"connectedDeviceCount"`
	CreatedAt              string  `pulumi:"createdAt"`
	DeviceAutoProvisioning bool    `pulumi:"deviceAutoProvisioning"`
	DeviceCount            int     `pulumi:"deviceCount"`
	DisableEvents          bool    `pulumi:"disableEvents"`
	Enabled                bool    `pulumi:"enabled"`
	Endpoint               string  `pulumi:"endpoint"`
	EventsTopicPrefix      string  `pulumi:"eventsTopicPrefix"`
	HubCa                  string  `pulumi:"hubCa"`
	HubCaChallenge         string  `pulumi:"hubCaChallenge"`
	HubId                  *string `pulumi:"hubId"`
	// The provider-assigned unique ID for this managed resource.
	Id             string  `pulumi:"id"`
	MqttCa         string  `pulumi:"mqttCa"`
	MqttCaUrl      string  `pulumi:"mqttCaUrl"`
	Name           *string `pulumi:"name"`
	OrganizationId string  `pulumi:"organizationId"`
	ProductPlan    string  `pulumi:"productPlan"`
	ProjectId      *string `pulumi:"projectId"`
	Region         *string `pulumi:"region"`
	Status         string  `pulumi:"status"`
	UpdatedAt      string  `pulumi:"updatedAt"`
}

func LookupHubOutput(ctx *pulumi.Context, args LookupHubOutputArgs, opts ...pulumi.InvokeOption) LookupHubResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupHubResultOutput, error) {
			args := v.(LookupHubArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway:iot/getHub:getHub", args, LookupHubResultOutput{}, options).(LookupHubResultOutput), nil
		}).(LookupHubResultOutput)
}

// A collection of arguments for invoking getHub.
type LookupHubOutputArgs struct {
	// The Hub ID.
	// Only one of the `name` and `hubId` should be specified.
	HubId pulumi.StringPtrInput `pulumi:"hubId"`
	// The name of the Hub.
	// Only one of the `name` and `hubId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the project the hub is associated with.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// `region`) The region in which the hub exists.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupHubOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHubArgs)(nil)).Elem()
}

// A collection of values returned by getHub.
type LookupHubResultOutput struct{ *pulumi.OutputState }

func (LookupHubResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHubResult)(nil)).Elem()
}

func (o LookupHubResultOutput) ToLookupHubResultOutput() LookupHubResultOutput {
	return o
}

func (o LookupHubResultOutput) ToLookupHubResultOutputWithContext(ctx context.Context) LookupHubResultOutput {
	return o
}

func (o LookupHubResultOutput) ConnectedDeviceCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupHubResult) int { return v.ConnectedDeviceCount }).(pulumi.IntOutput)
}

func (o LookupHubResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHubResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupHubResultOutput) DeviceAutoProvisioning() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHubResult) bool { return v.DeviceAutoProvisioning }).(pulumi.BoolOutput)
}

func (o LookupHubResultOutput) DeviceCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupHubResult) int { return v.DeviceCount }).(pulumi.IntOutput)
}

func (o LookupHubResultOutput) DisableEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHubResult) bool { return v.DisableEvents }).(pulumi.BoolOutput)
}

func (o LookupHubResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHubResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupHubResultOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHubResult) string { return v.Endpoint }).(pulumi.StringOutput)
}

func (o LookupHubResultOutput) EventsTopicPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHubResult) string { return v.EventsTopicPrefix }).(pulumi.StringOutput)
}

func (o LookupHubResultOutput) HubCa() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHubResult) string { return v.HubCa }).(pulumi.StringOutput)
}

func (o LookupHubResultOutput) HubCaChallenge() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHubResult) string { return v.HubCaChallenge }).(pulumi.StringOutput)
}

func (o LookupHubResultOutput) HubId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHubResult) *string { return v.HubId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupHubResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHubResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupHubResultOutput) MqttCa() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHubResult) string { return v.MqttCa }).(pulumi.StringOutput)
}

func (o LookupHubResultOutput) MqttCaUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHubResult) string { return v.MqttCaUrl }).(pulumi.StringOutput)
}

func (o LookupHubResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHubResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupHubResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHubResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupHubResultOutput) ProductPlan() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHubResult) string { return v.ProductPlan }).(pulumi.StringOutput)
}

func (o LookupHubResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHubResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o LookupHubResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHubResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupHubResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHubResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupHubResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHubResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHubResultOutput{})
}
