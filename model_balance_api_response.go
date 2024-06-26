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

// checks if the BalanceAPIResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BalanceAPIResponse{}

// BalanceAPIResponse struct for BalanceAPIResponse
type BalanceAPIResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Body *InputBody `json:"body,omitempty"`
	Address *string `json:"address,omitempty"`
	Data *BalanceResponse `json:"data,omitempty"`
}

type _BalanceAPIResponse BalanceAPIResponse

// NewBalanceAPIResponse instantiates a new BalanceAPIResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalanceAPIResponse(success bool, message string) *BalanceAPIResponse {
	this := BalanceAPIResponse{}
	this.Success = success
	this.Message = message
	return &this
}

// NewBalanceAPIResponseWithDefaults instantiates a new BalanceAPIResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceAPIResponseWithDefaults() *BalanceAPIResponse {
	this := BalanceAPIResponse{}
	return &this
}

// GetSuccess returns the Success field value
func (o *BalanceAPIResponse) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *BalanceAPIResponse) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *BalanceAPIResponse) SetSuccess(v bool) {
	o.Success = v
}

// GetMessage returns the Message field value
func (o *BalanceAPIResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *BalanceAPIResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *BalanceAPIResponse) SetMessage(v string) {
	o.Message = v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *BalanceAPIResponse) GetBody() InputBody {
	if o == nil || IsNil(o.Body) {
		var ret InputBody
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceAPIResponse) GetBodyOk() (*InputBody, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *BalanceAPIResponse) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given InputBody and assigns it to the Body field.
func (o *BalanceAPIResponse) SetBody(v InputBody) {
	o.Body = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *BalanceAPIResponse) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceAPIResponse) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *BalanceAPIResponse) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *BalanceAPIResponse) SetAddress(v string) {
	o.Address = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BalanceAPIResponse) GetData() BalanceResponse {
	if o == nil || IsNil(o.Data) {
		var ret BalanceResponse
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceAPIResponse) GetDataOk() (*BalanceResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BalanceAPIResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given BalanceResponse and assigns it to the Data field.
func (o *BalanceAPIResponse) SetData(v BalanceResponse) {
	o.Data = &v
}

func (o BalanceAPIResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BalanceAPIResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	toSerialize["message"] = o.Message
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

func (o *BalanceAPIResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"success",
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

	varBalanceAPIResponse := _BalanceAPIResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBalanceAPIResponse)

	if err != nil {
		return err
	}

	*o = BalanceAPIResponse(varBalanceAPIResponse)

	return err
}

type NullableBalanceAPIResponse struct {
	value *BalanceAPIResponse
	isSet bool
}

func (v NullableBalanceAPIResponse) Get() *BalanceAPIResponse {
	return v.value
}

func (v *NullableBalanceAPIResponse) Set(val *BalanceAPIResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBalanceAPIResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBalanceAPIResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalanceAPIResponse(val *BalanceAPIResponse) *NullableBalanceAPIResponse {
	return &NullableBalanceAPIResponse{value: val, isSet: true}
}

func (v NullableBalanceAPIResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalanceAPIResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


