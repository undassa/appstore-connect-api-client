# BetaAppReviewDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**AppStoreReviewDetailAttributes**](AppStoreReviewDetail_attributes.md) |  | [optional] 
**Relationships** | Pointer to [**AppEncryptionDeclarationRelationships**](AppEncryptionDeclaration_relationships.md) |  | [optional] 
**Links** | [**ResourceLinks**](ResourceLinks.md) |  | 

## Methods

### NewBetaAppReviewDetail

`func NewBetaAppReviewDetail(type_ string, id string, links ResourceLinks, ) *BetaAppReviewDetail`

NewBetaAppReviewDetail instantiates a new BetaAppReviewDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaAppReviewDetailWithDefaults

`func NewBetaAppReviewDetailWithDefaults() *BetaAppReviewDetail`

NewBetaAppReviewDetailWithDefaults instantiates a new BetaAppReviewDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BetaAppReviewDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BetaAppReviewDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BetaAppReviewDetail) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *BetaAppReviewDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BetaAppReviewDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BetaAppReviewDetail) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *BetaAppReviewDetail) GetAttributes() AppStoreReviewDetailAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BetaAppReviewDetail) GetAttributesOk() (*AppStoreReviewDetailAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BetaAppReviewDetail) SetAttributes(v AppStoreReviewDetailAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *BetaAppReviewDetail) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *BetaAppReviewDetail) GetRelationships() AppEncryptionDeclarationRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BetaAppReviewDetail) GetRelationshipsOk() (*AppEncryptionDeclarationRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BetaAppReviewDetail) SetRelationships(v AppEncryptionDeclarationRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BetaAppReviewDetail) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *BetaAppReviewDetail) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BetaAppReviewDetail) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BetaAppReviewDetail) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


