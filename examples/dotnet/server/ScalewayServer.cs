using Pulumi;
using Pulumi.Scaleway;

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
