/*
API цен и скидок

С помощью этих методов можно устанавливать цены и скидки. Максимум — 10 запросов за 6 секунд суммарно для всех методов раздела **Цены и скидки**.

API version:
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pricesapi

import (
	"encoding/json"
)

// checks if the ApiV2ListGoodsFilterGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV2ListGoodsFilterGet200ResponseData{}

// ApiV2ListGoodsFilterGet200ResponseData struct for ApiV2ListGoodsFilterGet200ResponseData
type ApiV2ListGoodsFilterGet200ResponseData struct {
	// Информация о товарах
	ListGoods []GoodsList `json:"listGoods,omitempty"`
}

// NewApiV2ListGoodsFilterGet200ResponseData instantiates a new ApiV2ListGoodsFilterGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV2ListGoodsFilterGet200ResponseData() *ApiV2ListGoodsFilterGet200ResponseData {
	this := ApiV2ListGoodsFilterGet200ResponseData{}
	return &this
}

// NewApiV2ListGoodsFilterGet200ResponseDataWithDefaults instantiates a new ApiV2ListGoodsFilterGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV2ListGoodsFilterGet200ResponseDataWithDefaults() *ApiV2ListGoodsFilterGet200ResponseData {
	this := ApiV2ListGoodsFilterGet200ResponseData{}
	return &this
}

// GetListGoods returns the ListGoods field value if set, zero value otherwise.
func (o *ApiV2ListGoodsFilterGet200ResponseData) GetListGoods() []GoodsList {
	if o == nil || IsNil(o.ListGoods) {
		var ret []GoodsList
		return ret
	}
	return o.ListGoods
}

// GetListGoodsOk returns a tuple with the ListGoods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV2ListGoodsFilterGet200ResponseData) GetListGoodsOk() ([]GoodsList, bool) {
	if o == nil || IsNil(o.ListGoods) {
		return nil, false
	}
	return o.ListGoods, true
}

// HasListGoods returns a boolean if a field has been set.
func (o *ApiV2ListGoodsFilterGet200ResponseData) HasListGoods() bool {
	if o != nil && !IsNil(o.ListGoods) {
		return true
	}

	return false
}

// SetListGoods gets a reference to the given []GoodsList and assigns it to the ListGoods field.
func (o *ApiV2ListGoodsFilterGet200ResponseData) SetListGoods(v []GoodsList) {
	o.ListGoods = v
}

func (o ApiV2ListGoodsFilterGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV2ListGoodsFilterGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ListGoods) {
		toSerialize["listGoods"] = o.ListGoods
	}
	return toSerialize, nil
}

type NullableApiV2ListGoodsFilterGet200ResponseData struct {
	value *ApiV2ListGoodsFilterGet200ResponseData
	isSet bool
}

func (v NullableApiV2ListGoodsFilterGet200ResponseData) Get() *ApiV2ListGoodsFilterGet200ResponseData {
	return v.value
}

func (v *NullableApiV2ListGoodsFilterGet200ResponseData) Set(val *ApiV2ListGoodsFilterGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV2ListGoodsFilterGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV2ListGoodsFilterGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV2ListGoodsFilterGet200ResponseData(val *ApiV2ListGoodsFilterGet200ResponseData) *NullableApiV2ListGoodsFilterGet200ResponseData {
	return &NullableApiV2ListGoodsFilterGet200ResponseData{value: val, isSet: true}
}

func (v NullableApiV2ListGoodsFilterGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV2ListGoodsFilterGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
