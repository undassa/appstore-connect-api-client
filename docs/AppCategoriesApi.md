# \AppCategoriesApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppCategoriesGetCollection**](AppCategoriesApi.md#AppCategoriesGetCollection) | **Get** /v1/appCategories | 
[**AppCategoriesGetInstance**](AppCategoriesApi.md#AppCategoriesGetInstance) | **Get** /v1/appCategories/{id} | 
[**AppCategoriesParentGetToOneRelated**](AppCategoriesApi.md#AppCategoriesParentGetToOneRelated) | **Get** /v1/appCategories/{id}/parent | 
[**AppCategoriesSubcategoriesGetToManyRelated**](AppCategoriesApi.md#AppCategoriesSubcategoriesGetToManyRelated) | **Get** /v1/appCategories/{id}/subcategories | 



## AppCategoriesGetCollection

> AppCategoriesResponse AppCategoriesGetCollection(ctx).FilterPlatforms(filterPlatforms).ExistsParent(existsParent).FieldsAppCategories(fieldsAppCategories).Limit(limit).Include(include).LimitSubcategories(limitSubcategories).Execute()



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
    filterPlatforms := []string{"FilterPlatforms_example"} // []string | filter by attribute 'platforms' (optional)
    existsParent := []string{"Inner_example"} // []string | filter by existence or non-existence of related 'parent' (optional)
    fieldsAppCategories := []string{"FieldsAppCategories_example"} // []string | the fields to include for returned resources of type appCategories (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    limitSubcategories := int32(56) // int32 | maximum number of related subcategories returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppCategoriesApi.AppCategoriesGetCollection(context.Background()).FilterPlatforms(filterPlatforms).ExistsParent(existsParent).FieldsAppCategories(fieldsAppCategories).Limit(limit).Include(include).LimitSubcategories(limitSubcategories).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppCategoriesApi.AppCategoriesGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppCategoriesGetCollection`: AppCategoriesResponse
    fmt.Fprintf(os.Stdout, "Response from `AppCategoriesApi.AppCategoriesGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppCategoriesGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterPlatforms** | **[]string** | filter by attribute &#39;platforms&#39; | 
 **existsParent** | **[]string** | filter by existence or non-existence of related &#39;parent&#39; | 
 **fieldsAppCategories** | **[]string** | the fields to include for returned resources of type appCategories | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitSubcategories** | **int32** | maximum number of related subcategories returned (when they are included) | 

### Return type

[**AppCategoriesResponse**](AppCategoriesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCategoriesGetInstance

> AppCategoryResponse AppCategoriesGetInstance(ctx, id).FieldsAppCategories(fieldsAppCategories).Include(include).LimitSubcategories(limitSubcategories).Execute()



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
    fieldsAppCategories := []string{"FieldsAppCategories_example"} // []string | the fields to include for returned resources of type appCategories (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    limitSubcategories := int32(56) // int32 | maximum number of related subcategories returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppCategoriesApi.AppCategoriesGetInstance(context.Background(), id).FieldsAppCategories(fieldsAppCategories).Include(include).LimitSubcategories(limitSubcategories).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppCategoriesApi.AppCategoriesGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppCategoriesGetInstance`: AppCategoryResponse
    fmt.Fprintf(os.Stdout, "Response from `AppCategoriesApi.AppCategoriesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppCategoriesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppCategories** | **[]string** | the fields to include for returned resources of type appCategories | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitSubcategories** | **int32** | maximum number of related subcategories returned (when they are included) | 

### Return type

[**AppCategoryResponse**](AppCategoryResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCategoriesParentGetToOneRelated

> AppCategoryResponse AppCategoriesParentGetToOneRelated(ctx, id).FieldsAppCategories(fieldsAppCategories).Execute()



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
    fieldsAppCategories := []string{"FieldsAppCategories_example"} // []string | the fields to include for returned resources of type appCategories (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppCategoriesApi.AppCategoriesParentGetToOneRelated(context.Background(), id).FieldsAppCategories(fieldsAppCategories).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppCategoriesApi.AppCategoriesParentGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppCategoriesParentGetToOneRelated`: AppCategoryResponse
    fmt.Fprintf(os.Stdout, "Response from `AppCategoriesApi.AppCategoriesParentGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppCategoriesParentGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppCategories** | **[]string** | the fields to include for returned resources of type appCategories | 

### Return type

[**AppCategoryResponse**](AppCategoryResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCategoriesSubcategoriesGetToManyRelated

> AppCategoriesResponse AppCategoriesSubcategoriesGetToManyRelated(ctx, id).FieldsAppCategories(fieldsAppCategories).Limit(limit).Execute()



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
    fieldsAppCategories := []string{"FieldsAppCategories_example"} // []string | the fields to include for returned resources of type appCategories (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppCategoriesApi.AppCategoriesSubcategoriesGetToManyRelated(context.Background(), id).FieldsAppCategories(fieldsAppCategories).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppCategoriesApi.AppCategoriesSubcategoriesGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppCategoriesSubcategoriesGetToManyRelated`: AppCategoriesResponse
    fmt.Fprintf(os.Stdout, "Response from `AppCategoriesApi.AppCategoriesSubcategoriesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppCategoriesSubcategoriesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppCategories** | **[]string** | the fields to include for returned resources of type appCategories | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**AppCategoriesResponse**](AppCategoriesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

