/*
moon-vault-api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package moonsdk

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type BitcoincashAPI interface {

	/*
	CreateBitcoinCashAccount Method for CreateBitcoinCashAccount

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return BitcoincashAPICreateBitcoinCashAccountRequest
	*/
	CreateBitcoinCashAccount(ctx context.Context) BitcoincashAPICreateBitcoinCashAccountRequest

	// CreateBitcoinCashAccountExecute executes the request
	//  @return AccountAPIResponse
	CreateBitcoinCashAccountExecute(r BitcoincashAPICreateBitcoinCashAccountRequest) (*AccountAPIResponse, *http.Response, error)

	/*
	GetBitcoinCashAccount Method for GetBitcoinCashAccount

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountName
	@return BitcoincashAPIGetBitcoinCashAccountRequest
	*/
	GetBitcoinCashAccount(ctx context.Context, accountName string) BitcoincashAPIGetBitcoinCashAccountRequest

	// GetBitcoinCashAccountExecute executes the request
	//  @return AccountAPIResponse
	GetBitcoinCashAccountExecute(r BitcoincashAPIGetBitcoinCashAccountRequest) (*AccountAPIResponse, *http.Response, error)

	/*
	ListBitcoinCashAccounts Method for ListBitcoinCashAccounts

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return BitcoincashAPIListBitcoinCashAccountsRequest
	*/
	ListBitcoinCashAccounts(ctx context.Context) BitcoincashAPIListBitcoinCashAccountsRequest

	// ListBitcoinCashAccountsExecute executes the request
	//  @return AccountAPIResponse
	ListBitcoinCashAccountsExecute(r BitcoincashAPIListBitcoinCashAccountsRequest) (*AccountAPIResponse, *http.Response, error)

	/*
	SignBitcoinCashTransaction Method for SignBitcoinCashTransaction

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountName
	@return BitcoincashAPISignBitcoinCashTransactionRequest
	*/
	SignBitcoinCashTransaction(ctx context.Context, accountName string) BitcoincashAPISignBitcoinCashTransactionRequest

	// SignBitcoinCashTransactionExecute executes the request
	//  @return BitcoinCashAPIResponse
	SignBitcoinCashTransactionExecute(r BitcoincashAPISignBitcoinCashTransactionRequest) (*BitcoinCashAPIResponse, *http.Response, error)
}

// BitcoincashAPIService BitcoincashAPI service
type BitcoincashAPIService service

type BitcoincashAPICreateBitcoinCashAccountRequest struct {
	ctx context.Context
	ApiService BitcoincashAPI
	authorization *string
	bitcoinCashInput *BitcoinCashInput
}

func (r BitcoincashAPICreateBitcoinCashAccountRequest) Authorization(authorization string) BitcoincashAPICreateBitcoinCashAccountRequest {
	r.authorization = &authorization
	return r
}

func (r BitcoincashAPICreateBitcoinCashAccountRequest) BitcoinCashInput(bitcoinCashInput BitcoinCashInput) BitcoincashAPICreateBitcoinCashAccountRequest {
	r.bitcoinCashInput = &bitcoinCashInput
	return r
}

func (r BitcoincashAPICreateBitcoinCashAccountRequest) Execute() (*AccountAPIResponse, *http.Response, error) {
	return r.ApiService.CreateBitcoinCashAccountExecute(r)
}

/*
CreateBitcoinCashAccount Method for CreateBitcoinCashAccount

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return BitcoincashAPICreateBitcoinCashAccountRequest
*/
func (a *BitcoincashAPIService) CreateBitcoinCashAccount(ctx context.Context) BitcoincashAPICreateBitcoinCashAccountRequest {
	return BitcoincashAPICreateBitcoinCashAccountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AccountAPIResponse
func (a *BitcoincashAPIService) CreateBitcoinCashAccountExecute(r BitcoincashAPICreateBitcoinCashAccountRequest) (*AccountAPIResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountAPIResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BitcoincashAPIService.CreateBitcoinCashAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bitcoincash"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.bitcoinCashInput == nil {
		return localVarReturnValue, nil, reportError("bitcoinCashInput is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	// body params
	localVarPostBody = r.bitcoinCashInput
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type BitcoincashAPIGetBitcoinCashAccountRequest struct {
	ctx context.Context
	ApiService BitcoincashAPI
	authorization *string
	accountName string
}

func (r BitcoincashAPIGetBitcoinCashAccountRequest) Authorization(authorization string) BitcoincashAPIGetBitcoinCashAccountRequest {
	r.authorization = &authorization
	return r
}

func (r BitcoincashAPIGetBitcoinCashAccountRequest) Execute() (*AccountAPIResponse, *http.Response, error) {
	return r.ApiService.GetBitcoinCashAccountExecute(r)
}

/*
GetBitcoinCashAccount Method for GetBitcoinCashAccount

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountName
 @return BitcoincashAPIGetBitcoinCashAccountRequest
*/
func (a *BitcoincashAPIService) GetBitcoinCashAccount(ctx context.Context, accountName string) BitcoincashAPIGetBitcoinCashAccountRequest {
	return BitcoincashAPIGetBitcoinCashAccountRequest{
		ApiService: a,
		ctx: ctx,
		accountName: accountName,
	}
}

// Execute executes the request
//  @return AccountAPIResponse
func (a *BitcoincashAPIService) GetBitcoinCashAccountExecute(r BitcoincashAPIGetBitcoinCashAccountRequest) (*AccountAPIResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountAPIResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BitcoincashAPIService.GetBitcoinCashAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bitcoincash/{accountName}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", url.PathEscape(parameterValueToString(r.accountName, "accountName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type BitcoincashAPIListBitcoinCashAccountsRequest struct {
	ctx context.Context
	ApiService BitcoincashAPI
	authorization *string
}

func (r BitcoincashAPIListBitcoinCashAccountsRequest) Authorization(authorization string) BitcoincashAPIListBitcoinCashAccountsRequest {
	r.authorization = &authorization
	return r
}

func (r BitcoincashAPIListBitcoinCashAccountsRequest) Execute() (*AccountAPIResponse, *http.Response, error) {
	return r.ApiService.ListBitcoinCashAccountsExecute(r)
}

/*
ListBitcoinCashAccounts Method for ListBitcoinCashAccounts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return BitcoincashAPIListBitcoinCashAccountsRequest
*/
func (a *BitcoincashAPIService) ListBitcoinCashAccounts(ctx context.Context) BitcoincashAPIListBitcoinCashAccountsRequest {
	return BitcoincashAPIListBitcoinCashAccountsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AccountAPIResponse
func (a *BitcoincashAPIService) ListBitcoinCashAccountsExecute(r BitcoincashAPIListBitcoinCashAccountsRequest) (*AccountAPIResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountAPIResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BitcoincashAPIService.ListBitcoinCashAccounts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bitcoincash"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type BitcoincashAPISignBitcoinCashTransactionRequest struct {
	ctx context.Context
	ApiService BitcoincashAPI
	authorization *string
	accountName string
	bitcoinCashTransactionInput *BitcoinCashTransactionInput
}

func (r BitcoincashAPISignBitcoinCashTransactionRequest) Authorization(authorization string) BitcoincashAPISignBitcoinCashTransactionRequest {
	r.authorization = &authorization
	return r
}

func (r BitcoincashAPISignBitcoinCashTransactionRequest) BitcoinCashTransactionInput(bitcoinCashTransactionInput BitcoinCashTransactionInput) BitcoincashAPISignBitcoinCashTransactionRequest {
	r.bitcoinCashTransactionInput = &bitcoinCashTransactionInput
	return r
}

func (r BitcoincashAPISignBitcoinCashTransactionRequest) Execute() (*BitcoinCashAPIResponse, *http.Response, error) {
	return r.ApiService.SignBitcoinCashTransactionExecute(r)
}

/*
SignBitcoinCashTransaction Method for SignBitcoinCashTransaction

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountName
 @return BitcoincashAPISignBitcoinCashTransactionRequest
*/
func (a *BitcoincashAPIService) SignBitcoinCashTransaction(ctx context.Context, accountName string) BitcoincashAPISignBitcoinCashTransactionRequest {
	return BitcoincashAPISignBitcoinCashTransactionRequest{
		ApiService: a,
		ctx: ctx,
		accountName: accountName,
	}
}

// Execute executes the request
//  @return BitcoinCashAPIResponse
func (a *BitcoincashAPIService) SignBitcoinCashTransactionExecute(r BitcoincashAPISignBitcoinCashTransactionRequest) (*BitcoinCashAPIResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BitcoinCashAPIResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BitcoincashAPIService.SignBitcoinCashTransaction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bitcoincash/{accountName}/sign-tx"
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", url.PathEscape(parameterValueToString(r.accountName, "accountName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.bitcoinCashTransactionInput == nil {
		return localVarReturnValue, nil, reportError("bitcoinCashTransactionInput is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	// body params
	localVarPostBody = r.bitcoinCashTransactionInput
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
