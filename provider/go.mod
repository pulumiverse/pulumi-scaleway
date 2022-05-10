module github.com/jaxxstorm/pulumi-scaleway/provider

go 1.16

require (
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.23.0
	github.com/pulumi/pulumi/sdk/v3 v3.32.1
	github.com/scaleway/terraform-provider-scaleway/v2 v2.2.1
)

require (
	github.com/ulikunitz/xz v0.5.8 // indirect
	golang.org/x/mod v0.5.1 // indirect
	golang.org/x/tools v0.1.7 // indirect
)

replace (
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.8.1
	github.com/hashicorp/go-getter => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20211230170131-3a7c83bfab87

)
