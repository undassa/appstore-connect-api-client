# \AppStoreVersionsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppStoreVersionsAgeRatingDeclarationGetToOneRelated**](AppStoreVersionsApi.md#AppStoreVersionsAgeRatingDeclarationGetToOneRelated) | **Get** /v1/appStoreVersions/{id}/ageRatingDeclaration | 
[**AppStoreVersionsAppStoreReviewDetailGetToOneRelated**](AppStoreVersionsApi.md#AppStoreVersionsAppStoreReviewDetailGetToOneRelated) | **Get** /v1/appStoreVersions/{id}/appStoreReviewDetail | 
[**AppStoreVersionsAppStoreVersionLocalizationsGetToManyRelated**](AppStoreVersionsApi.md#AppStoreVersionsAppStoreVersionLocalizationsGetToManyRelated) | **Get** /v1/appStoreVersions/{id}/appStoreVersionLocalizations | 
[**AppStoreVersionsAppStoreVersionPhasedReleaseGetToOneRelated**](AppStoreVersionsApi.md#AppStoreVersionsAppStoreVersionPhasedReleaseGetToOneRelated) | **Get** /v1/appStoreVersions/{id}/appStoreVersionPhasedRelease | 
[**AppStoreVersionsAppStoreVersionSubmissionGetToOneRelated**](AppStoreVersionsApi.md#AppStoreVersionsAppStoreVersionSubmissionGetToOneRelated) | **Get** /v1/appStoreVersions/{id}/appStoreVersionSubmission | 
[**AppStoreVersionsBuildGetToOneRelated**](AppStoreVersionsApi.md#AppStoreVersionsBuildGetToOneRelated) | **Get** /v1/appStoreVersions/{id}/build | 
[**AppStoreVersionsBuildGetToOneRelationship**](AppStoreVersionsApi.md#AppStoreVersionsBuildGetToOneRelationship) | **Get** /v1/appStoreVersions/{id}/relationships/build | 
[**AppStoreVersionsBuildUpdateToOneRelationship**](AppStoreVersionsApi.md#AppStoreVersionsBuildUpdateToOneRelationship) | **Patch** /v1/appStoreVersions/{id}/relationships/build | 
[**AppStoreVersionsCreateInstance**](AppStoreVersionsApi.md#AppStoreVersionsCreateInstance) | **Post** /v1/appStoreVersions | 
[**AppStoreVersionsDeleteInstance**](AppStoreVersionsApi.md#AppStoreVersionsDeleteInstance) | **Delete** /v1/appStoreVersions/{id} | 
[**AppStoreVersionsGetInstance**](AppStoreVersionsApi.md#AppStoreVersionsGetInstance) | **Get** /v1/appStoreVersions/{id} | 
[**AppStoreVersionsIdfaDeclarationGetToOneRelated**](AppStoreVersionsApi.md#AppStoreVersionsIdfaDeclarationGetToOneRelated) | **Get** /v1/appStoreVersions/{id}/idfaDeclaration | 
[**AppStoreVersionsRoutingAppCoverageGetToOneRelated**](AppStoreVersionsApi.md#AppStoreVersionsRoutingAppCoverageGetToOneRelated) | **Get** /v1/appStoreVersions/{id}/routingAppCoverage | 
[**AppStoreVersionsUpdateInstance**](AppStoreVersionsApi.md#AppStoreVersionsUpdateInstance) | **Patch** /v1/appStoreVersions/{id} | 



## AppStoreVersionsAgeRatingDeclarationGetToOneRelated

> AgeRatingDeclarationResponse AppStoreVersionsAgeRatingDeclarationGetToOneRelated(ctx, id).FieldsAgeRatingDeclarations(fieldsAgeRatingDeclarations).Execute()



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
    fieldsAgeRatingDeclarations := []string{"FieldsAgeRatingDeclarations_example"} // []string | the fields to include for returned resources of type ageRatingDeclarations (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppStoreVersionsApi.AppStoreVersionsAgeRatingDeclarationGetToOneRelated(context.Background(), id).FieldsAgeRatingDeclarations(fieldsAgeRatingDeclarations).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsApi.AppStoreVersionsAgeRatingDeclarationGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppStoreVersionsAgeRatingDeclarationGetToOneRelated`: AgeRatingDeclarationResponse
    fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsApi.AppStoreVersionsAgeRatingDeclarationGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsAgeRatingDeclarationGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAgeRatingDeclarations** | **[]string** | the fields to include for returned resources of type ageRatingDeclarations | 

### Return type

[**AgeRatingDeclarationResponse**](AgeRatingDeclarationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsAppStoreReviewDetailGetToOneRelated

> AppStoreReviewDetailResponse AppStoreVersionsAppStoreReviewDetailGetToOneRelated(ctx, id).FieldsAppStoreReviewDetails(fieldsAppStoreReviewDetails).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsAppStoreReviewAttachments(fieldsAppStoreReviewAttachments).Include(include).Execute()



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
    fieldsAppStoreReviewDetails := []string{"FieldsAppStoreReviewDetails_example"} // []string | the fields to include for returned resources of type appStoreReviewDetails (optional)
    fieldsAppStoreVersions := []string{"FieldsAppStoreVersions_example"} // []string | the fields to include for returned resources of type appStoreVersions (optional)
    fieldsAppStoreReviewAttachments := []string{"FieldsAppStoreReviewAttachments_example"} // []string | the fields to include for returned resources of type appStoreReviewAttachments (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppStoreVersionsApi.AppStoreVersionsAppStoreReviewDetailGetToOneRelated(context.Background(), id).FieldsAppStoreReviewDetails(fieldsAppStoreReviewDetails).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsAppStoreReviewAttachments(fieldsAppStoreReviewAttachments).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsApi.AppStoreVersionsAppStoreReviewDetailGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppStoreVersionsAppStoreReviewDetailGetToOneRelated`: AppStoreReviewDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsApi.AppStoreVersionsAppStoreReviewDetailGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsAppStoreReviewDetailGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppStoreReviewDetails** | **[]string** | the fields to include for returned resources of type appStoreReviewDetails | 
 **fieldsAppStoreVersions** | **[]string** | the fields to include for returned resources of type appStoreVersions | 
 **fieldsAppStoreReviewAttachments** | **[]string** | the fields to include for returned resources of type appStoreReviewAttachments | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**AppStoreReviewDetailResponse**](AppStoreReviewDetailResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsAppStoreVersionLocalizationsGetToManyRelated

> AppStoreVersionLocalizationsResponse AppStoreVersionsAppStoreVersionLocalizationsGetToManyRelated(ctx, id).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).Limit(limit).Execute()



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
    fieldsAppStoreVersionLocalizations := []string{"FieldsAppStoreVersionLocalizations_example"} // []string | the fields to include for returned resources of type appStoreVersionLocalizations (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppStoreVersionsApi.AppStoreVersionsAppStoreVersionLocalizationsGetToManyRelated(context.Background(), id).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsApi.AppStoreVersionsAppStoreVersionLocalizationsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppStoreVersionsAppStoreVersionLocalizationsGetToManyRelated`: AppStoreVersionLocalizationsResponse
    fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsApi.AppStoreVersionsAppStoreVersionLocalizationsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsAppStoreVersionLocalizationsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppStoreVersionLocalizations** | **[]string** | the fields to include for returned resources of type appStoreVersionLocalizations | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**AppStoreVersionLocalizationsResponse**](AppStoreVersionLocalizationsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsAppStoreVersionPhasedReleaseGetToOneRelated

> AppStoreVersionPhasedReleaseResponse AppStoreVersionsAppStoreVersionPhasedReleaseGetToOneRelated(ctx, id).FieldsAppStoreVersionPhasedReleases(fieldsAppStoreVersionPhasedReleases).Execute()



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
    fieldsAppStoreVersionPhasedReleases := []string{"FieldsAppStoreVersionPhasedReleases_example"} // []string | the fields to include for returned resources of type appStoreVersionPhasedReleases (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppStoreVersionsApi.AppStoreVersionsAppStoreVersionPhasedReleaseGetToOneRelated(context.Background(), id).FieldsAppStoreVersionPhasedReleases(fieldsAppStoreVersionPhasedReleases).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsApi.AppStoreVersionsAppStoreVersionPhasedReleaseGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppStoreVersionsAppStoreVersionPhasedReleaseGetToOneRelated`: AppStoreVersionPhasedReleaseResponse
    fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsApi.AppStoreVersionsAppStoreVersionPhasedReleaseGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsAppStoreVersionPhasedReleaseGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppStoreVersionPhasedReleases** | **[]string** | the fields to include for returned resources of type appStoreVersionPhasedReleases | 

### Return type

[**AppStoreVersionPhasedReleaseResponse**](AppStoreVersionPhasedReleaseResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsAppStoreVersionSubmissionGetToOneRelated

> AppStoreVersionSubmissionResponse AppStoreVersionsAppStoreVersionSubmissionGetToOneRelated(ctx, id).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsAppStoreVersionSubmissions(fieldsAppStoreVersionSubmissions).Include(include).Execute()



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
    fieldsAppStoreVersions := []string{"FieldsAppStoreVersions_example"} // []string | the fields to include for returned resources of type appStoreVersions (optional)
    fieldsAppStoreVersionSubmissions := []string{"FieldsAppStoreVersionSubmissions_example"} // []string | the fields to include for returned resources of type appStoreVersionSubmissions (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppStoreVersionsApi.AppStoreVersionsAppStoreVersionSubmissionGetToOneRelated(context.Background(), id).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsAppStoreVersionSubmissions(fieldsAppStoreVersionSubmissions).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsApi.AppStoreVersionsAppStoreVersionSubmissionGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppStoreVersionsAppStoreVersionSubmissionGetToOneRelated`: AppStoreVersionSubmissionResponse
    fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsApi.AppStoreVersionsAppStoreVersionSubmissionGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsAppStoreVersionSubmissionGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppStoreVersions** | **[]string** | the fields to include for returned resources of type appStoreVersions | 
 **fieldsAppStoreVersionSubmissions** | **[]string** | the fields to include for returned resources of type appStoreVersionSubmissions | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**AppStoreVersionSubmissionResponse**](AppStoreVersionSubmissionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsBuildGetToOneRelated

> BuildResponse AppStoreVersionsBuildGetToOneRelated(ctx, id).FieldsBuilds(fieldsBuilds).Execute()



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
    resp, r, err := api_client.AppStoreVersionsApi.AppStoreVersionsBuildGetToOneRelated(context.Background(), id).FieldsBuilds(fieldsBuilds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsApi.AppStoreVersionsBuildGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppStoreVersionsBuildGetToOneRelated`: BuildResponse
    fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsApi.AppStoreVersionsBuildGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsBuildGetToOneRelatedRequest struct via the builder pattern


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


## AppStoreVersionsBuildGetToOneRelationship

> AppStoreVersionBuildLinkageResponse AppStoreVersionsBuildGetToOneRelationship(ctx, id).Execute()



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
    resp, r, err := api_client.AppStoreVersionsApi.AppStoreVersionsBuildGetToOneRelationship(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsApi.AppStoreVersionsBuildGetToOneRelationship``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppStoreVersionsBuildGetToOneRelationship`: AppStoreVersionBuildLinkageResponse
    fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsApi.AppStoreVersionsBuildGetToOneRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsBuildGetToOneRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppStoreVersionBuildLinkageResponse**](AppStoreVersionBuildLinkageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsBuildUpdateToOneRelationship

> AppStoreVersionsBuildUpdateToOneRelationship(ctx, id).AppStoreVersionBuildLinkageRequest(appStoreVersionBuildLinkageRequest).Execute()



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
    appStoreVersionBuildLinkageRequest := *openapiclient.NewAppStoreVersionBuildLinkageRequest(*openapiclient.NewAppStoreVersionRelationshipsBuildData("Type_example", "Id_example")) // AppStoreVersionBuildLinkageRequest | Related linkage

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppStoreVersionsApi.AppStoreVersionsBuildUpdateToOneRelationship(context.Background(), id).AppStoreVersionBuildLinkageRequest(appStoreVersionBuildLinkageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsApi.AppStoreVersionsBuildUpdateToOneRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppStoreVersionsBuildUpdateToOneRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appStoreVersionBuildLinkageRequest** | [**AppStoreVersionBuildLinkageRequest**](AppStoreVersionBuildLinkageRequest.md) | Related linkage | 

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


## AppStoreVersionsCreateInstance

> AppStoreVersionResponse AppStoreVersionsCreateInstance(ctx).AppStoreVersionCreateRequest(appStoreVersionCreateRequest).Execute()



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
    appStoreVersionCreateRequest := *openapiclient.NewAppStoreVersionCreateRequest(*openapiclient.NewAppStoreVersionCreateRequestData("Type_example", *openapiclient.NewAppStoreVersionCreateRequestDataAttributes(openapiclient.Platform("IOS"), "VersionString_example"), *openapiclient.NewAppStoreVersionCreateRequestDataRelationships(*openapiclient.NewAppPreOrderCreateRequestDataRelationshipsApp(*openapiclient.NewAppEncryptionDeclarationRelationshipsAppData("Type_example", "Id_example"))))) // AppStoreVersionCreateRequest | AppStoreVersion representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppStoreVersionsApi.AppStoreVersionsCreateInstance(context.Background()).AppStoreVersionCreateRequest(appStoreVersionCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsApi.AppStoreVersionsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppStoreVersionsCreateInstance`: AppStoreVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsApi.AppStoreVersionsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appStoreVersionCreateRequest** | [**AppStoreVersionCreateRequest**](AppStoreVersionCreateRequest.md) | AppStoreVersion representation | 

### Return type

[**AppStoreVersionResponse**](AppStoreVersionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsDeleteInstance

> AppStoreVersionsDeleteInstance(ctx, id).Execute()



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
    resp, r, err := api_client.AppStoreVersionsApi.AppStoreVersionsDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsApi.AppStoreVersionsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppStoreVersionsDeleteInstanceRequest struct via the builder pattern


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


## AppStoreVersionsGetInstance

> AppStoreVersionResponse AppStoreVersionsGetInstance(ctx, id).FieldsAppStoreVersions(fieldsAppStoreVersions).Include(include).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).FieldsIdfaDeclarations(fieldsIdfaDeclarations).FieldsRoutingAppCoverages(fieldsRoutingAppCoverages).FieldsAppStoreVersionPhasedReleases(fieldsAppStoreVersionPhasedReleases).FieldsAgeRatingDeclarations(fieldsAgeRatingDeclarations).FieldsAppStoreReviewDetails(fieldsAppStoreReviewDetails).FieldsBuilds(fieldsBuilds).FieldsAppStoreVersionSubmissions(fieldsAppStoreVersionSubmissions).LimitAppStoreVersionLocalizations(limitAppStoreVersionLocalizations).Execute()



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
    fieldsAppStoreVersions := []string{"FieldsAppStoreVersions_example"} // []string | the fields to include for returned resources of type appStoreVersions (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsAppStoreVersionLocalizations := []string{"FieldsAppStoreVersionLocalizations_example"} // []string | the fields to include for returned resources of type appStoreVersionLocalizations (optional)
    fieldsIdfaDeclarations := []string{"FieldsIdfaDeclarations_example"} // []string | the fields to include for returned resources of type idfaDeclarations (optional)
    fieldsRoutingAppCoverages := []string{"FieldsRoutingAppCoverages_example"} // []string | the fields to include for returned resources of type routingAppCoverages (optional)
    fieldsAppStoreVersionPhasedReleases := []string{"FieldsAppStoreVersionPhasedReleases_example"} // []string | the fields to include for returned resources of type appStoreVersionPhasedReleases (optional)
    fieldsAgeRatingDeclarations := []string{"FieldsAgeRatingDeclarations_example"} // []string | the fields to include for returned resources of type ageRatingDeclarations (optional)
    fieldsAppStoreReviewDetails := []string{"FieldsAppStoreReviewDetails_example"} // []string | the fields to include for returned resources of type appStoreReviewDetails (optional)
    fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)
    fieldsAppStoreVersionSubmissions := []string{"FieldsAppStoreVersionSubmissions_example"} // []string | the fields to include for returned resources of type appStoreVersionSubmissions (optional)
    limitAppStoreVersionLocalizations := int32(56) // int32 | maximum number of related appStoreVersionLocalizations returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppStoreVersionsApi.AppStoreVersionsGetInstance(context.Background(), id).FieldsAppStoreVersions(fieldsAppStoreVersions).Include(include).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).FieldsIdfaDeclarations(fieldsIdfaDeclarations).FieldsRoutingAppCoverages(fieldsRoutingAppCoverages).FieldsAppStoreVersionPhasedReleases(fieldsAppStoreVersionPhasedReleases).FieldsAgeRatingDeclarations(fieldsAgeRatingDeclarations).FieldsAppStoreReviewDetails(fieldsAppStoreReviewDetails).FieldsBuilds(fieldsBuilds).FieldsAppStoreVersionSubmissions(fieldsAppStoreVersionSubmissions).LimitAppStoreVersionLocalizations(limitAppStoreVersionLocalizations).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsApi.AppStoreVersionsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppStoreVersionsGetInstance`: AppStoreVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsApi.AppStoreVersionsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppStoreVersions** | **[]string** | the fields to include for returned resources of type appStoreVersions | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsAppStoreVersionLocalizations** | **[]string** | the fields to include for returned resources of type appStoreVersionLocalizations | 
 **fieldsIdfaDeclarations** | **[]string** | the fields to include for returned resources of type idfaDeclarations | 
 **fieldsRoutingAppCoverages** | **[]string** | the fields to include for returned resources of type routingAppCoverages | 
 **fieldsAppStoreVersionPhasedReleases** | **[]string** | the fields to include for returned resources of type appStoreVersionPhasedReleases | 
 **fieldsAgeRatingDeclarations** | **[]string** | the fields to include for returned resources of type ageRatingDeclarations | 
 **fieldsAppStoreReviewDetails** | **[]string** | the fields to include for returned resources of type appStoreReviewDetails | 
 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 
 **fieldsAppStoreVersionSubmissions** | **[]string** | the fields to include for returned resources of type appStoreVersionSubmissions | 
 **limitAppStoreVersionLocalizations** | **int32** | maximum number of related appStoreVersionLocalizations returned (when they are included) | 

### Return type

[**AppStoreVersionResponse**](AppStoreVersionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsIdfaDeclarationGetToOneRelated

> IdfaDeclarationResponse AppStoreVersionsIdfaDeclarationGetToOneRelated(ctx, id).FieldsIdfaDeclarations(fieldsIdfaDeclarations).Execute()



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
    fieldsIdfaDeclarations := []string{"FieldsIdfaDeclarations_example"} // []string | the fields to include for returned resources of type idfaDeclarations (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppStoreVersionsApi.AppStoreVersionsIdfaDeclarationGetToOneRelated(context.Background(), id).FieldsIdfaDeclarations(fieldsIdfaDeclarations).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsApi.AppStoreVersionsIdfaDeclarationGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppStoreVersionsIdfaDeclarationGetToOneRelated`: IdfaDeclarationResponse
    fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsApi.AppStoreVersionsIdfaDeclarationGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsIdfaDeclarationGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsIdfaDeclarations** | **[]string** | the fields to include for returned resources of type idfaDeclarations | 

### Return type

[**IdfaDeclarationResponse**](IdfaDeclarationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsRoutingAppCoverageGetToOneRelated

> RoutingAppCoverageResponse AppStoreVersionsRoutingAppCoverageGetToOneRelated(ctx, id).FieldsRoutingAppCoverages(fieldsRoutingAppCoverages).Execute()



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
    fieldsRoutingAppCoverages := []string{"FieldsRoutingAppCoverages_example"} // []string | the fields to include for returned resources of type routingAppCoverages (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppStoreVersionsApi.AppStoreVersionsRoutingAppCoverageGetToOneRelated(context.Background(), id).FieldsRoutingAppCoverages(fieldsRoutingAppCoverages).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsApi.AppStoreVersionsRoutingAppCoverageGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppStoreVersionsRoutingAppCoverageGetToOneRelated`: RoutingAppCoverageResponse
    fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsApi.AppStoreVersionsRoutingAppCoverageGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsRoutingAppCoverageGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsRoutingAppCoverages** | **[]string** | the fields to include for returned resources of type routingAppCoverages | 

### Return type

[**RoutingAppCoverageResponse**](RoutingAppCoverageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsUpdateInstance

> AppStoreVersionResponse AppStoreVersionsUpdateInstance(ctx, id).AppStoreVersionUpdateRequest(appStoreVersionUpdateRequest).Execute()



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
    appStoreVersionUpdateRequest := *openapiclient.NewAppStoreVersionUpdateRequest(*openapiclient.NewAppStoreVersionUpdateRequestData("Type_example", "Id_example")) // AppStoreVersionUpdateRequest | AppStoreVersion representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppStoreVersionsApi.AppStoreVersionsUpdateInstance(context.Background(), id).AppStoreVersionUpdateRequest(appStoreVersionUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsApi.AppStoreVersionsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppStoreVersionsUpdateInstance`: AppStoreVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsApi.AppStoreVersionsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appStoreVersionUpdateRequest** | [**AppStoreVersionUpdateRequest**](AppStoreVersionUpdateRequest.md) | AppStoreVersion representation | 

### Return type

[**AppStoreVersionResponse**](AppStoreVersionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

