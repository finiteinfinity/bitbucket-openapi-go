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

// IssueJobStatus The status of an import or export job
type IssueJobStatus struct {
	Type *string `json:"type,omitempty"`
	// The status of the import/export job
	Status *string `json:"status,omitempty"`
	// The phase of the import/export job
	Phase *string `json:"phase,omitempty"`
	// The total number of issues being imported/exported
	Total *int32 `json:"total,omitempty"`
	// The total number of issues already imported/exported
	Count *int32 `json:"count,omitempty"`
	// The percentage of issues already imported/exported
	Pct *float32 `json:"pct,omitempty"`
}

// NewIssueJobStatus instantiates a new IssueJobStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueJobStatus() *IssueJobStatus {
	this := IssueJobStatus{}
	return &this
}

// NewIssueJobStatusWithDefaults instantiates a new IssueJobStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueJobStatusWithDefaults() *IssueJobStatus {
	this := IssueJobStatus{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IssueJobStatus) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueJobStatus) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IssueJobStatus) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IssueJobStatus) SetType(v string) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IssueJobStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueJobStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IssueJobStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IssueJobStatus) SetStatus(v string) {
	o.Status = &v
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *IssueJobStatus) GetPhase() string {
	if o == nil || o.Phase == nil {
		var ret string
		return ret
	}
	return *o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueJobStatus) GetPhaseOk() (*string, bool) {
	if o == nil || o.Phase == nil {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *IssueJobStatus) HasPhase() bool {
	if o != nil && o.Phase != nil {
		return true
	}

	return false
}

// SetPhase gets a reference to the given string and assigns it to the Phase field.
func (o *IssueJobStatus) SetPhase(v string) {
	o.Phase = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *IssueJobStatus) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueJobStatus) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *IssueJobStatus) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *IssueJobStatus) SetTotal(v int32) {
	o.Total = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *IssueJobStatus) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueJobStatus) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *IssueJobStatus) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *IssueJobStatus) SetCount(v int32) {
	o.Count = &v
}

// GetPct returns the Pct field value if set, zero value otherwise.
func (o *IssueJobStatus) GetPct() float32 {
	if o == nil || o.Pct == nil {
		var ret float32
		return ret
	}
	return *o.Pct
}

// GetPctOk returns a tuple with the Pct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueJobStatus) GetPctOk() (*float32, bool) {
	if o == nil || o.Pct == nil {
		return nil, false
	}
	return o.Pct, true
}

// HasPct returns a boolean if a field has been set.
func (o *IssueJobStatus) HasPct() bool {
	if o != nil && o.Pct != nil {
		return true
	}

	return false
}

// SetPct gets a reference to the given float32 and assigns it to the Pct field.
func (o *IssueJobStatus) SetPct(v float32) {
	o.Pct = &v
}

func (o IssueJobStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Phase != nil {
		toSerialize["phase"] = o.Phase
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Pct != nil {
		toSerialize["pct"] = o.Pct
	}
	return json.Marshal(toSerialize)
}

type NullableIssueJobStatus struct {
	value *IssueJobStatus
	isSet bool
}

func (v NullableIssueJobStatus) Get() *IssueJobStatus {
	return v.value
}

func (v *NullableIssueJobStatus) Set(val *IssueJobStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueJobStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueJobStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueJobStatus(val *IssueJobStatus) *NullableIssueJobStatus {
	return &NullableIssueJobStatus{value: val, isSet: true}
}

func (v NullableIssueJobStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueJobStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


