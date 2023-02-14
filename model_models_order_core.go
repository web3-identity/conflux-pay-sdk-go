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

// ModelsOrderCore struct for ModelsOrderCore
type ModelsOrderCore struct {
	// 单位为分
	Amount *int32 `json:"amount,omitempty"`
	AppName *string `json:"app_name,omitempty"`
	CodeUrl *string `json:"code_url,omitempty"`
	Description *string `json:"description,omitempty"`
	H5Url *string `json:"h5_url,omitempty"`
	// 二维码宽度。 只有alipay，且 trade type 为 h5 模式有效，qr pay mode 为4 时有效； 用法参考 https://opendocs.alipay.com/apis/api_1/alipay.trade.page.pay?scene=22
	QrCodeWidth *string `json:"qr_code_width,omitempty"`
	// 支付二维码模式。 只有alipay，且 trade type 为 h5 模式有效; 用法参考 https://opendocs.alipay.com/apis/api_1/alipay.trade.page.pay?scene=22
	QrPayMode *string `json:"qr_pay_mode,omitempty"`
	RefundState *string `json:"refund_state,omitempty"`
	// 付款成功后的跳转链接。只有alipay，且 trade type 为 h5 模式有效; 用法参考 https://opendocs.alipay.com/apis/api_1/alipay.trade.page.pay?scene=22
	ReturnUrl *string `json:"return_url,omitempty"`
	TimeExpire *string `json:"time_expire,omitempty"`
	TradeNo *string `json:"trade_no,omitempty"`
	TradeProvider *string `json:"trade_provider,omitempty"`
	TradeState *string `json:"trade_state,omitempty"`
	TradeType *string `json:"trade_type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsOrderCore ModelsOrderCore

// NewModelsOrderCore instantiates a new ModelsOrderCore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsOrderCore() *ModelsOrderCore {
	this := ModelsOrderCore{}
	return &this
}

// NewModelsOrderCoreWithDefaults instantiates a new ModelsOrderCore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsOrderCoreWithDefaults() *ModelsOrderCore {
	this := ModelsOrderCore{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ModelsOrderCore) GetAmount() int32 {
	if o == nil || isNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrderCore) GetAmountOk() (*int32, bool) {
	if o == nil || isNil(o.Amount) {
    return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ModelsOrderCore) HasAmount() bool {
	if o != nil && !isNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *ModelsOrderCore) SetAmount(v int32) {
	o.Amount = &v
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *ModelsOrderCore) GetAppName() string {
	if o == nil || isNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrderCore) GetAppNameOk() (*string, bool) {
	if o == nil || isNil(o.AppName) {
    return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *ModelsOrderCore) HasAppName() bool {
	if o != nil && !isNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *ModelsOrderCore) SetAppName(v string) {
	o.AppName = &v
}

// GetCodeUrl returns the CodeUrl field value if set, zero value otherwise.
func (o *ModelsOrderCore) GetCodeUrl() string {
	if o == nil || isNil(o.CodeUrl) {
		var ret string
		return ret
	}
	return *o.CodeUrl
}

// GetCodeUrlOk returns a tuple with the CodeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrderCore) GetCodeUrlOk() (*string, bool) {
	if o == nil || isNil(o.CodeUrl) {
    return nil, false
	}
	return o.CodeUrl, true
}

// HasCodeUrl returns a boolean if a field has been set.
func (o *ModelsOrderCore) HasCodeUrl() bool {
	if o != nil && !isNil(o.CodeUrl) {
		return true
	}

	return false
}

// SetCodeUrl gets a reference to the given string and assigns it to the CodeUrl field.
func (o *ModelsOrderCore) SetCodeUrl(v string) {
	o.CodeUrl = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModelsOrderCore) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrderCore) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModelsOrderCore) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModelsOrderCore) SetDescription(v string) {
	o.Description = &v
}

// GetH5Url returns the H5Url field value if set, zero value otherwise.
func (o *ModelsOrderCore) GetH5Url() string {
	if o == nil || isNil(o.H5Url) {
		var ret string
		return ret
	}
	return *o.H5Url
}

// GetH5UrlOk returns a tuple with the H5Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrderCore) GetH5UrlOk() (*string, bool) {
	if o == nil || isNil(o.H5Url) {
    return nil, false
	}
	return o.H5Url, true
}

// HasH5Url returns a boolean if a field has been set.
func (o *ModelsOrderCore) HasH5Url() bool {
	if o != nil && !isNil(o.H5Url) {
		return true
	}

	return false
}

// SetH5Url gets a reference to the given string and assigns it to the H5Url field.
func (o *ModelsOrderCore) SetH5Url(v string) {
	o.H5Url = &v
}

// GetQrCodeWidth returns the QrCodeWidth field value if set, zero value otherwise.
func (o *ModelsOrderCore) GetQrCodeWidth() string {
	if o == nil || isNil(o.QrCodeWidth) {
		var ret string
		return ret
	}
	return *o.QrCodeWidth
}

// GetQrCodeWidthOk returns a tuple with the QrCodeWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrderCore) GetQrCodeWidthOk() (*string, bool) {
	if o == nil || isNil(o.QrCodeWidth) {
    return nil, false
	}
	return o.QrCodeWidth, true
}

// HasQrCodeWidth returns a boolean if a field has been set.
func (o *ModelsOrderCore) HasQrCodeWidth() bool {
	if o != nil && !isNil(o.QrCodeWidth) {
		return true
	}

	return false
}

// SetQrCodeWidth gets a reference to the given string and assigns it to the QrCodeWidth field.
func (o *ModelsOrderCore) SetQrCodeWidth(v string) {
	o.QrCodeWidth = &v
}

// GetQrPayMode returns the QrPayMode field value if set, zero value otherwise.
func (o *ModelsOrderCore) GetQrPayMode() string {
	if o == nil || isNil(o.QrPayMode) {
		var ret string
		return ret
	}
	return *o.QrPayMode
}

// GetQrPayModeOk returns a tuple with the QrPayMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrderCore) GetQrPayModeOk() (*string, bool) {
	if o == nil || isNil(o.QrPayMode) {
    return nil, false
	}
	return o.QrPayMode, true
}

// HasQrPayMode returns a boolean if a field has been set.
func (o *ModelsOrderCore) HasQrPayMode() bool {
	if o != nil && !isNil(o.QrPayMode) {
		return true
	}

	return false
}

// SetQrPayMode gets a reference to the given string and assigns it to the QrPayMode field.
func (o *ModelsOrderCore) SetQrPayMode(v string) {
	o.QrPayMode = &v
}

// GetRefundState returns the RefundState field value if set, zero value otherwise.
func (o *ModelsOrderCore) GetRefundState() string {
	if o == nil || isNil(o.RefundState) {
		var ret string
		return ret
	}
	return *o.RefundState
}

// GetRefundStateOk returns a tuple with the RefundState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrderCore) GetRefundStateOk() (*string, bool) {
	if o == nil || isNil(o.RefundState) {
    return nil, false
	}
	return o.RefundState, true
}

// HasRefundState returns a boolean if a field has been set.
func (o *ModelsOrderCore) HasRefundState() bool {
	if o != nil && !isNil(o.RefundState) {
		return true
	}

	return false
}

// SetRefundState gets a reference to the given string and assigns it to the RefundState field.
func (o *ModelsOrderCore) SetRefundState(v string) {
	o.RefundState = &v
}

// GetReturnUrl returns the ReturnUrl field value if set, zero value otherwise.
func (o *ModelsOrderCore) GetReturnUrl() string {
	if o == nil || isNil(o.ReturnUrl) {
		var ret string
		return ret
	}
	return *o.ReturnUrl
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrderCore) GetReturnUrlOk() (*string, bool) {
	if o == nil || isNil(o.ReturnUrl) {
    return nil, false
	}
	return o.ReturnUrl, true
}

// HasReturnUrl returns a boolean if a field has been set.
func (o *ModelsOrderCore) HasReturnUrl() bool {
	if o != nil && !isNil(o.ReturnUrl) {
		return true
	}

	return false
}

// SetReturnUrl gets a reference to the given string and assigns it to the ReturnUrl field.
func (o *ModelsOrderCore) SetReturnUrl(v string) {
	o.ReturnUrl = &v
}

// GetTimeExpire returns the TimeExpire field value if set, zero value otherwise.
func (o *ModelsOrderCore) GetTimeExpire() string {
	if o == nil || isNil(o.TimeExpire) {
		var ret string
		return ret
	}
	return *o.TimeExpire
}

// GetTimeExpireOk returns a tuple with the TimeExpire field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrderCore) GetTimeExpireOk() (*string, bool) {
	if o == nil || isNil(o.TimeExpire) {
    return nil, false
	}
	return o.TimeExpire, true
}

// HasTimeExpire returns a boolean if a field has been set.
func (o *ModelsOrderCore) HasTimeExpire() bool {
	if o != nil && !isNil(o.TimeExpire) {
		return true
	}

	return false
}

// SetTimeExpire gets a reference to the given string and assigns it to the TimeExpire field.
func (o *ModelsOrderCore) SetTimeExpire(v string) {
	o.TimeExpire = &v
}

// GetTradeNo returns the TradeNo field value if set, zero value otherwise.
func (o *ModelsOrderCore) GetTradeNo() string {
	if o == nil || isNil(o.TradeNo) {
		var ret string
		return ret
	}
	return *o.TradeNo
}

// GetTradeNoOk returns a tuple with the TradeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrderCore) GetTradeNoOk() (*string, bool) {
	if o == nil || isNil(o.TradeNo) {
    return nil, false
	}
	return o.TradeNo, true
}

// HasTradeNo returns a boolean if a field has been set.
func (o *ModelsOrderCore) HasTradeNo() bool {
	if o != nil && !isNil(o.TradeNo) {
		return true
	}

	return false
}

// SetTradeNo gets a reference to the given string and assigns it to the TradeNo field.
func (o *ModelsOrderCore) SetTradeNo(v string) {
	o.TradeNo = &v
}

// GetTradeProvider returns the TradeProvider field value if set, zero value otherwise.
func (o *ModelsOrderCore) GetTradeProvider() string {
	if o == nil || isNil(o.TradeProvider) {
		var ret string
		return ret
	}
	return *o.TradeProvider
}

// GetTradeProviderOk returns a tuple with the TradeProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrderCore) GetTradeProviderOk() (*string, bool) {
	if o == nil || isNil(o.TradeProvider) {
    return nil, false
	}
	return o.TradeProvider, true
}

// HasTradeProvider returns a boolean if a field has been set.
func (o *ModelsOrderCore) HasTradeProvider() bool {
	if o != nil && !isNil(o.TradeProvider) {
		return true
	}

	return false
}

// SetTradeProvider gets a reference to the given string and assigns it to the TradeProvider field.
func (o *ModelsOrderCore) SetTradeProvider(v string) {
	o.TradeProvider = &v
}

// GetTradeState returns the TradeState field value if set, zero value otherwise.
func (o *ModelsOrderCore) GetTradeState() string {
	if o == nil || isNil(o.TradeState) {
		var ret string
		return ret
	}
	return *o.TradeState
}

// GetTradeStateOk returns a tuple with the TradeState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrderCore) GetTradeStateOk() (*string, bool) {
	if o == nil || isNil(o.TradeState) {
    return nil, false
	}
	return o.TradeState, true
}

// HasTradeState returns a boolean if a field has been set.
func (o *ModelsOrderCore) HasTradeState() bool {
	if o != nil && !isNil(o.TradeState) {
		return true
	}

	return false
}

// SetTradeState gets a reference to the given string and assigns it to the TradeState field.
func (o *ModelsOrderCore) SetTradeState(v string) {
	o.TradeState = &v
}

// GetTradeType returns the TradeType field value if set, zero value otherwise.
func (o *ModelsOrderCore) GetTradeType() string {
	if o == nil || isNil(o.TradeType) {
		var ret string
		return ret
	}
	return *o.TradeType
}

// GetTradeTypeOk returns a tuple with the TradeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrderCore) GetTradeTypeOk() (*string, bool) {
	if o == nil || isNil(o.TradeType) {
    return nil, false
	}
	return o.TradeType, true
}

// HasTradeType returns a boolean if a field has been set.
func (o *ModelsOrderCore) HasTradeType() bool {
	if o != nil && !isNil(o.TradeType) {
		return true
	}

	return false
}

// SetTradeType gets a reference to the given string and assigns it to the TradeType field.
func (o *ModelsOrderCore) SetTradeType(v string) {
	o.TradeType = &v
}

func (o ModelsOrderCore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !isNil(o.AppName) {
		toSerialize["app_name"] = o.AppName
	}
	if !isNil(o.CodeUrl) {
		toSerialize["code_url"] = o.CodeUrl
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.H5Url) {
		toSerialize["h5_url"] = o.H5Url
	}
	if !isNil(o.QrCodeWidth) {
		toSerialize["qr_code_width"] = o.QrCodeWidth
	}
	if !isNil(o.QrPayMode) {
		toSerialize["qr_pay_mode"] = o.QrPayMode
	}
	if !isNil(o.RefundState) {
		toSerialize["refund_state"] = o.RefundState
	}
	if !isNil(o.ReturnUrl) {
		toSerialize["return_url"] = o.ReturnUrl
	}
	if !isNil(o.TimeExpire) {
		toSerialize["time_expire"] = o.TimeExpire
	}
	if !isNil(o.TradeNo) {
		toSerialize["trade_no"] = o.TradeNo
	}
	if !isNil(o.TradeProvider) {
		toSerialize["trade_provider"] = o.TradeProvider
	}
	if !isNil(o.TradeState) {
		toSerialize["trade_state"] = o.TradeState
	}
	if !isNil(o.TradeType) {
		toSerialize["trade_type"] = o.TradeType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ModelsOrderCore) UnmarshalJSON(bytes []byte) (err error) {
	varModelsOrderCore := _ModelsOrderCore{}

	if err = json.Unmarshal(bytes, &varModelsOrderCore); err == nil {
		*o = ModelsOrderCore(varModelsOrderCore)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "app_name")
		delete(additionalProperties, "code_url")
		delete(additionalProperties, "description")
		delete(additionalProperties, "h5_url")
		delete(additionalProperties, "qr_code_width")
		delete(additionalProperties, "qr_pay_mode")
		delete(additionalProperties, "refund_state")
		delete(additionalProperties, "return_url")
		delete(additionalProperties, "time_expire")
		delete(additionalProperties, "trade_no")
		delete(additionalProperties, "trade_provider")
		delete(additionalProperties, "trade_state")
		delete(additionalProperties, "trade_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsOrderCore struct {
	value *ModelsOrderCore
	isSet bool
}

func (v NullableModelsOrderCore) Get() *ModelsOrderCore {
	return v.value
}

func (v *NullableModelsOrderCore) Set(val *ModelsOrderCore) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsOrderCore) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsOrderCore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsOrderCore(val *ModelsOrderCore) *NullableModelsOrderCore {
	return &NullableModelsOrderCore{value: val, isSet: true}
}

func (v NullableModelsOrderCore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsOrderCore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


