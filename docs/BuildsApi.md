# \BuildsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BuildsAppEncryptionDeclarationGetToOneRelated**](BuildsApi.md#BuildsAppEncryptionDeclarationGetToOneRelated) | **Get** /v1/builds/{id}/appEncryptionDeclaration | 
[**BuildsAppEncryptionDeclarationGetToOneRelationship**](BuildsApi.md#BuildsAppEncryptionDeclarationGetToOneRelationship) | **Get** /v1/builds/{id}/relationships/appEncryptionDeclaration | 
[**BuildsAppEncryptionDeclarationUpdateToOneRelationship**](BuildsApi.md#BuildsAppEncryptionDeclarationUpdateToOneRelationship) | **Patch** /v1/builds/{id}/relationships/appEncryptionDeclaration | 
[**BuildsAppGetToOneRelated**](BuildsApi.md#BuildsAppGetToOneRelated) | **Get** /v1/builds/{id}/app | 
[**BuildsAppStoreVersionGetToOneRelated**](BuildsApi.md#BuildsAppStoreVersionGetToOneRelated) | **Get** /v1/builds/{id}/appStoreVersion | 
[**BuildsBetaAppReviewSubmissionGetToOneRelated**](BuildsApi.md#BuildsBetaAppReviewSubmissionGetToOneRelated) | **Get** /v1/builds/{id}/betaAppReviewSubmission | 
[**BuildsBetaBuildLocalizationsGetToManyRelated**](BuildsApi.md#BuildsBetaBuildLocalizationsGetToManyRelated) | **Get** /v1/builds/{id}/betaBuildLocalizations | 
[**BuildsBetaGroupsCreateToManyRelationship**](BuildsApi.md#BuildsBetaGroupsCreateToManyRelationship) | **Post** /v1/builds/{id}/relationships/betaGroups | 
[**BuildsBetaGroupsDeleteToManyRelationship**](BuildsApi.md#BuildsBetaGroupsDeleteToManyRelationship) | **Delete** /v1/builds/{id}/relationships/betaGroups | 
[**BuildsBuildBetaDetailGetToOneRelated**](BuildsApi.md#BuildsBuildBetaDetailGetToOneRelated) | **Get** /v1/builds/{id}/buildBetaDetail | 
[**BuildsDiagnosticSignaturesGetToManyRelated**](BuildsApi.md#BuildsDiagnosticSignaturesGetToManyRelated) | **Get** /v1/builds/{id}/diagnosticSignatures | 
[**BuildsGetCollection**](BuildsApi.md#BuildsGetCollection) | **Get** /v1/builds | 
[**BuildsGetInstance**](BuildsApi.md#BuildsGetInstance) | **Get** /v1/builds/{id} | 
[**BuildsIconsGetToManyRelated**](BuildsApi.md#BuildsIconsGetToManyRelated) | **Get** /v1/builds/{id}/icons | 
[**BuildsIndividualTestersCreateToManyRelationship**](BuildsApi.md#BuildsIndividualTestersCreateToManyRelationship) | **Post** /v1/builds/{id}/relationships/individualTesters | 
[**BuildsIndividualTestersDeleteToManyRelationship**](BuildsApi.md#BuildsIndividualTestersDeleteToManyRelationship) | **Delete** /v1/builds/{id}/relationships/individualTesters | 
[**BuildsIndividualTestersGetToManyRelated**](BuildsApi.md#BuildsIndividualTestersGetToManyRelated) | **Get** /v1/builds/{id}/individualTesters | 
[**BuildsIndividualTestersGetToManyRelationship**](BuildsApi.md#BuildsIndividualTestersGetToManyRelationship) | **Get** /v1/builds/{id}/relationships/individualTesters | 
[**BuildsPerfPowerMetricsGetToManyRelated**](BuildsApi.md#BuildsPerfPowerMetricsGetToManyRelated) | **Get** /v1/builds/{id}/perfPowerMetrics | 
[**BuildsPreReleaseVersionGetToOneRelated**](BuildsApi.md#BuildsPreReleaseVersionGetToOneRelated) | **Get** /v1/builds/{id}/preReleaseVersion | 
[**BuildsUpdateInstance**](BuildsApi.md#BuildsUpdateInstance) | **Patch** /v1/builds/{id} | 



## BuildsAppEncryptionDeclarationGetToOneRelated

> AppEncryptionDeclarationResponse BuildsAppEncryptionDeclarationGetToOneRelated(ctx, id).FieldsAppEncryptionDeclarations(fieldsAppEncryptionDeclarations).Execute()



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
    fieldsAppEncryptionDeclarations := []string{"FieldsAppEncryptionDeclarations_example"} // []string | the fields to include for returned resources of type appEncryptionDeclarations (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BuildsApi.BuildsAppEncryptionDeclarationGetToOneRelated(context.Background(), id).FieldsAppEncryptionDeclarations(fieldsAppEncryptionDeclarations).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsApi.BuildsAppEncryptionDeclarationGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildsAppEncryptionDeclarationGetToOneRelated`: AppEncryptionDeclarationResponse
    fmt.Fprintf(os.Stdout, "Response from `BuildsApi.BuildsAppEncryptionDeclarationGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsAppEncryptionDeclarationGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppEncryptionDeclarations** | **[]string** | the fields to include for returned resources of type appEncryptionDeclarations | 

### Return type

[**AppEncryptionDeclarationResponse**](AppEncryptionDeclarationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsAppEncryptionDeclarationGetToOneRelationship

> BuildAppEncryptionDeclarationLinkageResponse BuildsAppEncryptionDeclarationGetToOneRelationship(ctx, id).Execute()



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
    resp, r, err := api_client.BuildsApi.BuildsAppEncryptionDeclarationGetToOneRelationship(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsApi.BuildsAppEncryptionDeclarationGetToOneRelationship``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildsAppEncryptionDeclarationGetToOneRelationship`: BuildAppEncryptionDeclarationLinkageResponse
    fmt.Fprintf(os.Stdout, "Response from `BuildsApi.BuildsAppEncryptionDeclarationGetToOneRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsAppEncryptionDeclarationGetToOneRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BuildAppEncryptionDeclarationLinkageResponse**](BuildAppEncryptionDeclarationLinkageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsAppEncryptionDeclarationUpdateToOneRelationship

> BuildsAppEncryptionDeclarationUpdateToOneRelationship(ctx, id).BuildAppEncryptionDeclarationLinkageRequest(buildAppEncryptionDeclarationLinkageRequest).Execute()



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
    buildAppEncryptionDeclarationLinkageRequest := *openapiclient.NewBuildAppEncryptionDeclarationLinkageRequest(*openapiclient.NewBuildRelationshipsAppEncryptionDeclarationData("Type_example", "Id_example")) // BuildAppEncryptionDeclarationLinkageRequest | Related linkage

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BuildsApi.BuildsAppEncryptionDeclarationUpdateToOneRelationship(context.Background(), id).BuildAppEncryptionDeclarationLinkageRequest(buildAppEncryptionDeclarationLinkageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsApi.BuildsAppEncryptionDeclarationUpdateToOneRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBuildsAppEncryptionDeclarationUpdateToOneRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **buildAppEncryptionDeclarationLinkageRequest** | [**BuildAppEncryptionDeclarationLinkageRequest**](BuildAppEncryptionDeclarationLinkageRequest.md) | Related linkage | 

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


## BuildsAppGetToOneRelated

> AppResponse BuildsAppGetToOneRelated(ctx, id).FieldsApps(fieldsApps).Execute()



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
    resp, r, err := api_client.BuildsApi.BuildsAppGetToOneRelated(context.Background(), id).FieldsApps(fieldsApps).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsApi.BuildsAppGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildsAppGetToOneRelated`: AppResponse
    fmt.Fprintf(os.Stdout, "Response from `BuildsApi.BuildsAppGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsAppGetToOneRelatedRequest struct via the builder pattern


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


## BuildsAppStoreVersionGetToOneRelated

> AppStoreVersionResponse BuildsAppStoreVersionGetToOneRelated(ctx, id).FieldsAppStoreVersions(fieldsAppStoreVersions).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BuildsApi.BuildsAppStoreVersionGetToOneRelated(context.Background(), id).FieldsAppStoreVersions(fieldsAppStoreVersions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsApi.BuildsAppStoreVersionGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildsAppStoreVersionGetToOneRelated`: AppStoreVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `BuildsApi.BuildsAppStoreVersionGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsAppStoreVersionGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppStoreVersions** | **[]string** | the fields to include for returned resources of type appStoreVersions | 

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


## BuildsBetaAppReviewSubmissionGetToOneRelated

> BetaAppReviewSubmissionResponse BuildsBetaAppReviewSubmissionGetToOneRelated(ctx, id).FieldsBetaAppReviewSubmissions(fieldsBetaAppReviewSubmissions).Execute()



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
    fieldsBetaAppReviewSubmissions := []string{"FieldsBetaAppReviewSubmissions_example"} // []string | the fields to include for returned resources of type betaAppReviewSubmissions (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BuildsApi.BuildsBetaAppReviewSubmissionGetToOneRelated(context.Background(), id).FieldsBetaAppReviewSubmissions(fieldsBetaAppReviewSubmissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsApi.BuildsBetaAppReviewSubmissionGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildsBetaAppReviewSubmissionGetToOneRelated`: BetaAppReviewSubmissionResponse
    fmt.Fprintf(os.Stdout, "Response from `BuildsApi.BuildsBetaAppReviewSubmissionGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsBetaAppReviewSubmissionGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaAppReviewSubmissions** | **[]string** | the fields to include for returned resources of type betaAppReviewSubmissions | 

### Return type

[**BetaAppReviewSubmissionResponse**](BetaAppReviewSubmissionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsBetaBuildLocalizationsGetToManyRelated

> BetaBuildLocalizationsResponse BuildsBetaBuildLocalizationsGetToManyRelated(ctx, id).FieldsBetaBuildLocalizations(fieldsBetaBuildLocalizations).Limit(limit).Execute()



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
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BuildsApi.BuildsBetaBuildLocalizationsGetToManyRelated(context.Background(), id).FieldsBetaBuildLocalizations(fieldsBetaBuildLocalizations).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsApi.BuildsBetaBuildLocalizationsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildsBetaBuildLocalizationsGetToManyRelated`: BetaBuildLocalizationsResponse
    fmt.Fprintf(os.Stdout, "Response from `BuildsApi.BuildsBetaBuildLocalizationsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsBetaBuildLocalizationsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaBuildLocalizations** | **[]string** | the fields to include for returned resources of type betaBuildLocalizations | 
 **limit** | **int32** | maximum resources per page | 

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


## BuildsBetaGroupsCreateToManyRelationship

> BuildsBetaGroupsCreateToManyRelationship(ctx, id).BuildBetaGroupsLinkagesRequest(buildBetaGroupsLinkagesRequest).Execute()



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
    buildBetaGroupsLinkagesRequest := *openapiclient.NewBuildBetaGroupsLinkagesRequest([]openapiclient.AppRelationshipsBetaGroupsData{*openapiclient.NewAppRelationshipsBetaGroupsData("Type_example", "Id_example")}) // BuildBetaGroupsLinkagesRequest | List of related linkages

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BuildsApi.BuildsBetaGroupsCreateToManyRelationship(context.Background(), id).BuildBetaGroupsLinkagesRequest(buildBetaGroupsLinkagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsApi.BuildsBetaGroupsCreateToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBuildsBetaGroupsCreateToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **buildBetaGroupsLinkagesRequest** | [**BuildBetaGroupsLinkagesRequest**](BuildBetaGroupsLinkagesRequest.md) | List of related linkages | 

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


## BuildsBetaGroupsDeleteToManyRelationship

> BuildsBetaGroupsDeleteToManyRelationship(ctx, id).BuildBetaGroupsLinkagesRequest(buildBetaGroupsLinkagesRequest).Execute()



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
    buildBetaGroupsLinkagesRequest := *openapiclient.NewBuildBetaGroupsLinkagesRequest([]openapiclient.AppRelationshipsBetaGroupsData{*openapiclient.NewAppRelationshipsBetaGroupsData("Type_example", "Id_example")}) // BuildBetaGroupsLinkagesRequest | List of related linkages

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BuildsApi.BuildsBetaGroupsDeleteToManyRelationship(context.Background(), id).BuildBetaGroupsLinkagesRequest(buildBetaGroupsLinkagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsApi.BuildsBetaGroupsDeleteToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBuildsBetaGroupsDeleteToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **buildBetaGroupsLinkagesRequest** | [**BuildBetaGroupsLinkagesRequest**](BuildBetaGroupsLinkagesRequest.md) | List of related linkages | 

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


## BuildsBuildBetaDetailGetToOneRelated

> BuildBetaDetailResponse BuildsBuildBetaDetailGetToOneRelated(ctx, id).FieldsBuildBetaDetails(fieldsBuildBetaDetails).Execute()



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
    fieldsBuildBetaDetails := []string{"FieldsBuildBetaDetails_example"} // []string | the fields to include for returned resources of type buildBetaDetails (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BuildsApi.BuildsBuildBetaDetailGetToOneRelated(context.Background(), id).FieldsBuildBetaDetails(fieldsBuildBetaDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsApi.BuildsBuildBetaDetailGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildsBuildBetaDetailGetToOneRelated`: BuildBetaDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `BuildsApi.BuildsBuildBetaDetailGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsBuildBetaDetailGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBuildBetaDetails** | **[]string** | the fields to include for returned resources of type buildBetaDetails | 

### Return type

[**BuildBetaDetailResponse**](BuildBetaDetailResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsDiagnosticSignaturesGetToManyRelated

> DiagnosticSignaturesResponse BuildsDiagnosticSignaturesGetToManyRelated(ctx, id).FilterDiagnosticType(filterDiagnosticType).FieldsDiagnosticSignatures(fieldsDiagnosticSignatures).Limit(limit).Execute()



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
    filterDiagnosticType := []string{"FilterDiagnosticType_example"} // []string | filter by attribute 'diagnosticType' (optional)
    fieldsDiagnosticSignatures := []string{"FieldsDiagnosticSignatures_example"} // []string | the fields to include for returned resources of type diagnosticSignatures (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BuildsApi.BuildsDiagnosticSignaturesGetToManyRelated(context.Background(), id).FilterDiagnosticType(filterDiagnosticType).FieldsDiagnosticSignatures(fieldsDiagnosticSignatures).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsApi.BuildsDiagnosticSignaturesGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildsDiagnosticSignaturesGetToManyRelated`: DiagnosticSignaturesResponse
    fmt.Fprintf(os.Stdout, "Response from `BuildsApi.BuildsDiagnosticSignaturesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsDiagnosticSignaturesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterDiagnosticType** | **[]string** | filter by attribute &#39;diagnosticType&#39; | 
 **fieldsDiagnosticSignatures** | **[]string** | the fields to include for returned resources of type diagnosticSignatures | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**DiagnosticSignaturesResponse**](DiagnosticSignaturesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsGetCollection

> BuildsResponse BuildsGetCollection(ctx).FilterBetaAppReviewSubmissionBetaReviewState(filterBetaAppReviewSubmissionBetaReviewState).FilterExpired(filterExpired).FilterPreReleaseVersionPlatform(filterPreReleaseVersionPlatform).FilterPreReleaseVersionVersion(filterPreReleaseVersionVersion).FilterProcessingState(filterProcessingState).FilterUsesNonExemptEncryption(filterUsesNonExemptEncryption).FilterVersion(filterVersion).FilterApp(filterApp).FilterAppStoreVersion(filterAppStoreVersion).FilterBetaGroups(filterBetaGroups).FilterPreReleaseVersion(filterPreReleaseVersion).FilterId(filterId).Sort(sort).FieldsBuilds(fieldsBuilds).Limit(limit).Include(include).FieldsAppEncryptionDeclarations(fieldsAppEncryptionDeclarations).FieldsBetaAppReviewSubmissions(fieldsBetaAppReviewSubmissions).FieldsBuildBetaDetails(fieldsBuildBetaDetails).FieldsBuildIcons(fieldsBuildIcons).FieldsPerfPowerMetrics(fieldsPerfPowerMetrics).FieldsPreReleaseVersions(fieldsPreReleaseVersions).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsDiagnosticSignatures(fieldsDiagnosticSignatures).FieldsBetaTesters(fieldsBetaTesters).FieldsBetaBuildLocalizations(fieldsBetaBuildLocalizations).FieldsApps(fieldsApps).LimitBetaBuildLocalizations(limitBetaBuildLocalizations).LimitIcons(limitIcons).LimitIndividualTesters(limitIndividualTesters).Execute()



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
    filterBetaAppReviewSubmissionBetaReviewState := []string{"FilterBetaAppReviewSubmissionBetaReviewState_example"} // []string | filter by attribute 'betaAppReviewSubmission.betaReviewState' (optional)
    filterExpired := []string{"Inner_example"} // []string | filter by attribute 'expired' (optional)
    filterPreReleaseVersionPlatform := []string{"FilterPreReleaseVersionPlatform_example"} // []string | filter by attribute 'preReleaseVersion.platform' (optional)
    filterPreReleaseVersionVersion := []string{"Inner_example"} // []string | filter by attribute 'preReleaseVersion.version' (optional)
    filterProcessingState := []string{"FilterProcessingState_example"} // []string | filter by attribute 'processingState' (optional)
    filterUsesNonExemptEncryption := []string{"Inner_example"} // []string | filter by attribute 'usesNonExemptEncryption' (optional)
    filterVersion := []string{"Inner_example"} // []string | filter by attribute 'version' (optional)
    filterApp := []string{"Inner_example"} // []string | filter by id(s) of related 'app' (optional)
    filterAppStoreVersion := []string{"Inner_example"} // []string | filter by id(s) of related 'appStoreVersion' (optional)
    filterBetaGroups := []string{"Inner_example"} // []string | filter by id(s) of related 'betaGroups' (optional)
    filterPreReleaseVersion := []string{"Inner_example"} // []string | filter by id(s) of related 'preReleaseVersion' (optional)
    filterId := []string{"Inner_example"} // []string | filter by id(s) (optional)
    sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; resources will be sorted as specified (optional)
    fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsAppEncryptionDeclarations := []string{"FieldsAppEncryptionDeclarations_example"} // []string | the fields to include for returned resources of type appEncryptionDeclarations (optional)
    fieldsBetaAppReviewSubmissions := []string{"FieldsBetaAppReviewSubmissions_example"} // []string | the fields to include for returned resources of type betaAppReviewSubmissions (optional)
    fieldsBuildBetaDetails := []string{"FieldsBuildBetaDetails_example"} // []string | the fields to include for returned resources of type buildBetaDetails (optional)
    fieldsBuildIcons := []string{"FieldsBuildIcons_example"} // []string | the fields to include for returned resources of type buildIcons (optional)
    fieldsPerfPowerMetrics := []string{"FieldsPerfPowerMetrics_example"} // []string | the fields to include for returned resources of type perfPowerMetrics (optional)
    fieldsPreReleaseVersions := []string{"FieldsPreReleaseVersions_example"} // []string | the fields to include for returned resources of type preReleaseVersions (optional)
    fieldsAppStoreVersions := []string{"FieldsAppStoreVersions_example"} // []string | the fields to include for returned resources of type appStoreVersions (optional)
    fieldsDiagnosticSignatures := []string{"FieldsDiagnosticSignatures_example"} // []string | the fields to include for returned resources of type diagnosticSignatures (optional)
    fieldsBetaTesters := []string{"FieldsBetaTesters_example"} // []string | the fields to include for returned resources of type betaTesters (optional)
    fieldsBetaBuildLocalizations := []string{"FieldsBetaBuildLocalizations_example"} // []string | the fields to include for returned resources of type betaBuildLocalizations (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
    limitBetaBuildLocalizations := int32(56) // int32 | maximum number of related betaBuildLocalizations returned (when they are included) (optional)
    limitIcons := int32(56) // int32 | maximum number of related icons returned (when they are included) (optional)
    limitIndividualTesters := int32(56) // int32 | maximum number of related individualTesters returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BuildsApi.BuildsGetCollection(context.Background()).FilterBetaAppReviewSubmissionBetaReviewState(filterBetaAppReviewSubmissionBetaReviewState).FilterExpired(filterExpired).FilterPreReleaseVersionPlatform(filterPreReleaseVersionPlatform).FilterPreReleaseVersionVersion(filterPreReleaseVersionVersion).FilterProcessingState(filterProcessingState).FilterUsesNonExemptEncryption(filterUsesNonExemptEncryption).FilterVersion(filterVersion).FilterApp(filterApp).FilterAppStoreVersion(filterAppStoreVersion).FilterBetaGroups(filterBetaGroups).FilterPreReleaseVersion(filterPreReleaseVersion).FilterId(filterId).Sort(sort).FieldsBuilds(fieldsBuilds).Limit(limit).Include(include).FieldsAppEncryptionDeclarations(fieldsAppEncryptionDeclarations).FieldsBetaAppReviewSubmissions(fieldsBetaAppReviewSubmissions).FieldsBuildBetaDetails(fieldsBuildBetaDetails).FieldsBuildIcons(fieldsBuildIcons).FieldsPerfPowerMetrics(fieldsPerfPowerMetrics).FieldsPreReleaseVersions(fieldsPreReleaseVersions).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsDiagnosticSignatures(fieldsDiagnosticSignatures).FieldsBetaTesters(fieldsBetaTesters).FieldsBetaBuildLocalizations(fieldsBetaBuildLocalizations).FieldsApps(fieldsApps).LimitBetaBuildLocalizations(limitBetaBuildLocalizations).LimitIcons(limitIcons).LimitIndividualTesters(limitIndividualTesters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsApi.BuildsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildsGetCollection`: BuildsResponse
    fmt.Fprintf(os.Stdout, "Response from `BuildsApi.BuildsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBuildsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterBetaAppReviewSubmissionBetaReviewState** | **[]string** | filter by attribute &#39;betaAppReviewSubmission.betaReviewState&#39; | 
 **filterExpired** | **[]string** | filter by attribute &#39;expired&#39; | 
 **filterPreReleaseVersionPlatform** | **[]string** | filter by attribute &#39;preReleaseVersion.platform&#39; | 
 **filterPreReleaseVersionVersion** | **[]string** | filter by attribute &#39;preReleaseVersion.version&#39; | 
 **filterProcessingState** | **[]string** | filter by attribute &#39;processingState&#39; | 
 **filterUsesNonExemptEncryption** | **[]string** | filter by attribute &#39;usesNonExemptEncryption&#39; | 
 **filterVersion** | **[]string** | filter by attribute &#39;version&#39; | 
 **filterApp** | **[]string** | filter by id(s) of related &#39;app&#39; | 
 **filterAppStoreVersion** | **[]string** | filter by id(s) of related &#39;appStoreVersion&#39; | 
 **filterBetaGroups** | **[]string** | filter by id(s) of related &#39;betaGroups&#39; | 
 **filterPreReleaseVersion** | **[]string** | filter by id(s) of related &#39;preReleaseVersion&#39; | 
 **filterId** | **[]string** | filter by id(s) | 
 **sort** | **[]string** | comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsAppEncryptionDeclarations** | **[]string** | the fields to include for returned resources of type appEncryptionDeclarations | 
 **fieldsBetaAppReviewSubmissions** | **[]string** | the fields to include for returned resources of type betaAppReviewSubmissions | 
 **fieldsBuildBetaDetails** | **[]string** | the fields to include for returned resources of type buildBetaDetails | 
 **fieldsBuildIcons** | **[]string** | the fields to include for returned resources of type buildIcons | 
 **fieldsPerfPowerMetrics** | **[]string** | the fields to include for returned resources of type perfPowerMetrics | 
 **fieldsPreReleaseVersions** | **[]string** | the fields to include for returned resources of type preReleaseVersions | 
 **fieldsAppStoreVersions** | **[]string** | the fields to include for returned resources of type appStoreVersions | 
 **fieldsDiagnosticSignatures** | **[]string** | the fields to include for returned resources of type diagnosticSignatures | 
 **fieldsBetaTesters** | **[]string** | the fields to include for returned resources of type betaTesters | 
 **fieldsBetaBuildLocalizations** | **[]string** | the fields to include for returned resources of type betaBuildLocalizations | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **limitBetaBuildLocalizations** | **int32** | maximum number of related betaBuildLocalizations returned (when they are included) | 
 **limitIcons** | **int32** | maximum number of related icons returned (when they are included) | 
 **limitIndividualTesters** | **int32** | maximum number of related individualTesters returned (when they are included) | 

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


## BuildsGetInstance

> BuildResponse BuildsGetInstance(ctx, id).FieldsBuilds(fieldsBuilds).Include(include).FieldsAppEncryptionDeclarations(fieldsAppEncryptionDeclarations).FieldsBetaAppReviewSubmissions(fieldsBetaAppReviewSubmissions).FieldsBuildBetaDetails(fieldsBuildBetaDetails).FieldsBuildIcons(fieldsBuildIcons).FieldsPerfPowerMetrics(fieldsPerfPowerMetrics).FieldsPreReleaseVersions(fieldsPreReleaseVersions).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsDiagnosticSignatures(fieldsDiagnosticSignatures).FieldsBetaTesters(fieldsBetaTesters).FieldsBetaBuildLocalizations(fieldsBetaBuildLocalizations).FieldsApps(fieldsApps).LimitBetaBuildLocalizations(limitBetaBuildLocalizations).LimitIcons(limitIcons).LimitIndividualTesters(limitIndividualTesters).Execute()



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
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsAppEncryptionDeclarations := []string{"FieldsAppEncryptionDeclarations_example"} // []string | the fields to include for returned resources of type appEncryptionDeclarations (optional)
    fieldsBetaAppReviewSubmissions := []string{"FieldsBetaAppReviewSubmissions_example"} // []string | the fields to include for returned resources of type betaAppReviewSubmissions (optional)
    fieldsBuildBetaDetails := []string{"FieldsBuildBetaDetails_example"} // []string | the fields to include for returned resources of type buildBetaDetails (optional)
    fieldsBuildIcons := []string{"FieldsBuildIcons_example"} // []string | the fields to include for returned resources of type buildIcons (optional)
    fieldsPerfPowerMetrics := []string{"FieldsPerfPowerMetrics_example"} // []string | the fields to include for returned resources of type perfPowerMetrics (optional)
    fieldsPreReleaseVersions := []string{"FieldsPreReleaseVersions_example"} // []string | the fields to include for returned resources of type preReleaseVersions (optional)
    fieldsAppStoreVersions := []string{"FieldsAppStoreVersions_example"} // []string | the fields to include for returned resources of type appStoreVersions (optional)
    fieldsDiagnosticSignatures := []string{"FieldsDiagnosticSignatures_example"} // []string | the fields to include for returned resources of type diagnosticSignatures (optional)
    fieldsBetaTesters := []string{"FieldsBetaTesters_example"} // []string | the fields to include for returned resources of type betaTesters (optional)
    fieldsBetaBuildLocalizations := []string{"FieldsBetaBuildLocalizations_example"} // []string | the fields to include for returned resources of type betaBuildLocalizations (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
    limitBetaBuildLocalizations := int32(56) // int32 | maximum number of related betaBuildLocalizations returned (when they are included) (optional)
    limitIcons := int32(56) // int32 | maximum number of related icons returned (when they are included) (optional)
    limitIndividualTesters := int32(56) // int32 | maximum number of related individualTesters returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BuildsApi.BuildsGetInstance(context.Background(), id).FieldsBuilds(fieldsBuilds).Include(include).FieldsAppEncryptionDeclarations(fieldsAppEncryptionDeclarations).FieldsBetaAppReviewSubmissions(fieldsBetaAppReviewSubmissions).FieldsBuildBetaDetails(fieldsBuildBetaDetails).FieldsBuildIcons(fieldsBuildIcons).FieldsPerfPowerMetrics(fieldsPerfPowerMetrics).FieldsPreReleaseVersions(fieldsPreReleaseVersions).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsDiagnosticSignatures(fieldsDiagnosticSignatures).FieldsBetaTesters(fieldsBetaTesters).FieldsBetaBuildLocalizations(fieldsBetaBuildLocalizations).FieldsApps(fieldsApps).LimitBetaBuildLocalizations(limitBetaBuildLocalizations).LimitIcons(limitIcons).LimitIndividualTesters(limitIndividualTesters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsApi.BuildsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildsGetInstance`: BuildResponse
    fmt.Fprintf(os.Stdout, "Response from `BuildsApi.BuildsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsAppEncryptionDeclarations** | **[]string** | the fields to include for returned resources of type appEncryptionDeclarations | 
 **fieldsBetaAppReviewSubmissions** | **[]string** | the fields to include for returned resources of type betaAppReviewSubmissions | 
 **fieldsBuildBetaDetails** | **[]string** | the fields to include for returned resources of type buildBetaDetails | 
 **fieldsBuildIcons** | **[]string** | the fields to include for returned resources of type buildIcons | 
 **fieldsPerfPowerMetrics** | **[]string** | the fields to include for returned resources of type perfPowerMetrics | 
 **fieldsPreReleaseVersions** | **[]string** | the fields to include for returned resources of type preReleaseVersions | 
 **fieldsAppStoreVersions** | **[]string** | the fields to include for returned resources of type appStoreVersions | 
 **fieldsDiagnosticSignatures** | **[]string** | the fields to include for returned resources of type diagnosticSignatures | 
 **fieldsBetaTesters** | **[]string** | the fields to include for returned resources of type betaTesters | 
 **fieldsBetaBuildLocalizations** | **[]string** | the fields to include for returned resources of type betaBuildLocalizations | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **limitBetaBuildLocalizations** | **int32** | maximum number of related betaBuildLocalizations returned (when they are included) | 
 **limitIcons** | **int32** | maximum number of related icons returned (when they are included) | 
 **limitIndividualTesters** | **int32** | maximum number of related individualTesters returned (when they are included) | 

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


## BuildsIconsGetToManyRelated

> BuildIconsResponse BuildsIconsGetToManyRelated(ctx, id).FieldsBuildIcons(fieldsBuildIcons).Limit(limit).Execute()



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
    fieldsBuildIcons := []string{"FieldsBuildIcons_example"} // []string | the fields to include for returned resources of type buildIcons (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BuildsApi.BuildsIconsGetToManyRelated(context.Background(), id).FieldsBuildIcons(fieldsBuildIcons).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsApi.BuildsIconsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildsIconsGetToManyRelated`: BuildIconsResponse
    fmt.Fprintf(os.Stdout, "Response from `BuildsApi.BuildsIconsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsIconsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBuildIcons** | **[]string** | the fields to include for returned resources of type buildIcons | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**BuildIconsResponse**](BuildIconsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsIndividualTestersCreateToManyRelationship

> BuildsIndividualTestersCreateToManyRelationship(ctx, id).BuildIndividualTestersLinkagesRequest(buildIndividualTestersLinkagesRequest).Execute()



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
    buildIndividualTestersLinkagesRequest := *openapiclient.NewBuildIndividualTestersLinkagesRequest([]openapiclient.BetaGroupRelationshipsBetaTestersData{*openapiclient.NewBetaGroupRelationshipsBetaTestersData("Type_example", "Id_example")}) // BuildIndividualTestersLinkagesRequest | List of related linkages

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BuildsApi.BuildsIndividualTestersCreateToManyRelationship(context.Background(), id).BuildIndividualTestersLinkagesRequest(buildIndividualTestersLinkagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsApi.BuildsIndividualTestersCreateToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBuildsIndividualTestersCreateToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **buildIndividualTestersLinkagesRequest** | [**BuildIndividualTestersLinkagesRequest**](BuildIndividualTestersLinkagesRequest.md) | List of related linkages | 

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


## BuildsIndividualTestersDeleteToManyRelationship

> BuildsIndividualTestersDeleteToManyRelationship(ctx, id).BuildIndividualTestersLinkagesRequest(buildIndividualTestersLinkagesRequest).Execute()



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
    buildIndividualTestersLinkagesRequest := *openapiclient.NewBuildIndividualTestersLinkagesRequest([]openapiclient.BetaGroupRelationshipsBetaTestersData{*openapiclient.NewBetaGroupRelationshipsBetaTestersData("Type_example", "Id_example")}) // BuildIndividualTestersLinkagesRequest | List of related linkages

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BuildsApi.BuildsIndividualTestersDeleteToManyRelationship(context.Background(), id).BuildIndividualTestersLinkagesRequest(buildIndividualTestersLinkagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsApi.BuildsIndividualTestersDeleteToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBuildsIndividualTestersDeleteToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **buildIndividualTestersLinkagesRequest** | [**BuildIndividualTestersLinkagesRequest**](BuildIndividualTestersLinkagesRequest.md) | List of related linkages | 

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


## BuildsIndividualTestersGetToManyRelated

> BetaTestersResponse BuildsIndividualTestersGetToManyRelated(ctx, id).FieldsBetaTesters(fieldsBetaTesters).Limit(limit).Execute()



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
    resp, r, err := api_client.BuildsApi.BuildsIndividualTestersGetToManyRelated(context.Background(), id).FieldsBetaTesters(fieldsBetaTesters).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsApi.BuildsIndividualTestersGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildsIndividualTestersGetToManyRelated`: BetaTestersResponse
    fmt.Fprintf(os.Stdout, "Response from `BuildsApi.BuildsIndividualTestersGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsIndividualTestersGetToManyRelatedRequest struct via the builder pattern


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


## BuildsIndividualTestersGetToManyRelationship

> BuildIndividualTestersLinkagesResponse BuildsIndividualTestersGetToManyRelationship(ctx, id).Limit(limit).Execute()



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
    resp, r, err := api_client.BuildsApi.BuildsIndividualTestersGetToManyRelationship(context.Background(), id).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsApi.BuildsIndividualTestersGetToManyRelationship``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildsIndividualTestersGetToManyRelationship`: BuildIndividualTestersLinkagesResponse
    fmt.Fprintf(os.Stdout, "Response from `BuildsApi.BuildsIndividualTestersGetToManyRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsIndividualTestersGetToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum resources per page | 

### Return type

[**BuildIndividualTestersLinkagesResponse**](BuildIndividualTestersLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsPerfPowerMetricsGetToManyRelated

> PerfPowerMetricsResponse BuildsPerfPowerMetricsGetToManyRelated(ctx, id).FilterDeviceType(filterDeviceType).FilterMetricType(filterMetricType).FilterPlatform(filterPlatform).Execute()



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
    filterDeviceType := []string{"Inner_example"} // []string | filter by attribute 'deviceType' (optional)
    filterMetricType := []string{"FilterMetricType_example"} // []string | filter by attribute 'metricType' (optional)
    filterPlatform := []string{"FilterPlatform_example"} // []string | filter by attribute 'platform' (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BuildsApi.BuildsPerfPowerMetricsGetToManyRelated(context.Background(), id).FilterDeviceType(filterDeviceType).FilterMetricType(filterMetricType).FilterPlatform(filterPlatform).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsApi.BuildsPerfPowerMetricsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildsPerfPowerMetricsGetToManyRelated`: PerfPowerMetricsResponse
    fmt.Fprintf(os.Stdout, "Response from `BuildsApi.BuildsPerfPowerMetricsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsPerfPowerMetricsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterDeviceType** | **[]string** | filter by attribute &#39;deviceType&#39; | 
 **filterMetricType** | **[]string** | filter by attribute &#39;metricType&#39; | 
 **filterPlatform** | **[]string** | filter by attribute &#39;platform&#39; | 

### Return type

[**PerfPowerMetricsResponse**](PerfPowerMetricsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsPreReleaseVersionGetToOneRelated

> PrereleaseVersionResponse BuildsPreReleaseVersionGetToOneRelated(ctx, id).FieldsPreReleaseVersions(fieldsPreReleaseVersions).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BuildsApi.BuildsPreReleaseVersionGetToOneRelated(context.Background(), id).FieldsPreReleaseVersions(fieldsPreReleaseVersions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsApi.BuildsPreReleaseVersionGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildsPreReleaseVersionGetToOneRelated`: PrereleaseVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `BuildsApi.BuildsPreReleaseVersionGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsPreReleaseVersionGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsPreReleaseVersions** | **[]string** | the fields to include for returned resources of type preReleaseVersions | 

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


## BuildsUpdateInstance

> BuildResponse BuildsUpdateInstance(ctx, id).BuildUpdateRequest(buildUpdateRequest).Execute()



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
    buildUpdateRequest := *openapiclient.NewBuildUpdateRequest(*openapiclient.NewBuildUpdateRequestData("Type_example", "Id_example")) // BuildUpdateRequest | Build representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BuildsApi.BuildsUpdateInstance(context.Background(), id).BuildUpdateRequest(buildUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsApi.BuildsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildsUpdateInstance`: BuildResponse
    fmt.Fprintf(os.Stdout, "Response from `BuildsApi.BuildsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **buildUpdateRequest** | [**BuildUpdateRequest**](BuildUpdateRequest.md) | Build representation | 

### Return type

[**BuildResponse**](BuildResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

