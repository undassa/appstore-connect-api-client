# PrereleaseVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**PrereleaseVersion**](PrereleaseVersion.md) |  | 
**Included** | Pointer to [**[]OneOfBuildApp**](OneOfBuildApp.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewPrereleaseVersionResponse

`func NewPrereleaseVersionResponse(data PrereleaseVersion, links DocumentLinks, ) *PrereleaseVersionResponse`

NewPrereleaseVersionResponse instantiates a new PrereleaseVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrereleaseVersionResponseWithDefaults

`func NewPrereleaseVersionResponseWithDefaults() *PrereleaseVersionResponse`

NewPrereleaseVersionResponseWithDefaults instantiates a new PrereleaseVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PrereleaseVersionResponse) GetData() PrereleaseVersion`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PrereleaseVersionResponse) GetDataOk() (*PrereleaseVersion, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PrereleaseVersionResponse) SetData(v PrereleaseVersion)`

SetData sets Data field to given value.


### GetIncluded

`func (o *PrereleaseVersionResponse) GetIncluded() []OneOfBuildApp`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *PrereleaseVersionResponse) GetIncludedOk() (*[]OneOfBuildApp, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *PrereleaseVersionResponse) SetIncluded(v []OneOfBuildApp)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *PrereleaseVersionResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *PrereleaseVersionResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PrereleaseVersionResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PrereleaseVersionResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


