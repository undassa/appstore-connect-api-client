# AppStoreVersionLocalization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**AppStoreVersionLocalizationAttributes**](AppStoreVersionLocalization_attributes.md) |  | [optional] 
**Relationships** | Pointer to [**AppStoreVersionLocalizationRelationships**](AppStoreVersionLocalization_relationships.md) |  | [optional] 
**Links** | [**ResourceLinks**](ResourceLinks.md) |  | 

## Methods

### NewAppStoreVersionLocalization

`func NewAppStoreVersionLocalization(type_ string, id string, links ResourceLinks, ) *AppStoreVersionLocalization`

NewAppStoreVersionLocalization instantiates a new AppStoreVersionLocalization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStoreVersionLocalizationWithDefaults

`func NewAppStoreVersionLocalizationWithDefaults() *AppStoreVersionLocalization`

NewAppStoreVersionLocalizationWithDefaults instantiates a new AppStoreVersionLocalization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppStoreVersionLocalization) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppStoreVersionLocalization) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppStoreVersionLocalization) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *AppStoreVersionLocalization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppStoreVersionLocalization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppStoreVersionLocalization) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *AppStoreVersionLocalization) GetAttributes() AppStoreVersionLocalizationAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppStoreVersionLocalization) GetAttributesOk() (*AppStoreVersionLocalizationAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppStoreVersionLocalization) SetAttributes(v AppStoreVersionLocalizationAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AppStoreVersionLocalization) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *AppStoreVersionLocalization) GetRelationships() AppStoreVersionLocalizationRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AppStoreVersionLocalization) GetRelationshipsOk() (*AppStoreVersionLocalizationRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AppStoreVersionLocalization) SetRelationships(v AppStoreVersionLocalizationRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AppStoreVersionLocalization) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *AppStoreVersionLocalization) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppStoreVersionLocalization) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppStoreVersionLocalization) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


