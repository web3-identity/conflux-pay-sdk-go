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

// checks if the ServicesMakeOrderReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServicesMakeOrderReq{}

// ServicesMakeOrderReq struct for ServicesMakeOrderReq
type ServicesMakeOrderReq struct {
	Amount int32 `json:"amount"`
	AppName string `json:"app_name"`
	Description string `json:"description"`
	NotifyUrl *string `json:"notify_url,omitempty"`
	// 二维码宽度。 只有alipay，且 trade type 为 h5 模式有效，qr pay mode 为4 时有效； 用法参考 https://opendocs.alipay.com/apis/api_1/alipay.trade.page.pay?scene=22
	QrCodeWidth *string `json:"qr_code_width,omitempty"`
	// 支付二维码模式。 只有alipay，且 trade type 为 h5 模式有效; 用法参考 https://opendocs.alipay.com/apis/api_1/alipay.trade.page.pay?scene=22
	QrPayMode *string `json:"qr_pay_mode,omitempty"`
	// 付款成功后的跳转链接。只有alipay，且 trade type 为 h5 模式有效; 用法参考 https://opendocs.alipay.com/apis/api_1/alipay.trade.page.pay?scene=22
	ReturnUrl *string `json:"return_url,omitempty"`
	// alipay 当面付无效，当面付固定过期时间为2小时
	TimeExpire int32 `json:"time_expire"`
	TradeProvider string `json:"trade_provider"`
	TradeType string `json:"trade_type"`
	AdditionalProperties map[string]interface{}
}

type _ServicesMakeOrderReq ServicesMakeOrderReq

// NewServicesMakeOrderReq instantiates a new ServicesMakeOrderReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicesMakeOrderReq(amount int32, appName string, description string, timeExpire int32, tradeProvider string, tradeType string) *ServicesMakeOrderReq {
	this := ServicesMakeOrderReq{}
	this.Amount = amount
	this.AppName = appName
	this.Description = description
	this.TimeExpire = timeExpire
	this.TradeProvider = tradeProvider
	this.TradeType = tradeType
	return &this
}

// NewServicesMakeOrderReqWithDefaults instantiates a new ServicesMakeOrderReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesMakeOrderReqWithDefaults() *ServicesMakeOrderReq {
	this := ServicesMakeOrderReq{}
	return &this
}

// GetAmount returns the Amount field value
func (o *ServicesMakeOrderReq) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ServicesMakeOrderReq) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ServicesMakeOrderReq) SetAmount(v int32) {
	o.Amount = v
}

// GetAppName returns the AppName field value
func (o *ServicesMakeOrderReq) GetAppName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value
// and a boolean to check if the value has been set.
func (o *ServicesMakeOrderReq) GetAppNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppName, true
}

// SetAppName sets field value
func (o *ServicesMakeOrderReq) SetAppName(v string) {
	o.AppName = v
}

// GetDescription returns the Description field value
func (o *ServicesMakeOrderReq) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ServicesMakeOrderReq) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ServicesMakeOrderReq) SetDescription(v string) {
	o.Description = v
}

// GetNotifyUrl returns the NotifyUrl field value if set, zero value otherwise.
func (o *ServicesMakeOrderReq) GetNotifyUrl() string {
	if o == nil || isNil(o.NotifyUrl) {
		var ret string
		return ret
	}
	return *o.NotifyUrl
}

// GetNotifyUrlOk returns a tuple with the NotifyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesMakeOrderReq) GetNotifyUrlOk() (*string, bool) {
	if o == nil || isNil(o.NotifyUrl) {
		return nil, false
	}
	return o.NotifyUrl, true
}

// HasNotifyUrl returns a boolean if a field has been set.
func (o *ServicesMakeOrderReq) HasNotifyUrl() bool {
	if o != nil && !isNil(o.NotifyUrl) {
		return true
	}

	return false
}

// SetNotifyUrl gets a reference to the given string and assigns it to the NotifyUrl field.
func (o *ServicesMakeOrderReq) SetNotifyUrl(v string) {
	o.NotifyUrl = &v
}

// GetQrCodeWidth returns the QrCodeWidth field value if set, zero value otherwise.
func (o *ServicesMakeOrderReq) GetQrCodeWidth() string {
	if o == nil || isNil(o.QrCodeWidth) {
		var ret string
		return ret
	}
	return *o.QrCodeWidth
}

// GetQrCodeWidthOk returns a tuple with the QrCodeWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesMakeOrderReq) GetQrCodeWidthOk() (*string, bool) {
	if o == nil || isNil(o.QrCodeWidth) {
		return nil, false
	}
	return o.QrCodeWidth, true
}

// HasQrCodeWidth returns a boolean if a field has been set.
func (o *ServicesMakeOrderReq) HasQrCodeWidth() bool {
	if o != nil && !isNil(o.QrCodeWidth) {
		return true
	}

	return false
}

// SetQrCodeWidth gets a reference to the given string and assigns it to the QrCodeWidth field.
func (o *ServicesMakeOrderReq) SetQrCodeWidth(v string) {
	o.QrCodeWidth = &v
}

// GetQrPayMode returns the QrPayMode field value if set, zero value otherwise.
func (o *ServicesMakeOrderReq) GetQrPayMode() string {
	if o == nil || isNil(o.QrPayMode) {
		var ret string
		return ret
	}
	return *o.QrPayMode
}

// GetQrPayModeOk returns a tuple with the QrPayMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesMakeOrderReq) GetQrPayModeOk() (*string, bool) {
	if o == nil || isNil(o.QrPayMode) {
		return nil, false
	}
	return o.QrPayMode, true
}

// HasQrPayMode returns a boolean if a field has been set.
func (o *ServicesMakeOrderReq) HasQrPayMode() bool {
	if o != nil && !isNil(o.QrPayMode) {
		return true
	}

	return false
}

// SetQrPayMode gets a reference to the given string and assigns it to the QrPayMode field.
func (o *ServicesMakeOrderReq) SetQrPayMode(v string) {
	o.QrPayMode = &v
}

// GetReturnUrl returns the ReturnUrl field value if set, zero value otherwise.
func (o *ServicesMakeOrderReq) GetReturnUrl() string {
	if o == nil || isNil(o.ReturnUrl) {
		var ret string
		return ret
	}
	return *o.ReturnUrl
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesMakeOrderReq) GetReturnUrlOk() (*string, bool) {
	if o == nil || isNil(o.ReturnUrl) {
		return nil, false
	}
	return o.ReturnUrl, true
}

// HasReturnUrl returns a boolean if a field has been set.
func (o *ServicesMakeOrderReq) HasReturnUrl() bool {
	if o != nil && !isNil(o.ReturnUrl) {
		return true
	}

	return false
}

// SetReturnUrl gets a reference to the given string and assigns it to the ReturnUrl field.
func (o *ServicesMakeOrderReq) SetReturnUrl(v string) {
	o.ReturnUrl = &v
}

// GetTimeExpire returns the TimeExpire field value
func (o *ServicesMakeOrderReq) GetTimeExpire() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TimeExpire
}

// GetTimeExpireOk returns a tuple with the TimeExpire field value
// and a boolean to check if the value has been set.
func (o *ServicesMakeOrderReq) GetTimeExpireOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeExpire, true
}

// SetTimeExpire sets field value
func (o *ServicesMakeOrderReq) SetTimeExpire(v int32) {
	o.TimeExpire = v
}

// GetTradeProvider returns the TradeProvider field value
func (o *ServicesMakeOrderReq) GetTradeProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TradeProvider
}

// GetTradeProviderOk returns a tuple with the TradeProvider field value
// and a boolean to check if the value has been set.
func (o *ServicesMakeOrderReq) GetTradeProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TradeProvider, true
}

// SetTradeProvider sets field value
func (o *ServicesMakeOrderReq) SetTradeProvider(v string) {
	o.TradeProvider = v
}

// GetTradeType returns the TradeType field value
func (o *ServicesMakeOrderReq) GetTradeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TradeType
}

// GetTradeTypeOk returns a tuple with the TradeType field value
// and a boolean to check if the value has been set.
func (o *ServicesMakeOrderReq) GetTradeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TradeType, true
}

// SetTradeType sets field value
func (o *ServicesMakeOrderReq) SetTradeType(v string) {
	o.TradeType = v
}

func (o ServicesMakeOrderReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServicesMakeOrderReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["app_name"] = o.AppName
	toSerialize["description"] = o.Description
	if !isNil(o.NotifyUrl) {
		toSerialize["notify_url"] = o.NotifyUrl
	}
	if !isNil(o.QrCodeWidth) {
		toSerialize["qr_code_width"] = o.QrCodeWidth
	}
	if !isNil(o.QrPayMode) {
		toSerialize["qr_pay_mode"] = o.QrPayMode
	}
	if !isNil(o.ReturnUrl) {
		toSerialize["return_url"] = o.ReturnUrl
	}
	toSerialize["time_expire"] = o.TimeExpire
	toSerialize["trade_provider"] = o.TradeProvider
	toSerialize["trade_type"] = o.TradeType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServicesMakeOrderReq) UnmarshalJSON(bytes []byte) (err error) {
	varServicesMakeOrderReq := _ServicesMakeOrderReq{}

	if err = json.Unmarshal(bytes, &varServicesMakeOrderReq); err == nil {
		*o = ServicesMakeOrderReq(varServicesMakeOrderReq)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "app_name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "notify_url")
		delete(additionalProperties, "qr_code_width")
		delete(additionalProperties, "qr_pay_mode")
		delete(additionalProperties, "return_url")
		delete(additionalProperties, "time_expire")
		delete(additionalProperties, "trade_provider")
		delete(additionalProperties, "trade_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServicesMakeOrderReq struct {
	value *ServicesMakeOrderReq
	isSet bool
}

func (v NullableServicesMakeOrderReq) Get() *ServicesMakeOrderReq {
	return v.value
}

func (v *NullableServicesMakeOrderReq) Set(val *ServicesMakeOrderReq) {
	v.value = val
	v.isSet = true
}

func (v NullableServicesMakeOrderReq) IsSet() bool {
	return v.isSet
}

func (v *NullableServicesMakeOrderReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicesMakeOrderReq(val *ServicesMakeOrderReq) *NullableServicesMakeOrderReq {
	return &NullableServicesMakeOrderReq{value: val, isSet: true}
}

func (v NullableServicesMakeOrderReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicesMakeOrderReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


