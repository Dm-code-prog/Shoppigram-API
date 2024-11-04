# ApiV2UploadTaskSizePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]SizeGoodReq**](SizeGoodReq.md) | Размеры и цены для них. Максимум 1 000 размеров &lt;br&gt;&lt;br&gt; Для товаров с поразмерной установкой цен карантин не применяется  | [optional] 

## Methods

### NewApiV2UploadTaskSizePostRequest

`func NewApiV2UploadTaskSizePostRequest() *ApiV2UploadTaskSizePostRequest`

NewApiV2UploadTaskSizePostRequest instantiates a new ApiV2UploadTaskSizePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2UploadTaskSizePostRequestWithDefaults

`func NewApiV2UploadTaskSizePostRequestWithDefaults() *ApiV2UploadTaskSizePostRequest`

NewApiV2UploadTaskSizePostRequestWithDefaults instantiates a new ApiV2UploadTaskSizePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ApiV2UploadTaskSizePostRequest) GetData() []SizeGoodReq`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApiV2UploadTaskSizePostRequest) GetDataOk() (*[]SizeGoodReq, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApiV2UploadTaskSizePostRequest) SetData(v []SizeGoodReq)`

SetData sets Data field to given value.

### HasData

`func (o *ApiV2UploadTaskSizePostRequest) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


