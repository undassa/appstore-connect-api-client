# AppStoreVersionLocalizationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AppStoreVersionLocalization**](AppStoreVersionLocalization.md) |  | 
**Included** | Pointer to [**[]OneOfAppScreenshotSetAppPreviewSet**](OneOfAppScreenshotSetAppPreviewSet.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewAppStoreVersionLocalizationsResponse

`func NewAppStoreVersionLocalizationsResponse(data []AppStoreVersionLocalization, links PagedDocumentLinks, ) *AppStoreVersionLocalizationsResponse`

NewAppStoreVersionLocalizationsResponse instantiates a new AppStoreVersionLocalizationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStoreVersionLocalizationsResponseWithDefaults

`func NewAppStoreVersionLocalizationsResponseWithDefaults() *AppStoreVersionLocalizationsResponse`

NewAppStoreVersionLocalizationsResponseWithDefaults instantiates a new AppStoreVersionLocalizationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppStoreVersionLocalizationsResponse) GetData() []AppStoreVersionLocalization`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppStoreVersionLocalizationsResponse) GetDataOk() (*[]AppStoreVersionLocalization, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppStoreVersionLocalizationsResponse) SetData(v []AppStoreVersionLocalization)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AppStoreVersionLocalizationsResponse) GetIncluded() []OneOfAppScreenshotSetAppPreviewSet`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppStoreVersionLocalizationsResponse) GetIncludedOk() (*[]OneOfAppScreenshotSetAppPreviewSet, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppStoreVersionLocalizationsResponse) SetIncluded(v []OneOfAppScreenshotSetAppPreviewSet)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppStoreVersionLocalizationsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *AppStoreVersionLocalizationsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppStoreVersionLocalizationsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppStoreVersionLocalizationsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *AppStoreVersionLocalizationsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AppStoreVersionLocalizationsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AppStoreVersionLocalizationsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AppStoreVersionLocalizationsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


