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

// BuildsResponse struct for BuildsResponse
type BuildsResponse struct {
	Data                 []Build                   `json:"data"`
	Included             *[]OneOfPrereleaseVersion `json:"included,omitempty"`
	Links                PagedDocumentLinks        `json:"links"`
	Meta                 *PagingInformation        `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BuildsResponse BuildsResponse

// NewBuildsResponse instantiates a new BuildsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildsResponse(data []Build, links PagedDocumentLinks) *BuildsResponse {
	this := BuildsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBuildsResponseWithDefaults instantiates a new BuildsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildsResponseWithDefaults() *BuildsResponse {
	this := BuildsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BuildsResponse) GetData() []Build {
	if o == nil {
		var ret []Build
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BuildsResponse) GetDataOk() (*[]Build, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BuildsResponse) SetData(v []Build) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *BuildsResponse) GetIncluded() []OneOfPrereleaseVersion {
	if o == nil || o.Included == nil {
		var ret []OneOfPrereleaseVersion
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildsResponse) GetIncludedOk() (*[]OneOfPrereleaseVersion, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *BuildsResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []OneOfPrereleaseVersion and assigns it to the Included field.
func (o *BuildsResponse) SetIncluded(v []OneOfPrereleaseVersion) {
	o.Included = &v
}

// GetLinks returns the Links field value
func (o *BuildsResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BuildsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BuildsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *BuildsResponse) GetMeta() PagingInformation {
	if o == nil || o.Meta == nil {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *BuildsResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *BuildsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o BuildsResponse) MarshalJSON() ([]byte, error) {
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
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BuildsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBuildsResponse := _BuildsResponse{}

	if err = json.Unmarshal(bytes, &varBuildsResponse); err == nil {
		*o = BuildsResponse(varBuildsResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "included")
		delete(additionalProperties, "links")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBuildsResponse struct {
	value *BuildsResponse
	isSet bool
}

func (v NullableBuildsResponse) Get() *BuildsResponse {
	return v.value
}

func (v *NullableBuildsResponse) Set(val *BuildsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildsResponse(val *BuildsResponse) *NullableBuildsResponse {
	return &NullableBuildsResponse{value: val, isSet: true}
}

func (v NullableBuildsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
