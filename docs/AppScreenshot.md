# AppScreenshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**AppScreenshotAttributes**](AppScreenshot_attributes.md) |  | [optional] 
**Relationships** | Pointer to [**AppScreenshotRelationships**](AppScreenshot_relationships.md) |  | [optional] 
**Links** | [**ResourceLinks**](ResourceLinks.md) |  | 

## Methods

### NewAppScreenshot

`func NewAppScreenshot(type_ string, id string, links ResourceLinks, ) *AppScreenshot`

NewAppScreenshot instantiates a new AppScreenshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppScreenshotWithDefaults

`func NewAppScreenshotWithDefaults() *AppScreenshot`

NewAppScreenshotWithDefaults instantiates a new AppScreenshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppScreenshot) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppScreenshot) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppScreenshot) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *AppScreenshot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppScreenshot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppScreenshot) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *AppScreenshot) GetAttributes() AppScreenshotAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppScreenshot) GetAttributesOk() (*AppScreenshotAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppScreenshot) SetAttributes(v AppScreenshotAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AppScreenshot) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *AppScreenshot) GetRelationships() AppScreenshotRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AppScreenshot) GetRelationshipsOk() (*AppScreenshotRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AppScreenshot) SetRelationships(v AppScreenshotRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AppScreenshot) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *AppScreenshot) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppScreenshot) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppScreenshot) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


