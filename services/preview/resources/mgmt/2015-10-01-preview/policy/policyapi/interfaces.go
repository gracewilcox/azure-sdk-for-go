package policyapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy"
	"github.com/Azure/go-autorest/autorest"
)

// AssignmentsClientAPI contains the set of methods on the AssignmentsClient type.
type AssignmentsClientAPI interface {
	Create(ctx context.Context, scope string, policyAssignmentName string, parameters policy.Assignment) (result policy.Assignment, err error)
	CreateByID(ctx context.Context, policyAssignmentID string, parameters policy.Assignment) (result policy.Assignment, err error)
	Delete(ctx context.Context, scope string, policyAssignmentName string) (result policy.Assignment, err error)
	DeleteByID(ctx context.Context, policyAssignmentID string) (result policy.Assignment, err error)
	Get(ctx context.Context, scope string, policyAssignmentName string) (result policy.Assignment, err error)
	GetByID(ctx context.Context, policyAssignmentID string) (result policy.Assignment, err error)
	List(ctx context.Context, filter string) (result policy.AssignmentListResultPage, err error)
	ListComplete(ctx context.Context, filter string) (result policy.AssignmentListResultIterator, err error)
	ListForResource(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string) (result policy.AssignmentListResultPage, err error)
	ListForResourceComplete(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string) (result policy.AssignmentListResultIterator, err error)
	ListForResourceGroup(ctx context.Context, resourceGroupName string, filter string) (result policy.AssignmentListResultPage, err error)
	ListForResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string) (result policy.AssignmentListResultIterator, err error)
}

var _ AssignmentsClientAPI = (*policy.AssignmentsClient)(nil)

// DefinitionsClientAPI contains the set of methods on the DefinitionsClient type.
type DefinitionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, policyDefinitionName string, parameters policy.Definition) (result policy.Definition, err error)
	Delete(ctx context.Context, policyDefinitionName string) (result autorest.Response, err error)
	Get(ctx context.Context, policyDefinitionName string) (result policy.Definition, err error)
	List(ctx context.Context, filter string) (result policy.DefinitionListResultPage, err error)
	ListComplete(ctx context.Context, filter string) (result policy.DefinitionListResultIterator, err error)
}

var _ DefinitionsClientAPI = (*policy.DefinitionsClient)(nil)
