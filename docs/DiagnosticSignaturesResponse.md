# DiagnosticSignaturesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DiagnosticSignature**](DiagnosticSignature.md) |  | 
**Included** | Pointer to [**[]DiagnosticLog**](DiagnosticLog.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewDiagnosticSignaturesResponse

`func NewDiagnosticSignaturesResponse(data []DiagnosticSignature, links PagedDocumentLinks, ) *DiagnosticSignaturesResponse`

NewDiagnosticSignaturesResponse instantiates a new DiagnosticSignaturesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiagnosticSignaturesResponseWithDefaults

`func NewDiagnosticSignaturesResponseWithDefaults() *DiagnosticSignaturesResponse`

NewDiagnosticSignaturesResponseWithDefaults instantiates a new DiagnosticSignaturesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DiagnosticSignaturesResponse) GetData() []DiagnosticSignature`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DiagnosticSignaturesResponse) GetDataOk() (*[]DiagnosticSignature, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DiagnosticSignaturesResponse) SetData(v []DiagnosticSignature)`

SetData sets Data field to given value.


### GetIncluded

`func (o *DiagnosticSignaturesResponse) GetIncluded() []DiagnosticLog`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *DiagnosticSignaturesResponse) GetIncludedOk() (*[]DiagnosticLog, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *DiagnosticSignaturesResponse) SetIncluded(v []DiagnosticLog)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *DiagnosticSignaturesResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *DiagnosticSignaturesResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DiagnosticSignaturesResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DiagnosticSignaturesResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *DiagnosticSignaturesResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DiagnosticSignaturesResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DiagnosticSignaturesResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DiagnosticSignaturesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


