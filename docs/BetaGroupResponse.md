# BetaGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**BetaGroup**](BetaGroup.md) |  | 
**Included** | Pointer to [**[]OneOfAppBuildBetaTester**](OneOfAppBuildBetaTester.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewBetaGroupResponse

`func NewBetaGroupResponse(data BetaGroup, links DocumentLinks, ) *BetaGroupResponse`

NewBetaGroupResponse instantiates a new BetaGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaGroupResponseWithDefaults

`func NewBetaGroupResponseWithDefaults() *BetaGroupResponse`

NewBetaGroupResponseWithDefaults instantiates a new BetaGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BetaGroupResponse) GetData() BetaGroup`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BetaGroupResponse) GetDataOk() (*BetaGroup, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BetaGroupResponse) SetData(v BetaGroup)`

SetData sets Data field to given value.


### GetIncluded

`func (o *BetaGroupResponse) GetIncluded() []OneOfAppBuildBetaTester`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *BetaGroupResponse) GetIncludedOk() (*[]OneOfAppBuildBetaTester, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *BetaGroupResponse) SetIncluded(v []OneOfAppBuildBetaTester)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *BetaGroupResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *BetaGroupResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BetaGroupResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BetaGroupResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


