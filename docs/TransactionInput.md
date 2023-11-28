# TransactionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedParams** | [**TransactionInputSupportedParams**](TransactionInputSupportedParams.md) |  | 
**Wallet** | [**TransactionInputWallet**](TransactionInputWallet.md) |  | 
**MetaData** | [**TransactionInputMetaData**](TransactionInputMetaData.md) |  | 
**OriginatingHost** | **string** |  | 
**PartnerContext** | **string** |  | 
**Uuid** | **string** |  | 
**Network** | **string** |  | 
**PaymentMethod** | **string** |  | 
**Type** | **string** |  | 
**Amount** | **float64** |  | 
**Destination** | **string** |  | 
**Source** | **string** |  | 
**Onramp** | **string** |  | 

## Methods

### NewTransactionInput

`func NewTransactionInput(supportedParams TransactionInputSupportedParams, wallet TransactionInputWallet, metaData TransactionInputMetaData, originatingHost string, partnerContext string, uuid string, network string, paymentMethod string, type_ string, amount float64, destination string, source string, onramp string, ) *TransactionInput`

NewTransactionInput instantiates a new TransactionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionInputWithDefaults

`func NewTransactionInputWithDefaults() *TransactionInput`

NewTransactionInputWithDefaults instantiates a new TransactionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedParams

`func (o *TransactionInput) GetSupportedParams() TransactionInputSupportedParams`

GetSupportedParams returns the SupportedParams field if non-nil, zero value otherwise.

### GetSupportedParamsOk

`func (o *TransactionInput) GetSupportedParamsOk() (*TransactionInputSupportedParams, bool)`

GetSupportedParamsOk returns a tuple with the SupportedParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedParams

`func (o *TransactionInput) SetSupportedParams(v TransactionInputSupportedParams)`

SetSupportedParams sets SupportedParams field to given value.


### GetWallet

`func (o *TransactionInput) GetWallet() TransactionInputWallet`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *TransactionInput) GetWalletOk() (*TransactionInputWallet, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *TransactionInput) SetWallet(v TransactionInputWallet)`

SetWallet sets Wallet field to given value.


### GetMetaData

`func (o *TransactionInput) GetMetaData() TransactionInputMetaData`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *TransactionInput) GetMetaDataOk() (*TransactionInputMetaData, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *TransactionInput) SetMetaData(v TransactionInputMetaData)`

SetMetaData sets MetaData field to given value.


### GetOriginatingHost

`func (o *TransactionInput) GetOriginatingHost() string`

GetOriginatingHost returns the OriginatingHost field if non-nil, zero value otherwise.

### GetOriginatingHostOk

`func (o *TransactionInput) GetOriginatingHostOk() (*string, bool)`

GetOriginatingHostOk returns a tuple with the OriginatingHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingHost

`func (o *TransactionInput) SetOriginatingHost(v string)`

SetOriginatingHost sets OriginatingHost field to given value.


### GetPartnerContext

`func (o *TransactionInput) GetPartnerContext() string`

GetPartnerContext returns the PartnerContext field if non-nil, zero value otherwise.

### GetPartnerContextOk

`func (o *TransactionInput) GetPartnerContextOk() (*string, bool)`

GetPartnerContextOk returns a tuple with the PartnerContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerContext

`func (o *TransactionInput) SetPartnerContext(v string)`

SetPartnerContext sets PartnerContext field to given value.


### GetUuid

`func (o *TransactionInput) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *TransactionInput) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *TransactionInput) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetNetwork

`func (o *TransactionInput) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *TransactionInput) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *TransactionInput) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetPaymentMethod

`func (o *TransactionInput) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *TransactionInput) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *TransactionInput) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetType

`func (o *TransactionInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionInput) SetType(v string)`

SetType sets Type field to given value.


### GetAmount

`func (o *TransactionInput) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionInput) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionInput) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetDestination

`func (o *TransactionInput) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *TransactionInput) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *TransactionInput) SetDestination(v string)`

SetDestination sets Destination field to given value.


### GetSource

`func (o *TransactionInput) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TransactionInput) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TransactionInput) SetSource(v string)`

SetSource sets Source field to given value.


### GetOnramp

`func (o *TransactionInput) GetOnramp() string`

GetOnramp returns the Onramp field if non-nil, zero value otherwise.

### GetOnrampOk

`func (o *TransactionInput) GetOnrampOk() (*string, bool)`

GetOnrampOk returns a tuple with the Onramp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnramp

`func (o *TransactionInput) SetOnramp(v string)`

SetOnramp sets Onramp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


