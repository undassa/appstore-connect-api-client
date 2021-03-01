# AppEncryptionDeclaration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**AppEncryptionDeclarationAttributes**](AppEncryptionDeclaration_attributes.md) |  | [optional] 
**Relationships** | Pointer to [**AppEncryptionDeclarationRelationships**](AppEncryptionDeclaration_relationships.md) |  | [optional] 
**Links** | [**ResourceLinks**](ResourceLinks.md) |  | 

## Methods

### NewAppEncryptionDeclaration

`func NewAppEncryptionDeclaration(type_ string, id string, links ResourceLinks, ) *AppEncryptionDeclaration`

NewAppEncryptionDeclaration instantiates a new AppEncryptionDeclaration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppEncryptionDeclarationWithDefaults

`func NewAppEncryptionDeclarationWithDefaults() *AppEncryptionDeclaration`

NewAppEncryptionDeclarationWithDefaults instantiates a new AppEncryptionDeclaration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppEncryptionDeclaration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppEncryptionDeclaration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppEncryptionDeclaration) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *AppEncryptionDeclaration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppEncryptionDeclaration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppEncryptionDeclaration) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *AppEncryptionDeclaration) GetAttributes() AppEncryptionDeclarationAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppEncryptionDeclaration) GetAttributesOk() (*AppEncryptionDeclarationAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppEncryptionDeclaration) SetAttributes(v AppEncryptionDeclarationAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AppEncryptionDeclaration) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *AppEncryptionDeclaration) GetRelationships() AppEncryptionDeclarationRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AppEncryptionDeclaration) GetRelationshipsOk() (*AppEncryptionDeclarationRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AppEncryptionDeclaration) SetRelationships(v AppEncryptionDeclarationRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AppEncryptionDeclaration) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *AppEncryptionDeclaration) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppEncryptionDeclaration) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppEncryptionDeclaration) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


