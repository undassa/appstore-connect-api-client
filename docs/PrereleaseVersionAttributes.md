# PrereleaseVersionAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to [**Platform**](Platform.md) |  | [optional] 

## Methods

### NewPrereleaseVersionAttributes

`func NewPrereleaseVersionAttributes() *PrereleaseVersionAttributes`

NewPrereleaseVersionAttributes instantiates a new PrereleaseVersionAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrereleaseVersionAttributesWithDefaults

`func NewPrereleaseVersionAttributesWithDefaults() *PrereleaseVersionAttributes`

NewPrereleaseVersionAttributesWithDefaults instantiates a new PrereleaseVersionAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *PrereleaseVersionAttributes) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PrereleaseVersionAttributes) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PrereleaseVersionAttributes) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PrereleaseVersionAttributes) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetPlatform

`func (o *PrereleaseVersionAttributes) GetPlatform() Platform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PrereleaseVersionAttributes) GetPlatformOk() (*Platform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PrereleaseVersionAttributes) SetPlatform(v Platform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *PrereleaseVersionAttributes) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


