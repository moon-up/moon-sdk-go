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

// checks if the LitecoinTransactionInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LitecoinTransactionInput{}

// LitecoinTransactionInput struct for LitecoinTransactionInput
type LitecoinTransactionInput struct {
	To *string `json:"to,omitempty"`
	Value *float64 `json:"value,omitempty"`
	Network *string `json:"network,omitempty"`
	Compress *bool `json:"compress,omitempty"`
}

// NewLitecoinTransactionInput instantiates a new LitecoinTransactionInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLitecoinTransactionInput() *LitecoinTransactionInput {
	this := LitecoinTransactionInput{}
	return &this
}

// NewLitecoinTransactionInputWithDefaults instantiates a new LitecoinTransactionInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLitecoinTransactionInputWithDefaults() *LitecoinTransactionInput {
	this := LitecoinTransactionInput{}
	return &this
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *LitecoinTransactionInput) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LitecoinTransactionInput) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *LitecoinTransactionInput) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *LitecoinTransactionInput) SetTo(v string) {
	o.To = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *LitecoinTransactionInput) GetValue() float64 {
	if o == nil || IsNil(o.Value) {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LitecoinTransactionInput) GetValueOk() (*float64, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *LitecoinTransactionInput) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *LitecoinTransactionInput) SetValue(v float64) {
	o.Value = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *LitecoinTransactionInput) GetNetwork() string {
	if o == nil || IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LitecoinTransactionInput) GetNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *LitecoinTransactionInput) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *LitecoinTransactionInput) SetNetwork(v string) {
	o.Network = &v
}

// GetCompress returns the Compress field value if set, zero value otherwise.
func (o *LitecoinTransactionInput) GetCompress() bool {
	if o == nil || IsNil(o.Compress) {
		var ret bool
		return ret
	}
	return *o.Compress
}

// GetCompressOk returns a tuple with the Compress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LitecoinTransactionInput) GetCompressOk() (*bool, bool) {
	if o == nil || IsNil(o.Compress) {
		return nil, false
	}
	return o.Compress, true
}

// HasCompress returns a boolean if a field has been set.
func (o *LitecoinTransactionInput) HasCompress() bool {
	if o != nil && !IsNil(o.Compress) {
		return true
	}

	return false
}

// SetCompress gets a reference to the given bool and assigns it to the Compress field.
func (o *LitecoinTransactionInput) SetCompress(v bool) {
	o.Compress = &v
}

func (o LitecoinTransactionInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LitecoinTransactionInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.Compress) {
		toSerialize["compress"] = o.Compress
	}
	return toSerialize, nil
}

type NullableLitecoinTransactionInput struct {
	value *LitecoinTransactionInput
	isSet bool
}

func (v NullableLitecoinTransactionInput) Get() *LitecoinTransactionInput {
	return v.value
}

func (v *NullableLitecoinTransactionInput) Set(val *LitecoinTransactionInput) {
	v.value = val
	v.isSet = true
}

func (v NullableLitecoinTransactionInput) IsSet() bool {
	return v.isSet
}

func (v *NullableLitecoinTransactionInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLitecoinTransactionInput(val *LitecoinTransactionInput) *NullableLitecoinTransactionInput {
	return &NullableLitecoinTransactionInput{value: val, isSet: true}
}

func (v NullableLitecoinTransactionInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLitecoinTransactionInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


