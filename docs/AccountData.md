# AccountData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | Pointer to **[]string** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountData

`func NewAccountData() *AccountData`

NewAccountData instantiates a new AccountData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDataWithDefaults

`func NewAccountDataWithDefaults() *AccountData`

NewAccountDataWithDefaults instantiates a new AccountData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *AccountData) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *AccountData) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *AccountData) SetKeys(v []string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *AccountData) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetAddress

`func (o *AccountData) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AccountData) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AccountData) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AccountData) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


