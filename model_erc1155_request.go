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

// checks if the Erc1155Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Erc1155Request{}

// Erc1155Request struct for Erc1155Request
type Erc1155Request struct {
	To *string `json:"to,omitempty"`
	Data *string `json:"data,omitempty"`
	Input *string `json:"input,omitempty"`
	Value *string `json:"value,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	Gas *string `json:"gas,omitempty"`
	GasPrice *string `json:"gasPrice,omitempty"`
	ChainId *string `json:"chain_id,omitempty"`
	Encoding *string `json:"encoding,omitempty"`
	EOA *bool `json:"EOA,omitempty"`
	ContractAddress *string `json:"contract_address,omitempty"`
	TokenId *string `json:"token_id,omitempty"`
	TokenIds *string `json:"token_ids,omitempty"`
	Approved *bool `json:"approved,omitempty"`
	Broadcast *bool `json:"broadcast,omitempty"`
}

// NewErc1155Request instantiates a new Erc1155Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErc1155Request() *Erc1155Request {
	this := Erc1155Request{}
	return &this
}

// NewErc1155RequestWithDefaults instantiates a new Erc1155Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErc1155RequestWithDefaults() *Erc1155Request {
	this := Erc1155Request{}
	return &this
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *Erc1155Request) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc1155Request) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *Erc1155Request) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *Erc1155Request) SetTo(v string) {
	o.To = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Erc1155Request) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc1155Request) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Erc1155Request) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *Erc1155Request) SetData(v string) {
	o.Data = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *Erc1155Request) GetInput() string {
	if o == nil || IsNil(o.Input) {
		var ret string
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc1155Request) GetInputOk() (*string, bool) {
	if o == nil || IsNil(o.Input) {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *Erc1155Request) HasInput() bool {
	if o != nil && !IsNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given string and assigns it to the Input field.
func (o *Erc1155Request) SetInput(v string) {
	o.Input = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Erc1155Request) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc1155Request) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Erc1155Request) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *Erc1155Request) SetValue(v string) {
	o.Value = &v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *Erc1155Request) GetNonce() string {
	if o == nil || IsNil(o.Nonce) {
		var ret string
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc1155Request) GetNonceOk() (*string, bool) {
	if o == nil || IsNil(o.Nonce) {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *Erc1155Request) HasNonce() bool {
	if o != nil && !IsNil(o.Nonce) {
		return true
	}

	return false
}

// SetNonce gets a reference to the given string and assigns it to the Nonce field.
func (o *Erc1155Request) SetNonce(v string) {
	o.Nonce = &v
}

// GetGas returns the Gas field value if set, zero value otherwise.
func (o *Erc1155Request) GetGas() string {
	if o == nil || IsNil(o.Gas) {
		var ret string
		return ret
	}
	return *o.Gas
}

// GetGasOk returns a tuple with the Gas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc1155Request) GetGasOk() (*string, bool) {
	if o == nil || IsNil(o.Gas) {
		return nil, false
	}
	return o.Gas, true
}

// HasGas returns a boolean if a field has been set.
func (o *Erc1155Request) HasGas() bool {
	if o != nil && !IsNil(o.Gas) {
		return true
	}

	return false
}

// SetGas gets a reference to the given string and assigns it to the Gas field.
func (o *Erc1155Request) SetGas(v string) {
	o.Gas = &v
}

// GetGasPrice returns the GasPrice field value if set, zero value otherwise.
func (o *Erc1155Request) GetGasPrice() string {
	if o == nil || IsNil(o.GasPrice) {
		var ret string
		return ret
	}
	return *o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc1155Request) GetGasPriceOk() (*string, bool) {
	if o == nil || IsNil(o.GasPrice) {
		return nil, false
	}
	return o.GasPrice, true
}

// HasGasPrice returns a boolean if a field has been set.
func (o *Erc1155Request) HasGasPrice() bool {
	if o != nil && !IsNil(o.GasPrice) {
		return true
	}

	return false
}

// SetGasPrice gets a reference to the given string and assigns it to the GasPrice field.
func (o *Erc1155Request) SetGasPrice(v string) {
	o.GasPrice = &v
}

// GetChainId returns the ChainId field value if set, zero value otherwise.
func (o *Erc1155Request) GetChainId() string {
	if o == nil || IsNil(o.ChainId) {
		var ret string
		return ret
	}
	return *o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc1155Request) GetChainIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChainId) {
		return nil, false
	}
	return o.ChainId, true
}

// HasChainId returns a boolean if a field has been set.
func (o *Erc1155Request) HasChainId() bool {
	if o != nil && !IsNil(o.ChainId) {
		return true
	}

	return false
}

// SetChainId gets a reference to the given string and assigns it to the ChainId field.
func (o *Erc1155Request) SetChainId(v string) {
	o.ChainId = &v
}

// GetEncoding returns the Encoding field value if set, zero value otherwise.
func (o *Erc1155Request) GetEncoding() string {
	if o == nil || IsNil(o.Encoding) {
		var ret string
		return ret
	}
	return *o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc1155Request) GetEncodingOk() (*string, bool) {
	if o == nil || IsNil(o.Encoding) {
		return nil, false
	}
	return o.Encoding, true
}

// HasEncoding returns a boolean if a field has been set.
func (o *Erc1155Request) HasEncoding() bool {
	if o != nil && !IsNil(o.Encoding) {
		return true
	}

	return false
}

// SetEncoding gets a reference to the given string and assigns it to the Encoding field.
func (o *Erc1155Request) SetEncoding(v string) {
	o.Encoding = &v
}

// GetEOA returns the EOA field value if set, zero value otherwise.
func (o *Erc1155Request) GetEOA() bool {
	if o == nil || IsNil(o.EOA) {
		var ret bool
		return ret
	}
	return *o.EOA
}

// GetEOAOk returns a tuple with the EOA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc1155Request) GetEOAOk() (*bool, bool) {
	if o == nil || IsNil(o.EOA) {
		return nil, false
	}
	return o.EOA, true
}

// HasEOA returns a boolean if a field has been set.
func (o *Erc1155Request) HasEOA() bool {
	if o != nil && !IsNil(o.EOA) {
		return true
	}

	return false
}

// SetEOA gets a reference to the given bool and assigns it to the EOA field.
func (o *Erc1155Request) SetEOA(v bool) {
	o.EOA = &v
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise.
func (o *Erc1155Request) GetContractAddress() string {
	if o == nil || IsNil(o.ContractAddress) {
		var ret string
		return ret
	}
	return *o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc1155Request) GetContractAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ContractAddress) {
		return nil, false
	}
	return o.ContractAddress, true
}

// HasContractAddress returns a boolean if a field has been set.
func (o *Erc1155Request) HasContractAddress() bool {
	if o != nil && !IsNil(o.ContractAddress) {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given string and assigns it to the ContractAddress field.
func (o *Erc1155Request) SetContractAddress(v string) {
	o.ContractAddress = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *Erc1155Request) GetTokenId() string {
	if o == nil || IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc1155Request) GetTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *Erc1155Request) HasTokenId() bool {
	if o != nil && !IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *Erc1155Request) SetTokenId(v string) {
	o.TokenId = &v
}

// GetTokenIds returns the TokenIds field value if set, zero value otherwise.
func (o *Erc1155Request) GetTokenIds() string {
	if o == nil || IsNil(o.TokenIds) {
		var ret string
		return ret
	}
	return *o.TokenIds
}

// GetTokenIdsOk returns a tuple with the TokenIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc1155Request) GetTokenIdsOk() (*string, bool) {
	if o == nil || IsNil(o.TokenIds) {
		return nil, false
	}
	return o.TokenIds, true
}

// HasTokenIds returns a boolean if a field has been set.
func (o *Erc1155Request) HasTokenIds() bool {
	if o != nil && !IsNil(o.TokenIds) {
		return true
	}

	return false
}

// SetTokenIds gets a reference to the given string and assigns it to the TokenIds field.
func (o *Erc1155Request) SetTokenIds(v string) {
	o.TokenIds = &v
}

// GetApproved returns the Approved field value if set, zero value otherwise.
func (o *Erc1155Request) GetApproved() bool {
	if o == nil || IsNil(o.Approved) {
		var ret bool
		return ret
	}
	return *o.Approved
}

// GetApprovedOk returns a tuple with the Approved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc1155Request) GetApprovedOk() (*bool, bool) {
	if o == nil || IsNil(o.Approved) {
		return nil, false
	}
	return o.Approved, true
}

// HasApproved returns a boolean if a field has been set.
func (o *Erc1155Request) HasApproved() bool {
	if o != nil && !IsNil(o.Approved) {
		return true
	}

	return false
}

// SetApproved gets a reference to the given bool and assigns it to the Approved field.
func (o *Erc1155Request) SetApproved(v bool) {
	o.Approved = &v
}

// GetBroadcast returns the Broadcast field value if set, zero value otherwise.
func (o *Erc1155Request) GetBroadcast() bool {
	if o == nil || IsNil(o.Broadcast) {
		var ret bool
		return ret
	}
	return *o.Broadcast
}

// GetBroadcastOk returns a tuple with the Broadcast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Erc1155Request) GetBroadcastOk() (*bool, bool) {
	if o == nil || IsNil(o.Broadcast) {
		return nil, false
	}
	return o.Broadcast, true
}

// HasBroadcast returns a boolean if a field has been set.
func (o *Erc1155Request) HasBroadcast() bool {
	if o != nil && !IsNil(o.Broadcast) {
		return true
	}

	return false
}

// SetBroadcast gets a reference to the given bool and assigns it to the Broadcast field.
func (o *Erc1155Request) SetBroadcast(v bool) {
	o.Broadcast = &v
}

func (o Erc1155Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Erc1155Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Input) {
		toSerialize["input"] = o.Input
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Nonce) {
		toSerialize["nonce"] = o.Nonce
	}
	if !IsNil(o.Gas) {
		toSerialize["gas"] = o.Gas
	}
	if !IsNil(o.GasPrice) {
		toSerialize["gasPrice"] = o.GasPrice
	}
	if !IsNil(o.ChainId) {
		toSerialize["chain_id"] = o.ChainId
	}
	if !IsNil(o.Encoding) {
		toSerialize["encoding"] = o.Encoding
	}
	if !IsNil(o.EOA) {
		toSerialize["EOA"] = o.EOA
	}
	if !IsNil(o.ContractAddress) {
		toSerialize["contract_address"] = o.ContractAddress
	}
	if !IsNil(o.TokenId) {
		toSerialize["token_id"] = o.TokenId
	}
	if !IsNil(o.TokenIds) {
		toSerialize["token_ids"] = o.TokenIds
	}
	if !IsNil(o.Approved) {
		toSerialize["approved"] = o.Approved
	}
	if !IsNil(o.Broadcast) {
		toSerialize["broadcast"] = o.Broadcast
	}
	return toSerialize, nil
}

type NullableErc1155Request struct {
	value *Erc1155Request
	isSet bool
}

func (v NullableErc1155Request) Get() *Erc1155Request {
	return v.value
}

func (v *NullableErc1155Request) Set(val *Erc1155Request) {
	v.value = val
	v.isSet = true
}

func (v NullableErc1155Request) IsSet() bool {
	return v.isSet
}

func (v *NullableErc1155Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErc1155Request(val *Erc1155Request) *NullableErc1155Request {
	return &NullableErc1155Request{value: val, isSet: true}
}

func (v NullableErc1155Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErc1155Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

