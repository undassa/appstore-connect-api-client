# UserAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to [**[]UserRole**](UserRole.md) |  | [optional] 
**AllAppsVisible** | Pointer to **bool** |  | [optional] 
**ProvisioningAllowed** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserAttributes

`func NewUserAttributes() *UserAttributes`

NewUserAttributes instantiates a new UserAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAttributesWithDefaults

`func NewUserAttributesWithDefaults() *UserAttributes`

NewUserAttributesWithDefaults instantiates a new UserAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserAttributes) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserAttributes) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserAttributes) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserAttributes) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetFirstName

`func (o *UserAttributes) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserAttributes) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserAttributes) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserAttributes) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UserAttributes) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserAttributes) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserAttributes) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserAttributes) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetRoles

`func (o *UserAttributes) GetRoles() []UserRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserAttributes) GetRolesOk() (*[]UserRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserAttributes) SetRoles(v []UserRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserAttributes) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetAllAppsVisible

`func (o *UserAttributes) GetAllAppsVisible() bool`

GetAllAppsVisible returns the AllAppsVisible field if non-nil, zero value otherwise.

### GetAllAppsVisibleOk

`func (o *UserAttributes) GetAllAppsVisibleOk() (*bool, bool)`

GetAllAppsVisibleOk returns a tuple with the AllAppsVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllAppsVisible

`func (o *UserAttributes) SetAllAppsVisible(v bool)`

SetAllAppsVisible sets AllAppsVisible field to given value.

### HasAllAppsVisible

`func (o *UserAttributes) HasAllAppsVisible() bool`

HasAllAppsVisible returns a boolean if a field has been set.

### GetProvisioningAllowed

`func (o *UserAttributes) GetProvisioningAllowed() bool`

GetProvisioningAllowed returns the ProvisioningAllowed field if non-nil, zero value otherwise.

### GetProvisioningAllowedOk

`func (o *UserAttributes) GetProvisioningAllowedOk() (*bool, bool)`

GetProvisioningAllowedOk returns a tuple with the ProvisioningAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningAllowed

`func (o *UserAttributes) SetProvisioningAllowed(v bool)`

SetProvisioningAllowed sets ProvisioningAllowed field to given value.

### HasProvisioningAllowed

`func (o *UserAttributes) HasProvisioningAllowed() bool`

HasProvisioningAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


