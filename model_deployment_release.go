/*
Bitbucket API

Code against the Bitbucket API to automate simple tasks, embed Bitbucket data into your own site, build mobile or desktop apps, or even add custom UI add-ons into Bitbucket itself using the Connect framework.

API version: 2.0
Contact: support@bitbucket.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bitbucket

import (
	"encoding/json"
	"time"
)

// DeploymentRelease A Bitbucket Deployment Release.
type DeploymentRelease struct {
	// The UUID identifying the release.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the release.
	Name *string `json:"name,omitempty"`
	// Link to the pipeline that produced the release.
	Url *string `json:"url,omitempty"`
	Commit *Commit `json:"commit,omitempty"`
	// The timestamp when the release was created.
	CreatedOn *time.Time `json:"created_on,omitempty"`
}

// NewDeploymentRelease instantiates a new DeploymentRelease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentRelease() *DeploymentRelease {
	this := DeploymentRelease{}
	return &this
}

// NewDeploymentReleaseWithDefaults instantiates a new DeploymentRelease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentReleaseWithDefaults() *DeploymentRelease {
	this := DeploymentRelease{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *DeploymentRelease) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRelease) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *DeploymentRelease) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *DeploymentRelease) SetUuid(v string) {
	o.Uuid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentRelease) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRelease) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentRelease) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentRelease) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *DeploymentRelease) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRelease) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *DeploymentRelease) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *DeploymentRelease) SetUrl(v string) {
	o.Url = &v
}

// GetCommit returns the Commit field value if set, zero value otherwise.
func (o *DeploymentRelease) GetCommit() Commit {
	if o == nil || o.Commit == nil {
		var ret Commit
		return ret
	}
	return *o.Commit
}

// GetCommitOk returns a tuple with the Commit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRelease) GetCommitOk() (*Commit, bool) {
	if o == nil || o.Commit == nil {
		return nil, false
	}
	return o.Commit, true
}

// HasCommit returns a boolean if a field has been set.
func (o *DeploymentRelease) HasCommit() bool {
	if o != nil && o.Commit != nil {
		return true
	}

	return false
}

// SetCommit gets a reference to the given Commit and assigns it to the Commit field.
func (o *DeploymentRelease) SetCommit(v Commit) {
	o.Commit = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *DeploymentRelease) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRelease) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *DeploymentRelease) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *DeploymentRelease) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

func (o DeploymentRelease) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Commit != nil {
		toSerialize["commit"] = o.Commit
	}
	if o.CreatedOn != nil {
		toSerialize["created_on"] = o.CreatedOn
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentRelease struct {
	value *DeploymentRelease
	isSet bool
}

func (v NullableDeploymentRelease) Get() *DeploymentRelease {
	return v.value
}

func (v *NullableDeploymentRelease) Set(val *DeploymentRelease) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentRelease) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentRelease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentRelease(val *DeploymentRelease) *NullableDeploymentRelease {
	return &NullableDeploymentRelease{value: val, isSet: true}
}

func (v NullableDeploymentRelease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentRelease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


