//go:generate tsp-client update --local-spec-repo /home/grace/code/azure-rest-api-specs/specification/keyvault/Security.KeyVault.Certificates
//go:generate go run internal/transforms.go
//go:generate goimports -w .

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azcertificates
