# BetaAppReviewDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**BetaAppReviewDetail**](BetaAppReviewDetail.md) |  | 
**Included** | Pointer to [**[]App**](App.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewBetaAppReviewDetailResponse

`func NewBetaAppReviewDetailResponse(data BetaAppReviewDetail, links DocumentLinks, ) *BetaAppReviewDetailResponse`

NewBetaAppReviewDetailResponse instantiates a new BetaAppReviewDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaAppReviewDetailResponseWithDefaults

`func NewBetaAppReviewDetailResponseWithDefaults() *BetaAppReviewDetailResponse`

NewBetaAppReviewDetailResponseWithDefaults instantiates a new BetaAppReviewDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BetaAppReviewDetailResponse) GetData() BetaAppReviewDetail`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BetaAppReviewDetailResponse) GetDataOk() (*BetaAppReviewDetail, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BetaAppReviewDetailResponse) SetData(v BetaAppReviewDetail)`

SetData sets Data field to given value.


### GetIncluded

`func (o *BetaAppReviewDetailResponse) GetIncluded() []App`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *BetaAppReviewDetailResponse) GetIncludedOk() (*[]App, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *BetaAppReviewDetailResponse) SetIncluded(v []App)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *BetaAppReviewDetailResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *BetaAppReviewDetailResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BetaAppReviewDetailResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BetaAppReviewDetailResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


