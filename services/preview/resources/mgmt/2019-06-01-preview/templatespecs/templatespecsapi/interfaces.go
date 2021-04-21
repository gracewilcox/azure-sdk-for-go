package templatespecsapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2019-06-01-preview/templatespecs"
	"github.com/Azure/go-autorest/autorest"
)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, templateSpecName string, templateSpec templatespecs.TemplateSpec) (result templatespecs.TemplateSpec, err error)
	Delete(ctx context.Context, resourceGroupName string, templateSpecName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, templateSpecName string, expand templatespecs.TemplateSpecExpandKind) (result templatespecs.TemplateSpec, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, expand templatespecs.TemplateSpecExpandKind) (result templatespecs.ListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, expand templatespecs.TemplateSpecExpandKind) (result templatespecs.ListResultIterator, err error)
	ListBySubscription(ctx context.Context, expand templatespecs.TemplateSpecExpandKind) (result templatespecs.ListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context, expand templatespecs.TemplateSpecExpandKind) (result templatespecs.ListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, templateSpecName string, templateSpec *templatespecs.UpdateModel) (result templatespecs.TemplateSpec, err error)
}

var _ ClientAPI = (*templatespecs.Client)(nil)

// VersionsClientAPI contains the set of methods on the VersionsClient type.
type VersionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, templateSpecName string, templateSpecVersion string, templateSpecVersionModel templatespecs.VersionTemplatespecs) (result templatespecs.VersionTemplatespecs, err error)
	Delete(ctx context.Context, resourceGroupName string, templateSpecName string, templateSpecVersion string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, templateSpecName string, templateSpecVersion string) (result templatespecs.VersionTemplatespecs, err error)
	List(ctx context.Context, resourceGroupName string, templateSpecName string) (result templatespecs.VersionsListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, templateSpecName string) (result templatespecs.VersionsListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, templateSpecName string, templateSpecVersion string, templateSpecVersionUpdateModel *templatespecs.VersionUpdateModel) (result templatespecs.VersionTemplatespecs, err error)
}

var _ VersionsClientAPI = (*templatespecs.VersionsClient)(nil)
