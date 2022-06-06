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

// WorkspaceLinks struct for WorkspaceLinks
type WorkspaceLinks struct {
	Avatar *Link `json:"avatar,omitempty"`
	Html *Link `json:"html,omitempty"`
	Members *Link `json:"members,omitempty"`
	Owners *Link `json:"owners,omitempty"`
	Projects *Link `json:"projects,omitempty"`
	Repositories *Link `json:"repositories,omitempty"`
	Snippets *Link `json:"snippets,omitempty"`
	Self *Link `json:"self,omitempty"`
}

// NewWorkspaceLinks instantiates a new WorkspaceLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkspaceLinks() *WorkspaceLinks {
	this := WorkspaceLinks{}
	return &this
}

// NewWorkspaceLinksWithDefaults instantiates a new WorkspaceLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkspaceLinksWithDefaults() *WorkspaceLinks {
	this := WorkspaceLinks{}
	return &this
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *WorkspaceLinks) GetAvatar() Link {
	if o == nil || o.Avatar == nil {
		var ret Link
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceLinks) GetAvatarOk() (*Link, bool) {
	if o == nil || o.Avatar == nil {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *WorkspaceLinks) HasAvatar() bool {
	if o != nil && o.Avatar != nil {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given Link and assigns it to the Avatar field.
func (o *WorkspaceLinks) SetAvatar(v Link) {
	o.Avatar = &v
}

// GetHtml returns the Html field value if set, zero value otherwise.
func (o *WorkspaceLinks) GetHtml() Link {
	if o == nil || o.Html == nil {
		var ret Link
		return ret
	}
	return *o.Html
}

// GetHtmlOk returns a tuple with the Html field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceLinks) GetHtmlOk() (*Link, bool) {
	if o == nil || o.Html == nil {
		return nil, false
	}
	return o.Html, true
}

// HasHtml returns a boolean if a field has been set.
func (o *WorkspaceLinks) HasHtml() bool {
	if o != nil && o.Html != nil {
		return true
	}

	return false
}

// SetHtml gets a reference to the given Link and assigns it to the Html field.
func (o *WorkspaceLinks) SetHtml(v Link) {
	o.Html = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *WorkspaceLinks) GetMembers() Link {
	if o == nil || o.Members == nil {
		var ret Link
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceLinks) GetMembersOk() (*Link, bool) {
	if o == nil || o.Members == nil {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *WorkspaceLinks) HasMembers() bool {
	if o != nil && o.Members != nil {
		return true
	}

	return false
}

// SetMembers gets a reference to the given Link and assigns it to the Members field.
func (o *WorkspaceLinks) SetMembers(v Link) {
	o.Members = &v
}

// GetOwners returns the Owners field value if set, zero value otherwise.
func (o *WorkspaceLinks) GetOwners() Link {
	if o == nil || o.Owners == nil {
		var ret Link
		return ret
	}
	return *o.Owners
}

// GetOwnersOk returns a tuple with the Owners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceLinks) GetOwnersOk() (*Link, bool) {
	if o == nil || o.Owners == nil {
		return nil, false
	}
	return o.Owners, true
}

// HasOwners returns a boolean if a field has been set.
func (o *WorkspaceLinks) HasOwners() bool {
	if o != nil && o.Owners != nil {
		return true
	}

	return false
}

// SetOwners gets a reference to the given Link and assigns it to the Owners field.
func (o *WorkspaceLinks) SetOwners(v Link) {
	o.Owners = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *WorkspaceLinks) GetProjects() Link {
	if o == nil || o.Projects == nil {
		var ret Link
		return ret
	}
	return *o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceLinks) GetProjectsOk() (*Link, bool) {
	if o == nil || o.Projects == nil {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *WorkspaceLinks) HasProjects() bool {
	if o != nil && o.Projects != nil {
		return true
	}

	return false
}

// SetProjects gets a reference to the given Link and assigns it to the Projects field.
func (o *WorkspaceLinks) SetProjects(v Link) {
	o.Projects = &v
}

// GetRepositories returns the Repositories field value if set, zero value otherwise.
func (o *WorkspaceLinks) GetRepositories() Link {
	if o == nil || o.Repositories == nil {
		var ret Link
		return ret
	}
	return *o.Repositories
}

// GetRepositoriesOk returns a tuple with the Repositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceLinks) GetRepositoriesOk() (*Link, bool) {
	if o == nil || o.Repositories == nil {
		return nil, false
	}
	return o.Repositories, true
}

// HasRepositories returns a boolean if a field has been set.
func (o *WorkspaceLinks) HasRepositories() bool {
	if o != nil && o.Repositories != nil {
		return true
	}

	return false
}

// SetRepositories gets a reference to the given Link and assigns it to the Repositories field.
func (o *WorkspaceLinks) SetRepositories(v Link) {
	o.Repositories = &v
}

// GetSnippets returns the Snippets field value if set, zero value otherwise.
func (o *WorkspaceLinks) GetSnippets() Link {
	if o == nil || o.Snippets == nil {
		var ret Link
		return ret
	}
	return *o.Snippets
}

// GetSnippetsOk returns a tuple with the Snippets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceLinks) GetSnippetsOk() (*Link, bool) {
	if o == nil || o.Snippets == nil {
		return nil, false
	}
	return o.Snippets, true
}

// HasSnippets returns a boolean if a field has been set.
func (o *WorkspaceLinks) HasSnippets() bool {
	if o != nil && o.Snippets != nil {
		return true
	}

	return false
}

// SetSnippets gets a reference to the given Link and assigns it to the Snippets field.
func (o *WorkspaceLinks) SetSnippets(v Link) {
	o.Snippets = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *WorkspaceLinks) GetSelf() Link {
	if o == nil || o.Self == nil {
		var ret Link
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceLinks) GetSelfOk() (*Link, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *WorkspaceLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given Link and assigns it to the Self field.
func (o *WorkspaceLinks) SetSelf(v Link) {
	o.Self = &v
}

func (o WorkspaceLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Avatar != nil {
		toSerialize["avatar"] = o.Avatar
	}
	if o.Html != nil {
		toSerialize["html"] = o.Html
	}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	if o.Owners != nil {
		toSerialize["owners"] = o.Owners
	}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	if o.Repositories != nil {
		toSerialize["repositories"] = o.Repositories
	}
	if o.Snippets != nil {
		toSerialize["snippets"] = o.Snippets
	}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	return json.Marshal(toSerialize)
}

type NullableWorkspaceLinks struct {
	value *WorkspaceLinks
	isSet bool
}

func (v NullableWorkspaceLinks) Get() *WorkspaceLinks {
	return v.value
}

func (v *NullableWorkspaceLinks) Set(val *WorkspaceLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspaceLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspaceLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspaceLinks(val *WorkspaceLinks) *NullableWorkspaceLinks {
	return &NullableWorkspaceLinks{value: val, isSet: true}
}

func (v NullableWorkspaceLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspaceLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


