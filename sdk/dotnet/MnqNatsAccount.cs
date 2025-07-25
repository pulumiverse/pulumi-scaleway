// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
    /// Creates and manages Scaleway Messaging and Queuing NATS accounts.
    /// For further information, see
    /// our [main documentation](https://www.scaleway.com/en/docs/messaging/reference-content/nats-overview/)
    /// To use the Scaleway provider with the official NATS JetStream provider, check out the corresponding guide.
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
    ///     var main = new Scaleway.Mnq.NatsAccount("main", new()
    ///     {
    ///         Name = "nats-account",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Namespaces can be imported using `{region}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/mnqNatsAccount:MnqNatsAccount main fr-par/11111111111111111111111111111111
    /// ```
    /// </summary>
    [Obsolete(@"scaleway.index/mnqnatsaccount.MnqNatsAccount has been deprecated in favor of scaleway.mnq/natsaccount.NatsAccount")]
    [ScalewayResourceType("scaleway:index/mnqNatsAccount:MnqNatsAccount")]
    public partial class MnqNatsAccount : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The endpoint of the NATS service for this account.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// The unique name of the NATS account.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the Project the
        /// account is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// `region`). The region
        /// in which the account should be created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a MnqNatsAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MnqNatsAccount(string name, MnqNatsAccountArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:index/mnqNatsAccount:MnqNatsAccount", name, args ?? new MnqNatsAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MnqNatsAccount(string name, Input<string> id, MnqNatsAccountState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/mnqNatsAccount:MnqNatsAccount", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MnqNatsAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MnqNatsAccount Get(string name, Input<string> id, MnqNatsAccountState? state = null, CustomResourceOptions? options = null)
        {
            return new MnqNatsAccount(name, id, state, options);
        }
    }

    public sealed class MnqNatsAccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique name of the NATS account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `project_id`) The ID of the Project the
        /// account is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region
        /// in which the account should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public MnqNatsAccountArgs()
        {
        }
        public static new MnqNatsAccountArgs Empty => new MnqNatsAccountArgs();
    }

    public sealed class MnqNatsAccountState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The endpoint of the NATS service for this account.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// The unique name of the NATS account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `project_id`) The ID of the Project the
        /// account is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region
        /// in which the account should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public MnqNatsAccountState()
        {
        }
        public static new MnqNatsAccountState Empty => new MnqNatsAccountState();
    }
}
