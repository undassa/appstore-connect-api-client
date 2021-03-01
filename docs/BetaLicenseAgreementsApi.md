# \BetaLicenseAgreementsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BetaLicenseAgreementsAppGetToOneRelated**](BetaLicenseAgreementsApi.md#BetaLicenseAgreementsAppGetToOneRelated) | **Get** /v1/betaLicenseAgreements/{id}/app | 
[**BetaLicenseAgreementsGetCollection**](BetaLicenseAgreementsApi.md#BetaLicenseAgreementsGetCollection) | **Get** /v1/betaLicenseAgreements | 
[**BetaLicenseAgreementsGetInstance**](BetaLicenseAgreementsApi.md#BetaLicenseAgreementsGetInstance) | **Get** /v1/betaLicenseAgreements/{id} | 
[**BetaLicenseAgreementsUpdateInstance**](BetaLicenseAgreementsApi.md#BetaLicenseAgreementsUpdateInstance) | **Patch** /v1/betaLicenseAgreements/{id} | 



## BetaLicenseAgreementsAppGetToOneRelated

> AppResponse BetaLicenseAgreementsAppGetToOneRelated(ctx, id).FieldsApps(fieldsApps).Execute()



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
    resp, r, err := api_client.BetaLicenseAgreementsApi.BetaLicenseAgreementsAppGetToOneRelated(context.Background(), id).FieldsApps(fieldsApps).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaLicenseAgreementsApi.BetaLicenseAgreementsAppGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaLicenseAgreementsAppGetToOneRelated`: AppResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaLicenseAgreementsApi.BetaLicenseAgreementsAppGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaLicenseAgreementsAppGetToOneRelatedRequest struct via the builder pattern


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


## BetaLicenseAgreementsGetCollection

> BetaLicenseAgreementsResponse BetaLicenseAgreementsGetCollection(ctx).FilterApp(filterApp).FieldsBetaLicenseAgreements(fieldsBetaLicenseAgreements).Limit(limit).Include(include).FieldsApps(fieldsApps).Execute()



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
    filterApp := []string{"Inner_example"} // []string | filter by id(s) of related 'app' (optional)
    fieldsBetaLicenseAgreements := []string{"FieldsBetaLicenseAgreements_example"} // []string | the fields to include for returned resources of type betaLicenseAgreements (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaLicenseAgreementsApi.BetaLicenseAgreementsGetCollection(context.Background()).FilterApp(filterApp).FieldsBetaLicenseAgreements(fieldsBetaLicenseAgreements).Limit(limit).Include(include).FieldsApps(fieldsApps).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaLicenseAgreementsApi.BetaLicenseAgreementsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaLicenseAgreementsGetCollection`: BetaLicenseAgreementsResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaLicenseAgreementsApi.BetaLicenseAgreementsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBetaLicenseAgreementsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterApp** | **[]string** | filter by id(s) of related &#39;app&#39; | 
 **fieldsBetaLicenseAgreements** | **[]string** | the fields to include for returned resources of type betaLicenseAgreements | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 

### Return type

[**BetaLicenseAgreementsResponse**](BetaLicenseAgreementsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaLicenseAgreementsGetInstance

> BetaLicenseAgreementResponse BetaLicenseAgreementsGetInstance(ctx, id).FieldsBetaLicenseAgreements(fieldsBetaLicenseAgreements).Include(include).FieldsApps(fieldsApps).Execute()



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
    fieldsBetaLicenseAgreements := []string{"FieldsBetaLicenseAgreements_example"} // []string | the fields to include for returned resources of type betaLicenseAgreements (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaLicenseAgreementsApi.BetaLicenseAgreementsGetInstance(context.Background(), id).FieldsBetaLicenseAgreements(fieldsBetaLicenseAgreements).Include(include).FieldsApps(fieldsApps).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaLicenseAgreementsApi.BetaLicenseAgreementsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaLicenseAgreementsGetInstance`: BetaLicenseAgreementResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaLicenseAgreementsApi.BetaLicenseAgreementsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaLicenseAgreementsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaLicenseAgreements** | **[]string** | the fields to include for returned resources of type betaLicenseAgreements | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 

### Return type

[**BetaLicenseAgreementResponse**](BetaLicenseAgreementResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaLicenseAgreementsUpdateInstance

> BetaLicenseAgreementResponse BetaLicenseAgreementsUpdateInstance(ctx, id).BetaLicenseAgreementUpdateRequest(betaLicenseAgreementUpdateRequest).Execute()



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
    betaLicenseAgreementUpdateRequest := *openapiclient.NewBetaLicenseAgreementUpdateRequest(*openapiclient.NewBetaLicenseAgreementUpdateRequestData("Type_example", "Id_example")) // BetaLicenseAgreementUpdateRequest | BetaLicenseAgreement representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BetaLicenseAgreementsApi.BetaLicenseAgreementsUpdateInstance(context.Background(), id).BetaLicenseAgreementUpdateRequest(betaLicenseAgreementUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaLicenseAgreementsApi.BetaLicenseAgreementsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaLicenseAgreementsUpdateInstance`: BetaLicenseAgreementResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaLicenseAgreementsApi.BetaLicenseAgreementsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaLicenseAgreementsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaLicenseAgreementUpdateRequest** | [**BetaLicenseAgreementUpdateRequest**](BetaLicenseAgreementUpdateRequest.md) | BetaLicenseAgreement representation | 

### Return type

[**BetaLicenseAgreementResponse**](BetaLicenseAgreementResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

