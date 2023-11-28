# InternalTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **NullableString** |  | 
**To** | **NullableString** |  | 
**Value** | **NullableString** |  | 
**TransactionHash** | **string** |  | 
**Gas** | **NullableString** |  | 

## Methods

### NewInternalTransaction

`func NewInternalTransaction(from NullableString, to NullableString, value NullableString, transactionHash string, gas NullableString, ) *InternalTransaction`

NewInternalTransaction instantiates a new InternalTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalTransactionWithDefaults

`func NewInternalTransactionWithDefaults() *InternalTransaction`

NewInternalTransactionWithDefaults instantiates a new InternalTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *InternalTransaction) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *InternalTransaction) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *InternalTransaction) SetFrom(v string)`

SetFrom sets From field to given value.


### SetFromNil

`func (o *InternalTransaction) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *InternalTransaction) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetTo

`func (o *InternalTransaction) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *InternalTransaction) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *InternalTransaction) SetTo(v string)`

SetTo sets To field to given value.


### SetToNil

`func (o *InternalTransaction) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *InternalTransaction) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetValue

`func (o *InternalTransaction) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InternalTransaction) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InternalTransaction) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *InternalTransaction) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *InternalTransaction) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetTransactionHash

`func (o *InternalTransaction) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *InternalTransaction) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *InternalTransaction) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetGas

`func (o *InternalTransaction) GetGas() string`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *InternalTransaction) GetGasOk() (*string, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *InternalTransaction) SetGas(v string)`

SetGas sets Gas field to given value.


### SetGasNil

`func (o *InternalTransaction) SetGasNil(b bool)`

 SetGasNil sets the value for Gas to be an explicit nil

### UnsetGas
`func (o *InternalTransaction) UnsetGas()`

UnsetGas ensures that no value is present for Gas, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


