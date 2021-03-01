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

// InAppPurchaseResponse struct for InAppPurchaseResponse
type InAppPurchaseResponse struct {
	Data InAppPurchase `json:"data"`
	Links DocumentLinks `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _InAppPurchaseResponse InAppPurchaseResponse

// NewInAppPurchaseResponse instantiates a new InAppPurchaseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseResponse(data InAppPurchase, links DocumentLinks, ) *InAppPurchaseResponse {
	this := InAppPurchaseResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewInAppPurchaseResponseWithDefaults instantiates a new InAppPurchaseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseResponseWithDefaults() *InAppPurchaseResponse {
	this := InAppPurchaseResponse{}
	return &this
}

// GetData returns the Data field value
func (o *InAppPurchaseResponse) GetData() InAppPurchase {
	if o == nil  {
		var ret InAppPurchase
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseResponse) GetDataOk() (*InAppPurchase, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InAppPurchaseResponse) SetData(v InAppPurchase) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *InAppPurchaseResponse) GetLinks() DocumentLinks {
	if o == nil  {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *InAppPurchaseResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o InAppPurchaseResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InAppPurchaseResponse) UnmarshalJSON(bytes []byte) (err error) {
	varInAppPurchaseResponse := _InAppPurchaseResponse{}

	if err = json.Unmarshal(bytes, &varInAppPurchaseResponse); err == nil {
		*o = InAppPurchaseResponse(varInAppPurchaseResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInAppPurchaseResponse struct {
	value *InAppPurchaseResponse
	isSet bool
}

func (v NullableInAppPurchaseResponse) Get() *InAppPurchaseResponse {
	return v.value
}

func (v *NullableInAppPurchaseResponse) Set(val *InAppPurchaseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseResponse(val *InAppPurchaseResponse) *NullableInAppPurchaseResponse {
	return &NullableInAppPurchaseResponse{value: val, isSet: true}
}

func (v NullableInAppPurchaseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

