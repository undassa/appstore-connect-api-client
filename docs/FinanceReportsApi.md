# \FinanceReportsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FinanceReportsGetCollection**](FinanceReportsApi.md#FinanceReportsGetCollection) | **Get** /v1/financeReports | 



## FinanceReportsGetCollection

> *os.File FinanceReportsGetCollection(ctx).FilterRegionCode(filterRegionCode).FilterReportDate(filterReportDate).FilterReportType(filterReportType).FilterVendorNumber(filterVendorNumber).Execute()



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
    filterRegionCode := []string{"Inner_example"} // []string | filter by attribute 'regionCode'
    filterReportDate := []string{"Inner_example"} // []string | filter by attribute 'reportDate'
    filterReportType := []string{"FilterReportType_example"} // []string | filter by attribute 'reportType'
    filterVendorNumber := []string{"Inner_example"} // []string | filter by attribute 'vendorNumber'

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FinanceReportsApi.FinanceReportsGetCollection(context.Background()).FilterRegionCode(filterRegionCode).FilterReportDate(filterReportDate).FilterReportType(filterReportType).FilterVendorNumber(filterVendorNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FinanceReportsApi.FinanceReportsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FinanceReportsGetCollection`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `FinanceReportsApi.FinanceReportsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFinanceReportsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterRegionCode** | **[]string** | filter by attribute &#39;regionCode&#39; | 
 **filterReportDate** | **[]string** | filter by attribute &#39;reportDate&#39; | 
 **filterReportType** | **[]string** | filter by attribute &#39;reportType&#39; | 
 **filterVendorNumber** | **[]string** | filter by attribute &#39;vendorNumber&#39; | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

