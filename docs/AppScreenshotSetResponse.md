# AppScreenshotSetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AppScreenshotSet**](AppScreenshotSet.md) |  | 
**Included** | Pointer to [**[]AppScreenshot**](AppScreenshot.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewAppScreenshotSetResponse

`func NewAppScreenshotSetResponse(data AppScreenshotSet, links DocumentLinks, ) *AppScreenshotSetResponse`

NewAppScreenshotSetResponse instantiates a new AppScreenshotSetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppScreenshotSetResponseWithDefaults

`func NewAppScreenshotSetResponseWithDefaults() *AppScreenshotSetResponse`

NewAppScreenshotSetResponseWithDefaults instantiates a new AppScreenshotSetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppScreenshotSetResponse) GetData() AppScreenshotSet`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppScreenshotSetResponse) GetDataOk() (*AppScreenshotSet, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppScreenshotSetResponse) SetData(v AppScreenshotSet)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AppScreenshotSetResponse) GetIncluded() []AppScreenshot`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppScreenshotSetResponse) GetIncludedOk() (*[]AppScreenshot, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppScreenshotSetResponse) SetIncluded(v []AppScreenshot)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppScreenshotSetResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *AppScreenshotSetResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppScreenshotSetResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppScreenshotSetResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


