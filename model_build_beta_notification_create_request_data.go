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

// BuildBetaNotificationCreateRequestData struct for BuildBetaNotificationCreateRequestData
type BuildBetaNotificationCreateRequestData struct {
	Type string `json:"type"`
	Relationships BetaAppReviewSubmissionCreateRequestDataRelationships `json:"relationships"`
	AdditionalProperties map[string]interface{}
}

type _BuildBetaNotificationCreateRequestData BuildBetaNotificationCreateRequestData

// NewBuildBetaNotificationCreateRequestData instantiates a new BuildBetaNotificationCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildBetaNotificationCreateRequestData(type_ string, relationships BetaAppReviewSubmissionCreateRequestDataRelationships, ) *BuildBetaNotificationCreateRequestData {
	this := BuildBetaNotificationCreateRequestData{}
	this.Type = type_
	this.Relationships = relationships
	return &this
}

// NewBuildBetaNotificationCreateRequestDataWithDefaults instantiates a new BuildBetaNotificationCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildBetaNotificationCreateRequestDataWithDefaults() *BuildBetaNotificationCreateRequestData {
	this := BuildBetaNotificationCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *BuildBetaNotificationCreateRequestData) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BuildBetaNotificationCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BuildBetaNotificationCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetRelationships returns the Relationships field value
func (o *BuildBetaNotificationCreateRequestData) GetRelationships() BetaAppReviewSubmissionCreateRequestDataRelationships {
	if o == nil  {
		var ret BetaAppReviewSubmissionCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *BuildBetaNotificationCreateRequestData) GetRelationshipsOk() (*BetaAppReviewSubmissionCreateRequestDataRelationships, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *BuildBetaNotificationCreateRequestData) SetRelationships(v BetaAppReviewSubmissionCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o BuildBetaNotificationCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["relationships"] = o.Relationships
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BuildBetaNotificationCreateRequestData) UnmarshalJSON(bytes []byte) (err error) {
	varBuildBetaNotificationCreateRequestData := _BuildBetaNotificationCreateRequestData{}

	if err = json.Unmarshal(bytes, &varBuildBetaNotificationCreateRequestData); err == nil {
		*o = BuildBetaNotificationCreateRequestData(varBuildBetaNotificationCreateRequestData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBuildBetaNotificationCreateRequestData struct {
	value *BuildBetaNotificationCreateRequestData
	isSet bool
}

func (v NullableBuildBetaNotificationCreateRequestData) Get() *BuildBetaNotificationCreateRequestData {
	return v.value
}

func (v *NullableBuildBetaNotificationCreateRequestData) Set(val *BuildBetaNotificationCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildBetaNotificationCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildBetaNotificationCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildBetaNotificationCreateRequestData(val *BuildBetaNotificationCreateRequestData) *NullableBuildBetaNotificationCreateRequestData {
	return &NullableBuildBetaNotificationCreateRequestData{value: val, isSet: true}
}

func (v NullableBuildBetaNotificationCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildBetaNotificationCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


