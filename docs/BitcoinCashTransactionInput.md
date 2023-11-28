# BitcoinCashTransactionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **float64** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**Compress** | Pointer to **bool** |  | [optional] 

## Methods

### NewBitcoinCashTransactionInput

`func NewBitcoinCashTransactionInput() *BitcoinCashTransactionInput`

NewBitcoinCashTransactionInput instantiates a new BitcoinCashTransactionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBitcoinCashTransactionInputWithDefaults

`func NewBitcoinCashTransactionInputWithDefaults() *BitcoinCashTransactionInput`

NewBitcoinCashTransactionInputWithDefaults instantiates a new BitcoinCashTransactionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *BitcoinCashTransactionInput) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *BitcoinCashTransactionInput) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *BitcoinCashTransactionInput) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *BitcoinCashTransactionInput) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetValue

`func (o *BitcoinCashTransactionInput) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BitcoinCashTransactionInput) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BitcoinCashTransactionInput) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *BitcoinCashTransactionInput) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetNetwork

`func (o *BitcoinCashTransactionInput) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *BitcoinCashTransactionInput) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *BitcoinCashTransactionInput) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *BitcoinCashTransactionInput) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCompress

`func (o *BitcoinCashTransactionInput) GetCompress() bool`

GetCompress returns the Compress field if non-nil, zero value otherwise.

### GetCompressOk

`func (o *BitcoinCashTransactionInput) GetCompressOk() (*bool, bool)`

GetCompressOk returns a tuple with the Compress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompress

`func (o *BitcoinCashTransactionInput) SetCompress(v bool)`

SetCompress sets Compress field to given value.

### HasCompress

`func (o *BitcoinCashTransactionInput) HasCompress() bool`

HasCompress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


