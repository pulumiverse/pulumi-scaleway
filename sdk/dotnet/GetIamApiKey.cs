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
    public static class GetIamApiKey
    {
        /// <summary>
        /// Gets information about an existing IAM API key. For more information, refer to the [IAM API documentation](https://www.scaleway.com/en/developers/api/iam/#api-keys-3665ae).
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
        ///     var main = Scaleway.GetIamApiKey.Invoke(new()
        ///     {
        ///         AccessKey = "SCWABCDEFGHIJKLMNOPQ",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetIamApiKeyResult> InvokeAsync(GetIamApiKeyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIamApiKeyResult>("scaleway:index/getIamApiKey:getIamApiKey", args ?? new GetIamApiKeyArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about an existing IAM API key. For more information, refer to the [IAM API documentation](https://www.scaleway.com/en/developers/api/iam/#api-keys-3665ae).
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
        ///     var main = Scaleway.GetIamApiKey.Invoke(new()
        ///     {
        ///         AccessKey = "SCWABCDEFGHIJKLMNOPQ",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetIamApiKeyResult> Invoke(GetIamApiKeyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIamApiKeyResult>("scaleway:index/getIamApiKey:getIamApiKey", args ?? new GetIamApiKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIamApiKeyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The access key of the IAM API key which is also the ID of the API key.
        /// </summary>
        [Input("accessKey", required: true)]
        public string AccessKey { get; set; } = null!;

        public GetIamApiKeyArgs()
        {
        }
        public static new GetIamApiKeyArgs Empty => new GetIamApiKeyArgs();
    }

    public sealed class GetIamApiKeyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The access key of the IAM API key which is also the ID of the API key.
        /// </summary>
        [Input("accessKey", required: true)]
        public Input<string> AccessKey { get; set; } = null!;

        public GetIamApiKeyInvokeArgs()
        {
        }
        public static new GetIamApiKeyInvokeArgs Empty => new GetIamApiKeyInvokeArgs();
    }


    [OutputType]
    public sealed class GetIamApiKeyResult
    {
        public readonly string AccessKey;
        public readonly string ApplicationId;
        public readonly string CreatedAt;
        public readonly string CreationIp;
        public readonly string DefaultProjectId;
        public readonly string Description;
        public readonly bool Editable;
        public readonly string ExpiresAt;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string UpdatedAt;
        public readonly string UserId;

        [OutputConstructor]
        private GetIamApiKeyResult(
            string accessKey,

            string applicationId,

            string createdAt,

            string creationIp,

            string defaultProjectId,

            string description,

            bool editable,

            string expiresAt,

            string id,

            string updatedAt,

            string userId)
        {
            AccessKey = accessKey;
            ApplicationId = applicationId;
            CreatedAt = createdAt;
            CreationIp = creationIp;
            DefaultProjectId = defaultProjectId;
            Description = description;
            Editable = editable;
            ExpiresAt = expiresAt;
            Id = id;
            UpdatedAt = updatedAt;
            UserId = userId;
        }
    }
}