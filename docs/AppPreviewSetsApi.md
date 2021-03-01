# \AppPreviewSetsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppPreviewSetsAppPreviewsGetToManyRelated**](AppPreviewSetsApi.md#AppPreviewSetsAppPreviewsGetToManyRelated) | **Get** /v1/appPreviewSets/{id}/appPreviews | 
[**AppPreviewSetsAppPreviewsGetToManyRelationship**](AppPreviewSetsApi.md#AppPreviewSetsAppPreviewsGetToManyRelationship) | **Get** /v1/appPreviewSets/{id}/relationships/appPreviews | 
[**AppPreviewSetsAppPreviewsReplaceToManyRelationship**](AppPreviewSetsApi.md#AppPreviewSetsAppPreviewsReplaceToManyRelationship) | **Patch** /v1/appPreviewSets/{id}/relationships/appPreviews | 
[**AppPreviewSetsCreateInstance**](AppPreviewSetsApi.md#AppPreviewSetsCreateInstance) | **Post** /v1/appPreviewSets | 
[**AppPreviewSetsDeleteInstance**](AppPreviewSetsApi.md#AppPreviewSetsDeleteInstance) | **Delete** /v1/appPreviewSets/{id} | 
[**AppPreviewSetsGetInstance**](AppPreviewSetsApi.md#AppPreviewSetsGetInstance) | **Get** /v1/appPreviewSets/{id} | 



## AppPreviewSetsAppPreviewsGetToManyRelated

> AppPreviewsResponse AppPreviewSetsAppPreviewsGetToManyRelated(ctx, id).FieldsAppPreviews(fieldsAppPreviews).FieldsAppPreviewSets(fieldsAppPreviewSets).Limit(limit).Include(include).Execute()



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
    fieldsAppPreviews := []string{"FieldsAppPreviews_example"} // []string | the fields to include for returned resources of type appPreviews (optional)
    fieldsAppPreviewSets := []string{"FieldsAppPreviewSets_example"} // []string | the fields to include for returned resources of type appPreviewSets (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppPreviewSetsApi.AppPreviewSetsAppPreviewsGetToManyRelated(context.Background(), id).FieldsAppPreviews(fieldsAppPreviews).FieldsAppPreviewSets(fieldsAppPreviewSets).Limit(limit).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppPreviewSetsApi.AppPreviewSetsAppPreviewsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppPreviewSetsAppPreviewsGetToManyRelated`: AppPreviewsResponse
    fmt.Fprintf(os.Stdout, "Response from `AppPreviewSetsApi.AppPreviewSetsAppPreviewsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppPreviewSetsAppPreviewsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppPreviews** | **[]string** | the fields to include for returned resources of type appPreviews | 
 **fieldsAppPreviewSets** | **[]string** | the fields to include for returned resources of type appPreviewSets | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**AppPreviewsResponse**](AppPreviewsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPreviewSetsAppPreviewsGetToManyRelationship

> AppPreviewSetAppPreviewsLinkagesResponse AppPreviewSetsAppPreviewsGetToManyRelationship(ctx, id).Limit(limit).Execute()



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
    resp, r, err := api_client.AppPreviewSetsApi.AppPreviewSetsAppPreviewsGetToManyRelationship(context.Background(), id).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppPreviewSetsApi.AppPreviewSetsAppPreviewsGetToManyRelationship``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppPreviewSetsAppPreviewsGetToManyRelationship`: AppPreviewSetAppPreviewsLinkagesResponse
    fmt.Fprintf(os.Stdout, "Response from `AppPreviewSetsApi.AppPreviewSetsAppPreviewsGetToManyRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppPreviewSetsAppPreviewsGetToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum resources per page | 

### Return type

[**AppPreviewSetAppPreviewsLinkagesResponse**](AppPreviewSetAppPreviewsLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPreviewSetsAppPreviewsReplaceToManyRelationship

> AppPreviewSetsAppPreviewsReplaceToManyRelationship(ctx, id).AppPreviewSetAppPreviewsLinkagesRequest(appPreviewSetAppPreviewsLinkagesRequest).Execute()



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
    appPreviewSetAppPreviewsLinkagesRequest := *openapiclient.NewAppPreviewSetAppPreviewsLinkagesRequest([]openapiclient.AppPreviewSetRelationshipsAppPreviewsData{*openapiclient.NewAppPreviewSetRelationshipsAppPreviewsData("Type_example", "Id_example")}) // AppPreviewSetAppPreviewsLinkagesRequest | List of related linkages

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppPreviewSetsApi.AppPreviewSetsAppPreviewsReplaceToManyRelationship(context.Background(), id).AppPreviewSetAppPreviewsLinkagesRequest(appPreviewSetAppPreviewsLinkagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppPreviewSetsApi.AppPreviewSetsAppPreviewsReplaceToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppPreviewSetsAppPreviewsReplaceToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appPreviewSetAppPreviewsLinkagesRequest** | [**AppPreviewSetAppPreviewsLinkagesRequest**](AppPreviewSetAppPreviewsLinkagesRequest.md) | List of related linkages | 

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


## AppPreviewSetsCreateInstance

> AppPreviewSetResponse AppPreviewSetsCreateInstance(ctx).AppPreviewSetCreateRequest(appPreviewSetCreateRequest).Execute()



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
    appPreviewSetCreateRequest := *openapiclient.NewAppPreviewSetCreateRequest(*openapiclient.NewAppPreviewSetCreateRequestData("Type_example", *openapiclient.NewAppPreviewSetCreateRequestDataAttributes(openapiclient.PreviewType("IPHONE_65")), *openapiclient.NewAppPreviewSetCreateRequestDataRelationships(*openapiclient.NewAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization(*openapiclient.NewAppPreviewSetRelationshipsAppStoreVersionLocalizationData("Type_example", "Id_example"))))) // AppPreviewSetCreateRequest | AppPreviewSet representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppPreviewSetsApi.AppPreviewSetsCreateInstance(context.Background()).AppPreviewSetCreateRequest(appPreviewSetCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppPreviewSetsApi.AppPreviewSetsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppPreviewSetsCreateInstance`: AppPreviewSetResponse
    fmt.Fprintf(os.Stdout, "Response from `AppPreviewSetsApi.AppPreviewSetsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppPreviewSetsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appPreviewSetCreateRequest** | [**AppPreviewSetCreateRequest**](AppPreviewSetCreateRequest.md) | AppPreviewSet representation | 

### Return type

[**AppPreviewSetResponse**](AppPreviewSetResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPreviewSetsDeleteInstance

> AppPreviewSetsDeleteInstance(ctx, id).Execute()



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
    resp, r, err := api_client.AppPreviewSetsApi.AppPreviewSetsDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppPreviewSetsApi.AppPreviewSetsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppPreviewSetsDeleteInstanceRequest struct via the builder pattern


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


## AppPreviewSetsGetInstance

> AppPreviewSetResponse AppPreviewSetsGetInstance(ctx, id).FieldsAppPreviewSets(fieldsAppPreviewSets).Include(include).FieldsAppPreviews(fieldsAppPreviews).LimitAppPreviews(limitAppPreviews).Execute()



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
    fieldsAppPreviewSets := []string{"FieldsAppPreviewSets_example"} // []string | the fields to include for returned resources of type appPreviewSets (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsAppPreviews := []string{"FieldsAppPreviews_example"} // []string | the fields to include for returned resources of type appPreviews (optional)
    limitAppPreviews := int32(56) // int32 | maximum number of related appPreviews returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppPreviewSetsApi.AppPreviewSetsGetInstance(context.Background(), id).FieldsAppPreviewSets(fieldsAppPreviewSets).Include(include).FieldsAppPreviews(fieldsAppPreviews).LimitAppPreviews(limitAppPreviews).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppPreviewSetsApi.AppPreviewSetsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppPreviewSetsGetInstance`: AppPreviewSetResponse
    fmt.Fprintf(os.Stdout, "Response from `AppPreviewSetsApi.AppPreviewSetsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppPreviewSetsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppPreviewSets** | **[]string** | the fields to include for returned resources of type appPreviewSets | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsAppPreviews** | **[]string** | the fields to include for returned resources of type appPreviews | 
 **limitAppPreviews** | **int32** | maximum number of related appPreviews returned (when they are included) | 

### Return type

[**AppPreviewSetResponse**](AppPreviewSetResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

