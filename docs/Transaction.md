# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionHash** | Pointer to **string** |  | [optional] 
**SignedTransaction** | Pointer to **string** |  | [optional] 
**RawTransaction** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **NullableString** |  | [optional] 
**Transactions** | Pointer to [**[]TransactionData**](TransactionData.md) |  | [optional] 
**MoonScanUrl** | Pointer to **string** |  | [optional] 
**Signature** | Pointer to **string** |  | [optional] 
**Transaction** | Pointer to [**Tx**](Tx.md) |  | [optional] 
**UserOps** | Pointer to [**[]TransactionRequest**](TransactionRequest.md) |  | [optional] 
**UseropTransaction** | Pointer to **string** |  | [optional] 

## Methods

### NewTransaction

`func NewTransaction() *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionHash

`func (o *Transaction) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *Transaction) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *Transaction) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.

### HasTransactionHash

`func (o *Transaction) HasTransactionHash() bool`

HasTransactionHash returns a boolean if a field has been set.

### GetSignedTransaction

`func (o *Transaction) GetSignedTransaction() string`

GetSignedTransaction returns the SignedTransaction field if non-nil, zero value otherwise.

### GetSignedTransactionOk

`func (o *Transaction) GetSignedTransactionOk() (*string, bool)`

GetSignedTransactionOk returns a tuple with the SignedTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedTransaction

`func (o *Transaction) SetSignedTransaction(v string)`

SetSignedTransaction sets SignedTransaction field to given value.

### HasSignedTransaction

`func (o *Transaction) HasSignedTransaction() bool`

HasSignedTransaction returns a boolean if a field has been set.

### GetRawTransaction

`func (o *Transaction) GetRawTransaction() string`

GetRawTransaction returns the RawTransaction field if non-nil, zero value otherwise.

### GetRawTransactionOk

`func (o *Transaction) GetRawTransactionOk() (*string, bool)`

GetRawTransactionOk returns a tuple with the RawTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawTransaction

`func (o *Transaction) SetRawTransaction(v string)`

SetRawTransaction sets RawTransaction field to given value.

### HasRawTransaction

`func (o *Transaction) HasRawTransaction() bool`

HasRawTransaction returns a boolean if a field has been set.

### GetData

`func (o *Transaction) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Transaction) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Transaction) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *Transaction) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *Transaction) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *Transaction) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTransactions

`func (o *Transaction) GetTransactions() []TransactionData`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *Transaction) GetTransactionsOk() (*[]TransactionData, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *Transaction) SetTransactions(v []TransactionData)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *Transaction) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetMoonScanUrl

`func (o *Transaction) GetMoonScanUrl() string`

GetMoonScanUrl returns the MoonScanUrl field if non-nil, zero value otherwise.

### GetMoonScanUrlOk

`func (o *Transaction) GetMoonScanUrlOk() (*string, bool)`

GetMoonScanUrlOk returns a tuple with the MoonScanUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoonScanUrl

`func (o *Transaction) SetMoonScanUrl(v string)`

SetMoonScanUrl sets MoonScanUrl field to given value.

### HasMoonScanUrl

`func (o *Transaction) HasMoonScanUrl() bool`

HasMoonScanUrl returns a boolean if a field has been set.

### GetSignature

`func (o *Transaction) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *Transaction) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *Transaction) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *Transaction) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetTransaction

`func (o *Transaction) GetTransaction() Tx`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *Transaction) GetTransactionOk() (*Tx, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *Transaction) SetTransaction(v Tx)`

SetTransaction sets Transaction field to given value.

### HasTransaction

`func (o *Transaction) HasTransaction() bool`

HasTransaction returns a boolean if a field has been set.

### GetUserOps

`func (o *Transaction) GetUserOps() []TransactionRequest`

GetUserOps returns the UserOps field if non-nil, zero value otherwise.

### GetUserOpsOk

`func (o *Transaction) GetUserOpsOk() (*[]TransactionRequest, bool)`

GetUserOpsOk returns a tuple with the UserOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserOps

`func (o *Transaction) SetUserOps(v []TransactionRequest)`

SetUserOps sets UserOps field to given value.

### HasUserOps

`func (o *Transaction) HasUserOps() bool`

HasUserOps returns a boolean if a field has been set.

### GetUseropTransaction

`func (o *Transaction) GetUseropTransaction() string`

GetUseropTransaction returns the UseropTransaction field if non-nil, zero value otherwise.

### GetUseropTransactionOk

`func (o *Transaction) GetUseropTransactionOk() (*string, bool)`

GetUseropTransactionOk returns a tuple with the UseropTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseropTransaction

`func (o *Transaction) SetUseropTransaction(v string)`

SetUseropTransaction sets UseropTransaction field to given value.

### HasUseropTransaction

`func (o *Transaction) HasUseropTransaction() bool`

HasUseropTransaction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


