# AppStoreVersionSubmission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Relationships** | Pointer to [**AppStoreVersionSubmissionRelationships**](AppStoreVersionSubmission_relationships.md) |  | [optional] 
**Links** | [**ResourceLinks**](ResourceLinks.md) |  | 

## Methods

### NewAppStoreVersionSubmission

`func NewAppStoreVersionSubmission(type_ string, id string, links ResourceLinks, ) *AppStoreVersionSubmission`

NewAppStoreVersionSubmission instantiates a new AppStoreVersionSubmission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStoreVersionSubmissionWithDefaults

`func NewAppStoreVersionSubmissionWithDefaults() *AppStoreVersionSubmission`

NewAppStoreVersionSubmissionWithDefaults instantiates a new AppStoreVersionSubmission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppStoreVersionSubmission) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppStoreVersionSubmission) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppStoreVersionSubmission) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *AppStoreVersionSubmission) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppStoreVersionSubmission) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppStoreVersionSubmission) SetId(v string)`

SetId sets Id field to given value.


### GetRelationships

`func (o *AppStoreVersionSubmission) GetRelationships() AppStoreVersionSubmissionRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AppStoreVersionSubmission) GetRelationshipsOk() (*AppStoreVersionSubmissionRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AppStoreVersionSubmission) SetRelationships(v AppStoreVersionSubmissionRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AppStoreVersionSubmission) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *AppStoreVersionSubmission) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppStoreVersionSubmission) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppStoreVersionSubmission) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


