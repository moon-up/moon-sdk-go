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

// checks if the Quote type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Quote{}

// Quote struct for Quote
type Quote struct {
	Recommendations []string `json:"recommendations"`
	PaymentMethod string `json:"paymentMethod"`
	QuoteId string `json:"quoteId"`
	Ramp string `json:"ramp"`
	AvailablePaymentMethods []AvailablePaymentMethod `json:"availablePaymentMethods"`
	Payout float64 `json:"payout"`
	TransactionFee float64 `json:"transactionFee"`
	NetworkFee float64 `json:"networkFee"`
	Rate float64 `json:"rate"`
}

type _Quote Quote

// NewQuote instantiates a new Quote object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuote(recommendations []string, paymentMethod string, quoteId string, ramp string, availablePaymentMethods []AvailablePaymentMethod, payout float64, transactionFee float64, networkFee float64, rate float64) *Quote {
	this := Quote{}
	this.Recommendations = recommendations
	this.PaymentMethod = paymentMethod
	this.QuoteId = quoteId
	this.Ramp = ramp
	this.AvailablePaymentMethods = availablePaymentMethods
	this.Payout = payout
	this.TransactionFee = transactionFee
	this.NetworkFee = networkFee
	this.Rate = rate
	return &this
}

// NewQuoteWithDefaults instantiates a new Quote object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteWithDefaults() *Quote {
	this := Quote{}
	return &this
}

// GetRecommendations returns the Recommendations field value
func (o *Quote) GetRecommendations() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Recommendations
}

// GetRecommendationsOk returns a tuple with the Recommendations field value
// and a boolean to check if the value has been set.
func (o *Quote) GetRecommendationsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recommendations, true
}

// SetRecommendations sets field value
func (o *Quote) SetRecommendations(v []string) {
	o.Recommendations = v
}

// GetPaymentMethod returns the PaymentMethod field value
func (o *Quote) GetPaymentMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value
// and a boolean to check if the value has been set.
func (o *Quote) GetPaymentMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentMethod, true
}

// SetPaymentMethod sets field value
func (o *Quote) SetPaymentMethod(v string) {
	o.PaymentMethod = v
}

// GetQuoteId returns the QuoteId field value
func (o *Quote) GetQuoteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QuoteId
}

// GetQuoteIdOk returns a tuple with the QuoteId field value
// and a boolean to check if the value has been set.
func (o *Quote) GetQuoteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuoteId, true
}

// SetQuoteId sets field value
func (o *Quote) SetQuoteId(v string) {
	o.QuoteId = v
}

// GetRamp returns the Ramp field value
func (o *Quote) GetRamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ramp
}

// GetRampOk returns a tuple with the Ramp field value
// and a boolean to check if the value has been set.
func (o *Quote) GetRampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ramp, true
}

// SetRamp sets field value
func (o *Quote) SetRamp(v string) {
	o.Ramp = v
}

// GetAvailablePaymentMethods returns the AvailablePaymentMethods field value
func (o *Quote) GetAvailablePaymentMethods() []AvailablePaymentMethod {
	if o == nil {
		var ret []AvailablePaymentMethod
		return ret
	}

	return o.AvailablePaymentMethods
}

// GetAvailablePaymentMethodsOk returns a tuple with the AvailablePaymentMethods field value
// and a boolean to check if the value has been set.
func (o *Quote) GetAvailablePaymentMethodsOk() ([]AvailablePaymentMethod, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailablePaymentMethods, true
}

// SetAvailablePaymentMethods sets field value
func (o *Quote) SetAvailablePaymentMethods(v []AvailablePaymentMethod) {
	o.AvailablePaymentMethods = v
}

// GetPayout returns the Payout field value
func (o *Quote) GetPayout() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Payout
}

// GetPayoutOk returns a tuple with the Payout field value
// and a boolean to check if the value has been set.
func (o *Quote) GetPayoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payout, true
}

// SetPayout sets field value
func (o *Quote) SetPayout(v float64) {
	o.Payout = v
}

// GetTransactionFee returns the TransactionFee field value
func (o *Quote) GetTransactionFee() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.TransactionFee
}

// GetTransactionFeeOk returns a tuple with the TransactionFee field value
// and a boolean to check if the value has been set.
func (o *Quote) GetTransactionFeeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionFee, true
}

// SetTransactionFee sets field value
func (o *Quote) SetTransactionFee(v float64) {
	o.TransactionFee = v
}

// GetNetworkFee returns the NetworkFee field value
func (o *Quote) GetNetworkFee() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.NetworkFee
}

// GetNetworkFeeOk returns a tuple with the NetworkFee field value
// and a boolean to check if the value has been set.
func (o *Quote) GetNetworkFeeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkFee, true
}

// SetNetworkFee sets field value
func (o *Quote) SetNetworkFee(v float64) {
	o.NetworkFee = v
}

// GetRate returns the Rate field value
func (o *Quote) GetRate() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value
// and a boolean to check if the value has been set.
func (o *Quote) GetRateOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rate, true
}

// SetRate sets field value
func (o *Quote) SetRate(v float64) {
	o.Rate = v
}

func (o Quote) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Quote) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["recommendations"] = o.Recommendations
	toSerialize["paymentMethod"] = o.PaymentMethod
	toSerialize["quoteId"] = o.QuoteId
	toSerialize["ramp"] = o.Ramp
	toSerialize["availablePaymentMethods"] = o.AvailablePaymentMethods
	toSerialize["payout"] = o.Payout
	toSerialize["transactionFee"] = o.TransactionFee
	toSerialize["networkFee"] = o.NetworkFee
	toSerialize["rate"] = o.Rate
	return toSerialize, nil
}

func (o *Quote) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"recommendations",
		"paymentMethod",
		"quoteId",
		"ramp",
		"availablePaymentMethods",
		"payout",
		"transactionFee",
		"networkFee",
		"rate",
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

	varQuote := _Quote{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQuote)

	if err != nil {
		return err
	}

	*o = Quote(varQuote)

	return err
}

type NullableQuote struct {
	value *Quote
	isSet bool
}

func (v NullableQuote) Get() *Quote {
	return v.value
}

func (v *NullableQuote) Set(val *Quote) {
	v.value = val
	v.isSet = true
}

func (v NullableQuote) IsSet() bool {
	return v.isSet
}

func (v *NullableQuote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuote(val *Quote) *NullableQuote {
	return &NullableQuote{value: val, isSet: true}
}

func (v NullableQuote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


