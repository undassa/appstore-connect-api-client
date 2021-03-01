# \BundleIdCapabilitiesApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BundleIdCapabilitiesCreateInstance**](BundleIdCapabilitiesApi.md#BundleIdCapabilitiesCreateInstance) | **Post** /v1/bundleIdCapabilities | 
[**BundleIdCapabilitiesDeleteInstance**](BundleIdCapabilitiesApi.md#BundleIdCapabilitiesDeleteInstance) | **Delete** /v1/bundleIdCapabilities/{id} | 
[**BundleIdCapabilitiesUpdateInstance**](BundleIdCapabilitiesApi.md#BundleIdCapabilitiesUpdateInstance) | **Patch** /v1/bundleIdCapabilities/{id} | 



## BundleIdCapabilitiesCreateInstance

> BundleIdCapabilityResponse BundleIdCapabilitiesCreateInstance(ctx).BundleIdCapabilityCreateRequest(bundleIdCapabilityCreateRequest).Execute()



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
    bundleIdCapabilityCreateRequest := *openapiclient.NewBundleIdCapabilityCreateRequest(*openapiclient.NewBundleIdCapabilityCreateRequestData("Type_example", *openapiclient.NewBundleIdCapabilityCreateRequestDataAttributes(openapiclient.CapabilityType("ICLOUD")), *openapiclient.NewBundleIdCapabilityCreateRequestDataRelationships(*openapiclient.NewBundleIdCapabilityCreateRequestDataRelationshipsBundleId(*openapiclient.NewBundleIdCapabilityCreateRequestDataRelationshipsBundleIdData("Type_example", "Id_example"))))) // BundleIdCapabilityCreateRequest | BundleIdCapability representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BundleIdCapabilitiesApi.BundleIdCapabilitiesCreateInstance(context.Background()).BundleIdCapabilityCreateRequest(bundleIdCapabilityCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundleIdCapabilitiesApi.BundleIdCapabilitiesCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BundleIdCapabilitiesCreateInstance`: BundleIdCapabilityResponse
    fmt.Fprintf(os.Stdout, "Response from `BundleIdCapabilitiesApi.BundleIdCapabilitiesCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBundleIdCapabilitiesCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bundleIdCapabilityCreateRequest** | [**BundleIdCapabilityCreateRequest**](BundleIdCapabilityCreateRequest.md) | BundleIdCapability representation | 

### Return type

[**BundleIdCapabilityResponse**](BundleIdCapabilityResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BundleIdCapabilitiesDeleteInstance

> BundleIdCapabilitiesDeleteInstance(ctx, id).Execute()



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
    resp, r, err := api_client.BundleIdCapabilitiesApi.BundleIdCapabilitiesDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundleIdCapabilitiesApi.BundleIdCapabilitiesDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBundleIdCapabilitiesDeleteInstanceRequest struct via the builder pattern


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


## BundleIdCapabilitiesUpdateInstance

> BundleIdCapabilityResponse BundleIdCapabilitiesUpdateInstance(ctx, id).BundleIdCapabilityUpdateRequest(bundleIdCapabilityUpdateRequest).Execute()



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
    bundleIdCapabilityUpdateRequest := *openapiclient.NewBundleIdCapabilityUpdateRequest(*openapiclient.NewBundleIdCapabilityUpdateRequestData("Type_example", "Id_example")) // BundleIdCapabilityUpdateRequest | BundleIdCapability representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BundleIdCapabilitiesApi.BundleIdCapabilitiesUpdateInstance(context.Background(), id).BundleIdCapabilityUpdateRequest(bundleIdCapabilityUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundleIdCapabilitiesApi.BundleIdCapabilitiesUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BundleIdCapabilitiesUpdateInstance`: BundleIdCapabilityResponse
    fmt.Fprintf(os.Stdout, "Response from `BundleIdCapabilitiesApi.BundleIdCapabilitiesUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBundleIdCapabilitiesUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bundleIdCapabilityUpdateRequest** | [**BundleIdCapabilityUpdateRequest**](BundleIdCapabilityUpdateRequest.md) | BundleIdCapability representation | 

### Return type

[**BundleIdCapabilityResponse**](BundleIdCapabilityResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

