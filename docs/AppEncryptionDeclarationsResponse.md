# AppEncryptionDeclarationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AppEncryptionDeclaration**](AppEncryptionDeclaration.md) |  | 
**Included** | Pointer to [**[]App**](App.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewAppEncryptionDeclarationsResponse

`func NewAppEncryptionDeclarationsResponse(data []AppEncryptionDeclaration, links PagedDocumentLinks, ) *AppEncryptionDeclarationsResponse`

NewAppEncryptionDeclarationsResponse instantiates a new AppEncryptionDeclarationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppEncryptionDeclarationsResponseWithDefaults

`func NewAppEncryptionDeclarationsResponseWithDefaults() *AppEncryptionDeclarationsResponse`

NewAppEncryptionDeclarationsResponseWithDefaults instantiates a new AppEncryptionDeclarationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppEncryptionDeclarationsResponse) GetData() []AppEncryptionDeclaration`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppEncryptionDeclarationsResponse) GetDataOk() (*[]AppEncryptionDeclaration, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppEncryptionDeclarationsResponse) SetData(v []AppEncryptionDeclaration)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AppEncryptionDeclarationsResponse) GetIncluded() []App`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppEncryptionDeclarationsResponse) GetIncludedOk() (*[]App, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppEncryptionDeclarationsResponse) SetIncluded(v []App)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppEncryptionDeclarationsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *AppEncryptionDeclarationsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppEncryptionDeclarationsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppEncryptionDeclarationsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *AppEncryptionDeclarationsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AppEncryptionDeclarationsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AppEncryptionDeclarationsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AppEncryptionDeclarationsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


