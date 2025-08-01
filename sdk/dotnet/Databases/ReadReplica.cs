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
    /// Creates and manages Read Replicas.
    /// For more information refer to the [API documentation](https://www.scaleway.com/en/developers/api/managed-database-postgre-mysql/).
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
    ///     var instance = new Scaleway.Databases.Instance("instance", new()
    ///     {
    ///         Name = "test-rdb-rr-update",
    ///         NodeType = "db-dev-s",
    ///         Engine = "PostgreSQL-14",
    ///         IsHaCluster = false,
    ///         DisableBackup = true,
    ///         UserName = "my_initial_user",
    ///         Password = "thiZ_is_v&amp;ry_s3cret",
    ///         Tags = new[]
    ///         {
    ///             "terraform-test",
    ///             "scaleway_rdb_read_replica",
    ///             "minimal",
    ///         },
    ///     });
    /// 
    ///     var replica = new Scaleway.Databases.ReadReplica("replica", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         DirectAccess = null,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Private network with static endpoint
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var instance = new Scaleway.Databases.Instance("instance", new()
    ///     {
    ///         Name = "rdb_instance",
    ///         NodeType = "db-dev-s",
    ///         Engine = "PostgreSQL-14",
    ///         IsHaCluster = false,
    ///         DisableBackup = true,
    ///         UserName = "my_initial_user",
    ///         Password = "thiZ_is_v&amp;ry_s3cret",
    ///     });
    /// 
    ///     var pn = new Scaleway.Network.PrivateNetwork("pn");
    /// 
    ///     var replica = new Scaleway.Databases.ReadReplica("replica", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         PrivateNetwork = new Scaleway.Databases.Inputs.ReadReplicaPrivateNetworkArgs
    ///         {
    ///             PrivateNetworkId = pn.Id,
    ///             ServiceIp = "192.168.1.254/24",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Private network with IPAM
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var instance = new Scaleway.Databases.Instance("instance", new()
    ///     {
    ///         Name = "rdb_instance",
    ///         NodeType = "db-dev-s",
    ///         Engine = "PostgreSQL-14",
    ///         IsHaCluster = false,
    ///         DisableBackup = true,
    ///         UserName = "my_initial_user",
    ///         Password = "thiZ_is_v&amp;ry_s3cret",
    ///     });
    /// 
    ///     var pn = new Scaleway.Network.PrivateNetwork("pn");
    /// 
    ///     var replica = new Scaleway.Databases.ReadReplica("replica", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         PrivateNetwork = new Scaleway.Databases.Inputs.ReadReplicaPrivateNetworkArgs
    ///         {
    ///             PrivateNetworkId = pn.Id,
    ///             EnableIpam = true,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Read Replicas can be imported using the `{region}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:databases/readReplica:ReadReplica rr fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:databases/readReplica:ReadReplica")]
    public partial class ReadReplica : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Creates a direct access endpoint to rdb replica.
        /// </summary>
        [Output("directAccess")]
        public Output<Outputs.ReadReplicaDirectAccess?> DirectAccess { get; private set; } = null!;

        /// <summary>
        /// UUID of the rdb instance.
        /// 
        /// &gt; **Important:** The replica musts contains at least one `direct_access` or `private_network`. It can contain both.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Create an endpoint in a Private Netork.
        /// </summary>
        [Output("privateNetwork")]
        public Output<Outputs.ReadReplicaPrivateNetwork?> PrivateNetwork { get; private set; } = null!;

        /// <summary>
        /// `region`) The region
        /// in which the Read Replica should be created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Defines whether to create the replica in the same availability zone as the main instance nodes or not.
        /// </summary>
        [Output("sameZone")]
        public Output<bool?> SameZone { get; private set; } = null!;


        /// <summary>
        /// Create a ReadReplica resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReadReplica(string name, ReadReplicaArgs args, CustomResourceOptions? options = null)
            : base("scaleway:databases/readReplica:ReadReplica", name, args ?? new ReadReplicaArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReadReplica(string name, Input<string> id, ReadReplicaState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:databases/readReplica:ReadReplica", name, state, MakeResourceOptions(options, id))
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
                    new global::Pulumi.Alias { Type = "scaleway:index/databaseReadReplica:DatabaseReadReplica" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ReadReplica resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReadReplica Get(string name, Input<string> id, ReadReplicaState? state = null, CustomResourceOptions? options = null)
        {
            return new ReadReplica(name, id, state, options);
        }
    }

    public sealed class ReadReplicaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creates a direct access endpoint to rdb replica.
        /// </summary>
        [Input("directAccess")]
        public Input<Inputs.ReadReplicaDirectAccessArgs>? DirectAccess { get; set; }

        /// <summary>
        /// UUID of the rdb instance.
        /// 
        /// &gt; **Important:** The replica musts contains at least one `direct_access` or `private_network`. It can contain both.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Create an endpoint in a Private Netork.
        /// </summary>
        [Input("privateNetwork")]
        public Input<Inputs.ReadReplicaPrivateNetworkArgs>? PrivateNetwork { get; set; }

        /// <summary>
        /// `region`) The region
        /// in which the Read Replica should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Defines whether to create the replica in the same availability zone as the main instance nodes or not.
        /// </summary>
        [Input("sameZone")]
        public Input<bool>? SameZone { get; set; }

        public ReadReplicaArgs()
        {
        }
        public static new ReadReplicaArgs Empty => new ReadReplicaArgs();
    }

    public sealed class ReadReplicaState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creates a direct access endpoint to rdb replica.
        /// </summary>
        [Input("directAccess")]
        public Input<Inputs.ReadReplicaDirectAccessGetArgs>? DirectAccess { get; set; }

        /// <summary>
        /// UUID of the rdb instance.
        /// 
        /// &gt; **Important:** The replica musts contains at least one `direct_access` or `private_network`. It can contain both.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Create an endpoint in a Private Netork.
        /// </summary>
        [Input("privateNetwork")]
        public Input<Inputs.ReadReplicaPrivateNetworkGetArgs>? PrivateNetwork { get; set; }

        /// <summary>
        /// `region`) The region
        /// in which the Read Replica should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Defines whether to create the replica in the same availability zone as the main instance nodes or not.
        /// </summary>
        [Input("sameZone")]
        public Input<bool>? SameZone { get; set; }

        public ReadReplicaState()
        {
        }
        public static new ReadReplicaState Empty => new ReadReplicaState();
    }
}
