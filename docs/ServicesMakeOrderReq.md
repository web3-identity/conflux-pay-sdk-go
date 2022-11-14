# ServicesMakeOrderReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** |  | 
**Description** | **string** |  | 
**NotifyUrl** | Pointer to **string** |  | [optional] 
**TimeExpire** | **int32** |  | 
**TradeType** | **string** |  | 

## Methods

### NewServicesMakeOrderReq

`func NewServicesMakeOrderReq(amount int32, description string, timeExpire int32, tradeType string, ) *ServicesMakeOrderReq`

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


