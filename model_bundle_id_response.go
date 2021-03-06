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

// BundleIdResponse struct for BundleIdResponse
type BundleIdResponse struct {
	Data                 BundleId              `json:"data"`
	Included             *[]OneOfProfileBundle `json:"included,omitempty"`
	Links                DocumentLinks         `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _BundleIdResponse BundleIdResponse
type OneOfProfileBundle struct{}

// NewBundleIdResponse instantiates a new BundleIdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleIdResponse(data BundleId, links DocumentLinks) *BundleIdResponse {
	this := BundleIdResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBundleIdResponseWithDefaults instantiates a new BundleIdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleIdResponseWithDefaults() *BundleIdResponse {
	this := BundleIdResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BundleIdResponse) GetData() BundleId {
	if o == nil {
		var ret BundleId
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BundleIdResponse) GetDataOk() (*BundleId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BundleIdResponse) SetData(v BundleId) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *BundleIdResponse) GetIncluded() []OneOfProfileBundle {
	if o == nil || o.Included == nil {
		var ret []OneOfProfileBundle
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleIdResponse) GetIncludedOk() (*[]OneOfProfileBundle, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *BundleIdResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []OneOfProfileBundle and assigns it to the Included field.
func (o *BundleIdResponse) SetIncluded(v []OneOfProfileBundle) {
	o.Included = &v
}

// GetLinks returns the Links field value
func (o *BundleIdResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BundleIdResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BundleIdResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o BundleIdResponse) MarshalJSON() ([]byte, error) {
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

func (o *BundleIdResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBundleIdResponse := _BundleIdResponse{}

	if err = json.Unmarshal(bytes, &varBundleIdResponse); err == nil {
		*o = BundleIdResponse(varBundleIdResponse)
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

type NullableBundleIdResponse struct {
	value *BundleIdResponse
	isSet bool
}

func (v NullableBundleIdResponse) Get() *BundleIdResponse {
	return v.value
}

func (v *NullableBundleIdResponse) Set(val *BundleIdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleIdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleIdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleIdResponse(val *BundleIdResponse) *NullableBundleIdResponse {
	return &NullableBundleIdResponse{value: val, isSet: true}
}

func (v NullableBundleIdResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleIdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
