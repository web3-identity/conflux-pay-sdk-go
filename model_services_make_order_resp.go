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

// ServicesMakeOrderResp struct for ServicesMakeOrderResp
type ServicesMakeOrderResp struct {
	CodeUrl *string `json:"code_url,omitempty"`
	H5Url *string `json:"h5_url,omitempty"`
	TradeNo *string `json:"trade_no,omitempty"`
	TradeProvider *string `json:"trade_provider,omitempty"`
	TradeType *string `json:"trade_type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServicesMakeOrderResp ServicesMakeOrderResp

// NewServicesMakeOrderResp instantiates a new ServicesMakeOrderResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicesMakeOrderResp() *ServicesMakeOrderResp {
	this := ServicesMakeOrderResp{}
	return &this
}

// NewServicesMakeOrderRespWithDefaults instantiates a new ServicesMakeOrderResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesMakeOrderRespWithDefaults() *ServicesMakeOrderResp {
	this := ServicesMakeOrderResp{}
	return &this
}

// GetCodeUrl returns the CodeUrl field value if set, zero value otherwise.
func (o *ServicesMakeOrderResp) GetCodeUrl() string {
	if o == nil || o.CodeUrl == nil {
		var ret string
		return ret
	}
	return *o.CodeUrl
}

// GetCodeUrlOk returns a tuple with the CodeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesMakeOrderResp) GetCodeUrlOk() (*string, bool) {
	if o == nil || o.CodeUrl == nil {
		return nil, false
	}
	return o.CodeUrl, true
}

// HasCodeUrl returns a boolean if a field has been set.
func (o *ServicesMakeOrderResp) HasCodeUrl() bool {
	if o != nil && o.CodeUrl != nil {
		return true
	}

	return false
}

// SetCodeUrl gets a reference to the given string and assigns it to the CodeUrl field.
func (o *ServicesMakeOrderResp) SetCodeUrl(v string) {
	o.CodeUrl = &v
}

// GetH5Url returns the H5Url field value if set, zero value otherwise.
func (o *ServicesMakeOrderResp) GetH5Url() string {
	if o == nil || o.H5Url == nil {
		var ret string
		return ret
	}
	return *o.H5Url
}

// GetH5UrlOk returns a tuple with the H5Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesMakeOrderResp) GetH5UrlOk() (*string, bool) {
	if o == nil || o.H5Url == nil {
		return nil, false
	}
	return o.H5Url, true
}

// HasH5Url returns a boolean if a field has been set.
func (o *ServicesMakeOrderResp) HasH5Url() bool {
	if o != nil && o.H5Url != nil {
		return true
	}

	return false
}

// SetH5Url gets a reference to the given string and assigns it to the H5Url field.
func (o *ServicesMakeOrderResp) SetH5Url(v string) {
	o.H5Url = &v
}

// GetTradeNo returns the TradeNo field value if set, zero value otherwise.
func (o *ServicesMakeOrderResp) GetTradeNo() string {
	if o == nil || o.TradeNo == nil {
		var ret string
		return ret
	}
	return *o.TradeNo
}

// GetTradeNoOk returns a tuple with the TradeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesMakeOrderResp) GetTradeNoOk() (*string, bool) {
	if o == nil || o.TradeNo == nil {
		return nil, false
	}
	return o.TradeNo, true
}

// HasTradeNo returns a boolean if a field has been set.
func (o *ServicesMakeOrderResp) HasTradeNo() bool {
	if o != nil && o.TradeNo != nil {
		return true
	}

	return false
}

// SetTradeNo gets a reference to the given string and assigns it to the TradeNo field.
func (o *ServicesMakeOrderResp) SetTradeNo(v string) {
	o.TradeNo = &v
}

// GetTradeProvider returns the TradeProvider field value if set, zero value otherwise.
func (o *ServicesMakeOrderResp) GetTradeProvider() string {
	if o == nil || o.TradeProvider == nil {
		var ret string
		return ret
	}
	return *o.TradeProvider
}

// GetTradeProviderOk returns a tuple with the TradeProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesMakeOrderResp) GetTradeProviderOk() (*string, bool) {
	if o == nil || o.TradeProvider == nil {
		return nil, false
	}
	return o.TradeProvider, true
}

// HasTradeProvider returns a boolean if a field has been set.
func (o *ServicesMakeOrderResp) HasTradeProvider() bool {
	if o != nil && o.TradeProvider != nil {
		return true
	}

	return false
}

// SetTradeProvider gets a reference to the given string and assigns it to the TradeProvider field.
func (o *ServicesMakeOrderResp) SetTradeProvider(v string) {
	o.TradeProvider = &v
}

// GetTradeType returns the TradeType field value if set, zero value otherwise.
func (o *ServicesMakeOrderResp) GetTradeType() string {
	if o == nil || o.TradeType == nil {
		var ret string
		return ret
	}
	return *o.TradeType
}

// GetTradeTypeOk returns a tuple with the TradeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesMakeOrderResp) GetTradeTypeOk() (*string, bool) {
	if o == nil || o.TradeType == nil {
		return nil, false
	}
	return o.TradeType, true
}

// HasTradeType returns a boolean if a field has been set.
func (o *ServicesMakeOrderResp) HasTradeType() bool {
	if o != nil && o.TradeType != nil {
		return true
	}

	return false
}

// SetTradeType gets a reference to the given string and assigns it to the TradeType field.
func (o *ServicesMakeOrderResp) SetTradeType(v string) {
	o.TradeType = &v
}

func (o ServicesMakeOrderResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CodeUrl != nil {
		toSerialize["code_url"] = o.CodeUrl
	}
	if o.H5Url != nil {
		toSerialize["h5_url"] = o.H5Url
	}
	if o.TradeNo != nil {
		toSerialize["trade_no"] = o.TradeNo
	}
	if o.TradeProvider != nil {
		toSerialize["trade_provider"] = o.TradeProvider
	}
	if o.TradeType != nil {
		toSerialize["trade_type"] = o.TradeType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServicesMakeOrderResp) UnmarshalJSON(bytes []byte) (err error) {
	varServicesMakeOrderResp := _ServicesMakeOrderResp{}

	if err = json.Unmarshal(bytes, &varServicesMakeOrderResp); err == nil {
		*o = ServicesMakeOrderResp(varServicesMakeOrderResp)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "code_url")
		delete(additionalProperties, "h5_url")
		delete(additionalProperties, "trade_no")
		delete(additionalProperties, "trade_provider")
		delete(additionalProperties, "trade_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServicesMakeOrderResp struct {
	value *ServicesMakeOrderResp
	isSet bool
}

func (v NullableServicesMakeOrderResp) Get() *ServicesMakeOrderResp {
	return v.value
}

func (v *NullableServicesMakeOrderResp) Set(val *ServicesMakeOrderResp) {
	v.value = val
	v.isSet = true
}

func (v NullableServicesMakeOrderResp) IsSet() bool {
	return v.isSet
}

func (v *NullableServicesMakeOrderResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicesMakeOrderResp(val *ServicesMakeOrderResp) *NullableServicesMakeOrderResp {
	return &NullableServicesMakeOrderResp{value: val, isSet: true}
}

func (v NullableServicesMakeOrderResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicesMakeOrderResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


