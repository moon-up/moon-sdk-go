# SupportedPaymentTypesMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Googlepay** | [**PaymentType**](PaymentType.md) |  | 
**Applepay** | [**PaymentType**](PaymentType.md) |  | 
**Creditcard** | [**PaymentType**](PaymentType.md) |  | 

## Methods

### NewSupportedPaymentTypesMessage

`func NewSupportedPaymentTypesMessage(googlepay PaymentType, applepay PaymentType, creditcard PaymentType, ) *SupportedPaymentTypesMessage`

NewSupportedPaymentTypesMessage instantiates a new SupportedPaymentTypesMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedPaymentTypesMessageWithDefaults

`func NewSupportedPaymentTypesMessageWithDefaults() *SupportedPaymentTypesMessage`

NewSupportedPaymentTypesMessageWithDefaults instantiates a new SupportedPaymentTypesMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGooglepay

`func (o *SupportedPaymentTypesMessage) GetGooglepay() PaymentType`

GetGooglepay returns the Googlepay field if non-nil, zero value otherwise.

### GetGooglepayOk

`func (o *SupportedPaymentTypesMessage) GetGooglepayOk() (*PaymentType, bool)`

GetGooglepayOk returns a tuple with the Googlepay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGooglepay

`func (o *SupportedPaymentTypesMessage) SetGooglepay(v PaymentType)`

SetGooglepay sets Googlepay field to given value.


### GetApplepay

`func (o *SupportedPaymentTypesMessage) GetApplepay() PaymentType`

GetApplepay returns the Applepay field if non-nil, zero value otherwise.

### GetApplepayOk

`func (o *SupportedPaymentTypesMessage) GetApplepayOk() (*PaymentType, bool)`

GetApplepayOk returns a tuple with the Applepay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplepay

`func (o *SupportedPaymentTypesMessage) SetApplepay(v PaymentType)`

SetApplepay sets Applepay field to given value.


### GetCreditcard

`func (o *SupportedPaymentTypesMessage) GetCreditcard() PaymentType`

GetCreditcard returns the Creditcard field if non-nil, zero value otherwise.

### GetCreditcardOk

`func (o *SupportedPaymentTypesMessage) GetCreditcardOk() (*PaymentType, bool)`

GetCreditcardOk returns a tuple with the Creditcard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditcard

`func (o *SupportedPaymentTypesMessage) SetCreditcard(v PaymentType)`

SetCreditcard sets Creditcard field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


