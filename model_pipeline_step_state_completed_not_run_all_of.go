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

// PipelineStepStateCompletedNotRunAllOf A Bitbucket Pipelines NOT_RUN pipeline step result.
type PipelineStepStateCompletedNotRunAllOf struct {
	// The name of the result (NOT_RUN)
	Name *string `json:"name,omitempty"`
}

// NewPipelineStepStateCompletedNotRunAllOf instantiates a new PipelineStepStateCompletedNotRunAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineStepStateCompletedNotRunAllOf() *PipelineStepStateCompletedNotRunAllOf {
	this := PipelineStepStateCompletedNotRunAllOf{}
	return &this
}

// NewPipelineStepStateCompletedNotRunAllOfWithDefaults instantiates a new PipelineStepStateCompletedNotRunAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineStepStateCompletedNotRunAllOfWithDefaults() *PipelineStepStateCompletedNotRunAllOf {
	this := PipelineStepStateCompletedNotRunAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PipelineStepStateCompletedNotRunAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineStepStateCompletedNotRunAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PipelineStepStateCompletedNotRunAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PipelineStepStateCompletedNotRunAllOf) SetName(v string) {
	o.Name = &v
}

func (o PipelineStepStateCompletedNotRunAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineStepStateCompletedNotRunAllOf struct {
	value *PipelineStepStateCompletedNotRunAllOf
	isSet bool
}

func (v NullablePipelineStepStateCompletedNotRunAllOf) Get() *PipelineStepStateCompletedNotRunAllOf {
	return v.value
}

func (v *NullablePipelineStepStateCompletedNotRunAllOf) Set(val *PipelineStepStateCompletedNotRunAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineStepStateCompletedNotRunAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineStepStateCompletedNotRunAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineStepStateCompletedNotRunAllOf(val *PipelineStepStateCompletedNotRunAllOf) *NullablePipelineStepStateCompletedNotRunAllOf {
	return &NullablePipelineStepStateCompletedNotRunAllOf{value: val, isSet: true}
}

func (v NullablePipelineStepStateCompletedNotRunAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineStepStateCompletedNotRunAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


