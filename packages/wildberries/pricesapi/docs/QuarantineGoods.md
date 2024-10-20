# QuarantineGoods

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NmID** | Pointer to **int32** | Артикул WB | [optional] 
**SizeID** | Pointer to **int32** | Не используется | [optional] 
**TechSizeName** | Pointer to **string** | Не используется | [optional] 
**CurrencyIsoCode4217** | Pointer to **string** | Валюта по стандарту ISO 4217 | [optional] 
**NewPrice** | Pointer to **float32** | Новая цена продавца до скидки | [optional] 
**OldPrice** | Pointer to **float32** | Текущая цена продавца до скидки | [optional] 
**NewDiscount** | Pointer to **int32** | Новая скидка продавца, % | [optional] 
**OldDiscount** | Pointer to **int32** | Текущая скидка продавца, % | [optional] 
**PriceDiff** | Pointer to **float32** | Разница: &#x60;newPrice&#x60; * (1 - &#x60;newDiscount&#x60; / 100) - &#x60;oldPrice&#x60; * (1 - &#x60;oldDiscount&#x60; / 100) | [optional] 

## Methods

### NewQuarantineGoods

`func NewQuarantineGoods() *QuarantineGoods`

NewQuarantineGoods instantiates a new QuarantineGoods object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuarantineGoodsWithDefaults

`func NewQuarantineGoodsWithDefaults() *QuarantineGoods`

NewQuarantineGoodsWithDefaults instantiates a new QuarantineGoods object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNmID

`func (o *QuarantineGoods) GetNmID() int32`

GetNmID returns the NmID field if non-nil, zero value otherwise.

### GetNmIDOk

`func (o *QuarantineGoods) GetNmIDOk() (*int32, bool)`

GetNmIDOk returns a tuple with the NmID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmID

`func (o *QuarantineGoods) SetNmID(v int32)`

SetNmID sets NmID field to given value.

### HasNmID

`func (o *QuarantineGoods) HasNmID() bool`

HasNmID returns a boolean if a field has been set.

### GetSizeID

`func (o *QuarantineGoods) GetSizeID() int32`

GetSizeID returns the SizeID field if non-nil, zero value otherwise.

### GetSizeIDOk

`func (o *QuarantineGoods) GetSizeIDOk() (*int32, bool)`

GetSizeIDOk returns a tuple with the SizeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeID

`func (o *QuarantineGoods) SetSizeID(v int32)`

SetSizeID sets SizeID field to given value.

### HasSizeID

`func (o *QuarantineGoods) HasSizeID() bool`

HasSizeID returns a boolean if a field has been set.

### GetTechSizeName

`func (o *QuarantineGoods) GetTechSizeName() string`

GetTechSizeName returns the TechSizeName field if non-nil, zero value otherwise.

### GetTechSizeNameOk

`func (o *QuarantineGoods) GetTechSizeNameOk() (*string, bool)`

GetTechSizeNameOk returns a tuple with the TechSizeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechSizeName

`func (o *QuarantineGoods) SetTechSizeName(v string)`

SetTechSizeName sets TechSizeName field to given value.

### HasTechSizeName

`func (o *QuarantineGoods) HasTechSizeName() bool`

HasTechSizeName returns a boolean if a field has been set.

### GetCurrencyIsoCode4217

`func (o *QuarantineGoods) GetCurrencyIsoCode4217() string`

GetCurrencyIsoCode4217 returns the CurrencyIsoCode4217 field if non-nil, zero value otherwise.

### GetCurrencyIsoCode4217Ok

`func (o *QuarantineGoods) GetCurrencyIsoCode4217Ok() (*string, bool)`

GetCurrencyIsoCode4217Ok returns a tuple with the CurrencyIsoCode4217 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyIsoCode4217

`func (o *QuarantineGoods) SetCurrencyIsoCode4217(v string)`

SetCurrencyIsoCode4217 sets CurrencyIsoCode4217 field to given value.

### HasCurrencyIsoCode4217

`func (o *QuarantineGoods) HasCurrencyIsoCode4217() bool`

HasCurrencyIsoCode4217 returns a boolean if a field has been set.

### GetNewPrice

`func (o *QuarantineGoods) GetNewPrice() float32`

GetNewPrice returns the NewPrice field if non-nil, zero value otherwise.

### GetNewPriceOk

`func (o *QuarantineGoods) GetNewPriceOk() (*float32, bool)`

GetNewPriceOk returns a tuple with the NewPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPrice

`func (o *QuarantineGoods) SetNewPrice(v float32)`

SetNewPrice sets NewPrice field to given value.

### HasNewPrice

`func (o *QuarantineGoods) HasNewPrice() bool`

HasNewPrice returns a boolean if a field has been set.

### GetOldPrice

`func (o *QuarantineGoods) GetOldPrice() float32`

GetOldPrice returns the OldPrice field if non-nil, zero value otherwise.

### GetOldPriceOk

`func (o *QuarantineGoods) GetOldPriceOk() (*float32, bool)`

GetOldPriceOk returns a tuple with the OldPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPrice

`func (o *QuarantineGoods) SetOldPrice(v float32)`

SetOldPrice sets OldPrice field to given value.

### HasOldPrice

`func (o *QuarantineGoods) HasOldPrice() bool`

HasOldPrice returns a boolean if a field has been set.

### GetNewDiscount

`func (o *QuarantineGoods) GetNewDiscount() int32`

GetNewDiscount returns the NewDiscount field if non-nil, zero value otherwise.

### GetNewDiscountOk

`func (o *QuarantineGoods) GetNewDiscountOk() (*int32, bool)`

GetNewDiscountOk returns a tuple with the NewDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDiscount

`func (o *QuarantineGoods) SetNewDiscount(v int32)`

SetNewDiscount sets NewDiscount field to given value.

### HasNewDiscount

`func (o *QuarantineGoods) HasNewDiscount() bool`

HasNewDiscount returns a boolean if a field has been set.

### GetOldDiscount

`func (o *QuarantineGoods) GetOldDiscount() int32`

GetOldDiscount returns the OldDiscount field if non-nil, zero value otherwise.

### GetOldDiscountOk

`func (o *QuarantineGoods) GetOldDiscountOk() (*int32, bool)`

GetOldDiscountOk returns a tuple with the OldDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldDiscount

`func (o *QuarantineGoods) SetOldDiscount(v int32)`

SetOldDiscount sets OldDiscount field to given value.

### HasOldDiscount

`func (o *QuarantineGoods) HasOldDiscount() bool`

HasOldDiscount returns a boolean if a field has been set.

### GetPriceDiff

`func (o *QuarantineGoods) GetPriceDiff() float32`

GetPriceDiff returns the PriceDiff field if non-nil, zero value otherwise.

### GetPriceDiffOk

`func (o *QuarantineGoods) GetPriceDiffOk() (*float32, bool)`

GetPriceDiffOk returns a tuple with the PriceDiff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDiff

`func (o *QuarantineGoods) SetPriceDiff(v float32)`

SetPriceDiff sets PriceDiff field to given value.

### HasPriceDiff

`func (o *QuarantineGoods) HasPriceDiff() bool`

HasPriceDiff returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


