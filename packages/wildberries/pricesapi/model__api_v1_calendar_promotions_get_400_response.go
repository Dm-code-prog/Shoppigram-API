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

// checks if the ApiV1CalendarPromotionsGet400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV1CalendarPromotionsGet400Response{}

// ApiV1CalendarPromotionsGet400Response struct for ApiV1CalendarPromotionsGet400Response
type ApiV1CalendarPromotionsGet400Response struct {
	// Текст ошибки
	ErrorText *string `json:"errorText,omitempty"`
}

// NewApiV1CalendarPromotionsGet400Response instantiates a new ApiV1CalendarPromotionsGet400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV1CalendarPromotionsGet400Response() *ApiV1CalendarPromotionsGet400Response {
	this := ApiV1CalendarPromotionsGet400Response{}
	return &this
}

// NewApiV1CalendarPromotionsGet400ResponseWithDefaults instantiates a new ApiV1CalendarPromotionsGet400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV1CalendarPromotionsGet400ResponseWithDefaults() *ApiV1CalendarPromotionsGet400Response {
	this := ApiV1CalendarPromotionsGet400Response{}
	return &this
}

// GetErrorText returns the ErrorText field value if set, zero value otherwise.
func (o *ApiV1CalendarPromotionsGet400Response) GetErrorText() string {
	if o == nil || IsNil(o.ErrorText) {
		var ret string
		return ret
	}
	return *o.ErrorText
}

// GetErrorTextOk returns a tuple with the ErrorText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1CalendarPromotionsGet400Response) GetErrorTextOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorText) {
		return nil, false
	}
	return o.ErrorText, true
}

// HasErrorText returns a boolean if a field has been set.
func (o *ApiV1CalendarPromotionsGet400Response) HasErrorText() bool {
	if o != nil && !IsNil(o.ErrorText) {
		return true
	}

	return false
}

// SetErrorText gets a reference to the given string and assigns it to the ErrorText field.
func (o *ApiV1CalendarPromotionsGet400Response) SetErrorText(v string) {
	o.ErrorText = &v
}

func (o ApiV1CalendarPromotionsGet400Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV1CalendarPromotionsGet400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorText) {
		toSerialize["errorText"] = o.ErrorText
	}
	return toSerialize, nil
}

type NullableApiV1CalendarPromotionsGet400Response struct {
	value *ApiV1CalendarPromotionsGet400Response
	isSet bool
}

func (v NullableApiV1CalendarPromotionsGet400Response) Get() *ApiV1CalendarPromotionsGet400Response {
	return v.value
}

func (v *NullableApiV1CalendarPromotionsGet400Response) Set(val *ApiV1CalendarPromotionsGet400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV1CalendarPromotionsGet400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV1CalendarPromotionsGet400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV1CalendarPromotionsGet400Response(val *ApiV1CalendarPromotionsGet400Response) *NullableApiV1CalendarPromotionsGet400Response {
	return &NullableApiV1CalendarPromotionsGet400Response{value: val, isSet: true}
}

func (v NullableApiV1CalendarPromotionsGet400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV1CalendarPromotionsGet400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
