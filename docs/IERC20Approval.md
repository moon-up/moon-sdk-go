# IERC20Approval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionHash** | **string** |  | 
**Contract** | **string** |  | 
**LogIndex** | **string** |  | 
**Owner** | **string** |  | 
**Spender** | **string** |  | 
**Value** | **string** |  | 
**TokenDecimals** | **string** |  | 
**TokenName** | **string** |  | 
**TokenSymbol** | **string** |  | 
**ValueWithDecimals** | Pointer to **string** |  | [optional] 
**Triggers** | Pointer to [**[]TriggerOutput**](TriggerOutput.md) |  | [optional] 

## Methods

### NewIERC20Approval

`func NewIERC20Approval(transactionHash string, contract string, logIndex string, owner string, spender string, value string, tokenDecimals string, tokenName string, tokenSymbol string, ) *IERC20Approval`

NewIERC20Approval instantiates a new IERC20Approval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIERC20ApprovalWithDefaults

`func NewIERC20ApprovalWithDefaults() *IERC20Approval`

NewIERC20ApprovalWithDefaults instantiates a new IERC20Approval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionHash

`func (o *IERC20Approval) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *IERC20Approval) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *IERC20Approval) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetContract

`func (o *IERC20Approval) GetContract() string`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *IERC20Approval) GetContractOk() (*string, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *IERC20Approval) SetContract(v string)`

SetContract sets Contract field to given value.


### GetLogIndex

`func (o *IERC20Approval) GetLogIndex() string`

GetLogIndex returns the LogIndex field if non-nil, zero value otherwise.

### GetLogIndexOk

`func (o *IERC20Approval) GetLogIndexOk() (*string, bool)`

GetLogIndexOk returns a tuple with the LogIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogIndex

`func (o *IERC20Approval) SetLogIndex(v string)`

SetLogIndex sets LogIndex field to given value.


### GetOwner

`func (o *IERC20Approval) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IERC20Approval) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IERC20Approval) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetSpender

`func (o *IERC20Approval) GetSpender() string`

GetSpender returns the Spender field if non-nil, zero value otherwise.

### GetSpenderOk

`func (o *IERC20Approval) GetSpenderOk() (*string, bool)`

GetSpenderOk returns a tuple with the Spender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpender

`func (o *IERC20Approval) SetSpender(v string)`

SetSpender sets Spender field to given value.


### GetValue

`func (o *IERC20Approval) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IERC20Approval) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IERC20Approval) SetValue(v string)`

SetValue sets Value field to given value.


### GetTokenDecimals

`func (o *IERC20Approval) GetTokenDecimals() string`

GetTokenDecimals returns the TokenDecimals field if non-nil, zero value otherwise.

### GetTokenDecimalsOk

`func (o *IERC20Approval) GetTokenDecimalsOk() (*string, bool)`

GetTokenDecimalsOk returns a tuple with the TokenDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenDecimals

`func (o *IERC20Approval) SetTokenDecimals(v string)`

SetTokenDecimals sets TokenDecimals field to given value.


### GetTokenName

`func (o *IERC20Approval) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *IERC20Approval) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *IERC20Approval) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.


### GetTokenSymbol

`func (o *IERC20Approval) GetTokenSymbol() string`

GetTokenSymbol returns the TokenSymbol field if non-nil, zero value otherwise.

### GetTokenSymbolOk

`func (o *IERC20Approval) GetTokenSymbolOk() (*string, bool)`

GetTokenSymbolOk returns a tuple with the TokenSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSymbol

`func (o *IERC20Approval) SetTokenSymbol(v string)`

SetTokenSymbol sets TokenSymbol field to given value.


### GetValueWithDecimals

`func (o *IERC20Approval) GetValueWithDecimals() string`

GetValueWithDecimals returns the ValueWithDecimals field if non-nil, zero value otherwise.

### GetValueWithDecimalsOk

`func (o *IERC20Approval) GetValueWithDecimalsOk() (*string, bool)`

GetValueWithDecimalsOk returns a tuple with the ValueWithDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueWithDecimals

`func (o *IERC20Approval) SetValueWithDecimals(v string)`

SetValueWithDecimals sets ValueWithDecimals field to given value.

### HasValueWithDecimals

`func (o *IERC20Approval) HasValueWithDecimals() bool`

HasValueWithDecimals returns a boolean if a field has been set.

### GetTriggers

`func (o *IERC20Approval) GetTriggers() []TriggerOutput`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *IERC20Approval) GetTriggersOk() (*[]TriggerOutput, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *IERC20Approval) SetTriggers(v []TriggerOutput)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *IERC20Approval) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


