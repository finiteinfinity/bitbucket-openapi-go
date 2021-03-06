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

// SearchSegment struct for SearchSegment
type SearchSegment struct {
	Text *string `json:"text,omitempty"`
	Match *bool `json:"match,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SearchSegment SearchSegment

// NewSearchSegment instantiates a new SearchSegment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchSegment() *SearchSegment {
	this := SearchSegment{}
	return &this
}

// NewSearchSegmentWithDefaults instantiates a new SearchSegment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchSegmentWithDefaults() *SearchSegment {
	this := SearchSegment{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *SearchSegment) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSegment) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *SearchSegment) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *SearchSegment) SetText(v string) {
	o.Text = &v
}

// GetMatch returns the Match field value if set, zero value otherwise.
func (o *SearchSegment) GetMatch() bool {
	if o == nil || o.Match == nil {
		var ret bool
		return ret
	}
	return *o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSegment) GetMatchOk() (*bool, bool) {
	if o == nil || o.Match == nil {
		return nil, false
	}
	return o.Match, true
}

// HasMatch returns a boolean if a field has been set.
func (o *SearchSegment) HasMatch() bool {
	if o != nil && o.Match != nil {
		return true
	}

	return false
}

// SetMatch gets a reference to the given bool and assigns it to the Match field.
func (o *SearchSegment) SetMatch(v bool) {
	o.Match = &v
}

func (o SearchSegment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	if o.Match != nil {
		toSerialize["match"] = o.Match
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SearchSegment) UnmarshalJSON(bytes []byte) (err error) {
	varSearchSegment := _SearchSegment{}

	if err = json.Unmarshal(bytes, &varSearchSegment); err == nil {
		*o = SearchSegment(varSearchSegment)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "text")
		delete(additionalProperties, "match")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSearchSegment struct {
	value *SearchSegment
	isSet bool
}

func (v NullableSearchSegment) Get() *SearchSegment {
	return v.value
}

func (v *NullableSearchSegment) Set(val *SearchSegment) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchSegment) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchSegment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchSegment(val *SearchSegment) *NullableSearchSegment {
	return &NullableSearchSegment{value: val, isSet: true}
}

func (v NullableSearchSegment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchSegment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


