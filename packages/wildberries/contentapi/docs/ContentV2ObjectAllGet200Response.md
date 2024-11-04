# ContentV2ObjectAllGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ContentV2ObjectAllGet200ResponseDataInner**](ContentV2ObjectAllGet200ResponseDataInner.md) | Подкатегории (предметы) | [optional] 
**Error** | Pointer to **bool** | Флаг наличия ошибки | [optional] 
**ErrorText** | Pointer to **string** | Текст ошибки | [optional] 
**AdditionalErrors** | Pointer to **string** | Дополнительные ошибки | [optional] 

## Methods

### NewContentV2ObjectAllGet200Response

`func NewContentV2ObjectAllGet200Response() *ContentV2ObjectAllGet200Response`

NewContentV2ObjectAllGet200Response instantiates a new ContentV2ObjectAllGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentV2ObjectAllGet200ResponseWithDefaults

`func NewContentV2ObjectAllGet200ResponseWithDefaults() *ContentV2ObjectAllGet200Response`

NewContentV2ObjectAllGet200ResponseWithDefaults instantiates a new ContentV2ObjectAllGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ContentV2ObjectAllGet200Response) GetData() []ContentV2ObjectAllGet200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ContentV2ObjectAllGet200Response) GetDataOk() (*[]ContentV2ObjectAllGet200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ContentV2ObjectAllGet200Response) SetData(v []ContentV2ObjectAllGet200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *ContentV2ObjectAllGet200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *ContentV2ObjectAllGet200Response) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ContentV2ObjectAllGet200Response) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ContentV2ObjectAllGet200Response) SetError(v bool)`

SetError sets Error field to given value.

### HasError

`func (o *ContentV2ObjectAllGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorText

`func (o *ContentV2ObjectAllGet200Response) GetErrorText() string`

GetErrorText returns the ErrorText field if non-nil, zero value otherwise.

### GetErrorTextOk

`func (o *ContentV2ObjectAllGet200Response) GetErrorTextOk() (*string, bool)`

GetErrorTextOk returns a tuple with the ErrorText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorText

`func (o *ContentV2ObjectAllGet200Response) SetErrorText(v string)`

SetErrorText sets ErrorText field to given value.

### HasErrorText

`func (o *ContentV2ObjectAllGet200Response) HasErrorText() bool`

HasErrorText returns a boolean if a field has been set.

### GetAdditionalErrors

`func (o *ContentV2ObjectAllGet200Response) GetAdditionalErrors() string`

GetAdditionalErrors returns the AdditionalErrors field if non-nil, zero value otherwise.

### GetAdditionalErrorsOk

`func (o *ContentV2ObjectAllGet200Response) GetAdditionalErrorsOk() (*string, bool)`

GetAdditionalErrorsOk returns a tuple with the AdditionalErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalErrors

`func (o *ContentV2ObjectAllGet200Response) SetAdditionalErrors(v string)`

SetAdditionalErrors sets AdditionalErrors field to given value.

### HasAdditionalErrors

`func (o *ContentV2ObjectAllGet200Response) HasAdditionalErrors() bool`

HasAdditionalErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


