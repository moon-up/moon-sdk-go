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

// checks if the TransactionInputSupportedParamsTheme type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionInputSupportedParamsTheme{}

// TransactionInputSupportedParamsTheme struct for TransactionInputSupportedParamsTheme
type TransactionInputSupportedParamsTheme struct {
	BorderRadius NullableFloat64 `json:"borderRadius"`
	CardColor string `json:"cardColor"`
	SecondaryTextColor string `json:"secondaryTextColor"`
	PrimaryTextColor string `json:"primaryTextColor"`
	SecondaryColor string `json:"secondaryColor"`
	PrimaryColor string `json:"primaryColor"`
	ThemeName string `json:"themeName"`
	IsDark bool `json:"isDark"`
}

type _TransactionInputSupportedParamsTheme TransactionInputSupportedParamsTheme

// NewTransactionInputSupportedParamsTheme instantiates a new TransactionInputSupportedParamsTheme object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionInputSupportedParamsTheme(borderRadius NullableFloat64, cardColor string, secondaryTextColor string, primaryTextColor string, secondaryColor string, primaryColor string, themeName string, isDark bool) *TransactionInputSupportedParamsTheme {
	this := TransactionInputSupportedParamsTheme{}
	this.BorderRadius = borderRadius
	this.CardColor = cardColor
	this.SecondaryTextColor = secondaryTextColor
	this.PrimaryTextColor = primaryTextColor
	this.SecondaryColor = secondaryColor
	this.PrimaryColor = primaryColor
	this.ThemeName = themeName
	this.IsDark = isDark
	return &this
}

// NewTransactionInputSupportedParamsThemeWithDefaults instantiates a new TransactionInputSupportedParamsTheme object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionInputSupportedParamsThemeWithDefaults() *TransactionInputSupportedParamsTheme {
	this := TransactionInputSupportedParamsTheme{}
	return &this
}

// GetBorderRadius returns the BorderRadius field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *TransactionInputSupportedParamsTheme) GetBorderRadius() float64 {
	if o == nil || o.BorderRadius.Get() == nil {
		var ret float64
		return ret
	}

	return *o.BorderRadius.Get()
}

// GetBorderRadiusOk returns a tuple with the BorderRadius field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionInputSupportedParamsTheme) GetBorderRadiusOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BorderRadius.Get(), o.BorderRadius.IsSet()
}

// SetBorderRadius sets field value
func (o *TransactionInputSupportedParamsTheme) SetBorderRadius(v float64) {
	o.BorderRadius.Set(&v)
}

// GetCardColor returns the CardColor field value
func (o *TransactionInputSupportedParamsTheme) GetCardColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardColor
}

// GetCardColorOk returns a tuple with the CardColor field value
// and a boolean to check if the value has been set.
func (o *TransactionInputSupportedParamsTheme) GetCardColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardColor, true
}

// SetCardColor sets field value
func (o *TransactionInputSupportedParamsTheme) SetCardColor(v string) {
	o.CardColor = v
}

// GetSecondaryTextColor returns the SecondaryTextColor field value
func (o *TransactionInputSupportedParamsTheme) GetSecondaryTextColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecondaryTextColor
}

// GetSecondaryTextColorOk returns a tuple with the SecondaryTextColor field value
// and a boolean to check if the value has been set.
func (o *TransactionInputSupportedParamsTheme) GetSecondaryTextColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecondaryTextColor, true
}

// SetSecondaryTextColor sets field value
func (o *TransactionInputSupportedParamsTheme) SetSecondaryTextColor(v string) {
	o.SecondaryTextColor = v
}

// GetPrimaryTextColor returns the PrimaryTextColor field value
func (o *TransactionInputSupportedParamsTheme) GetPrimaryTextColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrimaryTextColor
}

// GetPrimaryTextColorOk returns a tuple with the PrimaryTextColor field value
// and a boolean to check if the value has been set.
func (o *TransactionInputSupportedParamsTheme) GetPrimaryTextColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryTextColor, true
}

// SetPrimaryTextColor sets field value
func (o *TransactionInputSupportedParamsTheme) SetPrimaryTextColor(v string) {
	o.PrimaryTextColor = v
}

// GetSecondaryColor returns the SecondaryColor field value
func (o *TransactionInputSupportedParamsTheme) GetSecondaryColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecondaryColor
}

// GetSecondaryColorOk returns a tuple with the SecondaryColor field value
// and a boolean to check if the value has been set.
func (o *TransactionInputSupportedParamsTheme) GetSecondaryColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecondaryColor, true
}

// SetSecondaryColor sets field value
func (o *TransactionInputSupportedParamsTheme) SetSecondaryColor(v string) {
	o.SecondaryColor = v
}

// GetPrimaryColor returns the PrimaryColor field value
func (o *TransactionInputSupportedParamsTheme) GetPrimaryColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrimaryColor
}

// GetPrimaryColorOk returns a tuple with the PrimaryColor field value
// and a boolean to check if the value has been set.
func (o *TransactionInputSupportedParamsTheme) GetPrimaryColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryColor, true
}

// SetPrimaryColor sets field value
func (o *TransactionInputSupportedParamsTheme) SetPrimaryColor(v string) {
	o.PrimaryColor = v
}

// GetThemeName returns the ThemeName field value
func (o *TransactionInputSupportedParamsTheme) GetThemeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThemeName
}

// GetThemeNameOk returns a tuple with the ThemeName field value
// and a boolean to check if the value has been set.
func (o *TransactionInputSupportedParamsTheme) GetThemeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThemeName, true
}

// SetThemeName sets field value
func (o *TransactionInputSupportedParamsTheme) SetThemeName(v string) {
	o.ThemeName = v
}

// GetIsDark returns the IsDark field value
func (o *TransactionInputSupportedParamsTheme) GetIsDark() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDark
}

// GetIsDarkOk returns a tuple with the IsDark field value
// and a boolean to check if the value has been set.
func (o *TransactionInputSupportedParamsTheme) GetIsDarkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDark, true
}

// SetIsDark sets field value
func (o *TransactionInputSupportedParamsTheme) SetIsDark(v bool) {
	o.IsDark = v
}

func (o TransactionInputSupportedParamsTheme) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionInputSupportedParamsTheme) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["borderRadius"] = o.BorderRadius.Get()
	toSerialize["cardColor"] = o.CardColor
	toSerialize["secondaryTextColor"] = o.SecondaryTextColor
	toSerialize["primaryTextColor"] = o.PrimaryTextColor
	toSerialize["secondaryColor"] = o.SecondaryColor
	toSerialize["primaryColor"] = o.PrimaryColor
	toSerialize["themeName"] = o.ThemeName
	toSerialize["isDark"] = o.IsDark
	return toSerialize, nil
}

func (o *TransactionInputSupportedParamsTheme) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"borderRadius",
		"cardColor",
		"secondaryTextColor",
		"primaryTextColor",
		"secondaryColor",
		"primaryColor",
		"themeName",
		"isDark",
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

	varTransactionInputSupportedParamsTheme := _TransactionInputSupportedParamsTheme{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionInputSupportedParamsTheme)

	if err != nil {
		return err
	}

	*o = TransactionInputSupportedParamsTheme(varTransactionInputSupportedParamsTheme)

	return err
}

type NullableTransactionInputSupportedParamsTheme struct {
	value *TransactionInputSupportedParamsTheme
	isSet bool
}

func (v NullableTransactionInputSupportedParamsTheme) Get() *TransactionInputSupportedParamsTheme {
	return v.value
}

func (v *NullableTransactionInputSupportedParamsTheme) Set(val *TransactionInputSupportedParamsTheme) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionInputSupportedParamsTheme) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionInputSupportedParamsTheme) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionInputSupportedParamsTheme(val *TransactionInputSupportedParamsTheme) *NullableTransactionInputSupportedParamsTheme {
	return &NullableTransactionInputSupportedParamsTheme{value: val, isSet: true}
}

func (v NullableTransactionInputSupportedParamsTheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionInputSupportedParamsTheme) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


