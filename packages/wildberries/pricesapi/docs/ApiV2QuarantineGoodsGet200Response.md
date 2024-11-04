# ApiV2QuarantineGoodsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ApiV2QuarantineGoodsGet200ResponseData**](ApiV2QuarantineGoodsGet200ResponseData.md) |  | [optional] 
**Error** | Pointer to **bool** | Флаг ошибки | [optional] 
**ErrorText** | Pointer to **string** | Текст ошибки | [optional] 

## Methods

### NewApiV2QuarantineGoodsGet200Response

`func NewApiV2QuarantineGoodsGet200Response() *ApiV2QuarantineGoodsGet200Response`

NewApiV2QuarantineGoodsGet200Response instantiates a new ApiV2QuarantineGoodsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2QuarantineGoodsGet200ResponseWithDefaults

`func NewApiV2QuarantineGoodsGet200ResponseWithDefaults() *ApiV2QuarantineGoodsGet200Response`

NewApiV2QuarantineGoodsGet200ResponseWithDefaults instantiates a new ApiV2QuarantineGoodsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ApiV2QuarantineGoodsGet200Response) GetData() ApiV2QuarantineGoodsGet200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApiV2QuarantineGoodsGet200Response) GetDataOk() (*ApiV2QuarantineGoodsGet200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApiV2QuarantineGoodsGet200Response) SetData(v ApiV2QuarantineGoodsGet200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *ApiV2QuarantineGoodsGet200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *ApiV2QuarantineGoodsGet200Response) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiV2QuarantineGoodsGet200Response) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiV2QuarantineGoodsGet200Response) SetError(v bool)`

SetError sets Error field to given value.

### HasError

`func (o *ApiV2QuarantineGoodsGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorText

`func (o *ApiV2QuarantineGoodsGet200Response) GetErrorText() string`

GetErrorText returns the ErrorText field if non-nil, zero value otherwise.

### GetErrorTextOk

`func (o *ApiV2QuarantineGoodsGet200Response) GetErrorTextOk() (*string, bool)`

GetErrorTextOk returns a tuple with the ErrorText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorText

`func (o *ApiV2QuarantineGoodsGet200Response) SetErrorText(v string)`

SetErrorText sets ErrorText field to given value.

### HasErrorText

`func (o *ApiV2QuarantineGoodsGet200Response) HasErrorText() bool`

HasErrorText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


