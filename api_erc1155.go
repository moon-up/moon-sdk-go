/*
moon-vault-api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// ERC1155APIService ERC1155API service
type ERC1155APIService service

type ApiBalanceOfRequest struct {
	ctx context.Context
	ApiService *ERC1155APIService
	name string
	authorization *string
	erc1155Request *Erc1155Request
}

func (r ApiBalanceOfRequest) Authorization(authorization string) ApiBalanceOfRequest {
	r.authorization = &authorization
	return r
}

func (r ApiBalanceOfRequest) Erc1155Request(erc1155Request Erc1155Request) ApiBalanceOfRequest {
	r.erc1155Request = &erc1155Request
	return r
}

func (r ApiBalanceOfRequest) Execute() (*AccountControllerResponse, *http.Response, error) {
	return r.ApiService.BalanceOfExecute(r)
}

/*
BalanceOf Method for BalanceOf

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @return ApiBalanceOfRequest
*/
func (a *ERC1155APIService) BalanceOf(ctx context.Context, name string) ApiBalanceOfRequest {
	return ApiBalanceOfRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return AccountControllerResponse
func (a *ERC1155APIService) BalanceOfExecute(r ApiBalanceOfRequest) (*AccountControllerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountControllerResponse
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

type ApiBalanceOfBatchRequest struct {
	ctx context.Context
	ApiService *ERC1155APIService
	name string
	authorization *string
	erc1155Request *Erc1155Request
}

func (r ApiBalanceOfBatchRequest) Authorization(authorization string) ApiBalanceOfBatchRequest {
	r.authorization = &authorization
	return r
}

func (r ApiBalanceOfBatchRequest) Erc1155Request(erc1155Request Erc1155Request) ApiBalanceOfBatchRequest {
	r.erc1155Request = &erc1155Request
	return r
}

func (r ApiBalanceOfBatchRequest) Execute() (*AccountControllerResponse, *http.Response, error) {
	return r.ApiService.BalanceOfBatchExecute(r)
}

/*
BalanceOfBatch Method for BalanceOfBatch

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @return ApiBalanceOfBatchRequest
*/
func (a *ERC1155APIService) BalanceOfBatch(ctx context.Context, name string) ApiBalanceOfBatchRequest {
	return ApiBalanceOfBatchRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return AccountControllerResponse
func (a *ERC1155APIService) BalanceOfBatchExecute(r ApiBalanceOfBatchRequest) (*AccountControllerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountControllerResponse
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

type ApiIsApprovedForAllRequest struct {
	ctx context.Context
	ApiService *ERC1155APIService
	name string
	authorization *string
	erc1155Request *Erc1155Request
}

func (r ApiIsApprovedForAllRequest) Authorization(authorization string) ApiIsApprovedForAllRequest {
	r.authorization = &authorization
	return r
}

func (r ApiIsApprovedForAllRequest) Erc1155Request(erc1155Request Erc1155Request) ApiIsApprovedForAllRequest {
	r.erc1155Request = &erc1155Request
	return r
}

func (r ApiIsApprovedForAllRequest) Execute() (*AccountControllerResponse, *http.Response, error) {
	return r.ApiService.IsApprovedForAllExecute(r)
}

/*
IsApprovedForAll Method for IsApprovedForAll

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @return ApiIsApprovedForAllRequest
*/
func (a *ERC1155APIService) IsApprovedForAll(ctx context.Context, name string) ApiIsApprovedForAllRequest {
	return ApiIsApprovedForAllRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return AccountControllerResponse
func (a *ERC1155APIService) IsApprovedForAllExecute(r ApiIsApprovedForAllRequest) (*AccountControllerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountControllerResponse
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

type ApiSafeBatchTransferFromRequest struct {
	ctx context.Context
	ApiService *ERC1155APIService
	name string
	authorization *string
	erc1155Request *Erc1155Request
}

func (r ApiSafeBatchTransferFromRequest) Authorization(authorization string) ApiSafeBatchTransferFromRequest {
	r.authorization = &authorization
	return r
}

func (r ApiSafeBatchTransferFromRequest) Erc1155Request(erc1155Request Erc1155Request) ApiSafeBatchTransferFromRequest {
	r.erc1155Request = &erc1155Request
	return r
}

func (r ApiSafeBatchTransferFromRequest) Execute() (*AccountControllerResponse, *http.Response, error) {
	return r.ApiService.SafeBatchTransferFromExecute(r)
}

/*
SafeBatchTransferFrom Method for SafeBatchTransferFrom

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @return ApiSafeBatchTransferFromRequest
*/
func (a *ERC1155APIService) SafeBatchTransferFrom(ctx context.Context, name string) ApiSafeBatchTransferFromRequest {
	return ApiSafeBatchTransferFromRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return AccountControllerResponse
func (a *ERC1155APIService) SafeBatchTransferFromExecute(r ApiSafeBatchTransferFromRequest) (*AccountControllerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountControllerResponse
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

type ApiSafeTransferFromRequest struct {
	ctx context.Context
	ApiService *ERC1155APIService
	name string
	authorization *string
	erc1155Request *Erc1155Request
}

func (r ApiSafeTransferFromRequest) Authorization(authorization string) ApiSafeTransferFromRequest {
	r.authorization = &authorization
	return r
}

func (r ApiSafeTransferFromRequest) Erc1155Request(erc1155Request Erc1155Request) ApiSafeTransferFromRequest {
	r.erc1155Request = &erc1155Request
	return r
}

func (r ApiSafeTransferFromRequest) Execute() (*AccountControllerResponse, *http.Response, error) {
	return r.ApiService.SafeTransferFromExecute(r)
}

/*
SafeTransferFrom Method for SafeTransferFrom

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @return ApiSafeTransferFromRequest
*/
func (a *ERC1155APIService) SafeTransferFrom(ctx context.Context, name string) ApiSafeTransferFromRequest {
	return ApiSafeTransferFromRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return AccountControllerResponse
func (a *ERC1155APIService) SafeTransferFromExecute(r ApiSafeTransferFromRequest) (*AccountControllerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountControllerResponse
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

type ApiSetApprovalForAllRequest struct {
	ctx context.Context
	ApiService *ERC1155APIService
	name string
	authorization *string
	erc1155Request *Erc1155Request
}

func (r ApiSetApprovalForAllRequest) Authorization(authorization string) ApiSetApprovalForAllRequest {
	r.authorization = &authorization
	return r
}

func (r ApiSetApprovalForAllRequest) Erc1155Request(erc1155Request Erc1155Request) ApiSetApprovalForAllRequest {
	r.erc1155Request = &erc1155Request
	return r
}

func (r ApiSetApprovalForAllRequest) Execute() (*AccountControllerResponse, *http.Response, error) {
	return r.ApiService.SetApprovalForAllExecute(r)
}

/*
SetApprovalForAll Method for SetApprovalForAll

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @return ApiSetApprovalForAllRequest
*/
func (a *ERC1155APIService) SetApprovalForAll(ctx context.Context, name string) ApiSetApprovalForAllRequest {
	return ApiSetApprovalForAllRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return AccountControllerResponse
func (a *ERC1155APIService) SetApprovalForAllExecute(r ApiSetApprovalForAllRequest) (*AccountControllerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountControllerResponse
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
