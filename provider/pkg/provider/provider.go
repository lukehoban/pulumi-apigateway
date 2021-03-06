// Copyright 2016-2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"github.com/pkg/errors"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func makeProviderArgs(schema []byte) pulumi.ProviderArgs {
	return pulumi.ProviderArgs{
		Schema: schema,
		Construct: func(ctx *pulumi.Context, typ, name string, inputs *pulumi.ConstructInputs,
			options pulumi.ResourceOption) (pulumi.ConstructResult, error) {

			// TODO: Add support for additional component resources here.
			switch typ {
			case "apigateway:index:RestAPI":
				return constructRestAPI(ctx, name, inputs, options)
			default:
				return pulumi.ConstructResult{}, errors.Errorf("unknown resource type %s", typ)
			}
		},
	}
}

// constructRestAPI is an implementation of Construct for the example RestAPI component.
// It demonstrates converting the raw ConstructInputs to the component's args struct, creating
// the component, and returning its URN and state (outputs).
func constructRestAPI(ctx *pulumi.Context, name string, inputs *pulumi.ConstructInputs,
	options pulumi.ResourceOption) (pulumi.ConstructResult, error) {

	// Copy the raw inputs to RestAPIArgs. `inputs.SetArgs` uses the types and `pulumi:` tags
	// on the struct's fields to convert the raw values to the appropriate Input types.
	args := &RestAPIArgs{}
	if err := inputs.SetArgs(args); err != nil {
		return pulumi.ConstructResult{}, errors.Wrap(err, "setting args")
	}

	// Create the component resource.
	restAPI, err := NewRestAPI(ctx, name, args, options)
	if err != nil {
		return pulumi.ConstructResult{}, errors.Wrap(err, "creating component")
	}

	// Return the component resource's URN and outputs as its state.
	return pulumi.ConstructResult{
		URN: restAPI.URN(),
		State: pulumi.Map{
			"url":   restAPI.Url,
			"api":   restAPI.API,
			"stage": restAPI.Stage,
		},
	}, nil
}
