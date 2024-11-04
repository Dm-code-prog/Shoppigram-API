# GoodBufferHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NmID** | Pointer to **int32** | Артикул Wildberries | [optional] 
**VendorCode** | Pointer to **string** | Артикул продавца | [optional] 
**SizeID** | Pointer to **int32** | ID размера. В методах контента это поле &#x60;chrtID&#x60; | [optional] 
**TechSizeName** | Pointer to **string** | Размер | [optional] 
**Price** | Pointer to **int32** | Цена | [optional] 
**CurrencyIsoCode4217** | Pointer to **string** | Валюта, по стандарту ISO 4217 | [optional] 
**Discount** | Pointer to **int32** | Скидка, % | [optional] 
**Status** | Pointer to **int32** | Статус товара: &#x60;1&#x60; — в обработке  | [optional] 
**ErrorText** | Pointer to **string** | Текст ошибки | [optional] 

## Methods

### NewGoodBufferHistory

`func NewGoodBufferHistory() *GoodBufferHistory`

NewGoodBufferHistory instantiates a new GoodBufferHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoodBufferHistoryWithDefaults

`func NewGoodBufferHistoryWithDefaults() *GoodBufferHistory`

NewGoodBufferHistoryWithDefaults instantiates a new GoodBufferHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNmID

`func (o *GoodBufferHistory) GetNmID() int32`

GetNmID returns the NmID field if non-nil, zero value otherwise.

### GetNmIDOk

`func (o *GoodBufferHistory) GetNmIDOk() (*int32, bool)`

GetNmIDOk returns a tuple with the NmID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmID

`func (o *GoodBufferHistory) SetNmID(v int32)`

SetNmID sets NmID field to given value.

### HasNmID

`func (o *GoodBufferHistory) HasNmID() bool`

HasNmID returns a boolean if a field has been set.

### GetVendorCode

`func (o *GoodBufferHistory) GetVendorCode() string`

GetVendorCode returns the VendorCode field if non-nil, zero value otherwise.

### GetVendorCodeOk

`func (o *GoodBufferHistory) GetVendorCodeOk() (*string, bool)`

GetVendorCodeOk returns a tuple with the VendorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCode

`func (o *GoodBufferHistory) SetVendorCode(v string)`

SetVendorCode sets VendorCode field to given value.

### HasVendorCode

`func (o *GoodBufferHistory) HasVendorCode() bool`

HasVendorCode returns a boolean if a field has been set.

### GetSizeID

`func (o *GoodBufferHistory) GetSizeID() int32`

GetSizeID returns the SizeID field if non-nil, zero value otherwise.

### GetSizeIDOk

`func (o *GoodBufferHistory) GetSizeIDOk() (*int32, bool)`

GetSizeIDOk returns a tuple with the SizeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeID

`func (o *GoodBufferHistory) SetSizeID(v int32)`

SetSizeID sets SizeID field to given value.

### HasSizeID

`func (o *GoodBufferHistory) HasSizeID() bool`

HasSizeID returns a boolean if a field has been set.

### GetTechSizeName

`func (o *GoodBufferHistory) GetTechSizeName() string`

GetTechSizeName returns the TechSizeName field if non-nil, zero value otherwise.

### GetTechSizeNameOk

`func (o *GoodBufferHistory) GetTechSizeNameOk() (*string, bool)`

GetTechSizeNameOk returns a tuple with the TechSizeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechSizeName

`func (o *GoodBufferHistory) SetTechSizeName(v string)`

SetTechSizeName sets TechSizeName field to given value.

### HasTechSizeName

`func (o *GoodBufferHistory) HasTechSizeName() bool`

HasTechSizeName returns a boolean if a field has been set.

### GetPrice

`func (o *GoodBufferHistory) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GoodBufferHistory) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GoodBufferHistory) SetPrice(v int32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GoodBufferHistory) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCurrencyIsoCode4217

`func (o *GoodBufferHistory) GetCurrencyIsoCode4217() string`

GetCurrencyIsoCode4217 returns the CurrencyIsoCode4217 field if non-nil, zero value otherwise.

### GetCurrencyIsoCode4217Ok

`func (o *GoodBufferHistory) GetCurrencyIsoCode4217Ok() (*string, bool)`

GetCurrencyIsoCode4217Ok returns a tuple with the CurrencyIsoCode4217 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyIsoCode4217

`func (o *GoodBufferHistory) SetCurrencyIsoCode4217(v string)`

SetCurrencyIsoCode4217 sets CurrencyIsoCode4217 field to given value.

### HasCurrencyIsoCode4217

`func (o *GoodBufferHistory) HasCurrencyIsoCode4217() bool`

HasCurrencyIsoCode4217 returns a boolean if a field has been set.

### GetDiscount

`func (o *GoodBufferHistory) GetDiscount() int32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *GoodBufferHistory) GetDiscountOk() (*int32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *GoodBufferHistory) SetDiscount(v int32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *GoodBufferHistory) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetStatus

`func (o *GoodBufferHistory) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GoodBufferHistory) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GoodBufferHistory) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GoodBufferHistory) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorText

`func (o *GoodBufferHistory) GetErrorText() string`

GetErrorText returns the ErrorText field if non-nil, zero value otherwise.

### GetErrorTextOk

`func (o *GoodBufferHistory) GetErrorTextOk() (*string, bool)`

GetErrorTextOk returns a tuple with the ErrorText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorText

`func (o *GoodBufferHistory) SetErrorText(v string)`

SetErrorText sets ErrorText field to given value.

### HasErrorText

`func (o *GoodBufferHistory) HasErrorText() bool`

HasErrorText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


