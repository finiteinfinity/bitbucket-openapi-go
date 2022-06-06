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

// PipelineSchedule A Pipelines schedule.
type PipelineSchedule struct {
	// The UUID identifying the schedule.
	Uuid *string `json:"uuid,omitempty"`
	// Whether the schedule is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	Target *PipelineTarget `json:"target,omitempty"`
	Selector *PipelineSelector `json:"selector,omitempty"`
	// The cron expression that the schedule applies.
	CronPattern *string `json:"cron_pattern,omitempty"`
	// The timestamp when the schedule was created.
	CreatedOn *time.Time `json:"created_on,omitempty"`
	// The timestamp when the schedule was updated.
	UpdatedOn *time.Time `json:"updated_on,omitempty"`
}

// NewPipelineSchedule instantiates a new PipelineSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineSchedule() *PipelineSchedule {
	this := PipelineSchedule{}
	return &this
}

// NewPipelineScheduleWithDefaults instantiates a new PipelineSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineScheduleWithDefaults() *PipelineSchedule {
	this := PipelineSchedule{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *PipelineSchedule) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineSchedule) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *PipelineSchedule) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *PipelineSchedule) SetUuid(v string) {
	o.Uuid = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PipelineSchedule) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineSchedule) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PipelineSchedule) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PipelineSchedule) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *PipelineSchedule) GetTarget() PipelineTarget {
	if o == nil || o.Target == nil {
		var ret PipelineTarget
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineSchedule) GetTargetOk() (*PipelineTarget, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *PipelineSchedule) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given PipelineTarget and assigns it to the Target field.
func (o *PipelineSchedule) SetTarget(v PipelineTarget) {
	o.Target = &v
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *PipelineSchedule) GetSelector() PipelineSelector {
	if o == nil || o.Selector == nil {
		var ret PipelineSelector
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineSchedule) GetSelectorOk() (*PipelineSelector, bool) {
	if o == nil || o.Selector == nil {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *PipelineSchedule) HasSelector() bool {
	if o != nil && o.Selector != nil {
		return true
	}

	return false
}

// SetSelector gets a reference to the given PipelineSelector and assigns it to the Selector field.
func (o *PipelineSchedule) SetSelector(v PipelineSelector) {
	o.Selector = &v
}

// GetCronPattern returns the CronPattern field value if set, zero value otherwise.
func (o *PipelineSchedule) GetCronPattern() string {
	if o == nil || o.CronPattern == nil {
		var ret string
		return ret
	}
	return *o.CronPattern
}

// GetCronPatternOk returns a tuple with the CronPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineSchedule) GetCronPatternOk() (*string, bool) {
	if o == nil || o.CronPattern == nil {
		return nil, false
	}
	return o.CronPattern, true
}

// HasCronPattern returns a boolean if a field has been set.
func (o *PipelineSchedule) HasCronPattern() bool {
	if o != nil && o.CronPattern != nil {
		return true
	}

	return false
}

// SetCronPattern gets a reference to the given string and assigns it to the CronPattern field.
func (o *PipelineSchedule) SetCronPattern(v string) {
	o.CronPattern = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *PipelineSchedule) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineSchedule) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *PipelineSchedule) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *PipelineSchedule) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetUpdatedOn returns the UpdatedOn field value if set, zero value otherwise.
func (o *PipelineSchedule) GetUpdatedOn() time.Time {
	if o == nil || o.UpdatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedOn
}

// GetUpdatedOnOk returns a tuple with the UpdatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineSchedule) GetUpdatedOnOk() (*time.Time, bool) {
	if o == nil || o.UpdatedOn == nil {
		return nil, false
	}
	return o.UpdatedOn, true
}

// HasUpdatedOn returns a boolean if a field has been set.
func (o *PipelineSchedule) HasUpdatedOn() bool {
	if o != nil && o.UpdatedOn != nil {
		return true
	}

	return false
}

// SetUpdatedOn gets a reference to the given time.Time and assigns it to the UpdatedOn field.
func (o *PipelineSchedule) SetUpdatedOn(v time.Time) {
	o.UpdatedOn = &v
}

func (o PipelineSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.Selector != nil {
		toSerialize["selector"] = o.Selector
	}
	if o.CronPattern != nil {
		toSerialize["cron_pattern"] = o.CronPattern
	}
	if o.CreatedOn != nil {
		toSerialize["created_on"] = o.CreatedOn
	}
	if o.UpdatedOn != nil {
		toSerialize["updated_on"] = o.UpdatedOn
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineSchedule struct {
	value *PipelineSchedule
	isSet bool
}

func (v NullablePipelineSchedule) Get() *PipelineSchedule {
	return v.value
}

func (v *NullablePipelineSchedule) Set(val *PipelineSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineSchedule(val *PipelineSchedule) *NullablePipelineSchedule {
	return &NullablePipelineSchedule{value: val, isSet: true}
}

func (v NullablePipelineSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

