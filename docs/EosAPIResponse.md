# EosAPIResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Message** | **string** |  | 
**Body** | Pointer to [**InputBody**](InputBody.md) |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**EosTransactionOutput**](EosTransactionOutput.md) |  | [optional] 

## Methods

### NewEosAPIResponse

`func NewEosAPIResponse(success bool, message string, ) *EosAPIResponse`

NewEosAPIResponse instantiates a new EosAPIResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEosAPIResponseWithDefaults

`func NewEosAPIResponseWithDefaults() *EosAPIResponse`

NewEosAPIResponseWithDefaults instantiates a new EosAPIResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *EosAPIResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *EosAPIResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *EosAPIResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetMessage

`func (o *EosAPIResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EosAPIResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EosAPIResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetBody

`func (o *EosAPIResponse) GetBody() InputBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *EosAPIResponse) GetBodyOk() (*InputBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *EosAPIResponse) SetBody(v InputBody)`

SetBody sets Body field to given value.

### HasBody

`func (o *EosAPIResponse) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetAddress

`func (o *EosAPIResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *EosAPIResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *EosAPIResponse) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *EosAPIResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetData

`func (o *EosAPIResponse) GetData() EosTransactionOutput`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EosAPIResponse) GetDataOk() (*EosTransactionOutput, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EosAPIResponse) SetData(v EosTransactionOutput)`

SetData sets Data field to given value.

### HasData

`func (o *EosAPIResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


