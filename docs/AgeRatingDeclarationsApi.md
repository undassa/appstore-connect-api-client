# \AgeRatingDeclarationsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgeRatingDeclarationsUpdateInstance**](AgeRatingDeclarationsApi.md#AgeRatingDeclarationsUpdateInstance) | **Patch** /v1/ageRatingDeclarations/{id} | 



## AgeRatingDeclarationsUpdateInstance

> AgeRatingDeclarationResponse AgeRatingDeclarationsUpdateInstance(ctx, id).AgeRatingDeclarationUpdateRequest(ageRatingDeclarationUpdateRequest).Execute()



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
    ageRatingDeclarationUpdateRequest := *openapiclient.NewAgeRatingDeclarationUpdateRequest(*openapiclient.NewAgeRatingDeclarationUpdateRequestData("Type_example", "Id_example")) // AgeRatingDeclarationUpdateRequest | AgeRatingDeclaration representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AgeRatingDeclarationsApi.AgeRatingDeclarationsUpdateInstance(context.Background(), id).AgeRatingDeclarationUpdateRequest(ageRatingDeclarationUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgeRatingDeclarationsApi.AgeRatingDeclarationsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgeRatingDeclarationsUpdateInstance`: AgeRatingDeclarationResponse
    fmt.Fprintf(os.Stdout, "Response from `AgeRatingDeclarationsApi.AgeRatingDeclarationsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgeRatingDeclarationsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ageRatingDeclarationUpdateRequest** | [**AgeRatingDeclarationUpdateRequest**](AgeRatingDeclarationUpdateRequest.md) | AgeRatingDeclaration representation | 

### Return type

[**AgeRatingDeclarationResponse**](AgeRatingDeclarationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

