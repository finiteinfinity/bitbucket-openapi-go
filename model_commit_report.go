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

// CommitReport A report for a commit.
type CommitReport struct {
	// The UUID that can be used to identify the report.
	Uuid *string `json:"uuid,omitempty"`
	// The title of the report.
	Title *string `json:"title,omitempty"`
	// A string to describe the purpose of the report.
	Details *string `json:"details,omitempty"`
	// ID of the report provided by the report creator. It can be used to identify the report as an alternative to it's generated uuid. It is not used by Bitbucket, but only by the report creator for updating or deleting this specific report. Needs to be unique.
	ExternalId *string `json:"external_id,omitempty"`
	// A string to describe the tool or company who created the report.
	Reporter *string `json:"reporter,omitempty"`
	// A URL linking to the results of the report in an external tool.
	Link *string `json:"link,omitempty"`
	// If enabled, a remote link is created in Jira for the issue associated with the commit the report belongs to.
	RemoteLinkEnabled *bool `json:"remote_link_enabled,omitempty"`
	// A URL to the report logo. If none is provided, the default insights logo will be used.
	LogoUrl *string `json:"logo_url,omitempty"`
	// The type of the report.
	ReportType *string `json:"report_type,omitempty"`
	// The state of the report. May be set to PENDING and later updated.
	Result *string `json:"result,omitempty"`
	// An array of data fields to display information on the report. Maximum 10.
	Data []ReportData `json:"data,omitempty"`
	// The timestamp when the report was created.
	CreatedOn *time.Time `json:"created_on,omitempty"`
	// The timestamp when the report was updated.
	UpdatedOn *time.Time `json:"updated_on,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommitReport CommitReport

// NewCommitReport instantiates a new CommitReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommitReport() *CommitReport {
	this := CommitReport{}
	return &this
}

// NewCommitReportWithDefaults instantiates a new CommitReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitReportWithDefaults() *CommitReport {
	this := CommitReport{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *CommitReport) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitReport) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *CommitReport) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *CommitReport) SetUuid(v string) {
	o.Uuid = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CommitReport) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitReport) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CommitReport) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CommitReport) SetTitle(v string) {
	o.Title = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *CommitReport) GetDetails() string {
	if o == nil || o.Details == nil {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitReport) GetDetailsOk() (*string, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *CommitReport) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *CommitReport) SetDetails(v string) {
	o.Details = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *CommitReport) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitReport) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *CommitReport) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *CommitReport) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetReporter returns the Reporter field value if set, zero value otherwise.
func (o *CommitReport) GetReporter() string {
	if o == nil || o.Reporter == nil {
		var ret string
		return ret
	}
	return *o.Reporter
}

// GetReporterOk returns a tuple with the Reporter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitReport) GetReporterOk() (*string, bool) {
	if o == nil || o.Reporter == nil {
		return nil, false
	}
	return o.Reporter, true
}

// HasReporter returns a boolean if a field has been set.
func (o *CommitReport) HasReporter() bool {
	if o != nil && o.Reporter != nil {
		return true
	}

	return false
}

// SetReporter gets a reference to the given string and assigns it to the Reporter field.
func (o *CommitReport) SetReporter(v string) {
	o.Reporter = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *CommitReport) GetLink() string {
	if o == nil || o.Link == nil {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitReport) GetLinkOk() (*string, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *CommitReport) HasLink() bool {
	if o != nil && o.Link != nil {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *CommitReport) SetLink(v string) {
	o.Link = &v
}

// GetRemoteLinkEnabled returns the RemoteLinkEnabled field value if set, zero value otherwise.
func (o *CommitReport) GetRemoteLinkEnabled() bool {
	if o == nil || o.RemoteLinkEnabled == nil {
		var ret bool
		return ret
	}
	return *o.RemoteLinkEnabled
}

// GetRemoteLinkEnabledOk returns a tuple with the RemoteLinkEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitReport) GetRemoteLinkEnabledOk() (*bool, bool) {
	if o == nil || o.RemoteLinkEnabled == nil {
		return nil, false
	}
	return o.RemoteLinkEnabled, true
}

// HasRemoteLinkEnabled returns a boolean if a field has been set.
func (o *CommitReport) HasRemoteLinkEnabled() bool {
	if o != nil && o.RemoteLinkEnabled != nil {
		return true
	}

	return false
}

// SetRemoteLinkEnabled gets a reference to the given bool and assigns it to the RemoteLinkEnabled field.
func (o *CommitReport) SetRemoteLinkEnabled(v bool) {
	o.RemoteLinkEnabled = &v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *CommitReport) GetLogoUrl() string {
	if o == nil || o.LogoUrl == nil {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitReport) GetLogoUrlOk() (*string, bool) {
	if o == nil || o.LogoUrl == nil {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *CommitReport) HasLogoUrl() bool {
	if o != nil && o.LogoUrl != nil {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *CommitReport) SetLogoUrl(v string) {
	o.LogoUrl = &v
}

// GetReportType returns the ReportType field value if set, zero value otherwise.
func (o *CommitReport) GetReportType() string {
	if o == nil || o.ReportType == nil {
		var ret string
		return ret
	}
	return *o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitReport) GetReportTypeOk() (*string, bool) {
	if o == nil || o.ReportType == nil {
		return nil, false
	}
	return o.ReportType, true
}

// HasReportType returns a boolean if a field has been set.
func (o *CommitReport) HasReportType() bool {
	if o != nil && o.ReportType != nil {
		return true
	}

	return false
}

// SetReportType gets a reference to the given string and assigns it to the ReportType field.
func (o *CommitReport) SetReportType(v string) {
	o.ReportType = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CommitReport) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitReport) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CommitReport) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *CommitReport) SetResult(v string) {
	o.Result = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CommitReport) GetData() []ReportData {
	if o == nil || o.Data == nil {
		var ret []ReportData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitReport) GetDataOk() ([]ReportData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CommitReport) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []ReportData and assigns it to the Data field.
func (o *CommitReport) SetData(v []ReportData) {
	o.Data = v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *CommitReport) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitReport) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *CommitReport) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *CommitReport) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetUpdatedOn returns the UpdatedOn field value if set, zero value otherwise.
func (o *CommitReport) GetUpdatedOn() time.Time {
	if o == nil || o.UpdatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedOn
}

// GetUpdatedOnOk returns a tuple with the UpdatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitReport) GetUpdatedOnOk() (*time.Time, bool) {
	if o == nil || o.UpdatedOn == nil {
		return nil, false
	}
	return o.UpdatedOn, true
}

// HasUpdatedOn returns a boolean if a field has been set.
func (o *CommitReport) HasUpdatedOn() bool {
	if o != nil && o.UpdatedOn != nil {
		return true
	}

	return false
}

// SetUpdatedOn gets a reference to the given time.Time and assigns it to the UpdatedOn field.
func (o *CommitReport) SetUpdatedOn(v time.Time) {
	o.UpdatedOn = &v
}

func (o CommitReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.ExternalId != nil {
		toSerialize["external_id"] = o.ExternalId
	}
	if o.Reporter != nil {
		toSerialize["reporter"] = o.Reporter
	}
	if o.Link != nil {
		toSerialize["link"] = o.Link
	}
	if o.RemoteLinkEnabled != nil {
		toSerialize["remote_link_enabled"] = o.RemoteLinkEnabled
	}
	if o.LogoUrl != nil {
		toSerialize["logo_url"] = o.LogoUrl
	}
	if o.ReportType != nil {
		toSerialize["report_type"] = o.ReportType
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
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

func (o *CommitReport) UnmarshalJSON(bytes []byte) (err error) {
	varCommitReport := _CommitReport{}

	if err = json.Unmarshal(bytes, &varCommitReport); err == nil {
		*o = CommitReport(varCommitReport)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "title")
		delete(additionalProperties, "details")
		delete(additionalProperties, "external_id")
		delete(additionalProperties, "reporter")
		delete(additionalProperties, "link")
		delete(additionalProperties, "remote_link_enabled")
		delete(additionalProperties, "logo_url")
		delete(additionalProperties, "report_type")
		delete(additionalProperties, "result")
		delete(additionalProperties, "data")
		delete(additionalProperties, "created_on")
		delete(additionalProperties, "updated_on")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommitReport struct {
	value *CommitReport
	isSet bool
}

func (v NullableCommitReport) Get() *CommitReport {
	return v.value
}

func (v *NullableCommitReport) Set(val *CommitReport) {
	v.value = val
	v.isSet = true
}

func (v NullableCommitReport) IsSet() bool {
	return v.isSet
}

func (v *NullableCommitReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommitReport(val *CommitReport) *NullableCommitReport {
	return &NullableCommitReport{value: val, isSet: true}
}

func (v NullableCommitReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommitReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


