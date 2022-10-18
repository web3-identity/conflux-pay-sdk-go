# ModelsOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | 单位为分 | [optional] 
**AppName** | Pointer to **string** |  | [optional] 
**CodeUrl** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to [**GormDeletedAt**](GormDeletedAt.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**H5Url** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**TimeExpire** | Pointer to **string** |  | [optional] 
**TradeNo** | Pointer to **string** |  | [optional] 
**TradeProvider** | Pointer to **int32** |  | [optional] 
**TradeState** | Pointer to **int32** |  | [optional] 
**TradeType** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsOrder

`func NewModelsOrder() *ModelsOrder`

NewModelsOrder instantiates a new ModelsOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsOrderWithDefaults

`func NewModelsOrderWithDefaults() *ModelsOrder`

NewModelsOrderWithDefaults instantiates a new ModelsOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ModelsOrder) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ModelsOrder) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ModelsOrder) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ModelsOrder) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAppName

`func (o *ModelsOrder) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *ModelsOrder) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *ModelsOrder) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *ModelsOrder) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetCodeUrl

`func (o *ModelsOrder) GetCodeUrl() string`

GetCodeUrl returns the CodeUrl field if non-nil, zero value otherwise.

### GetCodeUrlOk

`func (o *ModelsOrder) GetCodeUrlOk() (*string, bool)`

GetCodeUrlOk returns a tuple with the CodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeUrl

`func (o *ModelsOrder) SetCodeUrl(v string)`

SetCodeUrl sets CodeUrl field to given value.

### HasCodeUrl

`func (o *ModelsOrder) HasCodeUrl() bool`

HasCodeUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsOrder) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsOrder) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsOrder) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsOrder) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ModelsOrder) GetDeletedAt() GormDeletedAt`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ModelsOrder) GetDeletedAtOk() (*GormDeletedAt, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ModelsOrder) SetDeletedAt(v GormDeletedAt)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ModelsOrder) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDescription

`func (o *ModelsOrder) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelsOrder) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelsOrder) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelsOrder) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetH5Url

`func (o *ModelsOrder) GetH5Url() string`

GetH5Url returns the H5Url field if non-nil, zero value otherwise.

### GetH5UrlOk

`func (o *ModelsOrder) GetH5UrlOk() (*string, bool)`

GetH5UrlOk returns a tuple with the H5Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH5Url

`func (o *ModelsOrder) SetH5Url(v string)`

SetH5Url sets H5Url field to given value.

### HasH5Url

`func (o *ModelsOrder) HasH5Url() bool`

HasH5Url returns a boolean if a field has been set.

### GetId

`func (o *ModelsOrder) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsOrder) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsOrder) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsOrder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimeExpire

`func (o *ModelsOrder) GetTimeExpire() string`

GetTimeExpire returns the TimeExpire field if non-nil, zero value otherwise.

### GetTimeExpireOk

`func (o *ModelsOrder) GetTimeExpireOk() (*string, bool)`

GetTimeExpireOk returns a tuple with the TimeExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeExpire

`func (o *ModelsOrder) SetTimeExpire(v string)`

SetTimeExpire sets TimeExpire field to given value.

### HasTimeExpire

`func (o *ModelsOrder) HasTimeExpire() bool`

HasTimeExpire returns a boolean if a field has been set.

### GetTradeNo

`func (o *ModelsOrder) GetTradeNo() string`

GetTradeNo returns the TradeNo field if non-nil, zero value otherwise.

### GetTradeNoOk

`func (o *ModelsOrder) GetTradeNoOk() (*string, bool)`

GetTradeNoOk returns a tuple with the TradeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeNo

`func (o *ModelsOrder) SetTradeNo(v string)`

SetTradeNo sets TradeNo field to given value.

### HasTradeNo

`func (o *ModelsOrder) HasTradeNo() bool`

HasTradeNo returns a boolean if a field has been set.

### GetTradeProvider

`func (o *ModelsOrder) GetTradeProvider() int32`

GetTradeProvider returns the TradeProvider field if non-nil, zero value otherwise.

### GetTradeProviderOk

`func (o *ModelsOrder) GetTradeProviderOk() (*int32, bool)`

GetTradeProviderOk returns a tuple with the TradeProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeProvider

`func (o *ModelsOrder) SetTradeProvider(v int32)`

SetTradeProvider sets TradeProvider field to given value.

### HasTradeProvider

`func (o *ModelsOrder) HasTradeProvider() bool`

HasTradeProvider returns a boolean if a field has been set.

### GetTradeState

`func (o *ModelsOrder) GetTradeState() int32`

GetTradeState returns the TradeState field if non-nil, zero value otherwise.

### GetTradeStateOk

`func (o *ModelsOrder) GetTradeStateOk() (*int32, bool)`

GetTradeStateOk returns a tuple with the TradeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeState

`func (o *ModelsOrder) SetTradeState(v int32)`

SetTradeState sets TradeState field to given value.

### HasTradeState

`func (o *ModelsOrder) HasTradeState() bool`

HasTradeState returns a boolean if a field has been set.

### GetTradeType

`func (o *ModelsOrder) GetTradeType() int32`

GetTradeType returns the TradeType field if non-nil, zero value otherwise.

### GetTradeTypeOk

`func (o *ModelsOrder) GetTradeTypeOk() (*int32, bool)`

GetTradeTypeOk returns a tuple with the TradeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeType

`func (o *ModelsOrder) SetTradeType(v int32)`

SetTradeType sets TradeType field to given value.

### HasTradeType

`func (o *ModelsOrder) HasTradeType() bool`

HasTradeType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsOrder) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsOrder) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsOrder) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsOrder) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


