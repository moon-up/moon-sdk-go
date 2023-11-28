# Erc721Response

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
**Name** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**BalanceOf** | Pointer to **string** |  | [optional] 
**OwnerOf** | Pointer to **string** |  | [optional] 
**TokenUri** | Pointer to **string** |  | [optional] 
**ContractAddress** | Pointer to **string** |  | [optional] 
**IsApprovedForAll** | Pointer to **string** |  | [optional] 

## Methods

### NewErc721Response

`func NewErc721Response(transactionHash string, signedTransaction string, ) *Erc721Response`

NewErc721Response instantiates a new Erc721Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErc721ResponseWithDefaults

`func NewErc721ResponseWithDefaults() *Erc721Response`

NewErc721ResponseWithDefaults instantiates a new Erc721Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoonScanUrl

`func (o *Erc721Response) GetMoonScanUrl() string`

GetMoonScanUrl returns the MoonScanUrl field if non-nil, zero value otherwise.

### GetMoonScanUrlOk

`func (o *Erc721Response) GetMoonScanUrlOk() (*string, bool)`

GetMoonScanUrlOk returns a tuple with the MoonScanUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoonScanUrl

`func (o *Erc721Response) SetMoonScanUrl(v string)`

SetMoonScanUrl sets MoonScanUrl field to given value.

### HasMoonScanUrl

`func (o *Erc721Response) HasMoonScanUrl() bool`

HasMoonScanUrl returns a boolean if a field has been set.

### GetTransactionHash

`func (o *Erc721Response) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *Erc721Response) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *Erc721Response) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetSignedTransaction

`func (o *Erc721Response) GetSignedTransaction() string`

GetSignedTransaction returns the SignedTransaction field if non-nil, zero value otherwise.

### GetSignedTransactionOk

`func (o *Erc721Response) GetSignedTransactionOk() (*string, bool)`

GetSignedTransactionOk returns a tuple with the SignedTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedTransaction

`func (o *Erc721Response) SetSignedTransaction(v string)`

SetSignedTransaction sets SignedTransaction field to given value.


### GetSignedMessage

`func (o *Erc721Response) GetSignedMessage() string`

GetSignedMessage returns the SignedMessage field if non-nil, zero value otherwise.

### GetSignedMessageOk

`func (o *Erc721Response) GetSignedMessageOk() (*string, bool)`

GetSignedMessageOk returns a tuple with the SignedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedMessage

`func (o *Erc721Response) SetSignedMessage(v string)`

SetSignedMessage sets SignedMessage field to given value.

### HasSignedMessage

`func (o *Erc721Response) HasSignedMessage() bool`

HasSignedMessage returns a boolean if a field has been set.

### GetRawTransaction

`func (o *Erc721Response) GetRawTransaction() string`

GetRawTransaction returns the RawTransaction field if non-nil, zero value otherwise.

### GetRawTransactionOk

`func (o *Erc721Response) GetRawTransactionOk() (*string, bool)`

GetRawTransactionOk returns a tuple with the RawTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawTransaction

`func (o *Erc721Response) SetRawTransaction(v string)`

SetRawTransaction sets RawTransaction field to given value.

### HasRawTransaction

`func (o *Erc721Response) HasRawTransaction() bool`

HasRawTransaction returns a boolean if a field has been set.

### GetSignature

`func (o *Erc721Response) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *Erc721Response) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *Erc721Response) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *Erc721Response) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetTransaction

`func (o *Erc721Response) GetTransaction() map[string]Tx`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *Erc721Response) GetTransactionOk() (*map[string]Tx, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *Erc721Response) SetTransaction(v map[string]Tx)`

SetTransaction sets Transaction field to given value.

### HasTransaction

`func (o *Erc721Response) HasTransaction() bool`

HasTransaction returns a boolean if a field has been set.

### GetUserOps

`func (o *Erc721Response) GetUserOps() []TransactionRequest`

GetUserOps returns the UserOps field if non-nil, zero value otherwise.

### GetUserOpsOk

`func (o *Erc721Response) GetUserOpsOk() (*[]TransactionRequest, bool)`

GetUserOpsOk returns a tuple with the UserOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserOps

`func (o *Erc721Response) SetUserOps(v []TransactionRequest)`

SetUserOps sets UserOps field to given value.

### HasUserOps

`func (o *Erc721Response) HasUserOps() bool`

HasUserOps returns a boolean if a field has been set.

### GetUseropTransaction

`func (o *Erc721Response) GetUseropTransaction() string`

GetUseropTransaction returns the UseropTransaction field if non-nil, zero value otherwise.

### GetUseropTransactionOk

`func (o *Erc721Response) GetUseropTransactionOk() (*string, bool)`

GetUseropTransactionOk returns a tuple with the UseropTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseropTransaction

`func (o *Erc721Response) SetUseropTransaction(v string)`

SetUseropTransaction sets UseropTransaction field to given value.

### HasUseropTransaction

`func (o *Erc721Response) HasUseropTransaction() bool`

HasUseropTransaction returns a boolean if a field has been set.

### GetName

`func (o *Erc721Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Erc721Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Erc721Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Erc721Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSymbol

`func (o *Erc721Response) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Erc721Response) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Erc721Response) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Erc721Response) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetBalanceOf

`func (o *Erc721Response) GetBalanceOf() string`

GetBalanceOf returns the BalanceOf field if non-nil, zero value otherwise.

### GetBalanceOfOk

`func (o *Erc721Response) GetBalanceOfOk() (*string, bool)`

GetBalanceOfOk returns a tuple with the BalanceOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceOf

`func (o *Erc721Response) SetBalanceOf(v string)`

SetBalanceOf sets BalanceOf field to given value.

### HasBalanceOf

`func (o *Erc721Response) HasBalanceOf() bool`

HasBalanceOf returns a boolean if a field has been set.

### GetOwnerOf

`func (o *Erc721Response) GetOwnerOf() string`

GetOwnerOf returns the OwnerOf field if non-nil, zero value otherwise.

### GetOwnerOfOk

`func (o *Erc721Response) GetOwnerOfOk() (*string, bool)`

GetOwnerOfOk returns a tuple with the OwnerOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerOf

`func (o *Erc721Response) SetOwnerOf(v string)`

SetOwnerOf sets OwnerOf field to given value.

### HasOwnerOf

`func (o *Erc721Response) HasOwnerOf() bool`

HasOwnerOf returns a boolean if a field has been set.

### GetTokenUri

`func (o *Erc721Response) GetTokenUri() string`

GetTokenUri returns the TokenUri field if non-nil, zero value otherwise.

### GetTokenUriOk

`func (o *Erc721Response) GetTokenUriOk() (*string, bool)`

GetTokenUriOk returns a tuple with the TokenUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUri

`func (o *Erc721Response) SetTokenUri(v string)`

SetTokenUri sets TokenUri field to given value.

### HasTokenUri

`func (o *Erc721Response) HasTokenUri() bool`

HasTokenUri returns a boolean if a field has been set.

### GetContractAddress

`func (o *Erc721Response) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *Erc721Response) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *Erc721Response) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.

### HasContractAddress

`func (o *Erc721Response) HasContractAddress() bool`

HasContractAddress returns a boolean if a field has been set.

### GetIsApprovedForAll

`func (o *Erc721Response) GetIsApprovedForAll() string`

GetIsApprovedForAll returns the IsApprovedForAll field if non-nil, zero value otherwise.

### GetIsApprovedForAllOk

`func (o *Erc721Response) GetIsApprovedForAllOk() (*string, bool)`

GetIsApprovedForAllOk returns a tuple with the IsApprovedForAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApprovedForAll

`func (o *Erc721Response) SetIsApprovedForAll(v string)`

SetIsApprovedForAll sets IsApprovedForAll field to given value.

### HasIsApprovedForAll

`func (o *Erc721Response) HasIsApprovedForAll() bool`

HasIsApprovedForAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


