/*
moon-vault-api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package moonsdk

import (
	"encoding/json"
)

// checks if the Erc721Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Erc721Response{}

// Erc721Response struct for Erc721Response
type Erc721Response struct {
	Type *float64 `json:"type,omitempty"`
	ChainId *float64 `json:"chain_id,omitempty"`
	Data *string `json:"data,omitempty"`
	Gas *string `json:"gas,omitempty"`
	GasPrice *string `json:"gas_price,omitempty"`
	GasTipCap NullableString `json:"gas_tip_cap,omitempty"`
	GasFeeCap NullableString `json:"gas_fee_cap,omitempty"`
	Value *string `json:"value,omitempty"`
	Nonce *float64 `json:"nonce,omitempty"`
	From *string `json:"from,omitempty"`
	To NullableString `json:"to,omitempty"`
	BlobGas NullableString `json:"blob_gas,omitempty"`
	BlobGasFeeCap NullableString `json:"blob_gas_fee_cap,omitempty"`
	BlobHashes []string `json:"blob_hashes,omitempty"`
	V *string `json:"v,omitempty"`
	R *string `json:"r,omitempty"`
	S *string `json:"s,omitempty"`
	Name *string `json:"name,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	BalanceOf *string `json:"balance_of,omitempty"`
	OwnerOf *string `json:"owner_of,omitempty"`
	TokenUri *string `json:"token_uri,omitempty"`
	ContractAddress *string `json:"contract_address,omitempty"`
	IsApprovedForAll *string `json:"isApprovedForAll,omitempty"`
}

// NewErc721Response instantiates a new Erc721Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErc721Response() *Erc721Response {
	this := Erc721Response{}
	return &this
}

// NewErc721ResponseWithDefaults instantiates a new Erc721Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErc721ResponseWithDefaults() *Erc721Response {
	this := Erc721Response{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Erc721Response) GetType() float64 {
	if o == nil || IsNil(o.Type) {
		var ret float64
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc721Response) GetTypeOk() (*float64, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Erc721Response) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given float64 and assigns it to the Type field.
func (o *Erc721Response) SetType(v float64) {
	o.Type = &v
}

// GetChainId returns the ChainId field value if set, zero value otherwise.
func (o *Erc721Response) GetChainId() float64 {
	if o == nil || IsNil(o.ChainId) {
		var ret float64
		return ret
	}
	return *o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc721Response) GetChainIdOk() (*float64, bool) {
	if o == nil || IsNil(o.ChainId) {
		return nil, false
	}
	return o.ChainId, true
}

// HasChainId returns a boolean if a field has been set.
func (o *Erc721Response) HasChainId() bool {
	if o != nil && !IsNil(o.ChainId) {
		return true
	}

	return false
}

// SetChainId gets a reference to the given float64 and assigns it to the ChainId field.
func (o *Erc721Response) SetChainId(v float64) {
	o.ChainId = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Erc721Response) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc721Response) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Erc721Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *Erc721Response) SetData(v string) {
	o.Data = &v
}

// GetGas returns the Gas field value if set, zero value otherwise.
func (o *Erc721Response) GetGas() string {
	if o == nil || IsNil(o.Gas) {
		var ret string
		return ret
	}
	return *o.Gas
}

// GetGasOk returns a tuple with the Gas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc721Response) GetGasOk() (*string, bool) {
	if o == nil || IsNil(o.Gas) {
		return nil, false
	}
	return o.Gas, true
}

// HasGas returns a boolean if a field has been set.
func (o *Erc721Response) HasGas() bool {
	if o != nil && !IsNil(o.Gas) {
		return true
	}

	return false
}

// SetGas gets a reference to the given string and assigns it to the Gas field.
func (o *Erc721Response) SetGas(v string) {
	o.Gas = &v
}

// GetGasPrice returns the GasPrice field value if set, zero value otherwise.
func (o *Erc721Response) GetGasPrice() string {
	if o == nil || IsNil(o.GasPrice) {
		var ret string
		return ret
	}
	return *o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc721Response) GetGasPriceOk() (*string, bool) {
	if o == nil || IsNil(o.GasPrice) {
		return nil, false
	}
	return o.GasPrice, true
}

// HasGasPrice returns a boolean if a field has been set.
func (o *Erc721Response) HasGasPrice() bool {
	if o != nil && !IsNil(o.GasPrice) {
		return true
	}

	return false
}

// SetGasPrice gets a reference to the given string and assigns it to the GasPrice field.
func (o *Erc721Response) SetGasPrice(v string) {
	o.GasPrice = &v
}

// GetGasTipCap returns the GasTipCap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Erc721Response) GetGasTipCap() string {
	if o == nil || IsNil(o.GasTipCap.Get()) {
		var ret string
		return ret
	}
	return *o.GasTipCap.Get()
}

// GetGasTipCapOk returns a tuple with the GasTipCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Erc721Response) GetGasTipCapOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GasTipCap.Get(), o.GasTipCap.IsSet()
}

// HasGasTipCap returns a boolean if a field has been set.
func (o *Erc721Response) HasGasTipCap() bool {
	if o != nil && o.GasTipCap.IsSet() {
		return true
	}

	return false
}

// SetGasTipCap gets a reference to the given NullableString and assigns it to the GasTipCap field.
func (o *Erc721Response) SetGasTipCap(v string) {
	o.GasTipCap.Set(&v)
}
// SetGasTipCapNil sets the value for GasTipCap to be an explicit nil
func (o *Erc721Response) SetGasTipCapNil() {
	o.GasTipCap.Set(nil)
}

// UnsetGasTipCap ensures that no value is present for GasTipCap, not even an explicit nil
func (o *Erc721Response) UnsetGasTipCap() {
	o.GasTipCap.Unset()
}

// GetGasFeeCap returns the GasFeeCap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Erc721Response) GetGasFeeCap() string {
	if o == nil || IsNil(o.GasFeeCap.Get()) {
		var ret string
		return ret
	}
	return *o.GasFeeCap.Get()
}

// GetGasFeeCapOk returns a tuple with the GasFeeCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Erc721Response) GetGasFeeCapOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GasFeeCap.Get(), o.GasFeeCap.IsSet()
}

// HasGasFeeCap returns a boolean if a field has been set.
func (o *Erc721Response) HasGasFeeCap() bool {
	if o != nil && o.GasFeeCap.IsSet() {
		return true
	}

	return false
}

// SetGasFeeCap gets a reference to the given NullableString and assigns it to the GasFeeCap field.
func (o *Erc721Response) SetGasFeeCap(v string) {
	o.GasFeeCap.Set(&v)
}
// SetGasFeeCapNil sets the value for GasFeeCap to be an explicit nil
func (o *Erc721Response) SetGasFeeCapNil() {
	o.GasFeeCap.Set(nil)
}

// UnsetGasFeeCap ensures that no value is present for GasFeeCap, not even an explicit nil
func (o *Erc721Response) UnsetGasFeeCap() {
	o.GasFeeCap.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Erc721Response) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc721Response) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Erc721Response) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *Erc721Response) SetValue(v string) {
	o.Value = &v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *Erc721Response) GetNonce() float64 {
	if o == nil || IsNil(o.Nonce) {
		var ret float64
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc721Response) GetNonceOk() (*float64, bool) {
	if o == nil || IsNil(o.Nonce) {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *Erc721Response) HasNonce() bool {
	if o != nil && !IsNil(o.Nonce) {
		return true
	}

	return false
}

// SetNonce gets a reference to the given float64 and assigns it to the Nonce field.
func (o *Erc721Response) SetNonce(v float64) {
	o.Nonce = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *Erc721Response) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc721Response) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *Erc721Response) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *Erc721Response) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Erc721Response) GetTo() string {
	if o == nil || IsNil(o.To.Get()) {
		var ret string
		return ret
	}
	return *o.To.Get()
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Erc721Response) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.To.Get(), o.To.IsSet()
}

// HasTo returns a boolean if a field has been set.
func (o *Erc721Response) HasTo() bool {
	if o != nil && o.To.IsSet() {
		return true
	}

	return false
}

// SetTo gets a reference to the given NullableString and assigns it to the To field.
func (o *Erc721Response) SetTo(v string) {
	o.To.Set(&v)
}
// SetToNil sets the value for To to be an explicit nil
func (o *Erc721Response) SetToNil() {
	o.To.Set(nil)
}

// UnsetTo ensures that no value is present for To, not even an explicit nil
func (o *Erc721Response) UnsetTo() {
	o.To.Unset()
}

// GetBlobGas returns the BlobGas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Erc721Response) GetBlobGas() string {
	if o == nil || IsNil(o.BlobGas.Get()) {
		var ret string
		return ret
	}
	return *o.BlobGas.Get()
}

// GetBlobGasOk returns a tuple with the BlobGas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Erc721Response) GetBlobGasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BlobGas.Get(), o.BlobGas.IsSet()
}

// HasBlobGas returns a boolean if a field has been set.
func (o *Erc721Response) HasBlobGas() bool {
	if o != nil && o.BlobGas.IsSet() {
		return true
	}

	return false
}

// SetBlobGas gets a reference to the given NullableString and assigns it to the BlobGas field.
func (o *Erc721Response) SetBlobGas(v string) {
	o.BlobGas.Set(&v)
}
// SetBlobGasNil sets the value for BlobGas to be an explicit nil
func (o *Erc721Response) SetBlobGasNil() {
	o.BlobGas.Set(nil)
}

// UnsetBlobGas ensures that no value is present for BlobGas, not even an explicit nil
func (o *Erc721Response) UnsetBlobGas() {
	o.BlobGas.Unset()
}

// GetBlobGasFeeCap returns the BlobGasFeeCap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Erc721Response) GetBlobGasFeeCap() string {
	if o == nil || IsNil(o.BlobGasFeeCap.Get()) {
		var ret string
		return ret
	}
	return *o.BlobGasFeeCap.Get()
}

// GetBlobGasFeeCapOk returns a tuple with the BlobGasFeeCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Erc721Response) GetBlobGasFeeCapOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BlobGasFeeCap.Get(), o.BlobGasFeeCap.IsSet()
}

// HasBlobGasFeeCap returns a boolean if a field has been set.
func (o *Erc721Response) HasBlobGasFeeCap() bool {
	if o != nil && o.BlobGasFeeCap.IsSet() {
		return true
	}

	return false
}

// SetBlobGasFeeCap gets a reference to the given NullableString and assigns it to the BlobGasFeeCap field.
func (o *Erc721Response) SetBlobGasFeeCap(v string) {
	o.BlobGasFeeCap.Set(&v)
}
// SetBlobGasFeeCapNil sets the value for BlobGasFeeCap to be an explicit nil
func (o *Erc721Response) SetBlobGasFeeCapNil() {
	o.BlobGasFeeCap.Set(nil)
}

// UnsetBlobGasFeeCap ensures that no value is present for BlobGasFeeCap, not even an explicit nil
func (o *Erc721Response) UnsetBlobGasFeeCap() {
	o.BlobGasFeeCap.Unset()
}

// GetBlobHashes returns the BlobHashes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Erc721Response) GetBlobHashes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.BlobHashes
}

// GetBlobHashesOk returns a tuple with the BlobHashes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Erc721Response) GetBlobHashesOk() ([]string, bool) {
	if o == nil || IsNil(o.BlobHashes) {
		return nil, false
	}
	return o.BlobHashes, true
}

// HasBlobHashes returns a boolean if a field has been set.
func (o *Erc721Response) HasBlobHashes() bool {
	if o != nil && IsNil(o.BlobHashes) {
		return true
	}

	return false
}

// SetBlobHashes gets a reference to the given []string and assigns it to the BlobHashes field.
func (o *Erc721Response) SetBlobHashes(v []string) {
	o.BlobHashes = v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *Erc721Response) GetV() string {
	if o == nil || IsNil(o.V) {
		var ret string
		return ret
	}
	return *o.V
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc721Response) GetVOk() (*string, bool) {
	if o == nil || IsNil(o.V) {
		return nil, false
	}
	return o.V, true
}

// HasV returns a boolean if a field has been set.
func (o *Erc721Response) HasV() bool {
	if o != nil && !IsNil(o.V) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *Erc721Response) SetV(v string) {
	o.V = &v
}

// GetR returns the R field value if set, zero value otherwise.
func (o *Erc721Response) GetR() string {
	if o == nil || IsNil(o.R) {
		var ret string
		return ret
	}
	return *o.R
}

// GetROk returns a tuple with the R field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc721Response) GetROk() (*string, bool) {
	if o == nil || IsNil(o.R) {
		return nil, false
	}
	return o.R, true
}

// HasR returns a boolean if a field has been set.
func (o *Erc721Response) HasR() bool {
	if o != nil && !IsNil(o.R) {
		return true
	}

	return false
}

// SetR gets a reference to the given string and assigns it to the R field.
func (o *Erc721Response) SetR(v string) {
	o.R = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *Erc721Response) GetS() string {
	if o == nil || IsNil(o.S) {
		var ret string
		return ret
	}
	return *o.S
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc721Response) GetSOk() (*string, bool) {
	if o == nil || IsNil(o.S) {
		return nil, false
	}
	return o.S, true
}

// HasS returns a boolean if a field has been set.
func (o *Erc721Response) HasS() bool {
	if o != nil && !IsNil(o.S) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *Erc721Response) SetS(v string) {
	o.S = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Erc721Response) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc721Response) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Erc721Response) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Erc721Response) SetName(v string) {
	o.Name = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *Erc721Response) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc721Response) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *Erc721Response) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *Erc721Response) SetSymbol(v string) {
	o.Symbol = &v
}

// GetBalanceOf returns the BalanceOf field value if set, zero value otherwise.
func (o *Erc721Response) GetBalanceOf() string {
	if o == nil || IsNil(o.BalanceOf) {
		var ret string
		return ret
	}
	return *o.BalanceOf
}

// GetBalanceOfOk returns a tuple with the BalanceOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc721Response) GetBalanceOfOk() (*string, bool) {
	if o == nil || IsNil(o.BalanceOf) {
		return nil, false
	}
	return o.BalanceOf, true
}

// HasBalanceOf returns a boolean if a field has been set.
func (o *Erc721Response) HasBalanceOf() bool {
	if o != nil && !IsNil(o.BalanceOf) {
		return true
	}

	return false
}

// SetBalanceOf gets a reference to the given string and assigns it to the BalanceOf field.
func (o *Erc721Response) SetBalanceOf(v string) {
	o.BalanceOf = &v
}

// GetOwnerOf returns the OwnerOf field value if set, zero value otherwise.
func (o *Erc721Response) GetOwnerOf() string {
	if o == nil || IsNil(o.OwnerOf) {
		var ret string
		return ret
	}
	return *o.OwnerOf
}

// GetOwnerOfOk returns a tuple with the OwnerOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc721Response) GetOwnerOfOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerOf) {
		return nil, false
	}
	return o.OwnerOf, true
}

// HasOwnerOf returns a boolean if a field has been set.
func (o *Erc721Response) HasOwnerOf() bool {
	if o != nil && !IsNil(o.OwnerOf) {
		return true
	}

	return false
}

// SetOwnerOf gets a reference to the given string and assigns it to the OwnerOf field.
func (o *Erc721Response) SetOwnerOf(v string) {
	o.OwnerOf = &v
}

// GetTokenUri returns the TokenUri field value if set, zero value otherwise.
func (o *Erc721Response) GetTokenUri() string {
	if o == nil || IsNil(o.TokenUri) {
		var ret string
		return ret
	}
	return *o.TokenUri
}

// GetTokenUriOk returns a tuple with the TokenUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc721Response) GetTokenUriOk() (*string, bool) {
	if o == nil || IsNil(o.TokenUri) {
		return nil, false
	}
	return o.TokenUri, true
}

// HasTokenUri returns a boolean if a field has been set.
func (o *Erc721Response) HasTokenUri() bool {
	if o != nil && !IsNil(o.TokenUri) {
		return true
	}

	return false
}

// SetTokenUri gets a reference to the given string and assigns it to the TokenUri field.
func (o *Erc721Response) SetTokenUri(v string) {
	o.TokenUri = &v
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise.
func (o *Erc721Response) GetContractAddress() string {
	if o == nil || IsNil(o.ContractAddress) {
		var ret string
		return ret
	}
	return *o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc721Response) GetContractAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ContractAddress) {
		return nil, false
	}
	return o.ContractAddress, true
}

// HasContractAddress returns a boolean if a field has been set.
func (o *Erc721Response) HasContractAddress() bool {
	if o != nil && !IsNil(o.ContractAddress) {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given string and assigns it to the ContractAddress field.
func (o *Erc721Response) SetContractAddress(v string) {
	o.ContractAddress = &v
}

// GetIsApprovedForAll returns the IsApprovedForAll field value if set, zero value otherwise.
func (o *Erc721Response) GetIsApprovedForAll() string {
	if o == nil || IsNil(o.IsApprovedForAll) {
		var ret string
		return ret
	}
	return *o.IsApprovedForAll
}

// GetIsApprovedForAllOk returns a tuple with the IsApprovedForAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc721Response) GetIsApprovedForAllOk() (*string, bool) {
	if o == nil || IsNil(o.IsApprovedForAll) {
		return nil, false
	}
	return o.IsApprovedForAll, true
}

// HasIsApprovedForAll returns a boolean if a field has been set.
func (o *Erc721Response) HasIsApprovedForAll() bool {
	if o != nil && !IsNil(o.IsApprovedForAll) {
		return true
	}

	return false
}

// SetIsApprovedForAll gets a reference to the given string and assigns it to the IsApprovedForAll field.
func (o *Erc721Response) SetIsApprovedForAll(v string) {
	o.IsApprovedForAll = &v
}

func (o Erc721Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Erc721Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ChainId) {
		toSerialize["chain_id"] = o.ChainId
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Gas) {
		toSerialize["gas"] = o.Gas
	}
	if !IsNil(o.GasPrice) {
		toSerialize["gas_price"] = o.GasPrice
	}
	if o.GasTipCap.IsSet() {
		toSerialize["gas_tip_cap"] = o.GasTipCap.Get()
	}
	if o.GasFeeCap.IsSet() {
		toSerialize["gas_fee_cap"] = o.GasFeeCap.Get()
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Nonce) {
		toSerialize["nonce"] = o.Nonce
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if o.To.IsSet() {
		toSerialize["to"] = o.To.Get()
	}
	if o.BlobGas.IsSet() {
		toSerialize["blob_gas"] = o.BlobGas.Get()
	}
	if o.BlobGasFeeCap.IsSet() {
		toSerialize["blob_gas_fee_cap"] = o.BlobGasFeeCap.Get()
	}
	if o.BlobHashes != nil {
		toSerialize["blob_hashes"] = o.BlobHashes
	}
	if !IsNil(o.V) {
		toSerialize["v"] = o.V
	}
	if !IsNil(o.R) {
		toSerialize["r"] = o.R
	}
	if !IsNil(o.S) {
		toSerialize["s"] = o.S
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.BalanceOf) {
		toSerialize["balance_of"] = o.BalanceOf
	}
	if !IsNil(o.OwnerOf) {
		toSerialize["owner_of"] = o.OwnerOf
	}
	if !IsNil(o.TokenUri) {
		toSerialize["token_uri"] = o.TokenUri
	}
	if !IsNil(o.ContractAddress) {
		toSerialize["contract_address"] = o.ContractAddress
	}
	if !IsNil(o.IsApprovedForAll) {
		toSerialize["isApprovedForAll"] = o.IsApprovedForAll
	}
	return toSerialize, nil
}

type NullableErc721Response struct {
	value *Erc721Response
	isSet bool
}

func (v NullableErc721Response) Get() *Erc721Response {
	return v.value
}

func (v *NullableErc721Response) Set(val *Erc721Response) {
	v.value = val
	v.isSet = true
}

func (v NullableErc721Response) IsSet() bool {
	return v.isSet
}

func (v *NullableErc721Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErc721Response(val *Erc721Response) *NullableErc721Response {
	return &NullableErc721Response{value: val, isSet: true}
}

func (v NullableErc721Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErc721Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


