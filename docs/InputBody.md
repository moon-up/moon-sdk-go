# InputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **string** |  | [optional] 
**Input** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Nonce** | Pointer to **string** |  | [optional] 
**Gas** | Pointer to **string** |  | [optional] 
**GasPrice** | Pointer to **string** |  | [optional] 
**ChainId** | Pointer to **string** |  | [optional] 
**Encoding** | Pointer to **string** |  | [optional] 
**EOA** | Pointer to **bool** |  | [optional] 
**ContractAddress** | Pointer to **string** |  | [optional] 
**TokenId** | Pointer to **string** |  | [optional] 
**TokenIds** | Pointer to **string** |  | [optional] 
**Approved** | Pointer to **bool** |  | [optional] 
**Broadcast** | Pointer to **bool** |  | [optional] 

## Methods

### NewInputBody

`func NewInputBody() *InputBody`

NewInputBody instantiates a new InputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputBodyWithDefaults

`func NewInputBodyWithDefaults() *InputBody`

NewInputBodyWithDefaults instantiates a new InputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *InputBody) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *InputBody) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *InputBody) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *InputBody) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetData

`func (o *InputBody) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InputBody) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InputBody) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *InputBody) HasData() bool`

HasData returns a boolean if a field has been set.

### GetInput

`func (o *InputBody) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *InputBody) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *InputBody) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *InputBody) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetValue

`func (o *InputBody) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InputBody) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InputBody) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *InputBody) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetNonce

`func (o *InputBody) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *InputBody) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *InputBody) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *InputBody) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetGas

`func (o *InputBody) GetGas() string`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *InputBody) GetGasOk() (*string, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *InputBody) SetGas(v string)`

SetGas sets Gas field to given value.

### HasGas

`func (o *InputBody) HasGas() bool`

HasGas returns a boolean if a field has been set.

### GetGasPrice

`func (o *InputBody) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *InputBody) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *InputBody) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *InputBody) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetChainId

`func (o *InputBody) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *InputBody) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *InputBody) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *InputBody) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetEncoding

`func (o *InputBody) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *InputBody) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *InputBody) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *InputBody) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetEOA

`func (o *InputBody) GetEOA() bool`

GetEOA returns the EOA field if non-nil, zero value otherwise.

### GetEOAOk

`func (o *InputBody) GetEOAOk() (*bool, bool)`

GetEOAOk returns a tuple with the EOA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEOA

`func (o *InputBody) SetEOA(v bool)`

SetEOA sets EOA field to given value.

### HasEOA

`func (o *InputBody) HasEOA() bool`

HasEOA returns a boolean if a field has been set.

### GetContractAddress

`func (o *InputBody) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *InputBody) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *InputBody) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.

### HasContractAddress

`func (o *InputBody) HasContractAddress() bool`

HasContractAddress returns a boolean if a field has been set.

### GetTokenId

`func (o *InputBody) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *InputBody) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *InputBody) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *InputBody) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetTokenIds

`func (o *InputBody) GetTokenIds() string`

GetTokenIds returns the TokenIds field if non-nil, zero value otherwise.

### GetTokenIdsOk

`func (o *InputBody) GetTokenIdsOk() (*string, bool)`

GetTokenIdsOk returns a tuple with the TokenIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIds

`func (o *InputBody) SetTokenIds(v string)`

SetTokenIds sets TokenIds field to given value.

### HasTokenIds

`func (o *InputBody) HasTokenIds() bool`

HasTokenIds returns a boolean if a field has been set.

### GetApproved

`func (o *InputBody) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *InputBody) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *InputBody) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *InputBody) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetBroadcast

`func (o *InputBody) GetBroadcast() bool`

GetBroadcast returns the Broadcast field if non-nil, zero value otherwise.

### GetBroadcastOk

`func (o *InputBody) GetBroadcastOk() (*bool, bool)`

GetBroadcastOk returns a tuple with the Broadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcast

`func (o *InputBody) SetBroadcast(v bool)`

SetBroadcast sets Broadcast field to given value.

### HasBroadcast

`func (o *InputBody) HasBroadcast() bool`

HasBroadcast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


