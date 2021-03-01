# AppStoreReviewDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AppStoreReviewDetail**](AppStoreReviewDetail.md) |  | 
**Included** | Pointer to [**[]AppStoreReviewAttachment**](AppStoreReviewAttachment.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewAppStoreReviewDetailResponse

`func NewAppStoreReviewDetailResponse(data AppStoreReviewDetail, links DocumentLinks, ) *AppStoreReviewDetailResponse`

NewAppStoreReviewDetailResponse instantiates a new AppStoreReviewDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStoreReviewDetailResponseWithDefaults

`func NewAppStoreReviewDetailResponseWithDefaults() *AppStoreReviewDetailResponse`

NewAppStoreReviewDetailResponseWithDefaults instantiates a new AppStoreReviewDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppStoreReviewDetailResponse) GetData() AppStoreReviewDetail`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppStoreReviewDetailResponse) GetDataOk() (*AppStoreReviewDetail, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppStoreReviewDetailResponse) SetData(v AppStoreReviewDetail)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AppStoreReviewDetailResponse) GetIncluded() []AppStoreReviewAttachment`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppStoreReviewDetailResponse) GetIncludedOk() (*[]AppStoreReviewAttachment, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppStoreReviewDetailResponse) SetIncluded(v []AppStoreReviewAttachment)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppStoreReviewDetailResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *AppStoreReviewDetailResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppStoreReviewDetailResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppStoreReviewDetailResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


