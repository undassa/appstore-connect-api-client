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
	"fmt"
)

// PhasedReleaseState the model 'PhasedReleaseState'
type PhasedReleaseState string

// List of PhasedReleaseState
const (
	INACTIVE PhasedReleaseState = "INACTIVE"
	ACTIVE PhasedReleaseState = "ACTIVE"
	PAUSED PhasedReleaseState = "PAUSED"
	COMPLETE PhasedReleaseState = "COMPLETE"
)

func (v *PhasedReleaseState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PhasedReleaseState(value)
	for _, existing := range []PhasedReleaseState{ "INACTIVE", "ACTIVE", "PAUSED", "COMPLETE",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PhasedReleaseState", value)
}

// Ptr returns reference to PhasedReleaseState value
func (v PhasedReleaseState) Ptr() *PhasedReleaseState {
	return &v
}

type NullablePhasedReleaseState struct {
	value *PhasedReleaseState
	isSet bool
}

func (v NullablePhasedReleaseState) Get() *PhasedReleaseState {
	return v.value
}

func (v *NullablePhasedReleaseState) Set(val *PhasedReleaseState) {
	v.value = val
	v.isSet = true
}

func (v NullablePhasedReleaseState) IsSet() bool {
	return v.isSet
}

func (v *NullablePhasedReleaseState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhasedReleaseState(val *PhasedReleaseState) *NullablePhasedReleaseState {
	return &NullablePhasedReleaseState{value: val, isSet: true}
}

func (v NullablePhasedReleaseState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhasedReleaseState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
