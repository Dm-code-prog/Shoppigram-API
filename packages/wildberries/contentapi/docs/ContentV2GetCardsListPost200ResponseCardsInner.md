# ContentV2GetCardsListPost200ResponseCardsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NmID** | Pointer to **int32** | Артикул WB | [optional] 
**ImtID** | Pointer to **int32** | Идентификатор КТ. &lt;br&gt; Артикулы WB из одной КТ будут иметь одинаковый imtID | [optional] 
**NmUUID** | Pointer to **string** | Внутренний технический идентификатор товара | [optional] 
**SubjectID** | Pointer to **int32** | Идентификатор предмета | [optional] 
**VendorCode** | Pointer to **string** | Артикул продавца | [optional] 
**SubjectName** | Pointer to **string** | Название предмета | [optional] 
**Brand** | Pointer to **string** | Бренд | [optional] 
**Title** | Pointer to **string** | Наименование товара | [optional] 
**Description** | Pointer to **string** | Описание товара | [optional] 
**Photos** | Pointer to [**[]ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner**](ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner.md) | Массив фото | [optional] 
**Video** | Pointer to **string** | URL видео | [optional] 
**Dimensions** | Pointer to [**ContentV2GetCardsListPost200ResponseCardsInnerDimensions**](ContentV2GetCardsListPost200ResponseCardsInnerDimensions.md) |  | [optional] 
**Characteristics** | Pointer to [**[]ContentV2GetCardsListPost200ResponseCardsInnerCharacteristicsInner**](ContentV2GetCardsListPost200ResponseCardsInnerCharacteristicsInner.md) | Характеристики | [optional] 
**Sizes** | Pointer to [**[]ContentV2GetCardsListPost200ResponseCardsInnerSizesInner**](ContentV2GetCardsListPost200ResponseCardsInnerSizesInner.md) | Размеры товара | [optional] 
**Tags** | Pointer to [**[]ContentV2GetCardsListPost200ResponseCardsInnerTagsInner**](ContentV2GetCardsListPost200ResponseCardsInnerTagsInner.md) | Теги | [optional] 
**CreatedAt** | Pointer to **string** | Дата создания | [optional] 
**UpdatedAt** | Pointer to **string** | Дата изменения | [optional] 

## Methods

### NewContentV2GetCardsListPost200ResponseCardsInner

`func NewContentV2GetCardsListPost200ResponseCardsInner() *ContentV2GetCardsListPost200ResponseCardsInner`

NewContentV2GetCardsListPost200ResponseCardsInner instantiates a new ContentV2GetCardsListPost200ResponseCardsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentV2GetCardsListPost200ResponseCardsInnerWithDefaults

`func NewContentV2GetCardsListPost200ResponseCardsInnerWithDefaults() *ContentV2GetCardsListPost200ResponseCardsInner`

NewContentV2GetCardsListPost200ResponseCardsInnerWithDefaults instantiates a new ContentV2GetCardsListPost200ResponseCardsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNmID

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetNmID() int32`

GetNmID returns the NmID field if non-nil, zero value otherwise.

### GetNmIDOk

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetNmIDOk() (*int32, bool)`

GetNmIDOk returns a tuple with the NmID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmID

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) SetNmID(v int32)`

SetNmID sets NmID field to given value.

### HasNmID

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) HasNmID() bool`

HasNmID returns a boolean if a field has been set.

### GetImtID

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetImtID() int32`

GetImtID returns the ImtID field if non-nil, zero value otherwise.

### GetImtIDOk

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetImtIDOk() (*int32, bool)`

GetImtIDOk returns a tuple with the ImtID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImtID

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) SetImtID(v int32)`

SetImtID sets ImtID field to given value.

### HasImtID

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) HasImtID() bool`

HasImtID returns a boolean if a field has been set.

### GetNmUUID

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetNmUUID() string`

GetNmUUID returns the NmUUID field if non-nil, zero value otherwise.

### GetNmUUIDOk

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetNmUUIDOk() (*string, bool)`

GetNmUUIDOk returns a tuple with the NmUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmUUID

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) SetNmUUID(v string)`

SetNmUUID sets NmUUID field to given value.

### HasNmUUID

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) HasNmUUID() bool`

HasNmUUID returns a boolean if a field has been set.

### GetSubjectID

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetSubjectID() int32`

GetSubjectID returns the SubjectID field if non-nil, zero value otherwise.

### GetSubjectIDOk

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetSubjectIDOk() (*int32, bool)`

GetSubjectIDOk returns a tuple with the SubjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectID

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) SetSubjectID(v int32)`

SetSubjectID sets SubjectID field to given value.

### HasSubjectID

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) HasSubjectID() bool`

HasSubjectID returns a boolean if a field has been set.

### GetVendorCode

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetVendorCode() string`

GetVendorCode returns the VendorCode field if non-nil, zero value otherwise.

### GetVendorCodeOk

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetVendorCodeOk() (*string, bool)`

GetVendorCodeOk returns a tuple with the VendorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCode

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) SetVendorCode(v string)`

SetVendorCode sets VendorCode field to given value.

### HasVendorCode

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) HasVendorCode() bool`

HasVendorCode returns a boolean if a field has been set.

### GetSubjectName

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetSubjectName() string`

GetSubjectName returns the SubjectName field if non-nil, zero value otherwise.

### GetSubjectNameOk

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetSubjectNameOk() (*string, bool)`

GetSubjectNameOk returns a tuple with the SubjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectName

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) SetSubjectName(v string)`

SetSubjectName sets SubjectName field to given value.

### HasSubjectName

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) HasSubjectName() bool`

HasSubjectName returns a boolean if a field has been set.

### GetBrand

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetTitle

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPhotos

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetPhotos() []ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner`

GetPhotos returns the Photos field if non-nil, zero value otherwise.

### GetPhotosOk

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetPhotosOk() (*[]ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner, bool)`

GetPhotosOk returns a tuple with the Photos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotos

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) SetPhotos(v []ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner)`

SetPhotos sets Photos field to given value.

### HasPhotos

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) HasPhotos() bool`

HasPhotos returns a boolean if a field has been set.

### GetVideo

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetVideo() string`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetVideoOk() (*string, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) SetVideo(v string)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetDimensions

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetDimensions() ContentV2GetCardsListPost200ResponseCardsInnerDimensions`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetDimensionsOk() (*ContentV2GetCardsListPost200ResponseCardsInnerDimensions, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) SetDimensions(v ContentV2GetCardsListPost200ResponseCardsInnerDimensions)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetCharacteristics

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetCharacteristics() []ContentV2GetCardsListPost200ResponseCardsInnerCharacteristicsInner`

GetCharacteristics returns the Characteristics field if non-nil, zero value otherwise.

### GetCharacteristicsOk

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetCharacteristicsOk() (*[]ContentV2GetCardsListPost200ResponseCardsInnerCharacteristicsInner, bool)`

GetCharacteristicsOk returns a tuple with the Characteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacteristics

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) SetCharacteristics(v []ContentV2GetCardsListPost200ResponseCardsInnerCharacteristicsInner)`

SetCharacteristics sets Characteristics field to given value.

### HasCharacteristics

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) HasCharacteristics() bool`

HasCharacteristics returns a boolean if a field has been set.

### GetSizes

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetSizes() []ContentV2GetCardsListPost200ResponseCardsInnerSizesInner`

GetSizes returns the Sizes field if non-nil, zero value otherwise.

### GetSizesOk

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetSizesOk() (*[]ContentV2GetCardsListPost200ResponseCardsInnerSizesInner, bool)`

GetSizesOk returns a tuple with the Sizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizes

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) SetSizes(v []ContentV2GetCardsListPost200ResponseCardsInnerSizesInner)`

SetSizes sets Sizes field to given value.

### HasSizes

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) HasSizes() bool`

HasSizes returns a boolean if a field has been set.

### GetTags

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetTags() []ContentV2GetCardsListPost200ResponseCardsInnerTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetTagsOk() (*[]ContentV2GetCardsListPost200ResponseCardsInnerTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) SetTags(v []ContentV2GetCardsListPost200ResponseCardsInnerTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ContentV2GetCardsListPost200ResponseCardsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


