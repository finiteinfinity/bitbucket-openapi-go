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

// BranchingModel A repository's branching model
type BranchingModel struct {
	// The active branch types.
	BranchTypes []BranchingModelBranchTypes `json:"branch_types,omitempty"`
	Development *BranchingModelDevelopment `json:"development,omitempty"`
	Production *BranchingModelDevelopment `json:"production,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BranchingModel BranchingModel

// NewBranchingModel instantiates a new BranchingModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBranchingModel() *BranchingModel {
	this := BranchingModel{}
	return &this
}

// NewBranchingModelWithDefaults instantiates a new BranchingModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBranchingModelWithDefaults() *BranchingModel {
	this := BranchingModel{}
	return &this
}

// GetBranchTypes returns the BranchTypes field value if set, zero value otherwise.
func (o *BranchingModel) GetBranchTypes() []BranchingModelBranchTypes {
	if o == nil || o.BranchTypes == nil {
		var ret []BranchingModelBranchTypes
		return ret
	}
	return o.BranchTypes
}

// GetBranchTypesOk returns a tuple with the BranchTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchingModel) GetBranchTypesOk() ([]BranchingModelBranchTypes, bool) {
	if o == nil || o.BranchTypes == nil {
		return nil, false
	}
	return o.BranchTypes, true
}

// HasBranchTypes returns a boolean if a field has been set.
func (o *BranchingModel) HasBranchTypes() bool {
	if o != nil && o.BranchTypes != nil {
		return true
	}

	return false
}

// SetBranchTypes gets a reference to the given []BranchingModelBranchTypes and assigns it to the BranchTypes field.
func (o *BranchingModel) SetBranchTypes(v []BranchingModelBranchTypes) {
	o.BranchTypes = v
}

// GetDevelopment returns the Development field value if set, zero value otherwise.
func (o *BranchingModel) GetDevelopment() BranchingModelDevelopment {
	if o == nil || o.Development == nil {
		var ret BranchingModelDevelopment
		return ret
	}
	return *o.Development
}

// GetDevelopmentOk returns a tuple with the Development field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchingModel) GetDevelopmentOk() (*BranchingModelDevelopment, bool) {
	if o == nil || o.Development == nil {
		return nil, false
	}
	return o.Development, true
}

// HasDevelopment returns a boolean if a field has been set.
func (o *BranchingModel) HasDevelopment() bool {
	if o != nil && o.Development != nil {
		return true
	}

	return false
}

// SetDevelopment gets a reference to the given BranchingModelDevelopment and assigns it to the Development field.
func (o *BranchingModel) SetDevelopment(v BranchingModelDevelopment) {
	o.Development = &v
}

// GetProduction returns the Production field value if set, zero value otherwise.
func (o *BranchingModel) GetProduction() BranchingModelDevelopment {
	if o == nil || o.Production == nil {
		var ret BranchingModelDevelopment
		return ret
	}
	return *o.Production
}

// GetProductionOk returns a tuple with the Production field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchingModel) GetProductionOk() (*BranchingModelDevelopment, bool) {
	if o == nil || o.Production == nil {
		return nil, false
	}
	return o.Production, true
}

// HasProduction returns a boolean if a field has been set.
func (o *BranchingModel) HasProduction() bool {
	if o != nil && o.Production != nil {
		return true
	}

	return false
}

// SetProduction gets a reference to the given BranchingModelDevelopment and assigns it to the Production field.
func (o *BranchingModel) SetProduction(v BranchingModelDevelopment) {
	o.Production = &v
}

func (o BranchingModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BranchTypes != nil {
		toSerialize["branch_types"] = o.BranchTypes
	}
	if o.Development != nil {
		toSerialize["development"] = o.Development
	}
	if o.Production != nil {
		toSerialize["production"] = o.Production
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BranchingModel) UnmarshalJSON(bytes []byte) (err error) {
	varBranchingModel := _BranchingModel{}

	if err = json.Unmarshal(bytes, &varBranchingModel); err == nil {
		*o = BranchingModel(varBranchingModel)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "branch_types")
		delete(additionalProperties, "development")
		delete(additionalProperties, "production")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBranchingModel struct {
	value *BranchingModel
	isSet bool
}

func (v NullableBranchingModel) Get() *BranchingModel {
	return v.value
}

func (v *NullableBranchingModel) Set(val *BranchingModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBranchingModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBranchingModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBranchingModel(val *BranchingModel) *NullableBranchingModel {
	return &NullableBranchingModel{value: val, isSet: true}
}

func (v NullableBranchingModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBranchingModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


