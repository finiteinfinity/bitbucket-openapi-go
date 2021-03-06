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

// PipelineScheduleExecutionExecuted struct for PipelineScheduleExecutionExecuted
type PipelineScheduleExecutionExecuted struct {
	PipelineScheduleExecution
	Pipeline *Pipeline `json:"pipeline,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PipelineScheduleExecutionExecuted PipelineScheduleExecutionExecuted

// NewPipelineScheduleExecutionExecuted instantiates a new PipelineScheduleExecutionExecuted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineScheduleExecutionExecuted() *PipelineScheduleExecutionExecuted {
	this := PipelineScheduleExecutionExecuted{}
	return &this
}

// NewPipelineScheduleExecutionExecutedWithDefaults instantiates a new PipelineScheduleExecutionExecuted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineScheduleExecutionExecutedWithDefaults() *PipelineScheduleExecutionExecuted {
	this := PipelineScheduleExecutionExecuted{}
	return &this
}

// GetPipeline returns the Pipeline field value if set, zero value otherwise.
func (o *PipelineScheduleExecutionExecuted) GetPipeline() Pipeline {
	if o == nil || o.Pipeline == nil {
		var ret Pipeline
		return ret
	}
	return *o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineScheduleExecutionExecuted) GetPipelineOk() (*Pipeline, bool) {
	if o == nil || o.Pipeline == nil {
		return nil, false
	}
	return o.Pipeline, true
}

// HasPipeline returns a boolean if a field has been set.
func (o *PipelineScheduleExecutionExecuted) HasPipeline() bool {
	if o != nil && o.Pipeline != nil {
		return true
	}

	return false
}

// SetPipeline gets a reference to the given Pipeline and assigns it to the Pipeline field.
func (o *PipelineScheduleExecutionExecuted) SetPipeline(v Pipeline) {
	o.Pipeline = &v
}

func (o PipelineScheduleExecutionExecuted) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPipelineScheduleExecution, errPipelineScheduleExecution := json.Marshal(o.PipelineScheduleExecution)
	if errPipelineScheduleExecution != nil {
		return []byte{}, errPipelineScheduleExecution
	}
	errPipelineScheduleExecution = json.Unmarshal([]byte(serializedPipelineScheduleExecution), &toSerialize)
	if errPipelineScheduleExecution != nil {
		return []byte{}, errPipelineScheduleExecution
	}
	if o.Pipeline != nil {
		toSerialize["pipeline"] = o.Pipeline
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PipelineScheduleExecutionExecuted) UnmarshalJSON(bytes []byte) (err error) {
	type PipelineScheduleExecutionExecutedWithoutEmbeddedStruct struct {
		Pipeline *Pipeline `json:"pipeline,omitempty"`
	}

	varPipelineScheduleExecutionExecutedWithoutEmbeddedStruct := PipelineScheduleExecutionExecutedWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPipelineScheduleExecutionExecutedWithoutEmbeddedStruct)
	if err == nil {
		varPipelineScheduleExecutionExecuted := _PipelineScheduleExecutionExecuted{}
		varPipelineScheduleExecutionExecuted.Pipeline = varPipelineScheduleExecutionExecutedWithoutEmbeddedStruct.Pipeline
		*o = PipelineScheduleExecutionExecuted(varPipelineScheduleExecutionExecuted)
	} else {
		return err
	}

	varPipelineScheduleExecutionExecuted := _PipelineScheduleExecutionExecuted{}

	err = json.Unmarshal(bytes, &varPipelineScheduleExecutionExecuted)
	if err == nil {
		o.PipelineScheduleExecution = varPipelineScheduleExecutionExecuted.PipelineScheduleExecution
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "pipeline")

		// remove fields from embedded structs
		reflectPipelineScheduleExecution := reflect.ValueOf(o.PipelineScheduleExecution)
		for i := 0; i < reflectPipelineScheduleExecution.Type().NumField(); i++ {
			t := reflectPipelineScheduleExecution.Type().Field(i)

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

type NullablePipelineScheduleExecutionExecuted struct {
	value *PipelineScheduleExecutionExecuted
	isSet bool
}

func (v NullablePipelineScheduleExecutionExecuted) Get() *PipelineScheduleExecutionExecuted {
	return v.value
}

func (v *NullablePipelineScheduleExecutionExecuted) Set(val *PipelineScheduleExecutionExecuted) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineScheduleExecutionExecuted) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineScheduleExecutionExecuted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineScheduleExecutionExecuted(val *PipelineScheduleExecutionExecuted) *NullablePipelineScheduleExecutionExecuted {
	return &NullablePipelineScheduleExecutionExecuted{value: val, isSet: true}
}

func (v NullablePipelineScheduleExecutionExecuted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineScheduleExecutionExecuted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


