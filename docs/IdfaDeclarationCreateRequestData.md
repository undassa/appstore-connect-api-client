# IdfaDeclarationCreateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Attributes** | [**IdfaDeclarationCreateRequestDataAttributes**](IdfaDeclarationCreateRequest_data_attributes.md) |  | 
**Relationships** | [**AppStoreReviewDetailCreateRequestDataRelationships**](AppStoreReviewDetailCreateRequest_data_relationships.md) |  | 

## Methods

### NewIdfaDeclarationCreateRequestData

`func NewIdfaDeclarationCreateRequestData(type_ string, attributes IdfaDeclarationCreateRequestDataAttributes, relationships AppStoreReviewDetailCreateRequestDataRelationships, ) *IdfaDeclarationCreateRequestData`

NewIdfaDeclarationCreateRequestData instantiates a new IdfaDeclarationCreateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdfaDeclarationCreateRequestDataWithDefaults

`func NewIdfaDeclarationCreateRequestDataWithDefaults() *IdfaDeclarationCreateRequestData`

NewIdfaDeclarationCreateRequestDataWithDefaults instantiates a new IdfaDeclarationCreateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IdfaDeclarationCreateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdfaDeclarationCreateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdfaDeclarationCreateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *IdfaDeclarationCreateRequestData) GetAttributes() IdfaDeclarationCreateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IdfaDeclarationCreateRequestData) GetAttributesOk() (*IdfaDeclarationCreateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IdfaDeclarationCreateRequestData) SetAttributes(v IdfaDeclarationCreateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *IdfaDeclarationCreateRequestData) GetRelationships() AppStoreReviewDetailCreateRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *IdfaDeclarationCreateRequestData) GetRelationshipsOk() (*AppStoreReviewDetailCreateRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *IdfaDeclarationCreateRequestData) SetRelationships(v AppStoreReviewDetailCreateRequestDataRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


