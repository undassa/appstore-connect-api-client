# \AppsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsAppInfosGetToManyRelated**](AppsApi.md#AppsAppInfosGetToManyRelated) | **Get** /v1/apps/{id}/appInfos | 
[**AppsAppStoreVersionsGetToManyRelated**](AppsApi.md#AppsAppStoreVersionsGetToManyRelated) | **Get** /v1/apps/{id}/appStoreVersions | 
[**AppsAvailableTerritoriesGetToManyRelated**](AppsApi.md#AppsAvailableTerritoriesGetToManyRelated) | **Get** /v1/apps/{id}/availableTerritories | 
[**AppsBetaAppLocalizationsGetToManyRelated**](AppsApi.md#AppsBetaAppLocalizationsGetToManyRelated) | **Get** /v1/apps/{id}/betaAppLocalizations | 
[**AppsBetaAppReviewDetailGetToOneRelated**](AppsApi.md#AppsBetaAppReviewDetailGetToOneRelated) | **Get** /v1/apps/{id}/betaAppReviewDetail | 
[**AppsBetaGroupsGetToManyRelated**](AppsApi.md#AppsBetaGroupsGetToManyRelated) | **Get** /v1/apps/{id}/betaGroups | 
[**AppsBetaLicenseAgreementGetToOneRelated**](AppsApi.md#AppsBetaLicenseAgreementGetToOneRelated) | **Get** /v1/apps/{id}/betaLicenseAgreement | 
[**AppsBetaTestersDeleteToManyRelationship**](AppsApi.md#AppsBetaTestersDeleteToManyRelationship) | **Delete** /v1/apps/{id}/relationships/betaTesters | 
[**AppsBuildsGetToManyRelated**](AppsApi.md#AppsBuildsGetToManyRelated) | **Get** /v1/apps/{id}/builds | 
[**AppsEndUserLicenseAgreementGetToOneRelated**](AppsApi.md#AppsEndUserLicenseAgreementGetToOneRelated) | **Get** /v1/apps/{id}/endUserLicenseAgreement | 
[**AppsGameCenterEnabledVersionsGetToManyRelated**](AppsApi.md#AppsGameCenterEnabledVersionsGetToManyRelated) | **Get** /v1/apps/{id}/gameCenterEnabledVersions | 
[**AppsGetCollection**](AppsApi.md#AppsGetCollection) | **Get** /v1/apps | 
[**AppsGetInstance**](AppsApi.md#AppsGetInstance) | **Get** /v1/apps/{id} | 
[**AppsInAppPurchasesGetToManyRelated**](AppsApi.md#AppsInAppPurchasesGetToManyRelated) | **Get** /v1/apps/{id}/inAppPurchases | 
[**AppsPerfPowerMetricsGetToManyRelated**](AppsApi.md#AppsPerfPowerMetricsGetToManyRelated) | **Get** /v1/apps/{id}/perfPowerMetrics | 
[**AppsPreOrderGetToOneRelated**](AppsApi.md#AppsPreOrderGetToOneRelated) | **Get** /v1/apps/{id}/preOrder | 
[**AppsPreReleaseVersionsGetToManyRelated**](AppsApi.md#AppsPreReleaseVersionsGetToManyRelated) | **Get** /v1/apps/{id}/preReleaseVersions | 
[**AppsPricesGetToManyRelated**](AppsApi.md#AppsPricesGetToManyRelated) | **Get** /v1/apps/{id}/prices | 
[**AppsUpdateInstance**](AppsApi.md#AppsUpdateInstance) | **Patch** /v1/apps/{id} | 



## AppsAppInfosGetToManyRelated

> AppInfosResponse AppsAppInfosGetToManyRelated(ctx, id).FieldsAppInfos(fieldsAppInfos).FieldsAppCategories(fieldsAppCategories).FieldsAppInfoLocalizations(fieldsAppInfoLocalizations).FieldsApps(fieldsApps).Limit(limit).Include(include).Execute()



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
    fieldsAppInfos := []string{"FieldsAppInfos_example"} // []string | the fields to include for returned resources of type appInfos (optional)
    fieldsAppCategories := []string{"FieldsAppCategories_example"} // []string | the fields to include for returned resources of type appCategories (optional)
    fieldsAppInfoLocalizations := []string{"FieldsAppInfoLocalizations_example"} // []string | the fields to include for returned resources of type appInfoLocalizations (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.AppsAppInfosGetToManyRelated(context.Background(), id).FieldsAppInfos(fieldsAppInfos).FieldsAppCategories(fieldsAppCategories).FieldsAppInfoLocalizations(fieldsAppInfoLocalizations).FieldsApps(fieldsApps).Limit(limit).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsAppInfosGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsAppInfosGetToManyRelated`: AppInfosResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsAppInfosGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAppInfosGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppInfos** | **[]string** | the fields to include for returned resources of type appInfos | 
 **fieldsAppCategories** | **[]string** | the fields to include for returned resources of type appCategories | 
 **fieldsAppInfoLocalizations** | **[]string** | the fields to include for returned resources of type appInfoLocalizations | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**AppInfosResponse**](AppInfosResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAppStoreVersionsGetToManyRelated

> AppStoreVersionsResponse AppsAppStoreVersionsGetToManyRelated(ctx, id).FilterAppStoreState(filterAppStoreState).FilterPlatform(filterPlatform).FilterVersionString(filterVersionString).FilterId(filterId).FieldsIdfaDeclarations(fieldsIdfaDeclarations).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).FieldsRoutingAppCoverages(fieldsRoutingAppCoverages).FieldsAppStoreVersionPhasedReleases(fieldsAppStoreVersionPhasedReleases).FieldsAgeRatingDeclarations(fieldsAgeRatingDeclarations).FieldsAppStoreReviewDetails(fieldsAppStoreReviewDetails).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsBuilds(fieldsBuilds).FieldsAppStoreVersionSubmissions(fieldsAppStoreVersionSubmissions).FieldsApps(fieldsApps).Limit(limit).Include(include).Execute()



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
    filterAppStoreState := []string{"FilterAppStoreState_example"} // []string | filter by attribute 'appStoreState' (optional)
    filterPlatform := []string{"FilterPlatform_example"} // []string | filter by attribute 'platform' (optional)
    filterVersionString := []string{"Inner_example"} // []string | filter by attribute 'versionString' (optional)
    filterId := []string{"Inner_example"} // []string | filter by id(s) (optional)
    fieldsIdfaDeclarations := []string{"FieldsIdfaDeclarations_example"} // []string | the fields to include for returned resources of type idfaDeclarations (optional)
    fieldsAppStoreVersionLocalizations := []string{"FieldsAppStoreVersionLocalizations_example"} // []string | the fields to include for returned resources of type appStoreVersionLocalizations (optional)
    fieldsRoutingAppCoverages := []string{"FieldsRoutingAppCoverages_example"} // []string | the fields to include for returned resources of type routingAppCoverages (optional)
    fieldsAppStoreVersionPhasedReleases := []string{"FieldsAppStoreVersionPhasedReleases_example"} // []string | the fields to include for returned resources of type appStoreVersionPhasedReleases (optional)
    fieldsAgeRatingDeclarations := []string{"FieldsAgeRatingDeclarations_example"} // []string | the fields to include for returned resources of type ageRatingDeclarations (optional)
    fieldsAppStoreReviewDetails := []string{"FieldsAppStoreReviewDetails_example"} // []string | the fields to include for returned resources of type appStoreReviewDetails (optional)
    fieldsAppStoreVersions := []string{"FieldsAppStoreVersions_example"} // []string | the fields to include for returned resources of type appStoreVersions (optional)
    fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)
    fieldsAppStoreVersionSubmissions := []string{"FieldsAppStoreVersionSubmissions_example"} // []string | the fields to include for returned resources of type appStoreVersionSubmissions (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.AppsAppStoreVersionsGetToManyRelated(context.Background(), id).FilterAppStoreState(filterAppStoreState).FilterPlatform(filterPlatform).FilterVersionString(filterVersionString).FilterId(filterId).FieldsIdfaDeclarations(fieldsIdfaDeclarations).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).FieldsRoutingAppCoverages(fieldsRoutingAppCoverages).FieldsAppStoreVersionPhasedReleases(fieldsAppStoreVersionPhasedReleases).FieldsAgeRatingDeclarations(fieldsAgeRatingDeclarations).FieldsAppStoreReviewDetails(fieldsAppStoreReviewDetails).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsBuilds(fieldsBuilds).FieldsAppStoreVersionSubmissions(fieldsAppStoreVersionSubmissions).FieldsApps(fieldsApps).Limit(limit).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsAppStoreVersionsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsAppStoreVersionsGetToManyRelated`: AppStoreVersionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsAppStoreVersionsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAppStoreVersionsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterAppStoreState** | **[]string** | filter by attribute &#39;appStoreState&#39; | 
 **filterPlatform** | **[]string** | filter by attribute &#39;platform&#39; | 
 **filterVersionString** | **[]string** | filter by attribute &#39;versionString&#39; | 
 **filterId** | **[]string** | filter by id(s) | 
 **fieldsIdfaDeclarations** | **[]string** | the fields to include for returned resources of type idfaDeclarations | 
 **fieldsAppStoreVersionLocalizations** | **[]string** | the fields to include for returned resources of type appStoreVersionLocalizations | 
 **fieldsRoutingAppCoverages** | **[]string** | the fields to include for returned resources of type routingAppCoverages | 
 **fieldsAppStoreVersionPhasedReleases** | **[]string** | the fields to include for returned resources of type appStoreVersionPhasedReleases | 
 **fieldsAgeRatingDeclarations** | **[]string** | the fields to include for returned resources of type ageRatingDeclarations | 
 **fieldsAppStoreReviewDetails** | **[]string** | the fields to include for returned resources of type appStoreReviewDetails | 
 **fieldsAppStoreVersions** | **[]string** | the fields to include for returned resources of type appStoreVersions | 
 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 
 **fieldsAppStoreVersionSubmissions** | **[]string** | the fields to include for returned resources of type appStoreVersionSubmissions | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**AppStoreVersionsResponse**](AppStoreVersionsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAvailableTerritoriesGetToManyRelated

> TerritoriesResponse AppsAvailableTerritoriesGetToManyRelated(ctx, id).FieldsTerritories(fieldsTerritories).Limit(limit).Execute()



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
    resp, r, err := api_client.AppsApi.AppsAvailableTerritoriesGetToManyRelated(context.Background(), id).FieldsTerritories(fieldsTerritories).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsAvailableTerritoriesGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsAvailableTerritoriesGetToManyRelated`: TerritoriesResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsAvailableTerritoriesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAvailableTerritoriesGetToManyRelatedRequest struct via the builder pattern


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


## AppsBetaAppLocalizationsGetToManyRelated

> BetaAppLocalizationsResponse AppsBetaAppLocalizationsGetToManyRelated(ctx, id).FieldsBetaAppLocalizations(fieldsBetaAppLocalizations).Limit(limit).Execute()



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
    fieldsBetaAppLocalizations := []string{"FieldsBetaAppLocalizations_example"} // []string | the fields to include for returned resources of type betaAppLocalizations (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.AppsBetaAppLocalizationsGetToManyRelated(context.Background(), id).FieldsBetaAppLocalizations(fieldsBetaAppLocalizations).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsBetaAppLocalizationsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsBetaAppLocalizationsGetToManyRelated`: BetaAppLocalizationsResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsBetaAppLocalizationsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsBetaAppLocalizationsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaAppLocalizations** | **[]string** | the fields to include for returned resources of type betaAppLocalizations | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**BetaAppLocalizationsResponse**](BetaAppLocalizationsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsBetaAppReviewDetailGetToOneRelated

> BetaAppReviewDetailResponse AppsBetaAppReviewDetailGetToOneRelated(ctx, id).FieldsBetaAppReviewDetails(fieldsBetaAppReviewDetails).Execute()



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
    fieldsBetaAppReviewDetails := []string{"FieldsBetaAppReviewDetails_example"} // []string | the fields to include for returned resources of type betaAppReviewDetails (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.AppsBetaAppReviewDetailGetToOneRelated(context.Background(), id).FieldsBetaAppReviewDetails(fieldsBetaAppReviewDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsBetaAppReviewDetailGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsBetaAppReviewDetailGetToOneRelated`: BetaAppReviewDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsBetaAppReviewDetailGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsBetaAppReviewDetailGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaAppReviewDetails** | **[]string** | the fields to include for returned resources of type betaAppReviewDetails | 

### Return type

[**BetaAppReviewDetailResponse**](BetaAppReviewDetailResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsBetaGroupsGetToManyRelated

> BetaGroupsResponse AppsBetaGroupsGetToManyRelated(ctx, id).FieldsBetaGroups(fieldsBetaGroups).Limit(limit).Execute()



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
    resp, r, err := api_client.AppsApi.AppsBetaGroupsGetToManyRelated(context.Background(), id).FieldsBetaGroups(fieldsBetaGroups).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsBetaGroupsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsBetaGroupsGetToManyRelated`: BetaGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsBetaGroupsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsBetaGroupsGetToManyRelatedRequest struct via the builder pattern


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


## AppsBetaLicenseAgreementGetToOneRelated

> BetaLicenseAgreementResponse AppsBetaLicenseAgreementGetToOneRelated(ctx, id).FieldsBetaLicenseAgreements(fieldsBetaLicenseAgreements).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.AppsBetaLicenseAgreementGetToOneRelated(context.Background(), id).FieldsBetaLicenseAgreements(fieldsBetaLicenseAgreements).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsBetaLicenseAgreementGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsBetaLicenseAgreementGetToOneRelated`: BetaLicenseAgreementResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsBetaLicenseAgreementGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsBetaLicenseAgreementGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaLicenseAgreements** | **[]string** | the fields to include for returned resources of type betaLicenseAgreements | 

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


## AppsBetaTestersDeleteToManyRelationship

> AppsBetaTestersDeleteToManyRelationship(ctx, id).AppBetaTestersLinkagesRequest(appBetaTestersLinkagesRequest).Execute()



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
    appBetaTestersLinkagesRequest := *openapiclient.NewAppBetaTestersLinkagesRequest([]openapiclient.BetaGroupRelationshipsBetaTestersData{*openapiclient.NewBetaGroupRelationshipsBetaTestersData("Type_example", "Id_example")}) // AppBetaTestersLinkagesRequest | List of related linkages

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.AppsBetaTestersDeleteToManyRelationship(context.Background(), id).AppBetaTestersLinkagesRequest(appBetaTestersLinkagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsBetaTestersDeleteToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppsBetaTestersDeleteToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appBetaTestersLinkagesRequest** | [**AppBetaTestersLinkagesRequest**](AppBetaTestersLinkagesRequest.md) | List of related linkages | 

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


## AppsBuildsGetToManyRelated

> BuildsResponse AppsBuildsGetToManyRelated(ctx, id).FieldsBuilds(fieldsBuilds).Limit(limit).Execute()



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
    resp, r, err := api_client.AppsApi.AppsBuildsGetToManyRelated(context.Background(), id).FieldsBuilds(fieldsBuilds).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsBuildsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsBuildsGetToManyRelated`: BuildsResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsBuildsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsBuildsGetToManyRelatedRequest struct via the builder pattern


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


## AppsEndUserLicenseAgreementGetToOneRelated

> EndUserLicenseAgreementResponse AppsEndUserLicenseAgreementGetToOneRelated(ctx, id).FieldsEndUserLicenseAgreements(fieldsEndUserLicenseAgreements).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.AppsEndUserLicenseAgreementGetToOneRelated(context.Background(), id).FieldsEndUserLicenseAgreements(fieldsEndUserLicenseAgreements).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsEndUserLicenseAgreementGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsEndUserLicenseAgreementGetToOneRelated`: EndUserLicenseAgreementResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsEndUserLicenseAgreementGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsEndUserLicenseAgreementGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsEndUserLicenseAgreements** | **[]string** | the fields to include for returned resources of type endUserLicenseAgreements | 

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


## AppsGameCenterEnabledVersionsGetToManyRelated

> GameCenterEnabledVersionsResponse AppsGameCenterEnabledVersionsGetToManyRelated(ctx, id).FilterPlatform(filterPlatform).FilterVersionString(filterVersionString).FilterId(filterId).Sort(sort).FieldsGameCenterEnabledVersions(fieldsGameCenterEnabledVersions).FieldsApps(fieldsApps).Limit(limit).Include(include).Execute()



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
    filterPlatform := []string{"FilterPlatform_example"} // []string | filter by attribute 'platform' (optional)
    filterVersionString := []string{"Inner_example"} // []string | filter by attribute 'versionString' (optional)
    filterId := []string{"Inner_example"} // []string | filter by id(s) (optional)
    sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; resources will be sorted as specified (optional)
    fieldsGameCenterEnabledVersions := []string{"FieldsGameCenterEnabledVersions_example"} // []string | the fields to include for returned resources of type gameCenterEnabledVersions (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.AppsGameCenterEnabledVersionsGetToManyRelated(context.Background(), id).FilterPlatform(filterPlatform).FilterVersionString(filterVersionString).FilterId(filterId).Sort(sort).FieldsGameCenterEnabledVersions(fieldsGameCenterEnabledVersions).FieldsApps(fieldsApps).Limit(limit).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsGameCenterEnabledVersionsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsGameCenterEnabledVersionsGetToManyRelated`: GameCenterEnabledVersionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsGameCenterEnabledVersionsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGameCenterEnabledVersionsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterPlatform** | **[]string** | filter by attribute &#39;platform&#39; | 
 **filterVersionString** | **[]string** | filter by attribute &#39;versionString&#39; | 
 **filterId** | **[]string** | filter by id(s) | 
 **sort** | **[]string** | comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsGameCenterEnabledVersions** | **[]string** | the fields to include for returned resources of type gameCenterEnabledVersions | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**GameCenterEnabledVersionsResponse**](GameCenterEnabledVersionsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsGetCollection

> AppsResponse AppsGetCollection(ctx).FilterAppStoreVersionsAppStoreState(filterAppStoreVersionsAppStoreState).FilterAppStoreVersionsPlatform(filterAppStoreVersionsPlatform).FilterBundleId(filterBundleId).FilterName(filterName).FilterSku(filterSku).FilterAppStoreVersions(filterAppStoreVersions).FilterId(filterId).ExistsGameCenterEnabledVersions(existsGameCenterEnabledVersions).Sort(sort).FieldsApps(fieldsApps).Limit(limit).Include(include).FieldsBetaGroups(fieldsBetaGroups).FieldsPerfPowerMetrics(fieldsPerfPowerMetrics).FieldsAppInfos(fieldsAppInfos).FieldsAppPreOrders(fieldsAppPreOrders).FieldsPreReleaseVersions(fieldsPreReleaseVersions).FieldsAppPrices(fieldsAppPrices).FieldsInAppPurchases(fieldsInAppPurchases).FieldsBetaAppReviewDetails(fieldsBetaAppReviewDetails).FieldsTerritories(fieldsTerritories).FieldsGameCenterEnabledVersions(fieldsGameCenterEnabledVersions).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsBuilds(fieldsBuilds).FieldsBetaAppLocalizations(fieldsBetaAppLocalizations).FieldsBetaLicenseAgreements(fieldsBetaLicenseAgreements).FieldsEndUserLicenseAgreements(fieldsEndUserLicenseAgreements).LimitAppInfos(limitAppInfos).LimitAppStoreVersions(limitAppStoreVersions).LimitAvailableTerritories(limitAvailableTerritories).LimitBetaAppLocalizations(limitBetaAppLocalizations).LimitBetaGroups(limitBetaGroups).LimitBuilds(limitBuilds).LimitGameCenterEnabledVersions(limitGameCenterEnabledVersions).LimitInAppPurchases(limitInAppPurchases).LimitPreReleaseVersions(limitPreReleaseVersions).LimitPrices(limitPrices).Execute()



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
    filterAppStoreVersionsAppStoreState := []string{"FilterAppStoreVersionsAppStoreState_example"} // []string | filter by attribute 'appStoreVersions.appStoreState' (optional)
    filterAppStoreVersionsPlatform := []string{"FilterAppStoreVersionsPlatform_example"} // []string | filter by attribute 'appStoreVersions.platform' (optional)
    filterBundleId := []string{"Inner_example"} // []string | filter by attribute 'bundleId' (optional)
    filterName := []string{"Inner_example"} // []string | filter by attribute 'name' (optional)
    filterSku := []string{"Inner_example"} // []string | filter by attribute 'sku' (optional)
    filterAppStoreVersions := []string{"Inner_example"} // []string | filter by id(s) of related 'appStoreVersions' (optional)
    filterId := []string{"Inner_example"} // []string | filter by id(s) (optional)
    existsGameCenterEnabledVersions := []string{"Inner_example"} // []string | filter by existence or non-existence of related 'gameCenterEnabledVersions' (optional)
    sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; resources will be sorted as specified (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsBetaGroups := []string{"FieldsBetaGroups_example"} // []string | the fields to include for returned resources of type betaGroups (optional)
    fieldsPerfPowerMetrics := []string{"FieldsPerfPowerMetrics_example"} // []string | the fields to include for returned resources of type perfPowerMetrics (optional)
    fieldsAppInfos := []string{"FieldsAppInfos_example"} // []string | the fields to include for returned resources of type appInfos (optional)
    fieldsAppPreOrders := []string{"FieldsAppPreOrders_example"} // []string | the fields to include for returned resources of type appPreOrders (optional)
    fieldsPreReleaseVersions := []string{"FieldsPreReleaseVersions_example"} // []string | the fields to include for returned resources of type preReleaseVersions (optional)
    fieldsAppPrices := []string{"FieldsAppPrices_example"} // []string | the fields to include for returned resources of type appPrices (optional)
    fieldsInAppPurchases := []string{"FieldsInAppPurchases_example"} // []string | the fields to include for returned resources of type inAppPurchases (optional)
    fieldsBetaAppReviewDetails := []string{"FieldsBetaAppReviewDetails_example"} // []string | the fields to include for returned resources of type betaAppReviewDetails (optional)
    fieldsTerritories := []string{"FieldsTerritories_example"} // []string | the fields to include for returned resources of type territories (optional)
    fieldsGameCenterEnabledVersions := []string{"FieldsGameCenterEnabledVersions_example"} // []string | the fields to include for returned resources of type gameCenterEnabledVersions (optional)
    fieldsAppStoreVersions := []string{"FieldsAppStoreVersions_example"} // []string | the fields to include for returned resources of type appStoreVersions (optional)
    fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)
    fieldsBetaAppLocalizations := []string{"FieldsBetaAppLocalizations_example"} // []string | the fields to include for returned resources of type betaAppLocalizations (optional)
    fieldsBetaLicenseAgreements := []string{"FieldsBetaLicenseAgreements_example"} // []string | the fields to include for returned resources of type betaLicenseAgreements (optional)
    fieldsEndUserLicenseAgreements := []string{"FieldsEndUserLicenseAgreements_example"} // []string | the fields to include for returned resources of type endUserLicenseAgreements (optional)
    limitAppInfos := int32(56) // int32 | maximum number of related appInfos returned (when they are included) (optional)
    limitAppStoreVersions := int32(56) // int32 | maximum number of related appStoreVersions returned (when they are included) (optional)
    limitAvailableTerritories := int32(56) // int32 | maximum number of related availableTerritories returned (when they are included) (optional)
    limitBetaAppLocalizations := int32(56) // int32 | maximum number of related betaAppLocalizations returned (when they are included) (optional)
    limitBetaGroups := int32(56) // int32 | maximum number of related betaGroups returned (when they are included) (optional)
    limitBuilds := int32(56) // int32 | maximum number of related builds returned (when they are included) (optional)
    limitGameCenterEnabledVersions := int32(56) // int32 | maximum number of related gameCenterEnabledVersions returned (when they are included) (optional)
    limitInAppPurchases := int32(56) // int32 | maximum number of related inAppPurchases returned (when they are included) (optional)
    limitPreReleaseVersions := int32(56) // int32 | maximum number of related preReleaseVersions returned (when they are included) (optional)
    limitPrices := int32(56) // int32 | maximum number of related prices returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.AppsGetCollection(context.Background()).FilterAppStoreVersionsAppStoreState(filterAppStoreVersionsAppStoreState).FilterAppStoreVersionsPlatform(filterAppStoreVersionsPlatform).FilterBundleId(filterBundleId).FilterName(filterName).FilterSku(filterSku).FilterAppStoreVersions(filterAppStoreVersions).FilterId(filterId).ExistsGameCenterEnabledVersions(existsGameCenterEnabledVersions).Sort(sort).FieldsApps(fieldsApps).Limit(limit).Include(include).FieldsBetaGroups(fieldsBetaGroups).FieldsPerfPowerMetrics(fieldsPerfPowerMetrics).FieldsAppInfos(fieldsAppInfos).FieldsAppPreOrders(fieldsAppPreOrders).FieldsPreReleaseVersions(fieldsPreReleaseVersions).FieldsAppPrices(fieldsAppPrices).FieldsInAppPurchases(fieldsInAppPurchases).FieldsBetaAppReviewDetails(fieldsBetaAppReviewDetails).FieldsTerritories(fieldsTerritories).FieldsGameCenterEnabledVersions(fieldsGameCenterEnabledVersions).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsBuilds(fieldsBuilds).FieldsBetaAppLocalizations(fieldsBetaAppLocalizations).FieldsBetaLicenseAgreements(fieldsBetaLicenseAgreements).FieldsEndUserLicenseAgreements(fieldsEndUserLicenseAgreements).LimitAppInfos(limitAppInfos).LimitAppStoreVersions(limitAppStoreVersions).LimitAvailableTerritories(limitAvailableTerritories).LimitBetaAppLocalizations(limitBetaAppLocalizations).LimitBetaGroups(limitBetaGroups).LimitBuilds(limitBuilds).LimitGameCenterEnabledVersions(limitGameCenterEnabledVersions).LimitInAppPurchases(limitInAppPurchases).LimitPreReleaseVersions(limitPreReleaseVersions).LimitPrices(limitPrices).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsGetCollection`: AppsResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterAppStoreVersionsAppStoreState** | **[]string** | filter by attribute &#39;appStoreVersions.appStoreState&#39; | 
 **filterAppStoreVersionsPlatform** | **[]string** | filter by attribute &#39;appStoreVersions.platform&#39; | 
 **filterBundleId** | **[]string** | filter by attribute &#39;bundleId&#39; | 
 **filterName** | **[]string** | filter by attribute &#39;name&#39; | 
 **filterSku** | **[]string** | filter by attribute &#39;sku&#39; | 
 **filterAppStoreVersions** | **[]string** | filter by id(s) of related &#39;appStoreVersions&#39; | 
 **filterId** | **[]string** | filter by id(s) | 
 **existsGameCenterEnabledVersions** | **[]string** | filter by existence or non-existence of related &#39;gameCenterEnabledVersions&#39; | 
 **sort** | **[]string** | comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsBetaGroups** | **[]string** | the fields to include for returned resources of type betaGroups | 
 **fieldsPerfPowerMetrics** | **[]string** | the fields to include for returned resources of type perfPowerMetrics | 
 **fieldsAppInfos** | **[]string** | the fields to include for returned resources of type appInfos | 
 **fieldsAppPreOrders** | **[]string** | the fields to include for returned resources of type appPreOrders | 
 **fieldsPreReleaseVersions** | **[]string** | the fields to include for returned resources of type preReleaseVersions | 
 **fieldsAppPrices** | **[]string** | the fields to include for returned resources of type appPrices | 
 **fieldsInAppPurchases** | **[]string** | the fields to include for returned resources of type inAppPurchases | 
 **fieldsBetaAppReviewDetails** | **[]string** | the fields to include for returned resources of type betaAppReviewDetails | 
 **fieldsTerritories** | **[]string** | the fields to include for returned resources of type territories | 
 **fieldsGameCenterEnabledVersions** | **[]string** | the fields to include for returned resources of type gameCenterEnabledVersions | 
 **fieldsAppStoreVersions** | **[]string** | the fields to include for returned resources of type appStoreVersions | 
 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 
 **fieldsBetaAppLocalizations** | **[]string** | the fields to include for returned resources of type betaAppLocalizations | 
 **fieldsBetaLicenseAgreements** | **[]string** | the fields to include for returned resources of type betaLicenseAgreements | 
 **fieldsEndUserLicenseAgreements** | **[]string** | the fields to include for returned resources of type endUserLicenseAgreements | 
 **limitAppInfos** | **int32** | maximum number of related appInfos returned (when they are included) | 
 **limitAppStoreVersions** | **int32** | maximum number of related appStoreVersions returned (when they are included) | 
 **limitAvailableTerritories** | **int32** | maximum number of related availableTerritories returned (when they are included) | 
 **limitBetaAppLocalizations** | **int32** | maximum number of related betaAppLocalizations returned (when they are included) | 
 **limitBetaGroups** | **int32** | maximum number of related betaGroups returned (when they are included) | 
 **limitBuilds** | **int32** | maximum number of related builds returned (when they are included) | 
 **limitGameCenterEnabledVersions** | **int32** | maximum number of related gameCenterEnabledVersions returned (when they are included) | 
 **limitInAppPurchases** | **int32** | maximum number of related inAppPurchases returned (when they are included) | 
 **limitPreReleaseVersions** | **int32** | maximum number of related preReleaseVersions returned (when they are included) | 
 **limitPrices** | **int32** | maximum number of related prices returned (when they are included) | 

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


## AppsGetInstance

> AppResponse AppsGetInstance(ctx, id).FieldsApps(fieldsApps).Include(include).FieldsBetaGroups(fieldsBetaGroups).FieldsPerfPowerMetrics(fieldsPerfPowerMetrics).FieldsAppInfos(fieldsAppInfos).FieldsAppPreOrders(fieldsAppPreOrders).FieldsPreReleaseVersions(fieldsPreReleaseVersions).FieldsAppPrices(fieldsAppPrices).FieldsInAppPurchases(fieldsInAppPurchases).FieldsBetaAppReviewDetails(fieldsBetaAppReviewDetails).FieldsTerritories(fieldsTerritories).FieldsGameCenterEnabledVersions(fieldsGameCenterEnabledVersions).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsBuilds(fieldsBuilds).FieldsBetaAppLocalizations(fieldsBetaAppLocalizations).FieldsBetaLicenseAgreements(fieldsBetaLicenseAgreements).FieldsEndUserLicenseAgreements(fieldsEndUserLicenseAgreements).LimitAppInfos(limitAppInfos).LimitAppStoreVersions(limitAppStoreVersions).LimitAvailableTerritories(limitAvailableTerritories).LimitBetaAppLocalizations(limitBetaAppLocalizations).LimitBetaGroups(limitBetaGroups).LimitBuilds(limitBuilds).LimitGameCenterEnabledVersions(limitGameCenterEnabledVersions).LimitInAppPurchases(limitInAppPurchases).LimitPreReleaseVersions(limitPreReleaseVersions).LimitPrices(limitPrices).Execute()



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
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsBetaGroups := []string{"FieldsBetaGroups_example"} // []string | the fields to include for returned resources of type betaGroups (optional)
    fieldsPerfPowerMetrics := []string{"FieldsPerfPowerMetrics_example"} // []string | the fields to include for returned resources of type perfPowerMetrics (optional)
    fieldsAppInfos := []string{"FieldsAppInfos_example"} // []string | the fields to include for returned resources of type appInfos (optional)
    fieldsAppPreOrders := []string{"FieldsAppPreOrders_example"} // []string | the fields to include for returned resources of type appPreOrders (optional)
    fieldsPreReleaseVersions := []string{"FieldsPreReleaseVersions_example"} // []string | the fields to include for returned resources of type preReleaseVersions (optional)
    fieldsAppPrices := []string{"FieldsAppPrices_example"} // []string | the fields to include for returned resources of type appPrices (optional)
    fieldsInAppPurchases := []string{"FieldsInAppPurchases_example"} // []string | the fields to include for returned resources of type inAppPurchases (optional)
    fieldsBetaAppReviewDetails := []string{"FieldsBetaAppReviewDetails_example"} // []string | the fields to include for returned resources of type betaAppReviewDetails (optional)
    fieldsTerritories := []string{"FieldsTerritories_example"} // []string | the fields to include for returned resources of type territories (optional)
    fieldsGameCenterEnabledVersions := []string{"FieldsGameCenterEnabledVersions_example"} // []string | the fields to include for returned resources of type gameCenterEnabledVersions (optional)
    fieldsAppStoreVersions := []string{"FieldsAppStoreVersions_example"} // []string | the fields to include for returned resources of type appStoreVersions (optional)
    fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)
    fieldsBetaAppLocalizations := []string{"FieldsBetaAppLocalizations_example"} // []string | the fields to include for returned resources of type betaAppLocalizations (optional)
    fieldsBetaLicenseAgreements := []string{"FieldsBetaLicenseAgreements_example"} // []string | the fields to include for returned resources of type betaLicenseAgreements (optional)
    fieldsEndUserLicenseAgreements := []string{"FieldsEndUserLicenseAgreements_example"} // []string | the fields to include for returned resources of type endUserLicenseAgreements (optional)
    limitAppInfos := int32(56) // int32 | maximum number of related appInfos returned (when they are included) (optional)
    limitAppStoreVersions := int32(56) // int32 | maximum number of related appStoreVersions returned (when they are included) (optional)
    limitAvailableTerritories := int32(56) // int32 | maximum number of related availableTerritories returned (when they are included) (optional)
    limitBetaAppLocalizations := int32(56) // int32 | maximum number of related betaAppLocalizations returned (when they are included) (optional)
    limitBetaGroups := int32(56) // int32 | maximum number of related betaGroups returned (when they are included) (optional)
    limitBuilds := int32(56) // int32 | maximum number of related builds returned (when they are included) (optional)
    limitGameCenterEnabledVersions := int32(56) // int32 | maximum number of related gameCenterEnabledVersions returned (when they are included) (optional)
    limitInAppPurchases := int32(56) // int32 | maximum number of related inAppPurchases returned (when they are included) (optional)
    limitPreReleaseVersions := int32(56) // int32 | maximum number of related preReleaseVersions returned (when they are included) (optional)
    limitPrices := int32(56) // int32 | maximum number of related prices returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.AppsGetInstance(context.Background(), id).FieldsApps(fieldsApps).Include(include).FieldsBetaGroups(fieldsBetaGroups).FieldsPerfPowerMetrics(fieldsPerfPowerMetrics).FieldsAppInfos(fieldsAppInfos).FieldsAppPreOrders(fieldsAppPreOrders).FieldsPreReleaseVersions(fieldsPreReleaseVersions).FieldsAppPrices(fieldsAppPrices).FieldsInAppPurchases(fieldsInAppPurchases).FieldsBetaAppReviewDetails(fieldsBetaAppReviewDetails).FieldsTerritories(fieldsTerritories).FieldsGameCenterEnabledVersions(fieldsGameCenterEnabledVersions).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsBuilds(fieldsBuilds).FieldsBetaAppLocalizations(fieldsBetaAppLocalizations).FieldsBetaLicenseAgreements(fieldsBetaLicenseAgreements).FieldsEndUserLicenseAgreements(fieldsEndUserLicenseAgreements).LimitAppInfos(limitAppInfos).LimitAppStoreVersions(limitAppStoreVersions).LimitAvailableTerritories(limitAvailableTerritories).LimitBetaAppLocalizations(limitBetaAppLocalizations).LimitBetaGroups(limitBetaGroups).LimitBuilds(limitBuilds).LimitGameCenterEnabledVersions(limitGameCenterEnabledVersions).LimitInAppPurchases(limitInAppPurchases).LimitPreReleaseVersions(limitPreReleaseVersions).LimitPrices(limitPrices).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsGetInstance`: AppResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsBetaGroups** | **[]string** | the fields to include for returned resources of type betaGroups | 
 **fieldsPerfPowerMetrics** | **[]string** | the fields to include for returned resources of type perfPowerMetrics | 
 **fieldsAppInfos** | **[]string** | the fields to include for returned resources of type appInfos | 
 **fieldsAppPreOrders** | **[]string** | the fields to include for returned resources of type appPreOrders | 
 **fieldsPreReleaseVersions** | **[]string** | the fields to include for returned resources of type preReleaseVersions | 
 **fieldsAppPrices** | **[]string** | the fields to include for returned resources of type appPrices | 
 **fieldsInAppPurchases** | **[]string** | the fields to include for returned resources of type inAppPurchases | 
 **fieldsBetaAppReviewDetails** | **[]string** | the fields to include for returned resources of type betaAppReviewDetails | 
 **fieldsTerritories** | **[]string** | the fields to include for returned resources of type territories | 
 **fieldsGameCenterEnabledVersions** | **[]string** | the fields to include for returned resources of type gameCenterEnabledVersions | 
 **fieldsAppStoreVersions** | **[]string** | the fields to include for returned resources of type appStoreVersions | 
 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 
 **fieldsBetaAppLocalizations** | **[]string** | the fields to include for returned resources of type betaAppLocalizations | 
 **fieldsBetaLicenseAgreements** | **[]string** | the fields to include for returned resources of type betaLicenseAgreements | 
 **fieldsEndUserLicenseAgreements** | **[]string** | the fields to include for returned resources of type endUserLicenseAgreements | 
 **limitAppInfos** | **int32** | maximum number of related appInfos returned (when they are included) | 
 **limitAppStoreVersions** | **int32** | maximum number of related appStoreVersions returned (when they are included) | 
 **limitAvailableTerritories** | **int32** | maximum number of related availableTerritories returned (when they are included) | 
 **limitBetaAppLocalizations** | **int32** | maximum number of related betaAppLocalizations returned (when they are included) | 
 **limitBetaGroups** | **int32** | maximum number of related betaGroups returned (when they are included) | 
 **limitBuilds** | **int32** | maximum number of related builds returned (when they are included) | 
 **limitGameCenterEnabledVersions** | **int32** | maximum number of related gameCenterEnabledVersions returned (when they are included) | 
 **limitInAppPurchases** | **int32** | maximum number of related inAppPurchases returned (when they are included) | 
 **limitPreReleaseVersions** | **int32** | maximum number of related preReleaseVersions returned (when they are included) | 
 **limitPrices** | **int32** | maximum number of related prices returned (when they are included) | 

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


## AppsInAppPurchasesGetToManyRelated

> InAppPurchasesResponse AppsInAppPurchasesGetToManyRelated(ctx, id).FilterInAppPurchaseType(filterInAppPurchaseType).FilterCanBeSubmitted(filterCanBeSubmitted).Sort(sort).FieldsInAppPurchases(fieldsInAppPurchases).FieldsApps(fieldsApps).Limit(limit).Include(include).Execute()



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
    filterInAppPurchaseType := []string{"FilterInAppPurchaseType_example"} // []string | filter by attribute 'inAppPurchaseType' (optional)
    filterCanBeSubmitted := []string{"Inner_example"} // []string | filter by canBeSubmitted (optional)
    sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; resources will be sorted as specified (optional)
    fieldsInAppPurchases := []string{"FieldsInAppPurchases_example"} // []string | the fields to include for returned resources of type inAppPurchases (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.AppsInAppPurchasesGetToManyRelated(context.Background(), id).FilterInAppPurchaseType(filterInAppPurchaseType).FilterCanBeSubmitted(filterCanBeSubmitted).Sort(sort).FieldsInAppPurchases(fieldsInAppPurchases).FieldsApps(fieldsApps).Limit(limit).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsInAppPurchasesGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsInAppPurchasesGetToManyRelated`: InAppPurchasesResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsInAppPurchasesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsInAppPurchasesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterInAppPurchaseType** | **[]string** | filter by attribute &#39;inAppPurchaseType&#39; | 
 **filterCanBeSubmitted** | **[]string** | filter by canBeSubmitted | 
 **sort** | **[]string** | comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsInAppPurchases** | **[]string** | the fields to include for returned resources of type inAppPurchases | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**InAppPurchasesResponse**](InAppPurchasesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsPerfPowerMetricsGetToManyRelated

> PerfPowerMetricsResponse AppsPerfPowerMetricsGetToManyRelated(ctx, id).FilterDeviceType(filterDeviceType).FilterMetricType(filterMetricType).FilterPlatform(filterPlatform).Execute()



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
    resp, r, err := api_client.AppsApi.AppsPerfPowerMetricsGetToManyRelated(context.Background(), id).FilterDeviceType(filterDeviceType).FilterMetricType(filterMetricType).FilterPlatform(filterPlatform).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsPerfPowerMetricsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsPerfPowerMetricsGetToManyRelated`: PerfPowerMetricsResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsPerfPowerMetricsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsPerfPowerMetricsGetToManyRelatedRequest struct via the builder pattern


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


## AppsPreOrderGetToOneRelated

> AppPreOrderResponse AppsPreOrderGetToOneRelated(ctx, id).FieldsAppPreOrders(fieldsAppPreOrders).Execute()



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
    fieldsAppPreOrders := []string{"FieldsAppPreOrders_example"} // []string | the fields to include for returned resources of type appPreOrders (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.AppsPreOrderGetToOneRelated(context.Background(), id).FieldsAppPreOrders(fieldsAppPreOrders).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsPreOrderGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsPreOrderGetToOneRelated`: AppPreOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsPreOrderGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsPreOrderGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppPreOrders** | **[]string** | the fields to include for returned resources of type appPreOrders | 

### Return type

[**AppPreOrderResponse**](AppPreOrderResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsPreReleaseVersionsGetToManyRelated

> PreReleaseVersionsResponse AppsPreReleaseVersionsGetToManyRelated(ctx, id).FieldsPreReleaseVersions(fieldsPreReleaseVersions).Limit(limit).Execute()



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
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.AppsPreReleaseVersionsGetToManyRelated(context.Background(), id).FieldsPreReleaseVersions(fieldsPreReleaseVersions).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsPreReleaseVersionsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsPreReleaseVersionsGetToManyRelated`: PreReleaseVersionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsPreReleaseVersionsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsPreReleaseVersionsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsPreReleaseVersions** | **[]string** | the fields to include for returned resources of type preReleaseVersions | 
 **limit** | **int32** | maximum resources per page | 

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


## AppsPricesGetToManyRelated

> AppPricesResponse AppsPricesGetToManyRelated(ctx, id).FieldsAppPrices(fieldsAppPrices).FieldsAppPriceTiers(fieldsAppPriceTiers).FieldsApps(fieldsApps).Limit(limit).Include(include).Execute()



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
    fieldsAppPrices := []string{"FieldsAppPrices_example"} // []string | the fields to include for returned resources of type appPrices (optional)
    fieldsAppPriceTiers := []string{"FieldsAppPriceTiers_example"} // []string | the fields to include for returned resources of type appPriceTiers (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.AppsPricesGetToManyRelated(context.Background(), id).FieldsAppPrices(fieldsAppPrices).FieldsAppPriceTiers(fieldsAppPriceTiers).FieldsApps(fieldsApps).Limit(limit).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsPricesGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsPricesGetToManyRelated`: AppPricesResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsPricesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsPricesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppPrices** | **[]string** | the fields to include for returned resources of type appPrices | 
 **fieldsAppPriceTiers** | **[]string** | the fields to include for returned resources of type appPriceTiers | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**AppPricesResponse**](AppPricesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsUpdateInstance

> AppResponse AppsUpdateInstance(ctx, id).AppUpdateRequest(appUpdateRequest).Execute()



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
    appUpdateRequest := *openapiclient.NewAppUpdateRequest(*openapiclient.NewAppUpdateRequestData("Type_example", "Id_example")) // AppUpdateRequest | App representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.AppsUpdateInstance(context.Background(), id).AppUpdateRequest(appUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsUpdateInstance`: AppResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appUpdateRequest** | [**AppUpdateRequest**](AppUpdateRequest.md) | App representation | 

### Return type

[**AppResponse**](AppResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

