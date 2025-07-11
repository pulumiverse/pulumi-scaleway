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
    public static class GetAcl
    {
        /// <summary>
        /// Gets information about the Database Instance network Access Control List.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     // Get the database ACL for the instance id 11111111-1111-1111-1111-111111111111 located in the default region e.g: fr-par
        ///     var myAcl = Scaleway.Databases.GetAcl.Invoke(new()
        ///     {
        ///         InstanceId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetAclResult> InvokeAsync(GetAclArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAclResult>("scaleway:databases/getAcl:getAcl", args ?? new GetAclArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about the Database Instance network Access Control List.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     // Get the database ACL for the instance id 11111111-1111-1111-1111-111111111111 located in the default region e.g: fr-par
        ///     var myAcl = Scaleway.Databases.GetAcl.Invoke(new()
        ///     {
        ///         InstanceId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAclResult> Invoke(GetAclInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAclResult>("scaleway:databases/getAcl:getAcl", args ?? new GetAclInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about the Database Instance network Access Control List.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     // Get the database ACL for the instance id 11111111-1111-1111-1111-111111111111 located in the default region e.g: fr-par
        ///     var myAcl = Scaleway.Databases.GetAcl.Invoke(new()
        ///     {
        ///         InstanceId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAclResult> Invoke(GetAclInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAclResult>("scaleway:databases/getAcl:getAcl", args ?? new GetAclInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAclArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The RDB instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// `region`) The region in which the Database Instance should be created.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetAclArgs()
        {
        }
        public static new GetAclArgs Empty => new GetAclArgs();
    }

    public sealed class GetAclInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The RDB instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// `region`) The region in which the Database Instance should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetAclInvokeArgs()
        {
        }
        public static new GetAclInvokeArgs Empty => new GetAclInvokeArgs();
    }


    [OutputType]
    public sealed class GetAclResult
    {
        /// <summary>
        /// A list of ACLs rules (structure is described below)
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAclAclRuleResult> AclRules;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? Region;

        [OutputConstructor]
        private GetAclResult(
            ImmutableArray<Outputs.GetAclAclRuleResult> aclRules,

            string id,

            string instanceId,

            string? region)
        {
            AclRules = aclRules;
            Id = id;
            InstanceId = instanceId;
            Region = region;
        }
    }
}
