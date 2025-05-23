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

// checks if the ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner{}

// ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner struct for ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner
type ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner struct {
	// ID акции
	Id *int32 `json:"id,omitempty"`
	// Название акции
	Name *string `json:"name,omitempty"`
	// Описание акции
	Description *string `json:"description,omitempty"`
	// Преимущества акции
	Advantages []string `json:"advantages,omitempty"`
	// Начало акции
	StartDateTime *string `json:"startDateTime,omitempty"`
	// Конец акции
	EndDateTime *string `json:"endDateTime,omitempty"`
	// Количество товаров с остатками, участвующих в акции
	InPromoActionLeftovers *int32 `json:"inPromoActionLeftovers,omitempty"`
	// Общее количество товаров, участвующих в акции
	InPromoActionTotal *int32 `json:"inPromoActionTotal,omitempty"`
	// Количество товаров с остатками, не участвующих в акции
	NotInPromoActionLeftovers *int32 `json:"notInPromoActionLeftovers,omitempty"`
	// Общее количество товаров, не участвующих в акции
	NotInPromoActionTotal *int32 `json:"notInPromoActionTotal,omitempty"`
	// Уже участвующие в акции товары, %. Рассчитывается по товарам в акции и с остатком
	ParticipationPercentage *int32 `json:"participationPercentage,omitempty"`
	// Тип акции:   - `regular` — акция   - `auto` — автоакция
	Type *string `json:"type,omitempty"`
	// Количество товаров, исключенных из автоакции до её старта. Только при `\"type\": \"auto\"`. <br>В момент старта акции эти товары автоматически будут без скидки
	ExceptionProductsCount *int32 `json:"exceptionProductsCount,omitempty"`
}

// NewApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner instantiates a new ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner() *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner {
	this := ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner{}
	return &this
}

// NewApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInnerWithDefaults instantiates a new ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInnerWithDefaults() *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner {
	this := ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) SetDescription(v string) {
	o.Description = &v
}

// GetAdvantages returns the Advantages field value if set, zero value otherwise.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetAdvantages() []string {
	if o == nil || IsNil(o.Advantages) {
		var ret []string
		return ret
	}
	return o.Advantages
}

// GetAdvantagesOk returns a tuple with the Advantages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetAdvantagesOk() ([]string, bool) {
	if o == nil || IsNil(o.Advantages) {
		return nil, false
	}
	return o.Advantages, true
}

// HasAdvantages returns a boolean if a field has been set.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) HasAdvantages() bool {
	if o != nil && !IsNil(o.Advantages) {
		return true
	}

	return false
}

// SetAdvantages gets a reference to the given []string and assigns it to the Advantages field.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) SetAdvantages(v []string) {
	o.Advantages = v
}

// GetStartDateTime returns the StartDateTime field value if set, zero value otherwise.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetStartDateTime() string {
	if o == nil || IsNil(o.StartDateTime) {
		var ret string
		return ret
	}
	return *o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetStartDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartDateTime) {
		return nil, false
	}
	return o.StartDateTime, true
}

// HasStartDateTime returns a boolean if a field has been set.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) HasStartDateTime() bool {
	if o != nil && !IsNil(o.StartDateTime) {
		return true
	}

	return false
}

// SetStartDateTime gets a reference to the given string and assigns it to the StartDateTime field.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) SetStartDateTime(v string) {
	o.StartDateTime = &v
}

// GetEndDateTime returns the EndDateTime field value if set, zero value otherwise.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetEndDateTime() string {
	if o == nil || IsNil(o.EndDateTime) {
		var ret string
		return ret
	}
	return *o.EndDateTime
}

// GetEndDateTimeOk returns a tuple with the EndDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetEndDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndDateTime) {
		return nil, false
	}
	return o.EndDateTime, true
}

// HasEndDateTime returns a boolean if a field has been set.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) HasEndDateTime() bool {
	if o != nil && !IsNil(o.EndDateTime) {
		return true
	}

	return false
}

// SetEndDateTime gets a reference to the given string and assigns it to the EndDateTime field.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) SetEndDateTime(v string) {
	o.EndDateTime = &v
}

// GetInPromoActionLeftovers returns the InPromoActionLeftovers field value if set, zero value otherwise.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetInPromoActionLeftovers() int32 {
	if o == nil || IsNil(o.InPromoActionLeftovers) {
		var ret int32
		return ret
	}
	return *o.InPromoActionLeftovers
}

// GetInPromoActionLeftoversOk returns a tuple with the InPromoActionLeftovers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetInPromoActionLeftoversOk() (*int32, bool) {
	if o == nil || IsNil(o.InPromoActionLeftovers) {
		return nil, false
	}
	return o.InPromoActionLeftovers, true
}

// HasInPromoActionLeftovers returns a boolean if a field has been set.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) HasInPromoActionLeftovers() bool {
	if o != nil && !IsNil(o.InPromoActionLeftovers) {
		return true
	}

	return false
}

// SetInPromoActionLeftovers gets a reference to the given int32 and assigns it to the InPromoActionLeftovers field.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) SetInPromoActionLeftovers(v int32) {
	o.InPromoActionLeftovers = &v
}

// GetInPromoActionTotal returns the InPromoActionTotal field value if set, zero value otherwise.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetInPromoActionTotal() int32 {
	if o == nil || IsNil(o.InPromoActionTotal) {
		var ret int32
		return ret
	}
	return *o.InPromoActionTotal
}

// GetInPromoActionTotalOk returns a tuple with the InPromoActionTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetInPromoActionTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.InPromoActionTotal) {
		return nil, false
	}
	return o.InPromoActionTotal, true
}

// HasInPromoActionTotal returns a boolean if a field has been set.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) HasInPromoActionTotal() bool {
	if o != nil && !IsNil(o.InPromoActionTotal) {
		return true
	}

	return false
}

// SetInPromoActionTotal gets a reference to the given int32 and assigns it to the InPromoActionTotal field.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) SetInPromoActionTotal(v int32) {
	o.InPromoActionTotal = &v
}

// GetNotInPromoActionLeftovers returns the NotInPromoActionLeftovers field value if set, zero value otherwise.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetNotInPromoActionLeftovers() int32 {
	if o == nil || IsNil(o.NotInPromoActionLeftovers) {
		var ret int32
		return ret
	}
	return *o.NotInPromoActionLeftovers
}

// GetNotInPromoActionLeftoversOk returns a tuple with the NotInPromoActionLeftovers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetNotInPromoActionLeftoversOk() (*int32, bool) {
	if o == nil || IsNil(o.NotInPromoActionLeftovers) {
		return nil, false
	}
	return o.NotInPromoActionLeftovers, true
}

// HasNotInPromoActionLeftovers returns a boolean if a field has been set.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) HasNotInPromoActionLeftovers() bool {
	if o != nil && !IsNil(o.NotInPromoActionLeftovers) {
		return true
	}

	return false
}

// SetNotInPromoActionLeftovers gets a reference to the given int32 and assigns it to the NotInPromoActionLeftovers field.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) SetNotInPromoActionLeftovers(v int32) {
	o.NotInPromoActionLeftovers = &v
}

// GetNotInPromoActionTotal returns the NotInPromoActionTotal field value if set, zero value otherwise.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetNotInPromoActionTotal() int32 {
	if o == nil || IsNil(o.NotInPromoActionTotal) {
		var ret int32
		return ret
	}
	return *o.NotInPromoActionTotal
}

// GetNotInPromoActionTotalOk returns a tuple with the NotInPromoActionTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetNotInPromoActionTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.NotInPromoActionTotal) {
		return nil, false
	}
	return o.NotInPromoActionTotal, true
}

// HasNotInPromoActionTotal returns a boolean if a field has been set.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) HasNotInPromoActionTotal() bool {
	if o != nil && !IsNil(o.NotInPromoActionTotal) {
		return true
	}

	return false
}

// SetNotInPromoActionTotal gets a reference to the given int32 and assigns it to the NotInPromoActionTotal field.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) SetNotInPromoActionTotal(v int32) {
	o.NotInPromoActionTotal = &v
}

// GetParticipationPercentage returns the ParticipationPercentage field value if set, zero value otherwise.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetParticipationPercentage() int32 {
	if o == nil || IsNil(o.ParticipationPercentage) {
		var ret int32
		return ret
	}
	return *o.ParticipationPercentage
}

// GetParticipationPercentageOk returns a tuple with the ParticipationPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetParticipationPercentageOk() (*int32, bool) {
	if o == nil || IsNil(o.ParticipationPercentage) {
		return nil, false
	}
	return o.ParticipationPercentage, true
}

// HasParticipationPercentage returns a boolean if a field has been set.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) HasParticipationPercentage() bool {
	if o != nil && !IsNil(o.ParticipationPercentage) {
		return true
	}

	return false
}

// SetParticipationPercentage gets a reference to the given int32 and assigns it to the ParticipationPercentage field.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) SetParticipationPercentage(v int32) {
	o.ParticipationPercentage = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) SetType(v string) {
	o.Type = &v
}

// GetExceptionProductsCount returns the ExceptionProductsCount field value if set, zero value otherwise.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetExceptionProductsCount() int32 {
	if o == nil || IsNil(o.ExceptionProductsCount) {
		var ret int32
		return ret
	}
	return *o.ExceptionProductsCount
}

// GetExceptionProductsCountOk returns a tuple with the ExceptionProductsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetExceptionProductsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ExceptionProductsCount) {
		return nil, false
	}
	return o.ExceptionProductsCount, true
}

// HasExceptionProductsCount returns a boolean if a field has been set.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) HasExceptionProductsCount() bool {
	if o != nil && !IsNil(o.ExceptionProductsCount) {
		return true
	}

	return false
}

// SetExceptionProductsCount gets a reference to the given int32 and assigns it to the ExceptionProductsCount field.
func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) SetExceptionProductsCount(v int32) {
	o.ExceptionProductsCount = &v
}

func (o ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Advantages) {
		toSerialize["advantages"] = o.Advantages
	}
	if !IsNil(o.StartDateTime) {
		toSerialize["startDateTime"] = o.StartDateTime
	}
	if !IsNil(o.EndDateTime) {
		toSerialize["endDateTime"] = o.EndDateTime
	}
	if !IsNil(o.InPromoActionLeftovers) {
		toSerialize["inPromoActionLeftovers"] = o.InPromoActionLeftovers
	}
	if !IsNil(o.InPromoActionTotal) {
		toSerialize["inPromoActionTotal"] = o.InPromoActionTotal
	}
	if !IsNil(o.NotInPromoActionLeftovers) {
		toSerialize["notInPromoActionLeftovers"] = o.NotInPromoActionLeftovers
	}
	if !IsNil(o.NotInPromoActionTotal) {
		toSerialize["notInPromoActionTotal"] = o.NotInPromoActionTotal
	}
	if !IsNil(o.ParticipationPercentage) {
		toSerialize["participationPercentage"] = o.ParticipationPercentage
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ExceptionProductsCount) {
		toSerialize["exceptionProductsCount"] = o.ExceptionProductsCount
	}
	return toSerialize, nil
}

type NullableApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner struct {
	value *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner
	isSet bool
}

func (v NullableApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) Get() *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner {
	return v.value
}

func (v *NullableApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) Set(val *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner(val *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) *NullableApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner {
	return &NullableApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner{value: val, isSet: true}
}

func (v NullableApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
