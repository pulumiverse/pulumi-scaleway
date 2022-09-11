package main

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	scaleway "github.com/lbrlabs/pulumi-scaleway/sdk/go/scaleway"
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
