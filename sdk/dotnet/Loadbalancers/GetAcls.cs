// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Loadbalancers
{
    public static class GetAcls
    {
        /// <summary>
        /// Gets information about multiple Load Balancer ACLs.
        /// 
        /// For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/reference-content/acls/) or [API reference](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-acls-get-an-acl).
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
        ///     // Find acls that share the same frontend ID
        ///     var byFrontID = Scaleway.Loadbalancers.GetAcls.Invoke(new()
        ///     {
        ///         FrontendId = frt01.Id,
        ///     });
        /// 
        ///     // Find acls by frontend ID and name
        ///     var byFrontIDAndName = Scaleway.Loadbalancers.GetAcls.Invoke(new()
        ///     {
        ///         FrontendId = frt01.Id,
        ///         Name = "tf-acls-datasource",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetAclsResult> InvokeAsync(GetAclsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAclsResult>("scaleway:loadbalancers/getAcls:getAcls", args ?? new GetAclsArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about multiple Load Balancer ACLs.
        /// 
        /// For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/reference-content/acls/) or [API reference](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-acls-get-an-acl).
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
        ///     // Find acls that share the same frontend ID
        ///     var byFrontID = Scaleway.Loadbalancers.GetAcls.Invoke(new()
        ///     {
        ///         FrontendId = frt01.Id,
        ///     });
        /// 
        ///     // Find acls by frontend ID and name
        ///     var byFrontIDAndName = Scaleway.Loadbalancers.GetAcls.Invoke(new()
        ///     {
        ///         FrontendId = frt01.Id,
        ///         Name = "tf-acls-datasource",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAclsResult> Invoke(GetAclsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAclsResult>("scaleway:loadbalancers/getAcls:getAcls", args ?? new GetAclsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about multiple Load Balancer ACLs.
        /// 
        /// For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/reference-content/acls/) or [API reference](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-acls-get-an-acl).
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
        ///     // Find acls that share the same frontend ID
        ///     var byFrontID = Scaleway.Loadbalancers.GetAcls.Invoke(new()
        ///     {
        ///         FrontendId = frt01.Id,
        ///     });
        /// 
        ///     // Find acls by frontend ID and name
        ///     var byFrontIDAndName = Scaleway.Loadbalancers.GetAcls.Invoke(new()
        ///     {
        ///         FrontendId = frt01.Id,
        ///         Name = "tf-acls-datasource",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAclsResult> Invoke(GetAclsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAclsResult>("scaleway:loadbalancers/getAcls:getAcls", args ?? new GetAclsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAclsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The frontend ID this ACL is attached to. ACLs with a matching frontend ID are listed.
        /// &gt; **Important:** LB frontend IDs are zoned, which means they are of the form `{zone}/{id}`, e.g. `fr-par-1/11111111-1111-1111-1111-111111111111`
        /// </summary>
        [Input("frontendId", required: true)]
        public string FrontendId { get; set; } = null!;

        /// <summary>
        /// The ACL name to filter for. ACLs with a matching name are listed.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// `zone`) The zone in which the ACLs exist.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetAclsArgs()
        {
        }
        public static new GetAclsArgs Empty => new GetAclsArgs();
    }

    public sealed class GetAclsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The frontend ID this ACL is attached to. ACLs with a matching frontend ID are listed.
        /// &gt; **Important:** LB frontend IDs are zoned, which means they are of the form `{zone}/{id}`, e.g. `fr-par-1/11111111-1111-1111-1111-111111111111`
        /// </summary>
        [Input("frontendId", required: true)]
        public Input<string> FrontendId { get; set; } = null!;

        /// <summary>
        /// The ACL name to filter for. ACLs with a matching name are listed.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `zone`) The zone in which the ACLs exist.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetAclsInvokeArgs()
        {
        }
        public static new GetAclsInvokeArgs Empty => new GetAclsInvokeArgs();
    }


    [OutputType]
    public sealed class GetAclsResult
    {
        /// <summary>
        /// List of retrieved ACLs
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAclsAclResult> Acls;
        public readonly string FrontendId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        public readonly string OrganizationId;
        public readonly string ProjectId;
        public readonly string Zone;

        [OutputConstructor]
        private GetAclsResult(
            ImmutableArray<Outputs.GetAclsAclResult> acls,

            string frontendId,

            string id,

            string? name,

            string organizationId,

            string projectId,

            string zone)
        {
            Acls = acls;
            FrontendId = frontendId;
            Id = id;
            Name = name;
            OrganizationId = organizationId;
            ProjectId = projectId;
            Zone = zone;
        }
    }
}
