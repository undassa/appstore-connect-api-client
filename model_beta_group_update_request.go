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

// BetaGroupUpdateRequest struct for BetaGroupUpdateRequest
type BetaGroupUpdateRequest struct {
	Data BetaGroupUpdateRequestData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _BetaGroupUpdateRequest BetaGroupUpdateRequest

// NewBetaGroupUpdateRequest instantiates a new BetaGroupUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaGroupUpdateRequest(data BetaGroupUpdateRequestData, ) *BetaGroupUpdateRequest {
	this := BetaGroupUpdateRequest{}
	this.Data = data
	return &this
}

// NewBetaGroupUpdateRequestWithDefaults instantiates a new BetaGroupUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaGroupUpdateRequestWithDefaults() *BetaGroupUpdateRequest {
	this := BetaGroupUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *BetaGroupUpdateRequest) GetData() BetaGroupUpdateRequestData {
	if o == nil  {
		var ret BetaGroupUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaGroupUpdateRequest) GetDataOk() (*BetaGroupUpdateRequestData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BetaGroupUpdateRequest) SetData(v BetaGroupUpdateRequestData) {
	o.Data = v
}

func (o BetaGroupUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BetaGroupUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varBetaGroupUpdateRequest := _BetaGroupUpdateRequest{}

	if err = json.Unmarshal(bytes, &varBetaGroupUpdateRequest); err == nil {
		*o = BetaGroupUpdateRequest(varBetaGroupUpdateRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBetaGroupUpdateRequest struct {
	value *BetaGroupUpdateRequest
	isSet bool
}

func (v NullableBetaGroupUpdateRequest) Get() *BetaGroupUpdateRequest {
	return v.value
}

func (v *NullableBetaGroupUpdateRequest) Set(val *BetaGroupUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaGroupUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaGroupUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaGroupUpdateRequest(val *BetaGroupUpdateRequest) *NullableBetaGroupUpdateRequest {
	return &NullableBetaGroupUpdateRequest{value: val, isSet: true}
}

func (v NullableBetaGroupUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaGroupUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

