# RippleTransactionOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignedTx** | Pointer to **string** |  | [optional] 
**TransactionHash** | Pointer to **string** |  | [optional] 

## Methods

### NewRippleTransactionOutput

`func NewRippleTransactionOutput() *RippleTransactionOutput`

NewRippleTransactionOutput instantiates a new RippleTransactionOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRippleTransactionOutputWithDefaults

`func NewRippleTransactionOutputWithDefaults() *RippleTransactionOutput`

NewRippleTransactionOutputWithDefaults instantiates a new RippleTransactionOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignedTx

`func (o *RippleTransactionOutput) GetSignedTx() string`

GetSignedTx returns the SignedTx field if non-nil, zero value otherwise.

### GetSignedTxOk

`func (o *RippleTransactionOutput) GetSignedTxOk() (*string, bool)`

GetSignedTxOk returns a tuple with the SignedTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedTx

`func (o *RippleTransactionOutput) SetSignedTx(v string)`

SetSignedTx sets SignedTx field to given value.

### HasSignedTx

`func (o *RippleTransactionOutput) HasSignedTx() bool`

HasSignedTx returns a boolean if a field has been set.

### GetTransactionHash

`func (o *RippleTransactionOutput) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *RippleTransactionOutput) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *RippleTransactionOutput) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.

### HasTransactionHash

`func (o *RippleTransactionOutput) HasTransactionHash() bool`

HasTransactionHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


