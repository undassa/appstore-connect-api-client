# AppResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**App**](App.md) |  | 
**Included** | Pointer to [**[]OneOfBetaGroupAppStoreVersionPrereleaseVersionBetaAppLocalizationBuildBetaLicenseAgreementBetaAppReviewDetailAppInfoEndUserLicenseAgreementAppPreOrderAppPriceTerritoryInAppPurchaseGameCenterEnabledVersionPerfPowerMetric**](OneOfBetaGroupAppStoreVersionPrereleaseVersionBetaAppLocalizationBuildBetaLicenseAgreementBetaAppReviewDetailAppInfoEndUserLicenseAgreementAppPreOrderAppPriceTerritoryInAppPurchaseGameCenterEnabledVersionPerfPowerMetric.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewAppResponse

`func NewAppResponse(data App, links DocumentLinks, ) *AppResponse`

NewAppResponse instantiates a new AppResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppResponseWithDefaults

`func NewAppResponseWithDefaults() *AppResponse`

NewAppResponseWithDefaults instantiates a new AppResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppResponse) GetData() App`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppResponse) GetDataOk() (*App, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppResponse) SetData(v App)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AppResponse) GetIncluded() []OneOfBetaGroupAppStoreVersionPrereleaseVersionBetaAppLocalizationBuildBetaLicenseAgreementBetaAppReviewDetailAppInfoEndUserLicenseAgreementAppPreOrderAppPriceTerritoryInAppPurchaseGameCenterEnabledVersionPerfPowerMetric`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppResponse) GetIncludedOk() (*[]OneOfBetaGroupAppStoreVersionPrereleaseVersionBetaAppLocalizationBuildBetaLicenseAgreementBetaAppReviewDetailAppInfoEndUserLicenseAgreementAppPreOrderAppPriceTerritoryInAppPurchaseGameCenterEnabledVersionPerfPowerMetric, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppResponse) SetIncluded(v []OneOfBetaGroupAppStoreVersionPrereleaseVersionBetaAppLocalizationBuildBetaLicenseAgreementBetaAppReviewDetailAppInfoEndUserLicenseAgreementAppPreOrderAppPriceTerritoryInAppPurchaseGameCenterEnabledVersionPerfPowerMetric)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *AppResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


