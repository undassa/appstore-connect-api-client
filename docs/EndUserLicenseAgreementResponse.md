# EndUserLicenseAgreementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**EndUserLicenseAgreement**](EndUserLicenseAgreement.md) |  | 
**Included** | Pointer to [**[]Territory**](Territory.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewEndUserLicenseAgreementResponse

`func NewEndUserLicenseAgreementResponse(data EndUserLicenseAgreement, links DocumentLinks, ) *EndUserLicenseAgreementResponse`

NewEndUserLicenseAgreementResponse instantiates a new EndUserLicenseAgreementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndUserLicenseAgreementResponseWithDefaults

`func NewEndUserLicenseAgreementResponseWithDefaults() *EndUserLicenseAgreementResponse`

NewEndUserLicenseAgreementResponseWithDefaults instantiates a new EndUserLicenseAgreementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EndUserLicenseAgreementResponse) GetData() EndUserLicenseAgreement`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EndUserLicenseAgreementResponse) GetDataOk() (*EndUserLicenseAgreement, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EndUserLicenseAgreementResponse) SetData(v EndUserLicenseAgreement)`

SetData sets Data field to given value.


### GetIncluded

`func (o *EndUserLicenseAgreementResponse) GetIncluded() []Territory`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *EndUserLicenseAgreementResponse) GetIncludedOk() (*[]Territory, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *EndUserLicenseAgreementResponse) SetIncluded(v []Territory)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *EndUserLicenseAgreementResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *EndUserLicenseAgreementResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EndUserLicenseAgreementResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EndUserLicenseAgreementResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


