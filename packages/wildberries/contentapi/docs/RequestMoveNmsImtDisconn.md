# RequestMoveNmsImtDisconn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NmIDs** | **[]int32** | &#x60;nmID&#x60;, которые необходимо разъединить (max 30) | 

## Methods

### NewRequestMoveNmsImtDisconn

`func NewRequestMoveNmsImtDisconn(nmIDs []int32, ) *RequestMoveNmsImtDisconn`

NewRequestMoveNmsImtDisconn instantiates a new RequestMoveNmsImtDisconn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestMoveNmsImtDisconnWithDefaults

`func NewRequestMoveNmsImtDisconnWithDefaults() *RequestMoveNmsImtDisconn`

NewRequestMoveNmsImtDisconnWithDefaults instantiates a new RequestMoveNmsImtDisconn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNmIDs

`func (o *RequestMoveNmsImtDisconn) GetNmIDs() []int32`

GetNmIDs returns the NmIDs field if non-nil, zero value otherwise.

### GetNmIDsOk

`func (o *RequestMoveNmsImtDisconn) GetNmIDsOk() (*[]int32, bool)`

GetNmIDsOk returns a tuple with the NmIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmIDs

`func (o *RequestMoveNmsImtDisconn) SetNmIDs(v []int32)`

SetNmIDs sets NmIDs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


