# Go API client for contentapi


<dl>
<dt>Словарь сокращений:</dt>
<dd>КТ — карточка товара</dd>
<dd>НМ — номенклатура</dd>
</dl>
Ограничения по количеству запросов:
<dd>Допускается максимум 100 запросов в минуту на методы контента в целом.</dd>

<br>
Публичное API Контента создано для синхронизации данных между серверами Wildberries и серверами продавцов.
<br>
Вы загружаете данные на свои носители, работаете с ними на своих мощностях и синхронизируетесь с нашими серверами по мере необходимости.
<br>
<code>Не допускается использование API Контента в качестве внешней базы данных. При превышении лимитов на запросы доступ к API будет ограничен.</code>
<br> 
<br>


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
import contentapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `contentapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), contentapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `contentapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), contentapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `contentapi.ContextOperationServerIndices` and `contentapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), contentapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), contentapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**ContentV2BarcodesPost**](docs/DefaultApi.md#contentv2barcodespost) | **Post** /content/v2/barcodes | Генерация баркодов
*DefaultApi* | [**ContentV2CardsDeleteTrashPost**](docs/DefaultApi.md#contentv2cardsdeletetrashpost) | **Post** /content/v2/cards/delete/trash | Перенос НМ в корзину
*DefaultApi* | [**ContentV2CardsErrorListGet**](docs/DefaultApi.md#contentv2cardserrorlistget) | **Get** /content/v2/cards/error/list | Список несозданных номенклатур (НМ) с ошибками
*DefaultApi* | [**ContentV2CardsLimitsGet**](docs/DefaultApi.md#contentv2cardslimitsget) | **Get** /content/v2/cards/limits | Лимиты по КТ
*DefaultApi* | [**ContentV2CardsMoveNmPost**](docs/DefaultApi.md#contentv2cardsmovenmpost) | **Post** /content/v2/cards/moveNm | Объединение / Разъединение НМ
*DefaultApi* | [**ContentV2CardsRecoverPost**](docs/DefaultApi.md#contentv2cardsrecoverpost) | **Post** /content/v2/cards/recover | Восстановление НМ из корзины
*DefaultApi* | [**ContentV2CardsUpdatePost**](docs/DefaultApi.md#contentv2cardsupdatepost) | **Post** /content/v2/cards/update | Редактирование КТ
*DefaultApi* | [**ContentV2CardsUploadAddPost**](docs/DefaultApi.md#contentv2cardsuploadaddpost) | **Post** /content/v2/cards/upload/add | Добавление НМ к КТ
*DefaultApi* | [**ContentV2CardsUploadPost**](docs/DefaultApi.md#contentv2cardsuploadpost) | **Post** /content/v2/cards/upload | Создание карточки товара
*DefaultApi* | [**ContentV2DirectoryColorsGet**](docs/DefaultApi.md#contentv2directorycolorsget) | **Get** /content/v2/directory/colors | Цвет
*DefaultApi* | [**ContentV2DirectoryCountriesGet**](docs/DefaultApi.md#contentv2directorycountriesget) | **Get** /content/v2/directory/countries | Страна Производства
*DefaultApi* | [**ContentV2DirectoryKindsGet**](docs/DefaultApi.md#contentv2directorykindsget) | **Get** /content/v2/directory/kinds | Пол
*DefaultApi* | [**ContentV2DirectorySeasonsGet**](docs/DefaultApi.md#contentv2directoryseasonsget) | **Get** /content/v2/directory/seasons | Сезон
*DefaultApi* | [**ContentV2DirectoryTnvedGet**](docs/DefaultApi.md#contentv2directorytnvedget) | **Get** /content/v2/directory/tnved | ТНВЭД код
*DefaultApi* | [**ContentV2DirectoryVatGet**](docs/DefaultApi.md#contentv2directoryvatget) | **Get** /content/v2/directory/vat | Ставка НДС
*DefaultApi* | [**ContentV2GetCardsListPost**](docs/DefaultApi.md#contentv2getcardslistpost) | **Post** /content/v2/get/cards/list | Список номенклатур (НМ)
*DefaultApi* | [**ContentV2GetCardsTrashPost**](docs/DefaultApi.md#contentv2getcardstrashpost) | **Post** /content/v2/get/cards/trash | Список НМ, находящихся в корзине
*DefaultApi* | [**ContentV2ObjectAllGet**](docs/DefaultApi.md#contentv2objectallget) | **Get** /content/v2/object/all | Список предметов (подкатегорий)
*DefaultApi* | [**ContentV2ObjectCharcsSubjectIdGet**](docs/DefaultApi.md#contentv2objectcharcssubjectidget) | **Get** /content/v2/object/charcs/{subjectId} | Характеристики предмета (подкатегории)
*DefaultApi* | [**ContentV2ObjectParentAllGet**](docs/DefaultApi.md#contentv2objectparentallget) | **Get** /content/v2/object/parent/all | Родительские категории товаров
*DefaultApi* | [**ContentV2TagIdDelete**](docs/DefaultApi.md#contentv2tagiddelete) | **Delete** /content/v2/tag/{id} | Удаление тега
*DefaultApi* | [**ContentV2TagIdPatch**](docs/DefaultApi.md#contentv2tagidpatch) | **Patch** /content/v2/tag/{id} | Изменение тега
*DefaultApi* | [**ContentV2TagNomenclatureLinkPost**](docs/DefaultApi.md#contentv2tagnomenclaturelinkpost) | **Post** /content/v2/tag/nomenclature/link | Управление тегами в КТ
*DefaultApi* | [**ContentV2TagPost**](docs/DefaultApi.md#contentv2tagpost) | **Post** /content/v2/tag | Создание тега
*DefaultApi* | [**ContentV2TagsGet**](docs/DefaultApi.md#contentv2tagsget) | **Get** /content/v2/tags | Список тегов
*DefaultApi* | [**ContentV3MediaFilePost**](docs/DefaultApi.md#contentv3mediafilepost) | **Post** /content/v3/media/file | Добавить медиафайлы
*DefaultApi* | [**ContentV3MediaSavePost**](docs/DefaultApi.md#contentv3mediasavepost) | **Post** /content/v3/media/save | Изменить медиафайлы


## Documentation For Models

 - [ContentV2BarcodesPost200Response](docs/ContentV2BarcodesPost200Response.md)
 - [ContentV2BarcodesPostRequest](docs/ContentV2BarcodesPostRequest.md)
 - [ContentV2CardsDeleteTrashPost200Response](docs/ContentV2CardsDeleteTrashPost200Response.md)
 - [ContentV2CardsDeleteTrashPostRequestInner](docs/ContentV2CardsDeleteTrashPostRequestInner.md)
 - [ContentV2CardsErrorListGet200Response](docs/ContentV2CardsErrorListGet200Response.md)
 - [ContentV2CardsErrorListGet200ResponseDataInner](docs/ContentV2CardsErrorListGet200ResponseDataInner.md)
 - [ContentV2CardsLimitsGet200Response](docs/ContentV2CardsLimitsGet200Response.md)
 - [ContentV2CardsLimitsGet200ResponseData](docs/ContentV2CardsLimitsGet200ResponseData.md)
 - [ContentV2CardsMoveNmPost400Response](docs/ContentV2CardsMoveNmPost400Response.md)
 - [ContentV2CardsMoveNmPostRequest](docs/ContentV2CardsMoveNmPostRequest.md)
 - [ContentV2CardsRecoverPostRequestInner](docs/ContentV2CardsRecoverPostRequestInner.md)
 - [ContentV2CardsUpdatePostRequestInner](docs/ContentV2CardsUpdatePostRequestInner.md)
 - [ContentV2CardsUpdatePostRequestInnerCharacteristicsInner](docs/ContentV2CardsUpdatePostRequestInnerCharacteristicsInner.md)
 - [ContentV2CardsUpdatePostRequestInnerDimensions](docs/ContentV2CardsUpdatePostRequestInnerDimensions.md)
 - [ContentV2CardsUpdatePostRequestInnerSizesInner](docs/ContentV2CardsUpdatePostRequestInnerSizesInner.md)
 - [ContentV2CardsUploadAddPostRequest](docs/ContentV2CardsUploadAddPostRequest.md)
 - [ContentV2CardsUploadAddPostRequestCardsToAddInner](docs/ContentV2CardsUploadAddPostRequestCardsToAddInner.md)
 - [ContentV2CardsUploadAddPostRequestCardsToAddInnerCharacteristicsInner](docs/ContentV2CardsUploadAddPostRequestCardsToAddInnerCharacteristicsInner.md)
 - [ContentV2CardsUploadAddPostRequestCardsToAddInnerSizesInner](docs/ContentV2CardsUploadAddPostRequestCardsToAddInnerSizesInner.md)
 - [ContentV2CardsUploadPost401Response](docs/ContentV2CardsUploadPost401Response.md)
 - [ContentV2CardsUploadPost413Response](docs/ContentV2CardsUploadPost413Response.md)
 - [ContentV2CardsUploadPostRequestInner](docs/ContentV2CardsUploadPostRequestInner.md)
 - [ContentV2CardsUploadPostRequestInnerVariantsInner](docs/ContentV2CardsUploadPostRequestInnerVariantsInner.md)
 - [ContentV2CardsUploadPostRequestInnerVariantsInnerCharacteristicsInner](docs/ContentV2CardsUploadPostRequestInnerVariantsInnerCharacteristicsInner.md)
 - [ContentV2CardsUploadPostRequestInnerVariantsInnerDimensions](docs/ContentV2CardsUploadPostRequestInnerVariantsInnerDimensions.md)
 - [ContentV2CardsUploadPostRequestInnerVariantsInnerSizesInner](docs/ContentV2CardsUploadPostRequestInnerVariantsInnerSizesInner.md)
 - [ContentV2DirectoryColorsGet200Response](docs/ContentV2DirectoryColorsGet200Response.md)
 - [ContentV2DirectoryColorsGet200ResponseDataInner](docs/ContentV2DirectoryColorsGet200ResponseDataInner.md)
 - [ContentV2DirectoryCountriesGet200Response](docs/ContentV2DirectoryCountriesGet200Response.md)
 - [ContentV2DirectoryCountriesGet200ResponseDataInner](docs/ContentV2DirectoryCountriesGet200ResponseDataInner.md)
 - [ContentV2DirectoryKindsGet200Response](docs/ContentV2DirectoryKindsGet200Response.md)
 - [ContentV2DirectorySeasonsGet200Response](docs/ContentV2DirectorySeasonsGet200Response.md)
 - [ContentV2DirectoryTnvedGet200Response](docs/ContentV2DirectoryTnvedGet200Response.md)
 - [ContentV2DirectoryTnvedGet200ResponseDataInner](docs/ContentV2DirectoryTnvedGet200ResponseDataInner.md)
 - [ContentV2DirectoryVatGet200Response](docs/ContentV2DirectoryVatGet200Response.md)
 - [ContentV2GetCardsListPost200Response](docs/ContentV2GetCardsListPost200Response.md)
 - [ContentV2GetCardsListPost200ResponseCardsInner](docs/ContentV2GetCardsListPost200ResponseCardsInner.md)
 - [ContentV2GetCardsListPost200ResponseCardsInnerCharacteristicsInner](docs/ContentV2GetCardsListPost200ResponseCardsInnerCharacteristicsInner.md)
 - [ContentV2GetCardsListPost200ResponseCardsInnerDimensions](docs/ContentV2GetCardsListPost200ResponseCardsInnerDimensions.md)
 - [ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner](docs/ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner.md)
 - [ContentV2GetCardsListPost200ResponseCardsInnerSizesInner](docs/ContentV2GetCardsListPost200ResponseCardsInnerSizesInner.md)
 - [ContentV2GetCardsListPost200ResponseCardsInnerTagsInner](docs/ContentV2GetCardsListPost200ResponseCardsInnerTagsInner.md)
 - [ContentV2GetCardsListPost200ResponseCursor](docs/ContentV2GetCardsListPost200ResponseCursor.md)
 - [ContentV2GetCardsListPostRequest](docs/ContentV2GetCardsListPostRequest.md)
 - [ContentV2GetCardsListPostRequestSettings](docs/ContentV2GetCardsListPostRequestSettings.md)
 - [ContentV2GetCardsListPostRequestSettingsCursor](docs/ContentV2GetCardsListPostRequestSettingsCursor.md)
 - [ContentV2GetCardsListPostRequestSettingsFilter](docs/ContentV2GetCardsListPostRequestSettingsFilter.md)
 - [ContentV2GetCardsListPostRequestSettingsSort](docs/ContentV2GetCardsListPostRequestSettingsSort.md)
 - [ContentV2GetCardsTrashPost200Response](docs/ContentV2GetCardsTrashPost200Response.md)
 - [ContentV2GetCardsTrashPost200ResponseCardsInner](docs/ContentV2GetCardsTrashPost200ResponseCardsInner.md)
 - [ContentV2GetCardsTrashPost200ResponseCardsInnerCharacteristicsInner](docs/ContentV2GetCardsTrashPost200ResponseCardsInnerCharacteristicsInner.md)
 - [ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions](docs/ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions.md)
 - [ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner](docs/ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner.md)
 - [ContentV2GetCardsTrashPost200ResponseCursor](docs/ContentV2GetCardsTrashPost200ResponseCursor.md)
 - [ContentV2GetCardsTrashPostRequest](docs/ContentV2GetCardsTrashPostRequest.md)
 - [ContentV2GetCardsTrashPostRequestSettings](docs/ContentV2GetCardsTrashPostRequestSettings.md)
 - [ContentV2GetCardsTrashPostRequestSettingsCursor](docs/ContentV2GetCardsTrashPostRequestSettingsCursor.md)
 - [ContentV2GetCardsTrashPostRequestSettingsFilter](docs/ContentV2GetCardsTrashPostRequestSettingsFilter.md)
 - [ContentV2GetCardsTrashPostRequestSettingsSort](docs/ContentV2GetCardsTrashPostRequestSettingsSort.md)
 - [ContentV2ObjectAllGet200Response](docs/ContentV2ObjectAllGet200Response.md)
 - [ContentV2ObjectAllGet200ResponseDataInner](docs/ContentV2ObjectAllGet200ResponseDataInner.md)
 - [ContentV2ObjectCharcsSubjectIdGet200Response](docs/ContentV2ObjectCharcsSubjectIdGet200Response.md)
 - [ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner](docs/ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner.md)
 - [ContentV2ObjectParentAllGet200Response](docs/ContentV2ObjectParentAllGet200Response.md)
 - [ContentV2ObjectParentAllGet200ResponseDataInner](docs/ContentV2ObjectParentAllGet200ResponseDataInner.md)
 - [ContentV2TagIdDelete200Response](docs/ContentV2TagIdDelete200Response.md)
 - [ContentV2TagIdDelete400Response](docs/ContentV2TagIdDelete400Response.md)
 - [ContentV2TagIdPatch200Response](docs/ContentV2TagIdPatch200Response.md)
 - [ContentV2TagIdPatch400Response](docs/ContentV2TagIdPatch400Response.md)
 - [ContentV2TagIdPatchRequest](docs/ContentV2TagIdPatchRequest.md)
 - [ContentV2TagNomenclatureLinkPostRequest](docs/ContentV2TagNomenclatureLinkPostRequest.md)
 - [ContentV2TagPost400Response](docs/ContentV2TagPost400Response.md)
 - [ContentV2TagPostRequest](docs/ContentV2TagPostRequest.md)
 - [ContentV2TagsGet200Response](docs/ContentV2TagsGet200Response.md)
 - [ContentV2TagsGet200ResponseData](docs/ContentV2TagsGet200ResponseData.md)
 - [ContentV3MediaSavePost200Response](docs/ContentV3MediaSavePost200Response.md)
 - [ContentV3MediaSavePostRequest](docs/ContentV3MediaSavePostRequest.md)
 - [MediaErrors](docs/MediaErrors.md)
 - [RequestMoveNmsImtConn](docs/RequestMoveNmsImtConn.md)
 - [RequestMoveNmsImtDisconn](docs/RequestMoveNmsImtDisconn.md)
 - [ResponseBodyContentError400](docs/ResponseBodyContentError400.md)
 - [ResponseBodyContentError400AdditionalErrors](docs/ResponseBodyContentError400AdditionalErrors.md)
 - [ResponseBodyContentError403](docs/ResponseBodyContentError403.md)
 - [ResponseCardCreate](docs/ResponseCardCreate.md)
 - [ResponseCardCreateAdditionalErrors](docs/ResponseCardCreateAdditionalErrors.md)
 - [ResponseCardCreateAdditionalErrorsOneOf](docs/ResponseCardCreateAdditionalErrorsOneOf.md)
 - [ResponseContentError1](docs/ResponseContentError1.md)
 - [ResponseContentError1AdditionalErrors](docs/ResponseContentError1AdditionalErrors.md)
 - [ResponseContentError4](docs/ResponseContentError4.md)
 - [ResponseContentError4AdditionalErrors](docs/ResponseContentError4AdditionalErrors.md)
 - [ResponseContentError5](docs/ResponseContentError5.md)
 - [ResponseContentError5AdditionalErrors](docs/ResponseContentError5AdditionalErrors.md)
 - [ResponseContentError6](docs/ResponseContentError6.md)
 - [ResponseIncorrectDate](docs/ResponseIncorrectDate.md)


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
		contentapi.ContextAPIKeys,
		map[string]contentapi.APIKey{
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



