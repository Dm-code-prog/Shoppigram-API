# ContentV2GetCardsListPost200ResponseCursor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdatedAt** | Pointer to **string** | Дата с которой надо запрашивать следующий список КТ | [optional] 
**NmID** | Pointer to **int32** | Номер Артикула WB с которой надо запрашивать следующий список КТ | [optional] 
**Total** | Pointer to **int32** | Кол-во возвращенных КТ | [optional] 

## Methods

### NewContentV2GetCardsListPost200ResponseCursor

`func NewContentV2GetCardsListPost200ResponseCursor() *ContentV2GetCardsListPost200ResponseCursor`

NewContentV2GetCardsListPost200ResponseCursor instantiates a new ContentV2GetCardsListPost200ResponseCursor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentV2GetCardsListPost200ResponseCursorWithDefaults

`func NewContentV2GetCardsListPost200ResponseCursorWithDefaults() *ContentV2GetCardsListPost200ResponseCursor`

NewContentV2GetCardsListPost200ResponseCursorWithDefaults instantiates a new ContentV2GetCardsListPost200ResponseCursor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdatedAt

`func (o *ContentV2GetCardsListPost200ResponseCursor) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ContentV2GetCardsListPost200ResponseCursor) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ContentV2GetCardsListPost200ResponseCursor) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ContentV2GetCardsListPost200ResponseCursor) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetNmID

`func (o *ContentV2GetCardsListPost200ResponseCursor) GetNmID() int32`

GetNmID returns the NmID field if non-nil, zero value otherwise.

### GetNmIDOk

`func (o *ContentV2GetCardsListPost200ResponseCursor) GetNmIDOk() (*int32, bool)`

GetNmIDOk returns a tuple with the NmID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmID

`func (o *ContentV2GetCardsListPost200ResponseCursor) SetNmID(v int32)`

SetNmID sets NmID field to given value.

### HasNmID

`func (o *ContentV2GetCardsListPost200ResponseCursor) HasNmID() bool`

HasNmID returns a boolean if a field has been set.

### GetTotal

`func (o *ContentV2GetCardsListPost200ResponseCursor) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ContentV2GetCardsListPost200ResponseCursor) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ContentV2GetCardsListPost200ResponseCursor) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ContentV2GetCardsListPost200ResponseCursor) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


