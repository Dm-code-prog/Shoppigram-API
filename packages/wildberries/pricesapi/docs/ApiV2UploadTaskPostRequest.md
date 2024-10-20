# ApiV2UploadTaskPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Good**](Good.md) | Товары, цены и скидки для них. Максимум 1 000 товаров. Цена и скидка не могут быть пустыми одновременно. &lt;br&gt;&lt;br&gt; Если новая цена со скидкой будет хотя бы в 3 раза меньше старой, она попадёт [в карантин](https://seller.wildberries.ru/discount-and-prices/quarantine) и товар будет продаваться по старой цене. Ошибка об этом будет в ответах методов [Состояния загрузок](./#tag/Sostoyaniya-zagruzok) &lt;br&gt;&lt;br&gt; Вы можете изменить цену или скидку с помощью API либо вывести товар из карантина [в личном кабинете](https://seller.wildberries.ru/discount-and-prices/quarantine)  | [optional] 

## Methods

### NewApiV2UploadTaskPostRequest

`func NewApiV2UploadTaskPostRequest() *ApiV2UploadTaskPostRequest`

NewApiV2UploadTaskPostRequest instantiates a new ApiV2UploadTaskPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2UploadTaskPostRequestWithDefaults

`func NewApiV2UploadTaskPostRequestWithDefaults() *ApiV2UploadTaskPostRequest`

NewApiV2UploadTaskPostRequestWithDefaults instantiates a new ApiV2UploadTaskPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ApiV2UploadTaskPostRequest) GetData() []Good`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApiV2UploadTaskPostRequest) GetDataOk() (*[]Good, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApiV2UploadTaskPostRequest) SetData(v []Good)`

SetData sets Data field to given value.

### HasData

`func (o *ApiV2UploadTaskPostRequest) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


