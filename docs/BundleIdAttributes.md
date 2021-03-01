# BundleIdAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to [**BundleIdPlatform**](BundleIdPlatform.md) |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**SeedId** | Pointer to **string** |  | [optional] 

## Methods

### NewBundleIdAttributes

`func NewBundleIdAttributes() *BundleIdAttributes`

NewBundleIdAttributes instantiates a new BundleIdAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleIdAttributesWithDefaults

`func NewBundleIdAttributesWithDefaults() *BundleIdAttributes`

NewBundleIdAttributesWithDefaults instantiates a new BundleIdAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BundleIdAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BundleIdAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BundleIdAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BundleIdAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatform

`func (o *BundleIdAttributes) GetPlatform() BundleIdPlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *BundleIdAttributes) GetPlatformOk() (*BundleIdPlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *BundleIdAttributes) SetPlatform(v BundleIdPlatform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *BundleIdAttributes) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetIdentifier

`func (o *BundleIdAttributes) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *BundleIdAttributes) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *BundleIdAttributes) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *BundleIdAttributes) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetSeedId

`func (o *BundleIdAttributes) GetSeedId() string`

GetSeedId returns the SeedId field if non-nil, zero value otherwise.

### GetSeedIdOk

`func (o *BundleIdAttributes) GetSeedIdOk() (*string, bool)`

GetSeedIdOk returns a tuple with the SeedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeedId

`func (o *BundleIdAttributes) SetSeedId(v string)`

SetSeedId sets SeedId field to given value.

### HasSeedId

`func (o *BundleIdAttributes) HasSeedId() bool`

HasSeedId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


