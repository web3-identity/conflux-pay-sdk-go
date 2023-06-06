# ModelsCmbRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccNbr** | Pointer to **string** |  | [optional] 
**AutFlg** | Pointer to **string** | useless | [optional] 
**CcyNbr** | Pointer to **string** | expected to be rmb | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to [**GormDeletedAt**](GormDeletedAt.md) |  | [optional] 
**DmaNam** | Pointer to **string** | sub unit name | [optional] 
**DmaNbr** | Pointer to **string** | sub unit number | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**NarInn** | Pointer to **string** | useless | [optional] 
**RpyAcc** | Pointer to **string** | useless | [optional] 
**RpyNam** | Pointer to **string** | useless | [optional] 
**TrxAmt** | Pointer to **float32** |  | [optional] 
**TrxDat** | Pointer to **string** |  | [optional] 
**TrxDir** | Pointer to **string** | tx direction | [optional] 
**TrxNbr** | Pointer to **string** | unique | [optional] 
**TrxTim** | Pointer to **string** |  | [optional] 
**TrxTxt** | Pointer to **string** | txt that sender appended | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsCmbRecord

`func NewModelsCmbRecord() *ModelsCmbRecord`

NewModelsCmbRecord instantiates a new ModelsCmbRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCmbRecordWithDefaults

`func NewModelsCmbRecordWithDefaults() *ModelsCmbRecord`

NewModelsCmbRecordWithDefaults instantiates a new ModelsCmbRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccNbr

`func (o *ModelsCmbRecord) GetAccNbr() string`

GetAccNbr returns the AccNbr field if non-nil, zero value otherwise.

### GetAccNbrOk

`func (o *ModelsCmbRecord) GetAccNbrOk() (*string, bool)`

GetAccNbrOk returns a tuple with the AccNbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccNbr

`func (o *ModelsCmbRecord) SetAccNbr(v string)`

SetAccNbr sets AccNbr field to given value.

### HasAccNbr

`func (o *ModelsCmbRecord) HasAccNbr() bool`

HasAccNbr returns a boolean if a field has been set.

### GetAutFlg

`func (o *ModelsCmbRecord) GetAutFlg() string`

GetAutFlg returns the AutFlg field if non-nil, zero value otherwise.

### GetAutFlgOk

`func (o *ModelsCmbRecord) GetAutFlgOk() (*string, bool)`

GetAutFlgOk returns a tuple with the AutFlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutFlg

`func (o *ModelsCmbRecord) SetAutFlg(v string)`

SetAutFlg sets AutFlg field to given value.

### HasAutFlg

`func (o *ModelsCmbRecord) HasAutFlg() bool`

HasAutFlg returns a boolean if a field has been set.

### GetCcyNbr

`func (o *ModelsCmbRecord) GetCcyNbr() string`

GetCcyNbr returns the CcyNbr field if non-nil, zero value otherwise.

### GetCcyNbrOk

`func (o *ModelsCmbRecord) GetCcyNbrOk() (*string, bool)`

GetCcyNbrOk returns a tuple with the CcyNbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcyNbr

`func (o *ModelsCmbRecord) SetCcyNbr(v string)`

SetCcyNbr sets CcyNbr field to given value.

### HasCcyNbr

`func (o *ModelsCmbRecord) HasCcyNbr() bool`

HasCcyNbr returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsCmbRecord) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsCmbRecord) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsCmbRecord) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsCmbRecord) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ModelsCmbRecord) GetDeletedAt() GormDeletedAt`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ModelsCmbRecord) GetDeletedAtOk() (*GormDeletedAt, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ModelsCmbRecord) SetDeletedAt(v GormDeletedAt)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ModelsCmbRecord) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDmaNam

`func (o *ModelsCmbRecord) GetDmaNam() string`

GetDmaNam returns the DmaNam field if non-nil, zero value otherwise.

### GetDmaNamOk

`func (o *ModelsCmbRecord) GetDmaNamOk() (*string, bool)`

GetDmaNamOk returns a tuple with the DmaNam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmaNam

`func (o *ModelsCmbRecord) SetDmaNam(v string)`

SetDmaNam sets DmaNam field to given value.

### HasDmaNam

`func (o *ModelsCmbRecord) HasDmaNam() bool`

HasDmaNam returns a boolean if a field has been set.

### GetDmaNbr

`func (o *ModelsCmbRecord) GetDmaNbr() string`

GetDmaNbr returns the DmaNbr field if non-nil, zero value otherwise.

### GetDmaNbrOk

`func (o *ModelsCmbRecord) GetDmaNbrOk() (*string, bool)`

GetDmaNbrOk returns a tuple with the DmaNbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmaNbr

`func (o *ModelsCmbRecord) SetDmaNbr(v string)`

SetDmaNbr sets DmaNbr field to given value.

### HasDmaNbr

`func (o *ModelsCmbRecord) HasDmaNbr() bool`

HasDmaNbr returns a boolean if a field has been set.

### GetId

`func (o *ModelsCmbRecord) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsCmbRecord) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsCmbRecord) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsCmbRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNarInn

`func (o *ModelsCmbRecord) GetNarInn() string`

GetNarInn returns the NarInn field if non-nil, zero value otherwise.

### GetNarInnOk

`func (o *ModelsCmbRecord) GetNarInnOk() (*string, bool)`

GetNarInnOk returns a tuple with the NarInn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNarInn

`func (o *ModelsCmbRecord) SetNarInn(v string)`

SetNarInn sets NarInn field to given value.

### HasNarInn

`func (o *ModelsCmbRecord) HasNarInn() bool`

HasNarInn returns a boolean if a field has been set.

### GetRpyAcc

`func (o *ModelsCmbRecord) GetRpyAcc() string`

GetRpyAcc returns the RpyAcc field if non-nil, zero value otherwise.

### GetRpyAccOk

`func (o *ModelsCmbRecord) GetRpyAccOk() (*string, bool)`

GetRpyAccOk returns a tuple with the RpyAcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpyAcc

`func (o *ModelsCmbRecord) SetRpyAcc(v string)`

SetRpyAcc sets RpyAcc field to given value.

### HasRpyAcc

`func (o *ModelsCmbRecord) HasRpyAcc() bool`

HasRpyAcc returns a boolean if a field has been set.

### GetRpyNam

`func (o *ModelsCmbRecord) GetRpyNam() string`

GetRpyNam returns the RpyNam field if non-nil, zero value otherwise.

### GetRpyNamOk

`func (o *ModelsCmbRecord) GetRpyNamOk() (*string, bool)`

GetRpyNamOk returns a tuple with the RpyNam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpyNam

`func (o *ModelsCmbRecord) SetRpyNam(v string)`

SetRpyNam sets RpyNam field to given value.

### HasRpyNam

`func (o *ModelsCmbRecord) HasRpyNam() bool`

HasRpyNam returns a boolean if a field has been set.

### GetTrxAmt

`func (o *ModelsCmbRecord) GetTrxAmt() float32`

GetTrxAmt returns the TrxAmt field if non-nil, zero value otherwise.

### GetTrxAmtOk

`func (o *ModelsCmbRecord) GetTrxAmtOk() (*float32, bool)`

GetTrxAmtOk returns a tuple with the TrxAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrxAmt

`func (o *ModelsCmbRecord) SetTrxAmt(v float32)`

SetTrxAmt sets TrxAmt field to given value.

### HasTrxAmt

`func (o *ModelsCmbRecord) HasTrxAmt() bool`

HasTrxAmt returns a boolean if a field has been set.

### GetTrxDat

`func (o *ModelsCmbRecord) GetTrxDat() string`

GetTrxDat returns the TrxDat field if non-nil, zero value otherwise.

### GetTrxDatOk

`func (o *ModelsCmbRecord) GetTrxDatOk() (*string, bool)`

GetTrxDatOk returns a tuple with the TrxDat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrxDat

`func (o *ModelsCmbRecord) SetTrxDat(v string)`

SetTrxDat sets TrxDat field to given value.

### HasTrxDat

`func (o *ModelsCmbRecord) HasTrxDat() bool`

HasTrxDat returns a boolean if a field has been set.

### GetTrxDir

`func (o *ModelsCmbRecord) GetTrxDir() string`

GetTrxDir returns the TrxDir field if non-nil, zero value otherwise.

### GetTrxDirOk

`func (o *ModelsCmbRecord) GetTrxDirOk() (*string, bool)`

GetTrxDirOk returns a tuple with the TrxDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrxDir

`func (o *ModelsCmbRecord) SetTrxDir(v string)`

SetTrxDir sets TrxDir field to given value.

### HasTrxDir

`func (o *ModelsCmbRecord) HasTrxDir() bool`

HasTrxDir returns a boolean if a field has been set.

### GetTrxNbr

`func (o *ModelsCmbRecord) GetTrxNbr() string`

GetTrxNbr returns the TrxNbr field if non-nil, zero value otherwise.

### GetTrxNbrOk

`func (o *ModelsCmbRecord) GetTrxNbrOk() (*string, bool)`

GetTrxNbrOk returns a tuple with the TrxNbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrxNbr

`func (o *ModelsCmbRecord) SetTrxNbr(v string)`

SetTrxNbr sets TrxNbr field to given value.

### HasTrxNbr

`func (o *ModelsCmbRecord) HasTrxNbr() bool`

HasTrxNbr returns a boolean if a field has been set.

### GetTrxTim

`func (o *ModelsCmbRecord) GetTrxTim() string`

GetTrxTim returns the TrxTim field if non-nil, zero value otherwise.

### GetTrxTimOk

`func (o *ModelsCmbRecord) GetTrxTimOk() (*string, bool)`

GetTrxTimOk returns a tuple with the TrxTim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrxTim

`func (o *ModelsCmbRecord) SetTrxTim(v string)`

SetTrxTim sets TrxTim field to given value.

### HasTrxTim

`func (o *ModelsCmbRecord) HasTrxTim() bool`

HasTrxTim returns a boolean if a field has been set.

### GetTrxTxt

`func (o *ModelsCmbRecord) GetTrxTxt() string`

GetTrxTxt returns the TrxTxt field if non-nil, zero value otherwise.

### GetTrxTxtOk

`func (o *ModelsCmbRecord) GetTrxTxtOk() (*string, bool)`

GetTrxTxtOk returns a tuple with the TrxTxt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrxTxt

`func (o *ModelsCmbRecord) SetTrxTxt(v string)`

SetTrxTxt sets TrxTxt field to given value.

### HasTrxTxt

`func (o *ModelsCmbRecord) HasTrxTxt() bool`

HasTrxTxt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsCmbRecord) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsCmbRecord) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsCmbRecord) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsCmbRecord) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


