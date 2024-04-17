# Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fiat** | [**[]FiatCurrency**](FiatCurrency.md) |  | 
**Crypto** | [**[]CryptoCurrency**](CryptoCurrency.md) |  | 

## Methods

### NewMessage

`func NewMessage(fiat []FiatCurrency, crypto []CryptoCurrency, ) *Message`

NewMessage instantiates a new Message object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageWithDefaults

`func NewMessageWithDefaults() *Message`

NewMessageWithDefaults instantiates a new Message object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiat

`func (o *Message) GetFiat() []FiatCurrency`

GetFiat returns the Fiat field if non-nil, zero value otherwise.

### GetFiatOk

`func (o *Message) GetFiatOk() (*[]FiatCurrency, bool)`

GetFiatOk returns a tuple with the Fiat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiat

`func (o *Message) SetFiat(v []FiatCurrency)`

SetFiat sets Fiat field to given value.


### GetCrypto

`func (o *Message) GetCrypto() []CryptoCurrency`

GetCrypto returns the Crypto field if non-nil, zero value otherwise.

### GetCryptoOk

`func (o *Message) GetCryptoOk() (*[]CryptoCurrency, bool)`

GetCryptoOk returns a tuple with the Crypto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrypto

`func (o *Message) SetCrypto(v []CryptoCurrency)`

SetCrypto sets Crypto field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


