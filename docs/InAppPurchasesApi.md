# \InAppPurchasesApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InAppPurchasesGetInstance**](InAppPurchasesApi.md#InAppPurchasesGetInstance) | **Get** /v1/inAppPurchases/{id} | 



## InAppPurchasesGetInstance

> InAppPurchaseResponse InAppPurchasesGetInstance(ctx, id).FieldsInAppPurchases(fieldsInAppPurchases).Include(include).LimitApps(limitApps).Execute()



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
    fieldsInAppPurchases := []string{"FieldsInAppPurchases_example"} // []string | the fields to include for returned resources of type inAppPurchases (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    limitApps := int32(56) // int32 | maximum number of related apps returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InAppPurchasesApi.InAppPurchasesGetInstance(context.Background(), id).FieldsInAppPurchases(fieldsInAppPurchases).Include(include).LimitApps(limitApps).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InAppPurchasesApi.InAppPurchasesGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InAppPurchasesGetInstance`: InAppPurchaseResponse
    fmt.Fprintf(os.Stdout, "Response from `InAppPurchasesApi.InAppPurchasesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiInAppPurchasesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsInAppPurchases** | **[]string** | the fields to include for returned resources of type inAppPurchases | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitApps** | **int32** | maximum number of related apps returned (when they are included) | 

### Return type

[**InAppPurchaseResponse**](InAppPurchaseResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

