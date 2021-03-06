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

// BuildUpdateRequestData struct for BuildUpdateRequestData
type BuildUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *BuildUpdateRequestDataAttributes `json:"attributes,omitempty"`
	Relationships *BuildUpdateRequestDataRelationships `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BuildUpdateRequestData BuildUpdateRequestData

// NewBuildUpdateRequestData instantiates a new BuildUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildUpdateRequestData(type_ string, id string, ) *BuildUpdateRequestData {
	this := BuildUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewBuildUpdateRequestDataWithDefaults instantiates a new BuildUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildUpdateRequestDataWithDefaults() *BuildUpdateRequestData {
	this := BuildUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *BuildUpdateRequestData) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BuildUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BuildUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *BuildUpdateRequestData) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BuildUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BuildUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *BuildUpdateRequestData) GetAttributes() BuildUpdateRequestDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret BuildUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildUpdateRequestData) GetAttributesOk() (*BuildUpdateRequestDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *BuildUpdateRequestData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given BuildUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *BuildUpdateRequestData) SetAttributes(v BuildUpdateRequestDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BuildUpdateRequestData) GetRelationships() BuildUpdateRequestDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret BuildUpdateRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildUpdateRequestData) GetRelationshipsOk() (*BuildUpdateRequestDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BuildUpdateRequestData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given BuildUpdateRequestDataRelationships and assigns it to the Relationships field.
func (o *BuildUpdateRequestData) SetRelationships(v BuildUpdateRequestDataRelationships) {
	o.Relationships = &v
}

func (o BuildUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BuildUpdateRequestData) UnmarshalJSON(bytes []byte) (err error) {
	varBuildUpdateRequestData := _BuildUpdateRequestData{}

	if err = json.Unmarshal(bytes, &varBuildUpdateRequestData); err == nil {
		*o = BuildUpdateRequestData(varBuildUpdateRequestData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBuildUpdateRequestData struct {
	value *BuildUpdateRequestData
	isSet bool
}

func (v NullableBuildUpdateRequestData) Get() *BuildUpdateRequestData {
	return v.value
}

func (v *NullableBuildUpdateRequestData) Set(val *BuildUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildUpdateRequestData(val *BuildUpdateRequestData) *NullableBuildUpdateRequestData {
	return &NullableBuildUpdateRequestData{value: val, isSet: true}
}

func (v NullableBuildUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


