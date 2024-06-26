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


type ERC1155API interface {

	/*
	BalanceOf Method for BalanceOf

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name
	@return ERC1155APIBalanceOfRequest
	*/
	BalanceOf(ctx context.Context, name string) ERC1155APIBalanceOfRequest

	// BalanceOfExecute executes the request
	//  @return TransactionAPIResponse
	BalanceOfExecute(r ERC1155APIBalanceOfRequest) (*TransactionAPIResponse, *http.Response, error)

	/*
	BalanceOfBatch Method for BalanceOfBatch

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name
	@return ERC1155APIBalanceOfBatchRequest
	*/
	BalanceOfBatch(ctx context.Context, name string) ERC1155APIBalanceOfBatchRequest

	// BalanceOfBatchExecute executes the request
	//  @return TransactionAPIResponse
	BalanceOfBatchExecute(r ERC1155APIBalanceOfBatchRequest) (*TransactionAPIResponse, *http.Response, error)

	/*
	IsApprovedForAll Method for IsApprovedForAll

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name
	@return ERC1155APIIsApprovedForAllRequest
	*/
	IsApprovedForAll(ctx context.Context, name string) ERC1155APIIsApprovedForAllRequest

	// IsApprovedForAllExecute executes the request
	//  @return TransactionAPIResponse
	IsApprovedForAllExecute(r ERC1155APIIsApprovedForAllRequest) (*TransactionAPIResponse, *http.Response, error)

	/*
	SafeBatchTransferFrom Method for SafeBatchTransferFrom

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name
	@return ERC1155APISafeBatchTransferFromRequest
	*/
	SafeBatchTransferFrom(ctx context.Context, name string) ERC1155APISafeBatchTransferFromRequest

	// SafeBatchTransferFromExecute executes the request
	//  @return TransactionAPIResponse
	SafeBatchTransferFromExecute(r ERC1155APISafeBatchTransferFromRequest) (*TransactionAPIResponse, *http.Response, error)

	/*
	SafeTransferFrom Method for SafeTransferFrom

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name
	@return ERC1155APISafeTransferFromRequest
	*/
	SafeTransferFrom(ctx context.Context, name string) ERC1155APISafeTransferFromRequest

	// SafeTransferFromExecute executes the request
	//  @return TransactionAPIResponse
	SafeTransferFromExecute(r ERC1155APISafeTransferFromRequest) (*TransactionAPIResponse, *http.Response, error)

	/*
	SetApprovalForAll Method for SetApprovalForAll

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name
	@return ERC1155APISetApprovalForAllRequest
	*/
	SetApprovalForAll(ctx context.Context, name string) ERC1155APISetApprovalForAllRequest

	// SetApprovalForAllExecute executes the request
	//  @return TransactionAPIResponse
	SetApprovalForAllExecute(r ERC1155APISetApprovalForAllRequest) (*TransactionAPIResponse, *http.Response, error)
}

// ERC1155APIService ERC1155API service
type ERC1155APIService service

type ERC1155APIBalanceOfRequest struct {
	ctx context.Context
	ApiService ERC1155API
	name string
	authorization *string
	erc1155Request *Erc1155Request
}

func (r ERC1155APIBalanceOfRequest) Authorization(authorization string) ERC1155APIBalanceOfRequest {
	r.authorization = &authorization
	return r
}

func (r ERC1155APIBalanceOfRequest) Erc1155Request(erc1155Request Erc1155Request) ERC1155APIBalanceOfRequest {
	r.erc1155Request = &erc1155Request
	return r
}

func (r ERC1155APIBalanceOfRequest) Execute() (*TransactionAPIResponse, *http.Response, error) {
	return r.ApiService.BalanceOfExecute(r)
}

/*
BalanceOf Method for BalanceOf

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @return ERC1155APIBalanceOfRequest
*/
func (a *ERC1155APIService) BalanceOf(ctx context.Context, name string) ERC1155APIBalanceOfRequest {
	return ERC1155APIBalanceOfRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return TransactionAPIResponse
func (a *ERC1155APIService) BalanceOfExecute(r ERC1155APIBalanceOfRequest) (*TransactionAPIResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TransactionAPIResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ERC1155APIService.BalanceOf")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/erc1155/{name}/balance-of"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.erc1155Request == nil {
		return localVarReturnValue, nil, reportError("erc1155Request is required and must be specified")
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
	localVarPostBody = r.erc1155Request
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

type ERC1155APIBalanceOfBatchRequest struct {
	ctx context.Context
	ApiService ERC1155API
	name string
	authorization *string
	erc1155Request *Erc1155Request
}

func (r ERC1155APIBalanceOfBatchRequest) Authorization(authorization string) ERC1155APIBalanceOfBatchRequest {
	r.authorization = &authorization
	return r
}

func (r ERC1155APIBalanceOfBatchRequest) Erc1155Request(erc1155Request Erc1155Request) ERC1155APIBalanceOfBatchRequest {
	r.erc1155Request = &erc1155Request
	return r
}

func (r ERC1155APIBalanceOfBatchRequest) Execute() (*TransactionAPIResponse, *http.Response, error) {
	return r.ApiService.BalanceOfBatchExecute(r)
}

/*
BalanceOfBatch Method for BalanceOfBatch

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @return ERC1155APIBalanceOfBatchRequest
*/
func (a *ERC1155APIService) BalanceOfBatch(ctx context.Context, name string) ERC1155APIBalanceOfBatchRequest {
	return ERC1155APIBalanceOfBatchRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return TransactionAPIResponse
func (a *ERC1155APIService) BalanceOfBatchExecute(r ERC1155APIBalanceOfBatchRequest) (*TransactionAPIResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TransactionAPIResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ERC1155APIService.BalanceOfBatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/erc1155/{name}/balance-of-batch"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.erc1155Request == nil {
		return localVarReturnValue, nil, reportError("erc1155Request is required and must be specified")
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
	localVarPostBody = r.erc1155Request
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

type ERC1155APIIsApprovedForAllRequest struct {
	ctx context.Context
	ApiService ERC1155API
	name string
	authorization *string
	erc1155Request *Erc1155Request
}

func (r ERC1155APIIsApprovedForAllRequest) Authorization(authorization string) ERC1155APIIsApprovedForAllRequest {
	r.authorization = &authorization
	return r
}

func (r ERC1155APIIsApprovedForAllRequest) Erc1155Request(erc1155Request Erc1155Request) ERC1155APIIsApprovedForAllRequest {
	r.erc1155Request = &erc1155Request
	return r
}

func (r ERC1155APIIsApprovedForAllRequest) Execute() (*TransactionAPIResponse, *http.Response, error) {
	return r.ApiService.IsApprovedForAllExecute(r)
}

/*
IsApprovedForAll Method for IsApprovedForAll

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @return ERC1155APIIsApprovedForAllRequest
*/
func (a *ERC1155APIService) IsApprovedForAll(ctx context.Context, name string) ERC1155APIIsApprovedForAllRequest {
	return ERC1155APIIsApprovedForAllRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return TransactionAPIResponse
func (a *ERC1155APIService) IsApprovedForAllExecute(r ERC1155APIIsApprovedForAllRequest) (*TransactionAPIResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TransactionAPIResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ERC1155APIService.IsApprovedForAll")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/erc1155/{name}/is-approved-for-all"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.erc1155Request == nil {
		return localVarReturnValue, nil, reportError("erc1155Request is required and must be specified")
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
	localVarPostBody = r.erc1155Request
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

type ERC1155APISafeBatchTransferFromRequest struct {
	ctx context.Context
	ApiService ERC1155API
	name string
	authorization *string
	erc1155Request *Erc1155Request
}

func (r ERC1155APISafeBatchTransferFromRequest) Authorization(authorization string) ERC1155APISafeBatchTransferFromRequest {
	r.authorization = &authorization
	return r
}

func (r ERC1155APISafeBatchTransferFromRequest) Erc1155Request(erc1155Request Erc1155Request) ERC1155APISafeBatchTransferFromRequest {
	r.erc1155Request = &erc1155Request
	return r
}

func (r ERC1155APISafeBatchTransferFromRequest) Execute() (*TransactionAPIResponse, *http.Response, error) {
	return r.ApiService.SafeBatchTransferFromExecute(r)
}

/*
SafeBatchTransferFrom Method for SafeBatchTransferFrom

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @return ERC1155APISafeBatchTransferFromRequest
*/
func (a *ERC1155APIService) SafeBatchTransferFrom(ctx context.Context, name string) ERC1155APISafeBatchTransferFromRequest {
	return ERC1155APISafeBatchTransferFromRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return TransactionAPIResponse
func (a *ERC1155APIService) SafeBatchTransferFromExecute(r ERC1155APISafeBatchTransferFromRequest) (*TransactionAPIResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TransactionAPIResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ERC1155APIService.SafeBatchTransferFrom")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/erc1155/{name}/safe-batch-transfer-from"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.erc1155Request == nil {
		return localVarReturnValue, nil, reportError("erc1155Request is required and must be specified")
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
	localVarPostBody = r.erc1155Request
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

type ERC1155APISafeTransferFromRequest struct {
	ctx context.Context
	ApiService ERC1155API
	name string
	authorization *string
	erc1155Request *Erc1155Request
}

func (r ERC1155APISafeTransferFromRequest) Authorization(authorization string) ERC1155APISafeTransferFromRequest {
	r.authorization = &authorization
	return r
}

func (r ERC1155APISafeTransferFromRequest) Erc1155Request(erc1155Request Erc1155Request) ERC1155APISafeTransferFromRequest {
	r.erc1155Request = &erc1155Request
	return r
}

func (r ERC1155APISafeTransferFromRequest) Execute() (*TransactionAPIResponse, *http.Response, error) {
	return r.ApiService.SafeTransferFromExecute(r)
}

/*
SafeTransferFrom Method for SafeTransferFrom

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @return ERC1155APISafeTransferFromRequest
*/
func (a *ERC1155APIService) SafeTransferFrom(ctx context.Context, name string) ERC1155APISafeTransferFromRequest {
	return ERC1155APISafeTransferFromRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return TransactionAPIResponse
func (a *ERC1155APIService) SafeTransferFromExecute(r ERC1155APISafeTransferFromRequest) (*TransactionAPIResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TransactionAPIResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ERC1155APIService.SafeTransferFrom")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/erc1155/{name}/safe-transfer-from"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.erc1155Request == nil {
		return localVarReturnValue, nil, reportError("erc1155Request is required and must be specified")
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
	localVarPostBody = r.erc1155Request
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

type ERC1155APISetApprovalForAllRequest struct {
	ctx context.Context
	ApiService ERC1155API
	name string
	authorization *string
	erc1155Request *Erc1155Request
}

func (r ERC1155APISetApprovalForAllRequest) Authorization(authorization string) ERC1155APISetApprovalForAllRequest {
	r.authorization = &authorization
	return r
}

func (r ERC1155APISetApprovalForAllRequest) Erc1155Request(erc1155Request Erc1155Request) ERC1155APISetApprovalForAllRequest {
	r.erc1155Request = &erc1155Request
	return r
}

func (r ERC1155APISetApprovalForAllRequest) Execute() (*TransactionAPIResponse, *http.Response, error) {
	return r.ApiService.SetApprovalForAllExecute(r)
}

/*
SetApprovalForAll Method for SetApprovalForAll

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @return ERC1155APISetApprovalForAllRequest
*/
func (a *ERC1155APIService) SetApprovalForAll(ctx context.Context, name string) ERC1155APISetApprovalForAllRequest {
	return ERC1155APISetApprovalForAllRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return TransactionAPIResponse
func (a *ERC1155APIService) SetApprovalForAllExecute(r ERC1155APISetApprovalForAllRequest) (*TransactionAPIResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TransactionAPIResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ERC1155APIService.SetApprovalForAll")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/erc1155/{name}/set-approval-for-all"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.erc1155Request == nil {
		return localVarReturnValue, nil, reportError("erc1155Request is required and must be specified")
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
	localVarPostBody = r.erc1155Request
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
