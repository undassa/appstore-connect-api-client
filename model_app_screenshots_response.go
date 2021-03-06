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

// AppScreenshotsResponse struct for AppScreenshotsResponse
type AppScreenshotsResponse struct {
	Data []AppScreenshot `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppScreenshotsResponse AppScreenshotsResponse

// NewAppScreenshotsResponse instantiates a new AppScreenshotsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppScreenshotsResponse(data []AppScreenshot, links PagedDocumentLinks, ) *AppScreenshotsResponse {
	this := AppScreenshotsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppScreenshotsResponseWithDefaults instantiates a new AppScreenshotsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppScreenshotsResponseWithDefaults() *AppScreenshotsResponse {
	this := AppScreenshotsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppScreenshotsResponse) GetData() []AppScreenshot {
	if o == nil  {
		var ret []AppScreenshot
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppScreenshotsResponse) GetDataOk() (*[]AppScreenshot, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppScreenshotsResponse) SetData(v []AppScreenshot) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *AppScreenshotsResponse) GetLinks() PagedDocumentLinks {
	if o == nil  {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppScreenshotsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppScreenshotsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppScreenshotsResponse) GetMeta() PagingInformation {
	if o == nil || o.Meta == nil {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppScreenshotsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppScreenshotsResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppScreenshotsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o AppScreenshotsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
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

func (o *AppScreenshotsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAppScreenshotsResponse := _AppScreenshotsResponse{}

	if err = json.Unmarshal(bytes, &varAppScreenshotsResponse); err == nil {
		*o = AppScreenshotsResponse(varAppScreenshotsResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "links")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppScreenshotsResponse struct {
	value *AppScreenshotsResponse
	isSet bool
}

func (v NullableAppScreenshotsResponse) Get() *AppScreenshotsResponse {
	return v.value
}

func (v *NullableAppScreenshotsResponse) Set(val *AppScreenshotsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppScreenshotsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppScreenshotsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppScreenshotsResponse(val *AppScreenshotsResponse) *NullableAppScreenshotsResponse {
	return &NullableAppScreenshotsResponse{value: val, isSet: true}
}

func (v NullableAppScreenshotsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppScreenshotsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


