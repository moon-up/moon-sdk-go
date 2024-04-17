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

// checks if the TransactionInputSupportedParamsPartnerDataRedirectUrl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionInputSupportedParamsPartnerDataRedirectUrl{}

// TransactionInputSupportedParamsPartnerDataRedirectUrl struct for TransactionInputSupportedParamsPartnerDataRedirectUrl
type TransactionInputSupportedParamsPartnerDataRedirectUrl struct {
	Success string `json:"success"`
}

type _TransactionInputSupportedParamsPartnerDataRedirectUrl TransactionInputSupportedParamsPartnerDataRedirectUrl

// NewTransactionInputSupportedParamsPartnerDataRedirectUrl instantiates a new TransactionInputSupportedParamsPartnerDataRedirectUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionInputSupportedParamsPartnerDataRedirectUrl(success string) *TransactionInputSupportedParamsPartnerDataRedirectUrl {
	this := TransactionInputSupportedParamsPartnerDataRedirectUrl{}
	this.Success = success
	return &this
}

// NewTransactionInputSupportedParamsPartnerDataRedirectUrlWithDefaults instantiates a new TransactionInputSupportedParamsPartnerDataRedirectUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionInputSupportedParamsPartnerDataRedirectUrlWithDefaults() *TransactionInputSupportedParamsPartnerDataRedirectUrl {
	this := TransactionInputSupportedParamsPartnerDataRedirectUrl{}
	return &this
}

// GetSuccess returns the Success field value
func (o *TransactionInputSupportedParamsPartnerDataRedirectUrl) GetSuccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *TransactionInputSupportedParamsPartnerDataRedirectUrl) GetSuccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *TransactionInputSupportedParamsPartnerDataRedirectUrl) SetSuccess(v string) {
	o.Success = v
}

func (o TransactionInputSupportedParamsPartnerDataRedirectUrl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionInputSupportedParamsPartnerDataRedirectUrl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	return toSerialize, nil
}

func (o *TransactionInputSupportedParamsPartnerDataRedirectUrl) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"success",
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

	varTransactionInputSupportedParamsPartnerDataRedirectUrl := _TransactionInputSupportedParamsPartnerDataRedirectUrl{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionInputSupportedParamsPartnerDataRedirectUrl)

	if err != nil {
		return err
	}

	*o = TransactionInputSupportedParamsPartnerDataRedirectUrl(varTransactionInputSupportedParamsPartnerDataRedirectUrl)

	return err
}

type NullableTransactionInputSupportedParamsPartnerDataRedirectUrl struct {
	value *TransactionInputSupportedParamsPartnerDataRedirectUrl
	isSet bool
}

func (v NullableTransactionInputSupportedParamsPartnerDataRedirectUrl) Get() *TransactionInputSupportedParamsPartnerDataRedirectUrl {
	return v.value
}

func (v *NullableTransactionInputSupportedParamsPartnerDataRedirectUrl) Set(val *TransactionInputSupportedParamsPartnerDataRedirectUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionInputSupportedParamsPartnerDataRedirectUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionInputSupportedParamsPartnerDataRedirectUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionInputSupportedParamsPartnerDataRedirectUrl(val *TransactionInputSupportedParamsPartnerDataRedirectUrl) *NullableTransactionInputSupportedParamsPartnerDataRedirectUrl {
	return &NullableTransactionInputSupportedParamsPartnerDataRedirectUrl{value: val, isSet: true}
}

func (v NullableTransactionInputSupportedParamsPartnerDataRedirectUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionInputSupportedParamsPartnerDataRedirectUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


