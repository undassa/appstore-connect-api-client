# BetaGroupBuildsLinkagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AppStoreVersionRelationshipsBuildData**](AppStoreVersionRelationshipsBuildData.md) |  | 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewBetaGroupBuildsLinkagesResponse

`func NewBetaGroupBuildsLinkagesResponse(data []AppStoreVersionRelationshipsBuildData, links PagedDocumentLinks, ) *BetaGroupBuildsLinkagesResponse`

NewBetaGroupBuildsLinkagesResponse instantiates a new BetaGroupBuildsLinkagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaGroupBuildsLinkagesResponseWithDefaults

`func NewBetaGroupBuildsLinkagesResponseWithDefaults() *BetaGroupBuildsLinkagesResponse`

NewBetaGroupBuildsLinkagesResponseWithDefaults instantiates a new BetaGroupBuildsLinkagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BetaGroupBuildsLinkagesResponse) GetData() []AppStoreVersionRelationshipsBuildData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BetaGroupBuildsLinkagesResponse) GetDataOk() (*[]AppStoreVersionRelationshipsBuildData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BetaGroupBuildsLinkagesResponse) SetData(v []AppStoreVersionRelationshipsBuildData)`

SetData sets Data field to given value.


### GetLinks

`func (o *BetaGroupBuildsLinkagesResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BetaGroupBuildsLinkagesResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BetaGroupBuildsLinkagesResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *BetaGroupBuildsLinkagesResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BetaGroupBuildsLinkagesResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BetaGroupBuildsLinkagesResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BetaGroupBuildsLinkagesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


