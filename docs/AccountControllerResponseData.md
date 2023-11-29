# AccountControllerResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nonce** | **float64** |  | 
**Balance** | **string** |  | 
**UseropTransaction** | Pointer to **string** |  | [optional] 
**UserOps** | Pointer to [**[]TransactionRequest**](TransactionRequest.md) |  | [optional] 
**Transaction** | Pointer to [**Tx**](Tx.md) |  | [optional] 
**Signature** | Pointer to **string** |  | [optional] 
**MoonScanUrl** | Pointer to **string** |  | [optional] 
**Transactions** | Pointer to [**[]TransactionData**](TransactionData.md) |  | [optional] 
**Data** | **string** |  | 
**RawTransaction** | Pointer to **string** |  | [optional] 
**SignedTransaction** | Pointer to **string** |  | [optional] 
**TransactionHash** | Pointer to **string** |  | [optional] 
**Keys** | Pointer to **[]string** |  | [optional] 
**Address** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Encoding** | Pointer to **string** |  | [optional] 
**Header** | Pointer to **bool** |  | [optional] 
**Signtype** | Pointer to **bool** |  | [optional] 
**Domain** | **string** |  | 
**CurrentAtokenBalance** | **string** |  | 
**CurrentBorrowBalance** | **string** |  | 
**PrincipalBorrowBalance** | **string** |  | 
**BorrowRateMode** | **string** |  | 
**BorrowRate** | **string** |  | 
**LiquidityRate** | **string** |  | 
**OriginationFee** | **string** |  | 
**VariableBorrowIndex** | **string** |  | 
**LastUpdateTimestamp** | **string** |  | 
**UsageAsCollateralEnabled** | **string** |  | 
**Type** | Pointer to **float64** |  | [optional] 
**ChainId** | Pointer to **float64** |  | [optional] 
**Gas** | Pointer to **string** |  | [optional] 
**GasPrice** | Pointer to **string** |  | [optional] 
**GasTipCap** | Pointer to **NullableString** |  | [optional] 
**GasFeeCap** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**To** | Pointer to **NullableString** |  | [optional] 
**BlobGas** | Pointer to **NullableString** |  | [optional] 
**BlobGasFeeCap** | Pointer to **NullableString** |  | [optional] 
**BlobHashes** | Pointer to **[]string** |  | [optional] 
**V** | Pointer to **string** |  | [optional] 
**R** | Pointer to **string** |  | [optional] 
**S** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Decimals** | Pointer to **string** |  | [optional] 
**TotalSupply** | Pointer to **string** |  | [optional] 
**ContractAddress** | Pointer to **string** |  | [optional] 
**BalanceOf** | Pointer to **string** |  | [optional] 
**Allowance** | Pointer to **string** |  | [optional] 
**BalanceOf** | Pointer to **string** |  | [optional] 
**BalanceOfBatch** | Pointer to **string** |  | [optional] 
**Success** | **bool** |  | 
**Message** | **string** |  | 
**SignedTx** | Pointer to **string** |  | [optional] 
**OwnerOf** | Pointer to **string** |  | [optional] 
**TokenUri** | Pointer to **string** |  | [optional] 
**IsApprovedForAll** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountControllerResponseData

`func NewAccountControllerResponseData(nonce float64, balance string, data string, address string, domain string, currentAtokenBalance string, currentBorrowBalance string, principalBorrowBalance string, borrowRateMode string, borrowRate string, liquidityRate string, originationFee string, variableBorrowIndex string, lastUpdateTimestamp string, usageAsCollateralEnabled string, success bool, message string, ) *AccountControllerResponseData`

NewAccountControllerResponseData instantiates a new AccountControllerResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountControllerResponseDataWithDefaults

`func NewAccountControllerResponseDataWithDefaults() *AccountControllerResponseData`

NewAccountControllerResponseDataWithDefaults instantiates a new AccountControllerResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNonce

`func (o *AccountControllerResponseData) GetNonce() float64`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *AccountControllerResponseData) GetNonceOk() (*float64, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *AccountControllerResponseData) SetNonce(v float64)`

SetNonce sets Nonce field to given value.


### GetBalance

`func (o *AccountControllerResponseData) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *AccountControllerResponseData) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *AccountControllerResponseData) SetBalance(v string)`

SetBalance sets Balance field to given value.


### GetUseropTransaction

`func (o *AccountControllerResponseData) GetUseropTransaction() string`

GetUseropTransaction returns the UseropTransaction field if non-nil, zero value otherwise.

### GetUseropTransactionOk

`func (o *AccountControllerResponseData) GetUseropTransactionOk() (*string, bool)`

GetUseropTransactionOk returns a tuple with the UseropTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseropTransaction

`func (o *AccountControllerResponseData) SetUseropTransaction(v string)`

SetUseropTransaction sets UseropTransaction field to given value.

### HasUseropTransaction

`func (o *AccountControllerResponseData) HasUseropTransaction() bool`

HasUseropTransaction returns a boolean if a field has been set.

### GetUserOps

`func (o *AccountControllerResponseData) GetUserOps() []TransactionRequest`

GetUserOps returns the UserOps field if non-nil, zero value otherwise.

### GetUserOpsOk

`func (o *AccountControllerResponseData) GetUserOpsOk() (*[]TransactionRequest, bool)`

GetUserOpsOk returns a tuple with the UserOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserOps

`func (o *AccountControllerResponseData) SetUserOps(v []TransactionRequest)`

SetUserOps sets UserOps field to given value.

### HasUserOps

`func (o *AccountControllerResponseData) HasUserOps() bool`

HasUserOps returns a boolean if a field has been set.

### GetTransaction

`func (o *AccountControllerResponseData) GetTransaction() Tx`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *AccountControllerResponseData) GetTransactionOk() (*Tx, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *AccountControllerResponseData) SetTransaction(v Tx)`

SetTransaction sets Transaction field to given value.

### HasTransaction

`func (o *AccountControllerResponseData) HasTransaction() bool`

HasTransaction returns a boolean if a field has been set.

### GetSignature

`func (o *AccountControllerResponseData) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *AccountControllerResponseData) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *AccountControllerResponseData) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *AccountControllerResponseData) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetMoonScanUrl

`func (o *AccountControllerResponseData) GetMoonScanUrl() string`

GetMoonScanUrl returns the MoonScanUrl field if non-nil, zero value otherwise.

### GetMoonScanUrlOk

`func (o *AccountControllerResponseData) GetMoonScanUrlOk() (*string, bool)`

GetMoonScanUrlOk returns a tuple with the MoonScanUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoonScanUrl

`func (o *AccountControllerResponseData) SetMoonScanUrl(v string)`

SetMoonScanUrl sets MoonScanUrl field to given value.

### HasMoonScanUrl

`func (o *AccountControllerResponseData) HasMoonScanUrl() bool`

HasMoonScanUrl returns a boolean if a field has been set.

### GetTransactions

`func (o *AccountControllerResponseData) GetTransactions() []TransactionData`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *AccountControllerResponseData) GetTransactionsOk() (*[]TransactionData, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *AccountControllerResponseData) SetTransactions(v []TransactionData)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *AccountControllerResponseData) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetData

`func (o *AccountControllerResponseData) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AccountControllerResponseData) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AccountControllerResponseData) SetData(v string)`

SetData sets Data field to given value.


### GetRawTransaction

`func (o *AccountControllerResponseData) GetRawTransaction() string`

GetRawTransaction returns the RawTransaction field if non-nil, zero value otherwise.

### GetRawTransactionOk

`func (o *AccountControllerResponseData) GetRawTransactionOk() (*string, bool)`

GetRawTransactionOk returns a tuple with the RawTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawTransaction

`func (o *AccountControllerResponseData) SetRawTransaction(v string)`

SetRawTransaction sets RawTransaction field to given value.

### HasRawTransaction

`func (o *AccountControllerResponseData) HasRawTransaction() bool`

HasRawTransaction returns a boolean if a field has been set.

### GetSignedTransaction

`func (o *AccountControllerResponseData) GetSignedTransaction() string`

GetSignedTransaction returns the SignedTransaction field if non-nil, zero value otherwise.

### GetSignedTransactionOk

`func (o *AccountControllerResponseData) GetSignedTransactionOk() (*string, bool)`

GetSignedTransactionOk returns a tuple with the SignedTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedTransaction

`func (o *AccountControllerResponseData) SetSignedTransaction(v string)`

SetSignedTransaction sets SignedTransaction field to given value.

### HasSignedTransaction

`func (o *AccountControllerResponseData) HasSignedTransaction() bool`

HasSignedTransaction returns a boolean if a field has been set.

### GetTransactionHash

`func (o *AccountControllerResponseData) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *AccountControllerResponseData) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *AccountControllerResponseData) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.

### HasTransactionHash

`func (o *AccountControllerResponseData) HasTransactionHash() bool`

HasTransactionHash returns a boolean if a field has been set.

### GetKeys

`func (o *AccountControllerResponseData) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *AccountControllerResponseData) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *AccountControllerResponseData) SetKeys(v []string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *AccountControllerResponseData) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetAddress

`func (o *AccountControllerResponseData) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AccountControllerResponseData) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AccountControllerResponseData) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetName

`func (o *AccountControllerResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountControllerResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountControllerResponseData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountControllerResponseData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEncoding

`func (o *AccountControllerResponseData) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *AccountControllerResponseData) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *AccountControllerResponseData) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *AccountControllerResponseData) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetHeader

`func (o *AccountControllerResponseData) GetHeader() bool`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *AccountControllerResponseData) GetHeaderOk() (*bool, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *AccountControllerResponseData) SetHeader(v bool)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *AccountControllerResponseData) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetSigntype

`func (o *AccountControllerResponseData) GetSigntype() bool`

GetSigntype returns the Signtype field if non-nil, zero value otherwise.

### GetSigntypeOk

`func (o *AccountControllerResponseData) GetSigntypeOk() (*bool, bool)`

GetSigntypeOk returns a tuple with the Signtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigntype

`func (o *AccountControllerResponseData) SetSigntype(v bool)`

SetSigntype sets Signtype field to given value.

### HasSigntype

`func (o *AccountControllerResponseData) HasSigntype() bool`

HasSigntype returns a boolean if a field has been set.

### GetDomain

`func (o *AccountControllerResponseData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *AccountControllerResponseData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *AccountControllerResponseData) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetCurrentAtokenBalance

`func (o *AccountControllerResponseData) GetCurrentAtokenBalance() string`

GetCurrentAtokenBalance returns the CurrentAtokenBalance field if non-nil, zero value otherwise.

### GetCurrentAtokenBalanceOk

`func (o *AccountControllerResponseData) GetCurrentAtokenBalanceOk() (*string, bool)`

GetCurrentAtokenBalanceOk returns a tuple with the CurrentAtokenBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAtokenBalance

`func (o *AccountControllerResponseData) SetCurrentAtokenBalance(v string)`

SetCurrentAtokenBalance sets CurrentAtokenBalance field to given value.


### GetCurrentBorrowBalance

`func (o *AccountControllerResponseData) GetCurrentBorrowBalance() string`

GetCurrentBorrowBalance returns the CurrentBorrowBalance field if non-nil, zero value otherwise.

### GetCurrentBorrowBalanceOk

`func (o *AccountControllerResponseData) GetCurrentBorrowBalanceOk() (*string, bool)`

GetCurrentBorrowBalanceOk returns a tuple with the CurrentBorrowBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBorrowBalance

`func (o *AccountControllerResponseData) SetCurrentBorrowBalance(v string)`

SetCurrentBorrowBalance sets CurrentBorrowBalance field to given value.


### GetPrincipalBorrowBalance

`func (o *AccountControllerResponseData) GetPrincipalBorrowBalance() string`

GetPrincipalBorrowBalance returns the PrincipalBorrowBalance field if non-nil, zero value otherwise.

### GetPrincipalBorrowBalanceOk

`func (o *AccountControllerResponseData) GetPrincipalBorrowBalanceOk() (*string, bool)`

GetPrincipalBorrowBalanceOk returns a tuple with the PrincipalBorrowBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalBorrowBalance

`func (o *AccountControllerResponseData) SetPrincipalBorrowBalance(v string)`

SetPrincipalBorrowBalance sets PrincipalBorrowBalance field to given value.


### GetBorrowRateMode

`func (o *AccountControllerResponseData) GetBorrowRateMode() string`

GetBorrowRateMode returns the BorrowRateMode field if non-nil, zero value otherwise.

### GetBorrowRateModeOk

`func (o *AccountControllerResponseData) GetBorrowRateModeOk() (*string, bool)`

GetBorrowRateModeOk returns a tuple with the BorrowRateMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorrowRateMode

`func (o *AccountControllerResponseData) SetBorrowRateMode(v string)`

SetBorrowRateMode sets BorrowRateMode field to given value.


### GetBorrowRate

`func (o *AccountControllerResponseData) GetBorrowRate() string`

GetBorrowRate returns the BorrowRate field if non-nil, zero value otherwise.

### GetBorrowRateOk

`func (o *AccountControllerResponseData) GetBorrowRateOk() (*string, bool)`

GetBorrowRateOk returns a tuple with the BorrowRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorrowRate

`func (o *AccountControllerResponseData) SetBorrowRate(v string)`

SetBorrowRate sets BorrowRate field to given value.


### GetLiquidityRate

`func (o *AccountControllerResponseData) GetLiquidityRate() string`

GetLiquidityRate returns the LiquidityRate field if non-nil, zero value otherwise.

### GetLiquidityRateOk

`func (o *AccountControllerResponseData) GetLiquidityRateOk() (*string, bool)`

GetLiquidityRateOk returns a tuple with the LiquidityRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidityRate

`func (o *AccountControllerResponseData) SetLiquidityRate(v string)`

SetLiquidityRate sets LiquidityRate field to given value.


### GetOriginationFee

`func (o *AccountControllerResponseData) GetOriginationFee() string`

GetOriginationFee returns the OriginationFee field if non-nil, zero value otherwise.

### GetOriginationFeeOk

`func (o *AccountControllerResponseData) GetOriginationFeeOk() (*string, bool)`

GetOriginationFeeOk returns a tuple with the OriginationFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginationFee

`func (o *AccountControllerResponseData) SetOriginationFee(v string)`

SetOriginationFee sets OriginationFee field to given value.


### GetVariableBorrowIndex

`func (o *AccountControllerResponseData) GetVariableBorrowIndex() string`

GetVariableBorrowIndex returns the VariableBorrowIndex field if non-nil, zero value otherwise.

### GetVariableBorrowIndexOk

`func (o *AccountControllerResponseData) GetVariableBorrowIndexOk() (*string, bool)`

GetVariableBorrowIndexOk returns a tuple with the VariableBorrowIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableBorrowIndex

`func (o *AccountControllerResponseData) SetVariableBorrowIndex(v string)`

SetVariableBorrowIndex sets VariableBorrowIndex field to given value.


### GetLastUpdateTimestamp

`func (o *AccountControllerResponseData) GetLastUpdateTimestamp() string`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *AccountControllerResponseData) GetLastUpdateTimestampOk() (*string, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *AccountControllerResponseData) SetLastUpdateTimestamp(v string)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.


### GetUsageAsCollateralEnabled

`func (o *AccountControllerResponseData) GetUsageAsCollateralEnabled() string`

GetUsageAsCollateralEnabled returns the UsageAsCollateralEnabled field if non-nil, zero value otherwise.

### GetUsageAsCollateralEnabledOk

`func (o *AccountControllerResponseData) GetUsageAsCollateralEnabledOk() (*string, bool)`

GetUsageAsCollateralEnabledOk returns a tuple with the UsageAsCollateralEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageAsCollateralEnabled

`func (o *AccountControllerResponseData) SetUsageAsCollateralEnabled(v string)`

SetUsageAsCollateralEnabled sets UsageAsCollateralEnabled field to given value.


### GetType

`func (o *AccountControllerResponseData) GetType() float64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountControllerResponseData) GetTypeOk() (*float64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountControllerResponseData) SetType(v float64)`

SetType sets Type field to given value.

### HasType

`func (o *AccountControllerResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetChainId

`func (o *AccountControllerResponseData) GetChainId() float64`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *AccountControllerResponseData) GetChainIdOk() (*float64, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *AccountControllerResponseData) SetChainId(v float64)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *AccountControllerResponseData) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetGas

`func (o *AccountControllerResponseData) GetGas() string`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *AccountControllerResponseData) GetGasOk() (*string, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *AccountControllerResponseData) SetGas(v string)`

SetGas sets Gas field to given value.

### HasGas

`func (o *AccountControllerResponseData) HasGas() bool`

HasGas returns a boolean if a field has been set.

### GetGasPrice

`func (o *AccountControllerResponseData) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *AccountControllerResponseData) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *AccountControllerResponseData) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *AccountControllerResponseData) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetGasTipCap

`func (o *AccountControllerResponseData) GetGasTipCap() string`

GetGasTipCap returns the GasTipCap field if non-nil, zero value otherwise.

### GetGasTipCapOk

`func (o *AccountControllerResponseData) GetGasTipCapOk() (*string, bool)`

GetGasTipCapOk returns a tuple with the GasTipCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasTipCap

`func (o *AccountControllerResponseData) SetGasTipCap(v string)`

SetGasTipCap sets GasTipCap field to given value.

### HasGasTipCap

`func (o *AccountControllerResponseData) HasGasTipCap() bool`

HasGasTipCap returns a boolean if a field has been set.

### SetGasTipCapNil

`func (o *AccountControllerResponseData) SetGasTipCapNil(b bool)`

 SetGasTipCapNil sets the value for GasTipCap to be an explicit nil

### UnsetGasTipCap
`func (o *AccountControllerResponseData) UnsetGasTipCap()`

UnsetGasTipCap ensures that no value is present for GasTipCap, not even an explicit nil
### GetGasFeeCap

`func (o *AccountControllerResponseData) GetGasFeeCap() string`

GetGasFeeCap returns the GasFeeCap field if non-nil, zero value otherwise.

### GetGasFeeCapOk

`func (o *AccountControllerResponseData) GetGasFeeCapOk() (*string, bool)`

GetGasFeeCapOk returns a tuple with the GasFeeCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasFeeCap

`func (o *AccountControllerResponseData) SetGasFeeCap(v string)`

SetGasFeeCap sets GasFeeCap field to given value.

### HasGasFeeCap

`func (o *AccountControllerResponseData) HasGasFeeCap() bool`

HasGasFeeCap returns a boolean if a field has been set.

### SetGasFeeCapNil

`func (o *AccountControllerResponseData) SetGasFeeCapNil(b bool)`

 SetGasFeeCapNil sets the value for GasFeeCap to be an explicit nil

### UnsetGasFeeCap
`func (o *AccountControllerResponseData) UnsetGasFeeCap()`

UnsetGasFeeCap ensures that no value is present for GasFeeCap, not even an explicit nil
### GetValue

`func (o *AccountControllerResponseData) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AccountControllerResponseData) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AccountControllerResponseData) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AccountControllerResponseData) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetFrom

`func (o *AccountControllerResponseData) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *AccountControllerResponseData) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *AccountControllerResponseData) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *AccountControllerResponseData) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *AccountControllerResponseData) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *AccountControllerResponseData) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *AccountControllerResponseData) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *AccountControllerResponseData) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *AccountControllerResponseData) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *AccountControllerResponseData) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetBlobGas

`func (o *AccountControllerResponseData) GetBlobGas() string`

GetBlobGas returns the BlobGas field if non-nil, zero value otherwise.

### GetBlobGasOk

`func (o *AccountControllerResponseData) GetBlobGasOk() (*string, bool)`

GetBlobGasOk returns a tuple with the BlobGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobGas

`func (o *AccountControllerResponseData) SetBlobGas(v string)`

SetBlobGas sets BlobGas field to given value.

### HasBlobGas

`func (o *AccountControllerResponseData) HasBlobGas() bool`

HasBlobGas returns a boolean if a field has been set.

### SetBlobGasNil

`func (o *AccountControllerResponseData) SetBlobGasNil(b bool)`

 SetBlobGasNil sets the value for BlobGas to be an explicit nil

### UnsetBlobGas
`func (o *AccountControllerResponseData) UnsetBlobGas()`

UnsetBlobGas ensures that no value is present for BlobGas, not even an explicit nil
### GetBlobGasFeeCap

`func (o *AccountControllerResponseData) GetBlobGasFeeCap() string`

GetBlobGasFeeCap returns the BlobGasFeeCap field if non-nil, zero value otherwise.

### GetBlobGasFeeCapOk

`func (o *AccountControllerResponseData) GetBlobGasFeeCapOk() (*string, bool)`

GetBlobGasFeeCapOk returns a tuple with the BlobGasFeeCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobGasFeeCap

`func (o *AccountControllerResponseData) SetBlobGasFeeCap(v string)`

SetBlobGasFeeCap sets BlobGasFeeCap field to given value.

### HasBlobGasFeeCap

`func (o *AccountControllerResponseData) HasBlobGasFeeCap() bool`

HasBlobGasFeeCap returns a boolean if a field has been set.

### SetBlobGasFeeCapNil

`func (o *AccountControllerResponseData) SetBlobGasFeeCapNil(b bool)`

 SetBlobGasFeeCapNil sets the value for BlobGasFeeCap to be an explicit nil

### UnsetBlobGasFeeCap
`func (o *AccountControllerResponseData) UnsetBlobGasFeeCap()`

UnsetBlobGasFeeCap ensures that no value is present for BlobGasFeeCap, not even an explicit nil
### GetBlobHashes

`func (o *AccountControllerResponseData) GetBlobHashes() []string`

GetBlobHashes returns the BlobHashes field if non-nil, zero value otherwise.

### GetBlobHashesOk

`func (o *AccountControllerResponseData) GetBlobHashesOk() (*[]string, bool)`

GetBlobHashesOk returns a tuple with the BlobHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobHashes

`func (o *AccountControllerResponseData) SetBlobHashes(v []string)`

SetBlobHashes sets BlobHashes field to given value.

### HasBlobHashes

`func (o *AccountControllerResponseData) HasBlobHashes() bool`

HasBlobHashes returns a boolean if a field has been set.

### SetBlobHashesNil

`func (o *AccountControllerResponseData) SetBlobHashesNil(b bool)`

 SetBlobHashesNil sets the value for BlobHashes to be an explicit nil

### UnsetBlobHashes
`func (o *AccountControllerResponseData) UnsetBlobHashes()`

UnsetBlobHashes ensures that no value is present for BlobHashes, not even an explicit nil
### GetV

`func (o *AccountControllerResponseData) GetV() string`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *AccountControllerResponseData) GetVOk() (*string, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *AccountControllerResponseData) SetV(v string)`

SetV sets V field to given value.

### HasV

`func (o *AccountControllerResponseData) HasV() bool`

HasV returns a boolean if a field has been set.

### GetR

`func (o *AccountControllerResponseData) GetR() string`

GetR returns the R field if non-nil, zero value otherwise.

### GetROk

`func (o *AccountControllerResponseData) GetROk() (*string, bool)`

GetROk returns a tuple with the R field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR

`func (o *AccountControllerResponseData) SetR(v string)`

SetR sets R field to given value.

### HasR

`func (o *AccountControllerResponseData) HasR() bool`

HasR returns a boolean if a field has been set.

### GetS

`func (o *AccountControllerResponseData) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *AccountControllerResponseData) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *AccountControllerResponseData) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *AccountControllerResponseData) HasS() bool`

HasS returns a boolean if a field has been set.

### GetSymbol

`func (o *AccountControllerResponseData) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *AccountControllerResponseData) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *AccountControllerResponseData) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *AccountControllerResponseData) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDecimals

`func (o *AccountControllerResponseData) GetDecimals() string`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *AccountControllerResponseData) GetDecimalsOk() (*string, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *AccountControllerResponseData) SetDecimals(v string)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *AccountControllerResponseData) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.

### GetTotalSupply

`func (o *AccountControllerResponseData) GetTotalSupply() string`

GetTotalSupply returns the TotalSupply field if non-nil, zero value otherwise.

### GetTotalSupplyOk

`func (o *AccountControllerResponseData) GetTotalSupplyOk() (*string, bool)`

GetTotalSupplyOk returns a tuple with the TotalSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSupply

`func (o *AccountControllerResponseData) SetTotalSupply(v string)`

SetTotalSupply sets TotalSupply field to given value.

### HasTotalSupply

`func (o *AccountControllerResponseData) HasTotalSupply() bool`

HasTotalSupply returns a boolean if a field has been set.

### GetContractAddress

`func (o *AccountControllerResponseData) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *AccountControllerResponseData) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *AccountControllerResponseData) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.

### HasContractAddress

`func (o *AccountControllerResponseData) HasContractAddress() bool`

HasContractAddress returns a boolean if a field has been set.

### GetBalanceOf

`func (o *AccountControllerResponseData) GetBalanceOf() string`

GetBalanceOf returns the BalanceOf field if non-nil, zero value otherwise.

### GetBalanceOfOk

`func (o *AccountControllerResponseData) GetBalanceOfOk() (*string, bool)`

GetBalanceOfOk returns a tuple with the BalanceOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceOf

`func (o *AccountControllerResponseData) SetBalanceOf(v string)`

SetBalanceOf sets BalanceOf field to given value.

### HasBalanceOf

`func (o *AccountControllerResponseData) HasBalanceOf() bool`

HasBalanceOf returns a boolean if a field has been set.

### GetAllowance

`func (o *AccountControllerResponseData) GetAllowance() string`

GetAllowance returns the Allowance field if non-nil, zero value otherwise.

### GetAllowanceOk

`func (o *AccountControllerResponseData) GetAllowanceOk() (*string, bool)`

GetAllowanceOk returns a tuple with the Allowance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowance

`func (o *AccountControllerResponseData) SetAllowance(v string)`

SetAllowance sets Allowance field to given value.

### HasAllowance

`func (o *AccountControllerResponseData) HasAllowance() bool`

HasAllowance returns a boolean if a field has been set.

### GetBalanceOf

`func (o *AccountControllerResponseData) GetBalanceOf() string`

GetBalanceOf returns the BalanceOf field if non-nil, zero value otherwise.

### GetBalanceOfOk

`func (o *AccountControllerResponseData) GetBalanceOfOk() (*string, bool)`

GetBalanceOfOk returns a tuple with the BalanceOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceOf

`func (o *AccountControllerResponseData) SetBalanceOf(v string)`

SetBalanceOf sets BalanceOf field to given value.

### HasBalanceOf

`func (o *AccountControllerResponseData) HasBalanceOf() bool`

HasBalanceOf returns a boolean if a field has been set.

### GetBalanceOfBatch

`func (o *AccountControllerResponseData) GetBalanceOfBatch() string`

GetBalanceOfBatch returns the BalanceOfBatch field if non-nil, zero value otherwise.

### GetBalanceOfBatchOk

`func (o *AccountControllerResponseData) GetBalanceOfBatchOk() (*string, bool)`

GetBalanceOfBatchOk returns a tuple with the BalanceOfBatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceOfBatch

`func (o *AccountControllerResponseData) SetBalanceOfBatch(v string)`

SetBalanceOfBatch sets BalanceOfBatch field to given value.

### HasBalanceOfBatch

`func (o *AccountControllerResponseData) HasBalanceOfBatch() bool`

HasBalanceOfBatch returns a boolean if a field has been set.

### GetSuccess

`func (o *AccountControllerResponseData) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AccountControllerResponseData) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AccountControllerResponseData) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetMessage

`func (o *AccountControllerResponseData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AccountControllerResponseData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AccountControllerResponseData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSignedTx

`func (o *AccountControllerResponseData) GetSignedTx() string`

GetSignedTx returns the SignedTx field if non-nil, zero value otherwise.

### GetSignedTxOk

`func (o *AccountControllerResponseData) GetSignedTxOk() (*string, bool)`

GetSignedTxOk returns a tuple with the SignedTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedTx

`func (o *AccountControllerResponseData) SetSignedTx(v string)`

SetSignedTx sets SignedTx field to given value.

### HasSignedTx

`func (o *AccountControllerResponseData) HasSignedTx() bool`

HasSignedTx returns a boolean if a field has been set.

### GetOwnerOf

`func (o *AccountControllerResponseData) GetOwnerOf() string`

GetOwnerOf returns the OwnerOf field if non-nil, zero value otherwise.

### GetOwnerOfOk

`func (o *AccountControllerResponseData) GetOwnerOfOk() (*string, bool)`

GetOwnerOfOk returns a tuple with the OwnerOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerOf

`func (o *AccountControllerResponseData) SetOwnerOf(v string)`

SetOwnerOf sets OwnerOf field to given value.

### HasOwnerOf

`func (o *AccountControllerResponseData) HasOwnerOf() bool`

HasOwnerOf returns a boolean if a field has been set.

### GetTokenUri

`func (o *AccountControllerResponseData) GetTokenUri() string`

GetTokenUri returns the TokenUri field if non-nil, zero value otherwise.

### GetTokenUriOk

`func (o *AccountControllerResponseData) GetTokenUriOk() (*string, bool)`

GetTokenUriOk returns a tuple with the TokenUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUri

`func (o *AccountControllerResponseData) SetTokenUri(v string)`

SetTokenUri sets TokenUri field to given value.

### HasTokenUri

`func (o *AccountControllerResponseData) HasTokenUri() bool`

HasTokenUri returns a boolean if a field has been set.

### GetIsApprovedForAll

`func (o *AccountControllerResponseData) GetIsApprovedForAll() string`

GetIsApprovedForAll returns the IsApprovedForAll field if non-nil, zero value otherwise.

### GetIsApprovedForAllOk

`func (o *AccountControllerResponseData) GetIsApprovedForAllOk() (*string, bool)`

GetIsApprovedForAllOk returns a tuple with the IsApprovedForAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApprovedForAll

`func (o *AccountControllerResponseData) SetIsApprovedForAll(v string)`

SetIsApprovedForAll sets IsApprovedForAll field to given value.

### HasIsApprovedForAll

`func (o *AccountControllerResponseData) HasIsApprovedForAll() bool`

HasIsApprovedForAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


