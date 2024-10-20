# ApiV1CalendarPromotionsUploadPostRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PromotionID** | Pointer to **int32** | ID акции | [optional] 
**UploadNow** | Pointer to **bool** | Установить скидку:   - &#x60;true&#x60; — сейчас   - &#x60;false&#x60; — в момент старта акции  | [optional] 
**Nomenclatures** | Pointer to **[]int32** | ID номенклатур, которые можно добавить в акцию | [optional] 

## Methods

### NewApiV1CalendarPromotionsUploadPostRequestData

`func NewApiV1CalendarPromotionsUploadPostRequestData() *ApiV1CalendarPromotionsUploadPostRequestData`

NewApiV1CalendarPromotionsUploadPostRequestData instantiates a new ApiV1CalendarPromotionsUploadPostRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1CalendarPromotionsUploadPostRequestDataWithDefaults

`func NewApiV1CalendarPromotionsUploadPostRequestDataWithDefaults() *ApiV1CalendarPromotionsUploadPostRequestData`

NewApiV1CalendarPromotionsUploadPostRequestDataWithDefaults instantiates a new ApiV1CalendarPromotionsUploadPostRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPromotionID

`func (o *ApiV1CalendarPromotionsUploadPostRequestData) GetPromotionID() int32`

GetPromotionID returns the PromotionID field if non-nil, zero value otherwise.

### GetPromotionIDOk

`func (o *ApiV1CalendarPromotionsUploadPostRequestData) GetPromotionIDOk() (*int32, bool)`

GetPromotionIDOk returns a tuple with the PromotionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionID

`func (o *ApiV1CalendarPromotionsUploadPostRequestData) SetPromotionID(v int32)`

SetPromotionID sets PromotionID field to given value.

### HasPromotionID

`func (o *ApiV1CalendarPromotionsUploadPostRequestData) HasPromotionID() bool`

HasPromotionID returns a boolean if a field has been set.

### GetUploadNow

`func (o *ApiV1CalendarPromotionsUploadPostRequestData) GetUploadNow() bool`

GetUploadNow returns the UploadNow field if non-nil, zero value otherwise.

### GetUploadNowOk

`func (o *ApiV1CalendarPromotionsUploadPostRequestData) GetUploadNowOk() (*bool, bool)`

GetUploadNowOk returns a tuple with the UploadNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadNow

`func (o *ApiV1CalendarPromotionsUploadPostRequestData) SetUploadNow(v bool)`

SetUploadNow sets UploadNow field to given value.

### HasUploadNow

`func (o *ApiV1CalendarPromotionsUploadPostRequestData) HasUploadNow() bool`

HasUploadNow returns a boolean if a field has been set.

### GetNomenclatures

`func (o *ApiV1CalendarPromotionsUploadPostRequestData) GetNomenclatures() []int32`

GetNomenclatures returns the Nomenclatures field if non-nil, zero value otherwise.

### GetNomenclaturesOk

`func (o *ApiV1CalendarPromotionsUploadPostRequestData) GetNomenclaturesOk() (*[]int32, bool)`

GetNomenclaturesOk returns a tuple with the Nomenclatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomenclatures

`func (o *ApiV1CalendarPromotionsUploadPostRequestData) SetNomenclatures(v []int32)`

SetNomenclatures sets Nomenclatures field to given value.

### HasNomenclatures

`func (o *ApiV1CalendarPromotionsUploadPostRequestData) HasNomenclatures() bool`

HasNomenclatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


