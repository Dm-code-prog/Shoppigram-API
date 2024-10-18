# ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CharcID** | Pointer to **int32** | Идентификатор характеристики | [optional] 
**SubjectName** | Pointer to **string** | Название предмета | [optional] 
**SubjectID** | Pointer to **int32** | Идентификатор предмета | [optional] 
**Name** | Pointer to **string** | Название характеристики | [optional] 
**Required** | Pointer to **bool** | true - характеристику необходимо обязательно указать в КТ. false - характеристику не обязательно указывать | [optional] 
**UnitName** | Pointer to **string** | Единица измерения | [optional] 
**MaxCount** | Pointer to **int32** | Максимальное кол-во значений, которое можно присвоить данной характеристике. Если 0, то нет ограничения.  | [optional] 
**Popular** | Pointer to **bool** | Характеристика популярна у пользователей (true - да, false - нет) | [optional] 
**CharcType** | Pointer to **int32** | Тип характеристики (1 и 0 - строка или массив строк; 4 - число или массив чисел) | [optional] 

## Methods

### NewContentV2ObjectCharcsSubjectIdGet200ResponseDataInner

`func NewContentV2ObjectCharcsSubjectIdGet200ResponseDataInner() *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner`

NewContentV2ObjectCharcsSubjectIdGet200ResponseDataInner instantiates a new ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentV2ObjectCharcsSubjectIdGet200ResponseDataInnerWithDefaults

`func NewContentV2ObjectCharcsSubjectIdGet200ResponseDataInnerWithDefaults() *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner`

NewContentV2ObjectCharcsSubjectIdGet200ResponseDataInnerWithDefaults instantiates a new ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCharcID

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) GetCharcID() int32`

GetCharcID returns the CharcID field if non-nil, zero value otherwise.

### GetCharcIDOk

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) GetCharcIDOk() (*int32, bool)`

GetCharcIDOk returns a tuple with the CharcID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharcID

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) SetCharcID(v int32)`

SetCharcID sets CharcID field to given value.

### HasCharcID

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) HasCharcID() bool`

HasCharcID returns a boolean if a field has been set.

### GetSubjectName

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) GetSubjectName() string`

GetSubjectName returns the SubjectName field if non-nil, zero value otherwise.

### GetSubjectNameOk

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) GetSubjectNameOk() (*string, bool)`

GetSubjectNameOk returns a tuple with the SubjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectName

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) SetSubjectName(v string)`

SetSubjectName sets SubjectName field to given value.

### HasSubjectName

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) HasSubjectName() bool`

HasSubjectName returns a boolean if a field has been set.

### GetSubjectID

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) GetSubjectID() int32`

GetSubjectID returns the SubjectID field if non-nil, zero value otherwise.

### GetSubjectIDOk

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) GetSubjectIDOk() (*int32, bool)`

GetSubjectIDOk returns a tuple with the SubjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectID

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) SetSubjectID(v int32)`

SetSubjectID sets SubjectID field to given value.

### HasSubjectID

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) HasSubjectID() bool`

HasSubjectID returns a boolean if a field has been set.

### GetName

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequired

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetUnitName

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) GetUnitName() string`

GetUnitName returns the UnitName field if non-nil, zero value otherwise.

### GetUnitNameOk

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) GetUnitNameOk() (*string, bool)`

GetUnitNameOk returns a tuple with the UnitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitName

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) SetUnitName(v string)`

SetUnitName sets UnitName field to given value.

### HasUnitName

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) HasUnitName() bool`

HasUnitName returns a boolean if a field has been set.

### GetMaxCount

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) GetMaxCount() int32`

GetMaxCount returns the MaxCount field if non-nil, zero value otherwise.

### GetMaxCountOk

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) GetMaxCountOk() (*int32, bool)`

GetMaxCountOk returns a tuple with the MaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCount

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) SetMaxCount(v int32)`

SetMaxCount sets MaxCount field to given value.

### HasMaxCount

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) HasMaxCount() bool`

HasMaxCount returns a boolean if a field has been set.

### GetPopular

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) GetPopular() bool`

GetPopular returns the Popular field if non-nil, zero value otherwise.

### GetPopularOk

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) GetPopularOk() (*bool, bool)`

GetPopularOk returns a tuple with the Popular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopular

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) SetPopular(v bool)`

SetPopular sets Popular field to given value.

### HasPopular

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) HasPopular() bool`

HasPopular returns a boolean if a field has been set.

### GetCharcType

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) GetCharcType() int32`

GetCharcType returns the CharcType field if non-nil, zero value otherwise.

### GetCharcTypeOk

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) GetCharcTypeOk() (*int32, bool)`

GetCharcTypeOk returns a tuple with the CharcType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharcType

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) SetCharcType(v int32)`

SetCharcType sets CharcType field to given value.

### HasCharcType

`func (o *ContentV2ObjectCharcsSubjectIdGet200ResponseDataInner) HasCharcType() bool`

HasCharcType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


