# ContentV2TagPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | Pointer to **string** | Цвет тега. &lt;dl&gt; &lt;dt&gt;Доступные цвета:&lt;/dt&gt; &lt;dd&gt;&lt;code&gt;D1CFD7&lt;/code&gt; - серый&lt;/dd&gt; &lt;dd&gt;&lt;code&gt;FEE0E0&lt;/code&gt; - красный&lt;/dd&gt; &lt;dd&gt;&lt;code&gt;ECDAFF&lt;/code&gt; - фиолетовый&lt;/dd&gt; &lt;dd&gt;&lt;code&gt;E4EAFF&lt;/code&gt; - синий&lt;/dd&gt; &lt;dd&gt;&lt;code&gt;DEF1DD&lt;/code&gt; - зеленый&lt;/dd&gt; &lt;dd&gt;&lt;code&gt;FFECC7&lt;/code&gt; - желтый&lt;/dd&gt; &lt;/dl&gt;  | [optional] 
**Name** | Pointer to **string** | Имя тега | [optional] 

## Methods

### NewContentV2TagPostRequest

`func NewContentV2TagPostRequest() *ContentV2TagPostRequest`

NewContentV2TagPostRequest instantiates a new ContentV2TagPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentV2TagPostRequestWithDefaults

`func NewContentV2TagPostRequestWithDefaults() *ContentV2TagPostRequest`

NewContentV2TagPostRequestWithDefaults instantiates a new ContentV2TagPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *ContentV2TagPostRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ContentV2TagPostRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ContentV2TagPostRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *ContentV2TagPostRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetName

`func (o *ContentV2TagPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentV2TagPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentV2TagPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContentV2TagPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


