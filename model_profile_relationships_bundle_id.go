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

// ProfileRelationshipsBundleId struct for ProfileRelationshipsBundleId
type ProfileRelationshipsBundleId struct {
	Links *AppCategoryRelationshipsSubcategoriesLinks `json:"links,omitempty"`
	Data *BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProfileRelationshipsBundleId ProfileRelationshipsBundleId

// NewProfileRelationshipsBundleId instantiates a new ProfileRelationshipsBundleId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileRelationshipsBundleId() *ProfileRelationshipsBundleId {
	this := ProfileRelationshipsBundleId{}
	return &this
}

// NewProfileRelationshipsBundleIdWithDefaults instantiates a new ProfileRelationshipsBundleId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileRelationshipsBundleIdWithDefaults() *ProfileRelationshipsBundleId {
	this := ProfileRelationshipsBundleId{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ProfileRelationshipsBundleId) GetLinks() AppCategoryRelationshipsSubcategoriesLinks {
	if o == nil || o.Links == nil {
		var ret AppCategoryRelationshipsSubcategoriesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileRelationshipsBundleId) GetLinksOk() (*AppCategoryRelationshipsSubcategoriesLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ProfileRelationshipsBundleId) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppCategoryRelationshipsSubcategoriesLinks and assigns it to the Links field.
func (o *ProfileRelationshipsBundleId) SetLinks(v AppCategoryRelationshipsSubcategoriesLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ProfileRelationshipsBundleId) GetData() BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData {
	if o == nil || o.Data == nil {
		var ret BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileRelationshipsBundleId) GetDataOk() (*BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ProfileRelationshipsBundleId) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData and assigns it to the Data field.
func (o *ProfileRelationshipsBundleId) SetData(v BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData) {
	o.Data = &v
}

func (o ProfileRelationshipsBundleId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProfileRelationshipsBundleId) UnmarshalJSON(bytes []byte) (err error) {
	varProfileRelationshipsBundleId := _ProfileRelationshipsBundleId{}

	if err = json.Unmarshal(bytes, &varProfileRelationshipsBundleId); err == nil {
		*o = ProfileRelationshipsBundleId(varProfileRelationshipsBundleId)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "links")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProfileRelationshipsBundleId struct {
	value *ProfileRelationshipsBundleId
	isSet bool
}

func (v NullableProfileRelationshipsBundleId) Get() *ProfileRelationshipsBundleId {
	return v.value
}

func (v *NullableProfileRelationshipsBundleId) Set(val *ProfileRelationshipsBundleId) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileRelationshipsBundleId) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileRelationshipsBundleId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileRelationshipsBundleId(val *ProfileRelationshipsBundleId) *NullableProfileRelationshipsBundleId {
	return &NullableProfileRelationshipsBundleId{value: val, isSet: true}
}

func (v NullableProfileRelationshipsBundleId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileRelationshipsBundleId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


