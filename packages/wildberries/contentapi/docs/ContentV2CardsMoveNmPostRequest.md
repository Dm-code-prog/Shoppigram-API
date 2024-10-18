# ContentV2CardsMoveNmPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetIMT** | **int32** | Существующий у продавца &#x60;imtID&#x60;, под которым необходимо объединить НМ | 
**NmIDs** | **[]int32** | &#x60;nmID&#x60;, которые необходимо разъединить (max 30) | 

## Methods

### NewContentV2CardsMoveNmPostRequest

`func NewContentV2CardsMoveNmPostRequest(targetIMT int32, nmIDs []int32, ) *ContentV2CardsMoveNmPostRequest`

NewContentV2CardsMoveNmPostRequest instantiates a new ContentV2CardsMoveNmPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentV2CardsMoveNmPostRequestWithDefaults

`func NewContentV2CardsMoveNmPostRequestWithDefaults() *ContentV2CardsMoveNmPostRequest`

NewContentV2CardsMoveNmPostRequestWithDefaults instantiates a new ContentV2CardsMoveNmPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetIMT

`func (o *ContentV2CardsMoveNmPostRequest) GetTargetIMT() int32`

GetTargetIMT returns the TargetIMT field if non-nil, zero value otherwise.

### GetTargetIMTOk

`func (o *ContentV2CardsMoveNmPostRequest) GetTargetIMTOk() (*int32, bool)`

GetTargetIMTOk returns a tuple with the TargetIMT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetIMT

`func (o *ContentV2CardsMoveNmPostRequest) SetTargetIMT(v int32)`

SetTargetIMT sets TargetIMT field to given value.


### GetNmIDs

`func (o *ContentV2CardsMoveNmPostRequest) GetNmIDs() []int32`

GetNmIDs returns the NmIDs field if non-nil, zero value otherwise.

### GetNmIDsOk

`func (o *ContentV2CardsMoveNmPostRequest) GetNmIDsOk() (*[]int32, bool)`

GetNmIDsOk returns a tuple with the NmIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmIDs

`func (o *ContentV2CardsMoveNmPostRequest) SetNmIDs(v []int32)`

SetNmIDs sets NmIDs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


