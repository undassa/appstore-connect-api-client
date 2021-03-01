# AppPreviewCreateRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileSize** | **int32** |  | 
**FileName** | **string** |  | 
**PreviewFrameTimeCode** | Pointer to **string** |  | [optional] 
**MimeType** | Pointer to **string** |  | [optional] 

## Methods

### NewAppPreviewCreateRequestDataAttributes

`func NewAppPreviewCreateRequestDataAttributes(fileSize int32, fileName string, ) *AppPreviewCreateRequestDataAttributes`

NewAppPreviewCreateRequestDataAttributes instantiates a new AppPreviewCreateRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPreviewCreateRequestDataAttributesWithDefaults

`func NewAppPreviewCreateRequestDataAttributesWithDefaults() *AppPreviewCreateRequestDataAttributes`

NewAppPreviewCreateRequestDataAttributesWithDefaults instantiates a new AppPreviewCreateRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileSize

`func (o *AppPreviewCreateRequestDataAttributes) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *AppPreviewCreateRequestDataAttributes) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *AppPreviewCreateRequestDataAttributes) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.


### GetFileName

`func (o *AppPreviewCreateRequestDataAttributes) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *AppPreviewCreateRequestDataAttributes) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *AppPreviewCreateRequestDataAttributes) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetPreviewFrameTimeCode

`func (o *AppPreviewCreateRequestDataAttributes) GetPreviewFrameTimeCode() string`

GetPreviewFrameTimeCode returns the PreviewFrameTimeCode field if non-nil, zero value otherwise.

### GetPreviewFrameTimeCodeOk

`func (o *AppPreviewCreateRequestDataAttributes) GetPreviewFrameTimeCodeOk() (*string, bool)`

GetPreviewFrameTimeCodeOk returns a tuple with the PreviewFrameTimeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewFrameTimeCode

`func (o *AppPreviewCreateRequestDataAttributes) SetPreviewFrameTimeCode(v string)`

SetPreviewFrameTimeCode sets PreviewFrameTimeCode field to given value.

### HasPreviewFrameTimeCode

`func (o *AppPreviewCreateRequestDataAttributes) HasPreviewFrameTimeCode() bool`

HasPreviewFrameTimeCode returns a boolean if a field has been set.

### GetMimeType

`func (o *AppPreviewCreateRequestDataAttributes) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *AppPreviewCreateRequestDataAttributes) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *AppPreviewCreateRequestDataAttributes) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *AppPreviewCreateRequestDataAttributes) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


