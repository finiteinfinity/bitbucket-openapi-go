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

// ProjectLinks struct for ProjectLinks
type ProjectLinks struct {
	Html *Link `json:"html,omitempty"`
	Avatar *Link `json:"avatar,omitempty"`
}

// NewProjectLinks instantiates a new ProjectLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectLinks() *ProjectLinks {
	this := ProjectLinks{}
	return &this
}

// NewProjectLinksWithDefaults instantiates a new ProjectLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectLinksWithDefaults() *ProjectLinks {
	this := ProjectLinks{}
	return &this
}

// GetHtml returns the Html field value if set, zero value otherwise.
func (o *ProjectLinks) GetHtml() Link {
	if o == nil || o.Html == nil {
		var ret Link
		return ret
	}
	return *o.Html
}

// GetHtmlOk returns a tuple with the Html field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectLinks) GetHtmlOk() (*Link, bool) {
	if o == nil || o.Html == nil {
		return nil, false
	}
	return o.Html, true
}

// HasHtml returns a boolean if a field has been set.
func (o *ProjectLinks) HasHtml() bool {
	if o != nil && o.Html != nil {
		return true
	}

	return false
}

// SetHtml gets a reference to the given Link and assigns it to the Html field.
func (o *ProjectLinks) SetHtml(v Link) {
	o.Html = &v
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *ProjectLinks) GetAvatar() Link {
	if o == nil || o.Avatar == nil {
		var ret Link
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectLinks) GetAvatarOk() (*Link, bool) {
	if o == nil || o.Avatar == nil {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *ProjectLinks) HasAvatar() bool {
	if o != nil && o.Avatar != nil {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given Link and assigns it to the Avatar field.
func (o *ProjectLinks) SetAvatar(v Link) {
	o.Avatar = &v
}

func (o ProjectLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Html != nil {
		toSerialize["html"] = o.Html
	}
	if o.Avatar != nil {
		toSerialize["avatar"] = o.Avatar
	}
	return json.Marshal(toSerialize)
}

type NullableProjectLinks struct {
	value *ProjectLinks
	isSet bool
}

func (v NullableProjectLinks) Get() *ProjectLinks {
	return v.value
}

func (v *NullableProjectLinks) Set(val *ProjectLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectLinks(val *ProjectLinks) *NullableProjectLinks {
	return &NullableProjectLinks{value: val, isSet: true}
}

func (v NullableProjectLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


