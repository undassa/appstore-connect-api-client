# AppInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AppInfo**](AppInfo.md) |  | 
**Included** | Pointer to [**[]OneOfAppInfoLocalizationAppCategoryAppCategoryAppCategoryAppCategoryAppCategoryAppCategory**](OneOfAppInfoLocalizationAppCategoryAppCategoryAppCategoryAppCategoryAppCategoryAppCategory.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewAppInfoResponse

`func NewAppInfoResponse(data AppInfo, links DocumentLinks, ) *AppInfoResponse`

NewAppInfoResponse instantiates a new AppInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppInfoResponseWithDefaults

`func NewAppInfoResponseWithDefaults() *AppInfoResponse`

NewAppInfoResponseWithDefaults instantiates a new AppInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppInfoResponse) GetData() AppInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppInfoResponse) GetDataOk() (*AppInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppInfoResponse) SetData(v AppInfo)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AppInfoResponse) GetIncluded() []OneOfAppInfoLocalizationAppCategoryAppCategoryAppCategoryAppCategoryAppCategoryAppCategory`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppInfoResponse) GetIncludedOk() (*[]OneOfAppInfoLocalizationAppCategoryAppCategoryAppCategoryAppCategoryAppCategoryAppCategory, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppInfoResponse) SetIncluded(v []OneOfAppInfoLocalizationAppCategoryAppCategoryAppCategoryAppCategoryAppCategoryAppCategory)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppInfoResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *AppInfoResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppInfoResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppInfoResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


