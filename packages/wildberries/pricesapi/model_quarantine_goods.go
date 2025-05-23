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

// checks if the QuarantineGoods type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuarantineGoods{}

// QuarantineGoods struct for QuarantineGoods
type QuarantineGoods struct {
	// Артикул WB
	NmID *int32 `json:"nmID,omitempty"`
	// Не используется
	SizeID *int32 `json:"sizeID,omitempty"`
	// Не используется
	TechSizeName *string `json:"techSizeName,omitempty"`
	// Валюта по стандарту ISO 4217
	CurrencyIsoCode4217 *string `json:"currencyIsoCode4217,omitempty"`
	// Новая цена продавца до скидки
	NewPrice *float32 `json:"newPrice,omitempty"`
	// Текущая цена продавца до скидки
	OldPrice *float32 `json:"oldPrice,omitempty"`
	// Новая скидка продавца, %
	NewDiscount *int32 `json:"newDiscount,omitempty"`
	// Текущая скидка продавца, %
	OldDiscount *int32 `json:"oldDiscount,omitempty"`
	// Разница: `newPrice` * (1 - `newDiscount` / 100) - `oldPrice` * (1 - `oldDiscount` / 100)
	PriceDiff *float32 `json:"priceDiff,omitempty"`
}

// NewQuarantineGoods instantiates a new QuarantineGoods object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuarantineGoods() *QuarantineGoods {
	this := QuarantineGoods{}
	return &this
}

// NewQuarantineGoodsWithDefaults instantiates a new QuarantineGoods object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuarantineGoodsWithDefaults() *QuarantineGoods {
	this := QuarantineGoods{}
	return &this
}

// GetNmID returns the NmID field value if set, zero value otherwise.
func (o *QuarantineGoods) GetNmID() int32 {
	if o == nil || IsNil(o.NmID) {
		var ret int32
		return ret
	}
	return *o.NmID
}

// GetNmIDOk returns a tuple with the NmID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuarantineGoods) GetNmIDOk() (*int32, bool) {
	if o == nil || IsNil(o.NmID) {
		return nil, false
	}
	return o.NmID, true
}

// HasNmID returns a boolean if a field has been set.
func (o *QuarantineGoods) HasNmID() bool {
	if o != nil && !IsNil(o.NmID) {
		return true
	}

	return false
}

// SetNmID gets a reference to the given int32 and assigns it to the NmID field.
func (o *QuarantineGoods) SetNmID(v int32) {
	o.NmID = &v
}

// GetSizeID returns the SizeID field value if set, zero value otherwise.
func (o *QuarantineGoods) GetSizeID() int32 {
	if o == nil || IsNil(o.SizeID) {
		var ret int32
		return ret
	}
	return *o.SizeID
}

// GetSizeIDOk returns a tuple with the SizeID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuarantineGoods) GetSizeIDOk() (*int32, bool) {
	if o == nil || IsNil(o.SizeID) {
		return nil, false
	}
	return o.SizeID, true
}

// HasSizeID returns a boolean if a field has been set.
func (o *QuarantineGoods) HasSizeID() bool {
	if o != nil && !IsNil(o.SizeID) {
		return true
	}

	return false
}

// SetSizeID gets a reference to the given int32 and assigns it to the SizeID field.
func (o *QuarantineGoods) SetSizeID(v int32) {
	o.SizeID = &v
}

// GetTechSizeName returns the TechSizeName field value if set, zero value otherwise.
func (o *QuarantineGoods) GetTechSizeName() string {
	if o == nil || IsNil(o.TechSizeName) {
		var ret string
		return ret
	}
	return *o.TechSizeName
}

// GetTechSizeNameOk returns a tuple with the TechSizeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuarantineGoods) GetTechSizeNameOk() (*string, bool) {
	if o == nil || IsNil(o.TechSizeName) {
		return nil, false
	}
	return o.TechSizeName, true
}

// HasTechSizeName returns a boolean if a field has been set.
func (o *QuarantineGoods) HasTechSizeName() bool {
	if o != nil && !IsNil(o.TechSizeName) {
		return true
	}

	return false
}

// SetTechSizeName gets a reference to the given string and assigns it to the TechSizeName field.
func (o *QuarantineGoods) SetTechSizeName(v string) {
	o.TechSizeName = &v
}

// GetCurrencyIsoCode4217 returns the CurrencyIsoCode4217 field value if set, zero value otherwise.
func (o *QuarantineGoods) GetCurrencyIsoCode4217() string {
	if o == nil || IsNil(o.CurrencyIsoCode4217) {
		var ret string
		return ret
	}
	return *o.CurrencyIsoCode4217
}

// GetCurrencyIsoCode4217Ok returns a tuple with the CurrencyIsoCode4217 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuarantineGoods) GetCurrencyIsoCode4217Ok() (*string, bool) {
	if o == nil || IsNil(o.CurrencyIsoCode4217) {
		return nil, false
	}
	return o.CurrencyIsoCode4217, true
}

// HasCurrencyIsoCode4217 returns a boolean if a field has been set.
func (o *QuarantineGoods) HasCurrencyIsoCode4217() bool {
	if o != nil && !IsNil(o.CurrencyIsoCode4217) {
		return true
	}

	return false
}

// SetCurrencyIsoCode4217 gets a reference to the given string and assigns it to the CurrencyIsoCode4217 field.
func (o *QuarantineGoods) SetCurrencyIsoCode4217(v string) {
	o.CurrencyIsoCode4217 = &v
}

// GetNewPrice returns the NewPrice field value if set, zero value otherwise.
func (o *QuarantineGoods) GetNewPrice() float32 {
	if o == nil || IsNil(o.NewPrice) {
		var ret float32
		return ret
	}
	return *o.NewPrice
}

// GetNewPriceOk returns a tuple with the NewPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuarantineGoods) GetNewPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.NewPrice) {
		return nil, false
	}
	return o.NewPrice, true
}

// HasNewPrice returns a boolean if a field has been set.
func (o *QuarantineGoods) HasNewPrice() bool {
	if o != nil && !IsNil(o.NewPrice) {
		return true
	}

	return false
}

// SetNewPrice gets a reference to the given float32 and assigns it to the NewPrice field.
func (o *QuarantineGoods) SetNewPrice(v float32) {
	o.NewPrice = &v
}

// GetOldPrice returns the OldPrice field value if set, zero value otherwise.
func (o *QuarantineGoods) GetOldPrice() float32 {
	if o == nil || IsNil(o.OldPrice) {
		var ret float32
		return ret
	}
	return *o.OldPrice
}

// GetOldPriceOk returns a tuple with the OldPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuarantineGoods) GetOldPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.OldPrice) {
		return nil, false
	}
	return o.OldPrice, true
}

// HasOldPrice returns a boolean if a field has been set.
func (o *QuarantineGoods) HasOldPrice() bool {
	if o != nil && !IsNil(o.OldPrice) {
		return true
	}

	return false
}

// SetOldPrice gets a reference to the given float32 and assigns it to the OldPrice field.
func (o *QuarantineGoods) SetOldPrice(v float32) {
	o.OldPrice = &v
}

// GetNewDiscount returns the NewDiscount field value if set, zero value otherwise.
func (o *QuarantineGoods) GetNewDiscount() int32 {
	if o == nil || IsNil(o.NewDiscount) {
		var ret int32
		return ret
	}
	return *o.NewDiscount
}

// GetNewDiscountOk returns a tuple with the NewDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuarantineGoods) GetNewDiscountOk() (*int32, bool) {
	if o == nil || IsNil(o.NewDiscount) {
		return nil, false
	}
	return o.NewDiscount, true
}

// HasNewDiscount returns a boolean if a field has been set.
func (o *QuarantineGoods) HasNewDiscount() bool {
	if o != nil && !IsNil(o.NewDiscount) {
		return true
	}

	return false
}

// SetNewDiscount gets a reference to the given int32 and assigns it to the NewDiscount field.
func (o *QuarantineGoods) SetNewDiscount(v int32) {
	o.NewDiscount = &v
}

// GetOldDiscount returns the OldDiscount field value if set, zero value otherwise.
func (o *QuarantineGoods) GetOldDiscount() int32 {
	if o == nil || IsNil(o.OldDiscount) {
		var ret int32
		return ret
	}
	return *o.OldDiscount
}

// GetOldDiscountOk returns a tuple with the OldDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuarantineGoods) GetOldDiscountOk() (*int32, bool) {
	if o == nil || IsNil(o.OldDiscount) {
		return nil, false
	}
	return o.OldDiscount, true
}

// HasOldDiscount returns a boolean if a field has been set.
func (o *QuarantineGoods) HasOldDiscount() bool {
	if o != nil && !IsNil(o.OldDiscount) {
		return true
	}

	return false
}

// SetOldDiscount gets a reference to the given int32 and assigns it to the OldDiscount field.
func (o *QuarantineGoods) SetOldDiscount(v int32) {
	o.OldDiscount = &v
}

// GetPriceDiff returns the PriceDiff field value if set, zero value otherwise.
func (o *QuarantineGoods) GetPriceDiff() float32 {
	if o == nil || IsNil(o.PriceDiff) {
		var ret float32
		return ret
	}
	return *o.PriceDiff
}

// GetPriceDiffOk returns a tuple with the PriceDiff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuarantineGoods) GetPriceDiffOk() (*float32, bool) {
	if o == nil || IsNil(o.PriceDiff) {
		return nil, false
	}
	return o.PriceDiff, true
}

// HasPriceDiff returns a boolean if a field has been set.
func (o *QuarantineGoods) HasPriceDiff() bool {
	if o != nil && !IsNil(o.PriceDiff) {
		return true
	}

	return false
}

// SetPriceDiff gets a reference to the given float32 and assigns it to the PriceDiff field.
func (o *QuarantineGoods) SetPriceDiff(v float32) {
	o.PriceDiff = &v
}

func (o QuarantineGoods) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuarantineGoods) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NmID) {
		toSerialize["nmID"] = o.NmID
	}
	if !IsNil(o.SizeID) {
		toSerialize["sizeID"] = o.SizeID
	}
	if !IsNil(o.TechSizeName) {
		toSerialize["techSizeName"] = o.TechSizeName
	}
	if !IsNil(o.CurrencyIsoCode4217) {
		toSerialize["currencyIsoCode4217"] = o.CurrencyIsoCode4217
	}
	if !IsNil(o.NewPrice) {
		toSerialize["newPrice"] = o.NewPrice
	}
	if !IsNil(o.OldPrice) {
		toSerialize["oldPrice"] = o.OldPrice
	}
	if !IsNil(o.NewDiscount) {
		toSerialize["newDiscount"] = o.NewDiscount
	}
	if !IsNil(o.OldDiscount) {
		toSerialize["oldDiscount"] = o.OldDiscount
	}
	if !IsNil(o.PriceDiff) {
		toSerialize["priceDiff"] = o.PriceDiff
	}
	return toSerialize, nil
}

type NullableQuarantineGoods struct {
	value *QuarantineGoods
	isSet bool
}

func (v NullableQuarantineGoods) Get() *QuarantineGoods {
	return v.value
}

func (v *NullableQuarantineGoods) Set(val *QuarantineGoods) {
	v.value = val
	v.isSet = true
}

func (v NullableQuarantineGoods) IsSet() bool {
	return v.isSet
}

func (v *NullableQuarantineGoods) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuarantineGoods(val *QuarantineGoods) *NullableQuarantineGoods {
	return &NullableQuarantineGoods{value: val, isSet: true}
}

func (v NullableQuarantineGoods) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuarantineGoods) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
