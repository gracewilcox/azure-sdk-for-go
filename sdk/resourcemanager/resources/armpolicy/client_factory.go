// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpolicy

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	internal       *arm.Client
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	internal, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID,
		internal:       internal,
	}, nil
}

// NewAssignmentsClient creates a new instance of AssignmentsClient.
func (c *ClientFactory) NewAssignmentsClient() *AssignmentsClient {
	return &AssignmentsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDataPolicyManifestsClient creates a new instance of DataPolicyManifestsClient.
func (c *ClientFactory) NewDataPolicyManifestsClient() *DataPolicyManifestsClient {
	return &DataPolicyManifestsClient{
		internal: c.internal,
	}
}

// NewDefinitionVersionsClient creates a new instance of DefinitionVersionsClient.
func (c *ClientFactory) NewDefinitionVersionsClient() *DefinitionVersionsClient {
	return &DefinitionVersionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDefinitionsClient creates a new instance of DefinitionsClient.
func (c *ClientFactory) NewDefinitionsClient() *DefinitionsClient {
	return &DefinitionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewExemptionsClient creates a new instance of ExemptionsClient.
func (c *ClientFactory) NewExemptionsClient() *ExemptionsClient {
	return &ExemptionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSetDefinitionVersionsClient creates a new instance of SetDefinitionVersionsClient.
func (c *ClientFactory) NewSetDefinitionVersionsClient() *SetDefinitionVersionsClient {
	return &SetDefinitionVersionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSetDefinitionsClient creates a new instance of SetDefinitionsClient.
func (c *ClientFactory) NewSetDefinitionsClient() *SetDefinitionsClient {
	return &SetDefinitionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewVariableValuesClient creates a new instance of VariableValuesClient.
func (c *ClientFactory) NewVariableValuesClient() *VariableValuesClient {
	return &VariableValuesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewVariablesClient creates a new instance of VariablesClient.
func (c *ClientFactory) NewVariablesClient() *VariablesClient {
	return &VariablesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
