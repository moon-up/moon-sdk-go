# TransactionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoonScanUrl** | Pointer to **string** |  | [optional] 
**TransactionHash** | **string** |  | 
**SignedTransaction** | **string** |  | 
**SignedMessage** | Pointer to **string** |  | [optional] 
**RawTransaction** | Pointer to **string** |  | [optional] 
**Signature** | Pointer to **string** |  | [optional] 
**Transaction** | Pointer to [**Tx**](Tx.md) |  | [optional] 
**UserOps** | Pointer to [**[]TransactionRequest**](TransactionRequest.md) |  | [optional] 
**UseropTransaction** | Pointer to **string** |  | [optional] 

## Methods

### NewTransactionData

`func NewTransactionData(transactionHash string, signedTransaction string, ) *TransactionData`

NewTransactionData instantiates a new TransactionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionDataWithDefaults

`func NewTransactionDataWithDefaults() *TransactionData`

NewTransactionDataWithDefaults instantiates a new TransactionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoonScanUrl

`func (o *TransactionData) GetMoonScanUrl() string`

GetMoonScanUrl returns the MoonScanUrl field if non-nil, zero value otherwise.

### GetMoonScanUrlOk

`func (o *TransactionData) GetMoonScanUrlOk() (*string, bool)`

GetMoonScanUrlOk returns a tuple with the MoonScanUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoonScanUrl

`func (o *TransactionData) SetMoonScanUrl(v string)`

SetMoonScanUrl sets MoonScanUrl field to given value.

### HasMoonScanUrl

`func (o *TransactionData) HasMoonScanUrl() bool`

HasMoonScanUrl returns a boolean if a field has been set.

### GetTransactionHash

`func (o *TransactionData) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *TransactionData) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *TransactionData) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetSignedTransaction

`func (o *TransactionData) GetSignedTransaction() string`

GetSignedTransaction returns the SignedTransaction field if non-nil, zero value otherwise.

### GetSignedTransactionOk

`func (o *TransactionData) GetSignedTransactionOk() (*string, bool)`

GetSignedTransactionOk returns a tuple with the SignedTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedTransaction

`func (o *TransactionData) SetSignedTransaction(v string)`

SetSignedTransaction sets SignedTransaction field to given value.


### GetSignedMessage

`func (o *TransactionData) GetSignedMessage() string`

GetSignedMessage returns the SignedMessage field if non-nil, zero value otherwise.

### GetSignedMessageOk

`func (o *TransactionData) GetSignedMessageOk() (*string, bool)`

GetSignedMessageOk returns a tuple with the SignedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedMessage

`func (o *TransactionData) SetSignedMessage(v string)`

SetSignedMessage sets SignedMessage field to given value.

### HasSignedMessage

`func (o *TransactionData) HasSignedMessage() bool`

HasSignedMessage returns a boolean if a field has been set.

### GetRawTransaction

`func (o *TransactionData) GetRawTransaction() string`

GetRawTransaction returns the RawTransaction field if non-nil, zero value otherwise.

### GetRawTransactionOk

`func (o *TransactionData) GetRawTransactionOk() (*string, bool)`

GetRawTransactionOk returns a tuple with the RawTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawTransaction

`func (o *TransactionData) SetRawTransaction(v string)`

SetRawTransaction sets RawTransaction field to given value.

### HasRawTransaction

`func (o *TransactionData) HasRawTransaction() bool`

HasRawTransaction returns a boolean if a field has been set.

### GetSignature

`func (o *TransactionData) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *TransactionData) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *TransactionData) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *TransactionData) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetTransaction

`func (o *TransactionData) GetTransaction() Tx`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *TransactionData) GetTransactionOk() (*Tx, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *TransactionData) SetTransaction(v Tx)`

SetTransaction sets Transaction field to given value.

### HasTransaction

`func (o *TransactionData) HasTransaction() bool`

HasTransaction returns a boolean if a field has been set.

### GetUserOps

`func (o *TransactionData) GetUserOps() []TransactionRequest`

GetUserOps returns the UserOps field if non-nil, zero value otherwise.

### GetUserOpsOk

`func (o *TransactionData) GetUserOpsOk() (*[]TransactionRequest, bool)`

GetUserOpsOk returns a tuple with the UserOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserOps

`func (o *TransactionData) SetUserOps(v []TransactionRequest)`

SetUserOps sets UserOps field to given value.

### HasUserOps

`func (o *TransactionData) HasUserOps() bool`

HasUserOps returns a boolean if a field has been set.

### GetUseropTransaction

`func (o *TransactionData) GetUseropTransaction() string`

GetUseropTransaction returns the UseropTransaction field if non-nil, zero value otherwise.

### GetUseropTransactionOk

`func (o *TransactionData) GetUseropTransactionOk() (*string, bool)`

GetUseropTransactionOk returns a tuple with the UseropTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseropTransaction

`func (o *TransactionData) SetUseropTransaction(v string)`

SetUseropTransaction sets UseropTransaction field to given value.

### HasUseropTransaction

`func (o *TransactionData) HasUseropTransaction() bool`

HasUseropTransaction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


