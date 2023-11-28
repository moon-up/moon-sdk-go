# AccountControllerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**AccountControllerResponseData**](AccountControllerResponseData.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountControllerResponse

`func NewAccountControllerResponse() *AccountControllerResponse`

NewAccountControllerResponse instantiates a new AccountControllerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountControllerResponseWithDefaults

`func NewAccountControllerResponseWithDefaults() *AccountControllerResponse`

NewAccountControllerResponseWithDefaults instantiates a new AccountControllerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AccountControllerResponse) GetData() AccountControllerResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AccountControllerResponse) GetDataOk() (*AccountControllerResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AccountControllerResponse) SetData(v AccountControllerResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *AccountControllerResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSuccess

`func (o *AccountControllerResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AccountControllerResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AccountControllerResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AccountControllerResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMessage

`func (o *AccountControllerResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AccountControllerResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AccountControllerResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AccountControllerResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


