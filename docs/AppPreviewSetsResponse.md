# AppPreviewSetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AppPreviewSet**](AppPreviewSet.md) |  | 
**Included** | Pointer to [**[]AppPreview**](AppPreview.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewAppPreviewSetsResponse

`func NewAppPreviewSetsResponse(data []AppPreviewSet, links PagedDocumentLinks, ) *AppPreviewSetsResponse`

NewAppPreviewSetsResponse instantiates a new AppPreviewSetsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPreviewSetsResponseWithDefaults

`func NewAppPreviewSetsResponseWithDefaults() *AppPreviewSetsResponse`

NewAppPreviewSetsResponseWithDefaults instantiates a new AppPreviewSetsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppPreviewSetsResponse) GetData() []AppPreviewSet`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppPreviewSetsResponse) GetDataOk() (*[]AppPreviewSet, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppPreviewSetsResponse) SetData(v []AppPreviewSet)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AppPreviewSetsResponse) GetIncluded() []AppPreview`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppPreviewSetsResponse) GetIncludedOk() (*[]AppPreview, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppPreviewSetsResponse) SetIncluded(v []AppPreview)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppPreviewSetsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *AppPreviewSetsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppPreviewSetsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppPreviewSetsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *AppPreviewSetsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AppPreviewSetsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AppPreviewSetsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AppPreviewSetsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


