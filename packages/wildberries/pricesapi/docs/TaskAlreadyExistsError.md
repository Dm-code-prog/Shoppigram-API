# TaskAlreadyExistsError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**TaskAlreadyExistsErrorData**](TaskAlreadyExistsErrorData.md) |  | [optional] 
**Error** | Pointer to **bool** | Флаг ошибки | [optional] 
**ErrorText** | Pointer to **string** | Текст ошибки | [optional] 

## Methods

### NewTaskAlreadyExistsError

`func NewTaskAlreadyExistsError() *TaskAlreadyExistsError`

NewTaskAlreadyExistsError instantiates a new TaskAlreadyExistsError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskAlreadyExistsErrorWithDefaults

`func NewTaskAlreadyExistsErrorWithDefaults() *TaskAlreadyExistsError`

NewTaskAlreadyExistsErrorWithDefaults instantiates a new TaskAlreadyExistsError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TaskAlreadyExistsError) GetData() TaskAlreadyExistsErrorData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TaskAlreadyExistsError) GetDataOk() (*TaskAlreadyExistsErrorData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TaskAlreadyExistsError) SetData(v TaskAlreadyExistsErrorData)`

SetData sets Data field to given value.

### HasData

`func (o *TaskAlreadyExistsError) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *TaskAlreadyExistsError) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TaskAlreadyExistsError) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TaskAlreadyExistsError) SetError(v bool)`

SetError sets Error field to given value.

### HasError

`func (o *TaskAlreadyExistsError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorText

`func (o *TaskAlreadyExistsError) GetErrorText() string`

GetErrorText returns the ErrorText field if non-nil, zero value otherwise.

### GetErrorTextOk

`func (o *TaskAlreadyExistsError) GetErrorTextOk() (*string, bool)`

GetErrorTextOk returns a tuple with the ErrorText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorText

`func (o *TaskAlreadyExistsError) SetErrorText(v string)`

SetErrorText sets ErrorText field to given value.

### HasErrorText

`func (o *TaskAlreadyExistsError) HasErrorText() bool`

HasErrorText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


