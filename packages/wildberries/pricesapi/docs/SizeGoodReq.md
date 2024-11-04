# SizeGoodReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NmID** | **int32** | Артикул Wildberries | 
**SizeID** | **int32** | ID размера. Можно получить с помощью метода [Получение списка товаров по артикулам](./#tag/Spiski-tovarov/paths/~1api~1v2~1list~1goods~1filter/get), поле &#x60;sizeID&#x60;. В методах контента это поле &#x60;chrtID&#x60; | 
**Price** | **int32** | Цена. Валюту можно получить с помощью метода [Получение списка товаров по артикулам](./#tag/Spiski-tovarov/paths/~1api~1v2~1list~1goods~1filter/get), поле &#x60;currencyIsoCode4217&#x60; | 

## Methods

### NewSizeGoodReq

`func NewSizeGoodReq(nmID int32, sizeID int32, price int32, ) *SizeGoodReq`

NewSizeGoodReq instantiates a new SizeGoodReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSizeGoodReqWithDefaults

`func NewSizeGoodReqWithDefaults() *SizeGoodReq`

NewSizeGoodReqWithDefaults instantiates a new SizeGoodReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNmID

`func (o *SizeGoodReq) GetNmID() int32`

GetNmID returns the NmID field if non-nil, zero value otherwise.

### GetNmIDOk

`func (o *SizeGoodReq) GetNmIDOk() (*int32, bool)`

GetNmIDOk returns a tuple with the NmID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmID

`func (o *SizeGoodReq) SetNmID(v int32)`

SetNmID sets NmID field to given value.


### GetSizeID

`func (o *SizeGoodReq) GetSizeID() int32`

GetSizeID returns the SizeID field if non-nil, zero value otherwise.

### GetSizeIDOk

`func (o *SizeGoodReq) GetSizeIDOk() (*int32, bool)`

GetSizeIDOk returns a tuple with the SizeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeID

`func (o *SizeGoodReq) SetSizeID(v int32)`

SetSizeID sets SizeID field to given value.


### GetPrice

`func (o *SizeGoodReq) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SizeGoodReq) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SizeGoodReq) SetPrice(v int32)`

SetPrice sets Price field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


