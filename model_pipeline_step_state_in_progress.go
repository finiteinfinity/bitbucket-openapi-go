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

// PipelineStepStateInProgress struct for PipelineStepStateInProgress
type PipelineStepStateInProgress struct {
	PipelineStepState
	// The name of pipeline step state (IN_PROGRESS).
	Name *string `json:"name,omitempty"`
}

// NewPipelineStepStateInProgress instantiates a new PipelineStepStateInProgress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineStepStateInProgress() *PipelineStepStateInProgress {
	this := PipelineStepStateInProgress{}
	return &this
}

// NewPipelineStepStateInProgressWithDefaults instantiates a new PipelineStepStateInProgress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineStepStateInProgressWithDefaults() *PipelineStepStateInProgress {
	this := PipelineStepStateInProgress{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PipelineStepStateInProgress) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineStepStateInProgress) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PipelineStepStateInProgress) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PipelineStepStateInProgress) SetName(v string) {
	o.Name = &v
}

func (o PipelineStepStateInProgress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPipelineStepState, errPipelineStepState := json.Marshal(o.PipelineStepState)
	if errPipelineStepState != nil {
		return []byte{}, errPipelineStepState
	}
	errPipelineStepState = json.Unmarshal([]byte(serializedPipelineStepState), &toSerialize)
	if errPipelineStepState != nil {
		return []byte{}, errPipelineStepState
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineStepStateInProgress struct {
	value *PipelineStepStateInProgress
	isSet bool
}

func (v NullablePipelineStepStateInProgress) Get() *PipelineStepStateInProgress {
	return v.value
}

func (v *NullablePipelineStepStateInProgress) Set(val *PipelineStepStateInProgress) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineStepStateInProgress) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineStepStateInProgress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineStepStateInProgress(val *PipelineStepStateInProgress) *NullablePipelineStepStateInProgress {
	return &NullablePipelineStepStateInProgress{value: val, isSet: true}
}

func (v NullablePipelineStepStateInProgress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineStepStateInProgress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


