# MediaErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalErrors** | Pointer to **map[string]interface{}** | Дополнительные ошибки | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Error** | Pointer to **bool** | Флаг ошибки | [optional] 
**ErrorText** | Pointer to **string** | Описание ошибки | [optional] 

## Methods

### NewMediaErrors

`func NewMediaErrors() *MediaErrors`

NewMediaErrors instantiates a new MediaErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaErrorsWithDefaults

`func NewMediaErrorsWithDefaults() *MediaErrors`

NewMediaErrorsWithDefaults instantiates a new MediaErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalErrors

`func (o *MediaErrors) GetAdditionalErrors() map[string]interface{}`

GetAdditionalErrors returns the AdditionalErrors field if non-nil, zero value otherwise.

### GetAdditionalErrorsOk

`func (o *MediaErrors) GetAdditionalErrorsOk() (*map[string]interface{}, bool)`

GetAdditionalErrorsOk returns a tuple with the AdditionalErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalErrors

`func (o *MediaErrors) SetAdditionalErrors(v map[string]interface{})`

SetAdditionalErrors sets AdditionalErrors field to given value.

### HasAdditionalErrors

`func (o *MediaErrors) HasAdditionalErrors() bool`

HasAdditionalErrors returns a boolean if a field has been set.

### GetData

`func (o *MediaErrors) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MediaErrors) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MediaErrors) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *MediaErrors) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *MediaErrors) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MediaErrors) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MediaErrors) SetError(v bool)`

SetError sets Error field to given value.

### HasError

`func (o *MediaErrors) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorText

`func (o *MediaErrors) GetErrorText() string`

GetErrorText returns the ErrorText field if non-nil, zero value otherwise.

### GetErrorTextOk

`func (o *MediaErrors) GetErrorTextOk() (*string, bool)`

GetErrorTextOk returns a tuple with the ErrorText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorText

`func (o *MediaErrors) SetErrorText(v string)`

SetErrorText sets ErrorText field to given value.

### HasErrorText

`func (o *MediaErrors) HasErrorText() bool`

HasErrorText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


