# ResponseContentError5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Error** | Pointer to **bool** | Флаг ошибки | [optional] 
**ErrorText** | Pointer to **string** | Описание ошибки | [optional] 
**AdditionalErrors** | Pointer to [**ResponseContentError5AdditionalErrors**](ResponseContentError5AdditionalErrors.md) |  | [optional] 

## Methods

### NewResponseContentError5

`func NewResponseContentError5() *ResponseContentError5`

NewResponseContentError5 instantiates a new ResponseContentError5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseContentError5WithDefaults

`func NewResponseContentError5WithDefaults() *ResponseContentError5`

NewResponseContentError5WithDefaults instantiates a new ResponseContentError5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResponseContentError5) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseContentError5) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseContentError5) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ResponseContentError5) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ResponseContentError5) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ResponseContentError5) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetError

`func (o *ResponseContentError5) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ResponseContentError5) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ResponseContentError5) SetError(v bool)`

SetError sets Error field to given value.

### HasError

`func (o *ResponseContentError5) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorText

`func (o *ResponseContentError5) GetErrorText() string`

GetErrorText returns the ErrorText field if non-nil, zero value otherwise.

### GetErrorTextOk

`func (o *ResponseContentError5) GetErrorTextOk() (*string, bool)`

GetErrorTextOk returns a tuple with the ErrorText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorText

`func (o *ResponseContentError5) SetErrorText(v string)`

SetErrorText sets ErrorText field to given value.

### HasErrorText

`func (o *ResponseContentError5) HasErrorText() bool`

HasErrorText returns a boolean if a field has been set.

### GetAdditionalErrors

`func (o *ResponseContentError5) GetAdditionalErrors() ResponseContentError5AdditionalErrors`

GetAdditionalErrors returns the AdditionalErrors field if non-nil, zero value otherwise.

### GetAdditionalErrorsOk

`func (o *ResponseContentError5) GetAdditionalErrorsOk() (*ResponseContentError5AdditionalErrors, bool)`

GetAdditionalErrorsOk returns a tuple with the AdditionalErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalErrors

`func (o *ResponseContentError5) SetAdditionalErrors(v ResponseContentError5AdditionalErrors)`

SetAdditionalErrors sets AdditionalErrors field to given value.

### HasAdditionalErrors

`func (o *ResponseContentError5) HasAdditionalErrors() bool`

HasAdditionalErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


