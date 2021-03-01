# AppScreenshotCreateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Attributes** | [**AppScreenshotCreateRequestDataAttributes**](AppScreenshotCreateRequest_data_attributes.md) |  | 
**Relationships** | [**AppScreenshotCreateRequestDataRelationships**](AppScreenshotCreateRequest_data_relationships.md) |  | 

## Methods

### NewAppScreenshotCreateRequestData

`func NewAppScreenshotCreateRequestData(type_ string, attributes AppScreenshotCreateRequestDataAttributes, relationships AppScreenshotCreateRequestDataRelationships, ) *AppScreenshotCreateRequestData`

NewAppScreenshotCreateRequestData instantiates a new AppScreenshotCreateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppScreenshotCreateRequestDataWithDefaults

`func NewAppScreenshotCreateRequestDataWithDefaults() *AppScreenshotCreateRequestData`

NewAppScreenshotCreateRequestDataWithDefaults instantiates a new AppScreenshotCreateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppScreenshotCreateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppScreenshotCreateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppScreenshotCreateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *AppScreenshotCreateRequestData) GetAttributes() AppScreenshotCreateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppScreenshotCreateRequestData) GetAttributesOk() (*AppScreenshotCreateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppScreenshotCreateRequestData) SetAttributes(v AppScreenshotCreateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *AppScreenshotCreateRequestData) GetRelationships() AppScreenshotCreateRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AppScreenshotCreateRequestData) GetRelationshipsOk() (*AppScreenshotCreateRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AppScreenshotCreateRequestData) SetRelationships(v AppScreenshotCreateRequestDataRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


