# PrereleaseVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**PrereleaseVersionAttributes**](PrereleaseVersion_attributes.md) |  | [optional] 
**Relationships** | Pointer to [**PrereleaseVersionRelationships**](PrereleaseVersion_relationships.md) |  | [optional] 
**Links** | [**ResourceLinks**](ResourceLinks.md) |  | 

## Methods

### NewPrereleaseVersion

`func NewPrereleaseVersion(type_ string, id string, links ResourceLinks, ) *PrereleaseVersion`

NewPrereleaseVersion instantiates a new PrereleaseVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrereleaseVersionWithDefaults

`func NewPrereleaseVersionWithDefaults() *PrereleaseVersion`

NewPrereleaseVersionWithDefaults instantiates a new PrereleaseVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PrereleaseVersion) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PrereleaseVersion) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PrereleaseVersion) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *PrereleaseVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrereleaseVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrereleaseVersion) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *PrereleaseVersion) GetAttributes() PrereleaseVersionAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PrereleaseVersion) GetAttributesOk() (*PrereleaseVersionAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PrereleaseVersion) SetAttributes(v PrereleaseVersionAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PrereleaseVersion) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *PrereleaseVersion) GetRelationships() PrereleaseVersionRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PrereleaseVersion) GetRelationshipsOk() (*PrereleaseVersionRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PrereleaseVersion) SetRelationships(v PrereleaseVersionRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PrereleaseVersion) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *PrereleaseVersion) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PrereleaseVersion) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PrereleaseVersion) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


