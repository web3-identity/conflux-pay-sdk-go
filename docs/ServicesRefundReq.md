# ServicesRefundReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifyUrl** | Pointer to **string** |  | [optional] 
**Reason** | **string** |  | 

## Methods

### NewServicesRefundReq

`func NewServicesRefundReq(reason string, ) *ServicesRefundReq`

NewServicesRefundReq instantiates a new ServicesRefundReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesRefundReqWithDefaults

`func NewServicesRefundReqWithDefaults() *ServicesRefundReq`

NewServicesRefundReqWithDefaults instantiates a new ServicesRefundReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifyUrl

`func (o *ServicesRefundReq) GetNotifyUrl() string`

GetNotifyUrl returns the NotifyUrl field if non-nil, zero value otherwise.

### GetNotifyUrlOk

`func (o *ServicesRefundReq) GetNotifyUrlOk() (*string, bool)`

GetNotifyUrlOk returns a tuple with the NotifyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUrl

`func (o *ServicesRefundReq) SetNotifyUrl(v string)`

SetNotifyUrl sets NotifyUrl field to given value.

### HasNotifyUrl

`func (o *ServicesRefundReq) HasNotifyUrl() bool`

HasNotifyUrl returns a boolean if a field has been set.

### GetReason

`func (o *ServicesRefundReq) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ServicesRefundReq) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ServicesRefundReq) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


