# \BetaAppLocalizationsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BetaAppLocalizationsAppGetToOneRelated**](BetaAppLocalizationsApi.md#BetaAppLocalizationsAppGetToOneRelated) | **Get** /v1/betaAppLocalizations/{id}/app | 
[**BetaAppLocalizationsCreateInstance**](BetaAppLocalizationsApi.md#BetaAppLocalizationsCreateInstance) | **Post** /v1/betaAppLocalizations | 
[**BetaAppLocalizationsDeleteInstance**](BetaAppLocalizationsApi.md#BetaAppLocalizationsDeleteInstance) | **Delete** /v1/betaAppLocalizations/{id} | 
[**BetaAppLocalizationsGetCollection**](BetaAppLocalizationsApi.md#BetaAppLocalizationsGetCollection) | **Get** /v1/betaAppLocalizations | 
[**BetaAppLocalizationsGetInstance**](BetaAppLocalizationsApi.md#BetaAppLocalizationsGetInstance) | **Get** /v1/betaAppLocalizations/{id} | 
[**BetaAppLocalizationsUpdateInstance**](BetaAppLocalizationsApi.md#BetaAppLocalizationsUpdateInstance) | **Patch** /v1/betaAppLocalizations/{id} | 



## BetaAppLocalizationsAppGetToOneRelated

> AppResponse BetaAppLocalizationsAppGetToOneRelated(ctx, id).FieldsApps(fieldsApps).Execute()



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
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaAppLocalizationsApi.BetaAppLocalizationsAppGetToOneRelated(context.Background(), id).FieldsApps(fieldsApps).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaAppLocalizationsApi.BetaAppLocalizationsAppGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaAppLocalizationsAppGetToOneRelated`: AppResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaAppLocalizationsApi.BetaAppLocalizationsAppGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaAppLocalizationsAppGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 

### Return type

[**AppResponse**](AppResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaAppLocalizationsCreateInstance

> BetaAppLocalizationResponse BetaAppLocalizationsCreateInstance(ctx).BetaAppLocalizationCreateRequest(betaAppLocalizationCreateRequest).Execute()



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
    betaAppLocalizationCreateRequest := *openapiclient.NewBetaAppLocalizationCreateRequest(*openapiclient.NewBetaAppLocalizationCreateRequestData("Type_example", *openapiclient.NewBetaAppLocalizationCreateRequestDataAttributes("Locale_example"), *openapiclient.NewAppPreOrderCreateRequestDataRelationships(*openapiclient.NewAppPreOrderCreateRequestDataRelationshipsApp(*openapiclient.NewAppEncryptionDeclarationRelationshipsAppData("Type_example", "Id_example"))))) // BetaAppLocalizationCreateRequest | BetaAppLocalization representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaAppLocalizationsApi.BetaAppLocalizationsCreateInstance(context.Background()).BetaAppLocalizationCreateRequest(betaAppLocalizationCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaAppLocalizationsApi.BetaAppLocalizationsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaAppLocalizationsCreateInstance`: BetaAppLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaAppLocalizationsApi.BetaAppLocalizationsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBetaAppLocalizationsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **betaAppLocalizationCreateRequest** | [**BetaAppLocalizationCreateRequest**](BetaAppLocalizationCreateRequest.md) | BetaAppLocalization representation | 

### Return type

[**BetaAppLocalizationResponse**](BetaAppLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaAppLocalizationsDeleteInstance

> BetaAppLocalizationsDeleteInstance(ctx, id).Execute()



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
    resp, r, err := api_client.BetaAppLocalizationsApi.BetaAppLocalizationsDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaAppLocalizationsApi.BetaAppLocalizationsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBetaAppLocalizationsDeleteInstanceRequest struct via the builder pattern


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


## BetaAppLocalizationsGetCollection

> BetaAppLocalizationsResponse BetaAppLocalizationsGetCollection(ctx).FilterLocale(filterLocale).FilterApp(filterApp).FieldsBetaAppLocalizations(fieldsBetaAppLocalizations).Limit(limit).Include(include).FieldsApps(fieldsApps).Execute()



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
    filterLocale := []string{"Inner_example"} // []string | filter by attribute 'locale' (optional)
    filterApp := []string{"Inner_example"} // []string | filter by id(s) of related 'app' (optional)
    fieldsBetaAppLocalizations := []string{"FieldsBetaAppLocalizations_example"} // []string | the fields to include for returned resources of type betaAppLocalizations (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaAppLocalizationsApi.BetaAppLocalizationsGetCollection(context.Background()).FilterLocale(filterLocale).FilterApp(filterApp).FieldsBetaAppLocalizations(fieldsBetaAppLocalizations).Limit(limit).Include(include).FieldsApps(fieldsApps).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaAppLocalizationsApi.BetaAppLocalizationsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaAppLocalizationsGetCollection`: BetaAppLocalizationsResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaAppLocalizationsApi.BetaAppLocalizationsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBetaAppLocalizationsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterLocale** | **[]string** | filter by attribute &#39;locale&#39; | 
 **filterApp** | **[]string** | filter by id(s) of related &#39;app&#39; | 
 **fieldsBetaAppLocalizations** | **[]string** | the fields to include for returned resources of type betaAppLocalizations | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 

### Return type

[**BetaAppLocalizationsResponse**](BetaAppLocalizationsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaAppLocalizationsGetInstance

> BetaAppLocalizationResponse BetaAppLocalizationsGetInstance(ctx, id).FieldsBetaAppLocalizations(fieldsBetaAppLocalizations).Include(include).FieldsApps(fieldsApps).Execute()



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
    fieldsBetaAppLocalizations := []string{"FieldsBetaAppLocalizations_example"} // []string | the fields to include for returned resources of type betaAppLocalizations (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaAppLocalizationsApi.BetaAppLocalizationsGetInstance(context.Background(), id).FieldsBetaAppLocalizations(fieldsBetaAppLocalizations).Include(include).FieldsApps(fieldsApps).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaAppLocalizationsApi.BetaAppLocalizationsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaAppLocalizationsGetInstance`: BetaAppLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaAppLocalizationsApi.BetaAppLocalizationsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaAppLocalizationsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaAppLocalizations** | **[]string** | the fields to include for returned resources of type betaAppLocalizations | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 

### Return type

[**BetaAppLocalizationResponse**](BetaAppLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaAppLocalizationsUpdateInstance

> BetaAppLocalizationResponse BetaAppLocalizationsUpdateInstance(ctx, id).BetaAppLocalizationUpdateRequest(betaAppLocalizationUpdateRequest).Execute()



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
    betaAppLocalizationUpdateRequest := *openapiclient.NewBetaAppLocalizationUpdateRequest(*openapiclient.NewBetaAppLocalizationUpdateRequestData("Type_example", "Id_example")) // BetaAppLocalizationUpdateRequest | BetaAppLocalization representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaAppLocalizationsApi.BetaAppLocalizationsUpdateInstance(context.Background(), id).BetaAppLocalizationUpdateRequest(betaAppLocalizationUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaAppLocalizationsApi.BetaAppLocalizationsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaAppLocalizationsUpdateInstance`: BetaAppLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaAppLocalizationsApi.BetaAppLocalizationsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaAppLocalizationsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaAppLocalizationUpdateRequest** | [**BetaAppLocalizationUpdateRequest**](BetaAppLocalizationUpdateRequest.md) | BetaAppLocalization representation | 

### Return type

[**BetaAppLocalizationResponse**](BetaAppLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

