module github.com/Azure/azure-sdk-for-go/sdk/azcore

go 1.18

replace github.com/Azure/azure-sdk-for-go/sdk/tscore => ../tscore

require (
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.5.2
	github.com/Azure/azure-sdk-for-go/sdk/tscore v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.8.4
	golang.org/x/net v0.20.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
