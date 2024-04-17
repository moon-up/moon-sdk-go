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

// checks if the AvailablePaymentMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailablePaymentMethod{}

// AvailablePaymentMethod struct for AvailablePaymentMethod
type AvailablePaymentMethod struct {
	Icon string `json:"icon"`
	Name string `json:"name"`
	PaymentTypeId string `json:"paymentTypeId"`
}

type _AvailablePaymentMethod AvailablePaymentMethod

// NewAvailablePaymentMethod instantiates a new AvailablePaymentMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailablePaymentMethod(icon string, name string, paymentTypeId string) *AvailablePaymentMethod {
	this := AvailablePaymentMethod{}
	this.Icon = icon
	this.Name = name
	this.PaymentTypeId = paymentTypeId
	return &this
}

// NewAvailablePaymentMethodWithDefaults instantiates a new AvailablePaymentMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailablePaymentMethodWithDefaults() *AvailablePaymentMethod {
	this := AvailablePaymentMethod{}
	return &this
}

// GetIcon returns the Icon field value
func (o *AvailablePaymentMethod) GetIcon() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Icon
}

// GetIconOk returns a tuple with the Icon field value
// and a boolean to check if the value has been set.
func (o *AvailablePaymentMethod) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Icon, true
}

// SetIcon sets field value
func (o *AvailablePaymentMethod) SetIcon(v string) {
	o.Icon = v
}

// GetName returns the Name field value
func (o *AvailablePaymentMethod) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AvailablePaymentMethod) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AvailablePaymentMethod) SetName(v string) {
	o.Name = v
}

// GetPaymentTypeId returns the PaymentTypeId field value
func (o *AvailablePaymentMethod) GetPaymentTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentTypeId
}

// GetPaymentTypeIdOk returns a tuple with the PaymentTypeId field value
// and a boolean to check if the value has been set.
func (o *AvailablePaymentMethod) GetPaymentTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentTypeId, true
}

// SetPaymentTypeId sets field value
func (o *AvailablePaymentMethod) SetPaymentTypeId(v string) {
	o.PaymentTypeId = v
}

func (o AvailablePaymentMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvailablePaymentMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["icon"] = o.Icon
	toSerialize["name"] = o.Name
	toSerialize["paymentTypeId"] = o.PaymentTypeId
	return toSerialize, nil
}

func (o *AvailablePaymentMethod) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"icon",
		"name",
		"paymentTypeId",
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

	varAvailablePaymentMethod := _AvailablePaymentMethod{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAvailablePaymentMethod)

	if err != nil {
		return err
	}

	*o = AvailablePaymentMethod(varAvailablePaymentMethod)

	return err
}

type NullableAvailablePaymentMethod struct {
	value *AvailablePaymentMethod
	isSet bool
}

func (v NullableAvailablePaymentMethod) Get() *AvailablePaymentMethod {
	return v.value
}

func (v *NullableAvailablePaymentMethod) Set(val *AvailablePaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailablePaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailablePaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailablePaymentMethod(val *AvailablePaymentMethod) *NullableAvailablePaymentMethod {
	return &NullableAvailablePaymentMethod{value: val, isSet: true}
}

func (v NullableAvailablePaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailablePaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


