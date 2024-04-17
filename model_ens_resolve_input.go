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

// checks if the EnsResolveInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnsResolveInput{}

// EnsResolveInput struct for EnsResolveInput
type EnsResolveInput struct {
	Domain string `json:"domain"`
	ChainId string `json:"chain_id"`
}

type _EnsResolveInput EnsResolveInput

// NewEnsResolveInput instantiates a new EnsResolveInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnsResolveInput(domain string, chainId string) *EnsResolveInput {
	this := EnsResolveInput{}
	this.Domain = domain
	this.ChainId = chainId
	return &this
}

// NewEnsResolveInputWithDefaults instantiates a new EnsResolveInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnsResolveInputWithDefaults() *EnsResolveInput {
	this := EnsResolveInput{}
	return &this
}

// GetDomain returns the Domain field value
func (o *EnsResolveInput) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *EnsResolveInput) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *EnsResolveInput) SetDomain(v string) {
	o.Domain = v
}

// GetChainId returns the ChainId field value
func (o *EnsResolveInput) GetChainId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value
// and a boolean to check if the value has been set.
func (o *EnsResolveInput) GetChainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainId, true
}

// SetChainId sets field value
func (o *EnsResolveInput) SetChainId(v string) {
	o.ChainId = v
}

func (o EnsResolveInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnsResolveInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["domain"] = o.Domain
	toSerialize["chain_id"] = o.ChainId
	return toSerialize, nil
}

func (o *EnsResolveInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"domain",
		"chain_id",
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

	varEnsResolveInput := _EnsResolveInput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEnsResolveInput)

	if err != nil {
		return err
	}

	*o = EnsResolveInput(varEnsResolveInput)

	return err
}

type NullableEnsResolveInput struct {
	value *EnsResolveInput
	isSet bool
}

func (v NullableEnsResolveInput) Get() *EnsResolveInput {
	return v.value
}

func (v *NullableEnsResolveInput) Set(val *EnsResolveInput) {
	v.value = val
	v.isSet = true
}

func (v NullableEnsResolveInput) IsSet() bool {
	return v.isSet
}

func (v *NullableEnsResolveInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnsResolveInput(val *EnsResolveInput) *NullableEnsResolveInput {
	return &NullableEnsResolveInput{value: val, isSet: true}
}

func (v NullableEnsResolveInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnsResolveInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

