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

// PipelineScheduleExecutionErroredAllOf A Pipelines schedule execution that failed to be executed.
type PipelineScheduleExecutionErroredAllOf struct {
	Error *PipelineError `json:"error,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PipelineScheduleExecutionErroredAllOf PipelineScheduleExecutionErroredAllOf

// NewPipelineScheduleExecutionErroredAllOf instantiates a new PipelineScheduleExecutionErroredAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineScheduleExecutionErroredAllOf() *PipelineScheduleExecutionErroredAllOf {
	this := PipelineScheduleExecutionErroredAllOf{}
	return &this
}

// NewPipelineScheduleExecutionErroredAllOfWithDefaults instantiates a new PipelineScheduleExecutionErroredAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineScheduleExecutionErroredAllOfWithDefaults() *PipelineScheduleExecutionErroredAllOf {
	this := PipelineScheduleExecutionErroredAllOf{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *PipelineScheduleExecutionErroredAllOf) GetError() PipelineError {
	if o == nil || o.Error == nil {
		var ret PipelineError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineScheduleExecutionErroredAllOf) GetErrorOk() (*PipelineError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *PipelineScheduleExecutionErroredAllOf) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given PipelineError and assigns it to the Error field.
func (o *PipelineScheduleExecutionErroredAllOf) SetError(v PipelineError) {
	o.Error = &v
}

func (o PipelineScheduleExecutionErroredAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PipelineScheduleExecutionErroredAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPipelineScheduleExecutionErroredAllOf := _PipelineScheduleExecutionErroredAllOf{}

	if err = json.Unmarshal(bytes, &varPipelineScheduleExecutionErroredAllOf); err == nil {
		*o = PipelineScheduleExecutionErroredAllOf(varPipelineScheduleExecutionErroredAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePipelineScheduleExecutionErroredAllOf struct {
	value *PipelineScheduleExecutionErroredAllOf
	isSet bool
}

func (v NullablePipelineScheduleExecutionErroredAllOf) Get() *PipelineScheduleExecutionErroredAllOf {
	return v.value
}

func (v *NullablePipelineScheduleExecutionErroredAllOf) Set(val *PipelineScheduleExecutionErroredAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineScheduleExecutionErroredAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineScheduleExecutionErroredAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineScheduleExecutionErroredAllOf(val *PipelineScheduleExecutionErroredAllOf) *NullablePipelineScheduleExecutionErroredAllOf {
	return &NullablePipelineScheduleExecutionErroredAllOf{value: val, isSet: true}
}

func (v NullablePipelineScheduleExecutionErroredAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineScheduleExecutionErroredAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


