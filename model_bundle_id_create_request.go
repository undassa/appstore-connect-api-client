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

// BundleIdCreateRequest struct for BundleIdCreateRequest
type BundleIdCreateRequest struct {
	Data BundleIdCreateRequestData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _BundleIdCreateRequest BundleIdCreateRequest

// NewBundleIdCreateRequest instantiates a new BundleIdCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleIdCreateRequest(data BundleIdCreateRequestData, ) *BundleIdCreateRequest {
	this := BundleIdCreateRequest{}
	this.Data = data
	return &this
}

// NewBundleIdCreateRequestWithDefaults instantiates a new BundleIdCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleIdCreateRequestWithDefaults() *BundleIdCreateRequest {
	this := BundleIdCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *BundleIdCreateRequest) GetData() BundleIdCreateRequestData {
	if o == nil  {
		var ret BundleIdCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BundleIdCreateRequest) GetDataOk() (*BundleIdCreateRequestData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BundleIdCreateRequest) SetData(v BundleIdCreateRequestData) {
	o.Data = v
}

func (o BundleIdCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BundleIdCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varBundleIdCreateRequest := _BundleIdCreateRequest{}

	if err = json.Unmarshal(bytes, &varBundleIdCreateRequest); err == nil {
		*o = BundleIdCreateRequest(varBundleIdCreateRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBundleIdCreateRequest struct {
	value *BundleIdCreateRequest
	isSet bool
}

func (v NullableBundleIdCreateRequest) Get() *BundleIdCreateRequest {
	return v.value
}

func (v *NullableBundleIdCreateRequest) Set(val *BundleIdCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleIdCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleIdCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleIdCreateRequest(val *BundleIdCreateRequest) *NullableBundleIdCreateRequest {
	return &NullableBundleIdCreateRequest{value: val, isSet: true}
}

func (v NullableBundleIdCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleIdCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

