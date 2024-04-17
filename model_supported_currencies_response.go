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

// checks if the SupportedCurrenciesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupportedCurrenciesResponse{}

// SupportedCurrenciesResponse struct for SupportedCurrenciesResponse
type SupportedCurrenciesResponse struct {
	Message Message `json:"message"`
}

type _SupportedCurrenciesResponse SupportedCurrenciesResponse

// NewSupportedCurrenciesResponse instantiates a new SupportedCurrenciesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportedCurrenciesResponse(message Message) *SupportedCurrenciesResponse {
	this := SupportedCurrenciesResponse{}
	this.Message = message
	return &this
}

// NewSupportedCurrenciesResponseWithDefaults instantiates a new SupportedCurrenciesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportedCurrenciesResponseWithDefaults() *SupportedCurrenciesResponse {
	this := SupportedCurrenciesResponse{}
	return &this
}

// GetMessage returns the Message field value
func (o *SupportedCurrenciesResponse) GetMessage() Message {
	if o == nil {
		var ret Message
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SupportedCurrenciesResponse) GetMessageOk() (*Message, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SupportedCurrenciesResponse) SetMessage(v Message) {
	o.Message = v
}

func (o SupportedCurrenciesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupportedCurrenciesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *SupportedCurrenciesResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
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

	varSupportedCurrenciesResponse := _SupportedCurrenciesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSupportedCurrenciesResponse)

	if err != nil {
		return err
	}

	*o = SupportedCurrenciesResponse(varSupportedCurrenciesResponse)

	return err
}

type NullableSupportedCurrenciesResponse struct {
	value *SupportedCurrenciesResponse
	isSet bool
}

func (v NullableSupportedCurrenciesResponse) Get() *SupportedCurrenciesResponse {
	return v.value
}

func (v *NullableSupportedCurrenciesResponse) Set(val *SupportedCurrenciesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportedCurrenciesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportedCurrenciesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportedCurrenciesResponse(val *SupportedCurrenciesResponse) *NullableSupportedCurrenciesResponse {
	return &NullableSupportedCurrenciesResponse{value: val, isSet: true}
}

func (v NullableSupportedCurrenciesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportedCurrenciesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

