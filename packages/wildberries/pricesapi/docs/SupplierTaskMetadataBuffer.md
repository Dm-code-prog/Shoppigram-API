# SupplierTaskMetadataBuffer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UploadID** | Pointer to **int32** | ID загрузки | [optional] 
**Status** | Pointer to **int32** | Статус загрузки: &#x60;1&#x60; — в обработке  | [optional] 
**UploadDate** | Pointer to **string** | Дата и время, когда загрузка создана | [optional] 
**ActivationDate** | Pointer to **string** | Дата и время, когда загрузка отправляется в обработку | [optional] 
**OverAllGoodsNumber** | Pointer to **int32** | Всего товаров | [optional] 
**SuccessGoodsNumber** | Pointer to **int32** | Товаров без ошибок (0, потому что загрузка в обработке) | [optional] 

## Methods

### NewSupplierTaskMetadataBuffer

`func NewSupplierTaskMetadataBuffer() *SupplierTaskMetadataBuffer`

NewSupplierTaskMetadataBuffer instantiates a new SupplierTaskMetadataBuffer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplierTaskMetadataBufferWithDefaults

`func NewSupplierTaskMetadataBufferWithDefaults() *SupplierTaskMetadataBuffer`

NewSupplierTaskMetadataBufferWithDefaults instantiates a new SupplierTaskMetadataBuffer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUploadID

`func (o *SupplierTaskMetadataBuffer) GetUploadID() int32`

GetUploadID returns the UploadID field if non-nil, zero value otherwise.

### GetUploadIDOk

`func (o *SupplierTaskMetadataBuffer) GetUploadIDOk() (*int32, bool)`

GetUploadIDOk returns a tuple with the UploadID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadID

`func (o *SupplierTaskMetadataBuffer) SetUploadID(v int32)`

SetUploadID sets UploadID field to given value.

### HasUploadID

`func (o *SupplierTaskMetadataBuffer) HasUploadID() bool`

HasUploadID returns a boolean if a field has been set.

### GetStatus

`func (o *SupplierTaskMetadataBuffer) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SupplierTaskMetadataBuffer) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SupplierTaskMetadataBuffer) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SupplierTaskMetadataBuffer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUploadDate

`func (o *SupplierTaskMetadataBuffer) GetUploadDate() string`

GetUploadDate returns the UploadDate field if non-nil, zero value otherwise.

### GetUploadDateOk

`func (o *SupplierTaskMetadataBuffer) GetUploadDateOk() (*string, bool)`

GetUploadDateOk returns a tuple with the UploadDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadDate

`func (o *SupplierTaskMetadataBuffer) SetUploadDate(v string)`

SetUploadDate sets UploadDate field to given value.

### HasUploadDate

`func (o *SupplierTaskMetadataBuffer) HasUploadDate() bool`

HasUploadDate returns a boolean if a field has been set.

### GetActivationDate

`func (o *SupplierTaskMetadataBuffer) GetActivationDate() string`

GetActivationDate returns the ActivationDate field if non-nil, zero value otherwise.

### GetActivationDateOk

`func (o *SupplierTaskMetadataBuffer) GetActivationDateOk() (*string, bool)`

GetActivationDateOk returns a tuple with the ActivationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationDate

`func (o *SupplierTaskMetadataBuffer) SetActivationDate(v string)`

SetActivationDate sets ActivationDate field to given value.

### HasActivationDate

`func (o *SupplierTaskMetadataBuffer) HasActivationDate() bool`

HasActivationDate returns a boolean if a field has been set.

### GetOverAllGoodsNumber

`func (o *SupplierTaskMetadataBuffer) GetOverAllGoodsNumber() int32`

GetOverAllGoodsNumber returns the OverAllGoodsNumber field if non-nil, zero value otherwise.

### GetOverAllGoodsNumberOk

`func (o *SupplierTaskMetadataBuffer) GetOverAllGoodsNumberOk() (*int32, bool)`

GetOverAllGoodsNumberOk returns a tuple with the OverAllGoodsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverAllGoodsNumber

`func (o *SupplierTaskMetadataBuffer) SetOverAllGoodsNumber(v int32)`

SetOverAllGoodsNumber sets OverAllGoodsNumber field to given value.

### HasOverAllGoodsNumber

`func (o *SupplierTaskMetadataBuffer) HasOverAllGoodsNumber() bool`

HasOverAllGoodsNumber returns a boolean if a field has been set.

### GetSuccessGoodsNumber

`func (o *SupplierTaskMetadataBuffer) GetSuccessGoodsNumber() int32`

GetSuccessGoodsNumber returns the SuccessGoodsNumber field if non-nil, zero value otherwise.

### GetSuccessGoodsNumberOk

`func (o *SupplierTaskMetadataBuffer) GetSuccessGoodsNumberOk() (*int32, bool)`

GetSuccessGoodsNumberOk returns a tuple with the SuccessGoodsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessGoodsNumber

`func (o *SupplierTaskMetadataBuffer) SetSuccessGoodsNumber(v int32)`

SetSuccessGoodsNumber sets SuccessGoodsNumber field to given value.

### HasSuccessGoodsNumber

`func (o *SupplierTaskMetadataBuffer) HasSuccessGoodsNumber() bool`

HasSuccessGoodsNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


