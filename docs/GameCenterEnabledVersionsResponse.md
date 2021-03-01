# GameCenterEnabledVersionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GameCenterEnabledVersion**](GameCenterEnabledVersion.md) |  | 
**Included** | Pointer to [**[]GameCenterEnabledVersion**](GameCenterEnabledVersion.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewGameCenterEnabledVersionsResponse

`func NewGameCenterEnabledVersionsResponse(data []GameCenterEnabledVersion, links PagedDocumentLinks, ) *GameCenterEnabledVersionsResponse`

NewGameCenterEnabledVersionsResponse instantiates a new GameCenterEnabledVersionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterEnabledVersionsResponseWithDefaults

`func NewGameCenterEnabledVersionsResponseWithDefaults() *GameCenterEnabledVersionsResponse`

NewGameCenterEnabledVersionsResponseWithDefaults instantiates a new GameCenterEnabledVersionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterEnabledVersionsResponse) GetData() []GameCenterEnabledVersion`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterEnabledVersionsResponse) GetDataOk() (*[]GameCenterEnabledVersion, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterEnabledVersionsResponse) SetData(v []GameCenterEnabledVersion)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GameCenterEnabledVersionsResponse) GetIncluded() []GameCenterEnabledVersion`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GameCenterEnabledVersionsResponse) GetIncludedOk() (*[]GameCenterEnabledVersion, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GameCenterEnabledVersionsResponse) SetIncluded(v []GameCenterEnabledVersion)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GameCenterEnabledVersionsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterEnabledVersionsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterEnabledVersionsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterEnabledVersionsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *GameCenterEnabledVersionsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GameCenterEnabledVersionsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GameCenterEnabledVersionsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GameCenterEnabledVersionsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


