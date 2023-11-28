# Erc20Response

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
**Decimals** | Pointer to **string** |  | [optional] 
**TotalSupply** | Pointer to **string** |  | [optional] 
**ContractAddress** | Pointer to **string** |  | [optional] 
**BalanceOf** | Pointer to **string** |  | [optional] 
**Allowance** | Pointer to **string** |  | [optional] 

## Methods

### NewErc20Response

`func NewErc20Response(transactionHash string, signedTransaction string, ) *Erc20Response`

NewErc20Response instantiates a new Erc20Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErc20ResponseWithDefaults

`func NewErc20ResponseWithDefaults() *Erc20Response`

NewErc20ResponseWithDefaults instantiates a new Erc20Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoonScanUrl

`func (o *Erc20Response) GetMoonScanUrl() string`

GetMoonScanUrl returns the MoonScanUrl field if non-nil, zero value otherwise.

### GetMoonScanUrlOk

`func (o *Erc20Response) GetMoonScanUrlOk() (*string, bool)`

GetMoonScanUrlOk returns a tuple with the MoonScanUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoonScanUrl

`func (o *Erc20Response) SetMoonScanUrl(v string)`

SetMoonScanUrl sets MoonScanUrl field to given value.

### HasMoonScanUrl

`func (o *Erc20Response) HasMoonScanUrl() bool`

HasMoonScanUrl returns a boolean if a field has been set.

### GetTransactionHash

`func (o *Erc20Response) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *Erc20Response) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *Erc20Response) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetSignedTransaction

`func (o *Erc20Response) GetSignedTransaction() string`

GetSignedTransaction returns the SignedTransaction field if non-nil, zero value otherwise.

### GetSignedTransactionOk

`func (o *Erc20Response) GetSignedTransactionOk() (*string, bool)`

GetSignedTransactionOk returns a tuple with the SignedTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedTransaction

`func (o *Erc20Response) SetSignedTransaction(v string)`

SetSignedTransaction sets SignedTransaction field to given value.


### GetSignedMessage

`func (o *Erc20Response) GetSignedMessage() string`

GetSignedMessage returns the SignedMessage field if non-nil, zero value otherwise.

### GetSignedMessageOk

`func (o *Erc20Response) GetSignedMessageOk() (*string, bool)`

GetSignedMessageOk returns a tuple with the SignedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedMessage

`func (o *Erc20Response) SetSignedMessage(v string)`

SetSignedMessage sets SignedMessage field to given value.

### HasSignedMessage

`func (o *Erc20Response) HasSignedMessage() bool`

HasSignedMessage returns a boolean if a field has been set.

### GetRawTransaction

`func (o *Erc20Response) GetRawTransaction() string`

GetRawTransaction returns the RawTransaction field if non-nil, zero value otherwise.

### GetRawTransactionOk

`func (o *Erc20Response) GetRawTransactionOk() (*string, bool)`

GetRawTransactionOk returns a tuple with the RawTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawTransaction

`func (o *Erc20Response) SetRawTransaction(v string)`

SetRawTransaction sets RawTransaction field to given value.

### HasRawTransaction

`func (o *Erc20Response) HasRawTransaction() bool`

HasRawTransaction returns a boolean if a field has been set.

### GetSignature

`func (o *Erc20Response) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *Erc20Response) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *Erc20Response) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *Erc20Response) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetTransaction

`func (o *Erc20Response) GetTransaction() map[string]Tx`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *Erc20Response) GetTransactionOk() (*map[string]Tx, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *Erc20Response) SetTransaction(v map[string]Tx)`

SetTransaction sets Transaction field to given value.

### HasTransaction

`func (o *Erc20Response) HasTransaction() bool`

HasTransaction returns a boolean if a field has been set.

### GetUserOps

`func (o *Erc20Response) GetUserOps() []TransactionRequest`

GetUserOps returns the UserOps field if non-nil, zero value otherwise.

### GetUserOpsOk

`func (o *Erc20Response) GetUserOpsOk() (*[]TransactionRequest, bool)`

GetUserOpsOk returns a tuple with the UserOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserOps

`func (o *Erc20Response) SetUserOps(v []TransactionRequest)`

SetUserOps sets UserOps field to given value.

### HasUserOps

`func (o *Erc20Response) HasUserOps() bool`

HasUserOps returns a boolean if a field has been set.

### GetUseropTransaction

`func (o *Erc20Response) GetUseropTransaction() string`

GetUseropTransaction returns the UseropTransaction field if non-nil, zero value otherwise.

### GetUseropTransactionOk

`func (o *Erc20Response) GetUseropTransactionOk() (*string, bool)`

GetUseropTransactionOk returns a tuple with the UseropTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseropTransaction

`func (o *Erc20Response) SetUseropTransaction(v string)`

SetUseropTransaction sets UseropTransaction field to given value.

### HasUseropTransaction

`func (o *Erc20Response) HasUseropTransaction() bool`

HasUseropTransaction returns a boolean if a field has been set.

### GetName

`func (o *Erc20Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Erc20Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Erc20Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Erc20Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSymbol

`func (o *Erc20Response) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Erc20Response) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Erc20Response) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Erc20Response) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDecimals

`func (o *Erc20Response) GetDecimals() string`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *Erc20Response) GetDecimalsOk() (*string, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *Erc20Response) SetDecimals(v string)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *Erc20Response) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.

### GetTotalSupply

`func (o *Erc20Response) GetTotalSupply() string`

GetTotalSupply returns the TotalSupply field if non-nil, zero value otherwise.

### GetTotalSupplyOk

`func (o *Erc20Response) GetTotalSupplyOk() (*string, bool)`

GetTotalSupplyOk returns a tuple with the TotalSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSupply

`func (o *Erc20Response) SetTotalSupply(v string)`

SetTotalSupply sets TotalSupply field to given value.

### HasTotalSupply

`func (o *Erc20Response) HasTotalSupply() bool`

HasTotalSupply returns a boolean if a field has been set.

### GetContractAddress

`func (o *Erc20Response) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *Erc20Response) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *Erc20Response) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.

### HasContractAddress

`func (o *Erc20Response) HasContractAddress() bool`

HasContractAddress returns a boolean if a field has been set.

### GetBalanceOf

`func (o *Erc20Response) GetBalanceOf() string`

GetBalanceOf returns the BalanceOf field if non-nil, zero value otherwise.

### GetBalanceOfOk

`func (o *Erc20Response) GetBalanceOfOk() (*string, bool)`

GetBalanceOfOk returns a tuple with the BalanceOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceOf

`func (o *Erc20Response) SetBalanceOf(v string)`

SetBalanceOf sets BalanceOf field to given value.

### HasBalanceOf

`func (o *Erc20Response) HasBalanceOf() bool`

HasBalanceOf returns a boolean if a field has been set.

### GetAllowance

`func (o *Erc20Response) GetAllowance() string`

GetAllowance returns the Allowance field if non-nil, zero value otherwise.

### GetAllowanceOk

`func (o *Erc20Response) GetAllowanceOk() (*string, bool)`

GetAllowanceOk returns a tuple with the Allowance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowance

`func (o *Erc20Response) SetAllowance(v string)`

SetAllowance sets Allowance field to given value.

### HasAllowance

`func (o *Erc20Response) HasAllowance() bool`

HasAllowance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


