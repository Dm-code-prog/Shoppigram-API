# SizeGood

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NmID** | Pointer to **int32** | Артикул Wildberries | [optional] 
**SizeID** | Pointer to **int32** | ID размера. Можно получить с помощью метода [Получение списка товаров по артикулам](./#tag/Spiski-tovarov/paths/~1api~1v2~1list~1goods~1filter/get), поле &#x60;sizeID&#x60;. В методах контента это поле &#x60;chrtID&#x60; | [optional] 
**VendorCode** | Pointer to **string** | Артикул продавца | [optional] 
**Price** | Pointer to **int32** | Цена | [optional] 
**CurrencyIsoCode4217** | Pointer to **string** | Валюта, по стандарту ISO 4217 | [optional] 
**DiscountedPrice** | Pointer to **float32** | Цена со скидкой | [optional] 
**Discount** | Pointer to **int32** | Скидка, % | [optional] 
**TechSizeName** | Pointer to **string** | Размер товара | [optional] 
**EditableSizePrice** | Pointer to **bool** | Можно ли устанавливать цены отдельно для разных размеров: &#x60;true&#x60; — можно, &#x60;false&#x60; — нельзя. Эта возможность зависит от категории товара  | [optional] 

## Methods

### NewSizeGood

`func NewSizeGood() *SizeGood`

NewSizeGood instantiates a new SizeGood object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSizeGoodWithDefaults

`func NewSizeGoodWithDefaults() *SizeGood`

NewSizeGoodWithDefaults instantiates a new SizeGood object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNmID

`func (o *SizeGood) GetNmID() int32`

GetNmID returns the NmID field if non-nil, zero value otherwise.

### GetNmIDOk

`func (o *SizeGood) GetNmIDOk() (*int32, bool)`

GetNmIDOk returns a tuple with the NmID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmID

`func (o *SizeGood) SetNmID(v int32)`

SetNmID sets NmID field to given value.

### HasNmID

`func (o *SizeGood) HasNmID() bool`

HasNmID returns a boolean if a field has been set.

### GetSizeID

`func (o *SizeGood) GetSizeID() int32`

GetSizeID returns the SizeID field if non-nil, zero value otherwise.

### GetSizeIDOk

`func (o *SizeGood) GetSizeIDOk() (*int32, bool)`

GetSizeIDOk returns a tuple with the SizeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeID

`func (o *SizeGood) SetSizeID(v int32)`

SetSizeID sets SizeID field to given value.

### HasSizeID

`func (o *SizeGood) HasSizeID() bool`

HasSizeID returns a boolean if a field has been set.

### GetVendorCode

`func (o *SizeGood) GetVendorCode() string`

GetVendorCode returns the VendorCode field if non-nil, zero value otherwise.

### GetVendorCodeOk

`func (o *SizeGood) GetVendorCodeOk() (*string, bool)`

GetVendorCodeOk returns a tuple with the VendorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCode

`func (o *SizeGood) SetVendorCode(v string)`

SetVendorCode sets VendorCode field to given value.

### HasVendorCode

`func (o *SizeGood) HasVendorCode() bool`

HasVendorCode returns a boolean if a field has been set.

### GetPrice

`func (o *SizeGood) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SizeGood) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SizeGood) SetPrice(v int32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *SizeGood) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCurrencyIsoCode4217

`func (o *SizeGood) GetCurrencyIsoCode4217() string`

GetCurrencyIsoCode4217 returns the CurrencyIsoCode4217 field if non-nil, zero value otherwise.

### GetCurrencyIsoCode4217Ok

`func (o *SizeGood) GetCurrencyIsoCode4217Ok() (*string, bool)`

GetCurrencyIsoCode4217Ok returns a tuple with the CurrencyIsoCode4217 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyIsoCode4217

`func (o *SizeGood) SetCurrencyIsoCode4217(v string)`

SetCurrencyIsoCode4217 sets CurrencyIsoCode4217 field to given value.

### HasCurrencyIsoCode4217

`func (o *SizeGood) HasCurrencyIsoCode4217() bool`

HasCurrencyIsoCode4217 returns a boolean if a field has been set.

### GetDiscountedPrice

`func (o *SizeGood) GetDiscountedPrice() float32`

GetDiscountedPrice returns the DiscountedPrice field if non-nil, zero value otherwise.

### GetDiscountedPriceOk

`func (o *SizeGood) GetDiscountedPriceOk() (*float32, bool)`

GetDiscountedPriceOk returns a tuple with the DiscountedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountedPrice

`func (o *SizeGood) SetDiscountedPrice(v float32)`

SetDiscountedPrice sets DiscountedPrice field to given value.

### HasDiscountedPrice

`func (o *SizeGood) HasDiscountedPrice() bool`

HasDiscountedPrice returns a boolean if a field has been set.

### GetDiscount

`func (o *SizeGood) GetDiscount() int32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *SizeGood) GetDiscountOk() (*int32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *SizeGood) SetDiscount(v int32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *SizeGood) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetTechSizeName

`func (o *SizeGood) GetTechSizeName() string`

GetTechSizeName returns the TechSizeName field if non-nil, zero value otherwise.

### GetTechSizeNameOk

`func (o *SizeGood) GetTechSizeNameOk() (*string, bool)`

GetTechSizeNameOk returns a tuple with the TechSizeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechSizeName

`func (o *SizeGood) SetTechSizeName(v string)`

SetTechSizeName sets TechSizeName field to given value.

### HasTechSizeName

`func (o *SizeGood) HasTechSizeName() bool`

HasTechSizeName returns a boolean if a field has been set.

### GetEditableSizePrice

`func (o *SizeGood) GetEditableSizePrice() bool`

GetEditableSizePrice returns the EditableSizePrice field if non-nil, zero value otherwise.

### GetEditableSizePriceOk

`func (o *SizeGood) GetEditableSizePriceOk() (*bool, bool)`

GetEditableSizePriceOk returns a tuple with the EditableSizePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableSizePrice

`func (o *SizeGood) SetEditableSizePrice(v bool)`

SetEditableSizePrice sets EditableSizePrice field to given value.

### HasEditableSizePrice

`func (o *SizeGood) HasEditableSizePrice() bool`

HasEditableSizePrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


