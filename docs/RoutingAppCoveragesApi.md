# \RoutingAppCoveragesApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RoutingAppCoveragesCreateInstance**](RoutingAppCoveragesApi.md#RoutingAppCoveragesCreateInstance) | **Post** /v1/routingAppCoverages | 
[**RoutingAppCoveragesDeleteInstance**](RoutingAppCoveragesApi.md#RoutingAppCoveragesDeleteInstance) | **Delete** /v1/routingAppCoverages/{id} | 
[**RoutingAppCoveragesGetInstance**](RoutingAppCoveragesApi.md#RoutingAppCoveragesGetInstance) | **Get** /v1/routingAppCoverages/{id} | 
[**RoutingAppCoveragesUpdateInstance**](RoutingAppCoveragesApi.md#RoutingAppCoveragesUpdateInstance) | **Patch** /v1/routingAppCoverages/{id} | 



## RoutingAppCoveragesCreateInstance

> RoutingAppCoverageResponse RoutingAppCoveragesCreateInstance(ctx).RoutingAppCoverageCreateRequest(routingAppCoverageCreateRequest).Execute()



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
    routingAppCoverageCreateRequest := *openapiclient.NewRoutingAppCoverageCreateRequest(*openapiclient.NewRoutingAppCoverageCreateRequestData("Type_example", *openapiclient.NewAppScreenshotCreateRequestDataAttributes(int32(123), "FileName_example"), *openapiclient.NewAppStoreReviewDetailCreateRequestDataRelationships(*openapiclient.NewAppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion(*openapiclient.NewAppStoreReviewDetailRelationshipsAppStoreVersionData("Type_example", "Id_example"))))) // RoutingAppCoverageCreateRequest | RoutingAppCoverage representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoutingAppCoveragesApi.RoutingAppCoveragesCreateInstance(context.Background()).RoutingAppCoverageCreateRequest(routingAppCoverageCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingAppCoveragesApi.RoutingAppCoveragesCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RoutingAppCoveragesCreateInstance`: RoutingAppCoverageResponse
    fmt.Fprintf(os.Stdout, "Response from `RoutingAppCoveragesApi.RoutingAppCoveragesCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRoutingAppCoveragesCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routingAppCoverageCreateRequest** | [**RoutingAppCoverageCreateRequest**](RoutingAppCoverageCreateRequest.md) | RoutingAppCoverage representation | 

### Return type

[**RoutingAppCoverageResponse**](RoutingAppCoverageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoutingAppCoveragesDeleteInstance

> RoutingAppCoveragesDeleteInstance(ctx, id).Execute()



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
    resp, r, err := api_client.RoutingAppCoveragesApi.RoutingAppCoveragesDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingAppCoveragesApi.RoutingAppCoveragesDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRoutingAppCoveragesDeleteInstanceRequest struct via the builder pattern


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


## RoutingAppCoveragesGetInstance

> RoutingAppCoverageResponse RoutingAppCoveragesGetInstance(ctx, id).FieldsRoutingAppCoverages(fieldsRoutingAppCoverages).Include(include).Execute()



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
    fieldsRoutingAppCoverages := []string{"FieldsRoutingAppCoverages_example"} // []string | the fields to include for returned resources of type routingAppCoverages (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoutingAppCoveragesApi.RoutingAppCoveragesGetInstance(context.Background(), id).FieldsRoutingAppCoverages(fieldsRoutingAppCoverages).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingAppCoveragesApi.RoutingAppCoveragesGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RoutingAppCoveragesGetInstance`: RoutingAppCoverageResponse
    fmt.Fprintf(os.Stdout, "Response from `RoutingAppCoveragesApi.RoutingAppCoveragesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiRoutingAppCoveragesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsRoutingAppCoverages** | **[]string** | the fields to include for returned resources of type routingAppCoverages | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**RoutingAppCoverageResponse**](RoutingAppCoverageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoutingAppCoveragesUpdateInstance

> RoutingAppCoverageResponse RoutingAppCoveragesUpdateInstance(ctx, id).RoutingAppCoverageUpdateRequest(routingAppCoverageUpdateRequest).Execute()



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
    routingAppCoverageUpdateRequest := *openapiclient.NewRoutingAppCoverageUpdateRequest(*openapiclient.NewRoutingAppCoverageUpdateRequestData("Type_example", "Id_example")) // RoutingAppCoverageUpdateRequest | RoutingAppCoverage representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoutingAppCoveragesApi.RoutingAppCoveragesUpdateInstance(context.Background(), id).RoutingAppCoverageUpdateRequest(routingAppCoverageUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingAppCoveragesApi.RoutingAppCoveragesUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RoutingAppCoveragesUpdateInstance`: RoutingAppCoverageResponse
    fmt.Fprintf(os.Stdout, "Response from `RoutingAppCoveragesApi.RoutingAppCoveragesUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiRoutingAppCoveragesUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **routingAppCoverageUpdateRequest** | [**RoutingAppCoverageUpdateRequest**](RoutingAppCoverageUpdateRequest.md) | RoutingAppCoverage representation | 

### Return type

[**RoutingAppCoverageResponse**](RoutingAppCoverageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

