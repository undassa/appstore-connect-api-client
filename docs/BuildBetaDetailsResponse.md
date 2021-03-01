# BuildBetaDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BuildBetaDetail**](BuildBetaDetail.md) |  | 
**Included** | Pointer to [**[]Build**](Build.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewBuildBetaDetailsResponse

`func NewBuildBetaDetailsResponse(data []BuildBetaDetail, links PagedDocumentLinks, ) *BuildBetaDetailsResponse`

NewBuildBetaDetailsResponse instantiates a new BuildBetaDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildBetaDetailsResponseWithDefaults

`func NewBuildBetaDetailsResponseWithDefaults() *BuildBetaDetailsResponse`

NewBuildBetaDetailsResponseWithDefaults instantiates a new BuildBetaDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BuildBetaDetailsResponse) GetData() []BuildBetaDetail`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BuildBetaDetailsResponse) GetDataOk() (*[]BuildBetaDetail, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BuildBetaDetailsResponse) SetData(v []BuildBetaDetail)`

SetData sets Data field to given value.


### GetIncluded

`func (o *BuildBetaDetailsResponse) GetIncluded() []Build`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *BuildBetaDetailsResponse) GetIncludedOk() (*[]Build, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *BuildBetaDetailsResponse) SetIncluded(v []Build)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *BuildBetaDetailsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *BuildBetaDetailsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BuildBetaDetailsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BuildBetaDetailsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *BuildBetaDetailsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BuildBetaDetailsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BuildBetaDetailsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BuildBetaDetailsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


