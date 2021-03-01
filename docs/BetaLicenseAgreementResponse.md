# BetaLicenseAgreementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**BetaLicenseAgreement**](BetaLicenseAgreement.md) |  | 
**Included** | Pointer to [**[]App**](App.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewBetaLicenseAgreementResponse

`func NewBetaLicenseAgreementResponse(data BetaLicenseAgreement, links DocumentLinks, ) *BetaLicenseAgreementResponse`

NewBetaLicenseAgreementResponse instantiates a new BetaLicenseAgreementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaLicenseAgreementResponseWithDefaults

`func NewBetaLicenseAgreementResponseWithDefaults() *BetaLicenseAgreementResponse`

NewBetaLicenseAgreementResponseWithDefaults instantiates a new BetaLicenseAgreementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BetaLicenseAgreementResponse) GetData() BetaLicenseAgreement`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BetaLicenseAgreementResponse) GetDataOk() (*BetaLicenseAgreement, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BetaLicenseAgreementResponse) SetData(v BetaLicenseAgreement)`

SetData sets Data field to given value.


### GetIncluded

`func (o *BetaLicenseAgreementResponse) GetIncluded() []App`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *BetaLicenseAgreementResponse) GetIncludedOk() (*[]App, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *BetaLicenseAgreementResponse) SetIncluded(v []App)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *BetaLicenseAgreementResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *BetaLicenseAgreementResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BetaLicenseAgreementResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BetaLicenseAgreementResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


