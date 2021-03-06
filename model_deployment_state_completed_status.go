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
	"reflect"
	"strings"
)

// DeploymentStateCompletedStatus struct for DeploymentStateCompletedStatus
type DeploymentStateCompletedStatus struct {
	Object
	AdditionalProperties map[string]interface{}
}

type _DeploymentStateCompletedStatus DeploymentStateCompletedStatus

// NewDeploymentStateCompletedStatus instantiates a new DeploymentStateCompletedStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentStateCompletedStatus() *DeploymentStateCompletedStatus {
	this := DeploymentStateCompletedStatus{}
	return &this
}

// NewDeploymentStateCompletedStatusWithDefaults instantiates a new DeploymentStateCompletedStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentStateCompletedStatusWithDefaults() *DeploymentStateCompletedStatus {
	this := DeploymentStateCompletedStatus{}
	return &this
}

func (o DeploymentStateCompletedStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedObject, errObject := json.Marshal(o.Object)
	if errObject != nil {
		return []byte{}, errObject
	}
	errObject = json.Unmarshal([]byte(serializedObject), &toSerialize)
	if errObject != nil {
		return []byte{}, errObject
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeploymentStateCompletedStatus) UnmarshalJSON(bytes []byte) (err error) {
	type DeploymentStateCompletedStatusWithoutEmbeddedStruct struct {
	}

	varDeploymentStateCompletedStatusWithoutEmbeddedStruct := DeploymentStateCompletedStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varDeploymentStateCompletedStatusWithoutEmbeddedStruct)
	if err == nil {
		varDeploymentStateCompletedStatus := _DeploymentStateCompletedStatus{}
		*o = DeploymentStateCompletedStatus(varDeploymentStateCompletedStatus)
	} else {
		return err
	}

	varDeploymentStateCompletedStatus := _DeploymentStateCompletedStatus{}

	err = json.Unmarshal(bytes, &varDeploymentStateCompletedStatus)
	if err == nil {
		o.Object = varDeploymentStateCompletedStatus.Object
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectObject := reflect.ValueOf(o.Object)
		for i := 0; i < reflectObject.Type().NumField(); i++ {
			t := reflectObject.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeploymentStateCompletedStatus struct {
	value *DeploymentStateCompletedStatus
	isSet bool
}

func (v NullableDeploymentStateCompletedStatus) Get() *DeploymentStateCompletedStatus {
	return v.value
}

func (v *NullableDeploymentStateCompletedStatus) Set(val *DeploymentStateCompletedStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentStateCompletedStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentStateCompletedStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentStateCompletedStatus(val *DeploymentStateCompletedStatus) *NullableDeploymentStateCompletedStatus {
	return &NullableDeploymentStateCompletedStatus{value: val, isSet: true}
}

func (v NullableDeploymentStateCompletedStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentStateCompletedStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


