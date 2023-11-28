# Log

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Triggers** | Pointer to [**[]TriggerOutput**](TriggerOutput.md) |  | [optional] 
**LogIndex** | **string** |  | 
**TransactionHash** | **string** |  | 
**Address** | **string** |  | 
**Data** | **string** |  | 
**Topic0** | **NullableString** |  | 
**Topic1** | **NullableString** |  | 
**Topic2** | **NullableString** |  | 
**Topic3** | **NullableString** |  | 

## Methods

### NewLog

`func NewLog(logIndex string, transactionHash string, address string, data string, topic0 NullableString, topic1 NullableString, topic2 NullableString, topic3 NullableString, ) *Log`

NewLog instantiates a new Log object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogWithDefaults

`func NewLogWithDefaults() *Log`

NewLogWithDefaults instantiates a new Log object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggers

`func (o *Log) GetTriggers() []TriggerOutput`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *Log) GetTriggersOk() (*[]TriggerOutput, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *Log) SetTriggers(v []TriggerOutput)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *Log) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetLogIndex

`func (o *Log) GetLogIndex() string`

GetLogIndex returns the LogIndex field if non-nil, zero value otherwise.

### GetLogIndexOk

`func (o *Log) GetLogIndexOk() (*string, bool)`

GetLogIndexOk returns a tuple with the LogIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogIndex

`func (o *Log) SetLogIndex(v string)`

SetLogIndex sets LogIndex field to given value.


### GetTransactionHash

`func (o *Log) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *Log) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *Log) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetAddress

`func (o *Log) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Log) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Log) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetData

`func (o *Log) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Log) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Log) SetData(v string)`

SetData sets Data field to given value.


### GetTopic0

`func (o *Log) GetTopic0() string`

GetTopic0 returns the Topic0 field if non-nil, zero value otherwise.

### GetTopic0Ok

`func (o *Log) GetTopic0Ok() (*string, bool)`

GetTopic0Ok returns a tuple with the Topic0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic0

`func (o *Log) SetTopic0(v string)`

SetTopic0 sets Topic0 field to given value.


### SetTopic0Nil

`func (o *Log) SetTopic0Nil(b bool)`

 SetTopic0Nil sets the value for Topic0 to be an explicit nil

### UnsetTopic0
`func (o *Log) UnsetTopic0()`

UnsetTopic0 ensures that no value is present for Topic0, not even an explicit nil
### GetTopic1

`func (o *Log) GetTopic1() string`

GetTopic1 returns the Topic1 field if non-nil, zero value otherwise.

### GetTopic1Ok

`func (o *Log) GetTopic1Ok() (*string, bool)`

GetTopic1Ok returns a tuple with the Topic1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic1

`func (o *Log) SetTopic1(v string)`

SetTopic1 sets Topic1 field to given value.


### SetTopic1Nil

`func (o *Log) SetTopic1Nil(b bool)`

 SetTopic1Nil sets the value for Topic1 to be an explicit nil

### UnsetTopic1
`func (o *Log) UnsetTopic1()`

UnsetTopic1 ensures that no value is present for Topic1, not even an explicit nil
### GetTopic2

`func (o *Log) GetTopic2() string`

GetTopic2 returns the Topic2 field if non-nil, zero value otherwise.

### GetTopic2Ok

`func (o *Log) GetTopic2Ok() (*string, bool)`

GetTopic2Ok returns a tuple with the Topic2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic2

`func (o *Log) SetTopic2(v string)`

SetTopic2 sets Topic2 field to given value.


### SetTopic2Nil

`func (o *Log) SetTopic2Nil(b bool)`

 SetTopic2Nil sets the value for Topic2 to be an explicit nil

### UnsetTopic2
`func (o *Log) UnsetTopic2()`

UnsetTopic2 ensures that no value is present for Topic2, not even an explicit nil
### GetTopic3

`func (o *Log) GetTopic3() string`

GetTopic3 returns the Topic3 field if non-nil, zero value otherwise.

### GetTopic3Ok

`func (o *Log) GetTopic3Ok() (*string, bool)`

GetTopic3Ok returns a tuple with the Topic3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic3

`func (o *Log) SetTopic3(v string)`

SetTopic3 sets Topic3 field to given value.


### SetTopic3Nil

`func (o *Log) SetTopic3Nil(b bool)`

 SetTopic3Nil sets the value for Topic3 to be an explicit nil

### UnsetTopic3
`func (o *Log) UnsetTopic3()`

UnsetTopic3 ensures that no value is present for Topic3, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


