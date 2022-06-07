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

// PipelineStatePending struct for PipelineStatePending
type PipelineStatePending struct {
	PipelineState
	// The name of pipeline state (PENDING).
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PipelineStatePending PipelineStatePending

// NewPipelineStatePending instantiates a new PipelineStatePending object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineStatePending() *PipelineStatePending {
	this := PipelineStatePending{}
	return &this
}

// NewPipelineStatePendingWithDefaults instantiates a new PipelineStatePending object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineStatePendingWithDefaults() *PipelineStatePending {
	this := PipelineStatePending{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PipelineStatePending) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineStatePending) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PipelineStatePending) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PipelineStatePending) SetName(v string) {
	o.Name = &v
}

func (o PipelineStatePending) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PipelineStatePending) UnmarshalJSON(bytes []byte) (err error) {
	type PipelineStatePendingWithoutEmbeddedStruct struct {
		// The name of pipeline state (PENDING).
		Name *string `json:"name,omitempty"`
	}

	varPipelineStatePendingWithoutEmbeddedStruct := PipelineStatePendingWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPipelineStatePendingWithoutEmbeddedStruct)
	if err == nil {
		varPipelineStatePending := _PipelineStatePending{}
		varPipelineStatePending.Name = varPipelineStatePendingWithoutEmbeddedStruct.Name
		*o = PipelineStatePending(varPipelineStatePending)
	} else {
		return err
	}

	varPipelineStatePending := _PipelineStatePending{}

	err = json.Unmarshal(bytes, &varPipelineStatePending)
	if err == nil {
		o.PipelineState = varPipelineStatePending.PipelineState
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")

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

type NullablePipelineStatePending struct {
	value *PipelineStatePending
	isSet bool
}

func (v NullablePipelineStatePending) Get() *PipelineStatePending {
	return v.value
}

func (v *NullablePipelineStatePending) Set(val *PipelineStatePending) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineStatePending) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineStatePending) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineStatePending(val *PipelineStatePending) *NullablePipelineStatePending {
	return &NullablePipelineStatePending{value: val, isSet: true}
}

func (v NullablePipelineStatePending) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineStatePending) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


