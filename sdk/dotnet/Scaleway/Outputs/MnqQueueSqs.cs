// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Scaleway.Outputs
{

    [OutputType]
    public sealed class MnqQueueSqs
    {
        public readonly string AccessKey;
        public readonly bool? ContentBasedDeduplication;
        public readonly string? Endpoint;
        public readonly bool? FifoQueue;
        public readonly int? ReceiveWaitTimeSeconds;
        public readonly string SecretKey;
        public readonly string? Url;
        public readonly int? VisibilityTimeoutSeconds;

        [OutputConstructor]
        private MnqQueueSqs(
            string accessKey,

            bool? contentBasedDeduplication,

            string? endpoint,

            bool? fifoQueue,

            int? receiveWaitTimeSeconds,

            string secretKey,

            string? url,

            int? visibilityTimeoutSeconds)
        {
            AccessKey = accessKey;
            ContentBasedDeduplication = contentBasedDeduplication;
            Endpoint = endpoint;
            FifoQueue = fifoQueue;
            ReceiveWaitTimeSeconds = receiveWaitTimeSeconds;
            SecretKey = secretKey;
            Url = url;
            VisibilityTimeoutSeconds = visibilityTimeoutSeconds;
        }
    }
}