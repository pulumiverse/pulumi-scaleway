// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Loadbalancers
{
    /// <summary>
    /// Creates and manages Scaleway Load Balancer backends.
    /// 
    /// or more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/reference-content/configuring-backends/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-backends).
    /// 
    /// ## Example Usage
    /// 
    /// ### Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var backend01 = new Scaleway.Loadbalancers.Backend("backend01", new()
    ///     {
    ///         LbId = lb01.Id,
    ///         Name = "backend01",
    ///         ForwardProtocol = "http",
    ///         ForwardPort = 80,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### With HTTP Health Check
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var backend01 = new Scaleway.Loadbalancers.Backend("backend01", new()
    ///     {
    ///         LbId = lb01.Id,
    ///         Name = "backend01",
    ///         ForwardProtocol = "http",
    ///         ForwardPort = 80,
    ///         HealthCheckHttp = new Scaleway.Loadbalancers.Inputs.BackendHealthCheckHttpArgs
    ///         {
    ///             Uri = "/health",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Load Balancer backends can be imported using `{zone}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:loadbalancers/backend:Backend backend01 fr-par-1/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:loadbalancers/backend:Backend")]
    public partial class Backend : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Scaleway S3 bucket website to be served in case all backend servers are down **NOTE** : Only the host part of the
        /// Scaleway S3 bucket website is expected. E.g. 'failover-website.s3-website.fr-par.scw.cloud' if your bucket website URL
        /// is 'https://failover-website.s3-website.fr-par.scw.cloud/'.
        /// </summary>
        [Output("failoverHost")]
        public Output<string?> FailoverHost { get; private set; } = null!;

        /// <summary>
        /// User sessions will be forwarded to this port of backend servers
        /// </summary>
        [Output("forwardPort")]
        public Output<int> ForwardPort { get; private set; } = null!;

        /// <summary>
        /// Load balancing algorithm
        /// </summary>
        [Output("forwardPortAlgorithm")]
        public Output<string?> ForwardPortAlgorithm { get; private set; } = null!;

        /// <summary>
        /// Backend protocol
        /// </summary>
        [Output("forwardProtocol")]
        public Output<string> ForwardProtocol { get; private set; } = null!;

        /// <summary>
        /// Interval between two HC requests
        /// </summary>
        [Output("healthCheckDelay")]
        public Output<string?> HealthCheckDelay { get; private set; } = null!;

        [Output("healthCheckHttp")]
        public Output<Outputs.BackendHealthCheckHttp?> HealthCheckHttp { get; private set; } = null!;

        [Output("healthCheckHttps")]
        public Output<Outputs.BackendHealthCheckHttps?> HealthCheckHttps { get; private set; } = null!;

        /// <summary>
        /// Number of allowed failed HC requests before the backend server is marked down
        /// </summary>
        [Output("healthCheckMaxRetries")]
        public Output<int?> HealthCheckMaxRetries { get; private set; } = null!;

        /// <summary>
        /// Port the HC requests will be send to. Default to `forward_port`
        /// </summary>
        [Output("healthCheckPort")]
        public Output<int> HealthCheckPort { get; private set; } = null!;

        /// <summary>
        /// Defines whether proxy protocol should be activated for the health check
        /// </summary>
        [Output("healthCheckSendProxy")]
        public Output<bool?> HealthCheckSendProxy { get; private set; } = null!;

        [Output("healthCheckTcp")]
        public Output<Outputs.BackendHealthCheckTcp> HealthCheckTcp { get; private set; } = null!;

        /// <summary>
        /// Timeout before we consider a HC request failed
        /// </summary>
        [Output("healthCheckTimeout")]
        public Output<string?> HealthCheckTimeout { get; private set; } = null!;

        /// <summary>
        /// Time to wait between two consecutive health checks when a backend server is in a transient state (going UP or DOWN)
        /// </summary>
        [Output("healthCheckTransientDelay")]
        public Output<string?> HealthCheckTransientDelay { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the Load Balancer should check the backend server’s certificate before initiating a connection
        /// </summary>
        [Output("ignoreSslServerVerify")]
        public Output<bool?> IgnoreSslServerVerify { get; private set; } = null!;

        /// <summary>
        /// The load-balancer ID
        /// </summary>
        [Output("lbId")]
        public Output<string> LbId { get; private set; } = null!;

        /// <summary>
        /// Maximum number of connections allowed per backend server
        /// </summary>
        [Output("maxConnections")]
        public Output<int?> MaxConnections { get; private set; } = null!;

        /// <summary>
        /// Number of retries when a backend server connection failed
        /// </summary>
        [Output("maxRetries")]
        public Output<int?> MaxRetries { get; private set; } = null!;

        /// <summary>
        /// The name of the backend
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Modify what occurs when a backend server is marked down
        /// </summary>
        [Output("onMarkedDownAction")]
        public Output<string?> OnMarkedDownAction { get; private set; } = null!;

        /// <summary>
        /// Type of PROXY protocol to enable
        /// </summary>
        [Output("proxyProtocol")]
        public Output<string?> ProxyProtocol { get; private set; } = null!;

        /// <summary>
        /// Whether to use another backend server on each attempt
        /// </summary>
        [Output("redispatchAttemptCount")]
        public Output<int?> RedispatchAttemptCount { get; private set; } = null!;

        /// <summary>
        /// Enables PROXY protocol version 2
        /// </summary>
        [Output("sendProxyV2")]
        public Output<bool> SendProxyV2 { get; private set; } = null!;

        /// <summary>
        /// Backend server IP addresses list (IPv4 or IPv6)
        /// </summary>
        [Output("serverIps")]
        public Output<ImmutableArray<string>> ServerIps { get; private set; } = null!;

        /// <summary>
        /// Enables SSL between load balancer and backend servers
        /// </summary>
        [Output("sslBridging")]
        public Output<bool?> SslBridging { get; private set; } = null!;

        /// <summary>
        /// The type of sticky sessions
        /// </summary>
        [Output("stickySessions")]
        public Output<string?> StickySessions { get; private set; } = null!;

        /// <summary>
        /// Cookie name for sticky sessions
        /// </summary>
        [Output("stickySessionsCookieName")]
        public Output<string?> StickySessionsCookieName { get; private set; } = null!;

        /// <summary>
        /// Maximum initial server connection establishment time
        /// </summary>
        [Output("timeoutConnect")]
        public Output<string?> TimeoutConnect { get; private set; } = null!;

        /// <summary>
        /// Maximum time (in seconds) for a request to be left pending in queue when `max_connections` is reached
        /// </summary>
        [Output("timeoutQueue")]
        public Output<string?> TimeoutQueue { get; private set; } = null!;

        /// <summary>
        /// Maximum server connection inactivity time
        /// </summary>
        [Output("timeoutServer")]
        public Output<string?> TimeoutServer { get; private set; } = null!;

        /// <summary>
        /// Maximum tunnel inactivity time
        /// </summary>
        [Output("timeoutTunnel")]
        public Output<string?> TimeoutTunnel { get; private set; } = null!;


        /// <summary>
        /// Create a Backend resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Backend(string name, BackendArgs args, CustomResourceOptions? options = null)
            : base("scaleway:loadbalancers/backend:Backend", name, args ?? new BackendArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Backend(string name, Input<string> id, BackendState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:loadbalancers/backend:Backend", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "scaleway:index/loadbalancerBackend:LoadbalancerBackend" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Backend resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Backend Get(string name, Input<string> id, BackendState? state = null, CustomResourceOptions? options = null)
        {
            return new Backend(name, id, state, options);
        }
    }

    public sealed class BackendArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Scaleway S3 bucket website to be served in case all backend servers are down **NOTE** : Only the host part of the
        /// Scaleway S3 bucket website is expected. E.g. 'failover-website.s3-website.fr-par.scw.cloud' if your bucket website URL
        /// is 'https://failover-website.s3-website.fr-par.scw.cloud/'.
        /// </summary>
        [Input("failoverHost")]
        public Input<string>? FailoverHost { get; set; }

        /// <summary>
        /// User sessions will be forwarded to this port of backend servers
        /// </summary>
        [Input("forwardPort", required: true)]
        public Input<int> ForwardPort { get; set; } = null!;

        /// <summary>
        /// Load balancing algorithm
        /// </summary>
        [Input("forwardPortAlgorithm")]
        public Input<string>? ForwardPortAlgorithm { get; set; }

        /// <summary>
        /// Backend protocol
        /// </summary>
        [Input("forwardProtocol", required: true)]
        public Input<string> ForwardProtocol { get; set; } = null!;

        /// <summary>
        /// Interval between two HC requests
        /// </summary>
        [Input("healthCheckDelay")]
        public Input<string>? HealthCheckDelay { get; set; }

        [Input("healthCheckHttp")]
        public Input<Inputs.BackendHealthCheckHttpArgs>? HealthCheckHttp { get; set; }

        [Input("healthCheckHttps")]
        public Input<Inputs.BackendHealthCheckHttpsArgs>? HealthCheckHttps { get; set; }

        /// <summary>
        /// Number of allowed failed HC requests before the backend server is marked down
        /// </summary>
        [Input("healthCheckMaxRetries")]
        public Input<int>? HealthCheckMaxRetries { get; set; }

        /// <summary>
        /// Port the HC requests will be send to. Default to `forward_port`
        /// </summary>
        [Input("healthCheckPort")]
        public Input<int>? HealthCheckPort { get; set; }

        /// <summary>
        /// Defines whether proxy protocol should be activated for the health check
        /// </summary>
        [Input("healthCheckSendProxy")]
        public Input<bool>? HealthCheckSendProxy { get; set; }

        [Input("healthCheckTcp")]
        public Input<Inputs.BackendHealthCheckTcpArgs>? HealthCheckTcp { get; set; }

        /// <summary>
        /// Timeout before we consider a HC request failed
        /// </summary>
        [Input("healthCheckTimeout")]
        public Input<string>? HealthCheckTimeout { get; set; }

        /// <summary>
        /// Time to wait between two consecutive health checks when a backend server is in a transient state (going UP or DOWN)
        /// </summary>
        [Input("healthCheckTransientDelay")]
        public Input<string>? HealthCheckTransientDelay { get; set; }

        /// <summary>
        /// Specifies whether the Load Balancer should check the backend server’s certificate before initiating a connection
        /// </summary>
        [Input("ignoreSslServerVerify")]
        public Input<bool>? IgnoreSslServerVerify { get; set; }

        /// <summary>
        /// The load-balancer ID
        /// </summary>
        [Input("lbId", required: true)]
        public Input<string> LbId { get; set; } = null!;

        /// <summary>
        /// Maximum number of connections allowed per backend server
        /// </summary>
        [Input("maxConnections")]
        public Input<int>? MaxConnections { get; set; }

        /// <summary>
        /// Number of retries when a backend server connection failed
        /// </summary>
        [Input("maxRetries")]
        public Input<int>? MaxRetries { get; set; }

        /// <summary>
        /// The name of the backend
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Modify what occurs when a backend server is marked down
        /// </summary>
        [Input("onMarkedDownAction")]
        public Input<string>? OnMarkedDownAction { get; set; }

        /// <summary>
        /// Type of PROXY protocol to enable
        /// </summary>
        [Input("proxyProtocol")]
        public Input<string>? ProxyProtocol { get; set; }

        /// <summary>
        /// Whether to use another backend server on each attempt
        /// </summary>
        [Input("redispatchAttemptCount")]
        public Input<int>? RedispatchAttemptCount { get; set; }

        /// <summary>
        /// Enables PROXY protocol version 2
        /// </summary>
        [Input("sendProxyV2")]
        public Input<bool>? SendProxyV2 { get; set; }

        [Input("serverIps")]
        private InputList<string>? _serverIps;

        /// <summary>
        /// Backend server IP addresses list (IPv4 or IPv6)
        /// </summary>
        public InputList<string> ServerIps
        {
            get => _serverIps ?? (_serverIps = new InputList<string>());
            set => _serverIps = value;
        }

        /// <summary>
        /// Enables SSL between load balancer and backend servers
        /// </summary>
        [Input("sslBridging")]
        public Input<bool>? SslBridging { get; set; }

        /// <summary>
        /// The type of sticky sessions
        /// </summary>
        [Input("stickySessions")]
        public Input<string>? StickySessions { get; set; }

        /// <summary>
        /// Cookie name for sticky sessions
        /// </summary>
        [Input("stickySessionsCookieName")]
        public Input<string>? StickySessionsCookieName { get; set; }

        /// <summary>
        /// Maximum initial server connection establishment time
        /// </summary>
        [Input("timeoutConnect")]
        public Input<string>? TimeoutConnect { get; set; }

        /// <summary>
        /// Maximum time (in seconds) for a request to be left pending in queue when `max_connections` is reached
        /// </summary>
        [Input("timeoutQueue")]
        public Input<string>? TimeoutQueue { get; set; }

        /// <summary>
        /// Maximum server connection inactivity time
        /// </summary>
        [Input("timeoutServer")]
        public Input<string>? TimeoutServer { get; set; }

        /// <summary>
        /// Maximum tunnel inactivity time
        /// </summary>
        [Input("timeoutTunnel")]
        public Input<string>? TimeoutTunnel { get; set; }

        public BackendArgs()
        {
        }
        public static new BackendArgs Empty => new BackendArgs();
    }

    public sealed class BackendState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Scaleway S3 bucket website to be served in case all backend servers are down **NOTE** : Only the host part of the
        /// Scaleway S3 bucket website is expected. E.g. 'failover-website.s3-website.fr-par.scw.cloud' if your bucket website URL
        /// is 'https://failover-website.s3-website.fr-par.scw.cloud/'.
        /// </summary>
        [Input("failoverHost")]
        public Input<string>? FailoverHost { get; set; }

        /// <summary>
        /// User sessions will be forwarded to this port of backend servers
        /// </summary>
        [Input("forwardPort")]
        public Input<int>? ForwardPort { get; set; }

        /// <summary>
        /// Load balancing algorithm
        /// </summary>
        [Input("forwardPortAlgorithm")]
        public Input<string>? ForwardPortAlgorithm { get; set; }

        /// <summary>
        /// Backend protocol
        /// </summary>
        [Input("forwardProtocol")]
        public Input<string>? ForwardProtocol { get; set; }

        /// <summary>
        /// Interval between two HC requests
        /// </summary>
        [Input("healthCheckDelay")]
        public Input<string>? HealthCheckDelay { get; set; }

        [Input("healthCheckHttp")]
        public Input<Inputs.BackendHealthCheckHttpGetArgs>? HealthCheckHttp { get; set; }

        [Input("healthCheckHttps")]
        public Input<Inputs.BackendHealthCheckHttpsGetArgs>? HealthCheckHttps { get; set; }

        /// <summary>
        /// Number of allowed failed HC requests before the backend server is marked down
        /// </summary>
        [Input("healthCheckMaxRetries")]
        public Input<int>? HealthCheckMaxRetries { get; set; }

        /// <summary>
        /// Port the HC requests will be send to. Default to `forward_port`
        /// </summary>
        [Input("healthCheckPort")]
        public Input<int>? HealthCheckPort { get; set; }

        /// <summary>
        /// Defines whether proxy protocol should be activated for the health check
        /// </summary>
        [Input("healthCheckSendProxy")]
        public Input<bool>? HealthCheckSendProxy { get; set; }

        [Input("healthCheckTcp")]
        public Input<Inputs.BackendHealthCheckTcpGetArgs>? HealthCheckTcp { get; set; }

        /// <summary>
        /// Timeout before we consider a HC request failed
        /// </summary>
        [Input("healthCheckTimeout")]
        public Input<string>? HealthCheckTimeout { get; set; }

        /// <summary>
        /// Time to wait between two consecutive health checks when a backend server is in a transient state (going UP or DOWN)
        /// </summary>
        [Input("healthCheckTransientDelay")]
        public Input<string>? HealthCheckTransientDelay { get; set; }

        /// <summary>
        /// Specifies whether the Load Balancer should check the backend server’s certificate before initiating a connection
        /// </summary>
        [Input("ignoreSslServerVerify")]
        public Input<bool>? IgnoreSslServerVerify { get; set; }

        /// <summary>
        /// The load-balancer ID
        /// </summary>
        [Input("lbId")]
        public Input<string>? LbId { get; set; }

        /// <summary>
        /// Maximum number of connections allowed per backend server
        /// </summary>
        [Input("maxConnections")]
        public Input<int>? MaxConnections { get; set; }

        /// <summary>
        /// Number of retries when a backend server connection failed
        /// </summary>
        [Input("maxRetries")]
        public Input<int>? MaxRetries { get; set; }

        /// <summary>
        /// The name of the backend
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Modify what occurs when a backend server is marked down
        /// </summary>
        [Input("onMarkedDownAction")]
        public Input<string>? OnMarkedDownAction { get; set; }

        /// <summary>
        /// Type of PROXY protocol to enable
        /// </summary>
        [Input("proxyProtocol")]
        public Input<string>? ProxyProtocol { get; set; }

        /// <summary>
        /// Whether to use another backend server on each attempt
        /// </summary>
        [Input("redispatchAttemptCount")]
        public Input<int>? RedispatchAttemptCount { get; set; }

        /// <summary>
        /// Enables PROXY protocol version 2
        /// </summary>
        [Input("sendProxyV2")]
        public Input<bool>? SendProxyV2 { get; set; }

        [Input("serverIps")]
        private InputList<string>? _serverIps;

        /// <summary>
        /// Backend server IP addresses list (IPv4 or IPv6)
        /// </summary>
        public InputList<string> ServerIps
        {
            get => _serverIps ?? (_serverIps = new InputList<string>());
            set => _serverIps = value;
        }

        /// <summary>
        /// Enables SSL between load balancer and backend servers
        /// </summary>
        [Input("sslBridging")]
        public Input<bool>? SslBridging { get; set; }

        /// <summary>
        /// The type of sticky sessions
        /// </summary>
        [Input("stickySessions")]
        public Input<string>? StickySessions { get; set; }

        /// <summary>
        /// Cookie name for sticky sessions
        /// </summary>
        [Input("stickySessionsCookieName")]
        public Input<string>? StickySessionsCookieName { get; set; }

        /// <summary>
        /// Maximum initial server connection establishment time
        /// </summary>
        [Input("timeoutConnect")]
        public Input<string>? TimeoutConnect { get; set; }

        /// <summary>
        /// Maximum time (in seconds) for a request to be left pending in queue when `max_connections` is reached
        /// </summary>
        [Input("timeoutQueue")]
        public Input<string>? TimeoutQueue { get; set; }

        /// <summary>
        /// Maximum server connection inactivity time
        /// </summary>
        [Input("timeoutServer")]
        public Input<string>? TimeoutServer { get; set; }

        /// <summary>
        /// Maximum tunnel inactivity time
        /// </summary>
        [Input("timeoutTunnel")]
        public Input<string>? TimeoutTunnel { get; set; }

        public BackendState()
        {
        }
        public static new BackendState Empty => new BackendState();
    }
}
