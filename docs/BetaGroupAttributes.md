# BetaGroupAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**CreatedDate** | Pointer to **time.Time** |  | [optional] 
**IsInternalGroup** | Pointer to **bool** |  | [optional] 
**PublicLinkEnabled** | Pointer to **bool** |  | [optional] 
**PublicLinkId** | Pointer to **string** |  | [optional] 
**PublicLinkLimitEnabled** | Pointer to **bool** |  | [optional] 
**PublicLinkLimit** | Pointer to **int32** |  | [optional] 
**PublicLink** | Pointer to **string** |  | [optional] 
**FeedbackEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewBetaGroupAttributes

`func NewBetaGroupAttributes() *BetaGroupAttributes`

NewBetaGroupAttributes instantiates a new BetaGroupAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaGroupAttributesWithDefaults

`func NewBetaGroupAttributesWithDefaults() *BetaGroupAttributes`

NewBetaGroupAttributesWithDefaults instantiates a new BetaGroupAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BetaGroupAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BetaGroupAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BetaGroupAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BetaGroupAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedDate

`func (o *BetaGroupAttributes) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *BetaGroupAttributes) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *BetaGroupAttributes) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *BetaGroupAttributes) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetIsInternalGroup

`func (o *BetaGroupAttributes) GetIsInternalGroup() bool`

GetIsInternalGroup returns the IsInternalGroup field if non-nil, zero value otherwise.

### GetIsInternalGroupOk

`func (o *BetaGroupAttributes) GetIsInternalGroupOk() (*bool, bool)`

GetIsInternalGroupOk returns a tuple with the IsInternalGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternalGroup

`func (o *BetaGroupAttributes) SetIsInternalGroup(v bool)`

SetIsInternalGroup sets IsInternalGroup field to given value.

### HasIsInternalGroup

`func (o *BetaGroupAttributes) HasIsInternalGroup() bool`

HasIsInternalGroup returns a boolean if a field has been set.

### GetPublicLinkEnabled

`func (o *BetaGroupAttributes) GetPublicLinkEnabled() bool`

GetPublicLinkEnabled returns the PublicLinkEnabled field if non-nil, zero value otherwise.

### GetPublicLinkEnabledOk

`func (o *BetaGroupAttributes) GetPublicLinkEnabledOk() (*bool, bool)`

GetPublicLinkEnabledOk returns a tuple with the PublicLinkEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicLinkEnabled

`func (o *BetaGroupAttributes) SetPublicLinkEnabled(v bool)`

SetPublicLinkEnabled sets PublicLinkEnabled field to given value.

### HasPublicLinkEnabled

`func (o *BetaGroupAttributes) HasPublicLinkEnabled() bool`

HasPublicLinkEnabled returns a boolean if a field has been set.

### GetPublicLinkId

`func (o *BetaGroupAttributes) GetPublicLinkId() string`

GetPublicLinkId returns the PublicLinkId field if non-nil, zero value otherwise.

### GetPublicLinkIdOk

`func (o *BetaGroupAttributes) GetPublicLinkIdOk() (*string, bool)`

GetPublicLinkIdOk returns a tuple with the PublicLinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicLinkId

`func (o *BetaGroupAttributes) SetPublicLinkId(v string)`

SetPublicLinkId sets PublicLinkId field to given value.

### HasPublicLinkId

`func (o *BetaGroupAttributes) HasPublicLinkId() bool`

HasPublicLinkId returns a boolean if a field has been set.

### GetPublicLinkLimitEnabled

`func (o *BetaGroupAttributes) GetPublicLinkLimitEnabled() bool`

GetPublicLinkLimitEnabled returns the PublicLinkLimitEnabled field if non-nil, zero value otherwise.

### GetPublicLinkLimitEnabledOk

`func (o *BetaGroupAttributes) GetPublicLinkLimitEnabledOk() (*bool, bool)`

GetPublicLinkLimitEnabledOk returns a tuple with the PublicLinkLimitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicLinkLimitEnabled

`func (o *BetaGroupAttributes) SetPublicLinkLimitEnabled(v bool)`

SetPublicLinkLimitEnabled sets PublicLinkLimitEnabled field to given value.

### HasPublicLinkLimitEnabled

`func (o *BetaGroupAttributes) HasPublicLinkLimitEnabled() bool`

HasPublicLinkLimitEnabled returns a boolean if a field has been set.

### GetPublicLinkLimit

`func (o *BetaGroupAttributes) GetPublicLinkLimit() int32`

GetPublicLinkLimit returns the PublicLinkLimit field if non-nil, zero value otherwise.

### GetPublicLinkLimitOk

`func (o *BetaGroupAttributes) GetPublicLinkLimitOk() (*int32, bool)`

GetPublicLinkLimitOk returns a tuple with the PublicLinkLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicLinkLimit

`func (o *BetaGroupAttributes) SetPublicLinkLimit(v int32)`

SetPublicLinkLimit sets PublicLinkLimit field to given value.

### HasPublicLinkLimit

`func (o *BetaGroupAttributes) HasPublicLinkLimit() bool`

HasPublicLinkLimit returns a boolean if a field has been set.

### GetPublicLink

`func (o *BetaGroupAttributes) GetPublicLink() string`

GetPublicLink returns the PublicLink field if non-nil, zero value otherwise.

### GetPublicLinkOk

`func (o *BetaGroupAttributes) GetPublicLinkOk() (*string, bool)`

GetPublicLinkOk returns a tuple with the PublicLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicLink

`func (o *BetaGroupAttributes) SetPublicLink(v string)`

SetPublicLink sets PublicLink field to given value.

### HasPublicLink

`func (o *BetaGroupAttributes) HasPublicLink() bool`

HasPublicLink returns a boolean if a field has been set.

### GetFeedbackEnabled

`func (o *BetaGroupAttributes) GetFeedbackEnabled() bool`

GetFeedbackEnabled returns the FeedbackEnabled field if non-nil, zero value otherwise.

### GetFeedbackEnabledOk

`func (o *BetaGroupAttributes) GetFeedbackEnabledOk() (*bool, bool)`

GetFeedbackEnabledOk returns a tuple with the FeedbackEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackEnabled

`func (o *BetaGroupAttributes) SetFeedbackEnabled(v bool)`

SetFeedbackEnabled sets FeedbackEnabled field to given value.

### HasFeedbackEnabled

`func (o *BetaGroupAttributes) HasFeedbackEnabled() bool`

HasFeedbackEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


