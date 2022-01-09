package main

import (
	"fmt"
	scaleway "github.com/jaxxstorm/pulumi-scaleway/sdk/go/scaleway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		publicIp, err := scaleway.NewInstanceIP(ctx, "example", &scaleway.InstanceIpArgs{})
		if err != nil {
			return fmt.Errorf("error creating public IP: %v", err)
		}

		_, err = scaleway.NewInstanceServer(ctx, "example", &scaleway.InstanceServerArgs{
			Image: pulumi.String("ubuntu_focal"),
			IpId:  publicIp.ID(),
			Type:  "DEV1-S",
		})
		if err != nil {
			return fmt.Errorf("error creating instance server: %v", err)
		}

		return nil
	})
}
