# \BetaBuildLocalizationsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BetaBuildLocalizationsBuildGetToOneRelated**](BetaBuildLocalizationsApi.md#BetaBuildLocalizationsBuildGetToOneRelated) | **Get** /v1/betaBuildLocalizations/{id}/build | 
[**BetaBuildLocalizationsCreateInstance**](BetaBuildLocalizationsApi.md#BetaBuildLocalizationsCreateInstance) | **Post** /v1/betaBuildLocalizations | 
[**BetaBuildLocalizationsDeleteInstance**](BetaBuildLocalizationsApi.md#BetaBuildLocalizationsDeleteInstance) | **Delete** /v1/betaBuildLocalizations/{id} | 
[**BetaBuildLocalizationsGetCollection**](BetaBuildLocalizationsApi.md#BetaBuildLocalizationsGetCollection) | **Get** /v1/betaBuildLocalizations | 
[**BetaBuildLocalizationsGetInstance**](BetaBuildLocalizationsApi.md#BetaBuildLocalizationsGetInstance) | **Get** /v1/betaBuildLocalizations/{id} | 
[**BetaBuildLocalizationsUpdateInstance**](BetaBuildLocalizationsApi.md#BetaBuildLocalizationsUpdateInstance) | **Patch** /v1/betaBuildLocalizations/{id} | 



## BetaBuildLocalizationsBuildGetToOneRelated

> BuildResponse BetaBuildLocalizationsBuildGetToOneRelated(ctx, id).FieldsBuilds(fieldsBuilds).Execute()



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
    fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaBuildLocalizationsApi.BetaBuildLocalizationsBuildGetToOneRelated(context.Background(), id).FieldsBuilds(fieldsBuilds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaBuildLocalizationsApi.BetaBuildLocalizationsBuildGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaBuildLocalizationsBuildGetToOneRelated`: BuildResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaBuildLocalizationsApi.BetaBuildLocalizationsBuildGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaBuildLocalizationsBuildGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 

### Return type

[**BuildResponse**](BuildResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaBuildLocalizationsCreateInstance

> BetaBuildLocalizationResponse BetaBuildLocalizationsCreateInstance(ctx).BetaBuildLocalizationCreateRequest(betaBuildLocalizationCreateRequest).Execute()



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
    betaBuildLocalizationCreateRequest := *openapiclient.NewBetaBuildLocalizationCreateRequest(*openapiclient.NewBetaBuildLocalizationCreateRequestData("Type_example", *openapiclient.NewBetaBuildLocalizationCreateRequestDataAttributes("Locale_example"), *openapiclient.NewBetaAppReviewSubmissionCreateRequestDataRelationships(*openapiclient.NewBetaAppReviewSubmissionCreateRequestDataRelationshipsBuild(*openapiclient.NewAppStoreVersionRelationshipsBuildData("Type_example", "Id_example"))))) // BetaBuildLocalizationCreateRequest | BetaBuildLocalization representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaBuildLocalizationsApi.BetaBuildLocalizationsCreateInstance(context.Background()).BetaBuildLocalizationCreateRequest(betaBuildLocalizationCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaBuildLocalizationsApi.BetaBuildLocalizationsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaBuildLocalizationsCreateInstance`: BetaBuildLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaBuildLocalizationsApi.BetaBuildLocalizationsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBetaBuildLocalizationsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **betaBuildLocalizationCreateRequest** | [**BetaBuildLocalizationCreateRequest**](BetaBuildLocalizationCreateRequest.md) | BetaBuildLocalization representation | 

### Return type

[**BetaBuildLocalizationResponse**](BetaBuildLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaBuildLocalizationsDeleteInstance

> BetaBuildLocalizationsDeleteInstance(ctx, id).Execute()



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
    resp, r, err := api_client.BetaBuildLocalizationsApi.BetaBuildLocalizationsDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaBuildLocalizationsApi.BetaBuildLocalizationsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBetaBuildLocalizationsDeleteInstanceRequest struct via the builder pattern


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


## BetaBuildLocalizationsGetCollection

> BetaBuildLocalizationsResponse BetaBuildLocalizationsGetCollection(ctx).FilterLocale(filterLocale).FilterBuild(filterBuild).FieldsBetaBuildLocalizations(fieldsBetaBuildLocalizations).Limit(limit).Include(include).FieldsBuilds(fieldsBuilds).Execute()



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
    filterBuild := []string{"Inner_example"} // []string | filter by id(s) of related 'build' (optional)
    fieldsBetaBuildLocalizations := []string{"FieldsBetaBuildLocalizations_example"} // []string | the fields to include for returned resources of type betaBuildLocalizations (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaBuildLocalizationsApi.BetaBuildLocalizationsGetCollection(context.Background()).FilterLocale(filterLocale).FilterBuild(filterBuild).FieldsBetaBuildLocalizations(fieldsBetaBuildLocalizations).Limit(limit).Include(include).FieldsBuilds(fieldsBuilds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaBuildLocalizationsApi.BetaBuildLocalizationsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaBuildLocalizationsGetCollection`: BetaBuildLocalizationsResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaBuildLocalizationsApi.BetaBuildLocalizationsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBetaBuildLocalizationsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterLocale** | **[]string** | filter by attribute &#39;locale&#39; | 
 **filterBuild** | **[]string** | filter by id(s) of related &#39;build&#39; | 
 **fieldsBetaBuildLocalizations** | **[]string** | the fields to include for returned resources of type betaBuildLocalizations | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 

### Return type

[**BetaBuildLocalizationsResponse**](BetaBuildLocalizationsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaBuildLocalizationsGetInstance

> BetaBuildLocalizationResponse BetaBuildLocalizationsGetInstance(ctx, id).FieldsBetaBuildLocalizations(fieldsBetaBuildLocalizations).Include(include).FieldsBuilds(fieldsBuilds).Execute()



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
    fieldsBetaBuildLocalizations := []string{"FieldsBetaBuildLocalizations_example"} // []string | the fields to include for returned resources of type betaBuildLocalizations (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaBuildLocalizationsApi.BetaBuildLocalizationsGetInstance(context.Background(), id).FieldsBetaBuildLocalizations(fieldsBetaBuildLocalizations).Include(include).FieldsBuilds(fieldsBuilds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaBuildLocalizationsApi.BetaBuildLocalizationsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaBuildLocalizationsGetInstance`: BetaBuildLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaBuildLocalizationsApi.BetaBuildLocalizationsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaBuildLocalizationsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaBuildLocalizations** | **[]string** | the fields to include for returned resources of type betaBuildLocalizations | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 

### Return type

[**BetaBuildLocalizationResponse**](BetaBuildLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaBuildLocalizationsUpdateInstance

> BetaBuildLocalizationResponse BetaBuildLocalizationsUpdateInstance(ctx, id).BetaBuildLocalizationUpdateRequest(betaBuildLocalizationUpdateRequest).Execute()



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
    betaBuildLocalizationUpdateRequest := *openapiclient.NewBetaBuildLocalizationUpdateRequest(*openapiclient.NewBetaBuildLocalizationUpdateRequestData("Type_example", "Id_example")) // BetaBuildLocalizationUpdateRequest | BetaBuildLocalization representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaBuildLocalizationsApi.BetaBuildLocalizationsUpdateInstance(context.Background(), id).BetaBuildLocalizationUpdateRequest(betaBuildLocalizationUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaBuildLocalizationsApi.BetaBuildLocalizationsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaBuildLocalizationsUpdateInstance`: BetaBuildLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaBuildLocalizationsApi.BetaBuildLocalizationsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaBuildLocalizationsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaBuildLocalizationUpdateRequest** | [**BetaBuildLocalizationUpdateRequest**](BetaBuildLocalizationUpdateRequest.md) | BetaBuildLocalization representation | 

### Return type

[**BetaBuildLocalizationResponse**](BetaBuildLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

