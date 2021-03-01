# \IdfaDeclarationsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdfaDeclarationsCreateInstance**](IdfaDeclarationsApi.md#IdfaDeclarationsCreateInstance) | **Post** /v1/idfaDeclarations | 
[**IdfaDeclarationsDeleteInstance**](IdfaDeclarationsApi.md#IdfaDeclarationsDeleteInstance) | **Delete** /v1/idfaDeclarations/{id} | 
[**IdfaDeclarationsUpdateInstance**](IdfaDeclarationsApi.md#IdfaDeclarationsUpdateInstance) | **Patch** /v1/idfaDeclarations/{id} | 



## IdfaDeclarationsCreateInstance

> IdfaDeclarationResponse IdfaDeclarationsCreateInstance(ctx).IdfaDeclarationCreateRequest(idfaDeclarationCreateRequest).Execute()



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
    idfaDeclarationCreateRequest := *openapiclient.NewIdfaDeclarationCreateRequest(*openapiclient.NewIdfaDeclarationCreateRequestData("Type_example", *openapiclient.NewIdfaDeclarationCreateRequestDataAttributes(false, false, false, false), *openapiclient.NewAppStoreReviewDetailCreateRequestDataRelationships(*openapiclient.NewAppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion(*openapiclient.NewAppStoreReviewDetailRelationshipsAppStoreVersionData("Type_example", "Id_example"))))) // IdfaDeclarationCreateRequest | IdfaDeclaration representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdfaDeclarationsApi.IdfaDeclarationsCreateInstance(context.Background()).IdfaDeclarationCreateRequest(idfaDeclarationCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdfaDeclarationsApi.IdfaDeclarationsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdfaDeclarationsCreateInstance`: IdfaDeclarationResponse
    fmt.Fprintf(os.Stdout, "Response from `IdfaDeclarationsApi.IdfaDeclarationsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdfaDeclarationsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idfaDeclarationCreateRequest** | [**IdfaDeclarationCreateRequest**](IdfaDeclarationCreateRequest.md) | IdfaDeclaration representation | 

### Return type

[**IdfaDeclarationResponse**](IdfaDeclarationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdfaDeclarationsDeleteInstance

> IdfaDeclarationsDeleteInstance(ctx, id).Execute()



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
    resp, r, err := api_client.IdfaDeclarationsApi.IdfaDeclarationsDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdfaDeclarationsApi.IdfaDeclarationsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiIdfaDeclarationsDeleteInstanceRequest struct via the builder pattern


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


## IdfaDeclarationsUpdateInstance

> IdfaDeclarationResponse IdfaDeclarationsUpdateInstance(ctx, id).IdfaDeclarationUpdateRequest(idfaDeclarationUpdateRequest).Execute()



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
    idfaDeclarationUpdateRequest := *openapiclient.NewIdfaDeclarationUpdateRequest(*openapiclient.NewIdfaDeclarationUpdateRequestData("Type_example", "Id_example")) // IdfaDeclarationUpdateRequest | IdfaDeclaration representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdfaDeclarationsApi.IdfaDeclarationsUpdateInstance(context.Background(), id).IdfaDeclarationUpdateRequest(idfaDeclarationUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdfaDeclarationsApi.IdfaDeclarationsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdfaDeclarationsUpdateInstance`: IdfaDeclarationResponse
    fmt.Fprintf(os.Stdout, "Response from `IdfaDeclarationsApi.IdfaDeclarationsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdfaDeclarationsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idfaDeclarationUpdateRequest** | [**IdfaDeclarationUpdateRequest**](IdfaDeclarationUpdateRequest.md) | IdfaDeclaration representation | 

### Return type

[**IdfaDeclarationResponse**](IdfaDeclarationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

