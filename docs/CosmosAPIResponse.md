# CosmosAPIResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Message** | **string** |  | 
**Body** | Pointer to [**InputBody**](InputBody.md) |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**CosmosTransactionOutput**](CosmosTransactionOutput.md) |  | [optional] 

## Methods

### NewCosmosAPIResponse

`func NewCosmosAPIResponse(success bool, message string, ) *CosmosAPIResponse`

NewCosmosAPIResponse instantiates a new CosmosAPIResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCosmosAPIResponseWithDefaults

`func NewCosmosAPIResponseWithDefaults() *CosmosAPIResponse`

NewCosmosAPIResponseWithDefaults instantiates a new CosmosAPIResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *CosmosAPIResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CosmosAPIResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CosmosAPIResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetMessage

`func (o *CosmosAPIResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CosmosAPIResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CosmosAPIResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetBody

`func (o *CosmosAPIResponse) GetBody() InputBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CosmosAPIResponse) GetBodyOk() (*InputBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CosmosAPIResponse) SetBody(v InputBody)`

SetBody sets Body field to given value.

### HasBody

`func (o *CosmosAPIResponse) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetAddress

`func (o *CosmosAPIResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CosmosAPIResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CosmosAPIResponse) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CosmosAPIResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetData

`func (o *CosmosAPIResponse) GetData() CosmosTransactionOutput`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CosmosAPIResponse) GetDataOk() (*CosmosTransactionOutput, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CosmosAPIResponse) SetData(v CosmosTransactionOutput)`

SetData sets Data field to given value.

### HasData

`func (o *CosmosAPIResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


