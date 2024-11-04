# ContentV2GetCardsTrashPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cards** | Pointer to [**[]ContentV2GetCardsTrashPost200ResponseCardsInner**](ContentV2GetCardsTrashPost200ResponseCardsInner.md) | Массив карточек товаров | [optional] 
**Cursor** | Pointer to [**ContentV2GetCardsTrashPost200ResponseCursor**](ContentV2GetCardsTrashPost200ResponseCursor.md) |  | [optional] 

## Methods

### NewContentV2GetCardsTrashPost200Response

`func NewContentV2GetCardsTrashPost200Response() *ContentV2GetCardsTrashPost200Response`

NewContentV2GetCardsTrashPost200Response instantiates a new ContentV2GetCardsTrashPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentV2GetCardsTrashPost200ResponseWithDefaults

`func NewContentV2GetCardsTrashPost200ResponseWithDefaults() *ContentV2GetCardsTrashPost200Response`

NewContentV2GetCardsTrashPost200ResponseWithDefaults instantiates a new ContentV2GetCardsTrashPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCards

`func (o *ContentV2GetCardsTrashPost200Response) GetCards() []ContentV2GetCardsTrashPost200ResponseCardsInner`

GetCards returns the Cards field if non-nil, zero value otherwise.

### GetCardsOk

`func (o *ContentV2GetCardsTrashPost200Response) GetCardsOk() (*[]ContentV2GetCardsTrashPost200ResponseCardsInner, bool)`

GetCardsOk returns a tuple with the Cards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCards

`func (o *ContentV2GetCardsTrashPost200Response) SetCards(v []ContentV2GetCardsTrashPost200ResponseCardsInner)`

SetCards sets Cards field to given value.

### HasCards

`func (o *ContentV2GetCardsTrashPost200Response) HasCards() bool`

HasCards returns a boolean if a field has been set.

### GetCursor

`func (o *ContentV2GetCardsTrashPost200Response) GetCursor() ContentV2GetCardsTrashPost200ResponseCursor`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ContentV2GetCardsTrashPost200Response) GetCursorOk() (*ContentV2GetCardsTrashPost200ResponseCursor, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ContentV2GetCardsTrashPost200Response) SetCursor(v ContentV2GetCardsTrashPost200ResponseCursor)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ContentV2GetCardsTrashPost200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


