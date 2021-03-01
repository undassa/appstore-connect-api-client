# BetaTesterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**BetaTester**](BetaTester.md) |  | 
**Included** | Pointer to [**[]OneOfAppBetaGroupBuild**](OneOfAppBetaGroupBuild.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewBetaTesterResponse

`func NewBetaTesterResponse(data BetaTester, links DocumentLinks, ) *BetaTesterResponse`

NewBetaTesterResponse instantiates a new BetaTesterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaTesterResponseWithDefaults

`func NewBetaTesterResponseWithDefaults() *BetaTesterResponse`

NewBetaTesterResponseWithDefaults instantiates a new BetaTesterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BetaTesterResponse) GetData() BetaTester`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BetaTesterResponse) GetDataOk() (*BetaTester, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BetaTesterResponse) SetData(v BetaTester)`

SetData sets Data field to given value.


### GetIncluded

`func (o *BetaTesterResponse) GetIncluded() []OneOfAppBetaGroupBuild`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *BetaTesterResponse) GetIncludedOk() (*[]OneOfAppBetaGroupBuild, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *BetaTesterResponse) SetIncluded(v []OneOfAppBetaGroupBuild)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *BetaTesterResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *BetaTesterResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BetaTesterResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BetaTesterResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


