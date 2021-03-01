# BetaLicenseAgreementsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BetaLicenseAgreement**](BetaLicenseAgreement.md) |  | 
**Included** | Pointer to [**[]App**](App.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewBetaLicenseAgreementsResponse

`func NewBetaLicenseAgreementsResponse(data []BetaLicenseAgreement, links PagedDocumentLinks, ) *BetaLicenseAgreementsResponse`

NewBetaLicenseAgreementsResponse instantiates a new BetaLicenseAgreementsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaLicenseAgreementsResponseWithDefaults

`func NewBetaLicenseAgreementsResponseWithDefaults() *BetaLicenseAgreementsResponse`

NewBetaLicenseAgreementsResponseWithDefaults instantiates a new BetaLicenseAgreementsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BetaLicenseAgreementsResponse) GetData() []BetaLicenseAgreement`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BetaLicenseAgreementsResponse) GetDataOk() (*[]BetaLicenseAgreement, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BetaLicenseAgreementsResponse) SetData(v []BetaLicenseAgreement)`

SetData sets Data field to given value.


### GetIncluded

`func (o *BetaLicenseAgreementsResponse) GetIncluded() []App`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *BetaLicenseAgreementsResponse) GetIncludedOk() (*[]App, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *BetaLicenseAgreementsResponse) SetIncluded(v []App)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *BetaLicenseAgreementsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *BetaLicenseAgreementsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BetaLicenseAgreementsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BetaLicenseAgreementsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *BetaLicenseAgreementsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BetaLicenseAgreementsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BetaLicenseAgreementsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BetaLicenseAgreementsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


