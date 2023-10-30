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

	"github.com/lbrlabs/pulumi-scaleway/provider/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/scaleway/terraform-provider-scaleway/v2/scaleway"
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
	providerConfig := scaleway.ProviderConfig{}
	p := shimv2.NewProvider(scaleway.Provider(&providerConfig)())
	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                       p,
		Name:                    "scaleway",
		Description:             "A Pulumi package for creating and managing scaleway cloud resources.",
		Keywords:                []string{"pulumi", "scaleway", "lbrlabs"},
		TFProviderLicense:       refProviderLicense(tfbridge.MITLicenseType),
		License:                 "Apache-2.0",
		LogoURL:                 "https://raw.githubusercontent.com/lbrlabs/pulumi-scaleway/master/assets/scaleway.png", //nolint:golint,lll
		Homepage:                "https://leebriggs.co.uk/projects#pulumi-scaleway",
		Repository:              "https://github.com/lbrlabs/pulumi-scaleway",
		PluginDownloadURL:       "github://api.github.com/lbrlabs",
		GitHubOrg:               "scaleway", // not in the terraform-providers repo
		Publisher:               "lbrlabs",
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
			"scaleway_account_ssh_key": {Tok: scalewayResource(scalewayMod, "AccountSshKey")},
			"scaleway_account_project": {Tok: scalewayResource(scalewayMod, "AccountProject")},
			"scaleway_apple_silicon_server": {
				Tok: scalewayResource(scalewayMod, "AppleSliconValleyServer"),
				Docs: &tfbridge.DocInfo{
					Source: "apple_silicon.md",
				},
			},
			"scaleway_baremetal_server":                    {Tok: scalewayResource(scalewayMod, "BaremetalServer")},
			"scaleway_cockpit":                             {Tok: scalewayResource(scalewayMod, "Cockpit")},
			"scaleway_cockpit_grafana_user":                {Tok: scalewayResource(scalewayMod, "CockpitGrafanaUser")},
			"scaleway_cockpit_token":                       {Tok: scalewayResource(scalewayMod, "CockpitToken")},
			"scaleway_container":                           {Tok: scalewayResource(scalewayMod, "Container")},
			"scaleway_container_namespace":                 {Tok: scalewayResource(scalewayMod, "ContainerNamespace")},
			"scaleway_container_cron":                      {Tok: scalewayResource(scalewayMod, "ContainerCron")},
			"scaleway_container_domain":                    {Tok: scalewayResource(scalewayMod, "ContainerDomain")},
			"scaleway_container_token":                     {Tok: scalewayResource(scalewayMod, "ContainerToken")},
			"scaleway_container_trigger":                   {Tok: scalewayResource(scalewayMod, "ContainerTrigger")},
			"scaleway_domain_record":                       {Tok: scalewayResource(scalewayMod, "DomainRecord")},
			"scaleway_domain_zone":                         {Tok: scalewayResource(scalewayMod, "DomainZone")},
			"scaleway_documentdb_database":                 {Tok: scalewayResource(scalewayMod, "DocumentdbDatabase")},
			"scaleway_documentdb_instance":                 {Tok: scalewayResource(scalewayMod, "DocumentdbInstance")},
			"scaleway_documentdb_private_network_endpoint": {Tok: scalewayResource(scalewayMod, "DocumentdbPrivateNetworkEndpoint")}, //nolint:golint,lll
			"scaleway_documentdb_privilege":                {Tok: scalewayResource(scalewayMod, "DocumentdbPrivilege")},
			"scaleway_documentdb_read_replica":             {Tok: scalewayResource(scalewayMod, "DocumentdbReadReplica")},
			"scaleway_documentdb_user":                     {Tok: scalewayResource(scalewayMod, "DocumentdbUser")},
			"scaleway_flexible_ip_mac_address":             {Tok: scalewayResource(scalewayMod, "FlexibleIpMacAddress")},
			"scaleway_flexible_ip":                         {Tok: scalewayResource(scalewayMod, "FlexibleIp")},
			"scaleway_function":                            {Tok: scalewayResource(scalewayMod, "Function")},
			"scaleway_function_cron":                       {Tok: scalewayResource(scalewayMod, "FunctionCron")},
			"scaleway_function_domain":                     {Tok: scalewayResource(scalewayMod, "FunctionDomain")},
			"scaleway_function_namespace":                  {Tok: scalewayResource(scalewayMod, "FunctionNamespace")},
			"scaleway_function_token":                      {Tok: scalewayResource(scalewayMod, "FunctionToken")},
			"scaleway_iam_api_key":                         {Tok: scalewayResource(scalewayMod, "IamApiKey")},
			"scaleway_iam_application":                     {Tok: scalewayResource(scalewayMod, "IamApplication")},
			"scaleway_iam_group":                           {Tok: scalewayResource(scalewayMod, "IamGroup")},
			"scaleway_iam_policy":                          {Tok: scalewayResource(scalewayMod, "IamPolicy")},
			"scaleway_iam_ssh_key":                         {Tok: scalewayResource(scalewayMod, "IamSshKey")},
			"scaleway_iam_group_membership":                {Tok: scalewayResource(scalewayMod, "IamGroupMembership")},
			"scaleway_iam_user":                            {Tok: scalewayResource(scalewayMod, "IamUser")},
			"scaleway_instance_image":                      {Tok: scalewayResource(scalewayMod, "InstanceImage")},
			"scaleway_instance_ip":                         {Tok: scalewayResource(scalewayMod, "InstanceIp")},
			"scaleway_instance_ip_reverse_dns":             {Tok: scalewayResource(scalewayMod, "InstanceIpReverseDns")},
			"scaleway_instance_placement_group":            {Tok: scalewayResource(scalewayMod, "InstancePlacementGroup")},
			"scaleway_instance_private_nic":                {Tok: scalewayResource(scalewayMod, "InstancePrivateNic")},
			"scaleway_instance_security_group":             {Tok: scalewayResource(scalewayMod, "InstanceSecurityGroup")},
			"scaleway_instance_security_group_rules":       {Tok: scalewayResource(scalewayMod, "InstanceSecurityGroupRules")},
			"scaleway_instance_server":                     {Tok: scalewayResource(scalewayMod, "InstanceServer")},
			"scaleway_instance_snapshot":                   {Tok: scalewayResource(scalewayMod, "InstanceSnapshot")},
			"scaleway_instance_user_data":                  {Tok: scalewayResource(scalewayMod, "InstanceUserData")},
			"scaleway_instance_volume":                     {Tok: scalewayResource(scalewayMod, "InstanceVolume")},
			"scaleway_iot_device":                          {Tok: scalewayResource(scalewayMod, "IotDevice")},
			"scaleway_iot_hub":                             {Tok: scalewayResource(scalewayMod, "IotHub")},
			"scaleway_iot_network":                         {Tok: scalewayResource(scalewayMod, "IotNetwork")},
			"scaleway_iot_route":                           {Tok: scalewayResource(scalewayMod, "IotRoute")},
			"scaleway_k8s_cluster":                         {Tok: scalewayResource(scalewayMod, "KubernetesCluster")},
			"scaleway_k8s_pool":                            {Tok: scalewayResource(scalewayMod, "KubernetesNodePool")},
			"scaleway_lb":                                  {Tok: scalewayResource(scalewayMod, "Loadbalancer")},
			"scaleway_lb_backend":                          {Tok: scalewayResource(scalewayMod, "LoadbalancerBackend")},
			"scaleway_lb_certificate":                      {Tok: scalewayResource(scalewayMod, "LoadbalancerCertificate")},
			"scaleway_lb_frontend":                         {Tok: scalewayResource(scalewayMod, "LoadbalancerFrontend")},
			"scaleway_lb_ip":                               {Tok: scalewayResource(scalewayMod, "LoadbalancerIp")},
			"scaleway_lb_route":                            {Tok: scalewayResource(scalewayMod, "LoadbalancerRoute")},
			"scaleway_mnq_credential":                      {Tok: scalewayResource(scalewayMod, "MnqCredential")},
			"scaleway_mnq_namespace":                       {Tok: scalewayResource(scalewayMod, "MnqNamespace")},
			"scaleway_mnq_nats_account":                    {Tok: scalewayResource(scalewayMod, "MnqNatsAccount")},
			"scaleway_mnq_nats_credentials":                {Tok: scalewayResource(scalewayMod, "MnqNatsCredentials")},
			"scaleway_mnq_sqs":                             {Tok: scalewayResource(scalewayMod, "MnqSqs")},
			"scaleway_mnq_sqs_credentials":                 {Tok: scalewayResource(scalewayMod, "MnqSqsCredentials")},
			"scaleway_mnq_sqs_queue":                       {Tok: scalewayResource(scalewayMod, "MnqSqsQueue")},
			"scaleway_object":                              {Tok: scalewayResource(scalewayMod, "ObjectItem")},
			"scaleway_object_bucket":                       {Tok: scalewayResource(scalewayMod, "ObjectBucket")},
			"scaleway_object_bucket_acl":                   {Tok: scalewayResource(scalewayMod, "ObjectBucketAcl")},
			"scaleway_object_bucket_policy":                {Tok: scalewayResource(scalewayMod, "ObjectBucketPolicy")},
			"scaleway_object_bucket_website_configuration": {Tok: scalewayResource(scalewayMod, "ObjectBucketWebsiteConfiguration")}, //nolint:golint,lll
			"scaleway_object_bucket_lock_configuration":    {Tok: scalewayResource(scalewayMod, "ObjectBucketLockConfiguration")}, //nolint:golint,lll
			"scaleway_rdb_acl":                             {Tok: scalewayResource(scalewayMod, "DatabaseAcl")},
			"scaleway_redis_cluster":                       {Tok: scalewayResource(scalewayMod, "RedisCluster")},
			"scaleway_rdb_database":                        {Tok: scalewayResource(scalewayMod, "Database")},
			"scaleway_rdb_database_backup":                 {Tok: scalewayResource(scalewayMod, "DatabaseBackup")},
			"scaleway_rdb_instance":                        {Tok: scalewayResource(scalewayMod, "DatabaseInstance")},
			"scaleway_rdb_privilege":                       {Tok: scalewayResource(scalewayMod, "DatabasePrivilege")},
			"scaleway_rdb_read_replica":                    {Tok: scalewayResource(scalewayMod, "DatabaseReadReplica")},
			"scaleway_rdb_user":                            {Tok: scalewayResource(scalewayMod, "DatabaseUser")},
			"scaleway_registry_namespace":                  {Tok: scalewayResource(scalewayMod, "RegistryNamespace")},
			"scaleway_secret":                              {Tok: scalewayResource(scalewayMod, "Secret")},
			"scaleway_secret_version":                      {Tok: scalewayResource(scalewayMod, "SecretVersion")},
			"scaleway_tem_domain":                          {Tok: scalewayResource(scalewayMod, "TemDomain")},
			"scaleway_vpc_gateway_network":                 {Tok: scalewayResource(scalewayMod, "VpcGatewayNetwork")},
			"scaleway_vpc_private_network":                 {Tok: scalewayResource(scalewayMod, "VpcPrivateNetwork")},
			"scaleway_vpc_public_gateway":                  {Tok: scalewayResource(scalewayMod, "VpcPublicGateway")},
			"scaleway_vpc_public_gateway_dhcp":             {Tok: scalewayResource(scalewayMod, "VpcPublicGatewayDhcp")},
			"scaleway_vpc_public_gateway_ip":               {Tok: scalewayResource(scalewayMod, "VpcPublicGatewayIp")},
			"scaleway_vpc_public_gateway_ip_reverse_dns":   {Tok: scalewayResource(scalewayMod, "VpcPublicGatewayIpReverseDns")},
			"scaleway_vpc_public_gateway_pat_rule":         {Tok: scalewayResource(scalewayMod, "VpcPublicGatewayPatRule")},
			"scaleway_vpc_public_gateway_dhcp_reservation": {Tok: scalewayResource(scalewayMod, "VpcPublicGatewayDhcpReservation")}, //nolint:golint,lll
			"scaleway_function_trigger":                    {Tok: scalewayResource(scalewayMod, "FunctionTrigger")},
			"scaleway_lb_acl":                              {Tok: scalewayResource(scalewayMod, "LoadbalancerAcl")},
			"scaleway_mnq_queue":                           {Tok: scalewayResource(scalewayMod, "MnqQueue")},
			"scaleway_webhosting":                          {Tok: scalewayResource(scalewayMod, "Webhosting")},
			"scaleway_vpc":                                 {Tok: scalewayResource(scalewayMod, "Vpc")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"scaleway_account_ssh_key": {Tok: scalewayDataSource(scalewayMod, "getAccountSshKey")},
			"scaleway_account_project": {Tok: scalewayDataSource(scalewayMod, "getAccountProject")},
			"scaleway_baremetal_offer": {
				Tok: scalewayDataSource(scalewayMod, "getBaremetalOffer"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"cpu": {
						MaxItemsOne: boolRef(true),
					},
				},
			},
			"scaleway_baremetal_os":                        {Tok: scalewayDataSource(scalewayMod, "getBaremetalOs")},
			"scaleway_baremetal_option":                    {Tok: scalewayDataSource(scalewayMod, "getBaremetalOption")},
			"scaleway_baremetal_server":                    {Tok: scalewayDataSource(scalewayMod, "getBaremetalServer")},
			"scaleway_billing_consumptions":                {Tok: scalewayDataSource(scalewayMod, "getBillingConsumptions")},
			"scaleway_billing_invoices":                    {Tok: scalewayDataSource(scalewayMod, "getBillingInvoices")},
			"scaleway_cockpit":                             {Tok: scalewayDataSource(scalewayMod, "getCockpit")},
			"scaleway_container":                           {Tok: scalewayDataSource(scalewayMod, "getContainer")},
			"scaleway_container_namespace":                 {Tok: scalewayDataSource(scalewayMod, "getContainerNamespace")},
			"scaleway_domain_record":                       {Tok: scalewayDataSource(scalewayMod, "getDomainRecord")},
			"scaleway_domain_zone":                         {Tok: scalewayDataSource(scalewayMod, "getDomainZone")},
			"scaleway_documentdb_database":                 {Tok: scalewayDataSource(scalewayMod, "getDocumentdbDatabase")},
			"scaleway_documentdb_instance":                 {Tok: scalewayDataSource(scalewayMod, "getDocumentdbInstance")},
			"scaleway_documentdb_load_balancer_endpoint":   {Tok: scalewayDataSource(scalewayMod, "getDocumentdbLoadBalancerEndpoint")}, //nolint:golint,lll
			"scaleway_flexible_ip":                         {Tok: scalewayDataSource(scalewayMod, "getFlexibleIp")},
			"scaleway_flexible_ips":                        {Tok: scalewayDataSource(scalewayMod, "getFlexibleIps")},
			"scaleway_function":                            {Tok: scalewayDataSource(scalewayMod, "getFunction")},
			"scaleway_function_namespace":                  {Tok: scalewayDataSource(scalewayMod, "getFunctionNamespace")},
			"scaleway_instance_image":                      {Tok: scalewayDataSource(scalewayMod, "getInstanceImage")},
			"scaleway_iam_application":                     {Tok: scalewayDataSource(scalewayMod, "getIamApplication")},
			"scaleway_iam_group":                           {Tok: scalewayDataSource(scalewayMod, "getIamGroup")},
			"scaleway_iam_ssh_key":                         {Tok: scalewayDataSource(scalewayMod, "getIamSshKey")},
			"scaleway_iam_user":                            {Tok: scalewayDataSource(scalewayMod, "getIamUser")},
			"scaleway_ipam_ip":                             {Tok: scalewayDataSource(scalewayMod, "getIpamIp")},
			"scaleway_instance_ip":                         {Tok: scalewayDataSource(scalewayMod, "getInstanceIp")},
			"scaleway_instance_private_nic":                {Tok: scalewayDataSource(scalewayMod, "getInstancePrivateNic")},
			"scaleway_instance_security_group":             {Tok: scalewayDataSource(scalewayMod, "getInstanceSecurityGroup")},
			"scaleway_instance_server":                     {Tok: scalewayDataSource(scalewayMod, "getInstanceServer")},
			"scaleway_instance_servers":                    {Tok: scalewayDataSource(scalewayMod, "getInstanceServers")},
			"scaleway_instance_snapshot":                   {Tok: scalewayDataSource(scalewayMod, "getInstanceSnapshot")},
			"scaleway_instance_volume":                     {Tok: scalewayDataSource(scalewayMod, "getInstanceVolume")},
			"scaleway_iot_device":                          {Tok: scalewayDataSource(scalewayMod, "getIotDevice")},
			"scaleway_iot_hub":                             {Tok: scalewayDataSource(scalewayMod, "getIotHub")},
			"scaleway_k8s_cluster":                         {Tok: scalewayDataSource(scalewayMod, "getKubernetesCluster")},
			"scaleway_k8s_pool":                            {Tok: scalewayDataSource(scalewayMod, "getKubernetesNodePool")},
			"scaleway_k8s_version":                         {Tok: scalewayDataSource(scalewayMod, "getK8sVersion")},
			"scaleway_lb":                                  {Tok: scalewayDataSource(scalewayMod, "getLoadbalancer")},
			"scaleway_lb_acls":                             {Tok: scalewayDataSource(scalewayMod, "getLbAcls")},
			"scaleway_lb_backend":                          {Tok: scalewayDataSource(scalewayMod, "getLbBackend")},
			"scaleway_lb_backends":                         {Tok: scalewayDataSource(scalewayMod, "getLbBackends")},
			"scaleway_lb_certificate":                      {Tok: scalewayDataSource(scalewayMod, "getLoadbalancerCertificate")},
			"scaleway_lb_frontend":                         {Tok: scalewayDataSource(scalewayMod, "getLbFrontend")},
			"scaleway_lb_frontends":                        {Tok: scalewayDataSource(scalewayMod, "getLbFrontends")},
			"scaleway_lb_ip":                               {Tok: scalewayDataSource(scalewayMod, "getLoadbalancerIp")},
			"scaleway_lb_ips":                              {Tok: scalewayDataSource(scalewayMod, "getLbIps")},
			"scaleway_lb_route":                            {Tok: scalewayDataSource(scalewayMod, "getLbRoute")},
			"scaleway_lb_routes":                           {Tok: scalewayDataSource(scalewayMod, "getLbRoutes")},
			"scaleway_lbs":                                 {Tok: scalewayDataSource(scalewayMod, "getLbs")},
			"scaleway_object_bucket":                       {Tok: scalewayDataSource(scalewayMod, "getObjectBucket")},
			"scaleway_marketplace_image":                   {Tok: scalewayDataSource(scalewayMod, "getMarketplaceImage")},
			"scaleway_mnq_sqs":                             {Tok: scalewayDataSource(scalewayMod, "getMnqSqs")},
			"scaleway_rdb_acl":                             {Tok: scalewayDataSource(scalewayMod, "getDatabaseAcl")},
			"scaleway_rdb_database":                        {Tok: scalewayDataSource(scalewayMod, "getDatabase")},
			"scaleway_rdb_database_backup":                 {Tok: scalewayDataSource(scalewayMod, "getDatabaseBackup")},
			"scaleway_rdb_instance":                        {Tok: scalewayDataSource(scalewayMod, "getDatabaseInstance")},
			"scaleway_rdb_privilege":                       {Tok: scalewayDataSource(scalewayMod, "getDatabasePrivilege")},
			"scaleway_redis_cluster":                       {Tok: scalewayDataSource(scalewayMod, "getRedisCluster")},
			"scaleway_registry_image":                      {Tok: scalewayDataSource(scalewayMod, "getRegistryImage")},
			"scaleway_registry_namespace":                  {Tok: scalewayDataSource(scalewayMod, "getRegistryNamespace")},
			"scaleway_secret":                              {Tok: scalewayDataSource(scalewayMod, "getSecret")},
			"scaleway_secret_version":                      {Tok: scalewayDataSource(scalewayMod, "getSecretVersion")},
			"scaleway_tem_domain":                          {Tok: scalewayDataSource(scalewayMod, "getTemDomain")},
			"scaleway_vpc_gateway_network":                 {Tok: scalewayDataSource(scalewayMod, "getVpcGatewayNetwork")},
			"scaleway_vpc_private_network":                 {Tok: scalewayDataSource(scalewayMod, "getVpcPrivateNetwork")},
			"scaleway_vpc_public_gateway":                  {Tok: scalewayDataSource(scalewayMod, "getVpcPublicGateway")},
			"scaleway_vpc_public_gateway_dhcp":             {Tok: scalewayDataSource(scalewayMod, "getVpcPublicGatewayDhcp")},
			"scaleway_vpc_public_gateway_dhcp_reservation": {Tok: scalewayDataSource(scalewayMod, "getVpcPublicGatewayDhcpReservation")}, //nolint:golint,lll
			"scaleway_vpc_public_gateway_ip":               {Tok: scalewayDataSource(scalewayMod, "getVpcPublicGatewayIp")},
			"scaleway_availability_zones":                  {Tok: scalewayDataSource(scalewayMod, "getAvailabilityZones")},
			"scaleway_cockpit_plan":                        {Tok: scalewayDataSource(scalewayMod, "getCockpitPlan")},
			"scaleway_object_bucket_policy":                {Tok: scalewayDataSource(scalewayMod, "getObjectBucketPolicy")},
			"scaleway_vpc":                                 {Tok: scalewayDataSource(scalewayMod, "getVpc")},
			"scaleway_vpcs":                                {Tok: scalewayDataSource(scalewayMod, "getVpcs")},
			"scaleway_vpc_public_gateway_pat_rule":         {Tok: scalewayDataSource(scalewayMod, "getVpcPublicPatRule")},
			"scaleway_webhosting_offer":                    {Tok: scalewayDataSource(scalewayMod, "getWebHostOffer")},
			"scaleway_webhosting":                          {Tok: scalewayDataSource(scalewayMod, "getWebhosting")},
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
			PackageName: "@lbrlabs/pulumi-scaleway",
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			// Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "lbrlabs_pulumi_scaleway",
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "Lbrlabs.PulumiPackage",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/lbrlabs/pulumi-%[1]s/sdk/", scalewayPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				scalewayPkg,
			),
			GenerateResourceContainerTypes: true,
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
