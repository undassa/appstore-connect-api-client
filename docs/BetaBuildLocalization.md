# BetaBuildLocalization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**BetaBuildLocalizationAttributes**](BetaBuildLocalization_attributes.md) |  | [optional] 
**Relationships** | Pointer to [**BetaAppReviewSubmissionRelationships**](BetaAppReviewSubmission_relationships.md) |  | [optional] 
**Links** | [**ResourceLinks**](ResourceLinks.md) |  | 

## Methods

### NewBetaBuildLocalization

`func NewBetaBuildLocalization(type_ string, id string, links ResourceLinks, ) *BetaBuildLocalization`

NewBetaBuildLocalization instantiates a new BetaBuildLocalization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaBuildLocalizationWithDefaults

`func NewBetaBuildLocalizationWithDefaults() *BetaBuildLocalization`

NewBetaBuildLocalizationWithDefaults instantiates a new BetaBuildLocalization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BetaBuildLocalization) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BetaBuildLocalization) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BetaBuildLocalization) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *BetaBuildLocalization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BetaBuildLocalization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BetaBuildLocalization) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *BetaBuildLocalization) GetAttributes() BetaBuildLocalizationAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BetaBuildLocalization) GetAttributesOk() (*BetaBuildLocalizationAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BetaBuildLocalization) SetAttributes(v BetaBuildLocalizationAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *BetaBuildLocalization) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *BetaBuildLocalization) GetRelationships() BetaAppReviewSubmissionRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BetaBuildLocalization) GetRelationshipsOk() (*BetaAppReviewSubmissionRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BetaBuildLocalization) SetRelationships(v BetaAppReviewSubmissionRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BetaBuildLocalization) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *BetaBuildLocalization) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BetaBuildLocalization) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BetaBuildLocalization) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


