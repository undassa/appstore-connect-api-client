# BetaGroupCreateRequestDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | [**AppPreOrderCreateRequestDataRelationshipsApp**](AppPreOrderCreateRequest_data_relationships_app.md) |  | 
**Builds** | Pointer to [**BetaGroupCreateRequestDataRelationshipsBuilds**](BetaGroupCreateRequest_data_relationships_builds.md) |  | [optional] 
**BetaTesters** | Pointer to [**BetaGroupCreateRequestDataRelationshipsBetaTesters**](BetaGroupCreateRequest_data_relationships_betaTesters.md) |  | [optional] 

## Methods

### NewBetaGroupCreateRequestDataRelationships

`func NewBetaGroupCreateRequestDataRelationships(app AppPreOrderCreateRequestDataRelationshipsApp, ) *BetaGroupCreateRequestDataRelationships`

NewBetaGroupCreateRequestDataRelationships instantiates a new BetaGroupCreateRequestDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaGroupCreateRequestDataRelationshipsWithDefaults

`func NewBetaGroupCreateRequestDataRelationshipsWithDefaults() *BetaGroupCreateRequestDataRelationships`

NewBetaGroupCreateRequestDataRelationshipsWithDefaults instantiates a new BetaGroupCreateRequestDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *BetaGroupCreateRequestDataRelationships) GetApp() AppPreOrderCreateRequestDataRelationshipsApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *BetaGroupCreateRequestDataRelationships) GetAppOk() (*AppPreOrderCreateRequestDataRelationshipsApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *BetaGroupCreateRequestDataRelationships) SetApp(v AppPreOrderCreateRequestDataRelationshipsApp)`

SetApp sets App field to given value.


### GetBuilds

`func (o *BetaGroupCreateRequestDataRelationships) GetBuilds() BetaGroupCreateRequestDataRelationshipsBuilds`

GetBuilds returns the Builds field if non-nil, zero value otherwise.

### GetBuildsOk

`func (o *BetaGroupCreateRequestDataRelationships) GetBuildsOk() (*BetaGroupCreateRequestDataRelationshipsBuilds, bool)`

GetBuildsOk returns a tuple with the Builds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilds

`func (o *BetaGroupCreateRequestDataRelationships) SetBuilds(v BetaGroupCreateRequestDataRelationshipsBuilds)`

SetBuilds sets Builds field to given value.

### HasBuilds

`func (o *BetaGroupCreateRequestDataRelationships) HasBuilds() bool`

HasBuilds returns a boolean if a field has been set.

### GetBetaTesters

`func (o *BetaGroupCreateRequestDataRelationships) GetBetaTesters() BetaGroupCreateRequestDataRelationshipsBetaTesters`

GetBetaTesters returns the BetaTesters field if non-nil, zero value otherwise.

### GetBetaTestersOk

`func (o *BetaGroupCreateRequestDataRelationships) GetBetaTestersOk() (*BetaGroupCreateRequestDataRelationshipsBetaTesters, bool)`

GetBetaTestersOk returns a tuple with the BetaTesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetaTesters

`func (o *BetaGroupCreateRequestDataRelationships) SetBetaTesters(v BetaGroupCreateRequestDataRelationshipsBetaTesters)`

SetBetaTesters sets BetaTesters field to given value.

### HasBetaTesters

`func (o *BetaGroupCreateRequestDataRelationships) HasBetaTesters() bool`

HasBetaTesters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


