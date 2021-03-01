# BuildBetaDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**BuildBetaDetail**](BuildBetaDetail.md) |  | 
**Included** | Pointer to [**[]Build**](Build.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewBuildBetaDetailResponse

`func NewBuildBetaDetailResponse(data BuildBetaDetail, links DocumentLinks, ) *BuildBetaDetailResponse`

NewBuildBetaDetailResponse instantiates a new BuildBetaDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildBetaDetailResponseWithDefaults

`func NewBuildBetaDetailResponseWithDefaults() *BuildBetaDetailResponse`

NewBuildBetaDetailResponseWithDefaults instantiates a new BuildBetaDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BuildBetaDetailResponse) GetData() BuildBetaDetail`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BuildBetaDetailResponse) GetDataOk() (*BuildBetaDetail, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BuildBetaDetailResponse) SetData(v BuildBetaDetail)`

SetData sets Data field to given value.


### GetIncluded

`func (o *BuildBetaDetailResponse) GetIncluded() []Build`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *BuildBetaDetailResponse) GetIncludedOk() (*[]Build, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *BuildBetaDetailResponse) SetIncluded(v []Build)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *BuildBetaDetailResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *BuildBetaDetailResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BuildBetaDetailResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BuildBetaDetailResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


