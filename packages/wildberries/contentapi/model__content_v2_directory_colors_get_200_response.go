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

// checks if the ContentV2DirectoryColorsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentV2DirectoryColorsGet200Response{}

// ContentV2DirectoryColorsGet200Response struct for ContentV2DirectoryColorsGet200Response
type ContentV2DirectoryColorsGet200Response struct {
	Data []ContentV2DirectoryColorsGet200ResponseDataInner `json:"data,omitempty"`
	// Флаг ошибки
	Error *bool `json:"error,omitempty"`
	// Описание ошибки
	ErrorText *string `json:"errorText,omitempty"`
	// Дополнительные ошибки
	AdditionalErrors *string `json:"additionalErrors,omitempty"`
}

// NewContentV2DirectoryColorsGet200Response instantiates a new ContentV2DirectoryColorsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentV2DirectoryColorsGet200Response() *ContentV2DirectoryColorsGet200Response {
	this := ContentV2DirectoryColorsGet200Response{}
	return &this
}

// NewContentV2DirectoryColorsGet200ResponseWithDefaults instantiates a new ContentV2DirectoryColorsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentV2DirectoryColorsGet200ResponseWithDefaults() *ContentV2DirectoryColorsGet200Response {
	this := ContentV2DirectoryColorsGet200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ContentV2DirectoryColorsGet200Response) GetData() []ContentV2DirectoryColorsGet200ResponseDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []ContentV2DirectoryColorsGet200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2DirectoryColorsGet200Response) GetDataOk() ([]ContentV2DirectoryColorsGet200ResponseDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ContentV2DirectoryColorsGet200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ContentV2DirectoryColorsGet200ResponseDataInner and assigns it to the Data field.
func (o *ContentV2DirectoryColorsGet200Response) SetData(v []ContentV2DirectoryColorsGet200ResponseDataInner) {
	o.Data = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ContentV2DirectoryColorsGet200Response) GetError() bool {
	if o == nil || IsNil(o.Error) {
		var ret bool
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2DirectoryColorsGet200Response) GetErrorOk() (*bool, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ContentV2DirectoryColorsGet200Response) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given bool and assigns it to the Error field.
func (o *ContentV2DirectoryColorsGet200Response) SetError(v bool) {
	o.Error = &v
}

// GetErrorText returns the ErrorText field value if set, zero value otherwise.
func (o *ContentV2DirectoryColorsGet200Response) GetErrorText() string {
	if o == nil || IsNil(o.ErrorText) {
		var ret string
		return ret
	}
	return *o.ErrorText
}

// GetErrorTextOk returns a tuple with the ErrorText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2DirectoryColorsGet200Response) GetErrorTextOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorText) {
		return nil, false
	}
	return o.ErrorText, true
}

// HasErrorText returns a boolean if a field has been set.
func (o *ContentV2DirectoryColorsGet200Response) HasErrorText() bool {
	if o != nil && !IsNil(o.ErrorText) {
		return true
	}

	return false
}

// SetErrorText gets a reference to the given string and assigns it to the ErrorText field.
func (o *ContentV2DirectoryColorsGet200Response) SetErrorText(v string) {
	o.ErrorText = &v
}

// GetAdditionalErrors returns the AdditionalErrors field value if set, zero value otherwise.
func (o *ContentV2DirectoryColorsGet200Response) GetAdditionalErrors() string {
	if o == nil || IsNil(o.AdditionalErrors) {
		var ret string
		return ret
	}
	return *o.AdditionalErrors
}

// GetAdditionalErrorsOk returns a tuple with the AdditionalErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2DirectoryColorsGet200Response) GetAdditionalErrorsOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalErrors) {
		return nil, false
	}
	return o.AdditionalErrors, true
}

// HasAdditionalErrors returns a boolean if a field has been set.
func (o *ContentV2DirectoryColorsGet200Response) HasAdditionalErrors() bool {
	if o != nil && !IsNil(o.AdditionalErrors) {
		return true
	}

	return false
}

// SetAdditionalErrors gets a reference to the given string and assigns it to the AdditionalErrors field.
func (o *ContentV2DirectoryColorsGet200Response) SetAdditionalErrors(v string) {
	o.AdditionalErrors = &v
}

func (o ContentV2DirectoryColorsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentV2DirectoryColorsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.ErrorText) {
		toSerialize["errorText"] = o.ErrorText
	}
	if !IsNil(o.AdditionalErrors) {
		toSerialize["additionalErrors"] = o.AdditionalErrors
	}
	return toSerialize, nil
}

type NullableContentV2DirectoryColorsGet200Response struct {
	value *ContentV2DirectoryColorsGet200Response
	isSet bool
}

func (v NullableContentV2DirectoryColorsGet200Response) Get() *ContentV2DirectoryColorsGet200Response {
	return v.value
}

func (v *NullableContentV2DirectoryColorsGet200Response) Set(val *ContentV2DirectoryColorsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableContentV2DirectoryColorsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableContentV2DirectoryColorsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentV2DirectoryColorsGet200Response(val *ContentV2DirectoryColorsGet200Response) *NullableContentV2DirectoryColorsGet200Response {
	return &NullableContentV2DirectoryColorsGet200Response{value: val, isSet: true}
}

func (v NullableContentV2DirectoryColorsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentV2DirectoryColorsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
