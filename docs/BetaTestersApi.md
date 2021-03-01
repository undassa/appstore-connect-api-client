# \BetaTestersApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BetaTestersAppsDeleteToManyRelationship**](BetaTestersApi.md#BetaTestersAppsDeleteToManyRelationship) | **Delete** /v1/betaTesters/{id}/relationships/apps | 
[**BetaTestersAppsGetToManyRelated**](BetaTestersApi.md#BetaTestersAppsGetToManyRelated) | **Get** /v1/betaTesters/{id}/apps | 
[**BetaTestersAppsGetToManyRelationship**](BetaTestersApi.md#BetaTestersAppsGetToManyRelationship) | **Get** /v1/betaTesters/{id}/relationships/apps | 
[**BetaTestersBetaGroupsCreateToManyRelationship**](BetaTestersApi.md#BetaTestersBetaGroupsCreateToManyRelationship) | **Post** /v1/betaTesters/{id}/relationships/betaGroups | 
[**BetaTestersBetaGroupsDeleteToManyRelationship**](BetaTestersApi.md#BetaTestersBetaGroupsDeleteToManyRelationship) | **Delete** /v1/betaTesters/{id}/relationships/betaGroups | 
[**BetaTestersBetaGroupsGetToManyRelated**](BetaTestersApi.md#BetaTestersBetaGroupsGetToManyRelated) | **Get** /v1/betaTesters/{id}/betaGroups | 
[**BetaTestersBetaGroupsGetToManyRelationship**](BetaTestersApi.md#BetaTestersBetaGroupsGetToManyRelationship) | **Get** /v1/betaTesters/{id}/relationships/betaGroups | 
[**BetaTestersBuildsCreateToManyRelationship**](BetaTestersApi.md#BetaTestersBuildsCreateToManyRelationship) | **Post** /v1/betaTesters/{id}/relationships/builds | 
[**BetaTestersBuildsDeleteToManyRelationship**](BetaTestersApi.md#BetaTestersBuildsDeleteToManyRelationship) | **Delete** /v1/betaTesters/{id}/relationships/builds | 
[**BetaTestersBuildsGetToManyRelated**](BetaTestersApi.md#BetaTestersBuildsGetToManyRelated) | **Get** /v1/betaTesters/{id}/builds | 
[**BetaTestersBuildsGetToManyRelationship**](BetaTestersApi.md#BetaTestersBuildsGetToManyRelationship) | **Get** /v1/betaTesters/{id}/relationships/builds | 
[**BetaTestersCreateInstance**](BetaTestersApi.md#BetaTestersCreateInstance) | **Post** /v1/betaTesters | 
[**BetaTestersDeleteInstance**](BetaTestersApi.md#BetaTestersDeleteInstance) | **Delete** /v1/betaTesters/{id} | 
[**BetaTestersGetCollection**](BetaTestersApi.md#BetaTestersGetCollection) | **Get** /v1/betaTesters | 
[**BetaTestersGetInstance**](BetaTestersApi.md#BetaTestersGetInstance) | **Get** /v1/betaTesters/{id} | 



## BetaTestersAppsDeleteToManyRelationship

> BetaTestersAppsDeleteToManyRelationship(ctx, id).BetaTesterAppsLinkagesRequest(betaTesterAppsLinkagesRequest).Execute()



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
    betaTesterAppsLinkagesRequest := *openapiclient.NewBetaTesterAppsLinkagesRequest([]openapiclient.AppEncryptionDeclarationRelationshipsAppData{*openapiclient.NewAppEncryptionDeclarationRelationshipsAppData("Type_example", "Id_example")}) // BetaTesterAppsLinkagesRequest | List of related linkages

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaTestersApi.BetaTestersAppsDeleteToManyRelationship(context.Background(), id).BetaTesterAppsLinkagesRequest(betaTesterAppsLinkagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersApi.BetaTestersAppsDeleteToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBetaTestersAppsDeleteToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaTesterAppsLinkagesRequest** | [**BetaTesterAppsLinkagesRequest**](BetaTesterAppsLinkagesRequest.md) | List of related linkages | 

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


## BetaTestersAppsGetToManyRelated

> AppsResponse BetaTestersAppsGetToManyRelated(ctx, id).FieldsApps(fieldsApps).Limit(limit).Execute()



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
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaTestersApi.BetaTestersAppsGetToManyRelated(context.Background(), id).FieldsApps(fieldsApps).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersApi.BetaTestersAppsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaTestersAppsGetToManyRelated`: AppsResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaTestersApi.BetaTestersAppsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaTestersAppsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**AppsResponse**](AppsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaTestersAppsGetToManyRelationship

> BetaTesterAppsLinkagesResponse BetaTestersAppsGetToManyRelationship(ctx, id).Limit(limit).Execute()



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
    resp, r, err := api_client.BetaTestersApi.BetaTestersAppsGetToManyRelationship(context.Background(), id).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersApi.BetaTestersAppsGetToManyRelationship``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaTestersAppsGetToManyRelationship`: BetaTesterAppsLinkagesResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaTestersApi.BetaTestersAppsGetToManyRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaTestersAppsGetToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum resources per page | 

### Return type

[**BetaTesterAppsLinkagesResponse**](BetaTesterAppsLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaTestersBetaGroupsCreateToManyRelationship

> BetaTestersBetaGroupsCreateToManyRelationship(ctx, id).BetaTesterBetaGroupsLinkagesRequest(betaTesterBetaGroupsLinkagesRequest).Execute()



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
    betaTesterBetaGroupsLinkagesRequest := *openapiclient.NewBetaTesterBetaGroupsLinkagesRequest([]openapiclient.AppRelationshipsBetaGroupsData{*openapiclient.NewAppRelationshipsBetaGroupsData("Type_example", "Id_example")}) // BetaTesterBetaGroupsLinkagesRequest | List of related linkages

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaTestersApi.BetaTestersBetaGroupsCreateToManyRelationship(context.Background(), id).BetaTesterBetaGroupsLinkagesRequest(betaTesterBetaGroupsLinkagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersApi.BetaTestersBetaGroupsCreateToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBetaTestersBetaGroupsCreateToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaTesterBetaGroupsLinkagesRequest** | [**BetaTesterBetaGroupsLinkagesRequest**](BetaTesterBetaGroupsLinkagesRequest.md) | List of related linkages | 

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


## BetaTestersBetaGroupsDeleteToManyRelationship

> BetaTestersBetaGroupsDeleteToManyRelationship(ctx, id).BetaTesterBetaGroupsLinkagesRequest(betaTesterBetaGroupsLinkagesRequest).Execute()



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
    betaTesterBetaGroupsLinkagesRequest := *openapiclient.NewBetaTesterBetaGroupsLinkagesRequest([]openapiclient.AppRelationshipsBetaGroupsData{*openapiclient.NewAppRelationshipsBetaGroupsData("Type_example", "Id_example")}) // BetaTesterBetaGroupsLinkagesRequest | List of related linkages

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaTestersApi.BetaTestersBetaGroupsDeleteToManyRelationship(context.Background(), id).BetaTesterBetaGroupsLinkagesRequest(betaTesterBetaGroupsLinkagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersApi.BetaTestersBetaGroupsDeleteToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBetaTestersBetaGroupsDeleteToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaTesterBetaGroupsLinkagesRequest** | [**BetaTesterBetaGroupsLinkagesRequest**](BetaTesterBetaGroupsLinkagesRequest.md) | List of related linkages | 

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


## BetaTestersBetaGroupsGetToManyRelated

> BetaGroupsResponse BetaTestersBetaGroupsGetToManyRelated(ctx, id).FieldsBetaGroups(fieldsBetaGroups).Limit(limit).Execute()



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
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaTestersApi.BetaTestersBetaGroupsGetToManyRelated(context.Background(), id).FieldsBetaGroups(fieldsBetaGroups).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersApi.BetaTestersBetaGroupsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaTestersBetaGroupsGetToManyRelated`: BetaGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaTestersApi.BetaTestersBetaGroupsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaTestersBetaGroupsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaGroups** | **[]string** | the fields to include for returned resources of type betaGroups | 
 **limit** | **int32** | maximum resources per page | 

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


## BetaTestersBetaGroupsGetToManyRelationship

> BetaTesterBetaGroupsLinkagesResponse BetaTestersBetaGroupsGetToManyRelationship(ctx, id).Limit(limit).Execute()



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
    resp, r, err := api_client.BetaTestersApi.BetaTestersBetaGroupsGetToManyRelationship(context.Background(), id).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersApi.BetaTestersBetaGroupsGetToManyRelationship``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaTestersBetaGroupsGetToManyRelationship`: BetaTesterBetaGroupsLinkagesResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaTestersApi.BetaTestersBetaGroupsGetToManyRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaTestersBetaGroupsGetToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum resources per page | 

### Return type

[**BetaTesterBetaGroupsLinkagesResponse**](BetaTesterBetaGroupsLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaTestersBuildsCreateToManyRelationship

> BetaTestersBuildsCreateToManyRelationship(ctx, id).BetaTesterBuildsLinkagesRequest(betaTesterBuildsLinkagesRequest).Execute()



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
    betaTesterBuildsLinkagesRequest := *openapiclient.NewBetaTesterBuildsLinkagesRequest([]openapiclient.AppStoreVersionRelationshipsBuildData{*openapiclient.NewAppStoreVersionRelationshipsBuildData("Type_example", "Id_example")}) // BetaTesterBuildsLinkagesRequest | List of related linkages

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaTestersApi.BetaTestersBuildsCreateToManyRelationship(context.Background(), id).BetaTesterBuildsLinkagesRequest(betaTesterBuildsLinkagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersApi.BetaTestersBuildsCreateToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBetaTestersBuildsCreateToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaTesterBuildsLinkagesRequest** | [**BetaTesterBuildsLinkagesRequest**](BetaTesterBuildsLinkagesRequest.md) | List of related linkages | 

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


## BetaTestersBuildsDeleteToManyRelationship

> BetaTestersBuildsDeleteToManyRelationship(ctx, id).BetaTesterBuildsLinkagesRequest(betaTesterBuildsLinkagesRequest).Execute()



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
    betaTesterBuildsLinkagesRequest := *openapiclient.NewBetaTesterBuildsLinkagesRequest([]openapiclient.AppStoreVersionRelationshipsBuildData{*openapiclient.NewAppStoreVersionRelationshipsBuildData("Type_example", "Id_example")}) // BetaTesterBuildsLinkagesRequest | List of related linkages

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaTestersApi.BetaTestersBuildsDeleteToManyRelationship(context.Background(), id).BetaTesterBuildsLinkagesRequest(betaTesterBuildsLinkagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersApi.BetaTestersBuildsDeleteToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBetaTestersBuildsDeleteToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaTesterBuildsLinkagesRequest** | [**BetaTesterBuildsLinkagesRequest**](BetaTesterBuildsLinkagesRequest.md) | List of related linkages | 

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


## BetaTestersBuildsGetToManyRelated

> BuildsResponse BetaTestersBuildsGetToManyRelated(ctx, id).FieldsBuilds(fieldsBuilds).Limit(limit).Execute()



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
    resp, r, err := api_client.BetaTestersApi.BetaTestersBuildsGetToManyRelated(context.Background(), id).FieldsBuilds(fieldsBuilds).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersApi.BetaTestersBuildsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaTestersBuildsGetToManyRelated`: BuildsResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaTestersApi.BetaTestersBuildsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaTestersBuildsGetToManyRelatedRequest struct via the builder pattern


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


## BetaTestersBuildsGetToManyRelationship

> BetaTesterBuildsLinkagesResponse BetaTestersBuildsGetToManyRelationship(ctx, id).Limit(limit).Execute()



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
    resp, r, err := api_client.BetaTestersApi.BetaTestersBuildsGetToManyRelationship(context.Background(), id).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersApi.BetaTestersBuildsGetToManyRelationship``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaTestersBuildsGetToManyRelationship`: BetaTesterBuildsLinkagesResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaTestersApi.BetaTestersBuildsGetToManyRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaTestersBuildsGetToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum resources per page | 

### Return type

[**BetaTesterBuildsLinkagesResponse**](BetaTesterBuildsLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaTestersCreateInstance

> BetaTesterResponse BetaTestersCreateInstance(ctx).BetaTesterCreateRequest(betaTesterCreateRequest).Execute()



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
    betaTesterCreateRequest := *openapiclient.NewBetaTesterCreateRequest(*openapiclient.NewBetaTesterCreateRequestData("Type_example", *openapiclient.NewBetaTesterCreateRequestDataAttributes("Email_example"))) // BetaTesterCreateRequest | BetaTester representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaTestersApi.BetaTestersCreateInstance(context.Background()).BetaTesterCreateRequest(betaTesterCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersApi.BetaTestersCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaTestersCreateInstance`: BetaTesterResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaTestersApi.BetaTestersCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBetaTestersCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **betaTesterCreateRequest** | [**BetaTesterCreateRequest**](BetaTesterCreateRequest.md) | BetaTester representation | 

### Return type

[**BetaTesterResponse**](BetaTesterResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaTestersDeleteInstance

> BetaTestersDeleteInstance(ctx, id).Execute()



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
    resp, r, err := api_client.BetaTestersApi.BetaTestersDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersApi.BetaTestersDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBetaTestersDeleteInstanceRequest struct via the builder pattern


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


## BetaTestersGetCollection

> BetaTestersResponse BetaTestersGetCollection(ctx).FilterEmail(filterEmail).FilterFirstName(filterFirstName).FilterInviteType(filterInviteType).FilterLastName(filterLastName).FilterApps(filterApps).FilterBetaGroups(filterBetaGroups).FilterBuilds(filterBuilds).Sort(sort).FieldsBetaTesters(fieldsBetaTesters).Limit(limit).Include(include).FieldsBetaGroups(fieldsBetaGroups).FieldsBuilds(fieldsBuilds).FieldsApps(fieldsApps).LimitApps(limitApps).LimitBetaGroups(limitBetaGroups).LimitBuilds(limitBuilds).Execute()



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
    filterEmail := []string{"Inner_example"} // []string | filter by attribute 'email' (optional)
    filterFirstName := []string{"Inner_example"} // []string | filter by attribute 'firstName' (optional)
    filterInviteType := []string{"FilterInviteType_example"} // []string | filter by attribute 'inviteType' (optional)
    filterLastName := []string{"Inner_example"} // []string | filter by attribute 'lastName' (optional)
    filterApps := []string{"Inner_example"} // []string | filter by id(s) of related 'apps' (optional)
    filterBetaGroups := []string{"Inner_example"} // []string | filter by id(s) of related 'betaGroups' (optional)
    filterBuilds := []string{"Inner_example"} // []string | filter by id(s) of related 'builds' (optional)
    sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; resources will be sorted as specified (optional)
    fieldsBetaTesters := []string{"FieldsBetaTesters_example"} // []string | the fields to include for returned resources of type betaTesters (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsBetaGroups := []string{"FieldsBetaGroups_example"} // []string | the fields to include for returned resources of type betaGroups (optional)
    fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
    limitApps := int32(56) // int32 | maximum number of related apps returned (when they are included) (optional)
    limitBetaGroups := int32(56) // int32 | maximum number of related betaGroups returned (when they are included) (optional)
    limitBuilds := int32(56) // int32 | maximum number of related builds returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaTestersApi.BetaTestersGetCollection(context.Background()).FilterEmail(filterEmail).FilterFirstName(filterFirstName).FilterInviteType(filterInviteType).FilterLastName(filterLastName).FilterApps(filterApps).FilterBetaGroups(filterBetaGroups).FilterBuilds(filterBuilds).Sort(sort).FieldsBetaTesters(fieldsBetaTesters).Limit(limit).Include(include).FieldsBetaGroups(fieldsBetaGroups).FieldsBuilds(fieldsBuilds).FieldsApps(fieldsApps).LimitApps(limitApps).LimitBetaGroups(limitBetaGroups).LimitBuilds(limitBuilds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersApi.BetaTestersGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaTestersGetCollection`: BetaTestersResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaTestersApi.BetaTestersGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBetaTestersGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterEmail** | **[]string** | filter by attribute &#39;email&#39; | 
 **filterFirstName** | **[]string** | filter by attribute &#39;firstName&#39; | 
 **filterInviteType** | **[]string** | filter by attribute &#39;inviteType&#39; | 
 **filterLastName** | **[]string** | filter by attribute &#39;lastName&#39; | 
 **filterApps** | **[]string** | filter by id(s) of related &#39;apps&#39; | 
 **filterBetaGroups** | **[]string** | filter by id(s) of related &#39;betaGroups&#39; | 
 **filterBuilds** | **[]string** | filter by id(s) of related &#39;builds&#39; | 
 **sort** | **[]string** | comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsBetaTesters** | **[]string** | the fields to include for returned resources of type betaTesters | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsBetaGroups** | **[]string** | the fields to include for returned resources of type betaGroups | 
 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **limitApps** | **int32** | maximum number of related apps returned (when they are included) | 
 **limitBetaGroups** | **int32** | maximum number of related betaGroups returned (when they are included) | 
 **limitBuilds** | **int32** | maximum number of related builds returned (when they are included) | 

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


## BetaTestersGetInstance

> BetaTesterResponse BetaTestersGetInstance(ctx, id).FieldsBetaTesters(fieldsBetaTesters).Include(include).FieldsBetaGroups(fieldsBetaGroups).FieldsBuilds(fieldsBuilds).FieldsApps(fieldsApps).LimitApps(limitApps).LimitBetaGroups(limitBetaGroups).LimitBuilds(limitBuilds).Execute()



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
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsBetaGroups := []string{"FieldsBetaGroups_example"} // []string | the fields to include for returned resources of type betaGroups (optional)
    fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
    limitApps := int32(56) // int32 | maximum number of related apps returned (when they are included) (optional)
    limitBetaGroups := int32(56) // int32 | maximum number of related betaGroups returned (when they are included) (optional)
    limitBuilds := int32(56) // int32 | maximum number of related builds returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaTestersApi.BetaTestersGetInstance(context.Background(), id).FieldsBetaTesters(fieldsBetaTesters).Include(include).FieldsBetaGroups(fieldsBetaGroups).FieldsBuilds(fieldsBuilds).FieldsApps(fieldsApps).LimitApps(limitApps).LimitBetaGroups(limitBetaGroups).LimitBuilds(limitBuilds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersApi.BetaTestersGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaTestersGetInstance`: BetaTesterResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaTestersApi.BetaTestersGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaTestersGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaTesters** | **[]string** | the fields to include for returned resources of type betaTesters | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsBetaGroups** | **[]string** | the fields to include for returned resources of type betaGroups | 
 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **limitApps** | **int32** | maximum number of related apps returned (when they are included) | 
 **limitBetaGroups** | **int32** | maximum number of related betaGroups returned (when they are included) | 
 **limitBuilds** | **int32** | maximum number of related builds returned (when they are included) | 

### Return type

[**BetaTesterResponse**](BetaTesterResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

