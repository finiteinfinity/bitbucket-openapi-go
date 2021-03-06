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

// CommitComment A commit comment.
type CommitComment struct {
	Commit *Commit `json:"commit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommitComment CommitComment

// NewCommitComment instantiates a new CommitComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommitComment() *CommitComment {
	this := CommitComment{}
	return &this
}

// NewCommitCommentWithDefaults instantiates a new CommitComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitCommentWithDefaults() *CommitComment {
	this := CommitComment{}
	return &this
}

// GetCommit returns the Commit field value if set, zero value otherwise.
func (o *CommitComment) GetCommit() Commit {
	if o == nil || o.Commit == nil {
		var ret Commit
		return ret
	}
	return *o.Commit
}

// GetCommitOk returns a tuple with the Commit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitComment) GetCommitOk() (*Commit, bool) {
	if o == nil || o.Commit == nil {
		return nil, false
	}
	return o.Commit, true
}

// HasCommit returns a boolean if a field has been set.
func (o *CommitComment) HasCommit() bool {
	if o != nil && o.Commit != nil {
		return true
	}

	return false
}

// SetCommit gets a reference to the given Commit and assigns it to the Commit field.
func (o *CommitComment) SetCommit(v Commit) {
	o.Commit = &v
}

func (o CommitComment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Commit != nil {
		toSerialize["commit"] = o.Commit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CommitComment) UnmarshalJSON(bytes []byte) (err error) {
	varCommitComment := _CommitComment{}

	if err = json.Unmarshal(bytes, &varCommitComment); err == nil {
		*o = CommitComment(varCommitComment)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "commit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommitComment struct {
	value *CommitComment
	isSet bool
}

func (v NullableCommitComment) Get() *CommitComment {
	return v.value
}

func (v *NullableCommitComment) Set(val *CommitComment) {
	v.value = val
	v.isSet = true
}

func (v NullableCommitComment) IsSet() bool {
	return v.isSet
}

func (v *NullableCommitComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommitComment(val *CommitComment) *NullableCommitComment {
	return &NullableCommitComment{value: val, isSet: true}
}

func (v NullableCommitComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommitComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


