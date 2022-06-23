// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway
{
    public static class GetFunctionNamespace
    {
        public static Task<GetFunctionNamespaceResult> InvokeAsync(GetFunctionNamespaceArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFunctionNamespaceResult>("scaleway:index/getFunctionNamespace:getFunctionNamespace", args ?? new GetFunctionNamespaceArgs(), options.WithDefaults());

        public static Output<GetFunctionNamespaceResult> Invoke(GetFunctionNamespaceInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFunctionNamespaceResult>("scaleway:index/getFunctionNamespace:getFunctionNamespace", args ?? new GetFunctionNamespaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFunctionNamespaceArgs : Pulumi.InvokeArgs
    {
        [Input("name")]
        public string? Name { get; set; }

        [Input("namespaceId")]
        public string? NamespaceId { get; set; }

        [Input("region")]
        public string? Region { get; set; }

        public GetFunctionNamespaceArgs()
        {
        }
    }

    public sealed class GetFunctionNamespaceInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("namespaceId")]
        public Input<string>? NamespaceId { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetFunctionNamespaceInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFunctionNamespaceResult
    {
        public readonly string Description;
        public readonly ImmutableDictionary<string, string> EnvironmentVariables;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        public readonly string? NamespaceId;
        public readonly string OrganizationId;
        public readonly string ProjectId;
        public readonly string? Region;
        public readonly string RegistryEndpoint;
        public readonly string RegistryNamespaceId;

        [OutputConstructor]
        private GetFunctionNamespaceResult(
            string description,

            ImmutableDictionary<string, string> environmentVariables,

            string id,

            string? name,

            string? namespaceId,

            string organizationId,

            string projectId,

            string? region,

            string registryEndpoint,

            string registryNamespaceId)
        {
            Description = description;
            EnvironmentVariables = environmentVariables;
            Id = id;
            Name = name;
            NamespaceId = namespaceId;
            OrganizationId = organizationId;
            ProjectId = projectId;
            Region = region;
            RegistryEndpoint = registryEndpoint;
            RegistryNamespaceId = registryNamespaceId;
        }
    }
}