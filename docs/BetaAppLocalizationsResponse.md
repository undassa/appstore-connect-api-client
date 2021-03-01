# BetaAppLocalizationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BetaAppLocalization**](BetaAppLocalization.md) |  | 
**Included** | Pointer to [**[]App**](App.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewBetaAppLocalizationsResponse

`func NewBetaAppLocalizationsResponse(data []BetaAppLocalization, links PagedDocumentLinks, ) *BetaAppLocalizationsResponse`

NewBetaAppLocalizationsResponse instantiates a new BetaAppLocalizationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaAppLocalizationsResponseWithDefaults

`func NewBetaAppLocalizationsResponseWithDefaults() *BetaAppLocalizationsResponse`

NewBetaAppLocalizationsResponseWithDefaults instantiates a new BetaAppLocalizationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BetaAppLocalizationsResponse) GetData() []BetaAppLocalization`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BetaAppLocalizationsResponse) GetDataOk() (*[]BetaAppLocalization, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BetaAppLocalizationsResponse) SetData(v []BetaAppLocalization)`

SetData sets Data field to given value.


### GetIncluded

`func (o *BetaAppLocalizationsResponse) GetIncluded() []App`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *BetaAppLocalizationsResponse) GetIncludedOk() (*[]App, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *BetaAppLocalizationsResponse) SetIncluded(v []App)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *BetaAppLocalizationsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *BetaAppLocalizationsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BetaAppLocalizationsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BetaAppLocalizationsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *BetaAppLocalizationsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BetaAppLocalizationsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BetaAppLocalizationsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BetaAppLocalizationsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


