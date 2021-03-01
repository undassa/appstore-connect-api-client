# EndUserLicenseAgreement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**BetaLicenseAgreementAttributes**](BetaLicenseAgreement_attributes.md) |  | [optional] 
**Relationships** | Pointer to [**EndUserLicenseAgreementRelationships**](EndUserLicenseAgreement_relationships.md) |  | [optional] 
**Links** | [**ResourceLinks**](ResourceLinks.md) |  | 

## Methods

### NewEndUserLicenseAgreement

`func NewEndUserLicenseAgreement(type_ string, id string, links ResourceLinks, ) *EndUserLicenseAgreement`

NewEndUserLicenseAgreement instantiates a new EndUserLicenseAgreement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndUserLicenseAgreementWithDefaults

`func NewEndUserLicenseAgreementWithDefaults() *EndUserLicenseAgreement`

NewEndUserLicenseAgreementWithDefaults instantiates a new EndUserLicenseAgreement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EndUserLicenseAgreement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EndUserLicenseAgreement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EndUserLicenseAgreement) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *EndUserLicenseAgreement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EndUserLicenseAgreement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EndUserLicenseAgreement) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *EndUserLicenseAgreement) GetAttributes() BetaLicenseAgreementAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EndUserLicenseAgreement) GetAttributesOk() (*BetaLicenseAgreementAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EndUserLicenseAgreement) SetAttributes(v BetaLicenseAgreementAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *EndUserLicenseAgreement) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *EndUserLicenseAgreement) GetRelationships() EndUserLicenseAgreementRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *EndUserLicenseAgreement) GetRelationshipsOk() (*EndUserLicenseAgreementRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *EndUserLicenseAgreement) SetRelationships(v EndUserLicenseAgreementRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *EndUserLicenseAgreement) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *EndUserLicenseAgreement) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EndUserLicenseAgreement) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EndUserLicenseAgreement) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


