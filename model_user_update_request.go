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

// UserUpdateRequest struct for UserUpdateRequest
type UserUpdateRequest struct {
	Data UserUpdateRequestData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _UserUpdateRequest UserUpdateRequest

// NewUserUpdateRequest instantiates a new UserUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserUpdateRequest(data UserUpdateRequestData, ) *UserUpdateRequest {
	this := UserUpdateRequest{}
	this.Data = data
	return &this
}

// NewUserUpdateRequestWithDefaults instantiates a new UserUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserUpdateRequestWithDefaults() *UserUpdateRequest {
	this := UserUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *UserUpdateRequest) GetData() UserUpdateRequestData {
	if o == nil  {
		var ret UserUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UserUpdateRequest) GetDataOk() (*UserUpdateRequestData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UserUpdateRequest) SetData(v UserUpdateRequestData) {
	o.Data = v
}

func (o UserUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varUserUpdateRequest := _UserUpdateRequest{}

	if err = json.Unmarshal(bytes, &varUserUpdateRequest); err == nil {
		*o = UserUpdateRequest(varUserUpdateRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserUpdateRequest struct {
	value *UserUpdateRequest
	isSet bool
}

func (v NullableUserUpdateRequest) Get() *UserUpdateRequest {
	return v.value
}

func (v *NullableUserUpdateRequest) Set(val *UserUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserUpdateRequest(val *UserUpdateRequest) *NullableUserUpdateRequest {
	return &NullableUserUpdateRequest{value: val, isSet: true}
}

func (v NullableUserUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


