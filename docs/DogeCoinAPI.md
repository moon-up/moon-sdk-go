# \DogeCoinAPI

All URIs are relative to *https://vault-api.usemoon.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDogeCoinAccount**](DogeCoinAPI.md#CreateDogeCoinAccount) | **Post** /dogecoin | 
[**GetDogeCoinAccount**](DogeCoinAPI.md#GetDogeCoinAccount) | **Get** /dogecoin/{accountName} | 
[**ListDogeCoinAccounts**](DogeCoinAPI.md#ListDogeCoinAccounts) | **Get** /dogecoin | 
[**SignDogeCoinTransaction**](DogeCoinAPI.md#SignDogeCoinTransaction) | **Post** /dogecoin/{accountName}/sign-tx | 



## CreateDogeCoinAccount

> AccountControllerResponse CreateDogeCoinAccount(ctx).Authorization(authorization).DogeCoinInput(dogeCoinInput).Execute()



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
    dogeCoinInput := *openapiclient.NewDogeCoinInput() // DogeCoinInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DogeCoinAPI.CreateDogeCoinAccount(context.Background()).Authorization(authorization).DogeCoinInput(dogeCoinInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DogeCoinAPI.CreateDogeCoinAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDogeCoinAccount`: AccountControllerResponse
    fmt.Fprintf(os.Stdout, "Response from `DogeCoinAPI.CreateDogeCoinAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDogeCoinAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **dogeCoinInput** | [**DogeCoinInput**](DogeCoinInput.md) |  | 

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


## GetDogeCoinAccount

> AccountControllerResponse GetDogeCoinAccount(ctx, accountName).Authorization(authorization).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DogeCoinAPI.GetDogeCoinAccount(context.Background(), accountName).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DogeCoinAPI.GetDogeCoinAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDogeCoinAccount`: AccountControllerResponse
    fmt.Fprintf(os.Stdout, "Response from `DogeCoinAPI.GetDogeCoinAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDogeCoinAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 


### Return type

[**AccountControllerResponse**](AccountControllerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDogeCoinAccounts

> AccountControllerResponse ListDogeCoinAccounts(ctx).Authorization(authorization).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DogeCoinAPI.ListDogeCoinAccounts(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DogeCoinAPI.ListDogeCoinAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDogeCoinAccounts`: AccountControllerResponse
    fmt.Fprintf(os.Stdout, "Response from `DogeCoinAPI.ListDogeCoinAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDogeCoinAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

### Return type

[**AccountControllerResponse**](AccountControllerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignDogeCoinTransaction

> AccountControllerResponse SignDogeCoinTransaction(ctx, accountName).Authorization(authorization).DogeCoinTransactionInput(dogeCoinTransactionInput).Execute()



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
    dogeCoinTransactionInput := *openapiclient.NewDogeCoinTransactionInput() // DogeCoinTransactionInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DogeCoinAPI.SignDogeCoinTransaction(context.Background(), accountName).Authorization(authorization).DogeCoinTransactionInput(dogeCoinTransactionInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DogeCoinAPI.SignDogeCoinTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SignDogeCoinTransaction`: AccountControllerResponse
    fmt.Fprintf(os.Stdout, "Response from `DogeCoinAPI.SignDogeCoinTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSignDogeCoinTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **dogeCoinTransactionInput** | [**DogeCoinTransactionInput**](DogeCoinTransactionInput.md) |  | 

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

