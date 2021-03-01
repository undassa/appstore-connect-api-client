# AppStoreReviewAttachmentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AppStoreReviewAttachment**](AppStoreReviewAttachment.md) |  | 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewAppStoreReviewAttachmentsResponse

`func NewAppStoreReviewAttachmentsResponse(data []AppStoreReviewAttachment, links PagedDocumentLinks, ) *AppStoreReviewAttachmentsResponse`

NewAppStoreReviewAttachmentsResponse instantiates a new AppStoreReviewAttachmentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStoreReviewAttachmentsResponseWithDefaults

`func NewAppStoreReviewAttachmentsResponseWithDefaults() *AppStoreReviewAttachmentsResponse`

NewAppStoreReviewAttachmentsResponseWithDefaults instantiates a new AppStoreReviewAttachmentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppStoreReviewAttachmentsResponse) GetData() []AppStoreReviewAttachment`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppStoreReviewAttachmentsResponse) GetDataOk() (*[]AppStoreReviewAttachment, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppStoreReviewAttachmentsResponse) SetData(v []AppStoreReviewAttachment)`

SetData sets Data field to given value.


### GetLinks

`func (o *AppStoreReviewAttachmentsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppStoreReviewAttachmentsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppStoreReviewAttachmentsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *AppStoreReviewAttachmentsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AppStoreReviewAttachmentsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AppStoreReviewAttachmentsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AppStoreReviewAttachmentsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


