# AppEncryptionDeclarationAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsesEncryption** | Pointer to **bool** |  | [optional] 
**Exempt** | Pointer to **bool** |  | [optional] 
**ContainsProprietaryCryptography** | Pointer to **bool** |  | [optional] 
**ContainsThirdPartyCryptography** | Pointer to **bool** |  | [optional] 
**AvailableOnFrenchStore** | Pointer to **bool** |  | [optional] 
**Platform** | Pointer to [**Platform**](Platform.md) |  | [optional] 
**UploadedDate** | Pointer to **time.Time** |  | [optional] 
**DocumentUrl** | Pointer to **string** |  | [optional] 
**DocumentName** | Pointer to **string** |  | [optional] 
**DocumentType** | Pointer to **string** |  | [optional] 
**AppEncryptionDeclarationState** | Pointer to [**AppEncryptionDeclarationState**](AppEncryptionDeclarationState.md) |  | [optional] 
**CodeValue** | Pointer to **string** |  | [optional] 

## Methods

### NewAppEncryptionDeclarationAttributes

`func NewAppEncryptionDeclarationAttributes() *AppEncryptionDeclarationAttributes`

NewAppEncryptionDeclarationAttributes instantiates a new AppEncryptionDeclarationAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppEncryptionDeclarationAttributesWithDefaults

`func NewAppEncryptionDeclarationAttributesWithDefaults() *AppEncryptionDeclarationAttributes`

NewAppEncryptionDeclarationAttributesWithDefaults instantiates a new AppEncryptionDeclarationAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsesEncryption

`func (o *AppEncryptionDeclarationAttributes) GetUsesEncryption() bool`

GetUsesEncryption returns the UsesEncryption field if non-nil, zero value otherwise.

### GetUsesEncryptionOk

`func (o *AppEncryptionDeclarationAttributes) GetUsesEncryptionOk() (*bool, bool)`

GetUsesEncryptionOk returns a tuple with the UsesEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesEncryption

`func (o *AppEncryptionDeclarationAttributes) SetUsesEncryption(v bool)`

SetUsesEncryption sets UsesEncryption field to given value.

### HasUsesEncryption

`func (o *AppEncryptionDeclarationAttributes) HasUsesEncryption() bool`

HasUsesEncryption returns a boolean if a field has been set.

### GetExempt

`func (o *AppEncryptionDeclarationAttributes) GetExempt() bool`

GetExempt returns the Exempt field if non-nil, zero value otherwise.

### GetExemptOk

`func (o *AppEncryptionDeclarationAttributes) GetExemptOk() (*bool, bool)`

GetExemptOk returns a tuple with the Exempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExempt

`func (o *AppEncryptionDeclarationAttributes) SetExempt(v bool)`

SetExempt sets Exempt field to given value.

### HasExempt

`func (o *AppEncryptionDeclarationAttributes) HasExempt() bool`

HasExempt returns a boolean if a field has been set.

### GetContainsProprietaryCryptography

`func (o *AppEncryptionDeclarationAttributes) GetContainsProprietaryCryptography() bool`

GetContainsProprietaryCryptography returns the ContainsProprietaryCryptography field if non-nil, zero value otherwise.

### GetContainsProprietaryCryptographyOk

`func (o *AppEncryptionDeclarationAttributes) GetContainsProprietaryCryptographyOk() (*bool, bool)`

GetContainsProprietaryCryptographyOk returns a tuple with the ContainsProprietaryCryptography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsProprietaryCryptography

`func (o *AppEncryptionDeclarationAttributes) SetContainsProprietaryCryptography(v bool)`

SetContainsProprietaryCryptography sets ContainsProprietaryCryptography field to given value.

### HasContainsProprietaryCryptography

`func (o *AppEncryptionDeclarationAttributes) HasContainsProprietaryCryptography() bool`

HasContainsProprietaryCryptography returns a boolean if a field has been set.

### GetContainsThirdPartyCryptography

`func (o *AppEncryptionDeclarationAttributes) GetContainsThirdPartyCryptography() bool`

GetContainsThirdPartyCryptography returns the ContainsThirdPartyCryptography field if non-nil, zero value otherwise.

### GetContainsThirdPartyCryptographyOk

`func (o *AppEncryptionDeclarationAttributes) GetContainsThirdPartyCryptographyOk() (*bool, bool)`

GetContainsThirdPartyCryptographyOk returns a tuple with the ContainsThirdPartyCryptography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsThirdPartyCryptography

`func (o *AppEncryptionDeclarationAttributes) SetContainsThirdPartyCryptography(v bool)`

SetContainsThirdPartyCryptography sets ContainsThirdPartyCryptography field to given value.

### HasContainsThirdPartyCryptography

`func (o *AppEncryptionDeclarationAttributes) HasContainsThirdPartyCryptography() bool`

HasContainsThirdPartyCryptography returns a boolean if a field has been set.

### GetAvailableOnFrenchStore

`func (o *AppEncryptionDeclarationAttributes) GetAvailableOnFrenchStore() bool`

GetAvailableOnFrenchStore returns the AvailableOnFrenchStore field if non-nil, zero value otherwise.

### GetAvailableOnFrenchStoreOk

`func (o *AppEncryptionDeclarationAttributes) GetAvailableOnFrenchStoreOk() (*bool, bool)`

GetAvailableOnFrenchStoreOk returns a tuple with the AvailableOnFrenchStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableOnFrenchStore

`func (o *AppEncryptionDeclarationAttributes) SetAvailableOnFrenchStore(v bool)`

SetAvailableOnFrenchStore sets AvailableOnFrenchStore field to given value.

### HasAvailableOnFrenchStore

`func (o *AppEncryptionDeclarationAttributes) HasAvailableOnFrenchStore() bool`

HasAvailableOnFrenchStore returns a boolean if a field has been set.

### GetPlatform

`func (o *AppEncryptionDeclarationAttributes) GetPlatform() Platform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AppEncryptionDeclarationAttributes) GetPlatformOk() (*Platform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AppEncryptionDeclarationAttributes) SetPlatform(v Platform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *AppEncryptionDeclarationAttributes) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetUploadedDate

`func (o *AppEncryptionDeclarationAttributes) GetUploadedDate() time.Time`

GetUploadedDate returns the UploadedDate field if non-nil, zero value otherwise.

### GetUploadedDateOk

`func (o *AppEncryptionDeclarationAttributes) GetUploadedDateOk() (*time.Time, bool)`

GetUploadedDateOk returns a tuple with the UploadedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedDate

`func (o *AppEncryptionDeclarationAttributes) SetUploadedDate(v time.Time)`

SetUploadedDate sets UploadedDate field to given value.

### HasUploadedDate

`func (o *AppEncryptionDeclarationAttributes) HasUploadedDate() bool`

HasUploadedDate returns a boolean if a field has been set.

### GetDocumentUrl

`func (o *AppEncryptionDeclarationAttributes) GetDocumentUrl() string`

GetDocumentUrl returns the DocumentUrl field if non-nil, zero value otherwise.

### GetDocumentUrlOk

`func (o *AppEncryptionDeclarationAttributes) GetDocumentUrlOk() (*string, bool)`

GetDocumentUrlOk returns a tuple with the DocumentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentUrl

`func (o *AppEncryptionDeclarationAttributes) SetDocumentUrl(v string)`

SetDocumentUrl sets DocumentUrl field to given value.

### HasDocumentUrl

`func (o *AppEncryptionDeclarationAttributes) HasDocumentUrl() bool`

HasDocumentUrl returns a boolean if a field has been set.

### GetDocumentName

`func (o *AppEncryptionDeclarationAttributes) GetDocumentName() string`

GetDocumentName returns the DocumentName field if non-nil, zero value otherwise.

### GetDocumentNameOk

`func (o *AppEncryptionDeclarationAttributes) GetDocumentNameOk() (*string, bool)`

GetDocumentNameOk returns a tuple with the DocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentName

`func (o *AppEncryptionDeclarationAttributes) SetDocumentName(v string)`

SetDocumentName sets DocumentName field to given value.

### HasDocumentName

`func (o *AppEncryptionDeclarationAttributes) HasDocumentName() bool`

HasDocumentName returns a boolean if a field has been set.

### GetDocumentType

`func (o *AppEncryptionDeclarationAttributes) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *AppEncryptionDeclarationAttributes) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *AppEncryptionDeclarationAttributes) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *AppEncryptionDeclarationAttributes) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetAppEncryptionDeclarationState

`func (o *AppEncryptionDeclarationAttributes) GetAppEncryptionDeclarationState() AppEncryptionDeclarationState`

GetAppEncryptionDeclarationState returns the AppEncryptionDeclarationState field if non-nil, zero value otherwise.

### GetAppEncryptionDeclarationStateOk

`func (o *AppEncryptionDeclarationAttributes) GetAppEncryptionDeclarationStateOk() (*AppEncryptionDeclarationState, bool)`

GetAppEncryptionDeclarationStateOk returns a tuple with the AppEncryptionDeclarationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEncryptionDeclarationState

`func (o *AppEncryptionDeclarationAttributes) SetAppEncryptionDeclarationState(v AppEncryptionDeclarationState)`

SetAppEncryptionDeclarationState sets AppEncryptionDeclarationState field to given value.

### HasAppEncryptionDeclarationState

`func (o *AppEncryptionDeclarationAttributes) HasAppEncryptionDeclarationState() bool`

HasAppEncryptionDeclarationState returns a boolean if a field has been set.

### GetCodeValue

`func (o *AppEncryptionDeclarationAttributes) GetCodeValue() string`

GetCodeValue returns the CodeValue field if non-nil, zero value otherwise.

### GetCodeValueOk

`func (o *AppEncryptionDeclarationAttributes) GetCodeValueOk() (*string, bool)`

GetCodeValueOk returns a tuple with the CodeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeValue

`func (o *AppEncryptionDeclarationAttributes) SetCodeValue(v string)`

SetCodeValue sets CodeValue field to given value.

### HasCodeValue

`func (o *AppEncryptionDeclarationAttributes) HasCodeValue() bool`

HasCodeValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


