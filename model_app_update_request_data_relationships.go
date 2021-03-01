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

// AppUpdateRequestDataRelationships struct for AppUpdateRequestDataRelationships
type AppUpdateRequestDataRelationships struct {
	Prices *AppUpdateRequestDataRelationshipsPrices `json:"prices,omitempty"`
	AvailableTerritories *AppUpdateRequestDataRelationshipsAvailableTerritories `json:"availableTerritories,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppUpdateRequestDataRelationships AppUpdateRequestDataRelationships

// NewAppUpdateRequestDataRelationships instantiates a new AppUpdateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppUpdateRequestDataRelationships() *AppUpdateRequestDataRelationships {
	this := AppUpdateRequestDataRelationships{}
	return &this
}

// NewAppUpdateRequestDataRelationshipsWithDefaults instantiates a new AppUpdateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppUpdateRequestDataRelationshipsWithDefaults() *AppUpdateRequestDataRelationships {
	this := AppUpdateRequestDataRelationships{}
	return &this
}

// GetPrices returns the Prices field value if set, zero value otherwise.
func (o *AppUpdateRequestDataRelationships) GetPrices() AppUpdateRequestDataRelationshipsPrices {
	if o == nil || o.Prices == nil {
		var ret AppUpdateRequestDataRelationshipsPrices
		return ret
	}
	return *o.Prices
}

// GetPricesOk returns a tuple with the Prices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUpdateRequestDataRelationships) GetPricesOk() (*AppUpdateRequestDataRelationshipsPrices, bool) {
	if o == nil || o.Prices == nil {
		return nil, false
	}
	return o.Prices, true
}

// HasPrices returns a boolean if a field has been set.
func (o *AppUpdateRequestDataRelationships) HasPrices() bool {
	if o != nil && o.Prices != nil {
		return true
	}

	return false
}

// SetPrices gets a reference to the given AppUpdateRequestDataRelationshipsPrices and assigns it to the Prices field.
func (o *AppUpdateRequestDataRelationships) SetPrices(v AppUpdateRequestDataRelationshipsPrices) {
	o.Prices = &v
}

// GetAvailableTerritories returns the AvailableTerritories field value if set, zero value otherwise.
func (o *AppUpdateRequestDataRelationships) GetAvailableTerritories() AppUpdateRequestDataRelationshipsAvailableTerritories {
	if o == nil || o.AvailableTerritories == nil {
		var ret AppUpdateRequestDataRelationshipsAvailableTerritories
		return ret
	}
	return *o.AvailableTerritories
}

// GetAvailableTerritoriesOk returns a tuple with the AvailableTerritories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUpdateRequestDataRelationships) GetAvailableTerritoriesOk() (*AppUpdateRequestDataRelationshipsAvailableTerritories, bool) {
	if o == nil || o.AvailableTerritories == nil {
		return nil, false
	}
	return o.AvailableTerritories, true
}

// HasAvailableTerritories returns a boolean if a field has been set.
func (o *AppUpdateRequestDataRelationships) HasAvailableTerritories() bool {
	if o != nil && o.AvailableTerritories != nil {
		return true
	}

	return false
}

// SetAvailableTerritories gets a reference to the given AppUpdateRequestDataRelationshipsAvailableTerritories and assigns it to the AvailableTerritories field.
func (o *AppUpdateRequestDataRelationships) SetAvailableTerritories(v AppUpdateRequestDataRelationshipsAvailableTerritories) {
	o.AvailableTerritories = &v
}

func (o AppUpdateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Prices != nil {
		toSerialize["prices"] = o.Prices
	}
	if o.AvailableTerritories != nil {
		toSerialize["availableTerritories"] = o.AvailableTerritories
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppUpdateRequestDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	varAppUpdateRequestDataRelationships := _AppUpdateRequestDataRelationships{}

	if err = json.Unmarshal(bytes, &varAppUpdateRequestDataRelationships); err == nil {
		*o = AppUpdateRequestDataRelationships(varAppUpdateRequestDataRelationships)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "prices")
		delete(additionalProperties, "availableTerritories")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppUpdateRequestDataRelationships struct {
	value *AppUpdateRequestDataRelationships
	isSet bool
}

func (v NullableAppUpdateRequestDataRelationships) Get() *AppUpdateRequestDataRelationships {
	return v.value
}

func (v *NullableAppUpdateRequestDataRelationships) Set(val *AppUpdateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppUpdateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppUpdateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppUpdateRequestDataRelationships(val *AppUpdateRequestDataRelationships) *NullableAppUpdateRequestDataRelationships {
	return &NullableAppUpdateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableAppUpdateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppUpdateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

