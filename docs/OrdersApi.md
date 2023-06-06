# \OrdersApi

All URIs are relative to *http://127.0.0.1:8080/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Close**](OrdersApi.md#Close) | **Put** /orders/close/{trade_no} | close order
[**MakeOrder**](OrdersApi.md#MakeOrder) | **Post** /orders | Make Order
[**QueryOrder**](OrdersApi.md#QueryOrder) | **Get** /orders/{trade_no} | query order by trade no
[**RefreshPayUrl**](OrdersApi.md#RefreshPayUrl) | **Put** /orders/refresh-url/{trade_no} | refresh pay url
[**Refund**](OrdersApi.md#Refund) | **Put** /orders/refund/{trade_no} | refund pay



## Close

> ModelsOrderCore Close(ctx, tradeNo).Execute()

close order



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
    tradeNo := "tradeNo_example" // string | trade no

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.Close(context.Background(), tradeNo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.Close``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Close`: ModelsOrderCore
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.Close`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tradeNo** | **string** | trade no | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsOrderCore**](ModelsOrderCore.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MakeOrder

> ModelsOrder MakeOrder(ctx).MakeOrdReq(makeOrdReq).Execute()

Make Order



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
    makeOrdReq := *openapiclient.NewServicesMakeOrderReq(int32(123), "AppName_example", "Description_example", int32(123), "TradeProvider_example", "TradeType_example") // ServicesMakeOrderReq | make_wechat_order_req

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.MakeOrder(context.Background()).MakeOrdReq(makeOrdReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.MakeOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MakeOrder`: ModelsOrder
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.MakeOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMakeOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **makeOrdReq** | [**ServicesMakeOrderReq**](ServicesMakeOrderReq.md) | make_wechat_order_req | 

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


## QueryOrder

> ModelsOrder QueryOrder(ctx, tradeNo).Execute()

query order by trade no



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
    tradeNo := "tradeNo_example" // string | trade no

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.QueryOrder(context.Background(), tradeNo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.QueryOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryOrder`: ModelsOrder
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.QueryOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tradeNo** | **string** | trade no | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryOrderRequest struct via the builder pattern


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


## RefreshPayUrl

> ModelsOrder RefreshPayUrl(ctx, tradeNo).Execute()

refresh pay url



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
    tradeNo := "tradeNo_example" // string | trade no

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.RefreshPayUrl(context.Background(), tradeNo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.RefreshPayUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshPayUrl`: ModelsOrder
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

[**ModelsOrder**](ModelsOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Refund

> ModelsOrderCore Refund(ctx, tradeNo).RefundReq(refundReq).Execute()

refund pay



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
    tradeNo := "tradeNo_example" // string | trade no
    refundReq := *openapiclient.NewServicesRefundReq("Reason_example") // ServicesRefundReq | refund_req

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.Refund(context.Background(), tradeNo).RefundReq(refundReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.Refund``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Refund`: ModelsOrderCore
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.Refund`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tradeNo** | **string** | trade no | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refundReq** | [**ServicesRefundReq**](ServicesRefundReq.md) | refund_req | 

### Return type

[**ModelsOrderCore**](ModelsOrderCore.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

