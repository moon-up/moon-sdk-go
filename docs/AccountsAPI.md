# \AccountsAPI

All URIs are relative to *https://vault-api.usemoon.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BroadcastTx**](AccountsAPI.md#BroadcastTx) | **Post** /accounts/{accountName}/broadcast-tx | 
[**CreateAccount**](AccountsAPI.md#CreateAccount) | **Post** /accounts | 
[**DeleteAccount**](AccountsAPI.md#DeleteAccount) | **Delete** /accounts/{accountName} | 
[**DeployContract**](AccountsAPI.md#DeployContract) | **Post** /accounts/{accountName}/deploy | 
[**GetAccount**](AccountsAPI.md#GetAccount) | **Get** /accounts/{accountName} | 
[**GetBalance**](AccountsAPI.md#GetBalance) | **Get** /accounts/{accountName}/balance | 
[**GetNonce**](AccountsAPI.md#GetNonce) | **Get** /accounts/{accountName}/nonce | 
[**ListAccounts**](AccountsAPI.md#ListAccounts) | **Get** /accounts | 
[**SignMessage**](AccountsAPI.md#SignMessage) | **Post** /accounts/{accountName}/sign-message | 
[**SignTransaction**](AccountsAPI.md#SignTransaction) | **Post** /accounts/{accountName}/sign-transaction | 
[**SignTypedData**](AccountsAPI.md#SignTypedData) | **Post** /accounts/{accountName}/sign-typed-data | 
[**TransferEth**](AccountsAPI.md#TransferEth) | **Post** /accounts/{accountName}/transfer-eth | 



## BroadcastTx

> AccountControllerResponse BroadcastTx(ctx, accountName).Authorization(authorization).BroadcastInput(broadcastInput).Execute()



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
	broadcastInput := *openapiclient.NewBroadcastInput("ChainId_example", "RawTransaction_example") // BroadcastInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.BroadcastTx(context.Background(), accountName).Authorization(authorization).BroadcastInput(broadcastInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.BroadcastTx``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BroadcastTx`: AccountControllerResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.BroadcastTx`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastTxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **broadcastInput** | [**BroadcastInput**](BroadcastInput.md) |  | 

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


## CreateAccount

> AccountControllerResponse CreateAccount(ctx).Authorization(authorization).CreateAccountInput(createAccountInput).Execute()



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
	createAccountInput := *openapiclient.NewCreateAccountInput() // CreateAccountInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.CreateAccount(context.Background()).Authorization(authorization).CreateAccountInput(createAccountInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.CreateAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccount`: AccountControllerResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.CreateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **createAccountInput** | [**CreateAccountInput**](CreateAccountInput.md) |  | 

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


## DeleteAccount

> AccountControllerResponse DeleteAccount(ctx, accountName).Authorization(authorization).Execute()



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
	resp, r, err := apiClient.AccountsAPI.DeleteAccount(context.Background(), accountName).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.DeleteAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAccount`: AccountControllerResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.DeleteAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountRequest struct via the builder pattern


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


## DeployContract

> AccountControllerResponse DeployContract(ctx, accountName).Authorization(authorization).DeployInput(deployInput).Execute()



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
	deployInput := *openapiclient.NewDeployInput("Abi_example", "Bytecode_example") // DeployInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.DeployContract(context.Background(), accountName).Authorization(authorization).DeployInput(deployInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.DeployContract``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeployContract`: AccountControllerResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.DeployContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeployContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **deployInput** | [**DeployInput**](DeployInput.md) |  | 

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


## GetAccount

> AccountControllerResponse GetAccount(ctx, accountName).Authorization(authorization).Execute()



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
	resp, r, err := apiClient.AccountsAPI.GetAccount(context.Background(), accountName).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccount`: AccountControllerResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


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


## GetBalance

> AccountControllerResponse GetBalance(ctx, accountName).Authorization(authorization).ChainId(chainId).Execute()



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
	accountName := "accountName_example" // string | 
	authorization := "authorization_example" // string | 
	chainId := "chainId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetBalance(context.Background(), accountName).Authorization(authorization).ChainId(chainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBalance`: AccountControllerResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **chainId** | **string** |  | 

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


## GetNonce

> AccountControllerResponse GetNonce(ctx, accountName).Authorization(authorization).Execute()



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
	accountName := "accountName_example" // string | 
	authorization := "authorization_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetNonce(context.Background(), accountName).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetNonce``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNonce`: AccountControllerResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetNonce`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNonceRequest struct via the builder pattern


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


## ListAccounts

> AccountControllerResponse ListAccounts(ctx).Authorization(authorization).Execute()



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
	resp, r, err := apiClient.AccountsAPI.ListAccounts(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ListAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccounts`: AccountControllerResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ListAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccountsRequest struct via the builder pattern


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


## SignMessage

> AccountControllerResponse SignMessage(ctx, accountName).Authorization(authorization).SignMessage(signMessage).Execute()



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
	accountName := "accountName_example" // string | 
	authorization := "authorization_example" // string | 
	signMessage := *openapiclient.NewSignMessage("Data_example") // SignMessage | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.SignMessage(context.Background(), accountName).Authorization(authorization).SignMessage(signMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.SignMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignMessage`: AccountControllerResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.SignMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSignMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **signMessage** | [**SignMessage**](SignMessage.md) |  | 

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


## SignTransaction

> AccountControllerResponse SignTransaction(ctx, accountName).Authorization(authorization).InputBody(inputBody).Execute()



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
	accountName := "accountName_example" // string | 
	authorization := "authorization_example" // string | 
	inputBody := *openapiclient.NewInputBody() // InputBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.SignTransaction(context.Background(), accountName).Authorization(authorization).InputBody(inputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.SignTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignTransaction`: AccountControllerResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.SignTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSignTransactionRequest struct via the builder pattern


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


## SignTypedData

> AccountControllerResponse SignTypedData(ctx, accountName).Authorization(authorization).SignTypedData(signTypedData).Execute()



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
	accountName := "accountName_example" // string | 
	authorization := "authorization_example" // string | 
	signTypedData := *openapiclient.NewSignTypedData("Data_example") // SignTypedData | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.SignTypedData(context.Background(), accountName).Authorization(authorization).SignTypedData(signTypedData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.SignTypedData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignTypedData`: AccountControllerResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.SignTypedData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSignTypedDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **signTypedData** | [**SignTypedData**](SignTypedData.md) |  | 

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


## TransferEth

> AccountControllerResponse TransferEth(ctx, accountName).Authorization(authorization).InputBody(inputBody).Execute()



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
	accountName := "accountName_example" // string | 
	authorization := "authorization_example" // string | 
	inputBody := *openapiclient.NewInputBody() // InputBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.TransferEth(context.Background(), accountName).Authorization(authorization).InputBody(inputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.TransferEth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferEth`: AccountControllerResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.TransferEth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferEthRequest struct via the builder pattern


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

