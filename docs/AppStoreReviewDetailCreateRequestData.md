# AppStoreReviewDetailCreateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Attributes** | Pointer to [**AppStoreReviewDetailAttributes**](AppStoreReviewDetail_attributes.md) |  | [optional] 
**Relationships** | [**AppStoreReviewDetailCreateRequestDataRelationships**](AppStoreReviewDetailCreateRequest_data_relationships.md) |  | 

## Methods

### NewAppStoreReviewDetailCreateRequestData

`func NewAppStoreReviewDetailCreateRequestData(type_ string, relationships AppStoreReviewDetailCreateRequestDataRelationships, ) *AppStoreReviewDetailCreateRequestData`

NewAppStoreReviewDetailCreateRequestData instantiates a new AppStoreReviewDetailCreateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStoreReviewDetailCreateRequestDataWithDefaults

`func NewAppStoreReviewDetailCreateRequestDataWithDefaults() *AppStoreReviewDetailCreateRequestData`

NewAppStoreReviewDetailCreateRequestDataWithDefaults instantiates a new AppStoreReviewDetailCreateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppStoreReviewDetailCreateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppStoreReviewDetailCreateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppStoreReviewDetailCreateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *AppStoreReviewDetailCreateRequestData) GetAttributes() AppStoreReviewDetailAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppStoreReviewDetailCreateRequestData) GetAttributesOk() (*AppStoreReviewDetailAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppStoreReviewDetailCreateRequestData) SetAttributes(v AppStoreReviewDetailAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AppStoreReviewDetailCreateRequestData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *AppStoreReviewDetailCreateRequestData) GetRelationships() AppStoreReviewDetailCreateRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AppStoreReviewDetailCreateRequestData) GetRelationshipsOk() (*AppStoreReviewDetailCreateRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AppStoreReviewDetailCreateRequestData) SetRelationships(v AppStoreReviewDetailCreateRequestDataRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


