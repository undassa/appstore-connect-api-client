# AppStoreVersionAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | Pointer to [**Platform**](Platform.md) |  | [optional] 
**VersionString** | Pointer to **string** |  | [optional] 
**AppStoreState** | Pointer to [**AppStoreVersionState**](AppStoreVersionState.md) |  | [optional] 
**Copyright** | Pointer to **string** |  | [optional] 
**ReleaseType** | Pointer to **string** |  | [optional] 
**EarliestReleaseDate** | Pointer to **time.Time** |  | [optional] 
**UsesIdfa** | Pointer to **bool** |  | [optional] 
**Downloadable** | Pointer to **bool** |  | [optional] 
**CreatedDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAppStoreVersionAttributes

`func NewAppStoreVersionAttributes() *AppStoreVersionAttributes`

NewAppStoreVersionAttributes instantiates a new AppStoreVersionAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStoreVersionAttributesWithDefaults

`func NewAppStoreVersionAttributesWithDefaults() *AppStoreVersionAttributes`

NewAppStoreVersionAttributesWithDefaults instantiates a new AppStoreVersionAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *AppStoreVersionAttributes) GetPlatform() Platform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AppStoreVersionAttributes) GetPlatformOk() (*Platform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AppStoreVersionAttributes) SetPlatform(v Platform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *AppStoreVersionAttributes) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetVersionString

`func (o *AppStoreVersionAttributes) GetVersionString() string`

GetVersionString returns the VersionString field if non-nil, zero value otherwise.

### GetVersionStringOk

`func (o *AppStoreVersionAttributes) GetVersionStringOk() (*string, bool)`

GetVersionStringOk returns a tuple with the VersionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionString

`func (o *AppStoreVersionAttributes) SetVersionString(v string)`

SetVersionString sets VersionString field to given value.

### HasVersionString

`func (o *AppStoreVersionAttributes) HasVersionString() bool`

HasVersionString returns a boolean if a field has been set.

### GetAppStoreState

`func (o *AppStoreVersionAttributes) GetAppStoreState() AppStoreVersionState`

GetAppStoreState returns the AppStoreState field if non-nil, zero value otherwise.

### GetAppStoreStateOk

`func (o *AppStoreVersionAttributes) GetAppStoreStateOk() (*AppStoreVersionState, bool)`

GetAppStoreStateOk returns a tuple with the AppStoreState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppStoreState

`func (o *AppStoreVersionAttributes) SetAppStoreState(v AppStoreVersionState)`

SetAppStoreState sets AppStoreState field to given value.

### HasAppStoreState

`func (o *AppStoreVersionAttributes) HasAppStoreState() bool`

HasAppStoreState returns a boolean if a field has been set.

### GetCopyright

`func (o *AppStoreVersionAttributes) GetCopyright() string`

GetCopyright returns the Copyright field if non-nil, zero value otherwise.

### GetCopyrightOk

`func (o *AppStoreVersionAttributes) GetCopyrightOk() (*string, bool)`

GetCopyrightOk returns a tuple with the Copyright field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyright

`func (o *AppStoreVersionAttributes) SetCopyright(v string)`

SetCopyright sets Copyright field to given value.

### HasCopyright

`func (o *AppStoreVersionAttributes) HasCopyright() bool`

HasCopyright returns a boolean if a field has been set.

### GetReleaseType

`func (o *AppStoreVersionAttributes) GetReleaseType() string`

GetReleaseType returns the ReleaseType field if non-nil, zero value otherwise.

### GetReleaseTypeOk

`func (o *AppStoreVersionAttributes) GetReleaseTypeOk() (*string, bool)`

GetReleaseTypeOk returns a tuple with the ReleaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseType

`func (o *AppStoreVersionAttributes) SetReleaseType(v string)`

SetReleaseType sets ReleaseType field to given value.

### HasReleaseType

`func (o *AppStoreVersionAttributes) HasReleaseType() bool`

HasReleaseType returns a boolean if a field has been set.

### GetEarliestReleaseDate

`func (o *AppStoreVersionAttributes) GetEarliestReleaseDate() time.Time`

GetEarliestReleaseDate returns the EarliestReleaseDate field if non-nil, zero value otherwise.

### GetEarliestReleaseDateOk

`func (o *AppStoreVersionAttributes) GetEarliestReleaseDateOk() (*time.Time, bool)`

GetEarliestReleaseDateOk returns a tuple with the EarliestReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestReleaseDate

`func (o *AppStoreVersionAttributes) SetEarliestReleaseDate(v time.Time)`

SetEarliestReleaseDate sets EarliestReleaseDate field to given value.

### HasEarliestReleaseDate

`func (o *AppStoreVersionAttributes) HasEarliestReleaseDate() bool`

HasEarliestReleaseDate returns a boolean if a field has been set.

### GetUsesIdfa

`func (o *AppStoreVersionAttributes) GetUsesIdfa() bool`

GetUsesIdfa returns the UsesIdfa field if non-nil, zero value otherwise.

### GetUsesIdfaOk

`func (o *AppStoreVersionAttributes) GetUsesIdfaOk() (*bool, bool)`

GetUsesIdfaOk returns a tuple with the UsesIdfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesIdfa

`func (o *AppStoreVersionAttributes) SetUsesIdfa(v bool)`

SetUsesIdfa sets UsesIdfa field to given value.

### HasUsesIdfa

`func (o *AppStoreVersionAttributes) HasUsesIdfa() bool`

HasUsesIdfa returns a boolean if a field has been set.

### GetDownloadable

`func (o *AppStoreVersionAttributes) GetDownloadable() bool`

GetDownloadable returns the Downloadable field if non-nil, zero value otherwise.

### GetDownloadableOk

`func (o *AppStoreVersionAttributes) GetDownloadableOk() (*bool, bool)`

GetDownloadableOk returns a tuple with the Downloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadable

`func (o *AppStoreVersionAttributes) SetDownloadable(v bool)`

SetDownloadable sets Downloadable field to given value.

### HasDownloadable

`func (o *AppStoreVersionAttributes) HasDownloadable() bool`

HasDownloadable returns a boolean if a field has been set.

### GetCreatedDate

`func (o *AppStoreVersionAttributes) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *AppStoreVersionAttributes) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *AppStoreVersionAttributes) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *AppStoreVersionAttributes) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


