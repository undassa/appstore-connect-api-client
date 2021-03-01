# AppPreviewSetCreateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Attributes** | [**AppPreviewSetCreateRequestDataAttributes**](AppPreviewSetCreateRequest_data_attributes.md) |  | 
**Relationships** | [**AppPreviewSetCreateRequestDataRelationships**](AppPreviewSetCreateRequest_data_relationships.md) |  | 

## Methods

### NewAppPreviewSetCreateRequestData

`func NewAppPreviewSetCreateRequestData(type_ string, attributes AppPreviewSetCreateRequestDataAttributes, relationships AppPreviewSetCreateRequestDataRelationships, ) *AppPreviewSetCreateRequestData`

NewAppPreviewSetCreateRequestData instantiates a new AppPreviewSetCreateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPreviewSetCreateRequestDataWithDefaults

`func NewAppPreviewSetCreateRequestDataWithDefaults() *AppPreviewSetCreateRequestData`

NewAppPreviewSetCreateRequestDataWithDefaults instantiates a new AppPreviewSetCreateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppPreviewSetCreateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppPreviewSetCreateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppPreviewSetCreateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *AppPreviewSetCreateRequestData) GetAttributes() AppPreviewSetCreateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppPreviewSetCreateRequestData) GetAttributesOk() (*AppPreviewSetCreateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppPreviewSetCreateRequestData) SetAttributes(v AppPreviewSetCreateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *AppPreviewSetCreateRequestData) GetRelationships() AppPreviewSetCreateRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AppPreviewSetCreateRequestData) GetRelationshipsOk() (*AppPreviewSetCreateRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AppPreviewSetCreateRequestData) SetRelationships(v AppPreviewSetCreateRequestDataRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


