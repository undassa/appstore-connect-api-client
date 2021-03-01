# UploadOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Length** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**RequestHeaders** | Pointer to [**[]UploadOperationHeader**](UploadOperationHeader.md) |  | [optional] 

## Methods

### NewUploadOperation

`func NewUploadOperation() *UploadOperation`

NewUploadOperation instantiates a new UploadOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadOperationWithDefaults

`func NewUploadOperationWithDefaults() *UploadOperation`

NewUploadOperationWithDefaults instantiates a new UploadOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *UploadOperation) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *UploadOperation) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *UploadOperation) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *UploadOperation) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetUrl

`func (o *UploadOperation) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UploadOperation) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UploadOperation) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UploadOperation) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetLength

`func (o *UploadOperation) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *UploadOperation) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *UploadOperation) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *UploadOperation) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetOffset

`func (o *UploadOperation) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *UploadOperation) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *UploadOperation) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *UploadOperation) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetRequestHeaders

`func (o *UploadOperation) GetRequestHeaders() []UploadOperationHeader`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *UploadOperation) GetRequestHeadersOk() (*[]UploadOperationHeader, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaders

`func (o *UploadOperation) SetRequestHeaders(v []UploadOperationHeader)`

SetRequestHeaders sets RequestHeaders field to given value.

### HasRequestHeaders

`func (o *UploadOperation) HasRequestHeaders() bool`

HasRequestHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


