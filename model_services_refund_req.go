/*
Rainbow-API

The responses of the open api in swagger focus on the data field rather than the code and the message fields

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package confluxpay

import (
	"encoding/json"
)

// ServicesRefundReq struct for ServicesRefundReq
type ServicesRefundReq struct {
	Reason string `json:"reason"`
	AdditionalProperties map[string]interface{}
}

type _ServicesRefundReq ServicesRefundReq

// NewServicesRefundReq instantiates a new ServicesRefundReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicesRefundReq(reason string) *ServicesRefundReq {
	this := ServicesRefundReq{}
	this.Reason = reason
	return &this
}

// NewServicesRefundReqWithDefaults instantiates a new ServicesRefundReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesRefundReqWithDefaults() *ServicesRefundReq {
	this := ServicesRefundReq{}
	return &this
}

// GetReason returns the Reason field value
func (o *ServicesRefundReq) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *ServicesRefundReq) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *ServicesRefundReq) SetReason(v string) {
	o.Reason = v
}

func (o ServicesRefundReq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServicesRefundReq) UnmarshalJSON(bytes []byte) (err error) {
	varServicesRefundReq := _ServicesRefundReq{}

	if err = json.Unmarshal(bytes, &varServicesRefundReq); err == nil {
		*o = ServicesRefundReq(varServicesRefundReq)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "reason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServicesRefundReq struct {
	value *ServicesRefundReq
	isSet bool
}

func (v NullableServicesRefundReq) Get() *ServicesRefundReq {
	return v.value
}

func (v *NullableServicesRefundReq) Set(val *ServicesRefundReq) {
	v.value = val
	v.isSet = true
}

func (v NullableServicesRefundReq) IsSet() bool {
	return v.isSet
}

func (v *NullableServicesRefundReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicesRefundReq(val *ServicesRefundReq) *NullableServicesRefundReq {
	return &NullableServicesRefundReq{value: val, isSet: true}
}

func (v NullableServicesRefundReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicesRefundReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

