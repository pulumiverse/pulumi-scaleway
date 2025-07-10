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
	"fmt"
	"path/filepath"
	"unicode"

	// embed is used to embed the schema files in the provider
	_ "embed"

	shim "github.com/scaleway/terraform-provider-scaleway/v2/shim"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tks "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	"github.com/pulumiverse/pulumi-scaleway/provider/pkg/version"
)

const (
	// packages:
	scalewayPkg = "scaleway"
	// modules:
	scalewayMod = "index" // the y module
	// further modules follow the grouping of the upstream TF provider
	// https://registry.terraform.io/providers/scaleway/scaleway/latest/docs
	accountMod       = "account"
	appleSiliconMod  = "applesilicon"
	autoscalingMod   = "autoscaling"
	billingMod       = "billing"
	blockMod         = "block"
	cockpitMod       = "observability"
	containersMod    = "containers"
	databasesMod     = "databases"
	domainMod        = "domain"
	elasticmetalMod  = "elasticmetal"
	functionsMod     = "functions"
	iamMod           = "iam"
	ipamMod          = "ipam"
	inferenceMod     = "inference"
	instanceMod      = "instance"
	iotMod           = "iot"
	jobMod           = "job"
	kubernetesMod    = "kubernetes"
	loadbalancersMod = "loadbalancers"
	mnqMod           = "mnq"
	mongodbMod       = "mongodb"
	objectMod        = "object"
	redisMod         = "redis"
	registryMod      = "registry"
	secretsMod       = "secrets"
	temMod           = "tem"
	vpcMod           = "network"
	webhostingMod    = "hosting"
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

//go:embed cmd/pulumi-resource-scaleway/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	p := shimv2.NewProvider(shim.NewProvider()())
	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                       p,
		Name:                    "scaleway",
		Description:             "A Pulumi package for creating and managing Scaleway cloud resources.",
		Keywords:                []string{"pulumi", "scaleway", "pulumiverse"},
		TFProviderLicense:       refProviderLicense(tfbridge.MITLicenseType),
		License:                 "Apache-2.0",
		LogoURL:                 "https://raw.githubusercontent.com/pulumiverse/pulumi-scaleway/master/assets/scaleway.png", //nolint:golint,lll
		Homepage:                "https://www.scaleway.com",
		Repository:              "https://github.com/pulumiverse/pulumi-scaleway",
		PluginDownloadURL:       "github://api.github.com/pulumiverse",
		MetadataInfo:            tfbridge.NewProviderMetadata(metadata),
		GitHubOrg:               "scaleway", // not in the terraform-providers repo
		Publisher:               "pulumiverse",
		DisplayName:             "Scaleway",
		TFProviderModuleVersion: "v2",
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
					EnvVars: []string{"SCW_ORGANIZATION_ID"},
				},
			},
			"project_id": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"SCW_DEFAULT_PROJECT_ID"},
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
			"scaleway_cockpit": {
				Tok: scalewayResource(cockpitMod, "Cockpit"),
			},
			"scaleway_container": {
				Tok: scalewayResource(containersMod, "Container"),
			},
			"scaleway_container_token": {
				Fields: map[string]*tfbridge.SchemaInfo{
					"token": {
						Name: "value",
					},
				},
			},
			"scaleway_function": {
				Tok: scalewayResource(functionsMod, "Function"),
			},
			"scaleway_function_token": {
				Fields: map[string]*tfbridge.SchemaInfo{
					"token": {
						Name: "value",
					},
				},
			},
			"scaleway_lb": {
				Tok: scalewayResource(loadbalancersMod, "LoadBalancer"),
				Docs: &tfbridge.DocInfo{
					Source: "lb.md",
				},
			},
			"scaleway_object": {
				Tok: scalewayResource(objectMod, "Item"),
			},
			"scaleway_sdb_sql_database": {
				Tok: scalewayResource(databasesMod, "ServerlessDatabase"),
				Docs: &tfbridge.DocInfo{
					Source: "sdb_sql_database.md",
				},
			},
			"scaleway_secret": {
				Tok: scalewayResource(secretsMod, "Secret"),
			},
			"scaleway_vpc": {
				Tok: scalewayResource(vpcMod, "Vpc"),
			},
			"scaleway_webhosting": {
				Tok: scalewayResource(webhostingMod, "Hosting"),
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"scaleway_availability_zones": {
				Tok: scalewayDataSource(accountMod, "getAvailabilityZones"),
			},
			"scaleway_baremetal_offer": {
				Fields: map[string]*tfbridge.SchemaInfo{
					"cpu": {
						MaxItemsOne: boolRef(true),
					},
				},
			},
			"scaleway_block_snapshot": {
				Docs: &tfbridge.DocInfo{
					Source: "bloc_snapshot.md",
				},
			},
			"scaleway_billing_consumptions": {
				Docs: &tfbridge.DocInfo{
					Source: "billing_consumption.md",
				},
			},
			"scaleway_cockpit": {
				Tok: scalewayDataSource(cockpitMod, "getInstance"),
			},
			"scaleway_config": {
				Docs: &tfbridge.DocInfo{
					// Source: "config.md", // File doesn't exist in upstream TF provider repo
					AllowMissing: true,
				},
			},
			"scaleway_container": {
				Tok: scalewayDataSource(containersMod, "getContainer"),
			},
			"scaleway_function": {
				Tok: scalewayDataSource(functionsMod, "getFunction"),
			},
			"scaleway_lb": {
				Tok: scalewayDataSource(loadbalancersMod, "getLoadBalancer"),
				Docs: &tfbridge.DocInfo{
					Source: "lb.md",
				},
			},
			"scaleway_lbs": {
				Tok: scalewayDataSource(loadbalancersMod, "getLoadBalancers"),
				Docs: &tfbridge.DocInfo{
					Source: "lbs.md",
				},
			},
			"scaleway_secret": {
				Tok: scalewayDataSource(secretsMod, "getSecret"),
			},
			"scaleway_vpc": {
				Tok: scalewayDataSource(vpcMod, "getVpc"),
			},
			"scaleway_vpcs": {
				Tok: scalewayDataSource(vpcMod, "getVpcs"),
			},
			"scaleway_webhosting": {
				Tok: scalewayDataSource(webhostingMod, "getHosting"),
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			PackageName: "@pulumiverse/scaleway",
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			// Overlay: &tfbridge.OverlayInfo{},
			RespectSchemaVersion: true,
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "pulumiverse_scaleway",
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
			PyProject:            struct{ Enabled bool }{true},
			RespectSchemaVersion: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "Pulumiverse",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			RespectSchemaVersion: true,
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumiverse/pulumi-%[1]s/sdk/", scalewayPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				scalewayPkg,
			),
			GenerateResourceContainerTypes: true,
			RespectSchemaVersion:           true,
		},
	}

	prov.MustComputeTokens(
		tks.MappedModules(
			"scaleway_",
			scalewayMod,
			map[string]string{
				"account":       accountMod,
				"apple_silicon": appleSiliconMod,
				"autoscaling":   autoscalingMod,
				"billing":       billingMod,
				"block":         blockMod,
				"cockpit":       cockpitMod,
				"container":     containersMod,
				"domain":        domainMod,
				"baremetal":     elasticmetalMod,
				"flexible":      elasticmetalMod,
				"function":      functionsMod,
				"iam":           iamMod,
				"ipam":          ipamMod,
				"inference":     inferenceMod,
				"instance":      instanceMod,
				"iot":           iotMod,
				"job":           jobMod,
				"k8s":           kubernetesMod,
				"lb":            loadbalancersMod,
				"mnq":           mnqMod,
				"mongodb":       mongodbMod,
				"object":        objectMod,
				"rdb":           databasesMod,
				"redis":         redisMod,
				"registry":      registryMod,
				"secret":        secretsMod,
				"tem":           temMod,
				"vpc":           vpcMod,
				"webhosting":    webhostingMod,
			},
			tks.MakeStandard(scalewayPkg),
		),
	)
	prov.SetAutonaming(255, "-")
	prov.MustApplyAutoAliases()

	return prov
}
