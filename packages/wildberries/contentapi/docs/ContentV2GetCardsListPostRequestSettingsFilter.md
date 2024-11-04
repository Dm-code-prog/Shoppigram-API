# ContentV2GetCardsListPostRequestSettingsFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WithPhoto** | Pointer to **int32** | Фильтр по фото:    * &#x60;0&#x60; — только карточки без фото   * &#x60;1&#x60; — только карточки с фото   * &#x60;-1&#x60; — все карточки товара  | [optional] 
**TextSearch** | Pointer to **string** | Поиск по артикулу продавца, артикулу WB, баркоду | [optional] 
**TagIDs** | Pointer to **[]int32** | Поиск по ID тегов | [optional] 
**AllowedCategoriesOnly** | Pointer to **bool** | Фильтр по категории. &#x60;true&#x60; - только разрешённые, &#x60;false&#x60; - все. Не используется в песочнице. | [optional] 
**ObjectIDs** | Pointer to **[]int32** | Поиск по id предметов | [optional] 
**Brands** | Pointer to **[]string** | Поиск по брендам | [optional] 
**ImtID** | Pointer to **int32** | Поиск по идентификатору КТ | [optional] 

## Methods

### NewContentV2GetCardsListPostRequestSettingsFilter

`func NewContentV2GetCardsListPostRequestSettingsFilter() *ContentV2GetCardsListPostRequestSettingsFilter`

NewContentV2GetCardsListPostRequestSettingsFilter instantiates a new ContentV2GetCardsListPostRequestSettingsFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentV2GetCardsListPostRequestSettingsFilterWithDefaults

`func NewContentV2GetCardsListPostRequestSettingsFilterWithDefaults() *ContentV2GetCardsListPostRequestSettingsFilter`

NewContentV2GetCardsListPostRequestSettingsFilterWithDefaults instantiates a new ContentV2GetCardsListPostRequestSettingsFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWithPhoto

`func (o *ContentV2GetCardsListPostRequestSettingsFilter) GetWithPhoto() int32`

GetWithPhoto returns the WithPhoto field if non-nil, zero value otherwise.

### GetWithPhotoOk

`func (o *ContentV2GetCardsListPostRequestSettingsFilter) GetWithPhotoOk() (*int32, bool)`

GetWithPhotoOk returns a tuple with the WithPhoto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithPhoto

`func (o *ContentV2GetCardsListPostRequestSettingsFilter) SetWithPhoto(v int32)`

SetWithPhoto sets WithPhoto field to given value.

### HasWithPhoto

`func (o *ContentV2GetCardsListPostRequestSettingsFilter) HasWithPhoto() bool`

HasWithPhoto returns a boolean if a field has been set.

### GetTextSearch

`func (o *ContentV2GetCardsListPostRequestSettingsFilter) GetTextSearch() string`

GetTextSearch returns the TextSearch field if non-nil, zero value otherwise.

### GetTextSearchOk

`func (o *ContentV2GetCardsListPostRequestSettingsFilter) GetTextSearchOk() (*string, bool)`

GetTextSearchOk returns a tuple with the TextSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextSearch

`func (o *ContentV2GetCardsListPostRequestSettingsFilter) SetTextSearch(v string)`

SetTextSearch sets TextSearch field to given value.

### HasTextSearch

`func (o *ContentV2GetCardsListPostRequestSettingsFilter) HasTextSearch() bool`

HasTextSearch returns a boolean if a field has been set.

### GetTagIDs

`func (o *ContentV2GetCardsListPostRequestSettingsFilter) GetTagIDs() []int32`

GetTagIDs returns the TagIDs field if non-nil, zero value otherwise.

### GetTagIDsOk

`func (o *ContentV2GetCardsListPostRequestSettingsFilter) GetTagIDsOk() (*[]int32, bool)`

GetTagIDsOk returns a tuple with the TagIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIDs

`func (o *ContentV2GetCardsListPostRequestSettingsFilter) SetTagIDs(v []int32)`

SetTagIDs sets TagIDs field to given value.

### HasTagIDs

`func (o *ContentV2GetCardsListPostRequestSettingsFilter) HasTagIDs() bool`

HasTagIDs returns a boolean if a field has been set.

### GetAllowedCategoriesOnly

`func (o *ContentV2GetCardsListPostRequestSettingsFilter) GetAllowedCategoriesOnly() bool`

GetAllowedCategoriesOnly returns the AllowedCategoriesOnly field if non-nil, zero value otherwise.

### GetAllowedCategoriesOnlyOk

`func (o *ContentV2GetCardsListPostRequestSettingsFilter) GetAllowedCategoriesOnlyOk() (*bool, bool)`

GetAllowedCategoriesOnlyOk returns a tuple with the AllowedCategoriesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCategoriesOnly

`func (o *ContentV2GetCardsListPostRequestSettingsFilter) SetAllowedCategoriesOnly(v bool)`

SetAllowedCategoriesOnly sets AllowedCategoriesOnly field to given value.

### HasAllowedCategoriesOnly

`func (o *ContentV2GetCardsListPostRequestSettingsFilter) HasAllowedCategoriesOnly() bool`

HasAllowedCategoriesOnly returns a boolean if a field has been set.

### GetObjectIDs

`func (o *ContentV2GetCardsListPostRequestSettingsFilter) GetObjectIDs() []int32`

GetObjectIDs returns the ObjectIDs field if non-nil, zero value otherwise.

### GetObjectIDsOk

`func (o *ContentV2GetCardsListPostRequestSettingsFilter) GetObjectIDsOk() (*[]int32, bool)`

GetObjectIDsOk returns a tuple with the ObjectIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIDs

`func (o *ContentV2GetCardsListPostRequestSettingsFilter) SetObjectIDs(v []int32)`

SetObjectIDs sets ObjectIDs field to given value.

### HasObjectIDs

`func (o *ContentV2GetCardsListPostRequestSettingsFilter) HasObjectIDs() bool`

HasObjectIDs returns a boolean if a field has been set.

### GetBrands

`func (o *ContentV2GetCardsListPostRequestSettingsFilter) GetBrands() []string`

GetBrands returns the Brands field if non-nil, zero value otherwise.

### GetBrandsOk

`func (o *ContentV2GetCardsListPostRequestSettingsFilter) GetBrandsOk() (*[]string, bool)`

GetBrandsOk returns a tuple with the Brands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrands

`func (o *ContentV2GetCardsListPostRequestSettingsFilter) SetBrands(v []string)`

SetBrands sets Brands field to given value.

### HasBrands

`func (o *ContentV2GetCardsListPostRequestSettingsFilter) HasBrands() bool`

HasBrands returns a boolean if a field has been set.

### GetImtID

`func (o *ContentV2GetCardsListPostRequestSettingsFilter) GetImtID() int32`

GetImtID returns the ImtID field if non-nil, zero value otherwise.

### GetImtIDOk

`func (o *ContentV2GetCardsListPostRequestSettingsFilter) GetImtIDOk() (*int32, bool)`

GetImtIDOk returns a tuple with the ImtID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImtID

`func (o *ContentV2GetCardsListPostRequestSettingsFilter) SetImtID(v int32)`

SetImtID sets ImtID field to given value.

### HasImtID

`func (o *ContentV2GetCardsListPostRequestSettingsFilter) HasImtID() bool`

HasImtID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


