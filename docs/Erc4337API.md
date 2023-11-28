# \Erc4337API

All URIs are relative to *https://vault-api.usemoon.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAddress**](Erc4337API.md#GetAddress) | **Post** /erc4337/{accountName}/address | 
[**SignBroadcastUserOpTx**](Erc4337API.md#SignBroadcastUserOpTx) | **Post** /erc4337/{accountName}/sign-broadcast-userop-tx | 



## GetAddress

> AccountControllerResponse GetAddress(ctx, accountName).Authorization(authorization).InputBody(inputBody).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "authorization_example" // string | 
    accountName := "accountName_example" // string | 
    inputBody := *openapiclient.NewInputBody() // InputBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Erc4337API.GetAddress(context.Background(), accountName).Authorization(authorization).InputBody(inputBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Erc4337API.GetAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAddress`: AccountControllerResponse
    fmt.Fprintf(os.Stdout, "Response from `Erc4337API.GetAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **inputBody** | [**InputBody**](InputBody.md) |  | 

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


## SignBroadcastUserOpTx

> AccountControllerResponse SignBroadcastUserOpTx(ctx, accountName).Authorization(authorization).InputBody(inputBody).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "authorization_example" // string | 
    accountName := "accountName_example" // string | 
    inputBody := *openapiclient.NewInputBody() // InputBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Erc4337API.SignBroadcastUserOpTx(context.Background(), accountName).Authorization(authorization).InputBody(inputBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Erc4337API.SignBroadcastUserOpTx``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SignBroadcastUserOpTx`: AccountControllerResponse
    fmt.Fprintf(os.Stdout, "Response from `Erc4337API.SignBroadcastUserOpTx`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSignBroadcastUserOpTxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **inputBody** | [**InputBody**](InputBody.md) |  | 

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

