# AppInfoRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to [**AppEncryptionDeclarationRelationshipsApp**](AppEncryptionDeclaration_relationships_app.md) |  | [optional] 
**AppInfoLocalizations** | Pointer to [**AppInfoRelationshipsAppInfoLocalizations**](AppInfo_relationships_appInfoLocalizations.md) |  | [optional] 
**PrimaryCategory** | Pointer to [**AppCategoryRelationshipsParent**](AppCategory_relationships_parent.md) |  | [optional] 
**PrimarySubcategoryOne** | Pointer to [**AppCategoryRelationshipsParent**](AppCategory_relationships_parent.md) |  | [optional] 
**PrimarySubcategoryTwo** | Pointer to [**AppCategoryRelationshipsParent**](AppCategory_relationships_parent.md) |  | [optional] 
**SecondaryCategory** | Pointer to [**AppCategoryRelationshipsParent**](AppCategory_relationships_parent.md) |  | [optional] 
**SecondarySubcategoryOne** | Pointer to [**AppCategoryRelationshipsParent**](AppCategory_relationships_parent.md) |  | [optional] 
**SecondarySubcategoryTwo** | Pointer to [**AppCategoryRelationshipsParent**](AppCategory_relationships_parent.md) |  | [optional] 

## Methods

### NewAppInfoRelationships

`func NewAppInfoRelationships() *AppInfoRelationships`

NewAppInfoRelationships instantiates a new AppInfoRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppInfoRelationshipsWithDefaults

`func NewAppInfoRelationshipsWithDefaults() *AppInfoRelationships`

NewAppInfoRelationshipsWithDefaults instantiates a new AppInfoRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *AppInfoRelationships) GetApp() AppEncryptionDeclarationRelationshipsApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *AppInfoRelationships) GetAppOk() (*AppEncryptionDeclarationRelationshipsApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *AppInfoRelationships) SetApp(v AppEncryptionDeclarationRelationshipsApp)`

SetApp sets App field to given value.

### HasApp

`func (o *AppInfoRelationships) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetAppInfoLocalizations

`func (o *AppInfoRelationships) GetAppInfoLocalizations() AppInfoRelationshipsAppInfoLocalizations`

GetAppInfoLocalizations returns the AppInfoLocalizations field if non-nil, zero value otherwise.

### GetAppInfoLocalizationsOk

`func (o *AppInfoRelationships) GetAppInfoLocalizationsOk() (*AppInfoRelationshipsAppInfoLocalizations, bool)`

GetAppInfoLocalizationsOk returns a tuple with the AppInfoLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInfoLocalizations

`func (o *AppInfoRelationships) SetAppInfoLocalizations(v AppInfoRelationshipsAppInfoLocalizations)`

SetAppInfoLocalizations sets AppInfoLocalizations field to given value.

### HasAppInfoLocalizations

`func (o *AppInfoRelationships) HasAppInfoLocalizations() bool`

HasAppInfoLocalizations returns a boolean if a field has been set.

### GetPrimaryCategory

`func (o *AppInfoRelationships) GetPrimaryCategory() AppCategoryRelationshipsParent`

GetPrimaryCategory returns the PrimaryCategory field if non-nil, zero value otherwise.

### GetPrimaryCategoryOk

`func (o *AppInfoRelationships) GetPrimaryCategoryOk() (*AppCategoryRelationshipsParent, bool)`

GetPrimaryCategoryOk returns a tuple with the PrimaryCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCategory

`func (o *AppInfoRelationships) SetPrimaryCategory(v AppCategoryRelationshipsParent)`

SetPrimaryCategory sets PrimaryCategory field to given value.

### HasPrimaryCategory

`func (o *AppInfoRelationships) HasPrimaryCategory() bool`

HasPrimaryCategory returns a boolean if a field has been set.

### GetPrimarySubcategoryOne

`func (o *AppInfoRelationships) GetPrimarySubcategoryOne() AppCategoryRelationshipsParent`

GetPrimarySubcategoryOne returns the PrimarySubcategoryOne field if non-nil, zero value otherwise.

### GetPrimarySubcategoryOneOk

`func (o *AppInfoRelationships) GetPrimarySubcategoryOneOk() (*AppCategoryRelationshipsParent, bool)`

GetPrimarySubcategoryOneOk returns a tuple with the PrimarySubcategoryOne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySubcategoryOne

`func (o *AppInfoRelationships) SetPrimarySubcategoryOne(v AppCategoryRelationshipsParent)`

SetPrimarySubcategoryOne sets PrimarySubcategoryOne field to given value.

### HasPrimarySubcategoryOne

`func (o *AppInfoRelationships) HasPrimarySubcategoryOne() bool`

HasPrimarySubcategoryOne returns a boolean if a field has been set.

### GetPrimarySubcategoryTwo

`func (o *AppInfoRelationships) GetPrimarySubcategoryTwo() AppCategoryRelationshipsParent`

GetPrimarySubcategoryTwo returns the PrimarySubcategoryTwo field if non-nil, zero value otherwise.

### GetPrimarySubcategoryTwoOk

`func (o *AppInfoRelationships) GetPrimarySubcategoryTwoOk() (*AppCategoryRelationshipsParent, bool)`

GetPrimarySubcategoryTwoOk returns a tuple with the PrimarySubcategoryTwo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySubcategoryTwo

`func (o *AppInfoRelationships) SetPrimarySubcategoryTwo(v AppCategoryRelationshipsParent)`

SetPrimarySubcategoryTwo sets PrimarySubcategoryTwo field to given value.

### HasPrimarySubcategoryTwo

`func (o *AppInfoRelationships) HasPrimarySubcategoryTwo() bool`

HasPrimarySubcategoryTwo returns a boolean if a field has been set.

### GetSecondaryCategory

`func (o *AppInfoRelationships) GetSecondaryCategory() AppCategoryRelationshipsParent`

GetSecondaryCategory returns the SecondaryCategory field if non-nil, zero value otherwise.

### GetSecondaryCategoryOk

`func (o *AppInfoRelationships) GetSecondaryCategoryOk() (*AppCategoryRelationshipsParent, bool)`

GetSecondaryCategoryOk returns a tuple with the SecondaryCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCategory

`func (o *AppInfoRelationships) SetSecondaryCategory(v AppCategoryRelationshipsParent)`

SetSecondaryCategory sets SecondaryCategory field to given value.

### HasSecondaryCategory

`func (o *AppInfoRelationships) HasSecondaryCategory() bool`

HasSecondaryCategory returns a boolean if a field has been set.

### GetSecondarySubcategoryOne

`func (o *AppInfoRelationships) GetSecondarySubcategoryOne() AppCategoryRelationshipsParent`

GetSecondarySubcategoryOne returns the SecondarySubcategoryOne field if non-nil, zero value otherwise.

### GetSecondarySubcategoryOneOk

`func (o *AppInfoRelationships) GetSecondarySubcategoryOneOk() (*AppCategoryRelationshipsParent, bool)`

GetSecondarySubcategoryOneOk returns a tuple with the SecondarySubcategoryOne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarySubcategoryOne

`func (o *AppInfoRelationships) SetSecondarySubcategoryOne(v AppCategoryRelationshipsParent)`

SetSecondarySubcategoryOne sets SecondarySubcategoryOne field to given value.

### HasSecondarySubcategoryOne

`func (o *AppInfoRelationships) HasSecondarySubcategoryOne() bool`

HasSecondarySubcategoryOne returns a boolean if a field has been set.

### GetSecondarySubcategoryTwo

`func (o *AppInfoRelationships) GetSecondarySubcategoryTwo() AppCategoryRelationshipsParent`

GetSecondarySubcategoryTwo returns the SecondarySubcategoryTwo field if non-nil, zero value otherwise.

### GetSecondarySubcategoryTwoOk

`func (o *AppInfoRelationships) GetSecondarySubcategoryTwoOk() (*AppCategoryRelationshipsParent, bool)`

GetSecondarySubcategoryTwoOk returns a tuple with the SecondarySubcategoryTwo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarySubcategoryTwo

`func (o *AppInfoRelationships) SetSecondarySubcategoryTwo(v AppCategoryRelationshipsParent)`

SetSecondarySubcategoryTwo sets SecondarySubcategoryTwo field to given value.

### HasSecondarySubcategoryTwo

`func (o *AppInfoRelationships) HasSecondarySubcategoryTwo() bool`

HasSecondarySubcategoryTwo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


