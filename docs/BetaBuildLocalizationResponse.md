# BetaBuildLocalizationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**BetaBuildLocalization**](BetaBuildLocalization.md) |  | 
**Included** | Pointer to [**[]Build**](Build.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewBetaBuildLocalizationResponse

`func NewBetaBuildLocalizationResponse(data BetaBuildLocalization, links DocumentLinks, ) *BetaBuildLocalizationResponse`

NewBetaBuildLocalizationResponse instantiates a new BetaBuildLocalizationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaBuildLocalizationResponseWithDefaults

`func NewBetaBuildLocalizationResponseWithDefaults() *BetaBuildLocalizationResponse`

NewBetaBuildLocalizationResponseWithDefaults instantiates a new BetaBuildLocalizationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BetaBuildLocalizationResponse) GetData() BetaBuildLocalization`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BetaBuildLocalizationResponse) GetDataOk() (*BetaBuildLocalization, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BetaBuildLocalizationResponse) SetData(v BetaBuildLocalization)`

SetData sets Data field to given value.


### GetIncluded

`func (o *BetaBuildLocalizationResponse) GetIncluded() []Build`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *BetaBuildLocalizationResponse) GetIncludedOk() (*[]Build, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *BetaBuildLocalizationResponse) SetIncluded(v []Build)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *BetaBuildLocalizationResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *BetaBuildLocalizationResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BetaBuildLocalizationResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BetaBuildLocalizationResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


