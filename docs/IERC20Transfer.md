# IERC20Transfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionHash** | **string** |  | 
**Contract** | **string** |  | 
**LogIndex** | **string** |  | 
**From** | **string** |  | 
**To** | **string** |  | 
**Value** | **string** |  | 
**TokenDecimals** | **string** |  | 
**TokenName** | **string** |  | 
**TokenSymbol** | **string** |  | 
**ValueWithDecimals** | Pointer to **string** |  | [optional] 
**Triggers** | Pointer to [**[]TriggerOutput**](TriggerOutput.md) |  | [optional] 

## Methods

### NewIERC20Transfer

`func NewIERC20Transfer(transactionHash string, contract string, logIndex string, from string, to string, value string, tokenDecimals string, tokenName string, tokenSymbol string, ) *IERC20Transfer`

NewIERC20Transfer instantiates a new IERC20Transfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIERC20TransferWithDefaults

`func NewIERC20TransferWithDefaults() *IERC20Transfer`

NewIERC20TransferWithDefaults instantiates a new IERC20Transfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionHash

`func (o *IERC20Transfer) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *IERC20Transfer) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *IERC20Transfer) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetContract

`func (o *IERC20Transfer) GetContract() string`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *IERC20Transfer) GetContractOk() (*string, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *IERC20Transfer) SetContract(v string)`

SetContract sets Contract field to given value.


### GetLogIndex

`func (o *IERC20Transfer) GetLogIndex() string`

GetLogIndex returns the LogIndex field if non-nil, zero value otherwise.

### GetLogIndexOk

`func (o *IERC20Transfer) GetLogIndexOk() (*string, bool)`

GetLogIndexOk returns a tuple with the LogIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogIndex

`func (o *IERC20Transfer) SetLogIndex(v string)`

SetLogIndex sets LogIndex field to given value.


### GetFrom

`func (o *IERC20Transfer) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *IERC20Transfer) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *IERC20Transfer) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *IERC20Transfer) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *IERC20Transfer) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *IERC20Transfer) SetTo(v string)`

SetTo sets To field to given value.


### GetValue

`func (o *IERC20Transfer) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IERC20Transfer) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IERC20Transfer) SetValue(v string)`

SetValue sets Value field to given value.


### GetTokenDecimals

`func (o *IERC20Transfer) GetTokenDecimals() string`

GetTokenDecimals returns the TokenDecimals field if non-nil, zero value otherwise.

### GetTokenDecimalsOk

`func (o *IERC20Transfer) GetTokenDecimalsOk() (*string, bool)`

GetTokenDecimalsOk returns a tuple with the TokenDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenDecimals

`func (o *IERC20Transfer) SetTokenDecimals(v string)`

SetTokenDecimals sets TokenDecimals field to given value.


### GetTokenName

`func (o *IERC20Transfer) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *IERC20Transfer) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *IERC20Transfer) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.


### GetTokenSymbol

`func (o *IERC20Transfer) GetTokenSymbol() string`

GetTokenSymbol returns the TokenSymbol field if non-nil, zero value otherwise.

### GetTokenSymbolOk

`func (o *IERC20Transfer) GetTokenSymbolOk() (*string, bool)`

GetTokenSymbolOk returns a tuple with the TokenSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSymbol

`func (o *IERC20Transfer) SetTokenSymbol(v string)`

SetTokenSymbol sets TokenSymbol field to given value.


### GetValueWithDecimals

`func (o *IERC20Transfer) GetValueWithDecimals() string`

GetValueWithDecimals returns the ValueWithDecimals field if non-nil, zero value otherwise.

### GetValueWithDecimalsOk

`func (o *IERC20Transfer) GetValueWithDecimalsOk() (*string, bool)`

GetValueWithDecimalsOk returns a tuple with the ValueWithDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueWithDecimals

`func (o *IERC20Transfer) SetValueWithDecimals(v string)`

SetValueWithDecimals sets ValueWithDecimals field to given value.

### HasValueWithDecimals

`func (o *IERC20Transfer) HasValueWithDecimals() bool`

HasValueWithDecimals returns a boolean if a field has been set.

### GetTriggers

`func (o *IERC20Transfer) GetTriggers() []TriggerOutput`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *IERC20Transfer) GetTriggersOk() (*[]TriggerOutput, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *IERC20Transfer) SetTriggers(v []TriggerOutput)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *IERC20Transfer) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


