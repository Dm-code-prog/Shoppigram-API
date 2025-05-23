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

// checks if the RequestMoveNmsImtDisconn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestMoveNmsImtDisconn{}

// RequestMoveNmsImtDisconn struct for RequestMoveNmsImtDisconn
type RequestMoveNmsImtDisconn struct {
	// `nmID`, которые необходимо разъединить (max 30)
	NmIDs []int32 `json:"nmIDs"`
}

type _RequestMoveNmsImtDisconn RequestMoveNmsImtDisconn

// NewRequestMoveNmsImtDisconn instantiates a new RequestMoveNmsImtDisconn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestMoveNmsImtDisconn(nmIDs []int32) *RequestMoveNmsImtDisconn {
	this := RequestMoveNmsImtDisconn{}
	this.NmIDs = nmIDs
	return &this
}

// NewRequestMoveNmsImtDisconnWithDefaults instantiates a new RequestMoveNmsImtDisconn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestMoveNmsImtDisconnWithDefaults() *RequestMoveNmsImtDisconn {
	this := RequestMoveNmsImtDisconn{}
	return &this
}

// GetNmIDs returns the NmIDs field value
func (o *RequestMoveNmsImtDisconn) GetNmIDs() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.NmIDs
}

// GetNmIDsOk returns a tuple with the NmIDs field value
// and a boolean to check if the value has been set.
func (o *RequestMoveNmsImtDisconn) GetNmIDsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NmIDs, true
}

// SetNmIDs sets field value
func (o *RequestMoveNmsImtDisconn) SetNmIDs(v []int32) {
	o.NmIDs = v
}

func (o RequestMoveNmsImtDisconn) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestMoveNmsImtDisconn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nmIDs"] = o.NmIDs
	return toSerialize, nil
}

func (o *RequestMoveNmsImtDisconn) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"nmIDs",
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

	varRequestMoveNmsImtDisconn := _RequestMoveNmsImtDisconn{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRequestMoveNmsImtDisconn)

	if err != nil {
		return err
	}

	*o = RequestMoveNmsImtDisconn(varRequestMoveNmsImtDisconn)

	return err
}

type NullableRequestMoveNmsImtDisconn struct {
	value *RequestMoveNmsImtDisconn
	isSet bool
}

func (v NullableRequestMoveNmsImtDisconn) Get() *RequestMoveNmsImtDisconn {
	return v.value
}

func (v *NullableRequestMoveNmsImtDisconn) Set(val *RequestMoveNmsImtDisconn) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestMoveNmsImtDisconn) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestMoveNmsImtDisconn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestMoveNmsImtDisconn(val *RequestMoveNmsImtDisconn) *NullableRequestMoveNmsImtDisconn {
	return &NullableRequestMoveNmsImtDisconn{value: val, isSet: true}
}

func (v NullableRequestMoveNmsImtDisconn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestMoveNmsImtDisconn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
