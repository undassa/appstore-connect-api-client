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

// AppPreviewSetCreateRequestData struct for AppPreviewSetCreateRequestData
type AppPreviewSetCreateRequestData struct {
	Type string `json:"type"`
	Attributes AppPreviewSetCreateRequestDataAttributes `json:"attributes"`
	Relationships AppPreviewSetCreateRequestDataRelationships `json:"relationships"`
	AdditionalProperties map[string]interface{}
}

type _AppPreviewSetCreateRequestData AppPreviewSetCreateRequestData

// NewAppPreviewSetCreateRequestData instantiates a new AppPreviewSetCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPreviewSetCreateRequestData(type_ string, attributes AppPreviewSetCreateRequestDataAttributes, relationships AppPreviewSetCreateRequestDataRelationships, ) *AppPreviewSetCreateRequestData {
	this := AppPreviewSetCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewAppPreviewSetCreateRequestDataWithDefaults instantiates a new AppPreviewSetCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPreviewSetCreateRequestDataWithDefaults() *AppPreviewSetCreateRequestData {
	this := AppPreviewSetCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppPreviewSetCreateRequestData) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppPreviewSetCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppPreviewSetCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AppPreviewSetCreateRequestData) GetAttributes() AppPreviewSetCreateRequestDataAttributes {
	if o == nil  {
		var ret AppPreviewSetCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AppPreviewSetCreateRequestData) GetAttributesOk() (*AppPreviewSetCreateRequestDataAttributes, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AppPreviewSetCreateRequestData) SetAttributes(v AppPreviewSetCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *AppPreviewSetCreateRequestData) GetRelationships() AppPreviewSetCreateRequestDataRelationships {
	if o == nil  {
		var ret AppPreviewSetCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AppPreviewSetCreateRequestData) GetRelationshipsOk() (*AppPreviewSetCreateRequestDataRelationships, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AppPreviewSetCreateRequestData) SetRelationships(v AppPreviewSetCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o AppPreviewSetCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["relationships"] = o.Relationships
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppPreviewSetCreateRequestData) UnmarshalJSON(bytes []byte) (err error) {
	varAppPreviewSetCreateRequestData := _AppPreviewSetCreateRequestData{}

	if err = json.Unmarshal(bytes, &varAppPreviewSetCreateRequestData); err == nil {
		*o = AppPreviewSetCreateRequestData(varAppPreviewSetCreateRequestData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppPreviewSetCreateRequestData struct {
	value *AppPreviewSetCreateRequestData
	isSet bool
}

func (v NullableAppPreviewSetCreateRequestData) Get() *AppPreviewSetCreateRequestData {
	return v.value
}

func (v *NullableAppPreviewSetCreateRequestData) Set(val *AppPreviewSetCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPreviewSetCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPreviewSetCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPreviewSetCreateRequestData(val *AppPreviewSetCreateRequestData) *NullableAppPreviewSetCreateRequestData {
	return &NullableAppPreviewSetCreateRequestData{value: val, isSet: true}
}

func (v NullableAppPreviewSetCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPreviewSetCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


