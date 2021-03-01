# AppCategoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AppCategory**](AppCategory.md) |  | 
**Included** | Pointer to [**[]OneOfAppCategoryAppCategory**](OneOfAppCategoryAppCategory.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewAppCategoryResponse

`func NewAppCategoryResponse(data AppCategory, links DocumentLinks, ) *AppCategoryResponse`

NewAppCategoryResponse instantiates a new AppCategoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppCategoryResponseWithDefaults

`func NewAppCategoryResponseWithDefaults() *AppCategoryResponse`

NewAppCategoryResponseWithDefaults instantiates a new AppCategoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppCategoryResponse) GetData() AppCategory`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppCategoryResponse) GetDataOk() (*AppCategory, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppCategoryResponse) SetData(v AppCategory)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AppCategoryResponse) GetIncluded() []OneOfAppCategoryAppCategory`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppCategoryResponse) GetIncludedOk() (*[]OneOfAppCategoryAppCategory, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppCategoryResponse) SetIncluded(v []OneOfAppCategoryAppCategory)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppCategoryResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *AppCategoryResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppCategoryResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppCategoryResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


