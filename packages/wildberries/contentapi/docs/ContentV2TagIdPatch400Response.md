# ContentV2TagIdPatch400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Error** | Pointer to **bool** | Флаг ошибки | [optional] 
**ErrorText** | Pointer to **string** | Описание ошибки | [optional] 
**AdditionalErrors** | Pointer to [**ResponseContentError4AdditionalErrors**](ResponseContentError4AdditionalErrors.md) |  | [optional] 

## Methods

### NewContentV2TagIdPatch400Response

`func NewContentV2TagIdPatch400Response() *ContentV2TagIdPatch400Response`

NewContentV2TagIdPatch400Response instantiates a new ContentV2TagIdPatch400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentV2TagIdPatch400ResponseWithDefaults

`func NewContentV2TagIdPatch400ResponseWithDefaults() *ContentV2TagIdPatch400Response`

NewContentV2TagIdPatch400ResponseWithDefaults instantiates a new ContentV2TagIdPatch400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ContentV2TagIdPatch400Response) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ContentV2TagIdPatch400Response) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ContentV2TagIdPatch400Response) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ContentV2TagIdPatch400Response) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ContentV2TagIdPatch400Response) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ContentV2TagIdPatch400Response) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetError

`func (o *ContentV2TagIdPatch400Response) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ContentV2TagIdPatch400Response) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ContentV2TagIdPatch400Response) SetError(v bool)`

SetError sets Error field to given value.

### HasError

`func (o *ContentV2TagIdPatch400Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorText

`func (o *ContentV2TagIdPatch400Response) GetErrorText() string`

GetErrorText returns the ErrorText field if non-nil, zero value otherwise.

### GetErrorTextOk

`func (o *ContentV2TagIdPatch400Response) GetErrorTextOk() (*string, bool)`

GetErrorTextOk returns a tuple with the ErrorText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorText

`func (o *ContentV2TagIdPatch400Response) SetErrorText(v string)`

SetErrorText sets ErrorText field to given value.

### HasErrorText

`func (o *ContentV2TagIdPatch400Response) HasErrorText() bool`

HasErrorText returns a boolean if a field has been set.

### GetAdditionalErrors

`func (o *ContentV2TagIdPatch400Response) GetAdditionalErrors() ResponseContentError4AdditionalErrors`

GetAdditionalErrors returns the AdditionalErrors field if non-nil, zero value otherwise.

### GetAdditionalErrorsOk

`func (o *ContentV2TagIdPatch400Response) GetAdditionalErrorsOk() (*ResponseContentError4AdditionalErrors, bool)`

GetAdditionalErrorsOk returns a tuple with the AdditionalErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalErrors

`func (o *ContentV2TagIdPatch400Response) SetAdditionalErrors(v ResponseContentError4AdditionalErrors)`

SetAdditionalErrors sets AdditionalErrors field to given value.

### HasAdditionalErrors

`func (o *ContentV2TagIdPatch400Response) HasAdditionalErrors() bool`

HasAdditionalErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


