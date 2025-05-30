//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v10"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8eb3f7a4f66d408152c32b9d647e59147172d533/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ExposureControl_GetFeatureValue.json
func ExampleExposureControlClient_GetFeatureValue() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExposureControlClient().GetFeatureValue(ctx, "WestEurope", armdatafactory.ExposureControlRequest{
		FeatureName: to.Ptr("ADFIntegrationRuntimeSharingRbac"),
		FeatureType: to.Ptr("Feature"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExposureControlResponse = armdatafactory.ExposureControlResponse{
	// 	FeatureName: to.Ptr("ADFIntegrationRuntimeSharingRbac"),
	// 	Value: to.Ptr("False"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8eb3f7a4f66d408152c32b9d647e59147172d533/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ExposureControl_GetFeatureValueByFactory.json
func ExampleExposureControlClient_GetFeatureValueByFactory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExposureControlClient().GetFeatureValueByFactory(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.ExposureControlRequest{
		FeatureName: to.Ptr("ADFIntegrationRuntimeSharingRbac"),
		FeatureType: to.Ptr("Feature"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExposureControlResponse = armdatafactory.ExposureControlResponse{
	// 	FeatureName: to.Ptr("ADFIntegrationRuntimeSharingRbac"),
	// 	Value: to.Ptr("False"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8eb3f7a4f66d408152c32b9d647e59147172d533/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ExposureControl_QueryFeatureValuesByFactory.json
func ExampleExposureControlClient_QueryFeatureValuesByFactory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExposureControlClient().QueryFeatureValuesByFactory(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.ExposureControlBatchRequest{
		ExposureControlRequests: []*armdatafactory.ExposureControlRequest{
			{
				FeatureName: to.Ptr("ADFIntegrationRuntimeSharingRbac"),
				FeatureType: to.Ptr("Feature"),
			},
			{
				FeatureName: to.Ptr("ADFSampleFeature"),
				FeatureType: to.Ptr("Feature"),
			}},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExposureControlBatchResponse = armdatafactory.ExposureControlBatchResponse{
	// 	ExposureControlResponses: []*armdatafactory.ExposureControlResponse{
	// 		{
	// 			FeatureName: to.Ptr("ADFIntegrationRuntimeSharingRbac"),
	// 			Value: to.Ptr("False"),
	// 		},
	// 		{
	// 			FeatureName: to.Ptr("ADFSampleFeature"),
	// 			Value: to.Ptr("True"),
	// 	}},
	// }
}
