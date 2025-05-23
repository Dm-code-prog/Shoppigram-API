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

// checks if the ContentV2ObjectAllGet200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentV2ObjectAllGet200ResponseDataInner{}

// ContentV2ObjectAllGet200ResponseDataInner struct for ContentV2ObjectAllGet200ResponseDataInner
type ContentV2ObjectAllGet200ResponseDataInner struct {
	// Идентификатор предмета
	SubjectID *int32 `json:"subjectID,omitempty"`
	// Идентификатор родительской категории
	ParentID *int32 `json:"parentID,omitempty"`
	// Название предмета
	SubjectName *string `json:"subjectName,omitempty"`
	// Название родительской категории
	ParentName *string `json:"parentName,omitempty"`
}

// NewContentV2ObjectAllGet200ResponseDataInner instantiates a new ContentV2ObjectAllGet200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentV2ObjectAllGet200ResponseDataInner() *ContentV2ObjectAllGet200ResponseDataInner {
	this := ContentV2ObjectAllGet200ResponseDataInner{}
	return &this
}

// NewContentV2ObjectAllGet200ResponseDataInnerWithDefaults instantiates a new ContentV2ObjectAllGet200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentV2ObjectAllGet200ResponseDataInnerWithDefaults() *ContentV2ObjectAllGet200ResponseDataInner {
	this := ContentV2ObjectAllGet200ResponseDataInner{}
	return &this
}

// GetSubjectID returns the SubjectID field value if set, zero value otherwise.
func (o *ContentV2ObjectAllGet200ResponseDataInner) GetSubjectID() int32 {
	if o == nil || IsNil(o.SubjectID) {
		var ret int32
		return ret
	}
	return *o.SubjectID
}

// GetSubjectIDOk returns a tuple with the SubjectID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2ObjectAllGet200ResponseDataInner) GetSubjectIDOk() (*int32, bool) {
	if o == nil || IsNil(o.SubjectID) {
		return nil, false
	}
	return o.SubjectID, true
}

// HasSubjectID returns a boolean if a field has been set.
func (o *ContentV2ObjectAllGet200ResponseDataInner) HasSubjectID() bool {
	if o != nil && !IsNil(o.SubjectID) {
		return true
	}

	return false
}

// SetSubjectID gets a reference to the given int32 and assigns it to the SubjectID field.
func (o *ContentV2ObjectAllGet200ResponseDataInner) SetSubjectID(v int32) {
	o.SubjectID = &v
}

// GetParentID returns the ParentID field value if set, zero value otherwise.
func (o *ContentV2ObjectAllGet200ResponseDataInner) GetParentID() int32 {
	if o == nil || IsNil(o.ParentID) {
		var ret int32
		return ret
	}
	return *o.ParentID
}

// GetParentIDOk returns a tuple with the ParentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2ObjectAllGet200ResponseDataInner) GetParentIDOk() (*int32, bool) {
	if o == nil || IsNil(o.ParentID) {
		return nil, false
	}
	return o.ParentID, true
}

// HasParentID returns a boolean if a field has been set.
func (o *ContentV2ObjectAllGet200ResponseDataInner) HasParentID() bool {
	if o != nil && !IsNil(o.ParentID) {
		return true
	}

	return false
}

// SetParentID gets a reference to the given int32 and assigns it to the ParentID field.
func (o *ContentV2ObjectAllGet200ResponseDataInner) SetParentID(v int32) {
	o.ParentID = &v
}

// GetSubjectName returns the SubjectName field value if set, zero value otherwise.
func (o *ContentV2ObjectAllGet200ResponseDataInner) GetSubjectName() string {
	if o == nil || IsNil(o.SubjectName) {
		var ret string
		return ret
	}
	return *o.SubjectName
}

// GetSubjectNameOk returns a tuple with the SubjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2ObjectAllGet200ResponseDataInner) GetSubjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectName) {
		return nil, false
	}
	return o.SubjectName, true
}

// HasSubjectName returns a boolean if a field has been set.
func (o *ContentV2ObjectAllGet200ResponseDataInner) HasSubjectName() bool {
	if o != nil && !IsNil(o.SubjectName) {
		return true
	}

	return false
}

// SetSubjectName gets a reference to the given string and assigns it to the SubjectName field.
func (o *ContentV2ObjectAllGet200ResponseDataInner) SetSubjectName(v string) {
	o.SubjectName = &v
}

// GetParentName returns the ParentName field value if set, zero value otherwise.
func (o *ContentV2ObjectAllGet200ResponseDataInner) GetParentName() string {
	if o == nil || IsNil(o.ParentName) {
		var ret string
		return ret
	}
	return *o.ParentName
}

// GetParentNameOk returns a tuple with the ParentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentV2ObjectAllGet200ResponseDataInner) GetParentNameOk() (*string, bool) {
	if o == nil || IsNil(o.ParentName) {
		return nil, false
	}
	return o.ParentName, true
}

// HasParentName returns a boolean if a field has been set.
func (o *ContentV2ObjectAllGet200ResponseDataInner) HasParentName() bool {
	if o != nil && !IsNil(o.ParentName) {
		return true
	}

	return false
}

// SetParentName gets a reference to the given string and assigns it to the ParentName field.
func (o *ContentV2ObjectAllGet200ResponseDataInner) SetParentName(v string) {
	o.ParentName = &v
}

func (o ContentV2ObjectAllGet200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentV2ObjectAllGet200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubjectID) {
		toSerialize["subjectID"] = o.SubjectID
	}
	if !IsNil(o.ParentID) {
		toSerialize["parentID"] = o.ParentID
	}
	if !IsNil(o.SubjectName) {
		toSerialize["subjectName"] = o.SubjectName
	}
	if !IsNil(o.ParentName) {
		toSerialize["parentName"] = o.ParentName
	}
	return toSerialize, nil
}

type NullableContentV2ObjectAllGet200ResponseDataInner struct {
	value *ContentV2ObjectAllGet200ResponseDataInner
	isSet bool
}

func (v NullableContentV2ObjectAllGet200ResponseDataInner) Get() *ContentV2ObjectAllGet200ResponseDataInner {
	return v.value
}

func (v *NullableContentV2ObjectAllGet200ResponseDataInner) Set(val *ContentV2ObjectAllGet200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableContentV2ObjectAllGet200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableContentV2ObjectAllGet200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentV2ObjectAllGet200ResponseDataInner(val *ContentV2ObjectAllGet200ResponseDataInner) *NullableContentV2ObjectAllGet200ResponseDataInner {
	return &NullableContentV2ObjectAllGet200ResponseDataInner{value: val, isSet: true}
}

func (v NullableContentV2ObjectAllGet200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentV2ObjectAllGet200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
