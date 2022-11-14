# Go API client for confluxpay

Conflux-Pay API documentation

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import confluxpay "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), confluxpay.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), confluxpay.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), confluxpay.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), confluxpay.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://127.0.0.1:8080/v0*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*OrdersApi* | [**Close**](docs/OrdersApi.md#close) | **Put** /orders/wechat/close/{trade_no} | close order
*OrdersApi* | [**MakeOrder**](docs/OrdersApi.md#makeorder) | **Post** /orders/wechat | Make Order
*OrdersApi* | [**QueryOrderSummary**](docs/OrdersApi.md#queryordersummary) | **Get** /orders/summary/{trade_no} | query order summary by trade no
*OrdersApi* | [**QueryWechatOrderDetail**](docs/OrdersApi.md#querywechatorderdetail) | **Get** /orders/wechat/{trade_no} | query order by trade no
*OrdersApi* | [**RefreshPayUrl**](docs/OrdersApi.md#refreshpayurl) | **Put** /orders/wechat/refresh-url/{trade_no} | refresh pay url
*OrdersApi* | [**Refund**](docs/OrdersApi.md#refund) | **Put** /orders/wechat/refund/{trade_no} | refund pay


## Documentation For Models

 - [CnsErrorsRainbowErrorDetailInfo](docs/CnsErrorsRainbowErrorDetailInfo.md)
 - [GormDeletedAt](docs/GormDeletedAt.md)
 - [ModelsOrder](docs/ModelsOrder.md)
 - [ModelsWechatOrderDetail](docs/ModelsWechatOrderDetail.md)
 - [ModelsWechatRefundDetail](docs/ModelsWechatRefundDetail.md)
 - [ServicesMakeOrderReq](docs/ServicesMakeOrderReq.md)
 - [ServicesMakeOrderResp](docs/ServicesMakeOrderResp.md)
 - [ServicesRefundReq](docs/ServicesRefundReq.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



