/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// AppEncryptionDeclarationsApiService AppEncryptionDeclarationsApi service
type AppEncryptionDeclarationsApiService service

type ApiAppEncryptionDeclarationsAppGetToOneRelatedRequest struct {
	ctx _context.Context
	ApiService *AppEncryptionDeclarationsApiService
	id string
	fieldsApps *[]string
}

func (r ApiAppEncryptionDeclarationsAppGetToOneRelatedRequest) FieldsApps(fieldsApps []string) ApiAppEncryptionDeclarationsAppGetToOneRelatedRequest {
	r.fieldsApps = &fieldsApps
	return r
}

func (r ApiAppEncryptionDeclarationsAppGetToOneRelatedRequest) Execute() (AppResponse, *_nethttp.Response, error) {
	return r.ApiService.AppEncryptionDeclarationsAppGetToOneRelatedExecute(r)
}

/*
 * AppEncryptionDeclarationsAppGetToOneRelated Method for AppEncryptionDeclarationsAppGetToOneRelated
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id the id of the requested resource
 * @return ApiAppEncryptionDeclarationsAppGetToOneRelatedRequest
 */
func (a *AppEncryptionDeclarationsApiService) AppEncryptionDeclarationsAppGetToOneRelated(ctx _context.Context, id string) ApiAppEncryptionDeclarationsAppGetToOneRelatedRequest {
	return ApiAppEncryptionDeclarationsAppGetToOneRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return AppResponse
 */
func (a *AppEncryptionDeclarationsApiService) AppEncryptionDeclarationsAppGetToOneRelatedExecute(r ApiAppEncryptionDeclarationsAppGetToOneRelatedRequest) (AppResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AppResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppEncryptionDeclarationsApiService.AppEncryptionDeclarationsAppGetToOneRelated")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appEncryptionDeclarations/{id}/app"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.fieldsApps != nil {
		localVarQueryParams.Add("fields[apps]", parameterToString(*r.fieldsApps, "csv"))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiAppEncryptionDeclarationsBuildsCreateToManyRelationshipRequest struct {
	ctx _context.Context
	ApiService *AppEncryptionDeclarationsApiService
	id string
	appEncryptionDeclarationBuildsLinkagesRequest *AppEncryptionDeclarationBuildsLinkagesRequest
}

func (r ApiAppEncryptionDeclarationsBuildsCreateToManyRelationshipRequest) AppEncryptionDeclarationBuildsLinkagesRequest(appEncryptionDeclarationBuildsLinkagesRequest AppEncryptionDeclarationBuildsLinkagesRequest) ApiAppEncryptionDeclarationsBuildsCreateToManyRelationshipRequest {
	r.appEncryptionDeclarationBuildsLinkagesRequest = &appEncryptionDeclarationBuildsLinkagesRequest
	return r
}

func (r ApiAppEncryptionDeclarationsBuildsCreateToManyRelationshipRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.AppEncryptionDeclarationsBuildsCreateToManyRelationshipExecute(r)
}

/*
 * AppEncryptionDeclarationsBuildsCreateToManyRelationship Method for AppEncryptionDeclarationsBuildsCreateToManyRelationship
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id the id of the requested resource
 * @return ApiAppEncryptionDeclarationsBuildsCreateToManyRelationshipRequest
 */
func (a *AppEncryptionDeclarationsApiService) AppEncryptionDeclarationsBuildsCreateToManyRelationship(ctx _context.Context, id string) ApiAppEncryptionDeclarationsBuildsCreateToManyRelationshipRequest {
	return ApiAppEncryptionDeclarationsBuildsCreateToManyRelationshipRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 */
func (a *AppEncryptionDeclarationsApiService) AppEncryptionDeclarationsBuildsCreateToManyRelationshipExecute(r ApiAppEncryptionDeclarationsBuildsCreateToManyRelationshipRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppEncryptionDeclarationsApiService.AppEncryptionDeclarationsBuildsCreateToManyRelationship")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appEncryptionDeclarations/{id}/relationships/builds"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.appEncryptionDeclarationBuildsLinkagesRequest == nil {
		return nil, reportError("appEncryptionDeclarationBuildsLinkagesRequest is required and must be specified")
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
	// body params
	localVarPostBody = r.appEncryptionDeclarationBuildsLinkagesRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiAppEncryptionDeclarationsGetCollectionRequest struct {
	ctx _context.Context
	ApiService *AppEncryptionDeclarationsApiService
	filterPlatform *[]string
	filterApp *[]string
	filterBuilds *[]string
	fieldsAppEncryptionDeclarations *[]string
	limit *int32
	include *[]string
	fieldsApps *[]string
}

func (r ApiAppEncryptionDeclarationsGetCollectionRequest) FilterPlatform(filterPlatform []string) ApiAppEncryptionDeclarationsGetCollectionRequest {
	r.filterPlatform = &filterPlatform
	return r
}
func (r ApiAppEncryptionDeclarationsGetCollectionRequest) FilterApp(filterApp []string) ApiAppEncryptionDeclarationsGetCollectionRequest {
	r.filterApp = &filterApp
	return r
}
func (r ApiAppEncryptionDeclarationsGetCollectionRequest) FilterBuilds(filterBuilds []string) ApiAppEncryptionDeclarationsGetCollectionRequest {
	r.filterBuilds = &filterBuilds
	return r
}
func (r ApiAppEncryptionDeclarationsGetCollectionRequest) FieldsAppEncryptionDeclarations(fieldsAppEncryptionDeclarations []string) ApiAppEncryptionDeclarationsGetCollectionRequest {
	r.fieldsAppEncryptionDeclarations = &fieldsAppEncryptionDeclarations
	return r
}
func (r ApiAppEncryptionDeclarationsGetCollectionRequest) Limit(limit int32) ApiAppEncryptionDeclarationsGetCollectionRequest {
	r.limit = &limit
	return r
}
func (r ApiAppEncryptionDeclarationsGetCollectionRequest) Include(include []string) ApiAppEncryptionDeclarationsGetCollectionRequest {
	r.include = &include
	return r
}
func (r ApiAppEncryptionDeclarationsGetCollectionRequest) FieldsApps(fieldsApps []string) ApiAppEncryptionDeclarationsGetCollectionRequest {
	r.fieldsApps = &fieldsApps
	return r
}

func (r ApiAppEncryptionDeclarationsGetCollectionRequest) Execute() (AppEncryptionDeclarationsResponse, *_nethttp.Response, error) {
	return r.ApiService.AppEncryptionDeclarationsGetCollectionExecute(r)
}

/*
 * AppEncryptionDeclarationsGetCollection Method for AppEncryptionDeclarationsGetCollection
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiAppEncryptionDeclarationsGetCollectionRequest
 */
func (a *AppEncryptionDeclarationsApiService) AppEncryptionDeclarationsGetCollection(ctx _context.Context) ApiAppEncryptionDeclarationsGetCollectionRequest {
	return ApiAppEncryptionDeclarationsGetCollectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return AppEncryptionDeclarationsResponse
 */
func (a *AppEncryptionDeclarationsApiService) AppEncryptionDeclarationsGetCollectionExecute(r ApiAppEncryptionDeclarationsGetCollectionRequest) (AppEncryptionDeclarationsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AppEncryptionDeclarationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppEncryptionDeclarationsApiService.AppEncryptionDeclarationsGetCollection")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appEncryptionDeclarations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.filterPlatform != nil {
		localVarQueryParams.Add("filter[platform]", parameterToString(*r.filterPlatform, "csv"))
	}
	if r.filterApp != nil {
		localVarQueryParams.Add("filter[app]", parameterToString(*r.filterApp, "csv"))
	}
	if r.filterBuilds != nil {
		localVarQueryParams.Add("filter[builds]", parameterToString(*r.filterBuilds, "csv"))
	}
	if r.fieldsAppEncryptionDeclarations != nil {
		localVarQueryParams.Add("fields[appEncryptionDeclarations]", parameterToString(*r.fieldsAppEncryptionDeclarations, "csv"))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, "csv"))
	}
	if r.fieldsApps != nil {
		localVarQueryParams.Add("fields[apps]", parameterToString(*r.fieldsApps, "csv"))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiAppEncryptionDeclarationsGetInstanceRequest struct {
	ctx _context.Context
	ApiService *AppEncryptionDeclarationsApiService
	id string
	fieldsAppEncryptionDeclarations *[]string
	include *[]string
	fieldsApps *[]string
}

func (r ApiAppEncryptionDeclarationsGetInstanceRequest) FieldsAppEncryptionDeclarations(fieldsAppEncryptionDeclarations []string) ApiAppEncryptionDeclarationsGetInstanceRequest {
	r.fieldsAppEncryptionDeclarations = &fieldsAppEncryptionDeclarations
	return r
}
func (r ApiAppEncryptionDeclarationsGetInstanceRequest) Include(include []string) ApiAppEncryptionDeclarationsGetInstanceRequest {
	r.include = &include
	return r
}
func (r ApiAppEncryptionDeclarationsGetInstanceRequest) FieldsApps(fieldsApps []string) ApiAppEncryptionDeclarationsGetInstanceRequest {
	r.fieldsApps = &fieldsApps
	return r
}

func (r ApiAppEncryptionDeclarationsGetInstanceRequest) Execute() (AppEncryptionDeclarationResponse, *_nethttp.Response, error) {
	return r.ApiService.AppEncryptionDeclarationsGetInstanceExecute(r)
}

/*
 * AppEncryptionDeclarationsGetInstance Method for AppEncryptionDeclarationsGetInstance
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id the id of the requested resource
 * @return ApiAppEncryptionDeclarationsGetInstanceRequest
 */
func (a *AppEncryptionDeclarationsApiService) AppEncryptionDeclarationsGetInstance(ctx _context.Context, id string) ApiAppEncryptionDeclarationsGetInstanceRequest {
	return ApiAppEncryptionDeclarationsGetInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return AppEncryptionDeclarationResponse
 */
func (a *AppEncryptionDeclarationsApiService) AppEncryptionDeclarationsGetInstanceExecute(r ApiAppEncryptionDeclarationsGetInstanceRequest) (AppEncryptionDeclarationResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AppEncryptionDeclarationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppEncryptionDeclarationsApiService.AppEncryptionDeclarationsGetInstance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appEncryptionDeclarations/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.fieldsAppEncryptionDeclarations != nil {
		localVarQueryParams.Add("fields[appEncryptionDeclarations]", parameterToString(*r.fieldsAppEncryptionDeclarations, "csv"))
	}
	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, "csv"))
	}
	if r.fieldsApps != nil {
		localVarQueryParams.Add("fields[apps]", parameterToString(*r.fieldsApps, "csv"))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}