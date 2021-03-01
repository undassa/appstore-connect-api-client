# AppPreviewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AppPreview**](AppPreview.md) |  | 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewAppPreviewResponse

`func NewAppPreviewResponse(data AppPreview, links DocumentLinks, ) *AppPreviewResponse`

NewAppPreviewResponse instantiates a new AppPreviewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPreviewResponseWithDefaults

`func NewAppPreviewResponseWithDefaults() *AppPreviewResponse`

NewAppPreviewResponseWithDefaults instantiates a new AppPreviewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppPreviewResponse) GetData() AppPreview`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppPreviewResponse) GetDataOk() (*AppPreview, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppPreviewResponse) SetData(v AppPreview)`

SetData sets Data field to given value.


### GetLinks

`func (o *AppPreviewResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppPreviewResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppPreviewResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


