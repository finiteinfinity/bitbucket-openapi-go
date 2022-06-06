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

// Object Base type for most resource objects. It defines the common `type` element that identifies an object's type. It also identifies the element as Swagger's `discriminator`.
type Object struct {
	Type string `json:"type"`
}

// NewObject instantiates a new Object object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObject() *Object {
	this := Object{}
	return &this
}

// NewObjectWithDefaults instantiates a new Object object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectWithDefaults() *Object {
	this := Object{}
	return &this
}

// GetType returns the Type field value
func (o *Object) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Object) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Object) SetType(v string) {
	o.Type = v
}

func (o Object) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableObject struct {
	value *Object
	isSet bool
}

func (v NullableObject) Get() *Object {
	return v.value
}

func (v *NullableObject) Set(val *Object) {
	v.value = val
	v.isSet = true
}

func (v NullableObject) IsSet() bool {
	return v.isSet
}

func (v *NullableObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObject(val *Object) *NullableObject {
	return &NullableObject{value: val, isSet: true}
}

func (v NullableObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


