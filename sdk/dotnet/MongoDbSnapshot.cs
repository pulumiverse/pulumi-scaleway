// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway
{
    /// <summary>
    /// Creates and manages Scaleway MongoDB® snapshots.
    /// For more information refer to [the API documentation](https://www.scaleway.com/en/docs/managed-databases/mongodb/).
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Scaleway.MongoDbSnapshot("main", new()
    ///     {
    ///         InstanceId = mainScalewayMongodbInstance.Id,
    ///         Name = "name-snapshot",
    ///         ExpiresAt = "2024-12-31T23:59:59Z",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// MongoDB® snapshots can be imported using the `{region}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/mongoDbSnapshot:MongoDbSnapshot main fr-par-1/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/mongoDbSnapshot:MongoDbSnapshot")]
    public partial class MongoDbSnapshot : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The date and time when the MongoDB® snapshot was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The expiration date of the MongoDB® snapshot in ISO 8601 format (e.g. `2024-12-31T23:59:59Z`).
        /// 
        /// &gt; **Important:** Once set, `expires_at` cannot be removed.
        /// </summary>
        [Output("expiresAt")]
        public Output<string> ExpiresAt { get; private set; } = null!;

        /// <summary>
        /// The ID of the MongoDB® instance from which the snapshot was created.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The name of the MongoDB® instance from which the snapshot was created.
        /// </summary>
        [Output("instanceName")]
        public Output<string> InstanceName { get; private set; } = null!;

        /// <summary>
        /// The name of the MongoDB® snapshot.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The type of node associated with the MongoDB® snapshot.
        /// </summary>
        [Output("nodeType")]
        public Output<string> NodeType { get; private set; } = null!;

        /// <summary>
        /// `region`) The region in which the MongoDB® snapshot should be created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The size of the MongoDB® snapshot in bytes.
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        /// <summary>
        /// The date and time of the last update of the MongoDB® snapshot.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// The type of volume used for the MongoDB® snapshot.
        /// </summary>
        [Output("volumeType")]
        public Output<string> VolumeType { get; private set; } = null!;


        /// <summary>
        /// Create a MongoDbSnapshot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MongoDbSnapshot(string name, MongoDbSnapshotArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/mongoDbSnapshot:MongoDbSnapshot", name, args ?? new MongoDbSnapshotArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MongoDbSnapshot(string name, Input<string> id, MongoDbSnapshotState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/mongoDbSnapshot:MongoDbSnapshot", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MongoDbSnapshot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MongoDbSnapshot Get(string name, Input<string> id, MongoDbSnapshotState? state = null, CustomResourceOptions? options = null)
        {
            return new MongoDbSnapshot(name, id, state, options);
        }
    }

    public sealed class MongoDbSnapshotArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The expiration date of the MongoDB® snapshot in ISO 8601 format (e.g. `2024-12-31T23:59:59Z`).
        /// 
        /// &gt; **Important:** Once set, `expires_at` cannot be removed.
        /// </summary>
        [Input("expiresAt", required: true)]
        public Input<string> ExpiresAt { get; set; } = null!;

        /// <summary>
        /// The ID of the MongoDB® instance from which the snapshot was created.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The name of the MongoDB® snapshot.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `region`) The region in which the MongoDB® snapshot should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public MongoDbSnapshotArgs()
        {
        }
        public static new MongoDbSnapshotArgs Empty => new MongoDbSnapshotArgs();
    }

    public sealed class MongoDbSnapshotState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The date and time when the MongoDB® snapshot was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The expiration date of the MongoDB® snapshot in ISO 8601 format (e.g. `2024-12-31T23:59:59Z`).
        /// 
        /// &gt; **Important:** Once set, `expires_at` cannot be removed.
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// The ID of the MongoDB® instance from which the snapshot was created.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The name of the MongoDB® instance from which the snapshot was created.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// The name of the MongoDB® snapshot.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The type of node associated with the MongoDB® snapshot.
        /// </summary>
        [Input("nodeType")]
        public Input<string>? NodeType { get; set; }

        /// <summary>
        /// `region`) The region in which the MongoDB® snapshot should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The size of the MongoDB® snapshot in bytes.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// The date and time of the last update of the MongoDB® snapshot.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// The type of volume used for the MongoDB® snapshot.
        /// </summary>
        [Input("volumeType")]
        public Input<string>? VolumeType { get; set; }

        public MongoDbSnapshotState()
        {
        }
        public static new MongoDbSnapshotState Empty => new MongoDbSnapshotState();
    }
}