# \AppStoreVersionLocalizationsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppStoreVersionLocalizationsAppPreviewSetsGetToManyRelated**](AppStoreVersionLocalizationsApi.md#AppStoreVersionLocalizationsAppPreviewSetsGetToManyRelated) | **Get** /v1/appStoreVersionLocalizations/{id}/appPreviewSets | 
[**AppStoreVersionLocalizationsAppScreenshotSetsGetToManyRelated**](AppStoreVersionLocalizationsApi.md#AppStoreVersionLocalizationsAppScreenshotSetsGetToManyRelated) | **Get** /v1/appStoreVersionLocalizations/{id}/appScreenshotSets | 
[**AppStoreVersionLocalizationsCreateInstance**](AppStoreVersionLocalizationsApi.md#AppStoreVersionLocalizationsCreateInstance) | **Post** /v1/appStoreVersionLocalizations | 
[**AppStoreVersionLocalizationsDeleteInstance**](AppStoreVersionLocalizationsApi.md#AppStoreVersionLocalizationsDeleteInstance) | **Delete** /v1/appStoreVersionLocalizations/{id} | 
[**AppStoreVersionLocalizationsGetInstance**](AppStoreVersionLocalizationsApi.md#AppStoreVersionLocalizationsGetInstance) | **Get** /v1/appStoreVersionLocalizations/{id} | 
[**AppStoreVersionLocalizationsUpdateInstance**](AppStoreVersionLocalizationsApi.md#AppStoreVersionLocalizationsUpdateInstance) | **Patch** /v1/appStoreVersionLocalizations/{id} | 



## AppStoreVersionLocalizationsAppPreviewSetsGetToManyRelated

> AppPreviewSetsResponse AppStoreVersionLocalizationsAppPreviewSetsGetToManyRelated(ctx, id).FilterPreviewType(filterPreviewType).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).FieldsAppPreviews(fieldsAppPreviews).FieldsAppPreviewSets(fieldsAppPreviewSets).Limit(limit).Include(include).Execute()



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
    filterPreviewType := []string{"FilterPreviewType_example"} // []string | filter by attribute 'previewType' (optional)
    fieldsAppStoreVersionLocalizations := []string{"FieldsAppStoreVersionLocalizations_example"} // []string | the fields to include for returned resources of type appStoreVersionLocalizations (optional)
    fieldsAppPreviews := []string{"FieldsAppPreviews_example"} // []string | the fields to include for returned resources of type appPreviews (optional)
    fieldsAppPreviewSets := []string{"FieldsAppPreviewSets_example"} // []string | the fields to include for returned resources of type appPreviewSets (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppStoreVersionLocalizationsApi.AppStoreVersionLocalizationsAppPreviewSetsGetToManyRelated(context.Background(), id).FilterPreviewType(filterPreviewType).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).FieldsAppPreviews(fieldsAppPreviews).FieldsAppPreviewSets(fieldsAppPreviewSets).Limit(limit).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionLocalizationsApi.AppStoreVersionLocalizationsAppPreviewSetsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppStoreVersionLocalizationsAppPreviewSetsGetToManyRelated`: AppPreviewSetsResponse
    fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionLocalizationsApi.AppStoreVersionLocalizationsAppPreviewSetsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionLocalizationsAppPreviewSetsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterPreviewType** | **[]string** | filter by attribute &#39;previewType&#39; | 
 **fieldsAppStoreVersionLocalizations** | **[]string** | the fields to include for returned resources of type appStoreVersionLocalizations | 
 **fieldsAppPreviews** | **[]string** | the fields to include for returned resources of type appPreviews | 
 **fieldsAppPreviewSets** | **[]string** | the fields to include for returned resources of type appPreviewSets | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**AppPreviewSetsResponse**](AppPreviewSetsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionLocalizationsAppScreenshotSetsGetToManyRelated

> AppScreenshotSetsResponse AppStoreVersionLocalizationsAppScreenshotSetsGetToManyRelated(ctx, id).FilterScreenshotDisplayType(filterScreenshotDisplayType).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).FieldsAppScreenshotSets(fieldsAppScreenshotSets).FieldsAppScreenshots(fieldsAppScreenshots).Limit(limit).Include(include).Execute()



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
    filterScreenshotDisplayType := []string{"FilterScreenshotDisplayType_example"} // []string | filter by attribute 'screenshotDisplayType' (optional)
    fieldsAppStoreVersionLocalizations := []string{"FieldsAppStoreVersionLocalizations_example"} // []string | the fields to include for returned resources of type appStoreVersionLocalizations (optional)
    fieldsAppScreenshotSets := []string{"FieldsAppScreenshotSets_example"} // []string | the fields to include for returned resources of type appScreenshotSets (optional)
    fieldsAppScreenshots := []string{"FieldsAppScreenshots_example"} // []string | the fields to include for returned resources of type appScreenshots (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppStoreVersionLocalizationsApi.AppStoreVersionLocalizationsAppScreenshotSetsGetToManyRelated(context.Background(), id).FilterScreenshotDisplayType(filterScreenshotDisplayType).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).FieldsAppScreenshotSets(fieldsAppScreenshotSets).FieldsAppScreenshots(fieldsAppScreenshots).Limit(limit).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionLocalizationsApi.AppStoreVersionLocalizationsAppScreenshotSetsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppStoreVersionLocalizationsAppScreenshotSetsGetToManyRelated`: AppScreenshotSetsResponse
    fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionLocalizationsApi.AppStoreVersionLocalizationsAppScreenshotSetsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionLocalizationsAppScreenshotSetsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterScreenshotDisplayType** | **[]string** | filter by attribute &#39;screenshotDisplayType&#39; | 
 **fieldsAppStoreVersionLocalizations** | **[]string** | the fields to include for returned resources of type appStoreVersionLocalizations | 
 **fieldsAppScreenshotSets** | **[]string** | the fields to include for returned resources of type appScreenshotSets | 
 **fieldsAppScreenshots** | **[]string** | the fields to include for returned resources of type appScreenshots | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**AppScreenshotSetsResponse**](AppScreenshotSetsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionLocalizationsCreateInstance

> AppStoreVersionLocalizationResponse AppStoreVersionLocalizationsCreateInstance(ctx).AppStoreVersionLocalizationCreateRequest(appStoreVersionLocalizationCreateRequest).Execute()



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
    appStoreVersionLocalizationCreateRequest := *openapiclient.NewAppStoreVersionLocalizationCreateRequest(*openapiclient.NewAppStoreVersionLocalizationCreateRequestData("Type_example", *openapiclient.NewAppStoreVersionLocalizationCreateRequestDataAttributes("Locale_example"), *openapiclient.NewAppStoreReviewDetailCreateRequestDataRelationships(*openapiclient.NewAppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion(*openapiclient.NewAppStoreReviewDetailRelationshipsAppStoreVersionData("Type_example", "Id_example"))))) // AppStoreVersionLocalizationCreateRequest | AppStoreVersionLocalization representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppStoreVersionLocalizationsApi.AppStoreVersionLocalizationsCreateInstance(context.Background()).AppStoreVersionLocalizationCreateRequest(appStoreVersionLocalizationCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionLocalizationsApi.AppStoreVersionLocalizationsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppStoreVersionLocalizationsCreateInstance`: AppStoreVersionLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionLocalizationsApi.AppStoreVersionLocalizationsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionLocalizationsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appStoreVersionLocalizationCreateRequest** | [**AppStoreVersionLocalizationCreateRequest**](AppStoreVersionLocalizationCreateRequest.md) | AppStoreVersionLocalization representation | 

### Return type

[**AppStoreVersionLocalizationResponse**](AppStoreVersionLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionLocalizationsDeleteInstance

> AppStoreVersionLocalizationsDeleteInstance(ctx, id).Execute()



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
    resp, r, err := api_client.AppStoreVersionLocalizationsApi.AppStoreVersionLocalizationsDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionLocalizationsApi.AppStoreVersionLocalizationsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppStoreVersionLocalizationsDeleteInstanceRequest struct via the builder pattern


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


## AppStoreVersionLocalizationsGetInstance

> AppStoreVersionLocalizationResponse AppStoreVersionLocalizationsGetInstance(ctx, id).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).Include(include).FieldsAppScreenshotSets(fieldsAppScreenshotSets).FieldsAppPreviewSets(fieldsAppPreviewSets).LimitAppPreviewSets(limitAppPreviewSets).LimitAppScreenshotSets(limitAppScreenshotSets).Execute()



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
    fieldsAppStoreVersionLocalizations := []string{"FieldsAppStoreVersionLocalizations_example"} // []string | the fields to include for returned resources of type appStoreVersionLocalizations (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsAppScreenshotSets := []string{"FieldsAppScreenshotSets_example"} // []string | the fields to include for returned resources of type appScreenshotSets (optional)
    fieldsAppPreviewSets := []string{"FieldsAppPreviewSets_example"} // []string | the fields to include for returned resources of type appPreviewSets (optional)
    limitAppPreviewSets := int32(56) // int32 | maximum number of related appPreviewSets returned (when they are included) (optional)
    limitAppScreenshotSets := int32(56) // int32 | maximum number of related appScreenshotSets returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppStoreVersionLocalizationsApi.AppStoreVersionLocalizationsGetInstance(context.Background(), id).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).Include(include).FieldsAppScreenshotSets(fieldsAppScreenshotSets).FieldsAppPreviewSets(fieldsAppPreviewSets).LimitAppPreviewSets(limitAppPreviewSets).LimitAppScreenshotSets(limitAppScreenshotSets).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionLocalizationsApi.AppStoreVersionLocalizationsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppStoreVersionLocalizationsGetInstance`: AppStoreVersionLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionLocalizationsApi.AppStoreVersionLocalizationsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionLocalizationsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppStoreVersionLocalizations** | **[]string** | the fields to include for returned resources of type appStoreVersionLocalizations | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsAppScreenshotSets** | **[]string** | the fields to include for returned resources of type appScreenshotSets | 
 **fieldsAppPreviewSets** | **[]string** | the fields to include for returned resources of type appPreviewSets | 
 **limitAppPreviewSets** | **int32** | maximum number of related appPreviewSets returned (when they are included) | 
 **limitAppScreenshotSets** | **int32** | maximum number of related appScreenshotSets returned (when they are included) | 

### Return type

[**AppStoreVersionLocalizationResponse**](AppStoreVersionLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionLocalizationsUpdateInstance

> AppStoreVersionLocalizationResponse AppStoreVersionLocalizationsUpdateInstance(ctx, id).AppStoreVersionLocalizationUpdateRequest(appStoreVersionLocalizationUpdateRequest).Execute()



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
    appStoreVersionLocalizationUpdateRequest := *openapiclient.NewAppStoreVersionLocalizationUpdateRequest(*openapiclient.NewAppStoreVersionLocalizationUpdateRequestData("Type_example", "Id_example")) // AppStoreVersionLocalizationUpdateRequest | AppStoreVersionLocalization representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppStoreVersionLocalizationsApi.AppStoreVersionLocalizationsUpdateInstance(context.Background(), id).AppStoreVersionLocalizationUpdateRequest(appStoreVersionLocalizationUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionLocalizationsApi.AppStoreVersionLocalizationsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppStoreVersionLocalizationsUpdateInstance`: AppStoreVersionLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionLocalizationsApi.AppStoreVersionLocalizationsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionLocalizationsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appStoreVersionLocalizationUpdateRequest** | [**AppStoreVersionLocalizationUpdateRequest**](AppStoreVersionLocalizationUpdateRequest.md) | AppStoreVersionLocalization representation | 

### Return type

[**AppStoreVersionLocalizationResponse**](AppStoreVersionLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

