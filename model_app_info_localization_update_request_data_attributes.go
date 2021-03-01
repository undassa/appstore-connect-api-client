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

// AppInfoLocalizationUpdateRequestDataAttributes struct for AppInfoLocalizationUpdateRequestDataAttributes
type AppInfoLocalizationUpdateRequestDataAttributes struct {
	Name *string `json:"name,omitempty"`
	Subtitle *string `json:"subtitle,omitempty"`
	PrivacyPolicyUrl *string `json:"privacyPolicyUrl,omitempty"`
	PrivacyPolicyText *string `json:"privacyPolicyText,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppInfoLocalizationUpdateRequestDataAttributes AppInfoLocalizationUpdateRequestDataAttributes

// NewAppInfoLocalizationUpdateRequestDataAttributes instantiates a new AppInfoLocalizationUpdateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppInfoLocalizationUpdateRequestDataAttributes() *AppInfoLocalizationUpdateRequestDataAttributes {
	this := AppInfoLocalizationUpdateRequestDataAttributes{}
	return &this
}

// NewAppInfoLocalizationUpdateRequestDataAttributesWithDefaults instantiates a new AppInfoLocalizationUpdateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppInfoLocalizationUpdateRequestDataAttributesWithDefaults() *AppInfoLocalizationUpdateRequestDataAttributes {
	this := AppInfoLocalizationUpdateRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetSubtitle returns the Subtitle field value if set, zero value otherwise.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) GetSubtitle() string {
	if o == nil || o.Subtitle == nil {
		var ret string
		return ret
	}
	return *o.Subtitle
}

// GetSubtitleOk returns a tuple with the Subtitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) GetSubtitleOk() (*string, bool) {
	if o == nil || o.Subtitle == nil {
		return nil, false
	}
	return o.Subtitle, true
}

// HasSubtitle returns a boolean if a field has been set.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) HasSubtitle() bool {
	if o != nil && o.Subtitle != nil {
		return true
	}

	return false
}

// SetSubtitle gets a reference to the given string and assigns it to the Subtitle field.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) SetSubtitle(v string) {
	o.Subtitle = &v
}

// GetPrivacyPolicyUrl returns the PrivacyPolicyUrl field value if set, zero value otherwise.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) GetPrivacyPolicyUrl() string {
	if o == nil || o.PrivacyPolicyUrl == nil {
		var ret string
		return ret
	}
	return *o.PrivacyPolicyUrl
}

// GetPrivacyPolicyUrlOk returns a tuple with the PrivacyPolicyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) GetPrivacyPolicyUrlOk() (*string, bool) {
	if o == nil || o.PrivacyPolicyUrl == nil {
		return nil, false
	}
	return o.PrivacyPolicyUrl, true
}

// HasPrivacyPolicyUrl returns a boolean if a field has been set.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) HasPrivacyPolicyUrl() bool {
	if o != nil && o.PrivacyPolicyUrl != nil {
		return true
	}

	return false
}

// SetPrivacyPolicyUrl gets a reference to the given string and assigns it to the PrivacyPolicyUrl field.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) SetPrivacyPolicyUrl(v string) {
	o.PrivacyPolicyUrl = &v
}

// GetPrivacyPolicyText returns the PrivacyPolicyText field value if set, zero value otherwise.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) GetPrivacyPolicyText() string {
	if o == nil || o.PrivacyPolicyText == nil {
		var ret string
		return ret
	}
	return *o.PrivacyPolicyText
}

// GetPrivacyPolicyTextOk returns a tuple with the PrivacyPolicyText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) GetPrivacyPolicyTextOk() (*string, bool) {
	if o == nil || o.PrivacyPolicyText == nil {
		return nil, false
	}
	return o.PrivacyPolicyText, true
}

// HasPrivacyPolicyText returns a boolean if a field has been set.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) HasPrivacyPolicyText() bool {
	if o != nil && o.PrivacyPolicyText != nil {
		return true
	}

	return false
}

// SetPrivacyPolicyText gets a reference to the given string and assigns it to the PrivacyPolicyText field.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) SetPrivacyPolicyText(v string) {
	o.PrivacyPolicyText = &v
}

func (o AppInfoLocalizationUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Subtitle != nil {
		toSerialize["subtitle"] = o.Subtitle
	}
	if o.PrivacyPolicyUrl != nil {
		toSerialize["privacyPolicyUrl"] = o.PrivacyPolicyUrl
	}
	if o.PrivacyPolicyText != nil {
		toSerialize["privacyPolicyText"] = o.PrivacyPolicyText
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppInfoLocalizationUpdateRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varAppInfoLocalizationUpdateRequestDataAttributes := _AppInfoLocalizationUpdateRequestDataAttributes{}

	if err = json.Unmarshal(bytes, &varAppInfoLocalizationUpdateRequestDataAttributes); err == nil {
		*o = AppInfoLocalizationUpdateRequestDataAttributes(varAppInfoLocalizationUpdateRequestDataAttributes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "subtitle")
		delete(additionalProperties, "privacyPolicyUrl")
		delete(additionalProperties, "privacyPolicyText")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppInfoLocalizationUpdateRequestDataAttributes struct {
	value *AppInfoLocalizationUpdateRequestDataAttributes
	isSet bool
}

func (v NullableAppInfoLocalizationUpdateRequestDataAttributes) Get() *AppInfoLocalizationUpdateRequestDataAttributes {
	return v.value
}

func (v *NullableAppInfoLocalizationUpdateRequestDataAttributes) Set(val *AppInfoLocalizationUpdateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppInfoLocalizationUpdateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppInfoLocalizationUpdateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppInfoLocalizationUpdateRequestDataAttributes(val *AppInfoLocalizationUpdateRequestDataAttributes) *NullableAppInfoLocalizationUpdateRequestDataAttributes {
	return &NullableAppInfoLocalizationUpdateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppInfoLocalizationUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppInfoLocalizationUpdateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

