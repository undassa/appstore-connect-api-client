# AppStoreReviewAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**AppStoreReviewAttachmentAttributes**](AppStoreReviewAttachment_attributes.md) |  | [optional] 
**Relationships** | Pointer to [**AppStoreReviewAttachmentRelationships**](AppStoreReviewAttachment_relationships.md) |  | [optional] 
**Links** | [**ResourceLinks**](ResourceLinks.md) |  | 

## Methods

### NewAppStoreReviewAttachment

`func NewAppStoreReviewAttachment(type_ string, id string, links ResourceLinks, ) *AppStoreReviewAttachment`

NewAppStoreReviewAttachment instantiates a new AppStoreReviewAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStoreReviewAttachmentWithDefaults

`func NewAppStoreReviewAttachmentWithDefaults() *AppStoreReviewAttachment`

NewAppStoreReviewAttachmentWithDefaults instantiates a new AppStoreReviewAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppStoreReviewAttachment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppStoreReviewAttachment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppStoreReviewAttachment) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *AppStoreReviewAttachment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppStoreReviewAttachment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppStoreReviewAttachment) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *AppStoreReviewAttachment) GetAttributes() AppStoreReviewAttachmentAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppStoreReviewAttachment) GetAttributesOk() (*AppStoreReviewAttachmentAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppStoreReviewAttachment) SetAttributes(v AppStoreReviewAttachmentAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AppStoreReviewAttachment) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *AppStoreReviewAttachment) GetRelationships() AppStoreReviewAttachmentRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AppStoreReviewAttachment) GetRelationshipsOk() (*AppStoreReviewAttachmentRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AppStoreReviewAttachment) SetRelationships(v AppStoreReviewAttachmentRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AppStoreReviewAttachment) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *AppStoreReviewAttachment) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppStoreReviewAttachment) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppStoreReviewAttachment) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


