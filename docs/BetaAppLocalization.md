# BetaAppLocalization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**BetaAppLocalizationAttributes**](BetaAppLocalization_attributes.md) |  | [optional] 
**Relationships** | Pointer to [**AppEncryptionDeclarationRelationships**](AppEncryptionDeclaration_relationships.md) |  | [optional] 
**Links** | [**ResourceLinks**](ResourceLinks.md) |  | 

## Methods

### NewBetaAppLocalization

`func NewBetaAppLocalization(type_ string, id string, links ResourceLinks, ) *BetaAppLocalization`

NewBetaAppLocalization instantiates a new BetaAppLocalization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaAppLocalizationWithDefaults

`func NewBetaAppLocalizationWithDefaults() *BetaAppLocalization`

NewBetaAppLocalizationWithDefaults instantiates a new BetaAppLocalization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BetaAppLocalization) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BetaAppLocalization) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BetaAppLocalization) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *BetaAppLocalization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BetaAppLocalization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BetaAppLocalization) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *BetaAppLocalization) GetAttributes() BetaAppLocalizationAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BetaAppLocalization) GetAttributesOk() (*BetaAppLocalizationAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BetaAppLocalization) SetAttributes(v BetaAppLocalizationAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *BetaAppLocalization) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *BetaAppLocalization) GetRelationships() AppEncryptionDeclarationRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BetaAppLocalization) GetRelationshipsOk() (*AppEncryptionDeclarationRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BetaAppLocalization) SetRelationships(v AppEncryptionDeclarationRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BetaAppLocalization) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *BetaAppLocalization) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BetaAppLocalization) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BetaAppLocalization) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


