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
    /// <summary>
    /// Creates and manages Scaleway Transactional Email Domains.
    /// For more information refer to the [API documentation](https://www.scaleway.com/en/developers/api/transactional-email).
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
    ///     var main = new Scaleway.Tem.Domain("main", new()
    ///     {
    ///         AcceptTos = true,
    ///         Name = "example.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Add the required records to your DNS zone
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var domainName = config.Require("domainName");
    ///     var main = new Scaleway.Tem.Domain("main", new()
    ///     {
    ///         Name = domainName,
    ///         AcceptTos = true,
    ///     });
    /// 
    ///     var spf = new Scaleway.Domain.Record("spf", new()
    ///     {
    ///         DnsZone = domainName,
    ///         Type = "TXT",
    ///         Data = main.SpfConfig.Apply(spfConfig =&gt; $"v=spf1 {spfConfig} -all"),
    ///     });
    /// 
    ///     var dkim = new Scaleway.Domain.Record("dkim", new()
    ///     {
    ///         DnsZone = domainName,
    ///         Name = main.ProjectId.Apply(projectId =&gt; $"{projectId}._domainkey"),
    ///         Type = "TXT",
    ///         Data = main.DkimConfig,
    ///     });
    /// 
    ///     var mx = new Scaleway.Domain.Record("mx", new()
    ///     {
    ///         DnsZone = domainName,
    ///         Type = "MX",
    ///         Data = ".",
    ///     });
    /// 
    ///     var dmarc = new Scaleway.Domain.Record("dmarc", new()
    ///     {
    ///         DnsZone = domainName,
    ///         Name = main.DmarcName,
    ///         Type = "TXT",
    ///         Data = main.DmarcConfig,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Automatically Configure DNS Settings for Your Domain
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var domainName = config.Require("domainName");
    ///     var main = new Scaleway.Tem.Domain("main", new()
    ///     {
    ///         Name = domainName,
    ///         AcceptTos = true,
    ///         Autoconfig = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Domains can be imported using the `{region}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/temDomain:TemDomain main fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [Obsolete(@"scaleway.index/temdomain.TemDomain has been deprecated in favor of scaleway.tem/domain.Domain")]
    [ScalewayResourceType("scaleway:index/temDomain:TemDomain")]
    public partial class TemDomain : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Acceptation of the [Term of Service](https://tem.s3.fr-par.scw.cloud/antispam_policy.pdf).
        /// &gt; **Important:** This attribute must be set to `true`.
        /// </summary>
        [Output("acceptTos")]
        public Output<bool> AcceptTos { get; private set; } = null!;

        /// <summary>
        /// Automatically configures DNS settings for the domain, simplifying the setup process by applying predefined configurations.
        /// </summary>
        [Output("autoconfig")]
        public Output<bool?> Autoconfig { get; private set; } = null!;

        /// <summary>
        /// The date and time of the Transaction Email Domain's creation (RFC 3339 format).
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The DKIM public key, as should be recorded in the DNS zone.
        /// </summary>
        [Output("dkimConfig")]
        public Output<string> DkimConfig { get; private set; } = null!;

        /// <summary>
        /// DMARC record for the domain, as should be recorded in the DNS zone.
        /// </summary>
        [Output("dmarcConfig")]
        public Output<string> DmarcConfig { get; private set; } = null!;

        /// <summary>
        /// DMARC name for the domain, as should be recorded in the DNS zone.
        /// </summary>
        [Output("dmarcName")]
        public Output<string> DmarcName { get; private set; } = null!;

        /// <summary>
        /// (Deprecated) The error message if the last check failed.
        /// </summary>
        [Output("lastError")]
        public Output<string> LastError { get; private set; } = null!;

        /// <summary>
        /// The date and time the domain was last found to be valid (RFC 3339 format).
        /// </summary>
        [Output("lastValidAt")]
        public Output<string> LastValidAt { get; private set; } = null!;

        /// <summary>
        /// The Scaleway's blackhole MX server to use if you do not have one.
        /// </summary>
        [Output("mxBlackhole")]
        public Output<string> MxBlackhole { get; private set; } = null!;

        /// <summary>
        /// The domain name, must not be used in another Transactional Email Domain.
        /// &gt; **Important:** Updates to `name` will recreate the domain.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The date and time of the next scheduled check (RFC 3339 format).
        /// </summary>
        [Output("nextCheckAt")]
        public Output<string> NextCheckAt { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the domain is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// `region`). The region in which the domain should be created.
        /// &gt; **Important:** Currently, only fr-par is supported. Specifying any other region will cause an error.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The domain's reputation.
        /// </summary>
        [Output("reputations")]
        public Output<ImmutableArray<Outputs.TemDomainReputation>> Reputations { get; private set; } = null!;

        /// <summary>
        /// The date and time of the revocation of the domain (RFC 3339 format).
        /// </summary>
        [Output("revokedAt")]
        public Output<string> RevokedAt { get; private set; } = null!;

        /// <summary>
        /// The SMTP host to use to send emails.
        /// </summary>
        [Output("smtpHost")]
        public Output<string> SmtpHost { get; private set; } = null!;

        /// <summary>
        /// The SMTP port to use to send emails over TLS.
        /// </summary>
        [Output("smtpPort")]
        public Output<int> SmtpPort { get; private set; } = null!;

        /// <summary>
        /// The SMTP port to use to send emails over TLS.
        /// </summary>
        [Output("smtpPortAlternative")]
        public Output<int> SmtpPortAlternative { get; private set; } = null!;

        /// <summary>
        /// The SMTP port to use to send emails.
        /// </summary>
        [Output("smtpPortUnsecure")]
        public Output<int> SmtpPortUnsecure { get; private set; } = null!;

        /// <summary>
        /// SMTPS auth user refers to the identifier for a user authorized to send emails via SMTPS, ensuring secure email transmission.
        /// </summary>
        [Output("smtpsAuthUser")]
        public Output<string> SmtpsAuthUser { get; private set; } = null!;

        /// <summary>
        /// The SMTPS port to use to send emails over TLS Wrapper.
        /// </summary>
        [Output("smtpsPort")]
        public Output<int> SmtpsPort { get; private set; } = null!;

        /// <summary>
        /// The SMTPS port to use to send emails over TLS Wrapper.
        /// </summary>
        [Output("smtpsPortAlternative")]
        public Output<int> SmtpsPortAlternative { get; private set; } = null!;

        /// <summary>
        /// The snippet of the SPF record that should be registered in the DNS zone.
        /// </summary>
        [Output("spfConfig")]
        public Output<string> SpfConfig { get; private set; } = null!;

        /// <summary>
        /// The status of the domain's reputation.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a TemDomain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TemDomain(string name, TemDomainArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/temDomain:TemDomain", name, args ?? new TemDomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TemDomain(string name, Input<string> id, TemDomainState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/temDomain:TemDomain", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TemDomain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TemDomain Get(string name, Input<string> id, TemDomainState? state = null, CustomResourceOptions? options = null)
        {
            return new TemDomain(name, id, state, options);
        }
    }

    public sealed class TemDomainArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Acceptation of the [Term of Service](https://tem.s3.fr-par.scw.cloud/antispam_policy.pdf).
        /// &gt; **Important:** This attribute must be set to `true`.
        /// </summary>
        [Input("acceptTos", required: true)]
        public Input<bool> AcceptTos { get; set; } = null!;

        /// <summary>
        /// Automatically configures DNS settings for the domain, simplifying the setup process by applying predefined configurations.
        /// </summary>
        [Input("autoconfig")]
        public Input<bool>? Autoconfig { get; set; }

        /// <summary>
        /// The domain name, must not be used in another Transactional Email Domain.
        /// &gt; **Important:** Updates to `name` will recreate the domain.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the domain is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region in which the domain should be created.
        /// &gt; **Important:** Currently, only fr-par is supported. Specifying any other region will cause an error.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public TemDomainArgs()
        {
        }
        public static new TemDomainArgs Empty => new TemDomainArgs();
    }

    public sealed class TemDomainState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Acceptation of the [Term of Service](https://tem.s3.fr-par.scw.cloud/antispam_policy.pdf).
        /// &gt; **Important:** This attribute must be set to `true`.
        /// </summary>
        [Input("acceptTos")]
        public Input<bool>? AcceptTos { get; set; }

        /// <summary>
        /// Automatically configures DNS settings for the domain, simplifying the setup process by applying predefined configurations.
        /// </summary>
        [Input("autoconfig")]
        public Input<bool>? Autoconfig { get; set; }

        /// <summary>
        /// The date and time of the Transaction Email Domain's creation (RFC 3339 format).
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The DKIM public key, as should be recorded in the DNS zone.
        /// </summary>
        [Input("dkimConfig")]
        public Input<string>? DkimConfig { get; set; }

        /// <summary>
        /// DMARC record for the domain, as should be recorded in the DNS zone.
        /// </summary>
        [Input("dmarcConfig")]
        public Input<string>? DmarcConfig { get; set; }

        /// <summary>
        /// DMARC name for the domain, as should be recorded in the DNS zone.
        /// </summary>
        [Input("dmarcName")]
        public Input<string>? DmarcName { get; set; }

        /// <summary>
        /// (Deprecated) The error message if the last check failed.
        /// </summary>
        [Input("lastError")]
        public Input<string>? LastError { get; set; }

        /// <summary>
        /// The date and time the domain was last found to be valid (RFC 3339 format).
        /// </summary>
        [Input("lastValidAt")]
        public Input<string>? LastValidAt { get; set; }

        /// <summary>
        /// The Scaleway's blackhole MX server to use if you do not have one.
        /// </summary>
        [Input("mxBlackhole")]
        public Input<string>? MxBlackhole { get; set; }

        /// <summary>
        /// The domain name, must not be used in another Transactional Email Domain.
        /// &gt; **Important:** Updates to `name` will recreate the domain.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The date and time of the next scheduled check (RFC 3339 format).
        /// </summary>
        [Input("nextCheckAt")]
        public Input<string>? NextCheckAt { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the domain is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region in which the domain should be created.
        /// &gt; **Important:** Currently, only fr-par is supported. Specifying any other region will cause an error.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("reputations")]
        private InputList<Inputs.TemDomainReputationGetArgs>? _reputations;

        /// <summary>
        /// The domain's reputation.
        /// </summary>
        public InputList<Inputs.TemDomainReputationGetArgs> Reputations
        {
            get => _reputations ?? (_reputations = new InputList<Inputs.TemDomainReputationGetArgs>());
            set => _reputations = value;
        }

        /// <summary>
        /// The date and time of the revocation of the domain (RFC 3339 format).
        /// </summary>
        [Input("revokedAt")]
        public Input<string>? RevokedAt { get; set; }

        /// <summary>
        /// The SMTP host to use to send emails.
        /// </summary>
        [Input("smtpHost")]
        public Input<string>? SmtpHost { get; set; }

        /// <summary>
        /// The SMTP port to use to send emails over TLS.
        /// </summary>
        [Input("smtpPort")]
        public Input<int>? SmtpPort { get; set; }

        /// <summary>
        /// The SMTP port to use to send emails over TLS.
        /// </summary>
        [Input("smtpPortAlternative")]
        public Input<int>? SmtpPortAlternative { get; set; }

        /// <summary>
        /// The SMTP port to use to send emails.
        /// </summary>
        [Input("smtpPortUnsecure")]
        public Input<int>? SmtpPortUnsecure { get; set; }

        /// <summary>
        /// SMTPS auth user refers to the identifier for a user authorized to send emails via SMTPS, ensuring secure email transmission.
        /// </summary>
        [Input("smtpsAuthUser")]
        public Input<string>? SmtpsAuthUser { get; set; }

        /// <summary>
        /// The SMTPS port to use to send emails over TLS Wrapper.
        /// </summary>
        [Input("smtpsPort")]
        public Input<int>? SmtpsPort { get; set; }

        /// <summary>
        /// The SMTPS port to use to send emails over TLS Wrapper.
        /// </summary>
        [Input("smtpsPortAlternative")]
        public Input<int>? SmtpsPortAlternative { get; set; }

        /// <summary>
        /// The snippet of the SPF record that should be registered in the DNS zone.
        /// </summary>
        [Input("spfConfig")]
        public Input<string>? SpfConfig { get; set; }

        /// <summary>
        /// The status of the domain's reputation.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public TemDomainState()
        {
        }
        public static new TemDomainState Empty => new TemDomainState();
    }
}
