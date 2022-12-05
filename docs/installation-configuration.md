---
title: Scaleway Installation & Configuration
meta_desc: Information on how to install the Scaleway provider.
layout: installation
---

## Installation

The Pulumi Scaleway provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@lbrlabs/scaleway`](https://www.npmjs.com/package/@lbrlabs/pulumi-scaleway)
* Python: [`lbrlabs_scaleway`](https://pypi.org/project/lbrlabs-pulumi-scaleway/)
* Go: [`github.com/lbrlabs/pulumi-scaleway/sdk/go/scaleway`](https://pkg.go.dev/github.com/lbrlabs/pulumi-scaleway/sdk)
* .NET: [`lbrlabs.Scaleway`](https://www.nuget.org/packages/Lbrlabs.PulumiPackage.Scaleway)

### Provider Binary

The Scaleway provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource scaleway <version> --server github://api.github.com/lbrlabs
```

Replace the version string with your desired version.

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

| Option | Required/Optional | Description |
|-----|------|----|
| `access_key`| Required [Scaleway access key](https://console.scaleway.com/project/credentials) |
| `secret_key`| Required | [Scaleway secret key](https://console.scaleway.com/project/credentials) |
| `project_id` | Required | The [project ID](https://console.scaleway.com/project/settings) that will be used as default value for all resources. |
| `region` | Optional | The [project ID](https://console.scaleway.com/project/settings) The [region](https://registry.terraform.io/providers/scaleway/scaleway/latest/guides/regions_and_zones#regions) that will be used as default value for all resources. (fr-par if none specified) |
| `zone` | Optional | The [project ID](https://console.scaleway.com/project/settings) The [zone](https://registry.terraform.io/providers/scaleway/scaleway/latest/guides/regions_and_zones#zones) that will be used as default value for all resources. (fr-par-1 if none specified)
