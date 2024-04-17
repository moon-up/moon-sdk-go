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

// checks if the Tx type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tx{}

// Tx struct for Tx
type Tx struct {
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
}

// NewTx instantiates a new Tx object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTx() *Tx {
	this := Tx{}
	return &this
}

// NewTxWithDefaults instantiates a new Tx object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTxWithDefaults() *Tx {
	this := Tx{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Tx) GetType() float64 {
	if o == nil || IsNil(o.Type) {
		var ret float64
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tx) GetTypeOk() (*float64, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Tx) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given float64 and assigns it to the Type field.
func (o *Tx) SetType(v float64) {
	o.Type = &v
}

// GetChainId returns the ChainId field value if set, zero value otherwise.
func (o *Tx) GetChainId() float64 {
	if o == nil || IsNil(o.ChainId) {
		var ret float64
		return ret
	}
	return *o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tx) GetChainIdOk() (*float64, bool) {
	if o == nil || IsNil(o.ChainId) {
		return nil, false
	}
	return o.ChainId, true
}

// HasChainId returns a boolean if a field has been set.
func (o *Tx) HasChainId() bool {
	if o != nil && !IsNil(o.ChainId) {
		return true
	}

	return false
}

// SetChainId gets a reference to the given float64 and assigns it to the ChainId field.
func (o *Tx) SetChainId(v float64) {
	o.ChainId = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Tx) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tx) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Tx) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *Tx) SetData(v string) {
	o.Data = &v
}

// GetGas returns the Gas field value if set, zero value otherwise.
func (o *Tx) GetGas() string {
	if o == nil || IsNil(o.Gas) {
		var ret string
		return ret
	}
	return *o.Gas
}

// GetGasOk returns a tuple with the Gas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tx) GetGasOk() (*string, bool) {
	if o == nil || IsNil(o.Gas) {
		return nil, false
	}
	return o.Gas, true
}

// HasGas returns a boolean if a field has been set.
func (o *Tx) HasGas() bool {
	if o != nil && !IsNil(o.Gas) {
		return true
	}

	return false
}

// SetGas gets a reference to the given string and assigns it to the Gas field.
func (o *Tx) SetGas(v string) {
	o.Gas = &v
}

// GetGasPrice returns the GasPrice field value if set, zero value otherwise.
func (o *Tx) GetGasPrice() string {
	if o == nil || IsNil(o.GasPrice) {
		var ret string
		return ret
	}
	return *o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tx) GetGasPriceOk() (*string, bool) {
	if o == nil || IsNil(o.GasPrice) {
		return nil, false
	}
	return o.GasPrice, true
}

// HasGasPrice returns a boolean if a field has been set.
func (o *Tx) HasGasPrice() bool {
	if o != nil && !IsNil(o.GasPrice) {
		return true
	}

	return false
}

// SetGasPrice gets a reference to the given string and assigns it to the GasPrice field.
func (o *Tx) SetGasPrice(v string) {
	o.GasPrice = &v
}

// GetGasTipCap returns the GasTipCap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tx) GetGasTipCap() string {
	if o == nil || IsNil(o.GasTipCap.Get()) {
		var ret string
		return ret
	}
	return *o.GasTipCap.Get()
}

// GetGasTipCapOk returns a tuple with the GasTipCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tx) GetGasTipCapOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GasTipCap.Get(), o.GasTipCap.IsSet()
}

// HasGasTipCap returns a boolean if a field has been set.
func (o *Tx) HasGasTipCap() bool {
	if o != nil && o.GasTipCap.IsSet() {
		return true
	}

	return false
}

// SetGasTipCap gets a reference to the given NullableString and assigns it to the GasTipCap field.
func (o *Tx) SetGasTipCap(v string) {
	o.GasTipCap.Set(&v)
}
// SetGasTipCapNil sets the value for GasTipCap to be an explicit nil
func (o *Tx) SetGasTipCapNil() {
	o.GasTipCap.Set(nil)
}

// UnsetGasTipCap ensures that no value is present for GasTipCap, not even an explicit nil
func (o *Tx) UnsetGasTipCap() {
	o.GasTipCap.Unset()
}

// GetGasFeeCap returns the GasFeeCap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tx) GetGasFeeCap() string {
	if o == nil || IsNil(o.GasFeeCap.Get()) {
		var ret string
		return ret
	}
	return *o.GasFeeCap.Get()
}

// GetGasFeeCapOk returns a tuple with the GasFeeCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tx) GetGasFeeCapOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GasFeeCap.Get(), o.GasFeeCap.IsSet()
}

// HasGasFeeCap returns a boolean if a field has been set.
func (o *Tx) HasGasFeeCap() bool {
	if o != nil && o.GasFeeCap.IsSet() {
		return true
	}

	return false
}

// SetGasFeeCap gets a reference to the given NullableString and assigns it to the GasFeeCap field.
func (o *Tx) SetGasFeeCap(v string) {
	o.GasFeeCap.Set(&v)
}
// SetGasFeeCapNil sets the value for GasFeeCap to be an explicit nil
func (o *Tx) SetGasFeeCapNil() {
	o.GasFeeCap.Set(nil)
}

// UnsetGasFeeCap ensures that no value is present for GasFeeCap, not even an explicit nil
func (o *Tx) UnsetGasFeeCap() {
	o.GasFeeCap.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Tx) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tx) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Tx) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *Tx) SetValue(v string) {
	o.Value = &v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *Tx) GetNonce() float64 {
	if o == nil || IsNil(o.Nonce) {
		var ret float64
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tx) GetNonceOk() (*float64, bool) {
	if o == nil || IsNil(o.Nonce) {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *Tx) HasNonce() bool {
	if o != nil && !IsNil(o.Nonce) {
		return true
	}

	return false
}

// SetNonce gets a reference to the given float64 and assigns it to the Nonce field.
func (o *Tx) SetNonce(v float64) {
	o.Nonce = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *Tx) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tx) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *Tx) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *Tx) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tx) GetTo() string {
	if o == nil || IsNil(o.To.Get()) {
		var ret string
		return ret
	}
	return *o.To.Get()
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tx) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.To.Get(), o.To.IsSet()
}

// HasTo returns a boolean if a field has been set.
func (o *Tx) HasTo() bool {
	if o != nil && o.To.IsSet() {
		return true
	}

	return false
}

// SetTo gets a reference to the given NullableString and assigns it to the To field.
func (o *Tx) SetTo(v string) {
	o.To.Set(&v)
}
// SetToNil sets the value for To to be an explicit nil
func (o *Tx) SetToNil() {
	o.To.Set(nil)
}

// UnsetTo ensures that no value is present for To, not even an explicit nil
func (o *Tx) UnsetTo() {
	o.To.Unset()
}

// GetBlobGas returns the BlobGas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tx) GetBlobGas() string {
	if o == nil || IsNil(o.BlobGas.Get()) {
		var ret string
		return ret
	}
	return *o.BlobGas.Get()
}

// GetBlobGasOk returns a tuple with the BlobGas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tx) GetBlobGasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BlobGas.Get(), o.BlobGas.IsSet()
}

// HasBlobGas returns a boolean if a field has been set.
func (o *Tx) HasBlobGas() bool {
	if o != nil && o.BlobGas.IsSet() {
		return true
	}

	return false
}

// SetBlobGas gets a reference to the given NullableString and assigns it to the BlobGas field.
func (o *Tx) SetBlobGas(v string) {
	o.BlobGas.Set(&v)
}
// SetBlobGasNil sets the value for BlobGas to be an explicit nil
func (o *Tx) SetBlobGasNil() {
	o.BlobGas.Set(nil)
}

// UnsetBlobGas ensures that no value is present for BlobGas, not even an explicit nil
func (o *Tx) UnsetBlobGas() {
	o.BlobGas.Unset()
}

// GetBlobGasFeeCap returns the BlobGasFeeCap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tx) GetBlobGasFeeCap() string {
	if o == nil || IsNil(o.BlobGasFeeCap.Get()) {
		var ret string
		return ret
	}
	return *o.BlobGasFeeCap.Get()
}

// GetBlobGasFeeCapOk returns a tuple with the BlobGasFeeCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tx) GetBlobGasFeeCapOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BlobGasFeeCap.Get(), o.BlobGasFeeCap.IsSet()
}

// HasBlobGasFeeCap returns a boolean if a field has been set.
func (o *Tx) HasBlobGasFeeCap() bool {
	if o != nil && o.BlobGasFeeCap.IsSet() {
		return true
	}

	return false
}

// SetBlobGasFeeCap gets a reference to the given NullableString and assigns it to the BlobGasFeeCap field.
func (o *Tx) SetBlobGasFeeCap(v string) {
	o.BlobGasFeeCap.Set(&v)
}
// SetBlobGasFeeCapNil sets the value for BlobGasFeeCap to be an explicit nil
func (o *Tx) SetBlobGasFeeCapNil() {
	o.BlobGasFeeCap.Set(nil)
}

// UnsetBlobGasFeeCap ensures that no value is present for BlobGasFeeCap, not even an explicit nil
func (o *Tx) UnsetBlobGasFeeCap() {
	o.BlobGasFeeCap.Unset()
}

// GetBlobHashes returns the BlobHashes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tx) GetBlobHashes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.BlobHashes
}

// GetBlobHashesOk returns a tuple with the BlobHashes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tx) GetBlobHashesOk() ([]string, bool) {
	if o == nil || IsNil(o.BlobHashes) {
		return nil, false
	}
	return o.BlobHashes, true
}

// HasBlobHashes returns a boolean if a field has been set.
func (o *Tx) HasBlobHashes() bool {
	if o != nil && !IsNil(o.BlobHashes) {
		return true
	}

	return false
}

// SetBlobHashes gets a reference to the given []string and assigns it to the BlobHashes field.
func (o *Tx) SetBlobHashes(v []string) {
	o.BlobHashes = v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *Tx) GetV() string {
	if o == nil || IsNil(o.V) {
		var ret string
		return ret
	}
	return *o.V
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tx) GetVOk() (*string, bool) {
	if o == nil || IsNil(o.V) {
		return nil, false
	}
	return o.V, true
}

// HasV returns a boolean if a field has been set.
func (o *Tx) HasV() bool {
	if o != nil && !IsNil(o.V) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *Tx) SetV(v string) {
	o.V = &v
}

// GetR returns the R field value if set, zero value otherwise.
func (o *Tx) GetR() string {
	if o == nil || IsNil(o.R) {
		var ret string
		return ret
	}
	return *o.R
}

// GetROk returns a tuple with the R field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tx) GetROk() (*string, bool) {
	if o == nil || IsNil(o.R) {
		return nil, false
	}
	return o.R, true
}

// HasR returns a boolean if a field has been set.
func (o *Tx) HasR() bool {
	if o != nil && !IsNil(o.R) {
		return true
	}

	return false
}

// SetR gets a reference to the given string and assigns it to the R field.
func (o *Tx) SetR(v string) {
	o.R = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *Tx) GetS() string {
	if o == nil || IsNil(o.S) {
		var ret string
		return ret
	}
	return *o.S
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tx) GetSOk() (*string, bool) {
	if o == nil || IsNil(o.S) {
		return nil, false
	}
	return o.S, true
}

// HasS returns a boolean if a field has been set.
func (o *Tx) HasS() bool {
	if o != nil && !IsNil(o.S) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *Tx) SetS(v string) {
	o.S = &v
}

func (o Tx) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tx) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableTx struct {
	value *Tx
	isSet bool
}

func (v NullableTx) Get() *Tx {
	return v.value
}

func (v *NullableTx) Set(val *Tx) {
	v.value = val
	v.isSet = true
}

func (v NullableTx) IsSet() bool {
	return v.isSet
}

func (v *NullableTx) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTx(val *Tx) *NullableTx {
	return &NullableTx{value: val, isSet: true}
}

func (v NullableTx) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTx) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

