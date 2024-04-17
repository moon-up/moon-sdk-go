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


type LitecoinAPI interface {

	/*
	CreateLitecoinAccount Method for CreateLitecoinAccount

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return LitecoinAPICreateLitecoinAccountRequest
	*/
	CreateLitecoinAccount(ctx context.Context) LitecoinAPICreateLitecoinAccountRequest

	// CreateLitecoinAccountExecute executes the request
	//  @return AccountAPIResponse
	CreateLitecoinAccountExecute(r LitecoinAPICreateLitecoinAccountRequest) (*AccountAPIResponse, *http.Response, error)

	/*
	GetLitecoinAccount Method for GetLitecoinAccount

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountName
	@return LitecoinAPIGetLitecoinAccountRequest
	*/
	GetLitecoinAccount(ctx context.Context, accountName string) LitecoinAPIGetLitecoinAccountRequest

	// GetLitecoinAccountExecute executes the request
	//  @return AccountAPIResponse
	GetLitecoinAccountExecute(r LitecoinAPIGetLitecoinAccountRequest) (*AccountAPIResponse, *http.Response, error)

	/*
	ListLitecoinAccounts Method for ListLitecoinAccounts

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return LitecoinAPIListLitecoinAccountsRequest
	*/
	ListLitecoinAccounts(ctx context.Context) LitecoinAPIListLitecoinAccountsRequest

	// ListLitecoinAccountsExecute executes the request
	//  @return AccountAPIResponse
	ListLitecoinAccountsExecute(r LitecoinAPIListLitecoinAccountsRequest) (*AccountAPIResponse, *http.Response, error)

	/*
	SignLitecoinTransaction Method for SignLitecoinTransaction

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountName
	@return LitecoinAPISignLitecoinTransactionRequest
	*/
	SignLitecoinTransaction(ctx context.Context, accountName string) LitecoinAPISignLitecoinTransactionRequest

	// SignLitecoinTransactionExecute executes the request
	//  @return LitecoinAPIResponse
	SignLitecoinTransactionExecute(r LitecoinAPISignLitecoinTransactionRequest) (*LitecoinAPIResponse, *http.Response, error)
}

// LitecoinAPIService LitecoinAPI service
type LitecoinAPIService service

type LitecoinAPICreateLitecoinAccountRequest struct {
	ctx context.Context
	ApiService LitecoinAPI
	authorization *string
	litecoinInput *LitecoinInput
}

func (r LitecoinAPICreateLitecoinAccountRequest) Authorization(authorization string) LitecoinAPICreateLitecoinAccountRequest {
	r.authorization = &authorization
	return r
}

func (r LitecoinAPICreateLitecoinAccountRequest) LitecoinInput(litecoinInput LitecoinInput) LitecoinAPICreateLitecoinAccountRequest {
	r.litecoinInput = &litecoinInput
	return r
}

func (r LitecoinAPICreateLitecoinAccountRequest) Execute() (*AccountAPIResponse, *http.Response, error) {
	return r.ApiService.CreateLitecoinAccountExecute(r)
}

/*
CreateLitecoinAccount Method for CreateLitecoinAccount

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return LitecoinAPICreateLitecoinAccountRequest
*/
func (a *LitecoinAPIService) CreateLitecoinAccount(ctx context.Context) LitecoinAPICreateLitecoinAccountRequest {
	return LitecoinAPICreateLitecoinAccountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AccountAPIResponse
func (a *LitecoinAPIService) CreateLitecoinAccountExecute(r LitecoinAPICreateLitecoinAccountRequest) (*AccountAPIResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountAPIResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LitecoinAPIService.CreateLitecoinAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/litecoin"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.litecoinInput == nil {
		return localVarReturnValue, nil, reportError("litecoinInput is required and must be specified")
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
	localVarPostBody = r.litecoinInput
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

type LitecoinAPIGetLitecoinAccountRequest struct {
	ctx context.Context
	ApiService LitecoinAPI
	authorization *string
	accountName string
}

func (r LitecoinAPIGetLitecoinAccountRequest) Authorization(authorization string) LitecoinAPIGetLitecoinAccountRequest {
	r.authorization = &authorization
	return r
}

func (r LitecoinAPIGetLitecoinAccountRequest) Execute() (*AccountAPIResponse, *http.Response, error) {
	return r.ApiService.GetLitecoinAccountExecute(r)
}

/*
GetLitecoinAccount Method for GetLitecoinAccount

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountName
 @return LitecoinAPIGetLitecoinAccountRequest
*/
func (a *LitecoinAPIService) GetLitecoinAccount(ctx context.Context, accountName string) LitecoinAPIGetLitecoinAccountRequest {
	return LitecoinAPIGetLitecoinAccountRequest{
		ApiService: a,
		ctx: ctx,
		accountName: accountName,
	}
}

// Execute executes the request
//  @return AccountAPIResponse
func (a *LitecoinAPIService) GetLitecoinAccountExecute(r LitecoinAPIGetLitecoinAccountRequest) (*AccountAPIResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountAPIResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LitecoinAPIService.GetLitecoinAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/litecoin/{accountName}"
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

type LitecoinAPIListLitecoinAccountsRequest struct {
	ctx context.Context
	ApiService LitecoinAPI
	authorization *string
}

func (r LitecoinAPIListLitecoinAccountsRequest) Authorization(authorization string) LitecoinAPIListLitecoinAccountsRequest {
	r.authorization = &authorization
	return r
}

func (r LitecoinAPIListLitecoinAccountsRequest) Execute() (*AccountAPIResponse, *http.Response, error) {
	return r.ApiService.ListLitecoinAccountsExecute(r)
}

/*
ListLitecoinAccounts Method for ListLitecoinAccounts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return LitecoinAPIListLitecoinAccountsRequest
*/
func (a *LitecoinAPIService) ListLitecoinAccounts(ctx context.Context) LitecoinAPIListLitecoinAccountsRequest {
	return LitecoinAPIListLitecoinAccountsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AccountAPIResponse
func (a *LitecoinAPIService) ListLitecoinAccountsExecute(r LitecoinAPIListLitecoinAccountsRequest) (*AccountAPIResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountAPIResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LitecoinAPIService.ListLitecoinAccounts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/litecoin"

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

type LitecoinAPISignLitecoinTransactionRequest struct {
	ctx context.Context
	ApiService LitecoinAPI
	authorization *string
	accountName string
	litecoinTransactionInput *LitecoinTransactionInput
}

func (r LitecoinAPISignLitecoinTransactionRequest) Authorization(authorization string) LitecoinAPISignLitecoinTransactionRequest {
	r.authorization = &authorization
	return r
}

func (r LitecoinAPISignLitecoinTransactionRequest) LitecoinTransactionInput(litecoinTransactionInput LitecoinTransactionInput) LitecoinAPISignLitecoinTransactionRequest {
	r.litecoinTransactionInput = &litecoinTransactionInput
	return r
}

func (r LitecoinAPISignLitecoinTransactionRequest) Execute() (*LitecoinAPIResponse, *http.Response, error) {
	return r.ApiService.SignLitecoinTransactionExecute(r)
}

/*
SignLitecoinTransaction Method for SignLitecoinTransaction

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountName
 @return LitecoinAPISignLitecoinTransactionRequest
*/
func (a *LitecoinAPIService) SignLitecoinTransaction(ctx context.Context, accountName string) LitecoinAPISignLitecoinTransactionRequest {
	return LitecoinAPISignLitecoinTransactionRequest{
		ApiService: a,
		ctx: ctx,
		accountName: accountName,
	}
}

// Execute executes the request
//  @return LitecoinAPIResponse
func (a *LitecoinAPIService) SignLitecoinTransactionExecute(r LitecoinAPISignLitecoinTransactionRequest) (*LitecoinAPIResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LitecoinAPIResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LitecoinAPIService.SignLitecoinTransaction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/litecoin/{accountName}/sign-tx"
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", url.PathEscape(parameterValueToString(r.accountName, "accountName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.litecoinTransactionInput == nil {
		return localVarReturnValue, nil, reportError("litecoinTransactionInput is required and must be specified")
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
	localVarPostBody = r.litecoinTransactionInput
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