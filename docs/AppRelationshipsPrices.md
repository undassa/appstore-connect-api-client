# AppRelationshipsPrices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AppCategoryRelationshipsSubcategoriesLinks**](AppCategory_relationships_subcategories_links.md) |  | [optional] 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 
**Data** | Pointer to [**[]AppRelationshipsPricesData**](AppRelationshipsPricesData.md) |  | [optional] 

## Methods

### NewAppRelationshipsPrices

`func NewAppRelationshipsPrices() *AppRelationshipsPrices`

NewAppRelationshipsPrices instantiates a new AppRelationshipsPrices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRelationshipsPricesWithDefaults

`func NewAppRelationshipsPricesWithDefaults() *AppRelationshipsPrices`

NewAppRelationshipsPricesWithDefaults instantiates a new AppRelationshipsPrices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AppRelationshipsPrices) GetLinks() AppCategoryRelationshipsSubcategoriesLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppRelationshipsPrices) GetLinksOk() (*AppCategoryRelationshipsSubcategoriesLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppRelationshipsPrices) SetLinks(v AppCategoryRelationshipsSubcategoriesLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AppRelationshipsPrices) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *AppRelationshipsPrices) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AppRelationshipsPrices) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AppRelationshipsPrices) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AppRelationshipsPrices) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *AppRelationshipsPrices) GetData() []AppRelationshipsPricesData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppRelationshipsPrices) GetDataOk() (*[]AppRelationshipsPricesData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppRelationshipsPrices) SetData(v []AppRelationshipsPricesData)`

SetData sets Data field to given value.

### HasData

`func (o *AppRelationshipsPrices) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


