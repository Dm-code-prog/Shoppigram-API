# ContentV2CardsUpdatePostRequestInnerSizesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChrtID** | Pointer to **int32** | Числовой идентификатор размера для данного артикула Wildberries Обязателен к заполнению для существующих размеров. Для добавляемых размеров не указывается.  | [optional] 
**TechSize** | Pointer to **string** | Размер товара (XL, S, 45 и др.) | [optional] 
**WbSize** | Pointer to **string** | Российский размер товара | [optional] 
**Skus** | Pointer to **[]string** | Баркод | [optional] 

## Methods

### NewContentV2CardsUpdatePostRequestInnerSizesInner

`func NewContentV2CardsUpdatePostRequestInnerSizesInner() *ContentV2CardsUpdatePostRequestInnerSizesInner`

NewContentV2CardsUpdatePostRequestInnerSizesInner instantiates a new ContentV2CardsUpdatePostRequestInnerSizesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentV2CardsUpdatePostRequestInnerSizesInnerWithDefaults

`func NewContentV2CardsUpdatePostRequestInnerSizesInnerWithDefaults() *ContentV2CardsUpdatePostRequestInnerSizesInner`

NewContentV2CardsUpdatePostRequestInnerSizesInnerWithDefaults instantiates a new ContentV2CardsUpdatePostRequestInnerSizesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChrtID

`func (o *ContentV2CardsUpdatePostRequestInnerSizesInner) GetChrtID() int32`

GetChrtID returns the ChrtID field if non-nil, zero value otherwise.

### GetChrtIDOk

`func (o *ContentV2CardsUpdatePostRequestInnerSizesInner) GetChrtIDOk() (*int32, bool)`

GetChrtIDOk returns a tuple with the ChrtID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChrtID

`func (o *ContentV2CardsUpdatePostRequestInnerSizesInner) SetChrtID(v int32)`

SetChrtID sets ChrtID field to given value.

### HasChrtID

`func (o *ContentV2CardsUpdatePostRequestInnerSizesInner) HasChrtID() bool`

HasChrtID returns a boolean if a field has been set.

### GetTechSize

`func (o *ContentV2CardsUpdatePostRequestInnerSizesInner) GetTechSize() string`

GetTechSize returns the TechSize field if non-nil, zero value otherwise.

### GetTechSizeOk

`func (o *ContentV2CardsUpdatePostRequestInnerSizesInner) GetTechSizeOk() (*string, bool)`

GetTechSizeOk returns a tuple with the TechSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechSize

`func (o *ContentV2CardsUpdatePostRequestInnerSizesInner) SetTechSize(v string)`

SetTechSize sets TechSize field to given value.

### HasTechSize

`func (o *ContentV2CardsUpdatePostRequestInnerSizesInner) HasTechSize() bool`

HasTechSize returns a boolean if a field has been set.

### GetWbSize

`func (o *ContentV2CardsUpdatePostRequestInnerSizesInner) GetWbSize() string`

GetWbSize returns the WbSize field if non-nil, zero value otherwise.

### GetWbSizeOk

`func (o *ContentV2CardsUpdatePostRequestInnerSizesInner) GetWbSizeOk() (*string, bool)`

GetWbSizeOk returns a tuple with the WbSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWbSize

`func (o *ContentV2CardsUpdatePostRequestInnerSizesInner) SetWbSize(v string)`

SetWbSize sets WbSize field to given value.

### HasWbSize

`func (o *ContentV2CardsUpdatePostRequestInnerSizesInner) HasWbSize() bool`

HasWbSize returns a boolean if a field has been set.

### GetSkus

`func (o *ContentV2CardsUpdatePostRequestInnerSizesInner) GetSkus() []string`

GetSkus returns the Skus field if non-nil, zero value otherwise.

### GetSkusOk

`func (o *ContentV2CardsUpdatePostRequestInnerSizesInner) GetSkusOk() (*[]string, bool)`

GetSkusOk returns a tuple with the Skus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkus

`func (o *ContentV2CardsUpdatePostRequestInnerSizesInner) SetSkus(v []string)`

SetSkus sets Skus field to given value.

### HasSkus

`func (o *ContentV2CardsUpdatePostRequestInnerSizesInner) HasSkus() bool`

HasSkus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


