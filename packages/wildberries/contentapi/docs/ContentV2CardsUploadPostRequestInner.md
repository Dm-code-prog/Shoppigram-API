# ContentV2CardsUploadPostRequestInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectID** | **int32** | ID предмета | 
**Variants** | [**[]ContentV2CardsUploadPostRequestInnerVariantsInner**](ContentV2CardsUploadPostRequestInnerVariantsInner.md) | Массив вариантов товара. В каждой КТ может быть не более 30 вариантов (НМ) | 

## Methods

### NewContentV2CardsUploadPostRequestInner

`func NewContentV2CardsUploadPostRequestInner(subjectID int32, variants []ContentV2CardsUploadPostRequestInnerVariantsInner, ) *ContentV2CardsUploadPostRequestInner`

NewContentV2CardsUploadPostRequestInner instantiates a new ContentV2CardsUploadPostRequestInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentV2CardsUploadPostRequestInnerWithDefaults

`func NewContentV2CardsUploadPostRequestInnerWithDefaults() *ContentV2CardsUploadPostRequestInner`

NewContentV2CardsUploadPostRequestInnerWithDefaults instantiates a new ContentV2CardsUploadPostRequestInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectID

`func (o *ContentV2CardsUploadPostRequestInner) GetSubjectID() int32`

GetSubjectID returns the SubjectID field if non-nil, zero value otherwise.

### GetSubjectIDOk

`func (o *ContentV2CardsUploadPostRequestInner) GetSubjectIDOk() (*int32, bool)`

GetSubjectIDOk returns a tuple with the SubjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectID

`func (o *ContentV2CardsUploadPostRequestInner) SetSubjectID(v int32)`

SetSubjectID sets SubjectID field to given value.


### GetVariants

`func (o *ContentV2CardsUploadPostRequestInner) GetVariants() []ContentV2CardsUploadPostRequestInnerVariantsInner`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *ContentV2CardsUploadPostRequestInner) GetVariantsOk() (*[]ContentV2CardsUploadPostRequestInnerVariantsInner, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *ContentV2CardsUploadPostRequestInner) SetVariants(v []ContentV2CardsUploadPostRequestInnerVariantsInner)`

SetVariants sets Variants field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


