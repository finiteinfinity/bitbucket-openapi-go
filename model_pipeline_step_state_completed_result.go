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

// PipelineStepStateCompletedResult struct for PipelineStepStateCompletedResult
type PipelineStepStateCompletedResult struct {
	Object
	AdditionalProperties map[string]interface{}
}

type _PipelineStepStateCompletedResult PipelineStepStateCompletedResult

// NewPipelineStepStateCompletedResult instantiates a new PipelineStepStateCompletedResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineStepStateCompletedResult() *PipelineStepStateCompletedResult {
	this := PipelineStepStateCompletedResult{}
	return &this
}

// NewPipelineStepStateCompletedResultWithDefaults instantiates a new PipelineStepStateCompletedResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineStepStateCompletedResultWithDefaults() *PipelineStepStateCompletedResult {
	this := PipelineStepStateCompletedResult{}
	return &this
}

func (o PipelineStepStateCompletedResult) MarshalJSON() ([]byte, error) {
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

func (o *PipelineStepStateCompletedResult) UnmarshalJSON(bytes []byte) (err error) {
	type PipelineStepStateCompletedResultWithoutEmbeddedStruct struct {
	}

	varPipelineStepStateCompletedResultWithoutEmbeddedStruct := PipelineStepStateCompletedResultWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPipelineStepStateCompletedResultWithoutEmbeddedStruct)
	if err == nil {
		varPipelineStepStateCompletedResult := _PipelineStepStateCompletedResult{}
		*o = PipelineStepStateCompletedResult(varPipelineStepStateCompletedResult)
	} else {
		return err
	}

	varPipelineStepStateCompletedResult := _PipelineStepStateCompletedResult{}

	err = json.Unmarshal(bytes, &varPipelineStepStateCompletedResult)
	if err == nil {
		o.Object = varPipelineStepStateCompletedResult.Object
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

type NullablePipelineStepStateCompletedResult struct {
	value *PipelineStepStateCompletedResult
	isSet bool
}

func (v NullablePipelineStepStateCompletedResult) Get() *PipelineStepStateCompletedResult {
	return v.value
}

func (v *NullablePipelineStepStateCompletedResult) Set(val *PipelineStepStateCompletedResult) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineStepStateCompletedResult) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineStepStateCompletedResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineStepStateCompletedResult(val *PipelineStepStateCompletedResult) *NullablePipelineStepStateCompletedResult {
	return &NullablePipelineStepStateCompletedResult{value: val, isSet: true}
}

func (v NullablePipelineStepStateCompletedResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineStepStateCompletedResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


