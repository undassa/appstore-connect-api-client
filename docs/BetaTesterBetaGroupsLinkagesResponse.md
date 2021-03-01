# BetaTesterBetaGroupsLinkagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AppRelationshipsBetaGroupsData**](AppRelationshipsBetaGroupsData.md) |  | 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewBetaTesterBetaGroupsLinkagesResponse

`func NewBetaTesterBetaGroupsLinkagesResponse(data []AppRelationshipsBetaGroupsData, links PagedDocumentLinks, ) *BetaTesterBetaGroupsLinkagesResponse`

NewBetaTesterBetaGroupsLinkagesResponse instantiates a new BetaTesterBetaGroupsLinkagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaTesterBetaGroupsLinkagesResponseWithDefaults

`func NewBetaTesterBetaGroupsLinkagesResponseWithDefaults() *BetaTesterBetaGroupsLinkagesResponse`

NewBetaTesterBetaGroupsLinkagesResponseWithDefaults instantiates a new BetaTesterBetaGroupsLinkagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BetaTesterBetaGroupsLinkagesResponse) GetData() []AppRelationshipsBetaGroupsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BetaTesterBetaGroupsLinkagesResponse) GetDataOk() (*[]AppRelationshipsBetaGroupsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BetaTesterBetaGroupsLinkagesResponse) SetData(v []AppRelationshipsBetaGroupsData)`

SetData sets Data field to given value.


### GetLinks

`func (o *BetaTesterBetaGroupsLinkagesResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BetaTesterBetaGroupsLinkagesResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BetaTesterBetaGroupsLinkagesResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *BetaTesterBetaGroupsLinkagesResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BetaTesterBetaGroupsLinkagesResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BetaTesterBetaGroupsLinkagesResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BetaTesterBetaGroupsLinkagesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


