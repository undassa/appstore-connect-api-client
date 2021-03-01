# UserUpdateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**UserUpdateRequestDataAttributes**](UserUpdateRequest_data_attributes.md) |  | [optional] 
**Relationships** | Pointer to [**UserInvitationCreateRequestDataRelationships**](UserInvitationCreateRequest_data_relationships.md) |  | [optional] 

## Methods

### NewUserUpdateRequestData

`func NewUserUpdateRequestData(type_ string, id string, ) *UserUpdateRequestData`

NewUserUpdateRequestData instantiates a new UserUpdateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUpdateRequestDataWithDefaults

`func NewUserUpdateRequestDataWithDefaults() *UserUpdateRequestData`

NewUserUpdateRequestDataWithDefaults instantiates a new UserUpdateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UserUpdateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserUpdateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserUpdateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *UserUpdateRequestData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserUpdateRequestData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserUpdateRequestData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *UserUpdateRequestData) GetAttributes() UserUpdateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UserUpdateRequestData) GetAttributesOk() (*UserUpdateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UserUpdateRequestData) SetAttributes(v UserUpdateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UserUpdateRequestData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *UserUpdateRequestData) GetRelationships() UserInvitationCreateRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *UserUpdateRequestData) GetRelationshipsOk() (*UserInvitationCreateRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *UserUpdateRequestData) SetRelationships(v UserInvitationCreateRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *UserUpdateRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


