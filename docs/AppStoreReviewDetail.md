# AppStoreReviewDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**AppStoreReviewDetailAttributes**](AppStoreReviewDetail_attributes.md) |  | [optional] 
**Relationships** | Pointer to [**AppStoreReviewDetailRelationships**](AppStoreReviewDetail_relationships.md) |  | [optional] 
**Links** | [**ResourceLinks**](ResourceLinks.md) |  | 

## Methods

### NewAppStoreReviewDetail

`func NewAppStoreReviewDetail(type_ string, id string, links ResourceLinks, ) *AppStoreReviewDetail`

NewAppStoreReviewDetail instantiates a new AppStoreReviewDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStoreReviewDetailWithDefaults

`func NewAppStoreReviewDetailWithDefaults() *AppStoreReviewDetail`

NewAppStoreReviewDetailWithDefaults instantiates a new AppStoreReviewDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppStoreReviewDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppStoreReviewDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppStoreReviewDetail) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *AppStoreReviewDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppStoreReviewDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppStoreReviewDetail) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *AppStoreReviewDetail) GetAttributes() AppStoreReviewDetailAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppStoreReviewDetail) GetAttributesOk() (*AppStoreReviewDetailAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppStoreReviewDetail) SetAttributes(v AppStoreReviewDetailAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AppStoreReviewDetail) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *AppStoreReviewDetail) GetRelationships() AppStoreReviewDetailRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AppStoreReviewDetail) GetRelationshipsOk() (*AppStoreReviewDetailRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AppStoreReviewDetail) SetRelationships(v AppStoreReviewDetailRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AppStoreReviewDetail) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *AppStoreReviewDetail) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppStoreReviewDetail) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppStoreReviewDetail) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


