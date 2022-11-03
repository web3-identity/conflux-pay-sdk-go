# ModelsWechatOrderDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** |  | [optional] 
**Appid** | Pointer to **string** |  | [optional] 
**Attach** | Pointer to **string** |  | [optional] 
**BankType** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to [**GormDeletedAt**](GormDeletedAt.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Mchid** | Pointer to **string** |  | [optional] 
**Payer** | Pointer to **string** |  | [optional] 
**RefreshStatus** | Pointer to **string** |  | [optional] 
**RefundNo** | Pointer to **string** |  | [optional] 
**SuccessTime** | Pointer to **string** |  | [optional] 
**TradeNo** | Pointer to **string** |  | [optional] 
**TradeState** | Pointer to **string** |  | [optional] 
**TradeStateDesc** | Pointer to **string** |  | [optional] 
**TradeType** | Pointer to **string** |  | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsWechatOrderDetail

`func NewModelsWechatOrderDetail() *ModelsWechatOrderDetail`

NewModelsWechatOrderDetail instantiates a new ModelsWechatOrderDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsWechatOrderDetailWithDefaults

`func NewModelsWechatOrderDetailWithDefaults() *ModelsWechatOrderDetail`

NewModelsWechatOrderDetailWithDefaults instantiates a new ModelsWechatOrderDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ModelsWechatOrderDetail) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ModelsWechatOrderDetail) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ModelsWechatOrderDetail) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ModelsWechatOrderDetail) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAppid

`func (o *ModelsWechatOrderDetail) GetAppid() string`

GetAppid returns the Appid field if non-nil, zero value otherwise.

### GetAppidOk

`func (o *ModelsWechatOrderDetail) GetAppidOk() (*string, bool)`

GetAppidOk returns a tuple with the Appid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppid

`func (o *ModelsWechatOrderDetail) SetAppid(v string)`

SetAppid sets Appid field to given value.

### HasAppid

`func (o *ModelsWechatOrderDetail) HasAppid() bool`

HasAppid returns a boolean if a field has been set.

### GetAttach

`func (o *ModelsWechatOrderDetail) GetAttach() string`

GetAttach returns the Attach field if non-nil, zero value otherwise.

### GetAttachOk

`func (o *ModelsWechatOrderDetail) GetAttachOk() (*string, bool)`

GetAttachOk returns a tuple with the Attach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttach

`func (o *ModelsWechatOrderDetail) SetAttach(v string)`

SetAttach sets Attach field to given value.

### HasAttach

`func (o *ModelsWechatOrderDetail) HasAttach() bool`

HasAttach returns a boolean if a field has been set.

### GetBankType

`func (o *ModelsWechatOrderDetail) GetBankType() string`

GetBankType returns the BankType field if non-nil, zero value otherwise.

### GetBankTypeOk

`func (o *ModelsWechatOrderDetail) GetBankTypeOk() (*string, bool)`

GetBankTypeOk returns a tuple with the BankType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankType

`func (o *ModelsWechatOrderDetail) SetBankType(v string)`

SetBankType sets BankType field to given value.

### HasBankType

`func (o *ModelsWechatOrderDetail) HasBankType() bool`

HasBankType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsWechatOrderDetail) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsWechatOrderDetail) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsWechatOrderDetail) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsWechatOrderDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ModelsWechatOrderDetail) GetDeletedAt() GormDeletedAt`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ModelsWechatOrderDetail) GetDeletedAtOk() (*GormDeletedAt, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ModelsWechatOrderDetail) SetDeletedAt(v GormDeletedAt)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ModelsWechatOrderDetail) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetId

`func (o *ModelsWechatOrderDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsWechatOrderDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsWechatOrderDetail) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsWechatOrderDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMchid

`func (o *ModelsWechatOrderDetail) GetMchid() string`

GetMchid returns the Mchid field if non-nil, zero value otherwise.

### GetMchidOk

`func (o *ModelsWechatOrderDetail) GetMchidOk() (*string, bool)`

GetMchidOk returns a tuple with the Mchid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMchid

`func (o *ModelsWechatOrderDetail) SetMchid(v string)`

SetMchid sets Mchid field to given value.

### HasMchid

`func (o *ModelsWechatOrderDetail) HasMchid() bool`

HasMchid returns a boolean if a field has been set.

### GetPayer

`func (o *ModelsWechatOrderDetail) GetPayer() string`

GetPayer returns the Payer field if non-nil, zero value otherwise.

### GetPayerOk

`func (o *ModelsWechatOrderDetail) GetPayerOk() (*string, bool)`

GetPayerOk returns a tuple with the Payer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayer

`func (o *ModelsWechatOrderDetail) SetPayer(v string)`

SetPayer sets Payer field to given value.

### HasPayer

`func (o *ModelsWechatOrderDetail) HasPayer() bool`

HasPayer returns a boolean if a field has been set.

### GetRefreshStatus

`func (o *ModelsWechatOrderDetail) GetRefreshStatus() string`

GetRefreshStatus returns the RefreshStatus field if non-nil, zero value otherwise.

### GetRefreshStatusOk

`func (o *ModelsWechatOrderDetail) GetRefreshStatusOk() (*string, bool)`

GetRefreshStatusOk returns a tuple with the RefreshStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshStatus

`func (o *ModelsWechatOrderDetail) SetRefreshStatus(v string)`

SetRefreshStatus sets RefreshStatus field to given value.

### HasRefreshStatus

`func (o *ModelsWechatOrderDetail) HasRefreshStatus() bool`

HasRefreshStatus returns a boolean if a field has been set.

### GetRefundNo

`func (o *ModelsWechatOrderDetail) GetRefundNo() string`

GetRefundNo returns the RefundNo field if non-nil, zero value otherwise.

### GetRefundNoOk

`func (o *ModelsWechatOrderDetail) GetRefundNoOk() (*string, bool)`

GetRefundNoOk returns a tuple with the RefundNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundNo

`func (o *ModelsWechatOrderDetail) SetRefundNo(v string)`

SetRefundNo sets RefundNo field to given value.

### HasRefundNo

`func (o *ModelsWechatOrderDetail) HasRefundNo() bool`

HasRefundNo returns a boolean if a field has been set.

### GetSuccessTime

`func (o *ModelsWechatOrderDetail) GetSuccessTime() string`

GetSuccessTime returns the SuccessTime field if non-nil, zero value otherwise.

### GetSuccessTimeOk

`func (o *ModelsWechatOrderDetail) GetSuccessTimeOk() (*string, bool)`

GetSuccessTimeOk returns a tuple with the SuccessTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessTime

`func (o *ModelsWechatOrderDetail) SetSuccessTime(v string)`

SetSuccessTime sets SuccessTime field to given value.

### HasSuccessTime

`func (o *ModelsWechatOrderDetail) HasSuccessTime() bool`

HasSuccessTime returns a boolean if a field has been set.

### GetTradeNo

`func (o *ModelsWechatOrderDetail) GetTradeNo() string`

GetTradeNo returns the TradeNo field if non-nil, zero value otherwise.

### GetTradeNoOk

`func (o *ModelsWechatOrderDetail) GetTradeNoOk() (*string, bool)`

GetTradeNoOk returns a tuple with the TradeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeNo

`func (o *ModelsWechatOrderDetail) SetTradeNo(v string)`

SetTradeNo sets TradeNo field to given value.

### HasTradeNo

`func (o *ModelsWechatOrderDetail) HasTradeNo() bool`

HasTradeNo returns a boolean if a field has been set.

### GetTradeState

`func (o *ModelsWechatOrderDetail) GetTradeState() string`

GetTradeState returns the TradeState field if non-nil, zero value otherwise.

### GetTradeStateOk

`func (o *ModelsWechatOrderDetail) GetTradeStateOk() (*string, bool)`

GetTradeStateOk returns a tuple with the TradeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeState

`func (o *ModelsWechatOrderDetail) SetTradeState(v string)`

SetTradeState sets TradeState field to given value.

### HasTradeState

`func (o *ModelsWechatOrderDetail) HasTradeState() bool`

HasTradeState returns a boolean if a field has been set.

### GetTradeStateDesc

`func (o *ModelsWechatOrderDetail) GetTradeStateDesc() string`

GetTradeStateDesc returns the TradeStateDesc field if non-nil, zero value otherwise.

### GetTradeStateDescOk

`func (o *ModelsWechatOrderDetail) GetTradeStateDescOk() (*string, bool)`

GetTradeStateDescOk returns a tuple with the TradeStateDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeStateDesc

`func (o *ModelsWechatOrderDetail) SetTradeStateDesc(v string)`

SetTradeStateDesc sets TradeStateDesc field to given value.

### HasTradeStateDesc

`func (o *ModelsWechatOrderDetail) HasTradeStateDesc() bool`

HasTradeStateDesc returns a boolean if a field has been set.

### GetTradeType

`func (o *ModelsWechatOrderDetail) GetTradeType() string`

GetTradeType returns the TradeType field if non-nil, zero value otherwise.

### GetTradeTypeOk

`func (o *ModelsWechatOrderDetail) GetTradeTypeOk() (*string, bool)`

GetTradeTypeOk returns a tuple with the TradeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeType

`func (o *ModelsWechatOrderDetail) SetTradeType(v string)`

SetTradeType sets TradeType field to given value.

### HasTradeType

`func (o *ModelsWechatOrderDetail) HasTradeType() bool`

HasTradeType returns a boolean if a field has been set.

### GetTransactionId

`func (o *ModelsWechatOrderDetail) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ModelsWechatOrderDetail) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ModelsWechatOrderDetail) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *ModelsWechatOrderDetail) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsWechatOrderDetail) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsWechatOrderDetail) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsWechatOrderDetail) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsWechatOrderDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


