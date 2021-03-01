# AppPreOrderCreateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Attributes** | Pointer to [**AppPreOrderCreateRequestDataAttributes**](AppPreOrderCreateRequest_data_attributes.md) |  | [optional] 
**Relationships** | [**AppPreOrderCreateRequestDataRelationships**](AppPreOrderCreateRequest_data_relationships.md) |  | 

## Methods

### NewAppPreOrderCreateRequestData

`func NewAppPreOrderCreateRequestData(type_ string, relationships AppPreOrderCreateRequestDataRelationships, ) *AppPreOrderCreateRequestData`

NewAppPreOrderCreateRequestData instantiates a new AppPreOrderCreateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPreOrderCreateRequestDataWithDefaults

`func NewAppPreOrderCreateRequestDataWithDefaults() *AppPreOrderCreateRequestData`

NewAppPreOrderCreateRequestDataWithDefaults instantiates a new AppPreOrderCreateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppPreOrderCreateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppPreOrderCreateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppPreOrderCreateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *AppPreOrderCreateRequestData) GetAttributes() AppPreOrderCreateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppPreOrderCreateRequestData) GetAttributesOk() (*AppPreOrderCreateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppPreOrderCreateRequestData) SetAttributes(v AppPreOrderCreateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AppPreOrderCreateRequestData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *AppPreOrderCreateRequestData) GetRelationships() AppPreOrderCreateRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AppPreOrderCreateRequestData) GetRelationshipsOk() (*AppPreOrderCreateRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AppPreOrderCreateRequestData) SetRelationships(v AppPreOrderCreateRequestDataRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


