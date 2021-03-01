# UserInvitationCreateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Attributes** | [**UserInvitationCreateRequestDataAttributes**](UserInvitationCreateRequest_data_attributes.md) |  | 
**Relationships** | Pointer to [**UserInvitationCreateRequestDataRelationships**](UserInvitationCreateRequest_data_relationships.md) |  | [optional] 

## Methods

### NewUserInvitationCreateRequestData

`func NewUserInvitationCreateRequestData(type_ string, attributes UserInvitationCreateRequestDataAttributes, ) *UserInvitationCreateRequestData`

NewUserInvitationCreateRequestData instantiates a new UserInvitationCreateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInvitationCreateRequestDataWithDefaults

`func NewUserInvitationCreateRequestDataWithDefaults() *UserInvitationCreateRequestData`

NewUserInvitationCreateRequestDataWithDefaults instantiates a new UserInvitationCreateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UserInvitationCreateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserInvitationCreateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserInvitationCreateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *UserInvitationCreateRequestData) GetAttributes() UserInvitationCreateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UserInvitationCreateRequestData) GetAttributesOk() (*UserInvitationCreateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UserInvitationCreateRequestData) SetAttributes(v UserInvitationCreateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *UserInvitationCreateRequestData) GetRelationships() UserInvitationCreateRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *UserInvitationCreateRequestData) GetRelationshipsOk() (*UserInvitationCreateRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *UserInvitationCreateRequestData) SetRelationships(v UserInvitationCreateRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *UserInvitationCreateRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


