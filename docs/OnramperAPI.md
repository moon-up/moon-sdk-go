# \OnramperAPI

All URIs are relative to *https://beta.usemoon.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OnRamperCheckout**](OnramperAPI.md#OnRamperCheckout) | **Post** /onramper/fund/${accountName} | 
[**OnRamperGetQuotesBuy**](OnramperAPI.md#OnRamperGetQuotesBuy) | **Get** /onramper/quotes/buy | 
[**OnRamperGetQuotesSell**](OnramperAPI.md#OnRamperGetQuotesSell) | **Get** /onramper/quotes/sell | 
[**OnRamperGetSupportedAssets**](OnramperAPI.md#OnRamperGetSupportedAssets) | **Get** /onramper/assets | 
[**OnRamperGetSupportedCurrencies**](OnramperAPI.md#OnRamperGetSupportedCurrencies) | **Get** /onramper/currencies | 
[**OnRamperGetSupportedDefaultsAll**](OnramperAPI.md#OnRamperGetSupportedDefaultsAll) | **Get** /onramper/defaults | 
[**OnRamperGetSupportedOnRampsAll**](OnramperAPI.md#OnRamperGetSupportedOnRampsAll) | **Get** /onramper/onramps | 
[**OnRamperGetSupportedPaymentTypes**](OnramperAPI.md#OnRamperGetSupportedPaymentTypes) | **Get** /onramper/payment-types | 
[**OnRamperGetSupportedPaymentTypesFiat**](OnramperAPI.md#OnRamperGetSupportedPaymentTypesFiat) | **Get** /onramper/payment-types/fiat | 



## OnRamperCheckout

> interface{} OnRamperCheckout(ctx, accountName).Authorization(authorization).TransactionInput(transactionInput).Execute()



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
	transactionInput := *openapiclient.NewTransactionInput(*openapiclient.NewTransactionInputSupportedParams(*openapiclient.NewTransactionInputSupportedParamsPartnerData(*openapiclient.NewTransactionInputSupportedParamsPartnerDataRedirectUrl("Success_example")), *openapiclient.NewTransactionInputSupportedParamsTheme(NullableFloat64(123), "CardColor_example", "SecondaryTextColor_example", "PrimaryTextColor_example", "SecondaryColor_example", "PrimaryColor_example", "ThemeName_example", false)), *openapiclient.NewTransactionInputWallet("Address_example"), *openapiclient.NewTransactionInputMetaData("QuoteId_example"), "OriginatingHost_example", "PartnerContext_example", "Uuid_example", "Network_example", "PaymentMethod_example", "Type_example", float64(123), "Destination_example", "Source_example", "Onramp_example") // TransactionInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnramperAPI.OnRamperCheckout(context.Background(), accountName).Authorization(authorization).TransactionInput(transactionInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnramperAPI.OnRamperCheckout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnRamperCheckout`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `OnramperAPI.OnRamperCheckout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOnRamperCheckoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **transactionInput** | [**TransactionInput**](TransactionInput.md) |  | 

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


## OnRamperGetQuotesBuy

> []Quote OnRamperGetQuotesBuy(ctx).Authorization(authorization).Fiat(fiat).Crypto(crypto).Amount(amount).PaymentMethod(paymentMethod).Uuid(uuid).ClientName(clientName).Country(country).Execute()



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
	fiat := "fiat_example" // string | 
	crypto := "crypto_example" // string | 
	amount := float64(1.2) // float64 | 
	paymentMethod := "paymentMethod_example" // string |  (optional) (default to "creditcard")
	uuid := "uuid_example" // string |  (optional) (default to "")
	clientName := "clientName_example" // string |  (optional) (default to "")
	country := "country_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnramperAPI.OnRamperGetQuotesBuy(context.Background()).Authorization(authorization).Fiat(fiat).Crypto(crypto).Amount(amount).PaymentMethod(paymentMethod).Uuid(uuid).ClientName(clientName).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnramperAPI.OnRamperGetQuotesBuy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnRamperGetQuotesBuy`: []Quote
	fmt.Fprintf(os.Stdout, "Response from `OnramperAPI.OnRamperGetQuotesBuy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnRamperGetQuotesBuyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **fiat** | **string** |  | 
 **crypto** | **string** |  | 
 **amount** | **float64** |  | 
 **paymentMethod** | **string** |  | [default to &quot;creditcard&quot;]
 **uuid** | **string** |  | [default to &quot;&quot;]
 **clientName** | **string** |  | [default to &quot;&quot;]
 **country** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]Quote**](Quote.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnRamperGetQuotesSell

> []SellQuote OnRamperGetQuotesSell(ctx).Authorization(authorization).Fiat(fiat).Crypto(crypto).Amount(amount).PaymentMethod(paymentMethod).Uuid(uuid).ClientName(clientName).Country(country).Execute()



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
	fiat := "fiat_example" // string | 
	crypto := "crypto_example" // string | 
	amount := float64(1.2) // float64 | 
	paymentMethod := "paymentMethod_example" // string |  (optional) (default to "creditcard")
	uuid := "uuid_example" // string |  (optional) (default to "")
	clientName := "clientName_example" // string |  (optional) (default to "")
	country := "country_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnramperAPI.OnRamperGetQuotesSell(context.Background()).Authorization(authorization).Fiat(fiat).Crypto(crypto).Amount(amount).PaymentMethod(paymentMethod).Uuid(uuid).ClientName(clientName).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnramperAPI.OnRamperGetQuotesSell``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnRamperGetQuotesSell`: []SellQuote
	fmt.Fprintf(os.Stdout, "Response from `OnramperAPI.OnRamperGetQuotesSell`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnRamperGetQuotesSellRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **fiat** | **string** |  | 
 **crypto** | **string** |  | 
 **amount** | **float64** |  | 
 **paymentMethod** | **string** |  | [default to &quot;creditcard&quot;]
 **uuid** | **string** |  | [default to &quot;&quot;]
 **clientName** | **string** |  | [default to &quot;&quot;]
 **country** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]SellQuote**](SellQuote.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnRamperGetSupportedAssets

> SupportedAssetResponse OnRamperGetSupportedAssets(ctx).Authorization(authorization).Source(source).Country(country).Execute()



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
	source := "source_example" // string | 
	country := "country_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnramperAPI.OnRamperGetSupportedAssets(context.Background()).Authorization(authorization).Source(source).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnramperAPI.OnRamperGetSupportedAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnRamperGetSupportedAssets`: SupportedAssetResponse
	fmt.Fprintf(os.Stdout, "Response from `OnramperAPI.OnRamperGetSupportedAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnRamperGetSupportedAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **source** | **string** |  | 
 **country** | **string** |  | 

### Return type

[**SupportedAssetResponse**](SupportedAssetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnRamperGetSupportedCurrencies

> SupportedCurrenciesResponse OnRamperGetSupportedCurrencies(ctx).Authorization(authorization).Type_(type_).Execute()



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
	type_ := "type__example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnramperAPI.OnRamperGetSupportedCurrencies(context.Background()).Authorization(authorization).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnramperAPI.OnRamperGetSupportedCurrencies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnRamperGetSupportedCurrencies`: SupportedCurrenciesResponse
	fmt.Fprintf(os.Stdout, "Response from `OnramperAPI.OnRamperGetSupportedCurrencies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnRamperGetSupportedCurrenciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **type_** | **string** |  | 

### Return type

[**SupportedCurrenciesResponse**](SupportedCurrenciesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnRamperGetSupportedDefaultsAll

> SupportedDefaultResponse OnRamperGetSupportedDefaultsAll(ctx).Authorization(authorization).Country(country).Type_(type_).Execute()



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
	country := "country_example" // string | 
	type_ := "type__example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnramperAPI.OnRamperGetSupportedDefaultsAll(context.Background()).Authorization(authorization).Country(country).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnramperAPI.OnRamperGetSupportedDefaultsAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnRamperGetSupportedDefaultsAll`: SupportedDefaultResponse
	fmt.Fprintf(os.Stdout, "Response from `OnramperAPI.OnRamperGetSupportedDefaultsAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnRamperGetSupportedDefaultsAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **country** | **string** |  | 
 **type_** | **string** |  | 

### Return type

[**SupportedDefaultResponse**](SupportedDefaultResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnRamperGetSupportedOnRampsAll

> GetSupportedOnRampsResponse OnRamperGetSupportedOnRampsAll(ctx).Authorization(authorization).Execute()



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
	resp, r, err := apiClient.OnramperAPI.OnRamperGetSupportedOnRampsAll(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnramperAPI.OnRamperGetSupportedOnRampsAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnRamperGetSupportedOnRampsAll`: GetSupportedOnRampsResponse
	fmt.Fprintf(os.Stdout, "Response from `OnramperAPI.OnRamperGetSupportedOnRampsAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnRamperGetSupportedOnRampsAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

### Return type

[**GetSupportedOnRampsResponse**](GetSupportedOnRampsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnRamperGetSupportedPaymentTypes

> SupportedPaymentTypesCurrencyResponse OnRamperGetSupportedPaymentTypes(ctx).Authorization(authorization).Fiat(fiat).Country(country).Type_(type_).Execute()



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
	fiat := "fiat_example" // string | 
	country := "country_example" // string | 
	type_ := "type__example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnramperAPI.OnRamperGetSupportedPaymentTypes(context.Background()).Authorization(authorization).Fiat(fiat).Country(country).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnramperAPI.OnRamperGetSupportedPaymentTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnRamperGetSupportedPaymentTypes`: SupportedPaymentTypesCurrencyResponse
	fmt.Fprintf(os.Stdout, "Response from `OnramperAPI.OnRamperGetSupportedPaymentTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnRamperGetSupportedPaymentTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **fiat** | **string** |  | 
 **country** | **string** |  | 
 **type_** | **string** |  | 

### Return type

[**SupportedPaymentTypesCurrencyResponse**](SupportedPaymentTypesCurrencyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnRamperGetSupportedPaymentTypesFiat

> SupportedPaymentTypesCurrencyResponse OnRamperGetSupportedPaymentTypesFiat(ctx).Authorization(authorization).Fiat(fiat).Country(country).Execute()



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
	fiat := "fiat_example" // string | 
	country := "country_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnramperAPI.OnRamperGetSupportedPaymentTypesFiat(context.Background()).Authorization(authorization).Fiat(fiat).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnramperAPI.OnRamperGetSupportedPaymentTypesFiat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnRamperGetSupportedPaymentTypesFiat`: SupportedPaymentTypesCurrencyResponse
	fmt.Fprintf(os.Stdout, "Response from `OnramperAPI.OnRamperGetSupportedPaymentTypesFiat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnRamperGetSupportedPaymentTypesFiatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **fiat** | **string** |  | 
 **country** | **string** |  | 

### Return type

[**SupportedPaymentTypesCurrencyResponse**](SupportedPaymentTypesCurrencyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

