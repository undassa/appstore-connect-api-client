# BuildBetaDetailAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoNotifyEnabled** | Pointer to **bool** |  | [optional] 
**InternalBuildState** | Pointer to [**InternalBetaState**](InternalBetaState.md) |  | [optional] 
**ExternalBuildState** | Pointer to [**ExternalBetaState**](ExternalBetaState.md) |  | [optional] 

## Methods

### NewBuildBetaDetailAttributes

`func NewBuildBetaDetailAttributes() *BuildBetaDetailAttributes`

NewBuildBetaDetailAttributes instantiates a new BuildBetaDetailAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildBetaDetailAttributesWithDefaults

`func NewBuildBetaDetailAttributesWithDefaults() *BuildBetaDetailAttributes`

NewBuildBetaDetailAttributesWithDefaults instantiates a new BuildBetaDetailAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoNotifyEnabled

`func (o *BuildBetaDetailAttributes) GetAutoNotifyEnabled() bool`

GetAutoNotifyEnabled returns the AutoNotifyEnabled field if non-nil, zero value otherwise.

### GetAutoNotifyEnabledOk

`func (o *BuildBetaDetailAttributes) GetAutoNotifyEnabledOk() (*bool, bool)`

GetAutoNotifyEnabledOk returns a tuple with the AutoNotifyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoNotifyEnabled

`func (o *BuildBetaDetailAttributes) SetAutoNotifyEnabled(v bool)`

SetAutoNotifyEnabled sets AutoNotifyEnabled field to given value.

### HasAutoNotifyEnabled

`func (o *BuildBetaDetailAttributes) HasAutoNotifyEnabled() bool`

HasAutoNotifyEnabled returns a boolean if a field has been set.

### GetInternalBuildState

`func (o *BuildBetaDetailAttributes) GetInternalBuildState() InternalBetaState`

GetInternalBuildState returns the InternalBuildState field if non-nil, zero value otherwise.

### GetInternalBuildStateOk

`func (o *BuildBetaDetailAttributes) GetInternalBuildStateOk() (*InternalBetaState, bool)`

GetInternalBuildStateOk returns a tuple with the InternalBuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalBuildState

`func (o *BuildBetaDetailAttributes) SetInternalBuildState(v InternalBetaState)`

SetInternalBuildState sets InternalBuildState field to given value.

### HasInternalBuildState

`func (o *BuildBetaDetailAttributes) HasInternalBuildState() bool`

HasInternalBuildState returns a boolean if a field has been set.

### GetExternalBuildState

`func (o *BuildBetaDetailAttributes) GetExternalBuildState() ExternalBetaState`

GetExternalBuildState returns the ExternalBuildState field if non-nil, zero value otherwise.

### GetExternalBuildStateOk

`func (o *BuildBetaDetailAttributes) GetExternalBuildStateOk() (*ExternalBetaState, bool)`

GetExternalBuildStateOk returns a tuple with the ExternalBuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalBuildState

`func (o *BuildBetaDetailAttributes) SetExternalBuildState(v ExternalBetaState)`

SetExternalBuildState sets ExternalBuildState field to given value.

### HasExternalBuildState

`func (o *BuildBetaDetailAttributes) HasExternalBuildState() bool`

HasExternalBuildState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


