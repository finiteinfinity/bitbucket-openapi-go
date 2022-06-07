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

// DeploymentStateCompletedStatusSuccessfulAllOf A SUCCESSFUL completed deployment status.
type DeploymentStateCompletedStatusSuccessfulAllOf struct {
	// The name of the completed deployment status (SUCCESSFUL).
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeploymentStateCompletedStatusSuccessfulAllOf DeploymentStateCompletedStatusSuccessfulAllOf

// NewDeploymentStateCompletedStatusSuccessfulAllOf instantiates a new DeploymentStateCompletedStatusSuccessfulAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentStateCompletedStatusSuccessfulAllOf() *DeploymentStateCompletedStatusSuccessfulAllOf {
	this := DeploymentStateCompletedStatusSuccessfulAllOf{}
	return &this
}

// NewDeploymentStateCompletedStatusSuccessfulAllOfWithDefaults instantiates a new DeploymentStateCompletedStatusSuccessfulAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentStateCompletedStatusSuccessfulAllOfWithDefaults() *DeploymentStateCompletedStatusSuccessfulAllOf {
	this := DeploymentStateCompletedStatusSuccessfulAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentStateCompletedStatusSuccessfulAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStateCompletedStatusSuccessfulAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentStateCompletedStatusSuccessfulAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentStateCompletedStatusSuccessfulAllOf) SetName(v string) {
	o.Name = &v
}

func (o DeploymentStateCompletedStatusSuccessfulAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeploymentStateCompletedStatusSuccessfulAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varDeploymentStateCompletedStatusSuccessfulAllOf := _DeploymentStateCompletedStatusSuccessfulAllOf{}

	if err = json.Unmarshal(bytes, &varDeploymentStateCompletedStatusSuccessfulAllOf); err == nil {
		*o = DeploymentStateCompletedStatusSuccessfulAllOf(varDeploymentStateCompletedStatusSuccessfulAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeploymentStateCompletedStatusSuccessfulAllOf struct {
	value *DeploymentStateCompletedStatusSuccessfulAllOf
	isSet bool
}

func (v NullableDeploymentStateCompletedStatusSuccessfulAllOf) Get() *DeploymentStateCompletedStatusSuccessfulAllOf {
	return v.value
}

func (v *NullableDeploymentStateCompletedStatusSuccessfulAllOf) Set(val *DeploymentStateCompletedStatusSuccessfulAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentStateCompletedStatusSuccessfulAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentStateCompletedStatusSuccessfulAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentStateCompletedStatusSuccessfulAllOf(val *DeploymentStateCompletedStatusSuccessfulAllOf) *NullableDeploymentStateCompletedStatusSuccessfulAllOf {
	return &NullableDeploymentStateCompletedStatusSuccessfulAllOf{value: val, isSet: true}
}

func (v NullableDeploymentStateCompletedStatusSuccessfulAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentStateCompletedStatusSuccessfulAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


