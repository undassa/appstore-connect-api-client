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

// BundleIdRelationships struct for BundleIdRelationships
type BundleIdRelationships struct {
	Profiles *BundleIdRelationshipsProfiles `json:"profiles,omitempty"`
	BundleIdCapabilities *BundleIdRelationshipsBundleIdCapabilities `json:"bundleIdCapabilities,omitempty"`
	App *AppEncryptionDeclarationRelationshipsApp `json:"app,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BundleIdRelationships BundleIdRelationships

// NewBundleIdRelationships instantiates a new BundleIdRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleIdRelationships() *BundleIdRelationships {
	this := BundleIdRelationships{}
	return &this
}

// NewBundleIdRelationshipsWithDefaults instantiates a new BundleIdRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleIdRelationshipsWithDefaults() *BundleIdRelationships {
	this := BundleIdRelationships{}
	return &this
}

// GetProfiles returns the Profiles field value if set, zero value otherwise.
func (o *BundleIdRelationships) GetProfiles() BundleIdRelationshipsProfiles {
	if o == nil || o.Profiles == nil {
		var ret BundleIdRelationshipsProfiles
		return ret
	}
	return *o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleIdRelationships) GetProfilesOk() (*BundleIdRelationshipsProfiles, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *BundleIdRelationships) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given BundleIdRelationshipsProfiles and assigns it to the Profiles field.
func (o *BundleIdRelationships) SetProfiles(v BundleIdRelationshipsProfiles) {
	o.Profiles = &v
}

// GetBundleIdCapabilities returns the BundleIdCapabilities field value if set, zero value otherwise.
func (o *BundleIdRelationships) GetBundleIdCapabilities() BundleIdRelationshipsBundleIdCapabilities {
	if o == nil || o.BundleIdCapabilities == nil {
		var ret BundleIdRelationshipsBundleIdCapabilities
		return ret
	}
	return *o.BundleIdCapabilities
}

// GetBundleIdCapabilitiesOk returns a tuple with the BundleIdCapabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleIdRelationships) GetBundleIdCapabilitiesOk() (*BundleIdRelationshipsBundleIdCapabilities, bool) {
	if o == nil || o.BundleIdCapabilities == nil {
		return nil, false
	}
	return o.BundleIdCapabilities, true
}

// HasBundleIdCapabilities returns a boolean if a field has been set.
func (o *BundleIdRelationships) HasBundleIdCapabilities() bool {
	if o != nil && o.BundleIdCapabilities != nil {
		return true
	}

	return false
}

// SetBundleIdCapabilities gets a reference to the given BundleIdRelationshipsBundleIdCapabilities and assigns it to the BundleIdCapabilities field.
func (o *BundleIdRelationships) SetBundleIdCapabilities(v BundleIdRelationshipsBundleIdCapabilities) {
	o.BundleIdCapabilities = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *BundleIdRelationships) GetApp() AppEncryptionDeclarationRelationshipsApp {
	if o == nil || o.App == nil {
		var ret AppEncryptionDeclarationRelationshipsApp
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleIdRelationships) GetAppOk() (*AppEncryptionDeclarationRelationshipsApp, bool) {
	if o == nil || o.App == nil {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *BundleIdRelationships) HasApp() bool {
	if o != nil && o.App != nil {
		return true
	}

	return false
}

// SetApp gets a reference to the given AppEncryptionDeclarationRelationshipsApp and assigns it to the App field.
func (o *BundleIdRelationships) SetApp(v AppEncryptionDeclarationRelationshipsApp) {
	o.App = &v
}

func (o BundleIdRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Profiles != nil {
		toSerialize["profiles"] = o.Profiles
	}
	if o.BundleIdCapabilities != nil {
		toSerialize["bundleIdCapabilities"] = o.BundleIdCapabilities
	}
	if o.App != nil {
		toSerialize["app"] = o.App
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BundleIdRelationships) UnmarshalJSON(bytes []byte) (err error) {
	varBundleIdRelationships := _BundleIdRelationships{}

	if err = json.Unmarshal(bytes, &varBundleIdRelationships); err == nil {
		*o = BundleIdRelationships(varBundleIdRelationships)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "profiles")
		delete(additionalProperties, "bundleIdCapabilities")
		delete(additionalProperties, "app")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBundleIdRelationships struct {
	value *BundleIdRelationships
	isSet bool
}

func (v NullableBundleIdRelationships) Get() *BundleIdRelationships {
	return v.value
}

func (v *NullableBundleIdRelationships) Set(val *BundleIdRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleIdRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleIdRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleIdRelationships(val *BundleIdRelationships) *NullableBundleIdRelationships {
	return &NullableBundleIdRelationships{value: val, isSet: true}
}

func (v NullableBundleIdRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleIdRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


