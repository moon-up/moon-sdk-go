# \ERC1155API

All URIs are relative to *https://beta.usemoon.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BalanceOf**](ERC1155API.md#BalanceOf) | **Post** /erc1155/{name}/balance-of | 
[**BalanceOfBatch**](ERC1155API.md#BalanceOfBatch) | **Post** /erc1155/{name}/balance-of-batch | 
[**IsApprovedForAll**](ERC1155API.md#IsApprovedForAll) | **Post** /erc1155/{name}/is-approved-for-all | 
[**SafeBatchTransferFrom**](ERC1155API.md#SafeBatchTransferFrom) | **Post** /erc1155/{name}/safe-batch-transfer-from | 
[**SafeTransferFrom**](ERC1155API.md#SafeTransferFrom) | **Post** /erc1155/{name}/safe-transfer-from | 
[**SetApprovalForAll**](ERC1155API.md#SetApprovalForAll) | **Post** /erc1155/{name}/set-approval-for-all | 



## BalanceOf

> TransactionAPIResponse BalanceOf(ctx, name).Authorization(authorization).Erc1155Request(erc1155Request).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/moon-up/moon-sdk-go"
)

func main() {
	name := "name_example" // string | 
	authorization := "authorization_example" // string | 
	erc1155Request := *openapiclient.NewErc1155Request() // Erc1155Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ERC1155API.BalanceOf(context.Background(), name).Authorization(authorization).Erc1155Request(erc1155Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ERC1155API.BalanceOf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BalanceOf`: TransactionAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ERC1155API.BalanceOf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBalanceOfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **erc1155Request** | [**Erc1155Request**](Erc1155Request.md) |  | 

### Return type

[**TransactionAPIResponse**](TransactionAPIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BalanceOfBatch

> TransactionAPIResponse BalanceOfBatch(ctx, name).Authorization(authorization).Erc1155Request(erc1155Request).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/moon-up/moon-sdk-go"
)

func main() {
	name := "name_example" // string | 
	authorization := "authorization_example" // string | 
	erc1155Request := *openapiclient.NewErc1155Request() // Erc1155Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ERC1155API.BalanceOfBatch(context.Background(), name).Authorization(authorization).Erc1155Request(erc1155Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ERC1155API.BalanceOfBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BalanceOfBatch`: TransactionAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ERC1155API.BalanceOfBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBalanceOfBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **erc1155Request** | [**Erc1155Request**](Erc1155Request.md) |  | 

### Return type

[**TransactionAPIResponse**](TransactionAPIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IsApprovedForAll

> TransactionAPIResponse IsApprovedForAll(ctx, name).Authorization(authorization).Erc1155Request(erc1155Request).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/moon-up/moon-sdk-go"
)

func main() {
	name := "name_example" // string | 
	authorization := "authorization_example" // string | 
	erc1155Request := *openapiclient.NewErc1155Request() // Erc1155Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ERC1155API.IsApprovedForAll(context.Background(), name).Authorization(authorization).Erc1155Request(erc1155Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ERC1155API.IsApprovedForAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IsApprovedForAll`: TransactionAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ERC1155API.IsApprovedForAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIsApprovedForAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **erc1155Request** | [**Erc1155Request**](Erc1155Request.md) |  | 

### Return type

[**TransactionAPIResponse**](TransactionAPIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SafeBatchTransferFrom

> TransactionAPIResponse SafeBatchTransferFrom(ctx, name).Authorization(authorization).Erc1155Request(erc1155Request).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/moon-up/moon-sdk-go"
)

func main() {
	name := "name_example" // string | 
	authorization := "authorization_example" // string | 
	erc1155Request := *openapiclient.NewErc1155Request() // Erc1155Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ERC1155API.SafeBatchTransferFrom(context.Background(), name).Authorization(authorization).Erc1155Request(erc1155Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ERC1155API.SafeBatchTransferFrom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SafeBatchTransferFrom`: TransactionAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ERC1155API.SafeBatchTransferFrom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSafeBatchTransferFromRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **erc1155Request** | [**Erc1155Request**](Erc1155Request.md) |  | 

### Return type

[**TransactionAPIResponse**](TransactionAPIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SafeTransferFrom

> TransactionAPIResponse SafeTransferFrom(ctx, name).Authorization(authorization).Erc1155Request(erc1155Request).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/moon-up/moon-sdk-go"
)

func main() {
	name := "name_example" // string | 
	authorization := "authorization_example" // string | 
	erc1155Request := *openapiclient.NewErc1155Request() // Erc1155Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ERC1155API.SafeTransferFrom(context.Background(), name).Authorization(authorization).Erc1155Request(erc1155Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ERC1155API.SafeTransferFrom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SafeTransferFrom`: TransactionAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ERC1155API.SafeTransferFrom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSafeTransferFromRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **erc1155Request** | [**Erc1155Request**](Erc1155Request.md) |  | 

### Return type

[**TransactionAPIResponse**](TransactionAPIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetApprovalForAll

> TransactionAPIResponse SetApprovalForAll(ctx, name).Authorization(authorization).Erc1155Request(erc1155Request).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/moon-up/moon-sdk-go"
)

func main() {
	name := "name_example" // string | 
	authorization := "authorization_example" // string | 
	erc1155Request := *openapiclient.NewErc1155Request() // Erc1155Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ERC1155API.SetApprovalForAll(context.Background(), name).Authorization(authorization).Erc1155Request(erc1155Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ERC1155API.SetApprovalForAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetApprovalForAll`: TransactionAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ERC1155API.SetApprovalForAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetApprovalForAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **erc1155Request** | [**Erc1155Request**](Erc1155Request.md) |  | 

### Return type

[**TransactionAPIResponse**](TransactionAPIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

