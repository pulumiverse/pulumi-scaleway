// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The `scaleway.domain.Record` data source is used to get information about an existing domain record.
 *
 * Refer to the Domains and DNS [product documentation](https://www.scaleway.com/en/docs/network/domains-and-dns/) and [API documentation](https://www.scaleway.com/en/developers/api/domains-and-dns/) for more information.
 *
 * ## Query domain records
 *
 * The following commands allow you to:
 *
 * - query a domain record specified by the DNS zone (`domain.tld`), the record name (`www`), the record type (`A`), and the record content (`1.2.3.4`).
 * - query a domain record specified by the DNS zone (`domain.tld`) and the unique record ID (`11111111-1111-1111-1111-111111111111`).
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * // Query record by DNS zone, record name, type and content
 * const byContent = scaleway.domain.getRecord({
 *     dnsZone: "domain.tld",
 *     name: "www",
 *     type: "A",
 *     data: "1.2.3.4",
 * });
 * // Query record by DNS zone and record ID
 * const byId = scaleway.domain.getRecord({
 *     dnsZone: "domain.tld",
 *     recordId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getRecord(args?: GetRecordArgs, opts?: pulumi.InvokeOptions): Promise<GetRecordResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:domain/getRecord:getRecord", {
        "data": args.data,
        "dnsZone": args.dnsZone,
        "name": args.name,
        "projectId": args.projectId,
        "recordId": args.recordId,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getRecord.
 */
export interface GetRecordArgs {
    /**
     * The content of the record (e.g., an IPv4 address for an `A` record or a string for a `TXT` record). Cannot be used with `recordId`.
     */
    data?: string;
    /**
     * The DNS zone (domain) to which the record belongs. This is a required field in both examples above but is optional in the context of defining the data source.
     */
    dnsZone?: string;
    /**
     * The name of the record, which can be an empty string for a root record. Cannot be used with `recordId`.
     */
    name?: string;
    /**
     * ). The ID of the Project associated with the domain.
     */
    projectId?: string;
    /**
     * The unique identifier of the record. Cannot be used with `name`, `type`, and `data`.
     */
    recordId?: string;
    /**
     * The type of the record (`A`, `AAAA`, `MX`, `CNAME`, etc.). Cannot be used with `recordId`.
     */
    type?: string;
}

/**
 * A collection of values returned by getRecord.
 */
export interface GetRecordResult {
    readonly data?: string;
    readonly dnsZone?: string;
    readonly fqdn: string;
    /**
     * Information about dynamic records based on user geolocation. Find out more about dynamic records.
     */
    readonly geoIps: outputs.domain.GetRecordGeoIp[];
    /**
     * Information about dynamic records based on URL resolution. Find out more about dynamic records.
     */
    readonly httpServices: outputs.domain.GetRecordHttpService[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name?: string;
    /**
     * The priority of the record, mainly used with `MX` records.
     */
    readonly priority: number;
    readonly projectId?: string;
    readonly recordId?: string;
    readonly rootZone: boolean;
    /**
     * The Time To Live (TTL) of the record in seconds.
     */
    readonly ttl: number;
    readonly type?: string;
    /**
     * Information about dynamic records based on the client’s (resolver) subnet. Find out more about dynamic records.
     */
    readonly views: outputs.domain.GetRecordView[];
    /**
     * Information about dynamic records based on IP weights. Find out more about dynamic records.
     */
    readonly weighteds: outputs.domain.GetRecordWeighted[];
}
/**
 * The `scaleway.domain.Record` data source is used to get information about an existing domain record.
 *
 * Refer to the Domains and DNS [product documentation](https://www.scaleway.com/en/docs/network/domains-and-dns/) and [API documentation](https://www.scaleway.com/en/developers/api/domains-and-dns/) for more information.
 *
 * ## Query domain records
 *
 * The following commands allow you to:
 *
 * - query a domain record specified by the DNS zone (`domain.tld`), the record name (`www`), the record type (`A`), and the record content (`1.2.3.4`).
 * - query a domain record specified by the DNS zone (`domain.tld`) and the unique record ID (`11111111-1111-1111-1111-111111111111`).
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * // Query record by DNS zone, record name, type and content
 * const byContent = scaleway.domain.getRecord({
 *     dnsZone: "domain.tld",
 *     name: "www",
 *     type: "A",
 *     data: "1.2.3.4",
 * });
 * // Query record by DNS zone and record ID
 * const byId = scaleway.domain.getRecord({
 *     dnsZone: "domain.tld",
 *     recordId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getRecordOutput(args?: GetRecordOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetRecordResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:domain/getRecord:getRecord", {
        "data": args.data,
        "dnsZone": args.dnsZone,
        "name": args.name,
        "projectId": args.projectId,
        "recordId": args.recordId,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getRecord.
 */
export interface GetRecordOutputArgs {
    /**
     * The content of the record (e.g., an IPv4 address for an `A` record or a string for a `TXT` record). Cannot be used with `recordId`.
     */
    data?: pulumi.Input<string>;
    /**
     * The DNS zone (domain) to which the record belongs. This is a required field in both examples above but is optional in the context of defining the data source.
     */
    dnsZone?: pulumi.Input<string>;
    /**
     * The name of the record, which can be an empty string for a root record. Cannot be used with `recordId`.
     */
    name?: pulumi.Input<string>;
    /**
     * ). The ID of the Project associated with the domain.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The unique identifier of the record. Cannot be used with `name`, `type`, and `data`.
     */
    recordId?: pulumi.Input<string>;
    /**
     * The type of the record (`A`, `AAAA`, `MX`, `CNAME`, etc.). Cannot be used with `recordId`.
     */
    type?: pulumi.Input<string>;
}
