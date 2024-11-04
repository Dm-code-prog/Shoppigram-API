# Good

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NmID** | **int32** | Артикул Wildberries | 
**Price** | Pointer to **int32** | Цена. Валюту можно получить с помощью метода [Получение списка товаров по артикулам](./#tag/Spiski-tovarov/paths/~1api~1v2~1list~1goods~1filter/get), поле &#x60;currencyIsoCode4217&#x60; | [optional] 
**Discount** | Pointer to **int32** | Скидка, % | [optional] 

## Methods

### NewGood

`func NewGood(nmID int32, ) *Good`

NewGood instantiates a new Good object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoodWithDefaults

`func NewGoodWithDefaults() *Good`

NewGoodWithDefaults instantiates a new Good object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNmID

`func (o *Good) GetNmID() int32`

GetNmID returns the NmID field if non-nil, zero value otherwise.

### GetNmIDOk

`func (o *Good) GetNmIDOk() (*int32, bool)`

GetNmIDOk returns a tuple with the NmID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmID

`func (o *Good) SetNmID(v int32)`

SetNmID sets NmID field to given value.


### GetPrice

`func (o *Good) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Good) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Good) SetPrice(v int32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Good) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetDiscount

`func (o *Good) GetDiscount() int32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *Good) GetDiscountOk() (*int32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *Good) SetDiscount(v int32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *Good) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


