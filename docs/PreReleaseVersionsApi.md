# \PreReleaseVersionsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PreReleaseVersionsAppGetToOneRelated**](PreReleaseVersionsApi.md#PreReleaseVersionsAppGetToOneRelated) | **Get** /v1/preReleaseVersions/{id}/app | 
[**PreReleaseVersionsBuildsGetToManyRelated**](PreReleaseVersionsApi.md#PreReleaseVersionsBuildsGetToManyRelated) | **Get** /v1/preReleaseVersions/{id}/builds | 
[**PreReleaseVersionsGetCollection**](PreReleaseVersionsApi.md#PreReleaseVersionsGetCollection) | **Get** /v1/preReleaseVersions | 
[**PreReleaseVersionsGetInstance**](PreReleaseVersionsApi.md#PreReleaseVersionsGetInstance) | **Get** /v1/preReleaseVersions/{id} | 



## PreReleaseVersionsAppGetToOneRelated

> AppResponse PreReleaseVersionsAppGetToOneRelated(ctx, id).FieldsApps(fieldsApps).Execute()



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
    resp, r, err := api_client.PreReleaseVersionsApi.PreReleaseVersionsAppGetToOneRelated(context.Background(), id).FieldsApps(fieldsApps).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreReleaseVersionsApi.PreReleaseVersionsAppGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreReleaseVersionsAppGetToOneRelated`: AppResponse
    fmt.Fprintf(os.Stdout, "Response from `PreReleaseVersionsApi.PreReleaseVersionsAppGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreReleaseVersionsAppGetToOneRelatedRequest struct via the builder pattern


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


## PreReleaseVersionsBuildsGetToManyRelated

> BuildsResponse PreReleaseVersionsBuildsGetToManyRelated(ctx, id).FieldsBuilds(fieldsBuilds).Limit(limit).Execute()



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
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PreReleaseVersionsApi.PreReleaseVersionsBuildsGetToManyRelated(context.Background(), id).FieldsBuilds(fieldsBuilds).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreReleaseVersionsApi.PreReleaseVersionsBuildsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreReleaseVersionsBuildsGetToManyRelated`: BuildsResponse
    fmt.Fprintf(os.Stdout, "Response from `PreReleaseVersionsApi.PreReleaseVersionsBuildsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreReleaseVersionsBuildsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**BuildsResponse**](BuildsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreReleaseVersionsGetCollection

> PreReleaseVersionsResponse PreReleaseVersionsGetCollection(ctx).FilterBuildsExpired(filterBuildsExpired).FilterBuildsProcessingState(filterBuildsProcessingState).FilterPlatform(filterPlatform).FilterVersion(filterVersion).FilterApp(filterApp).FilterBuilds(filterBuilds).Sort(sort).FieldsPreReleaseVersions(fieldsPreReleaseVersions).Limit(limit).Include(include).FieldsBuilds(fieldsBuilds).FieldsApps(fieldsApps).LimitBuilds(limitBuilds).Execute()



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
    filterBuildsExpired := []string{"Inner_example"} // []string | filter by attribute 'builds.expired' (optional)
    filterBuildsProcessingState := []string{"FilterBuildsProcessingState_example"} // []string | filter by attribute 'builds.processingState' (optional)
    filterPlatform := []string{"FilterPlatform_example"} // []string | filter by attribute 'platform' (optional)
    filterVersion := []string{"Inner_example"} // []string | filter by attribute 'version' (optional)
    filterApp := []string{"Inner_example"} // []string | filter by id(s) of related 'app' (optional)
    filterBuilds := []string{"Inner_example"} // []string | filter by id(s) of related 'builds' (optional)
    sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; resources will be sorted as specified (optional)
    fieldsPreReleaseVersions := []string{"FieldsPreReleaseVersions_example"} // []string | the fields to include for returned resources of type preReleaseVersions (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
    limitBuilds := int32(56) // int32 | maximum number of related builds returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PreReleaseVersionsApi.PreReleaseVersionsGetCollection(context.Background()).FilterBuildsExpired(filterBuildsExpired).FilterBuildsProcessingState(filterBuildsProcessingState).FilterPlatform(filterPlatform).FilterVersion(filterVersion).FilterApp(filterApp).FilterBuilds(filterBuilds).Sort(sort).FieldsPreReleaseVersions(fieldsPreReleaseVersions).Limit(limit).Include(include).FieldsBuilds(fieldsBuilds).FieldsApps(fieldsApps).LimitBuilds(limitBuilds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreReleaseVersionsApi.PreReleaseVersionsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreReleaseVersionsGetCollection`: PreReleaseVersionsResponse
    fmt.Fprintf(os.Stdout, "Response from `PreReleaseVersionsApi.PreReleaseVersionsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreReleaseVersionsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterBuildsExpired** | **[]string** | filter by attribute &#39;builds.expired&#39; | 
 **filterBuildsProcessingState** | **[]string** | filter by attribute &#39;builds.processingState&#39; | 
 **filterPlatform** | **[]string** | filter by attribute &#39;platform&#39; | 
 **filterVersion** | **[]string** | filter by attribute &#39;version&#39; | 
 **filterApp** | **[]string** | filter by id(s) of related &#39;app&#39; | 
 **filterBuilds** | **[]string** | filter by id(s) of related &#39;builds&#39; | 
 **sort** | **[]string** | comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsPreReleaseVersions** | **[]string** | the fields to include for returned resources of type preReleaseVersions | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **limitBuilds** | **int32** | maximum number of related builds returned (when they are included) | 

### Return type

[**PreReleaseVersionsResponse**](PreReleaseVersionsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreReleaseVersionsGetInstance

> PrereleaseVersionResponse PreReleaseVersionsGetInstance(ctx, id).FieldsPreReleaseVersions(fieldsPreReleaseVersions).Include(include).FieldsBuilds(fieldsBuilds).FieldsApps(fieldsApps).LimitBuilds(limitBuilds).Execute()



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
    fieldsPreReleaseVersions := []string{"FieldsPreReleaseVersions_example"} // []string | the fields to include for returned resources of type preReleaseVersions (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
    limitBuilds := int32(56) // int32 | maximum number of related builds returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PreReleaseVersionsApi.PreReleaseVersionsGetInstance(context.Background(), id).FieldsPreReleaseVersions(fieldsPreReleaseVersions).Include(include).FieldsBuilds(fieldsBuilds).FieldsApps(fieldsApps).LimitBuilds(limitBuilds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreReleaseVersionsApi.PreReleaseVersionsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreReleaseVersionsGetInstance`: PrereleaseVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `PreReleaseVersionsApi.PreReleaseVersionsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreReleaseVersionsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsPreReleaseVersions** | **[]string** | the fields to include for returned resources of type preReleaseVersions | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **limitBuilds** | **int32** | maximum number of related builds returned (when they are included) | 

### Return type

[**PrereleaseVersionResponse**](PrereleaseVersionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

