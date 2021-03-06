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
	"time"
)

// BuildAttributes struct for BuildAttributes
type BuildAttributes struct {
	Version *string `json:"version,omitempty"`
	UploadedDate *time.Time `json:"uploadedDate,omitempty"`
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`
	Expired *bool `json:"expired,omitempty"`
	MinOsVersion *string `json:"minOsVersion,omitempty"`
	IconAssetToken *ImageAsset `json:"iconAssetToken,omitempty"`
	ProcessingState *string `json:"processingState,omitempty"`
	UsesNonExemptEncryption *bool `json:"usesNonExemptEncryption,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BuildAttributes BuildAttributes

// NewBuildAttributes instantiates a new BuildAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildAttributes() *BuildAttributes {
	this := BuildAttributes{}
	return &this
}

// NewBuildAttributesWithDefaults instantiates a new BuildAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildAttributesWithDefaults() *BuildAttributes {
	this := BuildAttributes{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BuildAttributes) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildAttributes) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BuildAttributes) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *BuildAttributes) SetVersion(v string) {
	o.Version = &v
}

// GetUploadedDate returns the UploadedDate field value if set, zero value otherwise.
func (o *BuildAttributes) GetUploadedDate() time.Time {
	if o == nil || o.UploadedDate == nil {
		var ret time.Time
		return ret
	}
	return *o.UploadedDate
}

// GetUploadedDateOk returns a tuple with the UploadedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildAttributes) GetUploadedDateOk() (*time.Time, bool) {
	if o == nil || o.UploadedDate == nil {
		return nil, false
	}
	return o.UploadedDate, true
}

// HasUploadedDate returns a boolean if a field has been set.
func (o *BuildAttributes) HasUploadedDate() bool {
	if o != nil && o.UploadedDate != nil {
		return true
	}

	return false
}

// SetUploadedDate gets a reference to the given time.Time and assigns it to the UploadedDate field.
func (o *BuildAttributes) SetUploadedDate(v time.Time) {
	o.UploadedDate = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *BuildAttributes) GetExpirationDate() time.Time {
	if o == nil || o.ExpirationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildAttributes) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil || o.ExpirationDate == nil {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *BuildAttributes) HasExpirationDate() bool {
	if o != nil && o.ExpirationDate != nil {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given time.Time and assigns it to the ExpirationDate field.
func (o *BuildAttributes) SetExpirationDate(v time.Time) {
	o.ExpirationDate = &v
}

// GetExpired returns the Expired field value if set, zero value otherwise.
func (o *BuildAttributes) GetExpired() bool {
	if o == nil || o.Expired == nil {
		var ret bool
		return ret
	}
	return *o.Expired
}

// GetExpiredOk returns a tuple with the Expired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildAttributes) GetExpiredOk() (*bool, bool) {
	if o == nil || o.Expired == nil {
		return nil, false
	}
	return o.Expired, true
}

// HasExpired returns a boolean if a field has been set.
func (o *BuildAttributes) HasExpired() bool {
	if o != nil && o.Expired != nil {
		return true
	}

	return false
}

// SetExpired gets a reference to the given bool and assigns it to the Expired field.
func (o *BuildAttributes) SetExpired(v bool) {
	o.Expired = &v
}

// GetMinOsVersion returns the MinOsVersion field value if set, zero value otherwise.
func (o *BuildAttributes) GetMinOsVersion() string {
	if o == nil || o.MinOsVersion == nil {
		var ret string
		return ret
	}
	return *o.MinOsVersion
}

// GetMinOsVersionOk returns a tuple with the MinOsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildAttributes) GetMinOsVersionOk() (*string, bool) {
	if o == nil || o.MinOsVersion == nil {
		return nil, false
	}
	return o.MinOsVersion, true
}

// HasMinOsVersion returns a boolean if a field has been set.
func (o *BuildAttributes) HasMinOsVersion() bool {
	if o != nil && o.MinOsVersion != nil {
		return true
	}

	return false
}

// SetMinOsVersion gets a reference to the given string and assigns it to the MinOsVersion field.
func (o *BuildAttributes) SetMinOsVersion(v string) {
	o.MinOsVersion = &v
}

// GetIconAssetToken returns the IconAssetToken field value if set, zero value otherwise.
func (o *BuildAttributes) GetIconAssetToken() ImageAsset {
	if o == nil || o.IconAssetToken == nil {
		var ret ImageAsset
		return ret
	}
	return *o.IconAssetToken
}

// GetIconAssetTokenOk returns a tuple with the IconAssetToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildAttributes) GetIconAssetTokenOk() (*ImageAsset, bool) {
	if o == nil || o.IconAssetToken == nil {
		return nil, false
	}
	return o.IconAssetToken, true
}

// HasIconAssetToken returns a boolean if a field has been set.
func (o *BuildAttributes) HasIconAssetToken() bool {
	if o != nil && o.IconAssetToken != nil {
		return true
	}

	return false
}

// SetIconAssetToken gets a reference to the given ImageAsset and assigns it to the IconAssetToken field.
func (o *BuildAttributes) SetIconAssetToken(v ImageAsset) {
	o.IconAssetToken = &v
}

// GetProcessingState returns the ProcessingState field value if set, zero value otherwise.
func (o *BuildAttributes) GetProcessingState() string {
	if o == nil || o.ProcessingState == nil {
		var ret string
		return ret
	}
	return *o.ProcessingState
}

// GetProcessingStateOk returns a tuple with the ProcessingState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildAttributes) GetProcessingStateOk() (*string, bool) {
	if o == nil || o.ProcessingState == nil {
		return nil, false
	}
	return o.ProcessingState, true
}

// HasProcessingState returns a boolean if a field has been set.
func (o *BuildAttributes) HasProcessingState() bool {
	if o != nil && o.ProcessingState != nil {
		return true
	}

	return false
}

// SetProcessingState gets a reference to the given string and assigns it to the ProcessingState field.
func (o *BuildAttributes) SetProcessingState(v string) {
	o.ProcessingState = &v
}

// GetUsesNonExemptEncryption returns the UsesNonExemptEncryption field value if set, zero value otherwise.
func (o *BuildAttributes) GetUsesNonExemptEncryption() bool {
	if o == nil || o.UsesNonExemptEncryption == nil {
		var ret bool
		return ret
	}
	return *o.UsesNonExemptEncryption
}

// GetUsesNonExemptEncryptionOk returns a tuple with the UsesNonExemptEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildAttributes) GetUsesNonExemptEncryptionOk() (*bool, bool) {
	if o == nil || o.UsesNonExemptEncryption == nil {
		return nil, false
	}
	return o.UsesNonExemptEncryption, true
}

// HasUsesNonExemptEncryption returns a boolean if a field has been set.
func (o *BuildAttributes) HasUsesNonExemptEncryption() bool {
	if o != nil && o.UsesNonExemptEncryption != nil {
		return true
	}

	return false
}

// SetUsesNonExemptEncryption gets a reference to the given bool and assigns it to the UsesNonExemptEncryption field.
func (o *BuildAttributes) SetUsesNonExemptEncryption(v bool) {
	o.UsesNonExemptEncryption = &v
}

func (o BuildAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.UploadedDate != nil {
		toSerialize["uploadedDate"] = o.UploadedDate
	}
	if o.ExpirationDate != nil {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if o.Expired != nil {
		toSerialize["expired"] = o.Expired
	}
	if o.MinOsVersion != nil {
		toSerialize["minOsVersion"] = o.MinOsVersion
	}
	if o.IconAssetToken != nil {
		toSerialize["iconAssetToken"] = o.IconAssetToken
	}
	if o.ProcessingState != nil {
		toSerialize["processingState"] = o.ProcessingState
	}
	if o.UsesNonExemptEncryption != nil {
		toSerialize["usesNonExemptEncryption"] = o.UsesNonExemptEncryption
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BuildAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varBuildAttributes := _BuildAttributes{}

	if err = json.Unmarshal(bytes, &varBuildAttributes); err == nil {
		*o = BuildAttributes(varBuildAttributes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "version")
		delete(additionalProperties, "uploadedDate")
		delete(additionalProperties, "expirationDate")
		delete(additionalProperties, "expired")
		delete(additionalProperties, "minOsVersion")
		delete(additionalProperties, "iconAssetToken")
		delete(additionalProperties, "processingState")
		delete(additionalProperties, "usesNonExemptEncryption")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBuildAttributes struct {
	value *BuildAttributes
	isSet bool
}

func (v NullableBuildAttributes) Get() *BuildAttributes {
	return v.value
}

func (v *NullableBuildAttributes) Set(val *BuildAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildAttributes(val *BuildAttributes) *NullableBuildAttributes {
	return &NullableBuildAttributes{value: val, isSet: true}
}

func (v NullableBuildAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


