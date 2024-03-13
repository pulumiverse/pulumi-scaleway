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
    /// Creates and manages Scaleway Messaging and queuing SQS Credentials.
    /// For further information please check
    /// our [documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/sqs-overview/)
    /// 
    /// ## Example Usage
    /// 
    /// ### Basic
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var mainMnqSqs = new Scaleway.MnqSqs("mainMnqSqs");
    /// 
    ///     var mainMnqSqsCredentials = new Scaleway.MnqSqsCredentials("mainMnqSqsCredentials", new()
    ///     {
    ///         ProjectId = mainMnqSqs.ProjectId,
    ///         Permissions = new Scaleway.Inputs.MnqSqsCredentialsPermissionsArgs
    ///         {
    ///             CanManage = false,
    ///             CanReceive = true,
    ///             CanPublish = false,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// SQS credentials can be imported using the `{region}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/mnqSqsCredentials:MnqSqsCredentials main fr-par/11111111111111111111111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/mnqSqsCredentials:MnqSqsCredentials")]
    public partial class MnqSqsCredentials : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the key.
        /// </summary>
        [Output("accessKey")]
        public Output<string> AccessKey { get; private set; } = null!;

        /// <summary>
        /// The unique name of the sqs credentials.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// . List of permissions associated to these credentials. Only one of permissions may be set.
        /// </summary>
        [Output("permissions")]
        public Output<Outputs.MnqSqsCredentialsPermissions> Permissions { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the sqs is enabled for.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// `region`). The region in which sqs is enabled.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The secret value of the key.
        /// </summary>
        [Output("secretKey")]
        public Output<string> SecretKey { get; private set; } = null!;


        /// <summary>
        /// Create a MnqSqsCredentials resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MnqSqsCredentials(string name, MnqSqsCredentialsArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:index/mnqSqsCredentials:MnqSqsCredentials", name, args ?? new MnqSqsCredentialsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MnqSqsCredentials(string name, Input<string> id, MnqSqsCredentialsState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/mnqSqsCredentials:MnqSqsCredentials", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
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
        /// Get an existing MnqSqsCredentials resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MnqSqsCredentials Get(string name, Input<string> id, MnqSqsCredentialsState? state = null, CustomResourceOptions? options = null)
        {
            return new MnqSqsCredentials(name, id, state, options);
        }
    }

    public sealed class MnqSqsCredentialsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique name of the sqs credentials.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// . List of permissions associated to these credentials. Only one of permissions may be set.
        /// </summary>
        [Input("permissions")]
        public Input<Inputs.MnqSqsCredentialsPermissionsArgs>? Permissions { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the sqs is enabled for.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region in which sqs is enabled.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public MnqSqsCredentialsArgs()
        {
        }
        public static new MnqSqsCredentialsArgs Empty => new MnqSqsCredentialsArgs();
    }

    public sealed class MnqSqsCredentialsState : global::Pulumi.ResourceArgs
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
        /// The unique name of the sqs credentials.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// . List of permissions associated to these credentials. Only one of permissions may be set.
        /// </summary>
        [Input("permissions")]
        public Input<Inputs.MnqSqsCredentialsPermissionsGetArgs>? Permissions { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the sqs is enabled for.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region in which sqs is enabled.
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

        public MnqSqsCredentialsState()
        {
        }
        public static new MnqSqsCredentialsState Empty => new MnqSqsCredentialsState();
    }
}