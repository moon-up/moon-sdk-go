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

// checks if the GetSupportedOnRampsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSupportedOnRampsResponse{}

// GetSupportedOnRampsResponse struct for GetSupportedOnRampsResponse
type GetSupportedOnRampsResponse struct {
	Message []GetSupportedOnRampsResponseMessageInner `json:"message"`
}

type _GetSupportedOnRampsResponse GetSupportedOnRampsResponse

// NewGetSupportedOnRampsResponse instantiates a new GetSupportedOnRampsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSupportedOnRampsResponse(message []GetSupportedOnRampsResponseMessageInner) *GetSupportedOnRampsResponse {
	this := GetSupportedOnRampsResponse{}
	this.Message = message
	return &this
}

// NewGetSupportedOnRampsResponseWithDefaults instantiates a new GetSupportedOnRampsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSupportedOnRampsResponseWithDefaults() *GetSupportedOnRampsResponse {
	this := GetSupportedOnRampsResponse{}
	return &this
}

// GetMessage returns the Message field value
func (o *GetSupportedOnRampsResponse) GetMessage() []GetSupportedOnRampsResponseMessageInner {
	if o == nil {
		var ret []GetSupportedOnRampsResponseMessageInner
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *GetSupportedOnRampsResponse) GetMessageOk() ([]GetSupportedOnRampsResponseMessageInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message, true
}

// SetMessage sets field value
func (o *GetSupportedOnRampsResponse) SetMessage(v []GetSupportedOnRampsResponseMessageInner) {
	o.Message = v
}

func (o GetSupportedOnRampsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSupportedOnRampsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *GetSupportedOnRampsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetSupportedOnRampsResponse := _GetSupportedOnRampsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetSupportedOnRampsResponse)

	if err != nil {
		return err
	}

	*o = GetSupportedOnRampsResponse(varGetSupportedOnRampsResponse)

	return err
}

type NullableGetSupportedOnRampsResponse struct {
	value *GetSupportedOnRampsResponse
	isSet bool
}

func (v NullableGetSupportedOnRampsResponse) Get() *GetSupportedOnRampsResponse {
	return v.value
}

func (v *NullableGetSupportedOnRampsResponse) Set(val *GetSupportedOnRampsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSupportedOnRampsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSupportedOnRampsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSupportedOnRampsResponse(val *GetSupportedOnRampsResponse) *NullableGetSupportedOnRampsResponse {
	return &NullableGetSupportedOnRampsResponse{value: val, isSet: true}
}

func (v NullableGetSupportedOnRampsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSupportedOnRampsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


