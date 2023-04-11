// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Scaleway.Inputs
{

    public sealed class MnqCredentialNatsCredentialsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("content")]
        private Input<string>? _content;

        /// <summary>
        /// Raw content of the NATS credentials file.
        /// </summary>
        public Input<string>? Content
        {
            get => _content;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _content = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public MnqCredentialNatsCredentialsGetArgs()
        {
        }
        public static new MnqCredentialNatsCredentialsGetArgs Empty => new MnqCredentialNatsCredentialsGetArgs();
    }
}