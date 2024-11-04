# GoodHistory

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
**Status** | Pointer to **int32** | Статус товара:    * &#x60;2&#x60; — товар без ошибок, цена и/или скидка обновилась   * &#x60;3&#x60; — есть ошибки, данные не обновились  | [optional] 
**ErrorText** | Pointer to **string** | Текст ошибки  &lt;blockquote class&#x3D;\&quot;spoiler\&quot;&gt;   &lt;p class&#x3D;\&quot;descr\&quot;&gt;Ошибка &lt;code&gt;The product is in quarantine&lt;/code&gt; возникает, если новая цена со скидкой хотя бы в 3 раза меньше старой. Вы можете изменить цену или скидку с помощью API либо вывести товар из карантина &lt;a href&#x3D;\&quot;https://seller.wildberries.ru/discount-and-prices/quarantine\&quot;&gt;в личном кабинете&lt;/a&gt;&lt;/p&gt; &lt;/blockquote&gt;  | [optional] 

## Methods

### NewGoodHistory

`func NewGoodHistory() *GoodHistory`

NewGoodHistory instantiates a new GoodHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoodHistoryWithDefaults

`func NewGoodHistoryWithDefaults() *GoodHistory`

NewGoodHistoryWithDefaults instantiates a new GoodHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNmID

`func (o *GoodHistory) GetNmID() int32`

GetNmID returns the NmID field if non-nil, zero value otherwise.

### GetNmIDOk

`func (o *GoodHistory) GetNmIDOk() (*int32, bool)`

GetNmIDOk returns a tuple with the NmID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmID

`func (o *GoodHistory) SetNmID(v int32)`

SetNmID sets NmID field to given value.

### HasNmID

`func (o *GoodHistory) HasNmID() bool`

HasNmID returns a boolean if a field has been set.

### GetVendorCode

`func (o *GoodHistory) GetVendorCode() string`

GetVendorCode returns the VendorCode field if non-nil, zero value otherwise.

### GetVendorCodeOk

`func (o *GoodHistory) GetVendorCodeOk() (*string, bool)`

GetVendorCodeOk returns a tuple with the VendorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCode

`func (o *GoodHistory) SetVendorCode(v string)`

SetVendorCode sets VendorCode field to given value.

### HasVendorCode

`func (o *GoodHistory) HasVendorCode() bool`

HasVendorCode returns a boolean if a field has been set.

### GetSizeID

`func (o *GoodHistory) GetSizeID() int32`

GetSizeID returns the SizeID field if non-nil, zero value otherwise.

### GetSizeIDOk

`func (o *GoodHistory) GetSizeIDOk() (*int32, bool)`

GetSizeIDOk returns a tuple with the SizeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeID

`func (o *GoodHistory) SetSizeID(v int32)`

SetSizeID sets SizeID field to given value.

### HasSizeID

`func (o *GoodHistory) HasSizeID() bool`

HasSizeID returns a boolean if a field has been set.

### GetTechSizeName

`func (o *GoodHistory) GetTechSizeName() string`

GetTechSizeName returns the TechSizeName field if non-nil, zero value otherwise.

### GetTechSizeNameOk

`func (o *GoodHistory) GetTechSizeNameOk() (*string, bool)`

GetTechSizeNameOk returns a tuple with the TechSizeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechSizeName

`func (o *GoodHistory) SetTechSizeName(v string)`

SetTechSizeName sets TechSizeName field to given value.

### HasTechSizeName

`func (o *GoodHistory) HasTechSizeName() bool`

HasTechSizeName returns a boolean if a field has been set.

### GetPrice

`func (o *GoodHistory) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GoodHistory) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GoodHistory) SetPrice(v int32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GoodHistory) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCurrencyIsoCode4217

`func (o *GoodHistory) GetCurrencyIsoCode4217() string`

GetCurrencyIsoCode4217 returns the CurrencyIsoCode4217 field if non-nil, zero value otherwise.

### GetCurrencyIsoCode4217Ok

`func (o *GoodHistory) GetCurrencyIsoCode4217Ok() (*string, bool)`

GetCurrencyIsoCode4217Ok returns a tuple with the CurrencyIsoCode4217 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyIsoCode4217

`func (o *GoodHistory) SetCurrencyIsoCode4217(v string)`

SetCurrencyIsoCode4217 sets CurrencyIsoCode4217 field to given value.

### HasCurrencyIsoCode4217

`func (o *GoodHistory) HasCurrencyIsoCode4217() bool`

HasCurrencyIsoCode4217 returns a boolean if a field has been set.

### GetDiscount

`func (o *GoodHistory) GetDiscount() int32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *GoodHistory) GetDiscountOk() (*int32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *GoodHistory) SetDiscount(v int32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *GoodHistory) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetStatus

`func (o *GoodHistory) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GoodHistory) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GoodHistory) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GoodHistory) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorText

`func (o *GoodHistory) GetErrorText() string`

GetErrorText returns the ErrorText field if non-nil, zero value otherwise.

### GetErrorTextOk

`func (o *GoodHistory) GetErrorTextOk() (*string, bool)`

GetErrorTextOk returns a tuple with the ErrorText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorText

`func (o *GoodHistory) SetErrorText(v string)`

SetErrorText sets ErrorText field to given value.

### HasErrorText

`func (o *GoodHistory) HasErrorText() bool`

HasErrorText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


