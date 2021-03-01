# ProfilesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Profile**](Profile.md) |  | 
**Included** | Pointer to [**[]OneOfBundleIdDeviceCertificate**](OneOfBundleIdDeviceCertificate.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewProfilesResponse

`func NewProfilesResponse(data []Profile, links PagedDocumentLinks, ) *ProfilesResponse`

NewProfilesResponse instantiates a new ProfilesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfilesResponseWithDefaults

`func NewProfilesResponseWithDefaults() *ProfilesResponse`

NewProfilesResponseWithDefaults instantiates a new ProfilesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ProfilesResponse) GetData() []Profile`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProfilesResponse) GetDataOk() (*[]Profile, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProfilesResponse) SetData(v []Profile)`

SetData sets Data field to given value.


### GetIncluded

`func (o *ProfilesResponse) GetIncluded() []OneOfBundleIdDeviceCertificate`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *ProfilesResponse) GetIncludedOk() (*[]OneOfBundleIdDeviceCertificate, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *ProfilesResponse) SetIncluded(v []OneOfBundleIdDeviceCertificate)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *ProfilesResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *ProfilesResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProfilesResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProfilesResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *ProfilesResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ProfilesResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ProfilesResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ProfilesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


