# Erc721Request

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

### NewErc721Request

`func NewErc721Request() *Erc721Request`

NewErc721Request instantiates a new Erc721Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErc721RequestWithDefaults

`func NewErc721RequestWithDefaults() *Erc721Request`

NewErc721RequestWithDefaults instantiates a new Erc721Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *Erc721Request) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Erc721Request) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Erc721Request) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *Erc721Request) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetData

`func (o *Erc721Request) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Erc721Request) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Erc721Request) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *Erc721Request) HasData() bool`

HasData returns a boolean if a field has been set.

### GetInput

`func (o *Erc721Request) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *Erc721Request) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *Erc721Request) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *Erc721Request) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetValue

`func (o *Erc721Request) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Erc721Request) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Erc721Request) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Erc721Request) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetNonce

`func (o *Erc721Request) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *Erc721Request) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *Erc721Request) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *Erc721Request) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetGas

`func (o *Erc721Request) GetGas() string`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *Erc721Request) GetGasOk() (*string, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *Erc721Request) SetGas(v string)`

SetGas sets Gas field to given value.

### HasGas

`func (o *Erc721Request) HasGas() bool`

HasGas returns a boolean if a field has been set.

### GetGasPrice

`func (o *Erc721Request) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *Erc721Request) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *Erc721Request) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *Erc721Request) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetChainId

`func (o *Erc721Request) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *Erc721Request) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *Erc721Request) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *Erc721Request) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetEncoding

`func (o *Erc721Request) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *Erc721Request) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *Erc721Request) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *Erc721Request) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetEOA

`func (o *Erc721Request) GetEOA() bool`

GetEOA returns the EOA field if non-nil, zero value otherwise.

### GetEOAOk

`func (o *Erc721Request) GetEOAOk() (*bool, bool)`

GetEOAOk returns a tuple with the EOA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEOA

`func (o *Erc721Request) SetEOA(v bool)`

SetEOA sets EOA field to given value.

### HasEOA

`func (o *Erc721Request) HasEOA() bool`

HasEOA returns a boolean if a field has been set.

### GetContractAddress

`func (o *Erc721Request) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *Erc721Request) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *Erc721Request) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.

### HasContractAddress

`func (o *Erc721Request) HasContractAddress() bool`

HasContractAddress returns a boolean if a field has been set.

### GetTokenId

`func (o *Erc721Request) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *Erc721Request) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *Erc721Request) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *Erc721Request) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetTokenIds

`func (o *Erc721Request) GetTokenIds() string`

GetTokenIds returns the TokenIds field if non-nil, zero value otherwise.

### GetTokenIdsOk

`func (o *Erc721Request) GetTokenIdsOk() (*string, bool)`

GetTokenIdsOk returns a tuple with the TokenIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIds

`func (o *Erc721Request) SetTokenIds(v string)`

SetTokenIds sets TokenIds field to given value.

### HasTokenIds

`func (o *Erc721Request) HasTokenIds() bool`

HasTokenIds returns a boolean if a field has been set.

### GetApproved

`func (o *Erc721Request) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *Erc721Request) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *Erc721Request) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *Erc721Request) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetBroadcast

`func (o *Erc721Request) GetBroadcast() bool`

GetBroadcast returns the Broadcast field if non-nil, zero value otherwise.

### GetBroadcastOk

`func (o *Erc721Request) GetBroadcastOk() (*bool, bool)`

GetBroadcastOk returns a tuple with the Broadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcast

`func (o *Erc721Request) SetBroadcast(v bool)`

SetBroadcast sets Broadcast field to given value.

### HasBroadcast

`func (o *Erc721Request) HasBroadcast() bool`

HasBroadcast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


