# TaskAlreadyExistsErrorData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID загрузки | [optional] 
**AlreadyExists** | Pointer to **bool** | Флаг дублирования загрузки: &#x60;true&#x60; — такая загрузка уже есть  | [optional] 

## Methods

### NewTaskAlreadyExistsErrorData

`func NewTaskAlreadyExistsErrorData() *TaskAlreadyExistsErrorData`

NewTaskAlreadyExistsErrorData instantiates a new TaskAlreadyExistsErrorData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskAlreadyExistsErrorDataWithDefaults

`func NewTaskAlreadyExistsErrorDataWithDefaults() *TaskAlreadyExistsErrorData`

NewTaskAlreadyExistsErrorDataWithDefaults instantiates a new TaskAlreadyExistsErrorData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskAlreadyExistsErrorData) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskAlreadyExistsErrorData) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskAlreadyExistsErrorData) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TaskAlreadyExistsErrorData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAlreadyExists

`func (o *TaskAlreadyExistsErrorData) GetAlreadyExists() bool`

GetAlreadyExists returns the AlreadyExists field if non-nil, zero value otherwise.

### GetAlreadyExistsOk

`func (o *TaskAlreadyExistsErrorData) GetAlreadyExistsOk() (*bool, bool)`

GetAlreadyExistsOk returns a tuple with the AlreadyExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyExists

`func (o *TaskAlreadyExistsErrorData) SetAlreadyExists(v bool)`

SetAlreadyExists sets AlreadyExists field to given value.

### HasAlreadyExists

`func (o *TaskAlreadyExistsErrorData) HasAlreadyExists() bool`

HasAlreadyExists returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


