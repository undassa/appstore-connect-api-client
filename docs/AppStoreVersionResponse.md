# AppStoreVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AppStoreVersion**](AppStoreVersion.md) |  | 
**Included** | Pointer to [**[]OneOfAgeRatingDeclarationAppStoreVersionLocalizationBuildAppStoreVersionPhasedReleaseRoutingAppCoverageAppStoreReviewDetailAppStoreVersionSubmissionIdfaDeclaration**](OneOfAgeRatingDeclarationAppStoreVersionLocalizationBuildAppStoreVersionPhasedReleaseRoutingAppCoverageAppStoreReviewDetailAppStoreVersionSubmissionIdfaDeclaration.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewAppStoreVersionResponse

`func NewAppStoreVersionResponse(data AppStoreVersion, links DocumentLinks, ) *AppStoreVersionResponse`

NewAppStoreVersionResponse instantiates a new AppStoreVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStoreVersionResponseWithDefaults

`func NewAppStoreVersionResponseWithDefaults() *AppStoreVersionResponse`

NewAppStoreVersionResponseWithDefaults instantiates a new AppStoreVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppStoreVersionResponse) GetData() AppStoreVersion`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppStoreVersionResponse) GetDataOk() (*AppStoreVersion, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppStoreVersionResponse) SetData(v AppStoreVersion)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AppStoreVersionResponse) GetIncluded() []OneOfAgeRatingDeclarationAppStoreVersionLocalizationBuildAppStoreVersionPhasedReleaseRoutingAppCoverageAppStoreReviewDetailAppStoreVersionSubmissionIdfaDeclaration`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppStoreVersionResponse) GetIncludedOk() (*[]OneOfAgeRatingDeclarationAppStoreVersionLocalizationBuildAppStoreVersionPhasedReleaseRoutingAppCoverageAppStoreReviewDetailAppStoreVersionSubmissionIdfaDeclaration, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppStoreVersionResponse) SetIncluded(v []OneOfAgeRatingDeclarationAppStoreVersionLocalizationBuildAppStoreVersionPhasedReleaseRoutingAppCoverageAppStoreReviewDetailAppStoreVersionSubmissionIdfaDeclaration)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppStoreVersionResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *AppStoreVersionResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppStoreVersionResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppStoreVersionResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


