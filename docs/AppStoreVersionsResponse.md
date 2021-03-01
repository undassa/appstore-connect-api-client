# AppStoreVersionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AppStoreVersion**](AppStoreVersion.md) |  | 
**Included** | Pointer to [**[]OneOfAgeRatingDeclarationAppStoreVersionLocalizationBuildAppStoreVersionPhasedReleaseRoutingAppCoverageAppStoreReviewDetailAppStoreVersionSubmissionIdfaDeclaration**](OneOfAgeRatingDeclarationAppStoreVersionLocalizationBuildAppStoreVersionPhasedReleaseRoutingAppCoverageAppStoreReviewDetailAppStoreVersionSubmissionIdfaDeclaration.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewAppStoreVersionsResponse

`func NewAppStoreVersionsResponse(data []AppStoreVersion, links PagedDocumentLinks, ) *AppStoreVersionsResponse`

NewAppStoreVersionsResponse instantiates a new AppStoreVersionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStoreVersionsResponseWithDefaults

`func NewAppStoreVersionsResponseWithDefaults() *AppStoreVersionsResponse`

NewAppStoreVersionsResponseWithDefaults instantiates a new AppStoreVersionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppStoreVersionsResponse) GetData() []AppStoreVersion`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppStoreVersionsResponse) GetDataOk() (*[]AppStoreVersion, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppStoreVersionsResponse) SetData(v []AppStoreVersion)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AppStoreVersionsResponse) GetIncluded() []OneOfAgeRatingDeclarationAppStoreVersionLocalizationBuildAppStoreVersionPhasedReleaseRoutingAppCoverageAppStoreReviewDetailAppStoreVersionSubmissionIdfaDeclaration`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppStoreVersionsResponse) GetIncludedOk() (*[]OneOfAgeRatingDeclarationAppStoreVersionLocalizationBuildAppStoreVersionPhasedReleaseRoutingAppCoverageAppStoreReviewDetailAppStoreVersionSubmissionIdfaDeclaration, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppStoreVersionsResponse) SetIncluded(v []OneOfAgeRatingDeclarationAppStoreVersionLocalizationBuildAppStoreVersionPhasedReleaseRoutingAppCoverageAppStoreReviewDetailAppStoreVersionSubmissionIdfaDeclaration)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppStoreVersionsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *AppStoreVersionsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppStoreVersionsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppStoreVersionsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *AppStoreVersionsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AppStoreVersionsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AppStoreVersionsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AppStoreVersionsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


