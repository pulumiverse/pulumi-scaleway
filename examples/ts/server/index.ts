import * as pulumi from "@pulumi/pulumi";
import * as scaleway from "@jaxxstorm/pulumi-scaleway";

const publicIp = new scaleway.InstanceIp("example")
const server = new scaleway.InstanceServer("example", {
    type: "DEV1-S",
    image: "ubuntu_focal",
    ipId: publicIp.id,
})
