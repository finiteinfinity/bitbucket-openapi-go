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

// DeploymentStateCompletedStatusStopped struct for DeploymentStateCompletedStatusStopped
type DeploymentStateCompletedStatusStopped struct {
	DeploymentStateCompletedStatus
	// The name of the completed deployment status (STOPPED).
	Name *string `json:"name,omitempty"`
}

// NewDeploymentStateCompletedStatusStopped instantiates a new DeploymentStateCompletedStatusStopped object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentStateCompletedStatusStopped() *DeploymentStateCompletedStatusStopped {
	this := DeploymentStateCompletedStatusStopped{}
	return &this
}

// NewDeploymentStateCompletedStatusStoppedWithDefaults instantiates a new DeploymentStateCompletedStatusStopped object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentStateCompletedStatusStoppedWithDefaults() *DeploymentStateCompletedStatusStopped {
	this := DeploymentStateCompletedStatusStopped{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentStateCompletedStatusStopped) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStateCompletedStatusStopped) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentStateCompletedStatusStopped) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentStateCompletedStatusStopped) SetName(v string) {
	o.Name = &v
}

func (o DeploymentStateCompletedStatusStopped) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedDeploymentStateCompletedStatus, errDeploymentStateCompletedStatus := json.Marshal(o.DeploymentStateCompletedStatus)
	if errDeploymentStateCompletedStatus != nil {
		return []byte{}, errDeploymentStateCompletedStatus
	}
	errDeploymentStateCompletedStatus = json.Unmarshal([]byte(serializedDeploymentStateCompletedStatus), &toSerialize)
	if errDeploymentStateCompletedStatus != nil {
		return []byte{}, errDeploymentStateCompletedStatus
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentStateCompletedStatusStopped struct {
	value *DeploymentStateCompletedStatusStopped
	isSet bool
}

func (v NullableDeploymentStateCompletedStatusStopped) Get() *DeploymentStateCompletedStatusStopped {
	return v.value
}

func (v *NullableDeploymentStateCompletedStatusStopped) Set(val *DeploymentStateCompletedStatusStopped) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentStateCompletedStatusStopped) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentStateCompletedStatusStopped) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentStateCompletedStatusStopped(val *DeploymentStateCompletedStatusStopped) *NullableDeploymentStateCompletedStatusStopped {
	return &NullableDeploymentStateCompletedStatusStopped{value: val, isSet: true}
}

func (v NullableDeploymentStateCompletedStatusStopped) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentStateCompletedStatusStopped) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


