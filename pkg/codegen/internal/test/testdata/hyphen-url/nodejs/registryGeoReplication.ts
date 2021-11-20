// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

import * as pulumiAzureNative from "@pulumi/azure-native";

export class RegistryGeoReplication extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'registrygeoreplication:index:RegistryGeoReplication';

    /**
     * Returns true if the given object is an instance of RegistryGeoReplication.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegistryGeoReplication {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegistryGeoReplication.__pulumiType;
    }

    /**
     * The login server url
     */
    public /*out*/ readonly acrLoginServerOut!: pulumi.Output<string>;
    /**
     * The Registry
     */
    public /*out*/ readonly registry!: pulumi.Output<pulumiAzureNative.containerregistry.Registry>;
    /**
     * The replication policy
     */
    public /*out*/ readonly replication!: pulumi.Output<pulumiAzureNative.containerregistry.Replication>;

    /**
     * Create a RegistryGeoReplication resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegistryGeoReplicationArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceGroup === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroup'");
            }
            resourceInputs["resourceGroup"] = args ? args.resourceGroup : undefined;
            resourceInputs["acrLoginServerOut"] = undefined /*out*/;
            resourceInputs["registry"] = undefined /*out*/;
            resourceInputs["replication"] = undefined /*out*/;
        } else {
            resourceInputs["acrLoginServerOut"] = undefined /*out*/;
            resourceInputs["registry"] = undefined /*out*/;
            resourceInputs["replication"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RegistryGeoReplication.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a RegistryGeoReplication resource.
 */
export interface RegistryGeoReplicationArgs {
    /**
     * The resource group that hosts the component resource
     */
    resourceGroup: pulumi.Input<pulumiAzureNative.resources.ResourceGroup>;
}
