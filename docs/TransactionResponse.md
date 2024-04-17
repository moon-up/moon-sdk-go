# TransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Tx** | [**TransactionResponseTx**](TransactionResponseTx.md) |  | 
**Info** | [**TransactionResponseInfo**](TransactionResponseInfo.md) |  | 
**ChainId** | **float64** |  | 
**CurrentBlockNumber** | **float64** |  | 

## Methods

### NewTransactionResponse

`func NewTransactionResponse(message string, tx TransactionResponseTx, info TransactionResponseInfo, chainId float64, currentBlockNumber float64, ) *TransactionResponse`

NewTransactionResponse instantiates a new TransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionResponseWithDefaults

`func NewTransactionResponseWithDefaults() *TransactionResponse`

NewTransactionResponseWithDefaults instantiates a new TransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *TransactionResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TransactionResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TransactionResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTx

`func (o *TransactionResponse) GetTx() TransactionResponseTx`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *TransactionResponse) GetTxOk() (*TransactionResponseTx, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *TransactionResponse) SetTx(v TransactionResponseTx)`

SetTx sets Tx field to given value.


### GetInfo

`func (o *TransactionResponse) GetInfo() TransactionResponseInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TransactionResponse) GetInfoOk() (*TransactionResponseInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TransactionResponse) SetInfo(v TransactionResponseInfo)`

SetInfo sets Info field to given value.


### GetChainId

`func (o *TransactionResponse) GetChainId() float64`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *TransactionResponse) GetChainIdOk() (*float64, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *TransactionResponse) SetChainId(v float64)`

SetChainId sets ChainId field to given value.


### GetCurrentBlockNumber

`func (o *TransactionResponse) GetCurrentBlockNumber() float64`

GetCurrentBlockNumber returns the CurrentBlockNumber field if non-nil, zero value otherwise.

### GetCurrentBlockNumberOk

`func (o *TransactionResponse) GetCurrentBlockNumberOk() (*float64, bool)`

GetCurrentBlockNumberOk returns a tuple with the CurrentBlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBlockNumber

`func (o *TransactionResponse) SetCurrentBlockNumber(v float64)`

SetCurrentBlockNumber sets CurrentBlockNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


