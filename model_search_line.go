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

// SearchLine struct for SearchLine
type SearchLine struct {
	Line *int32 `json:"line,omitempty"`
	Segments []SearchSegment `json:"segments,omitempty"`
}

// NewSearchLine instantiates a new SearchLine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchLine() *SearchLine {
	this := SearchLine{}
	return &this
}

// NewSearchLineWithDefaults instantiates a new SearchLine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchLineWithDefaults() *SearchLine {
	this := SearchLine{}
	return &this
}

// GetLine returns the Line field value if set, zero value otherwise.
func (o *SearchLine) GetLine() int32 {
	if o == nil || o.Line == nil {
		var ret int32
		return ret
	}
	return *o.Line
}

// GetLineOk returns a tuple with the Line field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchLine) GetLineOk() (*int32, bool) {
	if o == nil || o.Line == nil {
		return nil, false
	}
	return o.Line, true
}

// HasLine returns a boolean if a field has been set.
func (o *SearchLine) HasLine() bool {
	if o != nil && o.Line != nil {
		return true
	}

	return false
}

// SetLine gets a reference to the given int32 and assigns it to the Line field.
func (o *SearchLine) SetLine(v int32) {
	o.Line = &v
}

// GetSegments returns the Segments field value if set, zero value otherwise.
func (o *SearchLine) GetSegments() []SearchSegment {
	if o == nil || o.Segments == nil {
		var ret []SearchSegment
		return ret
	}
	return o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchLine) GetSegmentsOk() ([]SearchSegment, bool) {
	if o == nil || o.Segments == nil {
		return nil, false
	}
	return o.Segments, true
}

// HasSegments returns a boolean if a field has been set.
func (o *SearchLine) HasSegments() bool {
	if o != nil && o.Segments != nil {
		return true
	}

	return false
}

// SetSegments gets a reference to the given []SearchSegment and assigns it to the Segments field.
func (o *SearchLine) SetSegments(v []SearchSegment) {
	o.Segments = v
}

func (o SearchLine) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Line != nil {
		toSerialize["line"] = o.Line
	}
	if o.Segments != nil {
		toSerialize["segments"] = o.Segments
	}
	return json.Marshal(toSerialize)
}

type NullableSearchLine struct {
	value *SearchLine
	isSet bool
}

func (v NullableSearchLine) Get() *SearchLine {
	return v.value
}

func (v *NullableSearchLine) Set(val *SearchLine) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchLine) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchLine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchLine(val *SearchLine) *NullableSearchLine {
	return &NullableSearchLine{value: val, isSet: true}
}

func (v NullableSearchLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchLine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


