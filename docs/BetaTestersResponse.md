# BetaTestersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BetaTester**](BetaTester.md) |  | 
**Included** | Pointer to [**[]OneOfAppBetaGroupBuild**](OneOfAppBetaGroupBuild.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewBetaTestersResponse

`func NewBetaTestersResponse(data []BetaTester, links PagedDocumentLinks, ) *BetaTestersResponse`

NewBetaTestersResponse instantiates a new BetaTestersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaTestersResponseWithDefaults

`func NewBetaTestersResponseWithDefaults() *BetaTestersResponse`

NewBetaTestersResponseWithDefaults instantiates a new BetaTestersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BetaTestersResponse) GetData() []BetaTester`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BetaTestersResponse) GetDataOk() (*[]BetaTester, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BetaTestersResponse) SetData(v []BetaTester)`

SetData sets Data field to given value.


### GetIncluded

`func (o *BetaTestersResponse) GetIncluded() []OneOfAppBetaGroupBuild`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *BetaTestersResponse) GetIncludedOk() (*[]OneOfAppBetaGroupBuild, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *BetaTestersResponse) SetIncluded(v []OneOfAppBetaGroupBuild)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *BetaTestersResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *BetaTestersResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BetaTestersResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BetaTestersResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *BetaTestersResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BetaTestersResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BetaTestersResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BetaTestersResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


