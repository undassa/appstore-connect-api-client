# ErrorResponseErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**Code** | **string** |  | 
**Title** | **string** |  | 
**Detail** | **string** |  | 
**Source** | Pointer to [**OneOfobjectobject**](oneOf&lt;object,object&gt;.md) |  | [optional] 

## Methods

### NewErrorResponseErrors

`func NewErrorResponseErrors(status string, code string, title string, detail string, ) *ErrorResponseErrors`

NewErrorResponseErrors instantiates a new ErrorResponseErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseErrorsWithDefaults

`func NewErrorResponseErrorsWithDefaults() *ErrorResponseErrors`

NewErrorResponseErrorsWithDefaults instantiates a new ErrorResponseErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ErrorResponseErrors) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ErrorResponseErrors) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ErrorResponseErrors) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ErrorResponseErrors) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ErrorResponseErrors) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ErrorResponseErrors) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ErrorResponseErrors) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCode

`func (o *ErrorResponseErrors) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorResponseErrors) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorResponseErrors) SetCode(v string)`

SetCode sets Code field to given value.


### GetTitle

`func (o *ErrorResponseErrors) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ErrorResponseErrors) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ErrorResponseErrors) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *ErrorResponseErrors) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ErrorResponseErrors) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ErrorResponseErrors) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetSource

`func (o *ErrorResponseErrors) GetSource() OneOfobjectobject`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ErrorResponseErrors) GetSourceOk() (*OneOfobjectobject, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ErrorResponseErrors) SetSource(v OneOfobjectobject)`

SetSource sets Source field to given value.

### HasSource

`func (o *ErrorResponseErrors) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


