# \DevicesApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesCreateInstance**](DevicesApi.md#DevicesCreateInstance) | **Post** /v1/devices | 
[**DevicesGetCollection**](DevicesApi.md#DevicesGetCollection) | **Get** /v1/devices | 
[**DevicesGetInstance**](DevicesApi.md#DevicesGetInstance) | **Get** /v1/devices/{id} | 
[**DevicesUpdateInstance**](DevicesApi.md#DevicesUpdateInstance) | **Patch** /v1/devices/{id} | 



## DevicesCreateInstance

> DeviceResponse DevicesCreateInstance(ctx).DeviceCreateRequest(deviceCreateRequest).Execute()



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
    deviceCreateRequest := *openapiclient.NewDeviceCreateRequest(*openapiclient.NewDeviceCreateRequestData("Type_example", *openapiclient.NewDeviceCreateRequestDataAttributes("Name_example", openapiclient.BundleIdPlatform("IOS"), "Udid_example"))) // DeviceCreateRequest | Device representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.DevicesCreateInstance(context.Background()).DeviceCreateRequest(deviceCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DevicesCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesCreateInstance`: DeviceResponse
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.DevicesCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDevicesCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceCreateRequest** | [**DeviceCreateRequest**](DeviceCreateRequest.md) | Device representation | 

### Return type

[**DeviceResponse**](DeviceResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesGetCollection

> DevicesResponse DevicesGetCollection(ctx).FilterName(filterName).FilterPlatform(filterPlatform).FilterStatus(filterStatus).FilterUdid(filterUdid).FilterId(filterId).Sort(sort).FieldsDevices(fieldsDevices).Limit(limit).Execute()



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
    filterName := []string{"Inner_example"} // []string | filter by attribute 'name' (optional)
    filterPlatform := []string{"FilterPlatform_example"} // []string | filter by attribute 'platform' (optional)
    filterStatus := []string{"FilterStatus_example"} // []string | filter by attribute 'status' (optional)
    filterUdid := []string{"Inner_example"} // []string | filter by attribute 'udid' (optional)
    filterId := []string{"Inner_example"} // []string | filter by id(s) (optional)
    sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; resources will be sorted as specified (optional)
    fieldsDevices := []string{"FieldsDevices_example"} // []string | the fields to include for returned resources of type devices (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.DevicesGetCollection(context.Background()).FilterName(filterName).FilterPlatform(filterPlatform).FilterStatus(filterStatus).FilterUdid(filterUdid).FilterId(filterId).Sort(sort).FieldsDevices(fieldsDevices).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DevicesGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesGetCollection`: DevicesResponse
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.DevicesGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDevicesGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterName** | **[]string** | filter by attribute &#39;name&#39; | 
 **filterPlatform** | **[]string** | filter by attribute &#39;platform&#39; | 
 **filterStatus** | **[]string** | filter by attribute &#39;status&#39; | 
 **filterUdid** | **[]string** | filter by attribute &#39;udid&#39; | 
 **filterId** | **[]string** | filter by id(s) | 
 **sort** | **[]string** | comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsDevices** | **[]string** | the fields to include for returned resources of type devices | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**DevicesResponse**](DevicesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesGetInstance

> DeviceResponse DevicesGetInstance(ctx, id).FieldsDevices(fieldsDevices).Execute()



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
    fieldsDevices := []string{"FieldsDevices_example"} // []string | the fields to include for returned resources of type devices (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.DevicesGetInstance(context.Background(), id).FieldsDevices(fieldsDevices).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DevicesGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesGetInstance`: DeviceResponse
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.DevicesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsDevices** | **[]string** | the fields to include for returned resources of type devices | 

### Return type

[**DeviceResponse**](DeviceResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesUpdateInstance

> DeviceResponse DevicesUpdateInstance(ctx, id).DeviceUpdateRequest(deviceUpdateRequest).Execute()



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
    deviceUpdateRequest := *openapiclient.NewDeviceUpdateRequest(*openapiclient.NewDeviceUpdateRequestData("Type_example", "Id_example")) // DeviceUpdateRequest | Device representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.DevicesUpdateInstance(context.Background(), id).DeviceUpdateRequest(deviceUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DevicesUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesUpdateInstance`: DeviceResponse
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.DevicesUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceUpdateRequest** | [**DeviceUpdateRequest**](DeviceUpdateRequest.md) | Device representation | 

### Return type

[**DeviceResponse**](DeviceResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

