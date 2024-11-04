# ContentV3MediaSavePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NmId** | Pointer to **int32** | Артикул Wildberries | [optional] 
**Data** | Pointer to **[]string** | Ссылки на изображения в том порядке, в котором они будут на карточке товара | [optional] 

## Methods

### NewContentV3MediaSavePostRequest

`func NewContentV3MediaSavePostRequest() *ContentV3MediaSavePostRequest`

NewContentV3MediaSavePostRequest instantiates a new ContentV3MediaSavePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentV3MediaSavePostRequestWithDefaults

`func NewContentV3MediaSavePostRequestWithDefaults() *ContentV3MediaSavePostRequest`

NewContentV3MediaSavePostRequestWithDefaults instantiates a new ContentV3MediaSavePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNmId

`func (o *ContentV3MediaSavePostRequest) GetNmId() int32`

GetNmId returns the NmId field if non-nil, zero value otherwise.

### GetNmIdOk

`func (o *ContentV3MediaSavePostRequest) GetNmIdOk() (*int32, bool)`

GetNmIdOk returns a tuple with the NmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmId

`func (o *ContentV3MediaSavePostRequest) SetNmId(v int32)`

SetNmId sets NmId field to given value.

### HasNmId

`func (o *ContentV3MediaSavePostRequest) HasNmId() bool`

HasNmId returns a boolean if a field has been set.

### GetData

`func (o *ContentV3MediaSavePostRequest) GetData() []string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ContentV3MediaSavePostRequest) GetDataOk() (*[]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ContentV3MediaSavePostRequest) SetData(v []string)`

SetData sets Data field to given value.

### HasData

`func (o *ContentV3MediaSavePostRequest) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


