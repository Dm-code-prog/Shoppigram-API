# ApiV2HistoryTasksGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**SupplierTaskMetadata**](SupplierTaskMetadata.md) |  | [optional] 
**Error** | Pointer to **bool** | Флаг ошибки | [optional] 
**ErrorText** | Pointer to **string** | Текст ошибки  &lt;blockquote class&#x3D;\&quot;spoiler\&quot;&gt;   &lt;p class&#x3D;\&quot;descr\&quot;&gt;Ошибка &lt;code&gt;The product is in quarantine&lt;/code&gt; возникает, если новая цена со скидкой хотя бы в 3 раза меньше старой. Вы можете изменить цену или скидку с помощью API либо вывести товар из карантина &lt;a href&#x3D;\&quot;https://seller.wildberries.ru/discount-and-prices/quarantine\&quot;&gt;в личном кабинете&lt;/a&gt;&lt;/p&gt; &lt;/blockquote&gt;  | [optional] 

## Methods

### NewApiV2HistoryTasksGet200Response

`func NewApiV2HistoryTasksGet200Response() *ApiV2HistoryTasksGet200Response`

NewApiV2HistoryTasksGet200Response instantiates a new ApiV2HistoryTasksGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2HistoryTasksGet200ResponseWithDefaults

`func NewApiV2HistoryTasksGet200ResponseWithDefaults() *ApiV2HistoryTasksGet200Response`

NewApiV2HistoryTasksGet200ResponseWithDefaults instantiates a new ApiV2HistoryTasksGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ApiV2HistoryTasksGet200Response) GetData() SupplierTaskMetadata`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApiV2HistoryTasksGet200Response) GetDataOk() (*SupplierTaskMetadata, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApiV2HistoryTasksGet200Response) SetData(v SupplierTaskMetadata)`

SetData sets Data field to given value.

### HasData

`func (o *ApiV2HistoryTasksGet200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *ApiV2HistoryTasksGet200Response) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiV2HistoryTasksGet200Response) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiV2HistoryTasksGet200Response) SetError(v bool)`

SetError sets Error field to given value.

### HasError

`func (o *ApiV2HistoryTasksGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorText

`func (o *ApiV2HistoryTasksGet200Response) GetErrorText() string`

GetErrorText returns the ErrorText field if non-nil, zero value otherwise.

### GetErrorTextOk

`func (o *ApiV2HistoryTasksGet200Response) GetErrorTextOk() (*string, bool)`

GetErrorTextOk returns a tuple with the ErrorText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorText

`func (o *ApiV2HistoryTasksGet200Response) SetErrorText(v string)`

SetErrorText sets ErrorText field to given value.

### HasErrorText

`func (o *ApiV2HistoryTasksGet200Response) HasErrorText() bool`

HasErrorText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


