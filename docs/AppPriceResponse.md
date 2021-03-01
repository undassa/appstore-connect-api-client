# AppPriceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AppPrice**](AppPrice.md) |  | 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewAppPriceResponse

`func NewAppPriceResponse(data AppPrice, links DocumentLinks, ) *AppPriceResponse`

NewAppPriceResponse instantiates a new AppPriceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPriceResponseWithDefaults

`func NewAppPriceResponseWithDefaults() *AppPriceResponse`

NewAppPriceResponseWithDefaults instantiates a new AppPriceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppPriceResponse) GetData() AppPrice`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppPriceResponse) GetDataOk() (*AppPrice, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppPriceResponse) SetData(v AppPrice)`

SetData sets Data field to given value.


### GetLinks

`func (o *AppPriceResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppPriceResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppPriceResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


