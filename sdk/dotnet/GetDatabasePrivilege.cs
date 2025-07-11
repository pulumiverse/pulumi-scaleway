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
    [Obsolete(@"scaleway.index/getdatabaseprivilege.getDatabasePrivilege has been deprecated in favor of scaleway.databases/getprivilege.getPrivilege")]
    public static class GetDatabasePrivilege
    {
        /// <summary>
        /// Gets information about the privileges in a database.
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
        ///     // Get the database privilege for the user "my-user" on the database "my-database" hosted on instance id 11111111-1111-1111-1111-111111111111 and on the default region. e.g: fr-par
        ///     var main = Scaleway.Databases.GetPrivilege.Invoke(new()
        ///     {
        ///         InstanceId = "11111111-1111-111111111111",
        ///         UserName = "my-user",
        ///         DatabaseName = "my-database",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetDatabasePrivilegeResult> InvokeAsync(GetDatabasePrivilegeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatabasePrivilegeResult>("scaleway:index/getDatabasePrivilege:getDatabasePrivilege", args ?? new GetDatabasePrivilegeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about the privileges in a database.
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
        ///     // Get the database privilege for the user "my-user" on the database "my-database" hosted on instance id 11111111-1111-1111-1111-111111111111 and on the default region. e.g: fr-par
        ///     var main = Scaleway.Databases.GetPrivilege.Invoke(new()
        ///     {
        ///         InstanceId = "11111111-1111-111111111111",
        ///         UserName = "my-user",
        ///         DatabaseName = "my-database",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetDatabasePrivilegeResult> Invoke(GetDatabasePrivilegeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabasePrivilegeResult>("scaleway:index/getDatabasePrivilege:getDatabasePrivilege", args ?? new GetDatabasePrivilegeInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about the privileges in a database.
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
        ///     // Get the database privilege for the user "my-user" on the database "my-database" hosted on instance id 11111111-1111-1111-1111-111111111111 and on the default region. e.g: fr-par
        ///     var main = Scaleway.Databases.GetPrivilege.Invoke(new()
        ///     {
        ///         InstanceId = "11111111-1111-111111111111",
        ///         UserName = "my-user",
        ///         DatabaseName = "my-database",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetDatabasePrivilegeResult> Invoke(GetDatabasePrivilegeInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabasePrivilegeResult>("scaleway:index/getDatabasePrivilege:getDatabasePrivilege", args ?? new GetDatabasePrivilegeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabasePrivilegeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The database name.
        /// </summary>
        [Input("databaseName", required: true)]
        public string DatabaseName { get; set; } = null!;

        /// <summary>
        /// The Database Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// `region`) The region in which the resource exists.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// The user name.
        /// </summary>
        [Input("userName", required: true)]
        public string UserName { get; set; } = null!;

        public GetDatabasePrivilegeArgs()
        {
        }
        public static new GetDatabasePrivilegeArgs Empty => new GetDatabasePrivilegeArgs();
    }

    public sealed class GetDatabasePrivilegeInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The database name.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// The Database Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// `region`) The region in which the resource exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The user name.
        /// </summary>
        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public GetDatabasePrivilegeInvokeArgs()
        {
        }
        public static new GetDatabasePrivilegeInvokeArgs Empty => new GetDatabasePrivilegeInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatabasePrivilegeResult
    {
        public readonly string DatabaseName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        /// <summary>
        /// The permission for this user on the database. Possible values are `readonly`, `readwrite`, `all`
        /// , `custom` and `none`.
        /// </summary>
        public readonly string Permission;
        public readonly string? Region;
        public readonly string UserName;

        [OutputConstructor]
        private GetDatabasePrivilegeResult(
            string databaseName,

            string id,

            string instanceId,

            string permission,

            string? region,

            string userName)
        {
            DatabaseName = databaseName;
            Id = id;
            InstanceId = instanceId;
            Permission = permission;
            Region = region;
            UserName = userName;
        }
    }
}
