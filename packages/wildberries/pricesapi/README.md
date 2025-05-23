# Go API client for pricesapi

С помощью этих методов можно устанавливать цены и скидки. Максимум — 10 запросов за 6 секунд суммарно для всех методов раздела **Цены и скидки**.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 
- Package version: 1.0.0
- Generator version: 7.9.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import pricesapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `pricesapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), pricesapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `pricesapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), pricesapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `pricesapi.ContextOperationServerIndices` and `pricesapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), pricesapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), pricesapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**ApiV1CalendarPromotionsDetailsGet**](docs/DefaultApi.md#apiv1calendarpromotionsdetailsget) | **Get** /api/v1/calendar/promotions/details | Детальная информация по акциям
*DefaultApi* | [**ApiV1CalendarPromotionsGet**](docs/DefaultApi.md#apiv1calendarpromotionsget) | **Get** /api/v1/calendar/promotions | Список акций
*DefaultApi* | [**ApiV1CalendarPromotionsNomenclaturesGet**](docs/DefaultApi.md#apiv1calendarpromotionsnomenclaturesget) | **Get** /api/v1/calendar/promotions/nomenclatures | Список товаров для участия в акции
*DefaultApi* | [**ApiV1CalendarPromotionsUploadPost**](docs/DefaultApi.md#apiv1calendarpromotionsuploadpost) | **Post** /api/v1/calendar/promotions/upload | Добавить товар в акцию
*DefaultApi* | [**ApiV2BufferGoodsTaskGet**](docs/DefaultApi.md#apiv2buffergoodstaskget) | **Get** /api/v2/buffer/goods/task | Детализация необработанной загрузки
*DefaultApi* | [**ApiV2BufferTasksGet**](docs/DefaultApi.md#apiv2buffertasksget) | **Get** /api/v2/buffer/tasks | Состояние необработанной загрузки
*DefaultApi* | [**ApiV2HistoryGoodsTaskGet**](docs/DefaultApi.md#apiv2historygoodstaskget) | **Get** /api/v2/history/goods/task | Детализация обработанной загрузки
*DefaultApi* | [**ApiV2HistoryTasksGet**](docs/DefaultApi.md#apiv2historytasksget) | **Get** /api/v2/history/tasks | Состояние обработанной загрузки
*DefaultApi* | [**ApiV2ListGoodsFilterGet**](docs/DefaultApi.md#apiv2listgoodsfilterget) | **Get** /api/v2/list/goods/filter | Получить товары
*DefaultApi* | [**ApiV2ListGoodsSizeNmGet**](docs/DefaultApi.md#apiv2listgoodssizenmget) | **Get** /api/v2/list/goods/size/nm | Получить размеры товара
*DefaultApi* | [**ApiV2QuarantineGoodsGet**](docs/DefaultApi.md#apiv2quarantinegoodsget) | **Get** /api/v2/quarantine/goods | Получить товары в карантине
*DefaultApi* | [**ApiV2UploadTaskPost**](docs/DefaultApi.md#apiv2uploadtaskpost) | **Post** /api/v2/upload/task | Установить цены и скидки
*DefaultApi* | [**ApiV2UploadTaskSizePost**](docs/DefaultApi.md#apiv2uploadtasksizepost) | **Post** /api/v2/upload/task/size | Установить цены для размеров


## Documentation For Models

 - [ApiV1CalendarPromotionsDetailsGet200Response](docs/ApiV1CalendarPromotionsDetailsGet200Response.md)
 - [ApiV1CalendarPromotionsDetailsGet200ResponseData](docs/ApiV1CalendarPromotionsDetailsGet200ResponseData.md)
 - [ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner](docs/ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner.md)
 - [ApiV1CalendarPromotionsGet200Response](docs/ApiV1CalendarPromotionsGet200Response.md)
 - [ApiV1CalendarPromotionsGet200ResponseData](docs/ApiV1CalendarPromotionsGet200ResponseData.md)
 - [ApiV1CalendarPromotionsGet200ResponseDataPromotionsInner](docs/ApiV1CalendarPromotionsGet200ResponseDataPromotionsInner.md)
 - [ApiV1CalendarPromotionsGet400Response](docs/ApiV1CalendarPromotionsGet400Response.md)
 - [ApiV1CalendarPromotionsNomenclaturesGet200Response](docs/ApiV1CalendarPromotionsNomenclaturesGet200Response.md)
 - [ApiV1CalendarPromotionsNomenclaturesGet200ResponseData](docs/ApiV1CalendarPromotionsNomenclaturesGet200ResponseData.md)
 - [ApiV1CalendarPromotionsNomenclaturesGet422Response](docs/ApiV1CalendarPromotionsNomenclaturesGet422Response.md)
 - [ApiV1CalendarPromotionsUploadPost200Response](docs/ApiV1CalendarPromotionsUploadPost200Response.md)
 - [ApiV1CalendarPromotionsUploadPost200ResponseData](docs/ApiV1CalendarPromotionsUploadPost200ResponseData.md)
 - [ApiV1CalendarPromotionsUploadPost422Response](docs/ApiV1CalendarPromotionsUploadPost422Response.md)
 - [ApiV1CalendarPromotionsUploadPostRequest](docs/ApiV1CalendarPromotionsUploadPostRequest.md)
 - [ApiV1CalendarPromotionsUploadPostRequestData](docs/ApiV1CalendarPromotionsUploadPostRequestData.md)
 - [ApiV2BufferGoodsTaskGet200Response](docs/ApiV2BufferGoodsTaskGet200Response.md)
 - [ApiV2BufferGoodsTaskGet200ResponseData](docs/ApiV2BufferGoodsTaskGet200ResponseData.md)
 - [ApiV2BufferTasksGet200Response](docs/ApiV2BufferTasksGet200Response.md)
 - [ApiV2HistoryGoodsTaskGet200Response](docs/ApiV2HistoryGoodsTaskGet200Response.md)
 - [ApiV2HistoryGoodsTaskGet200ResponseData](docs/ApiV2HistoryGoodsTaskGet200ResponseData.md)
 - [ApiV2HistoryTasksGet200Response](docs/ApiV2HistoryTasksGet200Response.md)
 - [ApiV2HistoryTasksGet400Response](docs/ApiV2HistoryTasksGet400Response.md)
 - [ApiV2HistoryTasksGet4XXResponse](docs/ApiV2HistoryTasksGet4XXResponse.md)
 - [ApiV2HistoryTasksGet5XXResponse](docs/ApiV2HistoryTasksGet5XXResponse.md)
 - [ApiV2ListGoodsFilterGet200Response](docs/ApiV2ListGoodsFilterGet200Response.md)
 - [ApiV2ListGoodsFilterGet200ResponseData](docs/ApiV2ListGoodsFilterGet200ResponseData.md)
 - [ApiV2ListGoodsSizeNmGet200Response](docs/ApiV2ListGoodsSizeNmGet200Response.md)
 - [ApiV2ListGoodsSizeNmGet200ResponseData](docs/ApiV2ListGoodsSizeNmGet200ResponseData.md)
 - [ApiV2QuarantineGoodsGet200Response](docs/ApiV2QuarantineGoodsGet200Response.md)
 - [ApiV2QuarantineGoodsGet200ResponseData](docs/ApiV2QuarantineGoodsGet200ResponseData.md)
 - [ApiV2UploadTaskPostRequest](docs/ApiV2UploadTaskPostRequest.md)
 - [ApiV2UploadTaskSizePostRequest](docs/ApiV2UploadTaskSizePostRequest.md)
 - [Good](docs/Good.md)
 - [GoodBufferHistory](docs/GoodBufferHistory.md)
 - [GoodHistory](docs/GoodHistory.md)
 - [GoodsList](docs/GoodsList.md)
 - [GoodsListSizesInner](docs/GoodsListSizesInner.md)
 - [PromotionsGoodsList](docs/PromotionsGoodsList.md)
 - [QuarantineGoods](docs/QuarantineGoods.md)
 - [ResponseError](docs/ResponseError.md)
 - [SizeGood](docs/SizeGood.md)
 - [SizeGoodReq](docs/SizeGoodReq.md)
 - [SupplierTaskMetadata](docs/SupplierTaskMetadata.md)
 - [SupplierTaskMetadataBuffer](docs/SupplierTaskMetadataBuffer.md)
 - [TaskAlreadyExistsError](docs/TaskAlreadyExistsError.md)
 - [TaskAlreadyExistsErrorData](docs/TaskAlreadyExistsErrorData.md)
 - [TaskCreated](docs/TaskCreated.md)
 - [TaskCreatedData](docs/TaskCreatedData.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### HeaderApiKey

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: HeaderApiKey and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		pricesapi.ContextAPIKeys,
		map[string]pricesapi.APIKey{
			"HeaderApiKey": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


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



