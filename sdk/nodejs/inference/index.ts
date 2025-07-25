// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { DeploymentArgs, DeploymentState } from "./deployment";
export type Deployment = import("./deployment").Deployment;
export const Deployment: typeof import("./deployment").Deployment = null as any;
utilities.lazyLoad(exports, ["Deployment"], () => require("./deployment"));

export { GetModelArgs, GetModelResult, GetModelOutputArgs } from "./getModel";
export const getModel: typeof import("./getModel").getModel = null as any;
export const getModelOutput: typeof import("./getModel").getModelOutput = null as any;
utilities.lazyLoad(exports, ["getModel","getModelOutput"], () => require("./getModel"));

export { ModelArgs, ModelState } from "./model";
export type Model = import("./model").Model;
export const Model: typeof import("./model").Model = null as any;
utilities.lazyLoad(exports, ["Model"], () => require("./model"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "scaleway:inference/deployment:Deployment":
                return new Deployment(name, <any>undefined, { urn })
            case "scaleway:inference/model:Model":
                return new Model(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("scaleway", "inference/deployment", _module)
pulumi.runtime.registerResourceModule("scaleway", "inference/model", _module)
