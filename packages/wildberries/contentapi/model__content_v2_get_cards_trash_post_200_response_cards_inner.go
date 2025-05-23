/*
Описание API Контента

 <dl> <dt>Словарь сокращений:</dt> <dd>КТ — карточка товара</dd> <dd>НМ — номенклатура</dd> </dl> Ограничения по количеству запросов: <dd>Допускается максимум 100 запросов в минуту на методы контента в целом.</dd>  <br> Публичное API Контента создано для синхронизации данных между серверами Wildberries и серверами продавцов. <br> Вы загружаете данные на свои носители, работаете с ними на своих мощностях и синхронизируетесь с нашими серверами по мере необходимости. <br> <code>Не допускается использование API Контента в качестве внешней базы данных. При превышении лимитов на запросы доступ к API будет ограничен.</code> <br>  <br>

API version:
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package contentapi

import (
	"encoding/json"
)

// checks if the ContentV2GetCardsTrashPost200ResponseCardsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentV2GetCardsTrashPost200ResponseCardsInner{}

// ContentV2GetCardsTrashPost200ResponseCardsInner struct for ContentV2GetCardsTrashPost200ResponseCardsInner
type ContentV2GetCardsTrashPost200ResponseCardsInner struct {
	// Артикул WB
	NmID *int32 `json:"nmID,omitempty"`
	// Артикул продавца
	VendorCode *string `json:"vendorCode,omitempty"`
	// Идентификатор предмета
	SubjectID *string `json:"subjectID,omitempty"`
	// Название предмета
	SubjectName *string `json:"subjectName,omitempty"`
	// Массив фото
	Photos []ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner `json:"photos,omitempty"`
	// URL видео
	Video *string `json:"video,omitempty"`
	// Массив размеров
	Sizes      []ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner `json:"sizes,omitempty"`
	Dimensions *ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions  `json:"dimensions,omitempty"`
	// Массив характеристик, при наличии
	Characteristics []ContentV2GetCardsTrashPost200ResponseCardsInnerCharacteristicsInner `json:"characteristics,omitempty"`
	CreatedAt       *string                                                               `json:"createdAt,omitempty"`
	TrashedAt       *string                                                               `json:"trashedAt,omitempty"`
}

// NewContentV2GetCardsTrashPost200ResponseCardsInner instantiates a new ContentV2GetCardsTrashPost200ResponseCardsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentV2GetCardsTrashPost200ResponseCardsInner() *ContentV2GetCardsTrashPost200ResponseCardsInner {
	this := ContentV2GetCardsTrashPost200ResponseCardsInner{}
	return &this
}

// NewContentV2GetCardsTrashPost200ResponseCardsInnerWithDefaults instantiates a new ContentV2GetCardsTrashPost200ResponseCardsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentV2GetCardsTrashPost200ResponseCardsInnerWithDefaults() *ContentV2GetCardsTrashPost200ResponseCardsInner {
	this := ContentV2GetCardsTrashPost200ResponseCardsInner{}
	return &this
}

// GetNmID returns the NmID field value if set, zero value otherwise.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetNmID() int32 {
	if o == nil || IsNil(o.NmID) {
		var ret int32
		return ret
	}
	return *o.NmID
}

// GetNmIDOk returns a tuple with the NmID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetNmIDOk() (*int32, bool) {
	if o == nil || IsNil(o.NmID) {
		return nil, false
	}
	return o.NmID, true
}

// HasNmID returns a boolean if a field has been set.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) HasNmID() bool {
	if o != nil && !IsNil(o.NmID) {
		return true
	}

	return false
}

// SetNmID gets a reference to the given int32 and assigns it to the NmID field.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) SetNmID(v int32) {
	o.NmID = &v
}

// GetVendorCode returns the VendorCode field value if set, zero value otherwise.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetVendorCode() string {
	if o == nil || IsNil(o.VendorCode) {
		var ret string
		return ret
	}
	return *o.VendorCode
}

// GetVendorCodeOk returns a tuple with the VendorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetVendorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.VendorCode) {
		return nil, false
	}
	return o.VendorCode, true
}

// HasVendorCode returns a boolean if a field has been set.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) HasVendorCode() bool {
	if o != nil && !IsNil(o.VendorCode) {
		return true
	}

	return false
}

// SetVendorCode gets a reference to the given string and assigns it to the VendorCode field.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) SetVendorCode(v string) {
	o.VendorCode = &v
}

// GetSubjectID returns the SubjectID field value if set, zero value otherwise.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetSubjectID() string {
	if o == nil || IsNil(o.SubjectID) {
		var ret string
		return ret
	}
	return *o.SubjectID
}

// GetSubjectIDOk returns a tuple with the SubjectID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetSubjectIDOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectID) {
		return nil, false
	}
	return o.SubjectID, true
}

// HasSubjectID returns a boolean if a field has been set.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) HasSubjectID() bool {
	if o != nil && !IsNil(o.SubjectID) {
		return true
	}

	return false
}

// SetSubjectID gets a reference to the given string and assigns it to the SubjectID field.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) SetSubjectID(v string) {
	o.SubjectID = &v
}

// GetSubjectName returns the SubjectName field value if set, zero value otherwise.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetSubjectName() string {
	if o == nil || IsNil(o.SubjectName) {
		var ret string
		return ret
	}
	return *o.SubjectName
}

// GetSubjectNameOk returns a tuple with the SubjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetSubjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectName) {
		return nil, false
	}
	return o.SubjectName, true
}

// HasSubjectName returns a boolean if a field has been set.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) HasSubjectName() bool {
	if o != nil && !IsNil(o.SubjectName) {
		return true
	}

	return false
}

// SetSubjectName gets a reference to the given string and assigns it to the SubjectName field.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) SetSubjectName(v string) {
	o.SubjectName = &v
}

// GetPhotos returns the Photos field value if set, zero value otherwise.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetPhotos() []ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner {
	if o == nil || IsNil(o.Photos) {
		var ret []ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner
		return ret
	}
	return o.Photos
}

// GetPhotosOk returns a tuple with the Photos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetPhotosOk() ([]ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner, bool) {
	if o == nil || IsNil(o.Photos) {
		return nil, false
	}
	return o.Photos, true
}

// HasPhotos returns a boolean if a field has been set.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) HasPhotos() bool {
	if o != nil && !IsNil(o.Photos) {
		return true
	}

	return false
}

// SetPhotos gets a reference to the given []ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner and assigns it to the Photos field.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) SetPhotos(v []ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner) {
	o.Photos = v
}

// GetVideo returns the Video field value if set, zero value otherwise.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetVideo() string {
	if o == nil || IsNil(o.Video) {
		var ret string
		return ret
	}
	return *o.Video
}

// GetVideoOk returns a tuple with the Video field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetVideoOk() (*string, bool) {
	if o == nil || IsNil(o.Video) {
		return nil, false
	}
	return o.Video, true
}

// HasVideo returns a boolean if a field has been set.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) HasVideo() bool {
	if o != nil && !IsNil(o.Video) {
		return true
	}

	return false
}

// SetVideo gets a reference to the given string and assigns it to the Video field.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) SetVideo(v string) {
	o.Video = &v
}

// GetSizes returns the Sizes field value if set, zero value otherwise.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetSizes() []ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner {
	if o == nil || IsNil(o.Sizes) {
		var ret []ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner
		return ret
	}
	return o.Sizes
}

// GetSizesOk returns a tuple with the Sizes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetSizesOk() ([]ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner, bool) {
	if o == nil || IsNil(o.Sizes) {
		return nil, false
	}
	return o.Sizes, true
}

// HasSizes returns a boolean if a field has been set.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) HasSizes() bool {
	if o != nil && !IsNil(o.Sizes) {
		return true
	}

	return false
}

// SetSizes gets a reference to the given []ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner and assigns it to the Sizes field.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) SetSizes(v []ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner) {
	o.Sizes = v
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetDimensions() ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions {
	if o == nil || IsNil(o.Dimensions) {
		var ret ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions
		return ret
	}
	return *o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetDimensionsOk() (*ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions, bool) {
	if o == nil || IsNil(o.Dimensions) {
		return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) HasDimensions() bool {
	if o != nil && !IsNil(o.Dimensions) {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions and assigns it to the Dimensions field.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) SetDimensions(v ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions) {
	o.Dimensions = &v
}

// GetCharacteristics returns the Characteristics field value if set, zero value otherwise.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetCharacteristics() []ContentV2GetCardsTrashPost200ResponseCardsInnerCharacteristicsInner {
	if o == nil || IsNil(o.Characteristics) {
		var ret []ContentV2GetCardsTrashPost200ResponseCardsInnerCharacteristicsInner
		return ret
	}
	return o.Characteristics
}

// GetCharacteristicsOk returns a tuple with the Characteristics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetCharacteristicsOk() ([]ContentV2GetCardsTrashPost200ResponseCardsInnerCharacteristicsInner, bool) {
	if o == nil || IsNil(o.Characteristics) {
		return nil, false
	}
	return o.Characteristics, true
}

// HasCharacteristics returns a boolean if a field has been set.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) HasCharacteristics() bool {
	if o != nil && !IsNil(o.Characteristics) {
		return true
	}

	return false
}

// SetCharacteristics gets a reference to the given []ContentV2GetCardsTrashPost200ResponseCardsInnerCharacteristicsInner and assigns it to the Characteristics field.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) SetCharacteristics(v []ContentV2GetCardsTrashPost200ResponseCardsInnerCharacteristicsInner) {
	o.Characteristics = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetTrashedAt returns the TrashedAt field value if set, zero value otherwise.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetTrashedAt() string {
	if o == nil || IsNil(o.TrashedAt) {
		var ret string
		return ret
	}
	return *o.TrashedAt
}

// GetTrashedAtOk returns a tuple with the TrashedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) GetTrashedAtOk() (*string, bool) {
	if o == nil || IsNil(o.TrashedAt) {
		return nil, false
	}
	return o.TrashedAt, true
}

// HasTrashedAt returns a boolean if a field has been set.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) HasTrashedAt() bool {
	if o != nil && !IsNil(o.TrashedAt) {
		return true
	}

	return false
}

// SetTrashedAt gets a reference to the given string and assigns it to the TrashedAt field.
func (o *ContentV2GetCardsTrashPost200ResponseCardsInner) SetTrashedAt(v string) {
	o.TrashedAt = &v
}

func (o ContentV2GetCardsTrashPost200ResponseCardsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentV2GetCardsTrashPost200ResponseCardsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NmID) {
		toSerialize["nmID"] = o.NmID
	}
	if !IsNil(o.VendorCode) {
		toSerialize["vendorCode"] = o.VendorCode
	}
	if !IsNil(o.SubjectID) {
		toSerialize["subjectID"] = o.SubjectID
	}
	if !IsNil(o.SubjectName) {
		toSerialize["subjectName"] = o.SubjectName
	}
	if !IsNil(o.Photos) {
		toSerialize["photos"] = o.Photos
	}
	if !IsNil(o.Video) {
		toSerialize["video"] = o.Video
	}
	if !IsNil(o.Sizes) {
		toSerialize["sizes"] = o.Sizes
	}
	if !IsNil(o.Dimensions) {
		toSerialize["dimensions"] = o.Dimensions
	}
	if !IsNil(o.Characteristics) {
		toSerialize["characteristics"] = o.Characteristics
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.TrashedAt) {
		toSerialize["trashedAt"] = o.TrashedAt
	}
	return toSerialize, nil
}

type NullableContentV2GetCardsTrashPost200ResponseCardsInner struct {
	value *ContentV2GetCardsTrashPost200ResponseCardsInner
	isSet bool
}

func (v NullableContentV2GetCardsTrashPost200ResponseCardsInner) Get() *ContentV2GetCardsTrashPost200ResponseCardsInner {
	return v.value
}

func (v *NullableContentV2GetCardsTrashPost200ResponseCardsInner) Set(val *ContentV2GetCardsTrashPost200ResponseCardsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableContentV2GetCardsTrashPost200ResponseCardsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableContentV2GetCardsTrashPost200ResponseCardsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentV2GetCardsTrashPost200ResponseCardsInner(val *ContentV2GetCardsTrashPost200ResponseCardsInner) *NullableContentV2GetCardsTrashPost200ResponseCardsInner {
	return &NullableContentV2GetCardsTrashPost200ResponseCardsInner{value: val, isSet: true}
}

func (v NullableContentV2GetCardsTrashPost200ResponseCardsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentV2GetCardsTrashPost200ResponseCardsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
