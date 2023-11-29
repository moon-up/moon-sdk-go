# Erc1155Response

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
**BalanceOf** | Pointer to **string** |  | [optional] 
**BalanceOfBatch** | Pointer to **string** |  | [optional] 

## Methods

### NewErc1155Response

`func NewErc1155Response() *Erc1155Response`

NewErc1155Response instantiates a new Erc1155Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErc1155ResponseWithDefaults

`func NewErc1155ResponseWithDefaults() *Erc1155Response`

NewErc1155ResponseWithDefaults instantiates a new Erc1155Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Erc1155Response) GetType() float64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Erc1155Response) GetTypeOk() (*float64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Erc1155Response) SetType(v float64)`

SetType sets Type field to given value.

### HasType

`func (o *Erc1155Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetChainId

`func (o *Erc1155Response) GetChainId() float64`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *Erc1155Response) GetChainIdOk() (*float64, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *Erc1155Response) SetChainId(v float64)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *Erc1155Response) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetData

`func (o *Erc1155Response) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Erc1155Response) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Erc1155Response) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *Erc1155Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetGas

`func (o *Erc1155Response) GetGas() string`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *Erc1155Response) GetGasOk() (*string, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *Erc1155Response) SetGas(v string)`

SetGas sets Gas field to given value.

### HasGas

`func (o *Erc1155Response) HasGas() bool`

HasGas returns a boolean if a field has been set.

### GetGasPrice

`func (o *Erc1155Response) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *Erc1155Response) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *Erc1155Response) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *Erc1155Response) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetGasTipCap

`func (o *Erc1155Response) GetGasTipCap() string`

GetGasTipCap returns the GasTipCap field if non-nil, zero value otherwise.

### GetGasTipCapOk

`func (o *Erc1155Response) GetGasTipCapOk() (*string, bool)`

GetGasTipCapOk returns a tuple with the GasTipCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasTipCap

`func (o *Erc1155Response) SetGasTipCap(v string)`

SetGasTipCap sets GasTipCap field to given value.

### HasGasTipCap

`func (o *Erc1155Response) HasGasTipCap() bool`

HasGasTipCap returns a boolean if a field has been set.

### SetGasTipCapNil

`func (o *Erc1155Response) SetGasTipCapNil(b bool)`

 SetGasTipCapNil sets the value for GasTipCap to be an explicit nil

### UnsetGasTipCap
`func (o *Erc1155Response) UnsetGasTipCap()`

UnsetGasTipCap ensures that no value is present for GasTipCap, not even an explicit nil
### GetGasFeeCap

`func (o *Erc1155Response) GetGasFeeCap() string`

GetGasFeeCap returns the GasFeeCap field if non-nil, zero value otherwise.

### GetGasFeeCapOk

`func (o *Erc1155Response) GetGasFeeCapOk() (*string, bool)`

GetGasFeeCapOk returns a tuple with the GasFeeCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasFeeCap

`func (o *Erc1155Response) SetGasFeeCap(v string)`

SetGasFeeCap sets GasFeeCap field to given value.

### HasGasFeeCap

`func (o *Erc1155Response) HasGasFeeCap() bool`

HasGasFeeCap returns a boolean if a field has been set.

### SetGasFeeCapNil

`func (o *Erc1155Response) SetGasFeeCapNil(b bool)`

 SetGasFeeCapNil sets the value for GasFeeCap to be an explicit nil

### UnsetGasFeeCap
`func (o *Erc1155Response) UnsetGasFeeCap()`

UnsetGasFeeCap ensures that no value is present for GasFeeCap, not even an explicit nil
### GetValue

`func (o *Erc1155Response) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Erc1155Response) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Erc1155Response) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Erc1155Response) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetNonce

`func (o *Erc1155Response) GetNonce() float64`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *Erc1155Response) GetNonceOk() (*float64, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *Erc1155Response) SetNonce(v float64)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *Erc1155Response) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetFrom

`func (o *Erc1155Response) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Erc1155Response) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Erc1155Response) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *Erc1155Response) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *Erc1155Response) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Erc1155Response) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Erc1155Response) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *Erc1155Response) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *Erc1155Response) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *Erc1155Response) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetBlobGas

`func (o *Erc1155Response) GetBlobGas() string`

GetBlobGas returns the BlobGas field if non-nil, zero value otherwise.

### GetBlobGasOk

`func (o *Erc1155Response) GetBlobGasOk() (*string, bool)`

GetBlobGasOk returns a tuple with the BlobGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobGas

`func (o *Erc1155Response) SetBlobGas(v string)`

SetBlobGas sets BlobGas field to given value.

### HasBlobGas

`func (o *Erc1155Response) HasBlobGas() bool`

HasBlobGas returns a boolean if a field has been set.

### SetBlobGasNil

`func (o *Erc1155Response) SetBlobGasNil(b bool)`

 SetBlobGasNil sets the value for BlobGas to be an explicit nil

### UnsetBlobGas
`func (o *Erc1155Response) UnsetBlobGas()`

UnsetBlobGas ensures that no value is present for BlobGas, not even an explicit nil
### GetBlobGasFeeCap

`func (o *Erc1155Response) GetBlobGasFeeCap() string`

GetBlobGasFeeCap returns the BlobGasFeeCap field if non-nil, zero value otherwise.

### GetBlobGasFeeCapOk

`func (o *Erc1155Response) GetBlobGasFeeCapOk() (*string, bool)`

GetBlobGasFeeCapOk returns a tuple with the BlobGasFeeCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobGasFeeCap

`func (o *Erc1155Response) SetBlobGasFeeCap(v string)`

SetBlobGasFeeCap sets BlobGasFeeCap field to given value.

### HasBlobGasFeeCap

`func (o *Erc1155Response) HasBlobGasFeeCap() bool`

HasBlobGasFeeCap returns a boolean if a field has been set.

### SetBlobGasFeeCapNil

`func (o *Erc1155Response) SetBlobGasFeeCapNil(b bool)`

 SetBlobGasFeeCapNil sets the value for BlobGasFeeCap to be an explicit nil

### UnsetBlobGasFeeCap
`func (o *Erc1155Response) UnsetBlobGasFeeCap()`

UnsetBlobGasFeeCap ensures that no value is present for BlobGasFeeCap, not even an explicit nil
### GetBlobHashes

`func (o *Erc1155Response) GetBlobHashes() []string`

GetBlobHashes returns the BlobHashes field if non-nil, zero value otherwise.

### GetBlobHashesOk

`func (o *Erc1155Response) GetBlobHashesOk() (*[]string, bool)`

GetBlobHashesOk returns a tuple with the BlobHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobHashes

`func (o *Erc1155Response) SetBlobHashes(v []string)`

SetBlobHashes sets BlobHashes field to given value.

### HasBlobHashes

`func (o *Erc1155Response) HasBlobHashes() bool`

HasBlobHashes returns a boolean if a field has been set.

### SetBlobHashesNil

`func (o *Erc1155Response) SetBlobHashesNil(b bool)`

 SetBlobHashesNil sets the value for BlobHashes to be an explicit nil

### UnsetBlobHashes
`func (o *Erc1155Response) UnsetBlobHashes()`

UnsetBlobHashes ensures that no value is present for BlobHashes, not even an explicit nil
### GetV

`func (o *Erc1155Response) GetV() string`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *Erc1155Response) GetVOk() (*string, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *Erc1155Response) SetV(v string)`

SetV sets V field to given value.

### HasV

`func (o *Erc1155Response) HasV() bool`

HasV returns a boolean if a field has been set.

### GetR

`func (o *Erc1155Response) GetR() string`

GetR returns the R field if non-nil, zero value otherwise.

### GetROk

`func (o *Erc1155Response) GetROk() (*string, bool)`

GetROk returns a tuple with the R field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR

`func (o *Erc1155Response) SetR(v string)`

SetR sets R field to given value.

### HasR

`func (o *Erc1155Response) HasR() bool`

HasR returns a boolean if a field has been set.

### GetS

`func (o *Erc1155Response) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *Erc1155Response) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *Erc1155Response) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *Erc1155Response) HasS() bool`

HasS returns a boolean if a field has been set.

### GetBalanceOf

`func (o *Erc1155Response) GetBalanceOf() string`

GetBalanceOf returns the BalanceOf field if non-nil, zero value otherwise.

### GetBalanceOfOk

`func (o *Erc1155Response) GetBalanceOfOk() (*string, bool)`

GetBalanceOfOk returns a tuple with the BalanceOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceOf

`func (o *Erc1155Response) SetBalanceOf(v string)`

SetBalanceOf sets BalanceOf field to given value.

### HasBalanceOf

`func (o *Erc1155Response) HasBalanceOf() bool`

HasBalanceOf returns a boolean if a field has been set.

### GetBalanceOfBatch

`func (o *Erc1155Response) GetBalanceOfBatch() string`

GetBalanceOfBatch returns the BalanceOfBatch field if non-nil, zero value otherwise.

### GetBalanceOfBatchOk

`func (o *Erc1155Response) GetBalanceOfBatchOk() (*string, bool)`

GetBalanceOfBatchOk returns a tuple with the BalanceOfBatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceOfBatch

`func (o *Erc1155Response) SetBalanceOfBatch(v string)`

SetBalanceOfBatch sets BalanceOfBatch field to given value.

### HasBalanceOfBatch

`func (o *Erc1155Response) HasBalanceOfBatch() bool`

HasBalanceOfBatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


