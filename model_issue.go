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

// Issue An issue.
type Issue struct {
	Links *IssueLinks `json:"links,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Repository *Repository `json:"repository,omitempty"`
	Title *string `json:"title,omitempty"`
	Reporter *User `json:"reporter,omitempty"`
	Assignee *User `json:"assignee,omitempty"`
	CreatedOn *time.Time `json:"created_on,omitempty"`
	UpdatedOn *time.Time `json:"updated_on,omitempty"`
	EditedOn *time.Time `json:"edited_on,omitempty"`
	State *string `json:"state,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Priority *string `json:"priority,omitempty"`
	Milestone *Milestone `json:"milestone,omitempty"`
	Version *Version `json:"version,omitempty"`
	Component *Component `json:"component,omitempty"`
	Votes *int32 `json:"votes,omitempty"`
	Content *RenderedPullRequestMarkupTitle `json:"content,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Issue Issue

// NewIssue instantiates a new Issue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssue() *Issue {
	this := Issue{}
	return &this
}

// NewIssueWithDefaults instantiates a new Issue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueWithDefaults() *Issue {
	this := Issue{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Issue) GetLinks() IssueLinks {
	if o == nil || o.Links == nil {
		var ret IssueLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetLinksOk() (*IssueLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Issue) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given IssueLinks and assigns it to the Links field.
func (o *Issue) SetLinks(v IssueLinks) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Issue) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Issue) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Issue) SetId(v int32) {
	o.Id = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *Issue) GetRepository() Repository {
	if o == nil || o.Repository == nil {
		var ret Repository
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetRepositoryOk() (*Repository, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *Issue) HasRepository() bool {
	if o != nil && o.Repository != nil {
		return true
	}

	return false
}

// SetRepository gets a reference to the given Repository and assigns it to the Repository field.
func (o *Issue) SetRepository(v Repository) {
	o.Repository = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Issue) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Issue) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Issue) SetTitle(v string) {
	o.Title = &v
}

// GetReporter returns the Reporter field value if set, zero value otherwise.
func (o *Issue) GetReporter() User {
	if o == nil || o.Reporter == nil {
		var ret User
		return ret
	}
	return *o.Reporter
}

// GetReporterOk returns a tuple with the Reporter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetReporterOk() (*User, bool) {
	if o == nil || o.Reporter == nil {
		return nil, false
	}
	return o.Reporter, true
}

// HasReporter returns a boolean if a field has been set.
func (o *Issue) HasReporter() bool {
	if o != nil && o.Reporter != nil {
		return true
	}

	return false
}

// SetReporter gets a reference to the given User and assigns it to the Reporter field.
func (o *Issue) SetReporter(v User) {
	o.Reporter = &v
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *Issue) GetAssignee() User {
	if o == nil || o.Assignee == nil {
		var ret User
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetAssigneeOk() (*User, bool) {
	if o == nil || o.Assignee == nil {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *Issue) HasAssignee() bool {
	if o != nil && o.Assignee != nil {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given User and assigns it to the Assignee field.
func (o *Issue) SetAssignee(v User) {
	o.Assignee = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *Issue) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *Issue) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *Issue) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetUpdatedOn returns the UpdatedOn field value if set, zero value otherwise.
func (o *Issue) GetUpdatedOn() time.Time {
	if o == nil || o.UpdatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedOn
}

// GetUpdatedOnOk returns a tuple with the UpdatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetUpdatedOnOk() (*time.Time, bool) {
	if o == nil || o.UpdatedOn == nil {
		return nil, false
	}
	return o.UpdatedOn, true
}

// HasUpdatedOn returns a boolean if a field has been set.
func (o *Issue) HasUpdatedOn() bool {
	if o != nil && o.UpdatedOn != nil {
		return true
	}

	return false
}

// SetUpdatedOn gets a reference to the given time.Time and assigns it to the UpdatedOn field.
func (o *Issue) SetUpdatedOn(v time.Time) {
	o.UpdatedOn = &v
}

// GetEditedOn returns the EditedOn field value if set, zero value otherwise.
func (o *Issue) GetEditedOn() time.Time {
	if o == nil || o.EditedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.EditedOn
}

// GetEditedOnOk returns a tuple with the EditedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetEditedOnOk() (*time.Time, bool) {
	if o == nil || o.EditedOn == nil {
		return nil, false
	}
	return o.EditedOn, true
}

// HasEditedOn returns a boolean if a field has been set.
func (o *Issue) HasEditedOn() bool {
	if o != nil && o.EditedOn != nil {
		return true
	}

	return false
}

// SetEditedOn gets a reference to the given time.Time and assigns it to the EditedOn field.
func (o *Issue) SetEditedOn(v time.Time) {
	o.EditedOn = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Issue) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Issue) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Issue) SetState(v string) {
	o.State = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *Issue) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *Issue) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *Issue) SetKind(v string) {
	o.Kind = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *Issue) GetPriority() string {
	if o == nil || o.Priority == nil {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetPriorityOk() (*string, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *Issue) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *Issue) SetPriority(v string) {
	o.Priority = &v
}

// GetMilestone returns the Milestone field value if set, zero value otherwise.
func (o *Issue) GetMilestone() Milestone {
	if o == nil || o.Milestone == nil {
		var ret Milestone
		return ret
	}
	return *o.Milestone
}

// GetMilestoneOk returns a tuple with the Milestone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetMilestoneOk() (*Milestone, bool) {
	if o == nil || o.Milestone == nil {
		return nil, false
	}
	return o.Milestone, true
}

// HasMilestone returns a boolean if a field has been set.
func (o *Issue) HasMilestone() bool {
	if o != nil && o.Milestone != nil {
		return true
	}

	return false
}

// SetMilestone gets a reference to the given Milestone and assigns it to the Milestone field.
func (o *Issue) SetMilestone(v Milestone) {
	o.Milestone = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Issue) GetVersion() Version {
	if o == nil || o.Version == nil {
		var ret Version
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetVersionOk() (*Version, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Issue) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given Version and assigns it to the Version field.
func (o *Issue) SetVersion(v Version) {
	o.Version = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *Issue) GetComponent() Component {
	if o == nil || o.Component == nil {
		var ret Component
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetComponentOk() (*Component, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *Issue) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given Component and assigns it to the Component field.
func (o *Issue) SetComponent(v Component) {
	o.Component = &v
}

// GetVotes returns the Votes field value if set, zero value otherwise.
func (o *Issue) GetVotes() int32 {
	if o == nil || o.Votes == nil {
		var ret int32
		return ret
	}
	return *o.Votes
}

// GetVotesOk returns a tuple with the Votes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetVotesOk() (*int32, bool) {
	if o == nil || o.Votes == nil {
		return nil, false
	}
	return o.Votes, true
}

// HasVotes returns a boolean if a field has been set.
func (o *Issue) HasVotes() bool {
	if o != nil && o.Votes != nil {
		return true
	}

	return false
}

// SetVotes gets a reference to the given int32 and assigns it to the Votes field.
func (o *Issue) SetVotes(v int32) {
	o.Votes = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *Issue) GetContent() RenderedPullRequestMarkupTitle {
	if o == nil || o.Content == nil {
		var ret RenderedPullRequestMarkupTitle
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetContentOk() (*RenderedPullRequestMarkupTitle, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *Issue) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given RenderedPullRequestMarkupTitle and assigns it to the Content field.
func (o *Issue) SetContent(v RenderedPullRequestMarkupTitle) {
	o.Content = &v
}

func (o Issue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Repository != nil {
		toSerialize["repository"] = o.Repository
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Reporter != nil {
		toSerialize["reporter"] = o.Reporter
	}
	if o.Assignee != nil {
		toSerialize["assignee"] = o.Assignee
	}
	if o.CreatedOn != nil {
		toSerialize["created_on"] = o.CreatedOn
	}
	if o.UpdatedOn != nil {
		toSerialize["updated_on"] = o.UpdatedOn
	}
	if o.EditedOn != nil {
		toSerialize["edited_on"] = o.EditedOn
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Milestone != nil {
		toSerialize["milestone"] = o.Milestone
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.Votes != nil {
		toSerialize["votes"] = o.Votes
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Issue) UnmarshalJSON(bytes []byte) (err error) {
	varIssue := _Issue{}

	if err = json.Unmarshal(bytes, &varIssue); err == nil {
		*o = Issue(varIssue)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "links")
		delete(additionalProperties, "id")
		delete(additionalProperties, "repository")
		delete(additionalProperties, "title")
		delete(additionalProperties, "reporter")
		delete(additionalProperties, "assignee")
		delete(additionalProperties, "created_on")
		delete(additionalProperties, "updated_on")
		delete(additionalProperties, "edited_on")
		delete(additionalProperties, "state")
		delete(additionalProperties, "kind")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "milestone")
		delete(additionalProperties, "version")
		delete(additionalProperties, "component")
		delete(additionalProperties, "votes")
		delete(additionalProperties, "content")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIssue struct {
	value *Issue
	isSet bool
}

func (v NullableIssue) Get() *Issue {
	return v.value
}

func (v *NullableIssue) Set(val *Issue) {
	v.value = val
	v.isSet = true
}

func (v NullableIssue) IsSet() bool {
	return v.isSet
}

func (v *NullableIssue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssue(val *Issue) *NullableIssue {
	return &NullableIssue{value: val, isSet: true}
}

func (v NullableIssue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


