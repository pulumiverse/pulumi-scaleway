// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { DefinitionArgs, DefinitionState } from "./definition";
export type Definition = import("./definition").Definition;
export const Definition: typeof import("./definition").Definition = null as any;
utilities.lazyLoad(exports, ["Definition"], () => require("./definition"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "scaleway:job/definition:Definition":
                return new Definition(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("scaleway", "job/definition", _module)
