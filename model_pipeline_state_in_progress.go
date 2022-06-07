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

// PipelineStateInProgress struct for PipelineStateInProgress
type PipelineStateInProgress struct {
	PipelineState
	// The name of pipeline state (IN_PROGRESS).
	Name *string `json:"name,omitempty"`
	Stage *PipelineStateInProgressStage `json:"stage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PipelineStateInProgress PipelineStateInProgress

// NewPipelineStateInProgress instantiates a new PipelineStateInProgress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineStateInProgress() *PipelineStateInProgress {
	this := PipelineStateInProgress{}
	return &this
}

// NewPipelineStateInProgressWithDefaults instantiates a new PipelineStateInProgress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineStateInProgressWithDefaults() *PipelineStateInProgress {
	this := PipelineStateInProgress{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PipelineStateInProgress) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineStateInProgress) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PipelineStateInProgress) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PipelineStateInProgress) SetName(v string) {
	o.Name = &v
}

// GetStage returns the Stage field value if set, zero value otherwise.
func (o *PipelineStateInProgress) GetStage() PipelineStateInProgressStage {
	if o == nil || o.Stage == nil {
		var ret PipelineStateInProgressStage
		return ret
	}
	return *o.Stage
}

// GetStageOk returns a tuple with the Stage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineStateInProgress) GetStageOk() (*PipelineStateInProgressStage, bool) {
	if o == nil || o.Stage == nil {
		return nil, false
	}
	return o.Stage, true
}

// HasStage returns a boolean if a field has been set.
func (o *PipelineStateInProgress) HasStage() bool {
	if o != nil && o.Stage != nil {
		return true
	}

	return false
}

// SetStage gets a reference to the given PipelineStateInProgressStage and assigns it to the Stage field.
func (o *PipelineStateInProgress) SetStage(v PipelineStateInProgressStage) {
	o.Stage = &v
}

func (o PipelineStateInProgress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPipelineState, errPipelineState := json.Marshal(o.PipelineState)
	if errPipelineState != nil {
		return []byte{}, errPipelineState
	}
	errPipelineState = json.Unmarshal([]byte(serializedPipelineState), &toSerialize)
	if errPipelineState != nil {
		return []byte{}, errPipelineState
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Stage != nil {
		toSerialize["stage"] = o.Stage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PipelineStateInProgress) UnmarshalJSON(bytes []byte) (err error) {
	type PipelineStateInProgressWithoutEmbeddedStruct struct {
		// The name of pipeline state (IN_PROGRESS).
		Name *string `json:"name,omitempty"`
		Stage *PipelineStateInProgressStage `json:"stage,omitempty"`
	}

	varPipelineStateInProgressWithoutEmbeddedStruct := PipelineStateInProgressWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPipelineStateInProgressWithoutEmbeddedStruct)
	if err == nil {
		varPipelineStateInProgress := _PipelineStateInProgress{}
		varPipelineStateInProgress.Name = varPipelineStateInProgressWithoutEmbeddedStruct.Name
		varPipelineStateInProgress.Stage = varPipelineStateInProgressWithoutEmbeddedStruct.Stage
		*o = PipelineStateInProgress(varPipelineStateInProgress)
	} else {
		return err
	}

	varPipelineStateInProgress := _PipelineStateInProgress{}

	err = json.Unmarshal(bytes, &varPipelineStateInProgress)
	if err == nil {
		o.PipelineState = varPipelineStateInProgress.PipelineState
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "stage")

		// remove fields from embedded structs
		reflectPipelineState := reflect.ValueOf(o.PipelineState)
		for i := 0; i < reflectPipelineState.Type().NumField(); i++ {
			t := reflectPipelineState.Type().Field(i)

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

type NullablePipelineStateInProgress struct {
	value *PipelineStateInProgress
	isSet bool
}

func (v NullablePipelineStateInProgress) Get() *PipelineStateInProgress {
	return v.value
}

func (v *NullablePipelineStateInProgress) Set(val *PipelineStateInProgress) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineStateInProgress) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineStateInProgress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineStateInProgress(val *PipelineStateInProgress) *NullablePipelineStateInProgress {
	return &NullablePipelineStateInProgress{value: val, isSet: true}
}

func (v NullablePipelineStateInProgress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineStateInProgress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


