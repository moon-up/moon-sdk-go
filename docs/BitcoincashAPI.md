# \BitcoincashAPI

All URIs are relative to *https://beta.usemoon.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBitcoinCashAccount**](BitcoincashAPI.md#CreateBitcoinCashAccount) | **Post** /bitcoincash | 
[**GetBitcoinCashAccount**](BitcoincashAPI.md#GetBitcoinCashAccount) | **Get** /bitcoincash/{accountName} | 
[**ListBitcoinCashAccounts**](BitcoincashAPI.md#ListBitcoinCashAccounts) | **Get** /bitcoincash | 
[**SignBitcoinCashTransaction**](BitcoincashAPI.md#SignBitcoinCashTransaction) | **Post** /bitcoincash/{accountName}/sign-tx | 



## CreateBitcoinCashAccount

> AccountAPIResponse CreateBitcoinCashAccount(ctx).Authorization(authorization).BitcoinCashInput(bitcoinCashInput).Execute()



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
	bitcoinCashInput := *openapiclient.NewBitcoinCashInput() // BitcoinCashInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BitcoincashAPI.CreateBitcoinCashAccount(context.Background()).Authorization(authorization).BitcoinCashInput(bitcoinCashInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BitcoincashAPI.CreateBitcoinCashAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBitcoinCashAccount`: AccountAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `BitcoincashAPI.CreateBitcoinCashAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBitcoinCashAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **bitcoinCashInput** | [**BitcoinCashInput**](BitcoinCashInput.md) |  | 

### Return type

[**AccountAPIResponse**](AccountAPIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBitcoinCashAccount

> AccountAPIResponse GetBitcoinCashAccount(ctx, accountName).Authorization(authorization).Execute()



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
	accountName := "accountName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BitcoincashAPI.GetBitcoinCashAccount(context.Background(), accountName).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BitcoincashAPI.GetBitcoinCashAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBitcoinCashAccount`: AccountAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `BitcoincashAPI.GetBitcoinCashAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBitcoinCashAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 


### Return type

[**AccountAPIResponse**](AccountAPIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBitcoinCashAccounts

> AccountAPIResponse ListBitcoinCashAccounts(ctx).Authorization(authorization).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BitcoincashAPI.ListBitcoinCashAccounts(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BitcoincashAPI.ListBitcoinCashAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBitcoinCashAccounts`: AccountAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `BitcoincashAPI.ListBitcoinCashAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBitcoinCashAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

### Return type

[**AccountAPIResponse**](AccountAPIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignBitcoinCashTransaction

> BitcoinCashAPIResponse SignBitcoinCashTransaction(ctx, accountName).Authorization(authorization).BitcoinCashTransactionInput(bitcoinCashTransactionInput).Execute()



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
	accountName := "accountName_example" // string | 
	bitcoinCashTransactionInput := *openapiclient.NewBitcoinCashTransactionInput() // BitcoinCashTransactionInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BitcoincashAPI.SignBitcoinCashTransaction(context.Background(), accountName).Authorization(authorization).BitcoinCashTransactionInput(bitcoinCashTransactionInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BitcoincashAPI.SignBitcoinCashTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignBitcoinCashTransaction`: BitcoinCashAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `BitcoincashAPI.SignBitcoinCashTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSignBitcoinCashTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **bitcoinCashTransactionInput** | [**BitcoinCashTransactionInput**](BitcoinCashTransactionInput.md) |  | 

### Return type

[**BitcoinCashAPIResponse**](BitcoinCashAPIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

