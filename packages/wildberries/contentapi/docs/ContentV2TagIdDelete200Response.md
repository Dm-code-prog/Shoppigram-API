# ContentV2TagIdDelete200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Error** | Pointer to **bool** | Флаг ошибки | [optional] 
**ErrorText** | Pointer to **string** | Описание ошибки | [optional] 
**AdditionalErrors** | Pointer to [**ResponseContentError5AdditionalErrors**](ResponseContentError5AdditionalErrors.md) |  | [optional] 

## Methods

### NewContentV2TagIdDelete200Response

`func NewContentV2TagIdDelete200Response() *ContentV2TagIdDelete200Response`

NewContentV2TagIdDelete200Response instantiates a new ContentV2TagIdDelete200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentV2TagIdDelete200ResponseWithDefaults

`func NewContentV2TagIdDelete200ResponseWithDefaults() *ContentV2TagIdDelete200Response`

NewContentV2TagIdDelete200ResponseWithDefaults instantiates a new ContentV2TagIdDelete200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ContentV2TagIdDelete200Response) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ContentV2TagIdDelete200Response) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ContentV2TagIdDelete200Response) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ContentV2TagIdDelete200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ContentV2TagIdDelete200Response) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ContentV2TagIdDelete200Response) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetError

`func (o *ContentV2TagIdDelete200Response) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ContentV2TagIdDelete200Response) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ContentV2TagIdDelete200Response) SetError(v bool)`

SetError sets Error field to given value.

### HasError

`func (o *ContentV2TagIdDelete200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorText

`func (o *ContentV2TagIdDelete200Response) GetErrorText() string`

GetErrorText returns the ErrorText field if non-nil, zero value otherwise.

### GetErrorTextOk

`func (o *ContentV2TagIdDelete200Response) GetErrorTextOk() (*string, bool)`

GetErrorTextOk returns a tuple with the ErrorText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorText

`func (o *ContentV2TagIdDelete200Response) SetErrorText(v string)`

SetErrorText sets ErrorText field to given value.

### HasErrorText

`func (o *ContentV2TagIdDelete200Response) HasErrorText() bool`

HasErrorText returns a boolean if a field has been set.

### GetAdditionalErrors

`func (o *ContentV2TagIdDelete200Response) GetAdditionalErrors() ResponseContentError5AdditionalErrors`

GetAdditionalErrors returns the AdditionalErrors field if non-nil, zero value otherwise.

### GetAdditionalErrorsOk

`func (o *ContentV2TagIdDelete200Response) GetAdditionalErrorsOk() (*ResponseContentError5AdditionalErrors, bool)`

GetAdditionalErrorsOk returns a tuple with the AdditionalErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalErrors

`func (o *ContentV2TagIdDelete200Response) SetAdditionalErrors(v ResponseContentError5AdditionalErrors)`

SetAdditionalErrors sets AdditionalErrors field to given value.

### HasAdditionalErrors

`func (o *ContentV2TagIdDelete200Response) HasAdditionalErrors() bool`

HasAdditionalErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


