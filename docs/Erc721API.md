# \Erc721API

All URIs are relative to *https://vault-api.usemoon.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Approve**](Erc721API.md#Approve) | **Post** /erc721/{name}/approve | 
[**BalanceOf**](Erc721API.md#BalanceOf) | **Post** /erc721/{name}/balance-of | 
[**GetApproved**](Erc721API.md#GetApproved) | **Post** /erc721/{name}/get-approved | 
[**IsApprovedForAll**](Erc721API.md#IsApprovedForAll) | **Post** /erc721/{name}/is-approved-for-all | 
[**Name**](Erc721API.md#Name) | **Post** /erc721/{name}/name | 
[**OwnerOf**](Erc721API.md#OwnerOf) | **Post** /erc721/{name}/owner-of | 
[**SafeTransferFrom**](Erc721API.md#SafeTransferFrom) | **Post** /erc721/{name}/safe-transfer-from | 
[**SetApprovalForAll**](Erc721API.md#SetApprovalForAll) | **Post** /erc721/{name}/set-approval-for-all | 
[**Symbol**](Erc721API.md#Symbol) | **Post** /erc721/{name}/symbol | 
[**TokenUri**](Erc721API.md#TokenUri) | **Post** /erc721/{name}/token-uri | 
[**Transfer**](Erc721API.md#Transfer) | **Post** /erc721/{name}/transfer | 
[**TransferFrom**](Erc721API.md#TransferFrom) | **Post** /erc721/{name}/transfer-from | 



## Approve

> AccountControllerResponse Approve(ctx, name).Authorization(authorization).Erc721Request(erc721Request).Execute()



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
	authorization := "authorization_example" // string | 
	name := "name_example" // string | 
	erc721Request := *openapiclient.NewErc721Request() // Erc721Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Erc721API.Approve(context.Background(), name).Authorization(authorization).Erc721Request(erc721Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Erc721API.Approve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Approve`: AccountControllerResponse
	fmt.Fprintf(os.Stdout, "Response from `Erc721API.Approve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **erc721Request** | [**Erc721Request**](Erc721Request.md) |  | 

### Return type

[**AccountControllerResponse**](AccountControllerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BalanceOf

> AccountControllerResponse BalanceOf(ctx, name).Authorization(authorization).Erc721Request(erc721Request).Execute()



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
	authorization := "authorization_example" // string | 
	name := "name_example" // string | 
	erc721Request := *openapiclient.NewErc721Request() // Erc721Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Erc721API.BalanceOf(context.Background(), name).Authorization(authorization).Erc721Request(erc721Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Erc721API.BalanceOf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BalanceOf`: AccountControllerResponse
	fmt.Fprintf(os.Stdout, "Response from `Erc721API.BalanceOf`: %v\n", resp)
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

 **erc721Request** | [**Erc721Request**](Erc721Request.md) |  | 

### Return type

[**AccountControllerResponse**](AccountControllerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApproved

> AccountControllerResponse GetApproved(ctx, name).Authorization(authorization).Erc721Request(erc721Request).Execute()



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
	authorization := "authorization_example" // string | 
	name := "name_example" // string | 
	erc721Request := *openapiclient.NewErc721Request() // Erc721Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Erc721API.GetApproved(context.Background(), name).Authorization(authorization).Erc721Request(erc721Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Erc721API.GetApproved``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApproved`: AccountControllerResponse
	fmt.Fprintf(os.Stdout, "Response from `Erc721API.GetApproved`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApprovedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **erc721Request** | [**Erc721Request**](Erc721Request.md) |  | 

### Return type

[**AccountControllerResponse**](AccountControllerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IsApprovedForAll

> AccountControllerResponse IsApprovedForAll(ctx, name).Authorization(authorization).Erc721Request(erc721Request).Execute()



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
	authorization := "authorization_example" // string | 
	name := "name_example" // string | 
	erc721Request := *openapiclient.NewErc721Request() // Erc721Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Erc721API.IsApprovedForAll(context.Background(), name).Authorization(authorization).Erc721Request(erc721Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Erc721API.IsApprovedForAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IsApprovedForAll`: AccountControllerResponse
	fmt.Fprintf(os.Stdout, "Response from `Erc721API.IsApprovedForAll`: %v\n", resp)
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

 **erc721Request** | [**Erc721Request**](Erc721Request.md) |  | 

### Return type

[**AccountControllerResponse**](AccountControllerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Name

> AccountControllerResponse Name(ctx, name).Authorization(authorization).Erc721Request(erc721Request).Execute()



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
	authorization := "authorization_example" // string | 
	name := "name_example" // string | 
	erc721Request := *openapiclient.NewErc721Request() // Erc721Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Erc721API.Name(context.Background(), name).Authorization(authorization).Erc721Request(erc721Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Erc721API.Name``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Name`: AccountControllerResponse
	fmt.Fprintf(os.Stdout, "Response from `Erc721API.Name`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **erc721Request** | [**Erc721Request**](Erc721Request.md) |  | 

### Return type

[**AccountControllerResponse**](AccountControllerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OwnerOf

> AccountControllerResponse OwnerOf(ctx, name).Authorization(authorization).Erc721Request(erc721Request).Execute()



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
	authorization := "authorization_example" // string | 
	name := "name_example" // string | 
	erc721Request := *openapiclient.NewErc721Request() // Erc721Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Erc721API.OwnerOf(context.Background(), name).Authorization(authorization).Erc721Request(erc721Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Erc721API.OwnerOf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OwnerOf`: AccountControllerResponse
	fmt.Fprintf(os.Stdout, "Response from `Erc721API.OwnerOf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOwnerOfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **erc721Request** | [**Erc721Request**](Erc721Request.md) |  | 

### Return type

[**AccountControllerResponse**](AccountControllerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SafeTransferFrom

> AccountControllerResponse SafeTransferFrom(ctx, name).Authorization(authorization).Erc721Request(erc721Request).Execute()



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
	authorization := "authorization_example" // string | 
	name := "name_example" // string | 
	erc721Request := *openapiclient.NewErc721Request() // Erc721Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Erc721API.SafeTransferFrom(context.Background(), name).Authorization(authorization).Erc721Request(erc721Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Erc721API.SafeTransferFrom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SafeTransferFrom`: AccountControllerResponse
	fmt.Fprintf(os.Stdout, "Response from `Erc721API.SafeTransferFrom`: %v\n", resp)
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

 **erc721Request** | [**Erc721Request**](Erc721Request.md) |  | 

### Return type

[**AccountControllerResponse**](AccountControllerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetApprovalForAll

> AccountControllerResponse SetApprovalForAll(ctx, name).Authorization(authorization).Erc721Request(erc721Request).Execute()



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
	authorization := "authorization_example" // string | 
	name := "name_example" // string | 
	erc721Request := *openapiclient.NewErc721Request() // Erc721Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Erc721API.SetApprovalForAll(context.Background(), name).Authorization(authorization).Erc721Request(erc721Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Erc721API.SetApprovalForAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetApprovalForAll`: AccountControllerResponse
	fmt.Fprintf(os.Stdout, "Response from `Erc721API.SetApprovalForAll`: %v\n", resp)
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

 **erc721Request** | [**Erc721Request**](Erc721Request.md) |  | 

### Return type

[**AccountControllerResponse**](AccountControllerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Symbol

> AccountControllerResponse Symbol(ctx, name).Authorization(authorization).Erc721Request(erc721Request).Execute()



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
	authorization := "authorization_example" // string | 
	name := "name_example" // string | 
	erc721Request := *openapiclient.NewErc721Request() // Erc721Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Erc721API.Symbol(context.Background(), name).Authorization(authorization).Erc721Request(erc721Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Erc721API.Symbol``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Symbol`: AccountControllerResponse
	fmt.Fprintf(os.Stdout, "Response from `Erc721API.Symbol`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSymbolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **erc721Request** | [**Erc721Request**](Erc721Request.md) |  | 

### Return type

[**AccountControllerResponse**](AccountControllerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenUri

> AccountControllerResponse TokenUri(ctx, name).Authorization(authorization).Erc721Request(erc721Request).Execute()



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
	authorization := "authorization_example" // string | 
	name := "name_example" // string | 
	erc721Request := *openapiclient.NewErc721Request() // Erc721Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Erc721API.TokenUri(context.Background(), name).Authorization(authorization).Erc721Request(erc721Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Erc721API.TokenUri``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokenUri`: AccountControllerResponse
	fmt.Fprintf(os.Stdout, "Response from `Erc721API.TokenUri`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTokenUriRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **erc721Request** | [**Erc721Request**](Erc721Request.md) |  | 

### Return type

[**AccountControllerResponse**](AccountControllerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Transfer

> AccountControllerResponse Transfer(ctx, name).Authorization(authorization).Erc721Request(erc721Request).Execute()



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
	authorization := "authorization_example" // string | 
	name := "name_example" // string | 
	erc721Request := *openapiclient.NewErc721Request() // Erc721Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Erc721API.Transfer(context.Background(), name).Authorization(authorization).Erc721Request(erc721Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Erc721API.Transfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Transfer`: AccountControllerResponse
	fmt.Fprintf(os.Stdout, "Response from `Erc721API.Transfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **erc721Request** | [**Erc721Request**](Erc721Request.md) |  | 

### Return type

[**AccountControllerResponse**](AccountControllerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferFrom

> AccountControllerResponse TransferFrom(ctx, name).Authorization(authorization).Erc721Request(erc721Request).Execute()



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
	authorization := "authorization_example" // string | 
	name := "name_example" // string | 
	erc721Request := *openapiclient.NewErc721Request() // Erc721Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Erc721API.TransferFrom(context.Background(), name).Authorization(authorization).Erc721Request(erc721Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Erc721API.TransferFrom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferFrom`: AccountControllerResponse
	fmt.Fprintf(os.Stdout, "Response from `Erc721API.TransferFrom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferFromRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **erc721Request** | [**Erc721Request**](Erc721Request.md) |  | 

### Return type

[**AccountControllerResponse**](AccountControllerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

