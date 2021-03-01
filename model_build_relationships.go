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

// BuildRelationships struct for BuildRelationships
type BuildRelationships struct {
	PreReleaseVersion *BuildRelationshipsPreReleaseVersion `json:"preReleaseVersion,omitempty"`
	IndividualTesters *BetaGroupRelationshipsBetaTesters `json:"individualTesters,omitempty"`
	BetaBuildLocalizations *BuildRelationshipsBetaBuildLocalizations `json:"betaBuildLocalizations,omitempty"`
	AppEncryptionDeclaration *BuildRelationshipsAppEncryptionDeclaration `json:"appEncryptionDeclaration,omitempty"`
	BetaAppReviewSubmission *BuildRelationshipsBetaAppReviewSubmission `json:"betaAppReviewSubmission,omitempty"`
	App *AppEncryptionDeclarationRelationshipsApp `json:"app,omitempty"`
	BuildBetaDetail *BuildRelationshipsBuildBetaDetail `json:"buildBetaDetail,omitempty"`
	AppStoreVersion *AppStoreReviewDetailRelationshipsAppStoreVersion `json:"appStoreVersion,omitempty"`
	Icons *BuildRelationshipsIcons `json:"icons,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BuildRelationships BuildRelationships

// NewBuildRelationships instantiates a new BuildRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildRelationships() *BuildRelationships {
	this := BuildRelationships{}
	return &this
}

// NewBuildRelationshipsWithDefaults instantiates a new BuildRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildRelationshipsWithDefaults() *BuildRelationships {
	this := BuildRelationships{}
	return &this
}

// GetPreReleaseVersion returns the PreReleaseVersion field value if set, zero value otherwise.
func (o *BuildRelationships) GetPreReleaseVersion() BuildRelationshipsPreReleaseVersion {
	if o == nil || o.PreReleaseVersion == nil {
		var ret BuildRelationshipsPreReleaseVersion
		return ret
	}
	return *o.PreReleaseVersion
}

// GetPreReleaseVersionOk returns a tuple with the PreReleaseVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildRelationships) GetPreReleaseVersionOk() (*BuildRelationshipsPreReleaseVersion, bool) {
	if o == nil || o.PreReleaseVersion == nil {
		return nil, false
	}
	return o.PreReleaseVersion, true
}

// HasPreReleaseVersion returns a boolean if a field has been set.
func (o *BuildRelationships) HasPreReleaseVersion() bool {
	if o != nil && o.PreReleaseVersion != nil {
		return true
	}

	return false
}

// SetPreReleaseVersion gets a reference to the given BuildRelationshipsPreReleaseVersion and assigns it to the PreReleaseVersion field.
func (o *BuildRelationships) SetPreReleaseVersion(v BuildRelationshipsPreReleaseVersion) {
	o.PreReleaseVersion = &v
}

// GetIndividualTesters returns the IndividualTesters field value if set, zero value otherwise.
func (o *BuildRelationships) GetIndividualTesters() BetaGroupRelationshipsBetaTesters {
	if o == nil || o.IndividualTesters == nil {
		var ret BetaGroupRelationshipsBetaTesters
		return ret
	}
	return *o.IndividualTesters
}

// GetIndividualTestersOk returns a tuple with the IndividualTesters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildRelationships) GetIndividualTestersOk() (*BetaGroupRelationshipsBetaTesters, bool) {
	if o == nil || o.IndividualTesters == nil {
		return nil, false
	}
	return o.IndividualTesters, true
}

// HasIndividualTesters returns a boolean if a field has been set.
func (o *BuildRelationships) HasIndividualTesters() bool {
	if o != nil && o.IndividualTesters != nil {
		return true
	}

	return false
}

// SetIndividualTesters gets a reference to the given BetaGroupRelationshipsBetaTesters and assigns it to the IndividualTesters field.
func (o *BuildRelationships) SetIndividualTesters(v BetaGroupRelationshipsBetaTesters) {
	o.IndividualTesters = &v
}

// GetBetaBuildLocalizations returns the BetaBuildLocalizations field value if set, zero value otherwise.
func (o *BuildRelationships) GetBetaBuildLocalizations() BuildRelationshipsBetaBuildLocalizations {
	if o == nil || o.BetaBuildLocalizations == nil {
		var ret BuildRelationshipsBetaBuildLocalizations
		return ret
	}
	return *o.BetaBuildLocalizations
}

// GetBetaBuildLocalizationsOk returns a tuple with the BetaBuildLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildRelationships) GetBetaBuildLocalizationsOk() (*BuildRelationshipsBetaBuildLocalizations, bool) {
	if o == nil || o.BetaBuildLocalizations == nil {
		return nil, false
	}
	return o.BetaBuildLocalizations, true
}

// HasBetaBuildLocalizations returns a boolean if a field has been set.
func (o *BuildRelationships) HasBetaBuildLocalizations() bool {
	if o != nil && o.BetaBuildLocalizations != nil {
		return true
	}

	return false
}

// SetBetaBuildLocalizations gets a reference to the given BuildRelationshipsBetaBuildLocalizations and assigns it to the BetaBuildLocalizations field.
func (o *BuildRelationships) SetBetaBuildLocalizations(v BuildRelationshipsBetaBuildLocalizations) {
	o.BetaBuildLocalizations = &v
}

// GetAppEncryptionDeclaration returns the AppEncryptionDeclaration field value if set, zero value otherwise.
func (o *BuildRelationships) GetAppEncryptionDeclaration() BuildRelationshipsAppEncryptionDeclaration {
	if o == nil || o.AppEncryptionDeclaration == nil {
		var ret BuildRelationshipsAppEncryptionDeclaration
		return ret
	}
	return *o.AppEncryptionDeclaration
}

// GetAppEncryptionDeclarationOk returns a tuple with the AppEncryptionDeclaration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildRelationships) GetAppEncryptionDeclarationOk() (*BuildRelationshipsAppEncryptionDeclaration, bool) {
	if o == nil || o.AppEncryptionDeclaration == nil {
		return nil, false
	}
	return o.AppEncryptionDeclaration, true
}

// HasAppEncryptionDeclaration returns a boolean if a field has been set.
func (o *BuildRelationships) HasAppEncryptionDeclaration() bool {
	if o != nil && o.AppEncryptionDeclaration != nil {
		return true
	}

	return false
}

// SetAppEncryptionDeclaration gets a reference to the given BuildRelationshipsAppEncryptionDeclaration and assigns it to the AppEncryptionDeclaration field.
func (o *BuildRelationships) SetAppEncryptionDeclaration(v BuildRelationshipsAppEncryptionDeclaration) {
	o.AppEncryptionDeclaration = &v
}

// GetBetaAppReviewSubmission returns the BetaAppReviewSubmission field value if set, zero value otherwise.
func (o *BuildRelationships) GetBetaAppReviewSubmission() BuildRelationshipsBetaAppReviewSubmission {
	if o == nil || o.BetaAppReviewSubmission == nil {
		var ret BuildRelationshipsBetaAppReviewSubmission
		return ret
	}
	return *o.BetaAppReviewSubmission
}

// GetBetaAppReviewSubmissionOk returns a tuple with the BetaAppReviewSubmission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildRelationships) GetBetaAppReviewSubmissionOk() (*BuildRelationshipsBetaAppReviewSubmission, bool) {
	if o == nil || o.BetaAppReviewSubmission == nil {
		return nil, false
	}
	return o.BetaAppReviewSubmission, true
}

// HasBetaAppReviewSubmission returns a boolean if a field has been set.
func (o *BuildRelationships) HasBetaAppReviewSubmission() bool {
	if o != nil && o.BetaAppReviewSubmission != nil {
		return true
	}

	return false
}

// SetBetaAppReviewSubmission gets a reference to the given BuildRelationshipsBetaAppReviewSubmission and assigns it to the BetaAppReviewSubmission field.
func (o *BuildRelationships) SetBetaAppReviewSubmission(v BuildRelationshipsBetaAppReviewSubmission) {
	o.BetaAppReviewSubmission = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *BuildRelationships) GetApp() AppEncryptionDeclarationRelationshipsApp {
	if o == nil || o.App == nil {
		var ret AppEncryptionDeclarationRelationshipsApp
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildRelationships) GetAppOk() (*AppEncryptionDeclarationRelationshipsApp, bool) {
	if o == nil || o.App == nil {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *BuildRelationships) HasApp() bool {
	if o != nil && o.App != nil {
		return true
	}

	return false
}

// SetApp gets a reference to the given AppEncryptionDeclarationRelationshipsApp and assigns it to the App field.
func (o *BuildRelationships) SetApp(v AppEncryptionDeclarationRelationshipsApp) {
	o.App = &v
}

// GetBuildBetaDetail returns the BuildBetaDetail field value if set, zero value otherwise.
func (o *BuildRelationships) GetBuildBetaDetail() BuildRelationshipsBuildBetaDetail {
	if o == nil || o.BuildBetaDetail == nil {
		var ret BuildRelationshipsBuildBetaDetail
		return ret
	}
	return *o.BuildBetaDetail
}

// GetBuildBetaDetailOk returns a tuple with the BuildBetaDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildRelationships) GetBuildBetaDetailOk() (*BuildRelationshipsBuildBetaDetail, bool) {
	if o == nil || o.BuildBetaDetail == nil {
		return nil, false
	}
	return o.BuildBetaDetail, true
}

// HasBuildBetaDetail returns a boolean if a field has been set.
func (o *BuildRelationships) HasBuildBetaDetail() bool {
	if o != nil && o.BuildBetaDetail != nil {
		return true
	}

	return false
}

// SetBuildBetaDetail gets a reference to the given BuildRelationshipsBuildBetaDetail and assigns it to the BuildBetaDetail field.
func (o *BuildRelationships) SetBuildBetaDetail(v BuildRelationshipsBuildBetaDetail) {
	o.BuildBetaDetail = &v
}

// GetAppStoreVersion returns the AppStoreVersion field value if set, zero value otherwise.
func (o *BuildRelationships) GetAppStoreVersion() AppStoreReviewDetailRelationshipsAppStoreVersion {
	if o == nil || o.AppStoreVersion == nil {
		var ret AppStoreReviewDetailRelationshipsAppStoreVersion
		return ret
	}
	return *o.AppStoreVersion
}

// GetAppStoreVersionOk returns a tuple with the AppStoreVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildRelationships) GetAppStoreVersionOk() (*AppStoreReviewDetailRelationshipsAppStoreVersion, bool) {
	if o == nil || o.AppStoreVersion == nil {
		return nil, false
	}
	return o.AppStoreVersion, true
}

// HasAppStoreVersion returns a boolean if a field has been set.
func (o *BuildRelationships) HasAppStoreVersion() bool {
	if o != nil && o.AppStoreVersion != nil {
		return true
	}

	return false
}

// SetAppStoreVersion gets a reference to the given AppStoreReviewDetailRelationshipsAppStoreVersion and assigns it to the AppStoreVersion field.
func (o *BuildRelationships) SetAppStoreVersion(v AppStoreReviewDetailRelationshipsAppStoreVersion) {
	o.AppStoreVersion = &v
}

// GetIcons returns the Icons field value if set, zero value otherwise.
func (o *BuildRelationships) GetIcons() BuildRelationshipsIcons {
	if o == nil || o.Icons == nil {
		var ret BuildRelationshipsIcons
		return ret
	}
	return *o.Icons
}

// GetIconsOk returns a tuple with the Icons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildRelationships) GetIconsOk() (*BuildRelationshipsIcons, bool) {
	if o == nil || o.Icons == nil {
		return nil, false
	}
	return o.Icons, true
}

// HasIcons returns a boolean if a field has been set.
func (o *BuildRelationships) HasIcons() bool {
	if o != nil && o.Icons != nil {
		return true
	}

	return false
}

// SetIcons gets a reference to the given BuildRelationshipsIcons and assigns it to the Icons field.
func (o *BuildRelationships) SetIcons(v BuildRelationshipsIcons) {
	o.Icons = &v
}

func (o BuildRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PreReleaseVersion != nil {
		toSerialize["preReleaseVersion"] = o.PreReleaseVersion
	}
	if o.IndividualTesters != nil {
		toSerialize["individualTesters"] = o.IndividualTesters
	}
	if o.BetaBuildLocalizations != nil {
		toSerialize["betaBuildLocalizations"] = o.BetaBuildLocalizations
	}
	if o.AppEncryptionDeclaration != nil {
		toSerialize["appEncryptionDeclaration"] = o.AppEncryptionDeclaration
	}
	if o.BetaAppReviewSubmission != nil {
		toSerialize["betaAppReviewSubmission"] = o.BetaAppReviewSubmission
	}
	if o.App != nil {
		toSerialize["app"] = o.App
	}
	if o.BuildBetaDetail != nil {
		toSerialize["buildBetaDetail"] = o.BuildBetaDetail
	}
	if o.AppStoreVersion != nil {
		toSerialize["appStoreVersion"] = o.AppStoreVersion
	}
	if o.Icons != nil {
		toSerialize["icons"] = o.Icons
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BuildRelationships) UnmarshalJSON(bytes []byte) (err error) {
	varBuildRelationships := _BuildRelationships{}

	if err = json.Unmarshal(bytes, &varBuildRelationships); err == nil {
		*o = BuildRelationships(varBuildRelationships)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "preReleaseVersion")
		delete(additionalProperties, "individualTesters")
		delete(additionalProperties, "betaBuildLocalizations")
		delete(additionalProperties, "appEncryptionDeclaration")
		delete(additionalProperties, "betaAppReviewSubmission")
		delete(additionalProperties, "app")
		delete(additionalProperties, "buildBetaDetail")
		delete(additionalProperties, "appStoreVersion")
		delete(additionalProperties, "icons")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBuildRelationships struct {
	value *BuildRelationships
	isSet bool
}

func (v NullableBuildRelationships) Get() *BuildRelationships {
	return v.value
}

func (v *NullableBuildRelationships) Set(val *BuildRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildRelationships(val *BuildRelationships) *NullableBuildRelationships {
	return &NullableBuildRelationships{value: val, isSet: true}
}

func (v NullableBuildRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

