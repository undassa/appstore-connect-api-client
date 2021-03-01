# \EndUserLicenseAgreementsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EndUserLicenseAgreementsCreateInstance**](EndUserLicenseAgreementsApi.md#EndUserLicenseAgreementsCreateInstance) | **Post** /v1/endUserLicenseAgreements | 
[**EndUserLicenseAgreementsDeleteInstance**](EndUserLicenseAgreementsApi.md#EndUserLicenseAgreementsDeleteInstance) | **Delete** /v1/endUserLicenseAgreements/{id} | 
[**EndUserLicenseAgreementsGetInstance**](EndUserLicenseAgreementsApi.md#EndUserLicenseAgreementsGetInstance) | **Get** /v1/endUserLicenseAgreements/{id} | 
[**EndUserLicenseAgreementsTerritoriesGetToManyRelated**](EndUserLicenseAgreementsApi.md#EndUserLicenseAgreementsTerritoriesGetToManyRelated) | **Get** /v1/endUserLicenseAgreements/{id}/territories | 
[**EndUserLicenseAgreementsUpdateInstance**](EndUserLicenseAgreementsApi.md#EndUserLicenseAgreementsUpdateInstance) | **Patch** /v1/endUserLicenseAgreements/{id} | 



## EndUserLicenseAgreementsCreateInstance

> EndUserLicenseAgreementResponse EndUserLicenseAgreementsCreateInstance(ctx).EndUserLicenseAgreementCreateRequest(endUserLicenseAgreementCreateRequest).Execute()



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
    endUserLicenseAgreementCreateRequest := *openapiclient.NewEndUserLicenseAgreementCreateRequest(*openapiclient.NewEndUserLicenseAgreementCreateRequestData("Type_example", *openapiclient.NewEndUserLicenseAgreementCreateRequestDataAttributes("AgreementText_example"), *openapiclient.NewEndUserLicenseAgreementCreateRequestDataRelationships(*openapiclient.NewAppPreOrderCreateRequestDataRelationshipsApp(*openapiclient.NewAppEncryptionDeclarationRelationshipsAppData("Type_example", "Id_example")), *openapiclient.NewEndUserLicenseAgreementCreateRequestDataRelationshipsTerritories([]openapiclient.AppPricePointRelationshipsTerritoryData{*openapiclient.NewAppPricePointRelationshipsTerritoryData("Type_example", "Id_example")})))) // EndUserLicenseAgreementCreateRequest | EndUserLicenseAgreement representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EndUserLicenseAgreementsApi.EndUserLicenseAgreementsCreateInstance(context.Background()).EndUserLicenseAgreementCreateRequest(endUserLicenseAgreementCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndUserLicenseAgreementsApi.EndUserLicenseAgreementsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EndUserLicenseAgreementsCreateInstance`: EndUserLicenseAgreementResponse
    fmt.Fprintf(os.Stdout, "Response from `EndUserLicenseAgreementsApi.EndUserLicenseAgreementsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEndUserLicenseAgreementsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endUserLicenseAgreementCreateRequest** | [**EndUserLicenseAgreementCreateRequest**](EndUserLicenseAgreementCreateRequest.md) | EndUserLicenseAgreement representation | 

### Return type

[**EndUserLicenseAgreementResponse**](EndUserLicenseAgreementResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndUserLicenseAgreementsDeleteInstance

> EndUserLicenseAgreementsDeleteInstance(ctx, id).Execute()



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
    resp, r, err := api_client.EndUserLicenseAgreementsApi.EndUserLicenseAgreementsDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndUserLicenseAgreementsApi.EndUserLicenseAgreementsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEndUserLicenseAgreementsDeleteInstanceRequest struct via the builder pattern


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


## EndUserLicenseAgreementsGetInstance

> EndUserLicenseAgreementResponse EndUserLicenseAgreementsGetInstance(ctx, id).FieldsEndUserLicenseAgreements(fieldsEndUserLicenseAgreements).Include(include).FieldsTerritories(fieldsTerritories).LimitTerritories(limitTerritories).Execute()



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
    fieldsEndUserLicenseAgreements := []string{"FieldsEndUserLicenseAgreements_example"} // []string | the fields to include for returned resources of type endUserLicenseAgreements (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsTerritories := []string{"FieldsTerritories_example"} // []string | the fields to include for returned resources of type territories (optional)
    limitTerritories := int32(56) // int32 | maximum number of related territories returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EndUserLicenseAgreementsApi.EndUserLicenseAgreementsGetInstance(context.Background(), id).FieldsEndUserLicenseAgreements(fieldsEndUserLicenseAgreements).Include(include).FieldsTerritories(fieldsTerritories).LimitTerritories(limitTerritories).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndUserLicenseAgreementsApi.EndUserLicenseAgreementsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EndUserLicenseAgreementsGetInstance`: EndUserLicenseAgreementResponse
    fmt.Fprintf(os.Stdout, "Response from `EndUserLicenseAgreementsApi.EndUserLicenseAgreementsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiEndUserLicenseAgreementsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsEndUserLicenseAgreements** | **[]string** | the fields to include for returned resources of type endUserLicenseAgreements | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsTerritories** | **[]string** | the fields to include for returned resources of type territories | 
 **limitTerritories** | **int32** | maximum number of related territories returned (when they are included) | 

### Return type

[**EndUserLicenseAgreementResponse**](EndUserLicenseAgreementResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndUserLicenseAgreementsTerritoriesGetToManyRelated

> TerritoriesResponse EndUserLicenseAgreementsTerritoriesGetToManyRelated(ctx, id).FieldsTerritories(fieldsTerritories).Limit(limit).Execute()



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
    fieldsTerritories := []string{"FieldsTerritories_example"} // []string | the fields to include for returned resources of type territories (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EndUserLicenseAgreementsApi.EndUserLicenseAgreementsTerritoriesGetToManyRelated(context.Background(), id).FieldsTerritories(fieldsTerritories).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndUserLicenseAgreementsApi.EndUserLicenseAgreementsTerritoriesGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EndUserLicenseAgreementsTerritoriesGetToManyRelated`: TerritoriesResponse
    fmt.Fprintf(os.Stdout, "Response from `EndUserLicenseAgreementsApi.EndUserLicenseAgreementsTerritoriesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiEndUserLicenseAgreementsTerritoriesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsTerritories** | **[]string** | the fields to include for returned resources of type territories | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**TerritoriesResponse**](TerritoriesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndUserLicenseAgreementsUpdateInstance

> EndUserLicenseAgreementResponse EndUserLicenseAgreementsUpdateInstance(ctx, id).EndUserLicenseAgreementUpdateRequest(endUserLicenseAgreementUpdateRequest).Execute()



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
    endUserLicenseAgreementUpdateRequest := *openapiclient.NewEndUserLicenseAgreementUpdateRequest(*openapiclient.NewEndUserLicenseAgreementUpdateRequestData("Type_example", "Id_example")) // EndUserLicenseAgreementUpdateRequest | EndUserLicenseAgreement representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EndUserLicenseAgreementsApi.EndUserLicenseAgreementsUpdateInstance(context.Background(), id).EndUserLicenseAgreementUpdateRequest(endUserLicenseAgreementUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndUserLicenseAgreementsApi.EndUserLicenseAgreementsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EndUserLicenseAgreementsUpdateInstance`: EndUserLicenseAgreementResponse
    fmt.Fprintf(os.Stdout, "Response from `EndUserLicenseAgreementsApi.EndUserLicenseAgreementsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiEndUserLicenseAgreementsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endUserLicenseAgreementUpdateRequest** | [**EndUserLicenseAgreementUpdateRequest**](EndUserLicenseAgreementUpdateRequest.md) | EndUserLicenseAgreement representation | 

### Return type

[**EndUserLicenseAgreementResponse**](EndUserLicenseAgreementResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

