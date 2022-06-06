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

// RepositoryUserPermission A user's direct permission for a given repository.
type RepositoryUserPermission struct {
	Type string `json:"type"`
	Permission *string `json:"permission,omitempty"`
	User *User `json:"user,omitempty"`
	Repository *Repository `json:"repository,omitempty"`
	Links *BranchingModelSettingsLinks `json:"links,omitempty"`
}

// NewRepositoryUserPermission instantiates a new RepositoryUserPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryUserPermission() *RepositoryUserPermission {
	this := RepositoryUserPermission{}
	return &this
}

// NewRepositoryUserPermissionWithDefaults instantiates a new RepositoryUserPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryUserPermissionWithDefaults() *RepositoryUserPermission {
	this := RepositoryUserPermission{}
	return &this
}

// GetType returns the Type field value
func (o *RepositoryUserPermission) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RepositoryUserPermission) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RepositoryUserPermission) SetType(v string) {
	o.Type = v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *RepositoryUserPermission) GetPermission() string {
	if o == nil || o.Permission == nil {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryUserPermission) GetPermissionOk() (*string, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *RepositoryUserPermission) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *RepositoryUserPermission) SetPermission(v string) {
	o.Permission = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *RepositoryUserPermission) GetUser() User {
	if o == nil || o.User == nil {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryUserPermission) GetUserOk() (*User, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *RepositoryUserPermission) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *RepositoryUserPermission) SetUser(v User) {
	o.User = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *RepositoryUserPermission) GetRepository() Repository {
	if o == nil || o.Repository == nil {
		var ret Repository
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryUserPermission) GetRepositoryOk() (*Repository, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *RepositoryUserPermission) HasRepository() bool {
	if o != nil && o.Repository != nil {
		return true
	}

	return false
}

// SetRepository gets a reference to the given Repository and assigns it to the Repository field.
func (o *RepositoryUserPermission) SetRepository(v Repository) {
	o.Repository = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RepositoryUserPermission) GetLinks() BranchingModelSettingsLinks {
	if o == nil || o.Links == nil {
		var ret BranchingModelSettingsLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryUserPermission) GetLinksOk() (*BranchingModelSettingsLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RepositoryUserPermission) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given BranchingModelSettingsLinks and assigns it to the Links field.
func (o *RepositoryUserPermission) SetLinks(v BranchingModelSettingsLinks) {
	o.Links = &v
}

func (o RepositoryUserPermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Permission != nil {
		toSerialize["permission"] = o.Permission
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Repository != nil {
		toSerialize["repository"] = o.Repository
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryUserPermission struct {
	value *RepositoryUserPermission
	isSet bool
}

func (v NullableRepositoryUserPermission) Get() *RepositoryUserPermission {
	return v.value
}

func (v *NullableRepositoryUserPermission) Set(val *RepositoryUserPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryUserPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryUserPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryUserPermission(val *RepositoryUserPermission) *NullableRepositoryUserPermission {
	return &NullableRepositoryUserPermission{value: val, isSet: true}
}

func (v NullableRepositoryUserPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryUserPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

