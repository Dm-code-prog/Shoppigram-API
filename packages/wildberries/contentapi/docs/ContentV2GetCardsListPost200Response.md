# ContentV2GetCardsListPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cards** | Pointer to [**[]ContentV2GetCardsListPost200ResponseCardsInner**](ContentV2GetCardsListPost200ResponseCardsInner.md) | Список КТ | [optional] 
**Cursor** | Pointer to [**ContentV2GetCardsListPost200ResponseCursor**](ContentV2GetCardsListPost200ResponseCursor.md) |  | [optional] 

## Methods

### NewContentV2GetCardsListPost200Response

`func NewContentV2GetCardsListPost200Response() *ContentV2GetCardsListPost200Response`

NewContentV2GetCardsListPost200Response instantiates a new ContentV2GetCardsListPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentV2GetCardsListPost200ResponseWithDefaults

`func NewContentV2GetCardsListPost200ResponseWithDefaults() *ContentV2GetCardsListPost200Response`

NewContentV2GetCardsListPost200ResponseWithDefaults instantiates a new ContentV2GetCardsListPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCards

`func (o *ContentV2GetCardsListPost200Response) GetCards() []ContentV2GetCardsListPost200ResponseCardsInner`

GetCards returns the Cards field if non-nil, zero value otherwise.

### GetCardsOk

`func (o *ContentV2GetCardsListPost200Response) GetCardsOk() (*[]ContentV2GetCardsListPost200ResponseCardsInner, bool)`

GetCardsOk returns a tuple with the Cards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCards

`func (o *ContentV2GetCardsListPost200Response) SetCards(v []ContentV2GetCardsListPost200ResponseCardsInner)`

SetCards sets Cards field to given value.

### HasCards

`func (o *ContentV2GetCardsListPost200Response) HasCards() bool`

HasCards returns a boolean if a field has been set.

### GetCursor

`func (o *ContentV2GetCardsListPost200Response) GetCursor() ContentV2GetCardsListPost200ResponseCursor`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ContentV2GetCardsListPost200Response) GetCursorOk() (*ContentV2GetCardsListPost200ResponseCursor, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ContentV2GetCardsListPost200Response) SetCursor(v ContentV2GetCardsListPost200ResponseCursor)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ContentV2GetCardsListPost200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


