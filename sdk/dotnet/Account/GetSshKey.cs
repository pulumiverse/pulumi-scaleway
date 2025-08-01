// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Account
{
    public static class GetSshKey
    {
        /// <summary>
        /// The `scaleway.account.SshKey` data source is used to retrieve information about a the SSH key of a Scaleway account.
        /// 
        /// Refer to the Organizations and Projects [documentation](https://www.scaleway.com/en/docs/organizations-and-projects/how-to/create-ssh-key/) and [API documentation](https://www.scaleway.com/en/developers/api/iam/#path-ssh-keys) for more information.
        /// </summary>
        public static Task<GetSshKeyResult> InvokeAsync(GetSshKeyArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSshKeyResult>("scaleway:account/getSshKey:getSshKey", args ?? new GetSshKeyArgs(), options.WithDefaults());

        /// <summary>
        /// The `scaleway.account.SshKey` data source is used to retrieve information about a the SSH key of a Scaleway account.
        /// 
        /// Refer to the Organizations and Projects [documentation](https://www.scaleway.com/en/docs/organizations-and-projects/how-to/create-ssh-key/) and [API documentation](https://www.scaleway.com/en/developers/api/iam/#path-ssh-keys) for more information.
        /// </summary>
        public static Output<GetSshKeyResult> Invoke(GetSshKeyInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSshKeyResult>("scaleway:account/getSshKey:getSshKey", args ?? new GetSshKeyInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The `scaleway.account.SshKey` data source is used to retrieve information about a the SSH key of a Scaleway account.
        /// 
        /// Refer to the Organizations and Projects [documentation](https://www.scaleway.com/en/docs/organizations-and-projects/how-to/create-ssh-key/) and [API documentation](https://www.scaleway.com/en/developers/api/iam/#path-ssh-keys) for more information.
        /// </summary>
        public static Output<GetSshKeyResult> Invoke(GetSshKeyInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSshKeyResult>("scaleway:account/getSshKey:getSshKey", args ?? new GetSshKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSshKeyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the SSH key.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// `project_id`) The unique identifier of the project with which the SSH key is associated.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// The unique identifier of the SSH key.
        /// 
        /// &gt; **Note** You must specify at least one: `name` and/or `ssh_key_id`.
        /// </summary>
        [Input("sshKeyId")]
        public string? SshKeyId { get; set; }

        public GetSshKeyArgs()
        {
        }
        public static new GetSshKeyArgs Empty => new GetSshKeyArgs();
    }

    public sealed class GetSshKeyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the SSH key.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `project_id`) The unique identifier of the project with which the SSH key is associated.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The unique identifier of the SSH key.
        /// 
        /// &gt; **Note** You must specify at least one: `name` and/or `ssh_key_id`.
        /// </summary>
        [Input("sshKeyId")]
        public Input<string>? SshKeyId { get; set; }

        public GetSshKeyInvokeArgs()
        {
        }
        public static new GetSshKeyInvokeArgs Empty => new GetSshKeyInvokeArgs();
    }


    [OutputType]
    public sealed class GetSshKeyResult
    {
        public readonly string CreatedAt;
        public readonly bool Disabled;
        public readonly string Fingerprint;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        /// <summary>
        /// The unique identifier of the Organization with which the SSH key is associated.
        /// </summary>
        public readonly string OrganizationId;
        public readonly string? ProjectId;
        /// <summary>
        /// The string of the SSH public key.
        /// </summary>
        public readonly string PublicKey;
        public readonly string? SshKeyId;
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetSshKeyResult(
            string createdAt,

            bool disabled,

            string fingerprint,

            string id,

            string? name,

            string organizationId,

            string? projectId,

            string publicKey,

            string? sshKeyId,

            string updatedAt)
        {
            CreatedAt = createdAt;
            Disabled = disabled;
            Fingerprint = fingerprint;
            Id = id;
            Name = name;
            OrganizationId = organizationId;
            ProjectId = projectId;
            PublicKey = publicKey;
            SshKeyId = sshKeyId;
            UpdatedAt = updatedAt;
        }
    }
}
