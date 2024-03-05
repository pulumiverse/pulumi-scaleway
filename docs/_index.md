---
title: Scaleway
meta_desc: Provides an overview of the Scaleway Provider for Pulumi.
layout: package
---

The [Scaleway](https://www.scaleway.com/) provider for Pulumi can be used to provision any of the cloud resources available in Pulumi.
The Scaleway provider must be configured with credentials to deploy and update resources in Scaleway.

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as scaleway from "@pulumiverse/scaleway";
const publicIp = new scaleway.InstanceIp("example")
const server = new scaleway.InstanceServer("example", {
    type: "DEV1-S",
    image: "ubuntu_focal",
    ipId: publicIp.id,
    tags: [
        "typescript"
    ]
})
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumiverse_scaleway as scaleway

public_ip = scaleway.InstanceIp("example")

server = scaleway.InstanceServer("example",
    image="ubuntu_focal",
    type="DEV1-S",
    ip_id=public_ip.id,
    tags=["python"]
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"fmt"
	scaleway "github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		publicIp, err := scaleway.NewInstanceIp(ctx, "example", &scaleway.InstanceIpArgs{})
		if err != nil {
			return fmt.Errorf("error creating public IP: %v", err)
		}

		server, err := scaleway.NewInstanceServer(ctx, "example", &scaleway.InstanceServerArgs{
			Image: pulumi.String("ubuntu_focal"),
			IpId:  publicIp.ID(),
			Type:  pulumi.String("DEV1-S"),
			Tags: pulumi.StringArray{
				pulumi.String("go"),
			},
		})
		if err != nil {
			return fmt.Errorf("error creating instance server: %v", err)
		}

		ctx.Export("server", server.Name)

		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using Pulumi;
using Pulumiverse.Scaleway;

class ScalewayServer : Stack
{
    public ScalewayServer()
    {
        var publicIp = new InstanceIp("example", new InstanceIpArgs{});

        var server = new InstanceServer("example", new InstanceServerArgs{
            Image = "ubuntu_focal",
            IpId = publicIp.Id,
            Type = "DEV1-S",
        });
    }
}
```

{{% /choosable %}}

{{< /chooser >}}

## Issues

This is a community maintained provider. Please file issues and feature requests here:

[pulumiverse/pulumi-scaleway](https://github.com/pulumiverse/pulumi-scaleway/issues)
