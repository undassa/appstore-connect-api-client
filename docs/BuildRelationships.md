# BuildRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreReleaseVersion** | Pointer to [**BuildRelationshipsPreReleaseVersion**](Build_relationships_preReleaseVersion.md) |  | [optional] 
**IndividualTesters** | Pointer to [**BetaGroupRelationshipsBetaTesters**](BetaGroup_relationships_betaTesters.md) |  | [optional] 
**BetaBuildLocalizations** | Pointer to [**BuildRelationshipsBetaBuildLocalizations**](Build_relationships_betaBuildLocalizations.md) |  | [optional] 
**AppEncryptionDeclaration** | Pointer to [**BuildRelationshipsAppEncryptionDeclaration**](Build_relationships_appEncryptionDeclaration.md) |  | [optional] 
**BetaAppReviewSubmission** | Pointer to [**BuildRelationshipsBetaAppReviewSubmission**](Build_relationships_betaAppReviewSubmission.md) |  | [optional] 
**App** | Pointer to [**AppEncryptionDeclarationRelationshipsApp**](AppEncryptionDeclaration_relationships_app.md) |  | [optional] 
**BuildBetaDetail** | Pointer to [**BuildRelationshipsBuildBetaDetail**](Build_relationships_buildBetaDetail.md) |  | [optional] 
**AppStoreVersion** | Pointer to [**AppStoreReviewDetailRelationshipsAppStoreVersion**](AppStoreReviewDetail_relationships_appStoreVersion.md) |  | [optional] 
**Icons** | Pointer to [**BuildRelationshipsIcons**](Build_relationships_icons.md) |  | [optional] 

## Methods

### NewBuildRelationships

`func NewBuildRelationships() *BuildRelationships`

NewBuildRelationships instantiates a new BuildRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildRelationshipsWithDefaults

`func NewBuildRelationshipsWithDefaults() *BuildRelationships`

NewBuildRelationshipsWithDefaults instantiates a new BuildRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreReleaseVersion

`func (o *BuildRelationships) GetPreReleaseVersion() BuildRelationshipsPreReleaseVersion`

GetPreReleaseVersion returns the PreReleaseVersion field if non-nil, zero value otherwise.

### GetPreReleaseVersionOk

`func (o *BuildRelationships) GetPreReleaseVersionOk() (*BuildRelationshipsPreReleaseVersion, bool)`

GetPreReleaseVersionOk returns a tuple with the PreReleaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreReleaseVersion

`func (o *BuildRelationships) SetPreReleaseVersion(v BuildRelationshipsPreReleaseVersion)`

SetPreReleaseVersion sets PreReleaseVersion field to given value.

### HasPreReleaseVersion

`func (o *BuildRelationships) HasPreReleaseVersion() bool`

HasPreReleaseVersion returns a boolean if a field has been set.

### GetIndividualTesters

`func (o *BuildRelationships) GetIndividualTesters() BetaGroupRelationshipsBetaTesters`

GetIndividualTesters returns the IndividualTesters field if non-nil, zero value otherwise.

### GetIndividualTestersOk

`func (o *BuildRelationships) GetIndividualTestersOk() (*BetaGroupRelationshipsBetaTesters, bool)`

GetIndividualTestersOk returns a tuple with the IndividualTesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualTesters

`func (o *BuildRelationships) SetIndividualTesters(v BetaGroupRelationshipsBetaTesters)`

SetIndividualTesters sets IndividualTesters field to given value.

### HasIndividualTesters

`func (o *BuildRelationships) HasIndividualTesters() bool`

HasIndividualTesters returns a boolean if a field has been set.

### GetBetaBuildLocalizations

`func (o *BuildRelationships) GetBetaBuildLocalizations() BuildRelationshipsBetaBuildLocalizations`

GetBetaBuildLocalizations returns the BetaBuildLocalizations field if non-nil, zero value otherwise.

### GetBetaBuildLocalizationsOk

`func (o *BuildRelationships) GetBetaBuildLocalizationsOk() (*BuildRelationshipsBetaBuildLocalizations, bool)`

GetBetaBuildLocalizationsOk returns a tuple with the BetaBuildLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetaBuildLocalizations

`func (o *BuildRelationships) SetBetaBuildLocalizations(v BuildRelationshipsBetaBuildLocalizations)`

SetBetaBuildLocalizations sets BetaBuildLocalizations field to given value.

### HasBetaBuildLocalizations

`func (o *BuildRelationships) HasBetaBuildLocalizations() bool`

HasBetaBuildLocalizations returns a boolean if a field has been set.

### GetAppEncryptionDeclaration

`func (o *BuildRelationships) GetAppEncryptionDeclaration() BuildRelationshipsAppEncryptionDeclaration`

GetAppEncryptionDeclaration returns the AppEncryptionDeclaration field if non-nil, zero value otherwise.

### GetAppEncryptionDeclarationOk

`func (o *BuildRelationships) GetAppEncryptionDeclarationOk() (*BuildRelationshipsAppEncryptionDeclaration, bool)`

GetAppEncryptionDeclarationOk returns a tuple with the AppEncryptionDeclaration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEncryptionDeclaration

`func (o *BuildRelationships) SetAppEncryptionDeclaration(v BuildRelationshipsAppEncryptionDeclaration)`

SetAppEncryptionDeclaration sets AppEncryptionDeclaration field to given value.

### HasAppEncryptionDeclaration

`func (o *BuildRelationships) HasAppEncryptionDeclaration() bool`

HasAppEncryptionDeclaration returns a boolean if a field has been set.

### GetBetaAppReviewSubmission

`func (o *BuildRelationships) GetBetaAppReviewSubmission() BuildRelationshipsBetaAppReviewSubmission`

GetBetaAppReviewSubmission returns the BetaAppReviewSubmission field if non-nil, zero value otherwise.

### GetBetaAppReviewSubmissionOk

`func (o *BuildRelationships) GetBetaAppReviewSubmissionOk() (*BuildRelationshipsBetaAppReviewSubmission, bool)`

GetBetaAppReviewSubmissionOk returns a tuple with the BetaAppReviewSubmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetaAppReviewSubmission

`func (o *BuildRelationships) SetBetaAppReviewSubmission(v BuildRelationshipsBetaAppReviewSubmission)`

SetBetaAppReviewSubmission sets BetaAppReviewSubmission field to given value.

### HasBetaAppReviewSubmission

`func (o *BuildRelationships) HasBetaAppReviewSubmission() bool`

HasBetaAppReviewSubmission returns a boolean if a field has been set.

### GetApp

`func (o *BuildRelationships) GetApp() AppEncryptionDeclarationRelationshipsApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *BuildRelationships) GetAppOk() (*AppEncryptionDeclarationRelationshipsApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *BuildRelationships) SetApp(v AppEncryptionDeclarationRelationshipsApp)`

SetApp sets App field to given value.

### HasApp

`func (o *BuildRelationships) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetBuildBetaDetail

`func (o *BuildRelationships) GetBuildBetaDetail() BuildRelationshipsBuildBetaDetail`

GetBuildBetaDetail returns the BuildBetaDetail field if non-nil, zero value otherwise.

### GetBuildBetaDetailOk

`func (o *BuildRelationships) GetBuildBetaDetailOk() (*BuildRelationshipsBuildBetaDetail, bool)`

GetBuildBetaDetailOk returns a tuple with the BuildBetaDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildBetaDetail

`func (o *BuildRelationships) SetBuildBetaDetail(v BuildRelationshipsBuildBetaDetail)`

SetBuildBetaDetail sets BuildBetaDetail field to given value.

### HasBuildBetaDetail

`func (o *BuildRelationships) HasBuildBetaDetail() bool`

HasBuildBetaDetail returns a boolean if a field has been set.

### GetAppStoreVersion

`func (o *BuildRelationships) GetAppStoreVersion() AppStoreReviewDetailRelationshipsAppStoreVersion`

GetAppStoreVersion returns the AppStoreVersion field if non-nil, zero value otherwise.

### GetAppStoreVersionOk

`func (o *BuildRelationships) GetAppStoreVersionOk() (*AppStoreReviewDetailRelationshipsAppStoreVersion, bool)`

GetAppStoreVersionOk returns a tuple with the AppStoreVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppStoreVersion

`func (o *BuildRelationships) SetAppStoreVersion(v AppStoreReviewDetailRelationshipsAppStoreVersion)`

SetAppStoreVersion sets AppStoreVersion field to given value.

### HasAppStoreVersion

`func (o *BuildRelationships) HasAppStoreVersion() bool`

HasAppStoreVersion returns a boolean if a field has been set.

### GetIcons

`func (o *BuildRelationships) GetIcons() BuildRelationshipsIcons`

GetIcons returns the Icons field if non-nil, zero value otherwise.

### GetIconsOk

`func (o *BuildRelationships) GetIconsOk() (*BuildRelationshipsIcons, bool)`

GetIconsOk returns a tuple with the Icons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcons

`func (o *BuildRelationships) SetIcons(v BuildRelationshipsIcons)`

SetIcons sets Icons field to given value.

### HasIcons

`func (o *BuildRelationships) HasIcons() bool`

HasIcons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


