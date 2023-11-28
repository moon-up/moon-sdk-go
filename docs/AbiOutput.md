# AbiOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**Components** | Pointer to [**[]AbiOutput**](AbiOutput.md) |  | [optional] 
**InternalType** | Pointer to **string** |  | [optional] 

## Methods

### NewAbiOutput

`func NewAbiOutput(name string, type_ string, ) *AbiOutput`

NewAbiOutput instantiates a new AbiOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbiOutputWithDefaults

`func NewAbiOutputWithDefaults() *AbiOutput`

NewAbiOutputWithDefaults instantiates a new AbiOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AbiOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AbiOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AbiOutput) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AbiOutput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AbiOutput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AbiOutput) SetType(v string)`

SetType sets Type field to given value.


### GetComponents

`func (o *AbiOutput) GetComponents() []AbiOutput`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *AbiOutput) GetComponentsOk() (*[]AbiOutput, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *AbiOutput) SetComponents(v []AbiOutput)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *AbiOutput) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetInternalType

`func (o *AbiOutput) GetInternalType() string`

GetInternalType returns the InternalType field if non-nil, zero value otherwise.

### GetInternalTypeOk

`func (o *AbiOutput) GetInternalTypeOk() (*string, bool)`

GetInternalTypeOk returns a tuple with the InternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalType

`func (o *AbiOutput) SetInternalType(v string)`

SetInternalType sets InternalType field to given value.

### HasInternalType

`func (o *AbiOutput) HasInternalType() bool`

HasInternalType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


