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

// PipelinesStgWestPipelineStep struct for PipelinesStgWestPipelineStep
type PipelinesStgWestPipelineStep struct {
	Object
}

// NewPipelinesStgWestPipelineStep instantiates a new PipelinesStgWestPipelineStep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelinesStgWestPipelineStep() *PipelinesStgWestPipelineStep {
	this := PipelinesStgWestPipelineStep{}
	return &this
}

// NewPipelinesStgWestPipelineStepWithDefaults instantiates a new PipelinesStgWestPipelineStep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelinesStgWestPipelineStepWithDefaults() *PipelinesStgWestPipelineStep {
	this := PipelinesStgWestPipelineStep{}
	return &this
}

func (o PipelinesStgWestPipelineStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedObject, errObject := json.Marshal(o.Object)
	if errObject != nil {
		return []byte{}, errObject
	}
	errObject = json.Unmarshal([]byte(serializedObject), &toSerialize)
	if errObject != nil {
		return []byte{}, errObject
	}
	return json.Marshal(toSerialize)
}

type NullablePipelinesStgWestPipelineStep struct {
	value *PipelinesStgWestPipelineStep
	isSet bool
}

func (v NullablePipelinesStgWestPipelineStep) Get() *PipelinesStgWestPipelineStep {
	return v.value
}

func (v *NullablePipelinesStgWestPipelineStep) Set(val *PipelinesStgWestPipelineStep) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelinesStgWestPipelineStep) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelinesStgWestPipelineStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelinesStgWestPipelineStep(val *PipelinesStgWestPipelineStep) *NullablePipelinesStgWestPipelineStep {
	return &NullablePipelinesStgWestPipelineStep{value: val, isSet: true}
}

func (v NullablePipelinesStgWestPipelineStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelinesStgWestPipelineStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

