# CreatePaymentIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | **map[string]string** |  | 
**Network** | Pointer to **string** |  | [optional] 
**Amount** | **float64** |  | 
**Currency** | Pointer to **string** |  | [optional] 

## Methods

### NewCreatePaymentIntentInput

`func NewCreatePaymentIntentInput(metadata map[string]string, amount float64, ) *CreatePaymentIntentInput`

NewCreatePaymentIntentInput instantiates a new CreatePaymentIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePaymentIntentInputWithDefaults

`func NewCreatePaymentIntentInputWithDefaults() *CreatePaymentIntentInput`

NewCreatePaymentIntentInputWithDefaults instantiates a new CreatePaymentIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *CreatePaymentIntentInput) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreatePaymentIntentInput) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreatePaymentIntentInput) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.


### GetNetwork

`func (o *CreatePaymentIntentInput) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CreatePaymentIntentInput) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CreatePaymentIntentInput) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *CreatePaymentIntentInput) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetAmount

`func (o *CreatePaymentIntentInput) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreatePaymentIntentInput) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreatePaymentIntentInput) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *CreatePaymentIntentInput) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreatePaymentIntentInput) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreatePaymentIntentInput) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreatePaymentIntentInput) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


