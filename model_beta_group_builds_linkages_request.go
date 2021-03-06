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

// BetaGroupBuildsLinkagesRequest struct for BetaGroupBuildsLinkagesRequest
type BetaGroupBuildsLinkagesRequest struct {
	Data []AppStoreVersionRelationshipsBuildData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _BetaGroupBuildsLinkagesRequest BetaGroupBuildsLinkagesRequest

// NewBetaGroupBuildsLinkagesRequest instantiates a new BetaGroupBuildsLinkagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaGroupBuildsLinkagesRequest(data []AppStoreVersionRelationshipsBuildData, ) *BetaGroupBuildsLinkagesRequest {
	this := BetaGroupBuildsLinkagesRequest{}
	this.Data = data
	return &this
}

// NewBetaGroupBuildsLinkagesRequestWithDefaults instantiates a new BetaGroupBuildsLinkagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaGroupBuildsLinkagesRequestWithDefaults() *BetaGroupBuildsLinkagesRequest {
	this := BetaGroupBuildsLinkagesRequest{}
	return &this
}

// GetData returns the Data field value
func (o *BetaGroupBuildsLinkagesRequest) GetData() []AppStoreVersionRelationshipsBuildData {
	if o == nil  {
		var ret []AppStoreVersionRelationshipsBuildData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaGroupBuildsLinkagesRequest) GetDataOk() (*[]AppStoreVersionRelationshipsBuildData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BetaGroupBuildsLinkagesRequest) SetData(v []AppStoreVersionRelationshipsBuildData) {
	o.Data = v
}

func (o BetaGroupBuildsLinkagesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BetaGroupBuildsLinkagesRequest) UnmarshalJSON(bytes []byte) (err error) {
	varBetaGroupBuildsLinkagesRequest := _BetaGroupBuildsLinkagesRequest{}

	if err = json.Unmarshal(bytes, &varBetaGroupBuildsLinkagesRequest); err == nil {
		*o = BetaGroupBuildsLinkagesRequest(varBetaGroupBuildsLinkagesRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBetaGroupBuildsLinkagesRequest struct {
	value *BetaGroupBuildsLinkagesRequest
	isSet bool
}

func (v NullableBetaGroupBuildsLinkagesRequest) Get() *BetaGroupBuildsLinkagesRequest {
	return v.value
}

func (v *NullableBetaGroupBuildsLinkagesRequest) Set(val *BetaGroupBuildsLinkagesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaGroupBuildsLinkagesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaGroupBuildsLinkagesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaGroupBuildsLinkagesRequest(val *BetaGroupBuildsLinkagesRequest) *NullableBetaGroupBuildsLinkagesRequest {
	return &NullableBetaGroupBuildsLinkagesRequest{value: val, isSet: true}
}

func (v NullableBetaGroupBuildsLinkagesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaGroupBuildsLinkagesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


