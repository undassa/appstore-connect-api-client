# \BetaGroupsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BetaGroupsAppGetToOneRelated**](BetaGroupsApi.md#BetaGroupsAppGetToOneRelated) | **Get** /v1/betaGroups/{id}/app | 
[**BetaGroupsBetaTestersCreateToManyRelationship**](BetaGroupsApi.md#BetaGroupsBetaTestersCreateToManyRelationship) | **Post** /v1/betaGroups/{id}/relationships/betaTesters | 
[**BetaGroupsBetaTestersDeleteToManyRelationship**](BetaGroupsApi.md#BetaGroupsBetaTestersDeleteToManyRelationship) | **Delete** /v1/betaGroups/{id}/relationships/betaTesters | 
[**BetaGroupsBetaTestersGetToManyRelated**](BetaGroupsApi.md#BetaGroupsBetaTestersGetToManyRelated) | **Get** /v1/betaGroups/{id}/betaTesters | 
[**BetaGroupsBetaTestersGetToManyRelationship**](BetaGroupsApi.md#BetaGroupsBetaTestersGetToManyRelationship) | **Get** /v1/betaGroups/{id}/relationships/betaTesters | 
[**BetaGroupsBuildsCreateToManyRelationship**](BetaGroupsApi.md#BetaGroupsBuildsCreateToManyRelationship) | **Post** /v1/betaGroups/{id}/relationships/builds | 
[**BetaGroupsBuildsDeleteToManyRelationship**](BetaGroupsApi.md#BetaGroupsBuildsDeleteToManyRelationship) | **Delete** /v1/betaGroups/{id}/relationships/builds | 
[**BetaGroupsBuildsGetToManyRelated**](BetaGroupsApi.md#BetaGroupsBuildsGetToManyRelated) | **Get** /v1/betaGroups/{id}/builds | 
[**BetaGroupsBuildsGetToManyRelationship**](BetaGroupsApi.md#BetaGroupsBuildsGetToManyRelationship) | **Get** /v1/betaGroups/{id}/relationships/builds | 
[**BetaGroupsCreateInstance**](BetaGroupsApi.md#BetaGroupsCreateInstance) | **Post** /v1/betaGroups | 
[**BetaGroupsDeleteInstance**](BetaGroupsApi.md#BetaGroupsDeleteInstance) | **Delete** /v1/betaGroups/{id} | 
[**BetaGroupsGetCollection**](BetaGroupsApi.md#BetaGroupsGetCollection) | **Get** /v1/betaGroups | 
[**BetaGroupsGetInstance**](BetaGroupsApi.md#BetaGroupsGetInstance) | **Get** /v1/betaGroups/{id} | 
[**BetaGroupsUpdateInstance**](BetaGroupsApi.md#BetaGroupsUpdateInstance) | **Patch** /v1/betaGroups/{id} | 



## BetaGroupsAppGetToOneRelated

> AppResponse BetaGroupsAppGetToOneRelated(ctx, id).FieldsApps(fieldsApps).Execute()



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
    resp, r, err := api_client.BetaGroupsApi.BetaGroupsAppGetToOneRelated(context.Background(), id).FieldsApps(fieldsApps).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaGroupsApi.BetaGroupsAppGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaGroupsAppGetToOneRelated`: AppResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaGroupsApi.BetaGroupsAppGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaGroupsAppGetToOneRelatedRequest struct via the builder pattern


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


## BetaGroupsBetaTestersCreateToManyRelationship

> BetaGroupsBetaTestersCreateToManyRelationship(ctx, id).BetaGroupBetaTestersLinkagesRequest(betaGroupBetaTestersLinkagesRequest).Execute()



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
    betaGroupBetaTestersLinkagesRequest := *openapiclient.NewBetaGroupBetaTestersLinkagesRequest([]openapiclient.BetaGroupRelationshipsBetaTestersData{*openapiclient.NewBetaGroupRelationshipsBetaTestersData("Type_example", "Id_example")}) // BetaGroupBetaTestersLinkagesRequest | List of related linkages

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaGroupsApi.BetaGroupsBetaTestersCreateToManyRelationship(context.Background(), id).BetaGroupBetaTestersLinkagesRequest(betaGroupBetaTestersLinkagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaGroupsApi.BetaGroupsBetaTestersCreateToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBetaGroupsBetaTestersCreateToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaGroupBetaTestersLinkagesRequest** | [**BetaGroupBetaTestersLinkagesRequest**](BetaGroupBetaTestersLinkagesRequest.md) | List of related linkages | 

### Return type

 (empty response body)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaGroupsBetaTestersDeleteToManyRelationship

> BetaGroupsBetaTestersDeleteToManyRelationship(ctx, id).BetaGroupBetaTestersLinkagesRequest(betaGroupBetaTestersLinkagesRequest).Execute()



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
    betaGroupBetaTestersLinkagesRequest := *openapiclient.NewBetaGroupBetaTestersLinkagesRequest([]openapiclient.BetaGroupRelationshipsBetaTestersData{*openapiclient.NewBetaGroupRelationshipsBetaTestersData("Type_example", "Id_example")}) // BetaGroupBetaTestersLinkagesRequest | List of related linkages

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaGroupsApi.BetaGroupsBetaTestersDeleteToManyRelationship(context.Background(), id).BetaGroupBetaTestersLinkagesRequest(betaGroupBetaTestersLinkagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaGroupsApi.BetaGroupsBetaTestersDeleteToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBetaGroupsBetaTestersDeleteToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaGroupBetaTestersLinkagesRequest** | [**BetaGroupBetaTestersLinkagesRequest**](BetaGroupBetaTestersLinkagesRequest.md) | List of related linkages | 

### Return type

 (empty response body)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaGroupsBetaTestersGetToManyRelated

> BetaTestersResponse BetaGroupsBetaTestersGetToManyRelated(ctx, id).FieldsBetaTesters(fieldsBetaTesters).Limit(limit).Execute()



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
    fieldsBetaTesters := []string{"FieldsBetaTesters_example"} // []string | the fields to include for returned resources of type betaTesters (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaGroupsApi.BetaGroupsBetaTestersGetToManyRelated(context.Background(), id).FieldsBetaTesters(fieldsBetaTesters).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaGroupsApi.BetaGroupsBetaTestersGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaGroupsBetaTestersGetToManyRelated`: BetaTestersResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaGroupsApi.BetaGroupsBetaTestersGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaGroupsBetaTestersGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaTesters** | **[]string** | the fields to include for returned resources of type betaTesters | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**BetaTestersResponse**](BetaTestersResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaGroupsBetaTestersGetToManyRelationship

> BetaGroupBetaTestersLinkagesResponse BetaGroupsBetaTestersGetToManyRelationship(ctx, id).Limit(limit).Execute()



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
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaGroupsApi.BetaGroupsBetaTestersGetToManyRelationship(context.Background(), id).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaGroupsApi.BetaGroupsBetaTestersGetToManyRelationship``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaGroupsBetaTestersGetToManyRelationship`: BetaGroupBetaTestersLinkagesResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaGroupsApi.BetaGroupsBetaTestersGetToManyRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaGroupsBetaTestersGetToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum resources per page | 

### Return type

[**BetaGroupBetaTestersLinkagesResponse**](BetaGroupBetaTestersLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaGroupsBuildsCreateToManyRelationship

> BetaGroupsBuildsCreateToManyRelationship(ctx, id).BetaGroupBuildsLinkagesRequest(betaGroupBuildsLinkagesRequest).Execute()



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
    betaGroupBuildsLinkagesRequest := *openapiclient.NewBetaGroupBuildsLinkagesRequest([]openapiclient.AppStoreVersionRelationshipsBuildData{*openapiclient.NewAppStoreVersionRelationshipsBuildData("Type_example", "Id_example")}) // BetaGroupBuildsLinkagesRequest | List of related linkages

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaGroupsApi.BetaGroupsBuildsCreateToManyRelationship(context.Background(), id).BetaGroupBuildsLinkagesRequest(betaGroupBuildsLinkagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaGroupsApi.BetaGroupsBuildsCreateToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBetaGroupsBuildsCreateToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaGroupBuildsLinkagesRequest** | [**BetaGroupBuildsLinkagesRequest**](BetaGroupBuildsLinkagesRequest.md) | List of related linkages | 

### Return type

 (empty response body)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaGroupsBuildsDeleteToManyRelationship

> BetaGroupsBuildsDeleteToManyRelationship(ctx, id).BetaGroupBuildsLinkagesRequest(betaGroupBuildsLinkagesRequest).Execute()



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
    betaGroupBuildsLinkagesRequest := *openapiclient.NewBetaGroupBuildsLinkagesRequest([]openapiclient.AppStoreVersionRelationshipsBuildData{*openapiclient.NewAppStoreVersionRelationshipsBuildData("Type_example", "Id_example")}) // BetaGroupBuildsLinkagesRequest | List of related linkages

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaGroupsApi.BetaGroupsBuildsDeleteToManyRelationship(context.Background(), id).BetaGroupBuildsLinkagesRequest(betaGroupBuildsLinkagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaGroupsApi.BetaGroupsBuildsDeleteToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBetaGroupsBuildsDeleteToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaGroupBuildsLinkagesRequest** | [**BetaGroupBuildsLinkagesRequest**](BetaGroupBuildsLinkagesRequest.md) | List of related linkages | 

### Return type

 (empty response body)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaGroupsBuildsGetToManyRelated

> BuildsResponse BetaGroupsBuildsGetToManyRelated(ctx, id).FieldsBuilds(fieldsBuilds).Limit(limit).Execute()



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
    resp, r, err := api_client.BetaGroupsApi.BetaGroupsBuildsGetToManyRelated(context.Background(), id).FieldsBuilds(fieldsBuilds).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaGroupsApi.BetaGroupsBuildsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaGroupsBuildsGetToManyRelated`: BuildsResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaGroupsApi.BetaGroupsBuildsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaGroupsBuildsGetToManyRelatedRequest struct via the builder pattern


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


## BetaGroupsBuildsGetToManyRelationship

> BetaGroupBuildsLinkagesResponse BetaGroupsBuildsGetToManyRelationship(ctx, id).Limit(limit).Execute()



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
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaGroupsApi.BetaGroupsBuildsGetToManyRelationship(context.Background(), id).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaGroupsApi.BetaGroupsBuildsGetToManyRelationship``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaGroupsBuildsGetToManyRelationship`: BetaGroupBuildsLinkagesResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaGroupsApi.BetaGroupsBuildsGetToManyRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaGroupsBuildsGetToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum resources per page | 

### Return type

[**BetaGroupBuildsLinkagesResponse**](BetaGroupBuildsLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaGroupsCreateInstance

> BetaGroupResponse BetaGroupsCreateInstance(ctx).BetaGroupCreateRequest(betaGroupCreateRequest).Execute()



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
    betaGroupCreateRequest := *openapiclient.NewBetaGroupCreateRequest(*openapiclient.NewBetaGroupCreateRequestData("Type_example", *openapiclient.NewBetaGroupCreateRequestDataAttributes("Name_example"), *openapiclient.NewBetaGroupCreateRequestDataRelationships(*openapiclient.NewAppPreOrderCreateRequestDataRelationshipsApp(*openapiclient.NewAppEncryptionDeclarationRelationshipsAppData("Type_example", "Id_example"))))) // BetaGroupCreateRequest | BetaGroup representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaGroupsApi.BetaGroupsCreateInstance(context.Background()).BetaGroupCreateRequest(betaGroupCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaGroupsApi.BetaGroupsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaGroupsCreateInstance`: BetaGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaGroupsApi.BetaGroupsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBetaGroupsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **betaGroupCreateRequest** | [**BetaGroupCreateRequest**](BetaGroupCreateRequest.md) | BetaGroup representation | 

### Return type

[**BetaGroupResponse**](BetaGroupResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaGroupsDeleteInstance

> BetaGroupsDeleteInstance(ctx, id).Execute()



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
    resp, r, err := api_client.BetaGroupsApi.BetaGroupsDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaGroupsApi.BetaGroupsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBetaGroupsDeleteInstanceRequest struct via the builder pattern


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


## BetaGroupsGetCollection

> BetaGroupsResponse BetaGroupsGetCollection(ctx).FilterIsInternalGroup(filterIsInternalGroup).FilterName(filterName).FilterPublicLink(filterPublicLink).FilterPublicLinkEnabled(filterPublicLinkEnabled).FilterPublicLinkLimitEnabled(filterPublicLinkLimitEnabled).FilterApp(filterApp).FilterBuilds(filterBuilds).FilterId(filterId).Sort(sort).FieldsBetaGroups(fieldsBetaGroups).Limit(limit).Include(include).FieldsBuilds(fieldsBuilds).FieldsBetaTesters(fieldsBetaTesters).FieldsApps(fieldsApps).LimitBetaTesters(limitBetaTesters).LimitBuilds(limitBuilds).Execute()



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
    filterIsInternalGroup := []string{"Inner_example"} // []string | filter by attribute 'isInternalGroup' (optional)
    filterName := []string{"Inner_example"} // []string | filter by attribute 'name' (optional)
    filterPublicLink := []string{"Inner_example"} // []string | filter by attribute 'publicLink' (optional)
    filterPublicLinkEnabled := []string{"Inner_example"} // []string | filter by attribute 'publicLinkEnabled' (optional)
    filterPublicLinkLimitEnabled := []string{"Inner_example"} // []string | filter by attribute 'publicLinkLimitEnabled' (optional)
    filterApp := []string{"Inner_example"} // []string | filter by id(s) of related 'app' (optional)
    filterBuilds := []string{"Inner_example"} // []string | filter by id(s) of related 'builds' (optional)
    filterId := []string{"Inner_example"} // []string | filter by id(s) (optional)
    sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; resources will be sorted as specified (optional)
    fieldsBetaGroups := []string{"FieldsBetaGroups_example"} // []string | the fields to include for returned resources of type betaGroups (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)
    fieldsBetaTesters := []string{"FieldsBetaTesters_example"} // []string | the fields to include for returned resources of type betaTesters (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
    limitBetaTesters := int32(56) // int32 | maximum number of related betaTesters returned (when they are included) (optional)
    limitBuilds := int32(56) // int32 | maximum number of related builds returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaGroupsApi.BetaGroupsGetCollection(context.Background()).FilterIsInternalGroup(filterIsInternalGroup).FilterName(filterName).FilterPublicLink(filterPublicLink).FilterPublicLinkEnabled(filterPublicLinkEnabled).FilterPublicLinkLimitEnabled(filterPublicLinkLimitEnabled).FilterApp(filterApp).FilterBuilds(filterBuilds).FilterId(filterId).Sort(sort).FieldsBetaGroups(fieldsBetaGroups).Limit(limit).Include(include).FieldsBuilds(fieldsBuilds).FieldsBetaTesters(fieldsBetaTesters).FieldsApps(fieldsApps).LimitBetaTesters(limitBetaTesters).LimitBuilds(limitBuilds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaGroupsApi.BetaGroupsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaGroupsGetCollection`: BetaGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaGroupsApi.BetaGroupsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBetaGroupsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterIsInternalGroup** | **[]string** | filter by attribute &#39;isInternalGroup&#39; | 
 **filterName** | **[]string** | filter by attribute &#39;name&#39; | 
 **filterPublicLink** | **[]string** | filter by attribute &#39;publicLink&#39; | 
 **filterPublicLinkEnabled** | **[]string** | filter by attribute &#39;publicLinkEnabled&#39; | 
 **filterPublicLinkLimitEnabled** | **[]string** | filter by attribute &#39;publicLinkLimitEnabled&#39; | 
 **filterApp** | **[]string** | filter by id(s) of related &#39;app&#39; | 
 **filterBuilds** | **[]string** | filter by id(s) of related &#39;builds&#39; | 
 **filterId** | **[]string** | filter by id(s) | 
 **sort** | **[]string** | comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsBetaGroups** | **[]string** | the fields to include for returned resources of type betaGroups | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 
 **fieldsBetaTesters** | **[]string** | the fields to include for returned resources of type betaTesters | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **limitBetaTesters** | **int32** | maximum number of related betaTesters returned (when they are included) | 
 **limitBuilds** | **int32** | maximum number of related builds returned (when they are included) | 

### Return type

[**BetaGroupsResponse**](BetaGroupsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaGroupsGetInstance

> BetaGroupResponse BetaGroupsGetInstance(ctx, id).FieldsBetaGroups(fieldsBetaGroups).Include(include).FieldsBuilds(fieldsBuilds).FieldsBetaTesters(fieldsBetaTesters).FieldsApps(fieldsApps).LimitBetaTesters(limitBetaTesters).LimitBuilds(limitBuilds).Execute()



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
    fieldsBetaGroups := []string{"FieldsBetaGroups_example"} // []string | the fields to include for returned resources of type betaGroups (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)
    fieldsBetaTesters := []string{"FieldsBetaTesters_example"} // []string | the fields to include for returned resources of type betaTesters (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
    limitBetaTesters := int32(56) // int32 | maximum number of related betaTesters returned (when they are included) (optional)
    limitBuilds := int32(56) // int32 | maximum number of related builds returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaGroupsApi.BetaGroupsGetInstance(context.Background(), id).FieldsBetaGroups(fieldsBetaGroups).Include(include).FieldsBuilds(fieldsBuilds).FieldsBetaTesters(fieldsBetaTesters).FieldsApps(fieldsApps).LimitBetaTesters(limitBetaTesters).LimitBuilds(limitBuilds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaGroupsApi.BetaGroupsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaGroupsGetInstance`: BetaGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaGroupsApi.BetaGroupsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaGroupsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaGroups** | **[]string** | the fields to include for returned resources of type betaGroups | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 
 **fieldsBetaTesters** | **[]string** | the fields to include for returned resources of type betaTesters | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **limitBetaTesters** | **int32** | maximum number of related betaTesters returned (when they are included) | 
 **limitBuilds** | **int32** | maximum number of related builds returned (when they are included) | 

### Return type

[**BetaGroupResponse**](BetaGroupResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaGroupsUpdateInstance

> BetaGroupResponse BetaGroupsUpdateInstance(ctx, id).BetaGroupUpdateRequest(betaGroupUpdateRequest).Execute()



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
    betaGroupUpdateRequest := *openapiclient.NewBetaGroupUpdateRequest(*openapiclient.NewBetaGroupUpdateRequestData("Type_example", "Id_example")) // BetaGroupUpdateRequest | BetaGroup representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaGroupsApi.BetaGroupsUpdateInstance(context.Background(), id).BetaGroupUpdateRequest(betaGroupUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaGroupsApi.BetaGroupsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaGroupsUpdateInstance`: BetaGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaGroupsApi.BetaGroupsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaGroupsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaGroupUpdateRequest** | [**BetaGroupUpdateRequest**](BetaGroupUpdateRequest.md) | BetaGroup representation | 

### Return type

[**BetaGroupResponse**](BetaGroupResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

