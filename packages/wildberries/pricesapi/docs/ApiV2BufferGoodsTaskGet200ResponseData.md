# ApiV2BufferGoodsTaskGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UploadID** | Pointer to **int32** | ID загрузки | [optional] 
**BufferGoods** | Pointer to [**[]GoodBufferHistory**](GoodBufferHistory.md) | Информация о товарах в загрузке | [optional] 

## Methods

### NewApiV2BufferGoodsTaskGet200ResponseData

`func NewApiV2BufferGoodsTaskGet200ResponseData() *ApiV2BufferGoodsTaskGet200ResponseData`

NewApiV2BufferGoodsTaskGet200ResponseData instantiates a new ApiV2BufferGoodsTaskGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2BufferGoodsTaskGet200ResponseDataWithDefaults

`func NewApiV2BufferGoodsTaskGet200ResponseDataWithDefaults() *ApiV2BufferGoodsTaskGet200ResponseData`

NewApiV2BufferGoodsTaskGet200ResponseDataWithDefaults instantiates a new ApiV2BufferGoodsTaskGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUploadID

`func (o *ApiV2BufferGoodsTaskGet200ResponseData) GetUploadID() int32`

GetUploadID returns the UploadID field if non-nil, zero value otherwise.

### GetUploadIDOk

`func (o *ApiV2BufferGoodsTaskGet200ResponseData) GetUploadIDOk() (*int32, bool)`

GetUploadIDOk returns a tuple with the UploadID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadID

`func (o *ApiV2BufferGoodsTaskGet200ResponseData) SetUploadID(v int32)`

SetUploadID sets UploadID field to given value.

### HasUploadID

`func (o *ApiV2BufferGoodsTaskGet200ResponseData) HasUploadID() bool`

HasUploadID returns a boolean if a field has been set.

### GetBufferGoods

`func (o *ApiV2BufferGoodsTaskGet200ResponseData) GetBufferGoods() []GoodBufferHistory`

GetBufferGoods returns the BufferGoods field if non-nil, zero value otherwise.

### GetBufferGoodsOk

`func (o *ApiV2BufferGoodsTaskGet200ResponseData) GetBufferGoodsOk() (*[]GoodBufferHistory, bool)`

GetBufferGoodsOk returns a tuple with the BufferGoods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferGoods

`func (o *ApiV2BufferGoodsTaskGet200ResponseData) SetBufferGoods(v []GoodBufferHistory)`

SetBufferGoods sets BufferGoods field to given value.

### HasBufferGoods

`func (o *ApiV2BufferGoodsTaskGet200ResponseData) HasBufferGoods() bool`

HasBufferGoods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


