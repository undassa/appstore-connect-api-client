# DiagnosticSignature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**DiagnosticSignatureAttributes**](DiagnosticSignature_attributes.md) |  | [optional] 
**Links** | [**ResourceLinks**](ResourceLinks.md) |  | 

## Methods

### NewDiagnosticSignature

`func NewDiagnosticSignature(type_ string, id string, links ResourceLinks, ) *DiagnosticSignature`

NewDiagnosticSignature instantiates a new DiagnosticSignature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiagnosticSignatureWithDefaults

`func NewDiagnosticSignatureWithDefaults() *DiagnosticSignature`

NewDiagnosticSignatureWithDefaults instantiates a new DiagnosticSignature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DiagnosticSignature) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DiagnosticSignature) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DiagnosticSignature) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *DiagnosticSignature) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiagnosticSignature) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiagnosticSignature) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *DiagnosticSignature) GetAttributes() DiagnosticSignatureAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DiagnosticSignature) GetAttributesOk() (*DiagnosticSignatureAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DiagnosticSignature) SetAttributes(v DiagnosticSignatureAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *DiagnosticSignature) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetLinks

`func (o *DiagnosticSignature) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DiagnosticSignature) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DiagnosticSignature) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


