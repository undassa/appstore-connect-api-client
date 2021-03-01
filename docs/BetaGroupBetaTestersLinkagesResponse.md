# BetaGroupBetaTestersLinkagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BetaGroupRelationshipsBetaTestersData**](BetaGroupRelationshipsBetaTestersData.md) |  | 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewBetaGroupBetaTestersLinkagesResponse

`func NewBetaGroupBetaTestersLinkagesResponse(data []BetaGroupRelationshipsBetaTestersData, links PagedDocumentLinks, ) *BetaGroupBetaTestersLinkagesResponse`

NewBetaGroupBetaTestersLinkagesResponse instantiates a new BetaGroupBetaTestersLinkagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaGroupBetaTestersLinkagesResponseWithDefaults

`func NewBetaGroupBetaTestersLinkagesResponseWithDefaults() *BetaGroupBetaTestersLinkagesResponse`

NewBetaGroupBetaTestersLinkagesResponseWithDefaults instantiates a new BetaGroupBetaTestersLinkagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BetaGroupBetaTestersLinkagesResponse) GetData() []BetaGroupRelationshipsBetaTestersData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BetaGroupBetaTestersLinkagesResponse) GetDataOk() (*[]BetaGroupRelationshipsBetaTestersData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BetaGroupBetaTestersLinkagesResponse) SetData(v []BetaGroupRelationshipsBetaTestersData)`

SetData sets Data field to given value.


### GetLinks

`func (o *BetaGroupBetaTestersLinkagesResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BetaGroupBetaTestersLinkagesResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BetaGroupBetaTestersLinkagesResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *BetaGroupBetaTestersLinkagesResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BetaGroupBetaTestersLinkagesResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BetaGroupBetaTestersLinkagesResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BetaGroupBetaTestersLinkagesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


