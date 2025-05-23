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

// checks if the ApiV2UploadTaskSizePostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV2UploadTaskSizePostRequest{}

// ApiV2UploadTaskSizePostRequest struct for ApiV2UploadTaskSizePostRequest
type ApiV2UploadTaskSizePostRequest struct {
	// Размеры и цены для них. Максимум 1 000 размеров <br><br> Для товаров с поразмерной установкой цен карантин не применяется
	Data []SizeGoodReq `json:"data,omitempty"`
}

// NewApiV2UploadTaskSizePostRequest instantiates a new ApiV2UploadTaskSizePostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV2UploadTaskSizePostRequest() *ApiV2UploadTaskSizePostRequest {
	this := ApiV2UploadTaskSizePostRequest{}
	return &this
}

// NewApiV2UploadTaskSizePostRequestWithDefaults instantiates a new ApiV2UploadTaskSizePostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV2UploadTaskSizePostRequestWithDefaults() *ApiV2UploadTaskSizePostRequest {
	this := ApiV2UploadTaskSizePostRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ApiV2UploadTaskSizePostRequest) GetData() []SizeGoodReq {
	if o == nil || IsNil(o.Data) {
		var ret []SizeGoodReq
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV2UploadTaskSizePostRequest) GetDataOk() ([]SizeGoodReq, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ApiV2UploadTaskSizePostRequest) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []SizeGoodReq and assigns it to the Data field.
func (o *ApiV2UploadTaskSizePostRequest) SetData(v []SizeGoodReq) {
	o.Data = v
}

func (o ApiV2UploadTaskSizePostRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV2UploadTaskSizePostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableApiV2UploadTaskSizePostRequest struct {
	value *ApiV2UploadTaskSizePostRequest
	isSet bool
}

func (v NullableApiV2UploadTaskSizePostRequest) Get() *ApiV2UploadTaskSizePostRequest {
	return v.value
}

func (v *NullableApiV2UploadTaskSizePostRequest) Set(val *ApiV2UploadTaskSizePostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV2UploadTaskSizePostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV2UploadTaskSizePostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV2UploadTaskSizePostRequest(val *ApiV2UploadTaskSizePostRequest) *NullableApiV2UploadTaskSizePostRequest {
	return &NullableApiV2UploadTaskSizePostRequest{value: val, isSet: true}
}

func (v NullableApiV2UploadTaskSizePostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV2UploadTaskSizePostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
