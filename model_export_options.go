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

// ExportOptions Options for issue export.
type ExportOptions struct {
	Type string `json:"type"`
	ProjectKey *string `json:"project_key,omitempty"`
	ProjectName *string `json:"project_name,omitempty"`
	SendEmail *bool `json:"send_email,omitempty"`
	IncludeAttachments *bool `json:"include_attachments,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExportOptions ExportOptions

// NewExportOptions instantiates a new ExportOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportOptions() *ExportOptions {
	this := ExportOptions{}
	return &this
}

// NewExportOptionsWithDefaults instantiates a new ExportOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportOptionsWithDefaults() *ExportOptions {
	this := ExportOptions{}
	return &this
}

// GetType returns the Type field value
func (o *ExportOptions) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExportOptions) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExportOptions) SetType(v string) {
	o.Type = v
}

// GetProjectKey returns the ProjectKey field value if set, zero value otherwise.
func (o *ExportOptions) GetProjectKey() string {
	if o == nil || o.ProjectKey == nil {
		var ret string
		return ret
	}
	return *o.ProjectKey
}

// GetProjectKeyOk returns a tuple with the ProjectKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportOptions) GetProjectKeyOk() (*string, bool) {
	if o == nil || o.ProjectKey == nil {
		return nil, false
	}
	return o.ProjectKey, true
}

// HasProjectKey returns a boolean if a field has been set.
func (o *ExportOptions) HasProjectKey() bool {
	if o != nil && o.ProjectKey != nil {
		return true
	}

	return false
}

// SetProjectKey gets a reference to the given string and assigns it to the ProjectKey field.
func (o *ExportOptions) SetProjectKey(v string) {
	o.ProjectKey = &v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise.
func (o *ExportOptions) GetProjectName() string {
	if o == nil || o.ProjectName == nil {
		var ret string
		return ret
	}
	return *o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportOptions) GetProjectNameOk() (*string, bool) {
	if o == nil || o.ProjectName == nil {
		return nil, false
	}
	return o.ProjectName, true
}

// HasProjectName returns a boolean if a field has been set.
func (o *ExportOptions) HasProjectName() bool {
	if o != nil && o.ProjectName != nil {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given string and assigns it to the ProjectName field.
func (o *ExportOptions) SetProjectName(v string) {
	o.ProjectName = &v
}

// GetSendEmail returns the SendEmail field value if set, zero value otherwise.
func (o *ExportOptions) GetSendEmail() bool {
	if o == nil || o.SendEmail == nil {
		var ret bool
		return ret
	}
	return *o.SendEmail
}

// GetSendEmailOk returns a tuple with the SendEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportOptions) GetSendEmailOk() (*bool, bool) {
	if o == nil || o.SendEmail == nil {
		return nil, false
	}
	return o.SendEmail, true
}

// HasSendEmail returns a boolean if a field has been set.
func (o *ExportOptions) HasSendEmail() bool {
	if o != nil && o.SendEmail != nil {
		return true
	}

	return false
}

// SetSendEmail gets a reference to the given bool and assigns it to the SendEmail field.
func (o *ExportOptions) SetSendEmail(v bool) {
	o.SendEmail = &v
}

// GetIncludeAttachments returns the IncludeAttachments field value if set, zero value otherwise.
func (o *ExportOptions) GetIncludeAttachments() bool {
	if o == nil || o.IncludeAttachments == nil {
		var ret bool
		return ret
	}
	return *o.IncludeAttachments
}

// GetIncludeAttachmentsOk returns a tuple with the IncludeAttachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportOptions) GetIncludeAttachmentsOk() (*bool, bool) {
	if o == nil || o.IncludeAttachments == nil {
		return nil, false
	}
	return o.IncludeAttachments, true
}

// HasIncludeAttachments returns a boolean if a field has been set.
func (o *ExportOptions) HasIncludeAttachments() bool {
	if o != nil && o.IncludeAttachments != nil {
		return true
	}

	return false
}

// SetIncludeAttachments gets a reference to the given bool and assigns it to the IncludeAttachments field.
func (o *ExportOptions) SetIncludeAttachments(v bool) {
	o.IncludeAttachments = &v
}

func (o ExportOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.ProjectKey != nil {
		toSerialize["project_key"] = o.ProjectKey
	}
	if o.ProjectName != nil {
		toSerialize["project_name"] = o.ProjectName
	}
	if o.SendEmail != nil {
		toSerialize["send_email"] = o.SendEmail
	}
	if o.IncludeAttachments != nil {
		toSerialize["include_attachments"] = o.IncludeAttachments
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ExportOptions) UnmarshalJSON(bytes []byte) (err error) {
	varExportOptions := _ExportOptions{}

	if err = json.Unmarshal(bytes, &varExportOptions); err == nil {
		*o = ExportOptions(varExportOptions)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "project_key")
		delete(additionalProperties, "project_name")
		delete(additionalProperties, "send_email")
		delete(additionalProperties, "include_attachments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExportOptions struct {
	value *ExportOptions
	isSet bool
}

func (v NullableExportOptions) Get() *ExportOptions {
	return v.value
}

func (v *NullableExportOptions) Set(val *ExportOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableExportOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableExportOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportOptions(val *ExportOptions) *NullableExportOptions {
	return &NullableExportOptions{value: val, isSet: true}
}

func (v NullableExportOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


