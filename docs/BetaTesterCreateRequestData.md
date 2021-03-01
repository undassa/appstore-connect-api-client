# BetaTesterCreateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Attributes** | [**BetaTesterCreateRequestDataAttributes**](BetaTesterCreateRequest_data_attributes.md) |  | 
**Relationships** | Pointer to [**BetaTesterCreateRequestDataRelationships**](BetaTesterCreateRequest_data_relationships.md) |  | [optional] 

## Methods

### NewBetaTesterCreateRequestData

`func NewBetaTesterCreateRequestData(type_ string, attributes BetaTesterCreateRequestDataAttributes, ) *BetaTesterCreateRequestData`

NewBetaTesterCreateRequestData instantiates a new BetaTesterCreateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaTesterCreateRequestDataWithDefaults

`func NewBetaTesterCreateRequestDataWithDefaults() *BetaTesterCreateRequestData`

NewBetaTesterCreateRequestDataWithDefaults instantiates a new BetaTesterCreateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BetaTesterCreateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BetaTesterCreateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BetaTesterCreateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BetaTesterCreateRequestData) GetAttributes() BetaTesterCreateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BetaTesterCreateRequestData) GetAttributesOk() (*BetaTesterCreateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BetaTesterCreateRequestData) SetAttributes(v BetaTesterCreateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *BetaTesterCreateRequestData) GetRelationships() BetaTesterCreateRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BetaTesterCreateRequestData) GetRelationshipsOk() (*BetaTesterCreateRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BetaTesterCreateRequestData) SetRelationships(v BetaTesterCreateRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BetaTesterCreateRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


