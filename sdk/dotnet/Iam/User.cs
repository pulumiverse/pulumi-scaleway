// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Iam
{
    /// <summary>
    /// Creates and manages Scaleway IAM Users.
    /// For more information, see the [API documentation](https://www.scaleway.com/en/developers/api/iam/#path-users-list-users-of-an-organization).
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
    ///     var basic = new Scaleway.Iam.User("basic", new()
    ///     {
    ///         Email = "test@test.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// IAM users can be imported using the `{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:iam/user:User basic 11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:iam/user:User")]
    public partial class User : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the account root user associated with the user.
        /// </summary>
        [Output("accountRootUserId")]
        public Output<string> AccountRootUserId { get; private set; } = null!;

        /// <summary>
        /// The date and time of the creation of the IAM user.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Whether the IAM user is deletable.
        /// </summary>
        [Output("deletable")]
        public Output<bool> Deletable { get; private set; } = null!;

        /// <summary>
        /// The email of the IAM user.
        /// </summary>
        [Output("email")]
        public Output<string> Email { get; private set; } = null!;

        /// <summary>
        /// The date of the last login.
        /// </summary>
        [Output("lastLoginAt")]
        public Output<string> LastLoginAt { get; private set; } = null!;

        /// <summary>
        /// Whether the MFA is enabled.
        /// </summary>
        [Output("mfa")]
        public Output<bool> Mfa { get; private set; } = null!;

        /// <summary>
        /// `organization_id`) The ID of the organization the user is associated with.
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// The status of user invitation. Check the possible values in the [API doc](https://www.scaleway.com/en/developers/api/iam/#path-users-get-a-given-user).
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The tags associated with the user.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of user. Check the possible values in the [API doc](https://www.scaleway.com/en/developers/api/iam/#path-users-get-a-given-user).
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The date and time of the last update of the IAM user.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs args, CustomResourceOptions? options = null)
            : base("scaleway:iam/user:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:iam/user:User", name, state, MakeResourceOptions(options, id))
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
                    new global::Pulumi.Alias { Type = "scaleway:index/iamUser:IamUser" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
        {
            return new User(name, id, state, options);
        }
    }

    public sealed class UserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The email of the IAM user.
        /// </summary>
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        /// <summary>
        /// `organization_id`) The ID of the organization the user is associated with.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the user.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public UserArgs()
        {
        }
        public static new UserArgs Empty => new UserArgs();
    }

    public sealed class UserState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the account root user associated with the user.
        /// </summary>
        [Input("accountRootUserId")]
        public Input<string>? AccountRootUserId { get; set; }

        /// <summary>
        /// The date and time of the creation of the IAM user.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Whether the IAM user is deletable.
        /// </summary>
        [Input("deletable")]
        public Input<bool>? Deletable { get; set; }

        /// <summary>
        /// The email of the IAM user.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// The date of the last login.
        /// </summary>
        [Input("lastLoginAt")]
        public Input<string>? LastLoginAt { get; set; }

        /// <summary>
        /// Whether the MFA is enabled.
        /// </summary>
        [Input("mfa")]
        public Input<bool>? Mfa { get; set; }

        /// <summary>
        /// `organization_id`) The ID of the organization the user is associated with.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// The status of user invitation. Check the possible values in the [API doc](https://www.scaleway.com/en/developers/api/iam/#path-users-get-a-given-user).
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the user.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of user. Check the possible values in the [API doc](https://www.scaleway.com/en/developers/api/iam/#path-users-get-a-given-user).
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The date and time of the last update of the IAM user.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public UserState()
        {
        }
        public static new UserState Empty => new UserState();
    }
}
