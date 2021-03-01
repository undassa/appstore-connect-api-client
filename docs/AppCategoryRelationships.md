# AppCategoryRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subcategories** | Pointer to [**AppCategoryRelationshipsSubcategories**](AppCategory_relationships_subcategories.md) |  | [optional] 
**Parent** | Pointer to [**AppCategoryRelationshipsParent**](AppCategory_relationships_parent.md) |  | [optional] 

## Methods

### NewAppCategoryRelationships

`func NewAppCategoryRelationships() *AppCategoryRelationships`

NewAppCategoryRelationships instantiates a new AppCategoryRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppCategoryRelationshipsWithDefaults

`func NewAppCategoryRelationshipsWithDefaults() *AppCategoryRelationships`

NewAppCategoryRelationshipsWithDefaults instantiates a new AppCategoryRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubcategories

`func (o *AppCategoryRelationships) GetSubcategories() AppCategoryRelationshipsSubcategories`

GetSubcategories returns the Subcategories field if non-nil, zero value otherwise.

### GetSubcategoriesOk

`func (o *AppCategoryRelationships) GetSubcategoriesOk() (*AppCategoryRelationshipsSubcategories, bool)`

GetSubcategoriesOk returns a tuple with the Subcategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubcategories

`func (o *AppCategoryRelationships) SetSubcategories(v AppCategoryRelationshipsSubcategories)`

SetSubcategories sets Subcategories field to given value.

### HasSubcategories

`func (o *AppCategoryRelationships) HasSubcategories() bool`

HasSubcategories returns a boolean if a field has been set.

### GetParent

`func (o *AppCategoryRelationships) GetParent() AppCategoryRelationshipsParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *AppCategoryRelationships) GetParentOk() (*AppCategoryRelationshipsParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *AppCategoryRelationships) SetParent(v AppCategoryRelationshipsParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *AppCategoryRelationships) HasParent() bool`

HasParent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


