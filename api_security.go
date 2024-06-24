/*
Cat as a service (CATAAS)

Cat as a service (CATAAS) is a REST API to spread peace and love (or not) thanks to cats.

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
)

// SecurityAPIService SecurityAPI service
type SecurityAPIService service

type ApiApiSecuritySignInRequest struct {
	ctx                      context.Context
	ApiService               *SecurityAPIService
	apiSecuritySignInRequest *ApiSecuritySignInRequest
}

func (r ApiApiSecuritySignInRequest) ApiSecuritySignInRequest(apiSecuritySignInRequest ApiSecuritySignInRequest) ApiApiSecuritySignInRequest {
	r.apiSecuritySignInRequest = &apiSecuritySignInRequest
	return r
}

func (r ApiApiSecuritySignInRequest) Execute() (*ApiSecuritySignIn200Response, *http.Response, error) {
	return r.ApiService.ApiSecuritySignInExecute(r)
}

/*
ApiSecuritySignIn Method for ApiSecuritySignIn

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiSecuritySignInRequest
*/
func (a *SecurityAPIService) ApiSecuritySignIn(ctx context.Context) ApiApiSecuritySignInRequest {
	return ApiApiSecuritySignInRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ApiSecuritySignIn200Response
func (a *SecurityAPIService) ApiSecuritySignInExecute(r ApiApiSecuritySignInRequest) (*ApiSecuritySignIn200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ApiSecuritySignIn200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityAPIService.ApiSecuritySignIn")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/security/sign-in"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	// body params
	localVarPostBody = r.apiSecuritySignInRequest
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

type ApiApiSecuritySignUpRequest struct {
	ctx                      context.Context
	ApiService               *SecurityAPIService
	apiSecuritySignUpRequest *ApiSecuritySignUpRequest
}

func (r ApiApiSecuritySignUpRequest) ApiSecuritySignUpRequest(apiSecuritySignUpRequest ApiSecuritySignUpRequest) ApiApiSecuritySignUpRequest {
	r.apiSecuritySignUpRequest = &apiSecuritySignUpRequest
	return r
}

func (r ApiApiSecuritySignUpRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiSecuritySignUpExecute(r)
}

/*
ApiSecuritySignUp Method for ApiSecuritySignUp

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiSecuritySignUpRequest
*/
func (a *SecurityAPIService) ApiSecuritySignUp(ctx context.Context) ApiApiSecuritySignUpRequest {
	return ApiApiSecuritySignUpRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *SecurityAPIService) ApiSecuritySignUpExecute(r ApiApiSecuritySignUpRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityAPIService.ApiSecuritySignUp")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/security/sign-up"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.apiSecuritySignUpRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
