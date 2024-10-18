# ResponseContentError4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Error** | Pointer to **bool** | Флаг ошибки | [optional] 
**ErrorText** | Pointer to **string** | Описание ошибки | [optional] 
**AdditionalErrors** | Pointer to [**ResponseContentError4AdditionalErrors**](ResponseContentError4AdditionalErrors.md) |  | [optional] 

## Methods

### NewResponseContentError4

`func NewResponseContentError4() *ResponseContentError4`

NewResponseContentError4 instantiates a new ResponseContentError4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseContentError4WithDefaults

`func NewResponseContentError4WithDefaults() *ResponseContentError4`

NewResponseContentError4WithDefaults instantiates a new ResponseContentError4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResponseContentError4) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseContentError4) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseContentError4) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ResponseContentError4) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ResponseContentError4) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ResponseContentError4) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetError

`func (o *ResponseContentError4) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ResponseContentError4) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ResponseContentError4) SetError(v bool)`

SetError sets Error field to given value.

### HasError

`func (o *ResponseContentError4) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorText

`func (o *ResponseContentError4) GetErrorText() string`

GetErrorText returns the ErrorText field if non-nil, zero value otherwise.

### GetErrorTextOk

`func (o *ResponseContentError4) GetErrorTextOk() (*string, bool)`

GetErrorTextOk returns a tuple with the ErrorText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorText

`func (o *ResponseContentError4) SetErrorText(v string)`

SetErrorText sets ErrorText field to given value.

### HasErrorText

`func (o *ResponseContentError4) HasErrorText() bool`

HasErrorText returns a boolean if a field has been set.

### GetAdditionalErrors

`func (o *ResponseContentError4) GetAdditionalErrors() ResponseContentError4AdditionalErrors`

GetAdditionalErrors returns the AdditionalErrors field if non-nil, zero value otherwise.

### GetAdditionalErrorsOk

`func (o *ResponseContentError4) GetAdditionalErrorsOk() (*ResponseContentError4AdditionalErrors, bool)`

GetAdditionalErrorsOk returns a tuple with the AdditionalErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalErrors

`func (o *ResponseContentError4) SetAdditionalErrors(v ResponseContentError4AdditionalErrors)`

SetAdditionalErrors sets AdditionalErrors field to given value.

### HasAdditionalErrors

`func (o *ResponseContentError4) HasAdditionalErrors() bool`

HasAdditionalErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


