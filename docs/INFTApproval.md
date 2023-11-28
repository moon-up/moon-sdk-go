# INFTApproval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionHash** | **string** |  | 
**Contract** | **string** |  | 
**LogIndex** | **string** |  | 
**TokenContractType** | **string** |  | 
**TokenName** | **string** |  | 
**TokenSymbol** | **string** |  | 
**Account** | **string** |  | 
**Operator** | **string** |  | 
**ApprovedAll** | **bool** |  | 
**TokenId** | **NullableString** |  | 

## Methods

### NewINFTApproval

`func NewINFTApproval(transactionHash string, contract string, logIndex string, tokenContractType string, tokenName string, tokenSymbol string, account string, operator string, approvedAll bool, tokenId NullableString, ) *INFTApproval`

NewINFTApproval instantiates a new INFTApproval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewINFTApprovalWithDefaults

`func NewINFTApprovalWithDefaults() *INFTApproval`

NewINFTApprovalWithDefaults instantiates a new INFTApproval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionHash

`func (o *INFTApproval) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *INFTApproval) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *INFTApproval) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetContract

`func (o *INFTApproval) GetContract() string`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *INFTApproval) GetContractOk() (*string, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *INFTApproval) SetContract(v string)`

SetContract sets Contract field to given value.


### GetLogIndex

`func (o *INFTApproval) GetLogIndex() string`

GetLogIndex returns the LogIndex field if non-nil, zero value otherwise.

### GetLogIndexOk

`func (o *INFTApproval) GetLogIndexOk() (*string, bool)`

GetLogIndexOk returns a tuple with the LogIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogIndex

`func (o *INFTApproval) SetLogIndex(v string)`

SetLogIndex sets LogIndex field to given value.


### GetTokenContractType

`func (o *INFTApproval) GetTokenContractType() string`

GetTokenContractType returns the TokenContractType field if non-nil, zero value otherwise.

### GetTokenContractTypeOk

`func (o *INFTApproval) GetTokenContractTypeOk() (*string, bool)`

GetTokenContractTypeOk returns a tuple with the TokenContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenContractType

`func (o *INFTApproval) SetTokenContractType(v string)`

SetTokenContractType sets TokenContractType field to given value.


### GetTokenName

`func (o *INFTApproval) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *INFTApproval) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *INFTApproval) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.


### GetTokenSymbol

`func (o *INFTApproval) GetTokenSymbol() string`

GetTokenSymbol returns the TokenSymbol field if non-nil, zero value otherwise.

### GetTokenSymbolOk

`func (o *INFTApproval) GetTokenSymbolOk() (*string, bool)`

GetTokenSymbolOk returns a tuple with the TokenSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSymbol

`func (o *INFTApproval) SetTokenSymbol(v string)`

SetTokenSymbol sets TokenSymbol field to given value.


### GetAccount

`func (o *INFTApproval) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *INFTApproval) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *INFTApproval) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetOperator

`func (o *INFTApproval) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *INFTApproval) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *INFTApproval) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetApprovedAll

`func (o *INFTApproval) GetApprovedAll() bool`

GetApprovedAll returns the ApprovedAll field if non-nil, zero value otherwise.

### GetApprovedAllOk

`func (o *INFTApproval) GetApprovedAllOk() (*bool, bool)`

GetApprovedAllOk returns a tuple with the ApprovedAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAll

`func (o *INFTApproval) SetApprovedAll(v bool)`

SetApprovedAll sets ApprovedAll field to given value.


### GetTokenId

`func (o *INFTApproval) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *INFTApproval) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *INFTApproval) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### SetTokenIdNil

`func (o *INFTApproval) SetTokenIdNil(b bool)`

 SetTokenIdNil sets the value for TokenId to be an explicit nil

### UnsetTokenId
`func (o *INFTApproval) UnsetTokenId()`

UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


