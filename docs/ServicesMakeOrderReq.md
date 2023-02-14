# ServicesMakeOrderReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** |  | 
**Description** | **string** |  | 
**NotifyUrl** | Pointer to **string** |  | [optional] 
**QrCodeWidth** | Pointer to **string** | 二维码宽度。 只有alipay，且 trade type 为 h5 模式有效，qr pay mode 为4 时有效； 用法参考 https://opendocs.alipay.com/apis/api_1/alipay.trade.page.pay?scene&#x3D;22 | [optional] 
**QrPayMode** | Pointer to **string** | 支付二维码模式。 只有alipay，且 trade type 为 h5 模式有效; 用法参考 https://opendocs.alipay.com/apis/api_1/alipay.trade.page.pay?scene&#x3D;22 | [optional] 
**ReturnUrl** | Pointer to **string** | 付款成功后的跳转链接。只有alipay，且 trade type 为 h5 模式有效; 用法参考 https://opendocs.alipay.com/apis/api_1/alipay.trade.page.pay?scene&#x3D;22 | [optional] 
**TimeExpire** | **int32** | alipay 当面付无效，当面付固定过期时间为2小时 | 
**TradeProvider** | **string** |  | 
**TradeType** | **string** |  | 

## Methods

### NewServicesMakeOrderReq

`func NewServicesMakeOrderReq(amount int32, description string, timeExpire int32, tradeProvider string, tradeType string, ) *ServicesMakeOrderReq`

NewServicesMakeOrderReq instantiates a new ServicesMakeOrderReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesMakeOrderReqWithDefaults

`func NewServicesMakeOrderReqWithDefaults() *ServicesMakeOrderReq`

NewServicesMakeOrderReqWithDefaults instantiates a new ServicesMakeOrderReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ServicesMakeOrderReq) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ServicesMakeOrderReq) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ServicesMakeOrderReq) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *ServicesMakeOrderReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServicesMakeOrderReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServicesMakeOrderReq) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetNotifyUrl

`func (o *ServicesMakeOrderReq) GetNotifyUrl() string`

GetNotifyUrl returns the NotifyUrl field if non-nil, zero value otherwise.

### GetNotifyUrlOk

`func (o *ServicesMakeOrderReq) GetNotifyUrlOk() (*string, bool)`

GetNotifyUrlOk returns a tuple with the NotifyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUrl

`func (o *ServicesMakeOrderReq) SetNotifyUrl(v string)`

SetNotifyUrl sets NotifyUrl field to given value.

### HasNotifyUrl

`func (o *ServicesMakeOrderReq) HasNotifyUrl() bool`

HasNotifyUrl returns a boolean if a field has been set.

### GetQrCodeWidth

`func (o *ServicesMakeOrderReq) GetQrCodeWidth() string`

GetQrCodeWidth returns the QrCodeWidth field if non-nil, zero value otherwise.

### GetQrCodeWidthOk

`func (o *ServicesMakeOrderReq) GetQrCodeWidthOk() (*string, bool)`

GetQrCodeWidthOk returns a tuple with the QrCodeWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCodeWidth

`func (o *ServicesMakeOrderReq) SetQrCodeWidth(v string)`

SetQrCodeWidth sets QrCodeWidth field to given value.

### HasQrCodeWidth

`func (o *ServicesMakeOrderReq) HasQrCodeWidth() bool`

HasQrCodeWidth returns a boolean if a field has been set.

### GetQrPayMode

`func (o *ServicesMakeOrderReq) GetQrPayMode() string`

GetQrPayMode returns the QrPayMode field if non-nil, zero value otherwise.

### GetQrPayModeOk

`func (o *ServicesMakeOrderReq) GetQrPayModeOk() (*string, bool)`

GetQrPayModeOk returns a tuple with the QrPayMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrPayMode

`func (o *ServicesMakeOrderReq) SetQrPayMode(v string)`

SetQrPayMode sets QrPayMode field to given value.

### HasQrPayMode

`func (o *ServicesMakeOrderReq) HasQrPayMode() bool`

HasQrPayMode returns a boolean if a field has been set.

### GetReturnUrl

`func (o *ServicesMakeOrderReq) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *ServicesMakeOrderReq) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *ServicesMakeOrderReq) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *ServicesMakeOrderReq) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetTimeExpire

`func (o *ServicesMakeOrderReq) GetTimeExpire() int32`

GetTimeExpire returns the TimeExpire field if non-nil, zero value otherwise.

### GetTimeExpireOk

`func (o *ServicesMakeOrderReq) GetTimeExpireOk() (*int32, bool)`

GetTimeExpireOk returns a tuple with the TimeExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeExpire

`func (o *ServicesMakeOrderReq) SetTimeExpire(v int32)`

SetTimeExpire sets TimeExpire field to given value.


### GetTradeProvider

`func (o *ServicesMakeOrderReq) GetTradeProvider() string`

GetTradeProvider returns the TradeProvider field if non-nil, zero value otherwise.

### GetTradeProviderOk

`func (o *ServicesMakeOrderReq) GetTradeProviderOk() (*string, bool)`

GetTradeProviderOk returns a tuple with the TradeProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeProvider

`func (o *ServicesMakeOrderReq) SetTradeProvider(v string)`

SetTradeProvider sets TradeProvider field to given value.


### GetTradeType

`func (o *ServicesMakeOrderReq) GetTradeType() string`

GetTradeType returns the TradeType field if non-nil, zero value otherwise.

### GetTradeTypeOk

`func (o *ServicesMakeOrderReq) GetTradeTypeOk() (*string, bool)`

GetTradeTypeOk returns a tuple with the TradeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeType

`func (o *ServicesMakeOrderReq) SetTradeType(v string)`

SetTradeType sets TradeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


