# ModelsOrderCore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | 单位为分 | [optional] 
**AppName** | Pointer to **string** |  | [optional] 
**CodeUrl** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**H5Url** | Pointer to **string** |  | [optional] 
**QrCodeWidth** | Pointer to **string** | 二维码宽度。 只有alipay，且 trade type 为 h5 模式有效，qr pay mode 为4 时有效； 用法参考 https://opendocs.alipay.com/apis/api_1/alipay.trade.page.pay?scene&#x3D;22 | [optional] 
**QrPayMode** | Pointer to **string** | 支付二维码模式。 只有alipay，且 trade type 为 h5 模式有效; 用法参考 https://opendocs.alipay.com/apis/api_1/alipay.trade.page.pay?scene&#x3D;22 | [optional] 
**RefundState** | Pointer to **string** |  | [optional] 
**ReturnUrl** | Pointer to **string** | 付款成功后的跳转链接。只有alipay，且 trade type 为 h5 模式有效; 用法参考 https://opendocs.alipay.com/apis/api_1/alipay.trade.page.pay?scene&#x3D;22 | [optional] 
**TimeExpire** | Pointer to **string** |  | [optional] 
**TradeNo** | Pointer to **string** |  | [optional] 
**TradeProvider** | Pointer to **string** |  | [optional] 
**TradeState** | Pointer to **string** |  | [optional] 
**TradeType** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsOrderCore

`func NewModelsOrderCore() *ModelsOrderCore`

NewModelsOrderCore instantiates a new ModelsOrderCore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsOrderCoreWithDefaults

`func NewModelsOrderCoreWithDefaults() *ModelsOrderCore`

NewModelsOrderCoreWithDefaults instantiates a new ModelsOrderCore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ModelsOrderCore) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ModelsOrderCore) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ModelsOrderCore) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ModelsOrderCore) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAppName

`func (o *ModelsOrderCore) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *ModelsOrderCore) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *ModelsOrderCore) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *ModelsOrderCore) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetCodeUrl

`func (o *ModelsOrderCore) GetCodeUrl() string`

GetCodeUrl returns the CodeUrl field if non-nil, zero value otherwise.

### GetCodeUrlOk

`func (o *ModelsOrderCore) GetCodeUrlOk() (*string, bool)`

GetCodeUrlOk returns a tuple with the CodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeUrl

`func (o *ModelsOrderCore) SetCodeUrl(v string)`

SetCodeUrl sets CodeUrl field to given value.

### HasCodeUrl

`func (o *ModelsOrderCore) HasCodeUrl() bool`

HasCodeUrl returns a boolean if a field has been set.

### GetDescription

`func (o *ModelsOrderCore) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelsOrderCore) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelsOrderCore) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelsOrderCore) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetH5Url

`func (o *ModelsOrderCore) GetH5Url() string`

GetH5Url returns the H5Url field if non-nil, zero value otherwise.

### GetH5UrlOk

`func (o *ModelsOrderCore) GetH5UrlOk() (*string, bool)`

GetH5UrlOk returns a tuple with the H5Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH5Url

`func (o *ModelsOrderCore) SetH5Url(v string)`

SetH5Url sets H5Url field to given value.

### HasH5Url

`func (o *ModelsOrderCore) HasH5Url() bool`

HasH5Url returns a boolean if a field has been set.

### GetQrCodeWidth

`func (o *ModelsOrderCore) GetQrCodeWidth() string`

GetQrCodeWidth returns the QrCodeWidth field if non-nil, zero value otherwise.

### GetQrCodeWidthOk

`func (o *ModelsOrderCore) GetQrCodeWidthOk() (*string, bool)`

GetQrCodeWidthOk returns a tuple with the QrCodeWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCodeWidth

`func (o *ModelsOrderCore) SetQrCodeWidth(v string)`

SetQrCodeWidth sets QrCodeWidth field to given value.

### HasQrCodeWidth

`func (o *ModelsOrderCore) HasQrCodeWidth() bool`

HasQrCodeWidth returns a boolean if a field has been set.

### GetQrPayMode

`func (o *ModelsOrderCore) GetQrPayMode() string`

GetQrPayMode returns the QrPayMode field if non-nil, zero value otherwise.

### GetQrPayModeOk

`func (o *ModelsOrderCore) GetQrPayModeOk() (*string, bool)`

GetQrPayModeOk returns a tuple with the QrPayMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrPayMode

`func (o *ModelsOrderCore) SetQrPayMode(v string)`

SetQrPayMode sets QrPayMode field to given value.

### HasQrPayMode

`func (o *ModelsOrderCore) HasQrPayMode() bool`

HasQrPayMode returns a boolean if a field has been set.

### GetRefundState

`func (o *ModelsOrderCore) GetRefundState() string`

GetRefundState returns the RefundState field if non-nil, zero value otherwise.

### GetRefundStateOk

`func (o *ModelsOrderCore) GetRefundStateOk() (*string, bool)`

GetRefundStateOk returns a tuple with the RefundState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundState

`func (o *ModelsOrderCore) SetRefundState(v string)`

SetRefundState sets RefundState field to given value.

### HasRefundState

`func (o *ModelsOrderCore) HasRefundState() bool`

HasRefundState returns a boolean if a field has been set.

### GetReturnUrl

`func (o *ModelsOrderCore) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *ModelsOrderCore) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *ModelsOrderCore) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *ModelsOrderCore) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetTimeExpire

`func (o *ModelsOrderCore) GetTimeExpire() string`

GetTimeExpire returns the TimeExpire field if non-nil, zero value otherwise.

### GetTimeExpireOk

`func (o *ModelsOrderCore) GetTimeExpireOk() (*string, bool)`

GetTimeExpireOk returns a tuple with the TimeExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeExpire

`func (o *ModelsOrderCore) SetTimeExpire(v string)`

SetTimeExpire sets TimeExpire field to given value.

### HasTimeExpire

`func (o *ModelsOrderCore) HasTimeExpire() bool`

HasTimeExpire returns a boolean if a field has been set.

### GetTradeNo

`func (o *ModelsOrderCore) GetTradeNo() string`

GetTradeNo returns the TradeNo field if non-nil, zero value otherwise.

### GetTradeNoOk

`func (o *ModelsOrderCore) GetTradeNoOk() (*string, bool)`

GetTradeNoOk returns a tuple with the TradeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeNo

`func (o *ModelsOrderCore) SetTradeNo(v string)`

SetTradeNo sets TradeNo field to given value.

### HasTradeNo

`func (o *ModelsOrderCore) HasTradeNo() bool`

HasTradeNo returns a boolean if a field has been set.

### GetTradeProvider

`func (o *ModelsOrderCore) GetTradeProvider() string`

GetTradeProvider returns the TradeProvider field if non-nil, zero value otherwise.

### GetTradeProviderOk

`func (o *ModelsOrderCore) GetTradeProviderOk() (*string, bool)`

GetTradeProviderOk returns a tuple with the TradeProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeProvider

`func (o *ModelsOrderCore) SetTradeProvider(v string)`

SetTradeProvider sets TradeProvider field to given value.

### HasTradeProvider

`func (o *ModelsOrderCore) HasTradeProvider() bool`

HasTradeProvider returns a boolean if a field has been set.

### GetTradeState

`func (o *ModelsOrderCore) GetTradeState() string`

GetTradeState returns the TradeState field if non-nil, zero value otherwise.

### GetTradeStateOk

`func (o *ModelsOrderCore) GetTradeStateOk() (*string, bool)`

GetTradeStateOk returns a tuple with the TradeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeState

`func (o *ModelsOrderCore) SetTradeState(v string)`

SetTradeState sets TradeState field to given value.

### HasTradeState

`func (o *ModelsOrderCore) HasTradeState() bool`

HasTradeState returns a boolean if a field has been set.

### GetTradeType

`func (o *ModelsOrderCore) GetTradeType() string`

GetTradeType returns the TradeType field if non-nil, zero value otherwise.

### GetTradeTypeOk

`func (o *ModelsOrderCore) GetTradeTypeOk() (*string, bool)`

GetTradeTypeOk returns a tuple with the TradeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeType

`func (o *ModelsOrderCore) SetTradeType(v string)`

SetTradeType sets TradeType field to given value.

### HasTradeType

`func (o *ModelsOrderCore) HasTradeType() bool`

HasTradeType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


