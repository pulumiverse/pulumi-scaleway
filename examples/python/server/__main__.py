"""A Python Pulumi program"""

import pulumi
import pulumi_scaleway as scaleway

public_ip = scaleway.InstanceIp("example")

server = scaleway.InstanceServer("example", 
    image="ubuntu_focal",
    type="type",
    ip_id=public_ip.id,
)

