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

// checks if the TaskAlreadyExistsError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskAlreadyExistsError{}

// TaskAlreadyExistsError struct for TaskAlreadyExistsError
type TaskAlreadyExistsError struct {
	Data *TaskAlreadyExistsErrorData `json:"data,omitempty"`
	// Флаг ошибки
	Error *bool `json:"error,omitempty"`
	// Текст ошибки
	ErrorText *string `json:"errorText,omitempty"`
}

// NewTaskAlreadyExistsError instantiates a new TaskAlreadyExistsError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskAlreadyExistsError() *TaskAlreadyExistsError {
	this := TaskAlreadyExistsError{}
	return &this
}

// NewTaskAlreadyExistsErrorWithDefaults instantiates a new TaskAlreadyExistsError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskAlreadyExistsErrorWithDefaults() *TaskAlreadyExistsError {
	this := TaskAlreadyExistsError{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TaskAlreadyExistsError) GetData() TaskAlreadyExistsErrorData {
	if o == nil || IsNil(o.Data) {
		var ret TaskAlreadyExistsErrorData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskAlreadyExistsError) GetDataOk() (*TaskAlreadyExistsErrorData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TaskAlreadyExistsError) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given TaskAlreadyExistsErrorData and assigns it to the Data field.
func (o *TaskAlreadyExistsError) SetData(v TaskAlreadyExistsErrorData) {
	o.Data = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *TaskAlreadyExistsError) GetError() bool {
	if o == nil || IsNil(o.Error) {
		var ret bool
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskAlreadyExistsError) GetErrorOk() (*bool, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *TaskAlreadyExistsError) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given bool and assigns it to the Error field.
func (o *TaskAlreadyExistsError) SetError(v bool) {
	o.Error = &v
}

// GetErrorText returns the ErrorText field value if set, zero value otherwise.
func (o *TaskAlreadyExistsError) GetErrorText() string {
	if o == nil || IsNil(o.ErrorText) {
		var ret string
		return ret
	}
	return *o.ErrorText
}

// GetErrorTextOk returns a tuple with the ErrorText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskAlreadyExistsError) GetErrorTextOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorText) {
		return nil, false
	}
	return o.ErrorText, true
}

// HasErrorText returns a boolean if a field has been set.
func (o *TaskAlreadyExistsError) HasErrorText() bool {
	if o != nil && !IsNil(o.ErrorText) {
		return true
	}

	return false
}

// SetErrorText gets a reference to the given string and assigns it to the ErrorText field.
func (o *TaskAlreadyExistsError) SetErrorText(v string) {
	o.ErrorText = &v
}

func (o TaskAlreadyExistsError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskAlreadyExistsError) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableTaskAlreadyExistsError struct {
	value *TaskAlreadyExistsError
	isSet bool
}

func (v NullableTaskAlreadyExistsError) Get() *TaskAlreadyExistsError {
	return v.value
}

func (v *NullableTaskAlreadyExistsError) Set(val *TaskAlreadyExistsError) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskAlreadyExistsError) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskAlreadyExistsError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskAlreadyExistsError(val *TaskAlreadyExistsError) *NullableTaskAlreadyExistsError {
	return &NullableTaskAlreadyExistsError{value: val, isSet: true}
}

func (v NullableTaskAlreadyExistsError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskAlreadyExistsError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
