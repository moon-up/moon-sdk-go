# AbiInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**Indexed** | Pointer to **bool** |  | [optional] 
**Components** | Pointer to [**[]AbiInput**](AbiInput.md) |  | [optional] 
**InternalType** | Pointer to **string** |  | [optional] 

## Methods

### NewAbiInput

`func NewAbiInput(name string, type_ string, ) *AbiInput`

NewAbiInput instantiates a new AbiInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbiInputWithDefaults

`func NewAbiInputWithDefaults() *AbiInput`

NewAbiInputWithDefaults instantiates a new AbiInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AbiInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AbiInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AbiInput) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AbiInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AbiInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AbiInput) SetType(v string)`

SetType sets Type field to given value.


### GetIndexed

`func (o *AbiInput) GetIndexed() bool`

GetIndexed returns the Indexed field if non-nil, zero value otherwise.

### GetIndexedOk

`func (o *AbiInput) GetIndexedOk() (*bool, bool)`

GetIndexedOk returns a tuple with the Indexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexed

`func (o *AbiInput) SetIndexed(v bool)`

SetIndexed sets Indexed field to given value.

### HasIndexed

`func (o *AbiInput) HasIndexed() bool`

HasIndexed returns a boolean if a field has been set.

### GetComponents

`func (o *AbiInput) GetComponents() []AbiInput`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *AbiInput) GetComponentsOk() (*[]AbiInput, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *AbiInput) SetComponents(v []AbiInput)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *AbiInput) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetInternalType

`func (o *AbiInput) GetInternalType() string`

GetInternalType returns the InternalType field if non-nil, zero value otherwise.

### GetInternalTypeOk

`func (o *AbiInput) GetInternalTypeOk() (*string, bool)`

GetInternalTypeOk returns a tuple with the InternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalType

`func (o *AbiInput) SetInternalType(v string)`

SetInternalType sets InternalType field to given value.

### HasInternalType

`func (o *AbiInput) HasInternalType() bool`

HasInternalType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


