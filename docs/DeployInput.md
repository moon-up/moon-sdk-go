# DeployInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | Pointer to **string** |  | [optional] 
**Abi** | **string** |  | 
**Bytecode** | **string** |  | 
**ConstructorArgs** | Pointer to **string** |  | [optional] 

## Methods

### NewDeployInput

`func NewDeployInput(abi string, bytecode string, ) *DeployInput`

NewDeployInput instantiates a new DeployInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployInputWithDefaults

`func NewDeployInputWithDefaults() *DeployInput`

NewDeployInputWithDefaults instantiates a new DeployInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *DeployInput) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *DeployInput) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *DeployInput) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *DeployInput) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetAbi

`func (o *DeployInput) GetAbi() string`

GetAbi returns the Abi field if non-nil, zero value otherwise.

### GetAbiOk

`func (o *DeployInput) GetAbiOk() (*string, bool)`

GetAbiOk returns a tuple with the Abi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbi

`func (o *DeployInput) SetAbi(v string)`

SetAbi sets Abi field to given value.


### GetBytecode

`func (o *DeployInput) GetBytecode() string`

GetBytecode returns the Bytecode field if non-nil, zero value otherwise.

### GetBytecodeOk

`func (o *DeployInput) GetBytecodeOk() (*string, bool)`

GetBytecodeOk returns a tuple with the Bytecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytecode

`func (o *DeployInput) SetBytecode(v string)`

SetBytecode sets Bytecode field to given value.


### GetConstructorArgs

`func (o *DeployInput) GetConstructorArgs() string`

GetConstructorArgs returns the ConstructorArgs field if non-nil, zero value otherwise.

### GetConstructorArgsOk

`func (o *DeployInput) GetConstructorArgsOk() (*string, bool)`

GetConstructorArgsOk returns a tuple with the ConstructorArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstructorArgs

`func (o *DeployInput) SetConstructorArgs(v string)`

SetConstructorArgs sets ConstructorArgs field to given value.

### HasConstructorArgs

`func (o *DeployInput) HasConstructorArgs() bool`

HasConstructorArgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


