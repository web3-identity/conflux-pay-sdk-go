# \CmbApi

All URIs are relative to *http://127.0.0.1:8080/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUnitAccount**](CmbApi.md#AddUnitAccount) | **Post** /cmb/unit-account | Add a unit account
[**QueryHistoryCmbRecords**](CmbApi.md#QueryHistoryCmbRecords) | **Get** /cmb/history | 查询历史交易
[**QueryRecentCmbRecords**](CmbApi.md#QueryRecentCmbRecords) | **Get** /cmb/history/recent | 查询昨天和今天汇入的交易
[**SetUnitAccountRelation**](CmbApi.md#SetUnitAccountRelation) | **Post** /cmb/unit-account/relation | Set a related bank account of a unit account



## AddUnitAccount

> ControllersAddUnitAccountResult AddUnitAccount(ctx).AddUnitAccountReq(addUnitAccountReq).Execute()

Add a unit account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    addUnitAccountReq := *openapiclient.NewControllersAddUnitAccountReq("UnitAccountName_example", "UnitAccountNbr_example") // ControllersAddUnitAccountReq | add_unit_account_req

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CmbApi.AddUnitAccount(context.Background()).AddUnitAccountReq(addUnitAccountReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CmbApi.AddUnitAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddUnitAccount`: ControllersAddUnitAccountResult
    fmt.Fprintf(os.Stdout, "Response from `CmbApi.AddUnitAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddUnitAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addUnitAccountReq** | [**ControllersAddUnitAccountReq**](ControllersAddUnitAccountReq.md) | add_unit_account_req | 

### Return type

[**ControllersAddUnitAccountResult**](ControllersAddUnitAccountResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryHistoryCmbRecords

> []ModelsCmbRecord QueryHistoryCmbRecords(ctx).Limit(limit).Offset(offset).UnitAccountNbr(unitAccountNbr).TransactionDate(transactionDate).TransactionDirection(transactionDirection).Execute()

查询历史交易



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    limit := int32(56) // int32 | limit
    offset := int32(56) // int32 | offset
    unitAccountNbr := "unitAccountNbr_example" // string | specified unit account number
    transactionDate := "transactionDate_example" // string | specified date, e.g. 20230523
    transactionDirection := "transactionDirection_example" // string | transaction direction, C for recieve and D for out

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CmbApi.QueryHistoryCmbRecords(context.Background()).Limit(limit).Offset(offset).UnitAccountNbr(unitAccountNbr).TransactionDate(transactionDate).TransactionDirection(transactionDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CmbApi.QueryHistoryCmbRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryHistoryCmbRecords`: []ModelsCmbRecord
    fmt.Fprintf(os.Stdout, "Response from `CmbApi.QueryHistoryCmbRecords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryHistoryCmbRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | limit | 
 **offset** | **int32** | offset | 
 **unitAccountNbr** | **string** | specified unit account number | 
 **transactionDate** | **string** | specified date, e.g. 20230523 | 
 **transactionDirection** | **string** | transaction direction, C for recieve and D for out | 

### Return type

[**[]ModelsCmbRecord**](ModelsCmbRecord.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryRecentCmbRecords

> []ModelsCmbRecord QueryRecentCmbRecords(ctx).Limit(limit).Offset(offset).Execute()

查询昨天和今天汇入的交易



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    limit := int32(56) // int32 | limit
    offset := int32(56) // int32 | offset

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CmbApi.QueryRecentCmbRecords(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CmbApi.QueryRecentCmbRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryRecentCmbRecords`: []ModelsCmbRecord
    fmt.Fprintf(os.Stdout, "Response from `CmbApi.QueryRecentCmbRecords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryRecentCmbRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | limit | 
 **offset** | **int32** | offset | 

### Return type

[**[]ModelsCmbRecord**](ModelsCmbRecord.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetUnitAccountRelation

> ControllersSetUnitAccountRelationResult SetUnitAccountRelation(ctx).SetUnitAccountRelationReq(setUnitAccountRelationReq).Execute()

Set a related bank account of a unit account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    setUnitAccountRelationReq := *openapiclient.NewControllersSetUnitAccountRelationReq("BankAccountNbr_example", "UnitAccountNbr_example") // ControllersSetUnitAccountRelationReq | set_unit_account_relation_req

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CmbApi.SetUnitAccountRelation(context.Background()).SetUnitAccountRelationReq(setUnitAccountRelationReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CmbApi.SetUnitAccountRelation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetUnitAccountRelation`: ControllersSetUnitAccountRelationResult
    fmt.Fprintf(os.Stdout, "Response from `CmbApi.SetUnitAccountRelation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetUnitAccountRelationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setUnitAccountRelationReq** | [**ControllersSetUnitAccountRelationReq**](ControllersSetUnitAccountRelationReq.md) | set_unit_account_relation_req | 

### Return type

[**ControllersSetUnitAccountRelationResult**](ControllersSetUnitAccountRelationResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

