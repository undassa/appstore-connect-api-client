# AppMediaAssetState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]AppMediaStateError**](AppMediaStateError.md) |  | [optional] 
**Warnings** | Pointer to [**[]AppMediaStateError**](AppMediaStateError.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewAppMediaAssetState

`func NewAppMediaAssetState() *AppMediaAssetState`

NewAppMediaAssetState instantiates a new AppMediaAssetState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppMediaAssetStateWithDefaults

`func NewAppMediaAssetStateWithDefaults() *AppMediaAssetState`

NewAppMediaAssetStateWithDefaults instantiates a new AppMediaAssetState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *AppMediaAssetState) GetErrors() []AppMediaStateError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AppMediaAssetState) GetErrorsOk() (*[]AppMediaStateError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *AppMediaAssetState) SetErrors(v []AppMediaStateError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *AppMediaAssetState) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *AppMediaAssetState) GetWarnings() []AppMediaStateError`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *AppMediaAssetState) GetWarningsOk() (*[]AppMediaStateError, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *AppMediaAssetState) SetWarnings(v []AppMediaStateError)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *AppMediaAssetState) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetState

`func (o *AppMediaAssetState) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppMediaAssetState) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppMediaAssetState) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AppMediaAssetState) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


