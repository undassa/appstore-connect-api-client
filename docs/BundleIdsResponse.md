# BundleIdsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BundleId**](BundleId.md) |  | 
**Included** | Pointer to [**[]OneOfProfileBundleIdCapabilityApp**](OneOfProfileBundleIdCapabilityApp.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewBundleIdsResponse

`func NewBundleIdsResponse(data []BundleId, links PagedDocumentLinks, ) *BundleIdsResponse`

NewBundleIdsResponse instantiates a new BundleIdsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleIdsResponseWithDefaults

`func NewBundleIdsResponseWithDefaults() *BundleIdsResponse`

NewBundleIdsResponseWithDefaults instantiates a new BundleIdsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BundleIdsResponse) GetData() []BundleId`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BundleIdsResponse) GetDataOk() (*[]BundleId, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BundleIdsResponse) SetData(v []BundleId)`

SetData sets Data field to given value.


### GetIncluded

`func (o *BundleIdsResponse) GetIncluded() []OneOfProfileBundleIdCapabilityApp`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *BundleIdsResponse) GetIncludedOk() (*[]OneOfProfileBundleIdCapabilityApp, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *BundleIdsResponse) SetIncluded(v []OneOfProfileBundleIdCapabilityApp)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *BundleIdsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *BundleIdsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BundleIdsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BundleIdsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *BundleIdsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BundleIdsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BundleIdsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BundleIdsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


