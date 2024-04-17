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

// checks if the TronInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TronInput{}

// TronInput struct for TronInput
type TronInput struct {
	Network *string `json:"network,omitempty"`
	PrivateKey *string `json:"private_key,omitempty"`
}

// NewTronInput instantiates a new TronInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTronInput() *TronInput {
	this := TronInput{}
	return &this
}

// NewTronInputWithDefaults instantiates a new TronInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTronInputWithDefaults() *TronInput {
	this := TronInput{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *TronInput) GetNetwork() string {
	if o == nil || IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TronInput) GetNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *TronInput) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *TronInput) SetNetwork(v string) {
	o.Network = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *TronInput) GetPrivateKey() string {
	if o == nil || IsNil(o.PrivateKey) {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TronInput) GetPrivateKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateKey) {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *TronInput) HasPrivateKey() bool {
	if o != nil && !IsNil(o.PrivateKey) {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *TronInput) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

func (o TronInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TronInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.PrivateKey) {
		toSerialize["private_key"] = o.PrivateKey
	}
	return toSerialize, nil
}

type NullableTronInput struct {
	value *TronInput
	isSet bool
}

func (v NullableTronInput) Get() *TronInput {
	return v.value
}

func (v *NullableTronInput) Set(val *TronInput) {
	v.value = val
	v.isSet = true
}

func (v NullableTronInput) IsSet() bool {
	return v.isSet
}

func (v *NullableTronInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTronInput(val *TronInput) *NullableTronInput {
	return &NullableTronInput{value: val, isSet: true}
}

func (v NullableTronInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTronInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

