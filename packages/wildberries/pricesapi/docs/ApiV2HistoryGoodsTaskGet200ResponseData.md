# ApiV2HistoryGoodsTaskGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UploadID** | Pointer to **int32** | ID загрузки | [optional] 
**HistoryGoods** | Pointer to [**[]GoodHistory**](GoodHistory.md) | Информация о товарах в загрузке | [optional] 

## Methods

### NewApiV2HistoryGoodsTaskGet200ResponseData

`func NewApiV2HistoryGoodsTaskGet200ResponseData() *ApiV2HistoryGoodsTaskGet200ResponseData`

NewApiV2HistoryGoodsTaskGet200ResponseData instantiates a new ApiV2HistoryGoodsTaskGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2HistoryGoodsTaskGet200ResponseDataWithDefaults

`func NewApiV2HistoryGoodsTaskGet200ResponseDataWithDefaults() *ApiV2HistoryGoodsTaskGet200ResponseData`

NewApiV2HistoryGoodsTaskGet200ResponseDataWithDefaults instantiates a new ApiV2HistoryGoodsTaskGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUploadID

`func (o *ApiV2HistoryGoodsTaskGet200ResponseData) GetUploadID() int32`

GetUploadID returns the UploadID field if non-nil, zero value otherwise.

### GetUploadIDOk

`func (o *ApiV2HistoryGoodsTaskGet200ResponseData) GetUploadIDOk() (*int32, bool)`

GetUploadIDOk returns a tuple with the UploadID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadID

`func (o *ApiV2HistoryGoodsTaskGet200ResponseData) SetUploadID(v int32)`

SetUploadID sets UploadID field to given value.

### HasUploadID

`func (o *ApiV2HistoryGoodsTaskGet200ResponseData) HasUploadID() bool`

HasUploadID returns a boolean if a field has been set.

### GetHistoryGoods

`func (o *ApiV2HistoryGoodsTaskGet200ResponseData) GetHistoryGoods() []GoodHistory`

GetHistoryGoods returns the HistoryGoods field if non-nil, zero value otherwise.

### GetHistoryGoodsOk

`func (o *ApiV2HistoryGoodsTaskGet200ResponseData) GetHistoryGoodsOk() (*[]GoodHistory, bool)`

GetHistoryGoodsOk returns a tuple with the HistoryGoods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryGoods

`func (o *ApiV2HistoryGoodsTaskGet200ResponseData) SetHistoryGoods(v []GoodHistory)`

SetHistoryGoods sets HistoryGoods field to given value.

### HasHistoryGoods

`func (o *ApiV2HistoryGoodsTaskGet200ResponseData) HasHistoryGoods() bool`

HasHistoryGoods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


