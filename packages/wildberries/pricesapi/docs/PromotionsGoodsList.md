# PromotionsGoodsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID номенклатуры | [optional] 
**InAction** | Pointer to **bool** | Участвует в акции:   - &#x60;true&#x60; — да   - &#x60;false&#x60; — нет  | [optional] 
**Price** | Pointer to **float32** | Текущая розничная цена | [optional] 
**CurrencyCode** | Pointer to **string** | Валюта в формате ISO 4217 | [optional] 
**PlanPrice** | Pointer to **float32** | Плановая цена (цена во время акции) | [optional] 
**Discount** | Pointer to **int32** | Текущая скидка | [optional] 
**PlanDiscount** | Pointer to **int32** | Рекомендуемая скидка для участия в акции | [optional] 

## Methods

### NewPromotionsGoodsList

`func NewPromotionsGoodsList() *PromotionsGoodsList`

NewPromotionsGoodsList instantiates a new PromotionsGoodsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromotionsGoodsListWithDefaults

`func NewPromotionsGoodsListWithDefaults() *PromotionsGoodsList`

NewPromotionsGoodsListWithDefaults instantiates a new PromotionsGoodsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PromotionsGoodsList) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PromotionsGoodsList) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PromotionsGoodsList) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PromotionsGoodsList) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInAction

`func (o *PromotionsGoodsList) GetInAction() bool`

GetInAction returns the InAction field if non-nil, zero value otherwise.

### GetInActionOk

`func (o *PromotionsGoodsList) GetInActionOk() (*bool, bool)`

GetInActionOk returns a tuple with the InAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInAction

`func (o *PromotionsGoodsList) SetInAction(v bool)`

SetInAction sets InAction field to given value.

### HasInAction

`func (o *PromotionsGoodsList) HasInAction() bool`

HasInAction returns a boolean if a field has been set.

### GetPrice

`func (o *PromotionsGoodsList) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PromotionsGoodsList) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PromotionsGoodsList) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PromotionsGoodsList) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *PromotionsGoodsList) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PromotionsGoodsList) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PromotionsGoodsList) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PromotionsGoodsList) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetPlanPrice

`func (o *PromotionsGoodsList) GetPlanPrice() float32`

GetPlanPrice returns the PlanPrice field if non-nil, zero value otherwise.

### GetPlanPriceOk

`func (o *PromotionsGoodsList) GetPlanPriceOk() (*float32, bool)`

GetPlanPriceOk returns a tuple with the PlanPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanPrice

`func (o *PromotionsGoodsList) SetPlanPrice(v float32)`

SetPlanPrice sets PlanPrice field to given value.

### HasPlanPrice

`func (o *PromotionsGoodsList) HasPlanPrice() bool`

HasPlanPrice returns a boolean if a field has been set.

### GetDiscount

`func (o *PromotionsGoodsList) GetDiscount() int32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *PromotionsGoodsList) GetDiscountOk() (*int32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *PromotionsGoodsList) SetDiscount(v int32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *PromotionsGoodsList) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetPlanDiscount

`func (o *PromotionsGoodsList) GetPlanDiscount() int32`

GetPlanDiscount returns the PlanDiscount field if non-nil, zero value otherwise.

### GetPlanDiscountOk

`func (o *PromotionsGoodsList) GetPlanDiscountOk() (*int32, bool)`

GetPlanDiscountOk returns a tuple with the PlanDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanDiscount

`func (o *PromotionsGoodsList) SetPlanDiscount(v int32)`

SetPlanDiscount sets PlanDiscount field to given value.

### HasPlanDiscount

`func (o *PromotionsGoodsList) HasPlanDiscount() bool`

HasPlanDiscount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


