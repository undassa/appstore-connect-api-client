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

// DiagnosticLogsResponse struct for DiagnosticLogsResponse
type DiagnosticLogsResponse struct {
	Data []DiagnosticLog `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DiagnosticLogsResponse DiagnosticLogsResponse

// NewDiagnosticLogsResponse instantiates a new DiagnosticLogsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiagnosticLogsResponse(data []DiagnosticLog, links PagedDocumentLinks, ) *DiagnosticLogsResponse {
	this := DiagnosticLogsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewDiagnosticLogsResponseWithDefaults instantiates a new DiagnosticLogsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiagnosticLogsResponseWithDefaults() *DiagnosticLogsResponse {
	this := DiagnosticLogsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *DiagnosticLogsResponse) GetData() []DiagnosticLog {
	if o == nil  {
		var ret []DiagnosticLog
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DiagnosticLogsResponse) GetDataOk() (*[]DiagnosticLog, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DiagnosticLogsResponse) SetData(v []DiagnosticLog) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *DiagnosticLogsResponse) GetLinks() PagedDocumentLinks {
	if o == nil  {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *DiagnosticLogsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *DiagnosticLogsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *DiagnosticLogsResponse) GetMeta() PagingInformation {
	if o == nil || o.Meta == nil {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticLogsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *DiagnosticLogsResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *DiagnosticLogsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o DiagnosticLogsResponse) MarshalJSON() ([]byte, error) {
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

func (o *DiagnosticLogsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varDiagnosticLogsResponse := _DiagnosticLogsResponse{}

	if err = json.Unmarshal(bytes, &varDiagnosticLogsResponse); err == nil {
		*o = DiagnosticLogsResponse(varDiagnosticLogsResponse)
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

type NullableDiagnosticLogsResponse struct {
	value *DiagnosticLogsResponse
	isSet bool
}

func (v NullableDiagnosticLogsResponse) Get() *DiagnosticLogsResponse {
	return v.value
}

func (v *NullableDiagnosticLogsResponse) Set(val *DiagnosticLogsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDiagnosticLogsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDiagnosticLogsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiagnosticLogsResponse(val *DiagnosticLogsResponse) *NullableDiagnosticLogsResponse {
	return &NullableDiagnosticLogsResponse{value: val, isSet: true}
}

func (v NullableDiagnosticLogsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiagnosticLogsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

