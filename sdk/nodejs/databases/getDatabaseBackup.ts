// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Gets information about an RDB backup.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const findByName = scaleway.databases.getDatabaseBackup({
 *     name: "mybackup",
 * });
 * const findByNameAndInstance = scaleway.databases.getDatabaseBackup({
 *     name: "mybackup",
 *     instanceId: "11111111-1111-1111-1111-111111111111",
 * });
 * const findById = scaleway.databases.getDatabaseBackup({
 *     backupId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getDatabaseBackup(args?: GetDatabaseBackupArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabaseBackupResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:databases/getDatabaseBackup:getDatabaseBackup", {
        "backupId": args.backupId,
        "instanceId": args.instanceId,
        "name": args.name,
        "projectId": args.projectId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatabaseBackup.
 */
export interface GetDatabaseBackupArgs {
    /**
     * The backup ID.
     */
    backupId?: string;
    /**
     * The Database Instance ID.
     */
    instanceId?: string;
    /**
     * The name of the RDB instance.
     *
     * > **Note** You must specify at least one: `name` and/or `backupId`.
     */
    name?: string;
    /**
     * The ID of the project the Database Backup is associated with.
     */
    projectId?: string;
    /**
     * `region`) The region in which the Database Backup is associated with.
     */
    region?: string;
}

/**
 * A collection of values returned by getDatabaseBackup.
 */
export interface GetDatabaseBackupResult {
    readonly backupId?: string;
    readonly createdAt: string;
    readonly databaseName: string;
    readonly expiresAt: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId?: string;
    readonly instanceName: string;
    readonly name?: string;
    readonly projectId?: string;
    readonly region?: string;
    readonly size: number;
    readonly updatedAt: string;
}
/**
 * Gets information about an RDB backup.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const findByName = scaleway.databases.getDatabaseBackup({
 *     name: "mybackup",
 * });
 * const findByNameAndInstance = scaleway.databases.getDatabaseBackup({
 *     name: "mybackup",
 *     instanceId: "11111111-1111-1111-1111-111111111111",
 * });
 * const findById = scaleway.databases.getDatabaseBackup({
 *     backupId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getDatabaseBackupOutput(args?: GetDatabaseBackupOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDatabaseBackupResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:databases/getDatabaseBackup:getDatabaseBackup", {
        "backupId": args.backupId,
        "instanceId": args.instanceId,
        "name": args.name,
        "projectId": args.projectId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatabaseBackup.
 */
export interface GetDatabaseBackupOutputArgs {
    /**
     * The backup ID.
     */
    backupId?: pulumi.Input<string>;
    /**
     * The Database Instance ID.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The name of the RDB instance.
     *
     * > **Note** You must specify at least one: `name` and/or `backupId`.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project the Database Backup is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region in which the Database Backup is associated with.
     */
    region?: pulumi.Input<string>;
}
