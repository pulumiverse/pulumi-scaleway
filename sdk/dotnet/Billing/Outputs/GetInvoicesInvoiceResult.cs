// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Billing.Outputs
{

    [OutputType]
    public sealed class GetInvoicesInvoiceResult
    {
        /// <summary>
        /// The billing period of the invoice in the YYYY-MM format.
        /// </summary>
        public readonly string BillingPeriod;
        /// <summary>
        /// The payment time limit, set according to the Organization's payment conditions (RFC 3339 format).
        /// </summary>
        public readonly string DueDate;
        /// <summary>
        /// The associated invoice ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Invoices with the given type are listed. Valid values are `periodic` and `purchase`.
        /// </summary>
        public readonly string InvoiceType;
        /// <summary>
        /// The date when the invoice was sent to the customer (RFC 3339 format).
        /// </summary>
        public readonly string IssuedDate;
        /// <summary>
        /// The invoice number.
        /// </summary>
        public readonly int Number;
        /// <summary>
        /// The organization name.
        /// </summary>
        public readonly string OrganizationName;
        /// <summary>
        /// The name of the seller (Scaleway).
        /// </summary>
        public readonly string SellerName;
        /// <summary>
        /// The start date of the billing period (RFC 3339 format).
        /// </summary>
        public readonly string StartDate;
        /// <summary>
        /// The state of the invoice.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The end date of the billing period (RFC 3339 format).
        /// </summary>
        public readonly string StopDate;
        /// <summary>
        /// The total discount amount of the invoice.
        /// </summary>
        public readonly string TotalDiscount;
        /// <summary>
        /// The total tax amount of the invoice.
        /// </summary>
        public readonly string TotalTax;
        /// <summary>
        /// The total amount, taxed.
        /// </summary>
        public readonly string TotalTaxed;
        /// <summary>
        /// The total amount of the invoice before applying the discount.
        /// </summary>
        public readonly string TotalUndiscount;
        /// <summary>
        /// The total amount, untaxed.
        /// </summary>
        public readonly string TotalUntaxed;

        [OutputConstructor]
        private GetInvoicesInvoiceResult(
            string billingPeriod,

            string dueDate,

            string id,

            string invoiceType,

            string issuedDate,

            int number,

            string organizationName,

            string sellerName,

            string startDate,

            string state,

            string stopDate,

            string totalDiscount,

            string totalTax,

            string totalTaxed,

            string totalUndiscount,

            string totalUntaxed)
        {
            BillingPeriod = billingPeriod;
            DueDate = dueDate;
            Id = id;
            InvoiceType = invoiceType;
            IssuedDate = issuedDate;
            Number = number;
            OrganizationName = organizationName;
            SellerName = sellerName;
            StartDate = startDate;
            State = state;
            StopDate = stopDate;
            TotalDiscount = totalDiscount;
            TotalTax = totalTax;
            TotalTaxed = totalTaxed;
            TotalUndiscount = totalUndiscount;
            TotalUntaxed = totalUntaxed;
        }
    }
}
