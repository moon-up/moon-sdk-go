# INFTTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionHash** | **string** |  | 
**Contract** | **string** |  | 
**LogIndex** | **string** |  | 
**TokenContractType** | **string** |  | 
**TokenName** | **string** |  | 
**TokenSymbol** | **string** |  | 
**Triggers** | Pointer to [**[]TriggerOutput**](TriggerOutput.md) |  | [optional] 
**Operator** | **NullableString** |  | 
**From** | **string** |  | 
**To** | **string** |  | 
**TokenId** | **string** |  | 
**Amount** | **string** |  | 

## Methods

### NewINFTTransfer

`func NewINFTTransfer(transactionHash string, contract string, logIndex string, tokenContractType string, tokenName string, tokenSymbol string, operator NullableString, from string, to string, tokenId string, amount string, ) *INFTTransfer`

NewINFTTransfer instantiates a new INFTTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewINFTTransferWithDefaults

`func NewINFTTransferWithDefaults() *INFTTransfer`

NewINFTTransferWithDefaults instantiates a new INFTTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionHash

`func (o *INFTTransfer) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *INFTTransfer) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *INFTTransfer) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetContract

`func (o *INFTTransfer) GetContract() string`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *INFTTransfer) GetContractOk() (*string, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *INFTTransfer) SetContract(v string)`

SetContract sets Contract field to given value.


### GetLogIndex

`func (o *INFTTransfer) GetLogIndex() string`

GetLogIndex returns the LogIndex field if non-nil, zero value otherwise.

### GetLogIndexOk

`func (o *INFTTransfer) GetLogIndexOk() (*string, bool)`

GetLogIndexOk returns a tuple with the LogIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogIndex

`func (o *INFTTransfer) SetLogIndex(v string)`

SetLogIndex sets LogIndex field to given value.


### GetTokenContractType

`func (o *INFTTransfer) GetTokenContractType() string`

GetTokenContractType returns the TokenContractType field if non-nil, zero value otherwise.

### GetTokenContractTypeOk

`func (o *INFTTransfer) GetTokenContractTypeOk() (*string, bool)`

GetTokenContractTypeOk returns a tuple with the TokenContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenContractType

`func (o *INFTTransfer) SetTokenContractType(v string)`

SetTokenContractType sets TokenContractType field to given value.


### GetTokenName

`func (o *INFTTransfer) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *INFTTransfer) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *INFTTransfer) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.


### GetTokenSymbol

`func (o *INFTTransfer) GetTokenSymbol() string`

GetTokenSymbol returns the TokenSymbol field if non-nil, zero value otherwise.

### GetTokenSymbolOk

`func (o *INFTTransfer) GetTokenSymbolOk() (*string, bool)`

GetTokenSymbolOk returns a tuple with the TokenSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSymbol

`func (o *INFTTransfer) SetTokenSymbol(v string)`

SetTokenSymbol sets TokenSymbol field to given value.


### GetTriggers

`func (o *INFTTransfer) GetTriggers() []TriggerOutput`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *INFTTransfer) GetTriggersOk() (*[]TriggerOutput, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *INFTTransfer) SetTriggers(v []TriggerOutput)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *INFTTransfer) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetOperator

`func (o *INFTTransfer) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *INFTTransfer) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *INFTTransfer) SetOperator(v string)`

SetOperator sets Operator field to given value.


### SetOperatorNil

`func (o *INFTTransfer) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *INFTTransfer) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetFrom

`func (o *INFTTransfer) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *INFTTransfer) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *INFTTransfer) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *INFTTransfer) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *INFTTransfer) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *INFTTransfer) SetTo(v string)`

SetTo sets To field to given value.


### GetTokenId

`func (o *INFTTransfer) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *INFTTransfer) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *INFTTransfer) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetAmount

`func (o *INFTTransfer) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *INFTTransfer) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *INFTTransfer) SetAmount(v string)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


