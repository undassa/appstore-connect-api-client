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

// PreReleaseVersionsResponse struct for PreReleaseVersionsResponse
type PreReleaseVersionsResponse struct {
	Data                 []PrereleaseVersion `json:"data"`
	Included             *[]OneOfBuildApp    `json:"included,omitempty"`
	Links                PagedDocumentLinks  `json:"links"`
	Meta                 *PagingInformation  `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PreReleaseVersionsResponse PreReleaseVersionsResponse
type OneOfBuildApp struct{}

// NewPreReleaseVersionsResponse instantiates a new PreReleaseVersionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreReleaseVersionsResponse(data []PrereleaseVersion, links PagedDocumentLinks) *PreReleaseVersionsResponse {
	this := PreReleaseVersionsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewPreReleaseVersionsResponseWithDefaults instantiates a new PreReleaseVersionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreReleaseVersionsResponseWithDefaults() *PreReleaseVersionsResponse {
	this := PreReleaseVersionsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *PreReleaseVersionsResponse) GetData() []PrereleaseVersion {
	if o == nil {
		var ret []PrereleaseVersion
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PreReleaseVersionsResponse) GetDataOk() (*[]PrereleaseVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PreReleaseVersionsResponse) SetData(v []PrereleaseVersion) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *PreReleaseVersionsResponse) GetIncluded() []OneOfBuildApp {
	if o == nil || o.Included == nil {
		var ret []OneOfBuildApp
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreReleaseVersionsResponse) GetIncludedOk() (*[]OneOfBuildApp, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *PreReleaseVersionsResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []OneOfBuildApp and assigns it to the Included field.
func (o *PreReleaseVersionsResponse) SetIncluded(v []OneOfBuildApp) {
	o.Included = &v
}

// GetLinks returns the Links field value
func (o *PreReleaseVersionsResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PreReleaseVersionsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *PreReleaseVersionsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *PreReleaseVersionsResponse) GetMeta() PagingInformation {
	if o == nil || o.Meta == nil {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreReleaseVersionsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *PreReleaseVersionsResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *PreReleaseVersionsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o PreReleaseVersionsResponse) MarshalJSON() ([]byte, error) {
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

func (o *PreReleaseVersionsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPreReleaseVersionsResponse := _PreReleaseVersionsResponse{}

	if err = json.Unmarshal(bytes, &varPreReleaseVersionsResponse); err == nil {
		*o = PreReleaseVersionsResponse(varPreReleaseVersionsResponse)
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

type NullablePreReleaseVersionsResponse struct {
	value *PreReleaseVersionsResponse
	isSet bool
}

func (v NullablePreReleaseVersionsResponse) Get() *PreReleaseVersionsResponse {
	return v.value
}

func (v *NullablePreReleaseVersionsResponse) Set(val *PreReleaseVersionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePreReleaseVersionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePreReleaseVersionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreReleaseVersionsResponse(val *PreReleaseVersionsResponse) *NullablePreReleaseVersionsResponse {
	return &NullablePreReleaseVersionsResponse{value: val, isSet: true}
}

func (v NullablePreReleaseVersionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreReleaseVersionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
