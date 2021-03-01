# BuildAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** |  | [optional] 
**UploadedDate** | Pointer to **time.Time** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**Expired** | Pointer to **bool** |  | [optional] 
**MinOsVersion** | Pointer to **string** |  | [optional] 
**IconAssetToken** | Pointer to [**ImageAsset**](ImageAsset.md) |  | [optional] 
**ProcessingState** | Pointer to **string** |  | [optional] 
**UsesNonExemptEncryption** | Pointer to **bool** |  | [optional] 

## Methods

### NewBuildAttributes

`func NewBuildAttributes() *BuildAttributes`

NewBuildAttributes instantiates a new BuildAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildAttributesWithDefaults

`func NewBuildAttributesWithDefaults() *BuildAttributes`

NewBuildAttributesWithDefaults instantiates a new BuildAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *BuildAttributes) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BuildAttributes) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BuildAttributes) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BuildAttributes) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUploadedDate

`func (o *BuildAttributes) GetUploadedDate() time.Time`

GetUploadedDate returns the UploadedDate field if non-nil, zero value otherwise.

### GetUploadedDateOk

`func (o *BuildAttributes) GetUploadedDateOk() (*time.Time, bool)`

GetUploadedDateOk returns a tuple with the UploadedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedDate

`func (o *BuildAttributes) SetUploadedDate(v time.Time)`

SetUploadedDate sets UploadedDate field to given value.

### HasUploadedDate

`func (o *BuildAttributes) HasUploadedDate() bool`

HasUploadedDate returns a boolean if a field has been set.

### GetExpirationDate

`func (o *BuildAttributes) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *BuildAttributes) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *BuildAttributes) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *BuildAttributes) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetExpired

`func (o *BuildAttributes) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *BuildAttributes) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *BuildAttributes) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *BuildAttributes) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetMinOsVersion

`func (o *BuildAttributes) GetMinOsVersion() string`

GetMinOsVersion returns the MinOsVersion field if non-nil, zero value otherwise.

### GetMinOsVersionOk

`func (o *BuildAttributes) GetMinOsVersionOk() (*string, bool)`

GetMinOsVersionOk returns a tuple with the MinOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOsVersion

`func (o *BuildAttributes) SetMinOsVersion(v string)`

SetMinOsVersion sets MinOsVersion field to given value.

### HasMinOsVersion

`func (o *BuildAttributes) HasMinOsVersion() bool`

HasMinOsVersion returns a boolean if a field has been set.

### GetIconAssetToken

`func (o *BuildAttributes) GetIconAssetToken() ImageAsset`

GetIconAssetToken returns the IconAssetToken field if non-nil, zero value otherwise.

### GetIconAssetTokenOk

`func (o *BuildAttributes) GetIconAssetTokenOk() (*ImageAsset, bool)`

GetIconAssetTokenOk returns a tuple with the IconAssetToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconAssetToken

`func (o *BuildAttributes) SetIconAssetToken(v ImageAsset)`

SetIconAssetToken sets IconAssetToken field to given value.

### HasIconAssetToken

`func (o *BuildAttributes) HasIconAssetToken() bool`

HasIconAssetToken returns a boolean if a field has been set.

### GetProcessingState

`func (o *BuildAttributes) GetProcessingState() string`

GetProcessingState returns the ProcessingState field if non-nil, zero value otherwise.

### GetProcessingStateOk

`func (o *BuildAttributes) GetProcessingStateOk() (*string, bool)`

GetProcessingStateOk returns a tuple with the ProcessingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingState

`func (o *BuildAttributes) SetProcessingState(v string)`

SetProcessingState sets ProcessingState field to given value.

### HasProcessingState

`func (o *BuildAttributes) HasProcessingState() bool`

HasProcessingState returns a boolean if a field has been set.

### GetUsesNonExemptEncryption

`func (o *BuildAttributes) GetUsesNonExemptEncryption() bool`

GetUsesNonExemptEncryption returns the UsesNonExemptEncryption field if non-nil, zero value otherwise.

### GetUsesNonExemptEncryptionOk

`func (o *BuildAttributes) GetUsesNonExemptEncryptionOk() (*bool, bool)`

GetUsesNonExemptEncryptionOk returns a tuple with the UsesNonExemptEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesNonExemptEncryption

`func (o *BuildAttributes) SetUsesNonExemptEncryption(v bool)`

SetUsesNonExemptEncryption sets UsesNonExemptEncryption field to given value.

### HasUsesNonExemptEncryption

`func (o *BuildAttributes) HasUsesNonExemptEncryption() bool`

HasUsesNonExemptEncryption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


