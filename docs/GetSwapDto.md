# GetSwapDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Src** | **string** |  | 
**Dst** | **string** |  | 
**Amount** | **string** |  | 
**From** | **string** |  | 
**Slippage** | **float64** |  | 
**Protocols** | Pointer to **string** |  | [optional] 
**Fee** | Pointer to **string** |  | [optional] 
**DisableEstimate** | Pointer to **bool** |  | [optional] 
**Permit** | Pointer to **string** |  | [optional] 
**IncludeTokensInfo** | Pointer to **bool** |  | [optional] 
**IncludeProtocols** | Pointer to **bool** |  | [optional] 
**Compatibility** | Pointer to **bool** |  | [optional] 
**AllowPartialFill** | Pointer to **bool** |  | [optional] 
**Parts** | Pointer to **string** |  | [optional] 
**MainRouteParts** | Pointer to **string** |  | [optional] 
**ConnectorTokens** | Pointer to **string** |  | [optional] 
**ComplexityLevel** | Pointer to **string** |  | [optional] 
**GasLimit** | Pointer to **string** |  | [optional] 
**GasPrice** | Pointer to **string** |  | [optional] 
**Referrer** | Pointer to **string** |  | [optional] 
**Receiver** | Pointer to **string** |  | [optional] 
**ChainId** | Pointer to **float64** |  | [optional] 

## Methods

### NewGetSwapDto

`func NewGetSwapDto(src string, dst string, amount string, from string, slippage float64, ) *GetSwapDto`

NewGetSwapDto instantiates a new GetSwapDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSwapDtoWithDefaults

`func NewGetSwapDtoWithDefaults() *GetSwapDto`

NewGetSwapDtoWithDefaults instantiates a new GetSwapDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSrc

`func (o *GetSwapDto) GetSrc() string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *GetSwapDto) GetSrcOk() (*string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *GetSwapDto) SetSrc(v string)`

SetSrc sets Src field to given value.


### GetDst

`func (o *GetSwapDto) GetDst() string`

GetDst returns the Dst field if non-nil, zero value otherwise.

### GetDstOk

`func (o *GetSwapDto) GetDstOk() (*string, bool)`

GetDstOk returns a tuple with the Dst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDst

`func (o *GetSwapDto) SetDst(v string)`

SetDst sets Dst field to given value.


### GetAmount

`func (o *GetSwapDto) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetSwapDto) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetSwapDto) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetFrom

`func (o *GetSwapDto) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GetSwapDto) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GetSwapDto) SetFrom(v string)`

SetFrom sets From field to given value.


### GetSlippage

`func (o *GetSwapDto) GetSlippage() float64`

GetSlippage returns the Slippage field if non-nil, zero value otherwise.

### GetSlippageOk

`func (o *GetSwapDto) GetSlippageOk() (*float64, bool)`

GetSlippageOk returns a tuple with the Slippage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlippage

`func (o *GetSwapDto) SetSlippage(v float64)`

SetSlippage sets Slippage field to given value.


### GetProtocols

`func (o *GetSwapDto) GetProtocols() string`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *GetSwapDto) GetProtocolsOk() (*string, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *GetSwapDto) SetProtocols(v string)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *GetSwapDto) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.

### GetFee

`func (o *GetSwapDto) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetSwapDto) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetSwapDto) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GetSwapDto) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetDisableEstimate

`func (o *GetSwapDto) GetDisableEstimate() bool`

GetDisableEstimate returns the DisableEstimate field if non-nil, zero value otherwise.

### GetDisableEstimateOk

`func (o *GetSwapDto) GetDisableEstimateOk() (*bool, bool)`

GetDisableEstimateOk returns a tuple with the DisableEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEstimate

`func (o *GetSwapDto) SetDisableEstimate(v bool)`

SetDisableEstimate sets DisableEstimate field to given value.

### HasDisableEstimate

`func (o *GetSwapDto) HasDisableEstimate() bool`

HasDisableEstimate returns a boolean if a field has been set.

### GetPermit

`func (o *GetSwapDto) GetPermit() string`

GetPermit returns the Permit field if non-nil, zero value otherwise.

### GetPermitOk

`func (o *GetSwapDto) GetPermitOk() (*string, bool)`

GetPermitOk returns a tuple with the Permit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermit

`func (o *GetSwapDto) SetPermit(v string)`

SetPermit sets Permit field to given value.

### HasPermit

`func (o *GetSwapDto) HasPermit() bool`

HasPermit returns a boolean if a field has been set.

### GetIncludeTokensInfo

`func (o *GetSwapDto) GetIncludeTokensInfo() bool`

GetIncludeTokensInfo returns the IncludeTokensInfo field if non-nil, zero value otherwise.

### GetIncludeTokensInfoOk

`func (o *GetSwapDto) GetIncludeTokensInfoOk() (*bool, bool)`

GetIncludeTokensInfoOk returns a tuple with the IncludeTokensInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTokensInfo

`func (o *GetSwapDto) SetIncludeTokensInfo(v bool)`

SetIncludeTokensInfo sets IncludeTokensInfo field to given value.

### HasIncludeTokensInfo

`func (o *GetSwapDto) HasIncludeTokensInfo() bool`

HasIncludeTokensInfo returns a boolean if a field has been set.

### GetIncludeProtocols

`func (o *GetSwapDto) GetIncludeProtocols() bool`

GetIncludeProtocols returns the IncludeProtocols field if non-nil, zero value otherwise.

### GetIncludeProtocolsOk

`func (o *GetSwapDto) GetIncludeProtocolsOk() (*bool, bool)`

GetIncludeProtocolsOk returns a tuple with the IncludeProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeProtocols

`func (o *GetSwapDto) SetIncludeProtocols(v bool)`

SetIncludeProtocols sets IncludeProtocols field to given value.

### HasIncludeProtocols

`func (o *GetSwapDto) HasIncludeProtocols() bool`

HasIncludeProtocols returns a boolean if a field has been set.

### GetCompatibility

`func (o *GetSwapDto) GetCompatibility() bool`

GetCompatibility returns the Compatibility field if non-nil, zero value otherwise.

### GetCompatibilityOk

`func (o *GetSwapDto) GetCompatibilityOk() (*bool, bool)`

GetCompatibilityOk returns a tuple with the Compatibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibility

`func (o *GetSwapDto) SetCompatibility(v bool)`

SetCompatibility sets Compatibility field to given value.

### HasCompatibility

`func (o *GetSwapDto) HasCompatibility() bool`

HasCompatibility returns a boolean if a field has been set.

### GetAllowPartialFill

`func (o *GetSwapDto) GetAllowPartialFill() bool`

GetAllowPartialFill returns the AllowPartialFill field if non-nil, zero value otherwise.

### GetAllowPartialFillOk

`func (o *GetSwapDto) GetAllowPartialFillOk() (*bool, bool)`

GetAllowPartialFillOk returns a tuple with the AllowPartialFill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPartialFill

`func (o *GetSwapDto) SetAllowPartialFill(v bool)`

SetAllowPartialFill sets AllowPartialFill field to given value.

### HasAllowPartialFill

`func (o *GetSwapDto) HasAllowPartialFill() bool`

HasAllowPartialFill returns a boolean if a field has been set.

### GetParts

`func (o *GetSwapDto) GetParts() string`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *GetSwapDto) GetPartsOk() (*string, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *GetSwapDto) SetParts(v string)`

SetParts sets Parts field to given value.

### HasParts

`func (o *GetSwapDto) HasParts() bool`

HasParts returns a boolean if a field has been set.

### GetMainRouteParts

`func (o *GetSwapDto) GetMainRouteParts() string`

GetMainRouteParts returns the MainRouteParts field if non-nil, zero value otherwise.

### GetMainRoutePartsOk

`func (o *GetSwapDto) GetMainRoutePartsOk() (*string, bool)`

GetMainRoutePartsOk returns a tuple with the MainRouteParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainRouteParts

`func (o *GetSwapDto) SetMainRouteParts(v string)`

SetMainRouteParts sets MainRouteParts field to given value.

### HasMainRouteParts

`func (o *GetSwapDto) HasMainRouteParts() bool`

HasMainRouteParts returns a boolean if a field has been set.

### GetConnectorTokens

`func (o *GetSwapDto) GetConnectorTokens() string`

GetConnectorTokens returns the ConnectorTokens field if non-nil, zero value otherwise.

### GetConnectorTokensOk

`func (o *GetSwapDto) GetConnectorTokensOk() (*string, bool)`

GetConnectorTokensOk returns a tuple with the ConnectorTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorTokens

`func (o *GetSwapDto) SetConnectorTokens(v string)`

SetConnectorTokens sets ConnectorTokens field to given value.

### HasConnectorTokens

`func (o *GetSwapDto) HasConnectorTokens() bool`

HasConnectorTokens returns a boolean if a field has been set.

### GetComplexityLevel

`func (o *GetSwapDto) GetComplexityLevel() string`

GetComplexityLevel returns the ComplexityLevel field if non-nil, zero value otherwise.

### GetComplexityLevelOk

`func (o *GetSwapDto) GetComplexityLevelOk() (*string, bool)`

GetComplexityLevelOk returns a tuple with the ComplexityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplexityLevel

`func (o *GetSwapDto) SetComplexityLevel(v string)`

SetComplexityLevel sets ComplexityLevel field to given value.

### HasComplexityLevel

`func (o *GetSwapDto) HasComplexityLevel() bool`

HasComplexityLevel returns a boolean if a field has been set.

### GetGasLimit

`func (o *GetSwapDto) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *GetSwapDto) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *GetSwapDto) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.

### HasGasLimit

`func (o *GetSwapDto) HasGasLimit() bool`

HasGasLimit returns a boolean if a field has been set.

### GetGasPrice

`func (o *GetSwapDto) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *GetSwapDto) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *GetSwapDto) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *GetSwapDto) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetReferrer

`func (o *GetSwapDto) GetReferrer() string`

GetReferrer returns the Referrer field if non-nil, zero value otherwise.

### GetReferrerOk

`func (o *GetSwapDto) GetReferrerOk() (*string, bool)`

GetReferrerOk returns a tuple with the Referrer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrer

`func (o *GetSwapDto) SetReferrer(v string)`

SetReferrer sets Referrer field to given value.

### HasReferrer

`func (o *GetSwapDto) HasReferrer() bool`

HasReferrer returns a boolean if a field has been set.

### GetReceiver

`func (o *GetSwapDto) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *GetSwapDto) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *GetSwapDto) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *GetSwapDto) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetChainId

`func (o *GetSwapDto) GetChainId() float64`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *GetSwapDto) GetChainIdOk() (*float64, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *GetSwapDto) SetChainId(v float64)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *GetSwapDto) HasChainId() bool`

HasChainId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


