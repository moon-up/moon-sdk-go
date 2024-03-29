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

// checks if the BroadcastInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BroadcastInput{}

// BroadcastInput struct for BroadcastInput
type BroadcastInput struct {
	ChainId string `json:"chainId"`
	RawTransaction string `json:"rawTransaction"`
}

type _BroadcastInput BroadcastInput

// NewBroadcastInput instantiates a new BroadcastInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBroadcastInput(chainId string, rawTransaction string) *BroadcastInput {
	this := BroadcastInput{}
	this.ChainId = chainId
	this.RawTransaction = rawTransaction
	return &this
}

// NewBroadcastInputWithDefaults instantiates a new BroadcastInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBroadcastInputWithDefaults() *BroadcastInput {
	this := BroadcastInput{}
	return &this
}

// GetChainId returns the ChainId field value
func (o *BroadcastInput) GetChainId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value
// and a boolean to check if the value has been set.
func (o *BroadcastInput) GetChainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainId, true
}

// SetChainId sets field value
func (o *BroadcastInput) SetChainId(v string) {
	o.ChainId = v
}

// GetRawTransaction returns the RawTransaction field value
func (o *BroadcastInput) GetRawTransaction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RawTransaction
}

// GetRawTransactionOk returns a tuple with the RawTransaction field value
// and a boolean to check if the value has been set.
func (o *BroadcastInput) GetRawTransactionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RawTransaction, true
}

// SetRawTransaction sets field value
func (o *BroadcastInput) SetRawTransaction(v string) {
	o.RawTransaction = v
}

func (o BroadcastInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BroadcastInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chainId"] = o.ChainId
	toSerialize["rawTransaction"] = o.RawTransaction
	return toSerialize, nil
}

func (o *BroadcastInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chainId",
		"rawTransaction",
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

	varBroadcastInput := _BroadcastInput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBroadcastInput)

	if err != nil {
		return err
	}

	*o = BroadcastInput(varBroadcastInput)

	return err
}

type NullableBroadcastInput struct {
	value *BroadcastInput
	isSet bool
}

func (v NullableBroadcastInput) Get() *BroadcastInput {
	return v.value
}

func (v *NullableBroadcastInput) Set(val *BroadcastInput) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastInput) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastInput(val *BroadcastInput) *NullableBroadcastInput {
	return &NullableBroadcastInput{value: val, isSet: true}
}

func (v NullableBroadcastInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


