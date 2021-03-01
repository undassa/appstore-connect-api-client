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

// AppCategoryAttributes struct for AppCategoryAttributes
type AppCategoryAttributes struct {
	Platforms *[]Platform `json:"platforms,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppCategoryAttributes AppCategoryAttributes

// NewAppCategoryAttributes instantiates a new AppCategoryAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCategoryAttributes() *AppCategoryAttributes {
	this := AppCategoryAttributes{}
	return &this
}

// NewAppCategoryAttributesWithDefaults instantiates a new AppCategoryAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCategoryAttributesWithDefaults() *AppCategoryAttributes {
	this := AppCategoryAttributes{}
	return &this
}

// GetPlatforms returns the Platforms field value if set, zero value otherwise.
func (o *AppCategoryAttributes) GetPlatforms() []Platform {
	if o == nil || o.Platforms == nil {
		var ret []Platform
		return ret
	}
	return *o.Platforms
}

// GetPlatformsOk returns a tuple with the Platforms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCategoryAttributes) GetPlatformsOk() (*[]Platform, bool) {
	if o == nil || o.Platforms == nil {
		return nil, false
	}
	return o.Platforms, true
}

// HasPlatforms returns a boolean if a field has been set.
func (o *AppCategoryAttributes) HasPlatforms() bool {
	if o != nil && o.Platforms != nil {
		return true
	}

	return false
}

// SetPlatforms gets a reference to the given []Platform and assigns it to the Platforms field.
func (o *AppCategoryAttributes) SetPlatforms(v []Platform) {
	o.Platforms = &v
}

func (o AppCategoryAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Platforms != nil {
		toSerialize["platforms"] = o.Platforms
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppCategoryAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varAppCategoryAttributes := _AppCategoryAttributes{}

	if err = json.Unmarshal(bytes, &varAppCategoryAttributes); err == nil {
		*o = AppCategoryAttributes(varAppCategoryAttributes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "platforms")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppCategoryAttributes struct {
	value *AppCategoryAttributes
	isSet bool
}

func (v NullableAppCategoryAttributes) Get() *AppCategoryAttributes {
	return v.value
}

func (v *NullableAppCategoryAttributes) Set(val *AppCategoryAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCategoryAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCategoryAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCategoryAttributes(val *AppCategoryAttributes) *NullableAppCategoryAttributes {
	return &NullableAppCategoryAttributes{value: val, isSet: true}
}

func (v NullableAppCategoryAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCategoryAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

