/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect

import (
	"encoding/json"
)

// UserInvitationCreateRequestData struct for UserInvitationCreateRequestData
type UserInvitationCreateRequestData struct {
	Type string `json:"type"`
	Attributes UserInvitationCreateRequestDataAttributes `json:"attributes"`
	Relationships *UserInvitationCreateRequestDataRelationships `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserInvitationCreateRequestData UserInvitationCreateRequestData

// NewUserInvitationCreateRequestData instantiates a new UserInvitationCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInvitationCreateRequestData(type_ string, attributes UserInvitationCreateRequestDataAttributes, ) *UserInvitationCreateRequestData {
	this := UserInvitationCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewUserInvitationCreateRequestDataWithDefaults instantiates a new UserInvitationCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInvitationCreateRequestDataWithDefaults() *UserInvitationCreateRequestData {
	this := UserInvitationCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *UserInvitationCreateRequestData) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UserInvitationCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UserInvitationCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *UserInvitationCreateRequestData) GetAttributes() UserInvitationCreateRequestDataAttributes {
	if o == nil  {
		var ret UserInvitationCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *UserInvitationCreateRequestData) GetAttributesOk() (*UserInvitationCreateRequestDataAttributes, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *UserInvitationCreateRequestData) SetAttributes(v UserInvitationCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *UserInvitationCreateRequestData) GetRelationships() UserInvitationCreateRequestDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret UserInvitationCreateRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvitationCreateRequestData) GetRelationshipsOk() (*UserInvitationCreateRequestDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *UserInvitationCreateRequestData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given UserInvitationCreateRequestDataRelationships and assigns it to the Relationships field.
func (o *UserInvitationCreateRequestData) SetRelationships(v UserInvitationCreateRequestDataRelationships) {
	o.Relationships = &v
}

func (o UserInvitationCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserInvitationCreateRequestData) UnmarshalJSON(bytes []byte) (err error) {
	varUserInvitationCreateRequestData := _UserInvitationCreateRequestData{}

	if err = json.Unmarshal(bytes, &varUserInvitationCreateRequestData); err == nil {
		*o = UserInvitationCreateRequestData(varUserInvitationCreateRequestData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserInvitationCreateRequestData struct {
	value *UserInvitationCreateRequestData
	isSet bool
}

func (v NullableUserInvitationCreateRequestData) Get() *UserInvitationCreateRequestData {
	return v.value
}

func (v *NullableUserInvitationCreateRequestData) Set(val *UserInvitationCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInvitationCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInvitationCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInvitationCreateRequestData(val *UserInvitationCreateRequestData) *NullableUserInvitationCreateRequestData {
	return &NullableUserInvitationCreateRequestData{value: val, isSet: true}
}

func (v NullableUserInvitationCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInvitationCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

