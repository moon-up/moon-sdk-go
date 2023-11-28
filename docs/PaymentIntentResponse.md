# PaymentIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 
**QrCode** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Destination** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **float64** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Success** | **bool** |  | 

## Methods

### NewPaymentIntentResponse

`func NewPaymentIntentResponse(success bool, ) *PaymentIntentResponse`

NewPaymentIntentResponse instantiates a new PaymentIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentIntentResponseWithDefaults

`func NewPaymentIntentResponseWithDefaults() *PaymentIntentResponse`

NewPaymentIntentResponseWithDefaults instantiates a new PaymentIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *PaymentIntentResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PaymentIntentResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PaymentIntentResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PaymentIntentResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetUri

`func (o *PaymentIntentResponse) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *PaymentIntentResponse) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *PaymentIntentResponse) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *PaymentIntentResponse) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetQrCode

`func (o *PaymentIntentResponse) GetQrCode() string`

GetQrCode returns the QrCode field if non-nil, zero value otherwise.

### GetQrCodeOk

`func (o *PaymentIntentResponse) GetQrCodeOk() (*string, bool)`

GetQrCodeOk returns a tuple with the QrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCode

`func (o *PaymentIntentResponse) SetQrCode(v string)`

SetQrCode sets QrCode field to given value.

### HasQrCode

`func (o *PaymentIntentResponse) HasQrCode() bool`

HasQrCode returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentIntentResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentIntentResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentIntentResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDestination

`func (o *PaymentIntentResponse) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *PaymentIntentResponse) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *PaymentIntentResponse) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *PaymentIntentResponse) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetAmount

`func (o *PaymentIntentResponse) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentIntentResponse) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentIntentResponse) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaymentIntentResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetNetwork

`func (o *PaymentIntentResponse) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *PaymentIntentResponse) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *PaymentIntentResponse) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *PaymentIntentResponse) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentIntentResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentIntentResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentIntentResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentIntentResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetId

`func (o *PaymentIntentResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentIntentResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentIntentResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentIntentResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *PaymentIntentResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PaymentIntentResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PaymentIntentResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PaymentIntentResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSuccess

`func (o *PaymentIntentResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *PaymentIntentResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *PaymentIntentResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


