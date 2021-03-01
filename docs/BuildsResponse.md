# BuildsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Build**](Build.md) |  | 
**Included** | Pointer to [**[]OneOfPrereleaseVersionBetaTesterBetaBuildLocalizationAppEncryptionDeclarationBetaAppReviewSubmissionAppBuildBetaDetailAppStoreVersionBuildIconPerfPowerMetricDiagnosticSignature**](OneOfPrereleaseVersionBetaTesterBetaBuildLocalizationAppEncryptionDeclarationBetaAppReviewSubmissionAppBuildBetaDetailAppStoreVersionBuildIconPerfPowerMetricDiagnosticSignature.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewBuildsResponse

`func NewBuildsResponse(data []Build, links PagedDocumentLinks, ) *BuildsResponse`

NewBuildsResponse instantiates a new BuildsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildsResponseWithDefaults

`func NewBuildsResponseWithDefaults() *BuildsResponse`

NewBuildsResponseWithDefaults instantiates a new BuildsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BuildsResponse) GetData() []Build`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BuildsResponse) GetDataOk() (*[]Build, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BuildsResponse) SetData(v []Build)`

SetData sets Data field to given value.


### GetIncluded

`func (o *BuildsResponse) GetIncluded() []OneOfPrereleaseVersionBetaTesterBetaBuildLocalizationAppEncryptionDeclarationBetaAppReviewSubmissionAppBuildBetaDetailAppStoreVersionBuildIconPerfPowerMetricDiagnosticSignature`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *BuildsResponse) GetIncludedOk() (*[]OneOfPrereleaseVersionBetaTesterBetaBuildLocalizationAppEncryptionDeclarationBetaAppReviewSubmissionAppBuildBetaDetailAppStoreVersionBuildIconPerfPowerMetricDiagnosticSignature, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *BuildsResponse) SetIncluded(v []OneOfPrereleaseVersionBetaTesterBetaBuildLocalizationAppEncryptionDeclarationBetaAppReviewSubmissionAppBuildBetaDetailAppStoreVersionBuildIconPerfPowerMetricDiagnosticSignature)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *BuildsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *BuildsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BuildsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BuildsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *BuildsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BuildsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BuildsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BuildsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


