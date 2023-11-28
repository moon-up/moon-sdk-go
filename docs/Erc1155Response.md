# Erc1155Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoonScanUrl** | Pointer to **string** |  | [optional] 
**TransactionHash** | **string** |  | 
**SignedTransaction** | **string** |  | 
**SignedMessage** | Pointer to **string** |  | [optional] 
**RawTransaction** | Pointer to **string** |  | [optional] 
**Signature** | Pointer to **string** |  | [optional] 
**Transaction** | Pointer to [**map[string]Tx**](Tx.md) |  | [optional] 
**UserOps** | Pointer to [**[]TransactionRequest**](TransactionRequest.md) |  | [optional] 
**UseropTransaction** | Pointer to **string** |  | [optional] 
**BalanceOf** | Pointer to **string** |  | [optional] 
**BalanceOfBatch** | Pointer to **string** |  | [optional] 

## Methods

### NewErc1155Response

`func NewErc1155Response(transactionHash string, signedTransaction string, ) *Erc1155Response`

NewErc1155Response instantiates a new Erc1155Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErc1155ResponseWithDefaults

`func NewErc1155ResponseWithDefaults() *Erc1155Response`

NewErc1155ResponseWithDefaults instantiates a new Erc1155Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoonScanUrl

`func (o *Erc1155Response) GetMoonScanUrl() string`

GetMoonScanUrl returns the MoonScanUrl field if non-nil, zero value otherwise.

### GetMoonScanUrlOk

`func (o *Erc1155Response) GetMoonScanUrlOk() (*string, bool)`

GetMoonScanUrlOk returns a tuple with the MoonScanUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoonScanUrl

`func (o *Erc1155Response) SetMoonScanUrl(v string)`

SetMoonScanUrl sets MoonScanUrl field to given value.

### HasMoonScanUrl

`func (o *Erc1155Response) HasMoonScanUrl() bool`

HasMoonScanUrl returns a boolean if a field has been set.

### GetTransactionHash

`func (o *Erc1155Response) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *Erc1155Response) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *Erc1155Response) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetSignedTransaction

`func (o *Erc1155Response) GetSignedTransaction() string`

GetSignedTransaction returns the SignedTransaction field if non-nil, zero value otherwise.

### GetSignedTransactionOk

`func (o *Erc1155Response) GetSignedTransactionOk() (*string, bool)`

GetSignedTransactionOk returns a tuple with the SignedTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedTransaction

`func (o *Erc1155Response) SetSignedTransaction(v string)`

SetSignedTransaction sets SignedTransaction field to given value.


### GetSignedMessage

`func (o *Erc1155Response) GetSignedMessage() string`

GetSignedMessage returns the SignedMessage field if non-nil, zero value otherwise.

### GetSignedMessageOk

`func (o *Erc1155Response) GetSignedMessageOk() (*string, bool)`

GetSignedMessageOk returns a tuple with the SignedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedMessage

`func (o *Erc1155Response) SetSignedMessage(v string)`

SetSignedMessage sets SignedMessage field to given value.

### HasSignedMessage

`func (o *Erc1155Response) HasSignedMessage() bool`

HasSignedMessage returns a boolean if a field has been set.

### GetRawTransaction

`func (o *Erc1155Response) GetRawTransaction() string`

GetRawTransaction returns the RawTransaction field if non-nil, zero value otherwise.

### GetRawTransactionOk

`func (o *Erc1155Response) GetRawTransactionOk() (*string, bool)`

GetRawTransactionOk returns a tuple with the RawTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawTransaction

`func (o *Erc1155Response) SetRawTransaction(v string)`

SetRawTransaction sets RawTransaction field to given value.

### HasRawTransaction

`func (o *Erc1155Response) HasRawTransaction() bool`

HasRawTransaction returns a boolean if a field has been set.

### GetSignature

`func (o *Erc1155Response) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *Erc1155Response) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *Erc1155Response) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *Erc1155Response) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetTransaction

`func (o *Erc1155Response) GetTransaction() map[string]Tx`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *Erc1155Response) GetTransactionOk() (*map[string]Tx, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *Erc1155Response) SetTransaction(v map[string]Tx)`

SetTransaction sets Transaction field to given value.

### HasTransaction

`func (o *Erc1155Response) HasTransaction() bool`

HasTransaction returns a boolean if a field has been set.

### GetUserOps

`func (o *Erc1155Response) GetUserOps() []TransactionRequest`

GetUserOps returns the UserOps field if non-nil, zero value otherwise.

### GetUserOpsOk

`func (o *Erc1155Response) GetUserOpsOk() (*[]TransactionRequest, bool)`

GetUserOpsOk returns a tuple with the UserOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserOps

`func (o *Erc1155Response) SetUserOps(v []TransactionRequest)`

SetUserOps sets UserOps field to given value.

### HasUserOps

`func (o *Erc1155Response) HasUserOps() bool`

HasUserOps returns a boolean if a field has been set.

### GetUseropTransaction

`func (o *Erc1155Response) GetUseropTransaction() string`

GetUseropTransaction returns the UseropTransaction field if non-nil, zero value otherwise.

### GetUseropTransactionOk

`func (o *Erc1155Response) GetUseropTransactionOk() (*string, bool)`

GetUseropTransactionOk returns a tuple with the UseropTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseropTransaction

`func (o *Erc1155Response) SetUseropTransaction(v string)`

SetUseropTransaction sets UseropTransaction field to given value.

### HasUseropTransaction

`func (o *Erc1155Response) HasUseropTransaction() bool`

HasUseropTransaction returns a boolean if a field has been set.

### GetBalanceOf

`func (o *Erc1155Response) GetBalanceOf() string`

GetBalanceOf returns the BalanceOf field if non-nil, zero value otherwise.

### GetBalanceOfOk

`func (o *Erc1155Response) GetBalanceOfOk() (*string, bool)`

GetBalanceOfOk returns a tuple with the BalanceOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceOf

`func (o *Erc1155Response) SetBalanceOf(v string)`

SetBalanceOf sets BalanceOf field to given value.

### HasBalanceOf

`func (o *Erc1155Response) HasBalanceOf() bool`

HasBalanceOf returns a boolean if a field has been set.

### GetBalanceOfBatch

`func (o *Erc1155Response) GetBalanceOfBatch() string`

GetBalanceOfBatch returns the BalanceOfBatch field if non-nil, zero value otherwise.

### GetBalanceOfBatchOk

`func (o *Erc1155Response) GetBalanceOfBatchOk() (*string, bool)`

GetBalanceOfBatchOk returns a tuple with the BalanceOfBatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceOfBatch

`func (o *Erc1155Response) SetBalanceOfBatch(v string)`

SetBalanceOfBatch sets BalanceOfBatch field to given value.

### HasBalanceOfBatch

`func (o *Erc1155Response) HasBalanceOfBatch() bool`

HasBalanceOfBatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


