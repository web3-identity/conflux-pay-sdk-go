/*
Rainbow-API

Conflux-Pay API documentation

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package confluxpay

import (
	"encoding/json"
)

// checks if the ControllersAddUnitAccountResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ControllersAddUnitAccountResult{}

// ControllersAddUnitAccountResult struct for ControllersAddUnitAccountResult
type ControllersAddUnitAccountResult struct {
	FullUnitAccountNbr string `json:"full_unit_account_nbr"`
	AdditionalProperties map[string]interface{}
}

type _ControllersAddUnitAccountResult ControllersAddUnitAccountResult

// NewControllersAddUnitAccountResult instantiates a new ControllersAddUnitAccountResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersAddUnitAccountResult(fullUnitAccountNbr string) *ControllersAddUnitAccountResult {
	this := ControllersAddUnitAccountResult{}
	this.FullUnitAccountNbr = fullUnitAccountNbr
	return &this
}

// NewControllersAddUnitAccountResultWithDefaults instantiates a new ControllersAddUnitAccountResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersAddUnitAccountResultWithDefaults() *ControllersAddUnitAccountResult {
	this := ControllersAddUnitAccountResult{}
	return &this
}

// GetFullUnitAccountNbr returns the FullUnitAccountNbr field value
func (o *ControllersAddUnitAccountResult) GetFullUnitAccountNbr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullUnitAccountNbr
}

// GetFullUnitAccountNbrOk returns a tuple with the FullUnitAccountNbr field value
// and a boolean to check if the value has been set.
func (o *ControllersAddUnitAccountResult) GetFullUnitAccountNbrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullUnitAccountNbr, true
}

// SetFullUnitAccountNbr sets field value
func (o *ControllersAddUnitAccountResult) SetFullUnitAccountNbr(v string) {
	o.FullUnitAccountNbr = v
}

func (o ControllersAddUnitAccountResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ControllersAddUnitAccountResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["full_unit_account_nbr"] = o.FullUnitAccountNbr

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ControllersAddUnitAccountResult) UnmarshalJSON(bytes []byte) (err error) {
	varControllersAddUnitAccountResult := _ControllersAddUnitAccountResult{}

	if err = json.Unmarshal(bytes, &varControllersAddUnitAccountResult); err == nil {
		*o = ControllersAddUnitAccountResult(varControllersAddUnitAccountResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "full_unit_account_nbr")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableControllersAddUnitAccountResult struct {
	value *ControllersAddUnitAccountResult
	isSet bool
}

func (v NullableControllersAddUnitAccountResult) Get() *ControllersAddUnitAccountResult {
	return v.value
}

func (v *NullableControllersAddUnitAccountResult) Set(val *ControllersAddUnitAccountResult) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersAddUnitAccountResult) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersAddUnitAccountResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersAddUnitAccountResult(val *ControllersAddUnitAccountResult) *NullableControllersAddUnitAccountResult {
	return &NullableControllersAddUnitAccountResult{value: val, isSet: true}
}

func (v NullableControllersAddUnitAccountResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersAddUnitAccountResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


