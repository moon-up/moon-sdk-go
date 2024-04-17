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

// checks if the LitecoinTransactionOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LitecoinTransactionOutput{}

// LitecoinTransactionOutput struct for LitecoinTransactionOutput
type LitecoinTransactionOutput struct {
	SignedTx *string `json:"signedTx,omitempty"`
	TransactionHash *string `json:"transaction_hash,omitempty"`
}

// NewLitecoinTransactionOutput instantiates a new LitecoinTransactionOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLitecoinTransactionOutput() *LitecoinTransactionOutput {
	this := LitecoinTransactionOutput{}
	return &this
}

// NewLitecoinTransactionOutputWithDefaults instantiates a new LitecoinTransactionOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLitecoinTransactionOutputWithDefaults() *LitecoinTransactionOutput {
	this := LitecoinTransactionOutput{}
	return &this
}

// GetSignedTx returns the SignedTx field value if set, zero value otherwise.
func (o *LitecoinTransactionOutput) GetSignedTx() string {
	if o == nil || IsNil(o.SignedTx) {
		var ret string
		return ret
	}
	return *o.SignedTx
}

// GetSignedTxOk returns a tuple with the SignedTx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LitecoinTransactionOutput) GetSignedTxOk() (*string, bool) {
	if o == nil || IsNil(o.SignedTx) {
		return nil, false
	}
	return o.SignedTx, true
}

// HasSignedTx returns a boolean if a field has been set.
func (o *LitecoinTransactionOutput) HasSignedTx() bool {
	if o != nil && !IsNil(o.SignedTx) {
		return true
	}

	return false
}

// SetSignedTx gets a reference to the given string and assigns it to the SignedTx field.
func (o *LitecoinTransactionOutput) SetSignedTx(v string) {
	o.SignedTx = &v
}

// GetTransactionHash returns the TransactionHash field value if set, zero value otherwise.
func (o *LitecoinTransactionOutput) GetTransactionHash() string {
	if o == nil || IsNil(o.TransactionHash) {
		var ret string
		return ret
	}
	return *o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LitecoinTransactionOutput) GetTransactionHashOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionHash) {
		return nil, false
	}
	return o.TransactionHash, true
}

// HasTransactionHash returns a boolean if a field has been set.
func (o *LitecoinTransactionOutput) HasTransactionHash() bool {
	if o != nil && !IsNil(o.TransactionHash) {
		return true
	}

	return false
}

// SetTransactionHash gets a reference to the given string and assigns it to the TransactionHash field.
func (o *LitecoinTransactionOutput) SetTransactionHash(v string) {
	o.TransactionHash = &v
}

func (o LitecoinTransactionOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LitecoinTransactionOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SignedTx) {
		toSerialize["signedTx"] = o.SignedTx
	}
	if !IsNil(o.TransactionHash) {
		toSerialize["transaction_hash"] = o.TransactionHash
	}
	return toSerialize, nil
}

type NullableLitecoinTransactionOutput struct {
	value *LitecoinTransactionOutput
	isSet bool
}

func (v NullableLitecoinTransactionOutput) Get() *LitecoinTransactionOutput {
	return v.value
}

func (v *NullableLitecoinTransactionOutput) Set(val *LitecoinTransactionOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableLitecoinTransactionOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableLitecoinTransactionOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLitecoinTransactionOutput(val *LitecoinTransactionOutput) *NullableLitecoinTransactionOutput {
	return &NullableLitecoinTransactionOutput{value: val, isSet: true}
}

func (v NullableLitecoinTransactionOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLitecoinTransactionOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

