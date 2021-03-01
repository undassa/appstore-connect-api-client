# BetaGroupsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BetaGroup**](BetaGroup.md) |  | 
**Included** | Pointer to [**[]OneOfAppBuildBetaTester**](OneOfAppBuildBetaTester.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewBetaGroupsResponse

`func NewBetaGroupsResponse(data []BetaGroup, links PagedDocumentLinks, ) *BetaGroupsResponse`

NewBetaGroupsResponse instantiates a new BetaGroupsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaGroupsResponseWithDefaults

`func NewBetaGroupsResponseWithDefaults() *BetaGroupsResponse`

NewBetaGroupsResponseWithDefaults instantiates a new BetaGroupsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BetaGroupsResponse) GetData() []BetaGroup`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BetaGroupsResponse) GetDataOk() (*[]BetaGroup, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BetaGroupsResponse) SetData(v []BetaGroup)`

SetData sets Data field to given value.


### GetIncluded

`func (o *BetaGroupsResponse) GetIncluded() []OneOfAppBuildBetaTester`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *BetaGroupsResponse) GetIncludedOk() (*[]OneOfAppBuildBetaTester, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *BetaGroupsResponse) SetIncluded(v []OneOfAppBuildBetaTester)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *BetaGroupsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *BetaGroupsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BetaGroupsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BetaGroupsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *BetaGroupsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BetaGroupsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BetaGroupsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BetaGroupsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


