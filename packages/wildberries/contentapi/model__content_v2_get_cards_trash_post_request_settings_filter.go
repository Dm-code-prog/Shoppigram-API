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

// checks if the ContentV2GetCardsTrashPostRequestSettingsFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentV2GetCardsTrashPostRequestSettingsFilter{}

// ContentV2GetCardsTrashPostRequestSettingsFilter Параметры фильтрации
type ContentV2GetCardsTrashPostRequestSettingsFilter struct {
	// Поиск по артикулу продавца, артикулу WB, баркоду.
	TextSearch *string `json:"textSearch,omitempty"`
}

// NewContentV2GetCardsTrashPostRequestSettingsFilter instantiates a new ContentV2GetCardsTrashPostRequestSettingsFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentV2GetCardsTrashPostRequestSettingsFilter() *ContentV2GetCardsTrashPostRequestSettingsFilter {
	this := ContentV2GetCardsTrashPostRequestSettingsFilter{}
	return &this
}

// NewContentV2GetCardsTrashPostRequestSettingsFilterWithDefaults instantiates a new ContentV2GetCardsTrashPostRequestSettingsFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentV2GetCardsTrashPostRequestSettingsFilterWithDefaults() *ContentV2GetCardsTrashPostRequestSettingsFilter {
	this := ContentV2GetCardsTrashPostRequestSettingsFilter{}
	return &this
}

// GetTextSearch returns the TextSearch field value if set, zero value otherwise.
func (o *ContentV2GetCardsTrashPostRequestSettingsFilter) GetTextSearch() string {
	if o == nil || IsNil(o.TextSearch) {
		var ret string
		return ret
	}
	return *o.TextSearch
}

// GetTextSearchOk returns a tuple with the TextSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2GetCardsTrashPostRequestSettingsFilter) GetTextSearchOk() (*string, bool) {
	if o == nil || IsNil(o.TextSearch) {
		return nil, false
	}
	return o.TextSearch, true
}

// HasTextSearch returns a boolean if a field has been set.
func (o *ContentV2GetCardsTrashPostRequestSettingsFilter) HasTextSearch() bool {
	if o != nil && !IsNil(o.TextSearch) {
		return true
	}

	return false
}

// SetTextSearch gets a reference to the given string and assigns it to the TextSearch field.
func (o *ContentV2GetCardsTrashPostRequestSettingsFilter) SetTextSearch(v string) {
	o.TextSearch = &v
}

func (o ContentV2GetCardsTrashPostRequestSettingsFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentV2GetCardsTrashPostRequestSettingsFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TextSearch) {
		toSerialize["textSearch"] = o.TextSearch
	}
	return toSerialize, nil
}

type NullableContentV2GetCardsTrashPostRequestSettingsFilter struct {
	value *ContentV2GetCardsTrashPostRequestSettingsFilter
	isSet bool
}

func (v NullableContentV2GetCardsTrashPostRequestSettingsFilter) Get() *ContentV2GetCardsTrashPostRequestSettingsFilter {
	return v.value
}

func (v *NullableContentV2GetCardsTrashPostRequestSettingsFilter) Set(val *ContentV2GetCardsTrashPostRequestSettingsFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableContentV2GetCardsTrashPostRequestSettingsFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableContentV2GetCardsTrashPostRequestSettingsFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentV2GetCardsTrashPostRequestSettingsFilter(val *ContentV2GetCardsTrashPostRequestSettingsFilter) *NullableContentV2GetCardsTrashPostRequestSettingsFilter {
	return &NullableContentV2GetCardsTrashPostRequestSettingsFilter{value: val, isSet: true}
}

func (v NullableContentV2GetCardsTrashPostRequestSettingsFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentV2GetCardsTrashPostRequestSettingsFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
