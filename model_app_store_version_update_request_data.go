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

// AppStoreVersionUpdateRequestData struct for AppStoreVersionUpdateRequestData
type AppStoreVersionUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppStoreVersionUpdateRequestDataAttributes `json:"attributes,omitempty"`
	Relationships *AppStoreVersionUpdateRequestDataRelationships `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppStoreVersionUpdateRequestData AppStoreVersionUpdateRequestData

// NewAppStoreVersionUpdateRequestData instantiates a new AppStoreVersionUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionUpdateRequestData(type_ string, id string, ) *AppStoreVersionUpdateRequestData {
	this := AppStoreVersionUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppStoreVersionUpdateRequestDataWithDefaults instantiates a new AppStoreVersionUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionUpdateRequestDataWithDefaults() *AppStoreVersionUpdateRequestData {
	this := AppStoreVersionUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppStoreVersionUpdateRequestData) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppStoreVersionUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppStoreVersionUpdateRequestData) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppStoreVersionUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppStoreVersionUpdateRequestData) GetAttributes() AppStoreVersionUpdateRequestDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret AppStoreVersionUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionUpdateRequestData) GetAttributesOk() (*AppStoreVersionUpdateRequestDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppStoreVersionUpdateRequestData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppStoreVersionUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *AppStoreVersionUpdateRequestData) SetAttributes(v AppStoreVersionUpdateRequestDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AppStoreVersionUpdateRequestData) GetRelationships() AppStoreVersionUpdateRequestDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AppStoreVersionUpdateRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionUpdateRequestData) GetRelationshipsOk() (*AppStoreVersionUpdateRequestDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AppStoreVersionUpdateRequestData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppStoreVersionUpdateRequestDataRelationships and assigns it to the Relationships field.
func (o *AppStoreVersionUpdateRequestData) SetRelationships(v AppStoreVersionUpdateRequestDataRelationships) {
	o.Relationships = &v
}

func (o AppStoreVersionUpdateRequestData) MarshalJSON() ([]byte, error) {
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

func (o *AppStoreVersionUpdateRequestData) UnmarshalJSON(bytes []byte) (err error) {
	varAppStoreVersionUpdateRequestData := _AppStoreVersionUpdateRequestData{}

	if err = json.Unmarshal(bytes, &varAppStoreVersionUpdateRequestData); err == nil {
		*o = AppStoreVersionUpdateRequestData(varAppStoreVersionUpdateRequestData)
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

type NullableAppStoreVersionUpdateRequestData struct {
	value *AppStoreVersionUpdateRequestData
	isSet bool
}

func (v NullableAppStoreVersionUpdateRequestData) Get() *AppStoreVersionUpdateRequestData {
	return v.value
}

func (v *NullableAppStoreVersionUpdateRequestData) Set(val *AppStoreVersionUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionUpdateRequestData(val *AppStoreVersionUpdateRequestData) *NullableAppStoreVersionUpdateRequestData {
	return &NullableAppStoreVersionUpdateRequestData{value: val, isSet: true}
}

func (v NullableAppStoreVersionUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

