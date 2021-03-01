# AppRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BetaGroups** | Pointer to [**AppRelationshipsBetaGroups**](App_relationships_betaGroups.md) |  | [optional] 
**AppStoreVersions** | Pointer to [**AppRelationshipsAppStoreVersions**](App_relationships_appStoreVersions.md) |  | [optional] 
**PreReleaseVersions** | Pointer to [**AppRelationshipsPreReleaseVersions**](App_relationships_preReleaseVersions.md) |  | [optional] 
**BetaAppLocalizations** | Pointer to [**AppRelationshipsBetaAppLocalizations**](App_relationships_betaAppLocalizations.md) |  | [optional] 
**Builds** | Pointer to [**AppRelationshipsBuilds**](App_relationships_builds.md) |  | [optional] 
**BetaLicenseAgreement** | Pointer to [**AppRelationshipsBetaLicenseAgreement**](App_relationships_betaLicenseAgreement.md) |  | [optional] 
**BetaAppReviewDetail** | Pointer to [**AppRelationshipsBetaAppReviewDetail**](App_relationships_betaAppReviewDetail.md) |  | [optional] 
**AppInfos** | Pointer to [**AppRelationshipsAppInfos**](App_relationships_appInfos.md) |  | [optional] 
**EndUserLicenseAgreement** | Pointer to [**AppRelationshipsEndUserLicenseAgreement**](App_relationships_endUserLicenseAgreement.md) |  | [optional] 
**PreOrder** | Pointer to [**AppRelationshipsPreOrder**](App_relationships_preOrder.md) |  | [optional] 
**Prices** | Pointer to [**AppRelationshipsPrices**](App_relationships_prices.md) |  | [optional] 
**AvailableTerritories** | Pointer to [**AppRelationshipsAvailableTerritories**](App_relationships_availableTerritories.md) |  | [optional] 
**InAppPurchases** | Pointer to [**AppRelationshipsInAppPurchases**](App_relationships_inAppPurchases.md) |  | [optional] 
**GameCenterEnabledVersions** | Pointer to [**AppRelationshipsGameCenterEnabledVersions**](App_relationships_gameCenterEnabledVersions.md) |  | [optional] 

## Methods

### NewAppRelationships

`func NewAppRelationships() *AppRelationships`

NewAppRelationships instantiates a new AppRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRelationshipsWithDefaults

`func NewAppRelationshipsWithDefaults() *AppRelationships`

NewAppRelationshipsWithDefaults instantiates a new AppRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBetaGroups

`func (o *AppRelationships) GetBetaGroups() AppRelationshipsBetaGroups`

GetBetaGroups returns the BetaGroups field if non-nil, zero value otherwise.

### GetBetaGroupsOk

`func (o *AppRelationships) GetBetaGroupsOk() (*AppRelationshipsBetaGroups, bool)`

GetBetaGroupsOk returns a tuple with the BetaGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetaGroups

`func (o *AppRelationships) SetBetaGroups(v AppRelationshipsBetaGroups)`

SetBetaGroups sets BetaGroups field to given value.

### HasBetaGroups

`func (o *AppRelationships) HasBetaGroups() bool`

HasBetaGroups returns a boolean if a field has been set.

### GetAppStoreVersions

`func (o *AppRelationships) GetAppStoreVersions() AppRelationshipsAppStoreVersions`

GetAppStoreVersions returns the AppStoreVersions field if non-nil, zero value otherwise.

### GetAppStoreVersionsOk

`func (o *AppRelationships) GetAppStoreVersionsOk() (*AppRelationshipsAppStoreVersions, bool)`

GetAppStoreVersionsOk returns a tuple with the AppStoreVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppStoreVersions

`func (o *AppRelationships) SetAppStoreVersions(v AppRelationshipsAppStoreVersions)`

SetAppStoreVersions sets AppStoreVersions field to given value.

### HasAppStoreVersions

`func (o *AppRelationships) HasAppStoreVersions() bool`

HasAppStoreVersions returns a boolean if a field has been set.

### GetPreReleaseVersions

`func (o *AppRelationships) GetPreReleaseVersions() AppRelationshipsPreReleaseVersions`

GetPreReleaseVersions returns the PreReleaseVersions field if non-nil, zero value otherwise.

### GetPreReleaseVersionsOk

`func (o *AppRelationships) GetPreReleaseVersionsOk() (*AppRelationshipsPreReleaseVersions, bool)`

GetPreReleaseVersionsOk returns a tuple with the PreReleaseVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreReleaseVersions

`func (o *AppRelationships) SetPreReleaseVersions(v AppRelationshipsPreReleaseVersions)`

SetPreReleaseVersions sets PreReleaseVersions field to given value.

### HasPreReleaseVersions

`func (o *AppRelationships) HasPreReleaseVersions() bool`

HasPreReleaseVersions returns a boolean if a field has been set.

### GetBetaAppLocalizations

`func (o *AppRelationships) GetBetaAppLocalizations() AppRelationshipsBetaAppLocalizations`

GetBetaAppLocalizations returns the BetaAppLocalizations field if non-nil, zero value otherwise.

### GetBetaAppLocalizationsOk

`func (o *AppRelationships) GetBetaAppLocalizationsOk() (*AppRelationshipsBetaAppLocalizations, bool)`

GetBetaAppLocalizationsOk returns a tuple with the BetaAppLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetaAppLocalizations

`func (o *AppRelationships) SetBetaAppLocalizations(v AppRelationshipsBetaAppLocalizations)`

SetBetaAppLocalizations sets BetaAppLocalizations field to given value.

### HasBetaAppLocalizations

`func (o *AppRelationships) HasBetaAppLocalizations() bool`

HasBetaAppLocalizations returns a boolean if a field has been set.

### GetBuilds

`func (o *AppRelationships) GetBuilds() AppRelationshipsBuilds`

GetBuilds returns the Builds field if non-nil, zero value otherwise.

### GetBuildsOk

`func (o *AppRelationships) GetBuildsOk() (*AppRelationshipsBuilds, bool)`

GetBuildsOk returns a tuple with the Builds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilds

`func (o *AppRelationships) SetBuilds(v AppRelationshipsBuilds)`

SetBuilds sets Builds field to given value.

### HasBuilds

`func (o *AppRelationships) HasBuilds() bool`

HasBuilds returns a boolean if a field has been set.

### GetBetaLicenseAgreement

`func (o *AppRelationships) GetBetaLicenseAgreement() AppRelationshipsBetaLicenseAgreement`

GetBetaLicenseAgreement returns the BetaLicenseAgreement field if non-nil, zero value otherwise.

### GetBetaLicenseAgreementOk

`func (o *AppRelationships) GetBetaLicenseAgreementOk() (*AppRelationshipsBetaLicenseAgreement, bool)`

GetBetaLicenseAgreementOk returns a tuple with the BetaLicenseAgreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetaLicenseAgreement

`func (o *AppRelationships) SetBetaLicenseAgreement(v AppRelationshipsBetaLicenseAgreement)`

SetBetaLicenseAgreement sets BetaLicenseAgreement field to given value.

### HasBetaLicenseAgreement

`func (o *AppRelationships) HasBetaLicenseAgreement() bool`

HasBetaLicenseAgreement returns a boolean if a field has been set.

### GetBetaAppReviewDetail

`func (o *AppRelationships) GetBetaAppReviewDetail() AppRelationshipsBetaAppReviewDetail`

GetBetaAppReviewDetail returns the BetaAppReviewDetail field if non-nil, zero value otherwise.

### GetBetaAppReviewDetailOk

`func (o *AppRelationships) GetBetaAppReviewDetailOk() (*AppRelationshipsBetaAppReviewDetail, bool)`

GetBetaAppReviewDetailOk returns a tuple with the BetaAppReviewDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetaAppReviewDetail

`func (o *AppRelationships) SetBetaAppReviewDetail(v AppRelationshipsBetaAppReviewDetail)`

SetBetaAppReviewDetail sets BetaAppReviewDetail field to given value.

### HasBetaAppReviewDetail

`func (o *AppRelationships) HasBetaAppReviewDetail() bool`

HasBetaAppReviewDetail returns a boolean if a field has been set.

### GetAppInfos

`func (o *AppRelationships) GetAppInfos() AppRelationshipsAppInfos`

GetAppInfos returns the AppInfos field if non-nil, zero value otherwise.

### GetAppInfosOk

`func (o *AppRelationships) GetAppInfosOk() (*AppRelationshipsAppInfos, bool)`

GetAppInfosOk returns a tuple with the AppInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInfos

`func (o *AppRelationships) SetAppInfos(v AppRelationshipsAppInfos)`

SetAppInfos sets AppInfos field to given value.

### HasAppInfos

`func (o *AppRelationships) HasAppInfos() bool`

HasAppInfos returns a boolean if a field has been set.

### GetEndUserLicenseAgreement

`func (o *AppRelationships) GetEndUserLicenseAgreement() AppRelationshipsEndUserLicenseAgreement`

GetEndUserLicenseAgreement returns the EndUserLicenseAgreement field if non-nil, zero value otherwise.

### GetEndUserLicenseAgreementOk

`func (o *AppRelationships) GetEndUserLicenseAgreementOk() (*AppRelationshipsEndUserLicenseAgreement, bool)`

GetEndUserLicenseAgreementOk returns a tuple with the EndUserLicenseAgreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserLicenseAgreement

`func (o *AppRelationships) SetEndUserLicenseAgreement(v AppRelationshipsEndUserLicenseAgreement)`

SetEndUserLicenseAgreement sets EndUserLicenseAgreement field to given value.

### HasEndUserLicenseAgreement

`func (o *AppRelationships) HasEndUserLicenseAgreement() bool`

HasEndUserLicenseAgreement returns a boolean if a field has been set.

### GetPreOrder

`func (o *AppRelationships) GetPreOrder() AppRelationshipsPreOrder`

GetPreOrder returns the PreOrder field if non-nil, zero value otherwise.

### GetPreOrderOk

`func (o *AppRelationships) GetPreOrderOk() (*AppRelationshipsPreOrder, bool)`

GetPreOrderOk returns a tuple with the PreOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreOrder

`func (o *AppRelationships) SetPreOrder(v AppRelationshipsPreOrder)`

SetPreOrder sets PreOrder field to given value.

### HasPreOrder

`func (o *AppRelationships) HasPreOrder() bool`

HasPreOrder returns a boolean if a field has been set.

### GetPrices

`func (o *AppRelationships) GetPrices() AppRelationshipsPrices`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *AppRelationships) GetPricesOk() (*AppRelationshipsPrices, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *AppRelationships) SetPrices(v AppRelationshipsPrices)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *AppRelationships) HasPrices() bool`

HasPrices returns a boolean if a field has been set.

### GetAvailableTerritories

`func (o *AppRelationships) GetAvailableTerritories() AppRelationshipsAvailableTerritories`

GetAvailableTerritories returns the AvailableTerritories field if non-nil, zero value otherwise.

### GetAvailableTerritoriesOk

`func (o *AppRelationships) GetAvailableTerritoriesOk() (*AppRelationshipsAvailableTerritories, bool)`

GetAvailableTerritoriesOk returns a tuple with the AvailableTerritories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableTerritories

`func (o *AppRelationships) SetAvailableTerritories(v AppRelationshipsAvailableTerritories)`

SetAvailableTerritories sets AvailableTerritories field to given value.

### HasAvailableTerritories

`func (o *AppRelationships) HasAvailableTerritories() bool`

HasAvailableTerritories returns a boolean if a field has been set.

### GetInAppPurchases

`func (o *AppRelationships) GetInAppPurchases() AppRelationshipsInAppPurchases`

GetInAppPurchases returns the InAppPurchases field if non-nil, zero value otherwise.

### GetInAppPurchasesOk

`func (o *AppRelationships) GetInAppPurchasesOk() (*AppRelationshipsInAppPurchases, bool)`

GetInAppPurchasesOk returns a tuple with the InAppPurchases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInAppPurchases

`func (o *AppRelationships) SetInAppPurchases(v AppRelationshipsInAppPurchases)`

SetInAppPurchases sets InAppPurchases field to given value.

### HasInAppPurchases

`func (o *AppRelationships) HasInAppPurchases() bool`

HasInAppPurchases returns a boolean if a field has been set.

### GetGameCenterEnabledVersions

`func (o *AppRelationships) GetGameCenterEnabledVersions() AppRelationshipsGameCenterEnabledVersions`

GetGameCenterEnabledVersions returns the GameCenterEnabledVersions field if non-nil, zero value otherwise.

### GetGameCenterEnabledVersionsOk

`func (o *AppRelationships) GetGameCenterEnabledVersionsOk() (*AppRelationshipsGameCenterEnabledVersions, bool)`

GetGameCenterEnabledVersionsOk returns a tuple with the GameCenterEnabledVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameCenterEnabledVersions

`func (o *AppRelationships) SetGameCenterEnabledVersions(v AppRelationshipsGameCenterEnabledVersions)`

SetGameCenterEnabledVersions sets GameCenterEnabledVersions field to given value.

### HasGameCenterEnabledVersions

`func (o *AppRelationships) HasGameCenterEnabledVersions() bool`

HasGameCenterEnabledVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


