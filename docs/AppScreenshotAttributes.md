# AppScreenshotAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileSize** | Pointer to **int32** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**SourceFileChecksum** | Pointer to **string** |  | [optional] 
**ImageAsset** | Pointer to [**ImageAsset**](ImageAsset.md) |  | [optional] 
**AssetToken** | Pointer to **string** |  | [optional] 
**AssetType** | Pointer to **string** |  | [optional] 
**UploadOperations** | Pointer to [**[]UploadOperation**](UploadOperation.md) |  | [optional] 
**AssetDeliveryState** | Pointer to [**AppMediaAssetState**](AppMediaAssetState.md) |  | [optional] 

## Methods

### NewAppScreenshotAttributes

`func NewAppScreenshotAttributes() *AppScreenshotAttributes`

NewAppScreenshotAttributes instantiates a new AppScreenshotAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppScreenshotAttributesWithDefaults

`func NewAppScreenshotAttributesWithDefaults() *AppScreenshotAttributes`

NewAppScreenshotAttributesWithDefaults instantiates a new AppScreenshotAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileSize

`func (o *AppScreenshotAttributes) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *AppScreenshotAttributes) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *AppScreenshotAttributes) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *AppScreenshotAttributes) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetFileName

`func (o *AppScreenshotAttributes) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *AppScreenshotAttributes) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *AppScreenshotAttributes) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *AppScreenshotAttributes) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetSourceFileChecksum

`func (o *AppScreenshotAttributes) GetSourceFileChecksum() string`

GetSourceFileChecksum returns the SourceFileChecksum field if non-nil, zero value otherwise.

### GetSourceFileChecksumOk

`func (o *AppScreenshotAttributes) GetSourceFileChecksumOk() (*string, bool)`

GetSourceFileChecksumOk returns a tuple with the SourceFileChecksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFileChecksum

`func (o *AppScreenshotAttributes) SetSourceFileChecksum(v string)`

SetSourceFileChecksum sets SourceFileChecksum field to given value.

### HasSourceFileChecksum

`func (o *AppScreenshotAttributes) HasSourceFileChecksum() bool`

HasSourceFileChecksum returns a boolean if a field has been set.

### GetImageAsset

`func (o *AppScreenshotAttributes) GetImageAsset() ImageAsset`

GetImageAsset returns the ImageAsset field if non-nil, zero value otherwise.

### GetImageAssetOk

`func (o *AppScreenshotAttributes) GetImageAssetOk() (*ImageAsset, bool)`

GetImageAssetOk returns a tuple with the ImageAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAsset

`func (o *AppScreenshotAttributes) SetImageAsset(v ImageAsset)`

SetImageAsset sets ImageAsset field to given value.

### HasImageAsset

`func (o *AppScreenshotAttributes) HasImageAsset() bool`

HasImageAsset returns a boolean if a field has been set.

### GetAssetToken

`func (o *AppScreenshotAttributes) GetAssetToken() string`

GetAssetToken returns the AssetToken field if non-nil, zero value otherwise.

### GetAssetTokenOk

`func (o *AppScreenshotAttributes) GetAssetTokenOk() (*string, bool)`

GetAssetTokenOk returns a tuple with the AssetToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetToken

`func (o *AppScreenshotAttributes) SetAssetToken(v string)`

SetAssetToken sets AssetToken field to given value.

### HasAssetToken

`func (o *AppScreenshotAttributes) HasAssetToken() bool`

HasAssetToken returns a boolean if a field has been set.

### GetAssetType

`func (o *AppScreenshotAttributes) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *AppScreenshotAttributes) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *AppScreenshotAttributes) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.

### HasAssetType

`func (o *AppScreenshotAttributes) HasAssetType() bool`

HasAssetType returns a boolean if a field has been set.

### GetUploadOperations

`func (o *AppScreenshotAttributes) GetUploadOperations() []UploadOperation`

GetUploadOperations returns the UploadOperations field if non-nil, zero value otherwise.

### GetUploadOperationsOk

`func (o *AppScreenshotAttributes) GetUploadOperationsOk() (*[]UploadOperation, bool)`

GetUploadOperationsOk returns a tuple with the UploadOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadOperations

`func (o *AppScreenshotAttributes) SetUploadOperations(v []UploadOperation)`

SetUploadOperations sets UploadOperations field to given value.

### HasUploadOperations

`func (o *AppScreenshotAttributes) HasUploadOperations() bool`

HasUploadOperations returns a boolean if a field has been set.

### GetAssetDeliveryState

`func (o *AppScreenshotAttributes) GetAssetDeliveryState() AppMediaAssetState`

GetAssetDeliveryState returns the AssetDeliveryState field if non-nil, zero value otherwise.

### GetAssetDeliveryStateOk

`func (o *AppScreenshotAttributes) GetAssetDeliveryStateOk() (*AppMediaAssetState, bool)`

GetAssetDeliveryStateOk returns a tuple with the AssetDeliveryState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetDeliveryState

`func (o *AppScreenshotAttributes) SetAssetDeliveryState(v AppMediaAssetState)`

SetAssetDeliveryState sets AssetDeliveryState field to given value.

### HasAssetDeliveryState

`func (o *AppScreenshotAttributes) HasAssetDeliveryState() bool`

HasAssetDeliveryState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


