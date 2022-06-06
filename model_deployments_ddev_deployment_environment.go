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
)

// DeploymentsDdevDeploymentEnvironment struct for DeploymentsDdevDeploymentEnvironment
type DeploymentsDdevDeploymentEnvironment struct {
	Object
	// The UUID identifying the environment.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the environment.
	Name *string `json:"name,omitempty"`
}

// NewDeploymentsDdevDeploymentEnvironment instantiates a new DeploymentsDdevDeploymentEnvironment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentsDdevDeploymentEnvironment() *DeploymentsDdevDeploymentEnvironment {
	this := DeploymentsDdevDeploymentEnvironment{}
	return &this
}

// NewDeploymentsDdevDeploymentEnvironmentWithDefaults instantiates a new DeploymentsDdevDeploymentEnvironment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentsDdevDeploymentEnvironmentWithDefaults() *DeploymentsDdevDeploymentEnvironment {
	this := DeploymentsDdevDeploymentEnvironment{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *DeploymentsDdevDeploymentEnvironment) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentsDdevDeploymentEnvironment) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *DeploymentsDdevDeploymentEnvironment) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *DeploymentsDdevDeploymentEnvironment) SetUuid(v string) {
	o.Uuid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentsDdevDeploymentEnvironment) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentsDdevDeploymentEnvironment) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentsDdevDeploymentEnvironment) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentsDdevDeploymentEnvironment) SetName(v string) {
	o.Name = &v
}

func (o DeploymentsDdevDeploymentEnvironment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedObject, errObject := json.Marshal(o.Object)
	if errObject != nil {
		return []byte{}, errObject
	}
	errObject = json.Unmarshal([]byte(serializedObject), &toSerialize)
	if errObject != nil {
		return []byte{}, errObject
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentsDdevDeploymentEnvironment struct {
	value *DeploymentsDdevDeploymentEnvironment
	isSet bool
}

func (v NullableDeploymentsDdevDeploymentEnvironment) Get() *DeploymentsDdevDeploymentEnvironment {
	return v.value
}

func (v *NullableDeploymentsDdevDeploymentEnvironment) Set(val *DeploymentsDdevDeploymentEnvironment) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentsDdevDeploymentEnvironment) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentsDdevDeploymentEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentsDdevDeploymentEnvironment(val *DeploymentsDdevDeploymentEnvironment) *NullableDeploymentsDdevDeploymentEnvironment {
	return &NullableDeploymentsDdevDeploymentEnvironment{value: val, isSet: true}
}

func (v NullableDeploymentsDdevDeploymentEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentsDdevDeploymentEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

