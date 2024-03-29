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

// checks if the BitcoinTransactionOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BitcoinTransactionOutput{}

// BitcoinTransactionOutput struct for BitcoinTransactionOutput
type BitcoinTransactionOutput struct {
	SignedTx *string `json:"signedTx,omitempty"`
	TransactionHash *string `json:"transaction_hash,omitempty"`
}

// NewBitcoinTransactionOutput instantiates a new BitcoinTransactionOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBitcoinTransactionOutput() *BitcoinTransactionOutput {
	this := BitcoinTransactionOutput{}
	return &this
}

// NewBitcoinTransactionOutputWithDefaults instantiates a new BitcoinTransactionOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBitcoinTransactionOutputWithDefaults() *BitcoinTransactionOutput {
	this := BitcoinTransactionOutput{}
	return &this
}

// GetSignedTx returns the SignedTx field value if set, zero value otherwise.
func (o *BitcoinTransactionOutput) GetSignedTx() string {
	if o == nil || IsNil(o.SignedTx) {
		var ret string
		return ret
	}
	return *o.SignedTx
}

// GetSignedTxOk returns a tuple with the SignedTx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BitcoinTransactionOutput) GetSignedTxOk() (*string, bool) {
	if o == nil || IsNil(o.SignedTx) {
		return nil, false
	}
	return o.SignedTx, true
}

// HasSignedTx returns a boolean if a field has been set.
func (o *BitcoinTransactionOutput) HasSignedTx() bool {
	if o != nil && !IsNil(o.SignedTx) {
		return true
	}

	return false
}

// SetSignedTx gets a reference to the given string and assigns it to the SignedTx field.
func (o *BitcoinTransactionOutput) SetSignedTx(v string) {
	o.SignedTx = &v
}

// GetTransactionHash returns the TransactionHash field value if set, zero value otherwise.
func (o *BitcoinTransactionOutput) GetTransactionHash() string {
	if o == nil || IsNil(o.TransactionHash) {
		var ret string
		return ret
	}
	return *o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BitcoinTransactionOutput) GetTransactionHashOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionHash) {
		return nil, false
	}
	return o.TransactionHash, true
}

// HasTransactionHash returns a boolean if a field has been set.
func (o *BitcoinTransactionOutput) HasTransactionHash() bool {
	if o != nil && !IsNil(o.TransactionHash) {
		return true
	}

	return false
}

// SetTransactionHash gets a reference to the given string and assigns it to the TransactionHash field.
func (o *BitcoinTransactionOutput) SetTransactionHash(v string) {
	o.TransactionHash = &v
}

func (o BitcoinTransactionOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BitcoinTransactionOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SignedTx) {
		toSerialize["signedTx"] = o.SignedTx
	}
	if !IsNil(o.TransactionHash) {
		toSerialize["transaction_hash"] = o.TransactionHash
	}
	return toSerialize, nil
}

type NullableBitcoinTransactionOutput struct {
	value *BitcoinTransactionOutput
	isSet bool
}

func (v NullableBitcoinTransactionOutput) Get() *BitcoinTransactionOutput {
	return v.value
}

func (v *NullableBitcoinTransactionOutput) Set(val *BitcoinTransactionOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableBitcoinTransactionOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableBitcoinTransactionOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBitcoinTransactionOutput(val *BitcoinTransactionOutput) *NullableBitcoinTransactionOutput {
	return &NullableBitcoinTransactionOutput{value: val, isSet: true}
}

func (v NullableBitcoinTransactionOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBitcoinTransactionOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


