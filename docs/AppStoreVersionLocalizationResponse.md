# AppStoreVersionLocalizationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AppStoreVersionLocalization**](AppStoreVersionLocalization.md) |  | 
**Included** | Pointer to [**[]OneOfAppScreenshotSetAppPreviewSet**](OneOfAppScreenshotSetAppPreviewSet.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewAppStoreVersionLocalizationResponse

`func NewAppStoreVersionLocalizationResponse(data AppStoreVersionLocalization, links DocumentLinks, ) *AppStoreVersionLocalizationResponse`

NewAppStoreVersionLocalizationResponse instantiates a new AppStoreVersionLocalizationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStoreVersionLocalizationResponseWithDefaults

`func NewAppStoreVersionLocalizationResponseWithDefaults() *AppStoreVersionLocalizationResponse`

NewAppStoreVersionLocalizationResponseWithDefaults instantiates a new AppStoreVersionLocalizationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppStoreVersionLocalizationResponse) GetData() AppStoreVersionLocalization`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppStoreVersionLocalizationResponse) GetDataOk() (*AppStoreVersionLocalization, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppStoreVersionLocalizationResponse) SetData(v AppStoreVersionLocalization)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AppStoreVersionLocalizationResponse) GetIncluded() []OneOfAppScreenshotSetAppPreviewSet`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppStoreVersionLocalizationResponse) GetIncludedOk() (*[]OneOfAppScreenshotSetAppPreviewSet, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppStoreVersionLocalizationResponse) SetIncluded(v []OneOfAppScreenshotSetAppPreviewSet)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppStoreVersionLocalizationResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *AppStoreVersionLocalizationResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppStoreVersionLocalizationResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppStoreVersionLocalizationResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


