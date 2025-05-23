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

// checks if the ContentV2GetCardsListPost200ResponseCardsInnerSizesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentV2GetCardsListPost200ResponseCardsInnerSizesInner{}

// ContentV2GetCardsListPost200ResponseCardsInnerSizesInner struct for ContentV2GetCardsListPost200ResponseCardsInnerSizesInner
type ContentV2GetCardsListPost200ResponseCardsInnerSizesInner struct {
	// Числовой идентификатор размера для данного артикула WB
	ChrtID *int32 `json:"chrtID,omitempty"`
	// Размер товара (А, XXL, 57 и др.)
	TechSize *string `json:"techSize,omitempty"`
	// Российский размер товара
	WbSize *string `json:"wbSize,omitempty"`
	// Баркод товара
	Skus []string `json:"skus,omitempty"`
}

// NewContentV2GetCardsListPost200ResponseCardsInnerSizesInner instantiates a new ContentV2GetCardsListPost200ResponseCardsInnerSizesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentV2GetCardsListPost200ResponseCardsInnerSizesInner() *ContentV2GetCardsListPost200ResponseCardsInnerSizesInner {
	this := ContentV2GetCardsListPost200ResponseCardsInnerSizesInner{}
	return &this
}

// NewContentV2GetCardsListPost200ResponseCardsInnerSizesInnerWithDefaults instantiates a new ContentV2GetCardsListPost200ResponseCardsInnerSizesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentV2GetCardsListPost200ResponseCardsInnerSizesInnerWithDefaults() *ContentV2GetCardsListPost200ResponseCardsInnerSizesInner {
	this := ContentV2GetCardsListPost200ResponseCardsInnerSizesInner{}
	return &this
}

// GetChrtID returns the ChrtID field value if set, zero value otherwise.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerSizesInner) GetChrtID() int32 {
	if o == nil || IsNil(o.ChrtID) {
		var ret int32
		return ret
	}
	return *o.ChrtID
}

// GetChrtIDOk returns a tuple with the ChrtID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerSizesInner) GetChrtIDOk() (*int32, bool) {
	if o == nil || IsNil(o.ChrtID) {
		return nil, false
	}
	return o.ChrtID, true
}

// HasChrtID returns a boolean if a field has been set.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerSizesInner) HasChrtID() bool {
	if o != nil && !IsNil(o.ChrtID) {
		return true
	}

	return false
}

// SetChrtID gets a reference to the given int32 and assigns it to the ChrtID field.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerSizesInner) SetChrtID(v int32) {
	o.ChrtID = &v
}

// GetTechSize returns the TechSize field value if set, zero value otherwise.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerSizesInner) GetTechSize() string {
	if o == nil || IsNil(o.TechSize) {
		var ret string
		return ret
	}
	return *o.TechSize
}

// GetTechSizeOk returns a tuple with the TechSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerSizesInner) GetTechSizeOk() (*string, bool) {
	if o == nil || IsNil(o.TechSize) {
		return nil, false
	}
	return o.TechSize, true
}

// HasTechSize returns a boolean if a field has been set.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerSizesInner) HasTechSize() bool {
	if o != nil && !IsNil(o.TechSize) {
		return true
	}

	return false
}

// SetTechSize gets a reference to the given string and assigns it to the TechSize field.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerSizesInner) SetTechSize(v string) {
	o.TechSize = &v
}

// GetWbSize returns the WbSize field value if set, zero value otherwise.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerSizesInner) GetWbSize() string {
	if o == nil || IsNil(o.WbSize) {
		var ret string
		return ret
	}
	return *o.WbSize
}

// GetWbSizeOk returns a tuple with the WbSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerSizesInner) GetWbSizeOk() (*string, bool) {
	if o == nil || IsNil(o.WbSize) {
		return nil, false
	}
	return o.WbSize, true
}

// HasWbSize returns a boolean if a field has been set.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerSizesInner) HasWbSize() bool {
	if o != nil && !IsNil(o.WbSize) {
		return true
	}

	return false
}

// SetWbSize gets a reference to the given string and assigns it to the WbSize field.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerSizesInner) SetWbSize(v string) {
	o.WbSize = &v
}

// GetSkus returns the Skus field value if set, zero value otherwise.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerSizesInner) GetSkus() []string {
	if o == nil || IsNil(o.Skus) {
		var ret []string
		return ret
	}
	return o.Skus
}

// GetSkusOk returns a tuple with the Skus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerSizesInner) GetSkusOk() ([]string, bool) {
	if o == nil || IsNil(o.Skus) {
		return nil, false
	}
	return o.Skus, true
}

// HasSkus returns a boolean if a field has been set.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerSizesInner) HasSkus() bool {
	if o != nil && !IsNil(o.Skus) {
		return true
	}

	return false
}

// SetSkus gets a reference to the given []string and assigns it to the Skus field.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerSizesInner) SetSkus(v []string) {
	o.Skus = v
}

func (o ContentV2GetCardsListPost200ResponseCardsInnerSizesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentV2GetCardsListPost200ResponseCardsInnerSizesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChrtID) {
		toSerialize["chrtID"] = o.ChrtID
	}
	if !IsNil(o.TechSize) {
		toSerialize["techSize"] = o.TechSize
	}
	if !IsNil(o.WbSize) {
		toSerialize["wbSize"] = o.WbSize
	}
	if !IsNil(o.Skus) {
		toSerialize["skus"] = o.Skus
	}
	return toSerialize, nil
}

type NullableContentV2GetCardsListPost200ResponseCardsInnerSizesInner struct {
	value *ContentV2GetCardsListPost200ResponseCardsInnerSizesInner
	isSet bool
}

func (v NullableContentV2GetCardsListPost200ResponseCardsInnerSizesInner) Get() *ContentV2GetCardsListPost200ResponseCardsInnerSizesInner {
	return v.value
}

func (v *NullableContentV2GetCardsListPost200ResponseCardsInnerSizesInner) Set(val *ContentV2GetCardsListPost200ResponseCardsInnerSizesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableContentV2GetCardsListPost200ResponseCardsInnerSizesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableContentV2GetCardsListPost200ResponseCardsInnerSizesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentV2GetCardsListPost200ResponseCardsInnerSizesInner(val *ContentV2GetCardsListPost200ResponseCardsInnerSizesInner) *NullableContentV2GetCardsListPost200ResponseCardsInnerSizesInner {
	return &NullableContentV2GetCardsListPost200ResponseCardsInnerSizesInner{value: val, isSet: true}
}

func (v NullableContentV2GetCardsListPost200ResponseCardsInnerSizesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentV2GetCardsListPost200ResponseCardsInnerSizesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
