// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

import * as aws from "@pulumi/aws";

export interface EventHandlerRoute {
    function?: pulumi.Input<aws.lambda.Function>;
    method?: pulumi.Input<string>;
    path?: pulumi.Input<string>;
}