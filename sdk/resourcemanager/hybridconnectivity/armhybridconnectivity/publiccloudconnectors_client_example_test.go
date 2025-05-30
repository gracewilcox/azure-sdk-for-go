// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armhybridconnectivity_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity"
	"log"
)

// Generated from example definition: 2024-12-01/PublicCloudConnectors_CreateOrUpdate.json
func ExamplePublicCloudConnectorsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPublicCloudConnectorsClient("5ACC4579-DB34-4C2F-8F8C-25061168F342").BeginCreateOrUpdate(ctx, "rgpublicCloud", "advjwoakdusalamomg", armhybridconnectivity.PublicCloudConnector{
		Properties: &armhybridconnectivity.PublicCloudConnectorProperties{
			AwsCloudProfile: &armhybridconnectivity.AwsCloudProfile{
				AccountID: to.Ptr("snbnuxckevyqpm"),
				ExcludedAccounts: []*string{
					to.Ptr("rwgqpukglvbqmogqcliqolucp"),
				},
				IsOrganizationalAccount: to.Ptr(true),
			},
			HostType: to.Ptr(armhybridconnectivity.HostTypeAWS),
		},
		Tags:     map[string]*string{},
		Location: to.Ptr("jpiglusfxynfcewcjwvvnn"),
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
	// res = armhybridconnectivity.PublicCloudConnectorsClientCreateOrUpdateResponse{
	// 	PublicCloudConnector: &armhybridconnectivity.PublicCloudConnector{
	// 		Properties: &armhybridconnectivity.PublicCloudConnectorProperties{
	// 			AwsCloudProfile: &armhybridconnectivity.AwsCloudProfile{
	// 				AccountID: to.Ptr("snbnuxckevyqpm"),
	// 				ExcludedAccounts: []*string{
	// 					to.Ptr("rwgqpukglvbqmogqcliqolucp"),
	// 				},
	// 				IsOrganizationalAccount: to.Ptr(true),
	// 			},
	// 			HostType: to.Ptr(armhybridconnectivity.HostTypeAWS),
	// 			ProvisioningState: to.Ptr(armhybridconnectivity.ResourceProvisioningStateSucceeded),
	// 			ConnectorPrimaryIdentifier: to.Ptr("20a4e2be-8158-4b9e-b512-7a1af6f827de"),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("jpiglusfxynfcewcjwvvnn"),
	// 		ID: to.Ptr("/subscriptions/5ACC4579-DB34-4C2F-8F8C-25061168F342/providers/Microsoft.HybridConnectivity/PublicCloudConnectors/esixipkbydb"),
	// 		Name: to.Ptr("esixipkbydb"),
	// 		Type: to.Ptr("eelsjvqvkdxdncptsobrswhulnm"),
	// 		SystemData: &armhybridconnectivity.SystemData{
	// 			CreatedBy: to.Ptr("rpxzkcrobprrdvuoqxz"),
	// 			CreatedByType: to.Ptr(armhybridconnectivity.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-18T22:52:07.890Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("jidegyskxi"),
	// 			LastModifiedByType: to.Ptr(armhybridconnectivity.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-18T22:52:07.890Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-12-01/PublicCloudConnectors_Delete.json
func ExamplePublicCloudConnectorsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPublicCloudConnectorsClient("5ACC4579-DB34-4C2F-8F8C-25061168F342").Delete(ctx, "rgpublicCloud", "skcfyjvflkhibdywjay", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhybridconnectivity.PublicCloudConnectorsClientDeleteResponse{
	// }
}

// Generated from example definition: 2024-12-01/PublicCloudConnectors_Get.json
func ExamplePublicCloudConnectorsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPublicCloudConnectorsClient("5ACC4579-DB34-4C2F-8F8C-25061168F342").Get(ctx, "rgpublicCloud", "rzygvnpsnrdylwzdbsscjazvamyxmh", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhybridconnectivity.PublicCloudConnectorsClientGetResponse{
	// 	PublicCloudConnector: &armhybridconnectivity.PublicCloudConnector{
	// 		Properties: &armhybridconnectivity.PublicCloudConnectorProperties{
	// 			AwsCloudProfile: &armhybridconnectivity.AwsCloudProfile{
	// 				AccountID: to.Ptr("snbnuxckevyqpm"),
	// 				ExcludedAccounts: []*string{
	// 					to.Ptr("rwgqpukglvbqmogqcliqolucp"),
	// 				},
	// 				IsOrganizationalAccount: to.Ptr(true),
	// 			},
	// 			HostType: to.Ptr(armhybridconnectivity.HostTypeAWS),
	// 			ProvisioningState: to.Ptr(armhybridconnectivity.ResourceProvisioningStateSucceeded),
	// 			ConnectorPrimaryIdentifier: to.Ptr("20a4e2be-8158-4b9e-b512-7a1af6f827de"),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("jpiglusfxynfcewcjwvvnn"),
	// 		ID: to.Ptr("/subscriptions/5ACC4579-DB34-4C2F-8F8C-25061168F342/providers/Microsoft.HybridConnectivity/PublicCloudConnectors/esixipkbydb"),
	// 		Name: to.Ptr("esixipkbydb"),
	// 		Type: to.Ptr("eelsjvqvkdxdncptsobrswhulnm"),
	// 		SystemData: &armhybridconnectivity.SystemData{
	// 			CreatedBy: to.Ptr("rpxzkcrobprrdvuoqxz"),
	// 			CreatedByType: to.Ptr(armhybridconnectivity.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-18T22:52:07.890Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("jidegyskxi"),
	// 			LastModifiedByType: to.Ptr(armhybridconnectivity.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-18T22:52:07.890Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-12-01/PublicCloudConnectors_ListByResourceGroup.json
func ExamplePublicCloudConnectorsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPublicCloudConnectorsClient("5ACC4579-DB34-4C2F-8F8C-25061168F342").NewListByResourceGroupPager("rgpublicCloud", nil)
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
		// page = armhybridconnectivity.PublicCloudConnectorsClientListByResourceGroupResponse{
		// 	PublicCloudConnectorListResult: armhybridconnectivity.PublicCloudConnectorListResult{
		// 		Value: []*armhybridconnectivity.PublicCloudConnector{
		// 			{
		// 				Properties: &armhybridconnectivity.PublicCloudConnectorProperties{
		// 					AwsCloudProfile: &armhybridconnectivity.AwsCloudProfile{
		// 						ExcludedAccounts: []*string{
		// 							to.Ptr("rwgqpukglvbqmogqcliqolucp"),
		// 						},
		// 						AccountID: to.Ptr("troiiavknxmcpczvxwjhrdue"),
		// 						IsOrganizationalAccount: to.Ptr(true),
		// 					},
		// 					HostType: to.Ptr(armhybridconnectivity.HostTypeAWS),
		// 					ProvisioningState: to.Ptr(armhybridconnectivity.ResourceProvisioningStateSucceeded),
		// 					ConnectorPrimaryIdentifier: to.Ptr("20a4e2be-8158-4b9e-b512-7a1af6f827de"),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 				Location: to.Ptr("jpiglusfxynfcewcjwvvnn"),
		// 				ID: to.Ptr("/subscriptions/5ACC4579-DB34-4C2F-8F8C-25061168F342/providers/Microsoft.HybridConnectivity/PublicCloudConnectors/esixipkbydb"),
		// 				Name: to.Ptr("esixipkbydb"),
		// 				Type: to.Ptr("eelsjvqvkdxdncptsobrswhulnm"),
		// 				SystemData: &armhybridconnectivity.SystemData{
		// 					CreatedBy: to.Ptr("rpxzkcrobprrdvuoqxz"),
		// 					CreatedByType: to.Ptr(armhybridconnectivity.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-18T22:52:07.890Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("jidegyskxi"),
		// 					LastModifiedByType: to.Ptr(armhybridconnectivity.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-18T22:52:07.890Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}

// Generated from example definition: 2024-12-01/PublicCloudConnectors_ListBySubscription.json
func ExamplePublicCloudConnectorsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPublicCloudConnectorsClient("5ACC4579-DB34-4C2F-8F8C-25061168F342").NewListBySubscriptionPager(nil)
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
		// page = armhybridconnectivity.PublicCloudConnectorsClientListBySubscriptionResponse{
		// 	PublicCloudConnectorListResult: armhybridconnectivity.PublicCloudConnectorListResult{
		// 		Value: []*armhybridconnectivity.PublicCloudConnector{
		// 			{
		// 				Properties: &armhybridconnectivity.PublicCloudConnectorProperties{
		// 					AwsCloudProfile: &armhybridconnectivity.AwsCloudProfile{
		// 						ExcludedAccounts: []*string{
		// 							to.Ptr("rwgqpukglvbqmogqcliqolucp"),
		// 						},
		// 						AccountID: to.Ptr("troiiavknxmcpczvxwjhrdue"),
		// 						IsOrganizationalAccount: to.Ptr(true),
		// 					},
		// 					HostType: to.Ptr(armhybridconnectivity.HostTypeAWS),
		// 					ProvisioningState: to.Ptr(armhybridconnectivity.ResourceProvisioningStateSucceeded),
		// 					ConnectorPrimaryIdentifier: to.Ptr("20a4e2be-8158-4b9e-b512-7a1af6f827de"),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 				Location: to.Ptr("jpiglusfxynfcewcjwvvnn"),
		// 				ID: to.Ptr("/subscriptions/5ACC4579-DB34-4C2F-8F8C-25061168F342/providers/Microsoft.HybridConnectivity/PublicCloudConnectors/esixipkbydb"),
		// 				Name: to.Ptr("esixipkbydb"),
		// 				Type: to.Ptr("eelsjvqvkdxdncptsobrswhulnm"),
		// 				SystemData: &armhybridconnectivity.SystemData{
		// 					CreatedBy: to.Ptr("rpxzkcrobprrdvuoqxz"),
		// 					CreatedByType: to.Ptr(armhybridconnectivity.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-18T22:52:07.890Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("jidegyskxi"),
		// 					LastModifiedByType: to.Ptr(armhybridconnectivity.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-18T22:52:07.890Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}

// Generated from example definition: 2024-12-01/PublicCloudConnectors_TestPermissions.json
func ExamplePublicCloudConnectorsClient_BeginTestPermissions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPublicCloudConnectorsClient("5ACC4579-DB34-4C2F-8F8C-25061168F342").BeginTestPermissions(ctx, "rgpublicCloud", "rzygvnpsnrdylwzdbsscjazvamyxmh", nil)
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
	// res = armhybridconnectivity.PublicCloudConnectorsClientTestPermissionsResponse{
	// 	OperationStatusResult: &armhybridconnectivity.OperationStatusResult{
	// 		ID: to.Ptr("/subscriptions/5ACC4579-DB34-4C2F-8F8C-25061168F342/providers/Microsoft.HybridConnectivity/PublicCloudConnectors/esixipkbydb"),
	// 		ResourceID: to.Ptr("/subscriptions/5ACC4579-DB34-4C2F-8F8C-25061168F342/providers/Microsoft.HybridConnectivity/PublicCloudConnectors/esixipkbydb"),
	// 		Name: to.Ptr("ppeygvsnaspxmpwalpmkqva"),
	// 		Status: to.Ptr("toyjllkvm"),
	// 		PercentComplete: to.Ptr[float64](81),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-02T18:38:19.143Z"); return t}()),
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-02T18:38:19.143Z"); return t}()),
	// 		Operations: []*armhybridconnectivity.OperationStatusResult{
	// 			{
	// 				ID: to.Ptr("/subscriptions/5ACC4579-DB34-4C2F-8F8C-25061168F342/providers/Microsoft.HybridConnectivity/PublicCloudConnectors/esixipkbydb"),
	// 				ResourceID: to.Ptr("/subscriptions/5ACC4579-DB34-4C2F-8F8C-25061168F342/providers/Microsoft.HybridConnectivity/PublicCloudConnectors/esixipkbydb"),
	// 				Name: to.Ptr("svqtraeuwvyvblujlvqilypwpdrt"),
	// 				Status: to.Ptr("bevmrejij"),
	// 				PercentComplete: to.Ptr[float64](15),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-02T18:38:19.143Z"); return t}()),
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-02T18:38:19.143Z"); return t}()),
	// 				Operations: []*armhybridconnectivity.OperationStatusResult{
	// 				},
	// 				Error: &armhybridconnectivity.ErrorDetail{
	// 					Code: to.Ptr("ykzvluyqiftfsumgvwzdh"),
	// 					Message: to.Ptr("krbjgtqkjgiux"),
	// 					Target: to.Ptr("nsaucxt"),
	// 					Details: []*armhybridconnectivity.ErrorDetail{
	// 					},
	// 					AdditionalInfo: []*armhybridconnectivity.ErrorAdditionalInfo{
	// 						{
	// 							Type: to.Ptr("qivvrewsjvcildjgwwytgimwklh"),
	// 							Info: &armhybridconnectivity.ErrorAdditionalInfoInfo{
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Error: &armhybridconnectivity.ErrorDetail{
	// 			Code: to.Ptr("ykzvluyqiftfsumgvwzdh"),
	// 			Message: to.Ptr("krbjgtqkjgiux"),
	// 			Target: to.Ptr("nsaucxt"),
	// 			Details: []*armhybridconnectivity.ErrorDetail{
	// 			},
	// 			AdditionalInfo: []*armhybridconnectivity.ErrorAdditionalInfo{
	// 				{
	// 					Type: to.Ptr("qivvrewsjvcildjgwwytgimwklh"),
	// 					Info: &armhybridconnectivity.ErrorAdditionalInfoInfo{
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-12-01/PublicCloudConnectors_Update.json
func ExamplePublicCloudConnectorsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPublicCloudConnectorsClient("5ACC4579-DB34-4C2F-8F8C-25061168F342").Update(ctx, "rgpublicCloud", "svtirlbyqpepbzyessjenlueeznhg", armhybridconnectivity.PublicCloudConnector{
		Tags: map[string]*string{},
		Properties: &armhybridconnectivity.PublicCloudConnectorProperties{
			AwsCloudProfile: &armhybridconnectivity.AwsCloudProfile{
				ExcludedAccounts: []*string{
					to.Ptr("zrbtd"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhybridconnectivity.PublicCloudConnectorsClientUpdateResponse{
	// 	PublicCloudConnector: &armhybridconnectivity.PublicCloudConnector{
	// 		Properties: &armhybridconnectivity.PublicCloudConnectorProperties{
	// 			AwsCloudProfile: &armhybridconnectivity.AwsCloudProfile{
	// 				AccountID: to.Ptr("snbnuxckevyqpm"),
	// 				ExcludedAccounts: []*string{
	// 					to.Ptr("zrbtd"),
	// 				},
	// 				IsOrganizationalAccount: to.Ptr(true),
	// 			},
	// 			HostType: to.Ptr(armhybridconnectivity.HostTypeAWS),
	// 			ProvisioningState: to.Ptr(armhybridconnectivity.ResourceProvisioningStateSucceeded),
	// 			ConnectorPrimaryIdentifier: to.Ptr("20a4e2be-8158-4b9e-b512-7a1af6f827de"),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("jpiglusfxynfcewcjwvvnn"),
	// 		ID: to.Ptr("/subscriptions/5ACC4579-DB34-4C2F-8F8C-25061168F342/providers/Microsoft.HybridConnectivity/PublicCloudConnectors/esixipkbydb"),
	// 		Name: to.Ptr("esixipkbydb"),
	// 		Type: to.Ptr("eelsjvqvkdxdncptsobrswhulnm"),
	// 		SystemData: &armhybridconnectivity.SystemData{
	// 			CreatedBy: to.Ptr("rpxzkcrobprrdvuoqxz"),
	// 			CreatedByType: to.Ptr(armhybridconnectivity.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-18T22:52:07.890Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("jidegyskxi"),
	// 			LastModifiedByType: to.Ptr(armhybridconnectivity.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-18T22:52:07.890Z"); return t}()),
	// 		},
	// 	},
	// }
}
