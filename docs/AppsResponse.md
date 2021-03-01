# AppsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]App**](App.md) |  | 
**Included** | Pointer to [**[]OneOfBetaGroupAppStoreVersionPrereleaseVersionBetaAppLocalizationBuildBetaLicenseAgreementBetaAppReviewDetailAppInfoEndUserLicenseAgreementAppPreOrderAppPriceTerritoryInAppPurchaseGameCenterEnabledVersionPerfPowerMetric**](OneOfBetaGroupAppStoreVersionPrereleaseVersionBetaAppLocalizationBuildBetaLicenseAgreementBetaAppReviewDetailAppInfoEndUserLicenseAgreementAppPreOrderAppPriceTerritoryInAppPurchaseGameCenterEnabledVersionPerfPowerMetric.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewAppsResponse

`func NewAppsResponse(data []App, links PagedDocumentLinks, ) *AppsResponse`

NewAppsResponse instantiates a new AppsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppsResponseWithDefaults

`func NewAppsResponseWithDefaults() *AppsResponse`

NewAppsResponseWithDefaults instantiates a new AppsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppsResponse) GetData() []App`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppsResponse) GetDataOk() (*[]App, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppsResponse) SetData(v []App)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AppsResponse) GetIncluded() []OneOfBetaGroupAppStoreVersionPrereleaseVersionBetaAppLocalizationBuildBetaLicenseAgreementBetaAppReviewDetailAppInfoEndUserLicenseAgreementAppPreOrderAppPriceTerritoryInAppPurchaseGameCenterEnabledVersionPerfPowerMetric`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppsResponse) GetIncludedOk() (*[]OneOfBetaGroupAppStoreVersionPrereleaseVersionBetaAppLocalizationBuildBetaLicenseAgreementBetaAppReviewDetailAppInfoEndUserLicenseAgreementAppPreOrderAppPriceTerritoryInAppPurchaseGameCenterEnabledVersionPerfPowerMetric, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppsResponse) SetIncluded(v []OneOfBetaGroupAppStoreVersionPrereleaseVersionBetaAppLocalizationBuildBetaLicenseAgreementBetaAppReviewDetailAppInfoEndUserLicenseAgreementAppPreOrderAppPriceTerritoryInAppPurchaseGameCenterEnabledVersionPerfPowerMetric)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *AppsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *AppsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AppsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AppsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AppsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


