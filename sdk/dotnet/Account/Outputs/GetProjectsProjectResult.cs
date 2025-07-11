// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Account.Outputs
{

    [OutputType]
    public sealed class GetProjectsProjectResult
    {
        /// <summary>
        /// (Computed) The date and time when the project was created.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// (Computed) The description of the project.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// (Computed) The unique identifier of the project.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// (Computed) The name of the project.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The unique identifier of the Organization with which the Projects are associated.
        /// If no default `organization_id` is set, one must be set explicitly in this datasource
        /// </summary>
        public readonly string OrganizationId;
        /// <summary>
        /// (Computed) The date and time when the project was updated.
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetProjectsProjectResult(
            string createdAt,

            string description,

            string id,

            string name,

            string organizationId,

            string updatedAt)
        {
            CreatedAt = createdAt;
            Description = description;
            Id = id;
            Name = name;
            OrganizationId = organizationId;
            UpdatedAt = updatedAt;
        }
    }
}
