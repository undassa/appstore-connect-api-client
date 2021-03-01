# AppPriceTiersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AppPriceTier**](AppPriceTier.md) |  | 
**Included** | Pointer to [**[]AppPricePoint**](AppPricePoint.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewAppPriceTiersResponse

`func NewAppPriceTiersResponse(data []AppPriceTier, links PagedDocumentLinks, ) *AppPriceTiersResponse`

NewAppPriceTiersResponse instantiates a new AppPriceTiersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPriceTiersResponseWithDefaults

`func NewAppPriceTiersResponseWithDefaults() *AppPriceTiersResponse`

NewAppPriceTiersResponseWithDefaults instantiates a new AppPriceTiersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppPriceTiersResponse) GetData() []AppPriceTier`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppPriceTiersResponse) GetDataOk() (*[]AppPriceTier, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppPriceTiersResponse) SetData(v []AppPriceTier)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AppPriceTiersResponse) GetIncluded() []AppPricePoint`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppPriceTiersResponse) GetIncludedOk() (*[]AppPricePoint, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppPriceTiersResponse) SetIncluded(v []AppPricePoint)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppPriceTiersResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *AppPriceTiersResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppPriceTiersResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppPriceTiersResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *AppPriceTiersResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AppPriceTiersResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AppPriceTiersResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AppPriceTiersResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


