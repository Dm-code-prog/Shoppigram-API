# ContentV2CardsUploadPostRequestInnerVariantsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | Pointer to **string** | Бренд | [optional] 
**Title** | Pointer to **string** | Наименование товара | [optional] 
**Description** | Pointer to **string** | Описание товара. Максимальное количество символов зависит от категории товара. Стандарт — 2000, минимум — 1000, максимум — 5000.&lt;br&gt; Подробно о правилах описания в **Правилах заполнения карточки товара** в разделе [Инструкции](https://seller.wildberries.ru/help-center/article/A-113#описание) на портале продавцов.  | [optional] 
**VendorCode** | **string** | Артикул продавца | 
**Dimensions** | Pointer to [**ContentV2CardsUploadPostRequestInnerVariantsInnerDimensions**](ContentV2CardsUploadPostRequestInnerVariantsInnerDimensions.md) |  | [optional] 
**Sizes** | Pointer to [**[]ContentV2CardsUploadPostRequestInnerVariantsInnerSizesInner**](ContentV2CardsUploadPostRequestInnerVariantsInnerSizesInner.md) | Массив с размерами. &lt;br&gt;  Если для размерного товара (обувь, одежда и др.) не указать этот массив, то системой в карточке он будет сгенерирован автоматически с &#x60;techSize&#x60; &#x3D; \&quot;A\&quot; и &#x60;wbSize&#x60; &#x3D; \&quot;1\&quot; и баркодом.  | [optional] 
**Characteristics** | Pointer to [**[]ContentV2CardsUploadPostRequestInnerVariantsInnerCharacteristicsInner**](ContentV2CardsUploadPostRequestInnerVariantsInnerCharacteristicsInner.md) | Массив характеристик товара | [optional] 

## Methods

### NewContentV2CardsUploadPostRequestInnerVariantsInner

`func NewContentV2CardsUploadPostRequestInnerVariantsInner(vendorCode string, ) *ContentV2CardsUploadPostRequestInnerVariantsInner`

NewContentV2CardsUploadPostRequestInnerVariantsInner instantiates a new ContentV2CardsUploadPostRequestInnerVariantsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentV2CardsUploadPostRequestInnerVariantsInnerWithDefaults

`func NewContentV2CardsUploadPostRequestInnerVariantsInnerWithDefaults() *ContentV2CardsUploadPostRequestInnerVariantsInner`

NewContentV2CardsUploadPostRequestInnerVariantsInnerWithDefaults instantiates a new ContentV2CardsUploadPostRequestInnerVariantsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *ContentV2CardsUploadPostRequestInnerVariantsInner) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *ContentV2CardsUploadPostRequestInnerVariantsInner) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *ContentV2CardsUploadPostRequestInnerVariantsInner) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *ContentV2CardsUploadPostRequestInnerVariantsInner) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetTitle

`func (o *ContentV2CardsUploadPostRequestInnerVariantsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ContentV2CardsUploadPostRequestInnerVariantsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ContentV2CardsUploadPostRequestInnerVariantsInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ContentV2CardsUploadPostRequestInnerVariantsInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *ContentV2CardsUploadPostRequestInnerVariantsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContentV2CardsUploadPostRequestInnerVariantsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContentV2CardsUploadPostRequestInnerVariantsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContentV2CardsUploadPostRequestInnerVariantsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVendorCode

`func (o *ContentV2CardsUploadPostRequestInnerVariantsInner) GetVendorCode() string`

GetVendorCode returns the VendorCode field if non-nil, zero value otherwise.

### GetVendorCodeOk

`func (o *ContentV2CardsUploadPostRequestInnerVariantsInner) GetVendorCodeOk() (*string, bool)`

GetVendorCodeOk returns a tuple with the VendorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCode

`func (o *ContentV2CardsUploadPostRequestInnerVariantsInner) SetVendorCode(v string)`

SetVendorCode sets VendorCode field to given value.


### GetDimensions

`func (o *ContentV2CardsUploadPostRequestInnerVariantsInner) GetDimensions() ContentV2CardsUploadPostRequestInnerVariantsInnerDimensions`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *ContentV2CardsUploadPostRequestInnerVariantsInner) GetDimensionsOk() (*ContentV2CardsUploadPostRequestInnerVariantsInnerDimensions, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *ContentV2CardsUploadPostRequestInnerVariantsInner) SetDimensions(v ContentV2CardsUploadPostRequestInnerVariantsInnerDimensions)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *ContentV2CardsUploadPostRequestInnerVariantsInner) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetSizes

`func (o *ContentV2CardsUploadPostRequestInnerVariantsInner) GetSizes() []ContentV2CardsUploadPostRequestInnerVariantsInnerSizesInner`

GetSizes returns the Sizes field if non-nil, zero value otherwise.

### GetSizesOk

`func (o *ContentV2CardsUploadPostRequestInnerVariantsInner) GetSizesOk() (*[]ContentV2CardsUploadPostRequestInnerVariantsInnerSizesInner, bool)`

GetSizesOk returns a tuple with the Sizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizes

`func (o *ContentV2CardsUploadPostRequestInnerVariantsInner) SetSizes(v []ContentV2CardsUploadPostRequestInnerVariantsInnerSizesInner)`

SetSizes sets Sizes field to given value.

### HasSizes

`func (o *ContentV2CardsUploadPostRequestInnerVariantsInner) HasSizes() bool`

HasSizes returns a boolean if a field has been set.

### GetCharacteristics

`func (o *ContentV2CardsUploadPostRequestInnerVariantsInner) GetCharacteristics() []ContentV2CardsUploadPostRequestInnerVariantsInnerCharacteristicsInner`

GetCharacteristics returns the Characteristics field if non-nil, zero value otherwise.

### GetCharacteristicsOk

`func (o *ContentV2CardsUploadPostRequestInnerVariantsInner) GetCharacteristicsOk() (*[]ContentV2CardsUploadPostRequestInnerVariantsInnerCharacteristicsInner, bool)`

GetCharacteristicsOk returns a tuple with the Characteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacteristics

`func (o *ContentV2CardsUploadPostRequestInnerVariantsInner) SetCharacteristics(v []ContentV2CardsUploadPostRequestInnerVariantsInnerCharacteristicsInner)`

SetCharacteristics sets Characteristics field to given value.

### HasCharacteristics

`func (o *ContentV2CardsUploadPostRequestInnerVariantsInner) HasCharacteristics() bool`

HasCharacteristics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


