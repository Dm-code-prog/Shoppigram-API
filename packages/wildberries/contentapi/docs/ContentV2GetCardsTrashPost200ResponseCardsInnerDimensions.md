# ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Length** | Pointer to **int32** | Длина, см | [optional] 
**Width** | Pointer to **int32** | Ширина, см | [optional] 
**Height** | Pointer to **int32** | Высота, см | [optional] 
**IsValid** | Pointer to **bool** | Потенциальная некорректность габаритов товара: - &#x60;true&#x60; — не выявлена. &#x60;\&quot;isValid\&quot;:true&#x60; не гарантирует, что размеры указаны корректно. В отдельных случаях (например, при создании новой категории товаров) &#x60;\&quot;isValid\&quot;:true&#x60; будет возвращаться при любых значениях, кроме нулевых. - &#x60;false&#x60; — указанные габариты значительно отличаются от средних по категории (предмету). Рекомендуется перепроверить, правильно ли указаны размеры товара в упаковке **в сантиметрах**. Функциональность карточки товара, в том числе начисление логистики и хранения, при этом ограничена не будет. Логистика и хранение продолжают начисляться — по текущим габаритам. Также &#x60;\&quot;isValid\&quot;:false&#x60; возвращается при отсутствии значений или нулевом значении любой стороны.  | [optional] 

## Methods

### NewContentV2GetCardsTrashPost200ResponseCardsInnerDimensions

`func NewContentV2GetCardsTrashPost200ResponseCardsInnerDimensions() *ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions`

NewContentV2GetCardsTrashPost200ResponseCardsInnerDimensions instantiates a new ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentV2GetCardsTrashPost200ResponseCardsInnerDimensionsWithDefaults

`func NewContentV2GetCardsTrashPost200ResponseCardsInnerDimensionsWithDefaults() *ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions`

NewContentV2GetCardsTrashPost200ResponseCardsInnerDimensionsWithDefaults instantiates a new ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLength

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetWidth

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetIsValid

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


