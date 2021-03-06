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

// AppStoreVersionUpdateRequest struct for AppStoreVersionUpdateRequest
type AppStoreVersionUpdateRequest struct {
	Data AppStoreVersionUpdateRequestData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _AppStoreVersionUpdateRequest AppStoreVersionUpdateRequest

// NewAppStoreVersionUpdateRequest instantiates a new AppStoreVersionUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionUpdateRequest(data AppStoreVersionUpdateRequestData, ) *AppStoreVersionUpdateRequest {
	this := AppStoreVersionUpdateRequest{}
	this.Data = data
	return &this
}

// NewAppStoreVersionUpdateRequestWithDefaults instantiates a new AppStoreVersionUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionUpdateRequestWithDefaults() *AppStoreVersionUpdateRequest {
	this := AppStoreVersionUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreVersionUpdateRequest) GetData() AppStoreVersionUpdateRequestData {
	if o == nil  {
		var ret AppStoreVersionUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionUpdateRequest) GetDataOk() (*AppStoreVersionUpdateRequestData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppStoreVersionUpdateRequest) SetData(v AppStoreVersionUpdateRequestData) {
	o.Data = v
}

func (o AppStoreVersionUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppStoreVersionUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varAppStoreVersionUpdateRequest := _AppStoreVersionUpdateRequest{}

	if err = json.Unmarshal(bytes, &varAppStoreVersionUpdateRequest); err == nil {
		*o = AppStoreVersionUpdateRequest(varAppStoreVersionUpdateRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppStoreVersionUpdateRequest struct {
	value *AppStoreVersionUpdateRequest
	isSet bool
}

func (v NullableAppStoreVersionUpdateRequest) Get() *AppStoreVersionUpdateRequest {
	return v.value
}

func (v *NullableAppStoreVersionUpdateRequest) Set(val *AppStoreVersionUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionUpdateRequest(val *AppStoreVersionUpdateRequest) *NullableAppStoreVersionUpdateRequest {
	return &NullableAppStoreVersionUpdateRequest{value: val, isSet: true}
}

func (v NullableAppStoreVersionUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


