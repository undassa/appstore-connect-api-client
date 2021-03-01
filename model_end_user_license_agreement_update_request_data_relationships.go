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

// EndUserLicenseAgreementUpdateRequestDataRelationships struct for EndUserLicenseAgreementUpdateRequestDataRelationships
type EndUserLicenseAgreementUpdateRequestDataRelationships struct {
	Territories *AppUpdateRequestDataRelationshipsAvailableTerritories `json:"territories,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EndUserLicenseAgreementUpdateRequestDataRelationships EndUserLicenseAgreementUpdateRequestDataRelationships

// NewEndUserLicenseAgreementUpdateRequestDataRelationships instantiates a new EndUserLicenseAgreementUpdateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndUserLicenseAgreementUpdateRequestDataRelationships() *EndUserLicenseAgreementUpdateRequestDataRelationships {
	this := EndUserLicenseAgreementUpdateRequestDataRelationships{}
	return &this
}

// NewEndUserLicenseAgreementUpdateRequestDataRelationshipsWithDefaults instantiates a new EndUserLicenseAgreementUpdateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndUserLicenseAgreementUpdateRequestDataRelationshipsWithDefaults() *EndUserLicenseAgreementUpdateRequestDataRelationships {
	this := EndUserLicenseAgreementUpdateRequestDataRelationships{}
	return &this
}

// GetTerritories returns the Territories field value if set, zero value otherwise.
func (o *EndUserLicenseAgreementUpdateRequestDataRelationships) GetTerritories() AppUpdateRequestDataRelationshipsAvailableTerritories {
	if o == nil || o.Territories == nil {
		var ret AppUpdateRequestDataRelationshipsAvailableTerritories
		return ret
	}
	return *o.Territories
}

// GetTerritoriesOk returns a tuple with the Territories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndUserLicenseAgreementUpdateRequestDataRelationships) GetTerritoriesOk() (*AppUpdateRequestDataRelationshipsAvailableTerritories, bool) {
	if o == nil || o.Territories == nil {
		return nil, false
	}
	return o.Territories, true
}

// HasTerritories returns a boolean if a field has been set.
func (o *EndUserLicenseAgreementUpdateRequestDataRelationships) HasTerritories() bool {
	if o != nil && o.Territories != nil {
		return true
	}

	return false
}

// SetTerritories gets a reference to the given AppUpdateRequestDataRelationshipsAvailableTerritories and assigns it to the Territories field.
func (o *EndUserLicenseAgreementUpdateRequestDataRelationships) SetTerritories(v AppUpdateRequestDataRelationshipsAvailableTerritories) {
	o.Territories = &v
}

func (o EndUserLicenseAgreementUpdateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Territories != nil {
		toSerialize["territories"] = o.Territories
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EndUserLicenseAgreementUpdateRequestDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	varEndUserLicenseAgreementUpdateRequestDataRelationships := _EndUserLicenseAgreementUpdateRequestDataRelationships{}

	if err = json.Unmarshal(bytes, &varEndUserLicenseAgreementUpdateRequestDataRelationships); err == nil {
		*o = EndUserLicenseAgreementUpdateRequestDataRelationships(varEndUserLicenseAgreementUpdateRequestDataRelationships)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "territories")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEndUserLicenseAgreementUpdateRequestDataRelationships struct {
	value *EndUserLicenseAgreementUpdateRequestDataRelationships
	isSet bool
}

func (v NullableEndUserLicenseAgreementUpdateRequestDataRelationships) Get() *EndUserLicenseAgreementUpdateRequestDataRelationships {
	return v.value
}

func (v *NullableEndUserLicenseAgreementUpdateRequestDataRelationships) Set(val *EndUserLicenseAgreementUpdateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableEndUserLicenseAgreementUpdateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableEndUserLicenseAgreementUpdateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndUserLicenseAgreementUpdateRequestDataRelationships(val *EndUserLicenseAgreementUpdateRequestDataRelationships) *NullableEndUserLicenseAgreementUpdateRequestDataRelationships {
	return &NullableEndUserLicenseAgreementUpdateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableEndUserLicenseAgreementUpdateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndUserLicenseAgreementUpdateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

