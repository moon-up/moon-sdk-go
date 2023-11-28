# SupportedAssetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | **string** |  | 
**Assets** | [**[]SupportedAssetResponseAssetsInner**](SupportedAssetResponseAssetsInner.md) |  | 

## Methods

### NewSupportedAssetResponse

`func NewSupportedAssetResponse(country string, assets []SupportedAssetResponseAssetsInner, ) *SupportedAssetResponse`

NewSupportedAssetResponse instantiates a new SupportedAssetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedAssetResponseWithDefaults

`func NewSupportedAssetResponseWithDefaults() *SupportedAssetResponse`

NewSupportedAssetResponseWithDefaults instantiates a new SupportedAssetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *SupportedAssetResponse) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SupportedAssetResponse) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SupportedAssetResponse) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetAssets

`func (o *SupportedAssetResponse) GetAssets() []SupportedAssetResponseAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *SupportedAssetResponse) GetAssetsOk() (*[]SupportedAssetResponseAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *SupportedAssetResponse) SetAssets(v []SupportedAssetResponseAssetsInner)`

SetAssets sets Assets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


