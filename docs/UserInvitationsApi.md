# \UserInvitationsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserInvitationsCreateInstance**](UserInvitationsApi.md#UserInvitationsCreateInstance) | **Post** /v1/userInvitations | 
[**UserInvitationsDeleteInstance**](UserInvitationsApi.md#UserInvitationsDeleteInstance) | **Delete** /v1/userInvitations/{id} | 
[**UserInvitationsGetCollection**](UserInvitationsApi.md#UserInvitationsGetCollection) | **Get** /v1/userInvitations | 
[**UserInvitationsGetInstance**](UserInvitationsApi.md#UserInvitationsGetInstance) | **Get** /v1/userInvitations/{id} | 
[**UserInvitationsVisibleAppsGetToManyRelated**](UserInvitationsApi.md#UserInvitationsVisibleAppsGetToManyRelated) | **Get** /v1/userInvitations/{id}/visibleApps | 



## UserInvitationsCreateInstance

> UserInvitationResponse UserInvitationsCreateInstance(ctx).UserInvitationCreateRequest(userInvitationCreateRequest).Execute()



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
    userInvitationCreateRequest := *openapiclient.NewUserInvitationCreateRequest(*openapiclient.NewUserInvitationCreateRequestData("Type_example", *openapiclient.NewUserInvitationCreateRequestDataAttributes("Email_example", "FirstName_example", "LastName_example", []openapiclient.UserRole{openapiclient.UserRole("ADMIN")}))) // UserInvitationCreateRequest | UserInvitation representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserInvitationsApi.UserInvitationsCreateInstance(context.Background()).UserInvitationCreateRequest(userInvitationCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserInvitationsApi.UserInvitationsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserInvitationsCreateInstance`: UserInvitationResponse
    fmt.Fprintf(os.Stdout, "Response from `UserInvitationsApi.UserInvitationsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserInvitationsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userInvitationCreateRequest** | [**UserInvitationCreateRequest**](UserInvitationCreateRequest.md) | UserInvitation representation | 

### Return type

[**UserInvitationResponse**](UserInvitationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserInvitationsDeleteInstance

> UserInvitationsDeleteInstance(ctx, id).Execute()



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
    resp, r, err := api_client.UserInvitationsApi.UserInvitationsDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserInvitationsApi.UserInvitationsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUserInvitationsDeleteInstanceRequest struct via the builder pattern


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


## UserInvitationsGetCollection

> UserInvitationsResponse UserInvitationsGetCollection(ctx).FilterEmail(filterEmail).FilterRoles(filterRoles).FilterVisibleApps(filterVisibleApps).Sort(sort).FieldsUserInvitations(fieldsUserInvitations).Limit(limit).Include(include).FieldsApps(fieldsApps).LimitVisibleApps(limitVisibleApps).Execute()



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
    filterRoles := []string{"FilterRoles_example"} // []string | filter by attribute 'roles' (optional)
    filterVisibleApps := []string{"Inner_example"} // []string | filter by id(s) of related 'visibleApps' (optional)
    sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; resources will be sorted as specified (optional)
    fieldsUserInvitations := []string{"FieldsUserInvitations_example"} // []string | the fields to include for returned resources of type userInvitations (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
    limitVisibleApps := int32(56) // int32 | maximum number of related visibleApps returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserInvitationsApi.UserInvitationsGetCollection(context.Background()).FilterEmail(filterEmail).FilterRoles(filterRoles).FilterVisibleApps(filterVisibleApps).Sort(sort).FieldsUserInvitations(fieldsUserInvitations).Limit(limit).Include(include).FieldsApps(fieldsApps).LimitVisibleApps(limitVisibleApps).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserInvitationsApi.UserInvitationsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserInvitationsGetCollection`: UserInvitationsResponse
    fmt.Fprintf(os.Stdout, "Response from `UserInvitationsApi.UserInvitationsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserInvitationsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterEmail** | **[]string** | filter by attribute &#39;email&#39; | 
 **filterRoles** | **[]string** | filter by attribute &#39;roles&#39; | 
 **filterVisibleApps** | **[]string** | filter by id(s) of related &#39;visibleApps&#39; | 
 **sort** | **[]string** | comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsUserInvitations** | **[]string** | the fields to include for returned resources of type userInvitations | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **limitVisibleApps** | **int32** | maximum number of related visibleApps returned (when they are included) | 

### Return type

[**UserInvitationsResponse**](UserInvitationsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserInvitationsGetInstance

> UserInvitationResponse UserInvitationsGetInstance(ctx, id).FieldsUserInvitations(fieldsUserInvitations).Include(include).FieldsApps(fieldsApps).LimitVisibleApps(limitVisibleApps).Execute()



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
    fieldsUserInvitations := []string{"FieldsUserInvitations_example"} // []string | the fields to include for returned resources of type userInvitations (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
    limitVisibleApps := int32(56) // int32 | maximum number of related visibleApps returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserInvitationsApi.UserInvitationsGetInstance(context.Background(), id).FieldsUserInvitations(fieldsUserInvitations).Include(include).FieldsApps(fieldsApps).LimitVisibleApps(limitVisibleApps).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserInvitationsApi.UserInvitationsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserInvitationsGetInstance`: UserInvitationResponse
    fmt.Fprintf(os.Stdout, "Response from `UserInvitationsApi.UserInvitationsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserInvitationsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsUserInvitations** | **[]string** | the fields to include for returned resources of type userInvitations | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **limitVisibleApps** | **int32** | maximum number of related visibleApps returned (when they are included) | 

### Return type

[**UserInvitationResponse**](UserInvitationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserInvitationsVisibleAppsGetToManyRelated

> AppsResponse UserInvitationsVisibleAppsGetToManyRelated(ctx, id).FieldsApps(fieldsApps).Limit(limit).Execute()



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
    resp, r, err := api_client.UserInvitationsApi.UserInvitationsVisibleAppsGetToManyRelated(context.Background(), id).FieldsApps(fieldsApps).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserInvitationsApi.UserInvitationsVisibleAppsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserInvitationsVisibleAppsGetToManyRelated`: AppsResponse
    fmt.Fprintf(os.Stdout, "Response from `UserInvitationsApi.UserInvitationsVisibleAppsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserInvitationsVisibleAppsGetToManyRelatedRequest struct via the builder pattern


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

