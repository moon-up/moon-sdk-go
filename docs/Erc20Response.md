# Erc20Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **float64** |  | [optional] 
**ChainId** | Pointer to **float64** |  | [optional] 
**Data** | Pointer to **string** |  | [optional] 
**Gas** | Pointer to **string** |  | [optional] 
**GasPrice** | Pointer to **string** |  | [optional] 
**GasTipCap** | Pointer to **NullableString** |  | [optional] 
**GasFeeCap** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Nonce** | Pointer to **float64** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**To** | Pointer to **NullableString** |  | [optional] 
**BlobGas** | Pointer to **NullableString** |  | [optional] 
**BlobGasFeeCap** | Pointer to **NullableString** |  | [optional] 
**BlobHashes** | Pointer to **[]string** |  | [optional] 
**V** | Pointer to **string** |  | [optional] 
**R** | Pointer to **string** |  | [optional] 
**S** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Decimals** | Pointer to **string** |  | [optional] 
**TotalSupply** | Pointer to **string** |  | [optional] 
**ContractAddress** | Pointer to **string** |  | [optional] 
**BalanceOf** | Pointer to **string** |  | [optional] 
**Allowance** | Pointer to **string** |  | [optional] 

## Methods

### NewErc20Response

`func NewErc20Response() *Erc20Response`

NewErc20Response instantiates a new Erc20Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErc20ResponseWithDefaults

`func NewErc20ResponseWithDefaults() *Erc20Response`

NewErc20ResponseWithDefaults instantiates a new Erc20Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Erc20Response) GetType() float64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Erc20Response) GetTypeOk() (*float64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Erc20Response) SetType(v float64)`

SetType sets Type field to given value.

### HasType

`func (o *Erc20Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetChainId

`func (o *Erc20Response) GetChainId() float64`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *Erc20Response) GetChainIdOk() (*float64, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *Erc20Response) SetChainId(v float64)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *Erc20Response) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetData

`func (o *Erc20Response) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Erc20Response) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Erc20Response) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *Erc20Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetGas

`func (o *Erc20Response) GetGas() string`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *Erc20Response) GetGasOk() (*string, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *Erc20Response) SetGas(v string)`

SetGas sets Gas field to given value.

### HasGas

`func (o *Erc20Response) HasGas() bool`

HasGas returns a boolean if a field has been set.

### GetGasPrice

`func (o *Erc20Response) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *Erc20Response) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *Erc20Response) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *Erc20Response) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetGasTipCap

`func (o *Erc20Response) GetGasTipCap() string`

GetGasTipCap returns the GasTipCap field if non-nil, zero value otherwise.

### GetGasTipCapOk

`func (o *Erc20Response) GetGasTipCapOk() (*string, bool)`

GetGasTipCapOk returns a tuple with the GasTipCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasTipCap

`func (o *Erc20Response) SetGasTipCap(v string)`

SetGasTipCap sets GasTipCap field to given value.

### HasGasTipCap

`func (o *Erc20Response) HasGasTipCap() bool`

HasGasTipCap returns a boolean if a field has been set.

### SetGasTipCapNil

`func (o *Erc20Response) SetGasTipCapNil(b bool)`

 SetGasTipCapNil sets the value for GasTipCap to be an explicit nil

### UnsetGasTipCap
`func (o *Erc20Response) UnsetGasTipCap()`

UnsetGasTipCap ensures that no value is present for GasTipCap, not even an explicit nil
### GetGasFeeCap

`func (o *Erc20Response) GetGasFeeCap() string`

GetGasFeeCap returns the GasFeeCap field if non-nil, zero value otherwise.

### GetGasFeeCapOk

`func (o *Erc20Response) GetGasFeeCapOk() (*string, bool)`

GetGasFeeCapOk returns a tuple with the GasFeeCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasFeeCap

`func (o *Erc20Response) SetGasFeeCap(v string)`

SetGasFeeCap sets GasFeeCap field to given value.

### HasGasFeeCap

`func (o *Erc20Response) HasGasFeeCap() bool`

HasGasFeeCap returns a boolean if a field has been set.

### SetGasFeeCapNil

`func (o *Erc20Response) SetGasFeeCapNil(b bool)`

 SetGasFeeCapNil sets the value for GasFeeCap to be an explicit nil

### UnsetGasFeeCap
`func (o *Erc20Response) UnsetGasFeeCap()`

UnsetGasFeeCap ensures that no value is present for GasFeeCap, not even an explicit nil
### GetValue

`func (o *Erc20Response) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Erc20Response) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Erc20Response) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Erc20Response) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetNonce

`func (o *Erc20Response) GetNonce() float64`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *Erc20Response) GetNonceOk() (*float64, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *Erc20Response) SetNonce(v float64)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *Erc20Response) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetFrom

`func (o *Erc20Response) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Erc20Response) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Erc20Response) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *Erc20Response) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *Erc20Response) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Erc20Response) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Erc20Response) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *Erc20Response) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *Erc20Response) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *Erc20Response) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetBlobGas

`func (o *Erc20Response) GetBlobGas() string`

GetBlobGas returns the BlobGas field if non-nil, zero value otherwise.

### GetBlobGasOk

`func (o *Erc20Response) GetBlobGasOk() (*string, bool)`

GetBlobGasOk returns a tuple with the BlobGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobGas

`func (o *Erc20Response) SetBlobGas(v string)`

SetBlobGas sets BlobGas field to given value.

### HasBlobGas

`func (o *Erc20Response) HasBlobGas() bool`

HasBlobGas returns a boolean if a field has been set.

### SetBlobGasNil

`func (o *Erc20Response) SetBlobGasNil(b bool)`

 SetBlobGasNil sets the value for BlobGas to be an explicit nil

### UnsetBlobGas
`func (o *Erc20Response) UnsetBlobGas()`

UnsetBlobGas ensures that no value is present for BlobGas, not even an explicit nil
### GetBlobGasFeeCap

`func (o *Erc20Response) GetBlobGasFeeCap() string`

GetBlobGasFeeCap returns the BlobGasFeeCap field if non-nil, zero value otherwise.

### GetBlobGasFeeCapOk

`func (o *Erc20Response) GetBlobGasFeeCapOk() (*string, bool)`

GetBlobGasFeeCapOk returns a tuple with the BlobGasFeeCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobGasFeeCap

`func (o *Erc20Response) SetBlobGasFeeCap(v string)`

SetBlobGasFeeCap sets BlobGasFeeCap field to given value.

### HasBlobGasFeeCap

`func (o *Erc20Response) HasBlobGasFeeCap() bool`

HasBlobGasFeeCap returns a boolean if a field has been set.

### SetBlobGasFeeCapNil

`func (o *Erc20Response) SetBlobGasFeeCapNil(b bool)`

 SetBlobGasFeeCapNil sets the value for BlobGasFeeCap to be an explicit nil

### UnsetBlobGasFeeCap
`func (o *Erc20Response) UnsetBlobGasFeeCap()`

UnsetBlobGasFeeCap ensures that no value is present for BlobGasFeeCap, not even an explicit nil
### GetBlobHashes

`func (o *Erc20Response) GetBlobHashes() []string`

GetBlobHashes returns the BlobHashes field if non-nil, zero value otherwise.

### GetBlobHashesOk

`func (o *Erc20Response) GetBlobHashesOk() (*[]string, bool)`

GetBlobHashesOk returns a tuple with the BlobHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobHashes

`func (o *Erc20Response) SetBlobHashes(v []string)`

SetBlobHashes sets BlobHashes field to given value.

### HasBlobHashes

`func (o *Erc20Response) HasBlobHashes() bool`

HasBlobHashes returns a boolean if a field has been set.

### SetBlobHashesNil

`func (o *Erc20Response) SetBlobHashesNil(b bool)`

 SetBlobHashesNil sets the value for BlobHashes to be an explicit nil

### UnsetBlobHashes
`func (o *Erc20Response) UnsetBlobHashes()`

UnsetBlobHashes ensures that no value is present for BlobHashes, not even an explicit nil
### GetV

`func (o *Erc20Response) GetV() string`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *Erc20Response) GetVOk() (*string, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *Erc20Response) SetV(v string)`

SetV sets V field to given value.

### HasV

`func (o *Erc20Response) HasV() bool`

HasV returns a boolean if a field has been set.

### GetR

`func (o *Erc20Response) GetR() string`

GetR returns the R field if non-nil, zero value otherwise.

### GetROk

`func (o *Erc20Response) GetROk() (*string, bool)`

GetROk returns a tuple with the R field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR

`func (o *Erc20Response) SetR(v string)`

SetR sets R field to given value.

### HasR

`func (o *Erc20Response) HasR() bool`

HasR returns a boolean if a field has been set.

### GetS

`func (o *Erc20Response) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *Erc20Response) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *Erc20Response) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *Erc20Response) HasS() bool`

HasS returns a boolean if a field has been set.

### GetName

`func (o *Erc20Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Erc20Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Erc20Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Erc20Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSymbol

`func (o *Erc20Response) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Erc20Response) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Erc20Response) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Erc20Response) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDecimals

`func (o *Erc20Response) GetDecimals() string`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *Erc20Response) GetDecimalsOk() (*string, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *Erc20Response) SetDecimals(v string)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *Erc20Response) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.

### GetTotalSupply

`func (o *Erc20Response) GetTotalSupply() string`

GetTotalSupply returns the TotalSupply field if non-nil, zero value otherwise.

### GetTotalSupplyOk

`func (o *Erc20Response) GetTotalSupplyOk() (*string, bool)`

GetTotalSupplyOk returns a tuple with the TotalSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSupply

`func (o *Erc20Response) SetTotalSupply(v string)`

SetTotalSupply sets TotalSupply field to given value.

### HasTotalSupply

`func (o *Erc20Response) HasTotalSupply() bool`

HasTotalSupply returns a boolean if a field has been set.

### GetContractAddress

`func (o *Erc20Response) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *Erc20Response) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *Erc20Response) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.

### HasContractAddress

`func (o *Erc20Response) HasContractAddress() bool`

HasContractAddress returns a boolean if a field has been set.

### GetBalanceOf

`func (o *Erc20Response) GetBalanceOf() string`

GetBalanceOf returns the BalanceOf field if non-nil, zero value otherwise.

### GetBalanceOfOk

`func (o *Erc20Response) GetBalanceOfOk() (*string, bool)`

GetBalanceOfOk returns a tuple with the BalanceOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceOf

`func (o *Erc20Response) SetBalanceOf(v string)`

SetBalanceOf sets BalanceOf field to given value.

### HasBalanceOf

`func (o *Erc20Response) HasBalanceOf() bool`

HasBalanceOf returns a boolean if a field has been set.

### GetAllowance

`func (o *Erc20Response) GetAllowance() string`

GetAllowance returns the Allowance field if non-nil, zero value otherwise.

### GetAllowanceOk

`func (o *Erc20Response) GetAllowanceOk() (*string, bool)`

GetAllowanceOk returns a tuple with the Allowance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowance

`func (o *Erc20Response) SetAllowance(v string)`

SetAllowance sets Allowance field to given value.

### HasAllowance

`func (o *Erc20Response) HasAllowance() bool`

HasAllowance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


