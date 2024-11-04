# ContentV2TagNomenclatureLinkPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NmID** | Pointer to **int32** | Артикул WB | [optional] 
**TagsIDs** | Pointer to **[]int32** | Массив числовых идентификаторов тегов.&lt;br&gt;   Что бы снять теги с КТ, необходимо передать пустой массив.&lt;br&gt; Чтобы добавить теги к уже имеющимся в КТ, необходимо в запросе передать новые теги и теги, которые уже есть в КТ.  | [optional] 

## Methods

### NewContentV2TagNomenclatureLinkPostRequest

`func NewContentV2TagNomenclatureLinkPostRequest() *ContentV2TagNomenclatureLinkPostRequest`

NewContentV2TagNomenclatureLinkPostRequest instantiates a new ContentV2TagNomenclatureLinkPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentV2TagNomenclatureLinkPostRequestWithDefaults

`func NewContentV2TagNomenclatureLinkPostRequestWithDefaults() *ContentV2TagNomenclatureLinkPostRequest`

NewContentV2TagNomenclatureLinkPostRequestWithDefaults instantiates a new ContentV2TagNomenclatureLinkPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNmID

`func (o *ContentV2TagNomenclatureLinkPostRequest) GetNmID() int32`

GetNmID returns the NmID field if non-nil, zero value otherwise.

### GetNmIDOk

`func (o *ContentV2TagNomenclatureLinkPostRequest) GetNmIDOk() (*int32, bool)`

GetNmIDOk returns a tuple with the NmID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmID

`func (o *ContentV2TagNomenclatureLinkPostRequest) SetNmID(v int32)`

SetNmID sets NmID field to given value.

### HasNmID

`func (o *ContentV2TagNomenclatureLinkPostRequest) HasNmID() bool`

HasNmID returns a boolean if a field has been set.

### GetTagsIDs

`func (o *ContentV2TagNomenclatureLinkPostRequest) GetTagsIDs() []int32`

GetTagsIDs returns the TagsIDs field if non-nil, zero value otherwise.

### GetTagsIDsOk

`func (o *ContentV2TagNomenclatureLinkPostRequest) GetTagsIDsOk() (*[]int32, bool)`

GetTagsIDsOk returns a tuple with the TagsIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsIDs

`func (o *ContentV2TagNomenclatureLinkPostRequest) SetTagsIDs(v []int32)`

SetTagsIDs sets TagsIDs field to given value.

### HasTagsIDs

`func (o *ContentV2TagNomenclatureLinkPostRequest) HasTagsIDs() bool`

HasTagsIDs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


