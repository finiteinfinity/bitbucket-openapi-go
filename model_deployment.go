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

// Deployment A Bitbucket Deployment.
type Deployment struct {
	// The UUID identifying the deployment.
	Uuid *string `json:"uuid,omitempty"`
	State *DeploymentState `json:"state,omitempty"`
	Environment *DeploymentEnvironment `json:"environment,omitempty"`
	Release *DeploymentRelease `json:"release,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Deployment Deployment

// NewDeployment instantiates a new Deployment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeployment() *Deployment {
	this := Deployment{}
	return &this
}

// NewDeploymentWithDefaults instantiates a new Deployment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentWithDefaults() *Deployment {
	this := Deployment{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *Deployment) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *Deployment) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *Deployment) SetUuid(v string) {
	o.Uuid = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Deployment) GetState() DeploymentState {
	if o == nil || o.State == nil {
		var ret DeploymentState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetStateOk() (*DeploymentState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Deployment) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given DeploymentState and assigns it to the State field.
func (o *Deployment) SetState(v DeploymentState) {
	o.State = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *Deployment) GetEnvironment() DeploymentEnvironment {
	if o == nil || o.Environment == nil {
		var ret DeploymentEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetEnvironmentOk() (*DeploymentEnvironment, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *Deployment) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given DeploymentEnvironment and assigns it to the Environment field.
func (o *Deployment) SetEnvironment(v DeploymentEnvironment) {
	o.Environment = &v
}

// GetRelease returns the Release field value if set, zero value otherwise.
func (o *Deployment) GetRelease() DeploymentRelease {
	if o == nil || o.Release == nil {
		var ret DeploymentRelease
		return ret
	}
	return *o.Release
}

// GetReleaseOk returns a tuple with the Release field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetReleaseOk() (*DeploymentRelease, bool) {
	if o == nil || o.Release == nil {
		return nil, false
	}
	return o.Release, true
}

// HasRelease returns a boolean if a field has been set.
func (o *Deployment) HasRelease() bool {
	if o != nil && o.Release != nil {
		return true
	}

	return false
}

// SetRelease gets a reference to the given DeploymentRelease and assigns it to the Release field.
func (o *Deployment) SetRelease(v DeploymentRelease) {
	o.Release = &v
}

func (o Deployment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Release != nil {
		toSerialize["release"] = o.Release
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Deployment) UnmarshalJSON(bytes []byte) (err error) {
	varDeployment := _Deployment{}

	if err = json.Unmarshal(bytes, &varDeployment); err == nil {
		*o = Deployment(varDeployment)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "state")
		delete(additionalProperties, "environment")
		delete(additionalProperties, "release")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeployment struct {
	value *Deployment
	isSet bool
}

func (v NullableDeployment) Get() *Deployment {
	return v.value
}

func (v *NullableDeployment) Set(val *Deployment) {
	v.value = val
	v.isSet = true
}

func (v NullableDeployment) IsSet() bool {
	return v.isSet
}

func (v *NullableDeployment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeployment(val *Deployment) *NullableDeployment {
	return &NullableDeployment{value: val, isSet: true}
}

func (v NullableDeployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeployment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


