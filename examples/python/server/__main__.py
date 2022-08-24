"""A Python Pulumi program"""

import pulumi
import pulumiverse_scaleway as scaleway

public_ip = scaleway.InstanceIp("example")

server = scaleway.InstanceServer("example", 
    image="ubuntu_focal",
    type="DEV1-S",
    ip_id=public_ip.id,
    tags=["python"]
)

