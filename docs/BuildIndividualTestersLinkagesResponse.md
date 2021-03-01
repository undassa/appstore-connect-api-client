# BuildIndividualTestersLinkagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BetaGroupRelationshipsBetaTestersData**](BetaGroupRelationshipsBetaTestersData.md) |  | 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewBuildIndividualTestersLinkagesResponse

`func NewBuildIndividualTestersLinkagesResponse(data []BetaGroupRelationshipsBetaTestersData, links PagedDocumentLinks, ) *BuildIndividualTestersLinkagesResponse`

NewBuildIndividualTestersLinkagesResponse instantiates a new BuildIndividualTestersLinkagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildIndividualTestersLinkagesResponseWithDefaults

`func NewBuildIndividualTestersLinkagesResponseWithDefaults() *BuildIndividualTestersLinkagesResponse`

NewBuildIndividualTestersLinkagesResponseWithDefaults instantiates a new BuildIndividualTestersLinkagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BuildIndividualTestersLinkagesResponse) GetData() []BetaGroupRelationshipsBetaTestersData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BuildIndividualTestersLinkagesResponse) GetDataOk() (*[]BetaGroupRelationshipsBetaTestersData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BuildIndividualTestersLinkagesResponse) SetData(v []BetaGroupRelationshipsBetaTestersData)`

SetData sets Data field to given value.


### GetLinks

`func (o *BuildIndividualTestersLinkagesResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BuildIndividualTestersLinkagesResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BuildIndividualTestersLinkagesResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *BuildIndividualTestersLinkagesResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BuildIndividualTestersLinkagesResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BuildIndividualTestersLinkagesResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BuildIndividualTestersLinkagesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


