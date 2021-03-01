# BetaAppReviewSubmissionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BetaAppReviewSubmission**](BetaAppReviewSubmission.md) |  | 
**Included** | Pointer to [**[]Build**](Build.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewBetaAppReviewSubmissionsResponse

`func NewBetaAppReviewSubmissionsResponse(data []BetaAppReviewSubmission, links PagedDocumentLinks, ) *BetaAppReviewSubmissionsResponse`

NewBetaAppReviewSubmissionsResponse instantiates a new BetaAppReviewSubmissionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaAppReviewSubmissionsResponseWithDefaults

`func NewBetaAppReviewSubmissionsResponseWithDefaults() *BetaAppReviewSubmissionsResponse`

NewBetaAppReviewSubmissionsResponseWithDefaults instantiates a new BetaAppReviewSubmissionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BetaAppReviewSubmissionsResponse) GetData() []BetaAppReviewSubmission`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BetaAppReviewSubmissionsResponse) GetDataOk() (*[]BetaAppReviewSubmission, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BetaAppReviewSubmissionsResponse) SetData(v []BetaAppReviewSubmission)`

SetData sets Data field to given value.


### GetIncluded

`func (o *BetaAppReviewSubmissionsResponse) GetIncluded() []Build`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *BetaAppReviewSubmissionsResponse) GetIncludedOk() (*[]Build, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *BetaAppReviewSubmissionsResponse) SetIncluded(v []Build)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *BetaAppReviewSubmissionsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *BetaAppReviewSubmissionsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BetaAppReviewSubmissionsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BetaAppReviewSubmissionsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *BetaAppReviewSubmissionsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BetaAppReviewSubmissionsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BetaAppReviewSubmissionsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BetaAppReviewSubmissionsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


