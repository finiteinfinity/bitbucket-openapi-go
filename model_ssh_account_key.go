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
	"reflect"
	"strings"
)

// SshAccountKey struct for SshAccountKey
type SshAccountKey struct {
	SshKey
	Owner *Account `json:"owner,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SshAccountKey SshAccountKey

// NewSshAccountKey instantiates a new SshAccountKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSshAccountKey() *SshAccountKey {
	this := SshAccountKey{}
	return &this
}

// NewSshAccountKeyWithDefaults instantiates a new SshAccountKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshAccountKeyWithDefaults() *SshAccountKey {
	this := SshAccountKey{}
	return &this
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *SshAccountKey) GetOwner() Account {
	if o == nil || o.Owner == nil {
		var ret Account
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshAccountKey) GetOwnerOk() (*Account, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *SshAccountKey) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given Account and assigns it to the Owner field.
func (o *SshAccountKey) SetOwner(v Account) {
	o.Owner = &v
}

func (o SshAccountKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedSshKey, errSshKey := json.Marshal(o.SshKey)
	if errSshKey != nil {
		return []byte{}, errSshKey
	}
	errSshKey = json.Unmarshal([]byte(serializedSshKey), &toSerialize)
	if errSshKey != nil {
		return []byte{}, errSshKey
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SshAccountKey) UnmarshalJSON(bytes []byte) (err error) {
	type SshAccountKeyWithoutEmbeddedStruct struct {
		Owner *Account `json:"owner,omitempty"`
	}

	varSshAccountKeyWithoutEmbeddedStruct := SshAccountKeyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSshAccountKeyWithoutEmbeddedStruct)
	if err == nil {
		varSshAccountKey := _SshAccountKey{}
		varSshAccountKey.Owner = varSshAccountKeyWithoutEmbeddedStruct.Owner
		*o = SshAccountKey(varSshAccountKey)
	} else {
		return err
	}

	varSshAccountKey := _SshAccountKey{}

	err = json.Unmarshal(bytes, &varSshAccountKey)
	if err == nil {
		o.SshKey = varSshAccountKey.SshKey
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "owner")

		// remove fields from embedded structs
		reflectSshKey := reflect.ValueOf(o.SshKey)
		for i := 0; i < reflectSshKey.Type().NumField(); i++ {
			t := reflectSshKey.Type().Field(i)

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

type NullableSshAccountKey struct {
	value *SshAccountKey
	isSet bool
}

func (v NullableSshAccountKey) Get() *SshAccountKey {
	return v.value
}

func (v *NullableSshAccountKey) Set(val *SshAccountKey) {
	v.value = val
	v.isSet = true
}

func (v NullableSshAccountKey) IsSet() bool {
	return v.isSet
}

func (v *NullableSshAccountKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSshAccountKey(val *SshAccountKey) *NullableSshAccountKey {
	return &NullableSshAccountKey{value: val, isSet: true}
}

func (v NullableSshAccountKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSshAccountKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


