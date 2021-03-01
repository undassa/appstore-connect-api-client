# GameCenterEnabledVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**GameCenterEnabledVersionAttributes**](GameCenterEnabledVersion_attributes.md) |  | [optional] 
**Relationships** | Pointer to [**GameCenterEnabledVersionRelationships**](GameCenterEnabledVersion_relationships.md) |  | [optional] 
**Links** | [**ResourceLinks**](ResourceLinks.md) |  | 

## Methods

### NewGameCenterEnabledVersion

`func NewGameCenterEnabledVersion(type_ string, id string, links ResourceLinks, ) *GameCenterEnabledVersion`

NewGameCenterEnabledVersion instantiates a new GameCenterEnabledVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterEnabledVersionWithDefaults

`func NewGameCenterEnabledVersionWithDefaults() *GameCenterEnabledVersion`

NewGameCenterEnabledVersionWithDefaults instantiates a new GameCenterEnabledVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GameCenterEnabledVersion) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GameCenterEnabledVersion) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GameCenterEnabledVersion) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *GameCenterEnabledVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GameCenterEnabledVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GameCenterEnabledVersion) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GameCenterEnabledVersion) GetAttributes() GameCenterEnabledVersionAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GameCenterEnabledVersion) GetAttributesOk() (*GameCenterEnabledVersionAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GameCenterEnabledVersion) SetAttributes(v GameCenterEnabledVersionAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GameCenterEnabledVersion) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *GameCenterEnabledVersion) GetRelationships() GameCenterEnabledVersionRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GameCenterEnabledVersion) GetRelationshipsOk() (*GameCenterEnabledVersionRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GameCenterEnabledVersion) SetRelationships(v GameCenterEnabledVersionRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GameCenterEnabledVersion) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterEnabledVersion) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterEnabledVersion) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterEnabledVersion) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


