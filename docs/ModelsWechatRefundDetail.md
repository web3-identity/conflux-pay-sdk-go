# ModelsWechatRefundDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | 金额详细信息 | [optional] 
**Channel** | Pointer to **string** | 枚举值： - ORIGINAL—原路退款 - BALANCE—退回到余额 - OTHER_BALANCE—原账户异常退到其他余额账户 - OTHER_BANKCARD—原银行卡异常退到其他银行卡 * &#x60;ORIGINAL&#x60; - 原路退款 * &#x60;BALANCE&#x60; - 退回到余额 * &#x60;OTHER_BALANCE&#x60; - 原账户异常退到其他余额账户 * &#x60;OTHER_BANKCARD&#x60; - 原银行卡异常退到其他银行卡 | [optional] 
**CreateTime** | Pointer to **string** | 退款受理时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日13点29分35秒。 | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to [**GormDeletedAt**](GormDeletedAt.md) |  | [optional] 
**FundsAccount** | Pointer to **string** | 退款所使用资金对应的资金账户类型 枚举值： - UNSETTLED : 未结算资金 - AVAILABLE : 可用余额 - UNAVAILABLE : 不可用余额 - OPERATION : 运营户 - BASIC : 基本账户（含可用余额和不可用余额） * &#x60;UNSETTLED&#x60; - 未结算资金 * &#x60;AVAILABLE&#x60; - 可用余额 * &#x60;UNAVAILABLE&#x60; - 不可用余额 * &#x60;OPERATION&#x60; - 运营户 * &#x60;BASIC&#x60; - 基本账户（含可用余额和不可用余额） | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**OutRefundNo** | Pointer to **string** | 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。 | [optional] 
**OutTradeNo** | Pointer to **string** | 原支付交易对应的商户订单号 | [optional] 
**RefundId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | 退款到银行发现用户的卡作废或者冻结了，导致原路退款银行卡失败，可前往商户平台（pay.weixin.qq.com）-交易中心，手动处理此笔退款。 枚举值： - SUCCESS—退款成功 - CLOSED—退款关闭 - PROCESSING—退款处理中 - ABNORMAL—退款异常 * &#x60;SUCCESS&#x60; - 退款成功 * &#x60;CLOSED&#x60; - 退款关闭 * &#x60;PROCESSING&#x60; - 退款处理中 * &#x60;ABNORMAL&#x60; - 退款异常 | [optional] 
**SuccessTime** | Pointer to **string** | 退款成功时间，退款状态status为SUCCESS（退款成功）时，返回该字段。遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日13点29分35秒。 | [optional] 
**TransactionId** | Pointer to **string** | 微信支付交易订单号 | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UserReceivedAccount** | Pointer to **string** | 取当前退款单的退款入账方，有以下几种情况： 1）退回银行卡：{银行名称}{卡类型}{卡尾号} 2）退回支付用户零钱:支付用户零钱 3）退还商户:商户基本账户商户结算银行账户 4）退回支付用户零钱通:支付用户零钱通 | [optional] 

## Methods

### NewModelsWechatRefundDetail

`func NewModelsWechatRefundDetail() *ModelsWechatRefundDetail`

NewModelsWechatRefundDetail instantiates a new ModelsWechatRefundDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsWechatRefundDetailWithDefaults

`func NewModelsWechatRefundDetailWithDefaults() *ModelsWechatRefundDetail`

NewModelsWechatRefundDetailWithDefaults instantiates a new ModelsWechatRefundDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ModelsWechatRefundDetail) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ModelsWechatRefundDetail) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ModelsWechatRefundDetail) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ModelsWechatRefundDetail) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetChannel

`func (o *ModelsWechatRefundDetail) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ModelsWechatRefundDetail) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ModelsWechatRefundDetail) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ModelsWechatRefundDetail) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreateTime

`func (o *ModelsWechatRefundDetail) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ModelsWechatRefundDetail) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ModelsWechatRefundDetail) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ModelsWechatRefundDetail) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsWechatRefundDetail) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsWechatRefundDetail) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsWechatRefundDetail) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsWechatRefundDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ModelsWechatRefundDetail) GetDeletedAt() GormDeletedAt`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ModelsWechatRefundDetail) GetDeletedAtOk() (*GormDeletedAt, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ModelsWechatRefundDetail) SetDeletedAt(v GormDeletedAt)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ModelsWechatRefundDetail) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetFundsAccount

`func (o *ModelsWechatRefundDetail) GetFundsAccount() string`

GetFundsAccount returns the FundsAccount field if non-nil, zero value otherwise.

### GetFundsAccountOk

`func (o *ModelsWechatRefundDetail) GetFundsAccountOk() (*string, bool)`

GetFundsAccountOk returns a tuple with the FundsAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundsAccount

`func (o *ModelsWechatRefundDetail) SetFundsAccount(v string)`

SetFundsAccount sets FundsAccount field to given value.

### HasFundsAccount

`func (o *ModelsWechatRefundDetail) HasFundsAccount() bool`

HasFundsAccount returns a boolean if a field has been set.

### GetId

`func (o *ModelsWechatRefundDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsWechatRefundDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsWechatRefundDetail) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsWechatRefundDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOutRefundNo

`func (o *ModelsWechatRefundDetail) GetOutRefundNo() string`

GetOutRefundNo returns the OutRefundNo field if non-nil, zero value otherwise.

### GetOutRefundNoOk

`func (o *ModelsWechatRefundDetail) GetOutRefundNoOk() (*string, bool)`

GetOutRefundNoOk returns a tuple with the OutRefundNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutRefundNo

`func (o *ModelsWechatRefundDetail) SetOutRefundNo(v string)`

SetOutRefundNo sets OutRefundNo field to given value.

### HasOutRefundNo

`func (o *ModelsWechatRefundDetail) HasOutRefundNo() bool`

HasOutRefundNo returns a boolean if a field has been set.

### GetOutTradeNo

`func (o *ModelsWechatRefundDetail) GetOutTradeNo() string`

GetOutTradeNo returns the OutTradeNo field if non-nil, zero value otherwise.

### GetOutTradeNoOk

`func (o *ModelsWechatRefundDetail) GetOutTradeNoOk() (*string, bool)`

GetOutTradeNoOk returns a tuple with the OutTradeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutTradeNo

`func (o *ModelsWechatRefundDetail) SetOutTradeNo(v string)`

SetOutTradeNo sets OutTradeNo field to given value.

### HasOutTradeNo

`func (o *ModelsWechatRefundDetail) HasOutTradeNo() bool`

HasOutTradeNo returns a boolean if a field has been set.

### GetRefundId

`func (o *ModelsWechatRefundDetail) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *ModelsWechatRefundDetail) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *ModelsWechatRefundDetail) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.

### HasRefundId

`func (o *ModelsWechatRefundDetail) HasRefundId() bool`

HasRefundId returns a boolean if a field has been set.

### GetStatus

`func (o *ModelsWechatRefundDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelsWechatRefundDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelsWechatRefundDetail) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelsWechatRefundDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuccessTime

`func (o *ModelsWechatRefundDetail) GetSuccessTime() string`

GetSuccessTime returns the SuccessTime field if non-nil, zero value otherwise.

### GetSuccessTimeOk

`func (o *ModelsWechatRefundDetail) GetSuccessTimeOk() (*string, bool)`

GetSuccessTimeOk returns a tuple with the SuccessTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessTime

`func (o *ModelsWechatRefundDetail) SetSuccessTime(v string)`

SetSuccessTime sets SuccessTime field to given value.

### HasSuccessTime

`func (o *ModelsWechatRefundDetail) HasSuccessTime() bool`

HasSuccessTime returns a boolean if a field has been set.

### GetTransactionId

`func (o *ModelsWechatRefundDetail) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ModelsWechatRefundDetail) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ModelsWechatRefundDetail) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *ModelsWechatRefundDetail) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsWechatRefundDetail) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsWechatRefundDetail) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsWechatRefundDetail) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsWechatRefundDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserReceivedAccount

`func (o *ModelsWechatRefundDetail) GetUserReceivedAccount() string`

GetUserReceivedAccount returns the UserReceivedAccount field if non-nil, zero value otherwise.

### GetUserReceivedAccountOk

`func (o *ModelsWechatRefundDetail) GetUserReceivedAccountOk() (*string, bool)`

GetUserReceivedAccountOk returns a tuple with the UserReceivedAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserReceivedAccount

`func (o *ModelsWechatRefundDetail) SetUserReceivedAccount(v string)`

SetUserReceivedAccount sets UserReceivedAccount field to given value.

### HasUserReceivedAccount

`func (o *ModelsWechatRefundDetail) HasUserReceivedAccount() bool`

HasUserReceivedAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


