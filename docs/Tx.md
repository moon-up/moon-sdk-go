# Tx

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

## Methods

### NewTx

`func NewTx() *Tx`

NewTx instantiates a new Tx object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTxWithDefaults

`func NewTxWithDefaults() *Tx`

NewTxWithDefaults instantiates a new Tx object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Tx) GetType() float64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Tx) GetTypeOk() (*float64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Tx) SetType(v float64)`

SetType sets Type field to given value.

### HasType

`func (o *Tx) HasType() bool`

HasType returns a boolean if a field has been set.

### GetChainId

`func (o *Tx) GetChainId() float64`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *Tx) GetChainIdOk() (*float64, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *Tx) SetChainId(v float64)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *Tx) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetData

`func (o *Tx) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Tx) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Tx) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *Tx) HasData() bool`

HasData returns a boolean if a field has been set.

### GetGas

`func (o *Tx) GetGas() string`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *Tx) GetGasOk() (*string, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *Tx) SetGas(v string)`

SetGas sets Gas field to given value.

### HasGas

`func (o *Tx) HasGas() bool`

HasGas returns a boolean if a field has been set.

### GetGasPrice

`func (o *Tx) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *Tx) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *Tx) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *Tx) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetGasTipCap

`func (o *Tx) GetGasTipCap() string`

GetGasTipCap returns the GasTipCap field if non-nil, zero value otherwise.

### GetGasTipCapOk

`func (o *Tx) GetGasTipCapOk() (*string, bool)`

GetGasTipCapOk returns a tuple with the GasTipCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasTipCap

`func (o *Tx) SetGasTipCap(v string)`

SetGasTipCap sets GasTipCap field to given value.

### HasGasTipCap

`func (o *Tx) HasGasTipCap() bool`

HasGasTipCap returns a boolean if a field has been set.

### SetGasTipCapNil

`func (o *Tx) SetGasTipCapNil(b bool)`

 SetGasTipCapNil sets the value for GasTipCap to be an explicit nil

### UnsetGasTipCap
`func (o *Tx) UnsetGasTipCap()`

UnsetGasTipCap ensures that no value is present for GasTipCap, not even an explicit nil
### GetGasFeeCap

`func (o *Tx) GetGasFeeCap() string`

GetGasFeeCap returns the GasFeeCap field if non-nil, zero value otherwise.

### GetGasFeeCapOk

`func (o *Tx) GetGasFeeCapOk() (*string, bool)`

GetGasFeeCapOk returns a tuple with the GasFeeCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasFeeCap

`func (o *Tx) SetGasFeeCap(v string)`

SetGasFeeCap sets GasFeeCap field to given value.

### HasGasFeeCap

`func (o *Tx) HasGasFeeCap() bool`

HasGasFeeCap returns a boolean if a field has been set.

### SetGasFeeCapNil

`func (o *Tx) SetGasFeeCapNil(b bool)`

 SetGasFeeCapNil sets the value for GasFeeCap to be an explicit nil

### UnsetGasFeeCap
`func (o *Tx) UnsetGasFeeCap()`

UnsetGasFeeCap ensures that no value is present for GasFeeCap, not even an explicit nil
### GetValue

`func (o *Tx) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Tx) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Tx) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Tx) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetNonce

`func (o *Tx) GetNonce() float64`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *Tx) GetNonceOk() (*float64, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *Tx) SetNonce(v float64)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *Tx) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetFrom

`func (o *Tx) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Tx) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Tx) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *Tx) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *Tx) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Tx) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Tx) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *Tx) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *Tx) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *Tx) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetBlobGas

`func (o *Tx) GetBlobGas() string`

GetBlobGas returns the BlobGas field if non-nil, zero value otherwise.

### GetBlobGasOk

`func (o *Tx) GetBlobGasOk() (*string, bool)`

GetBlobGasOk returns a tuple with the BlobGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobGas

`func (o *Tx) SetBlobGas(v string)`

SetBlobGas sets BlobGas field to given value.

### HasBlobGas

`func (o *Tx) HasBlobGas() bool`

HasBlobGas returns a boolean if a field has been set.

### SetBlobGasNil

`func (o *Tx) SetBlobGasNil(b bool)`

 SetBlobGasNil sets the value for BlobGas to be an explicit nil

### UnsetBlobGas
`func (o *Tx) UnsetBlobGas()`

UnsetBlobGas ensures that no value is present for BlobGas, not even an explicit nil
### GetBlobGasFeeCap

`func (o *Tx) GetBlobGasFeeCap() string`

GetBlobGasFeeCap returns the BlobGasFeeCap field if non-nil, zero value otherwise.

### GetBlobGasFeeCapOk

`func (o *Tx) GetBlobGasFeeCapOk() (*string, bool)`

GetBlobGasFeeCapOk returns a tuple with the BlobGasFeeCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobGasFeeCap

`func (o *Tx) SetBlobGasFeeCap(v string)`

SetBlobGasFeeCap sets BlobGasFeeCap field to given value.

### HasBlobGasFeeCap

`func (o *Tx) HasBlobGasFeeCap() bool`

HasBlobGasFeeCap returns a boolean if a field has been set.

### SetBlobGasFeeCapNil

`func (o *Tx) SetBlobGasFeeCapNil(b bool)`

 SetBlobGasFeeCapNil sets the value for BlobGasFeeCap to be an explicit nil

### UnsetBlobGasFeeCap
`func (o *Tx) UnsetBlobGasFeeCap()`

UnsetBlobGasFeeCap ensures that no value is present for BlobGasFeeCap, not even an explicit nil
### GetBlobHashes

`func (o *Tx) GetBlobHashes() []string`

GetBlobHashes returns the BlobHashes field if non-nil, zero value otherwise.

### GetBlobHashesOk

`func (o *Tx) GetBlobHashesOk() (*[]string, bool)`

GetBlobHashesOk returns a tuple with the BlobHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobHashes

`func (o *Tx) SetBlobHashes(v []string)`

SetBlobHashes sets BlobHashes field to given value.

### HasBlobHashes

`func (o *Tx) HasBlobHashes() bool`

HasBlobHashes returns a boolean if a field has been set.

### SetBlobHashesNil

`func (o *Tx) SetBlobHashesNil(b bool)`

 SetBlobHashesNil sets the value for BlobHashes to be an explicit nil

### UnsetBlobHashes
`func (o *Tx) UnsetBlobHashes()`

UnsetBlobHashes ensures that no value is present for BlobHashes, not even an explicit nil
### GetV

`func (o *Tx) GetV() string`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *Tx) GetVOk() (*string, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *Tx) SetV(v string)`

SetV sets V field to given value.

### HasV

`func (o *Tx) HasV() bool`

HasV returns a boolean if a field has been set.

### GetR

`func (o *Tx) GetR() string`

GetR returns the R field if non-nil, zero value otherwise.

### GetROk

`func (o *Tx) GetROk() (*string, bool)`

GetROk returns a tuple with the R field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR

`func (o *Tx) SetR(v string)`

SetR sets R field to given value.

### HasR

`func (o *Tx) HasR() bool`

HasR returns a boolean if a field has been set.

### GetS

`func (o *Tx) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *Tx) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *Tx) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *Tx) HasS() bool`

HasS returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


