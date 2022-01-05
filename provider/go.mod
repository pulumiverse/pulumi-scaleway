module github.com/jaxxstorm/pulumi-scaleway/provider

go 1.16

require (
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.15.0
	github.com/pulumi/pulumi/sdk/v3 v3.21.0
	github.com/scaleway/terraform-provider-scaleway v1.17.1-0.20211230161835-24df682dbb8e
	github.com/ulikunitz/xz v0.5.8 // indirect
)

replace (
	github.com/hashicorp/go-getter => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20211230170131-3a7c83bfab87
)
