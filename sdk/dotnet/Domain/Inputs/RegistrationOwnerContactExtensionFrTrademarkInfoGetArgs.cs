// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Domain.Inputs
{

    public sealed class RegistrationOwnerContactExtensionFrTrademarkInfoGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Trademark information from INPI (French extension).
        /// </summary>
        [Input("trademarkInpi")]
        public Input<string>? TrademarkInpi { get; set; }

        public RegistrationOwnerContactExtensionFrTrademarkInfoGetArgs()
        {
        }
        public static new RegistrationOwnerContactExtensionFrTrademarkInfoGetArgs Empty => new RegistrationOwnerContactExtensionFrTrademarkInfoGetArgs();
    }
}
