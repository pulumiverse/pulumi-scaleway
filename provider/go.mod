module github.com/jaxxstorm/pulumi-scaleway/provider

go 1.15

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible

require (
    github.com/hashicorp/terraform-plugin-sdk v1.14.0 
    github.com/pulumi/pulumi-terraform-bridge/v2 v2.11.0
    github.com/pulumi/pulumi/pkg/v2 v2.12.0
	github.com/pulumi/pulumi/sdk/v2 v2.12.0
	github.com/terraform-providers/terraform-provider-scaleway v1.16.0
)
