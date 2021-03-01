# AppPreviewAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileSize** | Pointer to **int32** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**SourceFileChecksum** | Pointer to **string** |  | [optional] 
**PreviewFrameTimeCode** | Pointer to **string** |  | [optional] 
**MimeType** | Pointer to **string** |  | [optional] 
**VideoUrl** | Pointer to **string** |  | [optional] 
**PreviewImage** | Pointer to [**ImageAsset**](ImageAsset.md) |  | [optional] 
**UploadOperations** | Pointer to [**[]UploadOperation**](UploadOperation.md) |  | [optional] 
**AssetDeliveryState** | Pointer to [**AppMediaAssetState**](AppMediaAssetState.md) |  | [optional] 

## Methods

### NewAppPreviewAttributes

`func NewAppPreviewAttributes() *AppPreviewAttributes`

NewAppPreviewAttributes instantiates a new AppPreviewAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPreviewAttributesWithDefaults

`func NewAppPreviewAttributesWithDefaults() *AppPreviewAttributes`

NewAppPreviewAttributesWithDefaults instantiates a new AppPreviewAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileSize

`func (o *AppPreviewAttributes) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *AppPreviewAttributes) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *AppPreviewAttributes) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *AppPreviewAttributes) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetFileName

`func (o *AppPreviewAttributes) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *AppPreviewAttributes) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *AppPreviewAttributes) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *AppPreviewAttributes) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetSourceFileChecksum

`func (o *AppPreviewAttributes) GetSourceFileChecksum() string`

GetSourceFileChecksum returns the SourceFileChecksum field if non-nil, zero value otherwise.

### GetSourceFileChecksumOk

`func (o *AppPreviewAttributes) GetSourceFileChecksumOk() (*string, bool)`

GetSourceFileChecksumOk returns a tuple with the SourceFileChecksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFileChecksum

`func (o *AppPreviewAttributes) SetSourceFileChecksum(v string)`

SetSourceFileChecksum sets SourceFileChecksum field to given value.

### HasSourceFileChecksum

`func (o *AppPreviewAttributes) HasSourceFileChecksum() bool`

HasSourceFileChecksum returns a boolean if a field has been set.

### GetPreviewFrameTimeCode

`func (o *AppPreviewAttributes) GetPreviewFrameTimeCode() string`

GetPreviewFrameTimeCode returns the PreviewFrameTimeCode field if non-nil, zero value otherwise.

### GetPreviewFrameTimeCodeOk

`func (o *AppPreviewAttributes) GetPreviewFrameTimeCodeOk() (*string, bool)`

GetPreviewFrameTimeCodeOk returns a tuple with the PreviewFrameTimeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewFrameTimeCode

`func (o *AppPreviewAttributes) SetPreviewFrameTimeCode(v string)`

SetPreviewFrameTimeCode sets PreviewFrameTimeCode field to given value.

### HasPreviewFrameTimeCode

`func (o *AppPreviewAttributes) HasPreviewFrameTimeCode() bool`

HasPreviewFrameTimeCode returns a boolean if a field has been set.

### GetMimeType

`func (o *AppPreviewAttributes) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *AppPreviewAttributes) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *AppPreviewAttributes) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *AppPreviewAttributes) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetVideoUrl

`func (o *AppPreviewAttributes) GetVideoUrl() string`

GetVideoUrl returns the VideoUrl field if non-nil, zero value otherwise.

### GetVideoUrlOk

`func (o *AppPreviewAttributes) GetVideoUrlOk() (*string, bool)`

GetVideoUrlOk returns a tuple with the VideoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoUrl

`func (o *AppPreviewAttributes) SetVideoUrl(v string)`

SetVideoUrl sets VideoUrl field to given value.

### HasVideoUrl

`func (o *AppPreviewAttributes) HasVideoUrl() bool`

HasVideoUrl returns a boolean if a field has been set.

### GetPreviewImage

`func (o *AppPreviewAttributes) GetPreviewImage() ImageAsset`

GetPreviewImage returns the PreviewImage field if non-nil, zero value otherwise.

### GetPreviewImageOk

`func (o *AppPreviewAttributes) GetPreviewImageOk() (*ImageAsset, bool)`

GetPreviewImageOk returns a tuple with the PreviewImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewImage

`func (o *AppPreviewAttributes) SetPreviewImage(v ImageAsset)`

SetPreviewImage sets PreviewImage field to given value.

### HasPreviewImage

`func (o *AppPreviewAttributes) HasPreviewImage() bool`

HasPreviewImage returns a boolean if a field has been set.

### GetUploadOperations

`func (o *AppPreviewAttributes) GetUploadOperations() []UploadOperation`

GetUploadOperations returns the UploadOperations field if non-nil, zero value otherwise.

### GetUploadOperationsOk

`func (o *AppPreviewAttributes) GetUploadOperationsOk() (*[]UploadOperation, bool)`

GetUploadOperationsOk returns a tuple with the UploadOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadOperations

`func (o *AppPreviewAttributes) SetUploadOperations(v []UploadOperation)`

SetUploadOperations sets UploadOperations field to given value.

### HasUploadOperations

`func (o *AppPreviewAttributes) HasUploadOperations() bool`

HasUploadOperations returns a boolean if a field has been set.

### GetAssetDeliveryState

`func (o *AppPreviewAttributes) GetAssetDeliveryState() AppMediaAssetState`

GetAssetDeliveryState returns the AssetDeliveryState field if non-nil, zero value otherwise.

### GetAssetDeliveryStateOk

`func (o *AppPreviewAttributes) GetAssetDeliveryStateOk() (*AppMediaAssetState, bool)`

GetAssetDeliveryStateOk returns a tuple with the AssetDeliveryState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetDeliveryState

`func (o *AppPreviewAttributes) SetAssetDeliveryState(v AppMediaAssetState)`

SetAssetDeliveryState sets AssetDeliveryState field to given value.

### HasAssetDeliveryState

`func (o *AppPreviewAttributes) HasAssetDeliveryState() bool`

HasAssetDeliveryState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


