# \AppStoreReviewAttachmentsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppStoreReviewAttachmentsCreateInstance**](AppStoreReviewAttachmentsApi.md#AppStoreReviewAttachmentsCreateInstance) | **Post** /v1/appStoreReviewAttachments | 
[**AppStoreReviewAttachmentsDeleteInstance**](AppStoreReviewAttachmentsApi.md#AppStoreReviewAttachmentsDeleteInstance) | **Delete** /v1/appStoreReviewAttachments/{id} | 
[**AppStoreReviewAttachmentsGetInstance**](AppStoreReviewAttachmentsApi.md#AppStoreReviewAttachmentsGetInstance) | **Get** /v1/appStoreReviewAttachments/{id} | 
[**AppStoreReviewAttachmentsUpdateInstance**](AppStoreReviewAttachmentsApi.md#AppStoreReviewAttachmentsUpdateInstance) | **Patch** /v1/appStoreReviewAttachments/{id} | 



## AppStoreReviewAttachmentsCreateInstance

> AppStoreReviewAttachmentResponse AppStoreReviewAttachmentsCreateInstance(ctx).AppStoreReviewAttachmentCreateRequest(appStoreReviewAttachmentCreateRequest).Execute()



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
    appStoreReviewAttachmentCreateRequest := *openapiclient.NewAppStoreReviewAttachmentCreateRequest(*openapiclient.NewAppStoreReviewAttachmentCreateRequestData("Type_example", *openapiclient.NewAppScreenshotCreateRequestDataAttributes(int32(123), "FileName_example"), *openapiclient.NewAppStoreReviewAttachmentCreateRequestDataRelationships(*openapiclient.NewAppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail(*openapiclient.NewAppStoreReviewAttachmentRelationshipsAppStoreReviewDetailData("Type_example", "Id_example"))))) // AppStoreReviewAttachmentCreateRequest | AppStoreReviewAttachment representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppStoreReviewAttachmentsApi.AppStoreReviewAttachmentsCreateInstance(context.Background()).AppStoreReviewAttachmentCreateRequest(appStoreReviewAttachmentCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppStoreReviewAttachmentsApi.AppStoreReviewAttachmentsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppStoreReviewAttachmentsCreateInstance`: AppStoreReviewAttachmentResponse
    fmt.Fprintf(os.Stdout, "Response from `AppStoreReviewAttachmentsApi.AppStoreReviewAttachmentsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreReviewAttachmentsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appStoreReviewAttachmentCreateRequest** | [**AppStoreReviewAttachmentCreateRequest**](AppStoreReviewAttachmentCreateRequest.md) | AppStoreReviewAttachment representation | 

### Return type

[**AppStoreReviewAttachmentResponse**](AppStoreReviewAttachmentResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreReviewAttachmentsDeleteInstance

> AppStoreReviewAttachmentsDeleteInstance(ctx, id).Execute()



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
    resp, r, err := api_client.AppStoreReviewAttachmentsApi.AppStoreReviewAttachmentsDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppStoreReviewAttachmentsApi.AppStoreReviewAttachmentsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppStoreReviewAttachmentsDeleteInstanceRequest struct via the builder pattern


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


## AppStoreReviewAttachmentsGetInstance

> AppStoreReviewAttachmentResponse AppStoreReviewAttachmentsGetInstance(ctx, id).FieldsAppStoreReviewAttachments(fieldsAppStoreReviewAttachments).Include(include).Execute()



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
    fieldsAppStoreReviewAttachments := []string{"FieldsAppStoreReviewAttachments_example"} // []string | the fields to include for returned resources of type appStoreReviewAttachments (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppStoreReviewAttachmentsApi.AppStoreReviewAttachmentsGetInstance(context.Background(), id).FieldsAppStoreReviewAttachments(fieldsAppStoreReviewAttachments).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppStoreReviewAttachmentsApi.AppStoreReviewAttachmentsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppStoreReviewAttachmentsGetInstance`: AppStoreReviewAttachmentResponse
    fmt.Fprintf(os.Stdout, "Response from `AppStoreReviewAttachmentsApi.AppStoreReviewAttachmentsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreReviewAttachmentsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppStoreReviewAttachments** | **[]string** | the fields to include for returned resources of type appStoreReviewAttachments | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**AppStoreReviewAttachmentResponse**](AppStoreReviewAttachmentResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreReviewAttachmentsUpdateInstance

> AppStoreReviewAttachmentResponse AppStoreReviewAttachmentsUpdateInstance(ctx, id).AppStoreReviewAttachmentUpdateRequest(appStoreReviewAttachmentUpdateRequest).Execute()



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
    appStoreReviewAttachmentUpdateRequest := *openapiclient.NewAppStoreReviewAttachmentUpdateRequest(*openapiclient.NewAppStoreReviewAttachmentUpdateRequestData("Type_example", "Id_example")) // AppStoreReviewAttachmentUpdateRequest | AppStoreReviewAttachment representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppStoreReviewAttachmentsApi.AppStoreReviewAttachmentsUpdateInstance(context.Background(), id).AppStoreReviewAttachmentUpdateRequest(appStoreReviewAttachmentUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppStoreReviewAttachmentsApi.AppStoreReviewAttachmentsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppStoreReviewAttachmentsUpdateInstance`: AppStoreReviewAttachmentResponse
    fmt.Fprintf(os.Stdout, "Response from `AppStoreReviewAttachmentsApi.AppStoreReviewAttachmentsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreReviewAttachmentsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appStoreReviewAttachmentUpdateRequest** | [**AppStoreReviewAttachmentUpdateRequest**](AppStoreReviewAttachmentUpdateRequest.md) | AppStoreReviewAttachment representation | 

### Return type

[**AppStoreReviewAttachmentResponse**](AppStoreReviewAttachmentResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

