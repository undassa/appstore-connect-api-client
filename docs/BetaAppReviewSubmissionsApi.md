# \BetaAppReviewSubmissionsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BetaAppReviewSubmissionsBuildGetToOneRelated**](BetaAppReviewSubmissionsApi.md#BetaAppReviewSubmissionsBuildGetToOneRelated) | **Get** /v1/betaAppReviewSubmissions/{id}/build | 
[**BetaAppReviewSubmissionsCreateInstance**](BetaAppReviewSubmissionsApi.md#BetaAppReviewSubmissionsCreateInstance) | **Post** /v1/betaAppReviewSubmissions | 
[**BetaAppReviewSubmissionsGetCollection**](BetaAppReviewSubmissionsApi.md#BetaAppReviewSubmissionsGetCollection) | **Get** /v1/betaAppReviewSubmissions | 
[**BetaAppReviewSubmissionsGetInstance**](BetaAppReviewSubmissionsApi.md#BetaAppReviewSubmissionsGetInstance) | **Get** /v1/betaAppReviewSubmissions/{id} | 



## BetaAppReviewSubmissionsBuildGetToOneRelated

> BuildResponse BetaAppReviewSubmissionsBuildGetToOneRelated(ctx, id).FieldsBuilds(fieldsBuilds).Execute()



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
    resp, r, err := api_client.BetaAppReviewSubmissionsApi.BetaAppReviewSubmissionsBuildGetToOneRelated(context.Background(), id).FieldsBuilds(fieldsBuilds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaAppReviewSubmissionsApi.BetaAppReviewSubmissionsBuildGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaAppReviewSubmissionsBuildGetToOneRelated`: BuildResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaAppReviewSubmissionsApi.BetaAppReviewSubmissionsBuildGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaAppReviewSubmissionsBuildGetToOneRelatedRequest struct via the builder pattern


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


## BetaAppReviewSubmissionsCreateInstance

> BetaAppReviewSubmissionResponse BetaAppReviewSubmissionsCreateInstance(ctx).BetaAppReviewSubmissionCreateRequest(betaAppReviewSubmissionCreateRequest).Execute()



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
    betaAppReviewSubmissionCreateRequest := *openapiclient.NewBetaAppReviewSubmissionCreateRequest(*openapiclient.NewBetaAppReviewSubmissionCreateRequestData("Type_example", *openapiclient.NewBetaAppReviewSubmissionCreateRequestDataRelationships(*openapiclient.NewBetaAppReviewSubmissionCreateRequestDataRelationshipsBuild(*openapiclient.NewAppStoreVersionRelationshipsBuildData("Type_example", "Id_example"))))) // BetaAppReviewSubmissionCreateRequest | BetaAppReviewSubmission representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaAppReviewSubmissionsApi.BetaAppReviewSubmissionsCreateInstance(context.Background()).BetaAppReviewSubmissionCreateRequest(betaAppReviewSubmissionCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaAppReviewSubmissionsApi.BetaAppReviewSubmissionsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaAppReviewSubmissionsCreateInstance`: BetaAppReviewSubmissionResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaAppReviewSubmissionsApi.BetaAppReviewSubmissionsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBetaAppReviewSubmissionsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **betaAppReviewSubmissionCreateRequest** | [**BetaAppReviewSubmissionCreateRequest**](BetaAppReviewSubmissionCreateRequest.md) | BetaAppReviewSubmission representation | 

### Return type

[**BetaAppReviewSubmissionResponse**](BetaAppReviewSubmissionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaAppReviewSubmissionsGetCollection

> BetaAppReviewSubmissionsResponse BetaAppReviewSubmissionsGetCollection(ctx).FilterBuild(filterBuild).FilterBetaReviewState(filterBetaReviewState).FieldsBetaAppReviewSubmissions(fieldsBetaAppReviewSubmissions).Limit(limit).Include(include).FieldsBuilds(fieldsBuilds).Execute()



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
    filterBuild := []string{"Inner_example"} // []string | filter by id(s) of related 'build'
    filterBetaReviewState := []string{"FilterBetaReviewState_example"} // []string | filter by attribute 'betaReviewState' (optional)
    fieldsBetaAppReviewSubmissions := []string{"FieldsBetaAppReviewSubmissions_example"} // []string | the fields to include for returned resources of type betaAppReviewSubmissions (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaAppReviewSubmissionsApi.BetaAppReviewSubmissionsGetCollection(context.Background()).FilterBuild(filterBuild).FilterBetaReviewState(filterBetaReviewState).FieldsBetaAppReviewSubmissions(fieldsBetaAppReviewSubmissions).Limit(limit).Include(include).FieldsBuilds(fieldsBuilds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaAppReviewSubmissionsApi.BetaAppReviewSubmissionsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaAppReviewSubmissionsGetCollection`: BetaAppReviewSubmissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaAppReviewSubmissionsApi.BetaAppReviewSubmissionsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBetaAppReviewSubmissionsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterBuild** | **[]string** | filter by id(s) of related &#39;build&#39; | 
 **filterBetaReviewState** | **[]string** | filter by attribute &#39;betaReviewState&#39; | 
 **fieldsBetaAppReviewSubmissions** | **[]string** | the fields to include for returned resources of type betaAppReviewSubmissions | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 

### Return type

[**BetaAppReviewSubmissionsResponse**](BetaAppReviewSubmissionsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaAppReviewSubmissionsGetInstance

> BetaAppReviewSubmissionResponse BetaAppReviewSubmissionsGetInstance(ctx, id).FieldsBetaAppReviewSubmissions(fieldsBetaAppReviewSubmissions).Include(include).FieldsBuilds(fieldsBuilds).Execute()



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
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaAppReviewSubmissionsApi.BetaAppReviewSubmissionsGetInstance(context.Background(), id).FieldsBetaAppReviewSubmissions(fieldsBetaAppReviewSubmissions).Include(include).FieldsBuilds(fieldsBuilds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaAppReviewSubmissionsApi.BetaAppReviewSubmissionsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaAppReviewSubmissionsGetInstance`: BetaAppReviewSubmissionResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaAppReviewSubmissionsApi.BetaAppReviewSubmissionsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaAppReviewSubmissionsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaAppReviewSubmissions** | **[]string** | the fields to include for returned resources of type betaAppReviewSubmissions | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 

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

