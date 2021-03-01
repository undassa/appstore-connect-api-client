# UserInvitationCreateRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**Roles** | [**[]UserRole**](UserRole.md) |  | 
**AllAppsVisible** | Pointer to **bool** |  | [optional] 
**ProvisioningAllowed** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserInvitationCreateRequestDataAttributes

`func NewUserInvitationCreateRequestDataAttributes(email string, firstName string, lastName string, roles []UserRole, ) *UserInvitationCreateRequestDataAttributes`

NewUserInvitationCreateRequestDataAttributes instantiates a new UserInvitationCreateRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInvitationCreateRequestDataAttributesWithDefaults

`func NewUserInvitationCreateRequestDataAttributesWithDefaults() *UserInvitationCreateRequestDataAttributes`

NewUserInvitationCreateRequestDataAttributesWithDefaults instantiates a new UserInvitationCreateRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserInvitationCreateRequestDataAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserInvitationCreateRequestDataAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserInvitationCreateRequestDataAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *UserInvitationCreateRequestDataAttributes) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserInvitationCreateRequestDataAttributes) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserInvitationCreateRequestDataAttributes) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *UserInvitationCreateRequestDataAttributes) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserInvitationCreateRequestDataAttributes) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserInvitationCreateRequestDataAttributes) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetRoles

`func (o *UserInvitationCreateRequestDataAttributes) GetRoles() []UserRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserInvitationCreateRequestDataAttributes) GetRolesOk() (*[]UserRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserInvitationCreateRequestDataAttributes) SetRoles(v []UserRole)`

SetRoles sets Roles field to given value.


### GetAllAppsVisible

`func (o *UserInvitationCreateRequestDataAttributes) GetAllAppsVisible() bool`

GetAllAppsVisible returns the AllAppsVisible field if non-nil, zero value otherwise.

### GetAllAppsVisibleOk

`func (o *UserInvitationCreateRequestDataAttributes) GetAllAppsVisibleOk() (*bool, bool)`

GetAllAppsVisibleOk returns a tuple with the AllAppsVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllAppsVisible

`func (o *UserInvitationCreateRequestDataAttributes) SetAllAppsVisible(v bool)`

SetAllAppsVisible sets AllAppsVisible field to given value.

### HasAllAppsVisible

`func (o *UserInvitationCreateRequestDataAttributes) HasAllAppsVisible() bool`

HasAllAppsVisible returns a boolean if a field has been set.

### GetProvisioningAllowed

`func (o *UserInvitationCreateRequestDataAttributes) GetProvisioningAllowed() bool`

GetProvisioningAllowed returns the ProvisioningAllowed field if non-nil, zero value otherwise.

### GetProvisioningAllowedOk

`func (o *UserInvitationCreateRequestDataAttributes) GetProvisioningAllowedOk() (*bool, bool)`

GetProvisioningAllowedOk returns a tuple with the ProvisioningAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningAllowed

`func (o *UserInvitationCreateRequestDataAttributes) SetProvisioningAllowed(v bool)`

SetProvisioningAllowed sets ProvisioningAllowed field to given value.

### HasProvisioningAllowed

`func (o *UserInvitationCreateRequestDataAttributes) HasProvisioningAllowed() bool`

HasProvisioningAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


