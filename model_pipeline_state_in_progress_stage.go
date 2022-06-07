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

// PipelineStateInProgressStage struct for PipelineStateInProgressStage
type PipelineStateInProgressStage struct {
	Object
	AdditionalProperties map[string]interface{}
}

type _PipelineStateInProgressStage PipelineStateInProgressStage

// NewPipelineStateInProgressStage instantiates a new PipelineStateInProgressStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineStateInProgressStage() *PipelineStateInProgressStage {
	this := PipelineStateInProgressStage{}
	return &this
}

// NewPipelineStateInProgressStageWithDefaults instantiates a new PipelineStateInProgressStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineStateInProgressStageWithDefaults() *PipelineStateInProgressStage {
	this := PipelineStateInProgressStage{}
	return &this
}

func (o PipelineStateInProgressStage) MarshalJSON() ([]byte, error) {
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

func (o *PipelineStateInProgressStage) UnmarshalJSON(bytes []byte) (err error) {
	type PipelineStateInProgressStageWithoutEmbeddedStruct struct {
	}

	varPipelineStateInProgressStageWithoutEmbeddedStruct := PipelineStateInProgressStageWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPipelineStateInProgressStageWithoutEmbeddedStruct)
	if err == nil {
		varPipelineStateInProgressStage := _PipelineStateInProgressStage{}
		*o = PipelineStateInProgressStage(varPipelineStateInProgressStage)
	} else {
		return err
	}

	varPipelineStateInProgressStage := _PipelineStateInProgressStage{}

	err = json.Unmarshal(bytes, &varPipelineStateInProgressStage)
	if err == nil {
		o.Object = varPipelineStateInProgressStage.Object
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

type NullablePipelineStateInProgressStage struct {
	value *PipelineStateInProgressStage
	isSet bool
}

func (v NullablePipelineStateInProgressStage) Get() *PipelineStateInProgressStage {
	return v.value
}

func (v *NullablePipelineStateInProgressStage) Set(val *PipelineStateInProgressStage) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineStateInProgressStage) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineStateInProgressStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineStateInProgressStage(val *PipelineStateInProgressStage) *NullablePipelineStateInProgressStage {
	return &NullablePipelineStateInProgressStage{value: val, isSet: true}
}

func (v NullablePipelineStateInProgressStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineStateInProgressStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


