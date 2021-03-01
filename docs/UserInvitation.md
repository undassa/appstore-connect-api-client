# UserInvitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**UserInvitationAttributes**](UserInvitation_attributes.md) |  | [optional] 
**Relationships** | Pointer to [**UserInvitationRelationships**](UserInvitation_relationships.md) |  | [optional] 
**Links** | [**ResourceLinks**](ResourceLinks.md) |  | 

## Methods

### NewUserInvitation

`func NewUserInvitation(type_ string, id string, links ResourceLinks, ) *UserInvitation`

NewUserInvitation instantiates a new UserInvitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInvitationWithDefaults

`func NewUserInvitationWithDefaults() *UserInvitation`

NewUserInvitationWithDefaults instantiates a new UserInvitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UserInvitation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserInvitation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserInvitation) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *UserInvitation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserInvitation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserInvitation) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *UserInvitation) GetAttributes() UserInvitationAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UserInvitation) GetAttributesOk() (*UserInvitationAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UserInvitation) SetAttributes(v UserInvitationAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UserInvitation) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *UserInvitation) GetRelationships() UserInvitationRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *UserInvitation) GetRelationshipsOk() (*UserInvitationRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *UserInvitation) SetRelationships(v UserInvitationRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *UserInvitation) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *UserInvitation) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserInvitation) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserInvitation) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


