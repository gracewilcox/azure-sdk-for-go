package storageapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2016-12-01/storage"
	"github.com/Azure/go-autorest/autorest"
)

// AccountsClientAPI contains the set of methods on the AccountsClient type.
type AccountsClientAPI interface {
	CheckNameAvailability(ctx context.Context, accountName storage.AccountCheckNameAvailabilityParameters) (result storage.CheckNameAvailabilityResult, err error)
	Create(ctx context.Context, resourceGroupName string, accountName string, parameters storage.AccountCreateParameters) (result storage.AccountsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string) (result autorest.Response, err error)
	GetProperties(ctx context.Context, resourceGroupName string, accountName string) (result storage.Account, err error)
	List(ctx context.Context) (result storage.AccountListResult, err error)
	ListAccountSAS(ctx context.Context, resourceGroupName string, accountName string, parameters storage.AccountSasParameters) (result storage.ListAccountSasResponse, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result storage.AccountListResult, err error)
	ListKeys(ctx context.Context, resourceGroupName string, accountName string) (result storage.AccountListKeysResult, err error)
	ListServiceSAS(ctx context.Context, resourceGroupName string, accountName string, parameters storage.ServiceSasParameters) (result storage.ListServiceSasResponse, err error)
	RegenerateKey(ctx context.Context, resourceGroupName string, accountName string, regenerateKey storage.AccountRegenerateKeyParameters) (result storage.AccountListKeysResult, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, parameters storage.AccountUpdateParameters) (result storage.Account, err error)
}

var _ AccountsClientAPI = (*storage.AccountsClient)(nil)

// UsageClientAPI contains the set of methods on the UsageClient type.
type UsageClientAPI interface {
	List(ctx context.Context) (result storage.UsageListResult, err error)
}

var _ UsageClientAPI = (*storage.UsageClient)(nil)
