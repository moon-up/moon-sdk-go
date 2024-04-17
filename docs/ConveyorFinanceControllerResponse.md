# ConveyorFinanceControllerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | Pointer to [**InputBody**](InputBody.md) |  | [optional] 
**Convey** | Pointer to [**TransactionResponse**](TransactionResponse.md) |  | [optional] 
**Data** | Pointer to [**TransactionData**](TransactionData.md) |  | [optional] 
**Tx** | Pointer to [**TransactionResponseTx**](TransactionResponseTx.md) |  | [optional] 
**Signed** | Pointer to [**Transaction**](Transaction.md) |  | [optional] 
**Success** | **bool** |  | 
**Message** | **string** |  | 

## Methods

### NewConveyorFinanceControllerResponse

`func NewConveyorFinanceControllerResponse(success bool, message string, ) *ConveyorFinanceControllerResponse`

NewConveyorFinanceControllerResponse instantiates a new ConveyorFinanceControllerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConveyorFinanceControllerResponseWithDefaults

`func NewConveyorFinanceControllerResponseWithDefaults() *ConveyorFinanceControllerResponse`

NewConveyorFinanceControllerResponseWithDefaults instantiates a new ConveyorFinanceControllerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *ConveyorFinanceControllerResponse) GetInput() InputBody`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ConveyorFinanceControllerResponse) GetInputOk() (*InputBody, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ConveyorFinanceControllerResponse) SetInput(v InputBody)`

SetInput sets Input field to given value.

### HasInput

`func (o *ConveyorFinanceControllerResponse) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetConvey

`func (o *ConveyorFinanceControllerResponse) GetConvey() TransactionResponse`

GetConvey returns the Convey field if non-nil, zero value otherwise.

### GetConveyOk

`func (o *ConveyorFinanceControllerResponse) GetConveyOk() (*TransactionResponse, bool)`

GetConveyOk returns a tuple with the Convey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvey

`func (o *ConveyorFinanceControllerResponse) SetConvey(v TransactionResponse)`

SetConvey sets Convey field to given value.

### HasConvey

`func (o *ConveyorFinanceControllerResponse) HasConvey() bool`

HasConvey returns a boolean if a field has been set.

### GetData

`func (o *ConveyorFinanceControllerResponse) GetData() TransactionData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConveyorFinanceControllerResponse) GetDataOk() (*TransactionData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConveyorFinanceControllerResponse) SetData(v TransactionData)`

SetData sets Data field to given value.

### HasData

`func (o *ConveyorFinanceControllerResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTx

`func (o *ConveyorFinanceControllerResponse) GetTx() TransactionResponseTx`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *ConveyorFinanceControllerResponse) GetTxOk() (*TransactionResponseTx, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *ConveyorFinanceControllerResponse) SetTx(v TransactionResponseTx)`

SetTx sets Tx field to given value.

### HasTx

`func (o *ConveyorFinanceControllerResponse) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetSigned

`func (o *ConveyorFinanceControllerResponse) GetSigned() Transaction`

GetSigned returns the Signed field if non-nil, zero value otherwise.

### GetSignedOk

`func (o *ConveyorFinanceControllerResponse) GetSignedOk() (*Transaction, bool)`

GetSignedOk returns a tuple with the Signed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigned

`func (o *ConveyorFinanceControllerResponse) SetSigned(v Transaction)`

SetSigned sets Signed field to given value.

### HasSigned

`func (o *ConveyorFinanceControllerResponse) HasSigned() bool`

HasSigned returns a boolean if a field has been set.

### GetSuccess

`func (o *ConveyorFinanceControllerResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ConveyorFinanceControllerResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ConveyorFinanceControllerResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetMessage

`func (o *ConveyorFinanceControllerResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConveyorFinanceControllerResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConveyorFinanceControllerResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


