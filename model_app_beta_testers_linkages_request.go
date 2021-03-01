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

// AppBetaTestersLinkagesRequest struct for AppBetaTestersLinkagesRequest
type AppBetaTestersLinkagesRequest struct {
	Data []BetaGroupRelationshipsBetaTestersData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _AppBetaTestersLinkagesRequest AppBetaTestersLinkagesRequest

// NewAppBetaTestersLinkagesRequest instantiates a new AppBetaTestersLinkagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppBetaTestersLinkagesRequest(data []BetaGroupRelationshipsBetaTestersData, ) *AppBetaTestersLinkagesRequest {
	this := AppBetaTestersLinkagesRequest{}
	this.Data = data
	return &this
}

// NewAppBetaTestersLinkagesRequestWithDefaults instantiates a new AppBetaTestersLinkagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppBetaTestersLinkagesRequestWithDefaults() *AppBetaTestersLinkagesRequest {
	this := AppBetaTestersLinkagesRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppBetaTestersLinkagesRequest) GetData() []BetaGroupRelationshipsBetaTestersData {
	if o == nil  {
		var ret []BetaGroupRelationshipsBetaTestersData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppBetaTestersLinkagesRequest) GetDataOk() (*[]BetaGroupRelationshipsBetaTestersData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppBetaTestersLinkagesRequest) SetData(v []BetaGroupRelationshipsBetaTestersData) {
	o.Data = v
}

func (o AppBetaTestersLinkagesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppBetaTestersLinkagesRequest) UnmarshalJSON(bytes []byte) (err error) {
	varAppBetaTestersLinkagesRequest := _AppBetaTestersLinkagesRequest{}

	if err = json.Unmarshal(bytes, &varAppBetaTestersLinkagesRequest); err == nil {
		*o = AppBetaTestersLinkagesRequest(varAppBetaTestersLinkagesRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppBetaTestersLinkagesRequest struct {
	value *AppBetaTestersLinkagesRequest
	isSet bool
}

func (v NullableAppBetaTestersLinkagesRequest) Get() *AppBetaTestersLinkagesRequest {
	return v.value
}

func (v *NullableAppBetaTestersLinkagesRequest) Set(val *AppBetaTestersLinkagesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppBetaTestersLinkagesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppBetaTestersLinkagesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppBetaTestersLinkagesRequest(val *AppBetaTestersLinkagesRequest) *NullableAppBetaTestersLinkagesRequest {
	return &NullableAppBetaTestersLinkagesRequest{value: val, isSet: true}
}

func (v NullableAppBetaTestersLinkagesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppBetaTestersLinkagesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

