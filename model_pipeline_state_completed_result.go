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

// PipelineStateCompletedResult struct for PipelineStateCompletedResult
type PipelineStateCompletedResult struct {
	Object
	AdditionalProperties map[string]interface{}
}

type _PipelineStateCompletedResult PipelineStateCompletedResult

// NewPipelineStateCompletedResult instantiates a new PipelineStateCompletedResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineStateCompletedResult() *PipelineStateCompletedResult {
	this := PipelineStateCompletedResult{}
	return &this
}

// NewPipelineStateCompletedResultWithDefaults instantiates a new PipelineStateCompletedResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineStateCompletedResultWithDefaults() *PipelineStateCompletedResult {
	this := PipelineStateCompletedResult{}
	return &this
}

func (o PipelineStateCompletedResult) MarshalJSON() ([]byte, error) {
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

func (o *PipelineStateCompletedResult) UnmarshalJSON(bytes []byte) (err error) {
	type PipelineStateCompletedResultWithoutEmbeddedStruct struct {
	}

	varPipelineStateCompletedResultWithoutEmbeddedStruct := PipelineStateCompletedResultWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPipelineStateCompletedResultWithoutEmbeddedStruct)
	if err == nil {
		varPipelineStateCompletedResult := _PipelineStateCompletedResult{}
		*o = PipelineStateCompletedResult(varPipelineStateCompletedResult)
	} else {
		return err
	}

	varPipelineStateCompletedResult := _PipelineStateCompletedResult{}

	err = json.Unmarshal(bytes, &varPipelineStateCompletedResult)
	if err == nil {
		o.Object = varPipelineStateCompletedResult.Object
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

type NullablePipelineStateCompletedResult struct {
	value *PipelineStateCompletedResult
	isSet bool
}

func (v NullablePipelineStateCompletedResult) Get() *PipelineStateCompletedResult {
	return v.value
}

func (v *NullablePipelineStateCompletedResult) Set(val *PipelineStateCompletedResult) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineStateCompletedResult) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineStateCompletedResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineStateCompletedResult(val *PipelineStateCompletedResult) *NullablePipelineStateCompletedResult {
	return &NullablePipelineStateCompletedResult{value: val, isSet: true}
}

func (v NullablePipelineStateCompletedResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineStateCompletedResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


