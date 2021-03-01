# AppInfosResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AppInfo**](AppInfo.md) |  | 
**Included** | Pointer to [**[]OneOfAppInfoLocalizationAppCategoryAppCategoryAppCategoryAppCategoryAppCategoryAppCategory**](OneOfAppInfoLocalizationAppCategoryAppCategoryAppCategoryAppCategoryAppCategoryAppCategory.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewAppInfosResponse

`func NewAppInfosResponse(data []AppInfo, links PagedDocumentLinks, ) *AppInfosResponse`

NewAppInfosResponse instantiates a new AppInfosResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppInfosResponseWithDefaults

`func NewAppInfosResponseWithDefaults() *AppInfosResponse`

NewAppInfosResponseWithDefaults instantiates a new AppInfosResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppInfosResponse) GetData() []AppInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppInfosResponse) GetDataOk() (*[]AppInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppInfosResponse) SetData(v []AppInfo)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AppInfosResponse) GetIncluded() []OneOfAppInfoLocalizationAppCategoryAppCategoryAppCategoryAppCategoryAppCategoryAppCategory`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppInfosResponse) GetIncludedOk() (*[]OneOfAppInfoLocalizationAppCategoryAppCategoryAppCategoryAppCategoryAppCategoryAppCategory, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppInfosResponse) SetIncluded(v []OneOfAppInfoLocalizationAppCategoryAppCategoryAppCategoryAppCategoryAppCategoryAppCategory)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppInfosResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *AppInfosResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppInfosResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppInfosResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *AppInfosResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AppInfosResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AppInfosResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AppInfosResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


