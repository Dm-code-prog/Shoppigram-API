# SupplierTaskMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UploadID** | Pointer to **int32** | ID загрузки | [optional] 
**Status** | Pointer to **int32** | Статус загрузки:    * &#x60;3&#x60; — обработана, в товарах нет ошибок, цены и скидки обновились   * &#x60;4&#x60; — отменена   * &#x60;5&#x60; — обработана, но в товарах есть ошибки. Для товаров без ошибок цены и скидки обновились, а ошибки в остальных товарах можно получить с помощью метода [Детализация обработанной загрузки](#tag/Istoriya-zagruzok/paths/~1api~1v2~1history~1goods~1task/get)   * &#x60;6&#x60; — обработана, но во всех товарах есть ошибки. Их тоже можно получить с помощью метода [Детализация обработанной загрузки](#tag/Istoriya-zagruzok/paths/~1api~1v2~1history~1goods~1task/get)  | [optional] 
**UploadDate** | Pointer to **string** | Дата и время, когда загрузка создана | [optional] 
**ActivationDate** | Pointer to **string** | Дата и время, когда загрузка отправляется в обработку | [optional] 
**OverAllGoodsNumber** | Pointer to **int32** | Всего товаров | [optional] 
**SuccessGoodsNumber** | Pointer to **int32** | Товаров без ошибок | [optional] 

## Methods

### NewSupplierTaskMetadata

`func NewSupplierTaskMetadata() *SupplierTaskMetadata`

NewSupplierTaskMetadata instantiates a new SupplierTaskMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplierTaskMetadataWithDefaults

`func NewSupplierTaskMetadataWithDefaults() *SupplierTaskMetadata`

NewSupplierTaskMetadataWithDefaults instantiates a new SupplierTaskMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUploadID

`func (o *SupplierTaskMetadata) GetUploadID() int32`

GetUploadID returns the UploadID field if non-nil, zero value otherwise.

### GetUploadIDOk

`func (o *SupplierTaskMetadata) GetUploadIDOk() (*int32, bool)`

GetUploadIDOk returns a tuple with the UploadID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadID

`func (o *SupplierTaskMetadata) SetUploadID(v int32)`

SetUploadID sets UploadID field to given value.

### HasUploadID

`func (o *SupplierTaskMetadata) HasUploadID() bool`

HasUploadID returns a boolean if a field has been set.

### GetStatus

`func (o *SupplierTaskMetadata) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SupplierTaskMetadata) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SupplierTaskMetadata) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SupplierTaskMetadata) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUploadDate

`func (o *SupplierTaskMetadata) GetUploadDate() string`

GetUploadDate returns the UploadDate field if non-nil, zero value otherwise.

### GetUploadDateOk

`func (o *SupplierTaskMetadata) GetUploadDateOk() (*string, bool)`

GetUploadDateOk returns a tuple with the UploadDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadDate

`func (o *SupplierTaskMetadata) SetUploadDate(v string)`

SetUploadDate sets UploadDate field to given value.

### HasUploadDate

`func (o *SupplierTaskMetadata) HasUploadDate() bool`

HasUploadDate returns a boolean if a field has been set.

### GetActivationDate

`func (o *SupplierTaskMetadata) GetActivationDate() string`

GetActivationDate returns the ActivationDate field if non-nil, zero value otherwise.

### GetActivationDateOk

`func (o *SupplierTaskMetadata) GetActivationDateOk() (*string, bool)`

GetActivationDateOk returns a tuple with the ActivationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationDate

`func (o *SupplierTaskMetadata) SetActivationDate(v string)`

SetActivationDate sets ActivationDate field to given value.

### HasActivationDate

`func (o *SupplierTaskMetadata) HasActivationDate() bool`

HasActivationDate returns a boolean if a field has been set.

### GetOverAllGoodsNumber

`func (o *SupplierTaskMetadata) GetOverAllGoodsNumber() int32`

GetOverAllGoodsNumber returns the OverAllGoodsNumber field if non-nil, zero value otherwise.

### GetOverAllGoodsNumberOk

`func (o *SupplierTaskMetadata) GetOverAllGoodsNumberOk() (*int32, bool)`

GetOverAllGoodsNumberOk returns a tuple with the OverAllGoodsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverAllGoodsNumber

`func (o *SupplierTaskMetadata) SetOverAllGoodsNumber(v int32)`

SetOverAllGoodsNumber sets OverAllGoodsNumber field to given value.

### HasOverAllGoodsNumber

`func (o *SupplierTaskMetadata) HasOverAllGoodsNumber() bool`

HasOverAllGoodsNumber returns a boolean if a field has been set.

### GetSuccessGoodsNumber

`func (o *SupplierTaskMetadata) GetSuccessGoodsNumber() int32`

GetSuccessGoodsNumber returns the SuccessGoodsNumber field if non-nil, zero value otherwise.

### GetSuccessGoodsNumberOk

`func (o *SupplierTaskMetadata) GetSuccessGoodsNumberOk() (*int32, bool)`

GetSuccessGoodsNumberOk returns a tuple with the SuccessGoodsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessGoodsNumber

`func (o *SupplierTaskMetadata) SetSuccessGoodsNumber(v int32)`

SetSuccessGoodsNumber sets SuccessGoodsNumber field to given value.

### HasSuccessGoodsNumber

`func (o *SupplierTaskMetadata) HasSuccessGoodsNumber() bool`

HasSuccessGoodsNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


