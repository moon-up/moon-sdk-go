# BroadcastInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | **string** |  | 
**RawTransaction** | **string** |  | 

## Methods

### NewBroadcastInput

`func NewBroadcastInput(chainId string, rawTransaction string, ) *BroadcastInput`

NewBroadcastInput instantiates a new BroadcastInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastInputWithDefaults

`func NewBroadcastInputWithDefaults() *BroadcastInput`

NewBroadcastInputWithDefaults instantiates a new BroadcastInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *BroadcastInput) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *BroadcastInput) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *BroadcastInput) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetRawTransaction

`func (o *BroadcastInput) GetRawTransaction() string`

GetRawTransaction returns the RawTransaction field if non-nil, zero value otherwise.

### GetRawTransactionOk

`func (o *BroadcastInput) GetRawTransactionOk() (*string, bool)`

GetRawTransactionOk returns a tuple with the RawTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawTransaction

`func (o *BroadcastInput) SetRawTransaction(v string)`

SetRawTransaction sets RawTransaction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


