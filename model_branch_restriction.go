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

// BranchRestriction A branch restriction rule.
type BranchRestriction struct {
	Links *BranchingModelSettingsLinks `json:"links,omitempty"`
	// The branch restriction status' id.
	Id *int32 `json:"id,omitempty"`
	// The type of restriction that is being applied.
	Kind string `json:"kind"`
	// Indicates how the restriction is matched against a branch. The default is `glob`.
	BranchMatchKind string `json:"branch_match_kind"`
	// Apply the restriction to branches of this type. Active when `branch_match_kind` is `branching_model`. The branch type will be calculated using the branching model configured for the repository.
	BranchType *string `json:"branch_type,omitempty"`
	// Apply the restriction to branches that match this pattern. Active when `branch_match_kind` is `glob`. Will be empty when `branch_match_kind` is `branching_model`.
	Pattern string `json:"pattern"`
	Users []Account `json:"users,omitempty"`
	Groups []Group `json:"groups,omitempty"`
	// <staticmethod object at 0x7f9f8727b5d0>
	Value *int32 `json:"value,omitempty"`
}

// NewBranchRestriction instantiates a new BranchRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBranchRestriction(kind string, branchMatchKind string, pattern string) *BranchRestriction {
	this := BranchRestriction{}
	this.Kind = kind
	this.BranchMatchKind = branchMatchKind
	this.Pattern = pattern
	return &this
}

// NewBranchRestrictionWithDefaults instantiates a new BranchRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBranchRestrictionWithDefaults() *BranchRestriction {
	this := BranchRestriction{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BranchRestriction) GetLinks() BranchingModelSettingsLinks {
	if o == nil || o.Links == nil {
		var ret BranchingModelSettingsLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchRestriction) GetLinksOk() (*BranchingModelSettingsLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BranchRestriction) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given BranchingModelSettingsLinks and assigns it to the Links field.
func (o *BranchRestriction) SetLinks(v BranchingModelSettingsLinks) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BranchRestriction) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchRestriction) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BranchRestriction) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *BranchRestriction) SetId(v int32) {
	o.Id = &v
}

// GetKind returns the Kind field value
func (o *BranchRestriction) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *BranchRestriction) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *BranchRestriction) SetKind(v string) {
	o.Kind = v
}

// GetBranchMatchKind returns the BranchMatchKind field value
func (o *BranchRestriction) GetBranchMatchKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BranchMatchKind
}

// GetBranchMatchKindOk returns a tuple with the BranchMatchKind field value
// and a boolean to check if the value has been set.
func (o *BranchRestriction) GetBranchMatchKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BranchMatchKind, true
}

// SetBranchMatchKind sets field value
func (o *BranchRestriction) SetBranchMatchKind(v string) {
	o.BranchMatchKind = v
}

// GetBranchType returns the BranchType field value if set, zero value otherwise.
func (o *BranchRestriction) GetBranchType() string {
	if o == nil || o.BranchType == nil {
		var ret string
		return ret
	}
	return *o.BranchType
}

// GetBranchTypeOk returns a tuple with the BranchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchRestriction) GetBranchTypeOk() (*string, bool) {
	if o == nil || o.BranchType == nil {
		return nil, false
	}
	return o.BranchType, true
}

// HasBranchType returns a boolean if a field has been set.
func (o *BranchRestriction) HasBranchType() bool {
	if o != nil && o.BranchType != nil {
		return true
	}

	return false
}

// SetBranchType gets a reference to the given string and assigns it to the BranchType field.
func (o *BranchRestriction) SetBranchType(v string) {
	o.BranchType = &v
}

// GetPattern returns the Pattern field value
func (o *BranchRestriction) GetPattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value
// and a boolean to check if the value has been set.
func (o *BranchRestriction) GetPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pattern, true
}

// SetPattern sets field value
func (o *BranchRestriction) SetPattern(v string) {
	o.Pattern = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *BranchRestriction) GetUsers() []Account {
	if o == nil || o.Users == nil {
		var ret []Account
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchRestriction) GetUsersOk() ([]Account, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *BranchRestriction) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []Account and assigns it to the Users field.
func (o *BranchRestriction) SetUsers(v []Account) {
	o.Users = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *BranchRestriction) GetGroups() []Group {
	if o == nil || o.Groups == nil {
		var ret []Group
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchRestriction) GetGroupsOk() ([]Group, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *BranchRestriction) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []Group and assigns it to the Groups field.
func (o *BranchRestriction) SetGroups(v []Group) {
	o.Groups = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BranchRestriction) GetValue() int32 {
	if o == nil || o.Value == nil {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchRestriction) GetValueOk() (*int32, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BranchRestriction) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *BranchRestriction) SetValue(v int32) {
	o.Value = &v
}

func (o BranchRestriction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["branch_match_kind"] = o.BranchMatchKind
	}
	if o.BranchType != nil {
		toSerialize["branch_type"] = o.BranchType
	}
	if true {
		toSerialize["pattern"] = o.Pattern
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBranchRestriction struct {
	value *BranchRestriction
	isSet bool
}

func (v NullableBranchRestriction) Get() *BranchRestriction {
	return v.value
}

func (v *NullableBranchRestriction) Set(val *BranchRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableBranchRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableBranchRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBranchRestriction(val *BranchRestriction) *NullableBranchRestriction {
	return &NullableBranchRestriction{value: val, isSet: true}
}

func (v NullableBranchRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBranchRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

