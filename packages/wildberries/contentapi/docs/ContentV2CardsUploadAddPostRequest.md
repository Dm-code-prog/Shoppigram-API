# ContentV2CardsUploadAddPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImtID** | Pointer to **int32** | imtID КТ, к которой добавляется НМ | [optional] 
**CardsToAdd** | Pointer to [**[]ContentV2CardsUploadAddPostRequestCardsToAddInner**](ContentV2CardsUploadAddPostRequestCardsToAddInner.md) | Структура добавляемой НМ | [optional] 

## Methods

### NewContentV2CardsUploadAddPostRequest

`func NewContentV2CardsUploadAddPostRequest() *ContentV2CardsUploadAddPostRequest`

NewContentV2CardsUploadAddPostRequest instantiates a new ContentV2CardsUploadAddPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentV2CardsUploadAddPostRequestWithDefaults

`func NewContentV2CardsUploadAddPostRequestWithDefaults() *ContentV2CardsUploadAddPostRequest`

NewContentV2CardsUploadAddPostRequestWithDefaults instantiates a new ContentV2CardsUploadAddPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImtID

`func (o *ContentV2CardsUploadAddPostRequest) GetImtID() int32`

GetImtID returns the ImtID field if non-nil, zero value otherwise.

### GetImtIDOk

`func (o *ContentV2CardsUploadAddPostRequest) GetImtIDOk() (*int32, bool)`

GetImtIDOk returns a tuple with the ImtID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImtID

`func (o *ContentV2CardsUploadAddPostRequest) SetImtID(v int32)`

SetImtID sets ImtID field to given value.

### HasImtID

`func (o *ContentV2CardsUploadAddPostRequest) HasImtID() bool`

HasImtID returns a boolean if a field has been set.

### GetCardsToAdd

`func (o *ContentV2CardsUploadAddPostRequest) GetCardsToAdd() []ContentV2CardsUploadAddPostRequestCardsToAddInner`

GetCardsToAdd returns the CardsToAdd field if non-nil, zero value otherwise.

### GetCardsToAddOk

`func (o *ContentV2CardsUploadAddPostRequest) GetCardsToAddOk() (*[]ContentV2CardsUploadAddPostRequestCardsToAddInner, bool)`

GetCardsToAddOk returns a tuple with the CardsToAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardsToAdd

`func (o *ContentV2CardsUploadAddPostRequest) SetCardsToAdd(v []ContentV2CardsUploadAddPostRequestCardsToAddInner)`

SetCardsToAdd sets CardsToAdd field to given value.

### HasCardsToAdd

`func (o *ContentV2CardsUploadAddPostRequest) HasCardsToAdd() bool`

HasCardsToAdd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


