# ContentV2CardsLimitsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ContentV2CardsLimitsGet200ResponseData**](ContentV2CardsLimitsGet200ResponseData.md) |  | [optional] 
**Error** | Pointer to **bool** | Флаг ошибки | [optional] 
**ErrorText** | Pointer to **string** | Описание ошибки | [optional] 
**AdditionalErrors** | Pointer to **NullableString** | Дополнительные ошибки | [optional] 

## Methods

### NewContentV2CardsLimitsGet200Response

`func NewContentV2CardsLimitsGet200Response() *ContentV2CardsLimitsGet200Response`

NewContentV2CardsLimitsGet200Response instantiates a new ContentV2CardsLimitsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentV2CardsLimitsGet200ResponseWithDefaults

`func NewContentV2CardsLimitsGet200ResponseWithDefaults() *ContentV2CardsLimitsGet200Response`

NewContentV2CardsLimitsGet200ResponseWithDefaults instantiates a new ContentV2CardsLimitsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ContentV2CardsLimitsGet200Response) GetData() ContentV2CardsLimitsGet200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ContentV2CardsLimitsGet200Response) GetDataOk() (*ContentV2CardsLimitsGet200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ContentV2CardsLimitsGet200Response) SetData(v ContentV2CardsLimitsGet200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *ContentV2CardsLimitsGet200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *ContentV2CardsLimitsGet200Response) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ContentV2CardsLimitsGet200Response) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ContentV2CardsLimitsGet200Response) SetError(v bool)`

SetError sets Error field to given value.

### HasError

`func (o *ContentV2CardsLimitsGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorText

`func (o *ContentV2CardsLimitsGet200Response) GetErrorText() string`

GetErrorText returns the ErrorText field if non-nil, zero value otherwise.

### GetErrorTextOk

`func (o *ContentV2CardsLimitsGet200Response) GetErrorTextOk() (*string, bool)`

GetErrorTextOk returns a tuple with the ErrorText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorText

`func (o *ContentV2CardsLimitsGet200Response) SetErrorText(v string)`

SetErrorText sets ErrorText field to given value.

### HasErrorText

`func (o *ContentV2CardsLimitsGet200Response) HasErrorText() bool`

HasErrorText returns a boolean if a field has been set.

### GetAdditionalErrors

`func (o *ContentV2CardsLimitsGet200Response) GetAdditionalErrors() string`

GetAdditionalErrors returns the AdditionalErrors field if non-nil, zero value otherwise.

### GetAdditionalErrorsOk

`func (o *ContentV2CardsLimitsGet200Response) GetAdditionalErrorsOk() (*string, bool)`

GetAdditionalErrorsOk returns a tuple with the AdditionalErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalErrors

`func (o *ContentV2CardsLimitsGet200Response) SetAdditionalErrors(v string)`

SetAdditionalErrors sets AdditionalErrors field to given value.

### HasAdditionalErrors

`func (o *ContentV2CardsLimitsGet200Response) HasAdditionalErrors() bool`

HasAdditionalErrors returns a boolean if a field has been set.

### SetAdditionalErrorsNil

`func (o *ContentV2CardsLimitsGet200Response) SetAdditionalErrorsNil(b bool)`

 SetAdditionalErrorsNil sets the value for AdditionalErrors to be an explicit nil

### UnsetAdditionalErrors
`func (o *ContentV2CardsLimitsGet200Response) UnsetAdditionalErrors()`

UnsetAdditionalErrors ensures that no value is present for AdditionalErrors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


