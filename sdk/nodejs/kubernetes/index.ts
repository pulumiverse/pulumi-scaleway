// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AclArgs, AclState } from "./acl";
export type Acl = import("./acl").Acl;
export const Acl: typeof import("./acl").Acl = null as any;
utilities.lazyLoad(exports, ["Acl"], () => require("./acl"));

export { ClusterArgs, ClusterState } from "./cluster";
export type Cluster = import("./cluster").Cluster;
export const Cluster: typeof import("./cluster").Cluster = null as any;
utilities.lazyLoad(exports, ["Cluster"], () => require("./cluster"));

export { GetClusterArgs, GetClusterResult, GetClusterOutputArgs } from "./getCluster";
export const getCluster: typeof import("./getCluster").getCluster = null as any;
export const getClusterOutput: typeof import("./getCluster").getClusterOutput = null as any;
utilities.lazyLoad(exports, ["getCluster","getClusterOutput"], () => require("./getCluster"));

export { GetPoolArgs, GetPoolResult, GetPoolOutputArgs } from "./getPool";
export const getPool: typeof import("./getPool").getPool = null as any;
export const getPoolOutput: typeof import("./getPool").getPoolOutput = null as any;
utilities.lazyLoad(exports, ["getPool","getPoolOutput"], () => require("./getPool"));

export { GetVersionArgs, GetVersionResult, GetVersionOutputArgs } from "./getVersion";
export const getVersion: typeof import("./getVersion").getVersion = null as any;
export const getVersionOutput: typeof import("./getVersion").getVersionOutput = null as any;
utilities.lazyLoad(exports, ["getVersion","getVersionOutput"], () => require("./getVersion"));

export { PoolArgs, PoolState } from "./pool";
export type Pool = import("./pool").Pool;
export const Pool: typeof import("./pool").Pool = null as any;
utilities.lazyLoad(exports, ["Pool"], () => require("./pool"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "scaleway:kubernetes/acl:Acl":
                return new Acl(name, <any>undefined, { urn })
            case "scaleway:kubernetes/cluster:Cluster":
                return new Cluster(name, <any>undefined, { urn })
            case "scaleway:kubernetes/pool:Pool":
                return new Pool(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("scaleway", "kubernetes/acl", _module)
pulumi.runtime.registerResourceModule("scaleway", "kubernetes/cluster", _module)
pulumi.runtime.registerResourceModule("scaleway", "kubernetes/pool", _module)
