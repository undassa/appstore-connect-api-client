# BetaGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**BetaGroupAttributes**](BetaGroup_attributes.md) |  | [optional] 
**Relationships** | Pointer to [**BetaGroupRelationships**](BetaGroup_relationships.md) |  | [optional] 
**Links** | [**ResourceLinks**](ResourceLinks.md) |  | 

## Methods

### NewBetaGroup

`func NewBetaGroup(type_ string, id string, links ResourceLinks, ) *BetaGroup`

NewBetaGroup instantiates a new BetaGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaGroupWithDefaults

`func NewBetaGroupWithDefaults() *BetaGroup`

NewBetaGroupWithDefaults instantiates a new BetaGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BetaGroup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BetaGroup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BetaGroup) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *BetaGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BetaGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BetaGroup) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *BetaGroup) GetAttributes() BetaGroupAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BetaGroup) GetAttributesOk() (*BetaGroupAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BetaGroup) SetAttributes(v BetaGroupAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *BetaGroup) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *BetaGroup) GetRelationships() BetaGroupRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BetaGroup) GetRelationshipsOk() (*BetaGroupRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BetaGroup) SetRelationships(v BetaGroupRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BetaGroup) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *BetaGroup) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BetaGroup) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BetaGroup) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


