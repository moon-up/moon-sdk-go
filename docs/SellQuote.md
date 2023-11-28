# SellQuote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recommendations** | **[]string** |  | 
**QuoteId** | **string** |  | 
**PaymentMethod** | **string** |  | 
**Ramp** | **string** |  | 
**Payout** | **float64** |  | 
**TransactionFee** | **float64** |  | 
**NetworkFee** | **float64** |  | 
**Rate** | **float64** |  | 

## Methods

### NewSellQuote

`func NewSellQuote(recommendations []string, quoteId string, paymentMethod string, ramp string, payout float64, transactionFee float64, networkFee float64, rate float64, ) *SellQuote`

NewSellQuote instantiates a new SellQuote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSellQuoteWithDefaults

`func NewSellQuoteWithDefaults() *SellQuote`

NewSellQuoteWithDefaults instantiates a new SellQuote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecommendations

`func (o *SellQuote) GetRecommendations() []string`

GetRecommendations returns the Recommendations field if non-nil, zero value otherwise.

### GetRecommendationsOk

`func (o *SellQuote) GetRecommendationsOk() (*[]string, bool)`

GetRecommendationsOk returns a tuple with the Recommendations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendations

`func (o *SellQuote) SetRecommendations(v []string)`

SetRecommendations sets Recommendations field to given value.


### GetQuoteId

`func (o *SellQuote) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *SellQuote) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *SellQuote) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetPaymentMethod

`func (o *SellQuote) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *SellQuote) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *SellQuote) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetRamp

`func (o *SellQuote) GetRamp() string`

GetRamp returns the Ramp field if non-nil, zero value otherwise.

### GetRampOk

`func (o *SellQuote) GetRampOk() (*string, bool)`

GetRampOk returns a tuple with the Ramp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamp

`func (o *SellQuote) SetRamp(v string)`

SetRamp sets Ramp field to given value.


### GetPayout

`func (o *SellQuote) GetPayout() float64`

GetPayout returns the Payout field if non-nil, zero value otherwise.

### GetPayoutOk

`func (o *SellQuote) GetPayoutOk() (*float64, bool)`

GetPayoutOk returns a tuple with the Payout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayout

`func (o *SellQuote) SetPayout(v float64)`

SetPayout sets Payout field to given value.


### GetTransactionFee

`func (o *SellQuote) GetTransactionFee() float64`

GetTransactionFee returns the TransactionFee field if non-nil, zero value otherwise.

### GetTransactionFeeOk

`func (o *SellQuote) GetTransactionFeeOk() (*float64, bool)`

GetTransactionFeeOk returns a tuple with the TransactionFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionFee

`func (o *SellQuote) SetTransactionFee(v float64)`

SetTransactionFee sets TransactionFee field to given value.


### GetNetworkFee

`func (o *SellQuote) GetNetworkFee() float64`

GetNetworkFee returns the NetworkFee field if non-nil, zero value otherwise.

### GetNetworkFeeOk

`func (o *SellQuote) GetNetworkFeeOk() (*float64, bool)`

GetNetworkFeeOk returns a tuple with the NetworkFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFee

`func (o *SellQuote) SetNetworkFee(v float64)`

SetNetworkFee sets NetworkFee field to given value.


### GetRate

`func (o *SellQuote) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *SellQuote) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *SellQuote) SetRate(v float64)`

SetRate sets Rate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


