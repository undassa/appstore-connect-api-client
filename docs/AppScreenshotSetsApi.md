# \AppScreenshotSetsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppScreenshotSetsAppScreenshotsGetToManyRelated**](AppScreenshotSetsApi.md#AppScreenshotSetsAppScreenshotsGetToManyRelated) | **Get** /v1/appScreenshotSets/{id}/appScreenshots | 
[**AppScreenshotSetsAppScreenshotsGetToManyRelationship**](AppScreenshotSetsApi.md#AppScreenshotSetsAppScreenshotsGetToManyRelationship) | **Get** /v1/appScreenshotSets/{id}/relationships/appScreenshots | 
[**AppScreenshotSetsAppScreenshotsReplaceToManyRelationship**](AppScreenshotSetsApi.md#AppScreenshotSetsAppScreenshotsReplaceToManyRelationship) | **Patch** /v1/appScreenshotSets/{id}/relationships/appScreenshots | 
[**AppScreenshotSetsCreateInstance**](AppScreenshotSetsApi.md#AppScreenshotSetsCreateInstance) | **Post** /v1/appScreenshotSets | 
[**AppScreenshotSetsDeleteInstance**](AppScreenshotSetsApi.md#AppScreenshotSetsDeleteInstance) | **Delete** /v1/appScreenshotSets/{id} | 
[**AppScreenshotSetsGetInstance**](AppScreenshotSetsApi.md#AppScreenshotSetsGetInstance) | **Get** /v1/appScreenshotSets/{id} | 



## AppScreenshotSetsAppScreenshotsGetToManyRelated

> AppScreenshotsResponse AppScreenshotSetsAppScreenshotsGetToManyRelated(ctx, id).FieldsAppScreenshotSets(fieldsAppScreenshotSets).FieldsAppScreenshots(fieldsAppScreenshots).Limit(limit).Include(include).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | the id of the requested resource
    fieldsAppScreenshotSets := []string{"FieldsAppScreenshotSets_example"} // []string | the fields to include for returned resources of type appScreenshotSets (optional)
    fieldsAppScreenshots := []string{"FieldsAppScreenshots_example"} // []string | the fields to include for returned resources of type appScreenshots (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppScreenshotSetsApi.AppScreenshotSetsAppScreenshotsGetToManyRelated(context.Background(), id).FieldsAppScreenshotSets(fieldsAppScreenshotSets).FieldsAppScreenshots(fieldsAppScreenshots).Limit(limit).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppScreenshotSetsApi.AppScreenshotSetsAppScreenshotsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppScreenshotSetsAppScreenshotsGetToManyRelated`: AppScreenshotsResponse
    fmt.Fprintf(os.Stdout, "Response from `AppScreenshotSetsApi.AppScreenshotSetsAppScreenshotsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppScreenshotSetsAppScreenshotsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppScreenshotSets** | **[]string** | the fields to include for returned resources of type appScreenshotSets | 
 **fieldsAppScreenshots** | **[]string** | the fields to include for returned resources of type appScreenshots | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**AppScreenshotsResponse**](AppScreenshotsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppScreenshotSetsAppScreenshotsGetToManyRelationship

> AppScreenshotSetAppScreenshotsLinkagesResponse AppScreenshotSetsAppScreenshotsGetToManyRelationship(ctx, id).Limit(limit).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | the id of the requested resource
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppScreenshotSetsApi.AppScreenshotSetsAppScreenshotsGetToManyRelationship(context.Background(), id).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppScreenshotSetsApi.AppScreenshotSetsAppScreenshotsGetToManyRelationship``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppScreenshotSetsAppScreenshotsGetToManyRelationship`: AppScreenshotSetAppScreenshotsLinkagesResponse
    fmt.Fprintf(os.Stdout, "Response from `AppScreenshotSetsApi.AppScreenshotSetsAppScreenshotsGetToManyRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppScreenshotSetsAppScreenshotsGetToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum resources per page | 

### Return type

[**AppScreenshotSetAppScreenshotsLinkagesResponse**](AppScreenshotSetAppScreenshotsLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppScreenshotSetsAppScreenshotsReplaceToManyRelationship

> AppScreenshotSetsAppScreenshotsReplaceToManyRelationship(ctx, id).AppScreenshotSetAppScreenshotsLinkagesRequest(appScreenshotSetAppScreenshotsLinkagesRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | the id of the requested resource
    appScreenshotSetAppScreenshotsLinkagesRequest := *openapiclient.NewAppScreenshotSetAppScreenshotsLinkagesRequest([]openapiclient.AppScreenshotSetRelationshipsAppScreenshotsData{*openapiclient.NewAppScreenshotSetRelationshipsAppScreenshotsData("Type_example", "Id_example")}) // AppScreenshotSetAppScreenshotsLinkagesRequest | List of related linkages

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppScreenshotSetsApi.AppScreenshotSetsAppScreenshotsReplaceToManyRelationship(context.Background(), id).AppScreenshotSetAppScreenshotsLinkagesRequest(appScreenshotSetAppScreenshotsLinkagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppScreenshotSetsApi.AppScreenshotSetsAppScreenshotsReplaceToManyRelationship``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppScreenshotSetsAppScreenshotsReplaceToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appScreenshotSetAppScreenshotsLinkagesRequest** | [**AppScreenshotSetAppScreenshotsLinkagesRequest**](AppScreenshotSetAppScreenshotsLinkagesRequest.md) | List of related linkages | 

### Return type

 (empty response body)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppScreenshotSetsCreateInstance

> AppScreenshotSetResponse AppScreenshotSetsCreateInstance(ctx).AppScreenshotSetCreateRequest(appScreenshotSetCreateRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appScreenshotSetCreateRequest := *openapiclient.NewAppScreenshotSetCreateRequest(*openapiclient.NewAppScreenshotSetCreateRequestData("Type_example", *openapiclient.NewAppScreenshotSetCreateRequestDataAttributes(openapiclient.ScreenshotDisplayType("APP_IPHONE_65")), *openapiclient.NewAppPreviewSetCreateRequestDataRelationships(*openapiclient.NewAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization(*openapiclient.NewAppPreviewSetRelationshipsAppStoreVersionLocalizationData("Type_example", "Id_example"))))) // AppScreenshotSetCreateRequest | AppScreenshotSet representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppScreenshotSetsApi.AppScreenshotSetsCreateInstance(context.Background()).AppScreenshotSetCreateRequest(appScreenshotSetCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppScreenshotSetsApi.AppScreenshotSetsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppScreenshotSetsCreateInstance`: AppScreenshotSetResponse
    fmt.Fprintf(os.Stdout, "Response from `AppScreenshotSetsApi.AppScreenshotSetsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppScreenshotSetsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appScreenshotSetCreateRequest** | [**AppScreenshotSetCreateRequest**](AppScreenshotSetCreateRequest.md) | AppScreenshotSet representation | 

### Return type

[**AppScreenshotSetResponse**](AppScreenshotSetResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppScreenshotSetsDeleteInstance

> AppScreenshotSetsDeleteInstance(ctx, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | the id of the requested resource

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppScreenshotSetsApi.AppScreenshotSetsDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppScreenshotSetsApi.AppScreenshotSetsDeleteInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppScreenshotSetsDeleteInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppScreenshotSetsGetInstance

> AppScreenshotSetResponse AppScreenshotSetsGetInstance(ctx, id).FieldsAppScreenshotSets(fieldsAppScreenshotSets).Include(include).FieldsAppScreenshots(fieldsAppScreenshots).LimitAppScreenshots(limitAppScreenshots).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | the id of the requested resource
    fieldsAppScreenshotSets := []string{"FieldsAppScreenshotSets_example"} // []string | the fields to include for returned resources of type appScreenshotSets (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsAppScreenshots := []string{"FieldsAppScreenshots_example"} // []string | the fields to include for returned resources of type appScreenshots (optional)
    limitAppScreenshots := int32(56) // int32 | maximum number of related appScreenshots returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppScreenshotSetsApi.AppScreenshotSetsGetInstance(context.Background(), id).FieldsAppScreenshotSets(fieldsAppScreenshotSets).Include(include).FieldsAppScreenshots(fieldsAppScreenshots).LimitAppScreenshots(limitAppScreenshots).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppScreenshotSetsApi.AppScreenshotSetsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppScreenshotSetsGetInstance`: AppScreenshotSetResponse
    fmt.Fprintf(os.Stdout, "Response from `AppScreenshotSetsApi.AppScreenshotSetsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppScreenshotSetsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppScreenshotSets** | **[]string** | the fields to include for returned resources of type appScreenshotSets | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsAppScreenshots** | **[]string** | the fields to include for returned resources of type appScreenshots | 
 **limitAppScreenshots** | **int32** | maximum number of related appScreenshots returned (when they are included) | 

### Return type

[**AppScreenshotSetResponse**](AppScreenshotSetResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

