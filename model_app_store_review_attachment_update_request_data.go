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

// AppStoreReviewAttachmentUpdateRequestData struct for AppStoreReviewAttachmentUpdateRequestData
type AppStoreReviewAttachmentUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppScreenshotUpdateRequestDataAttributes `json:"attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppStoreReviewAttachmentUpdateRequestData AppStoreReviewAttachmentUpdateRequestData

// NewAppStoreReviewAttachmentUpdateRequestData instantiates a new AppStoreReviewAttachmentUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreReviewAttachmentUpdateRequestData(type_ string, id string, ) *AppStoreReviewAttachmentUpdateRequestData {
	this := AppStoreReviewAttachmentUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppStoreReviewAttachmentUpdateRequestDataWithDefaults instantiates a new AppStoreReviewAttachmentUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreReviewAttachmentUpdateRequestDataWithDefaults() *AppStoreReviewAttachmentUpdateRequestData {
	this := AppStoreReviewAttachmentUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppStoreReviewAttachmentUpdateRequestData) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppStoreReviewAttachmentUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppStoreReviewAttachmentUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppStoreReviewAttachmentUpdateRequestData) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppStoreReviewAttachmentUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppStoreReviewAttachmentUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppStoreReviewAttachmentUpdateRequestData) GetAttributes() AppScreenshotUpdateRequestDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret AppScreenshotUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewAttachmentUpdateRequestData) GetAttributesOk() (*AppScreenshotUpdateRequestDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppStoreReviewAttachmentUpdateRequestData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppScreenshotUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *AppStoreReviewAttachmentUpdateRequestData) SetAttributes(v AppScreenshotUpdateRequestDataAttributes) {
	o.Attributes = &v
}

func (o AppStoreReviewAttachmentUpdateRequestData) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppStoreReviewAttachmentUpdateRequestData) UnmarshalJSON(bytes []byte) (err error) {
	varAppStoreReviewAttachmentUpdateRequestData := _AppStoreReviewAttachmentUpdateRequestData{}

	if err = json.Unmarshal(bytes, &varAppStoreReviewAttachmentUpdateRequestData); err == nil {
		*o = AppStoreReviewAttachmentUpdateRequestData(varAppStoreReviewAttachmentUpdateRequestData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppStoreReviewAttachmentUpdateRequestData struct {
	value *AppStoreReviewAttachmentUpdateRequestData
	isSet bool
}

func (v NullableAppStoreReviewAttachmentUpdateRequestData) Get() *AppStoreReviewAttachmentUpdateRequestData {
	return v.value
}

func (v *NullableAppStoreReviewAttachmentUpdateRequestData) Set(val *AppStoreReviewAttachmentUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreReviewAttachmentUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreReviewAttachmentUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreReviewAttachmentUpdateRequestData(val *AppStoreReviewAttachmentUpdateRequestData) *NullableAppStoreReviewAttachmentUpdateRequestData {
	return &NullableAppStoreReviewAttachmentUpdateRequestData{value: val, isSet: true}
}

func (v NullableAppStoreReviewAttachmentUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreReviewAttachmentUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


