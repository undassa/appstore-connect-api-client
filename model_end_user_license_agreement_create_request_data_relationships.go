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

// EndUserLicenseAgreementCreateRequestDataRelationships struct for EndUserLicenseAgreementCreateRequestDataRelationships
type EndUserLicenseAgreementCreateRequestDataRelationships struct {
	App AppPreOrderCreateRequestDataRelationshipsApp `json:"app"`
	Territories EndUserLicenseAgreementCreateRequestDataRelationshipsTerritories `json:"territories"`
	AdditionalProperties map[string]interface{}
}

type _EndUserLicenseAgreementCreateRequestDataRelationships EndUserLicenseAgreementCreateRequestDataRelationships

// NewEndUserLicenseAgreementCreateRequestDataRelationships instantiates a new EndUserLicenseAgreementCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndUserLicenseAgreementCreateRequestDataRelationships(app AppPreOrderCreateRequestDataRelationshipsApp, territories EndUserLicenseAgreementCreateRequestDataRelationshipsTerritories, ) *EndUserLicenseAgreementCreateRequestDataRelationships {
	this := EndUserLicenseAgreementCreateRequestDataRelationships{}
	this.App = app
	this.Territories = territories
	return &this
}

// NewEndUserLicenseAgreementCreateRequestDataRelationshipsWithDefaults instantiates a new EndUserLicenseAgreementCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndUserLicenseAgreementCreateRequestDataRelationshipsWithDefaults() *EndUserLicenseAgreementCreateRequestDataRelationships {
	this := EndUserLicenseAgreementCreateRequestDataRelationships{}
	return &this
}

// GetApp returns the App field value
func (o *EndUserLicenseAgreementCreateRequestDataRelationships) GetApp() AppPreOrderCreateRequestDataRelationshipsApp {
	if o == nil  {
		var ret AppPreOrderCreateRequestDataRelationshipsApp
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *EndUserLicenseAgreementCreateRequestDataRelationships) GetAppOk() (*AppPreOrderCreateRequestDataRelationshipsApp, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *EndUserLicenseAgreementCreateRequestDataRelationships) SetApp(v AppPreOrderCreateRequestDataRelationshipsApp) {
	o.App = v
}

// GetTerritories returns the Territories field value
func (o *EndUserLicenseAgreementCreateRequestDataRelationships) GetTerritories() EndUserLicenseAgreementCreateRequestDataRelationshipsTerritories {
	if o == nil  {
		var ret EndUserLicenseAgreementCreateRequestDataRelationshipsTerritories
		return ret
	}

	return o.Territories
}

// GetTerritoriesOk returns a tuple with the Territories field value
// and a boolean to check if the value has been set.
func (o *EndUserLicenseAgreementCreateRequestDataRelationships) GetTerritoriesOk() (*EndUserLicenseAgreementCreateRequestDataRelationshipsTerritories, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Territories, true
}

// SetTerritories sets field value
func (o *EndUserLicenseAgreementCreateRequestDataRelationships) SetTerritories(v EndUserLicenseAgreementCreateRequestDataRelationshipsTerritories) {
	o.Territories = v
}

func (o EndUserLicenseAgreementCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["app"] = o.App
	}
	if true {
		toSerialize["territories"] = o.Territories
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EndUserLicenseAgreementCreateRequestDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	varEndUserLicenseAgreementCreateRequestDataRelationships := _EndUserLicenseAgreementCreateRequestDataRelationships{}

	if err = json.Unmarshal(bytes, &varEndUserLicenseAgreementCreateRequestDataRelationships); err == nil {
		*o = EndUserLicenseAgreementCreateRequestDataRelationships(varEndUserLicenseAgreementCreateRequestDataRelationships)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "app")
		delete(additionalProperties, "territories")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEndUserLicenseAgreementCreateRequestDataRelationships struct {
	value *EndUserLicenseAgreementCreateRequestDataRelationships
	isSet bool
}

func (v NullableEndUserLicenseAgreementCreateRequestDataRelationships) Get() *EndUserLicenseAgreementCreateRequestDataRelationships {
	return v.value
}

func (v *NullableEndUserLicenseAgreementCreateRequestDataRelationships) Set(val *EndUserLicenseAgreementCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableEndUserLicenseAgreementCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableEndUserLicenseAgreementCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndUserLicenseAgreementCreateRequestDataRelationships(val *EndUserLicenseAgreementCreateRequestDataRelationships) *NullableEndUserLicenseAgreementCreateRequestDataRelationships {
	return &NullableEndUserLicenseAgreementCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableEndUserLicenseAgreementCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndUserLicenseAgreementCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


