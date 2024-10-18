# ResponseCardCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Error** | Pointer to **bool** | Флаг ошибки | [optional] 
**ErrorText** | Pointer to **string** | Описание ошибки | [optional] 
**AdditionalErrors** | Pointer to [**ResponseCardCreateAdditionalErrors**](ResponseCardCreateAdditionalErrors.md) |  | [optional] 

## Methods

### NewResponseCardCreate

`func NewResponseCardCreate() *ResponseCardCreate`

NewResponseCardCreate instantiates a new ResponseCardCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseCardCreateWithDefaults

`func NewResponseCardCreateWithDefaults() *ResponseCardCreate`

NewResponseCardCreateWithDefaults instantiates a new ResponseCardCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResponseCardCreate) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseCardCreate) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseCardCreate) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ResponseCardCreate) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ResponseCardCreate) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ResponseCardCreate) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetError

`func (o *ResponseCardCreate) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ResponseCardCreate) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ResponseCardCreate) SetError(v bool)`

SetError sets Error field to given value.

### HasError

`func (o *ResponseCardCreate) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorText

`func (o *ResponseCardCreate) GetErrorText() string`

GetErrorText returns the ErrorText field if non-nil, zero value otherwise.

### GetErrorTextOk

`func (o *ResponseCardCreate) GetErrorTextOk() (*string, bool)`

GetErrorTextOk returns a tuple with the ErrorText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorText

`func (o *ResponseCardCreate) SetErrorText(v string)`

SetErrorText sets ErrorText field to given value.

### HasErrorText

`func (o *ResponseCardCreate) HasErrorText() bool`

HasErrorText returns a boolean if a field has been set.

### GetAdditionalErrors

`func (o *ResponseCardCreate) GetAdditionalErrors() ResponseCardCreateAdditionalErrors`

GetAdditionalErrors returns the AdditionalErrors field if non-nil, zero value otherwise.

### GetAdditionalErrorsOk

`func (o *ResponseCardCreate) GetAdditionalErrorsOk() (*ResponseCardCreateAdditionalErrors, bool)`

GetAdditionalErrorsOk returns a tuple with the AdditionalErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalErrors

`func (o *ResponseCardCreate) SetAdditionalErrors(v ResponseCardCreateAdditionalErrors)`

SetAdditionalErrors sets AdditionalErrors field to given value.

### HasAdditionalErrors

`func (o *ResponseCardCreate) HasAdditionalErrors() bool`

HasAdditionalErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


