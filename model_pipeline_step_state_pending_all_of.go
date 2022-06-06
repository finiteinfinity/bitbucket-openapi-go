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

// PipelineStepStatePendingAllOf A Bitbucket Pipelines PENDING pipeline step state.
type PipelineStepStatePendingAllOf struct {
	// The name of pipeline step state (PENDING).
	Name *string `json:"name,omitempty"`
}

// NewPipelineStepStatePendingAllOf instantiates a new PipelineStepStatePendingAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineStepStatePendingAllOf() *PipelineStepStatePendingAllOf {
	this := PipelineStepStatePendingAllOf{}
	return &this
}

// NewPipelineStepStatePendingAllOfWithDefaults instantiates a new PipelineStepStatePendingAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineStepStatePendingAllOfWithDefaults() *PipelineStepStatePendingAllOf {
	this := PipelineStepStatePendingAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PipelineStepStatePendingAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineStepStatePendingAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PipelineStepStatePendingAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PipelineStepStatePendingAllOf) SetName(v string) {
	o.Name = &v
}

func (o PipelineStepStatePendingAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineStepStatePendingAllOf struct {
	value *PipelineStepStatePendingAllOf
	isSet bool
}

func (v NullablePipelineStepStatePendingAllOf) Get() *PipelineStepStatePendingAllOf {
	return v.value
}

func (v *NullablePipelineStepStatePendingAllOf) Set(val *PipelineStepStatePendingAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineStepStatePendingAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineStepStatePendingAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineStepStatePendingAllOf(val *PipelineStepStatePendingAllOf) *NullablePipelineStepStatePendingAllOf {
	return &NullablePipelineStepStatePendingAllOf{value: val, isSet: true}
}

func (v NullablePipelineStepStatePendingAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineStepStatePendingAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


