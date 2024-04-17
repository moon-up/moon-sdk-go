# \AaveAPI

All URIs are relative to *https://beta.usemoon.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Borrow**](AaveAPI.md#Borrow) | **Post** /aave/{name}/borrow | 
[**Lend**](AaveAPI.md#Lend) | **Post** /aave/{name}/lend | 
[**Repay**](AaveAPI.md#Repay) | **Post** /aave/{name}/repay | 
[**UserReserveData**](AaveAPI.md#UserReserveData) | **Post** /aave/{name}/user-reserve-data | 



## Borrow

> TransactionAPIResponse Borrow(ctx, name).Authorization(authorization).AaveInput(aaveInput).Execute()



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
	aaveInput := *openapiclient.NewAaveInput() // AaveInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AaveAPI.Borrow(context.Background(), name).Authorization(authorization).AaveInput(aaveInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AaveAPI.Borrow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Borrow`: TransactionAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `AaveAPI.Borrow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBorrowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **aaveInput** | [**AaveInput**](AaveInput.md) |  | 

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


## Lend

> TransactionAPIResponse Lend(ctx, name).Authorization(authorization).AaveInput(aaveInput).Execute()



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
	aaveInput := *openapiclient.NewAaveInput() // AaveInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AaveAPI.Lend(context.Background(), name).Authorization(authorization).AaveInput(aaveInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AaveAPI.Lend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Lend`: TransactionAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `AaveAPI.Lend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **aaveInput** | [**AaveInput**](AaveInput.md) |  | 

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


## Repay

> TransactionAPIResponse Repay(ctx, name).Authorization(authorization).AaveInput(aaveInput).Execute()



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
	aaveInput := *openapiclient.NewAaveInput() // AaveInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AaveAPI.Repay(context.Background(), name).Authorization(authorization).AaveInput(aaveInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AaveAPI.Repay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Repay`: TransactionAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `AaveAPI.Repay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **aaveInput** | [**AaveInput**](AaveInput.md) |  | 

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


## UserReserveData

> AaveReservesAPIResponse UserReserveData(ctx, name).Authorization(authorization).AaveInput(aaveInput).Execute()



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
	aaveInput := *openapiclient.NewAaveInput() // AaveInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AaveAPI.UserReserveData(context.Background(), name).Authorization(authorization).AaveInput(aaveInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AaveAPI.UserReserveData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserReserveData`: AaveReservesAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `AaveAPI.UserReserveData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserReserveDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **aaveInput** | [**AaveInput**](AaveInput.md) |  | 

### Return type

[**AaveReservesAPIResponse**](AaveReservesAPIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

