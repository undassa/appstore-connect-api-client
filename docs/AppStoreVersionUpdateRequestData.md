# AppStoreVersionUpdateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**AppStoreVersionUpdateRequestDataAttributes**](AppStoreVersionUpdateRequest_data_attributes.md) |  | [optional] 
**Relationships** | Pointer to [**AppStoreVersionUpdateRequestDataRelationships**](AppStoreVersionUpdateRequest_data_relationships.md) |  | [optional] 

## Methods

### NewAppStoreVersionUpdateRequestData

`func NewAppStoreVersionUpdateRequestData(type_ string, id string, ) *AppStoreVersionUpdateRequestData`

NewAppStoreVersionUpdateRequestData instantiates a new AppStoreVersionUpdateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStoreVersionUpdateRequestDataWithDefaults

`func NewAppStoreVersionUpdateRequestDataWithDefaults() *AppStoreVersionUpdateRequestData`

NewAppStoreVersionUpdateRequestDataWithDefaults instantiates a new AppStoreVersionUpdateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppStoreVersionUpdateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppStoreVersionUpdateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppStoreVersionUpdateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *AppStoreVersionUpdateRequestData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppStoreVersionUpdateRequestData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppStoreVersionUpdateRequestData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *AppStoreVersionUpdateRequestData) GetAttributes() AppStoreVersionUpdateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppStoreVersionUpdateRequestData) GetAttributesOk() (*AppStoreVersionUpdateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppStoreVersionUpdateRequestData) SetAttributes(v AppStoreVersionUpdateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AppStoreVersionUpdateRequestData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *AppStoreVersionUpdateRequestData) GetRelationships() AppStoreVersionUpdateRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AppStoreVersionUpdateRequestData) GetRelationshipsOk() (*AppStoreVersionUpdateRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AppStoreVersionUpdateRequestData) SetRelationships(v AppStoreVersionUpdateRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AppStoreVersionUpdateRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


