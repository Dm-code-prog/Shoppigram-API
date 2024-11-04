# GoodsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NmID** | Pointer to **int32** | Артикул Wildberries | [optional] 
**VendorCode** | Pointer to **string** | Артикул продавца | [optional] 
**Sizes** | Pointer to [**[]GoodsListSizesInner**](GoodsListSizesInner.md) | Размер | [optional] 
**CurrencyIsoCode4217** | Pointer to **string** | Валюта, по стандарту ISO 4217 | [optional] 
**Discount** | Pointer to **int32** | Скидка, % | [optional] 
**EditableSizePrice** | Pointer to **bool** | Можно ли устанавливать цены отдельно для разных размеров: &#x60;true&#x60; — можно, &#x60;false&#x60; — нельзя. Эта возможность зависит от категории товара  | [optional] 

## Methods

### NewGoodsList

`func NewGoodsList() *GoodsList`

NewGoodsList instantiates a new GoodsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoodsListWithDefaults

`func NewGoodsListWithDefaults() *GoodsList`

NewGoodsListWithDefaults instantiates a new GoodsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNmID

`func (o *GoodsList) GetNmID() int32`

GetNmID returns the NmID field if non-nil, zero value otherwise.

### GetNmIDOk

`func (o *GoodsList) GetNmIDOk() (*int32, bool)`

GetNmIDOk returns a tuple with the NmID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmID

`func (o *GoodsList) SetNmID(v int32)`

SetNmID sets NmID field to given value.

### HasNmID

`func (o *GoodsList) HasNmID() bool`

HasNmID returns a boolean if a field has been set.

### GetVendorCode

`func (o *GoodsList) GetVendorCode() string`

GetVendorCode returns the VendorCode field if non-nil, zero value otherwise.

### GetVendorCodeOk

`func (o *GoodsList) GetVendorCodeOk() (*string, bool)`

GetVendorCodeOk returns a tuple with the VendorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCode

`func (o *GoodsList) SetVendorCode(v string)`

SetVendorCode sets VendorCode field to given value.

### HasVendorCode

`func (o *GoodsList) HasVendorCode() bool`

HasVendorCode returns a boolean if a field has been set.

### GetSizes

`func (o *GoodsList) GetSizes() []GoodsListSizesInner`

GetSizes returns the Sizes field if non-nil, zero value otherwise.

### GetSizesOk

`func (o *GoodsList) GetSizesOk() (*[]GoodsListSizesInner, bool)`

GetSizesOk returns a tuple with the Sizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizes

`func (o *GoodsList) SetSizes(v []GoodsListSizesInner)`

SetSizes sets Sizes field to given value.

### HasSizes

`func (o *GoodsList) HasSizes() bool`

HasSizes returns a boolean if a field has been set.

### GetCurrencyIsoCode4217

`func (o *GoodsList) GetCurrencyIsoCode4217() string`

GetCurrencyIsoCode4217 returns the CurrencyIsoCode4217 field if non-nil, zero value otherwise.

### GetCurrencyIsoCode4217Ok

`func (o *GoodsList) GetCurrencyIsoCode4217Ok() (*string, bool)`

GetCurrencyIsoCode4217Ok returns a tuple with the CurrencyIsoCode4217 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyIsoCode4217

`func (o *GoodsList) SetCurrencyIsoCode4217(v string)`

SetCurrencyIsoCode4217 sets CurrencyIsoCode4217 field to given value.

### HasCurrencyIsoCode4217

`func (o *GoodsList) HasCurrencyIsoCode4217() bool`

HasCurrencyIsoCode4217 returns a boolean if a field has been set.

### GetDiscount

`func (o *GoodsList) GetDiscount() int32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *GoodsList) GetDiscountOk() (*int32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *GoodsList) SetDiscount(v int32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *GoodsList) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetEditableSizePrice

`func (o *GoodsList) GetEditableSizePrice() bool`

GetEditableSizePrice returns the EditableSizePrice field if non-nil, zero value otherwise.

### GetEditableSizePriceOk

`func (o *GoodsList) GetEditableSizePriceOk() (*bool, bool)`

GetEditableSizePriceOk returns a tuple with the EditableSizePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableSizePrice

`func (o *GoodsList) SetEditableSizePrice(v bool)`

SetEditableSizePrice sets EditableSizePrice field to given value.

### HasEditableSizePrice

`func (o *GoodsList) HasEditableSizePrice() bool`

HasEditableSizePrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


