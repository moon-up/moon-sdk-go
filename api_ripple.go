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


type RippleAPI interface {

	/*
	CreateRippleAccount Method for CreateRippleAccount

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RippleAPICreateRippleAccountRequest
	*/
	CreateRippleAccount(ctx context.Context) RippleAPICreateRippleAccountRequest

	// CreateRippleAccountExecute executes the request
	//  @return AccountAPIResponse
	CreateRippleAccountExecute(r RippleAPICreateRippleAccountRequest) (*AccountAPIResponse, *http.Response, error)

	/*
	GetRippleAccount Method for GetRippleAccount

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountName
	@return RippleAPIGetRippleAccountRequest
	*/
	GetRippleAccount(ctx context.Context, accountName string) RippleAPIGetRippleAccountRequest

	// GetRippleAccountExecute executes the request
	//  @return AccountAPIResponse
	GetRippleAccountExecute(r RippleAPIGetRippleAccountRequest) (*AccountAPIResponse, *http.Response, error)

	/*
	ListRippleAccounts Method for ListRippleAccounts

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RippleAPIListRippleAccountsRequest
	*/
	ListRippleAccounts(ctx context.Context) RippleAPIListRippleAccountsRequest

	// ListRippleAccountsExecute executes the request
	//  @return AccountAPIResponse
	ListRippleAccountsExecute(r RippleAPIListRippleAccountsRequest) (*AccountAPIResponse, *http.Response, error)

	/*
	SignRippleTransaction Method for SignRippleTransaction

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountName
	@return RippleAPISignRippleTransactionRequest
	*/
	SignRippleTransaction(ctx context.Context, accountName string) RippleAPISignRippleTransactionRequest

	// SignRippleTransactionExecute executes the request
	//  @return RippleAPIResponse
	SignRippleTransactionExecute(r RippleAPISignRippleTransactionRequest) (*RippleAPIResponse, *http.Response, error)
}

// RippleAPIService RippleAPI service
type RippleAPIService service

type RippleAPICreateRippleAccountRequest struct {
	ctx context.Context
	ApiService RippleAPI
	authorization *string
	rippleInput *RippleInput
}

func (r RippleAPICreateRippleAccountRequest) Authorization(authorization string) RippleAPICreateRippleAccountRequest {
	r.authorization = &authorization
	return r
}

func (r RippleAPICreateRippleAccountRequest) RippleInput(rippleInput RippleInput) RippleAPICreateRippleAccountRequest {
	r.rippleInput = &rippleInput
	return r
}

func (r RippleAPICreateRippleAccountRequest) Execute() (*AccountAPIResponse, *http.Response, error) {
	return r.ApiService.CreateRippleAccountExecute(r)
}

/*
CreateRippleAccount Method for CreateRippleAccount

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RippleAPICreateRippleAccountRequest
*/
func (a *RippleAPIService) CreateRippleAccount(ctx context.Context) RippleAPICreateRippleAccountRequest {
	return RippleAPICreateRippleAccountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AccountAPIResponse
func (a *RippleAPIService) CreateRippleAccountExecute(r RippleAPICreateRippleAccountRequest) (*AccountAPIResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountAPIResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RippleAPIService.CreateRippleAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ripple"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.rippleInput == nil {
		return localVarReturnValue, nil, reportError("rippleInput is required and must be specified")
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
	localVarPostBody = r.rippleInput
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

type RippleAPIGetRippleAccountRequest struct {
	ctx context.Context
	ApiService RippleAPI
	authorization *string
	accountName string
}

func (r RippleAPIGetRippleAccountRequest) Authorization(authorization string) RippleAPIGetRippleAccountRequest {
	r.authorization = &authorization
	return r
}

func (r RippleAPIGetRippleAccountRequest) Execute() (*AccountAPIResponse, *http.Response, error) {
	return r.ApiService.GetRippleAccountExecute(r)
}

/*
GetRippleAccount Method for GetRippleAccount

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountName
 @return RippleAPIGetRippleAccountRequest
*/
func (a *RippleAPIService) GetRippleAccount(ctx context.Context, accountName string) RippleAPIGetRippleAccountRequest {
	return RippleAPIGetRippleAccountRequest{
		ApiService: a,
		ctx: ctx,
		accountName: accountName,
	}
}

// Execute executes the request
//  @return AccountAPIResponse
func (a *RippleAPIService) GetRippleAccountExecute(r RippleAPIGetRippleAccountRequest) (*AccountAPIResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountAPIResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RippleAPIService.GetRippleAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ripple/{accountName}"
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

type RippleAPIListRippleAccountsRequest struct {
	ctx context.Context
	ApiService RippleAPI
	authorization *string
}

func (r RippleAPIListRippleAccountsRequest) Authorization(authorization string) RippleAPIListRippleAccountsRequest {
	r.authorization = &authorization
	return r
}

func (r RippleAPIListRippleAccountsRequest) Execute() (*AccountAPIResponse, *http.Response, error) {
	return r.ApiService.ListRippleAccountsExecute(r)
}

/*
ListRippleAccounts Method for ListRippleAccounts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RippleAPIListRippleAccountsRequest
*/
func (a *RippleAPIService) ListRippleAccounts(ctx context.Context) RippleAPIListRippleAccountsRequest {
	return RippleAPIListRippleAccountsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AccountAPIResponse
func (a *RippleAPIService) ListRippleAccountsExecute(r RippleAPIListRippleAccountsRequest) (*AccountAPIResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountAPIResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RippleAPIService.ListRippleAccounts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ripple"

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

type RippleAPISignRippleTransactionRequest struct {
	ctx context.Context
	ApiService RippleAPI
	authorization *string
	accountName string
	rippleTransactionInput *RippleTransactionInput
}

func (r RippleAPISignRippleTransactionRequest) Authorization(authorization string) RippleAPISignRippleTransactionRequest {
	r.authorization = &authorization
	return r
}

func (r RippleAPISignRippleTransactionRequest) RippleTransactionInput(rippleTransactionInput RippleTransactionInput) RippleAPISignRippleTransactionRequest {
	r.rippleTransactionInput = &rippleTransactionInput
	return r
}

func (r RippleAPISignRippleTransactionRequest) Execute() (*RippleAPIResponse, *http.Response, error) {
	return r.ApiService.SignRippleTransactionExecute(r)
}

/*
SignRippleTransaction Method for SignRippleTransaction

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountName
 @return RippleAPISignRippleTransactionRequest
*/
func (a *RippleAPIService) SignRippleTransaction(ctx context.Context, accountName string) RippleAPISignRippleTransactionRequest {
	return RippleAPISignRippleTransactionRequest{
		ApiService: a,
		ctx: ctx,
		accountName: accountName,
	}
}

// Execute executes the request
//  @return RippleAPIResponse
func (a *RippleAPIService) SignRippleTransactionExecute(r RippleAPISignRippleTransactionRequest) (*RippleAPIResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RippleAPIResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RippleAPIService.SignRippleTransaction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ripple/{accountName}/sign-tx"
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", url.PathEscape(parameterValueToString(r.accountName, "accountName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.rippleTransactionInput == nil {
		return localVarReturnValue, nil, reportError("rippleTransactionInput is required and must be specified")
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
	localVarPostBody = r.rippleTransactionInput
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
