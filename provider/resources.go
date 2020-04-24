// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package scaleway

import (
	"unicode"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfbridge"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
	"github.com/terraform-providers/terraform-provider-scaleway/scaleway"
)

// all of the token components used below.
const (
	// packages:
	scalewayPkg = "scaleway"
	// modules:
	scalewayMod = "index" // the y module
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(scalewayPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the data source's
// first character.
func makeDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeMember(mod+"/"+fn, res)
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

// boolRef returns a reference to the bool argument.
func boolRef(b bool) *bool {
	return &b
}

// stringValue gets a string value from a property map if present, else ""
func stringValue(vars resource.PropertyMap, prop resource.PropertyKey) string {
	val, ok := vars[prop]
	if ok && val.IsString() {
		return val.StringValue()
	}
	return ""
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c *terraform.ResourceConfig) error {
	return nil
}

// managedByPulumi is a default used for some managed resources, in the absence of something more meaningful.
var managedByPulumi = &tfbridge.DefaultInfo{Value: "Managed by Pulumi"}

// scalewayMember manufactures a type token for the Scaleway package and the given module and type.
func scalewayMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(scalewayPkg + ":" + mod + ":" + mem)
}

// scalewayType manufactures a type token for the Scaleway package and the given module and type.
func scalewayType(mod string, typ string) tokens.Type {
	return tokens.Type(scalewayMember(mod, typ))
}

// scalewayDataSource manufactures a standard resource token given a module and resource name.
// It automatically uses the Scaleway package and names the file by simply lower casing the data
// source's first character.
func scalewayDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return scalewayMember(mod+"/"+fn, res)
}

// scalewayResource manufactures a standard resource token given a module and resource name.
// It automatically uses the Scaleway package and names the file by simply lower casing the resource's
// first character.
func scalewayResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return scalewayType(mod+"/"+fn, res)
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := scaleway.Provider().(*schema.Provider)

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "scaleway",
		Description: "A Pulumi package for creating and managing scaleway cloud resources.",
		Keywords:    []string{"pulumi", "scaleway"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/jaxxstorm/pulumi-scaleway",
		Config: map[string]*tfbridge.SchemaInfo{
			"access_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"SCW_ACCESS_KEY"},
				},
			},
			"secret_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"SCW_SECRET_KEY"},
				},
			},
			"organization_id": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"SCW_DEFAULT_ORGANIZATION_ID"},
				},
			},
			"region": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"SCW_DEFAULT_REGION"},
				},
			},
			"zone": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"SCW_DEFAULT_ZONE"},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"scaleway_account_ssh_key":               {Tok: scalewayResource(scalewayMod, "AccountSshKey")},
			"scaleway_baremetal_server_beta":         {Tok: scalewayResource(scalewayMod, "BaremetalServerBeta")},
			"scaleway_bucket":                        {Tok: scalewayResource(scalewayMod, "Bucket")},
			"scaleway_instance_ip":                   {Tok: scalewayResource(scalewayMod, "InstanceIP")},
			"scaleway_instance_ip_reverse_dns":       {Tok: scalewayResource(scalewayMod, "InstanceIPReverseDNS")},
			"scaleway_instance_security_group":       {Tok: scalewayResource(scalewayMod, "InstanceSecurityGroup")},
			"scaleway_instance_security_group_rules": {Tok: scalewayResource(scalewayMod, "InstanceSecurityGroupRules")},
			"scaleway_instance_server":               {Tok: scalewayResource(scalewayMod, "InstanceServer")},
			"scaleway_instance_volume":               {Tok: scalewayResource(scalewayMod, "InstanceVolume")},
			"scaleway_ip":                            {Tok: scalewayResource(scalewayMod, "IP")},
			"scaleway_ip_reverse_dns":                {Tok: scalewayResource(scalewayMod, "IPReverseDNS")},
			"scaleway_k8s_cluster_beta":              {Tok: scalewayResource(scalewayMod, "KubernetesClusterBeta")},
			"scaleway_k8s_pool_beta":                 {Tok: scalewayResource(scalewayMod, "KubernetesNodePoolBeta")},
			"scaleway_lb_backend_beta":               {Tok: scalewayResource(scalewayMod, "LoadbalancerBackendBeta")},
			"scaleway_lb_beta":                       {Tok: scalewayResource(scalewayMod, "LoadbalancerBeta")},
			"scaleway_lb_certificate_beta":           {Tok: scalewayResource(scalewayMod, "LoadbalancerCertificateBeta")},
			"scaleway_lb_frontend_beta":              {Tok: scalewayResource(scalewayMod, "LoadbalancerFrontendBeta")},
			"scaleway_object_bucket":                 {Tok: scalewayResource(scalewayMod, "ObjectBucket")},
			"scaleway_instance_placement_group":      {Tok: scalewayResource(scalewayMod, "InstancePlacementGroup")},
			"scaleway_rdb_instance_beta":             {Tok: scalewayResource(scalewayMod, "DatabaseInstanceBeta")},
			"scaleway_registry_namespace_beta":       {Tok: scalewayResource(scalewayMod, "RegistryNamespaceBeta")},
			"scaleway_security_group":                {Tok: scalewayResource(scalewayMod, "SecurityGroup")},
			"scaleway_security_group_rule":           {Tok: scalewayResource(scalewayMod, "SecurityGroupRule")},
			"scaleway_server":                        {Tok: scalewayResource(scalewayMod, "Server")},
			"scaleway_ssh_key":                       {Tok: scalewayResource(scalewayMod, "SshKey")},
			"scaleway_token":                         {Tok: scalewayResource(scalewayMod, "Token")},
			"scaleway_user_data":                     {Tok: scalewayResource(scalewayMod, "UserData")},
			"scaleway_volume":                        {Tok: scalewayResource(scalewayMod, "Volume")},
			"scaleway_volume_attachment":             {Tok: scalewayResource(scalewayMod, "VolumeAttachment")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"scaleway_account_ssh_key": {Tok: scalewayDataSource(scalewayMod, "getAccountSshKey")},
			"scaleway_baremetal_offer_beta": {
				Tok: scalewayDataSource(scalewayMod, "getBaremetalOfferBeta"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"cpu": {
						MaxItemsOne: boolRef(true),
					},
				},
			},
			"scaleway_bootscript":              {Tok: scalewayDataSource(scalewayMod, "getBootscript")},
			"scaleway_image":                   {Tok: scalewayDataSource(scalewayMod, "getImage")},
			"scaleway_instance_image":          {Tok: scalewayDataSource(scalewayMod, "getInstanceImage")},
			"scaleway_instance_security_group": {Tok: scalewayDataSource(scalewayMod, "getInstanceSecurityGroup")},
			"scaleway_instance_server":         {Tok: scalewayDataSource(scalewayMod, "getInstanceServer")},
			"scaleway_instance_volume":         {Tok: scalewayDataSource(scalewayMod, "getInstanceVolume")},
			"scaleway_marketplace_image_beta":  {Tok: scalewayDataSource(scalewayMod, "getMarketplaceImageBeta")},
			"scaleway_registry_image_beta":     {Tok: scalewayDataSource(scalewayMod, "getRegistryImageBeta")},
			"scaleway_registry_namespace_beta": {Tok: scalewayDataSource(scalewayMod, "getRegistryNamespaceBeta")},
			"scaleway_security_group":          {Tok: scalewayDataSource(scalewayMod, "getSecurityGroup")},
			"scaleway_volume":                  {Tok: scalewayDataSource(scalewayMod, "getVolume")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			AsyncDataSources: true,
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^2.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			PackageName: "@jaxxstorm/pulumi-scaleway",
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=2.0.0,<3.0.0",
			},
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "2.*",
				"System.Collections.Immutable": "1.6.0",
			},
		},
	}

	// For all resources with name properties, we will add an auto-name property.  Make sure to skip those that
	// already have a name mapping entry, since those may have custom overrides set above (e.g., for length).
	const nameProperty = "name"
	for resname, res := range prov.Resources {
		if schema := p.ResourcesMap[resname]; schema != nil {
			// Only apply auto-name to input properties (Optional || Required) named `name`
			if tfs, has := schema.Schema[nameProperty]; has && (tfs.Optional || tfs.Required) {
				if _, hasfield := res.Fields[nameProperty]; !hasfield {
					if res.Fields == nil {
						res.Fields = make(map[string]*tfbridge.SchemaInfo)
					}
					res.Fields[nameProperty] = tfbridge.AutoName(nameProperty, 255)
				}
			}
		}
	}

	return prov
}
