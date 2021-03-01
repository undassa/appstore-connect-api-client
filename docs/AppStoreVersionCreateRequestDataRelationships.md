# AppStoreVersionCreateRequestDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | [**AppPreOrderCreateRequestDataRelationshipsApp**](AppPreOrderCreateRequest_data_relationships_app.md) |  | 
**Build** | Pointer to [**AppStoreVersionCreateRequestDataRelationshipsBuild**](AppStoreVersionCreateRequest_data_relationships_build.md) |  | [optional] 

## Methods

### NewAppStoreVersionCreateRequestDataRelationships

`func NewAppStoreVersionCreateRequestDataRelationships(app AppPreOrderCreateRequestDataRelationshipsApp, ) *AppStoreVersionCreateRequestDataRelationships`

NewAppStoreVersionCreateRequestDataRelationships instantiates a new AppStoreVersionCreateRequestDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStoreVersionCreateRequestDataRelationshipsWithDefaults

`func NewAppStoreVersionCreateRequestDataRelationshipsWithDefaults() *AppStoreVersionCreateRequestDataRelationships`

NewAppStoreVersionCreateRequestDataRelationshipsWithDefaults instantiates a new AppStoreVersionCreateRequestDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *AppStoreVersionCreateRequestDataRelationships) GetApp() AppPreOrderCreateRequestDataRelationshipsApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *AppStoreVersionCreateRequestDataRelationships) GetAppOk() (*AppPreOrderCreateRequestDataRelationshipsApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *AppStoreVersionCreateRequestDataRelationships) SetApp(v AppPreOrderCreateRequestDataRelationshipsApp)`

SetApp sets App field to given value.


### GetBuild

`func (o *AppStoreVersionCreateRequestDataRelationships) GetBuild() AppStoreVersionCreateRequestDataRelationshipsBuild`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *AppStoreVersionCreateRequestDataRelationships) GetBuildOk() (*AppStoreVersionCreateRequestDataRelationshipsBuild, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *AppStoreVersionCreateRequestDataRelationships) SetBuild(v AppStoreVersionCreateRequestDataRelationshipsBuild)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *AppStoreVersionCreateRequestDataRelationships) HasBuild() bool`

HasBuild returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


