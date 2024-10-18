# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentV2BarcodesPost**](DefaultApi.md#ContentV2BarcodesPost) | **Post** /content/v2/barcodes | Генерация баркодов
[**ContentV2CardsDeleteTrashPost**](DefaultApi.md#ContentV2CardsDeleteTrashPost) | **Post** /content/v2/cards/delete/trash | Перенос НМ в корзину
[**ContentV2CardsErrorListGet**](DefaultApi.md#ContentV2CardsErrorListGet) | **Get** /content/v2/cards/error/list | Список несозданных номенклатур (НМ) с ошибками
[**ContentV2CardsLimitsGet**](DefaultApi.md#ContentV2CardsLimitsGet) | **Get** /content/v2/cards/limits | Лимиты по КТ
[**ContentV2CardsMoveNmPost**](DefaultApi.md#ContentV2CardsMoveNmPost) | **Post** /content/v2/cards/moveNm | Объединение / Разъединение НМ
[**ContentV2CardsRecoverPost**](DefaultApi.md#ContentV2CardsRecoverPost) | **Post** /content/v2/cards/recover | Восстановление НМ из корзины
[**ContentV2CardsUpdatePost**](DefaultApi.md#ContentV2CardsUpdatePost) | **Post** /content/v2/cards/update | Редактирование КТ
[**ContentV2CardsUploadAddPost**](DefaultApi.md#ContentV2CardsUploadAddPost) | **Post** /content/v2/cards/upload/add | Добавление НМ к КТ
[**ContentV2CardsUploadPost**](DefaultApi.md#ContentV2CardsUploadPost) | **Post** /content/v2/cards/upload | Создание карточки товара
[**ContentV2DirectoryColorsGet**](DefaultApi.md#ContentV2DirectoryColorsGet) | **Get** /content/v2/directory/colors | Цвет
[**ContentV2DirectoryCountriesGet**](DefaultApi.md#ContentV2DirectoryCountriesGet) | **Get** /content/v2/directory/countries | Страна Производства
[**ContentV2DirectoryKindsGet**](DefaultApi.md#ContentV2DirectoryKindsGet) | **Get** /content/v2/directory/kinds | Пол
[**ContentV2DirectorySeasonsGet**](DefaultApi.md#ContentV2DirectorySeasonsGet) | **Get** /content/v2/directory/seasons | Сезон
[**ContentV2DirectoryTnvedGet**](DefaultApi.md#ContentV2DirectoryTnvedGet) | **Get** /content/v2/directory/tnved | ТНВЭД код
[**ContentV2DirectoryVatGet**](DefaultApi.md#ContentV2DirectoryVatGet) | **Get** /content/v2/directory/vat | Ставка НДС
[**ContentV2GetCardsListPost**](DefaultApi.md#ContentV2GetCardsListPost) | **Post** /content/v2/get/cards/list | Список номенклатур (НМ)
[**ContentV2GetCardsTrashPost**](DefaultApi.md#ContentV2GetCardsTrashPost) | **Post** /content/v2/get/cards/trash | Список НМ, находящихся в корзине
[**ContentV2ObjectAllGet**](DefaultApi.md#ContentV2ObjectAllGet) | **Get** /content/v2/object/all | Список предметов (подкатегорий)
[**ContentV2ObjectCharcsSubjectIdGet**](DefaultApi.md#ContentV2ObjectCharcsSubjectIdGet) | **Get** /content/v2/object/charcs/{subjectId} | Характеристики предмета (подкатегории)
[**ContentV2ObjectParentAllGet**](DefaultApi.md#ContentV2ObjectParentAllGet) | **Get** /content/v2/object/parent/all | Родительские категории товаров
[**ContentV2TagIdDelete**](DefaultApi.md#ContentV2TagIdDelete) | **Delete** /content/v2/tag/{id} | Удаление тега
[**ContentV2TagIdPatch**](DefaultApi.md#ContentV2TagIdPatch) | **Patch** /content/v2/tag/{id} | Изменение тега
[**ContentV2TagNomenclatureLinkPost**](DefaultApi.md#ContentV2TagNomenclatureLinkPost) | **Post** /content/v2/tag/nomenclature/link | Управление тегами в КТ
[**ContentV2TagPost**](DefaultApi.md#ContentV2TagPost) | **Post** /content/v2/tag | Создание тега
[**ContentV2TagsGet**](DefaultApi.md#ContentV2TagsGet) | **Get** /content/v2/tags | Список тегов
[**ContentV3MediaFilePost**](DefaultApi.md#ContentV3MediaFilePost) | **Post** /content/v3/media/file | Добавить медиафайлы
[**ContentV3MediaSavePost**](DefaultApi.md#ContentV3MediaSavePost) | **Post** /content/v3/media/save | Изменить медиафайлы



## ContentV2BarcodesPost

> ContentV2BarcodesPost200Response ContentV2BarcodesPost(ctx).ContentV2BarcodesPostRequest(contentV2BarcodesPostRequest).Execute()

Генерация баркодов



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
	contentV2BarcodesPostRequest := *openapiclient.NewContentV2BarcodesPostRequest() // ContentV2BarcodesPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ContentV2BarcodesPost(context.Background()).ContentV2BarcodesPostRequest(contentV2BarcodesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ContentV2BarcodesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentV2BarcodesPost`: ContentV2BarcodesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ContentV2BarcodesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentV2BarcodesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentV2BarcodesPostRequest** | [**ContentV2BarcodesPostRequest**](ContentV2BarcodesPostRequest.md) |  | 

### Return type

[**ContentV2BarcodesPost200Response**](ContentV2BarcodesPost200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentV2CardsDeleteTrashPost

> ContentV2CardsDeleteTrashPost200Response ContentV2CardsDeleteTrashPost(ctx).ContentV2CardsDeleteTrashPostRequestInner(contentV2CardsDeleteTrashPostRequestInner).Execute()

Перенос НМ в корзину



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
	contentV2CardsDeleteTrashPostRequestInner := []openapiclient.ContentV2CardsDeleteTrashPostRequestInner{*openapiclient.NewContentV2CardsDeleteTrashPostRequestInner()} // []ContentV2CardsDeleteTrashPostRequestInner | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ContentV2CardsDeleteTrashPost(context.Background()).ContentV2CardsDeleteTrashPostRequestInner(contentV2CardsDeleteTrashPostRequestInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ContentV2CardsDeleteTrashPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentV2CardsDeleteTrashPost`: ContentV2CardsDeleteTrashPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ContentV2CardsDeleteTrashPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentV2CardsDeleteTrashPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentV2CardsDeleteTrashPostRequestInner** | [**[]ContentV2CardsDeleteTrashPostRequestInner**](ContentV2CardsDeleteTrashPostRequestInner.md) |  | 

### Return type

[**ContentV2CardsDeleteTrashPost200Response**](ContentV2CardsDeleteTrashPost200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentV2CardsErrorListGet

> ContentV2CardsErrorListGet200Response ContentV2CardsErrorListGet(ctx).Locale(locale).Execute()

Список несозданных номенклатур (НМ) с ошибками



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
	locale := "en" // string |  Параметр выбора языка значений полей ответа (для которых предусмотрена мультиязычность).  Не используется в песочнице.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ContentV2CardsErrorListGet(context.Background()).Locale(locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ContentV2CardsErrorListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentV2CardsErrorListGet`: ContentV2CardsErrorListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ContentV2CardsErrorListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentV2CardsErrorListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locale** | **string** |  Параметр выбора языка значений полей ответа (для которых предусмотрена мультиязычность).  Не используется в песочнице.  | 

### Return type

[**ContentV2CardsErrorListGet200Response**](ContentV2CardsErrorListGet200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentV2CardsLimitsGet

> ContentV2CardsLimitsGet200Response ContentV2CardsLimitsGet(ctx).Execute()

Лимиты по КТ



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ContentV2CardsLimitsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ContentV2CardsLimitsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentV2CardsLimitsGet`: ContentV2CardsLimitsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ContentV2CardsLimitsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiContentV2CardsLimitsGetRequest struct via the builder pattern


### Return type

[**ContentV2CardsLimitsGet200Response**](ContentV2CardsLimitsGet200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentV2CardsMoveNmPost

> ResponseCardCreate ContentV2CardsMoveNmPost(ctx).ContentV2CardsMoveNmPostRequest(contentV2CardsMoveNmPostRequest).Execute()

Объединение / Разъединение НМ



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
	contentV2CardsMoveNmPostRequest := openapiclient._content_v2_cards_moveNm_post_request{RequestMoveNmsImtConn: openapiclient.NewRequestMoveNmsImtConn(int32(123), []int32{int32(123)})} // ContentV2CardsMoveNmPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ContentV2CardsMoveNmPost(context.Background()).ContentV2CardsMoveNmPostRequest(contentV2CardsMoveNmPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ContentV2CardsMoveNmPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentV2CardsMoveNmPost`: ResponseCardCreate
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ContentV2CardsMoveNmPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentV2CardsMoveNmPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentV2CardsMoveNmPostRequest** | [**ContentV2CardsMoveNmPostRequest**](ContentV2CardsMoveNmPostRequest.md) |  | 

### Return type

[**ResponseCardCreate**](ResponseCardCreate.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentV2CardsRecoverPost

> ContentV2CardsDeleteTrashPost200Response ContentV2CardsRecoverPost(ctx).ContentV2CardsRecoverPostRequestInner(contentV2CardsRecoverPostRequestInner).Execute()

Восстановление НМ из корзины



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
	contentV2CardsRecoverPostRequestInner := []openapiclient.ContentV2CardsRecoverPostRequestInner{*openapiclient.NewContentV2CardsRecoverPostRequestInner()} // []ContentV2CardsRecoverPostRequestInner | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ContentV2CardsRecoverPost(context.Background()).ContentV2CardsRecoverPostRequestInner(contentV2CardsRecoverPostRequestInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ContentV2CardsRecoverPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentV2CardsRecoverPost`: ContentV2CardsDeleteTrashPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ContentV2CardsRecoverPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentV2CardsRecoverPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentV2CardsRecoverPostRequestInner** | [**[]ContentV2CardsRecoverPostRequestInner**](ContentV2CardsRecoverPostRequestInner.md) |  | 

### Return type

[**ContentV2CardsDeleteTrashPost200Response**](ContentV2CardsDeleteTrashPost200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentV2CardsUpdatePost

> ResponseCardCreate ContentV2CardsUpdatePost(ctx).ContentV2CardsUpdatePostRequestInner(contentV2CardsUpdatePostRequestInner).Execute()

Редактирование КТ



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
	contentV2CardsUpdatePostRequestInner := []openapiclient.ContentV2CardsUpdatePostRequestInner{*openapiclient.NewContentV2CardsUpdatePostRequestInner(int32(123), "VendorCode_example", []openapiclient.ContentV2CardsUpdatePostRequestInnerSizesInner{*openapiclient.NewContentV2CardsUpdatePostRequestInnerSizesInner()})} // []ContentV2CardsUpdatePostRequestInner |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ContentV2CardsUpdatePost(context.Background()).ContentV2CardsUpdatePostRequestInner(contentV2CardsUpdatePostRequestInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ContentV2CardsUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentV2CardsUpdatePost`: ResponseCardCreate
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ContentV2CardsUpdatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentV2CardsUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentV2CardsUpdatePostRequestInner** | [**[]ContentV2CardsUpdatePostRequestInner**](ContentV2CardsUpdatePostRequestInner.md) |  | 

### Return type

[**ResponseCardCreate**](ResponseCardCreate.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentV2CardsUploadAddPost

> ResponseCardCreate ContentV2CardsUploadAddPost(ctx).ContentV2CardsUploadAddPostRequest(contentV2CardsUploadAddPostRequest).Execute()

Добавление НМ к КТ



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
	contentV2CardsUploadAddPostRequest := *openapiclient.NewContentV2CardsUploadAddPostRequest() // ContentV2CardsUploadAddPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ContentV2CardsUploadAddPost(context.Background()).ContentV2CardsUploadAddPostRequest(contentV2CardsUploadAddPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ContentV2CardsUploadAddPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentV2CardsUploadAddPost`: ResponseCardCreate
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ContentV2CardsUploadAddPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentV2CardsUploadAddPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentV2CardsUploadAddPostRequest** | [**ContentV2CardsUploadAddPostRequest**](ContentV2CardsUploadAddPostRequest.md) |  | 

### Return type

[**ResponseCardCreate**](ResponseCardCreate.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentV2CardsUploadPost

> ResponseCardCreate ContentV2CardsUploadPost(ctx).ContentV2CardsUploadPostRequestInner(contentV2CardsUploadPostRequestInner).Execute()

Создание карточки товара



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
	contentV2CardsUploadPostRequestInner := []openapiclient.ContentV2CardsUploadPostRequestInner{*openapiclient.NewContentV2CardsUploadPostRequestInner(int32(123), []openapiclient.ContentV2CardsUploadPostRequestInnerVariantsInner{*openapiclient.NewContentV2CardsUploadPostRequestInnerVariantsInner("VendorCode_example")})} // []ContentV2CardsUploadPostRequestInner |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ContentV2CardsUploadPost(context.Background()).ContentV2CardsUploadPostRequestInner(contentV2CardsUploadPostRequestInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ContentV2CardsUploadPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentV2CardsUploadPost`: ResponseCardCreate
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ContentV2CardsUploadPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentV2CardsUploadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentV2CardsUploadPostRequestInner** | [**[]ContentV2CardsUploadPostRequestInner**](ContentV2CardsUploadPostRequestInner.md) |  | 

### Return type

[**ResponseCardCreate**](ResponseCardCreate.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentV2DirectoryColorsGet

> ContentV2DirectoryColorsGet200Response ContentV2DirectoryColorsGet(ctx).Locale(locale).Execute()

Цвет



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
	locale := "en" // string | Параметр выбора языка (ru, en, zh) значений полей `subjectName`, `name`. Не используется в песочнице (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ContentV2DirectoryColorsGet(context.Background()).Locale(locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ContentV2DirectoryColorsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentV2DirectoryColorsGet`: ContentV2DirectoryColorsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ContentV2DirectoryColorsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentV2DirectoryColorsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locale** | **string** | Параметр выбора языка (ru, en, zh) значений полей &#x60;subjectName&#x60;, &#x60;name&#x60;. Не используется в песочнице | 

### Return type

[**ContentV2DirectoryColorsGet200Response**](ContentV2DirectoryColorsGet200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentV2DirectoryCountriesGet

> ContentV2DirectoryCountriesGet200Response ContentV2DirectoryCountriesGet(ctx).Locale(locale).Execute()

Страна Производства



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
	locale := "en" // string | Параметр выбора языка (ru, en, zh) значений полей `subjectName`, `name`. Не используется в песочнице (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ContentV2DirectoryCountriesGet(context.Background()).Locale(locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ContentV2DirectoryCountriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentV2DirectoryCountriesGet`: ContentV2DirectoryCountriesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ContentV2DirectoryCountriesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentV2DirectoryCountriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locale** | **string** | Параметр выбора языка (ru, en, zh) значений полей &#x60;subjectName&#x60;, &#x60;name&#x60;. Не используется в песочнице | 

### Return type

[**ContentV2DirectoryCountriesGet200Response**](ContentV2DirectoryCountriesGet200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentV2DirectoryKindsGet

> ContentV2DirectoryKindsGet200Response ContentV2DirectoryKindsGet(ctx).Locale(locale).Execute()

Пол



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
	locale := "en" // string | Параметр выбора языка (ru, en, zh) значений полей `subjectName`, `name`. Не используется в песочнице (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ContentV2DirectoryKindsGet(context.Background()).Locale(locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ContentV2DirectoryKindsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentV2DirectoryKindsGet`: ContentV2DirectoryKindsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ContentV2DirectoryKindsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentV2DirectoryKindsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locale** | **string** | Параметр выбора языка (ru, en, zh) значений полей &#x60;subjectName&#x60;, &#x60;name&#x60;. Не используется в песочнице | 

### Return type

[**ContentV2DirectoryKindsGet200Response**](ContentV2DirectoryKindsGet200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentV2DirectorySeasonsGet

> ContentV2DirectorySeasonsGet200Response ContentV2DirectorySeasonsGet(ctx).Locale(locale).Execute()

Сезон



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
	locale := "en" // string | Параметр выбора языка (ru, en, zh) значений полей `subjectName`, `name`. Не используется в песочнице (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ContentV2DirectorySeasonsGet(context.Background()).Locale(locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ContentV2DirectorySeasonsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentV2DirectorySeasonsGet`: ContentV2DirectorySeasonsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ContentV2DirectorySeasonsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentV2DirectorySeasonsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locale** | **string** | Параметр выбора языка (ru, en, zh) значений полей &#x60;subjectName&#x60;, &#x60;name&#x60;. Не используется в песочнице | 

### Return type

[**ContentV2DirectorySeasonsGet200Response**](ContentV2DirectorySeasonsGet200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentV2DirectoryTnvedGet

> ContentV2DirectoryTnvedGet200Response ContentV2DirectoryTnvedGet(ctx).SubjectID(subjectID).Search(search).Locale(locale).Execute()

ТНВЭД код



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
	subjectID := int32(105) // int32 | Идентификатор предмета
	search := int32(6106903000) // int32 | Поиск по ТНВЭД-коду. Работает только в паре с subjectID (optional)
	locale := "en" // string | Язык (`ru`, `en`, `zh`) для значений полей `subjectName`, `name`. Не используется в песочнице (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ContentV2DirectoryTnvedGet(context.Background()).SubjectID(subjectID).Search(search).Locale(locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ContentV2DirectoryTnvedGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentV2DirectoryTnvedGet`: ContentV2DirectoryTnvedGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ContentV2DirectoryTnvedGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentV2DirectoryTnvedGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subjectID** | **int32** | Идентификатор предмета | 
 **search** | **int32** | Поиск по ТНВЭД-коду. Работает только в паре с subjectID | 
 **locale** | **string** | Язык (&#x60;ru&#x60;, &#x60;en&#x60;, &#x60;zh&#x60;) для значений полей &#x60;subjectName&#x60;, &#x60;name&#x60;. Не используется в песочнице | 

### Return type

[**ContentV2DirectoryTnvedGet200Response**](ContentV2DirectoryTnvedGet200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentV2DirectoryVatGet

> ContentV2DirectoryVatGet200Response ContentV2DirectoryVatGet(ctx).Locale(locale).Execute()

Ставка НДС



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
	locale := "ru" // string | Язык значения элементов `data` (`ru`, `en`, `zh`). Не используется в песочнице

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ContentV2DirectoryVatGet(context.Background()).Locale(locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ContentV2DirectoryVatGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentV2DirectoryVatGet`: ContentV2DirectoryVatGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ContentV2DirectoryVatGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentV2DirectoryVatGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locale** | **string** | Язык значения элементов &#x60;data&#x60; (&#x60;ru&#x60;, &#x60;en&#x60;, &#x60;zh&#x60;). Не используется в песочнице | 

### Return type

[**ContentV2DirectoryVatGet200Response**](ContentV2DirectoryVatGet200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentV2GetCardsListPost

> ContentV2GetCardsListPost200Response ContentV2GetCardsListPost(ctx).ContentV2GetCardsListPostRequest(contentV2GetCardsListPostRequest).Locale(locale).Execute()

Список номенклатур (НМ)



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
	contentV2GetCardsListPostRequest := *openapiclient.NewContentV2GetCardsListPostRequest() // ContentV2GetCardsListPostRequest | 
	locale := "ru" // string | Язык для перевода полей ответа `name`, `value` и `object`:  `ru` - русский, `en` - английский, `zh` - китайский.  Не используется в песочнице.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ContentV2GetCardsListPost(context.Background()).ContentV2GetCardsListPostRequest(contentV2GetCardsListPostRequest).Locale(locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ContentV2GetCardsListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentV2GetCardsListPost`: ContentV2GetCardsListPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ContentV2GetCardsListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentV2GetCardsListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentV2GetCardsListPostRequest** | [**ContentV2GetCardsListPostRequest**](ContentV2GetCardsListPostRequest.md) |  | 
 **locale** | **string** | Язык для перевода полей ответа &#x60;name&#x60;, &#x60;value&#x60; и &#x60;object&#x60;:  &#x60;ru&#x60; - русский, &#x60;en&#x60; - английский, &#x60;zh&#x60; - китайский.  Не используется в песочнице.  | 

### Return type

[**ContentV2GetCardsListPost200Response**](ContentV2GetCardsListPost200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentV2GetCardsTrashPost

> ContentV2GetCardsTrashPost200Response ContentV2GetCardsTrashPost(ctx).ContentV2GetCardsTrashPostRequest(contentV2GetCardsTrashPostRequest).Locale(locale).Execute()

Список НМ, находящихся в корзине



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
	contentV2GetCardsTrashPostRequest := *openapiclient.NewContentV2GetCardsTrashPostRequest() // ContentV2GetCardsTrashPostRequest | 
	locale := "locale_example" // string | Язык полей ответа `name`, `value` и `object`: `ru`, `en`, `zh`.  Не используется в песочнице  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ContentV2GetCardsTrashPost(context.Background()).ContentV2GetCardsTrashPostRequest(contentV2GetCardsTrashPostRequest).Locale(locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ContentV2GetCardsTrashPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentV2GetCardsTrashPost`: ContentV2GetCardsTrashPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ContentV2GetCardsTrashPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentV2GetCardsTrashPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentV2GetCardsTrashPostRequest** | [**ContentV2GetCardsTrashPostRequest**](ContentV2GetCardsTrashPostRequest.md) |  | 
 **locale** | **string** | Язык полей ответа &#x60;name&#x60;, &#x60;value&#x60; и &#x60;object&#x60;: &#x60;ru&#x60;, &#x60;en&#x60;, &#x60;zh&#x60;.  Не используется в песочнице  | 

### Return type

[**ContentV2GetCardsTrashPost200Response**](ContentV2GetCardsTrashPost200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentV2ObjectAllGet

> ContentV2ObjectAllGet200Response ContentV2ObjectAllGet(ctx).Name(name).Limit(limit).Locale(locale).Offset(offset).ParentID(parentID).Execute()

Список предметов (подкатегорий)



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
	name := "Носки" // string | Поиск по наименованию предмета (Носки), поиск работает по подстроке, искать можно на любом из поддерживаемых языков. (optional)
	limit := int32(1000) // int32 | Количество подкатегорий (предметов), максимум 1 000 (optional)
	locale := "en" // string | Язык полей ответа (ru, en, zh). Не используется в песочнице (optional)
	offset := int32(5000) // int32 | Номер позиции, с которой необходимо получить ответ (optional)
	parentID := int32(1000) // int32 | Идентификатор родительской категории предмета (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ContentV2ObjectAllGet(context.Background()).Name(name).Limit(limit).Locale(locale).Offset(offset).ParentID(parentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ContentV2ObjectAllGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentV2ObjectAllGet`: ContentV2ObjectAllGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ContentV2ObjectAllGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentV2ObjectAllGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Поиск по наименованию предмета (Носки), поиск работает по подстроке, искать можно на любом из поддерживаемых языков. | 
 **limit** | **int32** | Количество подкатегорий (предметов), максимум 1 000 | 
 **locale** | **string** | Язык полей ответа (ru, en, zh). Не используется в песочнице | 
 **offset** | **int32** | Номер позиции, с которой необходимо получить ответ | 
 **parentID** | **int32** | Идентификатор родительской категории предмета | 

### Return type

[**ContentV2ObjectAllGet200Response**](ContentV2ObjectAllGet200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentV2ObjectCharcsSubjectIdGet

> ContentV2ObjectCharcsSubjectIdGet200Response ContentV2ObjectCharcsSubjectIdGet(ctx, subjectId).Locale(locale).Execute()

Характеристики предмета (подкатегории)



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
	subjectId := int32(105) // int32 | Идентификатор предмета
	locale := "en" // string | Параметр выбора языка (ru, en, zh) значений полей `subjectName`, `name`. Не используется в песочнице (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ContentV2ObjectCharcsSubjectIdGet(context.Background(), subjectId).Locale(locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ContentV2ObjectCharcsSubjectIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentV2ObjectCharcsSubjectIdGet`: ContentV2ObjectCharcsSubjectIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ContentV2ObjectCharcsSubjectIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectId** | **int32** | Идентификатор предмета | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentV2ObjectCharcsSubjectIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locale** | **string** | Параметр выбора языка (ru, en, zh) значений полей &#x60;subjectName&#x60;, &#x60;name&#x60;. Не используется в песочнице | 

### Return type

[**ContentV2ObjectCharcsSubjectIdGet200Response**](ContentV2ObjectCharcsSubjectIdGet200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentV2ObjectParentAllGet

> ContentV2ObjectParentAllGet200Response ContentV2ObjectParentAllGet(ctx).Locale(locale).Execute()

Родительские категории товаров



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
	locale := "en" // string | Параметр выбора языка (ru, en, zh) значений поля name. Не используется в песочнице (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ContentV2ObjectParentAllGet(context.Background()).Locale(locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ContentV2ObjectParentAllGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentV2ObjectParentAllGet`: ContentV2ObjectParentAllGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ContentV2ObjectParentAllGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentV2ObjectParentAllGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locale** | **string** | Параметр выбора языка (ru, en, zh) значений поля name. Не используется в песочнице | 

### Return type

[**ContentV2ObjectParentAllGet200Response**](ContentV2ObjectParentAllGet200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentV2TagIdDelete

> ContentV2TagIdDelete200Response ContentV2TagIdDelete(ctx, id).Execute()

Удаление тега



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
	id := int32(1) // int32 | Числовой идентификатор тега

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ContentV2TagIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ContentV2TagIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentV2TagIdDelete`: ContentV2TagIdDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ContentV2TagIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Числовой идентификатор тега | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentV2TagIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContentV2TagIdDelete200Response**](ContentV2TagIdDelete200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentV2TagIdPatch

> ContentV2TagIdPatch200Response ContentV2TagIdPatch(ctx, id).ContentV2TagIdPatchRequest(contentV2TagIdPatchRequest).Execute()

Изменение тега



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
	id := int32(1) // int32 | Числовой идентификатор тега
	contentV2TagIdPatchRequest := *openapiclient.NewContentV2TagIdPatchRequest() // ContentV2TagIdPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ContentV2TagIdPatch(context.Background(), id).ContentV2TagIdPatchRequest(contentV2TagIdPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ContentV2TagIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentV2TagIdPatch`: ContentV2TagIdPatch200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ContentV2TagIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Числовой идентификатор тега | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentV2TagIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentV2TagIdPatchRequest** | [**ContentV2TagIdPatchRequest**](ContentV2TagIdPatchRequest.md) |  | 

### Return type

[**ContentV2TagIdPatch200Response**](ContentV2TagIdPatch200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentV2TagNomenclatureLinkPost

> ResponseContentError6 ContentV2TagNomenclatureLinkPost(ctx).ContentV2TagNomenclatureLinkPostRequest(contentV2TagNomenclatureLinkPostRequest).Execute()

Управление тегами в КТ



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
	contentV2TagNomenclatureLinkPostRequest := *openapiclient.NewContentV2TagNomenclatureLinkPostRequest() // ContentV2TagNomenclatureLinkPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ContentV2TagNomenclatureLinkPost(context.Background()).ContentV2TagNomenclatureLinkPostRequest(contentV2TagNomenclatureLinkPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ContentV2TagNomenclatureLinkPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentV2TagNomenclatureLinkPost`: ResponseContentError6
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ContentV2TagNomenclatureLinkPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentV2TagNomenclatureLinkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentV2TagNomenclatureLinkPostRequest** | [**ContentV2TagNomenclatureLinkPostRequest**](ContentV2TagNomenclatureLinkPostRequest.md) |  | 

### Return type

[**ResponseContentError6**](ResponseContentError6.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentV2TagPost

> ResponseContentError6 ContentV2TagPost(ctx).ContentV2TagPostRequest(contentV2TagPostRequest).Execute()

Создание тега



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
	contentV2TagPostRequest := *openapiclient.NewContentV2TagPostRequest() // ContentV2TagPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ContentV2TagPost(context.Background()).ContentV2TagPostRequest(contentV2TagPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ContentV2TagPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentV2TagPost`: ResponseContentError6
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ContentV2TagPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentV2TagPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentV2TagPostRequest** | [**ContentV2TagPostRequest**](ContentV2TagPostRequest.md) |  | 

### Return type

[**ResponseContentError6**](ResponseContentError6.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentV2TagsGet

> ContentV2TagsGet200Response ContentV2TagsGet(ctx).Execute()

Список тегов



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ContentV2TagsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ContentV2TagsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentV2TagsGet`: ContentV2TagsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ContentV2TagsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiContentV2TagsGetRequest struct via the builder pattern


### Return type

[**ContentV2TagsGet200Response**](ContentV2TagsGet200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentV3MediaFilePost

> ContentV3MediaSavePost200Response ContentV3MediaFilePost(ctx).XNmId(xNmId).XPhotoNumber(xPhotoNumber).Uploadfile(uploadfile).Execute()

Добавить медиафайлы



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
	xNmId := "213864079" // string | Артикул Wildberries
	xPhotoNumber := int32(2) // int32 | Номер медиафайла на загрузку, начинается с `1`. При загрузке видео всегда указывайте `1`.  Чтобы добавить изображение к уже загруженным, номер медиафайла должен быть больше количества уже загруженных медиафайлов. 
	uploadfile := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ContentV3MediaFilePost(context.Background()).XNmId(xNmId).XPhotoNumber(xPhotoNumber).Uploadfile(uploadfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ContentV3MediaFilePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentV3MediaFilePost`: ContentV3MediaSavePost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ContentV3MediaFilePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentV3MediaFilePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xNmId** | **string** | Артикул Wildberries | 
 **xPhotoNumber** | **int32** | Номер медиафайла на загрузку, начинается с &#x60;1&#x60;. При загрузке видео всегда указывайте &#x60;1&#x60;.  Чтобы добавить изображение к уже загруженным, номер медиафайла должен быть больше количества уже загруженных медиафайлов.  | 
 **uploadfile** | ***os.File** |  | 

### Return type

[**ContentV3MediaSavePost200Response**](ContentV3MediaSavePost200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentV3MediaSavePost

> ContentV3MediaSavePost200Response ContentV3MediaSavePost(ctx).ContentV3MediaSavePostRequest(contentV3MediaSavePostRequest).Execute()

Изменить медиафайлы



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
	contentV3MediaSavePostRequest := *openapiclient.NewContentV3MediaSavePostRequest() // ContentV3MediaSavePostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ContentV3MediaSavePost(context.Background()).ContentV3MediaSavePostRequest(contentV3MediaSavePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ContentV3MediaSavePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentV3MediaSavePost`: ContentV3MediaSavePost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ContentV3MediaSavePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentV3MediaSavePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentV3MediaSavePostRequest** | [**ContentV3MediaSavePostRequest**](ContentV3MediaSavePostRequest.md) |  | 

### Return type

[**ContentV3MediaSavePost200Response**](ContentV3MediaSavePost200Response.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

