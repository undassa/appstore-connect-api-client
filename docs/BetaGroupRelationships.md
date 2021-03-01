# BetaGroupRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to [**AppEncryptionDeclarationRelationshipsApp**](AppEncryptionDeclaration_relationships_app.md) |  | [optional] 
**Builds** | Pointer to [**AppRelationshipsBuilds**](App_relationships_builds.md) |  | [optional] 
**BetaTesters** | Pointer to [**BetaGroupRelationshipsBetaTesters**](BetaGroup_relationships_betaTesters.md) |  | [optional] 

## Methods

### NewBetaGroupRelationships

`func NewBetaGroupRelationships() *BetaGroupRelationships`

NewBetaGroupRelationships instantiates a new BetaGroupRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaGroupRelationshipsWithDefaults

`func NewBetaGroupRelationshipsWithDefaults() *BetaGroupRelationships`

NewBetaGroupRelationshipsWithDefaults instantiates a new BetaGroupRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *BetaGroupRelationships) GetApp() AppEncryptionDeclarationRelationshipsApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *BetaGroupRelationships) GetAppOk() (*AppEncryptionDeclarationRelationshipsApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *BetaGroupRelationships) SetApp(v AppEncryptionDeclarationRelationshipsApp)`

SetApp sets App field to given value.

### HasApp

`func (o *BetaGroupRelationships) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetBuilds

`func (o *BetaGroupRelationships) GetBuilds() AppRelationshipsBuilds`

GetBuilds returns the Builds field if non-nil, zero value otherwise.

### GetBuildsOk

`func (o *BetaGroupRelationships) GetBuildsOk() (*AppRelationshipsBuilds, bool)`

GetBuildsOk returns a tuple with the Builds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilds

`func (o *BetaGroupRelationships) SetBuilds(v AppRelationshipsBuilds)`

SetBuilds sets Builds field to given value.

### HasBuilds

`func (o *BetaGroupRelationships) HasBuilds() bool`

HasBuilds returns a boolean if a field has been set.

### GetBetaTesters

`func (o *BetaGroupRelationships) GetBetaTesters() BetaGroupRelationshipsBetaTesters`

GetBetaTesters returns the BetaTesters field if non-nil, zero value otherwise.

### GetBetaTestersOk

`func (o *BetaGroupRelationships) GetBetaTestersOk() (*BetaGroupRelationshipsBetaTesters, bool)`

GetBetaTestersOk returns a tuple with the BetaTesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetaTesters

`func (o *BetaGroupRelationships) SetBetaTesters(v BetaGroupRelationshipsBetaTesters)`

SetBetaTesters sets BetaTesters field to given value.

### HasBetaTesters

`func (o *BetaGroupRelationships) HasBetaTesters() bool`

HasBetaTesters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


