# AppScreenshotSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**AppScreenshotSetAttributes**](AppScreenshotSet_attributes.md) |  | [optional] 
**Relationships** | Pointer to [**AppScreenshotSetRelationships**](AppScreenshotSet_relationships.md) |  | [optional] 
**Links** | [**ResourceLinks**](ResourceLinks.md) |  | 

## Methods

### NewAppScreenshotSet

`func NewAppScreenshotSet(type_ string, id string, links ResourceLinks, ) *AppScreenshotSet`

NewAppScreenshotSet instantiates a new AppScreenshotSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppScreenshotSetWithDefaults

`func NewAppScreenshotSetWithDefaults() *AppScreenshotSet`

NewAppScreenshotSetWithDefaults instantiates a new AppScreenshotSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppScreenshotSet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppScreenshotSet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppScreenshotSet) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *AppScreenshotSet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppScreenshotSet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppScreenshotSet) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *AppScreenshotSet) GetAttributes() AppScreenshotSetAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppScreenshotSet) GetAttributesOk() (*AppScreenshotSetAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppScreenshotSet) SetAttributes(v AppScreenshotSetAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AppScreenshotSet) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *AppScreenshotSet) GetRelationships() AppScreenshotSetRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AppScreenshotSet) GetRelationshipsOk() (*AppScreenshotSetRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AppScreenshotSet) SetRelationships(v AppScreenshotSetRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AppScreenshotSet) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *AppScreenshotSet) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppScreenshotSet) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppScreenshotSet) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


