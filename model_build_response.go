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

// BuildResponse struct for BuildResponse
type BuildResponse struct {
	Data                 Build                     `json:"data"`
	Included             *[]OneOfPrereleaseVersion `json:"included,omitempty"`
	Links                DocumentLinks             `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _BuildResponse BuildResponse
type OneOfPrereleaseVersion struct{}

// NewBuildResponse instantiates a new BuildResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildResponse(data Build, links DocumentLinks) *BuildResponse {
	this := BuildResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBuildResponseWithDefaults instantiates a new BuildResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildResponseWithDefaults() *BuildResponse {
	this := BuildResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BuildResponse) GetData() Build {
	if o == nil {
		var ret Build
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BuildResponse) GetDataOk() (*Build, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BuildResponse) SetData(v Build) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *BuildResponse) GetIncluded() []OneOfPrereleaseVersion {
	if o == nil || o.Included == nil {
		var ret []OneOfPrereleaseVersion
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildResponse) GetIncludedOk() (*[]OneOfPrereleaseVersion, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *BuildResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []OneOfPrereleaseVersion and assigns it to the Included field.
func (o *BuildResponse) SetIncluded(v []OneOfPrereleaseVersion) {
	o.Included = &v
}

// GetLinks returns the Links field value
func (o *BuildResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BuildResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BuildResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o BuildResponse) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BuildResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBuildResponse := _BuildResponse{}

	if err = json.Unmarshal(bytes, &varBuildResponse); err == nil {
		*o = BuildResponse(varBuildResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "included")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBuildResponse struct {
	value *BuildResponse
	isSet bool
}

func (v NullableBuildResponse) Get() *BuildResponse {
	return v.value
}

func (v *NullableBuildResponse) Set(val *BuildResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildResponse(val *BuildResponse) *NullableBuildResponse {
	return &NullableBuildResponse{value: val, isSet: true}
}

func (v NullableBuildResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
