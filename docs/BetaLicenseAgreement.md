# BetaLicenseAgreement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**BetaLicenseAgreementAttributes**](BetaLicenseAgreement_attributes.md) |  | [optional] 
**Relationships** | Pointer to [**AppEncryptionDeclarationRelationships**](AppEncryptionDeclaration_relationships.md) |  | [optional] 
**Links** | [**ResourceLinks**](ResourceLinks.md) |  | 

## Methods

### NewBetaLicenseAgreement

`func NewBetaLicenseAgreement(type_ string, id string, links ResourceLinks, ) *BetaLicenseAgreement`

NewBetaLicenseAgreement instantiates a new BetaLicenseAgreement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaLicenseAgreementWithDefaults

`func NewBetaLicenseAgreementWithDefaults() *BetaLicenseAgreement`

NewBetaLicenseAgreementWithDefaults instantiates a new BetaLicenseAgreement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BetaLicenseAgreement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BetaLicenseAgreement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BetaLicenseAgreement) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *BetaLicenseAgreement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BetaLicenseAgreement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BetaLicenseAgreement) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *BetaLicenseAgreement) GetAttributes() BetaLicenseAgreementAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BetaLicenseAgreement) GetAttributesOk() (*BetaLicenseAgreementAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BetaLicenseAgreement) SetAttributes(v BetaLicenseAgreementAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *BetaLicenseAgreement) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *BetaLicenseAgreement) GetRelationships() AppEncryptionDeclarationRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BetaLicenseAgreement) GetRelationshipsOk() (*AppEncryptionDeclarationRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BetaLicenseAgreement) SetRelationships(v AppEncryptionDeclarationRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BetaLicenseAgreement) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *BetaLicenseAgreement) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BetaLicenseAgreement) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BetaLicenseAgreement) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


