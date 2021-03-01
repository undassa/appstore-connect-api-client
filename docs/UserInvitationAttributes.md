# UserInvitationAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**Roles** | Pointer to [**[]UserRole**](UserRole.md) |  | [optional] 
**AllAppsVisible** | Pointer to **bool** |  | [optional] 
**ProvisioningAllowed** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserInvitationAttributes

`func NewUserInvitationAttributes() *UserInvitationAttributes`

NewUserInvitationAttributes instantiates a new UserInvitationAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInvitationAttributesWithDefaults

`func NewUserInvitationAttributesWithDefaults() *UserInvitationAttributes`

NewUserInvitationAttributesWithDefaults instantiates a new UserInvitationAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserInvitationAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserInvitationAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserInvitationAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserInvitationAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *UserInvitationAttributes) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserInvitationAttributes) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserInvitationAttributes) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserInvitationAttributes) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UserInvitationAttributes) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserInvitationAttributes) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserInvitationAttributes) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserInvitationAttributes) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetExpirationDate

`func (o *UserInvitationAttributes) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *UserInvitationAttributes) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *UserInvitationAttributes) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *UserInvitationAttributes) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetRoles

`func (o *UserInvitationAttributes) GetRoles() []UserRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserInvitationAttributes) GetRolesOk() (*[]UserRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserInvitationAttributes) SetRoles(v []UserRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserInvitationAttributes) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetAllAppsVisible

`func (o *UserInvitationAttributes) GetAllAppsVisible() bool`

GetAllAppsVisible returns the AllAppsVisible field if non-nil, zero value otherwise.

### GetAllAppsVisibleOk

`func (o *UserInvitationAttributes) GetAllAppsVisibleOk() (*bool, bool)`

GetAllAppsVisibleOk returns a tuple with the AllAppsVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllAppsVisible

`func (o *UserInvitationAttributes) SetAllAppsVisible(v bool)`

SetAllAppsVisible sets AllAppsVisible field to given value.

### HasAllAppsVisible

`func (o *UserInvitationAttributes) HasAllAppsVisible() bool`

HasAllAppsVisible returns a boolean if a field has been set.

### GetProvisioningAllowed

`func (o *UserInvitationAttributes) GetProvisioningAllowed() bool`

GetProvisioningAllowed returns the ProvisioningAllowed field if non-nil, zero value otherwise.

### GetProvisioningAllowedOk

`func (o *UserInvitationAttributes) GetProvisioningAllowedOk() (*bool, bool)`

GetProvisioningAllowedOk returns a tuple with the ProvisioningAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningAllowed

`func (o *UserInvitationAttributes) SetProvisioningAllowed(v bool)`

SetProvisioningAllowed sets ProvisioningAllowed field to given value.

### HasProvisioningAllowed

`func (o *UserInvitationAttributes) HasProvisioningAllowed() bool`

HasProvisioningAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


