# \OrdersApi

All URIs are relative to *http://127.0.0.1:8080/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MakeOrder**](OrdersApi.md#MakeOrder) | **Post** /orders/wechat | Make Order
[**QueryOrderSummary**](OrdersApi.md#QueryOrderSummary) | **Get** /orders/summary/{trade_no} | query order summary by trade no
[**QueryWechatOrderDetail**](OrdersApi.md#QueryWechatOrderDetail) | **Get** /orders/wechat/{trade_no} | query order by trade no
[**RefreshPayUrl**](OrdersApi.md#RefreshPayUrl) | **Post** /orders/wechat/refresh-url/{trade_no} | refresh pay url



## MakeOrder

> ServicesMakeOrderResp MakeOrder(ctx).WecahtOrdReq(wecahtOrdReq).Execute()

Make Order



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    wecahtOrdReq := *openapiclient.NewServicesMakeWechatOrderReq() // ServicesMakeWechatOrderReq | make_wechat_order_req

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.MakeOrder(context.Background()).WecahtOrdReq(wecahtOrdReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.MakeOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MakeOrder`: ServicesMakeOrderResp
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.MakeOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMakeOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wecahtOrdReq** | [**ServicesMakeWechatOrderReq**](ServicesMakeWechatOrderReq.md) | make_wechat_order_req | 

### Return type

[**ServicesMakeOrderResp**](ServicesMakeOrderResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryOrderSummary

> ModelsOrder QueryOrderSummary(ctx, tradeNo).Execute()

query order summary by trade no



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tradeNo := "tradeNo_example" // string | trade no

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.QueryOrderSummary(context.Background(), tradeNo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.QueryOrderSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryOrderSummary`: ModelsOrder
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.QueryOrderSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tradeNo** | **string** | trade no | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryOrderSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsOrder**](ModelsOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryWechatOrderDetail

> ModelsWechatOrderDetail QueryWechatOrderDetail(ctx, tradeNo).Execute()

query order by trade no



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tradeNo := "tradeNo_example" // string | trade no

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.QueryWechatOrderDetail(context.Background(), tradeNo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.QueryWechatOrderDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryWechatOrderDetail`: ModelsWechatOrderDetail
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.QueryWechatOrderDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tradeNo** | **string** | trade no | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryWechatOrderDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsWechatOrderDetail**](ModelsWechatOrderDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshPayUrl

> ServicesMakeOrderResp RefreshPayUrl(ctx, tradeNo).Execute()

refresh pay url



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tradeNo := "tradeNo_example" // string | trade no

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.RefreshPayUrl(context.Background(), tradeNo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.RefreshPayUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshPayUrl`: ServicesMakeOrderResp
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.RefreshPayUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tradeNo** | **string** | trade no | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshPayUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServicesMakeOrderResp**](ServicesMakeOrderResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

