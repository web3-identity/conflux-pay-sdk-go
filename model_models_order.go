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

// ModelsOrder struct for ModelsOrder
type ModelsOrder struct {
	// 单位为分
	Amount *int32 `json:"amount,omitempty"`
	AppName *string `json:"app_name,omitempty"`
	// 上层应用通知url
	AppPayNotifyUrl *string `json:"app_pay_notify_url,omitempty"`
	// 上层应用通知url
	AppRefundNotifyUrl *string `json:"app_refund_notify_url,omitempty"`
	CodeUrl *string `json:"code_url,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	DeletedAt *GormDeletedAt `json:"deleted_at,omitempty"`
	Description *string `json:"description,omitempty"`
	H5Url *string `json:"h5_url,omitempty"`
	Id *int32 `json:"id,omitempty"`
	IsPayNotifyCompleted *bool `json:"is_pay_notify_completed,omitempty"`
	IsRefundNotifyCompleted *bool `json:"is_refund_notify_completed,omitempty"`
	// PayNotifyNextTime    *time.Time `json:\"pay_notify_next_time\"`
	PayNotifyCount *int32 `json:"pay_notify_count,omitempty"`
	// 二维码宽度。 只有alipay，且 trade type 为 h5 模式有效，qr pay mode 为4 时有效； 用法参考 https://opendocs.alipay.com/apis/api_1/alipay.trade.page.pay?scene=22
	QrCodeWidth *string `json:"qr_code_width,omitempty"`
	// 支付二维码模式。 只有alipay，且 trade type 为 h5 模式有效; 用法参考 https://opendocs.alipay.com/apis/api_1/alipay.trade.page.pay?scene=22
	QrPayMode *string `json:"qr_pay_mode,omitempty"`
	// RefundNotifyNextTime    *time.Time `json:\"refund_notify_next_time\"`
	RefundNotifyCount *int32 `json:"refund_notify_count,omitempty"`
	RefundState *string `json:"refund_state,omitempty"`
	// 付款成功后的跳转链接。只有alipay，且 trade type 为 h5 模式有效; 用法参考 https://opendocs.alipay.com/apis/api_1/alipay.trade.page.pay?scene=22
	ReturnUrl *string `json:"return_url,omitempty"`
	TimeExpire *string `json:"time_expire,omitempty"`
	TradeNo *string `json:"trade_no,omitempty"`
	TradeProvider *string `json:"trade_provider,omitempty"`
	TradeState *string `json:"trade_state,omitempty"`
	TradeType *string `json:"trade_type,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	WapUrl *string `json:"wap_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsOrder ModelsOrder

// NewModelsOrder instantiates a new ModelsOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsOrder() *ModelsOrder {
	this := ModelsOrder{}
	return &this
}

// NewModelsOrderWithDefaults instantiates a new ModelsOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsOrderWithDefaults() *ModelsOrder {
	this := ModelsOrder{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ModelsOrder) GetAmount() int32 {
	if o == nil || isNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrder) GetAmountOk() (*int32, bool) {
	if o == nil || isNil(o.Amount) {
    return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ModelsOrder) HasAmount() bool {
	if o != nil && !isNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *ModelsOrder) SetAmount(v int32) {
	o.Amount = &v
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *ModelsOrder) GetAppName() string {
	if o == nil || isNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrder) GetAppNameOk() (*string, bool) {
	if o == nil || isNil(o.AppName) {
    return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *ModelsOrder) HasAppName() bool {
	if o != nil && !isNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *ModelsOrder) SetAppName(v string) {
	o.AppName = &v
}

// GetAppPayNotifyUrl returns the AppPayNotifyUrl field value if set, zero value otherwise.
func (o *ModelsOrder) GetAppPayNotifyUrl() string {
	if o == nil || isNil(o.AppPayNotifyUrl) {
		var ret string
		return ret
	}
	return *o.AppPayNotifyUrl
}

// GetAppPayNotifyUrlOk returns a tuple with the AppPayNotifyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrder) GetAppPayNotifyUrlOk() (*string, bool) {
	if o == nil || isNil(o.AppPayNotifyUrl) {
    return nil, false
	}
	return o.AppPayNotifyUrl, true
}

// HasAppPayNotifyUrl returns a boolean if a field has been set.
func (o *ModelsOrder) HasAppPayNotifyUrl() bool {
	if o != nil && !isNil(o.AppPayNotifyUrl) {
		return true
	}

	return false
}

// SetAppPayNotifyUrl gets a reference to the given string and assigns it to the AppPayNotifyUrl field.
func (o *ModelsOrder) SetAppPayNotifyUrl(v string) {
	o.AppPayNotifyUrl = &v
}

// GetAppRefundNotifyUrl returns the AppRefundNotifyUrl field value if set, zero value otherwise.
func (o *ModelsOrder) GetAppRefundNotifyUrl() string {
	if o == nil || isNil(o.AppRefundNotifyUrl) {
		var ret string
		return ret
	}
	return *o.AppRefundNotifyUrl
}

// GetAppRefundNotifyUrlOk returns a tuple with the AppRefundNotifyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrder) GetAppRefundNotifyUrlOk() (*string, bool) {
	if o == nil || isNil(o.AppRefundNotifyUrl) {
    return nil, false
	}
	return o.AppRefundNotifyUrl, true
}

// HasAppRefundNotifyUrl returns a boolean if a field has been set.
func (o *ModelsOrder) HasAppRefundNotifyUrl() bool {
	if o != nil && !isNil(o.AppRefundNotifyUrl) {
		return true
	}

	return false
}

// SetAppRefundNotifyUrl gets a reference to the given string and assigns it to the AppRefundNotifyUrl field.
func (o *ModelsOrder) SetAppRefundNotifyUrl(v string) {
	o.AppRefundNotifyUrl = &v
}

// GetCodeUrl returns the CodeUrl field value if set, zero value otherwise.
func (o *ModelsOrder) GetCodeUrl() string {
	if o == nil || isNil(o.CodeUrl) {
		var ret string
		return ret
	}
	return *o.CodeUrl
}

// GetCodeUrlOk returns a tuple with the CodeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrder) GetCodeUrlOk() (*string, bool) {
	if o == nil || isNil(o.CodeUrl) {
    return nil, false
	}
	return o.CodeUrl, true
}

// HasCodeUrl returns a boolean if a field has been set.
func (o *ModelsOrder) HasCodeUrl() bool {
	if o != nil && !isNil(o.CodeUrl) {
		return true
	}

	return false
}

// SetCodeUrl gets a reference to the given string and assigns it to the CodeUrl field.
func (o *ModelsOrder) SetCodeUrl(v string) {
	o.CodeUrl = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ModelsOrder) GetCreatedAt() string {
	if o == nil || isNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrder) GetCreatedAtOk() (*string, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ModelsOrder) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ModelsOrder) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *ModelsOrder) GetDeletedAt() GormDeletedAt {
	if o == nil || isNil(o.DeletedAt) {
		var ret GormDeletedAt
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrder) GetDeletedAtOk() (*GormDeletedAt, bool) {
	if o == nil || isNil(o.DeletedAt) {
    return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ModelsOrder) HasDeletedAt() bool {
	if o != nil && !isNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given GormDeletedAt and assigns it to the DeletedAt field.
func (o *ModelsOrder) SetDeletedAt(v GormDeletedAt) {
	o.DeletedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModelsOrder) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrder) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModelsOrder) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModelsOrder) SetDescription(v string) {
	o.Description = &v
}

// GetH5Url returns the H5Url field value if set, zero value otherwise.
func (o *ModelsOrder) GetH5Url() string {
	if o == nil || isNil(o.H5Url) {
		var ret string
		return ret
	}
	return *o.H5Url
}

// GetH5UrlOk returns a tuple with the H5Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrder) GetH5UrlOk() (*string, bool) {
	if o == nil || isNil(o.H5Url) {
    return nil, false
	}
	return o.H5Url, true
}

// HasH5Url returns a boolean if a field has been set.
func (o *ModelsOrder) HasH5Url() bool {
	if o != nil && !isNil(o.H5Url) {
		return true
	}

	return false
}

// SetH5Url gets a reference to the given string and assigns it to the H5Url field.
func (o *ModelsOrder) SetH5Url(v string) {
	o.H5Url = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsOrder) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrder) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsOrder) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModelsOrder) SetId(v int32) {
	o.Id = &v
}

// GetIsPayNotifyCompleted returns the IsPayNotifyCompleted field value if set, zero value otherwise.
func (o *ModelsOrder) GetIsPayNotifyCompleted() bool {
	if o == nil || isNil(o.IsPayNotifyCompleted) {
		var ret bool
		return ret
	}
	return *o.IsPayNotifyCompleted
}

// GetIsPayNotifyCompletedOk returns a tuple with the IsPayNotifyCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrder) GetIsPayNotifyCompletedOk() (*bool, bool) {
	if o == nil || isNil(o.IsPayNotifyCompleted) {
    return nil, false
	}
	return o.IsPayNotifyCompleted, true
}

// HasIsPayNotifyCompleted returns a boolean if a field has been set.
func (o *ModelsOrder) HasIsPayNotifyCompleted() bool {
	if o != nil && !isNil(o.IsPayNotifyCompleted) {
		return true
	}

	return false
}

// SetIsPayNotifyCompleted gets a reference to the given bool and assigns it to the IsPayNotifyCompleted field.
func (o *ModelsOrder) SetIsPayNotifyCompleted(v bool) {
	o.IsPayNotifyCompleted = &v
}

// GetIsRefundNotifyCompleted returns the IsRefundNotifyCompleted field value if set, zero value otherwise.
func (o *ModelsOrder) GetIsRefundNotifyCompleted() bool {
	if o == nil || isNil(o.IsRefundNotifyCompleted) {
		var ret bool
		return ret
	}
	return *o.IsRefundNotifyCompleted
}

// GetIsRefundNotifyCompletedOk returns a tuple with the IsRefundNotifyCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrder) GetIsRefundNotifyCompletedOk() (*bool, bool) {
	if o == nil || isNil(o.IsRefundNotifyCompleted) {
    return nil, false
	}
	return o.IsRefundNotifyCompleted, true
}

// HasIsRefundNotifyCompleted returns a boolean if a field has been set.
func (o *ModelsOrder) HasIsRefundNotifyCompleted() bool {
	if o != nil && !isNil(o.IsRefundNotifyCompleted) {
		return true
	}

	return false
}

// SetIsRefundNotifyCompleted gets a reference to the given bool and assigns it to the IsRefundNotifyCompleted field.
func (o *ModelsOrder) SetIsRefundNotifyCompleted(v bool) {
	o.IsRefundNotifyCompleted = &v
}

// GetPayNotifyCount returns the PayNotifyCount field value if set, zero value otherwise.
func (o *ModelsOrder) GetPayNotifyCount() int32 {
	if o == nil || isNil(o.PayNotifyCount) {
		var ret int32
		return ret
	}
	return *o.PayNotifyCount
}

// GetPayNotifyCountOk returns a tuple with the PayNotifyCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrder) GetPayNotifyCountOk() (*int32, bool) {
	if o == nil || isNil(o.PayNotifyCount) {
    return nil, false
	}
	return o.PayNotifyCount, true
}

// HasPayNotifyCount returns a boolean if a field has been set.
func (o *ModelsOrder) HasPayNotifyCount() bool {
	if o != nil && !isNil(o.PayNotifyCount) {
		return true
	}

	return false
}

// SetPayNotifyCount gets a reference to the given int32 and assigns it to the PayNotifyCount field.
func (o *ModelsOrder) SetPayNotifyCount(v int32) {
	o.PayNotifyCount = &v
}

// GetQrCodeWidth returns the QrCodeWidth field value if set, zero value otherwise.
func (o *ModelsOrder) GetQrCodeWidth() string {
	if o == nil || isNil(o.QrCodeWidth) {
		var ret string
		return ret
	}
	return *o.QrCodeWidth
}

// GetQrCodeWidthOk returns a tuple with the QrCodeWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrder) GetQrCodeWidthOk() (*string, bool) {
	if o == nil || isNil(o.QrCodeWidth) {
    return nil, false
	}
	return o.QrCodeWidth, true
}

// HasQrCodeWidth returns a boolean if a field has been set.
func (o *ModelsOrder) HasQrCodeWidth() bool {
	if o != nil && !isNil(o.QrCodeWidth) {
		return true
	}

	return false
}

// SetQrCodeWidth gets a reference to the given string and assigns it to the QrCodeWidth field.
func (o *ModelsOrder) SetQrCodeWidth(v string) {
	o.QrCodeWidth = &v
}

// GetQrPayMode returns the QrPayMode field value if set, zero value otherwise.
func (o *ModelsOrder) GetQrPayMode() string {
	if o == nil || isNil(o.QrPayMode) {
		var ret string
		return ret
	}
	return *o.QrPayMode
}

// GetQrPayModeOk returns a tuple with the QrPayMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrder) GetQrPayModeOk() (*string, bool) {
	if o == nil || isNil(o.QrPayMode) {
    return nil, false
	}
	return o.QrPayMode, true
}

// HasQrPayMode returns a boolean if a field has been set.
func (o *ModelsOrder) HasQrPayMode() bool {
	if o != nil && !isNil(o.QrPayMode) {
		return true
	}

	return false
}

// SetQrPayMode gets a reference to the given string and assigns it to the QrPayMode field.
func (o *ModelsOrder) SetQrPayMode(v string) {
	o.QrPayMode = &v
}

// GetRefundNotifyCount returns the RefundNotifyCount field value if set, zero value otherwise.
func (o *ModelsOrder) GetRefundNotifyCount() int32 {
	if o == nil || isNil(o.RefundNotifyCount) {
		var ret int32
		return ret
	}
	return *o.RefundNotifyCount
}

// GetRefundNotifyCountOk returns a tuple with the RefundNotifyCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrder) GetRefundNotifyCountOk() (*int32, bool) {
	if o == nil || isNil(o.RefundNotifyCount) {
    return nil, false
	}
	return o.RefundNotifyCount, true
}

// HasRefundNotifyCount returns a boolean if a field has been set.
func (o *ModelsOrder) HasRefundNotifyCount() bool {
	if o != nil && !isNil(o.RefundNotifyCount) {
		return true
	}

	return false
}

// SetRefundNotifyCount gets a reference to the given int32 and assigns it to the RefundNotifyCount field.
func (o *ModelsOrder) SetRefundNotifyCount(v int32) {
	o.RefundNotifyCount = &v
}

// GetRefundState returns the RefundState field value if set, zero value otherwise.
func (o *ModelsOrder) GetRefundState() string {
	if o == nil || isNil(o.RefundState) {
		var ret string
		return ret
	}
	return *o.RefundState
}

// GetRefundStateOk returns a tuple with the RefundState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrder) GetRefundStateOk() (*string, bool) {
	if o == nil || isNil(o.RefundState) {
    return nil, false
	}
	return o.RefundState, true
}

// HasRefundState returns a boolean if a field has been set.
func (o *ModelsOrder) HasRefundState() bool {
	if o != nil && !isNil(o.RefundState) {
		return true
	}

	return false
}

// SetRefundState gets a reference to the given string and assigns it to the RefundState field.
func (o *ModelsOrder) SetRefundState(v string) {
	o.RefundState = &v
}

// GetReturnUrl returns the ReturnUrl field value if set, zero value otherwise.
func (o *ModelsOrder) GetReturnUrl() string {
	if o == nil || isNil(o.ReturnUrl) {
		var ret string
		return ret
	}
	return *o.ReturnUrl
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrder) GetReturnUrlOk() (*string, bool) {
	if o == nil || isNil(o.ReturnUrl) {
    return nil, false
	}
	return o.ReturnUrl, true
}

// HasReturnUrl returns a boolean if a field has been set.
func (o *ModelsOrder) HasReturnUrl() bool {
	if o != nil && !isNil(o.ReturnUrl) {
		return true
	}

	return false
}

// SetReturnUrl gets a reference to the given string and assigns it to the ReturnUrl field.
func (o *ModelsOrder) SetReturnUrl(v string) {
	o.ReturnUrl = &v
}

// GetTimeExpire returns the TimeExpire field value if set, zero value otherwise.
func (o *ModelsOrder) GetTimeExpire() string {
	if o == nil || isNil(o.TimeExpire) {
		var ret string
		return ret
	}
	return *o.TimeExpire
}

// GetTimeExpireOk returns a tuple with the TimeExpire field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrder) GetTimeExpireOk() (*string, bool) {
	if o == nil || isNil(o.TimeExpire) {
    return nil, false
	}
	return o.TimeExpire, true
}

// HasTimeExpire returns a boolean if a field has been set.
func (o *ModelsOrder) HasTimeExpire() bool {
	if o != nil && !isNil(o.TimeExpire) {
		return true
	}

	return false
}

// SetTimeExpire gets a reference to the given string and assigns it to the TimeExpire field.
func (o *ModelsOrder) SetTimeExpire(v string) {
	o.TimeExpire = &v
}

// GetTradeNo returns the TradeNo field value if set, zero value otherwise.
func (o *ModelsOrder) GetTradeNo() string {
	if o == nil || isNil(o.TradeNo) {
		var ret string
		return ret
	}
	return *o.TradeNo
}

// GetTradeNoOk returns a tuple with the TradeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrder) GetTradeNoOk() (*string, bool) {
	if o == nil || isNil(o.TradeNo) {
    return nil, false
	}
	return o.TradeNo, true
}

// HasTradeNo returns a boolean if a field has been set.
func (o *ModelsOrder) HasTradeNo() bool {
	if o != nil && !isNil(o.TradeNo) {
		return true
	}

	return false
}

// SetTradeNo gets a reference to the given string and assigns it to the TradeNo field.
func (o *ModelsOrder) SetTradeNo(v string) {
	o.TradeNo = &v
}

// GetTradeProvider returns the TradeProvider field value if set, zero value otherwise.
func (o *ModelsOrder) GetTradeProvider() string {
	if o == nil || isNil(o.TradeProvider) {
		var ret string
		return ret
	}
	return *o.TradeProvider
}

// GetTradeProviderOk returns a tuple with the TradeProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrder) GetTradeProviderOk() (*string, bool) {
	if o == nil || isNil(o.TradeProvider) {
    return nil, false
	}
	return o.TradeProvider, true
}

// HasTradeProvider returns a boolean if a field has been set.
func (o *ModelsOrder) HasTradeProvider() bool {
	if o != nil && !isNil(o.TradeProvider) {
		return true
	}

	return false
}

// SetTradeProvider gets a reference to the given string and assigns it to the TradeProvider field.
func (o *ModelsOrder) SetTradeProvider(v string) {
	o.TradeProvider = &v
}

// GetTradeState returns the TradeState field value if set, zero value otherwise.
func (o *ModelsOrder) GetTradeState() string {
	if o == nil || isNil(o.TradeState) {
		var ret string
		return ret
	}
	return *o.TradeState
}

// GetTradeStateOk returns a tuple with the TradeState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrder) GetTradeStateOk() (*string, bool) {
	if o == nil || isNil(o.TradeState) {
    return nil, false
	}
	return o.TradeState, true
}

// HasTradeState returns a boolean if a field has been set.
func (o *ModelsOrder) HasTradeState() bool {
	if o != nil && !isNil(o.TradeState) {
		return true
	}

	return false
}

// SetTradeState gets a reference to the given string and assigns it to the TradeState field.
func (o *ModelsOrder) SetTradeState(v string) {
	o.TradeState = &v
}

// GetTradeType returns the TradeType field value if set, zero value otherwise.
func (o *ModelsOrder) GetTradeType() string {
	if o == nil || isNil(o.TradeType) {
		var ret string
		return ret
	}
	return *o.TradeType
}

// GetTradeTypeOk returns a tuple with the TradeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrder) GetTradeTypeOk() (*string, bool) {
	if o == nil || isNil(o.TradeType) {
    return nil, false
	}
	return o.TradeType, true
}

// HasTradeType returns a boolean if a field has been set.
func (o *ModelsOrder) HasTradeType() bool {
	if o != nil && !isNil(o.TradeType) {
		return true
	}

	return false
}

// SetTradeType gets a reference to the given string and assigns it to the TradeType field.
func (o *ModelsOrder) SetTradeType(v string) {
	o.TradeType = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ModelsOrder) GetUpdatedAt() string {
	if o == nil || isNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrder) GetUpdatedAtOk() (*string, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ModelsOrder) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ModelsOrder) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetWapUrl returns the WapUrl field value if set, zero value otherwise.
func (o *ModelsOrder) GetWapUrl() string {
	if o == nil || isNil(o.WapUrl) {
		var ret string
		return ret
	}
	return *o.WapUrl
}

// GetWapUrlOk returns a tuple with the WapUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsOrder) GetWapUrlOk() (*string, bool) {
	if o == nil || isNil(o.WapUrl) {
    return nil, false
	}
	return o.WapUrl, true
}

// HasWapUrl returns a boolean if a field has been set.
func (o *ModelsOrder) HasWapUrl() bool {
	if o != nil && !isNil(o.WapUrl) {
		return true
	}

	return false
}

// SetWapUrl gets a reference to the given string and assigns it to the WapUrl field.
func (o *ModelsOrder) SetWapUrl(v string) {
	o.WapUrl = &v
}

func (o ModelsOrder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !isNil(o.AppName) {
		toSerialize["app_name"] = o.AppName
	}
	if !isNil(o.AppPayNotifyUrl) {
		toSerialize["app_pay_notify_url"] = o.AppPayNotifyUrl
	}
	if !isNil(o.AppRefundNotifyUrl) {
		toSerialize["app_refund_notify_url"] = o.AppRefundNotifyUrl
	}
	if !isNil(o.CodeUrl) {
		toSerialize["code_url"] = o.CodeUrl
	}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.H5Url) {
		toSerialize["h5_url"] = o.H5Url
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.IsPayNotifyCompleted) {
		toSerialize["is_pay_notify_completed"] = o.IsPayNotifyCompleted
	}
	if !isNil(o.IsRefundNotifyCompleted) {
		toSerialize["is_refund_notify_completed"] = o.IsRefundNotifyCompleted
	}
	if !isNil(o.PayNotifyCount) {
		toSerialize["pay_notify_count"] = o.PayNotifyCount
	}
	if !isNil(o.QrCodeWidth) {
		toSerialize["qr_code_width"] = o.QrCodeWidth
	}
	if !isNil(o.QrPayMode) {
		toSerialize["qr_pay_mode"] = o.QrPayMode
	}
	if !isNil(o.RefundNotifyCount) {
		toSerialize["refund_notify_count"] = o.RefundNotifyCount
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
	if !isNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !isNil(o.WapUrl) {
		toSerialize["wap_url"] = o.WapUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ModelsOrder) UnmarshalJSON(bytes []byte) (err error) {
	varModelsOrder := _ModelsOrder{}

	if err = json.Unmarshal(bytes, &varModelsOrder); err == nil {
		*o = ModelsOrder(varModelsOrder)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "app_name")
		delete(additionalProperties, "app_pay_notify_url")
		delete(additionalProperties, "app_refund_notify_url")
		delete(additionalProperties, "code_url")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "description")
		delete(additionalProperties, "h5_url")
		delete(additionalProperties, "id")
		delete(additionalProperties, "is_pay_notify_completed")
		delete(additionalProperties, "is_refund_notify_completed")
		delete(additionalProperties, "pay_notify_count")
		delete(additionalProperties, "qr_code_width")
		delete(additionalProperties, "qr_pay_mode")
		delete(additionalProperties, "refund_notify_count")
		delete(additionalProperties, "refund_state")
		delete(additionalProperties, "return_url")
		delete(additionalProperties, "time_expire")
		delete(additionalProperties, "trade_no")
		delete(additionalProperties, "trade_provider")
		delete(additionalProperties, "trade_state")
		delete(additionalProperties, "trade_type")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "wap_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsOrder struct {
	value *ModelsOrder
	isSet bool
}

func (v NullableModelsOrder) Get() *ModelsOrder {
	return v.value
}

func (v *NullableModelsOrder) Set(val *ModelsOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsOrder(val *ModelsOrder) *NullableModelsOrder {
	return &NullableModelsOrder{value: val, isSet: true}
}

func (v NullableModelsOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


