# \CertificatesApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CertificatesCreateInstance**](CertificatesApi.md#CertificatesCreateInstance) | **Post** /v1/certificates | 
[**CertificatesDeleteInstance**](CertificatesApi.md#CertificatesDeleteInstance) | **Delete** /v1/certificates/{id} | 
[**CertificatesGetCollection**](CertificatesApi.md#CertificatesGetCollection) | **Get** /v1/certificates | 
[**CertificatesGetInstance**](CertificatesApi.md#CertificatesGetInstance) | **Get** /v1/certificates/{id} | 



## CertificatesCreateInstance

> CertificateResponse CertificatesCreateInstance(ctx).CertificateCreateRequest(certificateCreateRequest).Execute()



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
    certificateCreateRequest := *openapiclient.NewCertificateCreateRequest(*openapiclient.NewCertificateCreateRequestData("Type_example", *openapiclient.NewCertificateCreateRequestDataAttributes("CsrContent_example", openapiclient.CertificateType("IOS_DEVELOPMENT")))) // CertificateCreateRequest | Certificate representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificatesApi.CertificatesCreateInstance(context.Background()).CertificateCreateRequest(certificateCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.CertificatesCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificatesCreateInstance`: CertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificatesApi.CertificatesCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificateCreateRequest** | [**CertificateCreateRequest**](CertificateCreateRequest.md) | Certificate representation | 

### Return type

[**CertificateResponse**](CertificateResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificatesDeleteInstance

> CertificatesDeleteInstance(ctx, id).Execute()



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
    resp, r, err := api_client.CertificatesApi.CertificatesDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.CertificatesDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCertificatesDeleteInstanceRequest struct via the builder pattern


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


## CertificatesGetCollection

> CertificatesResponse CertificatesGetCollection(ctx).FilterCertificateType(filterCertificateType).FilterDisplayName(filterDisplayName).FilterSerialNumber(filterSerialNumber).FilterId(filterId).Sort(sort).FieldsCertificates(fieldsCertificates).Limit(limit).Execute()



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
    filterCertificateType := []string{"FilterCertificateType_example"} // []string | filter by attribute 'certificateType' (optional)
    filterDisplayName := []string{"Inner_example"} // []string | filter by attribute 'displayName' (optional)
    filterSerialNumber := []string{"Inner_example"} // []string | filter by attribute 'serialNumber' (optional)
    filterId := []string{"Inner_example"} // []string | filter by id(s) (optional)
    sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; resources will be sorted as specified (optional)
    fieldsCertificates := []string{"FieldsCertificates_example"} // []string | the fields to include for returned resources of type certificates (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificatesApi.CertificatesGetCollection(context.Background()).FilterCertificateType(filterCertificateType).FilterDisplayName(filterDisplayName).FilterSerialNumber(filterSerialNumber).FilterId(filterId).Sort(sort).FieldsCertificates(fieldsCertificates).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.CertificatesGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificatesGetCollection`: CertificatesResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificatesApi.CertificatesGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterCertificateType** | **[]string** | filter by attribute &#39;certificateType&#39; | 
 **filterDisplayName** | **[]string** | filter by attribute &#39;displayName&#39; | 
 **filterSerialNumber** | **[]string** | filter by attribute &#39;serialNumber&#39; | 
 **filterId** | **[]string** | filter by id(s) | 
 **sort** | **[]string** | comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsCertificates** | **[]string** | the fields to include for returned resources of type certificates | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**CertificatesResponse**](CertificatesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificatesGetInstance

> CertificateResponse CertificatesGetInstance(ctx, id).FieldsCertificates(fieldsCertificates).Execute()



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
    fieldsCertificates := []string{"FieldsCertificates_example"} // []string | the fields to include for returned resources of type certificates (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificatesApi.CertificatesGetInstance(context.Background(), id).FieldsCertificates(fieldsCertificates).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.CertificatesGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificatesGetInstance`: CertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificatesApi.CertificatesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsCertificates** | **[]string** | the fields to include for returned resources of type certificates | 

### Return type

[**CertificateResponse**](CertificateResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

