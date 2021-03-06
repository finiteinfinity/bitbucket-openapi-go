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

// StgWestReport struct for StgWestReport
type StgWestReport struct {
	Object
	AdditionalProperties map[string]interface{}
}

type _StgWestReport StgWestReport

// NewStgWestReport instantiates a new StgWestReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStgWestReport() *StgWestReport {
	this := StgWestReport{}
	return &this
}

// NewStgWestReportWithDefaults instantiates a new StgWestReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStgWestReportWithDefaults() *StgWestReport {
	this := StgWestReport{}
	return &this
}

func (o StgWestReport) MarshalJSON() ([]byte, error) {
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

func (o *StgWestReport) UnmarshalJSON(bytes []byte) (err error) {
	type StgWestReportWithoutEmbeddedStruct struct {
	}

	varStgWestReportWithoutEmbeddedStruct := StgWestReportWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStgWestReportWithoutEmbeddedStruct)
	if err == nil {
		varStgWestReport := _StgWestReport{}
		*o = StgWestReport(varStgWestReport)
	} else {
		return err
	}

	varStgWestReport := _StgWestReport{}

	err = json.Unmarshal(bytes, &varStgWestReport)
	if err == nil {
		o.Object = varStgWestReport.Object
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

type NullableStgWestReport struct {
	value *StgWestReport
	isSet bool
}

func (v NullableStgWestReport) Get() *StgWestReport {
	return v.value
}

func (v *NullableStgWestReport) Set(val *StgWestReport) {
	v.value = val
	v.isSet = true
}

func (v NullableStgWestReport) IsSet() bool {
	return v.isSet
}

func (v *NullableStgWestReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStgWestReport(val *StgWestReport) *NullableStgWestReport {
	return &NullableStgWestReport{value: val, isSet: true}
}

func (v NullableStgWestReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStgWestReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


