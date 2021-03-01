# AppUpdateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**AppUpdateRequestDataAttributes**](AppUpdateRequest_data_attributes.md) |  | [optional] 
**Relationships** | Pointer to [**AppUpdateRequestDataRelationships**](AppUpdateRequest_data_relationships.md) |  | [optional] 

## Methods

### NewAppUpdateRequestData

`func NewAppUpdateRequestData(type_ string, id string, ) *AppUpdateRequestData`

NewAppUpdateRequestData instantiates a new AppUpdateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUpdateRequestDataWithDefaults

`func NewAppUpdateRequestDataWithDefaults() *AppUpdateRequestData`

NewAppUpdateRequestDataWithDefaults instantiates a new AppUpdateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppUpdateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppUpdateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppUpdateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *AppUpdateRequestData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppUpdateRequestData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppUpdateRequestData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *AppUpdateRequestData) GetAttributes() AppUpdateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppUpdateRequestData) GetAttributesOk() (*AppUpdateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppUpdateRequestData) SetAttributes(v AppUpdateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AppUpdateRequestData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *AppUpdateRequestData) GetRelationships() AppUpdateRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AppUpdateRequestData) GetRelationshipsOk() (*AppUpdateRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AppUpdateRequestData) SetRelationships(v AppUpdateRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AppUpdateRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


