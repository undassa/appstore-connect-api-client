# AppInfoUpdateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Relationships** | Pointer to [**AppInfoUpdateRequestDataRelationships**](AppInfoUpdateRequest_data_relationships.md) |  | [optional] 

## Methods

### NewAppInfoUpdateRequestData

`func NewAppInfoUpdateRequestData(type_ string, id string, ) *AppInfoUpdateRequestData`

NewAppInfoUpdateRequestData instantiates a new AppInfoUpdateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppInfoUpdateRequestDataWithDefaults

`func NewAppInfoUpdateRequestDataWithDefaults() *AppInfoUpdateRequestData`

NewAppInfoUpdateRequestDataWithDefaults instantiates a new AppInfoUpdateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppInfoUpdateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppInfoUpdateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppInfoUpdateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *AppInfoUpdateRequestData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppInfoUpdateRequestData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppInfoUpdateRequestData) SetId(v string)`

SetId sets Id field to given value.


### GetRelationships

`func (o *AppInfoUpdateRequestData) GetRelationships() AppInfoUpdateRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AppInfoUpdateRequestData) GetRelationshipsOk() (*AppInfoUpdateRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AppInfoUpdateRequestData) SetRelationships(v AppInfoUpdateRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AppInfoUpdateRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


