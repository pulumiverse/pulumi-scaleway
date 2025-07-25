// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Databases
{
    /// <summary>
    /// Creates and manages Scaleway Database Instances.
    /// For more information, see refer to the [API documentation](https://www.scaleway.com/en/developers/api/managed-database-postgre-mysql/).
    /// 
    /// ## Example Usage
    /// 
    /// ### Example Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Scaleway.Databases.Instance("main", new()
    ///     {
    ///         Name = "test-rdb",
    ///         NodeType = "DB-DEV-S",
    ///         Engine = "PostgreSQL-15",
    ///         IsHaCluster = true,
    ///         DisableBackup = true,
    ///         UserName = "my_initial_user",
    ///         Password = "thiZ_is_v&amp;ry_s3cret",
    ///         EncryptionAtRest = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Example Block Storage Low Latency
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Scaleway.Databases.Instance("main", new()
    ///     {
    ///         Name = "test-rdb-sbs",
    ///         NodeType = "db-play2-pico",
    ///         Engine = "PostgreSQL-15",
    ///         IsHaCluster = true,
    ///         DisableBackup = true,
    ///         UserName = "my_initial_user",
    ///         Password = "thiZ_is_v&amp;ry_s3cret",
    ///         VolumeType = "sbs_15k",
    ///         VolumeSizeInGb = 10,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Example with Settings
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Scaleway.Databases.Instance("main", new()
    ///     {
    ///         Name = "test-rdb",
    ///         NodeType = "db-dev-s",
    ///         DisableBackup = true,
    ///         Engine = "MySQL-8",
    ///         UserName = "my_initial_user",
    ///         Password = "thiZ_is_v&amp;ry_s3cret",
    ///         InitSettings = 
    ///         {
    ///             { "lower_case_table_names", "1" },
    ///         },
    ///         Settings = 
    ///         {
    ///             { "max_connections", "350" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Example with backup schedule
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Scaleway.Databases.Instance("main", new()
    ///     {
    ///         Name = "test-rdb",
    ///         NodeType = "DB-DEV-S",
    ///         Engine = "PostgreSQL-15",
    ///         IsHaCluster = true,
    ///         UserName = "my_initial_user",
    ///         Password = "thiZ_is_v&amp;ry_s3cret",
    ///         DisableBackup = false,
    ///         BackupScheduleFrequency = 24,
    ///         BackupScheduleRetention = 7,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Examples of endpoint configuration
    /// 
    /// Database Instances can have a maximum of 1 public endpoint and 1 private endpoint. They can have both, or none.
    /// 
    /// ### 1 static Private Network endpoint
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var pn = new Scaleway.Network.PrivateNetwork("pn", new()
    ///     {
    ///         Ipv4Subnet = new Scaleway.Network.Inputs.PrivateNetworkIpv4SubnetArgs
    ///         {
    ///             Subnet = "172.16.20.0/22",
    ///         },
    ///     });
    /// 
    ///     var main = new Scaleway.Databases.Instance("main", new()
    ///     {
    ///         NodeType = "db-dev-s",
    ///         Engine = "PostgreSQL-15",
    ///         PrivateNetwork = new Scaleway.Databases.Inputs.InstancePrivateNetworkArgs
    ///         {
    ///             PnId = pn.Id,
    ///             IpNet = "172.16.20.4/22",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### 1 IPAM Private Network endpoint + 1 public endpoint
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var pn = new Scaleway.Network.PrivateNetwork("pn");
    /// 
    ///     var main = new Scaleway.Databases.Instance("main", new()
    ///     {
    ///         LoadBalancers = new[]
    ///         {
    ///             null,
    ///         },
    ///         NodeType = "DB-DEV-S",
    ///         Engine = "PostgreSQL-15",
    ///         PrivateNetwork = new Scaleway.Databases.Inputs.InstancePrivateNetworkArgs
    ///         {
    ///             PnId = pn.Id,
    ///             EnableIpam = true,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Default: 1 public endpoint
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Scaleway.Databases.Instance("main", new()
    ///     {
    ///         NodeType = "db-dev-s",
    ///         Engine = "PostgreSQL-15",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// &gt; **Note** If nothing is defined, your Database Instance will have a default public load-balancer endpoint.
    /// 
    /// ## Limitations
    /// 
    /// The Managed Database product is only compliant with the Private Network in the default availability zone (AZ).
    /// i.e. `fr-par-1`, `nl-ams-1`, `pl-waw-1`. To learn more, read our
    /// section [How to connect a PostgreSQL and MySQL Database Instance to a Private Network](https://www.scaleway.com/en/docs/managed-databases/postgresql-and-mysql/how-to/connect-database-private-network/)
    /// 
    /// ## Import
    /// 
    /// Database Instance can be imported using the `{region}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:databases/instance:Instance rdb01 fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:databases/instance:Instance")]
    public partial class Instance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Boolean to store logical backups in the same region as the database instance
        /// </summary>
        [Output("backupSameRegion")]
        public Output<bool> BackupSameRegion { get; private set; } = null!;

        /// <summary>
        /// Backup schedule frequency in hours
        /// </summary>
        [Output("backupScheduleFrequency")]
        public Output<int> BackupScheduleFrequency { get; private set; } = null!;

        /// <summary>
        /// Backup schedule retention in days
        /// </summary>
        [Output("backupScheduleRetention")]
        public Output<int> BackupScheduleRetention { get; private set; } = null!;

        /// <summary>
        /// Certificate of the Database Instance.
        /// </summary>
        [Output("certificate")]
        public Output<string> Certificate { get; private set; } = null!;

        /// <summary>
        /// Disable automated backup for the database instance
        /// </summary>
        [Output("disableBackup")]
        public Output<bool?> DisableBackup { get; private set; } = null!;

        /// <summary>
        /// Enable or disable encryption at rest for the Database Instance.
        /// </summary>
        [Output("encryptionAtRest")]
        public Output<bool?> EncryptionAtRest { get; private set; } = null!;

        /// <summary>
        /// (Deprecated) The IP of the Database Instance. Please use the private_network or the load_balancer attribute.
        /// </summary>
        [Output("endpointIp")]
        public Output<string> EndpointIp { get; private set; } = null!;

        /// <summary>
        /// (Deprecated) The port of the Database Instance. Please use the private_network or the load_balancer attribute.
        /// </summary>
        [Output("endpointPort")]
        public Output<int> EndpointPort { get; private set; } = null!;

        /// <summary>
        /// Database Instance's engine version (e.g. `PostgreSQL-11`).
        /// 
        /// &gt; **Important** Updates to `engine` will recreate the Database Instance.
        /// </summary>
        [Output("engine")]
        public Output<string> Engine { get; private set; } = null!;

        /// <summary>
        /// Map of engine settings to be set at database initialisation.
        /// </summary>
        [Output("initSettings")]
        public Output<ImmutableDictionary<string, string>?> InitSettings { get; private set; } = null!;

        /// <summary>
        /// Enable or disable high availability for the Database Instance.
        /// 
        /// &gt; **Important** Updates to `is_ha_cluster` will recreate the Database Instance.
        /// </summary>
        [Output("isHaCluster")]
        public Output<bool?> IsHaCluster { get; private set; } = null!;

        /// <summary>
        /// List of Load Balancer endpoints of the Database Instance.
        /// </summary>
        [Output("loadBalancers")]
        public Output<ImmutableArray<Outputs.InstanceLoadBalancer>> LoadBalancers { get; private set; } = null!;

        /// <summary>
        /// Logs policy configuration
        /// </summary>
        [Output("logsPolicy")]
        public Output<Outputs.InstanceLogsPolicy> LogsPolicy { get; private set; } = null!;

        /// <summary>
        /// The name of the Database Instance.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The type of Database Instance you want to create (e.g. `db-dev-s`).
        /// 
        /// &gt; **Important** Updates to `node_type` will upgrade the Database Instance to the desired `node_type` without any
        /// interruption.
        /// 
        /// &gt; **Important** Once your Database Instance reaches `disk_full` status, if you are using `lssd` storage, you should upgrade the `node_type`, and if you are using `bssd` storage, you should increase the volume size before making any other changes to your Database Instance.
        /// </summary>
        [Output("nodeType")]
        public Output<string> NodeType { get; private set; } = null!;

        /// <summary>
        /// The organization ID the Database Instance is associated with.
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// Password for the first user of the Database Instance.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// The private IPv4 address associated with the resource.
        /// </summary>
        [Output("privateIps")]
        public Output<ImmutableArray<Outputs.InstancePrivateIp>> PrivateIps { get; private set; } = null!;

        /// <summary>
        /// List of Private Networks endpoints of the Database Instance.
        /// </summary>
        [Output("privateNetwork")]
        public Output<Outputs.InstancePrivateNetwork?> PrivateNetwork { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the Database
        /// Instance is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// List of read replicas of the Database Instance.
        /// </summary>
        [Output("readReplicas")]
        public Output<ImmutableArray<Outputs.InstanceReadReplica>> ReadReplicas { get; private set; } = null!;

        /// <summary>
        /// `region`) The region
        /// in which the Database Instance should be created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Map of engine settings to be set on a running instance.
        /// </summary>
        [Output("settings")]
        public Output<ImmutableDictionary<string, string>> Settings { get; private set; } = null!;

        /// <summary>
        /// The ID of an existing snapshot to restore or create the Database Instance from. Conflicts with the `engine` parameter and backup settings.
        /// </summary>
        [Output("snapshotId")]
        public Output<string?> SnapshotId { get; private set; } = null!;

        /// <summary>
        /// The tags associated with the Database Instance.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Identifier for the first user of the Database Instance.
        /// 
        /// &gt; **Important** Updates to `user_name` will recreate the Database Instance.
        /// </summary>
        [Output("userName")]
        public Output<string> UserName { get; private set; } = null!;

        /// <summary>
        /// Volume size (in GB). Cannot be used when `volume_type` is set to `lssd`.
        /// 
        /// &gt; **Important** Once your Database Instance reaches `disk_full` status, you should increase the volume size before making any other change to your Database Instance.
        /// </summary>
        [Output("volumeSizeInGb")]
        public Output<int> VolumeSizeInGb { get; private set; } = null!;

        /// <summary>
        /// Type of volume where data are stored (`lssd`, `sbs_5k` or `sbs_15k`).
        /// </summary>
        [Output("volumeType")]
        public Output<string?> VolumeType { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("scaleway:databases/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:databases/instance:Instance", name, state, MakeResourceOptions(options, id))
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
                    new global::Pulumi.Alias { Type = "scaleway:index/databaseInstance:DatabaseInstance" },
                },
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean to store logical backups in the same region as the database instance
        /// </summary>
        [Input("backupSameRegion")]
        public Input<bool>? BackupSameRegion { get; set; }

        /// <summary>
        /// Backup schedule frequency in hours
        /// </summary>
        [Input("backupScheduleFrequency")]
        public Input<int>? BackupScheduleFrequency { get; set; }

        /// <summary>
        /// Backup schedule retention in days
        /// </summary>
        [Input("backupScheduleRetention")]
        public Input<int>? BackupScheduleRetention { get; set; }

        /// <summary>
        /// Disable automated backup for the database instance
        /// </summary>
        [Input("disableBackup")]
        public Input<bool>? DisableBackup { get; set; }

        /// <summary>
        /// Enable or disable encryption at rest for the Database Instance.
        /// </summary>
        [Input("encryptionAtRest")]
        public Input<bool>? EncryptionAtRest { get; set; }

        /// <summary>
        /// Database Instance's engine version (e.g. `PostgreSQL-11`).
        /// 
        /// &gt; **Important** Updates to `engine` will recreate the Database Instance.
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        [Input("initSettings")]
        private InputMap<string>? _initSettings;

        /// <summary>
        /// Map of engine settings to be set at database initialisation.
        /// </summary>
        public InputMap<string> InitSettings
        {
            get => _initSettings ?? (_initSettings = new InputMap<string>());
            set => _initSettings = value;
        }

        /// <summary>
        /// Enable or disable high availability for the Database Instance.
        /// 
        /// &gt; **Important** Updates to `is_ha_cluster` will recreate the Database Instance.
        /// </summary>
        [Input("isHaCluster")]
        public Input<bool>? IsHaCluster { get; set; }

        [Input("loadBalancers")]
        private InputList<Inputs.InstanceLoadBalancerArgs>? _loadBalancers;

        /// <summary>
        /// List of Load Balancer endpoints of the Database Instance.
        /// </summary>
        public InputList<Inputs.InstanceLoadBalancerArgs> LoadBalancers
        {
            get => _loadBalancers ?? (_loadBalancers = new InputList<Inputs.InstanceLoadBalancerArgs>());
            set => _loadBalancers = value;
        }

        /// <summary>
        /// Logs policy configuration
        /// </summary>
        [Input("logsPolicy")]
        public Input<Inputs.InstanceLogsPolicyArgs>? LogsPolicy { get; set; }

        /// <summary>
        /// The name of the Database Instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The type of Database Instance you want to create (e.g. `db-dev-s`).
        /// 
        /// &gt; **Important** Updates to `node_type` will upgrade the Database Instance to the desired `node_type` without any
        /// interruption.
        /// 
        /// &gt; **Important** Once your Database Instance reaches `disk_full` status, if you are using `lssd` storage, you should upgrade the `node_type`, and if you are using `bssd` storage, you should increase the volume size before making any other changes to your Database Instance.
        /// </summary>
        [Input("nodeType", required: true)]
        public Input<string> NodeType { get; set; } = null!;

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password for the first user of the Database Instance.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("privateIps")]
        private InputList<Inputs.InstancePrivateIpArgs>? _privateIps;

        /// <summary>
        /// The private IPv4 address associated with the resource.
        /// </summary>
        public InputList<Inputs.InstancePrivateIpArgs> PrivateIps
        {
            get => _privateIps ?? (_privateIps = new InputList<Inputs.InstancePrivateIpArgs>());
            set => _privateIps = value;
        }

        /// <summary>
        /// List of Private Networks endpoints of the Database Instance.
        /// </summary>
        [Input("privateNetwork")]
        public Input<Inputs.InstancePrivateNetworkArgs>? PrivateNetwork { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the Database
        /// Instance is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region
        /// in which the Database Instance should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("settings")]
        private InputMap<string>? _settings;

        /// <summary>
        /// Map of engine settings to be set on a running instance.
        /// </summary>
        public InputMap<string> Settings
        {
            get => _settings ?? (_settings = new InputMap<string>());
            set => _settings = value;
        }

        /// <summary>
        /// The ID of an existing snapshot to restore or create the Database Instance from. Conflicts with the `engine` parameter and backup settings.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the Database Instance.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Identifier for the first user of the Database Instance.
        /// 
        /// &gt; **Important** Updates to `user_name` will recreate the Database Instance.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        /// <summary>
        /// Volume size (in GB). Cannot be used when `volume_type` is set to `lssd`.
        /// 
        /// &gt; **Important** Once your Database Instance reaches `disk_full` status, you should increase the volume size before making any other change to your Database Instance.
        /// </summary>
        [Input("volumeSizeInGb")]
        public Input<int>? VolumeSizeInGb { get; set; }

        /// <summary>
        /// Type of volume where data are stored (`lssd`, `sbs_5k` or `sbs_15k`).
        /// </summary>
        [Input("volumeType")]
        public Input<string>? VolumeType { get; set; }

        public InstanceArgs()
        {
        }
        public static new InstanceArgs Empty => new InstanceArgs();
    }

    public sealed class InstanceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean to store logical backups in the same region as the database instance
        /// </summary>
        [Input("backupSameRegion")]
        public Input<bool>? BackupSameRegion { get; set; }

        /// <summary>
        /// Backup schedule frequency in hours
        /// </summary>
        [Input("backupScheduleFrequency")]
        public Input<int>? BackupScheduleFrequency { get; set; }

        /// <summary>
        /// Backup schedule retention in days
        /// </summary>
        [Input("backupScheduleRetention")]
        public Input<int>? BackupScheduleRetention { get; set; }

        /// <summary>
        /// Certificate of the Database Instance.
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        /// <summary>
        /// Disable automated backup for the database instance
        /// </summary>
        [Input("disableBackup")]
        public Input<bool>? DisableBackup { get; set; }

        /// <summary>
        /// Enable or disable encryption at rest for the Database Instance.
        /// </summary>
        [Input("encryptionAtRest")]
        public Input<bool>? EncryptionAtRest { get; set; }

        /// <summary>
        /// (Deprecated) The IP of the Database Instance. Please use the private_network or the load_balancer attribute.
        /// </summary>
        [Input("endpointIp")]
        public Input<string>? EndpointIp { get; set; }

        /// <summary>
        /// (Deprecated) The port of the Database Instance. Please use the private_network or the load_balancer attribute.
        /// </summary>
        [Input("endpointPort")]
        public Input<int>? EndpointPort { get; set; }

        /// <summary>
        /// Database Instance's engine version (e.g. `PostgreSQL-11`).
        /// 
        /// &gt; **Important** Updates to `engine` will recreate the Database Instance.
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        [Input("initSettings")]
        private InputMap<string>? _initSettings;

        /// <summary>
        /// Map of engine settings to be set at database initialisation.
        /// </summary>
        public InputMap<string> InitSettings
        {
            get => _initSettings ?? (_initSettings = new InputMap<string>());
            set => _initSettings = value;
        }

        /// <summary>
        /// Enable or disable high availability for the Database Instance.
        /// 
        /// &gt; **Important** Updates to `is_ha_cluster` will recreate the Database Instance.
        /// </summary>
        [Input("isHaCluster")]
        public Input<bool>? IsHaCluster { get; set; }

        [Input("loadBalancers")]
        private InputList<Inputs.InstanceLoadBalancerGetArgs>? _loadBalancers;

        /// <summary>
        /// List of Load Balancer endpoints of the Database Instance.
        /// </summary>
        public InputList<Inputs.InstanceLoadBalancerGetArgs> LoadBalancers
        {
            get => _loadBalancers ?? (_loadBalancers = new InputList<Inputs.InstanceLoadBalancerGetArgs>());
            set => _loadBalancers = value;
        }

        /// <summary>
        /// Logs policy configuration
        /// </summary>
        [Input("logsPolicy")]
        public Input<Inputs.InstanceLogsPolicyGetArgs>? LogsPolicy { get; set; }

        /// <summary>
        /// The name of the Database Instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The type of Database Instance you want to create (e.g. `db-dev-s`).
        /// 
        /// &gt; **Important** Updates to `node_type` will upgrade the Database Instance to the desired `node_type` without any
        /// interruption.
        /// 
        /// &gt; **Important** Once your Database Instance reaches `disk_full` status, if you are using `lssd` storage, you should upgrade the `node_type`, and if you are using `bssd` storage, you should increase the volume size before making any other changes to your Database Instance.
        /// </summary>
        [Input("nodeType")]
        public Input<string>? NodeType { get; set; }

        /// <summary>
        /// The organization ID the Database Instance is associated with.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password for the first user of the Database Instance.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("privateIps")]
        private InputList<Inputs.InstancePrivateIpGetArgs>? _privateIps;

        /// <summary>
        /// The private IPv4 address associated with the resource.
        /// </summary>
        public InputList<Inputs.InstancePrivateIpGetArgs> PrivateIps
        {
            get => _privateIps ?? (_privateIps = new InputList<Inputs.InstancePrivateIpGetArgs>());
            set => _privateIps = value;
        }

        /// <summary>
        /// List of Private Networks endpoints of the Database Instance.
        /// </summary>
        [Input("privateNetwork")]
        public Input<Inputs.InstancePrivateNetworkGetArgs>? PrivateNetwork { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the Database
        /// Instance is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("readReplicas")]
        private InputList<Inputs.InstanceReadReplicaGetArgs>? _readReplicas;

        /// <summary>
        /// List of read replicas of the Database Instance.
        /// </summary>
        public InputList<Inputs.InstanceReadReplicaGetArgs> ReadReplicas
        {
            get => _readReplicas ?? (_readReplicas = new InputList<Inputs.InstanceReadReplicaGetArgs>());
            set => _readReplicas = value;
        }

        /// <summary>
        /// `region`) The region
        /// in which the Database Instance should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("settings")]
        private InputMap<string>? _settings;

        /// <summary>
        /// Map of engine settings to be set on a running instance.
        /// </summary>
        public InputMap<string> Settings
        {
            get => _settings ?? (_settings = new InputMap<string>());
            set => _settings = value;
        }

        /// <summary>
        /// The ID of an existing snapshot to restore or create the Database Instance from. Conflicts with the `engine` parameter and backup settings.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the Database Instance.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Identifier for the first user of the Database Instance.
        /// 
        /// &gt; **Important** Updates to `user_name` will recreate the Database Instance.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        /// <summary>
        /// Volume size (in GB). Cannot be used when `volume_type` is set to `lssd`.
        /// 
        /// &gt; **Important** Once your Database Instance reaches `disk_full` status, you should increase the volume size before making any other change to your Database Instance.
        /// </summary>
        [Input("volumeSizeInGb")]
        public Input<int>? VolumeSizeInGb { get; set; }

        /// <summary>
        /// Type of volume where data are stored (`lssd`, `sbs_5k` or `sbs_15k`).
        /// </summary>
        [Input("volumeType")]
        public Input<string>? VolumeType { get; set; }

        public InstanceState()
        {
        }
        public static new InstanceState Empty => new InstanceState();
    }
}
