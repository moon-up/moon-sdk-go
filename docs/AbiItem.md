# AbiItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Anonymous** | Pointer to **bool** |  | [optional] 
**Constant** | Pointer to **bool** |  | [optional] 
**Inputs** | Pointer to [**[]AbiInput**](AbiInput.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Outputs** | Pointer to [**[]AbiOutput**](AbiOutput.md) |  | [optional] 
**Payable** | Pointer to **bool** |  | [optional] 
**StateMutability** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Gas** | Pointer to **float64** |  | [optional] 

## Methods

### NewAbiItem

`func NewAbiItem(type_ string, ) *AbiItem`

NewAbiItem instantiates a new AbiItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbiItemWithDefaults

`func NewAbiItemWithDefaults() *AbiItem`

NewAbiItemWithDefaults instantiates a new AbiItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnonymous

`func (o *AbiItem) GetAnonymous() bool`

GetAnonymous returns the Anonymous field if non-nil, zero value otherwise.

### GetAnonymousOk

`func (o *AbiItem) GetAnonymousOk() (*bool, bool)`

GetAnonymousOk returns a tuple with the Anonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymous

`func (o *AbiItem) SetAnonymous(v bool)`

SetAnonymous sets Anonymous field to given value.

### HasAnonymous

`func (o *AbiItem) HasAnonymous() bool`

HasAnonymous returns a boolean if a field has been set.

### GetConstant

`func (o *AbiItem) GetConstant() bool`

GetConstant returns the Constant field if non-nil, zero value otherwise.

### GetConstantOk

`func (o *AbiItem) GetConstantOk() (*bool, bool)`

GetConstantOk returns a tuple with the Constant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstant

`func (o *AbiItem) SetConstant(v bool)`

SetConstant sets Constant field to given value.

### HasConstant

`func (o *AbiItem) HasConstant() bool`

HasConstant returns a boolean if a field has been set.

### GetInputs

`func (o *AbiItem) GetInputs() []AbiInput`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *AbiItem) GetInputsOk() (*[]AbiInput, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *AbiItem) SetInputs(v []AbiInput)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *AbiItem) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetName

`func (o *AbiItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AbiItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AbiItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AbiItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputs

`func (o *AbiItem) GetOutputs() []AbiOutput`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *AbiItem) GetOutputsOk() (*[]AbiOutput, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *AbiItem) SetOutputs(v []AbiOutput)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *AbiItem) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetPayable

`func (o *AbiItem) GetPayable() bool`

GetPayable returns the Payable field if non-nil, zero value otherwise.

### GetPayableOk

`func (o *AbiItem) GetPayableOk() (*bool, bool)`

GetPayableOk returns a tuple with the Payable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayable

`func (o *AbiItem) SetPayable(v bool)`

SetPayable sets Payable field to given value.

### HasPayable

`func (o *AbiItem) HasPayable() bool`

HasPayable returns a boolean if a field has been set.

### GetStateMutability

`func (o *AbiItem) GetStateMutability() string`

GetStateMutability returns the StateMutability field if non-nil, zero value otherwise.

### GetStateMutabilityOk

`func (o *AbiItem) GetStateMutabilityOk() (*string, bool)`

GetStateMutabilityOk returns a tuple with the StateMutability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateMutability

`func (o *AbiItem) SetStateMutability(v string)`

SetStateMutability sets StateMutability field to given value.

### HasStateMutability

`func (o *AbiItem) HasStateMutability() bool`

HasStateMutability returns a boolean if a field has been set.

### GetType

`func (o *AbiItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AbiItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AbiItem) SetType(v string)`

SetType sets Type field to given value.


### GetGas

`func (o *AbiItem) GetGas() float64`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *AbiItem) GetGasOk() (*float64, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *AbiItem) SetGas(v float64)`

SetGas sets Gas field to given value.

### HasGas

`func (o *AbiItem) HasGas() bool`

HasGas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


