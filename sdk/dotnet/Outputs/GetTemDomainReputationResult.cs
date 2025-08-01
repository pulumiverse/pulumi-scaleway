// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Outputs
{

    [OutputType]
    public sealed class GetTemDomainReputationResult
    {
        /// <summary>
        /// The previously-calculated domain's reputation score
        /// </summary>
        public readonly int PreviousScore;
        /// <summary>
        /// Time and date the previous reputation score was calculated
        /// </summary>
        public readonly string PreviousScoredAt;
        /// <summary>
        /// A range from 0 to 100 that determines your domain's reputation score
        /// </summary>
        public readonly int Score;
        /// <summary>
        /// Time and date the score was calculated
        /// </summary>
        public readonly string ScoredAt;
        /// <summary>
        /// Status of the domain's reputation
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetTemDomainReputationResult(
            int previousScore,

            string previousScoredAt,

            int score,

            string scoredAt,

            string status)
        {
            PreviousScore = previousScore;
            PreviousScoredAt = previousScoredAt;
            Score = score;
            ScoredAt = scoredAt;
            Status = status;
        }
    }
}
