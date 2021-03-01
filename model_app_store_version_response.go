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

// AppStoreVersionResponse struct for AppStoreVersionResponse
type AppStoreVersionResponse struct {
	Data                 AppStoreVersion              `json:"data"`
	Included             *[]OneOfAgeRatingDeclaration `json:"included,omitempty"`
	Links                DocumentLinks                `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _AppStoreVersionResponse AppStoreVersionResponse
type OneOfAgeRatingDeclaration struct{}

// NewAppStoreVersionResponse instantiates a new AppStoreVersionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionResponse(data AppStoreVersion, links DocumentLinks) *AppStoreVersionResponse {
	this := AppStoreVersionResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppStoreVersionResponseWithDefaults instantiates a new AppStoreVersionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionResponseWithDefaults() *AppStoreVersionResponse {
	this := AppStoreVersionResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreVersionResponse) GetData() AppStoreVersion {
	if o == nil {
		var ret AppStoreVersion
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionResponse) GetDataOk() (*AppStoreVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppStoreVersionResponse) SetData(v AppStoreVersion) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppStoreVersionResponse) GetIncluded() []OneOfAgeRatingDeclaration {
	if o == nil || o.Included == nil {
		var ret []OneOfAgeRatingDeclaration
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionResponse) GetIncludedOk() (*[]OneOfAgeRatingDeclaration, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppStoreVersionResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []OneOfAgeRatingDeclaration and assigns it to the Included field.
func (o *AppStoreVersionResponse) SetIncluded(v []OneOfAgeRatingDeclaration) {
	o.Included = &v
}

// GetLinks returns the Links field value
func (o *AppStoreVersionResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppStoreVersionResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppStoreVersionResponse) MarshalJSON() ([]byte, error) {
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

func (o *AppStoreVersionResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAppStoreVersionResponse := _AppStoreVersionResponse{}

	if err = json.Unmarshal(bytes, &varAppStoreVersionResponse); err == nil {
		*o = AppStoreVersionResponse(varAppStoreVersionResponse)
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

type NullableAppStoreVersionResponse struct {
	value *AppStoreVersionResponse
	isSet bool
}

func (v NullableAppStoreVersionResponse) Get() *AppStoreVersionResponse {
	return v.value
}

func (v *NullableAppStoreVersionResponse) Set(val *AppStoreVersionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionResponse(val *AppStoreVersionResponse) *NullableAppStoreVersionResponse {
	return &NullableAppStoreVersionResponse{value: val, isSet: true}
}

func (v NullableAppStoreVersionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
