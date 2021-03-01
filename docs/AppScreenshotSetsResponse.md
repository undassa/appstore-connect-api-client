# AppScreenshotSetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AppScreenshotSet**](AppScreenshotSet.md) |  | 
**Included** | Pointer to [**[]AppScreenshot**](AppScreenshot.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewAppScreenshotSetsResponse

`func NewAppScreenshotSetsResponse(data []AppScreenshotSet, links PagedDocumentLinks, ) *AppScreenshotSetsResponse`

NewAppScreenshotSetsResponse instantiates a new AppScreenshotSetsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppScreenshotSetsResponseWithDefaults

`func NewAppScreenshotSetsResponseWithDefaults() *AppScreenshotSetsResponse`

NewAppScreenshotSetsResponseWithDefaults instantiates a new AppScreenshotSetsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppScreenshotSetsResponse) GetData() []AppScreenshotSet`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppScreenshotSetsResponse) GetDataOk() (*[]AppScreenshotSet, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppScreenshotSetsResponse) SetData(v []AppScreenshotSet)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AppScreenshotSetsResponse) GetIncluded() []AppScreenshot`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppScreenshotSetsResponse) GetIncludedOk() (*[]AppScreenshot, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppScreenshotSetsResponse) SetIncluded(v []AppScreenshot)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppScreenshotSetsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *AppScreenshotSetsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppScreenshotSetsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppScreenshotSetsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *AppScreenshotSetsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AppScreenshotSetsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AppScreenshotSetsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AppScreenshotSetsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


