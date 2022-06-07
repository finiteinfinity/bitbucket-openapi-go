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

// CommitFile A file object, representing a file at a commit in a repository
type CommitFile struct {
	Type string `json:"type"`
	// The path in the repository
	Path *string `json:"path,omitempty"`
	Commit *Commit `json:"commit,omitempty"`
	Attributes *string `json:"attributes,omitempty"`
	// The escaped version of the path as it appears in a diff. If the path does not require escaping this will be the same as path.
	EscapedPath *string `json:"escaped_path,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommitFile CommitFile

// NewCommitFile instantiates a new CommitFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommitFile() *CommitFile {
	this := CommitFile{}
	return &this
}

// NewCommitFileWithDefaults instantiates a new CommitFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitFileWithDefaults() *CommitFile {
	this := CommitFile{}
	return &this
}

// GetType returns the Type field value
func (o *CommitFile) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CommitFile) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CommitFile) SetType(v string) {
	o.Type = v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *CommitFile) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitFile) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *CommitFile) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *CommitFile) SetPath(v string) {
	o.Path = &v
}

// GetCommit returns the Commit field value if set, zero value otherwise.
func (o *CommitFile) GetCommit() Commit {
	if o == nil || o.Commit == nil {
		var ret Commit
		return ret
	}
	return *o.Commit
}

// GetCommitOk returns a tuple with the Commit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitFile) GetCommitOk() (*Commit, bool) {
	if o == nil || o.Commit == nil {
		return nil, false
	}
	return o.Commit, true
}

// HasCommit returns a boolean if a field has been set.
func (o *CommitFile) HasCommit() bool {
	if o != nil && o.Commit != nil {
		return true
	}

	return false
}

// SetCommit gets a reference to the given Commit and assigns it to the Commit field.
func (o *CommitFile) SetCommit(v Commit) {
	o.Commit = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CommitFile) GetAttributes() string {
	if o == nil || o.Attributes == nil {
		var ret string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitFile) GetAttributesOk() (*string, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CommitFile) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given string and assigns it to the Attributes field.
func (o *CommitFile) SetAttributes(v string) {
	o.Attributes = &v
}

// GetEscapedPath returns the EscapedPath field value if set, zero value otherwise.
func (o *CommitFile) GetEscapedPath() string {
	if o == nil || o.EscapedPath == nil {
		var ret string
		return ret
	}
	return *o.EscapedPath
}

// GetEscapedPathOk returns a tuple with the EscapedPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitFile) GetEscapedPathOk() (*string, bool) {
	if o == nil || o.EscapedPath == nil {
		return nil, false
	}
	return o.EscapedPath, true
}

// HasEscapedPath returns a boolean if a field has been set.
func (o *CommitFile) HasEscapedPath() bool {
	if o != nil && o.EscapedPath != nil {
		return true
	}

	return false
}

// SetEscapedPath gets a reference to the given string and assigns it to the EscapedPath field.
func (o *CommitFile) SetEscapedPath(v string) {
	o.EscapedPath = &v
}

func (o CommitFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Commit != nil {
		toSerialize["commit"] = o.Commit
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.EscapedPath != nil {
		toSerialize["escaped_path"] = o.EscapedPath
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CommitFile) UnmarshalJSON(bytes []byte) (err error) {
	varCommitFile := _CommitFile{}

	if err = json.Unmarshal(bytes, &varCommitFile); err == nil {
		*o = CommitFile(varCommitFile)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "path")
		delete(additionalProperties, "commit")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "escaped_path")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommitFile struct {
	value *CommitFile
	isSet bool
}

func (v NullableCommitFile) Get() *CommitFile {
	return v.value
}

func (v *NullableCommitFile) Set(val *CommitFile) {
	v.value = val
	v.isSet = true
}

func (v NullableCommitFile) IsSet() bool {
	return v.isSet
}

func (v *NullableCommitFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommitFile(val *CommitFile) *NullableCommitFile {
	return &NullableCommitFile{value: val, isSet: true}
}

func (v NullableCommitFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommitFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


