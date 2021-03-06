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

// Group A group object
type Group struct {
	Links *GroupLinks `json:"links,omitempty"`
	Owner *Account `json:"owner,omitempty"`
	Workspace *Workspace `json:"workspace,omitempty"`
	Name *string `json:"name,omitempty"`
	// The \"sluggified\" version of the group's name. This contains only ASCII characters and can therefore be slightly different than the name
	Slug *string `json:"slug,omitempty"`
	// The concatenation of the workspace's slug and the group's slug, separated with a colon (e.g. `acme:developers`) 
	FullSlug *string `json:"full_slug,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Group Group

// NewGroup instantiates a new Group object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroup() *Group {
	this := Group{}
	return &this
}

// NewGroupWithDefaults instantiates a new Group object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupWithDefaults() *Group {
	this := Group{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Group) GetLinks() GroupLinks {
	if o == nil || o.Links == nil {
		var ret GroupLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetLinksOk() (*GroupLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Group) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GroupLinks and assigns it to the Links field.
func (o *Group) SetLinks(v GroupLinks) {
	o.Links = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *Group) GetOwner() Account {
	if o == nil || o.Owner == nil {
		var ret Account
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetOwnerOk() (*Account, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *Group) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given Account and assigns it to the Owner field.
func (o *Group) SetOwner(v Account) {
	o.Owner = &v
}

// GetWorkspace returns the Workspace field value if set, zero value otherwise.
func (o *Group) GetWorkspace() Workspace {
	if o == nil || o.Workspace == nil {
		var ret Workspace
		return ret
	}
	return *o.Workspace
}

// GetWorkspaceOk returns a tuple with the Workspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetWorkspaceOk() (*Workspace, bool) {
	if o == nil || o.Workspace == nil {
		return nil, false
	}
	return o.Workspace, true
}

// HasWorkspace returns a boolean if a field has been set.
func (o *Group) HasWorkspace() bool {
	if o != nil && o.Workspace != nil {
		return true
	}

	return false
}

// SetWorkspace gets a reference to the given Workspace and assigns it to the Workspace field.
func (o *Group) SetWorkspace(v Workspace) {
	o.Workspace = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Group) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Group) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Group) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *Group) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *Group) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *Group) SetSlug(v string) {
	o.Slug = &v
}

// GetFullSlug returns the FullSlug field value if set, zero value otherwise.
func (o *Group) GetFullSlug() string {
	if o == nil || o.FullSlug == nil {
		var ret string
		return ret
	}
	return *o.FullSlug
}

// GetFullSlugOk returns a tuple with the FullSlug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetFullSlugOk() (*string, bool) {
	if o == nil || o.FullSlug == nil {
		return nil, false
	}
	return o.FullSlug, true
}

// HasFullSlug returns a boolean if a field has been set.
func (o *Group) HasFullSlug() bool {
	if o != nil && o.FullSlug != nil {
		return true
	}

	return false
}

// SetFullSlug gets a reference to the given string and assigns it to the FullSlug field.
func (o *Group) SetFullSlug(v string) {
	o.FullSlug = &v
}

func (o Group) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.Workspace != nil {
		toSerialize["workspace"] = o.Workspace
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	if o.FullSlug != nil {
		toSerialize["full_slug"] = o.FullSlug
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Group) UnmarshalJSON(bytes []byte) (err error) {
	varGroup := _Group{}

	if err = json.Unmarshal(bytes, &varGroup); err == nil {
		*o = Group(varGroup)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "links")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "workspace")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "full_slug")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroup struct {
	value *Group
	isSet bool
}

func (v NullableGroup) Get() *Group {
	return v.value
}

func (v *NullableGroup) Set(val *Group) {
	v.value = val
	v.isSet = true
}

func (v NullableGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroup(val *Group) *NullableGroup {
	return &NullableGroup{value: val, isSet: true}
}

func (v NullableGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


