# DiagnosticLogsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DiagnosticLog**](DiagnosticLog.md) |  | 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewDiagnosticLogsResponse

`func NewDiagnosticLogsResponse(data []DiagnosticLog, links PagedDocumentLinks, ) *DiagnosticLogsResponse`

NewDiagnosticLogsResponse instantiates a new DiagnosticLogsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiagnosticLogsResponseWithDefaults

`func NewDiagnosticLogsResponseWithDefaults() *DiagnosticLogsResponse`

NewDiagnosticLogsResponseWithDefaults instantiates a new DiagnosticLogsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DiagnosticLogsResponse) GetData() []DiagnosticLog`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DiagnosticLogsResponse) GetDataOk() (*[]DiagnosticLog, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DiagnosticLogsResponse) SetData(v []DiagnosticLog)`

SetData sets Data field to given value.


### GetLinks

`func (o *DiagnosticLogsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DiagnosticLogsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DiagnosticLogsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *DiagnosticLogsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DiagnosticLogsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DiagnosticLogsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DiagnosticLogsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


