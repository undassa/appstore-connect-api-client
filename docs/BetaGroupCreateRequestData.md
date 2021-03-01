# BetaGroupCreateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Attributes** | [**BetaGroupCreateRequestDataAttributes**](BetaGroupCreateRequest_data_attributes.md) |  | 
**Relationships** | [**BetaGroupCreateRequestDataRelationships**](BetaGroupCreateRequest_data_relationships.md) |  | 

## Methods

### NewBetaGroupCreateRequestData

`func NewBetaGroupCreateRequestData(type_ string, attributes BetaGroupCreateRequestDataAttributes, relationships BetaGroupCreateRequestDataRelationships, ) *BetaGroupCreateRequestData`

NewBetaGroupCreateRequestData instantiates a new BetaGroupCreateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaGroupCreateRequestDataWithDefaults

`func NewBetaGroupCreateRequestDataWithDefaults() *BetaGroupCreateRequestData`

NewBetaGroupCreateRequestDataWithDefaults instantiates a new BetaGroupCreateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BetaGroupCreateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BetaGroupCreateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BetaGroupCreateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BetaGroupCreateRequestData) GetAttributes() BetaGroupCreateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BetaGroupCreateRequestData) GetAttributesOk() (*BetaGroupCreateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BetaGroupCreateRequestData) SetAttributes(v BetaGroupCreateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *BetaGroupCreateRequestData) GetRelationships() BetaGroupCreateRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BetaGroupCreateRequestData) GetRelationshipsOk() (*BetaGroupCreateRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BetaGroupCreateRequestData) SetRelationships(v BetaGroupCreateRequestDataRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


