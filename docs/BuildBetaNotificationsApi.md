# \BuildBetaNotificationsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BuildBetaNotificationsCreateInstance**](BuildBetaNotificationsApi.md#BuildBetaNotificationsCreateInstance) | **Post** /v1/buildBetaNotifications | 



## BuildBetaNotificationsCreateInstance

> BuildBetaNotificationResponse BuildBetaNotificationsCreateInstance(ctx).BuildBetaNotificationCreateRequest(buildBetaNotificationCreateRequest).Execute()



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
    buildBetaNotificationCreateRequest := *openapiclient.NewBuildBetaNotificationCreateRequest(*openapiclient.NewBuildBetaNotificationCreateRequestData("Type_example", *openapiclient.NewBetaAppReviewSubmissionCreateRequestDataRelationships(*openapiclient.NewBetaAppReviewSubmissionCreateRequestDataRelationshipsBuild(*openapiclient.NewAppStoreVersionRelationshipsBuildData("Type_example", "Id_example"))))) // BuildBetaNotificationCreateRequest | BuildBetaNotification representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BuildBetaNotificationsApi.BuildBetaNotificationsCreateInstance(context.Background()).BuildBetaNotificationCreateRequest(buildBetaNotificationCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildBetaNotificationsApi.BuildBetaNotificationsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildBetaNotificationsCreateInstance`: BuildBetaNotificationResponse
    fmt.Fprintf(os.Stdout, "Response from `BuildBetaNotificationsApi.BuildBetaNotificationsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBuildBetaNotificationsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buildBetaNotificationCreateRequest** | [**BuildBetaNotificationCreateRequest**](BuildBetaNotificationCreateRequest.md) | BuildBetaNotification representation | 

### Return type

[**BuildBetaNotificationResponse**](BuildBetaNotificationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

