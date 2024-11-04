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
			"scaleway_apple_silicon_server": {
				Docs: &tfbridge.DocInfo{
					Source: "apple_silicon.md",
				},
			},
			"scaleway_k8s_cluster": {
				Tok: scalewayResource(scalewayMod, "KubernetesCluster"),
				Docs: &tfbridge.DocInfo{
					Source: "k8s_cluster.md",
				},
			},
			"scaleway_k8s_pool": {
				Tok: scalewayResource(scalewayMod, "KubernetesNodePool"),
				Docs: &tfbridge.DocInfo{
					Source: "k8s_pool.md",
				},
			},
			"scaleway_lb": {
				Tok: scalewayResource(scalewayMod, "Loadbalancer"),
				Docs: &tfbridge.DocInfo{
					Source: "lb.md",
				},
			},
			"scaleway_lb_acl": {
				Tok: scalewayResource(scalewayMod, "LoadbalancerAcl"),
				Docs: &tfbridge.DocInfo{
					Source: "lb_acl.md",
				},
			},
			"scaleway_lb_backend": {
				Tok: scalewayResource(scalewayMod, "LoadbalancerBackend"),
				Docs: &tfbridge.DocInfo{
					Source: "lb_backend.md",
				},
			},
			"scaleway_lb_certificate": {
				Tok: scalewayResource(scalewayMod, "LoadbalancerCertificate"),
				Docs: &tfbridge.DocInfo{
					Source: "lb_certificate.md",
				},
			},
			"scaleway_lb_frontend": {
				Tok: scalewayResource(scalewayMod, "LoadbalancerFrontend"),
				Docs: &tfbridge.DocInfo{
					Source: "lb_frontend.md",
				},
			},
			"scaleway_lb_ip": {
				Tok: scalewayResource(scalewayMod, "LoadbalancerIp"),
				Docs: &tfbridge.DocInfo{
					Source: "lb_ip.md",
				},
			},
			"scaleway_lb_route": {
				Tok: scalewayResource(scalewayMod, "LoadbalancerRoute"),
				Docs: &tfbridge.DocInfo{
					Source: "lb_route.md",
				},
			},
			"scaleway_mongodb_instance": {
				Tok: scalewayResource(scalewayMod, "MongoDbInstance"),
				Docs: &tfbridge.DocInfo{
					Source: "mongodb_instance.md",
				},
			},
			"scaleway_mongodb_snapshot": {
				Tok: scalewayResource(scalewayMod, "MongoDbSnapshot"),
				Docs: &tfbridge.DocInfo{
					Source: "mongodb_snapshot.md",
				},
			},
			"scaleway_object": {
				Tok: scalewayResource(scalewayMod, "ObjectItem"),
				Docs: &tfbridge.DocInfo{
					Source: "object.md",
				},
			},
			"scaleway_rdb_acl": {
				Tok: scalewayResource(scalewayMod, "DatabaseAcl"),
				Docs: &tfbridge.DocInfo{
					Source: "rdb_acl.md",
				},
			},
			"scaleway_redis_cluster": {
				Tok: scalewayResource(scalewayMod, "RedisCluster"),
				Docs: &tfbridge.DocInfo{
					Source: "redis_cluster.md",
				},
			},
			"scaleway_rdb_database": {
				Tok: scalewayResource(scalewayMod, "Database"),
				Docs: &tfbridge.DocInfo{
					Source: "rdb_database.md",
				},
			},
			"scaleway_rdb_database_backup": {
				Tok: scalewayResource(scalewayMod, "DatabaseBackup"),
				Docs: &tfbridge.DocInfo{
					Source: "rdb_database_backup.md",
				},
			},
			"scaleway_rdb_instance": {
				Tok: scalewayResource(scalewayMod, "DatabaseInstance"),
				Docs: &tfbridge.DocInfo{
					Source: "rdb_instance.md",
				},
			},
			"scaleway_rdb_privilege": {
				Tok: scalewayResource(scalewayMod, "DatabasePrivilege"),
				Docs: &tfbridge.DocInfo{
					Source: "rdb_privilege.md",
				},
			},
			"scaleway_rdb_read_replica": {
				Tok: scalewayResource(scalewayMod, "DatabaseReadReplica"),
				Docs: &tfbridge.DocInfo{
					Source: "rdb_read_replica.md",
				},
			},
			"scaleway_rdb_user": {
				Tok: scalewayResource(scalewayMod, "DatabaseUser"),
				Docs: &tfbridge.DocInfo{
					Source: "rdb_user.md",
				},
			},
			"scaleway_sdb_sql_database": {
				Tok: scalewayResource(scalewayMod, "SdbDatabase"),
				Docs: &tfbridge.DocInfo{
					Source: "sdb_sql_database.md",
				},
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
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
			"scaleway_config": {
				Docs: &tfbridge.DocInfo{
					// Source: "config.md", // File doesn't exist in upstream TF provider repo
					AllowMissing: true,
				},
			},
			"scaleway_k8s_cluster": {
				Tok: scalewayDataSource(scalewayMod, "getKubernetesCluster"),
				Docs: &tfbridge.DocInfo{
					Source: "k8s_cluster.md",
				},
			},
			"scaleway_k8s_pool": {
				Tok: scalewayDataSource(scalewayMod, "getKubernetesNodePool"),
				Docs: &tfbridge.DocInfo{
					Source: "k8s_pool.md",
				},
			},
			"scaleway_k8s_version": {
				Tok: scalewayDataSource(scalewayMod, "getK8sVersion"),
				Docs: &tfbridge.DocInfo{
					Source: "k8s_version.md",
				},
			},
			"scaleway_lb": {
				Tok: scalewayDataSource(scalewayMod, "getLoadbalancer"),
				Docs: &tfbridge.DocInfo{
					Source: "lb.md",
				},
			},
			"scaleway_lb_acls": {
				Tok: scalewayDataSource(scalewayMod, "getLbAcls"),
				Docs: &tfbridge.DocInfo{
					Source: "lb_acls.md",
				},
			},
			"scaleway_lb_backend": {
				Tok: scalewayDataSource(scalewayMod, "getLbBackend"),
				Docs: &tfbridge.DocInfo{
					Source: "lb_backend.md",
				},
			},
			"scaleway_lb_backends": {
				Tok: scalewayDataSource(scalewayMod, "getLbBackends"),
				Docs: &tfbridge.DocInfo{
					Source: "lb_backends.md",
				},
			},
			"scaleway_lb_certificate": {
				Tok: scalewayDataSource(scalewayMod, "getLoadbalancerCertificate"),
				Docs: &tfbridge.DocInfo{
					Source: "lb_certificate.md",
				},
			},
			"scaleway_lb_frontend": {
				Tok: scalewayDataSource(scalewayMod, "getLbFrontend"),
				Docs: &tfbridge.DocInfo{
					Source: "lb_frontend.md",
				},
			},
			"scaleway_lb_frontends": {
				Tok: scalewayDataSource(scalewayMod, "getLbFrontends"),
				Docs: &tfbridge.DocInfo{
					Source: "lb_frontends.md",
				},
			},
			"scaleway_lb_ip": {
				Tok: scalewayDataSource(scalewayMod, "getLoadbalancerIp"),
				Docs: &tfbridge.DocInfo{
					Source: "lb_ip.md",
				},
			},
			"scaleway_lb_ips": {
				Tok: scalewayDataSource(scalewayMod, "getLbIps"),
				Docs: &tfbridge.DocInfo{
					Source: "lb_ips.md",
				},
			},
			"scaleway_lb_route": {
				Tok: scalewayDataSource(scalewayMod, "getLbRoute"),
				Docs: &tfbridge.DocInfo{
					Source: "lb_route.md",
				},
			},
			"scaleway_lb_routes": {
				Tok: scalewayDataSource(scalewayMod, "getLbRoutes"),
				Docs: &tfbridge.DocInfo{
					Source: "lb_routes.md",
				},
			},
			"scaleway_lbs": {
				Tok: scalewayDataSource(scalewayMod, "getLbs"),
				Docs: &tfbridge.DocInfo{
					Source: "lbs.md",
				},
			},
			"scaleway_mongodb_instance": {
				Tok: scalewayDataSource(scalewayMod, "getMongoDbInstance"),
				Docs: &tfbridge.DocInfo{
					Source: "mongodb_instance.md",
				},
			},
			"scaleway_rdb_acl": {
				Tok: scalewayDataSource(scalewayMod, "getDatabaseAcl"),
				Docs: &tfbridge.DocInfo{
					Source: "rdb_acl.md",
				},
			},
			"scaleway_rdb_database": {
				Tok: scalewayDataSource(scalewayMod, "getDatabase"),
				Docs: &tfbridge.DocInfo{
					Source: "rdb_database.md",
				},
			},
			"scaleway_rdb_database_backup": {
				Tok: scalewayDataSource(scalewayMod, "getDatabaseBackup"),
				Docs: &tfbridge.DocInfo{
					Source: "rdb_database_backup.md",
				},
			},
			"scaleway_rdb_instance": {
				Tok: scalewayDataSource(scalewayMod, "getDatabaseInstance"),
				Docs: &tfbridge.DocInfo{
					Source: "rdb_instance.md",
				},
			},
			"scaleway_rdb_privilege": {
				Tok: scalewayDataSource(scalewayMod, "getDatabasePrivilege"),
				Docs: &tfbridge.DocInfo{
					Source: "rdb_privilege.md",
				},
			},
			"scaleway_vpc_public_gateway_pat_rule": {
				Tok: scalewayDataSource(scalewayMod, "getVpcPublicPatRule"),
				Docs: &tfbridge.DocInfo{
					Source: "vpc_public_gateway_pat_rule.md",
				},
			},
			"scaleway_webhosting_offer": {
				Tok: scalewayDataSource(scalewayMod, "getWebHostOffer"),
				Docs: &tfbridge.DocInfo{
					Source: "webhosting_offer.md",
				},
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

	prov.MustComputeTokens(tks.SingleModule("scaleway_", scalewayMod, tks.MakeStandard(scalewayPkg)))
	prov.SetAutonaming(255, "-")
	prov.MustApplyAutoAliases()

	return prov
}
