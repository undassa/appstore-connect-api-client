# BetaTesterRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to [**BetaTesterRelationshipsApps**](BetaTester_relationships_apps.md) |  | [optional] 
**BetaGroups** | Pointer to [**AppRelationshipsBetaGroups**](App_relationships_betaGroups.md) |  | [optional] 
**Builds** | Pointer to [**AppRelationshipsBuilds**](App_relationships_builds.md) |  | [optional] 

## Methods

### NewBetaTesterRelationships

`func NewBetaTesterRelationships() *BetaTesterRelationships`

NewBetaTesterRelationships instantiates a new BetaTesterRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaTesterRelationshipsWithDefaults

`func NewBetaTesterRelationshipsWithDefaults() *BetaTesterRelationships`

NewBetaTesterRelationshipsWithDefaults instantiates a new BetaTesterRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *BetaTesterRelationships) GetApps() BetaTesterRelationshipsApps`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *BetaTesterRelationships) GetAppsOk() (*BetaTesterRelationshipsApps, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *BetaTesterRelationships) SetApps(v BetaTesterRelationshipsApps)`

SetApps sets Apps field to given value.

### HasApps

`func (o *BetaTesterRelationships) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetBetaGroups

`func (o *BetaTesterRelationships) GetBetaGroups() AppRelationshipsBetaGroups`

GetBetaGroups returns the BetaGroups field if non-nil, zero value otherwise.

### GetBetaGroupsOk

`func (o *BetaTesterRelationships) GetBetaGroupsOk() (*AppRelationshipsBetaGroups, bool)`

GetBetaGroupsOk returns a tuple with the BetaGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetaGroups

`func (o *BetaTesterRelationships) SetBetaGroups(v AppRelationshipsBetaGroups)`

SetBetaGroups sets BetaGroups field to given value.

### HasBetaGroups

`func (o *BetaTesterRelationships) HasBetaGroups() bool`

HasBetaGroups returns a boolean if a field has been set.

### GetBuilds

`func (o *BetaTesterRelationships) GetBuilds() AppRelationshipsBuilds`

GetBuilds returns the Builds field if non-nil, zero value otherwise.

### GetBuildsOk

`func (o *BetaTesterRelationships) GetBuildsOk() (*AppRelationshipsBuilds, bool)`

GetBuildsOk returns a tuple with the Builds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilds

`func (o *BetaTesterRelationships) SetBuilds(v AppRelationshipsBuilds)`

SetBuilds sets Builds field to given value.

### HasBuilds

`func (o *BetaTesterRelationships) HasBuilds() bool`

HasBuilds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


