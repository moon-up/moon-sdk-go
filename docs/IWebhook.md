# IWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Block** | [**Block**](Block.md) |  | 
**ChainId** | **string** |  | 
**Logs** | [**[]Log**](Log.md) |  | 
**Txs** | [**[]Transaction**](Transaction.md) |  | 
**TxsInternal** | [**[]InternalTransaction**](InternalTransaction.md) |  | 
**Abi** | [**[]AbiItem**](AbiItem.md) |  | 
**Retries** | **float64** |  | 
**Confirmed** | **bool** |  | 
**Tag** | **string** |  | 
**StreamId** | **string** |  | 
**Erc20Transfers** | [**[]IERC20Transfer**](IERC20Transfer.md) |  | 
**Erc20Approvals** | [**[]IERC20Approval**](IERC20Approval.md) |  | 
**NftTransfers** | [**[]INFTTransfer**](INFTTransfer.md) |  | 
**NativeBalances** | [**[]INativeBalance**](INativeBalance.md) |  | 
**NftApprovals** | [**IOldNFTApproval**](IOldNFTApproval.md) |  | 
**NftTokenApprovals** | [**[]INFTApproval**](INFTApproval.md) |  | 

## Methods

### NewIWebhook

`func NewIWebhook(block Block, chainId string, logs []Log, txs []Transaction, txsInternal []InternalTransaction, abi []AbiItem, retries float64, confirmed bool, tag string, streamId string, erc20Transfers []IERC20Transfer, erc20Approvals []IERC20Approval, nftTransfers []INFTTransfer, nativeBalances []INativeBalance, nftApprovals IOldNFTApproval, nftTokenApprovals []INFTApproval, ) *IWebhook`

NewIWebhook instantiates a new IWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIWebhookWithDefaults

`func NewIWebhookWithDefaults() *IWebhook`

NewIWebhookWithDefaults instantiates a new IWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlock

`func (o *IWebhook) GetBlock() Block`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *IWebhook) GetBlockOk() (*Block, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *IWebhook) SetBlock(v Block)`

SetBlock sets Block field to given value.


### GetChainId

`func (o *IWebhook) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *IWebhook) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *IWebhook) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetLogs

`func (o *IWebhook) GetLogs() []Log`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *IWebhook) GetLogsOk() (*[]Log, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *IWebhook) SetLogs(v []Log)`

SetLogs sets Logs field to given value.


### GetTxs

`func (o *IWebhook) GetTxs() []Transaction`

GetTxs returns the Txs field if non-nil, zero value otherwise.

### GetTxsOk

`func (o *IWebhook) GetTxsOk() (*[]Transaction, bool)`

GetTxsOk returns a tuple with the Txs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxs

`func (o *IWebhook) SetTxs(v []Transaction)`

SetTxs sets Txs field to given value.


### GetTxsInternal

`func (o *IWebhook) GetTxsInternal() []InternalTransaction`

GetTxsInternal returns the TxsInternal field if non-nil, zero value otherwise.

### GetTxsInternalOk

`func (o *IWebhook) GetTxsInternalOk() (*[]InternalTransaction, bool)`

GetTxsInternalOk returns a tuple with the TxsInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxsInternal

`func (o *IWebhook) SetTxsInternal(v []InternalTransaction)`

SetTxsInternal sets TxsInternal field to given value.


### GetAbi

`func (o *IWebhook) GetAbi() []AbiItem`

GetAbi returns the Abi field if non-nil, zero value otherwise.

### GetAbiOk

`func (o *IWebhook) GetAbiOk() (*[]AbiItem, bool)`

GetAbiOk returns a tuple with the Abi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbi

`func (o *IWebhook) SetAbi(v []AbiItem)`

SetAbi sets Abi field to given value.


### GetRetries

`func (o *IWebhook) GetRetries() float64`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *IWebhook) GetRetriesOk() (*float64, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *IWebhook) SetRetries(v float64)`

SetRetries sets Retries field to given value.


### GetConfirmed

`func (o *IWebhook) GetConfirmed() bool`

GetConfirmed returns the Confirmed field if non-nil, zero value otherwise.

### GetConfirmedOk

`func (o *IWebhook) GetConfirmedOk() (*bool, bool)`

GetConfirmedOk returns a tuple with the Confirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmed

`func (o *IWebhook) SetConfirmed(v bool)`

SetConfirmed sets Confirmed field to given value.


### GetTag

`func (o *IWebhook) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *IWebhook) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *IWebhook) SetTag(v string)`

SetTag sets Tag field to given value.


### GetStreamId

`func (o *IWebhook) GetStreamId() string`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *IWebhook) GetStreamIdOk() (*string, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *IWebhook) SetStreamId(v string)`

SetStreamId sets StreamId field to given value.


### GetErc20Transfers

`func (o *IWebhook) GetErc20Transfers() []IERC20Transfer`

GetErc20Transfers returns the Erc20Transfers field if non-nil, zero value otherwise.

### GetErc20TransfersOk

`func (o *IWebhook) GetErc20TransfersOk() (*[]IERC20Transfer, bool)`

GetErc20TransfersOk returns a tuple with the Erc20Transfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErc20Transfers

`func (o *IWebhook) SetErc20Transfers(v []IERC20Transfer)`

SetErc20Transfers sets Erc20Transfers field to given value.


### GetErc20Approvals

`func (o *IWebhook) GetErc20Approvals() []IERC20Approval`

GetErc20Approvals returns the Erc20Approvals field if non-nil, zero value otherwise.

### GetErc20ApprovalsOk

`func (o *IWebhook) GetErc20ApprovalsOk() (*[]IERC20Approval, bool)`

GetErc20ApprovalsOk returns a tuple with the Erc20Approvals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErc20Approvals

`func (o *IWebhook) SetErc20Approvals(v []IERC20Approval)`

SetErc20Approvals sets Erc20Approvals field to given value.


### GetNftTransfers

`func (o *IWebhook) GetNftTransfers() []INFTTransfer`

GetNftTransfers returns the NftTransfers field if non-nil, zero value otherwise.

### GetNftTransfersOk

`func (o *IWebhook) GetNftTransfersOk() (*[]INFTTransfer, bool)`

GetNftTransfersOk returns a tuple with the NftTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftTransfers

`func (o *IWebhook) SetNftTransfers(v []INFTTransfer)`

SetNftTransfers sets NftTransfers field to given value.


### GetNativeBalances

`func (o *IWebhook) GetNativeBalances() []INativeBalance`

GetNativeBalances returns the NativeBalances field if non-nil, zero value otherwise.

### GetNativeBalancesOk

`func (o *IWebhook) GetNativeBalancesOk() (*[]INativeBalance, bool)`

GetNativeBalancesOk returns a tuple with the NativeBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeBalances

`func (o *IWebhook) SetNativeBalances(v []INativeBalance)`

SetNativeBalances sets NativeBalances field to given value.


### GetNftApprovals

`func (o *IWebhook) GetNftApprovals() IOldNFTApproval`

GetNftApprovals returns the NftApprovals field if non-nil, zero value otherwise.

### GetNftApprovalsOk

`func (o *IWebhook) GetNftApprovalsOk() (*IOldNFTApproval, bool)`

GetNftApprovalsOk returns a tuple with the NftApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftApprovals

`func (o *IWebhook) SetNftApprovals(v IOldNFTApproval)`

SetNftApprovals sets NftApprovals field to given value.


### GetNftTokenApprovals

`func (o *IWebhook) GetNftTokenApprovals() []INFTApproval`

GetNftTokenApprovals returns the NftTokenApprovals field if non-nil, zero value otherwise.

### GetNftTokenApprovalsOk

`func (o *IWebhook) GetNftTokenApprovalsOk() (*[]INFTApproval, bool)`

GetNftTokenApprovalsOk returns a tuple with the NftTokenApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftTokenApprovals

`func (o *IWebhook) SetNftTokenApprovals(v []INFTApproval)`

SetNftTokenApprovals sets NftTokenApprovals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


