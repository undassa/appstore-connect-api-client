# AppPriceTierResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AppPriceTier**](AppPriceTier.md) |  | 
**Included** | Pointer to [**[]AppPricePoint**](AppPricePoint.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewAppPriceTierResponse

`func NewAppPriceTierResponse(data AppPriceTier, links DocumentLinks, ) *AppPriceTierResponse`

NewAppPriceTierResponse instantiates a new AppPriceTierResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPriceTierResponseWithDefaults

`func NewAppPriceTierResponseWithDefaults() *AppPriceTierResponse`

NewAppPriceTierResponseWithDefaults instantiates a new AppPriceTierResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppPriceTierResponse) GetData() AppPriceTier`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppPriceTierResponse) GetDataOk() (*AppPriceTier, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppPriceTierResponse) SetData(v AppPriceTier)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AppPriceTierResponse) GetIncluded() []AppPricePoint`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppPriceTierResponse) GetIncludedOk() (*[]AppPricePoint, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppPriceTierResponse) SetIncluded(v []AppPricePoint)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppPriceTierResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *AppPriceTierResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppPriceTierResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppPriceTierResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


