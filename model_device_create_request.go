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

// DeviceCreateRequest struct for DeviceCreateRequest
type DeviceCreateRequest struct {
	Data DeviceCreateRequestData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _DeviceCreateRequest DeviceCreateRequest

// NewDeviceCreateRequest instantiates a new DeviceCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceCreateRequest(data DeviceCreateRequestData, ) *DeviceCreateRequest {
	this := DeviceCreateRequest{}
	this.Data = data
	return &this
}

// NewDeviceCreateRequestWithDefaults instantiates a new DeviceCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceCreateRequestWithDefaults() *DeviceCreateRequest {
	this := DeviceCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *DeviceCreateRequest) GetData() DeviceCreateRequestData {
	if o == nil  {
		var ret DeviceCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DeviceCreateRequest) GetDataOk() (*DeviceCreateRequestData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DeviceCreateRequest) SetData(v DeviceCreateRequestData) {
	o.Data = v
}

func (o DeviceCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeviceCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varDeviceCreateRequest := _DeviceCreateRequest{}

	if err = json.Unmarshal(bytes, &varDeviceCreateRequest); err == nil {
		*o = DeviceCreateRequest(varDeviceCreateRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceCreateRequest struct {
	value *DeviceCreateRequest
	isSet bool
}

func (v NullableDeviceCreateRequest) Get() *DeviceCreateRequest {
	return v.value
}

func (v *NullableDeviceCreateRequest) Set(val *DeviceCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceCreateRequest(val *DeviceCreateRequest) *NullableDeviceCreateRequest {
	return &NullableDeviceCreateRequest{value: val, isSet: true}
}

func (v NullableDeviceCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

