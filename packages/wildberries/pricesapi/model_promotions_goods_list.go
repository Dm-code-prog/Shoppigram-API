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

// checks if the PromotionsGoodsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotionsGoodsList{}

// PromotionsGoodsList struct for PromotionsGoodsList
type PromotionsGoodsList struct {
	// ID номенклатуры
	Id *int32 `json:"id,omitempty"`
	// Участвует в акции:   - `true` — да   - `false` — нет
	InAction *bool `json:"inAction,omitempty"`
	// Текущая розничная цена
	Price *float32 `json:"price,omitempty"`
	// Валюта в формате ISO 4217
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// Плановая цена (цена во время акции)
	PlanPrice *float32 `json:"planPrice,omitempty"`
	// Текущая скидка
	Discount *int32 `json:"discount,omitempty"`
	// Рекомендуемая скидка для участия в акции
	PlanDiscount *int32 `json:"planDiscount,omitempty"`
}

// NewPromotionsGoodsList instantiates a new PromotionsGoodsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotionsGoodsList() *PromotionsGoodsList {
	this := PromotionsGoodsList{}
	return &this
}

// NewPromotionsGoodsListWithDefaults instantiates a new PromotionsGoodsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotionsGoodsListWithDefaults() *PromotionsGoodsList {
	this := PromotionsGoodsList{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PromotionsGoodsList) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionsGoodsList) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PromotionsGoodsList) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PromotionsGoodsList) SetId(v int32) {
	o.Id = &v
}

// GetInAction returns the InAction field value if set, zero value otherwise.
func (o *PromotionsGoodsList) GetInAction() bool {
	if o == nil || IsNil(o.InAction) {
		var ret bool
		return ret
	}
	return *o.InAction
}

// GetInActionOk returns a tuple with the InAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionsGoodsList) GetInActionOk() (*bool, bool) {
	if o == nil || IsNil(o.InAction) {
		return nil, false
	}
	return o.InAction, true
}

// HasInAction returns a boolean if a field has been set.
func (o *PromotionsGoodsList) HasInAction() bool {
	if o != nil && !IsNil(o.InAction) {
		return true
	}

	return false
}

// SetInAction gets a reference to the given bool and assigns it to the InAction field.
func (o *PromotionsGoodsList) SetInAction(v bool) {
	o.InAction = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *PromotionsGoodsList) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionsGoodsList) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *PromotionsGoodsList) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *PromotionsGoodsList) SetPrice(v float32) {
	o.Price = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *PromotionsGoodsList) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionsGoodsList) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *PromotionsGoodsList) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *PromotionsGoodsList) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetPlanPrice returns the PlanPrice field value if set, zero value otherwise.
func (o *PromotionsGoodsList) GetPlanPrice() float32 {
	if o == nil || IsNil(o.PlanPrice) {
		var ret float32
		return ret
	}
	return *o.PlanPrice
}

// GetPlanPriceOk returns a tuple with the PlanPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionsGoodsList) GetPlanPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.PlanPrice) {
		return nil, false
	}
	return o.PlanPrice, true
}

// HasPlanPrice returns a boolean if a field has been set.
func (o *PromotionsGoodsList) HasPlanPrice() bool {
	if o != nil && !IsNil(o.PlanPrice) {
		return true
	}

	return false
}

// SetPlanPrice gets a reference to the given float32 and assigns it to the PlanPrice field.
func (o *PromotionsGoodsList) SetPlanPrice(v float32) {
	o.PlanPrice = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *PromotionsGoodsList) GetDiscount() int32 {
	if o == nil || IsNil(o.Discount) {
		var ret int32
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionsGoodsList) GetDiscountOk() (*int32, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *PromotionsGoodsList) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given int32 and assigns it to the Discount field.
func (o *PromotionsGoodsList) SetDiscount(v int32) {
	o.Discount = &v
}

// GetPlanDiscount returns the PlanDiscount field value if set, zero value otherwise.
func (o *PromotionsGoodsList) GetPlanDiscount() int32 {
	if o == nil || IsNil(o.PlanDiscount) {
		var ret int32
		return ret
	}
	return *o.PlanDiscount
}

// GetPlanDiscountOk returns a tuple with the PlanDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionsGoodsList) GetPlanDiscountOk() (*int32, bool) {
	if o == nil || IsNil(o.PlanDiscount) {
		return nil, false
	}
	return o.PlanDiscount, true
}

// HasPlanDiscount returns a boolean if a field has been set.
func (o *PromotionsGoodsList) HasPlanDiscount() bool {
	if o != nil && !IsNil(o.PlanDiscount) {
		return true
	}

	return false
}

// SetPlanDiscount gets a reference to the given int32 and assigns it to the PlanDiscount field.
func (o *PromotionsGoodsList) SetPlanDiscount(v int32) {
	o.PlanDiscount = &v
}

func (o PromotionsGoodsList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromotionsGoodsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InAction) {
		toSerialize["inAction"] = o.InAction
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.PlanPrice) {
		toSerialize["planPrice"] = o.PlanPrice
	}
	if !IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}
	if !IsNil(o.PlanDiscount) {
		toSerialize["planDiscount"] = o.PlanDiscount
	}
	return toSerialize, nil
}

type NullablePromotionsGoodsList struct {
	value *PromotionsGoodsList
	isSet bool
}

func (v NullablePromotionsGoodsList) Get() *PromotionsGoodsList {
	return v.value
}

func (v *NullablePromotionsGoodsList) Set(val *PromotionsGoodsList) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotionsGoodsList) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotionsGoodsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotionsGoodsList(val *PromotionsGoodsList) *NullablePromotionsGoodsList {
	return &NullablePromotionsGoodsList{value: val, isSet: true}
}

func (v NullablePromotionsGoodsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotionsGoodsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
