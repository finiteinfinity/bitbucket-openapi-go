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

// PullRequestCommit struct for PullRequestCommit
type PullRequestCommit struct {
	Hash *string `json:"hash,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PullRequestCommit PullRequestCommit

// NewPullRequestCommit instantiates a new PullRequestCommit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPullRequestCommit() *PullRequestCommit {
	this := PullRequestCommit{}
	return &this
}

// NewPullRequestCommitWithDefaults instantiates a new PullRequestCommit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPullRequestCommitWithDefaults() *PullRequestCommit {
	this := PullRequestCommit{}
	return &this
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *PullRequestCommit) GetHash() string {
	if o == nil || o.Hash == nil {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequestCommit) GetHashOk() (*string, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *PullRequestCommit) HasHash() bool {
	if o != nil && o.Hash != nil {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *PullRequestCommit) SetHash(v string) {
	o.Hash = &v
}

func (o PullRequestCommit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hash != nil {
		toSerialize["hash"] = o.Hash
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PullRequestCommit) UnmarshalJSON(bytes []byte) (err error) {
	varPullRequestCommit := _PullRequestCommit{}

	if err = json.Unmarshal(bytes, &varPullRequestCommit); err == nil {
		*o = PullRequestCommit(varPullRequestCommit)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "hash")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePullRequestCommit struct {
	value *PullRequestCommit
	isSet bool
}

func (v NullablePullRequestCommit) Get() *PullRequestCommit {
	return v.value
}

func (v *NullablePullRequestCommit) Set(val *PullRequestCommit) {
	v.value = val
	v.isSet = true
}

func (v NullablePullRequestCommit) IsSet() bool {
	return v.isSet
}

func (v *NullablePullRequestCommit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePullRequestCommit(val *PullRequestCommit) *NullablePullRequestCommit {
	return &NullablePullRequestCommit{value: val, isSet: true}
}

func (v NullablePullRequestCommit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePullRequestCommit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


