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

// SnippetCommit 
type SnippetCommit struct {
	Links *SnippetCommitLinks `json:"links,omitempty"`
	Snippet *Snippet `json:"snippet,omitempty"`
}

// NewSnippetCommit instantiates a new SnippetCommit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnippetCommit() *SnippetCommit {
	this := SnippetCommit{}
	return &this
}

// NewSnippetCommitWithDefaults instantiates a new SnippetCommit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnippetCommitWithDefaults() *SnippetCommit {
	this := SnippetCommit{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SnippetCommit) GetLinks() SnippetCommitLinks {
	if o == nil || o.Links == nil {
		var ret SnippetCommitLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnippetCommit) GetLinksOk() (*SnippetCommitLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SnippetCommit) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SnippetCommitLinks and assigns it to the Links field.
func (o *SnippetCommit) SetLinks(v SnippetCommitLinks) {
	o.Links = &v
}

// GetSnippet returns the Snippet field value if set, zero value otherwise.
func (o *SnippetCommit) GetSnippet() Snippet {
	if o == nil || o.Snippet == nil {
		var ret Snippet
		return ret
	}
	return *o.Snippet
}

// GetSnippetOk returns a tuple with the Snippet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnippetCommit) GetSnippetOk() (*Snippet, bool) {
	if o == nil || o.Snippet == nil {
		return nil, false
	}
	return o.Snippet, true
}

// HasSnippet returns a boolean if a field has been set.
func (o *SnippetCommit) HasSnippet() bool {
	if o != nil && o.Snippet != nil {
		return true
	}

	return false
}

// SetSnippet gets a reference to the given Snippet and assigns it to the Snippet field.
func (o *SnippetCommit) SetSnippet(v Snippet) {
	o.Snippet = &v
}

func (o SnippetCommit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Snippet != nil {
		toSerialize["snippet"] = o.Snippet
	}
	return json.Marshal(toSerialize)
}

type NullableSnippetCommit struct {
	value *SnippetCommit
	isSet bool
}

func (v NullableSnippetCommit) Get() *SnippetCommit {
	return v.value
}

func (v *NullableSnippetCommit) Set(val *SnippetCommit) {
	v.value = val
	v.isSet = true
}

func (v NullableSnippetCommit) IsSet() bool {
	return v.isSet
}

func (v *NullableSnippetCommit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnippetCommit(val *SnippetCommit) *NullableSnippetCommit {
	return &NullableSnippetCommit{value: val, isSet: true}
}

func (v NullableSnippetCommit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnippetCommit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


