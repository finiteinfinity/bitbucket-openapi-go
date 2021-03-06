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

// DeploymentStateInProgressAllOf A Bitbucket Deployment IN_PROGRESS deployment state.
type DeploymentStateInProgressAllOf struct {
	// The name of deployment state (IN_PROGRESS).
	Name *string `json:"name,omitempty"`
	// Link to the deployment result.
	Url *string `json:"url,omitempty"`
	Deployer *Account `json:"deployer,omitempty"`
	// The timestamp when the deployment was started.
	StartDate *time.Time `json:"start_date,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeploymentStateInProgressAllOf DeploymentStateInProgressAllOf

// NewDeploymentStateInProgressAllOf instantiates a new DeploymentStateInProgressAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentStateInProgressAllOf() *DeploymentStateInProgressAllOf {
	this := DeploymentStateInProgressAllOf{}
	return &this
}

// NewDeploymentStateInProgressAllOfWithDefaults instantiates a new DeploymentStateInProgressAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentStateInProgressAllOfWithDefaults() *DeploymentStateInProgressAllOf {
	this := DeploymentStateInProgressAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentStateInProgressAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStateInProgressAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentStateInProgressAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentStateInProgressAllOf) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *DeploymentStateInProgressAllOf) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStateInProgressAllOf) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *DeploymentStateInProgressAllOf) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *DeploymentStateInProgressAllOf) SetUrl(v string) {
	o.Url = &v
}

// GetDeployer returns the Deployer field value if set, zero value otherwise.
func (o *DeploymentStateInProgressAllOf) GetDeployer() Account {
	if o == nil || o.Deployer == nil {
		var ret Account
		return ret
	}
	return *o.Deployer
}

// GetDeployerOk returns a tuple with the Deployer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStateInProgressAllOf) GetDeployerOk() (*Account, bool) {
	if o == nil || o.Deployer == nil {
		return nil, false
	}
	return o.Deployer, true
}

// HasDeployer returns a boolean if a field has been set.
func (o *DeploymentStateInProgressAllOf) HasDeployer() bool {
	if o != nil && o.Deployer != nil {
		return true
	}

	return false
}

// SetDeployer gets a reference to the given Account and assigns it to the Deployer field.
func (o *DeploymentStateInProgressAllOf) SetDeployer(v Account) {
	o.Deployer = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *DeploymentStateInProgressAllOf) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStateInProgressAllOf) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *DeploymentStateInProgressAllOf) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *DeploymentStateInProgressAllOf) SetStartDate(v time.Time) {
	o.StartDate = &v
}

func (o DeploymentStateInProgressAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Deployer != nil {
		toSerialize["deployer"] = o.Deployer
	}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeploymentStateInProgressAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varDeploymentStateInProgressAllOf := _DeploymentStateInProgressAllOf{}

	if err = json.Unmarshal(bytes, &varDeploymentStateInProgressAllOf); err == nil {
		*o = DeploymentStateInProgressAllOf(varDeploymentStateInProgressAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "url")
		delete(additionalProperties, "deployer")
		delete(additionalProperties, "start_date")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeploymentStateInProgressAllOf struct {
	value *DeploymentStateInProgressAllOf
	isSet bool
}

func (v NullableDeploymentStateInProgressAllOf) Get() *DeploymentStateInProgressAllOf {
	return v.value
}

func (v *NullableDeploymentStateInProgressAllOf) Set(val *DeploymentStateInProgressAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentStateInProgressAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentStateInProgressAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentStateInProgressAllOf(val *DeploymentStateInProgressAllOf) *NullableDeploymentStateInProgressAllOf {
	return &NullableDeploymentStateInProgressAllOf{value: val, isSet: true}
}

func (v NullableDeploymentStateInProgressAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentStateInProgressAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


