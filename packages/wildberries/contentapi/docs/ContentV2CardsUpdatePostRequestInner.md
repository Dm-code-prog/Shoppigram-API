# ContentV2CardsUpdatePostRequestInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NmID** | **int32** | Артикул WB | 
**VendorCode** | **string** | Артикул продавца | 
**Brand** | Pointer to **string** | Бренд | [optional] 
**Title** | Pointer to **string** | Наименование товара | [optional] 
**Description** | Pointer to **string** | Описание товара. Максимальное количество символов зависит от категории товара. Стандарт — 2000, минимум — 1000, максимум — 5000.&lt;br&gt; Подробно о правилах описания в **Правилах заполнения карточки товара** в разделе [Инструкции](https://seller.wildberries.ru/help-center/article/A-113#описание) на портале продавцов.  | [optional] 
**Dimensions** | Pointer to [**ContentV2CardsUpdatePostRequestInnerDimensions**](ContentV2CardsUpdatePostRequestInnerDimensions.md) |  | [optional] 
**Characteristics** | Pointer to [**[]ContentV2CardsUpdatePostRequestInnerCharacteristicsInner**](ContentV2CardsUpdatePostRequestInnerCharacteristicsInner.md) | Характеристики товара | [optional] 
**Sizes** | [**[]ContentV2CardsUpdatePostRequestInnerSizesInner**](ContentV2CardsUpdatePostRequestInnerSizesInner.md) | Массив размеров артикула. &lt;br&gt; Для безразмерного товара все равно нужно передавать данный массив без параметров (wbSize и techSize), но с баркодом.  | 

## Methods

### NewContentV2CardsUpdatePostRequestInner

`func NewContentV2CardsUpdatePostRequestInner(nmID int32, vendorCode string, sizes []ContentV2CardsUpdatePostRequestInnerSizesInner, ) *ContentV2CardsUpdatePostRequestInner`

NewContentV2CardsUpdatePostRequestInner instantiates a new ContentV2CardsUpdatePostRequestInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentV2CardsUpdatePostRequestInnerWithDefaults

`func NewContentV2CardsUpdatePostRequestInnerWithDefaults() *ContentV2CardsUpdatePostRequestInner`

NewContentV2CardsUpdatePostRequestInnerWithDefaults instantiates a new ContentV2CardsUpdatePostRequestInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNmID

`func (o *ContentV2CardsUpdatePostRequestInner) GetNmID() int32`

GetNmID returns the NmID field if non-nil, zero value otherwise.

### GetNmIDOk

`func (o *ContentV2CardsUpdatePostRequestInner) GetNmIDOk() (*int32, bool)`

GetNmIDOk returns a tuple with the NmID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmID

`func (o *ContentV2CardsUpdatePostRequestInner) SetNmID(v int32)`

SetNmID sets NmID field to given value.


### GetVendorCode

`func (o *ContentV2CardsUpdatePostRequestInner) GetVendorCode() string`

GetVendorCode returns the VendorCode field if non-nil, zero value otherwise.

### GetVendorCodeOk

`func (o *ContentV2CardsUpdatePostRequestInner) GetVendorCodeOk() (*string, bool)`

GetVendorCodeOk returns a tuple with the VendorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCode

`func (o *ContentV2CardsUpdatePostRequestInner) SetVendorCode(v string)`

SetVendorCode sets VendorCode field to given value.


### GetBrand

`func (o *ContentV2CardsUpdatePostRequestInner) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *ContentV2CardsUpdatePostRequestInner) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *ContentV2CardsUpdatePostRequestInner) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *ContentV2CardsUpdatePostRequestInner) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetTitle

`func (o *ContentV2CardsUpdatePostRequestInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ContentV2CardsUpdatePostRequestInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ContentV2CardsUpdatePostRequestInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ContentV2CardsUpdatePostRequestInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *ContentV2CardsUpdatePostRequestInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContentV2CardsUpdatePostRequestInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContentV2CardsUpdatePostRequestInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContentV2CardsUpdatePostRequestInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDimensions

`func (o *ContentV2CardsUpdatePostRequestInner) GetDimensions() ContentV2CardsUpdatePostRequestInnerDimensions`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *ContentV2CardsUpdatePostRequestInner) GetDimensionsOk() (*ContentV2CardsUpdatePostRequestInnerDimensions, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *ContentV2CardsUpdatePostRequestInner) SetDimensions(v ContentV2CardsUpdatePostRequestInnerDimensions)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *ContentV2CardsUpdatePostRequestInner) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetCharacteristics

`func (o *ContentV2CardsUpdatePostRequestInner) GetCharacteristics() []ContentV2CardsUpdatePostRequestInnerCharacteristicsInner`

GetCharacteristics returns the Characteristics field if non-nil, zero value otherwise.

### GetCharacteristicsOk

`func (o *ContentV2CardsUpdatePostRequestInner) GetCharacteristicsOk() (*[]ContentV2CardsUpdatePostRequestInnerCharacteristicsInner, bool)`

GetCharacteristicsOk returns a tuple with the Characteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacteristics

`func (o *ContentV2CardsUpdatePostRequestInner) SetCharacteristics(v []ContentV2CardsUpdatePostRequestInnerCharacteristicsInner)`

SetCharacteristics sets Characteristics field to given value.

### HasCharacteristics

`func (o *ContentV2CardsUpdatePostRequestInner) HasCharacteristics() bool`

HasCharacteristics returns a boolean if a field has been set.

### GetSizes

`func (o *ContentV2CardsUpdatePostRequestInner) GetSizes() []ContentV2CardsUpdatePostRequestInnerSizesInner`

GetSizes returns the Sizes field if non-nil, zero value otherwise.

### GetSizesOk

`func (o *ContentV2CardsUpdatePostRequestInner) GetSizesOk() (*[]ContentV2CardsUpdatePostRequestInnerSizesInner, bool)`

GetSizesOk returns a tuple with the Sizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizes

`func (o *ContentV2CardsUpdatePostRequestInner) SetSizes(v []ContentV2CardsUpdatePostRequestInnerSizesInner)`

SetSizes sets Sizes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


