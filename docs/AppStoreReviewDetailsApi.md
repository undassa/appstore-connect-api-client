# \AppStoreReviewDetailsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelated**](AppStoreReviewDetailsApi.md#AppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelated) | **Get** /v1/appStoreReviewDetails/{id}/appStoreReviewAttachments | 
[**AppStoreReviewDetailsCreateInstance**](AppStoreReviewDetailsApi.md#AppStoreReviewDetailsCreateInstance) | **Post** /v1/appStoreReviewDetails | 
[**AppStoreReviewDetailsGetInstance**](AppStoreReviewDetailsApi.md#AppStoreReviewDetailsGetInstance) | **Get** /v1/appStoreReviewDetails/{id} | 
[**AppStoreReviewDetailsUpdateInstance**](AppStoreReviewDetailsApi.md#AppStoreReviewDetailsUpdateInstance) | **Patch** /v1/appStoreReviewDetails/{id} | 



## AppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelated

> AppStoreReviewAttachmentsResponse AppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelated(ctx, id).FieldsAppStoreReviewDetails(fieldsAppStoreReviewDetails).FieldsAppStoreReviewAttachments(fieldsAppStoreReviewAttachments).Limit(limit).Include(include).Execute()



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
    fieldsAppStoreReviewDetails := []string{"FieldsAppStoreReviewDetails_example"} // []string | the fields to include for returned resources of type appStoreReviewDetails (optional)
    fieldsAppStoreReviewAttachments := []string{"FieldsAppStoreReviewAttachments_example"} // []string | the fields to include for returned resources of type appStoreReviewAttachments (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppStoreReviewDetailsApi.AppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelated(context.Background(), id).FieldsAppStoreReviewDetails(fieldsAppStoreReviewDetails).FieldsAppStoreReviewAttachments(fieldsAppStoreReviewAttachments).Limit(limit).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppStoreReviewDetailsApi.AppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelated`: AppStoreReviewAttachmentsResponse
    fmt.Fprintf(os.Stdout, "Response from `AppStoreReviewDetailsApi.AppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppStoreReviewDetails** | **[]string** | the fields to include for returned resources of type appStoreReviewDetails | 
 **fieldsAppStoreReviewAttachments** | **[]string** | the fields to include for returned resources of type appStoreReviewAttachments | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**AppStoreReviewAttachmentsResponse**](AppStoreReviewAttachmentsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreReviewDetailsCreateInstance

> AppStoreReviewDetailResponse AppStoreReviewDetailsCreateInstance(ctx).AppStoreReviewDetailCreateRequest(appStoreReviewDetailCreateRequest).Execute()



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
    appStoreReviewDetailCreateRequest := *openapiclient.NewAppStoreReviewDetailCreateRequest(*openapiclient.NewAppStoreReviewDetailCreateRequestData("Type_example", *openapiclient.NewAppStoreReviewDetailCreateRequestDataRelationships(*openapiclient.NewAppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion(*openapiclient.NewAppStoreReviewDetailRelationshipsAppStoreVersionData("Type_example", "Id_example"))))) // AppStoreReviewDetailCreateRequest | AppStoreReviewDetail representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppStoreReviewDetailsApi.AppStoreReviewDetailsCreateInstance(context.Background()).AppStoreReviewDetailCreateRequest(appStoreReviewDetailCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppStoreReviewDetailsApi.AppStoreReviewDetailsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppStoreReviewDetailsCreateInstance`: AppStoreReviewDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `AppStoreReviewDetailsApi.AppStoreReviewDetailsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreReviewDetailsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appStoreReviewDetailCreateRequest** | [**AppStoreReviewDetailCreateRequest**](AppStoreReviewDetailCreateRequest.md) | AppStoreReviewDetail representation | 

### Return type

[**AppStoreReviewDetailResponse**](AppStoreReviewDetailResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreReviewDetailsGetInstance

> AppStoreReviewDetailResponse AppStoreReviewDetailsGetInstance(ctx, id).FieldsAppStoreReviewDetails(fieldsAppStoreReviewDetails).Include(include).FieldsAppStoreReviewAttachments(fieldsAppStoreReviewAttachments).LimitAppStoreReviewAttachments(limitAppStoreReviewAttachments).Execute()



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
    fieldsAppStoreReviewDetails := []string{"FieldsAppStoreReviewDetails_example"} // []string | the fields to include for returned resources of type appStoreReviewDetails (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsAppStoreReviewAttachments := []string{"FieldsAppStoreReviewAttachments_example"} // []string | the fields to include for returned resources of type appStoreReviewAttachments (optional)
    limitAppStoreReviewAttachments := int32(56) // int32 | maximum number of related appStoreReviewAttachments returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppStoreReviewDetailsApi.AppStoreReviewDetailsGetInstance(context.Background(), id).FieldsAppStoreReviewDetails(fieldsAppStoreReviewDetails).Include(include).FieldsAppStoreReviewAttachments(fieldsAppStoreReviewAttachments).LimitAppStoreReviewAttachments(limitAppStoreReviewAttachments).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppStoreReviewDetailsApi.AppStoreReviewDetailsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppStoreReviewDetailsGetInstance`: AppStoreReviewDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `AppStoreReviewDetailsApi.AppStoreReviewDetailsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreReviewDetailsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppStoreReviewDetails** | **[]string** | the fields to include for returned resources of type appStoreReviewDetails | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsAppStoreReviewAttachments** | **[]string** | the fields to include for returned resources of type appStoreReviewAttachments | 
 **limitAppStoreReviewAttachments** | **int32** | maximum number of related appStoreReviewAttachments returned (when they are included) | 

### Return type

[**AppStoreReviewDetailResponse**](AppStoreReviewDetailResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreReviewDetailsUpdateInstance

> AppStoreReviewDetailResponse AppStoreReviewDetailsUpdateInstance(ctx, id).AppStoreReviewDetailUpdateRequest(appStoreReviewDetailUpdateRequest).Execute()



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
    appStoreReviewDetailUpdateRequest := *openapiclient.NewAppStoreReviewDetailUpdateRequest(*openapiclient.NewAppStoreReviewDetailUpdateRequestData("Type_example", "Id_example")) // AppStoreReviewDetailUpdateRequest | AppStoreReviewDetail representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppStoreReviewDetailsApi.AppStoreReviewDetailsUpdateInstance(context.Background(), id).AppStoreReviewDetailUpdateRequest(appStoreReviewDetailUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppStoreReviewDetailsApi.AppStoreReviewDetailsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppStoreReviewDetailsUpdateInstance`: AppStoreReviewDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `AppStoreReviewDetailsApi.AppStoreReviewDetailsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreReviewDetailsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appStoreReviewDetailUpdateRequest** | [**AppStoreReviewDetailUpdateRequest**](AppStoreReviewDetailUpdateRequest.md) | AppStoreReviewDetail representation | 

### Return type

[**AppStoreReviewDetailResponse**](AppStoreReviewDetailResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

