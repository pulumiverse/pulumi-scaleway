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
	"github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfbridge"
	shimv1 "github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfshim/sdk-v1"
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

// boolRef returns a reference to the bool argument.
func boolRef(b bool) *bool {
	return &b
}

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

func refProviderLicense(license tfbridge.TFProviderLicense) *tfbridge.TFProviderLicense {
	return &license
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv1.NewProvider(scaleway.Provider().(*schema.Provider))
	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                 p,
		Name:              "scaleway",
		Description:       "A Pulumi package for creating and managing scaleway cloud resources.",
		Keywords:          []string{"pulumi", "scaleway"},
		TFProviderLicense: refProviderLicense(tfbridge.MITLicenseType),
		License:           "Apache-2.0",
		Homepage:          "https://pulumi.io",
		Repository:        "https://github.com/jaxxstorm/pulumi-scaleway",
		PluginDownloadURL: "https://bintray.com/jaxxstorm/pulumi/download_file?file_path=",
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
		Resources: map[string]*tfbridge.ResourceInfo{
			"scaleway_account_ssh_key":               {Tok: scalewayResource(scalewayMod, "AccountSshKey")},
			"scaleway_baremetal_server":              {Tok: scalewayResource(scalewayMod, "BaremetalServer")},
			"scaleway_instance_ip":                   {Tok: scalewayResource(scalewayMod, "InstanceIP")},
			"scaleway_instance_ip_reverse_dns":       {Tok: scalewayResource(scalewayMod, "InstanceIPReverseDNS")},
			"scaleway_instance_placement_group":      {Tok: scalewayResource(scalewayMod, "InstancePlacementGroup")},
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
			"scaleway_lb_ip_beta":                    {Tok: scalewayResource(scalewayMod, "LoadbalancerIPBeta")},
			"scaleway_object_bucket":                 {Tok: scalewayResource(scalewayMod, "ObjectBucket")},
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
			"scaleway_baremetal_offer": {
				Tok: scalewayDataSource(scalewayMod, "getBaremetalOffer"),
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
			"scaleway_lb_ip_beta":              {Tok: scalewayDataSource(scalewayMod, "getLoadbalancerIPBeta")},
			"scaleway_marketplace_image_beta":  {Tok: scalewayDataSource(scalewayMod, "getMarketplaceImageBeta")},
			"scaleway_registry_image_beta":     {Tok: scalewayDataSource(scalewayMod, "getRegistryImageBeta")},
			"scaleway_registry_namespace_beta": {Tok: scalewayDataSource(scalewayMod, "getRegistryNamespaceBeta")},
			"scaleway_security_group":          {Tok: scalewayDataSource(scalewayMod, "getSecurityGroup")},
			"scaleway_volume":                  {Tok: scalewayDataSource(scalewayMod, "getVolume")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
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

	prov.SetAutonaming(255, "-")

	return prov
}
