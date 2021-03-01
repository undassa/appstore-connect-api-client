# UserUpdateRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | Pointer to [**[]UserRole**](UserRole.md) |  | [optional] 
**AllAppsVisible** | Pointer to **bool** |  | [optional] 
**ProvisioningAllowed** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserUpdateRequestDataAttributes

`func NewUserUpdateRequestDataAttributes() *UserUpdateRequestDataAttributes`

NewUserUpdateRequestDataAttributes instantiates a new UserUpdateRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUpdateRequestDataAttributesWithDefaults

`func NewUserUpdateRequestDataAttributesWithDefaults() *UserUpdateRequestDataAttributes`

NewUserUpdateRequestDataAttributesWithDefaults instantiates a new UserUpdateRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *UserUpdateRequestDataAttributes) GetRoles() []UserRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserUpdateRequestDataAttributes) GetRolesOk() (*[]UserRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserUpdateRequestDataAttributes) SetRoles(v []UserRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserUpdateRequestDataAttributes) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetAllAppsVisible

`func (o *UserUpdateRequestDataAttributes) GetAllAppsVisible() bool`

GetAllAppsVisible returns the AllAppsVisible field if non-nil, zero value otherwise.

### GetAllAppsVisibleOk

`func (o *UserUpdateRequestDataAttributes) GetAllAppsVisibleOk() (*bool, bool)`

GetAllAppsVisibleOk returns a tuple with the AllAppsVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllAppsVisible

`func (o *UserUpdateRequestDataAttributes) SetAllAppsVisible(v bool)`

SetAllAppsVisible sets AllAppsVisible field to given value.

### HasAllAppsVisible

`func (o *UserUpdateRequestDataAttributes) HasAllAppsVisible() bool`

HasAllAppsVisible returns a boolean if a field has been set.

### GetProvisioningAllowed

`func (o *UserUpdateRequestDataAttributes) GetProvisioningAllowed() bool`

GetProvisioningAllowed returns the ProvisioningAllowed field if non-nil, zero value otherwise.

### GetProvisioningAllowedOk

`func (o *UserUpdateRequestDataAttributes) GetProvisioningAllowedOk() (*bool, bool)`

GetProvisioningAllowedOk returns a tuple with the ProvisioningAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningAllowed

`func (o *UserUpdateRequestDataAttributes) SetProvisioningAllowed(v bool)`

SetProvisioningAllowed sets ProvisioningAllowed field to given value.

### HasProvisioningAllowed

`func (o *UserUpdateRequestDataAttributes) HasProvisioningAllowed() bool`

HasProvisioningAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


