# TaskCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**TaskCreatedData**](TaskCreatedData.md) |  | [optional] 
**Error** | Pointer to **bool** | Флаг ошибки | [optional] 
**ErrorText** | Pointer to **string** | Текст ошибки | [optional] 

## Methods

### NewTaskCreated

`func NewTaskCreated() *TaskCreated`

NewTaskCreated instantiates a new TaskCreated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskCreatedWithDefaults

`func NewTaskCreatedWithDefaults() *TaskCreated`

NewTaskCreatedWithDefaults instantiates a new TaskCreated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TaskCreated) GetData() TaskCreatedData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TaskCreated) GetDataOk() (*TaskCreatedData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TaskCreated) SetData(v TaskCreatedData)`

SetData sets Data field to given value.

### HasData

`func (o *TaskCreated) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *TaskCreated) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TaskCreated) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TaskCreated) SetError(v bool)`

SetError sets Error field to given value.

### HasError

`func (o *TaskCreated) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorText

`func (o *TaskCreated) GetErrorText() string`

GetErrorText returns the ErrorText field if non-nil, zero value otherwise.

### GetErrorTextOk

`func (o *TaskCreated) GetErrorTextOk() (*string, bool)`

GetErrorTextOk returns a tuple with the ErrorText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorText

`func (o *TaskCreated) SetErrorText(v string)`

SetErrorText sets ErrorText field to given value.

### HasErrorText

`func (o *TaskCreated) HasErrorText() bool`

HasErrorText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


