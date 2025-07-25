// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Mnq
{
    /// <summary>
    /// Creates and manages Scaleway Messaging and Queuing SNS credentials.
    /// For further information, see
    /// our [main documentation](https://www.scaleway.com/en/docs/messaging/reference-content/sns-overview/)
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
    ///     var main = new Scaleway.Mnq.Sns("main");
    /// 
    ///     var mainSnsCredentials = new Scaleway.Mnq.SnsCredentials("main", new()
    ///     {
    ///         ProjectId = main.ProjectId,
    ///         Name = "sns-credentials",
    ///         Permissions = new Scaleway.Mnq.Inputs.SnsCredentialsPermissionsArgs
    ///         {
    ///             CanManage = false,
    ///             CanReceive = true,
    ///             CanPublish = false,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// SNS credentials can be imported using `{region}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:mnq/snsCredentials:SnsCredentials main fr-par/11111111111111111111111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:mnq/snsCredentials:SnsCredentials")]
    public partial class SnsCredentials : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the key.
        /// </summary>
        [Output("accessKey")]
        public Output<string> AccessKey { get; private set; } = null!;

        /// <summary>
        /// The unique name of the SNS credentials.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// . List of permissions associated with these credentials.
        /// </summary>
        [Output("permissions")]
        public Output<Outputs.SnsCredentialsPermissions> Permissions { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the Project in which SNS is enabled.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// `region`). The region in which SNS is enabled.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The secret value of the key.
        /// </summary>
        [Output("secretKey")]
        public Output<string> SecretKey { get; private set; } = null!;


        /// <summary>
        /// Create a SnsCredentials resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SnsCredentials(string name, SnsCredentialsArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:mnq/snsCredentials:SnsCredentials", name, args ?? new SnsCredentialsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SnsCredentials(string name, Input<string> id, SnsCredentialsState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:mnq/snsCredentials:SnsCredentials", name, state, MakeResourceOptions(options, id))
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
                    new global::Pulumi.Alias { Type = "scaleway:index/mnqSnsCredentials:MnqSnsCredentials" },
                },
                AdditionalSecretOutputs =
                {
                    "accessKey",
                    "secretKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SnsCredentials resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SnsCredentials Get(string name, Input<string> id, SnsCredentialsState? state = null, CustomResourceOptions? options = null)
        {
            return new SnsCredentials(name, id, state, options);
        }
    }

    public sealed class SnsCredentialsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique name of the SNS credentials.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// . List of permissions associated with these credentials.
        /// </summary>
        [Input("permissions")]
        public Input<Inputs.SnsCredentialsPermissionsArgs>? Permissions { get; set; }

        /// <summary>
        /// `project_id`) The ID of the Project in which SNS is enabled.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region in which SNS is enabled.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public SnsCredentialsArgs()
        {
        }
        public static new SnsCredentialsArgs Empty => new SnsCredentialsArgs();
    }

    public sealed class SnsCredentialsState : global::Pulumi.ResourceArgs
    {
        [Input("accessKey")]
        private Input<string>? _accessKey;

        /// <summary>
        /// The ID of the key.
        /// </summary>
        public Input<string>? AccessKey
        {
            get => _accessKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The unique name of the SNS credentials.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// . List of permissions associated with these credentials.
        /// </summary>
        [Input("permissions")]
        public Input<Inputs.SnsCredentialsPermissionsGetArgs>? Permissions { get; set; }

        /// <summary>
        /// `project_id`) The ID of the Project in which SNS is enabled.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region in which SNS is enabled.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("secretKey")]
        private Input<string>? _secretKey;

        /// <summary>
        /// The secret value of the key.
        /// </summary>
        public Input<string>? SecretKey
        {
            get => _secretKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public SnsCredentialsState()
        {
        }
        public static new SnsCredentialsState Empty => new SnsCredentialsState();
    }
}
