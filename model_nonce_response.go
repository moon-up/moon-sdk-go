/*
moon-vault-api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package moonsdk

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the NonceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonceResponse{}

// NonceResponse struct for NonceResponse
type NonceResponse struct {
	Nonce float64 `json:"nonce"`
}

type _NonceResponse NonceResponse

// NewNonceResponse instantiates a new NonceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonceResponse(nonce float64) *NonceResponse {
	this := NonceResponse{}
	this.Nonce = nonce
	return &this
}

// NewNonceResponseWithDefaults instantiates a new NonceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonceResponseWithDefaults() *NonceResponse {
	this := NonceResponse{}
	return &this
}

// GetNonce returns the Nonce field value
func (o *NonceResponse) GetNonce() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *NonceResponse) GetNonceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *NonceResponse) SetNonce(v float64) {
	o.Nonce = v
}

func (o NonceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nonce"] = o.Nonce
	return toSerialize, nil
}

func (o *NonceResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"nonce",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNonceResponse := _NonceResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNonceResponse)

	if err != nil {
		return err
	}

	*o = NonceResponse(varNonceResponse)

	return err
}

type NullableNonceResponse struct {
	value *NonceResponse
	isSet bool
}

func (v NullableNonceResponse) Get() *NonceResponse {
	return v.value
}

func (v *NullableNonceResponse) Set(val *NonceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNonceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNonceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonceResponse(val *NonceResponse) *NullableNonceResponse {
	return &NullableNonceResponse{value: val, isSet: true}
}

func (v NullableNonceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


