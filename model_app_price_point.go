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

// AppPricePoint struct for AppPricePoint
type AppPricePoint struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppPricePointAttributes `json:"attributes,omitempty"`
	Relationships *AppPricePointRelationships `json:"relationships,omitempty"`
	Links ResourceLinks `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _AppPricePoint AppPricePoint

// NewAppPricePoint instantiates a new AppPricePoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPricePoint(type_ string, id string, links ResourceLinks, ) *AppPricePoint {
	this := AppPricePoint{}
	this.Type = type_
	this.Id = id
	this.Links = links
	return &this
}

// NewAppPricePointWithDefaults instantiates a new AppPricePoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPricePointWithDefaults() *AppPricePoint {
	this := AppPricePoint{}
	return &this
}

// GetType returns the Type field value
func (o *AppPricePoint) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppPricePoint) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppPricePoint) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppPricePoint) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppPricePoint) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppPricePoint) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppPricePoint) GetAttributes() AppPricePointAttributes {
	if o == nil || o.Attributes == nil {
		var ret AppPricePointAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricePoint) GetAttributesOk() (*AppPricePointAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppPricePoint) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppPricePointAttributes and assigns it to the Attributes field.
func (o *AppPricePoint) SetAttributes(v AppPricePointAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AppPricePoint) GetRelationships() AppPricePointRelationships {
	if o == nil || o.Relationships == nil {
		var ret AppPricePointRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricePoint) GetRelationshipsOk() (*AppPricePointRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AppPricePoint) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppPricePointRelationships and assigns it to the Relationships field.
func (o *AppPricePoint) SetRelationships(v AppPricePointRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value
func (o *AppPricePoint) GetLinks() ResourceLinks {
	if o == nil  {
		var ret ResourceLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppPricePoint) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppPricePoint) SetLinks(v ResourceLinks) {
	o.Links = v
}

func (o AppPricePoint) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppPricePoint) UnmarshalJSON(bytes []byte) (err error) {
	varAppPricePoint := _AppPricePoint{}

	if err = json.Unmarshal(bytes, &varAppPricePoint); err == nil {
		*o = AppPricePoint(varAppPricePoint)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppPricePoint struct {
	value *AppPricePoint
	isSet bool
}

func (v NullableAppPricePoint) Get() *AppPricePoint {
	return v.value
}

func (v *NullableAppPricePoint) Set(val *AppPricePoint) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPricePoint) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPricePoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPricePoint(val *AppPricePoint) *NullableAppPricePoint {
	return &NullableAppPricePoint{value: val, isSet: true}
}

func (v NullableAppPricePoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPricePoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


