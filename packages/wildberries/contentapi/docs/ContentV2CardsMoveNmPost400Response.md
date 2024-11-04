# ContentV2CardsMoveNmPost400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**ErrorText** | Pointer to **string** | Описание ошибки | [optional] 
**AdditionalErrors** | Pointer to [**ResponseCardCreateAdditionalErrors**](ResponseCardCreateAdditionalErrors.md) |  | [optional] 

## Methods

### NewContentV2CardsMoveNmPost400Response

`func NewContentV2CardsMoveNmPost400Response() *ContentV2CardsMoveNmPost400Response`

NewContentV2CardsMoveNmPost400Response instantiates a new ContentV2CardsMoveNmPost400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentV2CardsMoveNmPost400ResponseWithDefaults

`func NewContentV2CardsMoveNmPost400ResponseWithDefaults() *ContentV2CardsMoveNmPost400Response`

NewContentV2CardsMoveNmPost400ResponseWithDefaults instantiates a new ContentV2CardsMoveNmPost400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ContentV2CardsMoveNmPost400Response) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ContentV2CardsMoveNmPost400Response) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ContentV2CardsMoveNmPost400Response) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ContentV2CardsMoveNmPost400Response) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ContentV2CardsMoveNmPost400Response) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ContentV2CardsMoveNmPost400Response) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetError

`func (o *ContentV2CardsMoveNmPost400Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ContentV2CardsMoveNmPost400Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ContentV2CardsMoveNmPost400Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ContentV2CardsMoveNmPost400Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorText

`func (o *ContentV2CardsMoveNmPost400Response) GetErrorText() string`

GetErrorText returns the ErrorText field if non-nil, zero value otherwise.

### GetErrorTextOk

`func (o *ContentV2CardsMoveNmPost400Response) GetErrorTextOk() (*string, bool)`

GetErrorTextOk returns a tuple with the ErrorText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorText

`func (o *ContentV2CardsMoveNmPost400Response) SetErrorText(v string)`

SetErrorText sets ErrorText field to given value.

### HasErrorText

`func (o *ContentV2CardsMoveNmPost400Response) HasErrorText() bool`

HasErrorText returns a boolean if a field has been set.

### GetAdditionalErrors

`func (o *ContentV2CardsMoveNmPost400Response) GetAdditionalErrors() ResponseCardCreateAdditionalErrors`

GetAdditionalErrors returns the AdditionalErrors field if non-nil, zero value otherwise.

### GetAdditionalErrorsOk

`func (o *ContentV2CardsMoveNmPost400Response) GetAdditionalErrorsOk() (*ResponseCardCreateAdditionalErrors, bool)`

GetAdditionalErrorsOk returns a tuple with the AdditionalErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalErrors

`func (o *ContentV2CardsMoveNmPost400Response) SetAdditionalErrors(v ResponseCardCreateAdditionalErrors)`

SetAdditionalErrors sets AdditionalErrors field to given value.

### HasAdditionalErrors

`func (o *ContentV2CardsMoveNmPost400Response) HasAdditionalErrors() bool`

HasAdditionalErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


