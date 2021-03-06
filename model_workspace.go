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

// Workspace A Bitbucket workspace.             Workspaces are used to organize repositories.
type Workspace struct {
	Links *WorkspaceLinks `json:"links,omitempty"`
	// The workspace's immutable id.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the workspace.
	Name *string `json:"name,omitempty"`
	// The short label that identifies this workspace.
	Slug *string `json:"slug,omitempty"`
	// Indicates whether the workspace is publicly accessible, or whether it is private to the members and consequently only visible to members.
	IsPrivate *bool `json:"is_private,omitempty"`
	CreatedOn *time.Time `json:"created_on,omitempty"`
	UpdatedOn *time.Time `json:"updated_on,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Workspace Workspace

// NewWorkspace instantiates a new Workspace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkspace() *Workspace {
	this := Workspace{}
	return &this
}

// NewWorkspaceWithDefaults instantiates a new Workspace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkspaceWithDefaults() *Workspace {
	this := Workspace{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Workspace) GetLinks() WorkspaceLinks {
	if o == nil || o.Links == nil {
		var ret WorkspaceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workspace) GetLinksOk() (*WorkspaceLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Workspace) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given WorkspaceLinks and assigns it to the Links field.
func (o *Workspace) SetLinks(v WorkspaceLinks) {
	o.Links = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *Workspace) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workspace) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *Workspace) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *Workspace) SetUuid(v string) {
	o.Uuid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Workspace) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workspace) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Workspace) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Workspace) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *Workspace) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workspace) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *Workspace) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *Workspace) SetSlug(v string) {
	o.Slug = &v
}

// GetIsPrivate returns the IsPrivate field value if set, zero value otherwise.
func (o *Workspace) GetIsPrivate() bool {
	if o == nil || o.IsPrivate == nil {
		var ret bool
		return ret
	}
	return *o.IsPrivate
}

// GetIsPrivateOk returns a tuple with the IsPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workspace) GetIsPrivateOk() (*bool, bool) {
	if o == nil || o.IsPrivate == nil {
		return nil, false
	}
	return o.IsPrivate, true
}

// HasIsPrivate returns a boolean if a field has been set.
func (o *Workspace) HasIsPrivate() bool {
	if o != nil && o.IsPrivate != nil {
		return true
	}

	return false
}

// SetIsPrivate gets a reference to the given bool and assigns it to the IsPrivate field.
func (o *Workspace) SetIsPrivate(v bool) {
	o.IsPrivate = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *Workspace) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workspace) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *Workspace) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *Workspace) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetUpdatedOn returns the UpdatedOn field value if set, zero value otherwise.
func (o *Workspace) GetUpdatedOn() time.Time {
	if o == nil || o.UpdatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedOn
}

// GetUpdatedOnOk returns a tuple with the UpdatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workspace) GetUpdatedOnOk() (*time.Time, bool) {
	if o == nil || o.UpdatedOn == nil {
		return nil, false
	}
	return o.UpdatedOn, true
}

// HasUpdatedOn returns a boolean if a field has been set.
func (o *Workspace) HasUpdatedOn() bool {
	if o != nil && o.UpdatedOn != nil {
		return true
	}

	return false
}

// SetUpdatedOn gets a reference to the given time.Time and assigns it to the UpdatedOn field.
func (o *Workspace) SetUpdatedOn(v time.Time) {
	o.UpdatedOn = &v
}

func (o Workspace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	if o.IsPrivate != nil {
		toSerialize["is_private"] = o.IsPrivate
	}
	if o.CreatedOn != nil {
		toSerialize["created_on"] = o.CreatedOn
	}
	if o.UpdatedOn != nil {
		toSerialize["updated_on"] = o.UpdatedOn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Workspace) UnmarshalJSON(bytes []byte) (err error) {
	varWorkspace := _Workspace{}

	if err = json.Unmarshal(bytes, &varWorkspace); err == nil {
		*o = Workspace(varWorkspace)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "links")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "is_private")
		delete(additionalProperties, "created_on")
		delete(additionalProperties, "updated_on")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkspace struct {
	value *Workspace
	isSet bool
}

func (v NullableWorkspace) Get() *Workspace {
	return v.value
}

func (v *NullableWorkspace) Set(val *Workspace) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspace) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspace(val *Workspace) *NullableWorkspace {
	return &NullableWorkspace{value: val, isSet: true}
}

func (v NullableWorkspace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


