# AgeRatingDeclaration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**AgeRatingDeclarationAttributes**](AgeRatingDeclaration_attributes.md) |  | [optional] 
**Links** | [**ResourceLinks**](ResourceLinks.md) |  | 

## Methods

### NewAgeRatingDeclaration

`func NewAgeRatingDeclaration(type_ string, id string, links ResourceLinks, ) *AgeRatingDeclaration`

NewAgeRatingDeclaration instantiates a new AgeRatingDeclaration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgeRatingDeclarationWithDefaults

`func NewAgeRatingDeclarationWithDefaults() *AgeRatingDeclaration`

NewAgeRatingDeclarationWithDefaults instantiates a new AgeRatingDeclaration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AgeRatingDeclaration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgeRatingDeclaration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgeRatingDeclaration) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *AgeRatingDeclaration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgeRatingDeclaration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgeRatingDeclaration) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *AgeRatingDeclaration) GetAttributes() AgeRatingDeclarationAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AgeRatingDeclaration) GetAttributesOk() (*AgeRatingDeclarationAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AgeRatingDeclaration) SetAttributes(v AgeRatingDeclarationAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AgeRatingDeclaration) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetLinks

`func (o *AgeRatingDeclaration) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AgeRatingDeclaration) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AgeRatingDeclaration) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


