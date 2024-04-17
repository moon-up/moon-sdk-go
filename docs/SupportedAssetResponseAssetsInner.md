# SupportedAssetResponseAssetsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Crypto** | **[]string** |  | 
**PaymentMethods** | **[]string** |  | 
**Fiat** | **string** |  | 

## Methods

### NewSupportedAssetResponseAssetsInner

`func NewSupportedAssetResponseAssetsInner(crypto []string, paymentMethods []string, fiat string, ) *SupportedAssetResponseAssetsInner`

NewSupportedAssetResponseAssetsInner instantiates a new SupportedAssetResponseAssetsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedAssetResponseAssetsInnerWithDefaults

`func NewSupportedAssetResponseAssetsInnerWithDefaults() *SupportedAssetResponseAssetsInner`

NewSupportedAssetResponseAssetsInnerWithDefaults instantiates a new SupportedAssetResponseAssetsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrypto

`func (o *SupportedAssetResponseAssetsInner) GetCrypto() []string`

GetCrypto returns the Crypto field if non-nil, zero value otherwise.

### GetCryptoOk

`func (o *SupportedAssetResponseAssetsInner) GetCryptoOk() (*[]string, bool)`

GetCryptoOk returns a tuple with the Crypto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrypto

`func (o *SupportedAssetResponseAssetsInner) SetCrypto(v []string)`

SetCrypto sets Crypto field to given value.


### GetPaymentMethods

`func (o *SupportedAssetResponseAssetsInner) GetPaymentMethods() []string`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *SupportedAssetResponseAssetsInner) GetPaymentMethodsOk() (*[]string, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *SupportedAssetResponseAssetsInner) SetPaymentMethods(v []string)`

SetPaymentMethods sets PaymentMethods field to given value.


### GetFiat

`func (o *SupportedAssetResponseAssetsInner) GetFiat() string`

GetFiat returns the Fiat field if non-nil, zero value otherwise.

### GetFiatOk

`func (o *SupportedAssetResponseAssetsInner) GetFiatOk() (*string, bool)`

GetFiatOk returns a tuple with the Fiat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiat

`func (o *SupportedAssetResponseAssetsInner) SetFiat(v string)`

SetFiat sets Fiat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


