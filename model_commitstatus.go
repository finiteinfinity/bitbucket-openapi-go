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
	"reflect"
	"strings"
)

// Commitstatus struct for Commitstatus
type Commitstatus struct {
	Object
	Links *CommitStatusLinks `json:"links,omitempty"`
	// The commit status' id.
	Uuid *string `json:"uuid,omitempty"`
	// An identifier for the status that's unique to         its type (current \"build\" is the only supported type) and the vendor,         e.g. BB-DEPLOY
	Key *string `json:"key,omitempty"`
	//  The name of the ref that pointed to this commit at the time the status object was created. Note that this the ref may since have moved off of the commit. This optional field can be useful for build systems whose build triggers and configuration are branch-dependent (e.g. a Pipeline build). It is legitimate for this field to not be set, or even apply (e.g. a static linting job).
	Refname *string `json:"refname,omitempty"`
	// A URL linking back to the vendor or build system, for providing more information about whatever process produced this status. Accepts context variables `repository` and `commit` that Bitbucket will evaluate at runtime whenever at runtime. For example, one could use https://foo.com/builds/{repository.full_name} which Bitbucket will turn into https://foo.com/builds/foo/bar at render time.
	Url *string `json:"url,omitempty"`
	// Provides some indication of the status of this commit
	State *string `json:"state,omitempty"`
	// An identifier for the build itself, e.g. BB-DEPLOY-1
	Name *string `json:"name,omitempty"`
	// A description of the build (e.g. \"Unit tests in Bamboo\")
	Description *string `json:"description,omitempty"`
	CreatedOn *time.Time `json:"created_on,omitempty"`
	UpdatedOn *time.Time `json:"updated_on,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Commitstatus Commitstatus

// NewCommitstatus instantiates a new Commitstatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommitstatus() *Commitstatus {
	this := Commitstatus{}
	return &this
}

// NewCommitstatusWithDefaults instantiates a new Commitstatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitstatusWithDefaults() *Commitstatus {
	this := Commitstatus{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Commitstatus) GetLinks() CommitStatusLinks {
	if o == nil || o.Links == nil {
		var ret CommitStatusLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commitstatus) GetLinksOk() (*CommitStatusLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Commitstatus) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given CommitStatusLinks and assigns it to the Links field.
func (o *Commitstatus) SetLinks(v CommitStatusLinks) {
	o.Links = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *Commitstatus) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commitstatus) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *Commitstatus) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *Commitstatus) SetUuid(v string) {
	o.Uuid = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *Commitstatus) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commitstatus) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *Commitstatus) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *Commitstatus) SetKey(v string) {
	o.Key = &v
}

// GetRefname returns the Refname field value if set, zero value otherwise.
func (o *Commitstatus) GetRefname() string {
	if o == nil || o.Refname == nil {
		var ret string
		return ret
	}
	return *o.Refname
}

// GetRefnameOk returns a tuple with the Refname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commitstatus) GetRefnameOk() (*string, bool) {
	if o == nil || o.Refname == nil {
		return nil, false
	}
	return o.Refname, true
}

// HasRefname returns a boolean if a field has been set.
func (o *Commitstatus) HasRefname() bool {
	if o != nil && o.Refname != nil {
		return true
	}

	return false
}

// SetRefname gets a reference to the given string and assigns it to the Refname field.
func (o *Commitstatus) SetRefname(v string) {
	o.Refname = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Commitstatus) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commitstatus) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Commitstatus) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Commitstatus) SetUrl(v string) {
	o.Url = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Commitstatus) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commitstatus) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Commitstatus) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Commitstatus) SetState(v string) {
	o.State = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Commitstatus) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commitstatus) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Commitstatus) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Commitstatus) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Commitstatus) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commitstatus) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Commitstatus) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Commitstatus) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *Commitstatus) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commitstatus) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *Commitstatus) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *Commitstatus) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetUpdatedOn returns the UpdatedOn field value if set, zero value otherwise.
func (o *Commitstatus) GetUpdatedOn() time.Time {
	if o == nil || o.UpdatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedOn
}

// GetUpdatedOnOk returns a tuple with the UpdatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commitstatus) GetUpdatedOnOk() (*time.Time, bool) {
	if o == nil || o.UpdatedOn == nil {
		return nil, false
	}
	return o.UpdatedOn, true
}

// HasUpdatedOn returns a boolean if a field has been set.
func (o *Commitstatus) HasUpdatedOn() bool {
	if o != nil && o.UpdatedOn != nil {
		return true
	}

	return false
}

// SetUpdatedOn gets a reference to the given time.Time and assigns it to the UpdatedOn field.
func (o *Commitstatus) SetUpdatedOn(v time.Time) {
	o.UpdatedOn = &v
}

func (o Commitstatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedObject, errObject := json.Marshal(o.Object)
	if errObject != nil {
		return []byte{}, errObject
	}
	errObject = json.Unmarshal([]byte(serializedObject), &toSerialize)
	if errObject != nil {
		return []byte{}, errObject
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Refname != nil {
		toSerialize["refname"] = o.Refname
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
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

func (o *Commitstatus) UnmarshalJSON(bytes []byte) (err error) {
	type CommitstatusWithoutEmbeddedStruct struct {
		Links *CommitStatusLinks `json:"links,omitempty"`
		// The commit status' id.
		Uuid *string `json:"uuid,omitempty"`
		// An identifier for the status that's unique to         its type (current \"build\" is the only supported type) and the vendor,         e.g. BB-DEPLOY
		Key *string `json:"key,omitempty"`
		//  The name of the ref that pointed to this commit at the time the status object was created. Note that this the ref may since have moved off of the commit. This optional field can be useful for build systems whose build triggers and configuration are branch-dependent (e.g. a Pipeline build). It is legitimate for this field to not be set, or even apply (e.g. a static linting job).
		Refname *string `json:"refname,omitempty"`
		// A URL linking back to the vendor or build system, for providing more information about whatever process produced this status. Accepts context variables `repository` and `commit` that Bitbucket will evaluate at runtime whenever at runtime. For example, one could use https://foo.com/builds/{repository.full_name} which Bitbucket will turn into https://foo.com/builds/foo/bar at render time.
		Url *string `json:"url,omitempty"`
		// Provides some indication of the status of this commit
		State *string `json:"state,omitempty"`
		// An identifier for the build itself, e.g. BB-DEPLOY-1
		Name *string `json:"name,omitempty"`
		// A description of the build (e.g. \"Unit tests in Bamboo\")
		Description *string `json:"description,omitempty"`
		CreatedOn *time.Time `json:"created_on,omitempty"`
		UpdatedOn *time.Time `json:"updated_on,omitempty"`
	}

	varCommitstatusWithoutEmbeddedStruct := CommitstatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCommitstatusWithoutEmbeddedStruct)
	if err == nil {
		varCommitstatus := _Commitstatus{}
		varCommitstatus.Links = varCommitstatusWithoutEmbeddedStruct.Links
		varCommitstatus.Uuid = varCommitstatusWithoutEmbeddedStruct.Uuid
		varCommitstatus.Key = varCommitstatusWithoutEmbeddedStruct.Key
		varCommitstatus.Refname = varCommitstatusWithoutEmbeddedStruct.Refname
		varCommitstatus.Url = varCommitstatusWithoutEmbeddedStruct.Url
		varCommitstatus.State = varCommitstatusWithoutEmbeddedStruct.State
		varCommitstatus.Name = varCommitstatusWithoutEmbeddedStruct.Name
		varCommitstatus.Description = varCommitstatusWithoutEmbeddedStruct.Description
		varCommitstatus.CreatedOn = varCommitstatusWithoutEmbeddedStruct.CreatedOn
		varCommitstatus.UpdatedOn = varCommitstatusWithoutEmbeddedStruct.UpdatedOn
		*o = Commitstatus(varCommitstatus)
	} else {
		return err
	}

	varCommitstatus := _Commitstatus{}

	err = json.Unmarshal(bytes, &varCommitstatus)
	if err == nil {
		o.Object = varCommitstatus.Object
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "links")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "key")
		delete(additionalProperties, "refname")
		delete(additionalProperties, "url")
		delete(additionalProperties, "state")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "created_on")
		delete(additionalProperties, "updated_on")

		// remove fields from embedded structs
		reflectObject := reflect.ValueOf(o.Object)
		for i := 0; i < reflectObject.Type().NumField(); i++ {
			t := reflectObject.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommitstatus struct {
	value *Commitstatus
	isSet bool
}

func (v NullableCommitstatus) Get() *Commitstatus {
	return v.value
}

func (v *NullableCommitstatus) Set(val *Commitstatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCommitstatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCommitstatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommitstatus(val *Commitstatus) *NullableCommitstatus {
	return &NullableCommitstatus{value: val, isSet: true}
}

func (v NullableCommitstatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommitstatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


