# AppInfoLocalization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**AppInfoLocalizationAttributes**](AppInfoLocalization_attributes.md) |  | [optional] 
**Relationships** | Pointer to [**AppInfoLocalizationRelationships**](AppInfoLocalization_relationships.md) |  | [optional] 
**Links** | [**ResourceLinks**](ResourceLinks.md) |  | 

## Methods

### NewAppInfoLocalization

`func NewAppInfoLocalization(type_ string, id string, links ResourceLinks, ) *AppInfoLocalization`

NewAppInfoLocalization instantiates a new AppInfoLocalization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppInfoLocalizationWithDefaults

`func NewAppInfoLocalizationWithDefaults() *AppInfoLocalization`

NewAppInfoLocalizationWithDefaults instantiates a new AppInfoLocalization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppInfoLocalization) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppInfoLocalization) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppInfoLocalization) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *AppInfoLocalization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppInfoLocalization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppInfoLocalization) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *AppInfoLocalization) GetAttributes() AppInfoLocalizationAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppInfoLocalization) GetAttributesOk() (*AppInfoLocalizationAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppInfoLocalization) SetAttributes(v AppInfoLocalizationAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AppInfoLocalization) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *AppInfoLocalization) GetRelationships() AppInfoLocalizationRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AppInfoLocalization) GetRelationshipsOk() (*AppInfoLocalizationRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AppInfoLocalization) SetRelationships(v AppInfoLocalizationRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AppInfoLocalization) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *AppInfoLocalization) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppInfoLocalization) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppInfoLocalization) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


