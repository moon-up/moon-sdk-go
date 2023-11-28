# TokenSwapParams

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
**TokenIn** | **string** |  | 
**TokenOut** | **string** |  | 
**TokenInDecimals** | **float64** |  | 
**TokenOutDecimals** | **float64** |  | 
**AmountIn** | **string** |  | 
**Slippage** | **string** |  | 
**Recipient** | **string** |  | 
**Referrer** | **string** |  | 

## Methods

### NewTokenSwapParams

`func NewTokenSwapParams(tokenIn string, tokenOut string, tokenInDecimals float64, tokenOutDecimals float64, amountIn string, slippage string, recipient string, referrer string, ) *TokenSwapParams`

NewTokenSwapParams instantiates a new TokenSwapParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenSwapParamsWithDefaults

`func NewTokenSwapParamsWithDefaults() *TokenSwapParams`

NewTokenSwapParamsWithDefaults instantiates a new TokenSwapParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *TokenSwapParams) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *TokenSwapParams) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *TokenSwapParams) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *TokenSwapParams) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetData

`func (o *TokenSwapParams) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TokenSwapParams) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TokenSwapParams) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *TokenSwapParams) HasData() bool`

HasData returns a boolean if a field has been set.

### GetInput

`func (o *TokenSwapParams) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *TokenSwapParams) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *TokenSwapParams) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *TokenSwapParams) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetValue

`func (o *TokenSwapParams) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TokenSwapParams) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TokenSwapParams) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TokenSwapParams) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetNonce

`func (o *TokenSwapParams) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *TokenSwapParams) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *TokenSwapParams) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *TokenSwapParams) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetGas

`func (o *TokenSwapParams) GetGas() string`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *TokenSwapParams) GetGasOk() (*string, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *TokenSwapParams) SetGas(v string)`

SetGas sets Gas field to given value.

### HasGas

`func (o *TokenSwapParams) HasGas() bool`

HasGas returns a boolean if a field has been set.

### GetGasPrice

`func (o *TokenSwapParams) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *TokenSwapParams) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *TokenSwapParams) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *TokenSwapParams) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetChainId

`func (o *TokenSwapParams) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *TokenSwapParams) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *TokenSwapParams) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *TokenSwapParams) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetEncoding

`func (o *TokenSwapParams) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *TokenSwapParams) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *TokenSwapParams) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *TokenSwapParams) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetEOA

`func (o *TokenSwapParams) GetEOA() bool`

GetEOA returns the EOA field if non-nil, zero value otherwise.

### GetEOAOk

`func (o *TokenSwapParams) GetEOAOk() (*bool, bool)`

GetEOAOk returns a tuple with the EOA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEOA

`func (o *TokenSwapParams) SetEOA(v bool)`

SetEOA sets EOA field to given value.

### HasEOA

`func (o *TokenSwapParams) HasEOA() bool`

HasEOA returns a boolean if a field has been set.

### GetContractAddress

`func (o *TokenSwapParams) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *TokenSwapParams) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *TokenSwapParams) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.

### HasContractAddress

`func (o *TokenSwapParams) HasContractAddress() bool`

HasContractAddress returns a boolean if a field has been set.

### GetTokenId

`func (o *TokenSwapParams) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TokenSwapParams) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TokenSwapParams) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *TokenSwapParams) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetTokenIds

`func (o *TokenSwapParams) GetTokenIds() string`

GetTokenIds returns the TokenIds field if non-nil, zero value otherwise.

### GetTokenIdsOk

`func (o *TokenSwapParams) GetTokenIdsOk() (*string, bool)`

GetTokenIdsOk returns a tuple with the TokenIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIds

`func (o *TokenSwapParams) SetTokenIds(v string)`

SetTokenIds sets TokenIds field to given value.

### HasTokenIds

`func (o *TokenSwapParams) HasTokenIds() bool`

HasTokenIds returns a boolean if a field has been set.

### GetApproved

`func (o *TokenSwapParams) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *TokenSwapParams) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *TokenSwapParams) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *TokenSwapParams) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetBroadcast

`func (o *TokenSwapParams) GetBroadcast() bool`

GetBroadcast returns the Broadcast field if non-nil, zero value otherwise.

### GetBroadcastOk

`func (o *TokenSwapParams) GetBroadcastOk() (*bool, bool)`

GetBroadcastOk returns a tuple with the Broadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcast

`func (o *TokenSwapParams) SetBroadcast(v bool)`

SetBroadcast sets Broadcast field to given value.

### HasBroadcast

`func (o *TokenSwapParams) HasBroadcast() bool`

HasBroadcast returns a boolean if a field has been set.

### GetTokenIn

`func (o *TokenSwapParams) GetTokenIn() string`

GetTokenIn returns the TokenIn field if non-nil, zero value otherwise.

### GetTokenInOk

`func (o *TokenSwapParams) GetTokenInOk() (*string, bool)`

GetTokenInOk returns a tuple with the TokenIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIn

`func (o *TokenSwapParams) SetTokenIn(v string)`

SetTokenIn sets TokenIn field to given value.


### GetTokenOut

`func (o *TokenSwapParams) GetTokenOut() string`

GetTokenOut returns the TokenOut field if non-nil, zero value otherwise.

### GetTokenOutOk

`func (o *TokenSwapParams) GetTokenOutOk() (*string, bool)`

GetTokenOutOk returns a tuple with the TokenOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenOut

`func (o *TokenSwapParams) SetTokenOut(v string)`

SetTokenOut sets TokenOut field to given value.


### GetTokenInDecimals

`func (o *TokenSwapParams) GetTokenInDecimals() float64`

GetTokenInDecimals returns the TokenInDecimals field if non-nil, zero value otherwise.

### GetTokenInDecimalsOk

`func (o *TokenSwapParams) GetTokenInDecimalsOk() (*float64, bool)`

GetTokenInDecimalsOk returns a tuple with the TokenInDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenInDecimals

`func (o *TokenSwapParams) SetTokenInDecimals(v float64)`

SetTokenInDecimals sets TokenInDecimals field to given value.


### GetTokenOutDecimals

`func (o *TokenSwapParams) GetTokenOutDecimals() float64`

GetTokenOutDecimals returns the TokenOutDecimals field if non-nil, zero value otherwise.

### GetTokenOutDecimalsOk

`func (o *TokenSwapParams) GetTokenOutDecimalsOk() (*float64, bool)`

GetTokenOutDecimalsOk returns a tuple with the TokenOutDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenOutDecimals

`func (o *TokenSwapParams) SetTokenOutDecimals(v float64)`

SetTokenOutDecimals sets TokenOutDecimals field to given value.


### GetAmountIn

`func (o *TokenSwapParams) GetAmountIn() string`

GetAmountIn returns the AmountIn field if non-nil, zero value otherwise.

### GetAmountInOk

`func (o *TokenSwapParams) GetAmountInOk() (*string, bool)`

GetAmountInOk returns a tuple with the AmountIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountIn

`func (o *TokenSwapParams) SetAmountIn(v string)`

SetAmountIn sets AmountIn field to given value.


### GetSlippage

`func (o *TokenSwapParams) GetSlippage() string`

GetSlippage returns the Slippage field if non-nil, zero value otherwise.

### GetSlippageOk

`func (o *TokenSwapParams) GetSlippageOk() (*string, bool)`

GetSlippageOk returns a tuple with the Slippage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlippage

`func (o *TokenSwapParams) SetSlippage(v string)`

SetSlippage sets Slippage field to given value.


### GetRecipient

`func (o *TokenSwapParams) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *TokenSwapParams) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *TokenSwapParams) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetReferrer

`func (o *TokenSwapParams) GetReferrer() string`

GetReferrer returns the Referrer field if non-nil, zero value otherwise.

### GetReferrerOk

`func (o *TokenSwapParams) GetReferrerOk() (*string, bool)`

GetReferrerOk returns a tuple with the Referrer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrer

`func (o *TokenSwapParams) SetReferrer(v string)`

SetReferrer sets Referrer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


