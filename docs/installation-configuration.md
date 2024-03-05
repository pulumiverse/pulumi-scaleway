---
title: Scaleway Installation & Configuration
meta_desc: Information on how to install the Scaleway provider.
layout: package
---

## Installation

The Pulumi Scaleway provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumiverse/scaleway`](https://www.npmjs.com/package/@pulumiverse/pulumi-scaleway)
* Python: [`pulumiverse_scaleway`](https://pypi.org/project/pulumiverse-scaleway/)
* Go: [`github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway`](https://pkg.go.dev/github.com/pulumiverse/pulumi-scaleway/sdk)
* .NET: [`Pulumiverse.Scaleway`](https://www.nuget.org/packages/Pulumiverse.Scaleway)

### Provider Binary

The Scaleway provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource scaleway <version> --server github://api.github.com/pulumiverse
```

Replace the version string with your desired version.

### Migrating from the LbrLabs package

The maintenance of this provider has been transferred from LbrLabs to Pulumiverse.
LbrLabs published up to v1.11.0, where Pulumiverse picks up with an initial v1.11.1
containing the renamed packages.

If you were using the LbrLabs edition, please update your dependencies to the
Pulumiverse edition:

| Programming Language | LbrLabs name | Pulumiverse name |
| -- | -- | -- |
| JavaScript/TypeScript | `@lbrlabs/pulumi-scaleway` | `@pulumiverse/scaleway` |
| Python | `lbrlabs_pulumi_scaleway` | `pulumiverse_scaleway` |
| Go | `github.com/lbrlabs/pulumi-scaleway/sdk/go/scaleway` | `github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway` |
| .NET | `Lbrlabs.PulumiPackage.Scaleway` | `Pulumiverse.Scaleway` |

## Setup

To provision resources with the Pulumi Scaleway provider, you need to have Scaleway credentials. Scaleway maintains documentation on how to create API keys [here](https://www.scaleway.com/en/docs/console/my-project/how-to/generate-api-key/)

### Set environment variables

Once you have provisioned these credentials, you can set environment variables to provision resources in Scaleway:

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export SCW_ACCESS_KEY=<SCW_ACCESS_KEY>
$ export SCW_SECRET_KEY=<SCW_SECRET_KEY>
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export SCW_ACCESS_KEY=<SCW_ACCESS_KEY>
$ export SCW_SECRET_KEY=<SCW_SECRET_KEY>
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:SCW_ACCESS_KEY = "<SCW_ACCESS_KEY>"
> $env:SCW_SECRET_KEY = "<SCW_SECRET_KEY>"
```

{{% /choosable %}}
{{< /chooser >}}

## Configuration Options

Use `pulumi config set scaleway:<option>` or pass options to the [constructor of `new scaleway.Provider`]({{< relref "/registry/packages/scaleway/api-docs/provider" >}}).

| Option | Environment Variables | Required/Optional | Description |
|-----|------|------|----|
| `access_key`| `SCW_ACCESS_KEY` | Required | [Scaleway access key](https://console.scaleway.com/project/credentials) |
| `secret_key`| `SCW_SECRET_KEY` | Required | [Scaleway secret key](https://console.scaleway.com/project/credentials) |
| `project_id` | `SCW_DEFAULT_PROJECT_ID` | Required | The [project ID](https://console.scaleway.com/project/settings) that will be used as default value for all resources. |
| `region` | `SCW_DEFAULT_REGION` | Optional | The [region](https://registry.terraform.io/providers/scaleway/scaleway/latest/guides/regions_and_zones#regions) that will be used as default value for all resources. (`fr-par` if none specified) |
| `zone` | `SCW_DEFAULT_ZONE` | Optional | The [zone](https://registry.terraform.io/providers/scaleway/scaleway/latest/guides/regions_and_zones#zones) that will be used as default value for all resources. (`fr-par-1` if none specified)
