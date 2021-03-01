# AppCategoriesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AppCategory**](AppCategory.md) |  | 
**Included** | Pointer to [**[]OneOfAppCategoryAppCategory**](OneOfAppCategoryAppCategory.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewAppCategoriesResponse

`func NewAppCategoriesResponse(data []AppCategory, links PagedDocumentLinks, ) *AppCategoriesResponse`

NewAppCategoriesResponse instantiates a new AppCategoriesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppCategoriesResponseWithDefaults

`func NewAppCategoriesResponseWithDefaults() *AppCategoriesResponse`

NewAppCategoriesResponseWithDefaults instantiates a new AppCategoriesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppCategoriesResponse) GetData() []AppCategory`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppCategoriesResponse) GetDataOk() (*[]AppCategory, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppCategoriesResponse) SetData(v []AppCategory)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AppCategoriesResponse) GetIncluded() []OneOfAppCategoryAppCategory`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppCategoriesResponse) GetIncludedOk() (*[]OneOfAppCategoryAppCategory, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppCategoriesResponse) SetIncluded(v []OneOfAppCategoryAppCategory)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppCategoriesResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *AppCategoriesResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppCategoriesResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppCategoriesResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *AppCategoriesResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AppCategoriesResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AppCategoriesResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AppCategoriesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


