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

// UserInvitationsResponse struct for UserInvitationsResponse
type UserInvitationsResponse struct {
	Data []UserInvitation `json:"data"`
	Included *[]App `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserInvitationsResponse UserInvitationsResponse

// NewUserInvitationsResponse instantiates a new UserInvitationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInvitationsResponse(data []UserInvitation, links PagedDocumentLinks, ) *UserInvitationsResponse {
	this := UserInvitationsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewUserInvitationsResponseWithDefaults instantiates a new UserInvitationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInvitationsResponseWithDefaults() *UserInvitationsResponse {
	this := UserInvitationsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UserInvitationsResponse) GetData() []UserInvitation {
	if o == nil  {
		var ret []UserInvitation
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UserInvitationsResponse) GetDataOk() (*[]UserInvitation, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UserInvitationsResponse) SetData(v []UserInvitation) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *UserInvitationsResponse) GetIncluded() []App {
	if o == nil || o.Included == nil {
		var ret []App
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvitationsResponse) GetIncludedOk() (*[]App, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *UserInvitationsResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []App and assigns it to the Included field.
func (o *UserInvitationsResponse) SetIncluded(v []App) {
	o.Included = &v
}

// GetLinks returns the Links field value
func (o *UserInvitationsResponse) GetLinks() PagedDocumentLinks {
	if o == nil  {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *UserInvitationsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *UserInvitationsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *UserInvitationsResponse) GetMeta() PagingInformation {
	if o == nil || o.Meta == nil {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvitationsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *UserInvitationsResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *UserInvitationsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o UserInvitationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	if true {
		toSerialize["links"] = o.Links
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserInvitationsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varUserInvitationsResponse := _UserInvitationsResponse{}

	if err = json.Unmarshal(bytes, &varUserInvitationsResponse); err == nil {
		*o = UserInvitationsResponse(varUserInvitationsResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "included")
		delete(additionalProperties, "links")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserInvitationsResponse struct {
	value *UserInvitationsResponse
	isSet bool
}

func (v NullableUserInvitationsResponse) Get() *UserInvitationsResponse {
	return v.value
}

func (v *NullableUserInvitationsResponse) Set(val *UserInvitationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInvitationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInvitationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInvitationsResponse(val *UserInvitationsResponse) *NullableUserInvitationsResponse {
	return &NullableUserInvitationsResponse{value: val, isSet: true}
}

func (v NullableUserInvitationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInvitationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


