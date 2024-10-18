# ContentV2GetCardsTrashPost200ResponseCardsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NmID** | Pointer to **int32** | Артикул WB | [optional] 
**VendorCode** | Pointer to **string** | Артикул продавца | [optional] 
**SubjectID** | Pointer to **string** | Идентификатор предмета | [optional] 
**SubjectName** | Pointer to **string** | Название предмета | [optional] 
**Photos** | Pointer to [**[]ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner**](ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner.md) | Массив фото | [optional] 
**Video** | Pointer to **string** | URL видео | [optional] 
**Sizes** | Pointer to [**[]ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner**](ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner.md) | Массив размеров | [optional] 
**Dimensions** | Pointer to [**ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions**](ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions.md) |  | [optional] 
**Characteristics** | Pointer to [**[]ContentV2GetCardsTrashPost200ResponseCardsInnerCharacteristicsInner**](ContentV2GetCardsTrashPost200ResponseCardsInnerCharacteristicsInner.md) | Массив характеристик, при наличии | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**TrashedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewContentV2GetCardsTrashPost200ResponseCardsInner

`func NewContentV2GetCardsTrashPost200ResponseCardsInner() *ContentV2GetCardsTrashPost200ResponseCardsInner`

NewContentV2GetCardsTrashPost200ResponseCardsInner instantiates a new ContentV2GetCardsTrashPost200ResponseCardsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentV2GetCardsTrashPost200ResponseCardsInnerWithDefaults

`func NewContentV2GetCardsTrashPost200ResponseCardsInnerWithDefaults() *ContentV2GetCardsTrashPost200ResponseCardsInner`

NewContentV2GetCardsTrashPost200ResponseCardsInnerWithDefaults instantiates a new ContentV2GetCardsTrashPost200ResponseCardsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNmID

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetNmID() int32`

GetNmID returns the NmID field if non-nil, zero value otherwise.

### GetNmIDOk

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetNmIDOk() (*int32, bool)`

GetNmIDOk returns a tuple with the NmID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmID

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) SetNmID(v int32)`

SetNmID sets NmID field to given value.

### HasNmID

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) HasNmID() bool`

HasNmID returns a boolean if a field has been set.

### GetVendorCode

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetVendorCode() string`

GetVendorCode returns the VendorCode field if non-nil, zero value otherwise.

### GetVendorCodeOk

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetVendorCodeOk() (*string, bool)`

GetVendorCodeOk returns a tuple with the VendorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCode

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) SetVendorCode(v string)`

SetVendorCode sets VendorCode field to given value.

### HasVendorCode

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) HasVendorCode() bool`

HasVendorCode returns a boolean if a field has been set.

### GetSubjectID

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetSubjectID() string`

GetSubjectID returns the SubjectID field if non-nil, zero value otherwise.

### GetSubjectIDOk

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetSubjectIDOk() (*string, bool)`

GetSubjectIDOk returns a tuple with the SubjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectID

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) SetSubjectID(v string)`

SetSubjectID sets SubjectID field to given value.

### HasSubjectID

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) HasSubjectID() bool`

HasSubjectID returns a boolean if a field has been set.

### GetSubjectName

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetSubjectName() string`

GetSubjectName returns the SubjectName field if non-nil, zero value otherwise.

### GetSubjectNameOk

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetSubjectNameOk() (*string, bool)`

GetSubjectNameOk returns a tuple with the SubjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectName

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) SetSubjectName(v string)`

SetSubjectName sets SubjectName field to given value.

### HasSubjectName

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) HasSubjectName() bool`

HasSubjectName returns a boolean if a field has been set.

### GetPhotos

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetPhotos() []ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner`

GetPhotos returns the Photos field if non-nil, zero value otherwise.

### GetPhotosOk

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetPhotosOk() (*[]ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner, bool)`

GetPhotosOk returns a tuple with the Photos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotos

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) SetPhotos(v []ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner)`

SetPhotos sets Photos field to given value.

### HasPhotos

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) HasPhotos() bool`

HasPhotos returns a boolean if a field has been set.

### GetVideo

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetVideo() string`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetVideoOk() (*string, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) SetVideo(v string)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetSizes

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetSizes() []ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner`

GetSizes returns the Sizes field if non-nil, zero value otherwise.

### GetSizesOk

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetSizesOk() (*[]ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner, bool)`

GetSizesOk returns a tuple with the Sizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizes

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) SetSizes(v []ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner)`

SetSizes sets Sizes field to given value.

### HasSizes

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) HasSizes() bool`

HasSizes returns a boolean if a field has been set.

### GetDimensions

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetDimensions() ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetDimensionsOk() (*ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) SetDimensions(v ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetCharacteristics

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetCharacteristics() []ContentV2GetCardsTrashPost200ResponseCardsInnerCharacteristicsInner`

GetCharacteristics returns the Characteristics field if non-nil, zero value otherwise.

### GetCharacteristicsOk

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetCharacteristicsOk() (*[]ContentV2GetCardsTrashPost200ResponseCardsInnerCharacteristicsInner, bool)`

GetCharacteristicsOk returns a tuple with the Characteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacteristics

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) SetCharacteristics(v []ContentV2GetCardsTrashPost200ResponseCardsInnerCharacteristicsInner)`

SetCharacteristics sets Characteristics field to given value.

### HasCharacteristics

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) HasCharacteristics() bool`

HasCharacteristics returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetTrashedAt

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetTrashedAt() string`

GetTrashedAt returns the TrashedAt field if non-nil, zero value otherwise.

### GetTrashedAtOk

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetTrashedAtOk() (*string, bool)`

GetTrashedAtOk returns a tuple with the TrashedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrashedAt

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) SetTrashedAt(v string)`

SetTrashedAt sets TrashedAt field to given value.

### HasTrashedAt

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) HasTrashedAt() bool`

HasTrashedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


