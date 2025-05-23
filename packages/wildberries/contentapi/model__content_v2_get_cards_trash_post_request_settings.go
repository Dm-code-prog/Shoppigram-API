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

// checks if the ContentV2GetCardsTrashPostRequestSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentV2GetCardsTrashPostRequestSettings{}

// ContentV2GetCardsTrashPostRequestSettings Настройки
type ContentV2GetCardsTrashPostRequestSettings struct {
	Sort   *ContentV2GetCardsTrashPostRequestSettingsSort   `json:"sort,omitempty"`
	Cursor *ContentV2GetCardsTrashPostRequestSettingsCursor `json:"cursor,omitempty"`
	Filter *ContentV2GetCardsTrashPostRequestSettingsFilter `json:"filter,omitempty"`
}

// NewContentV2GetCardsTrashPostRequestSettings instantiates a new ContentV2GetCardsTrashPostRequestSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentV2GetCardsTrashPostRequestSettings() *ContentV2GetCardsTrashPostRequestSettings {
	this := ContentV2GetCardsTrashPostRequestSettings{}
	return &this
}

// NewContentV2GetCardsTrashPostRequestSettingsWithDefaults instantiates a new ContentV2GetCardsTrashPostRequestSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentV2GetCardsTrashPostRequestSettingsWithDefaults() *ContentV2GetCardsTrashPostRequestSettings {
	this := ContentV2GetCardsTrashPostRequestSettings{}
	return &this
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *ContentV2GetCardsTrashPostRequestSettings) GetSort() ContentV2GetCardsTrashPostRequestSettingsSort {
	if o == nil || IsNil(o.Sort) {
		var ret ContentV2GetCardsTrashPostRequestSettingsSort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2GetCardsTrashPostRequestSettings) GetSortOk() (*ContentV2GetCardsTrashPostRequestSettingsSort, bool) {
	if o == nil || IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *ContentV2GetCardsTrashPostRequestSettings) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given ContentV2GetCardsTrashPostRequestSettingsSort and assigns it to the Sort field.
func (o *ContentV2GetCardsTrashPostRequestSettings) SetSort(v ContentV2GetCardsTrashPostRequestSettingsSort) {
	o.Sort = &v
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *ContentV2GetCardsTrashPostRequestSettings) GetCursor() ContentV2GetCardsTrashPostRequestSettingsCursor {
	if o == nil || IsNil(o.Cursor) {
		var ret ContentV2GetCardsTrashPostRequestSettingsCursor
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2GetCardsTrashPostRequestSettings) GetCursorOk() (*ContentV2GetCardsTrashPostRequestSettingsCursor, bool) {
	if o == nil || IsNil(o.Cursor) {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *ContentV2GetCardsTrashPostRequestSettings) HasCursor() bool {
	if o != nil && !IsNil(o.Cursor) {
		return true
	}

	return false
}

// SetCursor gets a reference to the given ContentV2GetCardsTrashPostRequestSettingsCursor and assigns it to the Cursor field.
func (o *ContentV2GetCardsTrashPostRequestSettings) SetCursor(v ContentV2GetCardsTrashPostRequestSettingsCursor) {
	o.Cursor = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *ContentV2GetCardsTrashPostRequestSettings) GetFilter() ContentV2GetCardsTrashPostRequestSettingsFilter {
	if o == nil || IsNil(o.Filter) {
		var ret ContentV2GetCardsTrashPostRequestSettingsFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2GetCardsTrashPostRequestSettings) GetFilterOk() (*ContentV2GetCardsTrashPostRequestSettingsFilter, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *ContentV2GetCardsTrashPostRequestSettings) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given ContentV2GetCardsTrashPostRequestSettingsFilter and assigns it to the Filter field.
func (o *ContentV2GetCardsTrashPostRequestSettings) SetFilter(v ContentV2GetCardsTrashPostRequestSettingsFilter) {
	o.Filter = &v
}

func (o ContentV2GetCardsTrashPostRequestSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentV2GetCardsTrashPostRequestSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}
	if !IsNil(o.Cursor) {
		toSerialize["cursor"] = o.Cursor
	}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	return toSerialize, nil
}

type NullableContentV2GetCardsTrashPostRequestSettings struct {
	value *ContentV2GetCardsTrashPostRequestSettings
	isSet bool
}

func (v NullableContentV2GetCardsTrashPostRequestSettings) Get() *ContentV2GetCardsTrashPostRequestSettings {
	return v.value
}

func (v *NullableContentV2GetCardsTrashPostRequestSettings) Set(val *ContentV2GetCardsTrashPostRequestSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableContentV2GetCardsTrashPostRequestSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableContentV2GetCardsTrashPostRequestSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentV2GetCardsTrashPostRequestSettings(val *ContentV2GetCardsTrashPostRequestSettings) *NullableContentV2GetCardsTrashPostRequestSettings {
	return &NullableContentV2GetCardsTrashPostRequestSettings{value: val, isSet: true}
}

func (v NullableContentV2GetCardsTrashPostRequestSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentV2GetCardsTrashPostRequestSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
