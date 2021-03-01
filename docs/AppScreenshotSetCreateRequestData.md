# AppScreenshotSetCreateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Attributes** | [**AppScreenshotSetCreateRequestDataAttributes**](AppScreenshotSetCreateRequest_data_attributes.md) |  | 
**Relationships** | [**AppPreviewSetCreateRequestDataRelationships**](AppPreviewSetCreateRequest_data_relationships.md) |  | 

## Methods

### NewAppScreenshotSetCreateRequestData

`func NewAppScreenshotSetCreateRequestData(type_ string, attributes AppScreenshotSetCreateRequestDataAttributes, relationships AppPreviewSetCreateRequestDataRelationships, ) *AppScreenshotSetCreateRequestData`

NewAppScreenshotSetCreateRequestData instantiates a new AppScreenshotSetCreateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppScreenshotSetCreateRequestDataWithDefaults

`func NewAppScreenshotSetCreateRequestDataWithDefaults() *AppScreenshotSetCreateRequestData`

NewAppScreenshotSetCreateRequestDataWithDefaults instantiates a new AppScreenshotSetCreateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppScreenshotSetCreateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppScreenshotSetCreateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppScreenshotSetCreateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *AppScreenshotSetCreateRequestData) GetAttributes() AppScreenshotSetCreateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppScreenshotSetCreateRequestData) GetAttributesOk() (*AppScreenshotSetCreateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppScreenshotSetCreateRequestData) SetAttributes(v AppScreenshotSetCreateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *AppScreenshotSetCreateRequestData) GetRelationships() AppPreviewSetCreateRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AppScreenshotSetCreateRequestData) GetRelationshipsOk() (*AppPreviewSetCreateRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AppScreenshotSetCreateRequestData) SetRelationships(v AppPreviewSetCreateRequestDataRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


