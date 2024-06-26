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

// checks if the EosTransactionOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EosTransactionOutput{}

// EosTransactionOutput struct for EosTransactionOutput
type EosTransactionOutput struct {
	SignedTx *string `json:"signedTx,omitempty"`
	TransactionHash *string `json:"transaction_hash,omitempty"`
}

// NewEosTransactionOutput instantiates a new EosTransactionOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEosTransactionOutput() *EosTransactionOutput {
	this := EosTransactionOutput{}
	return &this
}

// NewEosTransactionOutputWithDefaults instantiates a new EosTransactionOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEosTransactionOutputWithDefaults() *EosTransactionOutput {
	this := EosTransactionOutput{}
	return &this
}

// GetSignedTx returns the SignedTx field value if set, zero value otherwise.
func (o *EosTransactionOutput) GetSignedTx() string {
	if o == nil || IsNil(o.SignedTx) {
		var ret string
		return ret
	}
	return *o.SignedTx
}

// GetSignedTxOk returns a tuple with the SignedTx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EosTransactionOutput) GetSignedTxOk() (*string, bool) {
	if o == nil || IsNil(o.SignedTx) {
		return nil, false
	}
	return o.SignedTx, true
}

// HasSignedTx returns a boolean if a field has been set.
func (o *EosTransactionOutput) HasSignedTx() bool {
	if o != nil && !IsNil(o.SignedTx) {
		return true
	}

	return false
}

// SetSignedTx gets a reference to the given string and assigns it to the SignedTx field.
func (o *EosTransactionOutput) SetSignedTx(v string) {
	o.SignedTx = &v
}

// GetTransactionHash returns the TransactionHash field value if set, zero value otherwise.
func (o *EosTransactionOutput) GetTransactionHash() string {
	if o == nil || IsNil(o.TransactionHash) {
		var ret string
		return ret
	}
	return *o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EosTransactionOutput) GetTransactionHashOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionHash) {
		return nil, false
	}
	return o.TransactionHash, true
}

// HasTransactionHash returns a boolean if a field has been set.
func (o *EosTransactionOutput) HasTransactionHash() bool {
	if o != nil && !IsNil(o.TransactionHash) {
		return true
	}

	return false
}

// SetTransactionHash gets a reference to the given string and assigns it to the TransactionHash field.
func (o *EosTransactionOutput) SetTransactionHash(v string) {
	o.TransactionHash = &v
}

func (o EosTransactionOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EosTransactionOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SignedTx) {
		toSerialize["signedTx"] = o.SignedTx
	}
	if !IsNil(o.TransactionHash) {
		toSerialize["transaction_hash"] = o.TransactionHash
	}
	return toSerialize, nil
}

type NullableEosTransactionOutput struct {
	value *EosTransactionOutput
	isSet bool
}

func (v NullableEosTransactionOutput) Get() *EosTransactionOutput {
	return v.value
}

func (v *NullableEosTransactionOutput) Set(val *EosTransactionOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableEosTransactionOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableEosTransactionOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEosTransactionOutput(val *EosTransactionOutput) *NullableEosTransactionOutput {
	return &NullableEosTransactionOutput{value: val, isSet: true}
}

func (v NullableEosTransactionOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEosTransactionOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


