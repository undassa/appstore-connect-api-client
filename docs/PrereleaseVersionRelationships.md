# PrereleaseVersionRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Builds** | Pointer to [**AppRelationshipsBuilds**](App_relationships_builds.md) |  | [optional] 
**App** | Pointer to [**AppEncryptionDeclarationRelationshipsApp**](AppEncryptionDeclaration_relationships_app.md) |  | [optional] 

## Methods

### NewPrereleaseVersionRelationships

`func NewPrereleaseVersionRelationships() *PrereleaseVersionRelationships`

NewPrereleaseVersionRelationships instantiates a new PrereleaseVersionRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrereleaseVersionRelationshipsWithDefaults

`func NewPrereleaseVersionRelationshipsWithDefaults() *PrereleaseVersionRelationships`

NewPrereleaseVersionRelationshipsWithDefaults instantiates a new PrereleaseVersionRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuilds

`func (o *PrereleaseVersionRelationships) GetBuilds() AppRelationshipsBuilds`

GetBuilds returns the Builds field if non-nil, zero value otherwise.

### GetBuildsOk

`func (o *PrereleaseVersionRelationships) GetBuildsOk() (*AppRelationshipsBuilds, bool)`

GetBuildsOk returns a tuple with the Builds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilds

`func (o *PrereleaseVersionRelationships) SetBuilds(v AppRelationshipsBuilds)`

SetBuilds sets Builds field to given value.

### HasBuilds

`func (o *PrereleaseVersionRelationships) HasBuilds() bool`

HasBuilds returns a boolean if a field has been set.

### GetApp

`func (o *PrereleaseVersionRelationships) GetApp() AppEncryptionDeclarationRelationshipsApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *PrereleaseVersionRelationships) GetAppOk() (*AppEncryptionDeclarationRelationshipsApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *PrereleaseVersionRelationships) SetApp(v AppEncryptionDeclarationRelationshipsApp)`

SetApp sets App field to given value.

### HasApp

`func (o *PrereleaseVersionRelationships) HasApp() bool`

HasApp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


