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

// AppRelationshipsBetaAppReviewDetailData struct for AppRelationshipsBetaAppReviewDetailData
type AppRelationshipsBetaAppReviewDetailData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _AppRelationshipsBetaAppReviewDetailData AppRelationshipsBetaAppReviewDetailData

// NewAppRelationshipsBetaAppReviewDetailData instantiates a new AppRelationshipsBetaAppReviewDetailData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppRelationshipsBetaAppReviewDetailData(type_ string, id string, ) *AppRelationshipsBetaAppReviewDetailData {
	this := AppRelationshipsBetaAppReviewDetailData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppRelationshipsBetaAppReviewDetailDataWithDefaults instantiates a new AppRelationshipsBetaAppReviewDetailData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRelationshipsBetaAppReviewDetailDataWithDefaults() *AppRelationshipsBetaAppReviewDetailData {
	this := AppRelationshipsBetaAppReviewDetailData{}
	return &this
}

// GetType returns the Type field value
func (o *AppRelationshipsBetaAppReviewDetailData) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppRelationshipsBetaAppReviewDetailData) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppRelationshipsBetaAppReviewDetailData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppRelationshipsBetaAppReviewDetailData) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppRelationshipsBetaAppReviewDetailData) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppRelationshipsBetaAppReviewDetailData) SetId(v string) {
	o.Id = v
}

func (o AppRelationshipsBetaAppReviewDetailData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppRelationshipsBetaAppReviewDetailData) UnmarshalJSON(bytes []byte) (err error) {
	varAppRelationshipsBetaAppReviewDetailData := _AppRelationshipsBetaAppReviewDetailData{}

	if err = json.Unmarshal(bytes, &varAppRelationshipsBetaAppReviewDetailData); err == nil {
		*o = AppRelationshipsBetaAppReviewDetailData(varAppRelationshipsBetaAppReviewDetailData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppRelationshipsBetaAppReviewDetailData struct {
	value *AppRelationshipsBetaAppReviewDetailData
	isSet bool
}

func (v NullableAppRelationshipsBetaAppReviewDetailData) Get() *AppRelationshipsBetaAppReviewDetailData {
	return v.value
}

func (v *NullableAppRelationshipsBetaAppReviewDetailData) Set(val *AppRelationshipsBetaAppReviewDetailData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppRelationshipsBetaAppReviewDetailData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppRelationshipsBetaAppReviewDetailData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppRelationshipsBetaAppReviewDetailData(val *AppRelationshipsBetaAppReviewDetailData) *NullableAppRelationshipsBetaAppReviewDetailData {
	return &NullableAppRelationshipsBetaAppReviewDetailData{value: val, isSet: true}
}

func (v NullableAppRelationshipsBetaAppReviewDetailData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppRelationshipsBetaAppReviewDetailData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


