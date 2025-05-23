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

// checks if the ContentV2GetCardsListPost200ResponseCardsInnerDimensions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentV2GetCardsListPost200ResponseCardsInnerDimensions{}

// ContentV2GetCardsListPost200ResponseCardsInnerDimensions Габариты упаковки товара, см
type ContentV2GetCardsListPost200ResponseCardsInnerDimensions struct {
	// Длина, см
	Length *int32 `json:"length,omitempty"`
	// Ширина, см
	Width *int32 `json:"width,omitempty"`
	// Высота, см
	Height *int32 `json:"height,omitempty"`
	// Потенциальная некорректность габаритов товара: - `true` — не выявлена. `\"isValid\":true` не гарантирует, что размеры указаны корректно. В отдельных случаях (например, при создании новой категории товаров) `\"isValid\":true` будет возвращаться при любых значениях, кроме нулевых. - `false` — указанные габариты значительно отличаются от средних по категории (предмету). Рекомендуется перепроверить, правильно ли указаны размеры товара в упаковке **в сантиметрах**. Функциональность карточки товара, в том числе начисление логистики и хранения, при этом ограничена не будет. Логистика и хранение продолжают начисляться — по текущим габаритам. Также `\"isValid\":false` возвращается при отсутствии значений или нулевом значении любой стороны.
	IsValid *bool `json:"isValid,omitempty"`
}

// NewContentV2GetCardsListPost200ResponseCardsInnerDimensions instantiates a new ContentV2GetCardsListPost200ResponseCardsInnerDimensions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentV2GetCardsListPost200ResponseCardsInnerDimensions() *ContentV2GetCardsListPost200ResponseCardsInnerDimensions {
	this := ContentV2GetCardsListPost200ResponseCardsInnerDimensions{}
	return &this
}

// NewContentV2GetCardsListPost200ResponseCardsInnerDimensionsWithDefaults instantiates a new ContentV2GetCardsListPost200ResponseCardsInnerDimensions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentV2GetCardsListPost200ResponseCardsInnerDimensionsWithDefaults() *ContentV2GetCardsListPost200ResponseCardsInnerDimensions {
	this := ContentV2GetCardsListPost200ResponseCardsInnerDimensions{}
	return &this
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerDimensions) GetLength() int32 {
	if o == nil || IsNil(o.Length) {
		var ret int32
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerDimensions) GetLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.Length) {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerDimensions) HasLength() bool {
	if o != nil && !IsNil(o.Length) {
		return true
	}

	return false
}

// SetLength gets a reference to the given int32 and assigns it to the Length field.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerDimensions) SetLength(v int32) {
	o.Length = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerDimensions) GetWidth() int32 {
	if o == nil || IsNil(o.Width) {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerDimensions) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerDimensions) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerDimensions) SetWidth(v int32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerDimensions) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerDimensions) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerDimensions) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerDimensions) SetHeight(v int32) {
	o.Height = &v
}

// GetIsValid returns the IsValid field value if set, zero value otherwise.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerDimensions) GetIsValid() bool {
	if o == nil || IsNil(o.IsValid) {
		var ret bool
		return ret
	}
	return *o.IsValid
}

// GetIsValidOk returns a tuple with the IsValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerDimensions) GetIsValidOk() (*bool, bool) {
	if o == nil || IsNil(o.IsValid) {
		return nil, false
	}
	return o.IsValid, true
}

// HasIsValid returns a boolean if a field has been set.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerDimensions) HasIsValid() bool {
	if o != nil && !IsNil(o.IsValid) {
		return true
	}

	return false
}

// SetIsValid gets a reference to the given bool and assigns it to the IsValid field.
func (o *ContentV2GetCardsListPost200ResponseCardsInnerDimensions) SetIsValid(v bool) {
	o.IsValid = &v
}

func (o ContentV2GetCardsListPost200ResponseCardsInnerDimensions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentV2GetCardsListPost200ResponseCardsInnerDimensions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Length) {
		toSerialize["length"] = o.Length
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.IsValid) {
		toSerialize["isValid"] = o.IsValid
	}
	return toSerialize, nil
}

type NullableContentV2GetCardsListPost200ResponseCardsInnerDimensions struct {
	value *ContentV2GetCardsListPost200ResponseCardsInnerDimensions
	isSet bool
}

func (v NullableContentV2GetCardsListPost200ResponseCardsInnerDimensions) Get() *ContentV2GetCardsListPost200ResponseCardsInnerDimensions {
	return v.value
}

func (v *NullableContentV2GetCardsListPost200ResponseCardsInnerDimensions) Set(val *ContentV2GetCardsListPost200ResponseCardsInnerDimensions) {
	v.value = val
	v.isSet = true
}

func (v NullableContentV2GetCardsListPost200ResponseCardsInnerDimensions) IsSet() bool {
	return v.isSet
}

func (v *NullableContentV2GetCardsListPost200ResponseCardsInnerDimensions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentV2GetCardsListPost200ResponseCardsInnerDimensions(val *ContentV2GetCardsListPost200ResponseCardsInnerDimensions) *NullableContentV2GetCardsListPost200ResponseCardsInnerDimensions {
	return &NullableContentV2GetCardsListPost200ResponseCardsInnerDimensions{value: val, isSet: true}
}

func (v NullableContentV2GetCardsListPost200ResponseCardsInnerDimensions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentV2GetCardsListPost200ResponseCardsInnerDimensions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
