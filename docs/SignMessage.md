# SignMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Data** | **string** |  | 
**Encoding** | Pointer to **string** |  | [optional] 
**Header** | Pointer to **bool** |  | [optional] 
**Signtype** | Pointer to **bool** |  | [optional] 

## Methods

### NewSignMessage

`func NewSignMessage(data string, ) *SignMessage`

NewSignMessage instantiates a new SignMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignMessageWithDefaults

`func NewSignMessageWithDefaults() *SignMessage`

NewSignMessageWithDefaults instantiates a new SignMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SignMessage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SignMessage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SignMessage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SignMessage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetData

`func (o *SignMessage) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SignMessage) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SignMessage) SetData(v string)`

SetData sets Data field to given value.


### GetEncoding

`func (o *SignMessage) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *SignMessage) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *SignMessage) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *SignMessage) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetHeader

`func (o *SignMessage) GetHeader() bool`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *SignMessage) GetHeaderOk() (*bool, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *SignMessage) SetHeader(v bool)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *SignMessage) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetSigntype

`func (o *SignMessage) GetSigntype() bool`

GetSigntype returns the Signtype field if non-nil, zero value otherwise.

### GetSigntypeOk

`func (o *SignMessage) GetSigntypeOk() (*bool, bool)`

GetSigntypeOk returns a tuple with the Signtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigntype

`func (o *SignMessage) SetSigntype(v bool)`

SetSigntype sets Signtype field to given value.

### HasSigntype

`func (o *SignMessage) HasSigntype() bool`

HasSigntype returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


