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
    /// Creates and manages Scaleway Database instance authorized IPs.
    /// For more information refer to the [API documentation](https://www.scaleway.com/en/developers/api/managed-database-postgre-mysql/#acl-rules-allowed-ips).
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
    ///     var main = new Scaleway.Databases.Instance("main", new()
    ///     {
    ///         Name = "test-rdb",
    ///         NodeType = "DB-DEV-S",
    ///         Engine = "PostgreSQL-15",
    ///         IsHaCluster = true,
    ///         DisableBackup = true,
    ///         UserName = "my_initial_user",
    ///         Password = "thiZ_is_v&amp;ry_s3cret",
    ///     });
    /// 
    ///     var mainAcl = new Scaleway.Databases.Acl("main", new()
    ///     {
    ///         InstanceId = main.Id,
    ///         AclRules = new[]
    ///         {
    ///             new Scaleway.Databases.Inputs.AclAclRuleArgs
    ///             {
    ///                 Ip = "1.2.3.4/32",
    ///                 Description = "foo",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Database Instance can be imported using the `{region}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:databases/acl:Acl acl01 fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:databases/acl:Acl")]
    public partial class Acl : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of ACLs (structure is described below)
        /// </summary>
        [Output("aclRules")]
        public Output<ImmutableArray<Outputs.AclAclRule>> AclRules { get; private set; } = null!;

        /// <summary>
        /// UUID of the Database Instance.
        /// 
        /// &gt; **Important:** Updates to `instance_id` will recreate the Database ACL.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// `region`) The region in which the Database Instance should be created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a Acl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Acl(string name, AclArgs args, CustomResourceOptions? options = null)
            : base("scaleway:databases/acl:Acl", name, args ?? new AclArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Acl(string name, Input<string> id, AclState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:databases/acl:Acl", name, state, MakeResourceOptions(options, id))
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
                    new global::Pulumi.Alias { Type = "scaleway:index/databaseAcl:DatabaseAcl" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Acl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Acl Get(string name, Input<string> id, AclState? state = null, CustomResourceOptions? options = null)
        {
            return new Acl(name, id, state, options);
        }
    }

    public sealed class AclArgs : global::Pulumi.ResourceArgs
    {
        [Input("aclRules", required: true)]
        private InputList<Inputs.AclAclRuleArgs>? _aclRules;

        /// <summary>
        /// A list of ACLs (structure is described below)
        /// </summary>
        public InputList<Inputs.AclAclRuleArgs> AclRules
        {
            get => _aclRules ?? (_aclRules = new InputList<Inputs.AclAclRuleArgs>());
            set => _aclRules = value;
        }

        /// <summary>
        /// UUID of the Database Instance.
        /// 
        /// &gt; **Important:** Updates to `instance_id` will recreate the Database ACL.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// `region`) The region in which the Database Instance should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public AclArgs()
        {
        }
        public static new AclArgs Empty => new AclArgs();
    }

    public sealed class AclState : global::Pulumi.ResourceArgs
    {
        [Input("aclRules")]
        private InputList<Inputs.AclAclRuleGetArgs>? _aclRules;

        /// <summary>
        /// A list of ACLs (structure is described below)
        /// </summary>
        public InputList<Inputs.AclAclRuleGetArgs> AclRules
        {
            get => _aclRules ?? (_aclRules = new InputList<Inputs.AclAclRuleGetArgs>());
            set => _aclRules = value;
        }

        /// <summary>
        /// UUID of the Database Instance.
        /// 
        /// &gt; **Important:** Updates to `instance_id` will recreate the Database ACL.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// `region`) The region in which the Database Instance should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public AclState()
        {
        }
        public static new AclState Empty => new AclState();
    }
}
