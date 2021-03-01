# AppStoreVersionCreateRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | [**Platform**](Platform.md) |  | 
**VersionString** | **string** |  | 
**Copyright** | Pointer to **string** |  | [optional] 
**ReleaseType** | Pointer to **string** |  | [optional] 
**EarliestReleaseDate** | Pointer to **time.Time** |  | [optional] 
**UsesIdfa** | Pointer to **bool** |  | [optional] 

## Methods

### NewAppStoreVersionCreateRequestDataAttributes

`func NewAppStoreVersionCreateRequestDataAttributes(platform Platform, versionString string, ) *AppStoreVersionCreateRequestDataAttributes`

NewAppStoreVersionCreateRequestDataAttributes instantiates a new AppStoreVersionCreateRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStoreVersionCreateRequestDataAttributesWithDefaults

`func NewAppStoreVersionCreateRequestDataAttributesWithDefaults() *AppStoreVersionCreateRequestDataAttributes`

NewAppStoreVersionCreateRequestDataAttributesWithDefaults instantiates a new AppStoreVersionCreateRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *AppStoreVersionCreateRequestDataAttributes) GetPlatform() Platform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AppStoreVersionCreateRequestDataAttributes) GetPlatformOk() (*Platform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AppStoreVersionCreateRequestDataAttributes) SetPlatform(v Platform)`

SetPlatform sets Platform field to given value.


### GetVersionString

`func (o *AppStoreVersionCreateRequestDataAttributes) GetVersionString() string`

GetVersionString returns the VersionString field if non-nil, zero value otherwise.

### GetVersionStringOk

`func (o *AppStoreVersionCreateRequestDataAttributes) GetVersionStringOk() (*string, bool)`

GetVersionStringOk returns a tuple with the VersionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionString

`func (o *AppStoreVersionCreateRequestDataAttributes) SetVersionString(v string)`

SetVersionString sets VersionString field to given value.


### GetCopyright

`func (o *AppStoreVersionCreateRequestDataAttributes) GetCopyright() string`

GetCopyright returns the Copyright field if non-nil, zero value otherwise.

### GetCopyrightOk

`func (o *AppStoreVersionCreateRequestDataAttributes) GetCopyrightOk() (*string, bool)`

GetCopyrightOk returns a tuple with the Copyright field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyright

`func (o *AppStoreVersionCreateRequestDataAttributes) SetCopyright(v string)`

SetCopyright sets Copyright field to given value.

### HasCopyright

`func (o *AppStoreVersionCreateRequestDataAttributes) HasCopyright() bool`

HasCopyright returns a boolean if a field has been set.

### GetReleaseType

`func (o *AppStoreVersionCreateRequestDataAttributes) GetReleaseType() string`

GetReleaseType returns the ReleaseType field if non-nil, zero value otherwise.

### GetReleaseTypeOk

`func (o *AppStoreVersionCreateRequestDataAttributes) GetReleaseTypeOk() (*string, bool)`

GetReleaseTypeOk returns a tuple with the ReleaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseType

`func (o *AppStoreVersionCreateRequestDataAttributes) SetReleaseType(v string)`

SetReleaseType sets ReleaseType field to given value.

### HasReleaseType

`func (o *AppStoreVersionCreateRequestDataAttributes) HasReleaseType() bool`

HasReleaseType returns a boolean if a field has been set.

### GetEarliestReleaseDate

`func (o *AppStoreVersionCreateRequestDataAttributes) GetEarliestReleaseDate() time.Time`

GetEarliestReleaseDate returns the EarliestReleaseDate field if non-nil, zero value otherwise.

### GetEarliestReleaseDateOk

`func (o *AppStoreVersionCreateRequestDataAttributes) GetEarliestReleaseDateOk() (*time.Time, bool)`

GetEarliestReleaseDateOk returns a tuple with the EarliestReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestReleaseDate

`func (o *AppStoreVersionCreateRequestDataAttributes) SetEarliestReleaseDate(v time.Time)`

SetEarliestReleaseDate sets EarliestReleaseDate field to given value.

### HasEarliestReleaseDate

`func (o *AppStoreVersionCreateRequestDataAttributes) HasEarliestReleaseDate() bool`

HasEarliestReleaseDate returns a boolean if a field has been set.

### GetUsesIdfa

`func (o *AppStoreVersionCreateRequestDataAttributes) GetUsesIdfa() bool`

GetUsesIdfa returns the UsesIdfa field if non-nil, zero value otherwise.

### GetUsesIdfaOk

`func (o *AppStoreVersionCreateRequestDataAttributes) GetUsesIdfaOk() (*bool, bool)`

GetUsesIdfaOk returns a tuple with the UsesIdfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesIdfa

`func (o *AppStoreVersionCreateRequestDataAttributes) SetUsesIdfa(v bool)`

SetUsesIdfa sets UsesIdfa field to given value.

### HasUsesIdfa

`func (o *AppStoreVersionCreateRequestDataAttributes) HasUsesIdfa() bool`

HasUsesIdfa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


