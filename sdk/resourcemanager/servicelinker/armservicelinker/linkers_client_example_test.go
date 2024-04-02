//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armservicelinker_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicelinker/armservicelinker/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c16ce913afbdaa073e8ca5e480f3b465db2de542/specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2023-04-01-preview/examples/ListDryrun.json
func ExampleLinkersClient_NewListDryrunPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicelinker.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLinkersClient().NewListDryrunPager("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.DryrunList = armservicelinker.DryrunList{
		// 	Value: []*armservicelinker.DryrunResource{
		// 		{
		// 			Name: to.Ptr("dryrunName"),
		// 			Type: to.Ptr("Microsoft.ServiceLinker/dryruns"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app/providers/Microsoft.ServiceLinker/dryruns/dryrunName"),
		// 			SystemData: &armservicelinker.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-12T22:05:09.000Z"); return t}()),
		// 			},
		// 			Properties: &armservicelinker.DryrunProperties{
		// 				Parameters: &armservicelinker.CreateOrUpdateDryrunParameters{
		// 					ActionName: to.Ptr(armservicelinker.DryrunActionNameCreateOrUpdate),
		// 					AuthInfo: &armservicelinker.SecretAuthInfo{
		// 						AuthType: to.Ptr(armservicelinker.AuthTypeSecret),
		// 						Name: to.Ptr("username"),
		// 					},
		// 					TargetService: &armservicelinker.AzureResource{
		// 						Type: to.Ptr(armservicelinker.TargetServiceTypeAzureResource),
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc/mongodbDatabases/test-db"),
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c16ce913afbdaa073e8ca5e480f3b465db2de542/specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2023-04-01-preview/examples/GetDryrun.json
func ExampleLinkersClient_GetDryrun() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicelinker.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLinkersClient().GetDryrun(ctx, "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app", "dryrunName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DryrunResource = armservicelinker.DryrunResource{
	// 	Name: to.Ptr("dryrunName"),
	// 	Type: to.Ptr("Microsoft.ServiceLinker/dryruns"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app/providers/Microsoft.ServiceLinker/dryruns/dryrunName"),
	// 	SystemData: &armservicelinker.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-12T22:05:09.000Z"); return t}()),
	// 	},
	// 	Properties: &armservicelinker.DryrunProperties{
	// 		Parameters: &armservicelinker.CreateOrUpdateDryrunParameters{
	// 			ActionName: to.Ptr(armservicelinker.DryrunActionNameCreateOrUpdate),
	// 			AuthInfo: &armservicelinker.SecretAuthInfo{
	// 				AuthType: to.Ptr(armservicelinker.AuthTypeSecret),
	// 				Name: to.Ptr("username"),
	// 			},
	// 			TargetService: &armservicelinker.AzureResource{
	// 				Type: to.Ptr(armservicelinker.TargetServiceTypeAzureResource),
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc/mongodbDatabases/test-db"),
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c16ce913afbdaa073e8ca5e480f3b465db2de542/specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2023-04-01-preview/examples/PutDryrun.json
func ExampleLinkersClient_BeginCreateDryrun() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicelinker.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLinkersClient().BeginCreateDryrun(ctx, "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app", "dryrunName", armservicelinker.DryrunResource{
		Properties: &armservicelinker.DryrunProperties{
			Parameters: &armservicelinker.CreateOrUpdateDryrunParameters{
				ActionName: to.Ptr(armservicelinker.DryrunActionNameCreateOrUpdate),
				AuthInfo: &armservicelinker.SecretAuthInfo{
					AuthType: to.Ptr(armservicelinker.AuthTypeSecret),
					Name:     to.Ptr("name"),
					SecretInfo: &armservicelinker.ValueSecretInfo{
						SecretType: to.Ptr(armservicelinker.SecretTypeRawValue),
						Value:      to.Ptr("secret"),
					},
				},
				TargetService: &armservicelinker.AzureResource{
					Type: to.Ptr(armservicelinker.TargetServiceTypeAzureResource),
					ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc/mongodbDatabases/test-db"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DryrunResource = armservicelinker.DryrunResource{
	// 	Name: to.Ptr("dryrunName"),
	// 	Type: to.Ptr("Microsoft.ServiceLinker/dryruns"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app/providers/Microsoft.ServiceLinker/dryruns/dryrunName"),
	// 	Properties: &armservicelinker.DryrunProperties{
	// 		OperationPreviews: []*armservicelinker.DryrunOperationPreview{
	// 			{
	// 				Name: to.Ptr("configFirewallRule"),
	// 				Description: to.Ptr("Config firewall rule for target service to allow source service access"),
	// 				Action: to.Ptr("Microsoft.DocumentDb/databaseAccounts/write"),
	// 				OperationType: to.Ptr(armservicelinker.DryrunPreviewOperationTypeConfigNetwork),
	// 				Scope: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc"),
	// 		}},
	// 		Parameters: &armservicelinker.CreateOrUpdateDryrunParameters{
	// 			ActionName: to.Ptr(armservicelinker.DryrunActionNameCreateOrUpdate),
	// 			AuthInfo: &armservicelinker.SecretAuthInfo{
	// 				AuthType: to.Ptr(armservicelinker.AuthTypeSecret),
	// 				Name: to.Ptr("name"),
	// 			},
	// 			TargetService: &armservicelinker.AzureResource{
	// 				Type: to.Ptr(armservicelinker.TargetServiceTypeAzureResource),
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc/mongodbDatabases/test-db"),
	// 			},
	// 		},
	// 		PrerequisiteResults: []armservicelinker.DryrunPrerequisiteResultClassification{
	// 			&armservicelinker.BasicErrorDryrunPrerequisiteResult{
	// 				Type: to.Ptr(armservicelinker.DryrunPrerequisiteResultTypeBasicError),
	// 				Code: to.Ptr("ResourceNotFound"),
	// 				Message: to.Ptr("Target resource is not found"),
	// 			},
	// 			&armservicelinker.PermissionsMissingDryrunPrerequisiteResult{
	// 				Type: to.Ptr(armservicelinker.DryrunPrerequisiteResultTypePermissionsMissing),
	// 				Permissions: []*string{
	// 					to.Ptr("Microsoft.DocumentDb/databaseAccounts/write")},
	// 					Scope: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc"),
	// 			}},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c16ce913afbdaa073e8ca5e480f3b465db2de542/specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2023-04-01-preview/examples/PatchDryrun.json
func ExampleLinkersClient_BeginUpdateDryrun() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicelinker.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLinkersClient().BeginUpdateDryrun(ctx, "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app", "dryrunName", armservicelinker.DryrunPatch{
		Properties: &armservicelinker.DryrunProperties{
			Parameters: &armservicelinker.CreateOrUpdateDryrunParameters{
				ActionName: to.Ptr(armservicelinker.DryrunActionNameCreateOrUpdate),
				AuthInfo: &armservicelinker.SecretAuthInfo{
					AuthType: to.Ptr(armservicelinker.AuthTypeSecret),
					Name:     to.Ptr("name"),
					SecretInfo: &armservicelinker.ValueSecretInfo{
						SecretType: to.Ptr(armservicelinker.SecretTypeRawValue),
						Value:      to.Ptr("secret"),
					},
				},
				TargetService: &armservicelinker.AzureResource{
					Type: to.Ptr(armservicelinker.TargetServiceTypeAzureResource),
					ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc/mongodbDatabases/test-db"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DryrunResource = armservicelinker.DryrunResource{
	// 	Name: to.Ptr("dryrunName"),
	// 	Type: to.Ptr("Microsoft.ServiceLinker/dryruns"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app/providers/Microsoft.ServiceLinker/dryruns/dryrunName"),
	// 	Properties: &armservicelinker.DryrunProperties{
	// 		OperationPreviews: []*armservicelinker.DryrunOperationPreview{
	// 			{
	// 				Name: to.Ptr("configFirewallRule"),
	// 				Description: to.Ptr("Config firewall rule for target service to allow source service access"),
	// 				Action: to.Ptr("Microsoft.DocumentDb/databaseAccounts/write"),
	// 				OperationType: to.Ptr(armservicelinker.DryrunPreviewOperationTypeConfigNetwork),
	// 				Scope: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc"),
	// 		}},
	// 		Parameters: &armservicelinker.CreateOrUpdateDryrunParameters{
	// 			ActionName: to.Ptr(armservicelinker.DryrunActionNameCreateOrUpdate),
	// 			AuthInfo: &armservicelinker.SecretAuthInfo{
	// 				AuthType: to.Ptr(armservicelinker.AuthTypeSecret),
	// 				Name: to.Ptr("name"),
	// 			},
	// 			TargetService: &armservicelinker.AzureResource{
	// 				Type: to.Ptr(armservicelinker.TargetServiceTypeAzureResource),
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc/mongodbDatabases/test-db"),
	// 			},
	// 		},
	// 		PrerequisiteResults: []armservicelinker.DryrunPrerequisiteResultClassification{
	// 			&armservicelinker.BasicErrorDryrunPrerequisiteResult{
	// 				Type: to.Ptr(armservicelinker.DryrunPrerequisiteResultTypeBasicError),
	// 				Code: to.Ptr("ResourceNotFound"),
	// 				Message: to.Ptr("Target resource is not found"),
	// 			},
	// 			&armservicelinker.PermissionsMissingDryrunPrerequisiteResult{
	// 				Type: to.Ptr(armservicelinker.DryrunPrerequisiteResultTypePermissionsMissing),
	// 				Permissions: []*string{
	// 					to.Ptr("Microsoft.DocumentDb/databaseAccounts/write")},
	// 					Scope: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc"),
	// 			}},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c16ce913afbdaa073e8ca5e480f3b465db2de542/specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2023-04-01-preview/examples/DeleteDryrun.json
func ExampleLinkersClient_DeleteDryrun() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicelinker.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewLinkersClient().DeleteDryrun(ctx, "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app", "dryrunName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c16ce913afbdaa073e8ca5e480f3b465db2de542/specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2023-04-01-preview/examples/LinkerGenerateConfigurations.json
func ExampleLinkersClient_GenerateConfigurations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicelinker.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLinkersClient().GenerateConfigurations(ctx, "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app", "linkName", &armservicelinker.LinkersClientGenerateConfigurationsOptions{Parameters: &armservicelinker.ConfigurationInfo{
		CustomizedKeys: map[string]*string{
			"ASL_DocumentDb_ConnectionString": to.Ptr("MyConnectionstring"),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConfigurationResult = armservicelinker.ConfigurationResult{
	// 	Configurations: []*armservicelinker.SourceConfiguration{
	// 		{
	// 			Name: to.Ptr("MyConnectionstring"),
	// 			Value: to.Ptr("ConnectionString"),
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c16ce913afbdaa073e8ca5e480f3b465db2de542/specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2023-04-01-preview/examples/GetDaprConfigurations.json
func ExampleLinkersClient_NewListDaprConfigurationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicelinker.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLinkersClient().NewListDaprConfigurationsPager("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.DaprConfigurationList = armservicelinker.DaprConfigurationList{
		// 	Value: []*armservicelinker.DaprConfigurationResource{
		// 		{
		// 			Properties: &armservicelinker.DaprConfigurationProperties{
		// 				AuthType: to.Ptr(armservicelinker.AuthTypeSecret),
		// 				DaprProperties: &armservicelinker.DaprProperties{
		// 					BindingComponentDirection: to.Ptr(armservicelinker.DaprBindingComponentDirectionInput),
		// 					ComponentType: to.Ptr("bindings"),
		// 					Metadata: []*armservicelinker.DaprMetadata{
		// 						{
		// 							Name: to.Ptr("containerName"),
		// 							Description: to.Ptr("The name of the container to be used for Dapr state."),
		// 							Required: to.Ptr(armservicelinker.DaprMetadataRequiredTrue),
		// 					}},
		// 					RuntimeVersion: to.Ptr("1.10"),
		// 					Version: to.Ptr("v1"),
		// 				},
		// 				TargetType: to.Ptr("MICROSOFT.STORAGE/STORAGEACCOUNTS/BLOBSERVICES"),
		// 			},
		// 	}},
		// }
	}
}