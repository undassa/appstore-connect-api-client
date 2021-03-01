# \BundleIdsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BundleIdsAppGetToOneRelated**](BundleIdsApi.md#BundleIdsAppGetToOneRelated) | **Get** /v1/bundleIds/{id}/app | 
[**BundleIdsBundleIdCapabilitiesGetToManyRelated**](BundleIdsApi.md#BundleIdsBundleIdCapabilitiesGetToManyRelated) | **Get** /v1/bundleIds/{id}/bundleIdCapabilities | 
[**BundleIdsCreateInstance**](BundleIdsApi.md#BundleIdsCreateInstance) | **Post** /v1/bundleIds | 
[**BundleIdsDeleteInstance**](BundleIdsApi.md#BundleIdsDeleteInstance) | **Delete** /v1/bundleIds/{id} | 
[**BundleIdsGetCollection**](BundleIdsApi.md#BundleIdsGetCollection) | **Get** /v1/bundleIds | 
[**BundleIdsGetInstance**](BundleIdsApi.md#BundleIdsGetInstance) | **Get** /v1/bundleIds/{id} | 
[**BundleIdsProfilesGetToManyRelated**](BundleIdsApi.md#BundleIdsProfilesGetToManyRelated) | **Get** /v1/bundleIds/{id}/profiles | 
[**BundleIdsUpdateInstance**](BundleIdsApi.md#BundleIdsUpdateInstance) | **Patch** /v1/bundleIds/{id} | 



## BundleIdsAppGetToOneRelated

> AppResponse BundleIdsAppGetToOneRelated(ctx, id).FieldsApps(fieldsApps).Execute()



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
    resp, r, err := api_client.BundleIdsApi.BundleIdsAppGetToOneRelated(context.Background(), id).FieldsApps(fieldsApps).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundleIdsApi.BundleIdsAppGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BundleIdsAppGetToOneRelated`: AppResponse
    fmt.Fprintf(os.Stdout, "Response from `BundleIdsApi.BundleIdsAppGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBundleIdsAppGetToOneRelatedRequest struct via the builder pattern


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


## BundleIdsBundleIdCapabilitiesGetToManyRelated

> BundleIdCapabilitiesResponse BundleIdsBundleIdCapabilitiesGetToManyRelated(ctx, id).FieldsBundleIdCapabilities(fieldsBundleIdCapabilities).Limit(limit).Execute()



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
    fieldsBundleIdCapabilities := []string{"FieldsBundleIdCapabilities_example"} // []string | the fields to include for returned resources of type bundleIdCapabilities (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BundleIdsApi.BundleIdsBundleIdCapabilitiesGetToManyRelated(context.Background(), id).FieldsBundleIdCapabilities(fieldsBundleIdCapabilities).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundleIdsApi.BundleIdsBundleIdCapabilitiesGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BundleIdsBundleIdCapabilitiesGetToManyRelated`: BundleIdCapabilitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `BundleIdsApi.BundleIdsBundleIdCapabilitiesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBundleIdsBundleIdCapabilitiesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBundleIdCapabilities** | **[]string** | the fields to include for returned resources of type bundleIdCapabilities | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**BundleIdCapabilitiesResponse**](BundleIdCapabilitiesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BundleIdsCreateInstance

> BundleIdResponse BundleIdsCreateInstance(ctx).BundleIdCreateRequest(bundleIdCreateRequest).Execute()



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
    bundleIdCreateRequest := *openapiclient.NewBundleIdCreateRequest(*openapiclient.NewBundleIdCreateRequestData("Type_example", *openapiclient.NewBundleIdCreateRequestDataAttributes("Name_example", openapiclient.BundleIdPlatform("IOS"), "Identifier_example"))) // BundleIdCreateRequest | BundleId representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BundleIdsApi.BundleIdsCreateInstance(context.Background()).BundleIdCreateRequest(bundleIdCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundleIdsApi.BundleIdsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BundleIdsCreateInstance`: BundleIdResponse
    fmt.Fprintf(os.Stdout, "Response from `BundleIdsApi.BundleIdsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBundleIdsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bundleIdCreateRequest** | [**BundleIdCreateRequest**](BundleIdCreateRequest.md) | BundleId representation | 

### Return type

[**BundleIdResponse**](BundleIdResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BundleIdsDeleteInstance

> BundleIdsDeleteInstance(ctx, id).Execute()



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
    resp, r, err := api_client.BundleIdsApi.BundleIdsDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundleIdsApi.BundleIdsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBundleIdsDeleteInstanceRequest struct via the builder pattern


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


## BundleIdsGetCollection

> BundleIdsResponse BundleIdsGetCollection(ctx).FilterIdentifier(filterIdentifier).FilterName(filterName).FilterPlatform(filterPlatform).FilterSeedId(filterSeedId).FilterId(filterId).Sort(sort).FieldsBundleIds(fieldsBundleIds).Limit(limit).Include(include).FieldsBundleIdCapabilities(fieldsBundleIdCapabilities).FieldsProfiles(fieldsProfiles).FieldsApps(fieldsApps).LimitBundleIdCapabilities(limitBundleIdCapabilities).LimitProfiles(limitProfiles).Execute()



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
    filterIdentifier := []string{"Inner_example"} // []string | filter by attribute 'identifier' (optional)
    filterName := []string{"Inner_example"} // []string | filter by attribute 'name' (optional)
    filterPlatform := []string{"FilterPlatform_example"} // []string | filter by attribute 'platform' (optional)
    filterSeedId := []string{"Inner_example"} // []string | filter by attribute 'seedId' (optional)
    filterId := []string{"Inner_example"} // []string | filter by id(s) (optional)
    sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; resources will be sorted as specified (optional)
    fieldsBundleIds := []string{"FieldsBundleIds_example"} // []string | the fields to include for returned resources of type bundleIds (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsBundleIdCapabilities := []string{"FieldsBundleIdCapabilities_example"} // []string | the fields to include for returned resources of type bundleIdCapabilities (optional)
    fieldsProfiles := []string{"FieldsProfiles_example"} // []string | the fields to include for returned resources of type profiles (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
    limitBundleIdCapabilities := int32(56) // int32 | maximum number of related bundleIdCapabilities returned (when they are included) (optional)
    limitProfiles := int32(56) // int32 | maximum number of related profiles returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BundleIdsApi.BundleIdsGetCollection(context.Background()).FilterIdentifier(filterIdentifier).FilterName(filterName).FilterPlatform(filterPlatform).FilterSeedId(filterSeedId).FilterId(filterId).Sort(sort).FieldsBundleIds(fieldsBundleIds).Limit(limit).Include(include).FieldsBundleIdCapabilities(fieldsBundleIdCapabilities).FieldsProfiles(fieldsProfiles).FieldsApps(fieldsApps).LimitBundleIdCapabilities(limitBundleIdCapabilities).LimitProfiles(limitProfiles).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundleIdsApi.BundleIdsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BundleIdsGetCollection`: BundleIdsResponse
    fmt.Fprintf(os.Stdout, "Response from `BundleIdsApi.BundleIdsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBundleIdsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterIdentifier** | **[]string** | filter by attribute &#39;identifier&#39; | 
 **filterName** | **[]string** | filter by attribute &#39;name&#39; | 
 **filterPlatform** | **[]string** | filter by attribute &#39;platform&#39; | 
 **filterSeedId** | **[]string** | filter by attribute &#39;seedId&#39; | 
 **filterId** | **[]string** | filter by id(s) | 
 **sort** | **[]string** | comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsBundleIds** | **[]string** | the fields to include for returned resources of type bundleIds | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsBundleIdCapabilities** | **[]string** | the fields to include for returned resources of type bundleIdCapabilities | 
 **fieldsProfiles** | **[]string** | the fields to include for returned resources of type profiles | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **limitBundleIdCapabilities** | **int32** | maximum number of related bundleIdCapabilities returned (when they are included) | 
 **limitProfiles** | **int32** | maximum number of related profiles returned (when they are included) | 

### Return type

[**BundleIdsResponse**](BundleIdsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BundleIdsGetInstance

> BundleIdResponse BundleIdsGetInstance(ctx, id).FieldsBundleIds(fieldsBundleIds).Include(include).FieldsBundleIdCapabilities(fieldsBundleIdCapabilities).FieldsProfiles(fieldsProfiles).FieldsApps(fieldsApps).LimitBundleIdCapabilities(limitBundleIdCapabilities).LimitProfiles(limitProfiles).Execute()



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
    fieldsBundleIds := []string{"FieldsBundleIds_example"} // []string | the fields to include for returned resources of type bundleIds (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsBundleIdCapabilities := []string{"FieldsBundleIdCapabilities_example"} // []string | the fields to include for returned resources of type bundleIdCapabilities (optional)
    fieldsProfiles := []string{"FieldsProfiles_example"} // []string | the fields to include for returned resources of type profiles (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
    limitBundleIdCapabilities := int32(56) // int32 | maximum number of related bundleIdCapabilities returned (when they are included) (optional)
    limitProfiles := int32(56) // int32 | maximum number of related profiles returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BundleIdsApi.BundleIdsGetInstance(context.Background(), id).FieldsBundleIds(fieldsBundleIds).Include(include).FieldsBundleIdCapabilities(fieldsBundleIdCapabilities).FieldsProfiles(fieldsProfiles).FieldsApps(fieldsApps).LimitBundleIdCapabilities(limitBundleIdCapabilities).LimitProfiles(limitProfiles).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundleIdsApi.BundleIdsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BundleIdsGetInstance`: BundleIdResponse
    fmt.Fprintf(os.Stdout, "Response from `BundleIdsApi.BundleIdsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBundleIdsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBundleIds** | **[]string** | the fields to include for returned resources of type bundleIds | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsBundleIdCapabilities** | **[]string** | the fields to include for returned resources of type bundleIdCapabilities | 
 **fieldsProfiles** | **[]string** | the fields to include for returned resources of type profiles | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **limitBundleIdCapabilities** | **int32** | maximum number of related bundleIdCapabilities returned (when they are included) | 
 **limitProfiles** | **int32** | maximum number of related profiles returned (when they are included) | 

### Return type

[**BundleIdResponse**](BundleIdResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BundleIdsProfilesGetToManyRelated

> ProfilesResponse BundleIdsProfilesGetToManyRelated(ctx, id).FieldsProfiles(fieldsProfiles).Limit(limit).Execute()



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
    fieldsProfiles := []string{"FieldsProfiles_example"} // []string | the fields to include for returned resources of type profiles (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BundleIdsApi.BundleIdsProfilesGetToManyRelated(context.Background(), id).FieldsProfiles(fieldsProfiles).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundleIdsApi.BundleIdsProfilesGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BundleIdsProfilesGetToManyRelated`: ProfilesResponse
    fmt.Fprintf(os.Stdout, "Response from `BundleIdsApi.BundleIdsProfilesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBundleIdsProfilesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsProfiles** | **[]string** | the fields to include for returned resources of type profiles | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**ProfilesResponse**](ProfilesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BundleIdsUpdateInstance

> BundleIdResponse BundleIdsUpdateInstance(ctx, id).BundleIdUpdateRequest(bundleIdUpdateRequest).Execute()



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
    bundleIdUpdateRequest := *openapiclient.NewBundleIdUpdateRequest(*openapiclient.NewBundleIdUpdateRequestData("Type_example", "Id_example")) // BundleIdUpdateRequest | BundleId representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BundleIdsApi.BundleIdsUpdateInstance(context.Background(), id).BundleIdUpdateRequest(bundleIdUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundleIdsApi.BundleIdsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BundleIdsUpdateInstance`: BundleIdResponse
    fmt.Fprintf(os.Stdout, "Response from `BundleIdsApi.BundleIdsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBundleIdsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bundleIdUpdateRequest** | [**BundleIdUpdateRequest**](BundleIdUpdateRequest.md) | BundleId representation | 

### Return type

[**BundleIdResponse**](BundleIdResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

