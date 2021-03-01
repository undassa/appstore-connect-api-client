# AppStoreVersionPhasedReleaseCreateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Attributes** | Pointer to [**AppStoreVersionPhasedReleaseCreateRequestDataAttributes**](AppStoreVersionPhasedReleaseCreateRequest_data_attributes.md) |  | [optional] 
**Relationships** | [**AppStoreReviewDetailCreateRequestDataRelationships**](AppStoreReviewDetailCreateRequest_data_relationships.md) |  | 

## Methods

### NewAppStoreVersionPhasedReleaseCreateRequestData

`func NewAppStoreVersionPhasedReleaseCreateRequestData(type_ string, relationships AppStoreReviewDetailCreateRequestDataRelationships, ) *AppStoreVersionPhasedReleaseCreateRequestData`

NewAppStoreVersionPhasedReleaseCreateRequestData instantiates a new AppStoreVersionPhasedReleaseCreateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStoreVersionPhasedReleaseCreateRequestDataWithDefaults

`func NewAppStoreVersionPhasedReleaseCreateRequestDataWithDefaults() *AppStoreVersionPhasedReleaseCreateRequestData`

NewAppStoreVersionPhasedReleaseCreateRequestDataWithDefaults instantiates a new AppStoreVersionPhasedReleaseCreateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppStoreVersionPhasedReleaseCreateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppStoreVersionPhasedReleaseCreateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppStoreVersionPhasedReleaseCreateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *AppStoreVersionPhasedReleaseCreateRequestData) GetAttributes() AppStoreVersionPhasedReleaseCreateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppStoreVersionPhasedReleaseCreateRequestData) GetAttributesOk() (*AppStoreVersionPhasedReleaseCreateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppStoreVersionPhasedReleaseCreateRequestData) SetAttributes(v AppStoreVersionPhasedReleaseCreateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AppStoreVersionPhasedReleaseCreateRequestData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *AppStoreVersionPhasedReleaseCreateRequestData) GetRelationships() AppStoreReviewDetailCreateRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AppStoreVersionPhasedReleaseCreateRequestData) GetRelationshipsOk() (*AppStoreReviewDetailCreateRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AppStoreVersionPhasedReleaseCreateRequestData) SetRelationships(v AppStoreReviewDetailCreateRequestDataRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


