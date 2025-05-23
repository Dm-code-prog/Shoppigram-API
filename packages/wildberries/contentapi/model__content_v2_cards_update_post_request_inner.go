/*
Описание API Контента

 <dl> <dt>Словарь сокращений:</dt> <dd>КТ — карточка товара</dd> <dd>НМ — номенклатура</dd> </dl> Ограничения по количеству запросов: <dd>Допускается максимум 100 запросов в минуту на методы контента в целом.</dd>  <br> Публичное API Контента создано для синхронизации данных между серверами Wildberries и серверами продавцов. <br> Вы загружаете данные на свои носители, работаете с ними на своих мощностях и синхронизируетесь с нашими серверами по мере необходимости. <br> <code>Не допускается использование API Контента в качестве внешней базы данных. При превышении лимитов на запросы доступ к API будет ограничен.</code> <br>  <br>

API version:
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package contentapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ContentV2CardsUpdatePostRequestInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentV2CardsUpdatePostRequestInner{}

// ContentV2CardsUpdatePostRequestInner struct for ContentV2CardsUpdatePostRequestInner
type ContentV2CardsUpdatePostRequestInner struct {
	// Артикул WB
	NmID int32 `json:"nmID"`
	// Артикул продавца
	VendorCode string `json:"vendorCode"`
	// Бренд
	Brand *string `json:"brand,omitempty"`
	// Наименование товара
	Title *string `json:"title,omitempty"`
	// Описание товара. Максимальное количество символов зависит от категории товара. Стандарт — 2000, минимум — 1000, максимум — 5000.<br> Подробно о правилах описания в **Правилах заполнения карточки товара** в разделе [Инструкции](https://seller.wildberries.ru/help-center/article/A-113#описание) на портале продавцов.
	Description *string                                         `json:"description,omitempty"`
	Dimensions  *ContentV2CardsUpdatePostRequestInnerDimensions `json:"dimensions,omitempty"`
	// Характеристики товара
	Characteristics []ContentV2CardsUpdatePostRequestInnerCharacteristicsInner `json:"characteristics,omitempty"`
	// Массив размеров артикула. <br> Для безразмерного товара все равно нужно передавать данный массив без параметров (wbSize и techSize), но с баркодом.
	Sizes []ContentV2CardsUpdatePostRequestInnerSizesInner `json:"sizes"`
}

type _ContentV2CardsUpdatePostRequestInner ContentV2CardsUpdatePostRequestInner

// NewContentV2CardsUpdatePostRequestInner instantiates a new ContentV2CardsUpdatePostRequestInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentV2CardsUpdatePostRequestInner(nmID int32, vendorCode string, sizes []ContentV2CardsUpdatePostRequestInnerSizesInner) *ContentV2CardsUpdatePostRequestInner {
	this := ContentV2CardsUpdatePostRequestInner{}
	this.NmID = nmID
	this.VendorCode = vendorCode
	this.Sizes = sizes
	return &this
}

// NewContentV2CardsUpdatePostRequestInnerWithDefaults instantiates a new ContentV2CardsUpdatePostRequestInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentV2CardsUpdatePostRequestInnerWithDefaults() *ContentV2CardsUpdatePostRequestInner {
	this := ContentV2CardsUpdatePostRequestInner{}
	return &this
}

// GetNmID returns the NmID field value
func (o *ContentV2CardsUpdatePostRequestInner) GetNmID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NmID
}

// GetNmIDOk returns a tuple with the NmID field value
// and a boolean to check if the value has been set.
func (o *ContentV2CardsUpdatePostRequestInner) GetNmIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NmID, true
}

// SetNmID sets field value
func (o *ContentV2CardsUpdatePostRequestInner) SetNmID(v int32) {
	o.NmID = v
}

// GetVendorCode returns the VendorCode field value
func (o *ContentV2CardsUpdatePostRequestInner) GetVendorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VendorCode
}

// GetVendorCodeOk returns a tuple with the VendorCode field value
// and a boolean to check if the value has been set.
func (o *ContentV2CardsUpdatePostRequestInner) GetVendorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VendorCode, true
}

// SetVendorCode sets field value
func (o *ContentV2CardsUpdatePostRequestInner) SetVendorCode(v string) {
	o.VendorCode = v
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *ContentV2CardsUpdatePostRequestInner) GetBrand() string {
	if o == nil || IsNil(o.Brand) {
		var ret string
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2CardsUpdatePostRequestInner) GetBrandOk() (*string, bool) {
	if o == nil || IsNil(o.Brand) {
		return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *ContentV2CardsUpdatePostRequestInner) HasBrand() bool {
	if o != nil && !IsNil(o.Brand) {
		return true
	}

	return false
}

// SetBrand gets a reference to the given string and assigns it to the Brand field.
func (o *ContentV2CardsUpdatePostRequestInner) SetBrand(v string) {
	o.Brand = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ContentV2CardsUpdatePostRequestInner) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2CardsUpdatePostRequestInner) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ContentV2CardsUpdatePostRequestInner) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ContentV2CardsUpdatePostRequestInner) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ContentV2CardsUpdatePostRequestInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2CardsUpdatePostRequestInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ContentV2CardsUpdatePostRequestInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ContentV2CardsUpdatePostRequestInner) SetDescription(v string) {
	o.Description = &v
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *ContentV2CardsUpdatePostRequestInner) GetDimensions() ContentV2CardsUpdatePostRequestInnerDimensions {
	if o == nil || IsNil(o.Dimensions) {
		var ret ContentV2CardsUpdatePostRequestInnerDimensions
		return ret
	}
	return *o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2CardsUpdatePostRequestInner) GetDimensionsOk() (*ContentV2CardsUpdatePostRequestInnerDimensions, bool) {
	if o == nil || IsNil(o.Dimensions) {
		return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *ContentV2CardsUpdatePostRequestInner) HasDimensions() bool {
	if o != nil && !IsNil(o.Dimensions) {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given ContentV2CardsUpdatePostRequestInnerDimensions and assigns it to the Dimensions field.
func (o *ContentV2CardsUpdatePostRequestInner) SetDimensions(v ContentV2CardsUpdatePostRequestInnerDimensions) {
	o.Dimensions = &v
}

// GetCharacteristics returns the Characteristics field value if set, zero value otherwise.
func (o *ContentV2CardsUpdatePostRequestInner) GetCharacteristics() []ContentV2CardsUpdatePostRequestInnerCharacteristicsInner {
	if o == nil || IsNil(o.Characteristics) {
		var ret []ContentV2CardsUpdatePostRequestInnerCharacteristicsInner
		return ret
	}
	return o.Characteristics
}

// GetCharacteristicsOk returns a tuple with the Characteristics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2CardsUpdatePostRequestInner) GetCharacteristicsOk() ([]ContentV2CardsUpdatePostRequestInnerCharacteristicsInner, bool) {
	if o == nil || IsNil(o.Characteristics) {
		return nil, false
	}
	return o.Characteristics, true
}

// HasCharacteristics returns a boolean if a field has been set.
func (o *ContentV2CardsUpdatePostRequestInner) HasCharacteristics() bool {
	if o != nil && !IsNil(o.Characteristics) {
		return true
	}

	return false
}

// SetCharacteristics gets a reference to the given []ContentV2CardsUpdatePostRequestInnerCharacteristicsInner and assigns it to the Characteristics field.
func (o *ContentV2CardsUpdatePostRequestInner) SetCharacteristics(v []ContentV2CardsUpdatePostRequestInnerCharacteristicsInner) {
	o.Characteristics = v
}

// GetSizes returns the Sizes field value
func (o *ContentV2CardsUpdatePostRequestInner) GetSizes() []ContentV2CardsUpdatePostRequestInnerSizesInner {
	if o == nil {
		var ret []ContentV2CardsUpdatePostRequestInnerSizesInner
		return ret
	}

	return o.Sizes
}

// GetSizesOk returns a tuple with the Sizes field value
// and a boolean to check if the value has been set.
func (o *ContentV2CardsUpdatePostRequestInner) GetSizesOk() ([]ContentV2CardsUpdatePostRequestInnerSizesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sizes, true
}

// SetSizes sets field value
func (o *ContentV2CardsUpdatePostRequestInner) SetSizes(v []ContentV2CardsUpdatePostRequestInnerSizesInner) {
	o.Sizes = v
}

func (o ContentV2CardsUpdatePostRequestInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentV2CardsUpdatePostRequestInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nmID"] = o.NmID
	toSerialize["vendorCode"] = o.VendorCode
	if !IsNil(o.Brand) {
		toSerialize["brand"] = o.Brand
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Dimensions) {
		toSerialize["dimensions"] = o.Dimensions
	}
	if !IsNil(o.Characteristics) {
		toSerialize["characteristics"] = o.Characteristics
	}
	toSerialize["sizes"] = o.Sizes
	return toSerialize, nil
}

func (o *ContentV2CardsUpdatePostRequestInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"nmID",
		"vendorCode",
		"sizes",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varContentV2CardsUpdatePostRequestInner := _ContentV2CardsUpdatePostRequestInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContentV2CardsUpdatePostRequestInner)

	if err != nil {
		return err
	}

	*o = ContentV2CardsUpdatePostRequestInner(varContentV2CardsUpdatePostRequestInner)

	return err
}

type NullableContentV2CardsUpdatePostRequestInner struct {
	value *ContentV2CardsUpdatePostRequestInner
	isSet bool
}

func (v NullableContentV2CardsUpdatePostRequestInner) Get() *ContentV2CardsUpdatePostRequestInner {
	return v.value
}

func (v *NullableContentV2CardsUpdatePostRequestInner) Set(val *ContentV2CardsUpdatePostRequestInner) {
	v.value = val
	v.isSet = true
}

func (v NullableContentV2CardsUpdatePostRequestInner) IsSet() bool {
	return v.isSet
}

func (v *NullableContentV2CardsUpdatePostRequestInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentV2CardsUpdatePostRequestInner(val *ContentV2CardsUpdatePostRequestInner) *NullableContentV2CardsUpdatePostRequestInner {
	return &NullableContentV2CardsUpdatePostRequestInner{value: val, isSet: true}
}

func (v NullableContentV2CardsUpdatePostRequestInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentV2CardsUpdatePostRequestInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
