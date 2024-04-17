# \ConveyorFinanceAPI

All URIs are relative to *https://beta.usemoon.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Swap**](ConveyorFinanceAPI.md#Swap) | **Post** /conveyorfinance/{name}/swap | 



## Swap

> ConveyorFinanceControllerResponse Swap(ctx, name).Authorization(authorization).TokenSwapParams(tokenSwapParams).Execute()



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
	tokenSwapParams := *openapiclient.NewTokenSwapParams("TokenIn_example", "TokenOut_example", float64(123), float64(123), "AmountIn_example", "Slippage_example", "Recipient_example", "Referrer_example") // TokenSwapParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConveyorFinanceAPI.Swap(context.Background(), name).Authorization(authorization).TokenSwapParams(tokenSwapParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConveyorFinanceAPI.Swap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Swap`: ConveyorFinanceControllerResponse
	fmt.Fprintf(os.Stdout, "Response from `ConveyorFinanceAPI.Swap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSwapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **tokenSwapParams** | [**TokenSwapParams**](TokenSwapParams.md) |  | 

### Return type

[**ConveyorFinanceControllerResponse**](ConveyorFinanceControllerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

