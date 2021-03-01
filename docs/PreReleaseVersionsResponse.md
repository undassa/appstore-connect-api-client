# PreReleaseVersionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]PrereleaseVersion**](PrereleaseVersion.md) |  | 
**Included** | Pointer to [**[]OneOfBuildApp**](OneOfBuildApp.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewPreReleaseVersionsResponse

`func NewPreReleaseVersionsResponse(data []PrereleaseVersion, links PagedDocumentLinks, ) *PreReleaseVersionsResponse`

NewPreReleaseVersionsResponse instantiates a new PreReleaseVersionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreReleaseVersionsResponseWithDefaults

`func NewPreReleaseVersionsResponseWithDefaults() *PreReleaseVersionsResponse`

NewPreReleaseVersionsResponseWithDefaults instantiates a new PreReleaseVersionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PreReleaseVersionsResponse) GetData() []PrereleaseVersion`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PreReleaseVersionsResponse) GetDataOk() (*[]PrereleaseVersion, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PreReleaseVersionsResponse) SetData(v []PrereleaseVersion)`

SetData sets Data field to given value.


### GetIncluded

`func (o *PreReleaseVersionsResponse) GetIncluded() []OneOfBuildApp`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *PreReleaseVersionsResponse) GetIncludedOk() (*[]OneOfBuildApp, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *PreReleaseVersionsResponse) SetIncluded(v []OneOfBuildApp)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *PreReleaseVersionsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *PreReleaseVersionsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PreReleaseVersionsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PreReleaseVersionsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *PreReleaseVersionsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PreReleaseVersionsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PreReleaseVersionsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PreReleaseVersionsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


