# AppPreOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AppPreOrder**](AppPreOrder.md) |  | 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewAppPreOrderResponse

`func NewAppPreOrderResponse(data AppPreOrder, links DocumentLinks, ) *AppPreOrderResponse`

NewAppPreOrderResponse instantiates a new AppPreOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPreOrderResponseWithDefaults

`func NewAppPreOrderResponseWithDefaults() *AppPreOrderResponse`

NewAppPreOrderResponseWithDefaults instantiates a new AppPreOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppPreOrderResponse) GetData() AppPreOrder`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppPreOrderResponse) GetDataOk() (*AppPreOrder, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppPreOrderResponse) SetData(v AppPreOrder)`

SetData sets Data field to given value.


### GetLinks

`func (o *AppPreOrderResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppPreOrderResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppPreOrderResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


