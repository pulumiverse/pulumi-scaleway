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

	shim "github.com/scaleway/terraform-provider-scaleway/v2/shim"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
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
			"scaleway_account_ssh_key": {
				Tok: scalewayResource(scalewayMod, "AccountSshKey"),
				Docs: &tfbridge.DocInfo{
					Source: "account_ssh_key.md",
				},
			},
			"scaleway_account_project": {
				Tok: scalewayResource(scalewayMod, "AccountProject"),
				Docs: &tfbridge.DocInfo{
					Source: "account_project.md",
				},
			},
			"scaleway_apple_silicon_server": {
				Tok: scalewayResource(scalewayMod, "AppleSiliconServer"),
				Docs: &tfbridge.DocInfo{
					Source: "apple_silicon.md",
				},
			},
			"scaleway_baremetal_server": {
				Tok: scalewayResource(scalewayMod, "BaremetalServer"),
				Docs: &tfbridge.DocInfo{
					Source: "baremetal_server.md",
				},
			},
			"scaleway_block_snapshot": {
				Tok: scalewayResource(scalewayMod, "BlockSnapshot"),
				Docs: &tfbridge.DocInfo{
					Source: "block_snapshot.md",
				},
			},
			"scaleway_block_volume": {
				Tok: scalewayResource(scalewayMod, "BlockVolume"),
				Docs: &tfbridge.DocInfo{
					Source: "block_volume.md",
				},
			},
			"scaleway_cockpit": {
				Tok: scalewayResource(scalewayMod, "Cockpit"),
				Docs: &tfbridge.DocInfo{
					Source: "cockpit.md",
				},
			},
			"scaleway_cockpit_alert_manager": {
				Tok: scalewayResource(scalewayMod, "CockpitAlertManager"),
				Docs: &tfbridge.DocInfo{
					Source: "cockpit_alert_manager.md",
				},
			},
			"scaleway_cockpit_grafana_user": {
				Tok: scalewayResource(scalewayMod, "CockpitGrafanaUser"),
				Docs: &tfbridge.DocInfo{
					Source: "cockpit_grafana_user.md",
				},
			},
			"scaleway_cockpit_source": {
				Tok: scalewayResource(scalewayMod, "CockpitSource"),
				Docs: &tfbridge.DocInfo{
					Source: "cockpit_source.md",
				},
			},
			"scaleway_cockpit_token": {
				Tok: scalewayResource(scalewayMod, "CockpitToken"),
				Docs: &tfbridge.DocInfo{
					Source: "cockpit_token.md",
				},
			},
			"scaleway_container": {
				Tok: scalewayResource(scalewayMod, "Container"),
				Docs: &tfbridge.DocInfo{
					Source: "container.md",
				},
			},
			"scaleway_container_namespace": {
				Tok: scalewayResource(scalewayMod, "ContainerNamespace"),
				Docs: &tfbridge.DocInfo{
					Source: "container_namespace.md",
				},
			},
			"scaleway_container_cron": {
				Tok: scalewayResource(scalewayMod, "ContainerCron"),
				Docs: &tfbridge.DocInfo{
					Source: "container_cron.md",
				},
			},
			"scaleway_container_domain": {
				Tok: scalewayResource(scalewayMod, "ContainerDomain"),
				Docs: &tfbridge.DocInfo{
					Source: "container_domain.md",
				},
			},
			"scaleway_container_token": {
				Tok: scalewayResource(scalewayMod, "ContainerToken"),
				Docs: &tfbridge.DocInfo{
					Source: "container_token.md",
				},
			},
			"scaleway_container_trigger": {
				Tok: scalewayResource(scalewayMod, "ContainerTrigger"),
				Docs: &tfbridge.DocInfo{
					Source: "container_trigger.md",
				},
			},
			"scaleway_domain_record": {
				Tok: scalewayResource(scalewayMod, "DomainRecord"),
				Docs: &tfbridge.DocInfo{
					Source: "domain_record.md",
				},
			},
			"scaleway_domain_zone": {
				Tok: scalewayResource(scalewayMod, "DomainZone"),
				Docs: &tfbridge.DocInfo{
					Source: "domain_zone.md",
				},
			},
			"scaleway_documentdb_database": {
				Tok: scalewayResource(scalewayMod, "DocumentdbDatabase"),
				Docs: &tfbridge.DocInfo{
					Source: "documentdb_database.md",
				},
			},
			"scaleway_documentdb_instance": {
				Tok: scalewayResource(scalewayMod, "DocumentdbInstance"),
				Docs: &tfbridge.DocInfo{
					Source: "documentdb_instance.md",
				},
			},
			"scaleway_documentdb_private_network_endpoint": {
				Tok: scalewayResource(scalewayMod, "DocumentdbPrivateNetworkEndpoint"),
				Docs: &tfbridge.DocInfo{
					Source: "documentdb_private_network_endpoint.md",
				},
			},
			"scaleway_documentdb_privilege": {
				Tok: scalewayResource(scalewayMod, "DocumentdbPrivilege"),
				Docs: &tfbridge.DocInfo{
					Source: "documentdb_privilege.md",
				},
			},
			"scaleway_documentdb_read_replica": {
				Tok: scalewayResource(scalewayMod, "DocumentdbReadReplica"),
				Docs: &tfbridge.DocInfo{
					Source: "documentdb_read_replica.md",
				},
			},
			"scaleway_documentdb_user": {
				Tok: scalewayResource(scalewayMod, "DocumentdbUser"),
				Docs: &tfbridge.DocInfo{
					Source: "documentdb_user.md",
				},
			},
			"scaleway_flexible_ip_mac_address": {
				Tok: scalewayResource(scalewayMod, "FlexibleIpMacAddress"),
				Docs: &tfbridge.DocInfo{
					Source: "flexible_ip_mac_address.md",
				},
			},
			"scaleway_flexible_ip": {
				Tok: scalewayResource(scalewayMod, "FlexibleIp"),
				Docs: &tfbridge.DocInfo{
					Source: "flexible_ip.md",
				},
			},
			"scaleway_function": {
				Tok: scalewayResource(scalewayMod, "Function"),
				Docs: &tfbridge.DocInfo{
					Source: "function.md",
				},
			},
			"scaleway_function_cron": {
				Tok: scalewayResource(scalewayMod, "FunctionCron"),
				Docs: &tfbridge.DocInfo{
					Source: "function_cron.md",
				},
			},
			"scaleway_function_domain": {
				Tok: scalewayResource(scalewayMod, "FunctionDomain"),
				Docs: &tfbridge.DocInfo{
					Source: "function_domain.md",
				},
			},
			"scaleway_function_namespace": {
				Tok: scalewayResource(scalewayMod, "FunctionNamespace"),
				Docs: &tfbridge.DocInfo{
					Source: "function_namespace.md",
				},
			},
			"scaleway_function_token": {
				Tok: scalewayResource(scalewayMod, "FunctionToken"),
				Docs: &tfbridge.DocInfo{
					Source: "function_token.md",
				},
			},
			"scaleway_iam_api_key": {
				Tok: scalewayResource(scalewayMod, "IamApiKey"),
				Docs: &tfbridge.DocInfo{
					Source: "iam_api_key.md",
				},
			},
			"scaleway_iam_application": {
				Tok: scalewayResource(scalewayMod, "IamApplication"),
				Docs: &tfbridge.DocInfo{
					Source: "iam_application.md",
				},
			},
			"scaleway_iam_group": {
				Tok: scalewayResource(scalewayMod, "IamGroup"),
				Docs: &tfbridge.DocInfo{
					Source: "iam_group.md",
				},
			},
			"scaleway_iam_policy": {
				Tok: scalewayResource(scalewayMod, "IamPolicy"),
				Docs: &tfbridge.DocInfo{
					Source: "iam_policy.md",
				},
			},
			"scaleway_iam_ssh_key": {
				Tok: scalewayResource(scalewayMod, "IamSshKey"),
				Docs: &tfbridge.DocInfo{
					Source: "iam_ssh_key.md",
				},
			},
			"scaleway_iam_group_membership": {
				Tok: scalewayResource(scalewayMod, "IamGroupMembership"),
				Docs: &tfbridge.DocInfo{
					Source: "iam_group_membership.md",
				},
			},
			"scaleway_iam_user": {
				Tok: scalewayResource(scalewayMod, "IamUser"),
				Docs: &tfbridge.DocInfo{
					Source: "iam_user.md",
				},
			},
			"scaleway_instance_image": {
				Tok: scalewayResource(scalewayMod, "InstanceImage"),
				Docs: &tfbridge.DocInfo{
					Source: "instance_image.md",
				},
			},
			"scaleway_instance_ip": {
				Tok: scalewayResource(scalewayMod, "InstanceIp"),
				Docs: &tfbridge.DocInfo{
					Source: "instance_ip.md",
				},
			},
			"scaleway_instance_ip_reverse_dns": {
				Tok: scalewayResource(scalewayMod, "InstanceIpReverseDns"),
				Docs: &tfbridge.DocInfo{
					Source: "instance_ip_reverse_dns.md",
				},
			},
			"scaleway_instance_placement_group": {
				Tok: scalewayResource(scalewayMod, "InstancePlacementGroup"),
				Docs: &tfbridge.DocInfo{
					Source: "instance_placement_group.md",
				},
			},
			"scaleway_instance_private_nic": {
				Tok: scalewayResource(scalewayMod, "InstancePrivateNic"),
				Docs: &tfbridge.DocInfo{
					Source: "instance_private_nic.md",
				},
			},
			"scaleway_instance_security_group": {
				Tok: scalewayResource(scalewayMod, "InstanceSecurityGroup"),
				Docs: &tfbridge.DocInfo{
					Source: "instance_security_group.md",
				},
			},
			"scaleway_instance_security_group_rules": {
				Tok: scalewayResource(scalewayMod, "InstanceSecurityGroupRules"),
				Docs: &tfbridge.DocInfo{
					Source: "instance_security_group_rules.md",
				},
			},
			"scaleway_instance_server": {
				Tok: scalewayResource(scalewayMod, "InstanceServer"),
				Docs: &tfbridge.DocInfo{
					Source: "instance_server.md",
				},
			},
			"scaleway_instance_snapshot": {
				Tok: scalewayResource(scalewayMod, "InstanceSnapshot"),
				Docs: &tfbridge.DocInfo{
					Source: "instance_snapshot.md",
				},
			},
			"scaleway_instance_user_data": {
				Tok: scalewayResource(scalewayMod, "InstanceUserData"),
				Docs: &tfbridge.DocInfo{
					Source: "instance_user_data.md",
				},
			},
			"scaleway_instance_volume": {
				Tok: scalewayResource(scalewayMod, "InstanceVolume"),
				Docs: &tfbridge.DocInfo{
					Source: "instance_volume.md",
				},
			},
			"scaleway_iot_device": {
				Tok: scalewayResource(scalewayMod, "IotDevice"),
				Docs: &tfbridge.DocInfo{
					Source: "iot_device.md",
				},
			},
			"scaleway_iot_hub": {
				Tok: scalewayResource(scalewayMod, "IotHub"),
				Docs: &tfbridge.DocInfo{
					Source: "iot_hub.md",
				},
			},
			"scaleway_iot_network": {
				Tok: scalewayResource(scalewayMod, "IotNetwork"),
				Docs: &tfbridge.DocInfo{
					Source: "iot_network.md",
				},
			},
			"scaleway_iot_route": {
				Tok: scalewayResource(scalewayMod, "IotRoute"),
				Docs: &tfbridge.DocInfo{
					Source: "iot_route.md",
				},
			},
			"scaleway_ipam_ip": {
				Tok: scalewayResource(scalewayMod, "IpamIp"),
				Docs: &tfbridge.DocInfo{
					Source: "ipam_ip.md",
				},
			},
			"scaleway_ipam_ip_reverse_dns": {
				Tok: scalewayResource(scalewayMod, "IpamIpReverseDns"),
				Docs: &tfbridge.DocInfo{
					Source: "ipam_ip_reverse_dns.md",
				},
			},
			"scaleway_job_definition": {
				Tok: scalewayResource(scalewayMod, "JobDefinition"),
				Docs: &tfbridge.DocInfo{
					Source: "job_definition.md",
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
			"scaleway_mnq_nats_account": {
				Tok: scalewayResource(scalewayMod, "MnqNatsAccount"),
				Docs: &tfbridge.DocInfo{
					Source: "mnq_nats_account.md",
				},
			},
			"scaleway_mnq_nats_credentials": {
				Tok: scalewayResource(scalewayMod, "MnqNatsCredentials"),
				Docs: &tfbridge.DocInfo{
					Source: "mnq_nats_credentials.md",
				},
			},
			"scaleway_mnq_sns": {
				Tok: scalewayResource(scalewayMod, "MnqSns"),
				Docs: &tfbridge.DocInfo{
					Source: "mnq_sns.md",
				},
			},
			"scaleway_mnq_sns_credentials": {
				Tok: scalewayResource(scalewayMod, "MnqSnsCredentials"),
				Docs: &tfbridge.DocInfo{
					Source: "mnq_sns_credentials.md",
				},
			},
			"scaleway_mnq_sns_topic": {
				Tok: scalewayResource(scalewayMod, "MnqSnsTopic"),
				Docs: &tfbridge.DocInfo{
					Source: "mnq_sns_topic.md",
				},
			},
			"scaleway_mnq_sns_topic_subscription": {
				Tok: scalewayResource(scalewayMod, "MnqSnsTopicSubscription"),
				Docs: &tfbridge.DocInfo{
					Source: "mnq_sns_topic_subscriptioin.md",
				},
			},
			"scaleway_mnq_sqs": {
				Tok: scalewayResource(scalewayMod, "MnqSqs"),
				Docs: &tfbridge.DocInfo{
					Source: "mnq_sqs.md",
				},
			},
			"scaleway_mnq_sqs_credentials": {
				Tok: scalewayResource(scalewayMod, "MnqSqsCredentials"),
				Docs: &tfbridge.DocInfo{
					Source: "mnq_sqs_credentials.md",
				},
			},
			"scaleway_mnq_sqs_queue": {
				Tok: scalewayResource(scalewayMod, "MnqSqsQueue"),
				Docs: &tfbridge.DocInfo{
					Source: "mnq_sqs_queue.md",
				},
			},
			"scaleway_object": {
				Tok: scalewayResource(scalewayMod, "ObjectItem"),
				Docs: &tfbridge.DocInfo{
					Source: "object.md",
				},
			},
			"scaleway_object_bucket": {
				Tok: scalewayResource(scalewayMod, "ObjectBucket"),
				Docs: &tfbridge.DocInfo{
					Source: "object_bucket.md",
				},
			},
			"scaleway_object_bucket_acl": {
				Tok: scalewayResource(scalewayMod, "ObjectBucketAcl"),
				Docs: &tfbridge.DocInfo{
					Source: "object_bucket_acl.md",
				},
			},
			"scaleway_object_bucket_policy": {
				Tok: scalewayResource(scalewayMod, "ObjectBucketPolicy"),
				Docs: &tfbridge.DocInfo{
					Source: "object_bucket_policy.md",
				},
			},
			"scaleway_object_bucket_website_configuration": {
				Tok: scalewayResource(scalewayMod, "ObjectBucketWebsiteConfiguration"),
				Docs: &tfbridge.DocInfo{
					Source: "object_bucket_website_configuration.md",
				},
			},
			"scaleway_object_bucket_lock_configuration": {
				Tok: scalewayResource(scalewayMod, "ObjectBucketLockConfiguration"),
				Docs: &tfbridge.DocInfo{
					Source: "object_bucket_lock_configuration.md",
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
			"scaleway_registry_namespace": {
				Tok: scalewayResource(scalewayMod, "RegistryNamespace"),
				Docs: &tfbridge.DocInfo{
					Source: "registry_namespace.md",
				},
			},
			"scaleway_sdb_sql_database": {
				Tok: scalewayResource(scalewayMod, "SdbDatabase"),
				Docs: &tfbridge.DocInfo{
					Source: "sdb_sql_database.md",
				},
			},
			"scaleway_secret": {
				Tok: scalewayResource(scalewayMod, "Secret"),
				Docs: &tfbridge.DocInfo{
					Source: "secret.md",
				},
			},
			"scaleway_secret_version": {
				Tok: scalewayResource(scalewayMod, "SecretVersion"),
				Docs: &tfbridge.DocInfo{
					Source: "secret_version.md",
				},
			},
			"scaleway_tem_domain": {
				Tok: scalewayResource(scalewayMod, "TemDomain"),
				Docs: &tfbridge.DocInfo{
					Source: "tem_domain.md",
				},
			},
			"scaleway_tem_domain_validation": {
				Tok: scalewayResource(scalewayMod, "TemDomainValidation"),
				Docs: &tfbridge.DocInfo{
					Source: "tem_domain_validation.md",
				},
			},
			"scaleway_tem_webhook": {
				Tok: scalewayResource(scalewayMod, "TemWebhook"),
				Docs: &tfbridge.DocInfo{
					Source: "tem_webhook.md",
				},
			},
			"scaleway_vpc_gateway_network": {
				Tok: scalewayResource(scalewayMod, "VpcGatewayNetwork"),
				Docs: &tfbridge.DocInfo{
					Source: "vpc_gateway_network.md",
				},
			},
			"scaleway_vpc_private_network": {
				Tok: scalewayResource(scalewayMod, "VpcPrivateNetwork"),
				Docs: &tfbridge.DocInfo{
					Source: "vpc_private_network.md",
				},
			},
			"scaleway_vpc_public_gateway": {
				Tok: scalewayResource(scalewayMod, "VpcPublicGateway"),
				Docs: &tfbridge.DocInfo{
					Source: "vpc_public_gateway.md",
				},
			},
			"scaleway_vpc_public_gateway_dhcp": {
				Tok: scalewayResource(scalewayMod, "VpcPublicGatewayDhcp"),
				Docs: &tfbridge.DocInfo{
					Source: "vpc_public_gateway_dhcp.md",
				},
			},
			"scaleway_vpc_public_gateway_ip": {
				Tok: scalewayResource(scalewayMod, "VpcPublicGatewayIp"),
				Docs: &tfbridge.DocInfo{
					Source: "vpc_public_gateway_ip.md",
				},
			},
			"scaleway_vpc_public_gateway_ip_reverse_dns": {
				Tok: scalewayResource(scalewayMod, "VpcPublicGatewayIpReverseDns"),
				Docs: &tfbridge.DocInfo{
					Source: "vpc_public_gateway_ip_reverse_dns.md",
				},
			},
			"scaleway_vpc_public_gateway_pat_rule": {
				Tok: scalewayResource(scalewayMod, "VpcPublicGatewayPatRule"),
				Docs: &tfbridge.DocInfo{
					Source: "vpc_public_gateway_pat_rule.md",
				},
			},
			"scaleway_vpc_public_gateway_dhcp_reservation": {
				Tok: scalewayResource(scalewayMod, "VpcPublicGatewayDhcpReservation"),
				Docs: &tfbridge.DocInfo{
					Source: "vpc_public_gateway_dhcp_reservation.md",
				},
			},
			"scaleway_function_trigger": {
				Tok: scalewayResource(scalewayMod, "FunctionTrigger"),
				Docs: &tfbridge.DocInfo{
					Source: "function_trigger.md",
				},
			},
			"scaleway_lb_acl": {
				Tok: scalewayResource(scalewayMod, "LoadbalancerAcl"),
				Docs: &tfbridge.DocInfo{
					Source: "lb_acl.md",
				},
			},
			"scaleway_webhosting": {
				Tok: scalewayResource(scalewayMod, "Webhosting"),
				Docs: &tfbridge.DocInfo{
					Source: "webhosting.md",
				},
			},
			"scaleway_vpc": {
				Tok: scalewayResource(scalewayMod, "Vpc"),
				Docs: &tfbridge.DocInfo{
					Source: "vpc.md",
				},
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"scaleway_account_ssh_key": {
				Tok: scalewayDataSource(scalewayMod, "getAccountSshKey"),
				Docs: &tfbridge.DocInfo{
					Source: "account_ssh_key.md",
				},
			},
			"scaleway_account_project": {
				Tok: scalewayDataSource(scalewayMod, "getAccountProject"),
				Docs: &tfbridge.DocInfo{
					Source: "account_project.md",
				},
			},
			"scaleway_baremetal_offer": {
				Tok: scalewayDataSource(scalewayMod, "getBaremetalOffer"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"cpu": {
						MaxItemsOne: boolRef(true),
					},
				},
				Docs: &tfbridge.DocInfo{
					Source: "baremetal_offer.md",
				},
			},
			"scaleway_baremetal_os": {
				Tok: scalewayDataSource(scalewayMod, "getBaremetalOs"),
				Docs: &tfbridge.DocInfo{
					Source: "baremetal_os.md",
				},
			},
			"scaleway_baremetal_option": {
				Tok: scalewayDataSource(scalewayMod, "getBaremetalOption"),
				Docs: &tfbridge.DocInfo{
					Source: "baremetal_option.md",
				},
			},
			"scaleway_baremetal_server": {
				Tok: scalewayDataSource(scalewayMod, "getBaremetalServer"),
				Docs: &tfbridge.DocInfo{
					Source: "barmetal_server.md",
				},
			},
			"scaleway_block_snapshot": {
				Tok: scalewayDataSource(scalewayMod, "getBlockSnapshot"),
				Docs: &tfbridge.DocInfo{
					Source: "bloc_snapshot.md",
				},
			},
			"scaleway_block_volume": {
				Tok: scalewayDataSource(scalewayMod, "getBlockVolume"),
				Docs: &tfbridge.DocInfo{
					Source: "block_volume.md",
				},
			},
			"scaleway_billing_consumptions": {
				Tok: scalewayDataSource(scalewayMod, "getBillingConsumptions"),
				Docs: &tfbridge.DocInfo{
					Source: "billing_consumption.md",
				},
			},
			"scaleway_billing_invoices": {
				Tok: scalewayDataSource(scalewayMod, "getBillingInvoices"),
				Docs: &tfbridge.DocInfo{
					Source: "billing_invoices.md",
				},
			},
			"scaleway_cockpit": {
				Tok: scalewayDataSource(scalewayMod, "getCockpit"),
				Docs: &tfbridge.DocInfo{
					Source: "cockpit.md",
				},
			},
			"scaleway_config": {
				Tok: scalewayDataSource(scalewayMod, "getConfig"),
				Docs: &tfbridge.DocInfo{
					// Source: "config.md", // File doesn't exist in upstream TF provider repo
					AllowMissing: true,
				},
			},
			"scaleway_container": {
				Tok: scalewayDataSource(scalewayMod, "getContainer"),
				Docs: &tfbridge.DocInfo{
					Source: "container.md",
				},
			},
			"scaleway_container_namespace": {
				Tok: scalewayDataSource(scalewayMod, "getContainerNamespace"),
				Docs: &tfbridge.DocInfo{
					Source: "container_namespace.md",
				},
			},
			"scaleway_domain_record": {
				Tok: scalewayDataSource(scalewayMod, "getDomainRecord"),
				Docs: &tfbridge.DocInfo{
					Source: "domain_record.md",
				},
			},
			"scaleway_domain_zone": {
				Tok: scalewayDataSource(scalewayMod, "getDomainZone"),
				Docs: &tfbridge.DocInfo{
					Source: "domain_zone.md",
				},
			},
			"scaleway_documentdb_database": {
				Tok: scalewayDataSource(scalewayMod, "getDocumentdbDatabase"),
				Docs: &tfbridge.DocInfo{
					Source: "documentdb_database.md",
				},
			},
			"scaleway_documentdb_instance": {
				Tok: scalewayDataSource(scalewayMod, "getDocumentdbInstance"),
				Docs: &tfbridge.DocInfo{
					Source: "documentdb_instance.md",
				},
			},
			"scaleway_documentdb_load_balancer_endpoint": {
				Tok: scalewayDataSource(scalewayMod, "getDocumentdbLoadBalancerEndpoint"),
				Docs: &tfbridge.DocInfo{
					Source: "documentdb_load_balancer_endpoint.md",
				},
			},
			"scaleway_flexible_ip": {
				Tok: scalewayDataSource(scalewayMod, "getFlexibleIp"),
				Docs: &tfbridge.DocInfo{
					Source: "flexible_ip.md",
				},
			},
			"scaleway_flexible_ips": {
				Tok: scalewayDataSource(scalewayMod, "getFlexibleIps"),
				Docs: &tfbridge.DocInfo{
					Source: "flexible_ips.md",
				},
			},
			"scaleway_function": {
				Tok: scalewayDataSource(scalewayMod, "getFunction"),
				Docs: &tfbridge.DocInfo{
					Source: "function.md",
				},
			},
			"scaleway_function_namespace": {
				Tok: scalewayDataSource(scalewayMod, "getFunctionNamespace"),
				Docs: &tfbridge.DocInfo{
					Source: "function_namespace.md",
				},
			},
			"scaleway_instance_image": {
				Tok: scalewayDataSource(scalewayMod, "getInstanceImage"),
				Docs: &tfbridge.DocInfo{
					Source: "instance_image.md",
				},
			},
			"scaleway_instance_placement_group": {
				Tok: scalewayDataSource(scalewayMod, "getInstancePlacementGroup"),
				Docs: &tfbridge.DocInfo{
					Source: "instance_placement_group.md",
				},
			},
			"scaleway_iam_application": {
				Tok: scalewayDataSource(scalewayMod, "getIamApplication"),
				Docs: &tfbridge.DocInfo{
					Source: "iam_application.md",
				},
			},
			"scaleway_iam_group": {
				Tok: scalewayDataSource(scalewayMod, "getIamGroup"),
				Docs: &tfbridge.DocInfo{
					Source: "iam_group.md",
				},
			},
			"scaleway_iam_ssh_key": {
				Tok: scalewayDataSource(scalewayMod, "getIamSshKey"),
				Docs: &tfbridge.DocInfo{
					Source: "iam_ssh_key.md",
				},
			},
			"scaleway_iam_user": {
				Tok: scalewayDataSource(scalewayMod, "getIamUser"),
				Docs: &tfbridge.DocInfo{
					Source: "iam_user.md",
				},
			},
			"scaleway_ipam_ip": {
				Tok: scalewayDataSource(scalewayMod, "getIpamIp"),
				Docs: &tfbridge.DocInfo{
					Source: "ipam_ip.md",
				},
			},
			"scaleway_ipam_ips": {
				Tok: scalewayDataSource(scalewayMod, "getIpamIps"),
				Docs: &tfbridge.DocInfo{
					Source: "ipam_ips.md",
				},
			},
			"scaleway_instance_ip": {
				Tok: scalewayDataSource(scalewayMod, "getInstanceIp"),
				Docs: &tfbridge.DocInfo{
					Source: "instance_ip.md",
				},
			},
			"scaleway_instance_private_nic": {
				Tok: scalewayDataSource(scalewayMod, "getInstancePrivateNic"),
				Docs: &tfbridge.DocInfo{
					Source: "instance_private_nic.md",
				},
			},
			"scaleway_instance_security_group": {
				Tok: scalewayDataSource(scalewayMod, "getInstanceSecurityGroup"),
				Docs: &tfbridge.DocInfo{
					Source: "instance_security_group.md",
				},
			},
			"scaleway_instance_server": {
				Tok: scalewayDataSource(scalewayMod, "getInstanceServer"),
				Docs: &tfbridge.DocInfo{
					Source: "instance_server.md",
				},
			},
			"scaleway_instance_servers": {
				Tok: scalewayDataSource(scalewayMod, "getInstanceServers"),
				Docs: &tfbridge.DocInfo{
					Source: "instance_servers.md",
				},
			},
			"scaleway_instance_snapshot": {
				Tok: scalewayDataSource(scalewayMod, "getInstanceSnapshot"),
				Docs: &tfbridge.DocInfo{
					Source: "instance_snapshot.md",
				},
			},
			"scaleway_instance_volume": {
				Tok: scalewayDataSource(scalewayMod, "getInstanceVolume"),
				Docs: &tfbridge.DocInfo{
					Source: "instance_volume.md",
				},
			},
			"scaleway_iot_device": {
				Tok: scalewayDataSource(scalewayMod, "getIotDevice"),
				Docs: &tfbridge.DocInfo{
					Source: "iot_device.md",
				},
			},
			"scaleway_iot_hub": {
				Tok: scalewayDataSource(scalewayMod, "getIotHub"),
				Docs: &tfbridge.DocInfo{
					Source: "iot_hub.md",
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
			"scaleway_object_bucket": {
				Tok: scalewayDataSource(scalewayMod, "getObjectBucket"),
				Docs: &tfbridge.DocInfo{
					Source: "object_bucket.md",
				},
			},
			"scaleway_marketplace_image": {
				Tok: scalewayDataSource(scalewayMod, "getMarketplaceImage"),
				Docs: &tfbridge.DocInfo{
					Source: "marketplace_image.md",
				},
			},
			"scaleway_mnq_sns": {
				Tok: scalewayDataSource(scalewayMod, "getMnqSns"),
				Docs: &tfbridge.DocInfo{
					Source: "mnq_sns.md",
				},
			},
			"scaleway_mnq_sqs": {
				Tok: scalewayDataSource(scalewayMod, "getMnqSqs"),
				Docs: &tfbridge.DocInfo{
					Source: "mnq_sqs.md",
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
			"scaleway_redis_cluster": {
				Tok: scalewayDataSource(scalewayMod, "getRedisCluster"),
				Docs: &tfbridge.DocInfo{
					Source: "redis_cluster.md",
				},
			},
			"scaleway_registry_image": {
				Tok: scalewayDataSource(scalewayMod, "getRegistryImage"),
				Docs: &tfbridge.DocInfo{
					Source: "registry_image.md",
				},
			},
			"scaleway_registry_namespace": {
				Tok: scalewayDataSource(scalewayMod, "getRegistryNamespace"),
				Docs: &tfbridge.DocInfo{
					Source: "registry_namespace.md",
				},
			},
			"scaleway_secret": {
				Tok: scalewayDataSource(scalewayMod, "getSecret"),
				Docs: &tfbridge.DocInfo{
					Source: "secret.md",
				},
			},
			"scaleway_secret_version": {
				Tok: scalewayDataSource(scalewayMod, "getSecretVersion"),
				Docs: &tfbridge.DocInfo{
					Source: "secret_version.md",
				},
			},
			"scaleway_tem_domain": {
				Tok: scalewayDataSource(scalewayMod, "getTemDomain"),
				Docs: &tfbridge.DocInfo{
					Source: "tem_domain.md",
				},
			},
			"scaleway_vpc_gateway_network": {
				Tok: scalewayDataSource(scalewayMod, "getVpcGatewayNetwork"),
				Docs: &tfbridge.DocInfo{
					Source: "vpc_gateway_network.md",
				},
			},
			"scaleway_vpc_private_network": {
				Tok: scalewayDataSource(scalewayMod, "getVpcPrivateNetwork"),
				Docs: &tfbridge.DocInfo{
					Source: "vpc_private_network.md",
				},
			},
			"scaleway_vpc_public_gateway": {
				Tok: scalewayDataSource(scalewayMod, "getVpcPublicGateway"),
				Docs: &tfbridge.DocInfo{
					Source: "vpc_public_gateway.md",
				},
			},
			"scaleway_vpc_public_gateway_dhcp": {
				Tok: scalewayDataSource(scalewayMod, "getVpcPublicGatewayDhcp"),
				Docs: &tfbridge.DocInfo{
					Source: "vpc_public_gateway_dhcp.md",
				},
			},
			"scaleway_vpc_public_gateway_dhcp_reservation": {
				Tok: scalewayDataSource(scalewayMod, "getVpcPublicGatewayDhcpReservation"),
				Docs: &tfbridge.DocInfo{
					Source: "vpc_public_gateway_dhcp_reservation.md",
				},
			},
			"scaleway_vpc_public_gateway_ip": {
				Tok: scalewayDataSource(scalewayMod, "getVpcPublicGatewayIp"),
				Docs: &tfbridge.DocInfo{
					Source: "vpc_public_gateway_ip.md",
				},
			},
			"scaleway_availability_zones": {
				Tok: scalewayDataSource(scalewayMod, "getAvailabilityZones"),
				Docs: &tfbridge.DocInfo{
					Source: "availability_zones.md",
				},
			},
			"scaleway_cockpit_plan": {
				Tok: scalewayDataSource(scalewayMod, "getCockpitPlan"),
				Docs: &tfbridge.DocInfo{
					Source: "cockpit_plan.md",
				},
			},
			"scaleway_object_bucket_policy": {
				Tok: scalewayDataSource(scalewayMod, "getObjectBucketPolicy"),
				Docs: &tfbridge.DocInfo{
					Source: "object_bucket_policy.md",
				},
			},
			"scaleway_vpc": {
				Tok: scalewayDataSource(scalewayMod, "getVpc"),
				Docs: &tfbridge.DocInfo{
					Source: "vpc.md",
				},
			},
			"scaleway_vpcs": {
				Tok: scalewayDataSource(scalewayMod, "getVpcs"),
				Docs: &tfbridge.DocInfo{
					Source: "vpcs.md",
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
			"scaleway_webhosting": {
				Tok: scalewayDataSource(scalewayMod, "getWebhosting"),
				Docs: &tfbridge.DocInfo{
					Source: "webhosting.md",
				},
			},
			"scaleway_iam_api_key": {
				Tok: scalewayDataSource(scalewayMod, "getIamApiKey"),
				Docs: &tfbridge.DocInfo{
					Source: "iam_api_key.md",
				},
			},
			"scaleway_vpc_routes": {
				Tok: scalewayDataSource(scalewayMod, "getVpcRoutes"),
				Docs: &tfbridge.DocInfo{
					Source: "vpc_routes.md",
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

	prov.SetAutonaming(255, "-")

	return prov
}
