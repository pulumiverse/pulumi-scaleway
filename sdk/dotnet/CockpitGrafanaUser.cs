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
    /// Creates and manages Scaleway Cockpit Grafana Users.
    /// 
    /// For more information consult the [documentation](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#grafana-users).
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumi.Scaleway;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var mainCockpit = Scaleway.GetCockpit.Invoke();
    /// 
    ///     // Create an editor grafana user for the cockpit
    ///     var mainCockpitGrafanaUser = new Scaleway.CockpitGrafanaUser("mainCockpitGrafanaUser", new()
    ///     {
    ///         ProjectId = mainCockpit.Apply(getCockpitResult =&gt; getCockpitResult.ProjectId),
    ///         Login = "my-awesome-user",
    ///         Role = "editor",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Cockpits Grafana Users can be imported using the project ID and the grafana user ID formatted `{project_id}/{grafana_user_id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/cockpitGrafanaUser:CockpitGrafanaUser main 11111111-1111-1111-1111-111111111111/2
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/cockpitGrafanaUser:CockpitGrafanaUser")]
    public partial class CockpitGrafanaUser : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The login of the grafana user.
        /// </summary>
        [Output("login")]
        public Output<string> Login { get; private set; } = null!;

        /// <summary>
        /// The password of the grafana user
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the cockpit is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The role of the grafana user. Must be `editor` or `viewer`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a CockpitGrafanaUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CockpitGrafanaUser(string name, CockpitGrafanaUserArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/cockpitGrafanaUser:CockpitGrafanaUser", name, args ?? new CockpitGrafanaUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CockpitGrafanaUser(string name, Input<string> id, CockpitGrafanaUserState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/cockpitGrafanaUser:CockpitGrafanaUser", name, state, MakeResourceOptions(options, id))
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
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CockpitGrafanaUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CockpitGrafanaUser Get(string name, Input<string> id, CockpitGrafanaUserState? state = null, CustomResourceOptions? options = null)
        {
            return new CockpitGrafanaUser(name, id, state, options);
        }
    }

    public sealed class CockpitGrafanaUserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The login of the grafana user.
        /// </summary>
        [Input("login", required: true)]
        public Input<string> Login { get; set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the cockpit is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The role of the grafana user. Must be `editor` or `viewer`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public CockpitGrafanaUserArgs()
        {
        }
        public static new CockpitGrafanaUserArgs Empty => new CockpitGrafanaUserArgs();
    }

    public sealed class CockpitGrafanaUserState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The login of the grafana user.
        /// </summary>
        [Input("login")]
        public Input<string>? Login { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The password of the grafana user
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

        /// <summary>
        /// `project_id`) The ID of the project the cockpit is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The role of the grafana user. Must be `editor` or `viewer`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public CockpitGrafanaUserState()
        {
        }
        public static new CockpitGrafanaUserState Empty => new CockpitGrafanaUserState();
    }
}