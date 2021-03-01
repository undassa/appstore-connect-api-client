# AppPreviewUpdateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**AppPreviewUpdateRequestDataAttributes**](AppPreviewUpdateRequest_data_attributes.md) |  | [optional] 

## Methods

### NewAppPreviewUpdateRequestData

`func NewAppPreviewUpdateRequestData(type_ string, id string, ) *AppPreviewUpdateRequestData`

NewAppPreviewUpdateRequestData instantiates a new AppPreviewUpdateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPreviewUpdateRequestDataWithDefaults

`func NewAppPreviewUpdateRequestDataWithDefaults() *AppPreviewUpdateRequestData`

NewAppPreviewUpdateRequestDataWithDefaults instantiates a new AppPreviewUpdateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppPreviewUpdateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppPreviewUpdateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppPreviewUpdateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *AppPreviewUpdateRequestData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppPreviewUpdateRequestData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppPreviewUpdateRequestData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *AppPreviewUpdateRequestData) GetAttributes() AppPreviewUpdateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppPreviewUpdateRequestData) GetAttributesOk() (*AppPreviewUpdateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppPreviewUpdateRequestData) SetAttributes(v AppPreviewUpdateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AppPreviewUpdateRequestData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


