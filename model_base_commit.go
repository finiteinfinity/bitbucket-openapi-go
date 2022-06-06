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
	"time"
)

// BaseCommit The common base type for both repository and snippet commits.
type BaseCommit struct {
	Hash *string `json:"hash,omitempty"`
	Date *time.Time `json:"date,omitempty"`
	Author *Author `json:"author,omitempty"`
	Message *string `json:"message,omitempty"`
	Summary *RenderedPullRequestMarkupTitle `json:"summary,omitempty"`
	Parents []BaseCommit `json:"parents,omitempty"`
}

// NewBaseCommit instantiates a new BaseCommit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseCommit() *BaseCommit {
	this := BaseCommit{}
	return &this
}

// NewBaseCommitWithDefaults instantiates a new BaseCommit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseCommitWithDefaults() *BaseCommit {
	this := BaseCommit{}
	return &this
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *BaseCommit) GetHash() string {
	if o == nil || o.Hash == nil {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCommit) GetHashOk() (*string, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *BaseCommit) HasHash() bool {
	if o != nil && o.Hash != nil {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *BaseCommit) SetHash(v string) {
	o.Hash = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *BaseCommit) GetDate() time.Time {
	if o == nil || o.Date == nil {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCommit) GetDateOk() (*time.Time, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *BaseCommit) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *BaseCommit) SetDate(v time.Time) {
	o.Date = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *BaseCommit) GetAuthor() Author {
	if o == nil || o.Author == nil {
		var ret Author
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCommit) GetAuthorOk() (*Author, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *BaseCommit) HasAuthor() bool {
	if o != nil && o.Author != nil {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given Author and assigns it to the Author field.
func (o *BaseCommit) SetAuthor(v Author) {
	o.Author = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *BaseCommit) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCommit) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *BaseCommit) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *BaseCommit) SetMessage(v string) {
	o.Message = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *BaseCommit) GetSummary() RenderedPullRequestMarkupTitle {
	if o == nil || o.Summary == nil {
		var ret RenderedPullRequestMarkupTitle
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCommit) GetSummaryOk() (*RenderedPullRequestMarkupTitle, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *BaseCommit) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given RenderedPullRequestMarkupTitle and assigns it to the Summary field.
func (o *BaseCommit) SetSummary(v RenderedPullRequestMarkupTitle) {
	o.Summary = &v
}

// GetParents returns the Parents field value if set, zero value otherwise.
func (o *BaseCommit) GetParents() []BaseCommit {
	if o == nil || o.Parents == nil {
		var ret []BaseCommit
		return ret
	}
	return o.Parents
}

// GetParentsOk returns a tuple with the Parents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCommit) GetParentsOk() ([]BaseCommit, bool) {
	if o == nil || o.Parents == nil {
		return nil, false
	}
	return o.Parents, true
}

// HasParents returns a boolean if a field has been set.
func (o *BaseCommit) HasParents() bool {
	if o != nil && o.Parents != nil {
		return true
	}

	return false
}

// SetParents gets a reference to the given []BaseCommit and assigns it to the Parents field.
func (o *BaseCommit) SetParents(v []BaseCommit) {
	o.Parents = v
}

func (o BaseCommit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hash != nil {
		toSerialize["hash"] = o.Hash
	}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.Author != nil {
		toSerialize["author"] = o.Author
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.Parents != nil {
		toSerialize["parents"] = o.Parents
	}
	return json.Marshal(toSerialize)
}

type NullableBaseCommit struct {
	value *BaseCommit
	isSet bool
}

func (v NullableBaseCommit) Get() *BaseCommit {
	return v.value
}

func (v *NullableBaseCommit) Set(val *BaseCommit) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseCommit) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseCommit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseCommit(val *BaseCommit) *NullableBaseCommit {
	return &NullableBaseCommit{value: val, isSet: true}
}

func (v NullableBaseCommit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseCommit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

