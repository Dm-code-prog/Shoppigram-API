# ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID акции | [optional] 
**Name** | Pointer to **string** | Название акции | [optional] 
**Description** | Pointer to **string** | Описание акции | [optional] 
**Advantages** | Pointer to **[]string** | Преимущества акции | [optional] 
**StartDateTime** | Pointer to **string** | Начало акции | [optional] 
**EndDateTime** | Pointer to **string** | Конец акции | [optional] 
**InPromoActionLeftovers** | Pointer to **int32** | Количество товаров с остатками, участвующих в акции | [optional] 
**InPromoActionTotal** | Pointer to **int32** | Общее количество товаров, участвующих в акции | [optional] 
**NotInPromoActionLeftovers** | Pointer to **int32** | Количество товаров с остатками, не участвующих в акции | [optional] 
**NotInPromoActionTotal** | Pointer to **int32** | Общее количество товаров, не участвующих в акции | [optional] 
**ParticipationPercentage** | Pointer to **int32** | Уже участвующие в акции товары, %. Рассчитывается по товарам в акции и с остатком | [optional] 
**Type** | Pointer to **string** | Тип акции:   - &#x60;regular&#x60; — акция   - &#x60;auto&#x60; — автоакция  | [optional] 
**ExceptionProductsCount** | Pointer to **int32** | Количество товаров, исключенных из автоакции до её старта. Только при &#x60;\&quot;type\&quot;: \&quot;auto\&quot;&#x60;. &lt;br&gt;В момент старта акции эти товары автоматически будут без скидки  | [optional] 

## Methods

### NewApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner

`func NewApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner() *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner`

NewApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner instantiates a new ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInnerWithDefaults

`func NewApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInnerWithDefaults() *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner`

NewApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInnerWithDefaults instantiates a new ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAdvantages

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetAdvantages() []string`

GetAdvantages returns the Advantages field if non-nil, zero value otherwise.

### GetAdvantagesOk

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetAdvantagesOk() (*[]string, bool)`

GetAdvantagesOk returns a tuple with the Advantages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvantages

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) SetAdvantages(v []string)`

SetAdvantages sets Advantages field to given value.

### HasAdvantages

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) HasAdvantages() bool`

HasAdvantages returns a boolean if a field has been set.

### GetStartDateTime

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetStartDateTime() string`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetStartDateTimeOk() (*string, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) SetStartDateTime(v string)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetEndDateTime

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetEndDateTime() string`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetEndDateTimeOk() (*string, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) SetEndDateTime(v string)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetInPromoActionLeftovers

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetInPromoActionLeftovers() int32`

GetInPromoActionLeftovers returns the InPromoActionLeftovers field if non-nil, zero value otherwise.

### GetInPromoActionLeftoversOk

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetInPromoActionLeftoversOk() (*int32, bool)`

GetInPromoActionLeftoversOk returns a tuple with the InPromoActionLeftovers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInPromoActionLeftovers

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) SetInPromoActionLeftovers(v int32)`

SetInPromoActionLeftovers sets InPromoActionLeftovers field to given value.

### HasInPromoActionLeftovers

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) HasInPromoActionLeftovers() bool`

HasInPromoActionLeftovers returns a boolean if a field has been set.

### GetInPromoActionTotal

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetInPromoActionTotal() int32`

GetInPromoActionTotal returns the InPromoActionTotal field if non-nil, zero value otherwise.

### GetInPromoActionTotalOk

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetInPromoActionTotalOk() (*int32, bool)`

GetInPromoActionTotalOk returns a tuple with the InPromoActionTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInPromoActionTotal

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) SetInPromoActionTotal(v int32)`

SetInPromoActionTotal sets InPromoActionTotal field to given value.

### HasInPromoActionTotal

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) HasInPromoActionTotal() bool`

HasInPromoActionTotal returns a boolean if a field has been set.

### GetNotInPromoActionLeftovers

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetNotInPromoActionLeftovers() int32`

GetNotInPromoActionLeftovers returns the NotInPromoActionLeftovers field if non-nil, zero value otherwise.

### GetNotInPromoActionLeftoversOk

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetNotInPromoActionLeftoversOk() (*int32, bool)`

GetNotInPromoActionLeftoversOk returns a tuple with the NotInPromoActionLeftovers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotInPromoActionLeftovers

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) SetNotInPromoActionLeftovers(v int32)`

SetNotInPromoActionLeftovers sets NotInPromoActionLeftovers field to given value.

### HasNotInPromoActionLeftovers

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) HasNotInPromoActionLeftovers() bool`

HasNotInPromoActionLeftovers returns a boolean if a field has been set.

### GetNotInPromoActionTotal

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetNotInPromoActionTotal() int32`

GetNotInPromoActionTotal returns the NotInPromoActionTotal field if non-nil, zero value otherwise.

### GetNotInPromoActionTotalOk

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetNotInPromoActionTotalOk() (*int32, bool)`

GetNotInPromoActionTotalOk returns a tuple with the NotInPromoActionTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotInPromoActionTotal

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) SetNotInPromoActionTotal(v int32)`

SetNotInPromoActionTotal sets NotInPromoActionTotal field to given value.

### HasNotInPromoActionTotal

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) HasNotInPromoActionTotal() bool`

HasNotInPromoActionTotal returns a boolean if a field has been set.

### GetParticipationPercentage

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetParticipationPercentage() int32`

GetParticipationPercentage returns the ParticipationPercentage field if non-nil, zero value otherwise.

### GetParticipationPercentageOk

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetParticipationPercentageOk() (*int32, bool)`

GetParticipationPercentageOk returns a tuple with the ParticipationPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipationPercentage

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) SetParticipationPercentage(v int32)`

SetParticipationPercentage sets ParticipationPercentage field to given value.

### HasParticipationPercentage

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) HasParticipationPercentage() bool`

HasParticipationPercentage returns a boolean if a field has been set.

### GetType

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExceptionProductsCount

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetExceptionProductsCount() int32`

GetExceptionProductsCount returns the ExceptionProductsCount field if non-nil, zero value otherwise.

### GetExceptionProductsCountOk

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) GetExceptionProductsCountOk() (*int32, bool)`

GetExceptionProductsCountOk returns a tuple with the ExceptionProductsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionProductsCount

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) SetExceptionProductsCount(v int32)`

SetExceptionProductsCount sets ExceptionProductsCount field to given value.

### HasExceptionProductsCount

`func (o *ApiV1CalendarPromotionsDetailsGet200ResponseDataPromotionsInner) HasExceptionProductsCount() bool`

HasExceptionProductsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


