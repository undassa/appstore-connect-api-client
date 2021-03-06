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

// AppPricesResponse struct for AppPricesResponse
type AppPricesResponse struct {
	Data []AppPrice `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppPricesResponse AppPricesResponse

// NewAppPricesResponse instantiates a new AppPricesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPricesResponse(data []AppPrice, links PagedDocumentLinks, ) *AppPricesResponse {
	this := AppPricesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppPricesResponseWithDefaults instantiates a new AppPricesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPricesResponseWithDefaults() *AppPricesResponse {
	this := AppPricesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppPricesResponse) GetData() []AppPrice {
	if o == nil  {
		var ret []AppPrice
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppPricesResponse) GetDataOk() (*[]AppPrice, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppPricesResponse) SetData(v []AppPrice) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *AppPricesResponse) GetLinks() PagedDocumentLinks {
	if o == nil  {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppPricesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppPricesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppPricesResponse) GetMeta() PagingInformation {
	if o == nil || o.Meta == nil {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppPricesResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppPricesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o AppPricesResponse) MarshalJSON() ([]byte, error) {
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

func (o *AppPricesResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAppPricesResponse := _AppPricesResponse{}

	if err = json.Unmarshal(bytes, &varAppPricesResponse); err == nil {
		*o = AppPricesResponse(varAppPricesResponse)
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

type NullableAppPricesResponse struct {
	value *AppPricesResponse
	isSet bool
}

func (v NullableAppPricesResponse) Get() *AppPricesResponse {
	return v.value
}

func (v *NullableAppPricesResponse) Set(val *AppPricesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPricesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPricesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPricesResponse(val *AppPricesResponse) *NullableAppPricesResponse {
	return &NullableAppPricesResponse{value: val, isSet: true}
}

func (v NullableAppPricesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPricesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


