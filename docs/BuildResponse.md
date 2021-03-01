# BuildResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**Build**](Build.md) |  | 
**Included** | Pointer to [**[]OneOfPrereleaseVersionBetaTesterBetaBuildLocalizationAppEncryptionDeclarationBetaAppReviewSubmissionAppBuildBetaDetailAppStoreVersionBuildIconPerfPowerMetricDiagnosticSignature**](OneOfPrereleaseVersionBetaTesterBetaBuildLocalizationAppEncryptionDeclarationBetaAppReviewSubmissionAppBuildBetaDetailAppStoreVersionBuildIconPerfPowerMetricDiagnosticSignature.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewBuildResponse

`func NewBuildResponse(data Build, links DocumentLinks, ) *BuildResponse`

NewBuildResponse instantiates a new BuildResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildResponseWithDefaults

`func NewBuildResponseWithDefaults() *BuildResponse`

NewBuildResponseWithDefaults instantiates a new BuildResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BuildResponse) GetData() Build`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BuildResponse) GetDataOk() (*Build, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BuildResponse) SetData(v Build)`

SetData sets Data field to given value.


### GetIncluded

`func (o *BuildResponse) GetIncluded() []OneOfPrereleaseVersionBetaTesterBetaBuildLocalizationAppEncryptionDeclarationBetaAppReviewSubmissionAppBuildBetaDetailAppStoreVersionBuildIconPerfPowerMetricDiagnosticSignature`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *BuildResponse) GetIncludedOk() (*[]OneOfPrereleaseVersionBetaTesterBetaBuildLocalizationAppEncryptionDeclarationBetaAppReviewSubmissionAppBuildBetaDetailAppStoreVersionBuildIconPerfPowerMetricDiagnosticSignature, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *BuildResponse) SetIncluded(v []OneOfPrereleaseVersionBetaTesterBetaBuildLocalizationAppEncryptionDeclarationBetaAppReviewSubmissionAppBuildBetaDetailAppStoreVersionBuildIconPerfPowerMetricDiagnosticSignature)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *BuildResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *BuildResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BuildResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BuildResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


