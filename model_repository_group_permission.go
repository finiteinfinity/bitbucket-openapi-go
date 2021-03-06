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

// RepositoryGroupPermission A group's permission for a given repository.
type RepositoryGroupPermission struct {
	Type string `json:"type"`
	Links *BranchingModelSettingsLinks `json:"links,omitempty"`
	Permission *string `json:"permission,omitempty"`
	Group *Group `json:"group,omitempty"`
	Repository *Repository `json:"repository,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RepositoryGroupPermission RepositoryGroupPermission

// NewRepositoryGroupPermission instantiates a new RepositoryGroupPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryGroupPermission() *RepositoryGroupPermission {
	this := RepositoryGroupPermission{}
	return &this
}

// NewRepositoryGroupPermissionWithDefaults instantiates a new RepositoryGroupPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryGroupPermissionWithDefaults() *RepositoryGroupPermission {
	this := RepositoryGroupPermission{}
	return &this
}

// GetType returns the Type field value
func (o *RepositoryGroupPermission) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RepositoryGroupPermission) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RepositoryGroupPermission) SetType(v string) {
	o.Type = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RepositoryGroupPermission) GetLinks() BranchingModelSettingsLinks {
	if o == nil || o.Links == nil {
		var ret BranchingModelSettingsLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryGroupPermission) GetLinksOk() (*BranchingModelSettingsLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RepositoryGroupPermission) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given BranchingModelSettingsLinks and assigns it to the Links field.
func (o *RepositoryGroupPermission) SetLinks(v BranchingModelSettingsLinks) {
	o.Links = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *RepositoryGroupPermission) GetPermission() string {
	if o == nil || o.Permission == nil {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryGroupPermission) GetPermissionOk() (*string, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *RepositoryGroupPermission) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *RepositoryGroupPermission) SetPermission(v string) {
	o.Permission = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *RepositoryGroupPermission) GetGroup() Group {
	if o == nil || o.Group == nil {
		var ret Group
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryGroupPermission) GetGroupOk() (*Group, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *RepositoryGroupPermission) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given Group and assigns it to the Group field.
func (o *RepositoryGroupPermission) SetGroup(v Group) {
	o.Group = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *RepositoryGroupPermission) GetRepository() Repository {
	if o == nil || o.Repository == nil {
		var ret Repository
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryGroupPermission) GetRepositoryOk() (*Repository, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *RepositoryGroupPermission) HasRepository() bool {
	if o != nil && o.Repository != nil {
		return true
	}

	return false
}

// SetRepository gets a reference to the given Repository and assigns it to the Repository field.
func (o *RepositoryGroupPermission) SetRepository(v Repository) {
	o.Repository = &v
}

func (o RepositoryGroupPermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Permission != nil {
		toSerialize["permission"] = o.Permission
	}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.Repository != nil {
		toSerialize["repository"] = o.Repository
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RepositoryGroupPermission) UnmarshalJSON(bytes []byte) (err error) {
	varRepositoryGroupPermission := _RepositoryGroupPermission{}

	if err = json.Unmarshal(bytes, &varRepositoryGroupPermission); err == nil {
		*o = RepositoryGroupPermission(varRepositoryGroupPermission)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "links")
		delete(additionalProperties, "permission")
		delete(additionalProperties, "group")
		delete(additionalProperties, "repository")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRepositoryGroupPermission struct {
	value *RepositoryGroupPermission
	isSet bool
}

func (v NullableRepositoryGroupPermission) Get() *RepositoryGroupPermission {
	return v.value
}

func (v *NullableRepositoryGroupPermission) Set(val *RepositoryGroupPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryGroupPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryGroupPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryGroupPermission(val *RepositoryGroupPermission) *NullableRepositoryGroupPermission {
	return &NullableRepositoryGroupPermission{value: val, isSet: true}
}

func (v NullableRepositoryGroupPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryGroupPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


