# EosTransactionOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignedTx** | Pointer to **string** |  | [optional] 
**TransactionHash** | Pointer to **string** |  | [optional] 

## Methods

### NewEosTransactionOutput

`func NewEosTransactionOutput() *EosTransactionOutput`

NewEosTransactionOutput instantiates a new EosTransactionOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEosTransactionOutputWithDefaults

`func NewEosTransactionOutputWithDefaults() *EosTransactionOutput`

NewEosTransactionOutputWithDefaults instantiates a new EosTransactionOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignedTx

`func (o *EosTransactionOutput) GetSignedTx() string`

GetSignedTx returns the SignedTx field if non-nil, zero value otherwise.

### GetSignedTxOk

`func (o *EosTransactionOutput) GetSignedTxOk() (*string, bool)`

GetSignedTxOk returns a tuple with the SignedTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedTx

`func (o *EosTransactionOutput) SetSignedTx(v string)`

SetSignedTx sets SignedTx field to given value.

### HasSignedTx

`func (o *EosTransactionOutput) HasSignedTx() bool`

HasSignedTx returns a boolean if a field has been set.

### GetTransactionHash

`func (o *EosTransactionOutput) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *EosTransactionOutput) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *EosTransactionOutput) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.

### HasTransactionHash

`func (o *EosTransactionOutput) HasTransactionHash() bool`

HasTransactionHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


