# IdfaDeclaration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**IdfaDeclarationAttributes**](IdfaDeclaration_attributes.md) |  | [optional] 
**Relationships** | Pointer to [**AppStoreVersionSubmissionRelationships**](AppStoreVersionSubmission_relationships.md) |  | [optional] 
**Links** | [**ResourceLinks**](ResourceLinks.md) |  | 

## Methods

### NewIdfaDeclaration

`func NewIdfaDeclaration(type_ string, id string, links ResourceLinks, ) *IdfaDeclaration`

NewIdfaDeclaration instantiates a new IdfaDeclaration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdfaDeclarationWithDefaults

`func NewIdfaDeclarationWithDefaults() *IdfaDeclaration`

NewIdfaDeclarationWithDefaults instantiates a new IdfaDeclaration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IdfaDeclaration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdfaDeclaration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdfaDeclaration) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *IdfaDeclaration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdfaDeclaration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdfaDeclaration) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *IdfaDeclaration) GetAttributes() IdfaDeclarationAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IdfaDeclaration) GetAttributesOk() (*IdfaDeclarationAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IdfaDeclaration) SetAttributes(v IdfaDeclarationAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IdfaDeclaration) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *IdfaDeclaration) GetRelationships() AppStoreVersionSubmissionRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *IdfaDeclaration) GetRelationshipsOk() (*AppStoreVersionSubmissionRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *IdfaDeclaration) SetRelationships(v AppStoreVersionSubmissionRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *IdfaDeclaration) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *IdfaDeclaration) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IdfaDeclaration) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IdfaDeclaration) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


