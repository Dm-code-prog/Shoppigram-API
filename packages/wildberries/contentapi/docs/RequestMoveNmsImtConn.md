# RequestMoveNmsImtConn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetIMT** | **int32** | Существующий у продавца &#x60;imtID&#x60;, под которым необходимо объединить НМ | 
**NmIDs** | **[]int32** | &#x60;nmID&#x60;, которые необходимо объединить (максимум 30)  | 

## Methods

### NewRequestMoveNmsImtConn

`func NewRequestMoveNmsImtConn(targetIMT int32, nmIDs []int32, ) *RequestMoveNmsImtConn`

NewRequestMoveNmsImtConn instantiates a new RequestMoveNmsImtConn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestMoveNmsImtConnWithDefaults

`func NewRequestMoveNmsImtConnWithDefaults() *RequestMoveNmsImtConn`

NewRequestMoveNmsImtConnWithDefaults instantiates a new RequestMoveNmsImtConn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetIMT

`func (o *RequestMoveNmsImtConn) GetTargetIMT() int32`

GetTargetIMT returns the TargetIMT field if non-nil, zero value otherwise.

### GetTargetIMTOk

`func (o *RequestMoveNmsImtConn) GetTargetIMTOk() (*int32, bool)`

GetTargetIMTOk returns a tuple with the TargetIMT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetIMT

`func (o *RequestMoveNmsImtConn) SetTargetIMT(v int32)`

SetTargetIMT sets TargetIMT field to given value.


### GetNmIDs

`func (o *RequestMoveNmsImtConn) GetNmIDs() []int32`

GetNmIDs returns the NmIDs field if non-nil, zero value otherwise.

### GetNmIDsOk

`func (o *RequestMoveNmsImtConn) GetNmIDsOk() (*[]int32, bool)`

GetNmIDsOk returns a tuple with the NmIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmIDs

`func (o *RequestMoveNmsImtConn) SetNmIDs(v []int32)`

SetNmIDs sets NmIDs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


