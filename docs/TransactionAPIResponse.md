# TransactionAPIResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Message** | **string** |  | 
**Body** | Pointer to [**InputBody**](InputBody.md) |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**TransactionHash** | Pointer to **interface{}** |  | [optional] 
**SignedTx** | Pointer to **interface{}** |  | [optional] 
**Data** | Pointer to [**Transaction**](Transaction.md) |  | [optional] 

## Methods

### NewTransactionAPIResponse

`func NewTransactionAPIResponse(success bool, message string, ) *TransactionAPIResponse`

NewTransactionAPIResponse instantiates a new TransactionAPIResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionAPIResponseWithDefaults

`func NewTransactionAPIResponseWithDefaults() *TransactionAPIResponse`

NewTransactionAPIResponseWithDefaults instantiates a new TransactionAPIResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *TransactionAPIResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *TransactionAPIResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *TransactionAPIResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetMessage

`func (o *TransactionAPIResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TransactionAPIResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TransactionAPIResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetBody

`func (o *TransactionAPIResponse) GetBody() InputBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TransactionAPIResponse) GetBodyOk() (*InputBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TransactionAPIResponse) SetBody(v InputBody)`

SetBody sets Body field to given value.

### HasBody

`func (o *TransactionAPIResponse) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetAddress

`func (o *TransactionAPIResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TransactionAPIResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TransactionAPIResponse) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TransactionAPIResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetTransactionHash

`func (o *TransactionAPIResponse) GetTransactionHash() interface{}`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *TransactionAPIResponse) GetTransactionHashOk() (*interface{}, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *TransactionAPIResponse) SetTransactionHash(v interface{})`

SetTransactionHash sets TransactionHash field to given value.

### HasTransactionHash

`func (o *TransactionAPIResponse) HasTransactionHash() bool`

HasTransactionHash returns a boolean if a field has been set.

### SetTransactionHashNil

`func (o *TransactionAPIResponse) SetTransactionHashNil(b bool)`

 SetTransactionHashNil sets the value for TransactionHash to be an explicit nil

### UnsetTransactionHash
`func (o *TransactionAPIResponse) UnsetTransactionHash()`

UnsetTransactionHash ensures that no value is present for TransactionHash, not even an explicit nil
### GetSignedTx

`func (o *TransactionAPIResponse) GetSignedTx() interface{}`

GetSignedTx returns the SignedTx field if non-nil, zero value otherwise.

### GetSignedTxOk

`func (o *TransactionAPIResponse) GetSignedTxOk() (*interface{}, bool)`

GetSignedTxOk returns a tuple with the SignedTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedTx

`func (o *TransactionAPIResponse) SetSignedTx(v interface{})`

SetSignedTx sets SignedTx field to given value.

### HasSignedTx

`func (o *TransactionAPIResponse) HasSignedTx() bool`

HasSignedTx returns a boolean if a field has been set.

### SetSignedTxNil

`func (o *TransactionAPIResponse) SetSignedTxNil(b bool)`

 SetSignedTxNil sets the value for SignedTx to be an explicit nil

### UnsetSignedTx
`func (o *TransactionAPIResponse) UnsetSignedTx()`

UnsetSignedTx ensures that no value is present for SignedTx, not even an explicit nil
### GetData

`func (o *TransactionAPIResponse) GetData() Transaction`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TransactionAPIResponse) GetDataOk() (*Transaction, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TransactionAPIResponse) SetData(v Transaction)`

SetData sets Data field to given value.

### HasData

`func (o *TransactionAPIResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


