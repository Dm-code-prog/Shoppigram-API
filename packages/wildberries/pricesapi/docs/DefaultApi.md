# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1CalendarPromotionsDetailsGet**](DefaultApi.md#ApiV1CalendarPromotionsDetailsGet) | **Get** /api/v1/calendar/promotions/details | Детальная информация по акциям
[**ApiV1CalendarPromotionsGet**](DefaultApi.md#ApiV1CalendarPromotionsGet) | **Get** /api/v1/calendar/promotions | Список акций
[**ApiV1CalendarPromotionsNomenclaturesGet**](DefaultApi.md#ApiV1CalendarPromotionsNomenclaturesGet) | **Get** /api/v1/calendar/promotions/nomenclatures | Список товаров для участия в акции
[**ApiV1CalendarPromotionsUploadPost**](DefaultApi.md#ApiV1CalendarPromotionsUploadPost) | **Post** /api/v1/calendar/promotions/upload | Добавить товар в акцию
[**ApiV2BufferGoodsTaskGet**](DefaultApi.md#ApiV2BufferGoodsTaskGet) | **Get** /api/v2/buffer/goods/task | Детализация необработанной загрузки
[**ApiV2BufferTasksGet**](DefaultApi.md#ApiV2BufferTasksGet) | **Get** /api/v2/buffer/tasks | Состояние необработанной загрузки
[**ApiV2HistoryGoodsTaskGet**](DefaultApi.md#ApiV2HistoryGoodsTaskGet) | **Get** /api/v2/history/goods/task | Детализация обработанной загрузки
[**ApiV2HistoryTasksGet**](DefaultApi.md#ApiV2HistoryTasksGet) | **Get** /api/v2/history/tasks | Состояние обработанной загрузки
[**ApiV2ListGoodsFilterGet**](DefaultApi.md#ApiV2ListGoodsFilterGet) | **Get** /api/v2/list/goods/filter | Получить товары
[**ApiV2ListGoodsSizeNmGet**](DefaultApi.md#ApiV2ListGoodsSizeNmGet) | **Get** /api/v2/list/goods/size/nm | Получить размеры товара
[**ApiV2QuarantineGoodsGet**](DefaultApi.md#ApiV2QuarantineGoodsGet) | **Get** /api/v2/quarantine/goods | Получить товары в карантине
[**ApiV2UploadTaskPost**](DefaultApi.md#ApiV2UploadTaskPost) | **Post** /api/v2/upload/task | Установить цены и скидки
[**ApiV2UploadTaskSizePost**](DefaultApi.md#ApiV2UploadTaskSizePost) | **Post** /api/v2/upload/task/size | Установить цены для размеров



## ApiV1CalendarPromotionsDetailsGet

> ApiV1CalendarPromotionsDetailsGet200Response ApiV1CalendarPromotionsDetailsGet(ctx).PromotionIDs(promotionIDs).Execute()

Детальная информация по акциям



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
	promotionIDs := []int32{int32(123)} // []int32 | ID акций, по которым нужно вернуть информацию

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ApiV1CalendarPromotionsDetailsGet(context.Background()).PromotionIDs(promotionIDs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV1CalendarPromotionsDetailsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1CalendarPromotionsDetailsGet`: ApiV1CalendarPromotionsDetailsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV1CalendarPromotionsDetailsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1CalendarPromotionsDetailsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **promotionIDs** | **[]int32** | ID акций, по которым нужно вернуть информацию | 

### Return type

[**ApiV1CalendarPromotionsDetailsGet200Response**](ApiV1CalendarPromotionsDetailsGet200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1CalendarPromotionsGet

> ApiV1CalendarPromotionsGet200Response ApiV1CalendarPromotionsGet(ctx).StartDateTime(startDateTime).EndDateTime(endDateTime).AllPromo(allPromo).Limit(limit).Offset(offset).Execute()

Список акций



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
	startDateTime := "2023-09-01T00:00:00Z" // string | Начало периода, формат `YYYY-MM-DDTHH:MM:SSZ`
	endDateTime := "2024-08-01T23:59:59Z" // string | Конец периода, формат `YYYY-MM-DDTHH:MM:SSZ`
	allPromo := true // bool | Показать акции:   - `false` — доступные для участия   - `true` — все акции  (default to false)
	limit := int32(10) // int32 | Количество запрашиваемых акций (optional)
	offset := int32(0) // int32 | После какого элемента выдавать данные (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ApiV1CalendarPromotionsGet(context.Background()).StartDateTime(startDateTime).EndDateTime(endDateTime).AllPromo(allPromo).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV1CalendarPromotionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1CalendarPromotionsGet`: ApiV1CalendarPromotionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV1CalendarPromotionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1CalendarPromotionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDateTime** | **string** | Начало периода, формат &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60; | 
 **endDateTime** | **string** | Конец периода, формат &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60; | 
 **allPromo** | **bool** | Показать акции:   - &#x60;false&#x60; — доступные для участия   - &#x60;true&#x60; — все акции  | [default to false]
 **limit** | **int32** | Количество запрашиваемых акций | 
 **offset** | **int32** | После какого элемента выдавать данные | 

### Return type

[**ApiV1CalendarPromotionsGet200Response**](ApiV1CalendarPromotionsGet200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1CalendarPromotionsNomenclaturesGet

> ApiV1CalendarPromotionsNomenclaturesGet200Response ApiV1CalendarPromotionsNomenclaturesGet(ctx).PromotionID(promotionID).InAction(inAction).Limit(limit).Offset(offset).Execute()

Список товаров для участия в акции



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
	promotionID := int32(1) // int32 | ID акции
	inAction := true // bool | Участвует в акции:   - `true` — да   - `false` — нет  (default to false)
	limit := int32(10) // int32 | Количество запрашиваемых товаров (optional)
	offset := int32(0) // int32 | После какого элемента выдавать данные (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ApiV1CalendarPromotionsNomenclaturesGet(context.Background()).PromotionID(promotionID).InAction(inAction).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV1CalendarPromotionsNomenclaturesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1CalendarPromotionsNomenclaturesGet`: ApiV1CalendarPromotionsNomenclaturesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV1CalendarPromotionsNomenclaturesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1CalendarPromotionsNomenclaturesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **promotionID** | **int32** | ID акции | 
 **inAction** | **bool** | Участвует в акции:   - &#x60;true&#x60; — да   - &#x60;false&#x60; — нет  | [default to false]
 **limit** | **int32** | Количество запрашиваемых товаров | 
 **offset** | **int32** | После какого элемента выдавать данные | 

### Return type

[**ApiV1CalendarPromotionsNomenclaturesGet200Response**](ApiV1CalendarPromotionsNomenclaturesGet200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1CalendarPromotionsUploadPost

> ApiV1CalendarPromotionsUploadPost200Response ApiV1CalendarPromotionsUploadPost(ctx).ApiV1CalendarPromotionsUploadPostRequest(apiV1CalendarPromotionsUploadPostRequest).Execute()

Добавить товар в акцию



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
	apiV1CalendarPromotionsUploadPostRequest := *openapiclient.NewApiV1CalendarPromotionsUploadPostRequest() // ApiV1CalendarPromotionsUploadPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ApiV1CalendarPromotionsUploadPost(context.Background()).ApiV1CalendarPromotionsUploadPostRequest(apiV1CalendarPromotionsUploadPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV1CalendarPromotionsUploadPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1CalendarPromotionsUploadPost`: ApiV1CalendarPromotionsUploadPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV1CalendarPromotionsUploadPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1CalendarPromotionsUploadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiV1CalendarPromotionsUploadPostRequest** | [**ApiV1CalendarPromotionsUploadPostRequest**](ApiV1CalendarPromotionsUploadPostRequest.md) |  | 

### Return type

[**ApiV1CalendarPromotionsUploadPost200Response**](ApiV1CalendarPromotionsUploadPost200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2BufferGoodsTaskGet

> ApiV2BufferGoodsTaskGet200Response ApiV2BufferGoodsTaskGet(ctx).Limit(limit).UploadID(uploadID).Offset(offset).Execute()

Детализация необработанной загрузки



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
	limit := int32(10) // int32 | Сколько элементов вывести на одной странице (пагинация). Максимум 1 000 элементов
	uploadID := int32(146567) // int32 | ID загрузки
	offset := int32(0) // int32 | После какого элемента выдавать данные (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ApiV2BufferGoodsTaskGet(context.Background()).Limit(limit).UploadID(uploadID).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV2BufferGoodsTaskGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2BufferGoodsTaskGet`: ApiV2BufferGoodsTaskGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV2BufferGoodsTaskGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2BufferGoodsTaskGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Сколько элементов вывести на одной странице (пагинация). Максимум 1 000 элементов | 
 **uploadID** | **int32** | ID загрузки | 
 **offset** | **int32** | После какого элемента выдавать данные | 

### Return type

[**ApiV2BufferGoodsTaskGet200Response**](ApiV2BufferGoodsTaskGet200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2BufferTasksGet

> ApiV2BufferTasksGet200Response ApiV2BufferTasksGet(ctx).UploadID(uploadID).Execute()

Состояние необработанной загрузки



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
	uploadID := int32(146567) // int32 | ID загрузки

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ApiV2BufferTasksGet(context.Background()).UploadID(uploadID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV2BufferTasksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2BufferTasksGet`: ApiV2BufferTasksGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV2BufferTasksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2BufferTasksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadID** | **int32** | ID загрузки | 

### Return type

[**ApiV2BufferTasksGet200Response**](ApiV2BufferTasksGet200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2HistoryGoodsTaskGet

> ApiV2HistoryGoodsTaskGet200Response ApiV2HistoryGoodsTaskGet(ctx).Limit(limit).UploadID(uploadID).Offset(offset).Execute()

Детализация обработанной загрузки



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
	limit := int32(10) // int32 | Сколько элементов вывести на одной странице (пагинация). Максимум 1 000 элементов
	uploadID := int32(146567) // int32 | ID загрузки
	offset := int32(0) // int32 | После какого элемента выдавать данные (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ApiV2HistoryGoodsTaskGet(context.Background()).Limit(limit).UploadID(uploadID).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV2HistoryGoodsTaskGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2HistoryGoodsTaskGet`: ApiV2HistoryGoodsTaskGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV2HistoryGoodsTaskGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2HistoryGoodsTaskGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Сколько элементов вывести на одной странице (пагинация). Максимум 1 000 элементов | 
 **uploadID** | **int32** | ID загрузки | 
 **offset** | **int32** | После какого элемента выдавать данные | 

### Return type

[**ApiV2HistoryGoodsTaskGet200Response**](ApiV2HistoryGoodsTaskGet200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2HistoryTasksGet

> ApiV2HistoryTasksGet200Response ApiV2HistoryTasksGet(ctx).UploadID(uploadID).Execute()

Состояние обработанной загрузки



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
	uploadID := int32(146567) // int32 | ID загрузки

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ApiV2HistoryTasksGet(context.Background()).UploadID(uploadID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV2HistoryTasksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2HistoryTasksGet`: ApiV2HistoryTasksGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV2HistoryTasksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2HistoryTasksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadID** | **int32** | ID загрузки | 

### Return type

[**ApiV2HistoryTasksGet200Response**](ApiV2HistoryTasksGet200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ListGoodsFilterGet

> ApiV2ListGoodsFilterGet200Response ApiV2ListGoodsFilterGet(ctx).Limit(limit).Offset(offset).FilterNmID(filterNmID).Execute()

Получить товары



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
	limit := int32(10) // int32 | Сколько элементов вывести на одной странице (пагинация). Максимум 1 000 элементов
	offset := int32(0) // int32 | После какого элемента выдавать данные (optional)
	filterNmID := int32(44589768676) // int32 | Артикул Wildberries, по которому искать товар (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ApiV2ListGoodsFilterGet(context.Background()).Limit(limit).Offset(offset).FilterNmID(filterNmID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV2ListGoodsFilterGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ListGoodsFilterGet`: ApiV2ListGoodsFilterGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV2ListGoodsFilterGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ListGoodsFilterGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Сколько элементов вывести на одной странице (пагинация). Максимум 1 000 элементов | 
 **offset** | **int32** | После какого элемента выдавать данные | 
 **filterNmID** | **int32** | Артикул Wildberries, по которому искать товар | 

### Return type

[**ApiV2ListGoodsFilterGet200Response**](ApiV2ListGoodsFilterGet200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ListGoodsSizeNmGet

> ApiV2ListGoodsSizeNmGet200Response ApiV2ListGoodsSizeNmGet(ctx).Limit(limit).NmID(nmID).Offset(offset).Execute()

Получить размеры товара



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
	limit := int32(10) // int32 | Сколько элементов вывести на одной странице (пагинация). Максимум 1 000 элементов
	nmID := int32(1) // int32 | Артикул Wildberries
	offset := int32(0) // int32 | После какого элемента выдавать данные (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ApiV2ListGoodsSizeNmGet(context.Background()).Limit(limit).NmID(nmID).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV2ListGoodsSizeNmGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ListGoodsSizeNmGet`: ApiV2ListGoodsSizeNmGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV2ListGoodsSizeNmGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ListGoodsSizeNmGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Сколько элементов вывести на одной странице (пагинация). Максимум 1 000 элементов | 
 **nmID** | **int32** | Артикул Wildberries | 
 **offset** | **int32** | После какого элемента выдавать данные | 

### Return type

[**ApiV2ListGoodsSizeNmGet200Response**](ApiV2ListGoodsSizeNmGet200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2QuarantineGoodsGet

> ApiV2QuarantineGoodsGet200Response ApiV2QuarantineGoodsGet(ctx).Limit(limit).Offset(offset).Execute()

Получить товары в карантине



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
	limit := int32(10) // int32 | Сколько элементов вывести на одной странице (пагинация). Максимум 1 000 элементов
	offset := int32(0) // int32 | После какого элемента выдавать данные (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ApiV2QuarantineGoodsGet(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV2QuarantineGoodsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2QuarantineGoodsGet`: ApiV2QuarantineGoodsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV2QuarantineGoodsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2QuarantineGoodsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Сколько элементов вывести на одной странице (пагинация). Максимум 1 000 элементов | 
 **offset** | **int32** | После какого элемента выдавать данные | 

### Return type

[**ApiV2QuarantineGoodsGet200Response**](ApiV2QuarantineGoodsGet200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2UploadTaskPost

> TaskCreated ApiV2UploadTaskPost(ctx).ApiV2UploadTaskPostRequest(apiV2UploadTaskPostRequest).Execute()

Установить цены и скидки



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
	apiV2UploadTaskPostRequest := *openapiclient.NewApiV2UploadTaskPostRequest() // ApiV2UploadTaskPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ApiV2UploadTaskPost(context.Background()).ApiV2UploadTaskPostRequest(apiV2UploadTaskPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV2UploadTaskPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2UploadTaskPost`: TaskCreated
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV2UploadTaskPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2UploadTaskPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiV2UploadTaskPostRequest** | [**ApiV2UploadTaskPostRequest**](ApiV2UploadTaskPostRequest.md) |  | 

### Return type

[**TaskCreated**](TaskCreated.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2UploadTaskSizePost

> TaskCreated ApiV2UploadTaskSizePost(ctx).ApiV2UploadTaskSizePostRequest(apiV2UploadTaskSizePostRequest).Execute()

Установить цены для размеров



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
	apiV2UploadTaskSizePostRequest := *openapiclient.NewApiV2UploadTaskSizePostRequest() // ApiV2UploadTaskSizePostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ApiV2UploadTaskSizePost(context.Background()).ApiV2UploadTaskSizePostRequest(apiV2UploadTaskSizePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV2UploadTaskSizePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2UploadTaskSizePost`: TaskCreated
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV2UploadTaskSizePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2UploadTaskSizePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiV2UploadTaskSizePostRequest** | [**ApiV2UploadTaskSizePostRequest**](ApiV2UploadTaskSizePostRequest.md) |  | 

### Return type

[**TaskCreated**](TaskCreated.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

