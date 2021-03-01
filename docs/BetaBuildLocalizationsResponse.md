# BetaBuildLocalizationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BetaBuildLocalization**](BetaBuildLocalization.md) |  | 
**Included** | Pointer to [**[]Build**](Build.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewBetaBuildLocalizationsResponse

`func NewBetaBuildLocalizationsResponse(data []BetaBuildLocalization, links PagedDocumentLinks, ) *BetaBuildLocalizationsResponse`

NewBetaBuildLocalizationsResponse instantiates a new BetaBuildLocalizationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaBuildLocalizationsResponseWithDefaults

`func NewBetaBuildLocalizationsResponseWithDefaults() *BetaBuildLocalizationsResponse`

NewBetaBuildLocalizationsResponseWithDefaults instantiates a new BetaBuildLocalizationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BetaBuildLocalizationsResponse) GetData() []BetaBuildLocalization`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BetaBuildLocalizationsResponse) GetDataOk() (*[]BetaBuildLocalization, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BetaBuildLocalizationsResponse) SetData(v []BetaBuildLocalization)`

SetData sets Data field to given value.


### GetIncluded

`func (o *BetaBuildLocalizationsResponse) GetIncluded() []Build`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *BetaBuildLocalizationsResponse) GetIncludedOk() (*[]Build, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *BetaBuildLocalizationsResponse) SetIncluded(v []Build)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *BetaBuildLocalizationsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *BetaBuildLocalizationsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BetaBuildLocalizationsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BetaBuildLocalizationsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *BetaBuildLocalizationsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BetaBuildLocalizationsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BetaBuildLocalizationsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BetaBuildLocalizationsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


