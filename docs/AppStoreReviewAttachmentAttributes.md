# AppStoreReviewAttachmentAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileSize** | Pointer to **int32** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**SourceFileChecksum** | Pointer to **string** |  | [optional] 
**UploadOperations** | Pointer to [**[]UploadOperation**](UploadOperation.md) |  | [optional] 
**AssetDeliveryState** | Pointer to [**AppMediaAssetState**](AppMediaAssetState.md) |  | [optional] 

## Methods

### NewAppStoreReviewAttachmentAttributes

`func NewAppStoreReviewAttachmentAttributes() *AppStoreReviewAttachmentAttributes`

NewAppStoreReviewAttachmentAttributes instantiates a new AppStoreReviewAttachmentAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStoreReviewAttachmentAttributesWithDefaults

`func NewAppStoreReviewAttachmentAttributesWithDefaults() *AppStoreReviewAttachmentAttributes`

NewAppStoreReviewAttachmentAttributesWithDefaults instantiates a new AppStoreReviewAttachmentAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileSize

`func (o *AppStoreReviewAttachmentAttributes) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *AppStoreReviewAttachmentAttributes) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *AppStoreReviewAttachmentAttributes) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *AppStoreReviewAttachmentAttributes) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetFileName

`func (o *AppStoreReviewAttachmentAttributes) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *AppStoreReviewAttachmentAttributes) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *AppStoreReviewAttachmentAttributes) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *AppStoreReviewAttachmentAttributes) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetSourceFileChecksum

`func (o *AppStoreReviewAttachmentAttributes) GetSourceFileChecksum() string`

GetSourceFileChecksum returns the SourceFileChecksum field if non-nil, zero value otherwise.

### GetSourceFileChecksumOk

`func (o *AppStoreReviewAttachmentAttributes) GetSourceFileChecksumOk() (*string, bool)`

GetSourceFileChecksumOk returns a tuple with the SourceFileChecksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFileChecksum

`func (o *AppStoreReviewAttachmentAttributes) SetSourceFileChecksum(v string)`

SetSourceFileChecksum sets SourceFileChecksum field to given value.

### HasSourceFileChecksum

`func (o *AppStoreReviewAttachmentAttributes) HasSourceFileChecksum() bool`

HasSourceFileChecksum returns a boolean if a field has been set.

### GetUploadOperations

`func (o *AppStoreReviewAttachmentAttributes) GetUploadOperations() []UploadOperation`

GetUploadOperations returns the UploadOperations field if non-nil, zero value otherwise.

### GetUploadOperationsOk

`func (o *AppStoreReviewAttachmentAttributes) GetUploadOperationsOk() (*[]UploadOperation, bool)`

GetUploadOperationsOk returns a tuple with the UploadOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadOperations

`func (o *AppStoreReviewAttachmentAttributes) SetUploadOperations(v []UploadOperation)`

SetUploadOperations sets UploadOperations field to given value.

### HasUploadOperations

`func (o *AppStoreReviewAttachmentAttributes) HasUploadOperations() bool`

HasUploadOperations returns a boolean if a field has been set.

### GetAssetDeliveryState

`func (o *AppStoreReviewAttachmentAttributes) GetAssetDeliveryState() AppMediaAssetState`

GetAssetDeliveryState returns the AssetDeliveryState field if non-nil, zero value otherwise.

### GetAssetDeliveryStateOk

`func (o *AppStoreReviewAttachmentAttributes) GetAssetDeliveryStateOk() (*AppMediaAssetState, bool)`

GetAssetDeliveryStateOk returns a tuple with the AssetDeliveryState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetDeliveryState

`func (o *AppStoreReviewAttachmentAttributes) SetAssetDeliveryState(v AppMediaAssetState)`

SetAssetDeliveryState sets AssetDeliveryState field to given value.

### HasAssetDeliveryState

`func (o *AppStoreReviewAttachmentAttributes) HasAssetDeliveryState() bool`

HasAssetDeliveryState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


