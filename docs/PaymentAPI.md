# \PaymentAPI

All URIs are relative to *https://vault-api.usemoon.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MoralisWebhook**](PaymentAPI.md#MoralisWebhook) | **Post** /payment/webhook/{id} | 
[**PaymentCreatePaymentIntent**](PaymentAPI.md#PaymentCreatePaymentIntent) | **Post** /payment | 
[**PaymentDeletePaymentIntent**](PaymentAPI.md#PaymentDeletePaymentIntent) | **Delete** /payment/{id} | 
[**PaymentGetAllPaymentIntents**](PaymentAPI.md#PaymentGetAllPaymentIntents) | **Get** /payment | 
[**PaymentGetAvailableChains**](PaymentAPI.md#PaymentGetAvailableChains) | **Get** /payment/chains | 
[**PaymentGetPaymentIntent**](PaymentAPI.md#PaymentGetPaymentIntent) | **Get** /payment/{id} | 
[**PaymentUpdatePaymentIntent**](PaymentAPI.md#PaymentUpdatePaymentIntent) | **Put** /payment/{id} | 
[**TatumWebhook**](PaymentAPI.md#TatumWebhook) | **Post** /payment/tatum/webhook/{id} | 



## MoralisWebhook

> interface{} MoralisWebhook(ctx, id).IWebhook(iWebhook).Execute()



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
    id := "id_example" // string | 
    iWebhook := *openapiclient.NewIWebhook(*openapiclient.NewBlock("Number_example", "Hash_example", "Timestamp_example"), "ChainId_example", []openapiclient.Log{*openapiclient.NewLog("LogIndex_example", "TransactionHash_example", "Address_example", "Data_example", "Topic0_example", "Topic1_example", "Topic2_example", "Topic3_example")}, []openapiclient.Transaction{*openapiclient.NewTransaction("TransactionHash_example", "SignedTransaction_example")}, []openapiclient.InternalTransaction{*openapiclient.NewInternalTransaction("From_example", "To_example", "Value_example", "TransactionHash_example", "Gas_example")}, []openapiclient.AbiItem{*openapiclient.NewAbiItem("Type_example")}, float64(123), false, "Tag_example", "StreamId_example", []openapiclient.IERC20Transfer{*openapiclient.NewIERC20Transfer("TransactionHash_example", "Contract_example", "LogIndex_example", "From_example", "To_example", "Value_example", "TokenDecimals_example", "TokenName_example", "TokenSymbol_example")}, []openapiclient.IERC20Approval{*openapiclient.NewIERC20Approval("TransactionHash_example", "Contract_example", "LogIndex_example", "Owner_example", "Spender_example", "Value_example", "TokenDecimals_example", "TokenName_example", "TokenSymbol_example")}, []openapiclient.INFTTransfer{*openapiclient.NewINFTTransfer("TransactionHash_example", "Contract_example", "LogIndex_example", "TokenContractType_example", "TokenName_example", "TokenSymbol_example", "Operator_example", "From_example", "To_example", "TokenId_example", "Amount_example")}, []openapiclient.INativeBalance{*openapiclient.NewINativeBalance("Address_example", "Balance_example")}, *openapiclient.NewIOldNFTApproval([]openapiclient.INFTApprovalERC721{*openapiclient.NewINFTApprovalERC721("TransactionHash_example", "Contract_example", "LogIndex_example", "Owner_example", "Approved_example", "TokenId_example", "TokenContractType_example", "TokenName_example", "TokenSymbol_example")}, []openapiclient.INFTApprovalERC1155{*openapiclient.NewINFTApprovalERC1155("TransactionHash_example", "Contract_example", "LogIndex_example", "Account_example", "Operator_example", false, "TokenContractType_example", "TokenName_example", "TokenSymbol_example")}), []openapiclient.INFTApproval{*openapiclient.NewINFTApproval("TransactionHash_example", "Contract_example", "LogIndex_example", "TokenContractType_example", "TokenName_example", "TokenSymbol_example", "Account_example", "Operator_example", false, "TokenId_example")}) // IWebhook | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentAPI.MoralisWebhook(context.Background(), id).IWebhook(iWebhook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.MoralisWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoralisWebhook`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.MoralisWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoralisWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iWebhook** | [**IWebhook**](IWebhook.md) |  | 

### Return type

**interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentCreatePaymentIntent

> PaymentIntentResponse PaymentCreatePaymentIntent(ctx).Authorization(authorization).CreatePaymentIntentInput(createPaymentIntentInput).Execute()



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
    createPaymentIntentInput := *openapiclient.NewCreatePaymentIntentInput(map[string]string{"key": "Inner_example"}, float64(123)) // CreatePaymentIntentInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentAPI.PaymentCreatePaymentIntent(context.Background()).Authorization(authorization).CreatePaymentIntentInput(createPaymentIntentInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.PaymentCreatePaymentIntent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentCreatePaymentIntent`: PaymentIntentResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.PaymentCreatePaymentIntent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentCreatePaymentIntentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **createPaymentIntentInput** | [**CreatePaymentIntentInput**](CreatePaymentIntentInput.md) |  | 

### Return type

[**PaymentIntentResponse**](PaymentIntentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentDeletePaymentIntent

> PaymentIntentResponse PaymentDeletePaymentIntent(ctx, id).Authorization(authorization).Execute()



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentAPI.PaymentDeletePaymentIntent(context.Background(), id).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.PaymentDeletePaymentIntent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentDeletePaymentIntent`: PaymentIntentResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.PaymentDeletePaymentIntent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentDeletePaymentIntentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 


### Return type

[**PaymentIntentResponse**](PaymentIntentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentGetAllPaymentIntents

> []PaymentIntentResponse PaymentGetAllPaymentIntents(ctx).Authorization(authorization).Execute()



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
    resp, r, err := apiClient.PaymentAPI.PaymentGetAllPaymentIntents(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.PaymentGetAllPaymentIntents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentGetAllPaymentIntents`: []PaymentIntentResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.PaymentGetAllPaymentIntents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentGetAllPaymentIntentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

### Return type

[**[]PaymentIntentResponse**](PaymentIntentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentGetAvailableChains

> []string PaymentGetAvailableChains(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentAPI.PaymentGetAvailableChains(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.PaymentGetAvailableChains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentGetAvailableChains`: []string
    fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.PaymentGetAvailableChains`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentGetAvailableChainsRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentGetPaymentIntent

> PaymentIntentResponse PaymentGetPaymentIntent(ctx, id).Authorization(authorization).Execute()



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentAPI.PaymentGetPaymentIntent(context.Background(), id).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.PaymentGetPaymentIntent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentGetPaymentIntent`: PaymentIntentResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.PaymentGetPaymentIntent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentGetPaymentIntentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 


### Return type

[**PaymentIntentResponse**](PaymentIntentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentUpdatePaymentIntent

> PaymentIntentResponse PaymentUpdatePaymentIntent(ctx, id).Authorization(authorization).CreatePaymentIntentInput(createPaymentIntentInput).Execute()



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
    id := "id_example" // string | 
    createPaymentIntentInput := *openapiclient.NewCreatePaymentIntentInput(map[string]string{"key": "Inner_example"}, float64(123)) // CreatePaymentIntentInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentAPI.PaymentUpdatePaymentIntent(context.Background(), id).Authorization(authorization).CreatePaymentIntentInput(createPaymentIntentInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.PaymentUpdatePaymentIntent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentUpdatePaymentIntent`: PaymentIntentResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.PaymentUpdatePaymentIntent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentUpdatePaymentIntentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **createPaymentIntentInput** | [**CreatePaymentIntentInput**](CreatePaymentIntentInput.md) |  | 

### Return type

[**PaymentIntentResponse**](PaymentIntentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TatumWebhook

> interface{} TatumWebhook(ctx, id).Body(body).Execute()



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
    id := "id_example" // string | 
    body := interface{}(987) // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentAPI.TatumWebhook(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.TatumWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TatumWebhook`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.TatumWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTatumWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** |  | 

### Return type

**interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

