# AppPreviewSetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AppPreviewSet**](AppPreviewSet.md) |  | 
**Included** | Pointer to [**[]AppPreview**](AppPreview.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewAppPreviewSetResponse

`func NewAppPreviewSetResponse(data AppPreviewSet, links DocumentLinks, ) *AppPreviewSetResponse`

NewAppPreviewSetResponse instantiates a new AppPreviewSetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPreviewSetResponseWithDefaults

`func NewAppPreviewSetResponseWithDefaults() *AppPreviewSetResponse`

NewAppPreviewSetResponseWithDefaults instantiates a new AppPreviewSetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppPreviewSetResponse) GetData() AppPreviewSet`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppPreviewSetResponse) GetDataOk() (*AppPreviewSet, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppPreviewSetResponse) SetData(v AppPreviewSet)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AppPreviewSetResponse) GetIncluded() []AppPreview`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppPreviewSetResponse) GetIncludedOk() (*[]AppPreview, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppPreviewSetResponse) SetIncluded(v []AppPreview)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppPreviewSetResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *AppPreviewSetResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppPreviewSetResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppPreviewSetResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


