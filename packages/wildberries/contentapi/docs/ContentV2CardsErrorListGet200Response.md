# ContentV2CardsErrorListGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ContentV2CardsErrorListGet200ResponseDataInner**](ContentV2CardsErrorListGet200ResponseDataInner.md) |  | [optional] 
**Error** | Pointer to **bool** | Флаг ошибки. | [optional] 
**ErrorText** | Pointer to **string** | Описание ошибки. | [optional] 
**AdditionalErrors** | Pointer to **string** | Дополнительные ошибки. | [optional] 

## Methods

### NewContentV2CardsErrorListGet200Response

`func NewContentV2CardsErrorListGet200Response() *ContentV2CardsErrorListGet200Response`

NewContentV2CardsErrorListGet200Response instantiates a new ContentV2CardsErrorListGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentV2CardsErrorListGet200ResponseWithDefaults

`func NewContentV2CardsErrorListGet200ResponseWithDefaults() *ContentV2CardsErrorListGet200Response`

NewContentV2CardsErrorListGet200ResponseWithDefaults instantiates a new ContentV2CardsErrorListGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ContentV2CardsErrorListGet200Response) GetData() []ContentV2CardsErrorListGet200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ContentV2CardsErrorListGet200Response) GetDataOk() (*[]ContentV2CardsErrorListGet200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ContentV2CardsErrorListGet200Response) SetData(v []ContentV2CardsErrorListGet200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *ContentV2CardsErrorListGet200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *ContentV2CardsErrorListGet200Response) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ContentV2CardsErrorListGet200Response) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ContentV2CardsErrorListGet200Response) SetError(v bool)`

SetError sets Error field to given value.

### HasError

`func (o *ContentV2CardsErrorListGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorText

`func (o *ContentV2CardsErrorListGet200Response) GetErrorText() string`

GetErrorText returns the ErrorText field if non-nil, zero value otherwise.

### GetErrorTextOk

`func (o *ContentV2CardsErrorListGet200Response) GetErrorTextOk() (*string, bool)`

GetErrorTextOk returns a tuple with the ErrorText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorText

`func (o *ContentV2CardsErrorListGet200Response) SetErrorText(v string)`

SetErrorText sets ErrorText field to given value.

### HasErrorText

`func (o *ContentV2CardsErrorListGet200Response) HasErrorText() bool`

HasErrorText returns a boolean if a field has been set.

### GetAdditionalErrors

`func (o *ContentV2CardsErrorListGet200Response) GetAdditionalErrors() string`

GetAdditionalErrors returns the AdditionalErrors field if non-nil, zero value otherwise.

### GetAdditionalErrorsOk

`func (o *ContentV2CardsErrorListGet200Response) GetAdditionalErrorsOk() (*string, bool)`

GetAdditionalErrorsOk returns a tuple with the AdditionalErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalErrors

`func (o *ContentV2CardsErrorListGet200Response) SetAdditionalErrors(v string)`

SetAdditionalErrors sets AdditionalErrors field to given value.

### HasAdditionalErrors

`func (o *ContentV2CardsErrorListGet200Response) HasAdditionalErrors() bool`

HasAdditionalErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


