module github.com/jaxxstorm/pulumi-scaleway/provider

go 1.16

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible

require (
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.10.1 // indirect
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.11.0
	github.com/pulumi/pulumi/sdk/v2 v2.25.2
	github.com/pulumi/pulumi/sdk/v3 v3.17.0 // indirect
    github.com/scaleway/terraform-provider-scaleway v2.2.0+incompatible
)

replace (
	github.com/hashicorp/go-getter => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20201218231525-9cca98608a5e
) 
