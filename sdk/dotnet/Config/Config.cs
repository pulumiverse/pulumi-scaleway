// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumiverse.Scaleway
{
    public static class Config
    {
        [global::System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("scaleway");

        private static readonly __Value<string?> _accessKey = new __Value<string?>(() => __config.Get("accessKey") ?? Utilities.GetEnv("SCW_ACCESS_KEY"));
        /// <summary>
        /// The Scaleway access key.
        /// </summary>
        public static string? AccessKey
        {
            get => _accessKey.Get();
            set => _accessKey.Set(value);
        }

        private static readonly __Value<string?> _apiUrl = new __Value<string?>(() => __config.Get("apiUrl"));
        /// <summary>
        /// The Scaleway API URL to use.
        /// </summary>
        public static string? ApiUrl
        {
            get => _apiUrl.Get();
            set => _apiUrl.Set(value);
        }

        private static readonly __Value<string?> _organizationId = new __Value<string?>(() => __config.Get("organizationId") ?? Utilities.GetEnv("SCW_ORGANIZATION_ID"));
        /// <summary>
        /// The Scaleway organization ID.
        /// </summary>
        public static string? OrganizationId
        {
            get => _organizationId.Get();
            set => _organizationId.Set(value);
        }

        private static readonly __Value<string?> _profile = new __Value<string?>(() => __config.Get("profile"));
        /// <summary>
        /// The Scaleway profile to use.
        /// </summary>
        public static string? Profile
        {
            get => _profile.Get();
            set => _profile.Set(value);
        }

        private static readonly __Value<string?> _projectId = new __Value<string?>(() => __config.Get("projectId") ?? Utilities.GetEnv("SCW_DEFAULT_PROJECT_ID"));
        /// <summary>
        /// The Scaleway project ID.
        /// </summary>
        public static string? ProjectId
        {
            get => _projectId.Get();
            set => _projectId.Set(value);
        }

        private static readonly __Value<string?> _region = new __Value<string?>(() => __config.Get("region") ?? Utilities.GetEnv("SCW_DEFAULT_REGION"));
        /// <summary>
        /// The region you want to attach the resource to
        /// </summary>
        public static string? Region
        {
            get => _region.Get();
            set => _region.Set(value);
        }

        private static readonly __Value<string?> _secretKey = new __Value<string?>(() => __config.Get("secretKey") ?? Utilities.GetEnv("SCW_SECRET_KEY"));
        /// <summary>
        /// The Scaleway secret Key.
        /// </summary>
        public static string? SecretKey
        {
            get => _secretKey.Get();
            set => _secretKey.Set(value);
        }

        private static readonly __Value<string?> _zone = new __Value<string?>(() => __config.Get("zone") ?? Utilities.GetEnv("SCW_DEFAULT_ZONE"));
        /// <summary>
        /// The zone you want to attach the resource to
        /// </summary>
        public static string? Zone
        {
            get => _zone.Get();
            set => _zone.Set(value);
        }

    }
}
