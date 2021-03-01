# ProfileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**Profile**](Profile.md) |  | 
**Included** | Pointer to [**[]OneOfBundleIdDeviceCertificate**](OneOfBundleIdDeviceCertificate.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewProfileResponse

`func NewProfileResponse(data Profile, links DocumentLinks, ) *ProfileResponse`

NewProfileResponse instantiates a new ProfileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileResponseWithDefaults

`func NewProfileResponseWithDefaults() *ProfileResponse`

NewProfileResponseWithDefaults instantiates a new ProfileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ProfileResponse) GetData() Profile`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProfileResponse) GetDataOk() (*Profile, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProfileResponse) SetData(v Profile)`

SetData sets Data field to given value.


### GetIncluded

`func (o *ProfileResponse) GetIncluded() []OneOfBundleIdDeviceCertificate`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *ProfileResponse) GetIncludedOk() (*[]OneOfBundleIdDeviceCertificate, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *ProfileResponse) SetIncluded(v []OneOfBundleIdDeviceCertificate)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *ProfileResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *ProfileResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProfileResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProfileResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


