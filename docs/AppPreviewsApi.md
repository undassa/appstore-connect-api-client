# \AppPreviewsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppPreviewsCreateInstance**](AppPreviewsApi.md#AppPreviewsCreateInstance) | **Post** /v1/appPreviews | 
[**AppPreviewsDeleteInstance**](AppPreviewsApi.md#AppPreviewsDeleteInstance) | **Delete** /v1/appPreviews/{id} | 
[**AppPreviewsGetInstance**](AppPreviewsApi.md#AppPreviewsGetInstance) | **Get** /v1/appPreviews/{id} | 
[**AppPreviewsUpdateInstance**](AppPreviewsApi.md#AppPreviewsUpdateInstance) | **Patch** /v1/appPreviews/{id} | 



## AppPreviewsCreateInstance

> AppPreviewResponse AppPreviewsCreateInstance(ctx).AppPreviewCreateRequest(appPreviewCreateRequest).Execute()



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
    appPreviewCreateRequest := *openapiclient.NewAppPreviewCreateRequest(*openapiclient.NewAppPreviewCreateRequestData("Type_example", *openapiclient.NewAppPreviewCreateRequestDataAttributes(int32(123), "FileName_example"), *openapiclient.NewAppPreviewCreateRequestDataRelationships(*openapiclient.NewAppPreviewCreateRequestDataRelationshipsAppPreviewSet(*openapiclient.NewAppPreviewRelationshipsAppPreviewSetData("Type_example", "Id_example"))))) // AppPreviewCreateRequest | AppPreview representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppPreviewsApi.AppPreviewsCreateInstance(context.Background()).AppPreviewCreateRequest(appPreviewCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppPreviewsApi.AppPreviewsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppPreviewsCreateInstance`: AppPreviewResponse
    fmt.Fprintf(os.Stdout, "Response from `AppPreviewsApi.AppPreviewsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppPreviewsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appPreviewCreateRequest** | [**AppPreviewCreateRequest**](AppPreviewCreateRequest.md) | AppPreview representation | 

### Return type

[**AppPreviewResponse**](AppPreviewResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPreviewsDeleteInstance

> AppPreviewsDeleteInstance(ctx, id).Execute()



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
    resp, r, err := api_client.AppPreviewsApi.AppPreviewsDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppPreviewsApi.AppPreviewsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppPreviewsDeleteInstanceRequest struct via the builder pattern


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


## AppPreviewsGetInstance

> AppPreviewResponse AppPreviewsGetInstance(ctx, id).FieldsAppPreviews(fieldsAppPreviews).Include(include).Execute()



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
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppPreviewsApi.AppPreviewsGetInstance(context.Background(), id).FieldsAppPreviews(fieldsAppPreviews).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppPreviewsApi.AppPreviewsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppPreviewsGetInstance`: AppPreviewResponse
    fmt.Fprintf(os.Stdout, "Response from `AppPreviewsApi.AppPreviewsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppPreviewsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppPreviews** | **[]string** | the fields to include for returned resources of type appPreviews | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**AppPreviewResponse**](AppPreviewResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPreviewsUpdateInstance

> AppPreviewResponse AppPreviewsUpdateInstance(ctx, id).AppPreviewUpdateRequest(appPreviewUpdateRequest).Execute()



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
    appPreviewUpdateRequest := *openapiclient.NewAppPreviewUpdateRequest(*openapiclient.NewAppPreviewUpdateRequestData("Type_example", "Id_example")) // AppPreviewUpdateRequest | AppPreview representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppPreviewsApi.AppPreviewsUpdateInstance(context.Background(), id).AppPreviewUpdateRequest(appPreviewUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppPreviewsApi.AppPreviewsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppPreviewsUpdateInstance`: AppPreviewResponse
    fmt.Fprintf(os.Stdout, "Response from `AppPreviewsApi.AppPreviewsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppPreviewsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appPreviewUpdateRequest** | [**AppPreviewUpdateRequest**](AppPreviewUpdateRequest.md) | AppPreview representation | 

### Return type

[**AppPreviewResponse**](AppPreviewResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

